# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

from aid.solver import Solver


class ExampleSolver(Solver):
    def __init__(self):
        self.predictor = ''

    def infer(self, data):
        return {'data': [0, 1, 2, 3]}

sol = ExampleSolver()

sol.start(8080)
