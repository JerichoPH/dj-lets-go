package services

import (
	"fmt"
	
	"gorm.io/gorm"
)

// AccountService 用户服务
type AccountService struct{ BaseService }

// NewAccountService 构造函数
func NewAccountService(baseService BaseService) *AccountService {
	return &AccountService{BaseService: baseService}
}

// GetListByQuery 通过Query获取列表
func (receiver *AccountService) GetListByQuery() *gorm.DB {
	return (receiver.Model).
		SetWheresExtraExist(map[string]func(string, *gorm.DB) *gorm.DB{
			"username": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("a.username like ?", fmt.Sprintf("%%%s%%", value))
			},
			"nickname": func(value string, db *gorm.DB) *gorm.DB {
				return db.Where("a.nickname like ?", fmt.Sprintf("%%%s%%", value))
			},
		}).
		SetWheresDateBetween("created_at", "updated_at").
		SetCtx(receiver.Ctx).
		DbUseQuery("").
		Table("authorization_accounts as a")
}
