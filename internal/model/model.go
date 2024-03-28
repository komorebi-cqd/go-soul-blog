package model

import (
	"fmt"
	"time"

	"github.com/go-soul-blog/global"
	"github.com/go-soul-blog/pkg/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"`
	// CreatedBy  string `json:"created_by"`
	// ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", databaseSetting.Username, databaseSetting.Password, databaseSetting.Host, databaseSetting.DBName, databaseSetting.Charset, databaseSetting.ParseTime)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(0)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 回调
	db.Callback().Create().Before("gorm:create").Register("update_created_at", updateTimeStampForCreateCallback)
	db.Callback().Create().Before("gorm:update").Register("update_updated_at", updateTimeStampForUpdateCallback)

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

func updateTimeStampForCreateCallback(scope *gorm.DB) {
	fmt.Printf("1---updateTimeStampForCreateCallback \n")
	if scope.Statement.Schema != nil {
		nowTime := time.Now().Unix()
		fmt.Printf("2---updateTimeStampForCreateCallback \n")

		field := scope.Statement.Schema.LookUpField("CreatedOn")
		if field == nil {
			return
		}
		if _, isZero := field.ValueOf(scope.Statement.Context, scope.Statement.ReflectValue); isZero {
			scope.Statement.SetColumn("CreatedOn", nowTime)

			fmt.Printf("isZero %v cValue \n", scope.Statement.ReflectValue)
			// field.Set(scope.Statement.Context, scope.Statement.ReflectValue, nowTime)
			// _ = scope.InstanceSet("CreatedOn", time.Now().Unix())
		}

		field = scope.Statement.Schema.LookUpField("ModifiedOn")
		if _, isZero := field.ValueOf(scope.Statement.Context, scope.Statement.ReflectValue); isZero {
			field.Set(scope.Statement.Context, scope.Statement.ReflectValue, nowTime)
		}
	}
}

func updateTimeStampForUpdateCallback(scope *gorm.DB) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		fmt.Printf("111---updateTimeStampForUpdateCallback \n")
		_ = scope.InstanceSet("ModifiedOn", time.Now().Unix())
	}
}

// func deleteCallback(scope *gorm.DB) {
// 	if scope.Statement.Schema != nil {
// 		var extraOption string
// 		if str, ok := scope.Get("gorm:delete_option"); ok {
// 			extraOption = fmt.Sprint(str)
// 		}

// 		deletedOnField := scope.Statement.Schema.LookUpField("DeletedOn")
// 		isDelField := scope.Statement.Schema.LookUpField("IsDel")
// 		_, isDeletedOn := deletedOnField.ValueOf(scope.Statement.Context, scope.Statement.ReflectValue)
// 		_, isDel := isDelField.ValueOf(scope.Statement.Context, scope.Statement.ReflectValue)

// 		if isDeletedOn && isDel {

// 		} else {
// 			scope.Raw(fmt.Sprintf("DELETE FROM %v%v%v", deletedOnField.Schema.Table, addExtraSpaceIfExist(scope.CombinedConditionSql())), addExtraSpaceIfExist(extraOption))
// 		}
// 	}
// }

// func addExtraSpaceIfExist(str string) string {
// 	if str != "" {
// 		return " " + str
// 	}
// 	return ""
// }
