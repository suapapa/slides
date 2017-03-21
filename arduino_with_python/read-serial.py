#!/usr/bin/python

import os
import sys
import serial

try:
    serPort = sys.argv[1]
    serSpeed = sys.argv[2]
except:
    print "Usage: %s port speed"%sys.argv[0]
    sys.exit(1)

ser = serial.Serial(serPort, serSpeed)
for line in ser:
    print line.strip()

ser.close()
