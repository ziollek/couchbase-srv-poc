# couchbase-srv-poc

PoC connect to couchbase using service domain accessible via consul-dns service

## what does it contain?

1. coredns server that makes some rewrites and forwarding queries to consul agent 
2. dnsuti docker that uses above dns as a default and allows perform sample dns queries using dig
3. very simple go module that allows testing if it is possible to establish connection

## before run

Prepare your own configuration that should contain two variables

```
cat .env
CONSUL_AGENT=<IP>:8600
CONSUL_SUFFIX=suffix-name
```

## how to run


```
docker-compose up
```

### sample dns resolution (in another terminal than docker-compose)

```
docker-compose exec dnsutils bash
dig -t SRV _couchbase._tcp.5cache-client-cluster7.service.appengine.

; <<>> DiG 9.16.33-Debian <<>> -t SRV _couchbase._tcp.cluster.consul.suffix.
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 62378
;; flags: qr aa rd; QUERY: 1, ANSWER: 3, AUTHORITY: 0, ADDITIONAL: 7
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
;; QUESTION SECTION:
;_couchbase._tcp.cluster.consul.suffix.. IN SRV

;; ANSWER SECTION:
cluster.consul.suffix.. 0 IN SRV 1 1 11210 0a521042.addr.dc-xxx.suffix. 
```


### testing connection via tool

```
docker-compose exec connection-tester bash
go run . couchbase://my-cluster.consul.suffix bucket user pass
2023/10/25 20:28:34 connecting using following parameters connection=couchbase://my-cluster.consul.suffix, bucket=bucket, user=user, pass=***
2023/10/25 20:28:39 OK!
```
