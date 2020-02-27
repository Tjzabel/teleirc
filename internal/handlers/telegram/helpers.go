package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

/*
GetUsername returns the username of a user. If no username is set, their
first name is returned.
*/
func GetUsername(u *tgbotapi.User) string {
	if u.UserName == "" {
		return u.FirstName
	} else {
		return u.FirstName + " (@" + u.UserName + ")"
	}
}
