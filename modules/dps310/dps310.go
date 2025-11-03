package dps310

import "airsensors/modules"

type DPS310Decoder struct {
	calibration DPS310Calibration
}

func NewDecoder(calibration DPS310Calibration) *DPS310Decoder {
	return &DPS310Decoder{
		calibration: calibration,
	}
}

func (d *DPS310Decoder) Decode(raw modules.RawData) modules.Measurement {
	meas := modules.Measurement{}
	meas.Temperature = d.TemperatureFromRaw(raw.RawTemperature)
	meas.Pressure = d.PressureFromRaw(raw.RawPressure)
	return meas
}

func (d *DPS310Decoder) TemperatureFromRaw(rawTemperature uint) float64 {
	return 0.0
}

func (d *DPS310Decoder) PressureFromRaw(rawPressure uint) float64 {
	return 0.0
}

func (d *DPS310Decoder) SensorName() string {
	return "DPS310"
}
