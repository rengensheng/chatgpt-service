package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
)

const SessionTableName string = "Session"

type Session struct {
	CreatedAt string `form:"created_at" json:"created_at" xorm:"timestamp notnull"`

	CreatedBy string `form:"created_by" json:"created_by" xorm:"varchar(32)" binding:"max=32"`

	CreatedTime string `form:"created_time" json:"created_time" xorm:"timestamp notnull"`

	Id string `form:"id" json:"id" xorm:"varchar(32) pk notnull" binding:"max=32"`

	SessionName string `form:"session_name" json:"session_name" xorm:"varchar(255) notnull" binding:"required,max=255"`

	History int64 `form:"history" json:"history" xorm:"int"`

	UpdatedBy string `form:"updated_by" json:"updated_by" xorm:"varchar(32)" binding:"max=32"`

	UpdatedTime string `form:"updated_time" json:"updated_time" xorm:"timestamp notnull"`

	UserID string `form:"user_id" json:"user_id" xorm:"varchar(32)" binding:"max=32"`

	SessionType string `form:"session_type" json:"session_type" xorm:"varchar(32)" binding:"max=32"`

	Params string `form:"params" json:"params" xorm:"text"`
}

func (table *Session) Add() error {
	engine := database.GetXOrmEngine()
	table.Id = utils.GetUUID()
	table.CreatedAt = carbon.Now().ToDateTimeString()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(SessionTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Session) Update() error {
	engine := database.GetXOrmEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(SessionTableName).ID(table.Id).Cols("session_name", "history", "updated_time", "updated_by").Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Session) Delete() error {
	engine := database.GetXOrmEngine()
	affected, err := engine.Table(SessionTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *Session) GetOne(id string) error {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(SessionTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}
