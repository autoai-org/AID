# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

import socket


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
