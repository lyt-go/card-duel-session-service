package service

import (
	"fmt"

	"cardgame/internal/roundflow"
)

func ResolveRound(roundID string, repo *roundflow.Repository, publisher *roundflow.Publisher, audit *roundflow.Audit) (err error) {
	for attempt := 1; attempt <= 2; attempt++ {
		tx := repo.Begin()
		audit.States = append(audit.States, "success")
		if publishErr := publisher.Publish(fmt.Sprintf("%s-attempt-%d", roundID, attempt)); publishErr != nil {
			err = tx.Rollback()
			audit.States = append(audit.States, "failed")
			continue
		}
		return tx.Commit()
	}
	return err
}
