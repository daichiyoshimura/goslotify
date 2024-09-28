package timeslots_test

import (
	"os"
	"testing"
	"time"
)

var now time.Time

func TestMain(m *testing.M) {
	now = time.Now()

	code := m.Run()

	os.Exit(code)
}
