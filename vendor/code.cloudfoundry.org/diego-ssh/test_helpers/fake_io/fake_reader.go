// This file was generated by counterfeiter
package fake_io

import (
	"io"
	"sync"
)

type FakeReader struct {
	ReadStub        func(p []byte) (n int, err error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		p []byte
	}
	readReturns struct {
		result1 int
		result2 error
	}
}

func (fake *FakeReader) Read(p []byte) (n int, err error) {
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		p []byte
	}{p})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(p)
	} else {
		return fake.readReturns.result1, fake.readReturns.result2
	}
}

func (fake *FakeReader) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeReader) ReadArgsForCall(i int) []byte {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].p
}

func (fake *FakeReader) ReadReturns(result1 int, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

var _ io.Reader = new(FakeReader)
