FROM golang:1.16-alpine
WORKDIR /app

COPY aid ./aid
EXPOSE 17415
RUN (cd /app && ls -a)
CMD ["./aid","up"]
