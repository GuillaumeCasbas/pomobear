package usecases

import (
	"time"

	"github.com/guillaumecasbas/pomobear/domain"
)

type PomodoroUsecases struct {
	store domain.PomodoroStore
}

func NewPomodoroUsecases(repository domain.PomodoroRepository) PomodoroUsecases {
	s := domain.NewStore(repository)
	return PomodoroUsecases{s}
}

func (u PomodoroUsecases) Start() (bool, error) {
	_, ok := u.store.GetCurrent()

	if ok {
		return false, nil
	}

	err := u.store.Add(domain.NewPomodoro(time.Now()))

	if err != nil {
		return false, err
	}

	return true, nil
}

func (u PomodoroUsecases) Status() (int, error) {
	pomodoro, ok := u.store.GetCurrent()

	if !ok {
		return 0, nil
	}

	remainingTime := pomodoro.Endt.Sub(time.Now()) / time.Second

	return int(remainingTime), nil
}
