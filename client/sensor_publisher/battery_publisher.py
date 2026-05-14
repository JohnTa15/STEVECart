#https://github.com/e71828/pi_ina226/tree/main
#!/usr/bin/env python3
import json, logging, sys
from datetime import datetime
from ina226 import INA226
from time import sleep


ina = INA226(busnum=1, log_level=logging.INFO) 
ina.configure()
ina.set_low_battery(12.0) # setting low battery alert at 12.0V
while True:
    try:
        print(f"Bus voltage: {ina.voltage():.3f} V", file=sys.stderr)
        
        print(json.dumps({
            "client_id": "BatterySensor",
            "device": "ina226",
            "battery_level": round(ina.voltage(), 2), # Calculate battery percentage
            "battery_current": round(ina.current(), 2),
            "is_charging": False,
            "timestamp_battery" : datetime.now().isoformat()
        }), flush=True)
    except Exception as e:
        print(f"Error reading INA226: {e}", file=sys.stderr)
        
    sleep(5)
