package domain

import "time"

var (
	roundValue = time.Second
)

type Pomodoro struct {
	Startt time.Time
	Endt   time.Time
}

func NewPomodoro(startt time.Time, duration int) Pomodoro {
	d := time.Duration(duration) * time.Minute

	return Pomodoro{
		Startt: startt.Round(roundValue),
		Endt:   startt.Add(d).Round(roundValue),
	}
}

type PomodoroStore interface {
	GetCurrent() (*Pomodoro, bool)
	Add(p Pomodoro) error
	Save() error
}

type PomodoroRepository interface {
	GetAll() ([]Pomodoro, error)
	Save(p []Pomodoro) error
}
