/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-03 21:27:08
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-03 22:31:33
 */
package tcp

import (
	"context"
	"mygoredis/interface/tcp"
	"mygoredis/lib/logger"
	"net"
)

// Config stores tcp server properties
type Config struct {
	Address string
}

func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {
	closeChan := make(chan struct{})
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	logger.Info("start listen")
	return ListenAndServe(listener, handler, closeChan)
}

func ListenAndServe(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) error {
	ctx := context.Background()
	for true {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		logger.Info("accepted link")
		go func() {
			handler.Handle(ctx, conn)
		}()
	}

	return nil
}
