// Code generated by counterfeiter. DO NOT EDIT.
package userfakes

import (
	"context"
	"sync"

	usera "github.com/ebalkanski/goa/gen/user"
	"github.com/ebalkanski/goa/internal/service/user"
)

type FakeUserRepo struct {
	CreateStub        func(context.Context, *usera.User) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 *usera.User
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(context.Context, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	EditStub        func(context.Context, *usera.User) error
	editMutex       sync.RWMutex
	editArgsForCall []struct {
		arg1 context.Context
		arg2 *usera.User
	}
	editReturns struct {
		result1 error
	}
	editReturnsOnCall map[int]struct {
		result1 error
	}
	UserStub        func(context.Context, string) (*usera.User, error)
	userMutex       sync.RWMutex
	userArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	userReturns struct {
		result1 *usera.User
		result2 error
	}
	userReturnsOnCall map[int]struct {
		result1 *usera.User
		result2 error
	}
	UsersStub        func(context.Context) ([]*usera.User, error)
	usersMutex       sync.RWMutex
	usersArgsForCall []struct {
		arg1 context.Context
	}
	usersReturns struct {
		result1 []*usera.User
		result2 error
	}
	usersReturnsOnCall map[int]struct {
		result1 []*usera.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserRepo) Create(arg1 context.Context, arg2 *usera.User) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 *usera.User
	}{arg1, arg2})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserRepo) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeUserRepo) CreateCalls(stub func(context.Context, *usera.User) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeUserRepo) CreateArgsForCall(i int) (context.Context, *usera.User) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserRepo) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) Delete(arg1 context.Context, arg2 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserRepo) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeUserRepo) DeleteCalls(stub func(context.Context, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeUserRepo) DeleteArgsForCall(i int) (context.Context, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserRepo) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) Edit(arg1 context.Context, arg2 *usera.User) error {
	fake.editMutex.Lock()
	ret, specificReturn := fake.editReturnsOnCall[len(fake.editArgsForCall)]
	fake.editArgsForCall = append(fake.editArgsForCall, struct {
		arg1 context.Context
		arg2 *usera.User
	}{arg1, arg2})
	stub := fake.EditStub
	fakeReturns := fake.editReturns
	fake.recordInvocation("Edit", []interface{}{arg1, arg2})
	fake.editMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserRepo) EditCallCount() int {
	fake.editMutex.RLock()
	defer fake.editMutex.RUnlock()
	return len(fake.editArgsForCall)
}

func (fake *FakeUserRepo) EditCalls(stub func(context.Context, *usera.User) error) {
	fake.editMutex.Lock()
	defer fake.editMutex.Unlock()
	fake.EditStub = stub
}

func (fake *FakeUserRepo) EditArgsForCall(i int) (context.Context, *usera.User) {
	fake.editMutex.RLock()
	defer fake.editMutex.RUnlock()
	argsForCall := fake.editArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserRepo) EditReturns(result1 error) {
	fake.editMutex.Lock()
	defer fake.editMutex.Unlock()
	fake.EditStub = nil
	fake.editReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) EditReturnsOnCall(i int, result1 error) {
	fake.editMutex.Lock()
	defer fake.editMutex.Unlock()
	fake.EditStub = nil
	if fake.editReturnsOnCall == nil {
		fake.editReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.editReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) User(arg1 context.Context, arg2 string) (*usera.User, error) {
	fake.userMutex.Lock()
	ret, specificReturn := fake.userReturnsOnCall[len(fake.userArgsForCall)]
	fake.userArgsForCall = append(fake.userArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.UserStub
	fakeReturns := fake.userReturns
	fake.recordInvocation("User", []interface{}{arg1, arg2})
	fake.userMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserRepo) UserCallCount() int {
	fake.userMutex.RLock()
	defer fake.userMutex.RUnlock()
	return len(fake.userArgsForCall)
}

func (fake *FakeUserRepo) UserCalls(stub func(context.Context, string) (*usera.User, error)) {
	fake.userMutex.Lock()
	defer fake.userMutex.Unlock()
	fake.UserStub = stub
}

func (fake *FakeUserRepo) UserArgsForCall(i int) (context.Context, string) {
	fake.userMutex.RLock()
	defer fake.userMutex.RUnlock()
	argsForCall := fake.userArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserRepo) UserReturns(result1 *usera.User, result2 error) {
	fake.userMutex.Lock()
	defer fake.userMutex.Unlock()
	fake.UserStub = nil
	fake.userReturns = struct {
		result1 *usera.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) UserReturnsOnCall(i int, result1 *usera.User, result2 error) {
	fake.userMutex.Lock()
	defer fake.userMutex.Unlock()
	fake.UserStub = nil
	if fake.userReturnsOnCall == nil {
		fake.userReturnsOnCall = make(map[int]struct {
			result1 *usera.User
			result2 error
		})
	}
	fake.userReturnsOnCall[i] = struct {
		result1 *usera.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) Users(arg1 context.Context) ([]*usera.User, error) {
	fake.usersMutex.Lock()
	ret, specificReturn := fake.usersReturnsOnCall[len(fake.usersArgsForCall)]
	fake.usersArgsForCall = append(fake.usersArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.UsersStub
	fakeReturns := fake.usersReturns
	fake.recordInvocation("Users", []interface{}{arg1})
	fake.usersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserRepo) UsersCallCount() int {
	fake.usersMutex.RLock()
	defer fake.usersMutex.RUnlock()
	return len(fake.usersArgsForCall)
}

func (fake *FakeUserRepo) UsersCalls(stub func(context.Context) ([]*usera.User, error)) {
	fake.usersMutex.Lock()
	defer fake.usersMutex.Unlock()
	fake.UsersStub = stub
}

func (fake *FakeUserRepo) UsersArgsForCall(i int) context.Context {
	fake.usersMutex.RLock()
	defer fake.usersMutex.RUnlock()
	argsForCall := fake.usersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserRepo) UsersReturns(result1 []*usera.User, result2 error) {
	fake.usersMutex.Lock()
	defer fake.usersMutex.Unlock()
	fake.UsersStub = nil
	fake.usersReturns = struct {
		result1 []*usera.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) UsersReturnsOnCall(i int, result1 []*usera.User, result2 error) {
	fake.usersMutex.Lock()
	defer fake.usersMutex.Unlock()
	fake.UsersStub = nil
	if fake.usersReturnsOnCall == nil {
		fake.usersReturnsOnCall = make(map[int]struct {
			result1 []*usera.User
			result2 error
		})
	}
	fake.usersReturnsOnCall[i] = struct {
		result1 []*usera.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.editMutex.RLock()
	defer fake.editMutex.RUnlock()
	fake.userMutex.RLock()
	defer fake.userMutex.RUnlock()
	fake.usersMutex.RLock()
	defer fake.usersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserRepo) recordInvocation(key string, args []interface{}) {
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

var _ user.UserRepo = new(FakeUserRepo)
