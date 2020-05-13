#coding:utf-8
import os
import argparse
import sys
from mlpm.server import aidserver

os.chdir(sys.path[0])

from {{Package}}.{{Filename}} import {{Classname}}

solver = {{Classname}}()

aidserver.solver = solver