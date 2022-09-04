/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 16:26:29
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 16:33:24
 */
package resp

// Reply is the interface of redis serialization protocol message
type Reply interface {
	ToBytes() []byte //将回复的内容转成字节数组
}
