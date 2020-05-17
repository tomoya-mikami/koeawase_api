# app engineが対象としているバージョンが1.11 or 1.12+
FROM golang:1.13.11

RUN mkdir /workdir
WORKDIR /workdir
ADD . /workdir
RUN go mod download
ENV FIRESTORE_EMULATOR_HOST=localhost:8812

CMD ["/usr/local/go/bin/go", "run", "/workdir/main.go"]

