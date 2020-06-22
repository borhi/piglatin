FROM golang:1.14-alpine as build
WORKDIR /src/piglatin
COPY . .
RUN go build

FROM alpine:latest
COPY --from=build /src/piglatin .
CMD ["./piglatin", "-port", "3000"]
EXPOSE 3000
