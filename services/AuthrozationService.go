package services

import (
	"dj-lets-go/models"
	"dj-lets-go/tools"
	"dj-lets-go/types"
)

// AuthorizationService 权鉴服务
type AuthorizationService struct{ BaseService }

// NewAuthorizationService 构造函数
func NewAuthorizationService(baseService BaseService) *AuthorizationService {
	return &AuthorizationService{BaseService: baseService}
}

// GetRoles 获取当前用户角色
func (receiver AuthorizationService) GetRoles(preloads types.ListString) []*models.AuthorizationRole {
	var (
		account models.AuthorizationAccount
		roles   = make([]*models.AuthorizationRole, 0)
	)

	account = tools.GetAuthorization(receiver.Ctx).(models.AuthorizationAccount)

	receiver.Model.
		SetModel(models.AuthorizationRole{}).
		SetPreloads(preloads...).
		DB("", nil).
		Table("authorization_roles as r").
		Joins("join authorization_pivot_role_and_accounts raa on raa.role_uuid = r.uuid").
		Where("raa.account_uuid = ?", account.Uuid).
		Find(&roles)

	return roles
}

// GetPermissions 获取当前用户权限
func (receiver AuthorizationService) GetPermissions(preloads types.ListString) []*models.AuthorizationPermission {
	var (
		account     models.AuthorizationAccount
		permissions = make([]*models.AuthorizationPermission, 0)
	)

	account = tools.GetAuthorization(receiver.Ctx).(models.AuthorizationAccount)

	receiver.Model.
		SetModel(models.AuthorizationPermission{}).
		SetPreloads(preloads...).
		DB("", nil).
		Table("authorization_permissions as p").
		Joins("join authorization_pivot_role_and_permissions rap on rap.permission_uuid = p.uuid").
		Joins("join authorization_roles r on r.uuid = rap.role_uuid").
		Joins("join authorization_pivot_role_and_accounts raa on raa.role_uuid = rap.role_uuid").
		Where("raa.account_uuid = ?", account.Uuid).
		Find(&permissions)

	return permissions
}

// GetMenus 获取当前用户菜单
func (receiver AuthorizationService) GetMenus(preloads types.ListString) []*models.AuthorizationMenu {
	var (
		account models.AuthorizationAccount
		menus   = make([]*models.AuthorizationMenu, 0)
	)

	account = tools.GetAuthorization(receiver.Ctx).(models.AuthorizationAccount)

	receiver.Model.
		SetModel(models.AuthorizationMenu{}).
		SetPreloads(preloads...).
		DB("", nil).
		Table("authorization_menus as m").
		Joins("join authorization_pivot_role_and_menus ram on ram.menu_uuid = m.uuid").
		Joins("join authorization_pivot_role_and_accounts raa on raa.role_uuid = ram.role_uuid").
		Where("raa.account_uuid = ?", account.Uuid).
		Find(&menus)

	return menus
}
