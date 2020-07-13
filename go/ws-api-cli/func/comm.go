package _func

import (
	"crypto/md5"
	"fmt"
	"github.com/wangxudong123/assist"
	"io"
	"log"
	"time"
	"ws-api-cli/types"
)

func MD5(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

//return md5 timeUnixString key
func Sign() (md5, key, timeUnixString string) {
	if types.Skey == "" {
		log.Fatal("skey is not allowed to be empty")
	}

	if types.Key == "" {
		log.Fatal("key is not allowed to be empty")
	}

	if types.Id == "" {
		log.Fatal("id is not allowed to be empty")
	}

	timeUnixString = assist.Int64ToString(time.Now().Unix())
	s := types.Key + "_" + types.Id + "_" + types.Skey + "_" + timeUnixString
	md5 = MD5(s)
	key = types.Key
	return
}
