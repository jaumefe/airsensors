package bme680

import (
	"testing"

	"github.com/jaumefe/airsensors/modules"
)

var calibrationTestData = BME680Cablibration{
	ParT1: 25895,
	ParT2: 26617,
	ParT3: 3,

	ParP1:  37029,
	ParP2:  -10303,
	ParP3:  88,
	ParP4:  10168,
	ParP5:  -276,
	ParP6:  30,
	ParP7:  51,
	ParP8:  -3740,
	ParP9:  -2149,
	ParP10: 30,

	ParH1: 775,
	ParH2: 777,
	ParH3: 0,
	ParH4: 45,
	ParH5: 20,
	ParH6: 120,
	ParH7: -100,

	GasRange:      11,
	GasRangeSWErr: 10,
}

func TestNewDecoder(t *testing.T) {
	decoder := NewDecoder(calibrationTestData)
	if decoder == nil {
		t.Errorf("NewDecoder returned nil")
	}
}

func TestBME680DecoderDecode(t *testing.T) {
	decoder := NewDecoder(calibrationTestData)
	rawData := modules.RawData{
		RawTemperature:   485504,
		RawPressure:      282686,
		RawHumidity:      25196,
		RawGasResistance: 172,
	}
	measurement := decoder.Decode(rawData)
	tolerance := 1e-5
	expectedTemp := 22.589403203269466
	expectedHumidity := 50.972661917814946
	expectedPressure := 97477.57943543159
	expectedGasRes := 5179.55904558746
	if (measurement.Temperature < expectedTemp-tolerance) || (measurement.Temperature > expectedTemp+tolerance) {
		t.Errorf("Temperature out of expected range: got %f, expected %f", measurement.Temperature, expectedTemp)
	}
	if (measurement.Humidity < expectedHumidity-tolerance) || (measurement.Humidity > expectedHumidity+tolerance) {
		t.Errorf("Humidity out of expected range: got %f, expected %f", measurement.Humidity, expectedHumidity)
	}
	if (measurement.Pressure < expectedPressure-tolerance) || (measurement.Pressure > expectedPressure+tolerance) {
		t.Errorf("Pressure out of expected range: got %f, expected %f", measurement.Pressure, expectedPressure)
	}
	if (measurement.GasResistance < expectedGasRes-tolerance) || (measurement.GasResistance > expectedGasRes+tolerance) {
		t.Errorf("GasResistance out of expected range: got %f, expected %f", measurement.GasResistance, expectedGasRes)
	}
}

func TestSensorName(t *testing.T) {
	decoder := NewDecoder(calibrationTestData)
	sensorName := decoder.SensorName()
	expectedName := "BME680"
	if sensorName != expectedName {
		t.Errorf("SensorName returned %s, expected %s", sensorName, expectedName)
	}
}
