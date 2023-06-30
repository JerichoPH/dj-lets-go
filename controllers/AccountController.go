package controllers

import (
	"dj-lets-go/models"
	"dj-lets-go/services"
	"dj-lets-go/tools"
	"dj-lets-go/types"
	"dj-lets-go/wrongs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	// AccountController 用户控制器
	AccountController struct{}
	// accountStoreForm 用户表单
	accountStoreForm struct{}
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
		ret *gorm.DB
		// repeat models.AccountModel
	)

	// 新建
	account := &models.AccountModel{}
	if ret = models.NewGormModel().SetModel(models.AccountModel{}).
		DB("", nil).
		Create(&account); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Created(types.MapStringToAny{"account": account}).ToGinResponse())
}

// Delete 删除
func (AccountController) Delete(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AccountModel
	)

	// 查询
	ret = models.NewGormModel().SetModel(models.AccountModel{}).
		SetWheres(types.MapStringToAny{"uuid": ctx.Param("uuid")}).
		DB("", nil).
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	// 删除
	if ret := models.NewGormModel().SetModel(models.AccountModel{}).DB("", nil).Delete(&account); ret.Error != nil {
		wrongs.PanicForbidden(ret.Error.Error())
	}

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Deleted().ToGinResponse())
}

// Update 编辑
func (AccountController) Update(ctx *gin.Context) {
	var (
		ret     *gorm.DB
		account models.AccountModel
		// repeat  models.AccountModel
	)

	// 表单
	// form := new(accountStoreForm).ShouldBind(ctx)

	// 查询
	ret = models.NewGormModel().SetModel(models.AccountModel{}).
		SetWheres(types.MapStringToAny{"uuid": ctx.Param("uuid")}).
		DB("", nil).
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	// 编辑
	if ret = models.NewGormModel().SetModel(models.AccountModel{}).
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
		account models.AccountModel
	)
	ret = models.NewGormModel().SetModel(models.AccountModel{}).
		SetWheres(types.MapStringToAny{"uuid": ctx.Param("uuid")}).
		DB("", nil).
		First(&account)
	wrongs.PanicWhenIsEmpty(ret, "用户")

	ctx.JSON(tools.NewCorrectWithGinContext("", ctx).Datum(types.MapStringToAny{"account": account}).ToGinResponse())
}

func (AccountController) listByQuery(ctx *gin.Context) *gorm.DB {
	return services.NewAccountService(services.BaseService{Model: models.NewGormModel().SetModel(models.AccountModel{}), Ctx: ctx}).GetListByQuery()
}

// List 列表
func (receiver AccountController) List(ctx *gin.Context) {
	var accounts []models.AccountModel

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
	var accounts []models.AccountModel

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
