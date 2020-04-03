package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dctewi/tewi-hwboard/config"
	"github.com/dctewi/tewi-hwboard/core/route"
	"github.com/dctewi/tewi-hwboard/core/session"
	"github.com/dctewi/tewi-hwboard/core/util"
	"github.com/dctewi/tewi-hwboard/frame/controllers"

	log "unknwon.dev/clog/v2"

	_ "github.com/dctewi/tewi-hwboard/core/session/memory"
)

func init() {
	// Logger
	logfilepath := config.Log.Filepath
	util.MakeDirIfNessesary(logfilepath)
	if !strings.HasSuffix(logfilepath, "/") {
		logfilepath += "/"
	}
	logfilepath += "running.log"

	err := log.NewConsole(100, log.ConsoleConfig{
		Level: log.LevelInfo,
	})
	if err != nil {
		panic("unable to create logger: " + err.Error())
	}
	err = log.NewFile(100, log.FileConfig{
		Level:    log.LevelInfo,
		Filename: logfilepath,
		FileRotationConfig: log.FileRotationConfig{
			Rotate: true,
			Daily:  true,
		},
	})
	if err != nil {
		panic("unable to create logger: " + err.Error())
	}

	// Session manager
	session.GlobalSessions = session.NewManager("memory", "sid", 600)
	go session.GlobalSessions.GC()
}

func main() {
	defer func() {
		log.Info("Stopping application...")
		log.Stop()
	}()

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			log.Fatal(x.(string))
		}
	}()

	log.Info("Starting application...")

	mux := &route.Mux{}

	mux.RegiterController("/", &controllers.HomeController{})
	mux.RegiterController("/login", &controllers.LoginController{})
	mux.RegiterController("/admin", &controllers.AdminController{})
	mux.RegiterController("/history", &controllers.HistoryController{})
	mux.RegiterController("/submit", &controllers.SubmitController{})

	mux.RegisterStaticPath("/img", config.Path.StaticFoler+"/img")
	mux.RegisterStaticPath("/css", config.Path.StaticFoler+"/css")
	mux.RegisterStaticPath("/js", config.Path.StaticFoler+"/js")
	mux.RegisterStaticPath("/favicon.ico", config.Path.StaticFoler+"/img/favicon.ico")

	log.Info("Application started, listening...")

	var err error = nil
	if config.App.UseTLS {
		go http.ListenAndServe(config.App.HTTPPort, http.HandlerFunc(route.RedirectSSL))
		err = http.ListenAndServeTLS(config.App.SSLPort, config.App.TLSCrtPath, config.App.TLSKeyPath, mux)
	} else {
		err = http.ListenAndServe(config.App.HTTPPort, mux)
	}

	if err != nil {
		log.Fatal("Application fatal error: " + err.Error())
	}
}
