// Code generated by counterfeiter. DO NOT EDIT.
package v7actionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/uaa"
	"code.cloudfoundry.org/cli/api/uaa/constant"
)

type FakeUAAClient struct {
	APIVersionStub        func() string
	aPIVersionMutex       sync.RWMutex
	aPIVersionArgsForCall []struct {
	}
	aPIVersionReturns struct {
		result1 string
	}
	aPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	AuthenticateStub        func(map[string]string, string, constant.GrantType) (string, string, error)
	authenticateMutex       sync.RWMutex
	authenticateArgsForCall []struct {
		arg1 map[string]string
		arg2 string
		arg3 constant.GrantType
	}
	authenticateReturns struct {
		result1 string
		result2 string
		result3 error
	}
	authenticateReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	CreateUserStub        func(string, string, string) (uaa.User, error)
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createUserReturns struct {
		result1 uaa.User
		result2 error
	}
	createUserReturnsOnCall map[int]struct {
		result1 uaa.User
		result2 error
	}
	DeleteUserStub        func(string) (uaa.User, error)
	deleteUserMutex       sync.RWMutex
	deleteUserArgsForCall []struct {
		arg1 string
	}
	deleteUserReturns struct {
		result1 uaa.User
		result2 error
	}
	deleteUserReturnsOnCall map[int]struct {
		result1 uaa.User
		result2 error
	}
	GetSSHPasscodeStub        func(string, string) (string, error)
	getSSHPasscodeMutex       sync.RWMutex
	getSSHPasscodeArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSSHPasscodeReturns struct {
		result1 string
		result2 error
	}
	getSSHPasscodeReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ListUsersStub        func(string, string) ([]uaa.User, error)
	listUsersMutex       sync.RWMutex
	listUsersArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listUsersReturns struct {
		result1 []uaa.User
		result2 error
	}
	listUsersReturnsOnCall map[int]struct {
		result1 []uaa.User
		result2 error
	}
	RefreshAccessTokenStub        func(string) (uaa.RefreshedTokens, error)
	refreshAccessTokenMutex       sync.RWMutex
	refreshAccessTokenArgsForCall []struct {
		arg1 string
	}
	refreshAccessTokenReturns struct {
		result1 uaa.RefreshedTokens
		result2 error
	}
	refreshAccessTokenReturnsOnCall map[int]struct {
		result1 uaa.RefreshedTokens
		result2 error
	}
	ValidateClientUserStub        func(string) error
	validateClientUserMutex       sync.RWMutex
	validateClientUserArgsForCall []struct {
		arg1 string
	}
	validateClientUserReturns struct {
		result1 error
	}
	validateClientUserReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUAAClient) APIVersion() string {
	fake.aPIVersionMutex.Lock()
	ret, specificReturn := fake.aPIVersionReturnsOnCall[len(fake.aPIVersionArgsForCall)]
	fake.aPIVersionArgsForCall = append(fake.aPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("APIVersion", []interface{}{})
	fake.aPIVersionMutex.Unlock()
	if fake.APIVersionStub != nil {
		return fake.APIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.aPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeUAAClient) APIVersionCallCount() int {
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	return len(fake.aPIVersionArgsForCall)
}

func (fake *FakeUAAClient) APIVersionCalls(stub func() string) {
	fake.aPIVersionMutex.Lock()
	defer fake.aPIVersionMutex.Unlock()
	fake.APIVersionStub = stub
}

func (fake *FakeUAAClient) APIVersionReturns(result1 string) {
	fake.aPIVersionMutex.Lock()
	defer fake.aPIVersionMutex.Unlock()
	fake.APIVersionStub = nil
	fake.aPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeUAAClient) APIVersionReturnsOnCall(i int, result1 string) {
	fake.aPIVersionMutex.Lock()
	defer fake.aPIVersionMutex.Unlock()
	fake.APIVersionStub = nil
	if fake.aPIVersionReturnsOnCall == nil {
		fake.aPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.aPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeUAAClient) Authenticate(arg1 map[string]string, arg2 string, arg3 constant.GrantType) (string, string, error) {
	fake.authenticateMutex.Lock()
	ret, specificReturn := fake.authenticateReturnsOnCall[len(fake.authenticateArgsForCall)]
	fake.authenticateArgsForCall = append(fake.authenticateArgsForCall, struct {
		arg1 map[string]string
		arg2 string
		arg3 constant.GrantType
	}{arg1, arg2, arg3})
	fake.recordInvocation("Authenticate", []interface{}{arg1, arg2, arg3})
	fake.authenticateMutex.Unlock()
	if fake.AuthenticateStub != nil {
		return fake.AuthenticateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.authenticateReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeUAAClient) AuthenticateCallCount() int {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return len(fake.authenticateArgsForCall)
}

func (fake *FakeUAAClient) AuthenticateCalls(stub func(map[string]string, string, constant.GrantType) (string, string, error)) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = stub
}

func (fake *FakeUAAClient) AuthenticateArgsForCall(i int) (map[string]string, string, constant.GrantType) {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	argsForCall := fake.authenticateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUAAClient) AuthenticateReturns(result1 string, result2 string, result3 error) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = nil
	fake.authenticateReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUAAClient) AuthenticateReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = nil
	if fake.authenticateReturnsOnCall == nil {
		fake.authenticateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.authenticateReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeUAAClient) CreateUser(arg1 string, arg2 string, arg3 string) (uaa.User, error) {
	fake.createUserMutex.Lock()
	ret, specificReturn := fake.createUserReturnsOnCall[len(fake.createUserArgsForCall)]
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateUser", []interface{}{arg1, arg2, arg3})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeUAAClient) CreateUserCalls(stub func(string, string, string) (uaa.User, error)) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = stub
}

func (fake *FakeUAAClient) CreateUserArgsForCall(i int) (string, string, string) {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	argsForCall := fake.createUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUAAClient) CreateUserReturns(result1 uaa.User, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) CreateUserReturnsOnCall(i int, result1 uaa.User, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	if fake.createUserReturnsOnCall == nil {
		fake.createUserReturnsOnCall = make(map[int]struct {
			result1 uaa.User
			result2 error
		})
	}
	fake.createUserReturnsOnCall[i] = struct {
		result1 uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) DeleteUser(arg1 string) (uaa.User, error) {
	fake.deleteUserMutex.Lock()
	ret, specificReturn := fake.deleteUserReturnsOnCall[len(fake.deleteUserArgsForCall)]
	fake.deleteUserArgsForCall = append(fake.deleteUserArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteUser", []interface{}{arg1})
	fake.deleteUserMutex.Unlock()
	if fake.DeleteUserStub != nil {
		return fake.DeleteUserStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) DeleteUserCallCount() int {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return len(fake.deleteUserArgsForCall)
}

func (fake *FakeUAAClient) DeleteUserCalls(stub func(string) (uaa.User, error)) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = stub
}

func (fake *FakeUAAClient) DeleteUserArgsForCall(i int) string {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	argsForCall := fake.deleteUserArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUAAClient) DeleteUserReturns(result1 uaa.User, result2 error) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = nil
	fake.deleteUserReturns = struct {
		result1 uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) DeleteUserReturnsOnCall(i int, result1 uaa.User, result2 error) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = nil
	if fake.deleteUserReturnsOnCall == nil {
		fake.deleteUserReturnsOnCall = make(map[int]struct {
			result1 uaa.User
			result2 error
		})
	}
	fake.deleteUserReturnsOnCall[i] = struct {
		result1 uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) GetSSHPasscode(arg1 string, arg2 string) (string, error) {
	fake.getSSHPasscodeMutex.Lock()
	ret, specificReturn := fake.getSSHPasscodeReturnsOnCall[len(fake.getSSHPasscodeArgsForCall)]
	fake.getSSHPasscodeArgsForCall = append(fake.getSSHPasscodeArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSSHPasscode", []interface{}{arg1, arg2})
	fake.getSSHPasscodeMutex.Unlock()
	if fake.GetSSHPasscodeStub != nil {
		return fake.GetSSHPasscodeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSSHPasscodeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) GetSSHPasscodeCallCount() int {
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	return len(fake.getSSHPasscodeArgsForCall)
}

func (fake *FakeUAAClient) GetSSHPasscodeCalls(stub func(string, string) (string, error)) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = stub
}

func (fake *FakeUAAClient) GetSSHPasscodeArgsForCall(i int) (string, string) {
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	argsForCall := fake.getSSHPasscodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUAAClient) GetSSHPasscodeReturns(result1 string, result2 error) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = nil
	fake.getSSHPasscodeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) GetSSHPasscodeReturnsOnCall(i int, result1 string, result2 error) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = nil
	if fake.getSSHPasscodeReturnsOnCall == nil {
		fake.getSSHPasscodeReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getSSHPasscodeReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) ListUsers(arg1 string, arg2 string) ([]uaa.User, error) {
	fake.listUsersMutex.Lock()
	ret, specificReturn := fake.listUsersReturnsOnCall[len(fake.listUsersArgsForCall)]
	fake.listUsersArgsForCall = append(fake.listUsersArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ListUsers", []interface{}{arg1, arg2})
	fake.listUsersMutex.Unlock()
	if fake.ListUsersStub != nil {
		return fake.ListUsersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listUsersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) ListUsersCallCount() int {
	fake.listUsersMutex.RLock()
	defer fake.listUsersMutex.RUnlock()
	return len(fake.listUsersArgsForCall)
}

func (fake *FakeUAAClient) ListUsersCalls(stub func(string, string) ([]uaa.User, error)) {
	fake.listUsersMutex.Lock()
	defer fake.listUsersMutex.Unlock()
	fake.ListUsersStub = stub
}

func (fake *FakeUAAClient) ListUsersArgsForCall(i int) (string, string) {
	fake.listUsersMutex.RLock()
	defer fake.listUsersMutex.RUnlock()
	argsForCall := fake.listUsersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUAAClient) ListUsersReturns(result1 []uaa.User, result2 error) {
	fake.listUsersMutex.Lock()
	defer fake.listUsersMutex.Unlock()
	fake.ListUsersStub = nil
	fake.listUsersReturns = struct {
		result1 []uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) ListUsersReturnsOnCall(i int, result1 []uaa.User, result2 error) {
	fake.listUsersMutex.Lock()
	defer fake.listUsersMutex.Unlock()
	fake.ListUsersStub = nil
	if fake.listUsersReturnsOnCall == nil {
		fake.listUsersReturnsOnCall = make(map[int]struct {
			result1 []uaa.User
			result2 error
		})
	}
	fake.listUsersReturnsOnCall[i] = struct {
		result1 []uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) RefreshAccessToken(arg1 string) (uaa.RefreshedTokens, error) {
	fake.refreshAccessTokenMutex.Lock()
	ret, specificReturn := fake.refreshAccessTokenReturnsOnCall[len(fake.refreshAccessTokenArgsForCall)]
	fake.refreshAccessTokenArgsForCall = append(fake.refreshAccessTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RefreshAccessToken", []interface{}{arg1})
	fake.refreshAccessTokenMutex.Unlock()
	if fake.RefreshAccessTokenStub != nil {
		return fake.RefreshAccessTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.refreshAccessTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) RefreshAccessTokenCallCount() int {
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	return len(fake.refreshAccessTokenArgsForCall)
}

func (fake *FakeUAAClient) RefreshAccessTokenCalls(stub func(string) (uaa.RefreshedTokens, error)) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = stub
}

func (fake *FakeUAAClient) RefreshAccessTokenArgsForCall(i int) string {
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	argsForCall := fake.refreshAccessTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUAAClient) RefreshAccessTokenReturns(result1 uaa.RefreshedTokens, result2 error) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = nil
	fake.refreshAccessTokenReturns = struct {
		result1 uaa.RefreshedTokens
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) RefreshAccessTokenReturnsOnCall(i int, result1 uaa.RefreshedTokens, result2 error) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = nil
	if fake.refreshAccessTokenReturnsOnCall == nil {
		fake.refreshAccessTokenReturnsOnCall = make(map[int]struct {
			result1 uaa.RefreshedTokens
			result2 error
		})
	}
	fake.refreshAccessTokenReturnsOnCall[i] = struct {
		result1 uaa.RefreshedTokens
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) ValidateClientUser(arg1 string) error {
	fake.validateClientUserMutex.Lock()
	ret, specificReturn := fake.validateClientUserReturnsOnCall[len(fake.validateClientUserArgsForCall)]
	fake.validateClientUserArgsForCall = append(fake.validateClientUserArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ValidateClientUser", []interface{}{arg1})
	fake.validateClientUserMutex.Unlock()
	if fake.ValidateClientUserStub != nil {
		return fake.ValidateClientUserStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.validateClientUserReturns
	return fakeReturns.result1
}

func (fake *FakeUAAClient) ValidateClientUserCallCount() int {
	fake.validateClientUserMutex.RLock()
	defer fake.validateClientUserMutex.RUnlock()
	return len(fake.validateClientUserArgsForCall)
}

func (fake *FakeUAAClient) ValidateClientUserCalls(stub func(string) error) {
	fake.validateClientUserMutex.Lock()
	defer fake.validateClientUserMutex.Unlock()
	fake.ValidateClientUserStub = stub
}

func (fake *FakeUAAClient) ValidateClientUserArgsForCall(i int) string {
	fake.validateClientUserMutex.RLock()
	defer fake.validateClientUserMutex.RUnlock()
	argsForCall := fake.validateClientUserArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUAAClient) ValidateClientUserReturns(result1 error) {
	fake.validateClientUserMutex.Lock()
	defer fake.validateClientUserMutex.Unlock()
	fake.ValidateClientUserStub = nil
	fake.validateClientUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUAAClient) ValidateClientUserReturnsOnCall(i int, result1 error) {
	fake.validateClientUserMutex.Lock()
	defer fake.validateClientUserMutex.Unlock()
	fake.ValidateClientUserStub = nil
	if fake.validateClientUserReturnsOnCall == nil {
		fake.validateClientUserReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateClientUserReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.aPIVersionMutex.RLock()
	defer fake.aPIVersionMutex.RUnlock()
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	fake.listUsersMutex.RLock()
	defer fake.listUsersMutex.RUnlock()
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	fake.validateClientUserMutex.RLock()
	defer fake.validateClientUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUAAClient) recordInvocation(key string, args []interface{}) {
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

var _ v7action.UAAClient = new(FakeUAAClient)
