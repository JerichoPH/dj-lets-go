package wrongs

import (
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

type Wrong struct{ ErrorMessage string }

// EmptyWrong 空数据异常
type EmptyWrong struct{ Wrong }

// Error 获取异常信息
//
//	@receiver receiver
//	@return string
func (receiver *Wrong) Error() string {
	return receiver.ErrorMessage
}

// ValidateWrong 表单验证错误
type ValidateWrong struct{ Wrong }

// PanicValidate 421错误
//
//	@param text
func PanicValidate(text string, more ...interface{}) {
	panic(&ValidateWrong{Wrong{ErrorMessage: fmt.Sprintf(text, more...)}})
}

// PanicEmpty 404错误
//
//	@param text
//	@return error
func PanicEmpty(text string, more ...interface{}) {
	panic(&EmptyWrong{Wrong{ErrorMessage: fmt.Sprintf(text, more...)}})
}

// ForbiddenWrong 操作错误
type ForbiddenWrong struct{ Wrong }

// PanicForbidden 403错误
//
//	@param text
//	@return error
func PanicForbidden(text string, more ...interface{}) {
	panic(&ForbiddenWrong{Wrong{ErrorMessage: fmt.Sprintf(text, more...)}})
}

// UnAuthWrong 未授权异常
type UnAuthWrong struct{ Wrong }

// PanicUnAuth 未授权错误
//
//	@param text
//	@return error
func PanicUnAuth(text string, more ...interface{}) {
	panic(&UnAuthWrong{Wrong{ErrorMessage: fmt.Sprintf(text, more...)}})
}

// UnLoginWrong 未登录异常
type UnLoginWrong struct{ Wrong }

// PanicUnLogin 未登录错误
//
//	@param text
func PanicUnLogin(text string, more ...interface{}) {
	panic(&UnLoginWrong{Wrong{ErrorMessage: fmt.Sprintf(text, more...)}})
}

// PanicWhenIsNotInt 文字转整型
//
//	@param v
//	@param errMsg
//	@return intValue
func PanicWhenIsNotInt(strValue string, errorMessage string) (intValue int) {
	intValue, err := strconv.Atoi(strValue)
	if err != nil && errorMessage != "" {
		PanicForbidden(errorMessage)
	}
	return
}

// PanicWhenIsNotUint 文字转无符号整型
//
//	@param v
//	@param errMsg
//	@return uintValue
func PanicWhenIsNotUint(strValue string, errorMessage string) (uintValue uint) {
	intValue := PanicWhenIsNotInt(strValue, errorMessage)
	uintValue = uint(intValue)
	return
}

// PanicWhenIsEmpty 当数据库返回空则报错
//
//	@param db
//	@param name
//	@return bool
func PanicWhenIsEmpty(db *gorm.DB, errorField string) bool {
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			if errorField != "" {
				PanicEmpty("[%s] 不存在", errorField)
				return true
			} else {
				return true
			}
		}
	}
	return false
}

// PanicWhenIsRepeat 当数据库返回不空则报错
//
//	@param db
//	@param name
//	@return bool
func PanicWhenIsRepeat(db *gorm.DB, errorField string) bool {
	if db.Error == nil {
		if errorField != "" {
			PanicForbidden("[%s] 重复", errorField)
			return false
		} else {
			return false
		}
	}
	return true
}
