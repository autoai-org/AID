FROM python:3-slim-buster

RUN apt-get update && \
    apt-get install -y \
        zlib1g-dev gcc cmake build-essential libjpeg-dev libpng-dev && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*


COPY ./ /app

RUN pip install --upgrade \
        pip \
        -r /app/requirements.txt

WORKDIR /app

ENTRYPOINT ["python3"]

CMD ["runner_{{Solver_name}}.py", "8080"]