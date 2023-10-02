package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
)

const ConversationTableName string = "Conversation"

type Conversation struct {
	Id string `form:"id" json:"id" xorm:"varchar(32) pk notnull" binding:"max=32"`

	SessionId string `form:"session_id" json:"session_id" xorm:"varchar(32)" binding:"max=32"`

	TextContent string `form:"text_content" json:"text_content" xorm:"text" binding:"max=65535"`

	TextEncoding string `form:"text_encoding" json:"text_encoding" xorm:"text" binding:"max=65535"`

	UserId string `form:"user_id" json:"user_id" xorm:"varchar(32)" binding:"max=32"`

	CreatedBy string `form:"created_by" json:"created_by" xorm:"varchar(32)" binding:"max=32"`

	CreatedTime string `form:"created_time" json:"created_time" xorm:"timestamp notnull" binding:"required"`

	UpdatedBy string `form:"updated_by" json:"updated_by" xorm:"varchar(32)" binding:"max=32"`

	UpdatedTime string `form:"updated_time" json:"updated_time" xorm:"timestamp notnull" binding:"required"`
}

func (table *Conversation) Add() error {
	engine := database.GetXOrmEngine()
	table.Id = utils.GetUUID()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(ConversationTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Conversation) Update() error {
	engine := database.GetXOrmEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(ConversationTableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Conversation) Delete() error {
	engine := database.GetXOrmEngine()
	affected, err := engine.Table(ConversationTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *Conversation) GetOne(id string) error {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(ConversationTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}
