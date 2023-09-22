FROM alpine:3.18 AS runner

RUN set -x; \
 adduser -u 1001 --disabled-password --home /app app

USER app
WORKDIR /app

COPY --chown=1001:1001 ./bin/srv .
COPY --chown=1001:1001 ./config/*.yaml ./config/

# Docker build arguments to store metadata information about
# the git branch, etc. into the Docker image.
ARG BUILD_DATE
ARG GIT_BRANCH
ARG GIT_TAG
ARG GIT_HASH
ARG SERVICE_NAME

LABEL git-hash=${GIT_HASH} \
      git-tag=${GIT_TAG} \
      build-date=${BUILD_DATE} \
      branch-name=${GIT_BRANCH} \
      appname=${SERVICE_NAME} \
      maintainer=sre@adgear.com

CMD ["/app/srv"]