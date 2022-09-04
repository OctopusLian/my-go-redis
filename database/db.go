/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 18:17:03
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 18:18:15
 */
package database

import (
	"mygoredis/datastruct/dict"
	"mygoredis/interface/resp"
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
