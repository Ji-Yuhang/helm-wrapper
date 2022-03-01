# FROM centos:7
FROM alpine

RUN apk --no-cache add ca-certificates

ENV GIN_MODE=release

COPY config-example.yaml  /config.yaml
COPY bin/helm-wrapper /helm-wrapper

CMD [ "/helm-wrapper" ]
