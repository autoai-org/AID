import toml
class Pipeline(object):
    def __init__(self, name):
        self.name = name
    def load(self, filepath):
        