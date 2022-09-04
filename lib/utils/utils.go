/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-09-04 19:06:33
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-04 19:06:35
 */
package utils

// BytesEquals check whether the given bytes is equal
func BytesEquals(a []byte, b []byte) bool {
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	size := len(a)
	for i := 0; i < size; i++ {
		av := a[i]
		bv := b[i]
		if av != bv {
			return false
		}
	}
	return true
}
