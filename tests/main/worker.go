package main

import (
	"fmt"
	hack_cel "github.com/bufbuild/hack-cel"
	"sync"
)

func Work(wg *sync.WaitGroup, in <-chan TestCase, out chan<- TestResult) {
	for tc := range in {
		execTestCase(tc, out)
	}
	wg.Done()
}

func execTestCase(tc TestCase, out chan<- TestResult) {
	err := hack_cel.Validate(tc.Message)
	if err != nil {
		fmt.Println("got an error:", err)
		if tc.Failures == 0 {
			out <- TestResult{OK: false, Skipped: false}
		} else {
			out <- TestResult{OK: true, Skipped: false}
		}
	} else {
		if tc.Failures == 0 {
			out <- TestResult{OK: true, Skipped: false}
		} else {
			out <- TestResult{OK: false, Skipped: false}
		}
	}
	return
}
