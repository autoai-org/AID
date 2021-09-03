from celery import Celery

app = Celery("main", broker="")

