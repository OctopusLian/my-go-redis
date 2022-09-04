/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 17:23:24
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 17:28:54
 */
package parser

import (
	"io"
	"mygoredis/interface/resp"
)

type Payload struct {
	Data resp.Reply
	Err  error
}

type readState struct {
	readingMultiple   bool
	expectedArgsCount int
	msgType           byte
	args              [][]byte
	bulkLen           int64
}

func (s *readState) finished() bool {
	return s.expectedArgsCount > 0 && len(s.args) == s.expectedArgsCount
}

func ParserStream(reader io.Reader) <-chan *Payload {
	// 异步做协议解析
	ch := make(chan *Payload)
	go parse0(reader, ch)
	return ch
}

func parse0(reader io.Reader, ch chan<- *Payload) {

}
