/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 18:49:17
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 18:49:21
 */
package database

import (
	"mygoredis/interface/resp"
	"mygoredis/resp/reply"
)

// Ping the server
func Ping(db *DB, args [][]byte) resp.Reply {
	if len(args) == 0 {
		return &reply.PongReply{}
	} else if len(args) == 1 {
		return reply.MakeStatusReply(string(args[0]))
	} else {
		return reply.MakeErrReply("ERR wrong number of arguments for 'ping' command")
	}
}

func init() {
	RegisterCommand("ping", Ping, -1)
}
