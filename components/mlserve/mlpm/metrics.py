# Copyright (c) 2021 Xiaozhe Yao
#



# coding:utf-8
class Metrics(object):
    def __init__(self):
        self.data = []

    def add(self, value):
        # value is assumed to be a json object, rather than a plain text
        # e.g. {step: 0, acc:1}
        self.data.append(value)

    @property
    def output(self):
        return {}
