FROM scratch
COPY build/ca-certificates.crt /etc/ssl/certs/
COPY build/paladin /go/paladin

WORKDIR /go
VOLUME ["/go/database"]
CMD ["/go/paladin"]
