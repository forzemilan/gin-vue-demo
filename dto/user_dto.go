package dto

import (
	"ginessential/model"
)

// UserDto 用结构体定义返回字段
type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// ToUserDto 过滤敏感信息，仅返回普通信息
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
