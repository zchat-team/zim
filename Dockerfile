FROM alpine:3.16

WORKDIR /
COPY bin .
COPY manifest/scripts/run.sh .
COPY manifest/config config
ENTRYPOINT ["/run.sh"]
