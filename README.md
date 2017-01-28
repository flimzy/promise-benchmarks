This repo contains tests for interacting with JS Promises and callbacks from GopherJS.

To test, I created an empty promise, and an empty function which returns a callback.

```
$global.EmptyPromise = function() {
    return new Promise(function(resolve, reject) {
        resolve('Success!');
    });
};
$global.EmptyCallback = function(resolve, reject) {
    resolve('Success!');
};
```

This to ensure that I'm only testing the interaction between promises/callbacks
and GopherJS, and as little else as possible.

Then I constructed a few different ways of interacting with these functions:

- Using `sync.WaitGroup` (DoPromise1 & DoCallback1)
- Using bare channels (DoPromise2, DoCallback2, RawPromise & RawCallback)
- Waiting in a loop (DoPromise3)

Unsurprisingly, bare channels win (as the other two methods use channels in the
background, with additional overhead).

Further, the `RawPromise` and `RawCallback` benchmarks attempt to benchmark
only the channel operation (inspired by @r-l-x's [tests](https://gopherjs.github.io/playground/#/btm0GV_mQY)
on [GopherJS PR#558](https://github.com/gopherjs/gopherjs/pull/558)). These functions
provide the most surprising finding to me: That even when attempting to isolate
channel operation, callbacks are faster than promises. Something must be wrong
with these tests.

The most surprising finding, for me, was that using bare callbacks were far
more performant, on the order of 225x faster than promises.

```
BenchmarkChannel                  500000              2968 ns/op
BenchmarkChannelInGoroutine       300000              5056 ns/op
BenchmarkRawPromise                 2000           1179000 ns/op
BenchmarkRawCalback               300000              5153 ns/op
BenchmarkDoPromise1                 2000           1213000 ns/op
BenchmarkDoPromise2                 2000           1224000 ns/op
BenchmarkDoPromise3                 1000           2295000 ns/op
BenchmarkDoCallback1              200000              9760 ns/op
BenchmarkDoCallback2              200000              6670 ns/op
```

I ran my tests using node.js v6.9.3 on Debian stretch (testing), kernel
4.2.0-1-amd64 on a Lenovo Thinkpad T430s (Intel(R) Core(TM) i7-3520M CPU @ 2.90GHz x 4 cores).
I ran my benchmarks using the following command:

    gopherjs test github.com/flimzy/promise --bench=.
