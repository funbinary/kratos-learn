package main

import (
	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/hashicorp/consul/api"
	"log"
)

func main() {
	consulClient, err := api.NewClient(&api.Config{
		Address: "127.0.0.1:8500",
	})
	if err != nil {
		panic(err)
	}
	cs, err := consul.New(consulClient, consul.WithPath("http"))
	// consul中需要标注文件后缀，kratos读取配置需要适配文件后缀
	// The file suffix needs to be marked, and kratos needs to adapt the file suffix to read the configuration.
	if err != nil {
		panic(err)
	}
	c := config.New(config.WithSource(cs))
	defer c.Close()

	// load sources before get
	if err := c.Load(); err != nil {
		log.Fatalln(err)
	}

	// acquire config value
	foo, err := c.Value("http.server.port").String()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(foo)
}
