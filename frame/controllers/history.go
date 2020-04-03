package controllers

import (
	"net/http"

	"github.com/dctewi/tewi-hwboard/config"
	"github.com/dctewi/tewi-hwboard/core/database"
	"github.com/dctewi/tewi-hwboard/core/session"
	"github.com/dctewi/tewi-hwboard/frame/models"
)

// HistoryController /submit
type HistoryController struct {
}

// Get action
func (c *HistoryController) Get(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)

	userobject := sess.Get("userinfo")
	if userobject == nil {
		ToView(w, "history", models.HistoryModel{
			Model: models.Model{
				Title:    config.App.Title,
				Domain:   config.App.Domain,
				UserInfo: nil,
			},
			Historys: nil,
		})
	} else {
		userinfo := userobject.(database.UserInfo)

		ToView(w, "history", models.HistoryModel{
			Model: models.Model{
				Title:    config.App.Title,
				Domain:   config.App.Domain,
				UserInfo: &userinfo,
			},
			Historys: database.GetUploadLogByEMail(userinfo.Email),
		})
	}
}

// Post action
func (c *HistoryController) Post(w http.ResponseWriter, r *http.Request) {
	c.Get(w, r)
}
