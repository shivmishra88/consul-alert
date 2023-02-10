FROM alpine:edge

ENV GOPATH /go

RUN mkdir -p /go && \
    apk update && \
    apk add bash ca-certificates git go alpine-sdk && \
    go install github.com/AcalephStorage/consul-alerts@latest && \
    mv $GOPATH/bin/consul-alerts /bin && \
    apk upgrade --update --no-cache && \
    apk add --update --no-cache curl util-linux && \
    wget -O /tmp/consul.zip https://releases.hashicorp.com/consul/1.14.4/consul_1.14.4_linux_amd64.zip && \
    unzip -d /bin /tmp/consul.zip && \
    rm /tmp/consul.zip && \
    rm -rf /go && \
    apk del --purge go git alpine-sdk && \
    rm -rf /var/cache/apk/*

EXPOSE 9000
CMD []
ENTRYPOINT [ "/bin/consul-alerts", "--alert-addr=0.0.0.0:9000" ]
