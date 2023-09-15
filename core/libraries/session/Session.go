package session

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"
  "net/http"
)

// Types
type Session struct {
  Id string
  Data map[string]interface{}
  ExpireAt time.Time
}

type SessionManager struct {
  sessions map[string]Session
  mutex sync.Mutex
}

func NewSessionManager() *SessionManager {
  return &SessionManager{
    sessions: make(map[string]Session),
  }
}

func (manager *SessionManager) CreateSession() (*Session, error) {
  id, err := generateId()
  if err != nil {
    return nil, err
  }
  session := Session{
    Id: id,
    Data: make(map[string]interface{}),
    ExpireAt: time.Now().Add(6*time.Hour),
  }
  manager.mutex.Lock()
  manager.sessions[id] = session
  manager.mutex.Unlock()

  return &session, nil
}

func (manager *SessionManager) GetSession(id string) (*Session, bool) {
  manager.mutex.Lock()
  defer manager.mutex.Unlock()
  session, status := manager.sessions[id]
  if !status {
    return nil, false
  }
  if time.Now().After(session.ExpireAt) {
    delete(manager.sessions, id)
    return nil, false
  }
  return &session, true
}

func (manager *SessionManager) HasSession(id string) (bool) {
  _, status := manager.sessions[id]
  return status
}

func (manager *SessionManager) DeleteSession(id string) {
  manager.mutex.Lock()
  delete(manager.sessions, id)
  manager.mutex.Unlock()
}

func generateId() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
  session, err := NewSessionManager().CreateSession()
	if err != nil {
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	// Set a session cookie in the response to track the session.
	http.SetCookie(w, &http.Cookie{
		Name:     "sdf_sess",
		Value:    session.Id,
		Expires:  session.ExpireAt,
		HttpOnly: true,
	})
}

