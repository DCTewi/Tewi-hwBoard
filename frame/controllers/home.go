package controllers

import (
	"net/http"
	"sort"

	"github.com/dctewi/tewi-hwboard/config"
	"github.com/dctewi/tewi-hwboard/core/database"
	"github.com/dctewi/tewi-hwboard/core/session"
	"github.com/dctewi/tewi-hwboard/core/util"
	"github.com/dctewi/tewi-hwboard/frame/models"
)

// HomeController "/"
type HomeController struct {
}

// Get action
func (c *HomeController) Get(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)

	// Gen new token
	token := util.GenToken()
	if sess.Get("tokens") == nil {
		sess.Set("tokens", make(map[string]bool))
	}
	tokens := sess.Get("tokens").(map[string]bool)
	tokens[token] = true

	// Logged
	if info := sess.Get("userinfo"); info != nil {
		uinfo := info.(database.UserInfo)
		tasklist := database.GetAll("taskinfo")

		var tasks database.TaskInfoSlice
		for i := tasklist.Front(); i != nil; i = i.Next() {
			tasks = append(tasks, i.Value.(database.TaskInfo))
		}
		sort.Sort(&tasks)

		ToView(w, "home", models.HomeModel{
			Model: models.Model{
				Token:    token,
				Title:    config.App.Title,
				Domain:   config.App.Domain,
				UserInfo: &uinfo,
			},
			Tasks: tasks,
		})
	} else { // Not logged
		ToView(w, "home", models.HomeModel{
			Model: models.Model{
				Token:    token,
				Title:    config.App.Title,
				Domain:   config.App.Domain,
				UserInfo: nil,
			},
			Tasks: nil,
		})
	}
}

// Post action
func (c *HomeController) Post(w http.ResponseWriter, r *http.Request) {
	c.Get(w, r)
}
