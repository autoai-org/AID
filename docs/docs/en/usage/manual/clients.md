---
title: Clients
---

## Python Client

You can install the Python client by using ```pip install mlpm_client```.

An example of using the Python client to interact with AID is illustrated below:

```python title="Example of Python Client"
from mlpm_client import Client
c = Client('http://127.0.0.1')
with open('test.jpg', 'rb') as file:
    r = c.create_requests('aidmodels', 'image_encoding', 'encodingSolver',
                          {'file': file})
    result = r.do()
    print(result.text)
```