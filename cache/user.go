package cache

import (
	"errors"
	"fmt"
	"goredis/dao/table"
)

type UserRedisCache struct {
	UserId int
}

func (u *UserRedisCache) key() string {
	return fmt.Sprintf("user:%d", u.UserId)
}

// Load
func (u *UserRedisCache) Load() (table.TUser, error) {
	data := make(map[string]string, 32)
	err := Instance().HScanValues(u.key(), "*", 32, func(key, value string) {
		data[key] = value
	})
	if err != nil {
		return table.TUser{}, err
	}

	if len(data) == 0 {
		return table.TUser{}, errors.New("未查询到数据")
	}

	var user table.TUser
	if err = user.Decode(data); err != nil {
		return table.TUser{}, err
	}
	return user, nil
}

// Save
func (u *UserRedisCache) Save(user table.TUser) error {
	s, err := Instance().HMSet(u.key(), user.Encode())
	if err != nil {
		return err
	}
	if s == "OK" {

	}
	return nil
}
