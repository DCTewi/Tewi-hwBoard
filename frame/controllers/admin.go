package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/dctewi/tewi-hwboard/config"
	"github.com/dctewi/tewi-hwboard/core/database"
	"github.com/dctewi/tewi-hwboard/core/mail"
	"github.com/dctewi/tewi-hwboard/core/session"
	"github.com/dctewi/tewi-hwboard/core/util"
	"github.com/dctewi/tewi-hwboard/frame/models"

	log "unknwon.dev/clog/v2"
)

// AdminController /submit
type AdminController struct {
}

var adminKeys = make(map[string]time.Time)

// Get action
func (c *AdminController) Get(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)

	// Gen new token
	token := util.GenToken()
	if sess.Get("tokens") == nil {
		sess.Set("tokens", make(map[string]bool))
	}
	tokens := sess.Get("tokens").(map[string]bool)
	tokens[token] = true

	// Logged admin
	if sess.Get("admin") == true {
		msg := ""
		if wantidstr := r.URL.Query().Get("get"); wantidstr != "" { // get query
			if wantid, err := strconv.Atoi(wantidstr); err != nil { // query id
				msg = config.WebConstance["WantIDError"]
			} else if taskinfo := database.GetTaskByID(wantid); taskinfo == nil || taskinfo.ID == 0 { // query taskinfo
				msg = config.WebConstance["QueryIllegal"]
			} else if taskinfo.End.After(time.Now()) { // query available
				msg = config.WebConstance["TaskNotEnd"]
			} else { // query ok
				querypath := config.Path.UploadFolder + "/" + taskinfo.Subject + "-" + taskinfo.End.Format("060102")
				// zip folder
				err := util.ZipDir(querypath)
				if err != nil {
					log.Error("Zip error : " + err.Error())
					msg = config.WebConstance["NoAvailableFile"]
				} else {
					// Set header to zip
					w.Header().Set("Content-Type", "application/zip")
					w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", querypath+".zip"))

					// Get zip
					queryfile, err := os.Open(querypath + ".zip")
					if err != nil {
						log.Error("Open file error: " + querypath + ".zip")
					} else {
						// Trans
						defer queryfile.Close()

						io.Copy(w, queryfile)

						msg = config.WebConstance["QueryFileSuccess"]
					}
				}
				if msg == "" {
					msg = config.WebConstance["QueryFileUnknownError"]
				}
			}
		}

		ToView(w, "admin-logged", models.Model{
			Token:   token,
			Message: msg,
			Title:   config.App.Title,
			Domain:  config.App.Domain,
		})
		adminKeys = make(map[string]time.Time)

	} else if akey := r.URL.Query().Get("akey"); akey != "" { // Login attempt with admin key
		t := time.Now()
		expir, _ := time.ParseDuration("10m")
		if gentime, ok := adminKeys[akey]; ok { // key exists
			if t.Before(gentime.Add(expir)) { // key available
				log.Info("Admin logged at session: " + sess.SessionID() + " on IP:" + r.RemoteAddr)
				sess.Set("admin", true)
				ToView(w, "admin-logged", models.Model{
					Token:  token,
					Title:  config.App.Title,
					Domain: config.App.Domain,
				})
				delete(adminKeys, akey)
				return
			}

			log.Warn("Admin logged attempt with akey which time out on IP:" + r.RemoteAddr)
			delete(adminKeys, akey)
		}
		http.Redirect(w, r, "/", 303)

	} else { // New login attempt
		log.Info("Admin login attempt on IP:" + r.RemoteAddr)
		ToView(w, "admin", models.Model{
			Token:  token,
			Title:  config.App.Title,
			Domain: config.App.Domain,
		})
	}
}

// Post action
func (c *AdminController) Post(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)
	r.ParseForm()

	// Check
	token := r.Form.Get("token")
	tokens := sess.Get("tokens")
	if token != "" && tokens != nil { // token exists
		if _, ok := tokens.(map[string]bool)[token]; ok { // token available
			if sess.Get("admin") == true { // logged admin want to add task
				nsub := template.HTMLEscapeString(r.Form.Get("sub"))
				nttl := template.HTMLEscapeString(r.Form.Get("ttl"))
				nfmt := template.HTMLEscapeString(r.Form.Get("fmt"))
				ndat := template.HTMLEscapeString(r.Form.Get("dat"))

				usertimezone, err := time.LoadLocation(config.App.UserTimeZone)
				if err != nil {
					log.Fatal("APP CONFIG ERROR with config.App.UserTimeZone, error: " + err.Error())
					return
				}
				newsubmitdate, err := time.ParseInLocation("2006-01-02", ndat, usertimezone)
				if err != nil {
					log.Warn("Parse time error with form: " + fmt.Sprint(r.Form))
					return
				}
				ndatestr := newsubmitdate.Format("2006-01-02")
				loc, err := time.LoadLocation(config.App.UserTimeZone)
				if err != nil {
					log.Error("Parse TIMEZONE error: " + err.Error())
				}

				nsta, _ := time.ParseInLocation("2006-01-02 15:04", ndatestr+" 00:00", loc)
				nend, _ := time.ParseInLocation("2006-01-02 15:04", ndatestr+" 23:59", loc)

				newtaskinfo := database.TaskInfo{
					Subject:  nsub,
					SubTitle: nttl,
					FileType: nfmt,
					Start:    nsta,
					End:      nend,
				}
				database.Insert(newtaskinfo)
				log.Info("New task inserted on IP:" + r.RemoteAddr)
				http.Redirect(w, r, "/", 303)
			} else {
				if email := r.Form.Get("qml"); email != "" { // form available
					if checkIsAdmin(email) { // account available
						// Gen new admin key
						newadminkey := util.GenRandomString(24)
						adminKeys[newadminkey] = time.Now()

						err := mail.Send(mail.NewAdminMail([]string{email}, newadminkey))
						if err == nil {
							log.Info("Gened new admin login link for " + email + " with akey:" + newadminkey)
						} else {
							log.Error("Send mail error: " + err.Error())
						}
					} else {
						log.Error("Admin login failed with wrong email: " + email)
					}
				} else {
					log.Error("Admin login failed with wrong form")
				}
			}
			// Token used
			delete(tokens.(map[string]bool), token)
		} else {
			log.Error("Admin login failed with wrong token")
		}
	}
}

func checkIsAdmin(email string) bool {
	for _, i := range config.App.AdminEmails {
		if email == i {
			return true
		}
	}
	return false
}
