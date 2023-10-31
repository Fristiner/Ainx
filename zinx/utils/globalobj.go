package utils

import (
	"encoding/json"
	"os"

	"github.com/peter-matc/Ainx/zinx/ziface"
)

// 定义 存储一切有关zinx框架的全局参数， 供其他模块使用
// 一些参数可以通过zinx.json由用户进行配置

// GlobalObj
// @Description: 全局变量结构体
type GlobalObj struct {
	// Server
	TcpServer ziface.IServer
	// 当前主机监听的ip
	Host string
	// 当前主机监听的端口号
	TcpPort int
	// 服务器名称
	Name string

	// Zinx
	// 版本号
	Version string
	// 当前服务器主机允许的最大链接数
	MaxConn int

	// 当前Zinx框架数据包的最大值
	MaxPackageSize uint32
}

// 定义一个全局的对外对象GlobalObj

var GlobalObject *GlobalObj

// Reload 从zinx.json去加载用户自定义的参数
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	// 将json文件数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}

}

// 提供一个init方法，初始化当前的globalObject对象
func init() {
	// 如果没有加载配置文件 默认的值
	GlobalObject = &GlobalObj{
		Host:           "0.0.0.0",
		TcpPort:        8999,
		Name:           "ZinxServerApp",
		Version:        "V0.4.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// 从conf/zinx.json加载用户自定义的参数
	GlobalObject.Reload()
}