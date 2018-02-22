# 百行代码实现的ssh客户端和服务端

## 编译

`go get github.com/51reboot/golang-01-homework/lesson2/binggan/gossh`

**提示**:需要翻墙


## 使用

服务端:

`gossh -addr=:8080 -s`

客户端:
`gossh -addr=:8080`

或者:
`nc 127.0.0.1 8080`

**注意**:nc作为客户端top,vi之类的命令不能用


## 致有志于hack的同学

- 增加只读功能
- 增加替换shell程序的功能
