Hacking CEL and test harness and everything...

The test harness in *tests/main* is built based on v1. However, this test
harness only test against go (no other languages). Conformance tests in
the future may have a similar structure like this. Once we implement
more rules, we simply need to copy the corresponding cases into 
proto/tests/harness/cases from v1 and update *tests/main/cases.go*.