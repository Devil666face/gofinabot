package models

import (
	"errors"

	. "github.com/Devil666face/gofinabot/config"
	. "github.com/Devil666face/gofinabot/store/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TGID              uint   `gorm:"uniqueIndex;not null;index"`
	Username          string `gorm:"not null;default:noname;index"`
	IsAdmin           bool   `gorm:"default:false"`
	IsAllow           bool   `gorm:"default:false"`
	MoneyTransactions []MoneyTransaction
}

func (user User) String() string {
	return user.Username
}

func (user *User) Get(id uint) error {
	if err := Db.First(user, id); err != nil {
		return err.Error
	}
	return nil
}

func (user *User) GetUserByTgID(id int64) error {
	err := Db.Where("tg_id = ?", id).Take(user)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return err.Error
	}
	return nil
}

func (user *User) Create() error {
	if err := Db.Save(user); err != nil {
		return err.Error
	}
	return nil
}

func (user *User) Update() error {
	if err := Db.Save(user); err != nil {
		return err.Error
	}
	return nil
}

func getUsersForSelect(query string) ([]User, error) {
	var users = []User{}
	err := Db.Where(query, true).Find(&users)
	if err != nil {
		users = append(users, User{TGID: Cfg.SuperuserId})
		return users, err.Error
	}
	return users, nil
}

func GetAllAdmins() ([]User, error) {
	return getUsersForSelect("is_admin = ?")
}

func GetAllAllows() ([]User, error) {
	return getUsersForSelect("is_allow = ?")
}

func GetChatIdsForSelector(selector func() ([]User, error)) ([]int64, error) {
	var chats = []int64{}
	users, err := selector()
	if err != nil {
		return chats, err
	}
	for _, u := range users {
		chats = append(chats, int64(u.TGID))
	}
	return chats, nil
}
