# Copyright (c) 2021 Xiaozhe Yao
#

# coding:utf-8
import json
import os

from mlpm.config import GetConfig


class Logger(object):
    def __init__(self, config, target):
        self.target = target
        self.remote_log = config['remote_log']

    def write(self, logObject):
        with open(self.target, 'w+') as f:
            json.dump(logObject, f)

    def append(self, logObject):
        if os.stat(self.target).st_size == 0:
            self.write(logObject)
        else:
            with open(self.target, 'w+') as f:
                data = json.load(f)
                data.append(logObject)
                json.dump(data, f)
