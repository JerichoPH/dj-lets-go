package models

type (
	// AuthorizationAccount 用户模型
	AuthorizationAccount struct {
		GormModel
		Username string               `gorm:"type:varchar(64);unique;not null;comment:;" json:"username" json:"username"`
		Password string               `gorm:"type:varchar(64);comment:;" json:"password" json:"password"`
		Nickname string               `gorm:"type:varchar(64);unique;not null;comment:用户昵称;" json:"nickname" json:"nickname"`
		Roles    []*AuthorizationRole `gorm:"many2many:authorization_pivot_role_and_accounts;foreignKey:uuid;joinForeignKey:role_uuid;references:uuid;joinReferences:account_uuid;comment:相关角色;" json:"roles"`
		Videos   []*VideoFile         `gorm:"foreignKey:uploader_uuid;references:uuid;comment:相关视频;" json:"videos"`
	}

	// AuthorizationRole 角色模型
	AuthorizationRole struct {
		GormModel
		Name        string               `gorm:"type:varchar(64);unique;not nul;comment:角色名称;" json:"name"`
		Accounts    []*AuthorizationRole `gorm:"many2many:authorization_pivot_role_and_accounts;foreignKey:uuid;joinForeignKey:account_uuid;references:uuid;joinReferences:role_uuid;comment:相关用户;" json:"accounts"`
		Permissions []*AuthorizationRole `gorm:"many2many:authorization_pivot_role_and_permissions;foreignKey:uuid;joinForeignKey:permission_uuid;references:uuid;joinReferences:role_uuid;comment:相关权限;" json:"permissions"`
		Menus       []*AuthorizationRole `gorm:"many2many:authorization_pivot_role_and_menus;foreignKey:uuid;joinForeignKey:menu_uuid;references:uuid;joinReferences:role_uuid;comment:相关菜单;" json:"menus"`
	}

	// AuthorizationPermission 权限模型
	AuthorizationPermission struct {
		GormModel
		Name  string               `gorm:"type:varchar(64);unique;not null;comment:权限名称;" json:"name"`
		Roles []*AuthorizationRole `gorm:"many2many:authorization_pivot_role_and_permissions;foreignKey:uuid;joinForeignKey:role_uuid;references:uuid;joinReferences:permission_uuid;comment:相关角色;" json:"roles"`
	}

	// AuthorizationMenu 菜单模型
	AuthorizationMenu struct {
		GormModel
		Title      string               `gorm:"type:varchar(64);unique;not null;comment:菜单标题;" json:"title"`
		SubTitle   string               `gorm:"type:varchar(64);not null;default:'';comment:副标题;" json:"sub_title"`
		ParentUuid string               `gorm:"type:varchar(36);index;not null;default:'';comment:所属上级uuid;" json:"parent_uuid"`
		Parent     *AuthorizationMenu   `gorm:"foreignKey:parent_uuid;references:uuid;所属上级;" json:"parent"`
		Subs       []*AuthorizationMenu `gorm:"foreignKey:parent_uuid;references:uuid;相关子集;" json:"subs"`
		Roles      []*AuthorizationRole `gorm:"many2many:authorization_pivot_role_and_menus;foreignKey:uuid;joinForeignKey:role_uuid;references:uuid;joinReferences:menu_uuid;comment:相关角色;" json:"roles"`
	}
)

// TableName 用户表名称
func (receiver AuthorizationAccount) TableName() string {
	return "authorization_accounts"
}

// TableName 角色表名称
func (receiver AuthorizationRole) TableName() string {
	return "authorization_roles"
}

// TableName 权限表名称
func (receiver AuthorizationPermission) TableName() string {
	return "authorization_permissions"
}

// TableName 菜单表名称
func (receiver AuthorizationMenu) TableName() string {
	return "authorization_menus"
}
