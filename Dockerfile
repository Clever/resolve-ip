FROM debian:jessie

COPY bin/resolve-ip /usr/bin/resolve-ip
COPY GeoLiteCity /usr/bin/GeoLiteCity

WORKDIR /usr/bin

CMD ["/usr/bin/resolve-ip"]
