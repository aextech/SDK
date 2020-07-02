package main

import (
	"flag"
	"ws-api-cli/lib"
)

//TODO 此SDK仅为参考,发布到生产环境请做好充分测试,祝您用餐愉快 :)
var hosts = flag.String("h", "api.aex.zone", "input host")

var scheme = flag.String("s", "wss", "input scheme")

var path = flag.String("p", "/v3", "input path")

func main() {
	flag.Parse()

	lib.Client(*hosts, *path, *scheme)
}
