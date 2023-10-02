package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/models"
	"github.com/goylold/lowcode/validator"
)

const SessionTableName = "Session"

func SessionList(c *gin.Context) {
	engine := database.GetXOrmEngine()
	var requestParams common.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	if requestParams.Query == nil {
		requestParams.Query = make(map[string]string)
	}
	currentUserId := common.GetCurrentUser(c)
	requestParams.Query["user_i_d"] = currentUserId
	var tableEntities []models.Session
	err = requestParams.DisposeRequest(engine.Table(SessionTableName)).Find(&tableEntities)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	count, err := requestParams.DisposeRequest(engine.Table(SessionTableName)).Count()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccessList(tableEntities, count, c)
}

func SessionAdd(c *gin.Context) {
	var table models.Session
	err := c.ShouldBindJSON(&table)
	if err != nil {
		common.ResultError(500, validator.Translate(err), c)
		return
	}
	currentUserId := common.GetCurrentUser(c)
	table.CreatedBy = currentUserId
	table.UpdatedBy = currentUserId
	table.UserID = currentUserId
	err = table.Add()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccess(table, c)
}

func SessionUpdate(c *gin.Context) {
	var table models.Session
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

func SessionDelete(c *gin.Context) {
	var table models.Session
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

func SessionGetOne(c *gin.Context) {
	var table models.Session
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
