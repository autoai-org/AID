import os
import toml
from pathlib import Path


def getConfig():
    configFilePath = os.path.join(str(Path.home()), 'cvpm', 'config.toml')
    with open(configFilePath, 'r') as f:
        config = toml.load(f)
        return config


def getLogDir():
    return os.path.join(getConfig()['local']['LocalFolder'], 'logs')


if __name__ == '__main__':
    getConfig()
