FROM obraun/vss-protoactor-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o services/booking/main services/booking/main.go

FROM iron/go
COPY --from=builder /app/services/booking/main /app/services/booking
EXPOSE 8091
ENTRYPOINT [ "/app/services/booking/" ]
