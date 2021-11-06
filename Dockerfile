FROM golang:1.16 AS build

COPY . /go/src/
WORKDIR /go/src/
RUN make test
RUN make build

FROM centos:centos7
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" >> /etc/timezone

COPY --from=build /go/src/build /work/rpc-article
WORKDIR /work/rpc-article
CMD [ "bin/rpc-article", "-c", "config/base.json" ]
