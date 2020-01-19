# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

#coding:utf-8
import os
import traceback

from sanic import Sanic
from sanic.response import json
from sanic import response
from sanic import request

from aid.utility import get_available_port, str2bool
from werkzeug.datastructures import ImmutableMultiDict
from werkzeug.utils import secure_filename

aidserver = Sanic()

UPLOAD_FOLDER = './temp'

aidserver.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER

@aidserver.route("/", methods=["GET"])
async def ping(request):
    return response.text('Hello world!', status=200)

@aidserver.route("/infer", methods=["GET", "POST"])
async def infer(request):
    if request.method =='POST':
        results = {}
        config = ImmutableMultiDict(request.form)
        config = config.to_dict()
        data = config
        if 'file' in request.files:
            file = request.files['file']
            filename = secure_filename(file.filename)
            # make sure the UPLOAD_FOLDER exsits
            if not os.path.isdir(aidserver.config['UPLOAD_FOLDER']):
                os.makedirs(aidserver.config['UPLOAD_FOLDER'])
            file_abs_path = os.path.join(aidserver.config['UPLOAD_FOLDER'],
                                         filename)
            file.save(file_abs_path)
            data['input_file_path'] = file_abs_path
        try:
            results = aidserver.solver.infer(data)
            if 'delete_after_process' in config:
                if str2bool(config['delete_after_process']):
                    os.remove(file_abs_path)
            return response.json(results, status=200)
        except Exception as e:
            traceback.print_exc()
            return response.json({"error": str(e), "code": "500"}, status=500)

def run_server(solver, port=None):
    if port is None:
        port = get_available_port()
    aidserver.solver = solver
    aidserver.run(host='0.0.0.0', port=port, workers=4)
