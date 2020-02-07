package main

import "testing"

func TestPingFunction(t *testing.T) {
	testRequestStatusCode("POST", "/ping", nil, 200, t)
}
