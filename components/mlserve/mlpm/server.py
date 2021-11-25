# coding:utf-8
import os

from quart import request, send_from_directory, current_app, jsonify
from quart.wrappers.request import Request
from quart_openapi.resource import Resource

from mlpm.app import aidserver
from mlpm.schemas import prediction_schema
from mlpm.handler import (
    handle_batch_infer_request,
    handle_post_solver_train_or_infer,
    handle_change
)
from mlpm.utility import get_available_port

UPLOAD_INFER_FOLDER = os.path.join("./", "temp", "infer")
UPLOAD_TRAIN_FOLDER = os.path.join("./", "temp", "train")
PUBLIC_FOLDER = os.path.join("./", "temp", "public")
@aidserver.route("/infer", methods=["GET", "POST"])
@aidserver.expect((prediction_schema, 'multipart/form-data'))
class Inference(Resource):
    async def post(self):
        return await handle_post_solver_train_or_infer(
            request,
            UPLOAD_INFER_FOLDER,
            "infer",
            PUBLIC_FOLDER
        )

@aidserver.route("/train", methods=["GET", "POST"])
async def train():
    if request.method == "POST":
        return await handle_post_solver_train_or_infer(
            request,
            UPLOAD_TRAIN_FOLDER,
            "train",
            PUBLIC_FOLDER
        )

@aidserver.route("/batch", methods=["POST"])
async def batch_infer():
    if request.method == 'POST':
        return await handle_batch_infer_request(request, UPLOAD_INFER_FOLDER,
                                                PUBLIC_FOLDER)

@aidserver.route("/change", methods=["POST"])
async def change():
    if request.method == "POST":
        return await handle_change(request)

@aidserver.route("/<path:path>", methods=["GET"])
async def view_static(path: str):
    try:
        return await current_app.send_static_file(path)
    except Exception as e:
        return await current_app.send_static_file("404.html")

@aidserver.route("/static/<filename>")
async def send_static(filename):
    return await send_from_directory(os.path.abspath(PUBLIC_FOLDER), filename)

@aidserver.route("/", methods=["GET"])
async def openapi():
  # add other logic if desired
  return jsonify(aidserver.__schema__)

def run_server(aidserver, port=None):
    if port is None:
        port = get_available_port()
    aidserver.run(host='0.0.0.0', port=port, debug=False)
