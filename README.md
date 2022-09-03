<!--
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-09-03 21:19:18
 * @LastEditors: neozhang
 * @LastEditTime: 2022-09-03 23:31:06
-->
# Go实现Redis中间件  

## Redis网络协议  

### RESP  

#### 正常回复  

以“+” 开头，以“\r\n”结尾的字符串形式，例如 `+OK\r\n`  

#### 错误回复  

以“-” 开头，以“\r\n”结尾的字符串形式，例如 `-Error message\r\n`  

#### 整数  

以“:”开头，以“\r\n”结尾的字符串形式，例如 `:123456\r\n`  

#### 多行字符串  

以“$”开头，后跟实际发送的字节数，以“\r\n”结尾。  
例如 `$9\r\nimooc.com\r\n` 等于 `"imooc.com"`  
例如 `$0\r\n\r\n` 等于 ""

#### 数组  

以“*”开头，后跟成员个数  