FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/cinema/main services/cinema/main.go

FROM iron/go
COPY --from=builder /app/services/cinema/main /app/services/cinema
EXPOSE 8090
ENTRYPOINT ["/app/services/cinema"]
