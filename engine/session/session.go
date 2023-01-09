package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"
	//"sevian.com/memory"
)

var machines = make(map[string]Machine)

func Register(name string, machine Machine) {
	if machine == nil {
		panic("machine is nil")
	}
	if _, ok := machines[name]; ok {
		panic("machine yet exists")
	}

	machines[name] = machine
}

type Session interface {
	Set(key string, value interface{}) //set session value
	Get(key string) interface{}        //get session value
	//Delete(key interface{}) error     //delete session value
	SessionID() string //back current sessionID
}
type Machine interface {
	Test()
	Init(string) Session
	Read(string) Session
}
type Store interface{}

type ConfigManager struct {
	MachineType string
	CookieName  string
	MaxLifeTime int64
}
type Manager struct {
	machineType string
	cookieName  string
	maxLifeTime int64
	machine     Machine
}

func Create(config ConfigManager) (*Manager, error) {
	machine, ok := machines[config.MachineType]

	if !ok {
		return nil, fmt.Errorf("error was not created %v", config.MachineType)
	}

	return &Manager{machineType: config.MachineType, cookieName: config.CookieName, maxLifeTime: config.MaxLifeTime, machine: machine}, nil

}

func (manager *Manager) Start(w http.ResponseWriter, req *http.Request) Session {

	cookieName := manager.cookieName
	var session Session
	sid, err := req.Cookie(cookieName)
	if err != nil {

		sessionId := manager.sessionId()
		session = manager.machine.Init(sessionId)

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{
			Name:    cookieName,
			Value:   sessionId,
			Expires: expiration,
			Path:    "/",
		}

		http.SetCookie(w, &cookie)

	} else {

		//session = manager.machine.Read(sid.Value)
		session = manager.machine.Init(sid.Value)
	}

	return session
}

func Init() {
	//Register("memory", &memory.Memory{ /*sessions: make(map[string]Store)*/ })
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (m *Manager) Test() {
	fmt.Println("es un test ", m.cookieName)

	m.machine.Test()

}
