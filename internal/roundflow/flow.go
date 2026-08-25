package roundflow

import "errors"

var ErrTemporary = errors.New("回合广播暂时失败")
var ErrRollback = errors.New("回滚连接已关闭")

type Repository struct {
	CommitCount   int
	RollbackCount int
}
type Transaction struct{ repo *Repository }

func (r *Repository) Begin() *Transaction { return &Transaction{repo: r} }
func (t *Transaction) Commit() error      { t.repo.CommitCount++; return nil }
func (t *Transaction) Rollback() error    { t.repo.RollbackCount++; return ErrRollback }

type Publisher struct {
	FailFirst bool
	Calls     int
	Events    []string
	seen      map[string]bool
}

func (p *Publisher) Publish(eventID string) error {
	if p.seen == nil {
		p.seen = make(map[string]bool)
	}
	p.Calls++
	if !p.seen[eventID] {
		p.seen[eventID] = true
		p.Events = append(p.Events, eventID)
	}
	if p.FailFirst && p.Calls == 1 {
		return ErrTemporary
	}
	return nil
}

type Audit struct{ States []string }
