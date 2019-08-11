FROM golang:1.12-stretch AS build
ENV DEBIAN_FRONTEND=noninteractive
ENV CGO_ENABLED=0
RUN apt-get update && apt-get install -y gcc git make
RUN mkdir -p /build
WORKDIR /build
COPY . .
RUN go build .

FROM debian:stretch-slim
COPY --from=build /build/recurse-world /usr/bin/recurse-world
EXPOSE 8080
CMD ["/usr/bin/recurse-world"]
