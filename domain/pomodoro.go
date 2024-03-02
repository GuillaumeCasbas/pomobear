package domain

import "time"

var (
    duration = 25 * time.Minute
    roundValue = 60 * time.Second
)

type Pomodoro struct {
	Startt time.Time
	Endt   time.Time
}

func NewPomodoro(startt time.Time) Pomodoro {
	return Pomodoro{
		Startt: startt.Round(roundValue),
		Endt:   startt.Add(duration).Round(roundValue),
	}
}

type PomodoroRepository interface {
	GetCurrent() (Pomodoro, bool)
	Save(p Pomodoro) error
}
