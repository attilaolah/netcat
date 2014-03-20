## NetCat: a hackable Heroku server ![Public Domain](https://pypip.in/license/intperm/badge.png)

NetCat is an HTTP server written in Go, hosted on Heroku, under the following
address:

* `https://netcat.herokuapp.com`

It is *very* similar to [`httpbin`][1], except it only implements a few of the
features, but combines them and exposes them on all URLs. 

It also adds a [PNG bomb](#png-bombs) feature.

[1]: http://httpbin.org


### Slowness

To make it slow, set the `t` parameter to the desired timeout in seconds, e.g.
the following will take 28 seconds to return a response:

* [`https://netcat.herokuapp.com/?t=28`][4]

[4]: https://netcat.herokuapp.com/?t=28

**NOTE:** the default Heroku timeout of 30 seconds still applies. If you
specify a timeout longer than that, you'll also get a 503 response from Heroku,
regardless of what you've set in the `c` parameter.

The `t` parameter applies to all paths.

### Status code

By default, it returns empty *200 OK* responses for all paths except `/loop`,
which defaults to *302 Found*.

To make it return a custom status code, set the `c` parameter to the desired
status code, e.g. the following will return a *418 I'm a teapot* response.

* [`https://netcat.herokuapp.com/?c=418`][5]

[5]: https://netcat.herokuapp.com/?c=418

The `c` parameter applies to all paths.

### Custom headers

All query parameters except the first `c`, `b` and `t` params will be set as
response headers. The following sets a cookie in the response:

* [`https://netcat.herokuapp.com/t=2&set-cookie=foo=bar`][2]

[2]: https://netcat.herokuapp.com/t=2&set-cookie=foo=bar

### Custom body

To set the response body to a custom payload, send the payload in the `b`
parameter.

The `b` parameter applies to all paths.

### Redirect loops

The path `/loop` causes an infinite redirect loop. The status code defaults to
302, but it can be overridden with the `c` parameter, just like for any other
URL. The following is a very slow redirect loop:

* [`https://netcat.herokuapp.com/loop?t=28`][3]

[3]: https://netcat.herokuapp.com/loop?t=28

### PNG bombs

The paths `/bomb.png` and `/bomb/**.png` (where `**` matches anything) serve a
special, 16384×16384 PNG image. It is all white and it is compressed to fit in
32k (32761 bytes).

This path will have a `Content-Type` header set to `image/png` by default, but
you can override it with the `content-type` param, just like any other header.
