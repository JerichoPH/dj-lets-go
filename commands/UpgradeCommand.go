package commands

import (
	"dj-lets-go/database"
	"dj-lets-go/types"
)

// UpgradeCommand 程序升级
type UpgradeCommand struct{}

// NewUpgradeCommand 构造函数
func NewUpgradeCommand() UpgradeCommand {
	return UpgradeCommand{}
}

// 执行数据库管理语句
func (UpgradeCommand) execStatements(sql types.ListString) {
	if len(sql) > 0 {
		for _, s := range sql {
			database.NewGormLauncher().GetConn("").Exec(s)
		}
	}
}

// Do 执行命令
func (receiver UpgradeCommand) Do(params types.ListString) types.ListString {
	switch params[0] {
	}

	return types.ListString{"执行完成"}
}
