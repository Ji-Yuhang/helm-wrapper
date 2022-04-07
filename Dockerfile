# FROM centos:7
# FROM alpine
FROM registry.cn-beijing.aliyuncs.com/ijx-public/helm-cm-push:3.8.1
WORKDIR /

RUN apk --no-cache add ca-certificates

ENV GIN_MODE=release


COPY config-example.yaml  /config.yaml
COPY bin/helm-wrapper /helm-wrapper

CMD [ "/helm-wrapper" ]
