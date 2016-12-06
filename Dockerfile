FROM alpine:edge

ADD bin /capserv

ENTRYPOINT /capserv/capserv

EXPOSE 8888

