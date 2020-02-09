# Installing
```shell
# install nsq (see https://nsq.io/deployment/installing.html)
brew install nsq

# start nsqlookup daemon
nsqlookupd

# start nsq daemon
nsqd --lookupd-tcp-address=127.0.0.1:4160 --broadcast-address=127.0.0.1

# (optional) start browser-based admin ui at http://127.0.0.1:4171
nsqadmin --lookupd-http-address=127.0.0.1:4161

# (optional) dump messages to file (automatically creates `test` topic)
nsq_to_file --topic=test --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161

# send messages to `test` topic
curl -d 'test 1' 'http://127.0.0.1:4151/pub?topic=test'
curl -d 'test 2' 'http://127.0.0.1:4151/pub?topic=test'
curl -d 'test 3' 'http://127.0.0.1:4151/pub?topic=test'
curl -d 'test 4' 'http://127.0.0.1:4151/pub?topic=test'
curl -d 'test 5' 'http://127.0.0.1:4151/pub?topic=test'
```

# Debugging
```shell
# run nsq subscriber
go run nsq-subscriber/main.go nsq-subscriber/tail.go

# run http server
go run http-server/main.go http-server/nsq.go http-server/page.go
```

# Links
- https://golang.org/doc/articles/wiki/
- https://github.com/nsqio/nsq/tree/master/apps/nsq_tail
