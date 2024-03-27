package usecases

import (
	"time"

	"github.com/guillaumecasbas/pomobear/domain"
)

type ConfFunc func(*PomodoroConfig)

type PomodoroConfig struct {
	duration int
}

func defaultConf() PomodoroConfig {
	return PomodoroConfig{
		duration: 25,
	}
}

func WithDuration(duration int) ConfFunc {
	return func(c *PomodoroConfig) {
		c.duration = duration
	}
}

type PomodoroUsecases struct {
	conf  PomodoroConfig
	store domain.PomodoroStore
}

func NewPomodoroUsecases(repository domain.PomodoroRepository, configs ...ConfFunc) PomodoroUsecases {
	c := defaultConf()

	for _, fn := range configs {
		fn(&c)
	}
	s := domain.NewStore(repository)
	return PomodoroUsecases{
		conf:  c,
		store: s,
	}
}

func (u PomodoroUsecases) Start() (bool, error) {
	_, ok := u.store.GetCurrent()

	if ok {
		return false, nil
	}

	err := u.store.Add(domain.NewPomodoro(time.Now(),u.conf.duration))

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

func (u PomodoroUsecases) Stop() (bool, error) {
	pomodoro, ok := u.store.GetCurrent()

	if !ok {
		return false, nil
	}

	pomodoro.Endt = time.Now()

	err := u.store.Save()
	if err != nil {
		ok = false
	}

	return ok, err
}
