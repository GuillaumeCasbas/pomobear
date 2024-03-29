package domain

import (
	"testing"
	"time"
)

func TestNewPomodoro(t *testing.T) {
	t.Run("create a pomodoro rounded to second", func(t *testing.T) {
        duration := 40
		expectedStart, _ := time.Parse("2006-01-02 15:04:05", "1991-08-11 20:58:06")
		expectedEnd, _ := time.Parse("2006-01-02 15:04:05", "1991-08-11 21:38:06")
		time := time.Date(1991, time.August, 11, 20, 58, 6, 0, time.UTC)

		p := NewPomodoro(time, duration)

		if expectedStart != p.Startt {
			t.Errorf("expect %v, got %v", expectedStart, p.Startt)
		}

		if expectedEnd != p.Endt {
			t.Errorf("expect %v, got %v", expectedEnd, p.Endt)
		}
	})

}

