package tirepressuremonitoringsystem

const lowPressureThreshold float32 = 17
const highPressureThreshold float32 = 21

type Alarm struct {
	AlarmOn bool
	sensor  *Sensor
}

func NewAlarm() Alarm {
	return Alarm{
		sensor: &Sensor{},
	}
}
