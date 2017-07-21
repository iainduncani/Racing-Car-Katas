package tirepressuremonitoringsystem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	alarm := NewAlarm()

	assert.False(t, alarm.AlarmOn)
}

func TestSomethingElse(t *testing.T) {
	alarm := NewAlarm()

	for i := 0; i < 100; i++ {
		alarm.check()
		if alarm.AlarmOn {
			fmt.Printf("Triggered alarm on run %d", i)
			return
		}
	}

	assert.False(t, alarm.AlarmOn)
}
