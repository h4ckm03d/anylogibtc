// Code generated by counterfeiter. DO NOT EDIT.
package transactionfakes

import (
	"anylogibtc/dto"
	"anylogibtc/services/transaction"
	"context"
	"sync"
	"time"
)

type FakeTransactionService struct {
	HistoryStub        func(context.Context, time.Time, time.Time) (transaction.HistoriesDTO, error)
	historyMutex       sync.RWMutex
	historyArgsForCall []struct {
		arg1 context.Context
		arg2 time.Time
		arg3 time.Time
	}
	historyReturns struct {
		result1 transaction.HistoriesDTO
		result2 error
	}
	historyReturnsOnCall map[int]struct {
		result1 transaction.HistoriesDTO
		result2 error
	}
	SendStub        func(context.Context, dto.TransactionDTO) error
	sendMutex       sync.RWMutex
	sendArgsForCall []struct {
		arg1 context.Context
		arg2 dto.TransactionDTO
	}
	sendReturns struct {
		result1 error
	}
	sendReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTransactionService) History(arg1 context.Context, arg2 time.Time, arg3 time.Time) (transaction.HistoriesDTO, error) {
	fake.historyMutex.Lock()
	ret, specificReturn := fake.historyReturnsOnCall[len(fake.historyArgsForCall)]
	fake.historyArgsForCall = append(fake.historyArgsForCall, struct {
		arg1 context.Context
		arg2 time.Time
		arg3 time.Time
	}{arg1, arg2, arg3})
	stub := fake.HistoryStub
	fakeReturns := fake.historyReturns
	fake.recordInvocation("History", []interface{}{arg1, arg2, arg3})
	fake.historyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTransactionService) HistoryCallCount() int {
	fake.historyMutex.RLock()
	defer fake.historyMutex.RUnlock()
	return len(fake.historyArgsForCall)
}

func (fake *FakeTransactionService) HistoryCalls(stub func(context.Context, time.Time, time.Time) (transaction.HistoriesDTO, error)) {
	fake.historyMutex.Lock()
	defer fake.historyMutex.Unlock()
	fake.HistoryStub = stub
}

func (fake *FakeTransactionService) HistoryArgsForCall(i int) (context.Context, time.Time, time.Time) {
	fake.historyMutex.RLock()
	defer fake.historyMutex.RUnlock()
	argsForCall := fake.historyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeTransactionService) HistoryReturns(result1 transaction.HistoriesDTO, result2 error) {
	fake.historyMutex.Lock()
	defer fake.historyMutex.Unlock()
	fake.HistoryStub = nil
	fake.historyReturns = struct {
		result1 transaction.HistoriesDTO
		result2 error
	}{result1, result2}
}

func (fake *FakeTransactionService) HistoryReturnsOnCall(i int, result1 transaction.HistoriesDTO, result2 error) {
	fake.historyMutex.Lock()
	defer fake.historyMutex.Unlock()
	fake.HistoryStub = nil
	if fake.historyReturnsOnCall == nil {
		fake.historyReturnsOnCall = make(map[int]struct {
			result1 transaction.HistoriesDTO
			result2 error
		})
	}
	fake.historyReturnsOnCall[i] = struct {
		result1 transaction.HistoriesDTO
		result2 error
	}{result1, result2}
}

func (fake *FakeTransactionService) Send(arg1 context.Context, arg2 dto.TransactionDTO) error {
	fake.sendMutex.Lock()
	ret, specificReturn := fake.sendReturnsOnCall[len(fake.sendArgsForCall)]
	fake.sendArgsForCall = append(fake.sendArgsForCall, struct {
		arg1 context.Context
		arg2 dto.TransactionDTO
	}{arg1, arg2})
	stub := fake.SendStub
	fakeReturns := fake.sendReturns
	fake.recordInvocation("Send", []interface{}{arg1, arg2})
	fake.sendMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeTransactionService) SendCallCount() int {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	return len(fake.sendArgsForCall)
}

func (fake *FakeTransactionService) SendCalls(stub func(context.Context, dto.TransactionDTO) error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = stub
}

func (fake *FakeTransactionService) SendArgsForCall(i int) (context.Context, dto.TransactionDTO) {
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	argsForCall := fake.sendArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTransactionService) SendReturns(result1 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	fake.sendReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTransactionService) SendReturnsOnCall(i int, result1 error) {
	fake.sendMutex.Lock()
	defer fake.sendMutex.Unlock()
	fake.SendStub = nil
	if fake.sendReturnsOnCall == nil {
		fake.sendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTransactionService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.historyMutex.RLock()
	defer fake.historyMutex.RUnlock()
	fake.sendMutex.RLock()
	defer fake.sendMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTransactionService) recordInvocation(key string, args []interface{}) {
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

var _ transaction.TransactionService = new(FakeTransactionService)
