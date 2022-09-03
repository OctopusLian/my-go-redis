/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-03 22:55:12
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-03 23:13:55
 */
package tcp

import (
	"bufio"
	"context"
	"io"
	"mygoredis/lib/logger"
	"mygoredis/lib/sync/atomic"
	"mygoredis/lib/sync/wait"
	"net"
	"sync"
	"time"
)

type EchoHandler struct {
	activeConn sync.Map //保存，记录有多少连接数
	closing    atomic.Boolean
}

// MakeEchoHandler creates EchoHandler
func MakeHandler() *EchoHandler {
	return &EchoHandler{}
}

// EchoClient is client for EchoHandler, using for test
type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

// Close close connection
func (c *EchoClient) Close() error {
	// 把 socket关闭
	c.Waiting.WaitWithTimeout(10 * time.Second)
	c.Conn.Close()
	return nil
}

// Handle echos received line to client
func (h *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if h.closing.Get() {
		// 已经是关闭状态
		// closing handler refuse new connection
		_ = conn.Close()
	}
	// 初始化
	client := &EchoClient{
		Conn: conn,
	}
	// 存一个空结构体：不占任何空间
	h.activeConn.Store(client, struct{}{}) // 把一个客户端存进去

	reader := bufio.NewReader(conn)
	for {
		// may occurs: client EOF, client timeout, server early close
		msg, err := reader.ReadString('\n')
		if err != nil { //接收数据出现问题
			if err == io.EOF {
				// 客户端退出
				logger.Info("connection close")
				h.activeConn.Delete(client)
			} else {
				// 警告
				logger.Warn(err)
			}
			return
		}
		client.Waiting.Add(1) // +1
		b := []byte(msg)
		_, _ = conn.Write(b)
		client.Waiting.Done() // -1
	}
}

// Close stops echo handler
func (h *EchoHandler) Close() error {
	logger.Info("handler shutting down...")
	h.closing.Set(true)
	//关闭所有客户端
	h.activeConn.Range(func(key interface{}, val interface{}) bool {
		client := key.(*EchoClient)
		_ = client.Close()
		return true // 重要！
	})
	return nil
}
