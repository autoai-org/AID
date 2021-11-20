# Copyright (c) 2021 Xiaozhe Yao
#

#coding:utf-8

from quart import Quart

aidserver = Quart(__name__)
# Explicitly declare solver in the aidserver
aidserver.solver = None
