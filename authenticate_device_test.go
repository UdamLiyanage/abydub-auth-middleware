package main

import "testing"

func TestAuthenticateDeviceFunction(t *testing.T) {
	testRequestStatusCode("POST", "/auth", nil, 200, t)
}
