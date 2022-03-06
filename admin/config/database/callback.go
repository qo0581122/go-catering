package database

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

func UpdateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Format("2006-01-02 15:04:05")
		if createTimeField, ok := scope.FieldByName("CreatedTime"); ok {
			err := createTimeField.Set(nowTime)
			fmt.Println(err)
		}
		if modifyTimeField, ok := scope.FieldByName("UpdatedTime"); ok {
			modifyTimeField.Set(nowTime)
		}
	}
}
func UpdateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		err := scope.SetColumn("UpdatedTime", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(err)
	}
}

func DeleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedTime")

		if !scope.Search.Unscoped && hasDeletedOnField {
			nowTime := time.Now().Format("2006-01-02 15:04:05")
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(nowTime),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
