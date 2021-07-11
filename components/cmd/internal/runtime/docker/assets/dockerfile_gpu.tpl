FROM nvidia/cuda:11.4.0-runtime-ubuntu20.04

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install -y \
        software-properties-common zlib1g-dev gcc cmake build-essential python3 python3-pip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY ./ /app

RUN {{PrePIP|safe}}

RUN pip install -r /app/requirements.txt

RUN {{Setup|safe}}

WORKDIR /app

ENTRYPOINT ["gunicorn"]

CMD ["runner_{{Solvername}}:aidserver","-b", "0.0.0.0:8080","-k","uvicorn.workers.UvicornWorker"]