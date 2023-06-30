package commands

import (
	"context"
	"dj-lets-go/database"
	"dj-lets-go/models"
	"dj-lets-go/settings"
	"dj-lets-go/tools"
	"dj-lets-go/types"
	"dj-lets-go/wrongs"
	"fmt"
	"strings"
	"time"

	"github.com/gosuri/uiprogress"
	uuid "github.com/satori/go.uuid"
	"github.com/schollz/progressbar/v3"
)

// TestCommand 测试用
type TestCommand struct{}

// NewTestCommand 构造函数
func NewTestCommand() *TestCommand {
	return &TestCommand{}
}

func (receiver TestCommand) uuid() types.ListString {
	c := make(chan string)
	go func(c chan string) {
		uuidStr := uuid.NewV4().String()
		c <- uuidStr
	}(c)
	go tools.NewTimer(5).Ticker()
	return types.ListString{<-c}
}

func (receiver TestCommand) ls() types.ListString {
	_, res := (&tools.Cmd{}).Process("ls", "-la")
	return types.ListString{res}
}

func (receiver TestCommand) redis() types.ListString {
	ctx := context.Background()

	if err := database.NewRedis(0).Set(ctx, "test", "AAA", 15*time.Minute).Err(); err != nil {
		wrongs.PanicForbidden(err.Error())
	}

	for i := 0; i < 100000; i++ {
		if val, err := database.NewRedis(0).Get(ctx, "test").Result(); err != nil {
			wrongs.PanicForbidden(err.Error())
		} else {
			fmt.Println(i, val)
		}
	}

	return types.ListString{""}
}

func (receiver TestCommand) uiProcesses() types.ListString {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		err := bar.Add(1)
		if err != nil {
			return nil
		}
		time.Sleep(50 * time.Millisecond)
	}

	return types.ListString{}
}

func (receiver TestCommand) uiProcesses2() types.ListString {
	uiprogress.Start()
	bar := uiprogress.AddBar(100).AppendCompleted().PrependElapsed()
	for i := 0; i < 100; i++ {
		bar.Incr()
		time.Sleep(50 * time.Millisecond)
	}
	uiprogress.Stop()

	return types.ListString{}
}

func worker(ctx context.Context, done chan bool) {
	// 模拟一个需要执行5秒的任务
	time.Sleep(5 * time.Second)

	select {
	case <-ctx.Done():
		fmt.Println("Worker stopped due to timeout")
	default:
		fmt.Println("Worker completed successfully")
	}

	done <- true
}

func (receiver TestCommand) timeout() types.ListString {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	done := make(chan bool, 1)
	go worker(ctx, done)

	select {
	case <-done:
		fmt.Println("All workers completed")
	case <-ctx.Done():
		fmt.Println("Program stopped due to timeout")
	}

	return types.ListString{}
}

// sToExcel 设备统计到 Excel
func (receiver TestCommand) sToExcel(paragraphUniqueCode string) types.ListString {

	var (
		items         []types.MapStringToAny
		paragraphName string
		excelWriter   *tools.ExcelWriter
	)

	models.
		NewGormModel().
		SetModel(models.AccountModel{}).
		DB(fmt.Sprintf("fix-%s-rnv", paragraphUniqueCode), nil).
		Joins("join entire_models em on ei.entire_model_unique_code = em.unique_code").
		Joins("join categories c on em.category_unique_code = c.unique_code").
		Joins("left join `lines` l on ei.line_unique_code = l.unique_code").
		Table("entire_instances as ei").
		Select("ei.identity_code", "c.name as cn", "em.name as emn", "ei.serial_number", "ei.factory_name", "ei.factory_device_code", "ei.maintain_workshop_name", "ei.maintain_station_name", "l.name as ln").
		Order("c.unique_code, em.unique_code").
		Where("ei.identity_code like 'S%'").
		Find(&items)

	// 获取段编号对应的段名称
	paragraphName = settings.NewSetting().App.Section("paragraph-unique-codes").Key(strings.ToUpper(paragraphUniqueCode)).String()

	// 生成 Excel 表头
	excelWriter = tools.NewExcelWriter("./static/%s-设备.xlsx", paragraphName).ActiveSheetByIndex(0)
	excelWriter.AddRow(tools.NewExcelRow().SetRowNumber(1).SetCells([]*tools.ExcelCell{
		tools.NewExcelCellAny("唯一编号"),
		tools.NewExcelCellAny("种类"),
		tools.NewExcelCellAny("类型"),
		tools.NewExcelCellAny("所编号"),
		tools.NewExcelCellAny("供应商"),
		tools.NewExcelCellAny("厂编号"),
		tools.NewExcelCellAny("车间"),
		tools.NewExcelCellAny("车站"),
		tools.NewExcelCellAny("线别"),
	}))

	for idx, item := range items {
		rowNumber := idx + 2
		excelWriter.AddRow(tools.NewExcelRow().SetRowNumber(uint64(rowNumber)).SetCells([]*tools.ExcelCell{
			tools.NewExcelCellAny(item["identity_code"]),
			tools.NewExcelCellAny(item["cn"]),
			tools.NewExcelCellAny(item["emn"]),
			tools.NewExcelCellAny(item["serial_number"]),
			tools.NewExcelCellAny(item["factory_name"]),
			tools.NewExcelCellAny(item["factory_device_code"]),
			tools.NewExcelCellAny(item["maintain_workshop_name"]),
			tools.NewExcelCellAny(item["maintain_station_name"]),
			tools.NewExcelCellAny(item["ln"]),
		}))

		tools.StdoutDebug("设备信息", "").EchoLine(fmt.Sprintf("%s %s %s", item["cn"], item["emn"], item["identity_code"]))
	}

	excelWriter.Save()

	return types.ListString{}
}

// Do 执行命令
func (receiver TestCommand) Do(params types.ListString) types.ListString {
	switch params[0] {
	case "uuid":
		return receiver.uuid()
	case "ls":
		return receiver.ls()
	case "redis":
		return receiver.redis()
	case "uiProcesses":
		return receiver.uiProcesses()
	case "uiProcesses2":
		return receiver.uiProcesses2()
	case "timeout":
		return receiver.timeout()
	case "sToExcel":
		return receiver.sToExcel(params[1])
	default:
		return types.ListString{"没有找到命令"}
	}
}
