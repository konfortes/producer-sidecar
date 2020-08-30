FROM golang:alpine AS builder
LABEL stage=builder
RUN apk add --no-cache gcc libc-dev
WORKDIR /workspace
# TODO: optimize
COPY . .
RUN go build -a

FROM alpine AS final
WORKDIR /
COPY --from=builder /workspace/tbd .
EXPOSE 3000
CMD [ "./tbd" ]