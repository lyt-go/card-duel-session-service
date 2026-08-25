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
		// 先调用下游；失败时一律不落盘任何状态。
		if err := client.Challenge(mode); err != nil {
			tx.Rollback()

			// 明确拒绝属于业务终态：立即返回，不重试、不留状态。
			// 调用方可通过 errors.As 识别为拒绝。
			var remote *gateway.RemoteError
			if errors.As(err, &remote) && remote.Kind == "rejected" {
				return err
			}
			last = err
			continue
		}
		tx.Commit()
		return nil
	}
	return last
}
