package modules

type SensorDecoder interface {
	Decode(raw RawData) Measurement
	SensorName() string
}
