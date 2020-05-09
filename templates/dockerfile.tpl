FROM python:3.7-slim-buster

RUN apt-get update && \
    apt-get install -y \
        zlib1g-dev gcc cmake build-essential && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


COPY ./ /app

RUN pip install -r /app/requirements.txt

RUN {{Setup|safe}}

WORKDIR /app

ENTRYPOINT ["python3"]

CMD ["gunicorn", "runner_{{Solvername}}:aidserver", "-b" "127.0.0.1:8080","-k", "uvicorn.workers.UvicornWorker"]