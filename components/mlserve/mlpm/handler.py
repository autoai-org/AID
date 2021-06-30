# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT
# coding:utf-8
import os
import traceback
import uuid
from csv import DictReader, DictWriter

from werkzeug.datastructures import ImmutableMultiDict
from werkzeug.utils import secure_filename

from mlpm.app import aidserver
from mlpm.response import json_resp
from mlpm.utility import str2bool


async def handle_post_solver_train_or_infer(request, upload_folder,
                                            request_type, target_folder):
    config = ImmutableMultiDict(await request.form)
    data = config.to_dict()
    results = {}
    req_files = await request.files
    if 'file' in req_files:
        uploaded_file = req_files['file']
        filename = secure_filename(uploaded_file.filename)
        print(filename)
        # make sure the UPLOAD_FOLDER exsits
        if not os.path.isdir(upload_folder):
            os.makedirs(upload_folder)
        file_abs_path = os.path.join(upload_folder, filename)
        await uploaded_file.save(file_abs_path)
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


async def handle_batch_infer_request(request, upload_folder, target_folder):
    req_files = await request.files
    if 'file' in req_files:
        uploaded_file = req_files['file']
        filename = secure_filename(uploaded_file.filename)
        if not os.path.isdir(upload_folder):
            os.makedirs(upload_folder)
        file_abs_path = os.path.join(upload_folder, filename)
        uploaded_file.save(file_abs_path)
    try:
        with open(file_abs_path, 'r') as file_obj:
            csv_dict = list(DictReader(file_obj))
            results = [aidserver.solver.infer(row) for row in csv_dict]
            # merging results
            output = []
            for index, each in enumerate(csv_dict):
                each.update(results[index])
                output.append(each)
        if not os.path.isdir(target_folder):
            os.makedirs(target_folder)
        file_identifier = str(uuid.uuid4())
        output_file_path = os.path.join(target_folder,
                                        file_identifier + ".csv")
        head = output[0].keys()
        with open(output_file_path, 'w') as file_obj:
            writer = DictWriter(file_obj, fieldnames=head)
            writer.writeheader()
            for each in output:
                writer.writerow(each)
        return json_resp({'filename': file_identifier + ".csv"}, status=200)
    except Exception as e:
        traceback.print_exc()
        return json_resp({"error": str(e), "code": "500"}, status=500)
