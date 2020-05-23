#coding:utf-8
import os
import argparse
import sys
from mlpm.server import aidserver

os.chdir(sys.path[0])

from {{Package}}.{{Filename}} import {{Classname}}

solver = {{Classname}}()

aidserver.solver = solver
if __name__=="__main__":
    aidserver.start("0.0.0.0", port=8080)
