from storage.database import DB
from celery import Celery
from schemas.common.event import Event
import ast 
import requests

queue = Celery('tasks', backend='redis://localhost:63791/1', broker='redis://localhost:63791/1')

@queue.task
def webhook_task(url, payload):
    print(payload)
    response = requests.post(url, json=payload)
    res = response.text
    print(res)
    return res

class Pipeline(object):
    def __init__(self):
        super(Pipeline, self).__init__()
    
    def run_event_through_all_extensions(self, event: Event):
        # first we get all extensions
        print(event)
        extensions = DB().fetch_extensions()
        for each in extensions:
            conditions = ast.literal_eval(each.condition)
            if event.event_type in conditions:
                # we run the extension
                payload = event.event_payload
                if each.extension_type == "WEBHOOK":
                    webhook_task.delay(each.entrypoint, ast.literal_eval(payload))
        return event