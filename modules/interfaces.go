package modules

type SensorDecoder interface {
	Decode(raw RawData) (Measurement, error)
	SensorName() string
}
