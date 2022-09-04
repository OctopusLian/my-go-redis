/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 16:25:24
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 16:32:26
 */
package resp

// Connection represents a connection with redis client
type Connection interface {
	Write([]byte) error //回复消息
	// used for multi database
	GetDBIndex() int
	SelectDB(int)
}
