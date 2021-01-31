// Code generated by counterfeiter. DO NOT EDIT.
package routingfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/routing"
	"github.com/livekit/livekit-server/proto/livekit"
)

type FakeRouter struct {
	ClearRoomStateStub        func(string) error
	clearRoomStateMutex       sync.RWMutex
	clearRoomStateArgsForCall []struct {
		arg1 string
	}
	clearRoomStateReturns struct {
		result1 error
	}
	clearRoomStateReturnsOnCall map[int]struct {
		result1 error
	}
	GetNodeStub        func(string) (*livekit.Node, error)
	getNodeMutex       sync.RWMutex
	getNodeArgsForCall []struct {
		arg1 string
	}
	getNodeReturns struct {
		result1 *livekit.Node
		result2 error
	}
	getNodeReturnsOnCall map[int]struct {
		result1 *livekit.Node
		result2 error
	}
	GetNodeForRoomStub        func(string) (string, error)
	getNodeForRoomMutex       sync.RWMutex
	getNodeForRoomArgsForCall []struct {
		arg1 string
	}
	getNodeForRoomReturns struct {
		result1 string
		result2 error
	}
	getNodeForRoomReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ListNodesStub        func() ([]*livekit.Node, error)
	listNodesMutex       sync.RWMutex
	listNodesArgsForCall []struct {
	}
	listNodesReturns struct {
		result1 []*livekit.Node
		result2 error
	}
	listNodesReturnsOnCall map[int]struct {
		result1 []*livekit.Node
		result2 error
	}
	OnNewParticipantRTCStub        func(routing.ParticipantCallback)
	onNewParticipantRTCMutex       sync.RWMutex
	onNewParticipantRTCArgsForCall []struct {
		arg1 routing.ParticipantCallback
	}
	RegisterNodeStub        func() error
	registerNodeMutex       sync.RWMutex
	registerNodeArgsForCall []struct {
	}
	registerNodeReturns struct {
		result1 error
	}
	registerNodeReturnsOnCall map[int]struct {
		result1 error
	}
	SetNodeForRoomStub        func(string, string) error
	setNodeForRoomMutex       sync.RWMutex
	setNodeForRoomArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setNodeForRoomReturns struct {
		result1 error
	}
	setNodeForRoomReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StartParticipantSignalStub        func(string, string, bool) (routing.MessageSink, routing.MessageSource, error)
	startParticipantSignalMutex       sync.RWMutex
	startParticipantSignalArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	startParticipantSignalReturns struct {
		result1 routing.MessageSink
		result2 routing.MessageSource
		result3 error
	}
	startParticipantSignalReturnsOnCall map[int]struct {
		result1 routing.MessageSink
		result2 routing.MessageSource
		result3 error
	}
	StopStub        func()
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	UnregisterNodeStub        func() error
	unregisterNodeMutex       sync.RWMutex
	unregisterNodeArgsForCall []struct {
	}
	unregisterNodeReturns struct {
		result1 error
	}
	unregisterNodeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRouter) ClearRoomState(arg1 string) error {
	fake.clearRoomStateMutex.Lock()
	ret, specificReturn := fake.clearRoomStateReturnsOnCall[len(fake.clearRoomStateArgsForCall)]
	fake.clearRoomStateArgsForCall = append(fake.clearRoomStateArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ClearRoomStateStub
	fakeReturns := fake.clearRoomStateReturns
	fake.recordInvocation("ClearRoomState", []interface{}{arg1})
	fake.clearRoomStateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) ClearRoomStateCallCount() int {
	fake.clearRoomStateMutex.RLock()
	defer fake.clearRoomStateMutex.RUnlock()
	return len(fake.clearRoomStateArgsForCall)
}

func (fake *FakeRouter) ClearRoomStateCalls(stub func(string) error) {
	fake.clearRoomStateMutex.Lock()
	defer fake.clearRoomStateMutex.Unlock()
	fake.ClearRoomStateStub = stub
}

func (fake *FakeRouter) ClearRoomStateArgsForCall(i int) string {
	fake.clearRoomStateMutex.RLock()
	defer fake.clearRoomStateMutex.RUnlock()
	argsForCall := fake.clearRoomStateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) ClearRoomStateReturns(result1 error) {
	fake.clearRoomStateMutex.Lock()
	defer fake.clearRoomStateMutex.Unlock()
	fake.ClearRoomStateStub = nil
	fake.clearRoomStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) ClearRoomStateReturnsOnCall(i int, result1 error) {
	fake.clearRoomStateMutex.Lock()
	defer fake.clearRoomStateMutex.Unlock()
	fake.ClearRoomStateStub = nil
	if fake.clearRoomStateReturnsOnCall == nil {
		fake.clearRoomStateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.clearRoomStateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) GetNode(arg1 string) (*livekit.Node, error) {
	fake.getNodeMutex.Lock()
	ret, specificReturn := fake.getNodeReturnsOnCall[len(fake.getNodeArgsForCall)]
	fake.getNodeArgsForCall = append(fake.getNodeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetNodeStub
	fakeReturns := fake.getNodeReturns
	fake.recordInvocation("GetNode", []interface{}{arg1})
	fake.getNodeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRouter) GetNodeCallCount() int {
	fake.getNodeMutex.RLock()
	defer fake.getNodeMutex.RUnlock()
	return len(fake.getNodeArgsForCall)
}

func (fake *FakeRouter) GetNodeCalls(stub func(string) (*livekit.Node, error)) {
	fake.getNodeMutex.Lock()
	defer fake.getNodeMutex.Unlock()
	fake.GetNodeStub = stub
}

func (fake *FakeRouter) GetNodeArgsForCall(i int) string {
	fake.getNodeMutex.RLock()
	defer fake.getNodeMutex.RUnlock()
	argsForCall := fake.getNodeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) GetNodeReturns(result1 *livekit.Node, result2 error) {
	fake.getNodeMutex.Lock()
	defer fake.getNodeMutex.Unlock()
	fake.GetNodeStub = nil
	fake.getNodeReturns = struct {
		result1 *livekit.Node
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) GetNodeReturnsOnCall(i int, result1 *livekit.Node, result2 error) {
	fake.getNodeMutex.Lock()
	defer fake.getNodeMutex.Unlock()
	fake.GetNodeStub = nil
	if fake.getNodeReturnsOnCall == nil {
		fake.getNodeReturnsOnCall = make(map[int]struct {
			result1 *livekit.Node
			result2 error
		})
	}
	fake.getNodeReturnsOnCall[i] = struct {
		result1 *livekit.Node
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) GetNodeForRoom(arg1 string) (string, error) {
	fake.getNodeForRoomMutex.Lock()
	ret, specificReturn := fake.getNodeForRoomReturnsOnCall[len(fake.getNodeForRoomArgsForCall)]
	fake.getNodeForRoomArgsForCall = append(fake.getNodeForRoomArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetNodeForRoomStub
	fakeReturns := fake.getNodeForRoomReturns
	fake.recordInvocation("GetNodeForRoom", []interface{}{arg1})
	fake.getNodeForRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRouter) GetNodeForRoomCallCount() int {
	fake.getNodeForRoomMutex.RLock()
	defer fake.getNodeForRoomMutex.RUnlock()
	return len(fake.getNodeForRoomArgsForCall)
}

func (fake *FakeRouter) GetNodeForRoomCalls(stub func(string) (string, error)) {
	fake.getNodeForRoomMutex.Lock()
	defer fake.getNodeForRoomMutex.Unlock()
	fake.GetNodeForRoomStub = stub
}

func (fake *FakeRouter) GetNodeForRoomArgsForCall(i int) string {
	fake.getNodeForRoomMutex.RLock()
	defer fake.getNodeForRoomMutex.RUnlock()
	argsForCall := fake.getNodeForRoomArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) GetNodeForRoomReturns(result1 string, result2 error) {
	fake.getNodeForRoomMutex.Lock()
	defer fake.getNodeForRoomMutex.Unlock()
	fake.GetNodeForRoomStub = nil
	fake.getNodeForRoomReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) GetNodeForRoomReturnsOnCall(i int, result1 string, result2 error) {
	fake.getNodeForRoomMutex.Lock()
	defer fake.getNodeForRoomMutex.Unlock()
	fake.GetNodeForRoomStub = nil
	if fake.getNodeForRoomReturnsOnCall == nil {
		fake.getNodeForRoomReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getNodeForRoomReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) ListNodes() ([]*livekit.Node, error) {
	fake.listNodesMutex.Lock()
	ret, specificReturn := fake.listNodesReturnsOnCall[len(fake.listNodesArgsForCall)]
	fake.listNodesArgsForCall = append(fake.listNodesArgsForCall, struct {
	}{})
	stub := fake.ListNodesStub
	fakeReturns := fake.listNodesReturns
	fake.recordInvocation("ListNodes", []interface{}{})
	fake.listNodesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRouter) ListNodesCallCount() int {
	fake.listNodesMutex.RLock()
	defer fake.listNodesMutex.RUnlock()
	return len(fake.listNodesArgsForCall)
}

func (fake *FakeRouter) ListNodesCalls(stub func() ([]*livekit.Node, error)) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = stub
}

func (fake *FakeRouter) ListNodesReturns(result1 []*livekit.Node, result2 error) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = nil
	fake.listNodesReturns = struct {
		result1 []*livekit.Node
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) ListNodesReturnsOnCall(i int, result1 []*livekit.Node, result2 error) {
	fake.listNodesMutex.Lock()
	defer fake.listNodesMutex.Unlock()
	fake.ListNodesStub = nil
	if fake.listNodesReturnsOnCall == nil {
		fake.listNodesReturnsOnCall = make(map[int]struct {
			result1 []*livekit.Node
			result2 error
		})
	}
	fake.listNodesReturnsOnCall[i] = struct {
		result1 []*livekit.Node
		result2 error
	}{result1, result2}
}

func (fake *FakeRouter) OnNewParticipantRTC(arg1 routing.ParticipantCallback) {
	fake.onNewParticipantRTCMutex.Lock()
	fake.onNewParticipantRTCArgsForCall = append(fake.onNewParticipantRTCArgsForCall, struct {
		arg1 routing.ParticipantCallback
	}{arg1})
	stub := fake.OnNewParticipantRTCStub
	fake.recordInvocation("OnNewParticipantRTC", []interface{}{arg1})
	fake.onNewParticipantRTCMutex.Unlock()
	if stub != nil {
		fake.OnNewParticipantRTCStub(arg1)
	}
}

func (fake *FakeRouter) OnNewParticipantRTCCallCount() int {
	fake.onNewParticipantRTCMutex.RLock()
	defer fake.onNewParticipantRTCMutex.RUnlock()
	return len(fake.onNewParticipantRTCArgsForCall)
}

func (fake *FakeRouter) OnNewParticipantRTCCalls(stub func(routing.ParticipantCallback)) {
	fake.onNewParticipantRTCMutex.Lock()
	defer fake.onNewParticipantRTCMutex.Unlock()
	fake.OnNewParticipantRTCStub = stub
}

func (fake *FakeRouter) OnNewParticipantRTCArgsForCall(i int) routing.ParticipantCallback {
	fake.onNewParticipantRTCMutex.RLock()
	defer fake.onNewParticipantRTCMutex.RUnlock()
	argsForCall := fake.onNewParticipantRTCArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRouter) RegisterNode() error {
	fake.registerNodeMutex.Lock()
	ret, specificReturn := fake.registerNodeReturnsOnCall[len(fake.registerNodeArgsForCall)]
	fake.registerNodeArgsForCall = append(fake.registerNodeArgsForCall, struct {
	}{})
	stub := fake.RegisterNodeStub
	fakeReturns := fake.registerNodeReturns
	fake.recordInvocation("RegisterNode", []interface{}{})
	fake.registerNodeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) RegisterNodeCallCount() int {
	fake.registerNodeMutex.RLock()
	defer fake.registerNodeMutex.RUnlock()
	return len(fake.registerNodeArgsForCall)
}

func (fake *FakeRouter) RegisterNodeCalls(stub func() error) {
	fake.registerNodeMutex.Lock()
	defer fake.registerNodeMutex.Unlock()
	fake.RegisterNodeStub = stub
}

func (fake *FakeRouter) RegisterNodeReturns(result1 error) {
	fake.registerNodeMutex.Lock()
	defer fake.registerNodeMutex.Unlock()
	fake.RegisterNodeStub = nil
	fake.registerNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) RegisterNodeReturnsOnCall(i int, result1 error) {
	fake.registerNodeMutex.Lock()
	defer fake.registerNodeMutex.Unlock()
	fake.RegisterNodeStub = nil
	if fake.registerNodeReturnsOnCall == nil {
		fake.registerNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.registerNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) SetNodeForRoom(arg1 string, arg2 string) error {
	fake.setNodeForRoomMutex.Lock()
	ret, specificReturn := fake.setNodeForRoomReturnsOnCall[len(fake.setNodeForRoomArgsForCall)]
	fake.setNodeForRoomArgsForCall = append(fake.setNodeForRoomArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.SetNodeForRoomStub
	fakeReturns := fake.setNodeForRoomReturns
	fake.recordInvocation("SetNodeForRoom", []interface{}{arg1, arg2})
	fake.setNodeForRoomMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) SetNodeForRoomCallCount() int {
	fake.setNodeForRoomMutex.RLock()
	defer fake.setNodeForRoomMutex.RUnlock()
	return len(fake.setNodeForRoomArgsForCall)
}

func (fake *FakeRouter) SetNodeForRoomCalls(stub func(string, string) error) {
	fake.setNodeForRoomMutex.Lock()
	defer fake.setNodeForRoomMutex.Unlock()
	fake.SetNodeForRoomStub = stub
}

func (fake *FakeRouter) SetNodeForRoomArgsForCall(i int) (string, string) {
	fake.setNodeForRoomMutex.RLock()
	defer fake.setNodeForRoomMutex.RUnlock()
	argsForCall := fake.setNodeForRoomArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRouter) SetNodeForRoomReturns(result1 error) {
	fake.setNodeForRoomMutex.Lock()
	defer fake.setNodeForRoomMutex.Unlock()
	fake.SetNodeForRoomStub = nil
	fake.setNodeForRoomReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) SetNodeForRoomReturnsOnCall(i int, result1 error) {
	fake.setNodeForRoomMutex.Lock()
	defer fake.setNodeForRoomMutex.Unlock()
	fake.SetNodeForRoomStub = nil
	if fake.setNodeForRoomReturnsOnCall == nil {
		fake.setNodeForRoomReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setNodeForRoomReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	stub := fake.StartStub
	fakeReturns := fake.startReturns
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeRouter) StartCalls(stub func() error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeRouter) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) StartParticipantSignal(arg1 string, arg2 string, arg3 bool) (routing.MessageSink, routing.MessageSource, error) {
	fake.startParticipantSignalMutex.Lock()
	ret, specificReturn := fake.startParticipantSignalReturnsOnCall[len(fake.startParticipantSignalArgsForCall)]
	fake.startParticipantSignalArgsForCall = append(fake.startParticipantSignalArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.StartParticipantSignalStub
	fakeReturns := fake.startParticipantSignalReturns
	fake.recordInvocation("StartParticipantSignal", []interface{}{arg1, arg2, arg3})
	fake.startParticipantSignalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeRouter) StartParticipantSignalCallCount() int {
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	return len(fake.startParticipantSignalArgsForCall)
}

func (fake *FakeRouter) StartParticipantSignalCalls(stub func(string, string, bool) (routing.MessageSink, routing.MessageSource, error)) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = stub
}

func (fake *FakeRouter) StartParticipantSignalArgsForCall(i int) (string, string, bool) {
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	argsForCall := fake.startParticipantSignalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRouter) StartParticipantSignalReturns(result1 routing.MessageSink, result2 routing.MessageSource, result3 error) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = nil
	fake.startParticipantSignalReturns = struct {
		result1 routing.MessageSink
		result2 routing.MessageSource
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRouter) StartParticipantSignalReturnsOnCall(i int, result1 routing.MessageSink, result2 routing.MessageSource, result3 error) {
	fake.startParticipantSignalMutex.Lock()
	defer fake.startParticipantSignalMutex.Unlock()
	fake.StartParticipantSignalStub = nil
	if fake.startParticipantSignalReturnsOnCall == nil {
		fake.startParticipantSignalReturnsOnCall = make(map[int]struct {
			result1 routing.MessageSink
			result2 routing.MessageSource
			result3 error
		})
	}
	fake.startParticipantSignalReturnsOnCall[i] = struct {
		result1 routing.MessageSink
		result2 routing.MessageSource
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeRouter) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	stub := fake.StopStub
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if stub != nil {
		fake.StopStub()
	}
}

func (fake *FakeRouter) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeRouter) StopCalls(stub func()) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeRouter) UnregisterNode() error {
	fake.unregisterNodeMutex.Lock()
	ret, specificReturn := fake.unregisterNodeReturnsOnCall[len(fake.unregisterNodeArgsForCall)]
	fake.unregisterNodeArgsForCall = append(fake.unregisterNodeArgsForCall, struct {
	}{})
	stub := fake.UnregisterNodeStub
	fakeReturns := fake.unregisterNodeReturns
	fake.recordInvocation("UnregisterNode", []interface{}{})
	fake.unregisterNodeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRouter) UnregisterNodeCallCount() int {
	fake.unregisterNodeMutex.RLock()
	defer fake.unregisterNodeMutex.RUnlock()
	return len(fake.unregisterNodeArgsForCall)
}

func (fake *FakeRouter) UnregisterNodeCalls(stub func() error) {
	fake.unregisterNodeMutex.Lock()
	defer fake.unregisterNodeMutex.Unlock()
	fake.UnregisterNodeStub = stub
}

func (fake *FakeRouter) UnregisterNodeReturns(result1 error) {
	fake.unregisterNodeMutex.Lock()
	defer fake.unregisterNodeMutex.Unlock()
	fake.UnregisterNodeStub = nil
	fake.unregisterNodeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) UnregisterNodeReturnsOnCall(i int, result1 error) {
	fake.unregisterNodeMutex.Lock()
	defer fake.unregisterNodeMutex.Unlock()
	fake.UnregisterNodeStub = nil
	if fake.unregisterNodeReturnsOnCall == nil {
		fake.unregisterNodeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unregisterNodeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clearRoomStateMutex.RLock()
	defer fake.clearRoomStateMutex.RUnlock()
	fake.getNodeMutex.RLock()
	defer fake.getNodeMutex.RUnlock()
	fake.getNodeForRoomMutex.RLock()
	defer fake.getNodeForRoomMutex.RUnlock()
	fake.listNodesMutex.RLock()
	defer fake.listNodesMutex.RUnlock()
	fake.onNewParticipantRTCMutex.RLock()
	defer fake.onNewParticipantRTCMutex.RUnlock()
	fake.registerNodeMutex.RLock()
	defer fake.registerNodeMutex.RUnlock()
	fake.setNodeForRoomMutex.RLock()
	defer fake.setNodeForRoomMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.startParticipantSignalMutex.RLock()
	defer fake.startParticipantSignalMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.unregisterNodeMutex.RLock()
	defer fake.unregisterNodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRouter) recordInvocation(key string, args []interface{}) {
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

var _ routing.Router = new(FakeRouter)
