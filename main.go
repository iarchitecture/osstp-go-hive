package main

import "osstp-go-hive/routers"

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//或 		  			GOPROXY='https://proxy.golang.com.cn,direct'
//go:generate go mod tidy
//go:generate go mod download

// @Summary	创建文章

func main() {
	r := routers.InitRouter()
	r.Run()
}
