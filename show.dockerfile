FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/show/main services/show/main.go

FROM iron/go
COPY --from=builder /app/services/show/main /app/services/show
EXPOSE 8091
ENTRYPOINT [ "/app/services/show" ]
