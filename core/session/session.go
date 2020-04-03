package session

import (
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/dctewi/tewi-hwboard/core/util"

	log "unknwon.dev/clog/v2"
)

// Manager is the global session manager.
type Manager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

// Provider is the abstract struct of session.
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

// Session interface.
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

var provides = make(map[string]Provider)

// GlobalSessions manager
var GlobalSessions *Manager

// NewManager instantiate a new session manager.
func NewManager(provideName, cookieName string, maxLifeTime int64) *Manager {
	provider, ok := provides[provideName]
	if !ok {
		log.Fatal("session: unknown provide " + provideName + "(forgotten import?)")
		return nil
	}
	log.Info("Session manager inited: " + provideName + " - " + cookieName + " - " + strconv.FormatInt(maxLifeTime, 10))
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}
}

// Register makes a session provide available by provided name.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

// SessionStart start the session
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" || url.QueryEscape(cookie.Value) != cookie.Value {

		sid := manager.genSessionID()

		log.Info("Try read session: " + sid + " on IP:" + r.RemoteAddr + " from init.")
		session, _ = manager.provider.SessionInit(sid)

		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    sid,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   0, //int(manager.maxLifeTime),
		}

		http.SetCookie(w, &cookie)
	} else {
		sid := url.QueryEscape(cookie.Value)

		log.Info("Try read session: " + sid + " on IP:" + r.RemoteAddr + " from cookie.")
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

// SessionDestroy destroy sessions
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)

	if err != nil || cookie.Value == "" {
		// Won't happen, do nothing
	} else {
		manager.lock.Lock()
		defer manager.lock.Unlock()

		log.Info("Session destroy attempt on IP:" + r.RemoteAddr)
		manager.provider.SessionDestroy(cookie.Value)

		expiration := time.Now()
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Expires:  expiration,
			MaxAge:   -1,
		}
		http.SetCookie(w, &cookie)
	}
}

// GC collects expired sessions, should be called and only called
func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	log.Info("Session GC triggered")
	manager.provider.SessionGC(manager.maxLifeTime)

	time.AfterFunc(time.Duration(manager.maxLifeTime)*time.Second, func() {
		manager.GC()
	})
}

func (manager *Manager) genSessionID() string {
	// b := make([]byte, 32)
	// if _, err := rand.Read(b); err != nil {
	// 	return ""
	// }
	// return url.QueryEscape(base64.URLEncoding.EncodeToString(b))
	return util.GenSessionID()
}
