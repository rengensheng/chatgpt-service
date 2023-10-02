package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/models"
	"github.com/goylold/lowcode/utils"
	"github.com/goylold/lowcode/validator"
)

const UserTableName = "User"

func UserList(c *gin.Context) {
	engine := database.GetXOrmEngine()
	var requestParams common.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	var tableEntities []models.User
	err = requestParams.DisposeRequest(engine.Table(UserTableName)).Find(&tableEntities)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	count, err := requestParams.DisposeRequest(engine.Table(UserTableName)).Count()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccessList(tableEntities, count, c)
}

func UserAdd(c *gin.Context) {
	var table models.User
	err := c.ShouldBindJSON(&table)
	if err != nil {
		common.ResultError(500, validator.Translate(err), c)
		return
	}
	if table.Pwd != "" {
		table.Pwd = utils.MD5(table.Pwd)
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

func UserUpdate(c *gin.Context) {
	var table models.User
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

func UserDelete(c *gin.Context) {
	var table models.User
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

func UserGetOne(c *gin.Context) {
	var table models.User
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

func AccountExist(c *gin.Context) {
	engine := database.GetXOrmEngine()
	var requestParams common.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	var tableEntities []models.User
	err = requestParams.DisposeRequest(engine.Table(UserTableName)).Find(&tableEntities)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	if len(tableEntities) > 0 {
		common.ResultSuccess(true, c)
		return
	}
	common.ResultSuccess(false, c)
}
