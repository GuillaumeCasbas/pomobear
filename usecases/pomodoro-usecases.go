package usecases

import (
	"time"

	"github.com/guillaumecasbas/pomobear/domain"
)

type PomodoroUsecases struct {
	Repo domain.PomodoroRepository
}

func NewPomodoroUsecases(repository domain.PomodoroRepository) PomodoroUsecases {
	return PomodoroUsecases{repository}
}

func (u PomodoroUsecases) Start() (bool, error) {
	_, ok := u.Repo.GetCurrent()

	if ok {
		return false, nil
	}

	err := u.Repo.Save(domain.NewPomodoro(time.Now()))

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u PomodoroUsecases) Status() (int, error) {
	pomodoro, ok := u.Repo.GetCurrent()

	if !ok {
		return 0, nil
	}

	remainingTime := pomodoro.Endt.Sub(time.Now()) / time.Second

	return int(remainingTime), nil
}
