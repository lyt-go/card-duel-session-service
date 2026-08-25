package archive

import "errors"

var ErrHandleLimit = errors.New("战报写入句柄已耗尽")

type HandlePool struct {
	Open int
	Max  int
}

type Handle struct{ pool *HandlePool }

func (p *HandlePool) Acquire() (*Handle, error) {
	if p.Open >= p.Max {
		return nil, ErrHandleLimit
	}
	p.Open++
	return &Handle{pool: p}, nil
}

func (h *Handle) Close() error {
	h.pool.Open--
	return nil
}

type Transaction struct {
	Committed  bool
	RolledBack bool
}

func (t *Transaction) Commit() error   { t.Committed = true; return nil }
func (t *Transaction) Rollback() error { t.RolledBack = true; return nil }
