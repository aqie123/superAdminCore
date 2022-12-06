package main

import (
	"fmt"
	"github.com/aqie123/superAdminCore/cache"
	"github.com/aqie123/superAdminCore/core"
	"github.com/aqie123/superAdminCore/test/conf"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		panic("配置文件加载失败")
	}
}

func main() {

	c := core.NewCore()

	//加载配置(这是第一步)
	c.LoadConf(conf.Conf)

	err := cache.Cache().Put("ppp", "789789", 0)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(cache.Cache().Get("ppp"))

}
