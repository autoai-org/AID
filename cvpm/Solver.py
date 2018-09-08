import json

import requests
import toml
import tqdm

from cvpm.Server import run_server
from cvpm.Utility import BundleAnalyzer, Downloader


class Solver(object):
    def __init__(self, toml_file=None):
        self._isReady = False
        self.bundle = {}
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
        ba = BundleAnalyzer(self.bundle)
        return json.dumps(ba.load())

    def _prepare_models(self, toml_file):
        parsed_toml = toml.load(toml_file)
        downloader = Downloader()
        for each in parsed_toml["models"]:
            downloader.download(each["url"], "pretrained")

    def set_ready(self):
        self._isReady = True

    def infer(self, input, config):
        pass

    def train(self, train_x, train_y, **kwargs):
        pass

    def start(self):
        run_server(self)
