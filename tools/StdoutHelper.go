package tools

import (
	"fmt"
	"strings"
)

const (
	ORIGINAL = "\033["   // 起始标识符
	FINISHED = "\033[0m" // 终止标识符

	COLOR_BLACK  = "0" // 黑色
	COLOR_RED    = "1" //  红色
	COLOR_GREEN  = "2" // 绿色
	COLOR_YELLOW = "3" // 黄色
	COLOR_BLUE   = "4" // 蓝色
	COLOR_PURPLE = "5" // 紫色
	COLOR_CYAN   = "6" // 青色
	COLOR_WHITE  = "7" // 白色

	STYLE_DEFAULT     = "0" // 终端默认
	STYLE_DARK        = "1" // 变暗
	STYLE_HIGHLIGHT   = "2" // 高亮
	STYLE_ITALIC      = "3" // 倾斜
	STYLE_UNDERLINE   = "4" // 下横线
	STYLE_BLINK       = "5" // 闪烁
	STYLE_INVERSE     = "7" // 反白
	STYLE_INVISIBLE   = "8" // 不可见
	STYLE_DELETE_LINE = "9" // 删除线
)

type StdoutHelper struct {
	content interface{}
	fgColor string
	bgColor string
	style   string
}

// GetContent 获取内容
func (receiver *StdoutHelper) GetContent() string {
	return receiver.GetContentAndNext("")
}

func (receiver *StdoutHelper) GetContentAndNext(next interface{}) string {
	r := fmt.Sprintf("%s%v%s%v", receiver.getOriginal(), receiver.content, receiver.getFinished(), next)
	return r
}

func (receiver *StdoutHelper) Echo(next interface{}) {
	fmt.Print(receiver.GetContentAndNext(next))
}

func (receiver *StdoutHelper) EchoLine(next interface{}) {
	fmt.Println(receiver.GetContentAndNext(next))
}

// SetContent 设置内容
func (receiver *StdoutHelper) SetContent(content interface{}) *StdoutHelper {
	receiver.content = content
	return receiver
}

// GetFgColor 获取前景色
func (receiver *StdoutHelper) GetFgColor() string {
	return "3" + receiver.fgColor
}

// SetFgColor 设置前景色
func (receiver *StdoutHelper) SetFgColor(fgColor string) *StdoutHelper {
	receiver.fgColor = fgColor
	return receiver
}

// GetBgColor 获取背景色
func (receiver *StdoutHelper) GetBgColor() string {
	return "4" + receiver.bgColor
}

// SetBgColor 设置背景色
func (receiver *StdoutHelper) SetBgColor(bgColor string) *StdoutHelper {
	receiver.bgColor = bgColor
	return receiver
}

// GetStyle 获取样式
func (receiver *StdoutHelper) GetStyle() string {
	return receiver.style
}

// SetStyle 设置样式
func (receiver *StdoutHelper) SetStyle(style string) *StdoutHelper {
	receiver.style = style
	return receiver
}

// 获取起始标识符
func (receiver *StdoutHelper) getOriginal() string {
	var options []string

	if receiver.GetFgColor() != "" {
		options = append(options, receiver.GetFgColor())
	}
	if receiver.GetBgColor() != "" {
		options = append(options, receiver.GetBgColor())
	}
	if receiver.GetStyle() != "" {
		options = append(options, receiver.GetStyle())
	}

	r := ORIGINAL + strings.Join(options, ";") + "m"
	return r
}

// 获取终止标识符
func (receiver *StdoutHelper) getFinished() string {
	return FINISHED
}

// StdoutSuccess 成功格式
func StdoutSuccess(content interface{}, style string) *StdoutHelper {
	receiver := &StdoutHelper{fgColor: COLOR_GREEN, content: content}
	if style != "" {
		receiver.SetStyle(style)
	}
	return receiver
}

// StdoutInfo 消息
func StdoutInfo(content interface{}, style string) *StdoutHelper {
	receiver := &StdoutHelper{fgColor: COLOR_BLUE, content: content}
	if style != "" {
		receiver.SetStyle(style)
	}
	return receiver
}

// StdoutDebug 注释
func StdoutDebug(content interface{}, style string) *StdoutHelper {
	receiver := &StdoutHelper{fgColor: COLOR_YELLOW, content: content}
	if style != "" {
		receiver.SetStyle(style)
	}
	return receiver
}

// StdoutWarning 警告
func StdoutWarning(content interface{}, style string) *StdoutHelper {
	receiver := &StdoutHelper{fgColor: COLOR_PURPLE, content: content}
	if style != "" {
		receiver.SetStyle(style)
	}
	return receiver
}

// StdoutWrong 错误
func StdoutWrong(content interface{}, style string) *StdoutHelper {
	receiver := &StdoutHelper{fgColor: COLOR_RED, content: content}
	if style != "" {
		receiver.SetStyle(style)
	}
	return receiver
}
