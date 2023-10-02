package models

import (
	"errors"
	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
)

const WebUserTableName string = "WebUser"

type WebUser struct {
	CreatedAt string `form:"created_at" json:"created_t" xorm:"timestamp notnull"`

	CreatedBy string `form:"created_by" json:"created_by" xorm:"varchar(32)" binding:"max=32"`

	CreatedTime string `form:"created_time" json:"created_time" xorm:"timestamp notnull"`

	Email string `form:"email" json:"email" xorm:"varchar(255) notnull" binding:"max=255"`

	Id string `form:"id" json:"id" xorm:"varchar(32) pk notnull" binding:"max=32"`

	Password string `form:"password" json:"password" xorm:"varchar(255) notnull" binding:"required,max=255"`

	Phone string `form:"phone" json:"phone" xorm:"varchar(255) notnull" binding:"max=255"`

	UpdatedBy string `form:"updated_by" json:"updated_by" xorm:"varchar(32)" binding:"max=32"`

	UpdatedTime string `form:"updated_time" json:"updated_time" xorm:"timestamp notnull"`

	Username string `form:"username" json:"username" xorm:"varchar(255) notnull" binding:"required,max=255"`

	WeChatID string `form:"wechat_id" json:"wechat_id" xorm:"varchar(255) notnull" binding:"max=255"`
}

func (table *WebUser) Add() error {
	engine := database.GetXOrmEngine()
	table.Id = utils.GetUUID()
	table.CreatedAt = carbon.Now().ToDateTimeString()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	table.Password = utils.Sha256(table.Password)
	_, err := engine.Table(WebUserTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *WebUser) Update() error {
	engine := database.GetXOrmEngine()
	table.Password = utils.Sha256(table.Password)
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(WebUserTableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *WebUser) Delete() error {
	engine := database.GetXOrmEngine()
	affected, err := engine.Table(WebUserTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *WebUser) GetOne(id string) error {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(WebUserTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *WebUser) GetUserByUsername() error {
	engine := database.GetXOrmEngine()
	newWebUser := WebUser{Username: table.Username}
	_, err := engine.Table(WebUserTableName).Get(&newWebUser)
	table.Id = newWebUser.Id
	if err != nil {
		return err
	}
	return nil
}

func (table *WebUser) GetUserByUsernameAndPassword() error {
	engine := database.GetXOrmEngine()
	table.Password = utils.Sha256(table.Password)
	_, err := engine.Table(WebUserTableName).Where("password=?", table.Password).And("username=?", table.Username).Get(table)
	if err != nil {
		return err
	}
	return nil
}
