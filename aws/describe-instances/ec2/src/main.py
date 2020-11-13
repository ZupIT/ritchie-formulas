# -*- coding: utf-8 -*-

import os

from formula import formula

region = os.environ.get("REGION")
access_key = os.environ.get("ACCESS_KEY")
secret_key = os.environ.get("SECRET_ACCESS_KEY")

formula.Run()
