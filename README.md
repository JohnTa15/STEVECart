# S.T.E.V.E. – Smart Telemetry Edge Virtualized Engine

SmartCarts for supermarket tracking and management

> The project will be deployed on a Raspberry Pi 5

# Client/SmartCart Side
## The sensors that will be used:
- UWB Sensor (tag/anchors): MaUWB ESP32S3
- NFC: PN532
- Ultrasonic Sensor: HC-SR04
- Load Cell: HX711 Amplifier

## Other components: 
- Speakers
- Touchscreen Display
- SD card(for the OS)
- Inbuilt flashlight

# Notes
> The UWB Sensor is a microcontroller based sensor so it will be connected to the Raspberry Pi 5 via USB. In order to get this to work we need to install some anchors in the supermarket. The anchors will be connected to the main server.