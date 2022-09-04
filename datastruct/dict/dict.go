/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 17:32:45
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 17:35:37
 */
package dict

// Consumer is used to traversal dict, if it returns false the traversal will be break
type Consumer func(key string, val interface{}) bool

// Dict is interface of a key-value data structure
type Dict interface {
	Get(key string) (val interface{}, exists bool) //取出 key 对应的value
	Len() int
	Put(key string, val interface{}) (result int)
	PutIfAbsent(key string, val interface{}) (result int)
	PutIfExists(key string, val interface{}) (result int)
	Remove(key string) (result int)
	ForEach(consumer Consumer) //返回true继续遍历下一个key，返回false停止遍历
	Keys() []string
	RandomKeys(limit int) []string
	RandomDistinctKeys(limit int) []string // 返回不重复的key
	Clear()
}
