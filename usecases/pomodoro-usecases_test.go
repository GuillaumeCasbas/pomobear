package usecases

import (
	"errors"
	"testing"
	"time"

	"github.com/guillaumecasbas/pomobear/domain"
)

var currentPomodoro = domain.Pomodoro{Startt: time.Now().Add(-25 * time.Second), Endt: time.Now().Add(25 * time.Minute)}

type PomoRepoMock struct {
	Calls         [][]domain.Pomodoro
	HasOneRunning bool
	SaveWillThrow bool
}

func (r *PomoRepoMock) Save(p []domain.Pomodoro) error {
	if r.SaveWillThrow {
		return errors.New("foo bar")
	}

    r.Calls = append(r.Calls, p)

	return nil
}

func (r *PomoRepoMock) GetAll() ([]domain.Pomodoro, error) {
	pomodoros := []domain.Pomodoro{}

	if r.HasOneRunning {
		pomodoros = append(pomodoros, currentPomodoro)

	}

	return pomodoros, nil
}

func TestStart(t *testing.T) {
	t.Run("creates a new pomodoro and returns true", func(t *testing.T) {
		r := &PomoRepoMock{}
		u := NewPomodoroUsecases(r)
		expectedStartt := time.Now().Round(60 * time.Second)
		expectedEndt := time.Now().Add(25 * time.Minute).Round(60 * time.Second)

		created, err := u.Start()

		if err != nil {
			t.Fatal("expect 0 error, got one")
		}

		if !created {
			t.Error("expect true got false")
		}

		if len(r.Calls) != 1 {
			t.Fatalf("expect %d calls, got %d", 1, len(r.Calls))
		}

		pomodoro := r.Calls[0][0]

		if pomodoro.Startt != expectedStartt {
			t.Errorf("expect %s, got %s", expectedStartt, pomodoro.Startt)
		}

		if pomodoro.Endt != expectedEndt {
			t.Errorf("expect %s, got %s", expectedEndt, pomodoro.Endt)
		}
	})

	t.Run("skips the pomodoro creation and returns false if a pomodoro is already running", func(t *testing.T) {
		r := &PomoRepoMock{HasOneRunning: true}
		u := NewPomodoroUsecases(r)

		created, err := u.Start()

		if err != nil {
			t.Fatal("expect no error, got one")
		}

		if created {
			t.Error("expect false, got true")
		}

		if len(r.Calls) != 0 {
			t.Fatal("expect 0 calls, got one")
		}
	})

	t.Run("returns the error on throws", func(t *testing.T) {
		r := &PomoRepoMock{SaveWillThrow: true}
		u := NewPomodoroUsecases(r)

		created, err := u.Start()

		if err == nil {
			t.Fatal("expect an error, got none")
		}

		if created {
			t.Error("expect false, got true")
		}
	})
}

func TestStatus(t *testing.T) {
	t.Run("returns the current pomodoro remaining time", func(t *testing.T) {
		r := &PomoRepoMock{HasOneRunning: true}
		u := NewPomodoroUsecases(r)
		expectedDuration := 1499
		time, err := u.Status()

		if err != nil {
			t.Fatal("got an error, want none")
		}

		if time != expectedDuration {
			t.Errorf("got %d, want %d", time, expectedDuration)
		}
	})

	t.Run("returns 0 when there is no pomodoro running", func(t *testing.T) {
		r := &PomoRepoMock{}
		u := NewPomodoroUsecases(r)

		time, err := u.Status()

		if err != nil {
			t.Fatal("expect no error, got one")
		}

		if time != 0 {
			t.Errorf("expect 0, got %d", time)
		}
	})
}
