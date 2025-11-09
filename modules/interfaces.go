package modules

/*
SensorDecoder is the interface that wraps the basic Decode method.
It also includes the SensorName method to identify the sensor type.
*/
type SensorDecoder interface {
	Decode(raw RawData) Measurement
	SensorName() string
}
