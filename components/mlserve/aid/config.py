# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

#coding:utf-8
import os
from pathlib import Path


def GetHomeDir():
    return os.path.join(str(Path.home()), '.autoai', '.aid')


def GetLogDir():
    return os.path.join(GetHomDir(), 'logs')


def GetConfig():
    with open(os.path.join(GetHomeDir(), 'config.toml'), 'r') as f:
        config = toml.load(f)
        return config
