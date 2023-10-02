package services

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/models"
	"github.com/goylold/lowcode/utils"
)

func UserLogin(c *gin.Context) {
	json := make(map[string]string)
	err := c.ShouldBindJSON(&json)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	if json["username"] == "" || json["password"] == "" {
		common.ResultError(500, "用户名或密码是空！", c)
		return
	}
	user := models.User{}
	password := utils.MD5(json["password"])
	isExits := user.UserFindByUsernameAndPassword(json["username"], password)
	if !isExits {
		common.ResultError(500, "登录失败，用户名或密码错误", c)
		return
	}
	token, _ := common.UserLogin(user, c)
	user.Pwd = ""

	userLoginRes := models.UserLoginEntity{
		Desc:     user.Remark,
		RealName: user.Nickname,
		Token:    token,
		UserId:   user.Id,
		Username: user.Username,
	}
	common.ResultSuccess(userLoginRes, c)
	// common.ResultSuccess(`{"code":0,"result":{"userId":"1","username":"vben","realName":"自动化测试平台","avatar":"https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640","desc":"manager","password":"123456","token":"fakeToken1","homePath":"/dashboard/analysis","roles":[{"roleName":"Super Admin","value":"super"}]},"message":"ok","type":"success"}`, c)
}

func GetUserInfo(c *gin.Context) {
	user := common.GetCurrentUserInfo(c)
	user.Pwd = ""
	role := models.Role{}
	roleList, _ := role.GetRoleByRoleValues(strings.Split(user.Role, ","))
	simpleRoles := []models.SimpleRole{}
	for _, role := range roleList {
		simpleRoles = append(simpleRoles, models.SimpleRole{
			RoleName: role.RoleName,
			Value:    role.RoleValue,
		})
	}
	token := ""
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		token = c.GetHeader("Authorization")
	} else {
		token = cookie.Value
	}
	userLoginRes := models.UserLoginEntity{
		Desc:     user.Remark,
		RealName: user.Nickname,
		Roles:    simpleRoles,
		UserId:   user.Id,
		Username: user.Username,
		Token:    token,
		Avatar:   "https://q1.qlogo.cn/g?b=qq&nk=190848757&s=640",
	}
	common.ResultSuccess(userLoginRes, c)
}

func GetPermCode(c *gin.Context) {
	codes := common.GetPermCode(c)
	common.ResultSuccess(codes, c)
}

func GetMenuList(c *gin.Context) {
	var MenuItems []models.MenuItem
	user := common.GetCurrentUserInfo(c)
	role := models.Role{}
	roles, _ := role.GetRoleByRoleValues(strings.Split(user.Role, ","))
	menusIds := []string{}
	for _, role := range roles {
		menus := strings.Split(role.Menu, ",")
		menusIds = append(menusIds, menus...)
	}
	menuModel := models.Menu{}
	menuList, err := menuModel.GetMenuListByIds(menusIds)
	if err != nil {
		log.Println(err.Error())
		common.ResultError(500, "获取菜单列表失败", c)
		return
	}
	for _, menu := range menuList {
		if menu.Status != "0" {
			continue
		}
		MenuItems = append(MenuItems, models.MenuItem{
			Id:         menu.Id,
			Path:       menu.RoutePath,
			Name:       menu.MenuName,
			Icon:       menu.Icon,
			ParentMenu: menu.ParentMenu,
			Component:  menu.Component,
			Children:   []models.MenuItem{},
			OrderNo:    menu.OrderNo,
			Meta: models.MenuMeta{
				Icon:            menu.Icon,
				Title:           menu.MenuName,
				HideMenu:        menu.Show == "1",
				ShowMenu:        menu.Show == "0",
				IgnoreKeepAlive: menu.Keepalive == "0",
			},
		})
	}
	for idx, i := range MenuItems {
		for _, j := range MenuItems {
			if i.Id == j.ParentMenu {
				MenuItems[idx].Children = append(MenuItems[idx].Children, j)
			}
		}
	}
	var menuItemResult []models.MenuItem
	for _, i := range MenuItems {
		if i.ParentMenu == "" {
			menuItemResult = append(menuItemResult, i)
		}
	}
	common.ResultSuccess(menuItemResult, c)
}
