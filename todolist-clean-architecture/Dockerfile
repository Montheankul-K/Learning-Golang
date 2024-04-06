FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
ENV GOARCH=amd64
RUN go build \
    -ldflags "-X main.buildcommit=`git rev-parse --short HEAD` \
    -X main.buildtime=`date "+%Y-%m-%dT%H:%M:%S%Z:00"`" \
    -o app

FROM gcr.io/distroless/base-debian11
COPY --from=build /app/app /app
EXPOSE 8081
USER nonroot:nonroot
CMD ["/app"]