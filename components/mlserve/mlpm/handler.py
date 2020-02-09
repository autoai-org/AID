# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
# coding:utf-8
import os
import traceback
from sanic import Sanic, request, response
from sanic.response import json
from werkzeug.datastructures import ImmutableMultiDict
from werkzeug.utils import secure_filename

from mlpm.app import aidserver

def handle_post_solver_train_or_infer(request, upload_folder, request_type):
    config = ImmutableMultiDict(request.form)
    data = config.to_dict()
    results = {}
    if 'file' in request.files:
        file = request.files['file']
        filename = secure_filename(file.filename)
        # make sure the UPLOAD_FOLDER exsits
        if not os.path.isdir(upload_folder):
            os.makedirs(upload_folder)
        file_abs_path = os.path.join(upload_folder, filename)
        file.save(file_abs_path)
        data['input_file_path'] = file_abs_path
    try:
        if request_type == "infer":
            results = aidserver.solver.infer(data)
        elif request_type == "train":
            results = aidserver.solver.train(data)
        else:
            raise NotImplementedError
        if 'delete_after_process' in data:
            if str2bool(data['delete_after_process']):
                os.remove(file_abs_path)
        print(results)
        return response.json(results, status=200)
    except Exception as e:
        traceback.print_exc()
        return response.json({"error": str(e), "code": "500"}, status=500)
