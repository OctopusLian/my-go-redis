/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-03 21:27:08
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-03 22:53:16
 */
package tcp

import (
	"context"
	"mygoredis/interface/tcp"
	"mygoredis/lib/logger"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Config stores tcp server properties
type Config struct {
	Address string
}

func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {
	closeChan := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigChan
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeChan <- struct{}{}
		}
	}()
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	logger.Info("start listen")
	return ListenAndServe(listener, handler, closeChan)
}

func ListenAndServe(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) error {
	go func() {
		<-closeChan
		logger.Info("shutting down")
		listener.Close()
		handler.Close()
	}()

	defer func() {
		listener.Close()
		handler.Close()
	}()
	ctx := context.Background()
	var waitDone sync.WaitGroup
	for true {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		logger.Info("accepted link")
		waitDone.Add(1) // 开始服务，+1
		go func() {
			defer func() {
				waitDone.Done() //服务完，-1
			}()
			handler.Handle(ctx, conn)
		}()
	}
	waitDone.Wait()

	return nil
}
