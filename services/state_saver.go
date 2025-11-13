package services

import (
	"encoding/gob"
	"os"
	"sync"
	"time"
)

type LinkSet struct {
	Links   map[string]bool `json:"links"`
	Created time.Time       `json:"created"`
}

type AppState struct {
	LinksData map[int]LinkSet `json:"links_data"`
	NextID    int             `json:"next_id"`
}

var (
	stateMutex sync.RWMutex
	appState   AppState
)

func init() {
	if loadedState, err := LoadStateGob(); err == nil {
		appState = loadedState
	} else {
		appState = AppState{
			LinksData: make(map[int]LinkSet),
			NextID:    1,
		}
	}
}

func SaveLinksSet(links map[string]bool) (int, error) {
	stateMutex.Lock()
	defer stateMutex.Unlock()

	if appState.LinksData == nil {
		appState.LinksData = make(map[int]LinkSet)
	}

	setID := appState.NextID
	appState.LinksData[setID] = LinkSet{
		Links:   links,
		Created: time.Now(),
	}

	appState.NextID++

	err := SaveStateGob(appState)
	return setID, err
}

func GetLinksSet(setID int) (LinkSet, bool) {
	stateMutex.RLock()
	defer stateMutex.RUnlock()

	linkSet, exists := appState.LinksData[setID]
	return linkSet, exists
}

func SaveStateGob(state AppState) error {
	file, err := os.Create("state.gob")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(state)
}

func LoadStateGob() (AppState, error) {
	var state AppState
	file, err := os.Open("state.gob")
	if err != nil {
		return state, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&state)
	return state, err
}
