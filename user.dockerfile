FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/user/main services/user/main.go

FROM iron/go
COPY --from=builder /app/services/user/main /app/services/user
EXPOSE 8091
ENTRYPOINT [ "/app/services/user" ]
