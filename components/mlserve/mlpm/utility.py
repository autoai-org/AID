# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
# coding:utf-8
import socket

from mlpm.response import json_resp


def str2bool(v):
    return str(v).lower() in ("true", "false", "yes", "t", "1")


def _isPortOpen(port):
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    result = s.connect_ex(('127.0.0.1', port))
    s.close()
    if result == 0:
        return True
    else:
        return False


def get_available_port(start=8080):
    port = start
    while True:
        if _isPortOpen(port):
            port = port + 1
        else:
            break
    return port


def process_unready_requests(solver):
    if not solver.is_ready:
        return json_resp(
            {
                "error": "Solver is not ready yet, please wait patiently...",
                "code": "500"
            },
            status=500)
