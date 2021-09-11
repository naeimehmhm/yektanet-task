FROM golang:alpine3.14 AS build

WORKDIR /service
COPY . ./
RUN apk add make binutils
RUN make

FROM scratch
WORKDIR /
COPY --from=build /service/bin/app /app
EXPOSE 8080
USER nonroot:nonroot
CMD ["/app", "/liveness", "/readiness"]
