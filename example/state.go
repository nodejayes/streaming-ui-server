package example

import di "github.com/nodejayes/generic-di"

func init() {
	di.Injectable(NewAppState)
}

type AppState struct {
	Counter int
}

func NewAppState() *AppState {
	return &AppState{
		Counter: 0,
	}
}
