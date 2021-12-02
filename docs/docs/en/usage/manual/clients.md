---
title: Clients
---

:::caution
These clients are not thoroughly tested yet. There might be problems.
:::

## cURL

You probably already have curl installed. You can try to interact with AID as below:

```sh
curl -X POST -F file=@test.jpg http://127.0.0.1:8080/infer
```

## Python Client

You can install the Python client by using `pip install mlpm_client`.

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

## NodeJS/JavaScript Client

```javascript title="Example of Javascript Client"
let httpc = new HTTPClient("http://127.0.0.1:17415/runnings/38a64faa/infer");
httpc.addPayload("text", this.state.textValue);
httpc.send().then(function (res) {
  // do something
});
```
