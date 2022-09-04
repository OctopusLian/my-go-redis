/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 18:19:51
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 18:22:22
 */
package database

import "strings"

// 记录该系统里所有的指令和 command的关系
var cmdTable = make(map[string]*command)

type command struct {
	executor ExecFunc
	arity    int // allow number of args, arity < 0 means len(args) >= -arity
}

// RegisterCommand registers a new command
// arity means allowed number of cmdArgs, arity < 0 means len(args) >= -arity.
// for example: the arity of `get` is 2, `mget` is -2
func RegisterCommand(name string, executor ExecFunc, arity int) {
	name = strings.ToLower(name) //统一转换为小写
	cmdTable[name] = &command{
		executor: executor,
		arity:    arity,
	}
}
