# Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
#
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

from mlpm.server import aidserver

app = aidserver


def test_index_returns_200():
    request, response = app.test_client.get('/')
    assert response.status == 200


def test_index_put_not_allowed():
    request, response = app.test_client.put('/')
    assert response.status == 405


if __name__ == '__main__':
    test_index_returns_200()
    test_index_put_not_allowed()
