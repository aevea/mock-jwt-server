FROM golang:1.21.4-alpine as builder
RUN mkdir /app
WORKDIR /app

RUN apk add --no-cache make git gcc musl-dev

# This step is done separately than `COPY . /app/` in order to
# cache dependencies.
COPY go.mod go.sum Makefile /app/
RUN go mod download

COPY . /app/
RUN make build/docker

FROM alpine:3.18.4

LABEL repository="https://github.com/aevea/commitsar"
LABEL homepage="https://github.com/aevea/commitsar"
LABEL maintainer="Simon Prochazka <simon@aevea.ee>"

LABEL com.github.actions.name="Mock JWT Server"
LABEL com.github.actions.description="JWT generator with JWK support"
LABEL com.github.actions.icon="code"
LABEL com.github.actions.color="blue"

RUN  apk add --no-cache --virtual=.run-deps ca-certificates git &&\
    mkdir /app

WORKDIR /app
COPY --from=builder /app/build/jwt-server ./jwt-server

RUN ln -s $PWD/jwt-server /usr/local/bin

CMD ["jwt-server"]
EXPOSE 8090
