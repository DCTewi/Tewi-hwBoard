package controllers

import (
	"net/http"
	"text/template"
	"time"

	"github.com/dctewi/tewi-hwboard/config"
	"github.com/dctewi/tewi-hwboard/core/database"

	log "unknwon.dev/clog/v2"
)

// IController interface
type IController interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

// ToView of the controller
func ToView(w http.ResponseWriter, name string, model interface{}) {
	t := template.New(name)
	t = t.Funcs(template.FuncMap{
		"tr":          DealWebConstance,
		"time":        DealTimeFormat,
		"checktime":   DealTimeAvailable,
		"availdcolor": DealAvailableColor,
		"s2bool":      DealAvailablePipeline,
	})
	t, err := t.ParseFiles("frame/views/"+name+".tpl", "frame/views/header.tpl", "frame/views/footer.tpl")
	if err != nil {
		log.Error("Convert view error: " + err.Error())
	}

	t.ExecuteTemplate(w, "layout", model)
}

// DealWebConstance {{"Key" | tr}}
func DealWebConstance(args ...interface{}) string {
	if len(args) >= 1 {
		res, _ := config.WebConstance[args[0].(string)]

		return res
	}
	return ""
}

// DealTimeFormat {{.Time | time}}
func DealTimeFormat(args ...interface{}) string {
	if len(args) >= 1 {
		return args[0].(time.Time).Format("2006-01-02 15:04:05")
	}
	return ""
}

// DealTimeAvailable {{checktime .}}
func DealTimeAvailable(args ...interface{}) string {
	if len(args) >= 1 {
		info := args[0].(database.TaskInfo)

		l, r := info.Start, info.End
		now := time.Now()

		if l.Before(now) && r.After(now) {
			return DealWebConstance("Available")
		}
		return DealWebConstance("Unavailable")
	}
	return "Error"
}

// DealAvailableColor {{checktime . | availdcolor}}
func DealAvailableColor(args ...interface{}) string {
	if len(args) >= 1 {
		if args[0] == DealWebConstance("Available") {
			return "text-success"
		}
		return "text-danger"
	}
	return ""
}

// DealAvailablePipeline {{checktime . | s2bool}}
func DealAvailablePipeline(args ...interface{}) string {
	if len(args) >= 1 {
		if args[0] == DealWebConstance("Available") {
			return "True"
		}
	}
	return ""
}
