package promise

import "testing"

var expected = "Success!"

func TestDoPromise1(t *testing.T) {
	result, _ := DoPromise1()
	if result != expected {
		t.Errorf("DoPromise1 failed\n\tExpected: %s\n\t  Actual: %s\n", expected, result)
	}
}

func BenchmarkDoPromise1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = DoPromise1()
	}
}

func TestDoPromise2(t *testing.T) {
	result, _ := DoPromise2()
	if result != expected {
		t.Errorf("DoPromise2 failed\n\tExpected: %s\n\t  Actual: %s\n", expected, result)
	}
}

func BenchmarkDoPromise2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = DoPromise2()
	}
}

func TestDoPromise3(t *testing.T) {
	result, _ := DoPromise3()
	if result != expected {
		t.Errorf("DoPromise3 failed\n\tExpected: %s\n\t  Actual: %s\n", expected, result)
	}
}

func BenchmarkDoPromise3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = DoPromise3()
	}
}

func TestDoCallback1(t *testing.T) {
	result, _ := DoCallback1()
	if result != expected {
		t.Errorf("DoCallback1 failed\n\tExpected: %s\n\t  Actual: %s\n", expected, result)
	}
}

func BenchmarkDoCallback1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = DoCallback1()
	}
}

func TestDoCallback2(t *testing.T) {
	result, _ := DoCallback2()
	if result != expected {
		t.Errorf("DoCallback2 failed\n\tExpected: %s\n\t  Actual: %s\n", expected, result)
	}
}

func BenchmarkDoCallback2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = DoCallback2()
	}
}
