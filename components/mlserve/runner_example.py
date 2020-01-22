# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

from aid.solver import Solver
from aid.metrics import Metrics

class ExampleSolver(Solver):
    def __init__(self):
        self.predictor = ''
        self.acc = Metrics()
    def infer(self, data):
        return {'data': [0, 1, 2, 3]}
    def train(self, data):
        epochs = int(data['epochs'])
        for i in range(epochs):
            print(i)
        print(data)
        return data

sol = ExampleSolver()

sol.start(8080)
