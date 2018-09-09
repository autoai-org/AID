FROM library/golang

WORKDIR /go/src/cvpm

COPY cli/ .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

EXPOSE 10590

CMD [ "cvpm", "daemon", "run" ]