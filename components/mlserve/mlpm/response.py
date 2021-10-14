# coding:utf-8
from quart import Response, jsonify


def json_resp(data, status):
    return jsonify(data), status
