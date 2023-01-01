# resolve-ip
`resolve-ip` is a service that takes in an IP address and converts it to a latitude and longitude.
It includes GeoLite data created by MaxMind, available from 
[http://www.maxmind.com](http://www.maxmind.com).

## Endpoints

### `GET /healthcheck`

Returns a 200, used for load balancing healthchecks

### `GET /ip/{ip}`

If the IP address is found in the data set, it returns a 200 and a JSON body of the form:

``` json
{"lat":37.3394,"lon":-121.895}
```

If the IP address isn't found, it returns a 404.
test
