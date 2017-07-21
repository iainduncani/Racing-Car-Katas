package tirepressuremonitoringsystem

import (
	"math/rand"
	"time"
)

const offset float32 = 16

// The reading of the pressure value from the sensor is simulated in this implementation.
// Because the focus of the exercise is on the other class.

type Sensor struct {
}

func (s *Sensor) PopNextPressurePsiValue() float32 {
	var pressureTelemetryValue float32
	pressureTelemetryValue = samplePressure()

	return offset + pressureTelemetryValue
}

func samplePressure() float32 {
	// placeholder implementation that simulate a real sensor in a real tire
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	pressureTelemetryValue := 6 * random.Float32() * random.Float32()
	return pressureTelemetryValue
}
