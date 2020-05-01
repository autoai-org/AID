# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
# coding:utf-8
import os
import traceback

from mlpm.app import aidserver
from mlpm.utility import str2bool
from mlpm.response import json_resp
from werkzeug.datastructures import ImmutableMultiDict
from werkzeug.utils import secure_filename


def handle_post_solver_train_or_infer(request, upload_folder, request_type):
    config = ImmutableMultiDict(request.form)
    data = config.to_dict()
    results = {}
    if 'file' in request.files:
        uploaded_file = request.files['file']
        filename = secure_filename(uploaded_file.filename)
        # make sure the UPLOAD_FOLDER exsits
        if not os.path.isdir(upload_folder):
            os.makedirs(upload_folder)
        file_abs_path = os.path.join(upload_folder, filename)
        uploaded_file.save(file_abs_path)
        data['input_file_path'] = file_abs_path
    try:
        if request_type == "infer":
            results = aidserver.solver.infer(data)
        else:
            raise NotImplementedError
        if 'delete_after_process' in data:
            if str2bool(data['delete_after_process']):
                os.remove(file_abs_path)
        print(results)
        return json_resp(results, status=200)
    except Exception as e:
        traceback.print_exc()
        return json_resp({"error": str(e), "code": "500"}, status=500)
