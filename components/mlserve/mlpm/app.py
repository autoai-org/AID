# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
#coding:utf-8

from quart import Quart

aidserver = Quart(__name__)
# Explicitly declare solver in the aidserver
aidserver.solver = None
