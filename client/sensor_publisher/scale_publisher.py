#https://github.com/gandalf15/HX711
#!/usr/bin/env python3
import json, time, sys
from datetime import datetime
import RPi.GPIO as GPIO
from hx711 import HX711

try:
    GPIO.setmode(GPIO.BCM) #setting the mode to BCM
    hx = HX711(dout_pin=21, pd_sck_pin=20) 
    
    print("Tare checking.. please empty the cart..", file=sys.stderr)
    err = hx.zero() #resetting the scale to zero
    if err: 
        raise ValueError("Tare is unsuccessful")
    print("Tare is ready!", file=sys.stderr)

except Exception as e:
    print(f"Tare error: {e}", file=sys.stderr)
    sys.exit(1)

while True:
    try:
        weight = hx.get_weight_mean(20)
        
        print(f"Weight: {weight} kg", file=sys.stderr)
        
        print(json.dumps({
            "client_id": "ScaleSensor",
            "device": "HX711",
            "weight": round(weight, 3), 
            "timestamp_weight": datetime.now().isoformat()
        }), flush=True)
    except Exception as e:
        print(f"Error reading HX711: {e}", file=sys.stderr)
        
    time.sleep(1)
