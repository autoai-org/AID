# Copyright (c) 2021 Xiaozhe Yao
#

# coding:utf-8
from mlpm.server import aidserver, run_server


class Solver(object):
    def __init__(self, pretrained_toml=None):
        self._isReady = False
        self.bundle = None
        self._hyperparameters = {}
        self._enable_train = False
        self.server = aidserver
        if hasattr(self.__class__, 'train') and callable(
                getattr(self.__class__, 'train')):
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

    def change(self):
        raise NotImplementedError