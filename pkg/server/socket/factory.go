package socket

import (
	"fmt"
	"sync"

	di "github.com/nodejayes/generic-di"
)

func init() {
	di.Injectable(NewFactory)
}

type factory struct {
	sessions map[string][]Instance
	mutex    *sync.RWMutex
}

func Factory() *factory {
	return di.Inject[factory]()
}

func NewFactory() *factory {
	return &factory{
		sessions: make(map[string][]Instance),
		mutex:    &sync.RWMutex{},
	}
}

func (ctx *factory) GetClientSessions(clientID string) ([]Instance, error) {
	ctx.mutex.RLock()
	defer ctx.mutex.RUnlock()
	session, ok := ctx.sessions[clientID]
	if !ok {
		return nil, fmt.Errorf("no session found")
	}
	return session, nil
}

func (ctx *factory) GetSession(socketID string) (Instance, error) {
	ctx.mutex.RLock()
	defer ctx.mutex.RUnlock()
	for clientID := range ctx.sessions {
		for _, i := range ctx.sessions[clientID] {
			if i.GetSocketID() == socketID {
				return i, nil
			}
		}
	}
	return nil, fmt.Errorf("no session found")
}

func (ctx *factory) GetSessions(selector func(socket Instance) bool) []Instance {
	ctx.mutex.RLock()
	defer ctx.mutex.RUnlock()
	sessions := make([]Instance, 0)
	for _, s := range ctx.sessions {
		for _, i := range s {
			if selector(i) {
				sessions = append(sessions, i)
			}
		}
	}
	return sessions
}

func (ctx *factory) AddSession(clientID string, socket Instance) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	if _, ok := ctx.sessions[clientID]; !ok {
		ctx.sessions[clientID] = make([]Instance, 0)
	}
	ctx.sessions[clientID] = append(ctx.sessions[clientID], socket)
}

func (ctx *factory) RemoveSession(socketID string) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	for clientID, s := range ctx.sessions {
		for idx, i := range s {
			if i.GetSocketID() == socketID {
				s = append(s[:idx], s[idx+1:]...)
			}
		}
		if len(s) < 1 {
			delete(ctx.sessions, clientID)
		}
	}
}

func (ctx *factory) RemoveSessions(clientID string) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	delete(ctx.sessions, clientID)
}
