package controllers

import (
	"fmt"

	"dj-lets-go/providers"
	"dj-lets-go/tools"
	"dj-lets-go/types"

	"github.com/gin-gonic/gin"
)

type (
	// TestController 测试控制器
	TestController struct{}

	// 发送消息表单
	sendMessageForm struct {
		ReceiverUuid   string `json:"receiver_uuid"`
		MessageTitle   string `json:"message_title"`
		MessageContent string `json:"message_content"`
	}
)

// NewTestController 构造函数
func NewTestController() *TestController {
	return &TestController{}
}

// ShouldBind 表单绑定
func (receiver sendMessageForm) ShouldBind(ctx *gin.Context) sendMessageForm {
	err := ctx.ShouldBind(&receiver)
	if err != nil {
		return sendMessageForm{}
	}

	return receiver
}

// AnySendToWebsocket 接收并转发消息(websocket)
func (receiver TestController) AnySendToWebsocket(ctx *gin.Context) {
	form := sendMessageForm{}.ShouldBind(ctx)

	providers.WebsocketSendMessageByUuid(tools.NewCorrectWithBusiness(form.MessageTitle, "message", "").Datum(form.MessageContent).ToJsonStr(), form.ReceiverUuid)

	ctx.JSON(tools.NewCorrectWithGinContext("OK", ctx).Datum(types.MapStringToAny{"content": fmt.Sprintf("%s ==> %s %v", form.ReceiverUuid, form.MessageTitle, form.MessageContent)}).ToGinResponse())
}

// AnySendToTcpServer 接收并转发消息(tcp server)
func (receiver TestController) AnySendToTcpServer(ctx *gin.Context) {
	form := sendMessageForm{}.ShouldBind(ctx)

	providers.TcpServerSendMessageByUuid(tools.NewCorrectWithBusiness(form.MessageTitle, "message", "").Datum(form.MessageContent).ToJsonStr(), form.ReceiverUuid)

	ctx.JSON(tools.NewCorrectWithGinContext("OK", ctx).Datum(types.MapStringToAny{"content": fmt.Sprintf("%s ==> %s %v", form.ReceiverUuid, form.MessageTitle, form.MessageContent)}).ToGinResponse())
}

// AnySendToTcpClient 接收病转发消息(tcp client)
func (receiver TestController) AnySendToTcpClient(ctx *gin.Context) {
	form := sendMessageForm{}.ShouldBind(ctx)

	providers.TcpClientSendMessage(tools.NewCorrectWithBusiness(form.MessageTitle, "message", "").Datum(types.MapStringToAny{"content": fmt.Sprintf("%s %v", form.MessageTitle, form.MessageContent)}).ToJsonStr())

	ctx.JSON(tools.NewCorrectWithGinContext("OK", ctx).Datum(types.MapStringToAny{"content": fmt.Sprintf("%s %v", form.MessageTitle, form.MessageContent)}).ToGinResponse())
}
