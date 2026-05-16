# https://learn.adafruit.com/adafruit-pn532-rfid-nfc/python-circuitpython
#!/usr/bin/env python3
import busio, board, json, time, sys 
from digitalio import DigitalInOut 
from adafruit_pn532.spi import PN532_SPI
from datetime import datetime

#SPI Setup, pins setup, and initialize pn532
spi = busio.SPI(board.SCK, board.MOSI, board.MISO)
cs_pin = DigitalInOut(board.D5)
pn532 = PN532_SPI(spi, cs_pin)

#Running..
ic, ver, rev, support = pn532.firmware_version
print(f"Found chip with firmware version {ic}.{ver}.{rev} supporting {support}", file=sys.stderr)
pn532.SAMConfiguration()

print("Waiting for a card...", file=sys.stderr)

while True: 
    try:
        uid = pn532.read_passive_target(timeout=1.0) #Reading the NFC card
        if uid is not None:
            uid_str = "".join([f"{x:02x}" for x in uid]) #Converting uid into string
            print(json.dumps({
              "client_id": "NFCSensor",
              "device": "PN532",
              "NFC_data"  : uid_str,
              "timestamp_NFC" : datetime.now().isoformat()
            }), flush=True)
            time.sleep(2)
    except Exception as e:
        print(f"NFC Error: {e}", file=sys.stderr)
        time.sleep(1)
