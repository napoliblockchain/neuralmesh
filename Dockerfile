FROM golang:1.22-alpine AS build

WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN go build -o /out/neuralmesh-demo ./cmd/demo

FROM alpine:3.20

WORKDIR /app
COPY --from=build /out/neuralmesh-demo /usr/local/bin/neuralmesh-demo
CMD ["neuralmesh-demo"]
