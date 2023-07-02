package tools

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	
	"dj-lets-go/constants"
	"dj-lets-go/wrongs"
	
	"github.com/gin-gonic/gin"
)

// GetAuthorization 获取登陆信息
func GetAuthorization(ctx *gin.Context) any {
	authorization, exist := ctx.Get(constants.AccountAuthorizationFieldName)
	if !exist {
		wrongs.PanicUnLogin("")
	}
	
	return authorization
}

func GetRootPath() string {
	rootPath, _ := filepath.Abs(".")
	return rootPath
}

func GetStaticPath() string {
	return GetRootPath() + "/static"
}

// GetCurrentPath 最终方案-全兼容
func GetCurrentPath() string {
	dir := getGoBuildPath()
	if strings.Contains(dir, getTmpDir()) {
		return getGoRunPath()
	}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getGoBuildPath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getGoRunPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// ToJson json序列化
func ToJson(v interface{}) string {
	jsonBytes, _ := json.Marshal(v)
	return string(jsonBytes)
}

// ToJsonFormat json序列化 格式化
func ToJsonFormat(v interface{}) string {
	jsonBytes, _ := json.MarshalIndent(v, "", "  ")
	return string(jsonBytes)
}

// DeepCopyByGob 深拷贝
func DeepCopyByGob(dst, src interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	
	return gob.NewDecoder(&buffer).Decode(dst)
}

// JoinWithoutEmpty 去掉空值然后合并
func JoinWithoutEmpty(values []string, sep string) string {
	newValues := make([]string, 0)
	
	for _, value := range values {
		if value != "" {
			newValues = append(newValues, value)
		}
	}
	
	return strings.Join(newValues, sep)
}

// AddPrefix 给字符串增加前缀，否则返回默认值
func AddPrefix(value, prefix, defaultValue string) string {
	return Operation{}.Ternary(value != "", fmt.Sprintf("%s%s", prefix, value), defaultValue).(string)
}
