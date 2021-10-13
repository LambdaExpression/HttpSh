package main

import (
	"encoding/json"
	"flag"
	"github.com/kataras/iris/v12"
	"log"
	"os/exec"
	"strings"
)

var config string
var prot string
var configMap = map[string]string{}

func main() {

	if err := start(); err != nil {
		log.Println(err)
		return
	}

	app := iris.New()

	app.Get("{uri}", run)

	app.Listen(":" + prot)

}

func start() error {
	flag.StringVar(&config, "c", "{}", "--c {\"Access to the address/访问地址 1\":\"Execute script absolute path/执行脚本绝对路径1\",\"Access to the address/访问地址 2\":\"Execute script absolute path/执行脚本绝对路径 2\"}")
	flag.StringVar(&prot, "p", "8088", "--p port/端口")
	flag.Parse()

	return json.Unmarshal([]byte(config), &configMap)
}

// 运行接口
func run(ctx iris.Context) {
	uri := ctx.Params().Get("uri")
	// type 接口返回类型
	// 1 执行shell脚本，返回脚本返回值； 2 执行shell脚本，不返回脚本返回值
	t := ctx.URLParamDefault("type", "1")
	if uri == "favicon.ico" {
		ctx.WriteString("favicon.ico")
		return
	}
	shell := configMap[uri]
	if shell != "" {
		result, err := sh(shell, t)
		if err != nil {
			log.Println(uri, "error", err)
			ctx.WriteString("error " + err.Error())
			return
		}
		log.Println(uri, "run", shell, result)
		ctx.WriteString("success " + result)
		return
	}
	log.Println(uri, "nil")
	ctx.WriteString("nil")
}

// 执行脚本
func sh(shell string, t string) (result string, err error) {
	cmd := exec.Command("/bin/sh", shell)
	switch strings.TrimSpace(t) {
	case "1":
		bytes, err := cmd.Output()
		if err == nil {
			result = string(bytes)
		}
		break
	case "2":
		err = cmd.Run()
	default:
		err = cmd.Run()
	}
	return result, err
}
