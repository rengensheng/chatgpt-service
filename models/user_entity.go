package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
)

const UserTableName string = "User"

type User struct {
	Account string `form:"account" json:"account" xorm:"varchar(255)" binding:"max=255"`

	Avatar string `form:"avatar" json:"avatar" xorm:"varchar(32)" binding:"max=32"` // 头像

	CreatedBy string `form:"created_by" json:"created_by" xorm:"varchar(255)" binding:"min=0,max=255"`

	CreatedTime string `form:"created_time" json:"created_time" xorm:"datetime"` // 创建时间

	Dept string `form:"dept" json:"dept" xorm:"varchar(255)" binding:"min=0,max=255"`

	Email string `form:"email" json:"email" xorm:"varchar(255)" binding:"min=0,max=255"`

	Id string `form:"id" json:"id" xorm:"varchar(255) pk notnull" binding:"min=0,max=255"`

	LoginId string `form:"loginId" json:"loginId" xorm:"varchar(32)" binding:"min=0,max=32"` // 登录名

	Nickname string `form:"nickname" json:"nickname" xorm:"varchar(255)" binding:"min=0,max=255"`

	Password string `form:"password" json:"password" xorm:"varchar(32)" binding:"min=0,max=32"` // 登陆密码

	Phone string `form:"phone" json:"phone" xorm:"varchar(32)" binding:"min=0,max=32"` // 手机号

	Pwd string `form:"pwd" json:"pwd" xorm:"varchar(255)" binding:"min=0,max=255"`

	Remark string `form:"remark" json:"remark" xorm:"varchar(255)" binding:"min=0,max=255"`

	Revision int64 `form:"revision" json:"revision" xorm:"int"` // 乐观锁

	Role string `form:"role" json:"role" xorm:"varchar(255)" binding:"min=0,max=255"`

	UpdatedBy string `form:"updated_by" json:"updated_by" xorm:"varchar(255)" binding:"max=255"`

	UpdatedTime string `form:"updated_time" json:"updated_time" xorm:"datetime"` // 更新时间

	Username string `form:"username" json:"username" xorm:"varchar(32)" binding:"max=32"` // 用名

}

type UserLoginEntity struct {
	Desc     string       `json:"desc"`
	RealName string       `json:"realName"`
	Token    string       `json:"token"`
	UserId   string       `json:"userId"`
	Username string       `json:"username"`
	Avatar   string       `json:"avatar"`
	Roles    []SimpleRole `json:"roles"`
}

func (table *User) Add() error {
	engine := database.GetXOrmEngine()
	table.Id = utils.GetUUID()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(UserTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *User) Update() error {
	engine := database.GetXOrmEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(UserTableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *User) Delete() error {
	engine := database.GetXOrmEngine()
	affected, err := engine.Table(UserTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *User) GetOne(id string) error {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(UserTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *User) UserFindByUsernameAndPassword(username string, password string) bool {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(UserTableName).Where("account = ? and pwd = ?", username, password).Desc("id").Get(table)
	if err != nil {
		return false
	}
	if table.Id == "" {
		return false
	}
	return true
}
