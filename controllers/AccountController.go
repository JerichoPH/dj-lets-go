package controllers

import (
	"dj-lets-go/models"
	"dj-lets-go/services"
	"dj-lets-go/tools"
	"dj-lets-go/types"
	"dj-lets-go/wrongs"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// AccountController 用户控制器
	AccountController struct{}
	// accountStoreForm 用户表单
	accountStoreForm struct {
		Username             string `form:"username" json:"username" binding:"required"`
		Password             string `form:"password" json:"password" binding:"required"`
		PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
		Nickname             string `form:"nickname" json:"nickname" binding:"required"`
	}
)

// NewAccountController 构造函数
func NewAccountController() *AccountController {
	return &AccountController{}
}

// ShouldBind 表单绑定
func (receiver accountStoreForm) ShouldBind(ctx *gin.Context) accountStoreForm {
	if err := ctx.ShouldBind(&receiver); err != nil {
		wrongs.PanicValidate(err.Error())
	}
	
	return receiver
}

// Store 新建
func (AccountController) Store(ctx *gin.Context) {
	var (
		ret    *gorm.DB
		repeat models.AuthorizationAccount
	)
	
	// 表单验证
	form := (&accountStoreForm{}).ShouldBind(ctx)
	
	// 检查重复项（用户名）
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
	
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(types.MapStringToAny{"account": account}).ToGinResponse())
}

// Delete 删除
func (AccountController) Delete(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AuthorizationAccount
	)
	
	// 查询
	ret = models.NewGormModel().SetModel(models.AuthorizationAccount{}).
		SetWheres(types.MapStringToAny{"uuid": ctx.Param("uuid")}).
		DB("", nil).
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")
	
	// 删除
	if ret := models.NewGormModel().SetModel(models.AuthorizationAccount{}).DB("", nil).Delete(&account); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}
	
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (AccountController) Update(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AuthorizationAccount
		// repeat  models.AuthorizationAccount
	)
	
	// 表单
	// form := new(accountStoreForm).ShouldBind(ctx)
	
	// 查询
	ret = models.NewGormModel().SetModel(models.AuthorizationAccount{}).
		SetWheres(types.MapStringToAny{"uuid": ctx.Param("uuid")}).
		DB("", nil).
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")
	
	// 编辑
	if ret = models.NewGormModel().SetModel(models.AuthorizationAccount{}).
		DB("", nil).
		Where("uuid = ?", ctx.Param("uuid")).
		Save(&account); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}
	
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Updated(types.MapStringToAny{"account": account}).ToGinResponse())
}

// Detail 详情
func (AccountController) Detail(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AuthorizationAccount
	)
	ret = models.NewGormModel().SetModel(models.AuthorizationAccount{}).
		SetWheres(types.MapStringToAny{"uuid": ctx.Param("uuid")}).
		DB("", nil).
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")
	
	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(types.MapStringToAny{"account": account}).ToGinResponse())
}

func (AccountController) listByQuery(ctx *gin.Context) *gorm.DB {
	return services.NewAccountService(services.BaseService{Model: models.NewGormModel().SetModel(models.AuthorizationAccount{}), Ctx: ctx}).GetListByQuery()
}

// List 列表
func (receiver AccountController) List(ctx *gin.Context) {
	var accounts []models.AuthorizationAccount
	
	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForPager(
				receiver.listByQuery(ctx),
				func(db *gorm.DB) types.MapStringToAny {
					db.Find(&accounts)
					return types.MapStringToAny{"accounts": accounts}
				},
			).ToGinResponse(),
	)
}

// ListJdt jquery-dataTable分页列表
func (receiver AccountController) ListJdt(ctx *gin.Context) {
	var accounts []models.AuthorizationAccount
	
	ctx.JSON(
		tools.NewCorrectWithGinContext("", ctx).
			DataForJqueryDataTable(
				receiver.listByQuery(ctx),
				func(db *gorm.DB) types.MapStringToAny {
					db.Find(&accounts)
					return types.MapStringToAny{"accounts": accounts}
				},
			).ToGinResponse(),
	)
}
