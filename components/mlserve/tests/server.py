# Copyright (c) 2021 Xiaozhe Yao

from mlpm.server import aidserver, run_server
from mlpm.solver import Solver


class ExampleSolver(Solver):
    def __init__(self, toml_file=None):
        super().__init__(toml_file)
        self.predictor = ""
        self.acc = [1, 2, 3, 4]

    def infer(self, data):
        return {"data": [0, 1, 2, 3]}

    def train(self, data):
        epochs = int(data["epochs"])
        for i in range(epochs):
            print(i)
        print(data)
        return data


solver = ExampleSolver()
aidserver.solver = solver
