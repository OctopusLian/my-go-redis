/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 19:19:43
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-05 22:02:06
 */
package database

import (
	"fmt"
	"mygoredis/aof"
	"mygoredis/config"
	"mygoredis/interface/resp"
	"mygoredis/lib/logger"
	"mygoredis/resp/reply"
	"runtime/debug"
	"strconv"
	"strings"
)

type Database struct {
	dbSet      []*DB
	aofHandler *aof.AofHandler
}

// NewDatabase creates a redis database,
func NewDatabase() *Database {
	mdb := &Database{}
	if config.Properties.Databases == 0 {
		config.Properties.Databases = 16
	}
	mdb.dbSet = make([]*DB, config.Properties.Databases)
	for i := range mdb.dbSet {
		singleDB := makeDB()
		singleDB.index = i
		mdb.dbSet[i] = singleDB
	}
	// 初始化aof
	if config.Properties.AppendOnly {
		aofHandler, err := aof.NewAOFHandler(mdb)
		if err != nil {
			panic(err)
		}
		mdb.aofHandler = aofHandler
	}
	return mdb
}

// Exec executes command
// parameter `cmdLine` contains command and its arguments, for example: "set key value"
// 执行Redis命令的逻辑
func (mdb *Database) Exec(c resp.Connection, cmdLine [][]byte) (result resp.Reply) {
	defer func() {
		if err := recover(); err != nil {
			logger.Warn(fmt.Sprintf("error occurs: %v\n%s", err, string(debug.Stack())))
		}
	}()

	cmdName := strings.ToLower(string(cmdLine[0])) // 转小写
	if cmdName == "select" {
		if len(cmdLine) != 2 {
			return reply.MakeArgNumErrReply("select")
		}
		return execSelect(c, mdb, cmdLine[1:])
	}
	// normal commands
	dbIndex := c.GetDBIndex()
	selectedDB := mdb.dbSet[dbIndex]
	return selectedDB.Exec(c, cmdLine)
}

// Close graceful shutdown database
func (mdb *Database) Close() {
}

func (mdb *Database) AfterClientClose(c resp.Connection) {
}

//用户循环DB时执行的逻辑
func execSelect(c resp.Connection, mdb *Database, args [][]byte) resp.Reply {
	dbIndex, err := strconv.Atoi(string(args[0]))
	if err != nil {
		return reply.MakeErrReply("ERR invalid DB index")
	}
	if dbIndex >= len(mdb.dbSet) {
		return reply.MakeErrReply("ERR DB index is out of range")
	}
	c.SelectDB(dbIndex)
	return reply.MakeOkReply()
}
