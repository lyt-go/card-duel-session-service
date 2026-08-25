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
func NewEffect() *Effect { return &Effect{seen: make(map[string]bool)} }
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

// WriteCacheIfVersion 仅当新版本号严格更高时才写入缓存，
// 防止迟到的首轮回调把已经到达终态（succeeded）的列表覆盖回进行中（running）。
func (s *Store) WriteCacheIfVersion(id string, state State) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if state.Version > s.cache[id].Version {
		s.cache[id] = state
	}
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
