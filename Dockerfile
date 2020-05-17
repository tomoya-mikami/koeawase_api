# 一部ライブラリが1.13.11しか使えない
FROM golang:1.13.11

RUN mkdir /workdir
WORKDIR /workdir
ADD . /workdir
RUN go mod download
ENV FIRESTORE_EMULATOR_HOST=localhost:8812

CMD ["/usr/local/go/bin/go", "run", "/workdir/main.go", "/workdir/wire_gen.go", "/workdir/sample.go"]

