package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
	"github.com/hyahm/golog"
)

const RoleTableName string = "Role"

type Role struct {
	CreatedBy string `form:"created_by" json:"created_by" xorm:"varchar(255)" binding:"max=255"`

	CreatedTime string `form:"created_time" json:"created_time" xorm:"varchar(255)" binding:"max=255"`

	Id string `form:"id" json:"id" xorm:"varchar(255) pk notnull" binding:"max=255"`

	Menu string `form:"menu" json:"menu" xorm:"text" binding:"max=65535"`

	OrderNo string `form:"order_no" json:"order_no" xorm:"varchar(255)" binding:"max=255"`

	Remark string `form:"remark" json:"remark" xorm:"varchar(255)" binding:"max=255"`

	Revision int64 `form:"revision" json:"revision" xorm:"bigint"`

	RoleName string `form:"role_name" json:"role_name" xorm:"varchar(255)" binding:"max=255"`

	RoleValue string `form:"role_value" json:"role_value" xorm:"varchar(255)" binding:"max=255"`

	Status string `form:"status" json:"status" xorm:"varchar(255)" binding:"max=255"`

	UpdatedBy string `form:"updated_by" json:"updated_by" xorm:"varchar(255)" binding:"max=255"`

	UpdatedTime string `form:"updated_time" json:"updated_time" xorm:"varchar(255)" binding:"max=255"`
}

type SimpleRole struct {
	RoleName string `json:"roleName"`
	Value    string `json:"value"`
}

func (table *Role) Add() error {
	engine := database.GetXOrmEngine()
	table.Id = utils.GetUUID()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(RoleTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Role) Update() error {
	engine := database.GetXOrmEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(RoleTableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Role) Delete() error {
	engine := database.GetXOrmEngine()
	affected, err := engine.Table(RoleTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *Role) GetOne(id string) error {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(RoleTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Role) GetRoleByRoleValue() error {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(RoleTableName).Where("role_value = ?", table.RoleValue).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Role) GetRoleByRoleValues(ids []string) ([]Role, error) {
	engine := database.GetXOrmEngine()
	var roles []Role
	err := engine.Table(RoleTableName).In("role_value", ids).Desc("id").Find(&roles)
	if err != nil {
		golog.Info(err.Error())
		return nil, err
	}
	return roles, nil
}
