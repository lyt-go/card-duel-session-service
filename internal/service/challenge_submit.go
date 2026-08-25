package service

import (
	"errors"

	"cardgame/internal/challengerepo"
	"cardgame/internal/gateway"
)

func SubmitChallenge(mode string, client *gateway.Client, repo *challengerepo.Repository) error {
	var last error
	for attempt := 0; attempt < 2; attempt++ {
		tx := repo.Begin()
		tx.Stage()
		if err := client.Challenge(mode); err != nil {
			tx.Commit()
			last = err
			var remote *gateway.RemoteError
			if errors.As(err, &remote) && remote.Kind == "rejected" {
				return err
			}
			continue
		}
		tx.Commit()
		return nil
	}
	return last
}
