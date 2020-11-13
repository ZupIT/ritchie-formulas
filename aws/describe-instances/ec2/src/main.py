# -*- coding: utf-8 -*-

import os

from services import describe

PARAMS = {
    "region": os.environ.get("REGION"),
    "access_key": os.environ.get("ACCESS_KEY"),
    "secret_key": os.environ.get("SECRET_ACCESS_KEY"),
    "instance_state": os.environ.get("INSTANCE_STATE"),
}

describe.handle(PARAMS)
