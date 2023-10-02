package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
)

const BalanceTableName string = "Balance"

type Balance struct {
	Amount string `form:"amount" json:"amount" xorm:"decimal(10,2) notnull" binding:"required"`

	CreatedBy string `form:"created_by" json:"created_by" xorm:"varchar(32)" binding:"max=32"`

	CreatedTime string `form:"created_time" json:"created_time" xorm:"timestamp notnull" binding:"required"`

	Id string `form:"id" json:"id" xorm:"int pk notnull"`

	UpdatedAt string `form:"updated_at" json:"updated_at" xorm:"timestamp notnull" binding:"required"`

	UpdatedBy string `form:"updated_by" json:"updated_by" xorm:"varchar(32)" binding:"max=32"`

	UpdatedTime string `form:"updated_time" json:"updated_time" xorm:"timestamp notnull" binding:"required"`

	UserID string `form:"user_id" json:"user_id" xorm:"varchar(32)" binding:"max=32"`
}

func (table *Balance) Add() error {
	engine := database.GetXOrmEngine()
	table.Id = utils.GetUUID()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(BalanceTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Balance) Update() error {
	engine := database.GetXOrmEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(BalanceTableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Balance) Delete() error {
	engine := database.GetXOrmEngine()
	affected, err := engine.Table(BalanceTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *Balance) GetOne(id string) error {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(BalanceTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}
