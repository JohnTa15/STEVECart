# S.T.E.V.E. – Smart Telemetry Edge Virtualized Engine

SmartCarts for supermarket tracking and management

> The project will be deployed on a Raspberry Pi 5

# Client/SmartCart Side
## The sensors that will be used:
- UWB Sensor (tag/anchors): MaUWB ESP32S3
- NFC: PN532
- Ultrasonic Sensor: HC-SR04
- Load Cell: HX711 Amplifier
- AI camera: OAK-D Lite

## Other components: 
- Speakers
- Touchscreen Display
- SD card(for the OS)
- Inbuilt flashlight

# Notes
> The UWB Sensor is a microcontroller based sensor so it will be connected to the Raspberry Pi 5 via USB. In order to get this to work we need to install some anchors in the supermarket. The anchors will be connected to the main server.


# Running process steps
1. In strong computer the CVAT must be installed and setup in order to train the AI model with annotated images(OAK-D-LITE train). 


# CVAT tutorial
- create superadmin
- create project -> add product_labels -> add tasks (with options picture quality 100, segment size 60, validation mode ground truth)
- do the annotation process
- click project -> actions -> export dataset -> Ultralytics YOLO Detection 1.0, click save images
