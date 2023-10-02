package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
)

const MessageTableName string = "Message"

type Message struct {
	Content string `form:"content" json:"content" xorm:"text notnull" binding:"required,max=65535"`

	CreatedBy string `form:"created_by" json:"created_by" xorm:"varchar(32)" binding:"max=32"`

	CreatedTime string `form:"created_time" json:"created_time" xorm:"timestamp notnull" binding:"required"`

	Id string `form:"id" json:"id" xorm:"varchar(32) pk notnull" binding:"max=32"`

	Role string `form:"role" json:"role" xorm:"varchar(32)" binding:"max=32"`

	SenderID string `form:"sender_id" json:"sender_id" xorm:"varchar(32)" binding:"max=32"`

	SentAt string `form:"sent_at" json:"sent_at" xorm:"timestamp notnull" binding:"required"`

	SessionID string `form:"session_id" json:"session_id" xorm:"varchar(32)" binding:"max=32"`

	UpdatedBy string `form:"updated_by" json:"updated_by" xorm:"varchar(32)" binding:"max=32"`

	UpdatedTime string `form:"updated_time" json:"updated_time" xorm:"timestamp notnull" binding:"required"`
}

func (table *Message) Add() error {
	engine := database.GetXOrmEngine()
	table.Id = utils.GetUUID()
	table.SentAt = carbon.Now().ToDateTimeString()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(MessageTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Message) Update() error {
	engine := database.GetXOrmEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(MessageTableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Message) Delete() error {
	engine := database.GetXOrmEngine()
	affected, err := engine.Table(MessageTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *Message) GetOne(id string) error {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(MessageTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}
