package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/dctewi/tewi-hwboard/core/database"
	"github.com/dctewi/tewi-hwboard/core/session"
	"github.com/dctewi/tewi-hwboard/core/util"

	log "unknwon.dev/clog/v2"
)

// SubmitController /submit
type SubmitController struct {
}

// Get action
func (c *SubmitController) Get(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 301)
}

// Post action
func (c *SubmitController) Post(w http.ResponseWriter, r *http.Request) {
	sess := session.GlobalSessions.SessionStart(w, r)

	userobject := sess.Get("userinfo")
	if userobject != nil {
		userinfo := userobject.(database.UserInfo)
		r.ParseMultipartForm(5 << 10)

		// Check
		token := r.Form.Get("token")
		tokens := sess.Get("tokens")
		if token != "" && tokens != nil { // token exists
			if _, ok := tokens.(map[string]bool)[token]; ok { // token available
				submmitto := template.HTMLEscapeString(r.Form.Get("submitto"))
				submmittoint, err := strconv.Atoi(submmitto)
				if err != nil {
					log.Error("File submit with Atoi err: " + err.Error())
					return
				}

				taskinfo := database.GetTaskByID(submmittoint)
				if taskinfo == nil {
					log.Error("File submit with task not exists: submitto:" + submmitto)
					return
				}

				file, handler, err := r.FormFile("file")
				if err != nil {
					log.Error("File submit with r.FormFile err: " + err.Error())
					return
				}
				defer file.Close()
				if !strings.HasSuffix(handler.Filename, taskinfo.FileType) {
					log.Error("File submit with wrong file type: submit:" + handler.Filename + " want:" + taskinfo.FileType)
					return
				}

				savepath := util.GetUploadFolerName(taskinfo)
				err = util.MakeDirIfNessesary(savepath)
				if err != nil {
					log.Error("File submit with MakeDir err: " + err.Error())
					return
				}

				nameMap, classMap := util.GetNameListMap()
				username, ok := nameMap[userinfo.StudentID]
				if !ok {
					username = "[Mystery]"
				}

				clsno, ok := classMap[userinfo.StudentID]
				if !ok || clsno != taskinfo.ClassNO {
					ok = false
				}

				filename := savepath + "/" + util.GetUploadFileName(taskinfo, userinfo, username, ok)

				f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					log.Error("File submit with Openfile err: " + err.Error())
					return
				}

				io.Copy(f, file)
				f.Close()

				fmd5, err := GetFileMD5(filename)
				if err != nil {
					log.Error("File submit with MD5 err: " + err.Error())
				}

				loginfo := database.UploadLog{
					SubmitTo: submmittoint,
					Email:    userinfo.Email,
					Time:     time.Now(),
					FileMD5:  fmd5,
				}
				database.Insert(loginfo)
				log.Info("File submit success at session: " + sess.SessionID() + " on IP:" + r.RemoteAddr + " of file:" + fmt.Sprint(loginfo))
			}
		}
	} else {
		log.Warn("File submit attempt with no account on IP:" + r.RemoteAddr)
	}

	http.Redirect(w, r, "/history", 303)
}

// GetFileMD5 to get md5
func GetFileMD5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	md5hash := md5.New()

	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(md5hash.Sum(nil)), nil
}
