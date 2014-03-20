# SlowServer: a very slow server ![Public Domain](https://pypip.in/license/intperm/badge.png)

It is hosted on Heroku as `https://slowserver.herokuapp.com`.

## Slowness

To make it slow, set the `t` parameter to the desired timeout in seconds, e.g.
`https://slowserver.herokuapp.com/?t=28` will take 28 seconds to return a
response.


**NOTE:** the default Heroku timeout of 30 seconds still applies. If you
specify a timeout longer than that, you'll also get a 503 response from Heroku,
regardless of what you've set in the `c` parameter.

## Status code

By default, it returns empty **200 OK** responses for all paths except `/loop`,
which defaults to **302 Found**.

To make it return a custom status code, set the `c` parameter to the desired
status code, e.g. `https://slowserver.herokuapp.com/?t=28&c=418` will return a
`418 I'm a teapot` response.

## Custom headers

All query parameters except the first `c` and `t` params will be set as
response headers. The following sets a cookie in the response:

* [`https://slowserver.herokuapp.com/t=2&set-cookie=foo=bar`][2]

[2]: https://slowserver.herokuapp.com/t=2&set-cookie=foo=bar

## Redirect loops

The path `/loop` causes an infinite redirect loop. The status code defaults to
302, but it can be overridden with the `c` parameter, just like for any other
URL. All other parameters (`c` and `t`) are also valid. The following is a very
slow redirect loop:

* [`https://slowserver.herokuapp.com/loop?t=28`][1]

[1]: https://slowserver.herokuapp.com/loop?t=28

## TODO

* custom payload
* redirect loop countown, e.g. `/loop/2` → `loop/1` → `/`
* TCP version that uses sockets and blocks on the protocol level
