package table

import (
	"errors"
	"strconv"
)

type TUser struct {
	UserId  int    `json:"userId"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Sex     int    `json:"sex"`
	Address string `json:"address"`
}

func (u TUser) Encode() map[string]interface{} {
	return map[string]interface{}{
		"userId":  u.UserId,
		"name":    u.Name,
		"age":     u.Age,
		"sex":     u.Sex,
		"address": u.Address,
	}
}

func (u *TUser) Decode(value map[string]string) error {
	var (
		err error
		ok  = false
	)
	if u.Name, ok = value["name"]; !ok {
		return errors.New("name is empty")
	}
	if u.Address, ok = value["address"]; !ok {
		return errors.New("address is empty")
	}

	var userId string
	if userId, ok = value["userId"]; !ok {
		return errors.New("address is empty")
	}
	u.UserId, err = strconv.Atoi(userId)
	if err != nil {
		return errors.New("userId error")
	}

	var age string
	if age, ok = value["age"]; !ok {
		return errors.New("age is empty")
	}
	u.Age, err = strconv.Atoi(age)
	if err != nil {
		return errors.New("age error")
	}

	var sex string
	if sex, ok = value["sex"]; !ok {
		return errors.New("sex is empty")
	}
	u.Sex, err = strconv.Atoi(sex)
	if err != nil {
		return errors.New("sex error")
	}
	return nil
}
