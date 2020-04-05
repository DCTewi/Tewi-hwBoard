package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/dctewi/tewi-hwboard/config"
	"github.com/dctewi/tewi-hwboard/core/database"
	"github.com/dctewi/tewi-hwboard/core/mail"
	"github.com/dctewi/tewi-hwboard/core/session"
	"github.com/dctewi/tewi-hwboard/core/util"
	"github.com/dctewi/tewi-hwboard/frame/models"

	log "unknwon.dev/clog/v2"
)

// LoginController "/login"
type LoginController struct {
}

// {qq: ckey}
var ckeymap = make(map[string]string)

// Get action
func (c *LoginController) Get(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)

	// Logout action
	if r.URL.Query().Get("logout") == "true" {
		q := r.URL.Query()
		q.Set("logout", "")
		r.URL.RawQuery = q.Encode()
		session.GlobalSessions.SessionDestroy(w, r)
		http.Redirect(w, r, "/", 303)
	}

	// Logged, don't need login
	if info := sess.Get("userinfo"); info != nil {
		http.Redirect(w, r, "/", 303)
		return
	}

	// Gen new token
	token := util.GenToken()
	if sess.Get("tokens") == nil {
		sess.Set("tokens", make(map[string]bool))
	}
	tokens := sess.Get("tokens").(map[string]bool)
	tokens[token] = true

	mod := models.Model{
		Token:  token,
		Title:  config.App.Title,
		Domain: config.App.Domain,
	}

	switch w := r.URL.Query().Get("error"); w {
	case "":
		break
	case "multireg":
		mod.Message = config.WebConstance["MultiReg"]
	case "mailerr":
		mod.Message = config.WebConstance["MailError"]
	default:
		mod.Message = config.WebConstance["UnknownError"]
	}

	ToView(w, "login", mod)
	//http.Redirect(w, r, "/", 303)
}

// Post action
func (c *LoginController) Post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	sess := session.GlobalSessions.SessionStart(w, r)

	// Query for ckey
	if r.URL.Query().Get("q") == "ckey" {
		ckey := util.GenCKey()
		// Get userq
		userq := template.HTMLEscapeString(r.Form.Get("userq"))
		if okqq := util.CheckUserQ(userq); okqq {
			// Try to send
			if err := mail.Send(mail.NewMail([]string{userq + "@qq.com"}, ckey)); err == nil {
				ckeymap[userq] = ckey

				log.Info("Gened ckey for " + userq + " as " + ckey + " on session " + sess.SessionID() + " IP: " + r.RemoteAddr)
			} else {
				fmt.Fprintf(w, config.WebConstance["MailError"])
				log.Error("Mail send failed with error: " + err.Error())
			}
		}
		return
	}

	// Try to login
	token := r.Form.Get("token")
	tokens := sess.Get("tokens")
	// Check token
	if token != "" && tokens != nil {
		if _, ok := tokens.(map[string]bool)[token]; ok {
			userq := template.HTMLEscapeString(r.Form.Get("userq"))
			stuid := template.HTMLEscapeString(r.Form.Get("stuid"))
			uckey := template.HTMLEscapeString(r.Form.Get("cikey"))

			log.Info("Login attempt with q:" + userq + " stid:" + stuid + " uckey:" + uckey + " IP:" + r.RemoteAddr)
			// Check ckey
			if ckey, ok := ckeymap[userq]; ok && ckey == uckey {
				// Check form
				if okqq, okid := util.CheckUserQ(userq), util.CheckStuID(stuid); okqq && okid {
					// Check database
					infoByStuID := database.GetUserInfoBySID(stuid)
					infoByEmail := database.GetUserInfoByEmail(userq)

					if infoByEmail == infoByStuID {
						if infoByStuID.Email == "" { // New login
							info := database.UserInfo{}
							info.Email = userq
							info.StudentID = stuid
							database.Insert(info)
							infoByStuID, infoByEmail = info, info
							log.Info("Login with register success: userq:" + userq + " stuid:" + stuid)
						}
						// Correct login
						sess.Set("userinfo", infoByStuID)
						http.Redirect(w, r, "/", 303)
						log.Info("Login success: userq:" + userq + " stuid:" + stuid)

					} else { // Wrong login
						q := r.URL.Query()
						q.Set("error", "multireg")
						r.URL.RawQuery = q.Encode()
						http.Redirect(w, r, "/login?error=multireg", 303)
						log.Warn("Login failed with userq:" + userq + " stuid:" + stuid + "(Multi Reg)")
					}

				}
			}
		}
		delete(tokens.(map[string]bool), token)
	}
}
