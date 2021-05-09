package dto

import "first_go/model"

type UserDto struct {
	Name      string `string:"name"`
	Telephone string `string:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
