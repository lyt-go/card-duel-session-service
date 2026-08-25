package service

import (
	"errors"

	"cardgame/internal/archive"
)

type ExportAudit struct{ States []string }

func ExportMatchReports(ids []string, pool *archive.HandlePool, tx *archive.Transaction, audit *ExportAudit) (err error) {
	defer func() { err = tx.Commit() }()
	for _, id := range ids {
		h, acquireErr := pool.Acquire()
		if acquireErr != nil {
			audit.States = append(audit.States, "failed")
			return acquireErr
		}
		defer h.Close()
		if id == "broken" {
			audit.States = append(audit.States, "success")
			return errors.New("战报内容损坏")
		}
	}
	audit.States = append(audit.States, "success")
	return nil
}
