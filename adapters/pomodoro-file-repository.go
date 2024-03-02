package adapters

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/guillaumecasbas/pomobear/domain"
)

// TODO: this file is a bit messy
var folder = "pomo-bear"
var fileName = "sessions.json"

type PomodoroFileRepository struct {
	homeDir   string
	storePath string
}

func NewPomodoroFileRepository() (*PomodoroFileRepository, error) {
	r := &PomodoroFileRepository{}
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return r, err
	}

	r.homeDir = homeDir
	r.storePath = filepath.Join(homeDir, folder, fileName)

	err = r.createStoreFileIfNotExist()

	if err != nil {
		return r, err
	}

	return r, nil
}

func (r *PomodoroFileRepository) Save(pomodoro domain.Pomodoro) error {
	pomodoros, err := r.getPomodoros()
	if err != nil {
		return err
	}
	pomodoros = append(pomodoros, pomodoro)
	r.writeFile(pomodoros)
	return nil
}

func (r *PomodoroFileRepository) GetCurrent() (domain.Pomodoro, bool) {
	ok := false
	currentPomo := domain.Pomodoro{}

	pomodoros, err := r.getPomodoros()
	if err != nil {
		return domain.Pomodoro{}, ok
	}

	for _, v := range pomodoros {
		if v.Endt.After(time.Now()) {
			currentPomo = v
			ok = true
			break
		}
	}
	return currentPomo, ok
}

func (r *PomodoroFileRepository) writeFile(pomodoros []domain.Pomodoro) error {
	file, err := os.OpenFile(r.storePath, os.O_RDWR, 0644)

	defer file.Close()
	if err != nil {
		return err
	}

	return json.NewEncoder(file).Encode(pomodoros)
}

func (r *PomodoroFileRepository) createStoreFileIfNotExist() error {
	storeFolder := filepath.Join(r.homeDir, folder)
	storePath := filepath.Join(storeFolder, fileName)
	if _, err := os.Stat(storeFolder); os.IsNotExist(err) {
		err := os.Mkdir(storeFolder, 0755)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(storePath); os.IsNotExist(err) {
		file, err := os.Create(storePath)
		file.Close()
		if err != nil {
			return err
		}

		r.writeFile([]domain.Pomodoro{})
	}

	return nil
}

func (r *PomodoroFileRepository) getPomodoros() ([]domain.Pomodoro, error) {
	var sessions = []domain.Pomodoro{}
	fileContent, err := os.ReadFile(r.storePath)

	if err != nil {
		return sessions, err
	}

	err = json.Unmarshal(fileContent, &sessions)

	return sessions, err
}
