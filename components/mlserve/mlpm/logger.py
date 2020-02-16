# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
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

class StepLogger(object):
    def __init__(self, logFile):
        self.logFile = logFile
    
    def append(self, logObject):
        if os.stat(self.target).st_size == 0:
            self.write(logObject)
        else:
            with open(self.target, 'w+') as f:
                data = json.load(f)
                data.append(logObject)
                json.dump(data, f)
