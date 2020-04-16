# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

#coding:utf-8

from mlpm.app import aidserver
from mlpm.response import json_resp
from mlpm.handler import handle_post_solver_train_or_infer
from mlpm.utility import get_available_port, str2bool
from werkzeug.datastructures import ImmutableMultiDict
from werkzeug.utils import secure_filename
from flask import request

UPLOAD_INFER_FOLDER = './temp/infer'
UPLOAD_TRAIN_FOLDER = './temp/train'

@aidserver.route("/", methods=["GET"])
def ping():
    return json_resp({"status":"OK"}, status=200)


@aidserver.route("/infer", methods=["GET", "POST"])
def infer():
    if request.method == 'POST':
        return handle_post_solver_train_or_infer(request, UPLOAD_INFER_FOLDER,
                                          "infer")


@aidserver.route("/train", methods=["GET", "POST"])
def train():
    if request.method == "POST":
        return handle_post_solver_train_or_infer(request, UPLOAD_TRAIN_FOLDER,
                                          "train")


def run_server(solver, port=None):
    if port is None:
        port = get_available_port()
    aidserver.solver = solver
    aidserver.run(host='0.0.0.0', port=port, debug=False)
