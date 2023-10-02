package common

import (
	"github.com/goylold/lowcode/config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/models"
	"github.com/hyahm/golog"
)

type CustomClaims struct {
	ID         string
	Username   string
	Permission []string
	Roles      []string
	jwt.StandardClaims
}

func GetCurrentUserClaims(c *gin.Context) (*CustomClaims, error) {
	cookie, err := c.Request.Cookie("token")
	token := ""
	if err != nil {
		token = c.GetHeader("Authorization")
	} else {
		token = cookie.Value
	}
	userClaims, err := ParseToken(token)
	return userClaims, err
}

func GetCurrentUser(c *gin.Context) string {
	cookie, err := c.Request.Cookie("token")
	token := ""
	if err != nil {
		token = c.GetHeader("Authorization")
	} else {
		token = cookie.Value
	}
	userClaims, err := ParseToken(token)
	if err != nil {
		return "none"
	}
	return userClaims.Username
}

func GetCurrentUserInfo(c *gin.Context) models.User {
	user := models.User{}
	token := ""
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		golog.Info(err.Error())
		token = c.GetHeader("Authorization")
	} else {
		token = cookie.Value
	}
	userClaims, err := ParseToken(token)

	if err != nil {
		golog.Info(err.Error())
		return user
	}
	golog.Info(userClaims.Username, userClaims.ID)
	user.GetOne(userClaims.ID)
	return user
}

func GetPermCode(c *gin.Context) []string {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		golog.Info(err.Error())
		return []string{}
	}
	userClaims, err := ParseToken(cookie.Value)

	if err != nil {
		golog.Info(err.Error())
		return []string{}
	}
	return userClaims.Permission
}

func UserLogin(user models.User, c *gin.Context) (string, error) {
	token, err := GenerateToken(user)
	if err != nil {
		return "", err
	}
	c.SetCookie("token", token, 3600000, "/", "localhost", false, true)
	return token, nil
}

func GenerateToken(user models.User) (string, error) {
	var permission []string
	var roleValues []string
	role := models.Role{}
	roles, _ := role.GetRoleByRoleValues(strings.Split(user.Role, ","))
	var menusIds []string
	for _, role := range roles {
		roleValues = append(roleValues, role.RoleValue)
		menus := strings.Split(role.Menu, ",")
		menusIds = append(menusIds, menus...)
	}
	menuModel := models.Menu{}
	menuList, err := menuModel.GetMenuListByIds(menusIds)
	if err == nil {
		for _, v := range menuList {
			golog.Info(v, v.Permission)
			if v.Permission != "" {
				permission = append(permission, v.Permission)
			}
		}
	}
	nowTime := time.Now()
	expireTime := nowTime.Add(3600000 * time.Second)
	golog.Info("permission", permission)
	issuer := user.Account
	claims := CustomClaims{
		ID:         user.Id,
		Username:   user.Account,
		Permission: permission,
		Roles:      roleValues,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.GetConConfig().SecretKey))
	return token, err
}

func ParseToken(token string) (*CustomClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConConfig().SecretKey), nil
	})
	if err != nil {
		return &CustomClaims{}, err
	}
	if jwtToken != nil {
		if claims, ok := jwtToken.Claims.(*CustomClaims); ok && jwtToken.Valid {
			return claims, nil
		}
	}
	return &CustomClaims{}, nil
}
