FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/server ./

FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/server /app/server
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app/server"]
