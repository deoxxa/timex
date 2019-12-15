package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDefaults(t *testing.T) {
	for _, p := range []struct {
		s string
		t time.Time
		e error
	}{
		{"2019-01-01 12:00:00", time.Date(2019, time.January, 1, 12, 0, 0, 0, time.UTC), nil},
	} {
		t.Run(p.s, func(t *testing.T) {
			a := assert.New(t)
			v, err := ParseDefaults(p.s)
			a.Equal(v, p.t)
			a.Equal(err, p.e)
		})
	}
}
