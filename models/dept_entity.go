package models

import (
	"errors"

	"github.com/golang-module/carbon"
	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
)

const DeptTableName string = "Dept"

type Dept struct {
	Id string `form:"id" json:"id" xorm:"pk notnull"`

	Revision int64 `form:"revision" json:"revision"`

	CreatedBy string `form:"created_by" json:"created_by"`

	CreatedTime string `form:"created_time" json:"created_time"`

	UpdatedBy string `form:"updated_by" json:"updated_by"`

	UpdatedTime string `form:"updated_time" json:"updated_time"`

	DeptName string `form:"dept_name" json:"dept_name"`

	ParentDept string `form:"parent_dept" json:"parent_dept"`

	OrderNo int64 `form:"order_no" json:"order_no"`

	Remark string `form:"remark" json:"remark"`

	Status string `form:"status" json:"status"`
}

func (table *Dept) Add() error {
	engine := database.GetXOrmEngine()
	table.Id = utils.GetUUID()
	table.CreatedTime = carbon.Now().ToDateTimeString()
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(DeptTableName).Insert(&table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Dept) Update() error {
	engine := database.GetXOrmEngine()
	table.CreatedBy = ""
	table.CreatedTime = ""
	table.UpdatedTime = carbon.Now().ToDateTimeString()
	_, err := engine.Table(DeptTableName).Where("id = ?", table.Id).Update(table)
	if err != nil {
		return err
	}
	return nil
}

func (table *Dept) Delete() error {
	engine := database.GetXOrmEngine()
	affected, err := engine.Table(DeptTableName).Where("id = ?", table.Id).Delete()
	if affected == 0 {
		return errors.New("没有找到删除的数据")
	}
	if err != nil {
		return err
	}
	return nil
}

func (table *Dept) GetOne(id string) error {
	engine := database.GetXOrmEngine()
	_, err := engine.Table(DeptTableName).Where("id = ?", id).Desc("id").Get(table)
	if err != nil {
		return err
	}
	return nil
}
