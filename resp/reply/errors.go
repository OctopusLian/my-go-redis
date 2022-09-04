/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 16:42:51
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 17:16:18
 */
package reply

// UnknownErrReply represents UnknownErr
type UnknownErrReply struct{}

var unknownErrBytes = []byte("-Err unknown\r\n")

// ToBytes marshals redis.Reply
func (r *UnknownErrReply) ToBytes() []byte {
	return unknownErrBytes
}

func (r *UnknownErrReply) Error() string {
	return "Err unknown"
}

// ArgNumErrReply represents wrong number of arguments for command
type ArgNumErrReply struct {
	Cmd string
}

// ToBytes marshals redis.Reply
func (r *ArgNumErrReply) ToBytes() []byte {
	return []byte("-ERR wrong number of arguments for '" + r.Cmd + "' command\r\n")
}

func (r *ArgNumErrReply) Error() string {
	return "ERR wrong number of arguments for '" + r.Cmd + "' command"
}

// MakeArgNumErrReply represents wrong number of arguments for command
func MakeArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{
		Cmd: cmd,
	}
}

// SyntaxErrReply represents meeting unexpected arguments
//语法错误
type SyntaxErrReply struct{}

var syntaxErrBytes = []byte("-Err syntax error\r\n")
var theSyntaxErrReply = &SyntaxErrReply{}

// MakeSyntaxErrReply creates syntax error
func MakeSyntaxErrReply() *SyntaxErrReply {
	return theSyntaxErrReply
}

// ToBytes marshals redis.Reply
func (r *SyntaxErrReply) ToBytes() []byte {
	return syntaxErrBytes
}

func (r *SyntaxErrReply) Error() string {
	return "Err syntax error"
}

// WrongTypeErrReply represents operation against a key holding the wrong kind of value
// 数据类型错误
type WrongTypeErrReply struct{}

var wrongTypeErrBytes = []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\r\n")

// ToBytes marshals redis.Reply
func (r *WrongTypeErrReply) ToBytes() []byte {
	return wrongTypeErrBytes
}

func (r *WrongTypeErrReply) Error() string {
	return "WRONGTYPE Operation against a key holding the wrong kind of value"
}

// ProtocolErr

// ProtocolErrReply represents meeting unexpected byte during parse requests
// 接口协议错误
type ProtocolErrReply struct {
	Msg string
}

// ToBytes marshals redis.Reply
func (r *ProtocolErrReply) ToBytes() []byte {
	return []byte("-ERR Protocol error: '" + r.Msg + "'\r\n")
}

func (r *ProtocolErrReply) Error() string {
	return "ERR Protocol error: '" + r.Msg
}
