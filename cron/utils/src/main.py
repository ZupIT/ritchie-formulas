#!/usr/bin/python3
import os
from formula import formula


action = os.environ.get("RIT_INPUT_ACTION")
job = os.environ.get("RIT_INPUT_JOB")
freq = os.environ.get("RIT_INPUT_FREQUENCY")
weekday = os.environ.get("RIT_INPUT_DAY_OF_WEEK")
month = os.environ.get("RIT_INPUT_DAY_OF_MONTH")
time = os.environ.get("RIT_INPUT_HOUR")

formula.Run(action, params=(job, freq, weekday, month, time))
