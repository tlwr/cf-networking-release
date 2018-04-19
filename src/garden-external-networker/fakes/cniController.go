// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/containernetworking/cni/pkg/types"
)

type CNIController struct {
	UpStub        func(namespacePath, handle string, metadata map[string]interface{}, legacyNetConf map[string]interface{}) (types.Result, error)
	upMutex       sync.RWMutex
	upArgsForCall []struct {
		namespacePath string
		handle        string
		metadata      map[string]interface{}
		legacyNetConf map[string]interface{}
	}
	upReturns struct {
		result1 types.Result
		result2 error
	}
	upReturnsOnCall map[int]struct {
		result1 types.Result
		result2 error
	}
	DownStub        func(namespacePath, handle string) error
	downMutex       sync.RWMutex
	downArgsForCall []struct {
		namespacePath string
		handle        string
	}
	downReturns struct {
		result1 error
	}
	downReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CNIController) Up(namespacePath string, handle string, metadata map[string]interface{}, legacyNetConf map[string]interface{}) (types.Result, error) {
	fake.upMutex.Lock()
	ret, specificReturn := fake.upReturnsOnCall[len(fake.upArgsForCall)]
	fake.upArgsForCall = append(fake.upArgsForCall, struct {
		namespacePath string
		handle        string
		metadata      map[string]interface{}
		legacyNetConf map[string]interface{}
	}{namespacePath, handle, metadata, legacyNetConf})
	fake.recordInvocation("Up", []interface{}{namespacePath, handle, metadata, legacyNetConf})
	fake.upMutex.Unlock()
	if fake.UpStub != nil {
		return fake.UpStub(namespacePath, handle, metadata, legacyNetConf)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.upReturns.result1, fake.upReturns.result2
}

func (fake *CNIController) UpCallCount() int {
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	return len(fake.upArgsForCall)
}

func (fake *CNIController) UpArgsForCall(i int) (string, string, map[string]interface{}, map[string]interface{}) {
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	return fake.upArgsForCall[i].namespacePath, fake.upArgsForCall[i].handle, fake.upArgsForCall[i].metadata, fake.upArgsForCall[i].legacyNetConf
}

func (fake *CNIController) UpReturns(result1 types.Result, result2 error) {
	fake.UpStub = nil
	fake.upReturns = struct {
		result1 types.Result
		result2 error
	}{result1, result2}
}

func (fake *CNIController) UpReturnsOnCall(i int, result1 types.Result, result2 error) {
	fake.UpStub = nil
	if fake.upReturnsOnCall == nil {
		fake.upReturnsOnCall = make(map[int]struct {
			result1 types.Result
			result2 error
		})
	}
	fake.upReturnsOnCall[i] = struct {
		result1 types.Result
		result2 error
	}{result1, result2}
}

func (fake *CNIController) Down(namespacePath string, handle string) error {
	fake.downMutex.Lock()
	ret, specificReturn := fake.downReturnsOnCall[len(fake.downArgsForCall)]
	fake.downArgsForCall = append(fake.downArgsForCall, struct {
		namespacePath string
		handle        string
	}{namespacePath, handle})
	fake.recordInvocation("Down", []interface{}{namespacePath, handle})
	fake.downMutex.Unlock()
	if fake.DownStub != nil {
		return fake.DownStub(namespacePath, handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.downReturns.result1
}

func (fake *CNIController) DownCallCount() int {
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	return len(fake.downArgsForCall)
}

func (fake *CNIController) DownArgsForCall(i int) (string, string) {
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	return fake.downArgsForCall[i].namespacePath, fake.downArgsForCall[i].handle
}

func (fake *CNIController) DownReturns(result1 error) {
	fake.DownStub = nil
	fake.downReturns = struct {
		result1 error
	}{result1}
}

func (fake *CNIController) DownReturnsOnCall(i int, result1 error) {
	fake.DownStub = nil
	if fake.downReturnsOnCall == nil {
		fake.downReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *CNIController) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.upMutex.RLock()
	defer fake.upMutex.RUnlock()
	fake.downMutex.RLock()
	defer fake.downMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CNIController) recordInvocation(key string, args []interface{}) {
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
