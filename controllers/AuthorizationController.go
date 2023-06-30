package controllers

import (
	"dj-lets-go/models"
	"dj-lets-go/tools"
	"dj-lets-go/types"
	"dj-lets-go/wrongs"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	// AuthorizationController 权鉴控制器
	AuthorizationController struct{}
	// authorizationRegisterForm 注册表单
	authorizationRegisterForm struct {
		Account              string `form:"account" json:"account" binding:"required"`
		Password             string `form:"password" json:"password" binding:"required"`
		PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
		Nickname             string `form:"nickname" json:"nickname" binding:"required"`
	}
)

// NewAuthorizationController 构造函数
func NewAuthorizationController() *AuthorizationController {
	return &AuthorizationController{}
}

// ShouldBind 绑定表单
//
//	@receiver ins
//	@param ctx
//	@return authorizationRegisterForm
func (receiver authorizationRegisterForm) ShouldBind(ctx *gin.Context) authorizationRegisterForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if receiver.Account == "" {
		wrongs.PanicValidate("账号必填")
	}
	if receiver.Password == "" {
		wrongs.PanicValidate("密码必填")
	}
	if len(receiver.Password) < 6 || len(receiver.Password) > 18 {
		wrongs.PanicValidate("密码不可小于6位或大于18位")
	}
	if receiver.Password != receiver.PasswordConfirmation {
		wrongs.PanicValidate("两次密码输入不一致")
	}

	return receiver
}

// AuthorizationLoginForm 登录表单
type AuthorizationLoginForm struct {
	Username string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// ShouldBind 绑定表单
func (receiver AuthorizationLoginForm) ShouldBind(ctx *gin.Context) AuthorizationLoginForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	if receiver.Username == "" {
		wrongs.PanicValidate("账号必填")
	}
	if receiver.Password == "" {
		wrongs.PanicValidate("密码必填")
	}
	if len(receiver.Password) < 6 || len(receiver.Password) > 18 {
		wrongs.PanicValidate("密码不可小于6位或大于18位")
	}

	return receiver
}

// Register 注册
func (AuthorizationController) Register(ctx *gin.Context) {
	// 表单验证
	form := (&authorizationRegisterForm{}).ShouldBind(ctx)

	// 检查重复项（用户名）
	var repeat models.AuthorizationAccount
	var ret *gorm.DB
	ret = (&models.GormModel{}).
		SetWheres(types.MapStringToAny{"username": form.Account}).
		DB("", nil).
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "用户名")
	ret = (&models.GormModel{}).
		SetWheres(types.MapStringToAny{"nickname": form.Nickname}).
		DB("", nil).
		First(&repeat)
	wrongs.PanicWhenIsRepeat(ret, "昵称")

	// 密码加密
	bytes, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 14)

	// 保存新用户
	account := &models.AuthorizationAccount{
		GormModel: models.GormModel{Uuid: uuid.NewV4().String()},
		Username:  form.Account,
		Password:  string(bytes),
		Nickname:  form.Nickname,
	}
	if ret = models.NewGormModel().SetModel(models.AuthorizationAccount{}).
		SetOmits(clause.Associations).
		DB("", nil).
		Create(&account); ret.Error != nil {
		wrongs.PanicForbidden("创建失败：" + ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("注册成功", ctx).Created(types.MapStringToAny{"account": account}).ToGinResponse())
}

// Login 登录
func (AuthorizationController) Login(ctx *gin.Context) {
	// 表单验证
	form := (&AuthorizationLoginForm{}).ShouldBind(ctx)

	var (
		account models.AuthorizationAccount
		ret     *gorm.DB
	)
	// 获取用户
	ret = models.NewGormModel().
		SetModel(models.AuthorizationAccount{}).
		SetPreloads("Workshop", "Station", "WorkAreaByUniqueCode").
		SetWheres(types.MapStringToAny{"account": form.Username}).
		DB("", nil).
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(form.Password)); err != nil {
		wrongs.PanicUnAuth("账号或密码错误")
	}

	// 生成Jwt
	if token, err := tools.GenerateJwt(
		account.Username,
		account.Nickname,
	); err != nil {
		// 生成jwt错误
		wrongs.PanicForbidden(err.Error())
	} else {
		ctx.JSON(tools.NewCorrectWithGinContext("登陆成功", ctx).Datum(types.MapStringToAny{
			"token": token,
			"account": types.MapStringToAny{
				"username": account.Username,
				"nickname": account.Nickname,
			},
		}).ToGinResponse())
	}
}

// GetMenus 获取当前用户菜单
// func (AuthorizationController) GetMenus(ctx *gin.Context) {
//	var ret *gorm.DB
//	if accountUuid, exists := ctx.Get(tools.AccountOpenIdFieldName); !exists {
//		wrongs.PanicUnLogin("用户未登录")
//	} else {
//		// 获取当前用户信息
//		var account models.AuthorizationAccount
//		ret = models.NewGormModel().SetModel(models.AuthorizationAccount{}).
//			SetWheres(types.MapStringToAny{"uuid": accountUuid}).
//			SetPreloads("RbacRoles", "RbacRoles.Menus").
//			DB("",nil).
//			FindOneUseQuery(&account)
//		if !wrongs.PanicWhenIsEmpty(ret, "") {
//			wrongs.PanicUnLogin("当前令牌指向用户不存在")
//		}
//
//		var menus []models.MenuModel
//		models.NewGormModel().SetModel(models.MenuModel{}).
//			DB("",nil).
//			Joins("join pivot_rbac_role_and_menus prram on menus.uuid = prram.menu_uuid").
//			Joins("join rbac_roles r on prram.rbac_role_uuid = r.uuid").
//			Joins("join pivot_rbac_role_and_accounts prraa on r.uuid = prraa.rbac_role_uuid").
//			Joins("join accounts a on prraa.account_uuid = a.uuid").
//			Where("a.uuid = ?", account.GormModel.Uuid).
//			Where("menus.deleted_at is null").
//			Where("menus.parent_uuid = ''").
//			Order("menus.sort asc").
//			Order("menus.id asc").
//			Preload("Subs").
//			Find(&menus)
//
//		ctx.JSON(tools.CorrectInit("", ctx).Datum(types.MapStringToAny{"menus": menus}))
//	}
// }
