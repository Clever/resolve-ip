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

## Generating the Server

- Update swagger.yml with your endpoints

- Run `make generate`

- Write a controller that implements the generated Wag interface: gen-go/server/interface.go

- Implement your main function. For example:
```
package main

import (
	"flag"
	"log"

	"github.com/Clever/resolve-ip/gen-go/server"
)

func main() {
	addr := flag.String("addr", ":8080", "Address to listen at")
	flag.Parse()

	myController := MyController{} // Implements the generated interface
	s := server.New(myController, *addr)
	// Serve should not return
	log.Fatal(s.Serve())
}
```

- Create your glide.yaml and run `make install_deps`. See [the dev handbook](https://github.com/Clever/dev-handbook/blob/master/golang/glide.md) for help on how to use Glide.

- Add your environment variables to github.com/Clever/ark-config (LIGHTSTEP_ACCESS_TOKEN plus any you added)

For more details on Wag check out https://github.com/Clever/wag#usage
