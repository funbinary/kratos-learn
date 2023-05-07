package main

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"kratos-learn/pkg/bfile"
	"log"
	"os"
)

func main() {
	os.Setenv("KRATOS_PORT", "9001")
	path := bfile.Join(bfile.SelfDir(), "config.yaml")
	// 初始化配置源
	c := config.New(
		config.WithSource(
			env.NewSource("KRATOS_"),
			file.NewSource(path),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	// Get a value associated with the key
	port, err := c.Value("PORT").String()
	if err != nil {
		panic(err)
	}
	log.Printf("port: %s", port)

	name, err := c.Value("http.server.name").String()
	if err != nil {
		panic(err)
	}
	log.Printf("http server name: %s", name)

	hport, err := c.Value("http.server.port").String()
	if err != nil {
		panic(err)
	}
	log.Printf("http server port: %s", hport)

}
