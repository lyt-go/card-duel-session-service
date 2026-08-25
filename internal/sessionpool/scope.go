package sessionpool

type Scope struct {
	PlayerID string
	Labels   []string
}
type Pool struct{ cached *Scope }

func New() *Pool { return &Pool{} }
func (p *Pool) Get() *Scope {
	if p.cached == nil {
		return &Scope{}
	}
	scope := p.cached
	p.cached = nil
	return scope
}
func (p *Pool) Put(scope *Scope) { p.cached = scope }
