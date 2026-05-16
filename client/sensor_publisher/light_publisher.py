#https://github.com/adafruit/Adafruit_CircuitPython_TSL2591/blob/main/examples/tsl2591_simpletest.py
#!/usr/bin/env python3
import board, json, time, sys
from datetime import datetime
import adafruit_tsl2591

# Initialize I2C bus (board.I2C is standard for Adafruit)
i2c = board.I2C()

# Initialize the sensor
sensor = adafruit_tsl2591.TSL2591(i2c)

while True:
    try:
        lux = sensor.lux
        
        print(f"Light level: {lux} lux", file=sys.stderr)
        
        print(json.dumps({
            "client_id": "LightSensor",
            "device": "TSL2591",
            "lux": round(lux, 2), 
            "timestamp_light": datetime.now().isoformat(),
        }), flush=True)
    except Exception as e:
        print(f"Error reading TSL2591: {e}", file=sys.stderr)
        
    time.sleep(5) 
