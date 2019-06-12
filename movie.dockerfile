FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/movie/main services/movie/main.go

FROM iron/go
COPY --from=builder /app/services/movie/main /app/services/movie
EXPOSE 8090
ENTRYPOINT ["/app/services/movie"]
