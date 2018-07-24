FROM golang:1.10-stretch AS build
ENV DEBIAN_FRONTEND=noninteractive
ENV CGO_ENABLED=0
RUN apt-get update && apt-get install -y gcc git make
RUN go get golang.org/x/vgo
RUN mkdir -p /go/src/github.com/ripta/recurse-world
WORKDIR /go/src/github.com/ripta/recurse-world
COPY . .
RUN vgo install .

FROM debian:stretch-slim
COPY --from=build /go/bin/recurse-world /usr/bin/recurse-world
EXPOSE 8080
CMD ["/usr/bin/recurse-world"]
