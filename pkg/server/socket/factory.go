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
	sessions map[string]Instance
	mutex    *sync.RWMutex
}

func Factory() *factory {
	return di.Inject[factory]()
}

func NewFactory() *factory {
	return &factory{
		sessions: make(map[string]Instance),
		mutex:    &sync.RWMutex{},
	}
}

func (ctx *factory) GetSession(id string) (Instance, error) {
	ctx.mutex.RLock()
	defer ctx.mutex.RUnlock()
	session, ok := ctx.sessions[id]
	if !ok {
		return nil, fmt.Errorf("no session found")
	}
	return session, nil
}

func (ctx *factory) GetSessions(selector func(socket Instance) bool) []Instance {
	ctx.mutex.RLock()
	defer ctx.mutex.RUnlock()
	sessions := make([]Instance, 0)
	for _, s := range ctx.sessions {
		if selector(s) {
			sessions = append(sessions, s)
		}
	}
	return sessions
}

func (ctx *factory) AddSession(id string, socket Instance) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	ctx.sessions[id] = socket
}

func (ctx *factory) RemoveSession(id string) {
	ctx.mutex.Lock()
	defer ctx.mutex.Unlock()
	delete(ctx.sessions, id)
}
