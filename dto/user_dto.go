package dto

import "brianwchen/ginessential/model"

type UserDto struct {
	Name      string      `json:"name"`
	Telephone string      `json:"telephone"`
	Group     model.Group `json:"group"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
		Group:     user.Group,
	}
}
