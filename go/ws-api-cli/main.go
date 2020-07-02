package main

import (
	"flag"
	"ws-api-cli/lib"
)

var hosts = flag.String("h", "api.aex.zone", "input host")

var scheme = flag.String("s", "wss", "input scheme")

var path = flag.String("p", "/v3", "input path")

func main() {
	flag.Parse()

	lib.Client(*hosts, *path, *scheme)
}
