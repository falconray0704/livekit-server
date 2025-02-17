// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc/types"
)

type FakeRoom struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	UpdateSubscriptionsStub        func(types.Participant, []string, bool) error
	updateSubscriptionsMutex       sync.RWMutex
	updateSubscriptionsArgsForCall []struct {
		arg1 types.Participant
		arg2 []string
		arg3 bool
	}
	updateSubscriptionsReturns struct {
		result1 error
	}
	updateSubscriptionsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRoom) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRoom) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeRoom) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeRoom) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRoom) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRoom) UpdateSubscriptions(arg1 types.Participant, arg2 []string, arg3 bool) error {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.updateSubscriptionsMutex.Lock()
	ret, specificReturn := fake.updateSubscriptionsReturnsOnCall[len(fake.updateSubscriptionsArgsForCall)]
	fake.updateSubscriptionsArgsForCall = append(fake.updateSubscriptionsArgsForCall, struct {
		arg1 types.Participant
		arg2 []string
		arg3 bool
	}{arg1, arg2Copy, arg3})
	stub := fake.UpdateSubscriptionsStub
	fakeReturns := fake.updateSubscriptionsReturns
	fake.recordInvocation("UpdateSubscriptions", []interface{}{arg1, arg2Copy, arg3})
	fake.updateSubscriptionsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRoom) UpdateSubscriptionsCallCount() int {
	fake.updateSubscriptionsMutex.RLock()
	defer fake.updateSubscriptionsMutex.RUnlock()
	return len(fake.updateSubscriptionsArgsForCall)
}

func (fake *FakeRoom) UpdateSubscriptionsCalls(stub func(types.Participant, []string, bool) error) {
	fake.updateSubscriptionsMutex.Lock()
	defer fake.updateSubscriptionsMutex.Unlock()
	fake.UpdateSubscriptionsStub = stub
}

func (fake *FakeRoom) UpdateSubscriptionsArgsForCall(i int) (types.Participant, []string, bool) {
	fake.updateSubscriptionsMutex.RLock()
	defer fake.updateSubscriptionsMutex.RUnlock()
	argsForCall := fake.updateSubscriptionsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRoom) UpdateSubscriptionsReturns(result1 error) {
	fake.updateSubscriptionsMutex.Lock()
	defer fake.updateSubscriptionsMutex.Unlock()
	fake.UpdateSubscriptionsStub = nil
	fake.updateSubscriptionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoom) UpdateSubscriptionsReturnsOnCall(i int, result1 error) {
	fake.updateSubscriptionsMutex.Lock()
	defer fake.updateSubscriptionsMutex.Unlock()
	fake.UpdateSubscriptionsStub = nil
	if fake.updateSubscriptionsReturnsOnCall == nil {
		fake.updateSubscriptionsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSubscriptionsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRoom) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.updateSubscriptionsMutex.RLock()
	defer fake.updateSubscriptionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRoom) recordInvocation(key string, args []interface{}) {
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

var _ types.Room = new(FakeRoom)
