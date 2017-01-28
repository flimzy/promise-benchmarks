package promise

import (
	"runtime"
	"sync"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

var emptyPromise = js.Global.Get("EmptyPromise")
var emptyCallback = js.Global.Get("EmptyCallback")

type promiseResult struct {
	result string
	err    error
}

// DoPromise1 executes a promise with sync.WaitGroup
func DoPromise1() (string, error) {
	var result string
	var wg sync.WaitGroup
	wg.Add(1)
	emptyPromise.Call("then", func(r *js.Object) {
		defer wg.Done()
		result = r.String()
	})
	wg.Wait()
	return result, nil
}

// DoPromise2 executes a promise with a channel
func DoPromise2() (string, error) {
	ch := make(chan promiseResult)
	defer close(ch)
	emptyPromise.Call("then", func(r *js.Object) {
		ch <- promiseResult{result: r.String()}
	})
	r := <-ch
	return r.result, r.err
}

const sleepTime time.Duration = 1315666 / 3 // 1315666ns is the time it takes to execute DoPromise2

// DoPromise3 executes a promise, waiting in a loop
func DoPromise3() (string, error) {
	var resolved bool
	var result string
	emptyPromise.Call("then", func(r *js.Object) {
		result = r.String()
		resolved = true
	})
	for {
		if resolved {
			return result, nil
		}
		runtime.Gosched()
	}
}

// DoCallback1 executes a callback with sync.WaitGroup
func DoCallback1() (string, error) {
	var result string
	var wg sync.WaitGroup
	wg.Add(1)
	js.Global.Call("EmptyCallback", func(r *js.Object) {
		defer wg.Done()
		result = r.String()
	})
	wg.Wait()
	return result, nil
}

// DoCallback2 executes a callback with a channel
func DoCallback2() (string, error) {
	ch := make(chan promiseResult, 1)
	defer close(ch)
	js.Global.Call("EmptyCallback", func(r *js.Object) {
		ch <- promiseResult{result: r.String()}
	})
	r := <-ch
	return r.result, nil
}
