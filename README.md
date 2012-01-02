gofinger
========

A minimal finger client and server in go.  Targets the original and simple [RFC 742](http://tools.ietf.org/html/rfc742).

Building
========

1. Install [go](http://golang.org/) if you haven't already.
2. Check out the gofinger source or download a tarball.
3. Build from the main directory: `gomake all`.

Running
=======

To run the server, run `sudo src/cmd/fingerd/fingerd`. Sudo is required because the finger protocol runs on a privileged port. By
default it listens on `127.0.0.1`. To listen on all interfaces (not recommended!) you can run `sudo src/cmd/fingerd/fingerd -a 0.0.0.0`.
To listen on a specific interface, specify the IP address of that interface.

Formatting
==========

Gofinger has been formatted with [gofmt](http://golang.org/cmd/gofmt/).  To re-run gofmt, run `make gofmt` from the main directory.

License
=======

Gofinger is released under the 3-clause BSD license.