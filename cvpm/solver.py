import json

import toml

from cvpm.bundle import Bundle
from cvpm.server import run_server
from cvpm.utility import Downloader


class Solver(object):
    def __init__(self, toml_file=None):
        self._isReady = False
        self.bundle = None
        self._enable_train = False
        if toml_file is None:
            toml_file = "./pretrained/pretrained.toml"
        self._prepare_models(toml_file)

    @property
    def enable_train(self):
        return self._enable_train

    @property
    def is_ready(self):
        return self._isReady

    @property
    def help_message(self):
        if self.is_ready:
            members = self.bundle.members()
            return json.dumps(members)
        else:
            return json.dumps({"error": "Failed to start", "code": "101"}), 101

    def _prepare_models(self, toml_file):
        parsed_toml = toml.load(toml_file)
        downloader = Downloader()
        if "models" in parsed_toml.keys():
            for each in parsed_toml["models"]:
                downloader.download(each["url"], "pretrained")

    def set_ready(self):
        self._isReady = True

    def set_bundle(self, bundle):
        if issubclass(bundle, Bundle):
            self.bundle = bundle
            solver = self
            bundle.add_solver(self=bundle, solver=solver)

    def infer(self, input, config):
        pass

    def train(self, train_x, train_y, **kwargs):
        pass

    def start(self, port=None):
        print('Server will run on port: ' + str(port))
        run_server(self, port)
