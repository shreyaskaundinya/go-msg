# go-msg => PubSub based message queue in golang [WIP]

## Questions ?

-   Use quic and webtransport vs tcp socket
-   Make a comparison of quic webtransport server throughput and latencies vs
    tcp socket throughput
-   look into sequential IO // append only log
-   zero copy using sendfile syscall / ioWriter :
    https://itnext.io/optimizing-large-file-transfers-in-linux-with-go-an-exploration-of-tcp-and-syscall-ebe1b93fb72f

## Design

### Publisher

-   Publishing

### Controller

### Worker

-   Worker groups
-   Topics
-   Partition
-   Partition management | Replication
-   Zero copy using Sendfile syscall

### Subscriber Group

-   Subscriber
-   Partition splitting
