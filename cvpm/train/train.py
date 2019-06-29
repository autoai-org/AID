from threading import Thread

class TrainTask:
    def __init__(self, func, args):
        self._running = False
        self._func = func
        self._args = args

    def terminate(self):
        self._running = False

    def run(self):
        self._running = True
        self._func(*self._args)