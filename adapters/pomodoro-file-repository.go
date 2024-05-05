package adapters

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/guillaumecasbas/pomobear/domain"
)

var folder = "pomobear"
var fileName = "pomodoros.json"

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

func (r *PomodoroFileRepository) Save(pomodoros []domain.Pomodoro) error {
	err := r.writeFile(pomodoros)

	return err
}

func (r *PomodoroFileRepository) GetAll() ([]domain.Pomodoro, error) {
	pomodoros, err := r.getPomodoros()

	if err != nil {
		return []domain.Pomodoro{}, err
	}

	return pomodoros, nil
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
	var pomodoros = []domain.Pomodoro{}
	fileContent, err := os.ReadFile(r.storePath)

	if err != nil {
		return pomodoros, err
	}

	err = json.Unmarshal(fileContent, &pomodoros)

	return pomodoros, err
}
