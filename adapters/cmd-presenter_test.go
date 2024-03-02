package adapters

import (
	"fmt"
	"testing"
)

func TestDisplayStatus(t *testing.T) {
	t.Run("return empty when remaining time is 0", func(t *testing.T) {
		p := NewCmdPresenter()

		expected := "💤 Idle"
		result := p.DisplayStatus(0)

		if result != expected {
			t.Errorf("want %s, got %s", result, expected)
		}
	})

	cases := []struct {
		Remaining int
		Expected  string
	}{
		{1500, "🍅 25:00"},
		{1, "🍅 00:01"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("returns %q when the remaining time is %d", test.Expected, test.Remaining), func(t *testing.T) {
			p := NewCmdPresenter()

			result := p.DisplayStatus(test.Remaining)

			if result != test.Expected {
				t.Errorf("want %s, got %s", test.Expected, result)
			}
		})
	}

}
