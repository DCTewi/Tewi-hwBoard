package route

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dctewi/tewi-hwboard/config"
	"github.com/dctewi/tewi-hwboard/frame/controllers"

	log "unknwon.dev/clog/v2"
)

// Mux is the custom mux of hwBoard
type Mux struct {
	routeMap  map[string]controllers.IController
	staticMap map[string]string
}

func (p *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Handler crashed: " + err.(string))
			panic(err)
		}
	}()

	foundPage := false

	for prefix, staticDir := range p.staticMap {
		if strings.HasPrefix(r.URL.Path, prefix) {
			file := staticDir + r.URL.Path[len(prefix):]
			http.ServeFile(w, r, file)
			log.Info("Get static file request on IP:" + r.RemoteAddr + " of file:" + r.URL.Path)
			foundPage = true
			return
		}
	}

	if controller, ok := p.routeMap[r.URL.Path]; ok {
		if r.Method == "GET" {
			log.Info("Get GET request on IP:" + r.RemoteAddr + " of path:" + r.URL.Path)
			controller.Get(w, r)
		} else {
			r.ParseForm()
			log.Info("Get POST request on IP:" + r.RemoteAddr + " of path:" + r.URL.Path + " with form:" + fmt.Sprint(r.Form))
			controller.Post(w, r)
		}
		foundPage = true
	}

	if !foundPage {
		http.NotFound(w, r)
	}
}

// RegiterController to bind path and controller
func (p *Mux) RegiterController(path string, controller controllers.IController) {
	if p.routeMap == nil {
		p.routeMap = make(map[string]controllers.IController)
	}

	p.routeMap[path] = controller

	if _, ok := p.routeMap[path]; !ok {
		panic("register error " + path)
	}
}

// RegisterStaticPath to bind path and static path
func (p *Mux) RegisterStaticPath(urlpath, filepath string) {
	if p.staticMap == nil {
		p.staticMap = make(map[string]string)
	}

	p.staticMap[urlpath] = filepath

	if _, ok := p.staticMap[urlpath]; !ok {
		panic("register error: " + urlpath)
	}
}

// RedirectSSL redirect to ssl
func RedirectSSL(w http.ResponseWriter, r *http.Request) {
	host := strings.Split(r.Host, ":")
	if len(host) > 1 {
		host[1] = config.App.SSLPort[1:]
	}

	log.Info("Get HTTP request with host:" + r.URL.Host + " path:" + r.URL.Path)
	target := "https://" + strings.Join(host, ":") + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	log.Info("Redirect to " + target)
	http.Redirect(w, r, target, http.StatusMovedPermanently)
}
