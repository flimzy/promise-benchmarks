This repo contains tests for interacting with JS Promises and callbacks from GopherJS.

To test, I created an empty promise, and an empty function which returns a callback.

```
$global.EmptyPromise = new Promise(function(resolve, reject) {
    resolve('Success!');
});

$global.EmptyCallback = function(resolve, reject) {
    resolve('Success!');
};
```

This to ensure that I'm only testing the interaction between promises/callbacks
and GopherJS, and as little else as possible.

Then I constructed a few different ways of interacting with these functions:

- Using `sync.WaitGroup`
- Using bare channels
- Waiting in a loop

Unsurprisingly, bare channels win (as the other two methods use channels in the
background, with additional overhead).

The most surprising finding, for me, was that using bare callbacks were far
more performant, on the order of 225x faster than promises. To a JS developer,
this likely is not surprising at all, but alas, I'm a Go developer.

```
BenchmarkDoPromise1        10000           1347100 ns/op
BenchmarkDoPromise2        10000           1303800 ns/op
BenchmarkDoPromise3        10000           2379500 ns/op
BenchmarkDoCallback1     2000000             11437 ns/op
BenchmarkDoCallback2     2000000              5706 ns/op
```

I ran my tests using node.js v6.9.3 on Debian stretch (testing), kernel
4.2.0-1-amd64 on a Lenovo Thinkpad T430s (Intel(R) Core(TM) i7-3520M CPU @ 2.90GHz x 4 cores).
I ran my benchmarks using the following command:

    gopherjs test github.com/flimzy/promise --bench=. --benchtime=10s

With [PR 578](https://github.com/gopherjs/gopherjs/pull/578) merged to support
the `--benchtime` flag.
