# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

#coding:utf-8

from mlpm.app import aidserver
from mlpm.handler import handle_post_solver_train_or_infer
from mlpm.utility import get_available_port, str2bool
from sanic import request, response
from sanic.response import json
from werkzeug.datastructures import ImmutableMultiDict
from werkzeug.utils import secure_filename

UPLOAD_INFER_FOLDER = './temp/infer'
UPLOAD_TRAIN_FOLDER = './temp/train'


@aidserver.route("/", methods=["GET"])
async def ping(request):
    return response.text('Hello world!', status=200)


@aidserver.route("/infer", methods=["GET", "POST"])
async def infer(request):
    if request.method == 'POST':
        return handle_post_solver_train_or_infer(request, UPLOAD_INFER_FOLDER,
                                          "infer")


@aidserver.route("/train", methods=["GET", "POST"])
async def train(request):
    if request.method == "POST":
        return handle_post_solver_train_or_infer(request, UPLOAD_TRAIN_FOLDER,
                                          "train")


def run_server(solver, port=None):
    if port is None:
        port = get_available_port()
    aidserver.solver = solver
    aidserver.run(host='0.0.0.0', port=port, workers=4)
