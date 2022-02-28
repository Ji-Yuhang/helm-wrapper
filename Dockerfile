# FROM centos:7
FROM scratch

ENV GIN_MODE=release

COPY config-example.yaml  /config.yaml
COPY bin/helm-wrapper /helm-wrapper

CMD [ "/helm-wrapper" ]
