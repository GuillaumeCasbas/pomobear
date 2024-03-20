package domain

import (
	"time"
)

type Store struct {
	repository PomodoroRepository
	Pomodoros  []Pomodoro
}

func NewStore(r PomodoroRepository) *Store {
	data, _ := r.GetAll()

	return &Store{
		repository: r,
		Pomodoros:  data,
	}
}

func (s *Store) Add(pomodoro Pomodoro) error {
	s.Pomodoros = append(s.Pomodoros, pomodoro)

	err := s.repository.Save(s.Pomodoros)

	return err
}

func (s *Store) GetCurrent() (*Pomodoro, bool) {
	for i := range s.Pomodoros {
		p := &s.Pomodoros[i]

		if p.Endt.After(time.Now()) {
			return p, true
		}
	}
	return &Pomodoro{}, false
}
