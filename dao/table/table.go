package table

type TUser struct {
	UserId  int32  `json:"userId"`
	Name    string `json:"name"`
	Age     int32  `json:"age"`
	Sex     int32  `json:"sex"`
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

func (u TUser) Decode(data map[string]string) error {
	return nil
}
