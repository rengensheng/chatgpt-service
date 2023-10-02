package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
)

const TokenConsumptionTableName string = "TokenConsumption"

type TokenConsumption struct {
	Amount int64 `form:"amount" json:"amount" xorm:"int notnull" binding:"required"`

	ConsumedAt string `form:"consumed_at" json:"consumed_at" xorm:"timestamp notnull" binding:"required"`

	CreatedBy string `form:"created_by" json:"created_by" xorm:"varchar(32)" binding:"max=32"`

	CreatedTime string `form:"created_time" json:"created_time" xorm:"timestamp notnull" binding:"required"`

	Id string `form:"id" json:"id" xorm:"varchar(32) pk notnull" binding:"max=32"`

	UpdatedBy string `form:"updated_by" json:"updated_by" xorm:"varchar(32)" binding:"max=32"`

	UpdatedTime string `form:"updated_time" json:"updated_time" xorm:"timestamp notnull" binding:"required"`

	Message string `form:"message" json:"message" xorm:"text"`

	UserID string `form:"user_id" json:"user_id" xorm:"varchar(32)" binding:"max=32"`
}

func (table *TokenConsumption) Add() error {
	engine := database.GetXOrmEngine()
	table.Id = utils.GetUUID()
	table.ConsumedAt = carbon.Now().ToDateTimeString()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(TokenConsumptionTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *TokenConsumption) Update() error {
	engine := database.GetXOrmEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(TokenConsumptionTableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *TokenConsumption) Delete() error {
	engine := database.GetXOrmEngine()
	affected, err := engine.Table(TokenConsumptionTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *TokenConsumption) GetOne(id string) error {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(TokenConsumptionTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}
