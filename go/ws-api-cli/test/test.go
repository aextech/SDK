package main

import (
	"fmt"
	_func "ws-api-cli/func"
	"ws-api-cli/lib"
)

func main() {
	lib.Parse()

	fmt.Println(_func.GetMyBalance())
}
