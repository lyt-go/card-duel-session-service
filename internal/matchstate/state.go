package matchstate

import "sync"

type State struct {
	Status  string
	Version int
}

type Store struct {
	mu     sync.RWMutex
	detail map[string]State
	cache  map[string]State
}

func NewStore() *Store { return &Store{detail: make(map[string]State), cache: make(map[string]State)} }
func (s *Store) WriteDetail(id string, state State) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.detail[id] = state
}
func (s *Store) WriteCache(id string, state State) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache[id] = state
}
func (s *Store) Detail(id string) State { s.mu.RLock(); defer s.mu.RUnlock(); return s.detail[id] }
func (s *Store) Listed(id string) State { s.mu.RLock(); defer s.mu.RUnlock(); return s.cache[id] }

type Effect struct {
	mu      sync.Mutex
	Calls   int
	Actions []string
	seen    map[string]bool
}

func (e *Effect) Apply(key string) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.seen == nil {
		e.seen = make(map[string]bool)
	}
	e.Calls++
	if !e.seen[key] {
		e.seen[key] = true
		e.Actions = append(e.Actions, key)
	}
	if e.Calls == 1 {
		return errTemporary{}
	}
	return nil
}

type errTemporary struct{}

func (errTemporary) Error() string { return "临时失败" }
