package middlewares

import (
	"dj-lets-go/constants"
	"dj-lets-go/models"
	"dj-lets-go/tools"
	"dj-lets-go/types"
	"dj-lets-go/wrongs"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CheckAuthorization 检查Jwt是否合法
func CheckAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取令牌
		split := strings.Split(tools.GetJwtFromHeader(ctx), " ")
		if len(split) != 2 {
			wrongs.PanicUnAuth("令牌格式错误")
		}
		tokenType := split[0]
		token := split[1]

		var (
			account models.AuthorizationAccount
			ret     *gorm.DB
		)
		if token == "" {
			wrongs.PanicUnAuth("令牌不存在")
		} else {
			switch tokenType {
			case "JWT":
				claims, err := tools.ParseJwt(token)

				// 判断令牌是否有效
				if err != nil {
					wrongs.PanicUnAuth("令牌解析失败")
				} else if time.Now().Unix() > claims.ExpiresAt {
					wrongs.PanicUnAuth("令牌过期")
				}

				// 判断用户是否存在
				if reflect.DeepEqual(claims, tools.Claims{}) {
					wrongs.PanicUnAuth("令牌解析失败：用户不存在")
				}

				// 获取用户信息
				ret = models.NewGormModel().SetModel(models.AuthorizationAccount{}).DB("", nil).Where("uuid", claims.Uuid).First(&account)
				wrongs.PanicWhenIsEmpty(ret, fmt.Sprintf("令牌指向用户(JWT) %s %v ", token, claims))
			case "AU":
				ret = models.NewGormModel().SetModel(models.AuthorizationAccount{}).SetWheres(types.MapStringToAny{"uuid": token}).DB("", nil).First(&account)
				wrongs.PanicWhenIsEmpty(ret, fmt.Sprintf("令牌指向用户(AU) %s", token))
			default:
				wrongs.PanicForbidden("权鉴认证方式不支持")
			}

			ctx.Set(constants.AccountIdFieldName, account.Id)             // 设置用户编号
			ctx.Set(constants.AccountAccountFieldName, account.Username)  // 设置用户账号
			ctx.Set(constants.AccountNicknameFieldName, account.Nickname) // 设置用户昵称
			ctx.Set(constants.AccountAuthorizationFieldName, account)     // 设置用户信息
		}

		ctx.Next()
	}
}
