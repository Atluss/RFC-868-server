Time server RFC 868
=====================

Description
-----------------------------------
This program run time protocol server as described in [RFC 868](https://tools.ietf.org/html/rfc868)

How to run
-----------------------------------
`server -p 11037` in server folder arguments is necessary. Server runs in localhost

How to test server
-----------------------------------
`printf "%d\n" "0x$(nc time.nist.gov 37 | xxd -p)"` - in terminal where `time.nist.gov 37` is address and server port.

How to run client
-----------------------------------
`client localhost 11037` in client folder arguments is necessary

Build
-----------------------------------
* server /server/: go build server.go
* client /client/: go build client.go