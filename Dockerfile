FROM gliderlabs/alpine:3.3

COPY GeoLiteCity /GeoLiteCity

RUN apk-install ca-certificates
COPY bin/resolve-ip /bin/resolve-ip

CMD ["/bin/resolve-ip", "--addr=0.0.0.0:80"]
