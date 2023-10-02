package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/models"
	"github.com/goylold/lowcode/validator"
)

const WebUserTableName = "WebUser"

func WebUserList(c *gin.Context) {
	engine := database.GetXOrmEngine()
	var requestParams common.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	var tableEntities []models.WebUser
	err = requestParams.DisposeRequest(engine.Table(WebUserTableName)).Find(&tableEntities)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	count, err := requestParams.DisposeRequest(engine.Table(WebUserTableName)).Count()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccessList(tableEntities, count, c)
}

func WebUserAdd(c *gin.Context) {
	var table models.WebUser
	err := c.ShouldBindJSON(&table)
	if err != nil {
		common.ResultError(500, validator.Translate(err), c)
		return
	}
	table.CreatedBy = common.GetCurrentUser(c)
	table.UpdatedBy = common.GetCurrentUser(c)
	err = table.Add()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(table, c)
}

func WebUserUpdate(c *gin.Context) {
	var table models.WebUser
	err := c.ShouldBindJSON(&table)
	if err != nil {
		common.ResultError(500, validator.Translate(err), c)
		return
	}
	if table.Id == "" {
		common.ResultError(500, "Id不能为空", c)
		return
	}
	table.UpdatedBy = common.GetCurrentUser(c)
	err = table.Update()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(table, c)
}

func WebUserDelete(c *gin.Context) {
	var table models.WebUser
	id := c.Param("id")
	if id == "" {
		common.ResultError(500, "Id不能为空", c)
		return
	}
	table.Id = id
	err := table.Delete()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(table, c)
}

func WebUserGetOne(c *gin.Context) {
	var table models.WebUser
	id := c.Param("id")
	if id == "" {
		common.ResultError(500, "Id不能为空", c)
		return
	}
	err := table.GetOne(id)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(table, c)
}

func WebUserLogin(c *gin.Context) {
	var table models.WebUser
	err := c.ShouldBindJSON(&table)
	if err != nil {
		common.ResultError(500, validator.Translate(err), c)
		return
	}
	if table.Username == "" || table.Password == "" {
		common.ResultError(500, "请输入用户名或密码", c)
		return
	}
	err = table.GetUserByUsernameAndPassword()
	if err != nil {
		common.ResultError(500, "登录错误:"+err.Error(), c)
		return
	}
	if table.Id == "" {
		common.ResultError(403, "用户名或密码错误", c)
		return
	}
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	user := models.User{
		Account: table.Username,
		Id:      table.Id,
		Role:    "user",
	}
	token, err := common.UserLogin(user, c)
	if err != nil {
		common.ResultError(500, "注册用户登录失败，请联系管理员", c)
		return
	}
	c.SetCookie("token", token, 3600000, "/", "*", false, true)
	common.ResultSuccess(token, c)
}

func WebUserRegister(c *gin.Context) {
	var table models.WebUser
	err := c.ShouldBindJSON(&table)
	if err != nil {
		common.ResultError(500, validator.Translate(err), c)
		return
	}
	if table.Username == "" || table.Password == "" {
		common.ResultError(500, "请输入用户名或密码", c)
		return
	}
	// 查找是否存在用户
	err = table.GetUserByUsername()
	if err != nil {
		common.ResultError(500, "登录错误", c)
		return
	}
	if table.Id != "" {
		common.ResultError(401, "用户已存在", c)
		return
	}
	// 开始注册
	err = table.Add()
	if err != nil {
		common.ResultError(500, "注册用户失败，请联系管理员"+err.Error(), c)
		return
	}
	// 注册成功，直接登录
	user := models.User{
		Account: table.Username,
		Id:      table.Id,
		Role:    "user",
	}
	token, err := common.UserLogin(user, c)
	if err != nil {
		common.ResultError(500, "注册用户登录失败，请联系管理员", c)
		return
	}
	c.SetCookie("token", token, 3600000, "/", "*", false, true)
	common.ResultSuccess(token, c)
}
