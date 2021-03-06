// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/gorouter/accesslog"
	"code.cloudfoundry.org/gorouter/accesslog/schema"
)

type FakeAccessLogger struct {
	RunStub         func()
	runMutex        sync.RWMutex
	runArgsForCall  []struct{}
	StopStub        func()
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	LogStub         func(record schema.AccessLogRecord)
	logMutex        sync.RWMutex
	logArgsForCall  []struct {
		record schema.AccessLogRecord
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAccessLogger) Run() {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct{}{})
	fake.recordInvocation("Run", []interface{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		fake.RunStub()
	}
}

func (fake *FakeAccessLogger) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeAccessLogger) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		fake.StopStub()
	}
}

func (fake *FakeAccessLogger) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeAccessLogger) Log(record schema.AccessLogRecord) {
	fake.logMutex.Lock()
	fake.logArgsForCall = append(fake.logArgsForCall, struct {
		record schema.AccessLogRecord
	}{record})
	fake.recordInvocation("Log", []interface{}{record})
	fake.logMutex.Unlock()
	if fake.LogStub != nil {
		fake.LogStub(record)
	}
}

func (fake *FakeAccessLogger) LogCallCount() int {
	fake.logMutex.RLock()
	defer fake.logMutex.RUnlock()
	return len(fake.logArgsForCall)
}

func (fake *FakeAccessLogger) LogArgsForCall(i int) schema.AccessLogRecord {
	fake.logMutex.RLock()
	defer fake.logMutex.RUnlock()
	return fake.logArgsForCall[i].record
}

func (fake *FakeAccessLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.logMutex.RLock()
	defer fake.logMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAccessLogger) recordInvocation(key string, args []interface{}) {
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

var _ accesslog.AccessLogger = new(FakeAccessLogger)
