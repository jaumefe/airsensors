package sht4x

import (
	"testing"

	"github.com/jaumefe/airsensors/modules"
)

func TestNewDecoder(t *testing.T) {
	decoder := NewDecoder()
	if decoder == nil {
		t.Errorf("NewDecoder returned nil")
	}
}

func TestSHT4XDecoderDecode(t *testing.T) {
	decoder := NewDecoder()
	rawData := modules.RawData{
		RawTemperature: 25278,
		RawHumidity:    33159,
	}
	measurement := decoder.Decode(rawData)
	tolerance := 1e-5
	expectedTemp := 22.500572213321135
	expectedHumidity := 57.24673838406958
	if (measurement.Temperature < expectedTemp-tolerance) || (measurement.Temperature > expectedTemp+tolerance) {
		t.Errorf("Temperature out of expected range: got %f, expected %f", measurement.Temperature, expectedTemp)
	}
	if (measurement.Humidity < expectedHumidity-tolerance) || (measurement.Humidity > expectedHumidity+tolerance) {
		t.Errorf("Humidity out of expected range: got %f, expected %f", measurement.Humidity, expectedHumidity)
	}
}

func TestSensorName(t *testing.T) {
	decoder := NewDecoder()
	sensorName := decoder.SensorName()
	expectedName := "SHT4X"
	if sensorName != expectedName {
		t.Errorf("SensorName returned %s, expected %s", sensorName, expectedName)
	}
}
