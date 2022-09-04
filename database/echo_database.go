/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 17:31:17
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 17:31:28
 */
package database

import (
	"mygoredis/interface/resp"
	"mygoredis/lib/logger"
	"mygoredis/resp/reply"
)

type EchoDatabase struct {
}

func NewEchoDatabase() *EchoDatabase {
	return &EchoDatabase{}
}

func (e EchoDatabase) Exec(client resp.Connection, args [][]byte) resp.Reply {
	return reply.MakeMultiBulkReply(args)
}

func (e EchoDatabase) AfterClientClose(c resp.Connection) {
	logger.Info("EchoDatabase AfterClientClose")
}

func (e EchoDatabase) Close() {
	logger.Info("EchoDatabase Close")

}
