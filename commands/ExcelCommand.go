package commands

import (
	"dj-lets-go/tools"
	"dj-lets-go/types"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/go-gota/gota/dataframe"
)

type ExcelCommand struct{}

// NewExcelCommand 构造函数
func NewExcelCommand() *ExcelCommand {
	return &ExcelCommand{}
}

// json转excel
func (receiver ExcelCommand) fromJson(args ...string) types.ListString {
	var (
		argJsonFilename, argExcelFilename string
		excelWriter                       *tools.ExcelWriter
		rowNumber                         uint64
		excelRows                         []*tools.ExcelRow
		err                               error
		f                                 []byte
		jsonData                          []types.MapStringToAny
		excelTitle                        []*tools.ExcelCell
		title                             types.ListString
	)
	argJsonFilename = args[0]
	argExcelFilename = args[1]
	rowNumber = 1
	excelRows = make([]*tools.ExcelRow, 0)
	excelTitle = make([]*tools.ExcelCell, 0)
	title = make(types.ListString, 0)

	if f, err = ioutil.ReadFile(fmt.Sprintf("./static/%s", argJsonFilename)); err != nil {
		tools.StdoutWrong("读取json文件失败", "").EchoLine(err.Error())
		return nil
	} else {
		if err = json.Unmarshal(f, &jsonData); err != nil {
			tools.StdoutWrong("反序列化json文件失败", "").EchoLine(err.Error())
			return nil
		}
	}

	if len(jsonData) > 0 {
		// 制作表头
		for key := range jsonData[0] {
			excelTitle = append(excelTitle, tools.NewExcelCellAny(key))
			title = append(title, key)
		}
		excelRows = append(excelRows, tools.NewExcelRow().SetRowNumber(rowNumber).SetCells(excelTitle))

		// 填充内容
		for _, jsonDatum := range jsonData {
			rowNumber += 1
			cells := make([]*tools.ExcelCell, 0)

			for _, item := range title {
				cells = append(cells, tools.NewExcelCellAny(jsonDatum[item]))
			}
			excelRows = append(excelRows, tools.NewExcelRow().SetRowNumber(rowNumber).SetCells(cells))
		}

		// 保存文件
		excelWriter = tools.NewExcelWriter(fmt.Sprintf("./static/%s", argExcelFilename)).ActiveSheetByIndex(0)
		if err = excelWriter.SetRows(excelRows).Save(); err != nil {
			tools.StdoutWrong("保存excel文件错误：%s", err.Error())
			return nil
		}
	}

	return types.ListString{fmt.Sprintf("保存excel文件成功：./static/%s", argExcelFilename)}
}

// excel转json
func (receiver ExcelCommand) toJson(args ...string) types.ListString {
	var (
		argExcelFilename string
		argJsonFilename  string
		excelData        dataframe.DataFrame
		jsonData         []types.MapStringToAny
	)
	argExcelFilename = args[0]
	argJsonFilename = args[1]
	jsonData = make([]types.MapStringToAny, 0)

	excelData = tools.NewExcelReader().AutoRead("./static/%s", argExcelFilename).ToDataFrameDetectType()

	for _, excelDatum := range excelData.Maps() {
		jsonData = append(jsonData, excelDatum)
	}

	if marshal, err := json.Marshal(excelData.Maps()); err != nil {
		tools.StdoutWrong("json序列化失败：%s", err.Error())
		return nil
	} else {
		tools.NewFileSystem(fmt.Sprintf("./static/%s", argJsonFilename)).WriteString(string(marshal))
		return types.ListString{fmt.Sprintf("保存json文件成功：./static/%s", argJsonFilename)}
	}
}

// Do 执行命令
func (receiver ExcelCommand) Do(params types.ListString) types.ListString {
	switch params[0] {
	default:
		return types.ListString{"没有找到命令"}
	case "fromJson":
		return receiver.fromJson(params[1:]...)
	case "toJson":
		return receiver.toJson(params[1:]...)
	}
}
