package models

type (
	// AccountModel 用户模型
	AccountModel struct {
		GormModel
		Username string `gorm:"type:varchar(191);unique;not null;comment:;" json:"username"`
		Password string `gorm:"type:varchar(191);comment:;" json:"password"`
		Nickname string `gorm:"type:varchar(64);unique;not null;comment:用户昵称;" json:"nickname"`
	}
)

// TableName 表名称
func (receiver *AccountModel) TableName() string {
	return "accounts"
}
