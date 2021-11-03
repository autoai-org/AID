# Copyright (c) 2021 Xiaozhe Yao
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
# coding:utf-8
from mlpm.solver import Solver
from mlpm.metrics import Metrics
from mlpm.server import aidserver

class ExampleSolver(Solver):
    def __init__(self,toml_file=None):
        super().__init__(toml_file)
        self.predictor = ''
        self.acc = Metrics()

    def infer(self, data):
        print(data)
        return {'data': [0, 1, 2, 3]}

    def train(self, data):
        epochs = int(data['epochs'])
        for i in range(epochs):
            print(i)
        print(data)
        return data

solver = ExampleSolver()
aidserver.solver = solver
