package adapters

import (
	"fmt"
)

type CmdPresenter struct{}

func NewCmdPresenter() CmdPresenter {
	return CmdPresenter{}
}

func (p CmdPresenter) DisplayStatus(remainingTime int) string {
	if remainingTime > 0 {
		minutes := remainingTime / 60
		seconds := remainingTime % 60

		return fmt.Sprintf("🍅 %02d:%02d", minutes, seconds)
	}

	return "💤 Idle"
}
