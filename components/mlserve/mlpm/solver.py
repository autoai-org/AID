# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

import os

from mlpm.logger import StepLogger
from mlpm.server import run_server


class Solver(object):
    def __init__(self, pretrained_toml=None):
        self._isReady = False
        self.bundle = None
        self._hyperparameters = {}
        self._enable_train = False
        if hasattr(self.__class__, 'train') and callable(getattr(self.__class__, 'train')):
            self._enable_train = True

    @property
    def enable_train(self):
        return self._enable_train

    @property
    def is_ready(self):
        return self._isReady

    @property
    def hyperparamters(self):
        return self._hyperparameters

    def ready(self):
        self._isReady = True

    def start(self, port=None):
        print('Server will run on port: ' + str(port))
        run_server(self, port)

    def infer(self, data):
        raise NotImplementedError

    def train(self, data):
        self.logger = StepLogger(os.environ('LOG_FILE'))
