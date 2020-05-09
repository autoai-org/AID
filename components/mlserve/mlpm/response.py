from quart import jsonify, Response

def json_resp(data, status):
    return jsonify(data), status