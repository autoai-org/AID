# Copyright (c) 2021 Xiaozhe Yao
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
#coding:utf-8

from quart import Quart

aidserver = Quart(__name__)
# Explicitly declare solver in the aidserver
aidserver.solver = None
