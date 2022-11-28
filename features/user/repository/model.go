package repository

import (
	"main.go/features/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string
	HP       string
	Password string
}

func FromDomain(du domain.Core) User {
	return User{
		Model:    gorm.Model{ID: du.ID},
		Nama:     du.Nama,
		HP:       du.HP,
		Password: du.Password,
	}
}

func ToDomain(u User) domain.Core {
	return domain.Core{

		ID:       u.ID,
		Nama:     u.Nama,
		HP:       u.HP,
		Password: u.Password,
	}
}

func ToDomainArray(au []User) []domain.Core {
	var res []domain.Core
	for _, val := range au {
		res = append(res, domain.Core{ID: val.ID, Nama: val.Nama, HP: val.HP, Password: val.Password})
	}

	return res
}
