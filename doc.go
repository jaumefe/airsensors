package airquality

/*
Package designed to obtain air quality measurements from sensors.
It is not intended to directly interface this library with the sensors: no I2C support, no checksum calculation...
The inputs of these methods are the raw data from the sensors externally processed.

Compatible sensors:
- BME680
- DPS310
- SHT4x
*/
