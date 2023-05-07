package main

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"kratos-learn/pkg/bfile"
	"log"
	"time"
)

func main() {
	path := bfile.Join(bfile.SelfDir(), "config.yaml")
	// 初始化配置源
	c := config.New(
		config.WithSource(
			file.NewSource(path),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	// Defines the config JSON Field
	var v struct {
		Service struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"service"`
	}

	// Unmarshal the config to struct
	if err := c.Scan(&v); err != nil {
		panic(err)
	}
	log.Printf("config: %+v", v)

	// Get a value associated with the key
	name, err := c.Value("service.name").String()
	if err != nil {
		panic(err)
	}
	log.Printf("service: %s", name)

	// watch key
	if err := c.Watch("service.name", func(key string, value config.Value) {
		log.Printf("config changed: %s = %v\n", key, value)
	}); err != nil {
		panic(err)
	}

	for {
		version, err := c.Value("service.version").String()
		if err != nil {
			panic(err)
		}
		log.Printf("version: %s", version)
		time.Sleep(time.Second)
	}
}
