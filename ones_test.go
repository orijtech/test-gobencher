package exhibit

import "testing"

func TestCountOnes(t *testing.T) {
	for _, tt := range values {
		got := countOnes(tt.in)
		if got != tt.want {
			t.Errorf("%d:%x:: got=%d want=%d", tt.in, tt.in, got, tt.want)
		}
	}
}
