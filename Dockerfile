FROM golang:1.27.1 AS build
WORKDIR /app
COPY ["algorithms/Linear Regression/", "./"]
RUN CGO_ENABLED=0 go build -trimpath -o /pykitzoid-api ./cmd/api

FROM scratch
COPY --from=build /pykitzoid-api /pykitzoid-api
USER 65532:65532
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/pykitzoid-api"]
