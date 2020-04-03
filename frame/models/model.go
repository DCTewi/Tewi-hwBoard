package models

import (
	"github.com/dctewi/tewi-hwboard/core/database"
)

// Model of base
type Model struct {
	Token    string
	Title    string
	Domain   string
	Message  string
	UserInfo *database.UserInfo
}
