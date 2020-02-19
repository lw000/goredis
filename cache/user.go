package cache

import (
	"errors"
	"fmt"
	"goredis/dao/table"
	"log"
)

type UserRedisCache struct {
	UserId int
}

func (u *UserRedisCache) key() string {
	return fmt.Sprintf("user:%d", u.UserId)
}

// Load
func (u *UserRedisCache) Load() (table.TUser, error) {
	data := make(map[string]string, 64)
	err := Instance().HScanValues(u.key(), "*", 20, func(key, value string) {
		data[key] = value
	})
	if err != nil {
		log.Println(err)
		return table.TUser{}, err
	}

	if len(data) == 0 {
		return table.TUser{}, errors.New("未查询到数据")
	}

	var reply table.TUser

	return reply, nil
}

// Save
func (u *UserRedisCache) Save(data table.TUser) error {
	s, err := Instance().HMSet(u.key(), data.Encode())
	if err != nil {
		log.Println(err)
		return err
	}
	if s == "OK" {

	}
	return nil
}
