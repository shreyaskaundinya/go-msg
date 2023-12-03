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

## Resources

-   https://blog.cloudflare.com/the-road-to-quic/
-   https://medium.com/@f_yuki/quic-is-now-available-for-ios15-bidirectional-stream-with-ios-and-quic-go-7b9178b7b4da
-   https://github.com/donbright/rust-lang-cheat-sheet
-   https://blog.logrocket.com/understanding-ownership-in-rust/
-   https://blog.logrocket.com/introducing-the-rust-borrow-checker/
-   https://www.cs.cmu.edu/~410-s05/lectures/L31_LockFree.pdf
-   https://docs.quic.tech/quiche/index.html
-   https://kafka.apache.org/documentation.html#design
-   https://kafka.apache.org/0100/protocol.html#:~:text=Kafka%20uses%20a%20binary%20protocol,as%20request%20response%20message%20pairs.
