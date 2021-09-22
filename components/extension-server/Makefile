start:
	uvicorn rest.main:app --reload --port 17416
celery:
	celery -A runner.pipeline worker --loglevel=INFO