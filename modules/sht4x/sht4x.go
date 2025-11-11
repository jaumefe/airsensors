package sht4x

import "github.com/jaumefe/airsensors/modules"

type SHT4XDecoder struct{}

func NewDecoder() *SHT4XDecoder {
	return &SHT4XDecoder{}
}

func (d *SHT4XDecoder) Decode(raw modules.RawData) modules.Measurement {
	meas := modules.Measurement{}
	meas.Temperature = d.temperatureFromRaw(raw.RawTemperature)
	meas.Humidity = d.humidityFromRaw(raw.RawHumidity)
	return meas
}

func (d *SHT4XDecoder) temperatureFromRaw(rawTemperature int) float64 {
	temperature := -45.0 + 175.0*(float64(rawTemperature)/65535.0)
	return temperature
}

func (d *SHT4XDecoder) humidityFromRaw(rawHumidity int) float64 {
	humidity := -6 + 125*float64(rawHumidity)/65535.0
	return humidity
}

func (d *SHT4XDecoder) SensorName() string {
	return "SHT4X"
}
