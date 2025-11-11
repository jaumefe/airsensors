package dps310

import (
	"testing"

	"github.com/jaumefe/airsensors/modules"
)

var calibrationTestData = DPS310Calibration{
	C0:  222,
	C1:  -292,
	C00: 79161,
	C10: -52411,
	C01: -4085,
	C11: 1389,
	C20: -9454,
	C21: -33,
	C30: -1214,
	KT:  1572864.0,
	KP:  1572864.0,
}

func TestNewDecoder(t *testing.T) {
	decoder := NewDecoder(calibrationTestData)
	if decoder == nil {
		t.Errorf("NewDecoder returned nil")
	}
}

func TestDPS310DecoderDecode(t *testing.T) {
	decoder := NewDecoder(calibrationTestData)
	rawData := modules.RawData{
		RawTemperature: 635464,
		RawPressure:    -778518,
	}
	measurement := decoder.Decode(rawData)
	tolerance := 1e-5
	expectedTemp := -6.973002115885421
	expectedPressure := 101002.39087031863
	if (measurement.Temperature < expectedTemp-tolerance) || (measurement.Temperature > expectedTemp+tolerance) {
		t.Errorf("Temperature out of expected range: got %f, expected %f", measurement.Temperature, expectedTemp)
	}
	if (measurement.Pressure < expectedPressure-tolerance) || (measurement.Pressure > expectedPressure+tolerance) {
		t.Errorf("Pressure out of expected range: got %f, expected %f", measurement.Pressure, expectedPressure)
	}
}

func TestSensorName(t *testing.T) {
	decoder := NewDecoder(calibrationTestData)
	sensorName := decoder.SensorName()
	expectedName := "DPS310"
	if sensorName != expectedName {
		t.Errorf("SensorName returned %s, expected %s", sensorName, expectedName)
	}
}
