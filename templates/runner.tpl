#coding:utf-8
import os
import argparse
import sys
from mlpm.server import aidserver

os.chdir(sys.path[0])

parser = argparse.ArgumentParser(description='CVPM Runner')
parser.add_argument('port', type=int, help='Listening on Port')

args = parser.parse_args()

from {{Package}}.{{Filename}} import {{Classname}}

solver = {{Classname}}()

aidserver.solver = solver