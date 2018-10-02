// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/alphagov/paas-go/provider"
	"github.com/pivotal-cf/brokerapi"
)

type FakeServiceProvider struct {
	ProvisionStub        func(context.Context, provider.ProvisionData) (dashboardURL, operationData string, err error)
	provisionMutex       sync.RWMutex
	provisionArgsForCall []struct {
		arg1 context.Context
		arg2 provider.ProvisionData
	}
	provisionReturns struct {
		result1 string
		result2 string
		result3 error
	}
	provisionReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	DeprovisionStub        func(context.Context, provider.DeprovisionData) (operationData string, err error)
	deprovisionMutex       sync.RWMutex
	deprovisionArgsForCall []struct {
		arg1 context.Context
		arg2 provider.DeprovisionData
	}
	deprovisionReturns struct {
		result1 string
		result2 error
	}
	deprovisionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	BindStub        func(context.Context, provider.BindData) (binding brokerapi.Binding, err error)
	bindMutex       sync.RWMutex
	bindArgsForCall []struct {
		arg1 context.Context
		arg2 provider.BindData
	}
	bindReturns struct {
		result1 brokerapi.Binding
		result2 error
	}
	bindReturnsOnCall map[int]struct {
		result1 brokerapi.Binding
		result2 error
	}
	UnbindStub        func(context.Context, provider.UnbindData) (err error)
	unbindMutex       sync.RWMutex
	unbindArgsForCall []struct {
		arg1 context.Context
		arg2 provider.UnbindData
	}
	unbindReturns struct {
		result1 error
	}
	unbindReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(context.Context, provider.UpdateData) (operationData string, err error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 provider.UpdateData
	}
	updateReturns struct {
		result1 string
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	LastOperationStub        func(context.Context, provider.LastOperationData) (state brokerapi.LastOperationState, description string, err error)
	lastOperationMutex       sync.RWMutex
	lastOperationArgsForCall []struct {
		arg1 context.Context
		arg2 provider.LastOperationData
	}
	lastOperationReturns struct {
		result1 brokerapi.LastOperationState
		result2 string
		result3 error
	}
	lastOperationReturnsOnCall map[int]struct {
		result1 brokerapi.LastOperationState
		result2 string
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceProvider) Provision(arg1 context.Context, arg2 provider.ProvisionData) (dashboardURL, operationData string, err error) {
	fake.provisionMutex.Lock()
	ret, specificReturn := fake.provisionReturnsOnCall[len(fake.provisionArgsForCall)]
	fake.provisionArgsForCall = append(fake.provisionArgsForCall, struct {
		arg1 context.Context
		arg2 provider.ProvisionData
	}{arg1, arg2})
	fake.recordInvocation("Provision", []interface{}{arg1, arg2})
	fake.provisionMutex.Unlock()
	if fake.ProvisionStub != nil {
		return fake.ProvisionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.provisionReturns.result1, fake.provisionReturns.result2, fake.provisionReturns.result3
}

func (fake *FakeServiceProvider) ProvisionCallCount() int {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return len(fake.provisionArgsForCall)
}

func (fake *FakeServiceProvider) ProvisionArgsForCall(i int) (context.Context, provider.ProvisionData) {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return fake.provisionArgsForCall[i].arg1, fake.provisionArgsForCall[i].arg2
}

func (fake *FakeServiceProvider) ProvisionReturns(result1 string, result2 string, result3 error) {
	fake.ProvisionStub = nil
	fake.provisionReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServiceProvider) ProvisionReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.ProvisionStub = nil
	if fake.provisionReturnsOnCall == nil {
		fake.provisionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.provisionReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServiceProvider) Deprovision(arg1 context.Context, arg2 provider.DeprovisionData) (operationData string, err error) {
	fake.deprovisionMutex.Lock()
	ret, specificReturn := fake.deprovisionReturnsOnCall[len(fake.deprovisionArgsForCall)]
	fake.deprovisionArgsForCall = append(fake.deprovisionArgsForCall, struct {
		arg1 context.Context
		arg2 provider.DeprovisionData
	}{arg1, arg2})
	fake.recordInvocation("Deprovision", []interface{}{arg1, arg2})
	fake.deprovisionMutex.Unlock()
	if fake.DeprovisionStub != nil {
		return fake.DeprovisionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deprovisionReturns.result1, fake.deprovisionReturns.result2
}

func (fake *FakeServiceProvider) DeprovisionCallCount() int {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return len(fake.deprovisionArgsForCall)
}

func (fake *FakeServiceProvider) DeprovisionArgsForCall(i int) (context.Context, provider.DeprovisionData) {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return fake.deprovisionArgsForCall[i].arg1, fake.deprovisionArgsForCall[i].arg2
}

func (fake *FakeServiceProvider) DeprovisionReturns(result1 string, result2 error) {
	fake.DeprovisionStub = nil
	fake.deprovisionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceProvider) DeprovisionReturnsOnCall(i int, result1 string, result2 error) {
	fake.DeprovisionStub = nil
	if fake.deprovisionReturnsOnCall == nil {
		fake.deprovisionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.deprovisionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceProvider) Bind(arg1 context.Context, arg2 provider.BindData) (binding brokerapi.Binding, err error) {
	fake.bindMutex.Lock()
	ret, specificReturn := fake.bindReturnsOnCall[len(fake.bindArgsForCall)]
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		arg1 context.Context
		arg2 provider.BindData
	}{arg1, arg2})
	fake.recordInvocation("Bind", []interface{}{arg1, arg2})
	fake.bindMutex.Unlock()
	if fake.BindStub != nil {
		return fake.BindStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bindReturns.result1, fake.bindReturns.result2
}

func (fake *FakeServiceProvider) BindCallCount() int {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}

func (fake *FakeServiceProvider) BindArgsForCall(i int) (context.Context, provider.BindData) {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return fake.bindArgsForCall[i].arg1, fake.bindArgsForCall[i].arg2
}

func (fake *FakeServiceProvider) BindReturns(result1 brokerapi.Binding, result2 error) {
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1 brokerapi.Binding
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceProvider) BindReturnsOnCall(i int, result1 brokerapi.Binding, result2 error) {
	fake.BindStub = nil
	if fake.bindReturnsOnCall == nil {
		fake.bindReturnsOnCall = make(map[int]struct {
			result1 brokerapi.Binding
			result2 error
		})
	}
	fake.bindReturnsOnCall[i] = struct {
		result1 brokerapi.Binding
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceProvider) Unbind(arg1 context.Context, arg2 provider.UnbindData) (err error) {
	fake.unbindMutex.Lock()
	ret, specificReturn := fake.unbindReturnsOnCall[len(fake.unbindArgsForCall)]
	fake.unbindArgsForCall = append(fake.unbindArgsForCall, struct {
		arg1 context.Context
		arg2 provider.UnbindData
	}{arg1, arg2})
	fake.recordInvocation("Unbind", []interface{}{arg1, arg2})
	fake.unbindMutex.Unlock()
	if fake.UnbindStub != nil {
		return fake.UnbindStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unbindReturns.result1
}

func (fake *FakeServiceProvider) UnbindCallCount() int {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return len(fake.unbindArgsForCall)
}

func (fake *FakeServiceProvider) UnbindArgsForCall(i int) (context.Context, provider.UnbindData) {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return fake.unbindArgsForCall[i].arg1, fake.unbindArgsForCall[i].arg2
}

func (fake *FakeServiceProvider) UnbindReturns(result1 error) {
	fake.UnbindStub = nil
	fake.unbindReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceProvider) UnbindReturnsOnCall(i int, result1 error) {
	fake.UnbindStub = nil
	if fake.unbindReturnsOnCall == nil {
		fake.unbindReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unbindReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceProvider) Update(arg1 context.Context, arg2 provider.UpdateData) (operationData string, err error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 provider.UpdateData
	}{arg1, arg2})
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateReturns.result1, fake.updateReturns.result2
}

func (fake *FakeServiceProvider) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeServiceProvider) UpdateArgsForCall(i int) (context.Context, provider.UpdateData) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1, fake.updateArgsForCall[i].arg2
}

func (fake *FakeServiceProvider) UpdateReturns(result1 string, result2 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceProvider) UpdateReturnsOnCall(i int, result1 string, result2 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceProvider) LastOperation(arg1 context.Context, arg2 provider.LastOperationData) (state brokerapi.LastOperationState, description string, err error) {
	fake.lastOperationMutex.Lock()
	ret, specificReturn := fake.lastOperationReturnsOnCall[len(fake.lastOperationArgsForCall)]
	fake.lastOperationArgsForCall = append(fake.lastOperationArgsForCall, struct {
		arg1 context.Context
		arg2 provider.LastOperationData
	}{arg1, arg2})
	fake.recordInvocation("LastOperation", []interface{}{arg1, arg2})
	fake.lastOperationMutex.Unlock()
	if fake.LastOperationStub != nil {
		return fake.LastOperationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.lastOperationReturns.result1, fake.lastOperationReturns.result2, fake.lastOperationReturns.result3
}

func (fake *FakeServiceProvider) LastOperationCallCount() int {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return len(fake.lastOperationArgsForCall)
}

func (fake *FakeServiceProvider) LastOperationArgsForCall(i int) (context.Context, provider.LastOperationData) {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return fake.lastOperationArgsForCall[i].arg1, fake.lastOperationArgsForCall[i].arg2
}

func (fake *FakeServiceProvider) LastOperationReturns(result1 brokerapi.LastOperationState, result2 string, result3 error) {
	fake.LastOperationStub = nil
	fake.lastOperationReturns = struct {
		result1 brokerapi.LastOperationState
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServiceProvider) LastOperationReturnsOnCall(i int, result1 brokerapi.LastOperationState, result2 string, result3 error) {
	fake.LastOperationStub = nil
	if fake.lastOperationReturnsOnCall == nil {
		fake.lastOperationReturnsOnCall = make(map[int]struct {
			result1 brokerapi.LastOperationState
			result2 string
			result3 error
		})
	}
	fake.lastOperationReturnsOnCall[i] = struct {
		result1 brokerapi.LastOperationState
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeServiceProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceProvider) recordInvocation(key string, args []interface{}) {
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

var _ provider.ServiceProvider = new(FakeServiceProvider)
