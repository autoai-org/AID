import re
import requests
from requests.api import head
from requests_toolbelt.multipart.encoder import MultipartEncoder

URL_REGEX = re.compile(
    r'^(?:http)s?://'  # http:// or https://
    r'(?:(?:[A-Z0-9](?:[A-Z0-9-]{0,61}[A-Z0-9])?\.)+(?:[A-Z]{2,6}\.?|[A-Z0-9-]{2,}\.?)|'  #domain...
    r'localhost|'  #localhost...
    r'\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})'  # ...or ip
    r'(?::\d+)?'  # optional port
    r'(?:/?|[/?]\S+)$',
    re.IGNORECASE)


class Request(object):
    def __init__(self, endpoint, payload) -> None:
        super().__init__()
        self.endpoint = endpoint
        self.payload = payload

    def do(self):
        if 'file' in self.payload:
            self.payload['file'] = (self.payload['file'].name,
                                    self.payload['file'], 'text/plain')
        multipart_data = MultipartEncoder(self.payload)
        return requests.post(
            self.endpoint,
            data=multipart_data,
            headers={'Content-Type': multipart_data.content_type})


class Client(object):
    def __init__(self, host: str) -> None:
        super().__init__()
        self.endpoints = {}
        if (re.match(URL_REGEX, host) is not None):
            self.host = host
        else:
            raise ValueError(
                "Malformed URL for the host, please check if the format is http[s]://example.com"
            )

    def flush(self):
        self.endpoints = {}

    def create_requests(self, vendor, package, solver, payload):
        if vendor + package + solver in self.endpoints:
            # already know the given endpoints
            endpoint = self.endpoints[vendor + package + solver]
        else:
            params = {'vendor': vendor, 'package': package, 'solver': solver}
            resp = requests.post(self.host + ":17415/api/preflight",
                                 json=params,
                                 headers={'Content-type': "application/json"})
            data = resp.json()
            endpoint = data['port']
            self.endpoints[vendor + package + solver] = endpoint
        endpoint = self.host + ":" + endpoint + "/infer"
        print(endpoint)
        return Request(endpoint, payload)


'''
The following main function works as an example.
'''
if __name__ == "__main__":
    c = Client('http://127.0.0.1')
    with open('test.jpg', 'rb') as file:
        r = c.create_requests('aidmodels', 'image_encoding', 'encodingSolver',
                              {'file': file})
        result = r.do()
        print(result.text)
