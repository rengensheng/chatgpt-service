package services

import (
	"github.com/gin-gonic/gin"
	"github.com/goylold/lowcode/common"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/models"
	"github.com/goylold/lowcode/validator"
)

const RechargeTableName = "Recharge"

func RechargeList(c *gin.Context) {
	engine := database.GetXOrmEngine()
	var requestParams common.Request
	err := c.ShouldBindJSON(&requestParams)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	var tableEntities []models.Recharge
	err = requestParams.DisposeRequest(engine.Table(RechargeTableName)).Find(&tableEntities)
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	count, err := requestParams.DisposeRequest(engine.Table(RechargeTableName)).Count()
	if err != nil {
		common.ResultError(500, err.Error(), c)
		return
	}
	common.ResultSuccessList(tableEntities, count, c)
}

func RechargeAdd(c *gin.Context) {
	var table models.Recharge
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

func RechargeUpdate(c *gin.Context) {
	var table models.Recharge
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

func RechargeDelete(c *gin.Context) {
	var table models.Recharge
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

func RechargeGetOne(c *gin.Context) {
	var table models.Recharge
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
