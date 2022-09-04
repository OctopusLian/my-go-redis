/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 18:17:03
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 18:44:34
 */
package database

import (
	"mygoredis/datastruct/dict"
	"mygoredis/interface/resp"
	"mygoredis/resp/reply"
	"strings"
)

type DB struct {
	index int
	data  dict.Dict
}

// ExecFunc is interface for command executor
// args don't include cmd line
type ExecFunc func(db *DB, args [][]byte) resp.Reply

// CmdLine is alias for [][]byte, represents a command line
type CmdLine = [][]byte

func makeDB() *DB {
	db := &DB{
		data: dict.MakeSyncDict(),
	}
	return db
}

// Exec executes command within one database
func (db *DB) Exec(c resp.Connection, cmdLine [][]byte) resp.Reply {
	cmdName := strings.ToLower(string(cmdLine[0])) //统一转换为小写
	cmd, ok := cmdTable[cmdName]
	if !ok {
		return reply.MakeErrReply("ERR unknown command '" + cmdName + "'")
	}
	if !validateArity(cmd.arity, cmdLine) {
		return reply.MakeArgNumErrReply(cmdName)
	}
	fun := cmd.executor
	return fun(db, cmdLine[1:])
}

func validateArity(arity int, cmdArgs [][]byte) bool {
	argNum := len(cmdArgs)
	if arity >= 0 {
		return argNum == arity
	}
	return argNum >= -arity
}

/* ---- data Access ----- */
