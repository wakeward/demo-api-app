FROM golang:1.24.3-bullseye AS build

WORKDIR /

COPY . .
RUN go build -o demo-api -ldflags='-w -s' .

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /demo-api .

EXPOSE 8080

CMD ["./demo-api"]