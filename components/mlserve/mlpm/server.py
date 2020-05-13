# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

# coding:utf-8
import os

from quart import request, send_from_directory
from werkzeug.datastructures import ImmutableMultiDict
from werkzeug.utils import secure_filename

from mlpm.app import aidserver
from mlpm.handler import (handle_batch_infer_request,
                          handle_post_solver_train_or_infer)
from mlpm.response import json_resp
from mlpm.utility import get_available_port, str2bool

UPLOAD_INFER_FOLDER = os.path.join("./", "temp", "infer")
UPLOAD_TRAIN_FOLDER = os.path.join("./", "temp", "train")
PUBLIC_FOLDER = os.path.join("./", "temp", "public")


@aidserver.route("/", methods=["GET"])
async def ping():
    return await json_resp({"status": "OK"}, status=200)


@aidserver.route("/infer", methods=["GET", "POST"])
async def infer():
    if request.method == 'POST':
        return await handle_post_solver_train_or_infer(request,
                                                       UPLOAD_INFER_FOLDER,
                                                       "infer", PUBLIC_FOLDER)


@aidserver.route("/train", methods=["GET", "POST"])
async def train():
    if request.method == "POST":
        return await handle_post_solver_train_or_infer(request,
                                                       UPLOAD_TRAIN_FOLDER,
                                                       "train", PUBLIC_FOLDER)


@aidserver.route("/batch", methods=["POST"])
async def batch_infer():
    if request.method == 'POST':
        return await handle_batch_infer_request(request, UPLOAD_INFER_FOLDER,
                                                PUBLIC_FOLDER)


@aidserver.route("/static/<filename>")
async def send_static(filename):
    return await send_from_directory(os.path.abspath(PUBLIC_FOLDER), filename)


def run_server(solver, port=None):
    if port is None:
        port = get_available_port()
    aidserver.solver = solver
    aidserver.run(host='0.0.0.0', port=port, debug=False)
