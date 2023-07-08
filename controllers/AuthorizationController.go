package controllers

import (
	"dj-lets-go/services"
	"log"

	"dj-lets-go/constants"
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
		Username             string `form:"username" json:"username" binding:"required"`
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
	if receiver.Username == "" {
		wrongs.PanicValidate("账号必填")
	}
	if receiver.Password == "" {
		wrongs.PanicValidate("密码必填")
	}
	if receiver.Nickname == "" {
		wrongs.PanicValidate("昵称不能为空")
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
	Username string `form:"username" json:"username" binding:"required"`
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
		SetWheres(types.MapStringToAny{"username": form.Username}).
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
		Username:  form.Username,
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
		SetWheres(types.MapStringToAny{"username": form.Username}).
		DB("", nil).
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	// 验证密码
	log.Printf("%s,%s", account.Password, form.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(form.Password)); err != nil {
		wrongs.PanicUnAuth("账号或密码错误")
	}

	// 生成Jwt
	if token, err := tools.GenerateJwt(
		account.Uuid,
		account.Username,
		account.Nickname,
	); err != nil {
		// 生成jwt错误
		wrongs.PanicForbidden(err.Error())
	} else {
		ctx.JSON(
			tools.NewCorrectWithGinContext("登陆成功", ctx).
				Datum(
					types.MapStringToAny{
						"token": token,
						"account": types.MapStringToAny{
							"created_at": account.CreatedAt,
							"updated_at": account.UpdatedAt,
							"username":   account.Username,
							"nickname":   account.Nickname,
						},
					},
				).
				ToGinResponse(),
		)
	}
}

// AnyCheckIsLogin 检查是否登录
func (receiver AuthorizationController) AnyCheckIsLogin(ctx *gin.Context) {
	_, exist := ctx.Get(constants.AccountAuthorizationFieldName)
	if !exist {
		wrongs.PanicUnLogin("")
	}

	ctx.JSON(tools.NewCorrectWithGinContext("登录成功", ctx).Datum(types.MapStringToAny{}).ToGinResponse())
}

// GetRoles 获取当前用户角色
func (AuthorizationController) GetRoles(ctx *gin.Context) {
	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			Datum(
				types.MapStringToAny{
					"roles": services.NewAuthorizationService(
						services.BaseService{
							Ctx: ctx,
							Model: models.
								NewGormModel().
								SetModel(models.AuthorizationRole{}),
						},
					).
						GetRoles(types.ListString{})}).
			ToGinResponse(),
	)
}

// GetPermissions 获取当前用户权限
func (AuthorizationController) GetPermissions(ctx *gin.Context) {
	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			Datum(
				types.MapStringToAny{
					"permissions": services.NewAuthorizationService(
						services.BaseService{
							Ctx: ctx,
							Model: models.
								NewGormModel().
								SetModel(models.AuthorizationPermission{}),
						},
					).
						GetPermissions(types.ListString{}),
				},
			).
			ToGinResponse(),
	)
}

// GetMenus 获取当前用户菜单
func (AuthorizationController) GetMenus(ctx *gin.Context) {
	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			Datum(
				types.MapStringToAny{
					"menus": services.NewAuthorizationService(
						services.BaseService{
							Ctx: ctx,
							Model: models.
								NewGormModel().
								SetModel(models.AuthorizationMenu{}),
						},
					).
						GetMenus(types.ListString{}),
				},
			).
			ToGinResponse(),
	)
}
