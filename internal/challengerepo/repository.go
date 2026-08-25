package challengerepo

type Repository struct {
	Writes    int
	Rollbacks int
}
type Tx struct {
	repo   *Repository
	staged int
}

func (r *Repository) Begin() *Tx { return &Tx{repo: r} }
func (t *Tx) Stage()             { t.staged++ }
func (t *Tx) Commit()            { t.repo.Writes += t.staged }
func (t *Tx) Rollback()          { t.repo.Rollbacks++ }
