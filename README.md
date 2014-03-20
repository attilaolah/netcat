# SlowServer is a slow server. Nothing more. ![Public Domain](https://pypip.in/license/intperm/badge.png)

It is hosted on Heroku as `https://slowserver.herokuapp.com`.

By default, it returns empty 200 OK responses.

To make it slow, set the `t` parameter to the desired timeout in seconds, e.g.
`https://slowserver.herokuapp.com/?t=28` will take 28 seconds to return
nothing.

To make it return a custom status code, set the `c` parameter to the desired
status code, e.g. `https://slowserver.herokuapp.com/?t=28&c=418` will return a
`418 I'm a teapot` response.

**NOTE:** the default Heroku timeout of 30 seconds still applies. If you
specify a timeout longer than that, you'll also get a 503 response from Heroku,
regardless of what you've set in the `c` parameter.

## TODO

* custom headers
* custom payload
* TCP version that uses sockets and blocks on the protocol level
