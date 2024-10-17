FROM golang:1.22.5-alpine3.19

WORKDIR /test

COPY . /test

RUN go build /test

EXPOSE 8000

ENTRYPOINT [ "./go-todo" ]