FROM golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o main ./cmd


FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE 3000
CMD ["/app/main"]
