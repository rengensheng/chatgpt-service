package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/models"
	"github.com/goylold/lowcode/validator"
)

const RoleTableName = "Role"

func RoleList(c *gin.Context) {
	engine := database.GetXOrmEngine()
	var requestParams common.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	var tableEntities []models.Role
	err = requestParams.DisposeRequest(engine.Table(RoleTableName)).Find(&tableEntities)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	count, err := requestParams.DisposeRequest(engine.Table(RoleTableName)).Count()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccessList(tableEntities, count, c)
}

func RoleAdd(c *gin.Context) {
	var table models.Role
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

func RoleUpdate(c *gin.Context) {
	var table models.Role
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

func RoleDelete(c *gin.Context) {
	var table models.Role
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

func RoleGetOne(c *gin.Context) {
	var table models.Role
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
