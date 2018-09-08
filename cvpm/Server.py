import os
import json
import socket
import traceback
import logging
import gevent.pywsgi
from flask import g
from flask import Flask
from flask import request
from werkzeug.utils import secure_filename

logger = logging.getLogger()
logger.setLevel("INFO")
server = Flask(__name__)

ALLOWED_EXTENSIONS_TRAIN = set(['zip'])
ALLOWED_EXTENSIONS_INFER = set(['jpg', 'jpeg', 'png'])
UPLOAD_FOLDER = './temp'

server.config['UPLOAD_FOLDER'] = UPLOAD_FOLDER

def _isPortOpen(port):
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    result = s.connect_ex(('127.0.0.1', port))
    if result == 0:
        return True
    else:
        return False
    s.close()


def get_available_port(start=8080):
    port = start
    while True:
        if _isPortOpen(port):
            port = port + 1
        else:
            break
    return port

def allowed_file(filename, phase):
    ALLOWED_EXTENSIONS = ALLOWED_EXTENSIONS_TRAIN
    if phase == 'infer':
        ALLOWED_EXTENSIONS = ALLOWED_EXTENSIONS_INFER 
    return '.' in filename and \
        filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS


@server.route("/")
def help():
    help_message = server.solver.help_message
    return help_message


@server.route("/infer", methods=["GET", "POST"])
def infer():
    if request.method == 'POST':
        results = {}
        config = request.json
        print(request)
        if 'file' not in request.files:
            return json.dumps({"error": "no file part!", "code": "400"}), 400
        file = request.files['file']
        if file and allowed_file(file.filename, 'infer'):
            filename = secure_filename(file.filename)
            file_abs_path = os.path.join(server.config['UPLOAD_FOLDER'], filename)
            file.save(file_abs_path)
            try:
                results = server.solver.infer(file_abs_path, config)
                return json.dumps(results), 200
            except Exception as e:
                traceback.print_exc()
                return json.dumps({"error": str(e), "code":"500"}), 500
        else:
            return json.dumps({"error": "Forbidden Filename!", "code": "400"}), 400

@server.route("/train", methods=["GET", "POST"])
def train():
    if server.solver.enable_train:
        pass
    else:
        return json.dumps({"error": "not supported!", "code": "404"}), 404


def run_server(solver):
    port = get_available_port()
    logger.info("Server Running On: " + str(port))
    with server.app_context():
        server.solver = solver
    gevent_server = gevent.pywsgi.WSGIServer(('', port), server)
    gevent_server.serve_forever()
