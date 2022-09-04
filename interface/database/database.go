/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 16:28:06
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 16:28:16
 */
package database

import "mygoredis/interface/resp"

// CmdLine is alias for [][]byte, represents a command line
type CmdLine = [][]byte

// Database is the interface for redis style storage engine
type Database interface {
	Exec(client resp.Connection, args [][]byte) resp.Reply
	AfterClientClose(c resp.Connection)
	Close()
}

// DataEntity stores data bound to a key, including a string, list, hash, set and so on
type DataEntity struct {
	Data interface{}
}
