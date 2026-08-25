package sessionpool

// Scope 保存一次对局请求范围内的玩家身份与标签。
// Scope 会被 Pool 缓存复用，因此归还时必须清空身份字段，
// 否则下一次 Get 会拿到上一个玩家的 PlayerID 与累积的 Labels。
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

// Put 将 Scope 归还到池中复用，并清空身份字段，
// 保证下一次 Get 返回的是一个干净、无残留状态的 Scope。
func (p *Pool) Put(scope *Scope) {
	scope.PlayerID = ""
	scope.Labels = scope.Labels[:0]
	p.cached = scope
}
