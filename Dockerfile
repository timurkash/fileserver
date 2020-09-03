FROM golang:1.15 as builder

MAINTAINER timurkash@yandex.ru

RUN mkdir -p /images/

WORKDIR /images

COPY . .

RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
    CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags "-s -w \
    -X gitlab.mcsolutions.ru/find-psy/back/users/images/pkg/version.REVISION=${GIT_COMMIT}" \
    -a -o bin/images cmd/images/*

FROM alpine:latest

RUN addgroup -S app \
    && adduser -S -g app app \
    && apk --no-cache add \
    curl openssl netcat-openbsd mc

WORKDIR /home/app

COPY --from=builder /images/bin/images .
COPY --from=builder /images/docs/swagger.yaml ./docs/swagger.yaml

RUN chown -R app:app ./

USER app

EXPOSE 3027

CMD ["./images"]
