package _func

import (
	"io/ioutil"
	"log"
	"strconv"
	"ws-api-cli/types"
)

func GetMyBalance() string {
	p := make(map[string]string)
	p["md5"], p["time"], p["key"] = Sign()
	resp, err := Post("https://"+types.U.Host+"/v3/getMyBalance.php", nil, p, nil)
	if err != nil {
		log.Println("get balance error [Error]:" + err.Error())
		return ""
	}
	if resp.StatusCode != 200 {
		log.Println("get balance StatusCode" + strconv.Itoa(resp.StatusCode))
		return ""
	}

	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
