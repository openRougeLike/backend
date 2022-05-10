package database

import "github.com/openRougeLike/backend/user"

var GUser = user.User{
	ID:           "123",
	Username:     "Idiot",
	XP:           0,
	Lvl:          0,
	Money:        0,
	Base:         user.Upgradable{},
	Hp:           0,
	Mana:         0,
	BaseElements: user.Elements{},
	Statistics:   user.Statistics{},
	State:        user.STATE_NORMAL,
	Map:          user.Map{},
	Inventory:    user.Inventory{},
	Fight:        user.Fight{},
	Token:        "",
}

func FetchUser(token string) (user user.User, ok bool) {
	ok = true
	user = GUser
	return
}
