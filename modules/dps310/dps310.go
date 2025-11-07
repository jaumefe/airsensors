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

func (d *DPS310Decoder) Decode(raw modules.RawData) (modules.Measurement, error) {
	meas := modules.Measurement{}
	meas.Temperature = d.temperatureFromRaw(raw.RawTemperature)
	meas.Pressure = d.pressureFromRaw(raw.RawPressure, raw.RawTemperature)
	return meas, nil
}

func (d *DPS310Decoder) temperatureFromRaw(rawTemperature uint) float64 {
	temperatureScaledRaw := (float64)(rawTemperature) / d.calibration.KT
	temperature := (float64)(d.calibration.C0)*0.5 + (float64)(d.calibration.C1)*(float64)(temperatureScaledRaw)
	return temperature
}

func (d *DPS310Decoder) pressureFromRaw(rawPressure uint, rawTemperature uint) float64 {
	temperatureScaledRaw := (float64)(rawTemperature) / d.calibration.KT
	pressureScaledRaw := (float64)(rawPressure) / d.calibration.KP
	pressure := (float64)(d.calibration.C00) + pressureScaledRaw*((float64)(d.calibration.C10)+pressureScaledRaw*((float64)(d.calibration.C20)+pressureScaledRaw*(float64)(d.calibration.C30))) + temperatureScaledRaw*(float64)(d.calibration.C01) + temperatureScaledRaw*pressureScaledRaw*((float64)(d.calibration.C11)+pressureScaledRaw*(float64)(d.calibration.C21))
	return pressure
}

func (d *DPS310Decoder) SensorName() string {
	return "DPS310"
}
