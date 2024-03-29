package domain

import (
	"testing"
	"time"
)

type RepoMock struct {
	storage     []Pomodoro
	SaveCalls   [][]Pomodoro
	throwOnSave bool
}

func NewRepoMock(storage *[]Pomodoro, throwOnSave bool) *RepoMock {
	return &RepoMock{
		storage:     *storage,
		SaveCalls:   [][]Pomodoro{},
		throwOnSave: throwOnSave,
	}
}
func (r *RepoMock) GetAll() ([]Pomodoro, error) {
	return r.storage, nil
}

func (r *RepoMock) Save(p []Pomodoro) error {
	r.SaveCalls = append(r.SaveCalls, p)

	return nil
}

func TestNew(t *testing.T) {
	t.Run("initializes the store from the repository", func(t *testing.T) {
		emptyStorage := &[]Pomodoro{}
		storage := &[]Pomodoro{
			NewPomodoro(time.Now(), 29),
			NewPomodoro(time.Now().Add(2 * time.Minute), 13),
		}
		r1 := NewRepoMock(emptyStorage, false)
		r2 := NewRepoMock(storage, false)
		s1 := NewStore(r1)
		s2 := NewStore(r2)

		if len(s1.Pomodoros) != 0 {
			t.Errorf("expect 0, got %d", len(s1.Pomodoros))
		}

		if len(s2.Pomodoros) != 2 {
			t.Errorf("expect 2, got %d", len(s2.Pomodoros))
		}
	})
}

func TestStoreAdd(t *testing.T) {
	t.Run("adds a new pomodoro", func(t *testing.T) {
		r := &RepoMock{}
		p := NewPomodoro(time.Now(), 25)
		s := NewStore(r)

		err := s.Add(p)

		if err != nil {
			t.Fatal("exepct no error, got one")
		}

		if len(r.SaveCalls) != 1 {
			t.Errorf("expect 1 call, got %d", len(r.SaveCalls))
		}
	})
}

func TestGetCurrent(t *testing.T) {
	t.Run("returns false if there is no current pomodoro", func(t *testing.T) {
		storage := &[]Pomodoro{
			{
				Startt: time.Now().AddDate(0, -1, 0),
				Endt:   time.Now().AddDate(0, -1, 3),
			},
		}

		r := NewRepoMock(storage, false)
		s := NewStore(r)

		_, ok := s.GetCurrent()

		if ok {
			t.Errorf("expected false, got %v", ok)
		}

	})

	t.Run("returns true and the current runnning pomodoro", func(t *testing.T) {
		expectedPomodoro := &Pomodoro{
			Startt: time.Now(),
			Endt:   time.Now().Add(10 * time.Minute),
		}
		storage := &[]Pomodoro{
			*expectedPomodoro,
		}

		r := NewRepoMock(storage, false)
		s := NewStore(r)

		p, ok := s.GetCurrent()
		if !ok {
			t.Fatalf("expected true, got %v", ok)
		}

		if p.Endt != expectedPomodoro.Endt {
			t.Errorf("expect %v, got %v", expectedPomodoro.Endt, p.Endt)
		}
	})
}

func TestSave(t *testing.T) {
	t.Run("saves the store", func(t *testing.T) {
		storage := &[]Pomodoro{
			NewPomodoro(time.Now(), 25),
			NewPomodoro(time.Now().Add(2 * time.Minute), 25),
		}
		r := NewRepoMock(storage, false)
		s := NewStore(r)

		s.Save()

		if len(r.SaveCalls) != 1 {
			t.Errorf("expect 1 call, got %d", len(r.SaveCalls))
		}
	})

	t.Run("returns the error on throw", func(t *testing.T) {
		storage := &[]Pomodoro{
			NewPomodoro(time.Now(), 25),
			NewPomodoro(time.Now().Add(2 * time.Minute), 10),
		}
		r := NewRepoMock(storage, true)
		s := NewStore(r)

		err := s.Save()

		if err != nil {
			t.Error("expect an error, got none")
		}
	})
}
