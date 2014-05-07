package fake_graph_driver

import "sync"

type FakeGraphDriver struct {
	created     []CreatedGraph
	CreateError error

	removed     []string
	RemoveError error

	GetResult string
	GetError  error

	putted []string

	exists map[string]bool

	status [][2]string

	CleanupError error
	cleanedUp    bool

	sync.RWMutex
}

type CreatedGraph struct {
	ID     string
	Parent string
}

func New() *FakeGraphDriver {
	return &FakeGraphDriver{
		exists: make(map[string]bool),
	}
}

func (graph *FakeGraphDriver) String() string {
	return "fake"
}

func (graph *FakeGraphDriver) Create(id string, parent string) error {
	if graph.CreateError != nil {
		return graph.CreateError
	}

	graph.Lock()

	graph.created = append(graph.created, CreatedGraph{
		ID:     id,
		Parent: parent,
	})

	graph.Unlock()

	return nil
}

func (graph *FakeGraphDriver) Created() []CreatedGraph {
	graph.RLock()

	created := make([]CreatedGraph, len(graph.created))
	copy(created, graph.created)

	graph.RUnlock()

	return created
}

func (graph *FakeGraphDriver) Remove(id string) error {
	if graph.RemoveError != nil {
		return graph.RemoveError
	}

	graph.Lock()

	graph.removed = append(graph.removed, id)

	graph.Unlock()

	return nil
}

func (graph *FakeGraphDriver) Removed() []string {
	graph.RLock()

	removed := make([]string, len(graph.removed))
	copy(removed, graph.removed)

	graph.RUnlock()

	return removed
}

func (graph *FakeGraphDriver) Get(id string, mountLabel string) (string, error) {
	if graph.GetError != nil {
		return "", graph.GetError
	}

	return graph.GetResult, nil
}

func (graph *FakeGraphDriver) Put(id string) {
	graph.Lock()

	graph.putted = append(graph.putted, id)

	graph.Unlock()
}

func (graph *FakeGraphDriver) Putted() []string {
	graph.RLock()

	putted := make([]string, len(graph.putted))
	copy(putted, graph.putted)

	graph.RUnlock()

	return putted
}

func (graph *FakeGraphDriver) Exists(id string) bool {
	graph.RLock()
	defer graph.RUnlock()

	return graph.exists[id]
}

func (graph *FakeGraphDriver) SetExists(id string, exists bool) {
	graph.Lock()

	graph.exists[id] = exists

	graph.Unlock()
}

func (graph *FakeGraphDriver) Status() [][2]string {
	graph.RLock()
	defer graph.RUnlock()

	return graph.status
}

func (graph *FakeGraphDriver) SetStatus(status [][2]string) {
	graph.Lock()

	graph.status = status

	graph.Unlock()
}

func (graph *FakeGraphDriver) Cleanup() error {
	if graph.CleanupError != nil {
		return graph.CleanupError
	}

	graph.Lock()

	graph.cleanedUp = true

	graph.Unlock()

	return nil
}

func (graph *FakeGraphDriver) CleanedUp() bool {
	graph.RLock()
	defer graph.RUnlock()

	return graph.cleanedUp
}
