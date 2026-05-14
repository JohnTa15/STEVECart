#https://blog.naver.com/roboholic84/220319850312
import RPi.GPIO as GPIO
from datetime import datetime
import json, time, sys


GPIO.setmode(GPIO.BCM)

trig_pin = 13
echo_pin = 19

GPIO.setup(trig_pin, GPIO.OUT)
GPIO.setup(echo_pin, GPIO.IN)

try:
    while True:
        GPIO.output(trig_pin, False)
        time.sleep(0.5)

        GPIO.output(trig_pin, True)
        time.sleep(0.00001)

        while GPIO.input(echo_pin) == 0:
            pulse_start = time.time()

        while GPIO.input(echo_pin) == 1:
            pulse_end = time.time()
        
        pulse_duration = pulse_end - pulse_start
        distance = round(pulse_duration * 17000, 2)

        print(f"Distance: {distance} cm", file=sys.stderr)
        
        print(json.dumps({
            "client_id": "UltrasonicSensor",
            "device": "HC-SR04",
            "distance": distance,
            "timestamp_distance": datetime.now().isoformat()
        }), flush=True)
        
        time.sleep(4.5) # Add sleep to avoid measuring 1000 times per second!
except Exception as e:
    print(f"Error: {e}", file=sys.stderr)
finally:
    GPIO.cleanup()