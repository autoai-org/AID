# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD & AutoAI.org
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
import os
import json

class Logger(object):
    def __init__(self, target):
        self.target = target
    
    def write(self, jsonObject):
        with open(self.target, 'w+') as f:
            json.dump([jsonObject], f)
    
    def append(self, jsonObject):
        if os.stat(self.target).st_size == 0:
            self.write(jsonObject)
        else:
            with open(self.target, 'w+') as f:
                data = json.load(f)
                data.append(jsonObject)
                json.dump(data, f)