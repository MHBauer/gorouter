// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"

	"code.cloudfoundry.org/gorouter/proxy/round_tripper"
)

type FakeProxyRoundTripper struct {
	RoundTripStub        func(*http.Request) (*http.Response, error)
	roundTripMutex       sync.RWMutex
	roundTripArgsForCall []struct {
		arg1 *http.Request
	}
	roundTripReturns struct {
		result1 *http.Response
		result2 error
	}
	roundTripReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	CancelRequestStub        func(*http.Request)
	cancelRequestMutex       sync.RWMutex
	cancelRequestArgsForCall []struct {
		arg1 *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProxyRoundTripper) RoundTrip(arg1 *http.Request) (*http.Response, error) {
	fake.roundTripMutex.Lock()
	ret, specificReturn := fake.roundTripReturnsOnCall[len(fake.roundTripArgsForCall)]
	fake.roundTripArgsForCall = append(fake.roundTripArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.recordInvocation("RoundTrip", []interface{}{arg1})
	fake.roundTripMutex.Unlock()
	if fake.RoundTripStub != nil {
		return fake.RoundTripStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.roundTripReturns.result1, fake.roundTripReturns.result2
}

func (fake *FakeProxyRoundTripper) RoundTripCallCount() int {
	fake.roundTripMutex.RLock()
	defer fake.roundTripMutex.RUnlock()
	return len(fake.roundTripArgsForCall)
}

func (fake *FakeProxyRoundTripper) RoundTripArgsForCall(i int) *http.Request {
	fake.roundTripMutex.RLock()
	defer fake.roundTripMutex.RUnlock()
	return fake.roundTripArgsForCall[i].arg1
}

func (fake *FakeProxyRoundTripper) RoundTripReturns(result1 *http.Response, result2 error) {
	fake.RoundTripStub = nil
	fake.roundTripReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeProxyRoundTripper) RoundTripReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.RoundTripStub = nil
	if fake.roundTripReturnsOnCall == nil {
		fake.roundTripReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.roundTripReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeProxyRoundTripper) CancelRequest(arg1 *http.Request) {
	fake.cancelRequestMutex.Lock()
	fake.cancelRequestArgsForCall = append(fake.cancelRequestArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.recordInvocation("CancelRequest", []interface{}{arg1})
	fake.cancelRequestMutex.Unlock()
	if fake.CancelRequestStub != nil {
		fake.CancelRequestStub(arg1)
	}
}

func (fake *FakeProxyRoundTripper) CancelRequestCallCount() int {
	fake.cancelRequestMutex.RLock()
	defer fake.cancelRequestMutex.RUnlock()
	return len(fake.cancelRequestArgsForCall)
}

func (fake *FakeProxyRoundTripper) CancelRequestArgsForCall(i int) *http.Request {
	fake.cancelRequestMutex.RLock()
	defer fake.cancelRequestMutex.RUnlock()
	return fake.cancelRequestArgsForCall[i].arg1
}

func (fake *FakeProxyRoundTripper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.roundTripMutex.RLock()
	defer fake.roundTripMutex.RUnlock()
	fake.cancelRequestMutex.RLock()
	defer fake.cancelRequestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProxyRoundTripper) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ round_tripper.ProxyRoundTripper = new(FakeProxyRoundTripper)
