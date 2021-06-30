FROM python:3.9-slim-buster

RUN apt-get update && \
    apt-get install -y \
        zlib1g-dev gcc cmake build-essential && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


COPY ./ /app

RUN {{PrePIP|safe}}

RUN pip install -r /app/requirements.txt

RUN {{Setup|safe}}

WORKDIR /app

ENTRYPOINT ["gunicorn"]

CMD ["runner_{{Solvername}}:aidserver","-b", "0.0.0.0:8080","-k","uvicorn.workers.UvicornWorker"]