package memory

import (
	"fmt"

	"sevian.com/whendy/session"
)

type Memory struct {
	sessions map[string]session.Session
}

func (m *Memory) Init(sid string) session.Session {

	if m.sessions == nil {
		m.sessions = make(map[string]session.Session)
	}
	_, ok := m.sessions[sid]
	if !ok {
		m.sessions[sid] = Session{data: make(map[string]interface{})}
	}

	return m.sessions[sid]
}

func (m *Memory) Read(sid string) session.Session {

	return m.sessions[sid]
}

func (Memory) Test() {

	fmt.Println("session Test")
}

type Session struct {
	data map[string]interface{}
}

func (s Session) Set(key string, value interface{}) {
	s.data[key] = value

	fmt.Println("el primer valor de la sesion es ", s.data[key])
}

func (s Session) Get(key string) interface{} {
	return s.data[key]
}

func (s *Session) Delete(key string) {
	delete(s.data, key)
}

func (s Session) SessionID() string {
	return "xxx"
}
