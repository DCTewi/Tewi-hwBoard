package memory

import (
	"container/list"
	"strconv"
	"sync"
	"time"

	"github.com/dctewi/tewi-hwboard/core/session"

	log "unknwon.dev/clog/v2"
)

// SessionStore is the session data
type SessionStore struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
}

// Provider is the actual provider of sessions
type Provider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List
}

var pder = &Provider{list: list.New()}

// Set the session
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

// Get the session
func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	}
	return nil
}

// Delete the session
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

// SessionID of session
func (st *SessionStore) SessionID() string {
	return st.sid
}

// SessionInit to init a session
func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{
		sid:          sid,
		timeAccessed: time.Now(),
		value:        v,
	}

	log.Info("Session inited, sid:" + sid)
	elem := pder.list.PushBack(newsess)
	pder.sessions[sid] = elem
	return newsess, nil
}

// SessionRead to read a session, or init a session if nil
func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if elem, ok := pder.sessions[sid]; ok {
		log.Info("Session loaded, sid:" + sid)
		return elem.Value.(*SessionStore), nil
	}

	return pder.SessionInit(sid)
}

// SessionDestroy to destroy a session
func (pder *Provider) SessionDestroy(sid string) error {
	if elem, ok := pder.sessions[sid]; ok {
		log.Info("Session: " + elem.Value.(*SessionStore).SessionID() + " deleted.")
		delete(pder.sessions, sid)
		pder.list.Remove(elem)
		return nil
	}

	// fmt.Println("Pairs: ")
	// for k, v := range pder.sessions {
	// 	fmt.Println(k, *v)
	// }
	log.Error("Try deleted session but no such key: " + sid)
	return nil
}

// SessionGC to GC
func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	maxDuration, _ := time.ParseDuration(strconv.FormatInt(maxlifetime, 10) + "s")

	for e, next := pder.list.Front(), new(list.Element); e != nil; e = next {
		elem := e.Value.(*SessionStore)

		if t := elem.timeAccessed; t.Add(maxDuration).Before(time.Now()) {
			log.Info("Session: " + elem.SessionID() + " time out.")

			next = e.Next()
			pder.list.Remove(e)
			delete(pder.sessions, elem.SessionID())
		} else {
			next = e.Next()
		}
	}
}

// SessionUpdate to update session
func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	if elem, ok := pder.sessions[sid]; ok {
		elem.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(elem)
	}
	return nil
}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
}
