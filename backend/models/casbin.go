package models

import (
	"backend/utils/logging"
	"backend/utils/setting"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/util"
	gormadapter "github.com/casbin/gorm-adapter"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type CasbinModel struct {
	PType  string `json:"p_type" gorm:"column:p_type" description:"策略类型"`
	RoleId string `json:"role_id" gorm:"column:v0" description:"角色ID"`
	Path   string `json:"path" gorm:"column:v1" description:"api路径"`
	Method string `json:"method" gorm:"column:v2" description:"访问方法"`
}

func (c *CasbinModel) TableName() string {
	return "casbin_rule"
}

func (c *CasbinModel) Create() error {
	e := Casbin()
	success := e.AddPolicy(c.RoleId, c.Path, c.Method)

	if success == false {
		return nil
	}
	return nil
}

func (c *CasbinModel) List() [][]string {
	e := Casbin()
	policy := e.GetFilteredPolicy(0, c.RoleId)
	return policy
}

func Casbin() *casbin.Enforcer {

	adapter := gormadapter.NewAdapter("mysql", GetMysqlConnUrl(), true)

	enforcer := casbin.NewEnforcer(setting.CasbinSetting.ModelPath, adapter)
	enforcer.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = enforcer.LoadPolicy()
	return enforcer
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	logging.Info(key1, key2)
	return util.KeyMatch2(key1, key2)
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}
