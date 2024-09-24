package main

import (
	"fmt"
	// "log"
	"net/url"
)

func main() {
	u, _ := url.Parse("http//localhost:8080?page=1&page=2&a=1&")
	//uのQueryメソッドから得られた返り値queryMapにはクエリパラメーターが入っている
	queryMap := u.Query()

	// パラメーターに対応する値を取り出す
	fmt.Println(queryMap["page"])
	fmt.Println(queryMap["a"])
	fmt.Println(queryMap["b"])
}