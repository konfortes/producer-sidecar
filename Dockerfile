FROM golang:alpine AS builder
LABEL stage=builder
RUN apk add --no-cache gcc libc-dev
WORKDIR /workspace
# TODO: optimize
COPY . .
RUN go build -a

FROM alpine AS final
WORKDIR /
RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
COPY --from=builder /workspace/tbd .
EXPOSE 3000
CMD [ "./tbd" ]