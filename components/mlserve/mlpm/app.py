#coding:utf-8

from quart_openapi import Pint

aidserver = Pint(
    __name__, 
    title="AID Solver API",
    no_openapi=True,
    static_folder='static'
)