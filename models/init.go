package models

import (
	"fmt"
	"os"

	"github.com/goylold/lowcode/database"
	"github.com/goylold/lowcode/utils"
	"github.com/hyahm/golog"
)

func init() {
	golog.Info("同步数据库结构.........")
	engine := database.GetXOrmEngine()
	engine.Table(MenuTableName).Sync(new(Menu))
	engine.Table(RoleTableName).Sync(new(Role))
	engine.Table(UserTableName).Sync(new(User))
	engine.Table(DeptTableName).Sync(new(Dept))
	engine.Table(BalanceTableName).Sync(new(Balance))
	engine.Table(MessageTableName).Sync(new(Message))
	engine.Table(RechargeTableName).Sync(new(Recharge))
	engine.Table(SensitiveWordTableName).Sync(new(SensitiveWord))
	engine.Table(SessionTableName).Sync(new(Session))
	engine.Table(TokenConsumptionTableName).Sync(new(TokenConsumption))
	engine.Table(WebUserTableName).Sync(new(WebUser))
	engine.Table(ConversationTableName).Sync(new(Conversation))
	golog.Info("同步数据库结构完成.........")
	if utils.IsExists(".lock") {
		return
	}
	golog.Info("添加默认管理员 admin----123456")
	role := Role{
		RoleName:  "超级管理员",
		RoleValue: "admin",
		Status:    "0",
	}
	role.Add()
	userRole := Role{
		RoleName:  "普通用户",
		RoleValue: "user",
		Status:    "0",
	}
	userRole.Add()
	dept := Dept{
		DeptName: "默认部门",
		Status:   "0",
	}
	dept.Add()

	adminUser := User{
		Account:  "admin",
		Username: "超级管理员",
		Pwd:      utils.MD5("123456"),
		Nickname: "超级管理员",
		LoginId:  "admin",
		Role:     role.RoleValue,
		Dept:     dept.Id,
		Email:    "goylord2@gmail.com",
	}
	err := adminUser.Add()
	if err != nil {
		golog.Info("添加默认管理员账号失败", err.Error())
	} else {
		golog.Info("添加默认管理员账号成功...")
	}
	dashboardRootMenu := Menu{
		MenuName:  "仪表盘",
		Status:    "0",
		Component: "LAYOUT",
		Icon:      "ant-design:dashboard-outlined",
		RoutePath: "/dashboard",
		Show:      "0",
		OrderNo:   1,
		Type:      "0",
		IsExt:     "0",
		Keepalive: "0",
	}
	dashboardRootMenu.Add()
	analysisMenu := Menu{
		MenuName:   "分析页",
		Status:     "0",
		Component:  "dashboard/analysis/index.vue",
		Icon:       "ant-design:area-chart-outlined",
		RoutePath:  "analysis",
		ParentMenu: dashboardRootMenu.Id,
		Show:       "0",
		OrderNo:    0,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	analysisMenu.Add()
	workbenchMenu := Menu{
		MenuName:   "工作台",
		Status:     "0",
		Component:  "dashboard/workbench/index.vue",
		Icon:       "ant-design:calendar-twotone",
		RoutePath:  "workbench",
		ParentMenu: dashboardRootMenu.Id,
		Show:       "0",
		OrderNo:    1,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	workbenchMenu.Add()
	// 创建菜单
	systemMenu := Menu{
		MenuName:  "系统管理",
		Status:    "0",
		Component: "LAYOUT",
		Icon:      "ant-design:setting-outlined",
		RoutePath: "/system",
		Show:      "0",
		OrderNo:   5,
		Type:      "0",
		IsExt:     "0",
		Keepalive: "0",
	}
	systemMenu.Add()
	accountMenu := Menu{
		MenuName:   "账号管理",
		Status:     "0",
		Component:  "system/account/index.vue",
		Icon:       "ant-design:user-add-outlined",
		RoutePath:  "account",
		ParentMenu: systemMenu.Id,
		Show:       "0",
		OrderNo:    0,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	accountMenu.Add()
	accountDetailMenu := Menu{
		MenuName:   "账号详情",
		Status:     "0",
		Component:  "system/account/AccountDetail.vue",
		Icon:       "ant-design:appstore-outlined",
		RoutePath:  "account_detail/:id",
		ParentMenu: systemMenu.Id,
		Show:       "1",
		OrderNo:    2,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	accountDetailMenu.Add()
	roleMenu := Menu{
		MenuName:   "角色管理",
		Status:     "0",
		Component:  "system/role/index.vue",
		Icon:       "ant-design:android-outlined",
		RoutePath:  "role",
		ParentMenu: systemMenu.Id,
		Show:       "0",
		OrderNo:    1,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	roleMenu.Add()

	menuMenu := Menu{
		MenuName:   "菜单管理",
		Status:     "0",
		Component:  "system/menu/index.vue",
		Icon:       "ant-design:menu-outlined",
		RoutePath:  "menu",
		ParentMenu: systemMenu.Id,
		Show:       "0",
		OrderNo:    2,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	menuMenu.Add()
	deptMenu := Menu{
		MenuName:   "部门管理",
		Status:     "0",
		Component:  "system/dept/index.vue",
		Icon:       "ant-design:deployment-unit-outlined",
		RoutePath:  "dept",
		ParentMenu: systemMenu.Id,
		Show:       "0",
		OrderNo:    3,
		Type:       "0",
		IsExt:      "0",
		Keepalive:  "0",
	}
	deptMenu.Add()
	role.Menu = fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s",
		dashboardRootMenu.Id, analysisMenu.Id, workbenchMenu.Id,
		systemMenu.Id, accountMenu.Id, roleMenu.Id, menuMenu.Id, deptMenu.Id)
	role.Update()
	os.Create(".lock")
}
