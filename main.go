package main

import (
	"goredis/cache"
	"goredis/config"
	"goredis/dao/table"
	"log"
)

func main() {
	conf, err := config.LoadJSONConfig("conf/conf.json")
	if err != nil {
		log.Fatal(err)
	}
	err = cache.OpenRedis(&conf.RdsCfg)
	if err != nil {
		log.Fatal(err)
	}

	var s string
	s, err = cache.Instance().Set("name", "liwei", 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("s", s)

	log.Println(cache.Instance().Get("name"))

	u := cache.UserRedisCache{UserId: 10000}
	err = u.Save(table.TUser{UserId: 10000, Name: "levi", Age: 0, Sex: 0, Address: ""})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(u.Load())
}
