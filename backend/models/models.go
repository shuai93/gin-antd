package models

import (
	"backend/utils/logging"
	"backend/utils/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

func GetMysqlConnUrl() string {

	var (
		dbName, user, password, host string
	)

	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
}

func Setup() {
	var err error

	db, err = gorm.Open(mysql.Open(GetMysqlConnUrl()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Println(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Println(err)
	}

	//  设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	//  设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.Callback().Create().Before("gorm:create").Register("update_created_at", updateTimeStampForCreateCallback)

}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(db *gorm.DB) {
	if db.Statement.Schema != nil {
		nowTime := time.Now().Unix()
		field := db.Statement.Schema.LookUpField("created_on")

		if field != nil {
			_ = field.Set(db.Statement.ReflectValue, nowTime)
			logging.Info("%v", field)
		}

	}
}

//// updateTimeStampForUpdateCallback will set `ModifyTime` when updating
//func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
//	if _, ok := scope.Get("gorm:update_column"); !ok {
//		scope.SetColumn("ModifiedOn", time.Now().Unix())
//	}
//}
//
//func deleteCallback(scope *gorm.Scope) {
//	if !scope.HasError() {
//		var extraOption string
//		if str, ok := scope.Get("gorm:delete_option"); ok {
//			extraOption = fmt.Sprint(str)
//		}
//
//		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
//
//		if !scope.Search.Unscoped && hasDeletedOnField {
//			scope.Raw(fmt.Sprintf(
//				"UPDATE %v SET %v=%v%v%v",
//				scope.QuotedTableName(),
//				scope.Quote(deletedOnField.DBName),
//				scope.AddToVars(time.Now().Unix()),
//				addExtraSpaceIfExist(scope.CombinedConditionSql()),
//				addExtraSpaceIfExist(extraOption),
//			)).Exec()
//		} else {
//			scope.Raw(fmt.Sprintf(
//				"DELETE FROM %v%v%v",
//				scope.QuotedTableName(),
//				addExtraSpaceIfExist(scope.CombinedConditionSql()),
//				addExtraSpaceIfExist(extraOption),
//			)).Exec()
//		}
//	}
//}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
