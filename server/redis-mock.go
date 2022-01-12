package main

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/juev/counter/server.IRedis -o ./redis-mock.go -n RedisMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/go-redis/redis/v8"
	"github.com/gojuno/minimock/v3"
)

// RedisMock implements IRedis
type RedisMock struct {
	t minimock.Tester

	funcDel          func(ctx context.Context, keys ...string) (ip1 *redis.IntCmd)
	inspectFuncDel   func(ctx context.Context, keys ...string)
	afterDelCounter  uint64
	beforeDelCounter uint64
	DelMock          mRedisMockDel

	funcGet          func(ctx context.Context, key string) (sp1 *redis.StringCmd)
	inspectFuncGet   func(ctx context.Context, key string)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mRedisMockGet

	funcIncr          func(ctx context.Context, key string) (ip1 *redis.IntCmd)
	inspectFuncIncr   func(ctx context.Context, key string)
	afterIncrCounter  uint64
	beforeIncrCounter uint64
	IncrMock          mRedisMockIncr

	funcSet          func(ctx context.Context, key string, value interface{}) (sp1 *redis.StatusCmd)
	inspectFuncSet   func(ctx context.Context, key string, value interface{})
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mRedisMockSet
}

// NewRedisMock returns a mock for IRedis
func NewRedisMock(t minimock.Tester) *RedisMock {
	m := &RedisMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DelMock = mRedisMockDel{mock: m}
	m.DelMock.callArgs = []*RedisMockDelParams{}

	m.GetMock = mRedisMockGet{mock: m}
	m.GetMock.callArgs = []*RedisMockGetParams{}

	m.IncrMock = mRedisMockIncr{mock: m}
	m.IncrMock.callArgs = []*RedisMockIncrParams{}

	m.SetMock = mRedisMockSet{mock: m}
	m.SetMock.callArgs = []*RedisMockSetParams{}

	return m
}

type mRedisMockDel struct {
	mock               *RedisMock
	defaultExpectation *RedisMockDelExpectation
	expectations       []*RedisMockDelExpectation

	callArgs []*RedisMockDelParams
	mutex    sync.RWMutex
}

// RedisMockDelExpectation specifies expectation struct of the IRedis.Del
type RedisMockDelExpectation struct {
	mock    *RedisMock
	params  *RedisMockDelParams
	results *RedisMockDelResults
	Counter uint64
}

// RedisMockDelParams contains parameters of the IRedis.Del
type RedisMockDelParams struct {
	ctx  context.Context
	keys []string
}

// RedisMockDelResults contains results of the IRedis.Del
type RedisMockDelResults struct {
	ip1 *redis.IntCmd
}

// Expect sets up expected params for IRedis.Del
func (mmDel *mRedisMockDel) Expect(ctx context.Context, keys ...string) *mRedisMockDel {
	if mmDel.mock.funcDel != nil {
		mmDel.mock.t.Fatalf("RedisMock.Del mock is already set by Set")
	}

	if mmDel.defaultExpectation == nil {
		mmDel.defaultExpectation = &RedisMockDelExpectation{}
	}

	mmDel.defaultExpectation.params = &RedisMockDelParams{ctx, keys}
	for _, e := range mmDel.expectations {
		if minimock.Equal(e.params, mmDel.defaultExpectation.params) {
			mmDel.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDel.defaultExpectation.params)
		}
	}

	return mmDel
}

// Inspect accepts an inspector function that has same arguments as the IRedis.Del
func (mmDel *mRedisMockDel) Inspect(f func(ctx context.Context, keys ...string)) *mRedisMockDel {
	if mmDel.mock.inspectFuncDel != nil {
		mmDel.mock.t.Fatalf("Inspect function is already set for RedisMock.Del")
	}

	mmDel.mock.inspectFuncDel = f

	return mmDel
}

// Return sets up results that will be returned by IRedis.Del
func (mmDel *mRedisMockDel) Return(ip1 *redis.IntCmd) *RedisMock {
	if mmDel.mock.funcDel != nil {
		mmDel.mock.t.Fatalf("RedisMock.Del mock is already set by Set")
	}

	if mmDel.defaultExpectation == nil {
		mmDel.defaultExpectation = &RedisMockDelExpectation{mock: mmDel.mock}
	}
	mmDel.defaultExpectation.results = &RedisMockDelResults{ip1}
	return mmDel.mock
}

//Set uses given function f to mock the IRedis.Del method
func (mmDel *mRedisMockDel) Set(f func(ctx context.Context, keys ...string) (ip1 *redis.IntCmd)) *RedisMock {
	if mmDel.defaultExpectation != nil {
		mmDel.mock.t.Fatalf("Default expectation is already set for the IRedis.Del method")
	}

	if len(mmDel.expectations) > 0 {
		mmDel.mock.t.Fatalf("Some expectations are already set for the IRedis.Del method")
	}

	mmDel.mock.funcDel = f
	return mmDel.mock
}

// When sets expectation for the IRedis.Del which will trigger the result defined by the following
// Then helper
func (mmDel *mRedisMockDel) When(ctx context.Context, keys ...string) *RedisMockDelExpectation {
	if mmDel.mock.funcDel != nil {
		mmDel.mock.t.Fatalf("RedisMock.Del mock is already set by Set")
	}

	expectation := &RedisMockDelExpectation{
		mock:   mmDel.mock,
		params: &RedisMockDelParams{ctx, keys},
	}
	mmDel.expectations = append(mmDel.expectations, expectation)
	return expectation
}

// Then sets up IRedis.Del return parameters for the expectation previously defined by the When method
func (e *RedisMockDelExpectation) Then(ip1 *redis.IntCmd) *RedisMock {
	e.results = &RedisMockDelResults{ip1}
	return e.mock
}

// Del implements IRedis
func (mmDel *RedisMock) Del(ctx context.Context, keys ...string) (ip1 *redis.IntCmd) {
	mm_atomic.AddUint64(&mmDel.beforeDelCounter, 1)
	defer mm_atomic.AddUint64(&mmDel.afterDelCounter, 1)

	if mmDel.inspectFuncDel != nil {
		mmDel.inspectFuncDel(ctx, keys...)
	}

	mm_params := &RedisMockDelParams{ctx, keys}

	// Record call args
	mmDel.DelMock.mutex.Lock()
	mmDel.DelMock.callArgs = append(mmDel.DelMock.callArgs, mm_params)
	mmDel.DelMock.mutex.Unlock()

	for _, e := range mmDel.DelMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ip1
		}
	}

	if mmDel.DelMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDel.DelMock.defaultExpectation.Counter, 1)
		mm_want := mmDel.DelMock.defaultExpectation.params
		mm_got := RedisMockDelParams{ctx, keys}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDel.t.Errorf("RedisMock.Del got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDel.DelMock.defaultExpectation.results
		if mm_results == nil {
			mmDel.t.Fatal("No results are set for the RedisMock.Del")
		}
		return (*mm_results).ip1
	}
	if mmDel.funcDel != nil {
		return mmDel.funcDel(ctx, keys...)
	}
	mmDel.t.Fatalf("Unexpected call to RedisMock.Del. %v %v", ctx, keys)
	return
}

// DelAfterCounter returns a count of finished RedisMock.Del invocations
func (mmDel *RedisMock) DelAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDel.afterDelCounter)
}

// DelBeforeCounter returns a count of RedisMock.Del invocations
func (mmDel *RedisMock) DelBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDel.beforeDelCounter)
}

// Calls returns a list of arguments used in each call to RedisMock.Del.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDel *mRedisMockDel) Calls() []*RedisMockDelParams {
	mmDel.mutex.RLock()

	argCopy := make([]*RedisMockDelParams, len(mmDel.callArgs))
	copy(argCopy, mmDel.callArgs)

	mmDel.mutex.RUnlock()

	return argCopy
}

// MinimockDelDone returns true if the count of the Del invocations corresponds
// the number of defined expectations
func (m *RedisMock) MinimockDelDone() bool {
	for _, e := range m.DelMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DelMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDelCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDel != nil && mm_atomic.LoadUint64(&m.afterDelCounter) < 1 {
		return false
	}
	return true
}

// MinimockDelInspect logs each unmet expectation
func (m *RedisMock) MinimockDelInspect() {
	for _, e := range m.DelMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RedisMock.Del with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DelMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDelCounter) < 1 {
		if m.DelMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RedisMock.Del")
		} else {
			m.t.Errorf("Expected call to RedisMock.Del with params: %#v", *m.DelMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDel != nil && mm_atomic.LoadUint64(&m.afterDelCounter) < 1 {
		m.t.Error("Expected call to RedisMock.Del")
	}
}

type mRedisMockGet struct {
	mock               *RedisMock
	defaultExpectation *RedisMockGetExpectation
	expectations       []*RedisMockGetExpectation

	callArgs []*RedisMockGetParams
	mutex    sync.RWMutex
}

// RedisMockGetExpectation specifies expectation struct of the IRedis.Get
type RedisMockGetExpectation struct {
	mock    *RedisMock
	params  *RedisMockGetParams
	results *RedisMockGetResults
	Counter uint64
}

// RedisMockGetParams contains parameters of the IRedis.Get
type RedisMockGetParams struct {
	ctx context.Context
	key string
}

// RedisMockGetResults contains results of the IRedis.Get
type RedisMockGetResults struct {
	sp1 *redis.StringCmd
}

// Expect sets up expected params for IRedis.Get
func (mmGet *mRedisMockGet) Expect(ctx context.Context, key string) *mRedisMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &RedisMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &RedisMockGetParams{ctx, key}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the IRedis.Get
func (mmGet *mRedisMockGet) Inspect(f func(ctx context.Context, key string)) *mRedisMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for RedisMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by IRedis.Get
func (mmGet *mRedisMockGet) Return(sp1 *redis.StringCmd) *RedisMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &RedisMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &RedisMockGetResults{sp1}
	return mmGet.mock
}

//Set uses given function f to mock the IRedis.Get method
func (mmGet *mRedisMockGet) Set(f func(ctx context.Context, key string) (sp1 *redis.StringCmd)) *RedisMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the IRedis.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the IRedis.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the IRedis.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mRedisMockGet) When(ctx context.Context, key string) *RedisMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("RedisMock.Get mock is already set by Set")
	}

	expectation := &RedisMockGetExpectation{
		mock:   mmGet.mock,
		params: &RedisMockGetParams{ctx, key},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up IRedis.Get return parameters for the expectation previously defined by the When method
func (e *RedisMockGetExpectation) Then(sp1 *redis.StringCmd) *RedisMock {
	e.results = &RedisMockGetResults{sp1}
	return e.mock
}

// Get implements IRedis
func (mmGet *RedisMock) Get(ctx context.Context, key string) (sp1 *redis.StringCmd) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, key)
	}

	mm_params := &RedisMockGetParams{ctx, key}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.sp1
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := RedisMockGetParams{ctx, key}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("RedisMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the RedisMock.Get")
		}
		return (*mm_results).sp1
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, key)
	}
	mmGet.t.Fatalf("Unexpected call to RedisMock.Get. %v %v", ctx, key)
	return
}

// GetAfterCounter returns a count of finished RedisMock.Get invocations
func (mmGet *RedisMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of RedisMock.Get invocations
func (mmGet *RedisMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to RedisMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mRedisMockGet) Calls() []*RedisMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*RedisMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *RedisMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *RedisMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RedisMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RedisMock.Get")
		} else {
			m.t.Errorf("Expected call to RedisMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to RedisMock.Get")
	}
}

type mRedisMockIncr struct {
	mock               *RedisMock
	defaultExpectation *RedisMockIncrExpectation
	expectations       []*RedisMockIncrExpectation

	callArgs []*RedisMockIncrParams
	mutex    sync.RWMutex
}

// RedisMockIncrExpectation specifies expectation struct of the IRedis.Incr
type RedisMockIncrExpectation struct {
	mock    *RedisMock
	params  *RedisMockIncrParams
	results *RedisMockIncrResults
	Counter uint64
}

// RedisMockIncrParams contains parameters of the IRedis.Incr
type RedisMockIncrParams struct {
	ctx context.Context
	key string
}

// RedisMockIncrResults contains results of the IRedis.Incr
type RedisMockIncrResults struct {
	ip1 *redis.IntCmd
}

// Expect sets up expected params for IRedis.Incr
func (mmIncr *mRedisMockIncr) Expect(ctx context.Context, key string) *mRedisMockIncr {
	if mmIncr.mock.funcIncr != nil {
		mmIncr.mock.t.Fatalf("RedisMock.Incr mock is already set by Set")
	}

	if mmIncr.defaultExpectation == nil {
		mmIncr.defaultExpectation = &RedisMockIncrExpectation{}
	}

	mmIncr.defaultExpectation.params = &RedisMockIncrParams{ctx, key}
	for _, e := range mmIncr.expectations {
		if minimock.Equal(e.params, mmIncr.defaultExpectation.params) {
			mmIncr.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmIncr.defaultExpectation.params)
		}
	}

	return mmIncr
}

// Inspect accepts an inspector function that has same arguments as the IRedis.Incr
func (mmIncr *mRedisMockIncr) Inspect(f func(ctx context.Context, key string)) *mRedisMockIncr {
	if mmIncr.mock.inspectFuncIncr != nil {
		mmIncr.mock.t.Fatalf("Inspect function is already set for RedisMock.Incr")
	}

	mmIncr.mock.inspectFuncIncr = f

	return mmIncr
}

// Return sets up results that will be returned by IRedis.Incr
func (mmIncr *mRedisMockIncr) Return(ip1 *redis.IntCmd) *RedisMock {
	if mmIncr.mock.funcIncr != nil {
		mmIncr.mock.t.Fatalf("RedisMock.Incr mock is already set by Set")
	}

	if mmIncr.defaultExpectation == nil {
		mmIncr.defaultExpectation = &RedisMockIncrExpectation{mock: mmIncr.mock}
	}
	mmIncr.defaultExpectation.results = &RedisMockIncrResults{ip1}
	return mmIncr.mock
}

//Set uses given function f to mock the IRedis.Incr method
func (mmIncr *mRedisMockIncr) Set(f func(ctx context.Context, key string) (ip1 *redis.IntCmd)) *RedisMock {
	if mmIncr.defaultExpectation != nil {
		mmIncr.mock.t.Fatalf("Default expectation is already set for the IRedis.Incr method")
	}

	if len(mmIncr.expectations) > 0 {
		mmIncr.mock.t.Fatalf("Some expectations are already set for the IRedis.Incr method")
	}

	mmIncr.mock.funcIncr = f
	return mmIncr.mock
}

// When sets expectation for the IRedis.Incr which will trigger the result defined by the following
// Then helper
func (mmIncr *mRedisMockIncr) When(ctx context.Context, key string) *RedisMockIncrExpectation {
	if mmIncr.mock.funcIncr != nil {
		mmIncr.mock.t.Fatalf("RedisMock.Incr mock is already set by Set")
	}

	expectation := &RedisMockIncrExpectation{
		mock:   mmIncr.mock,
		params: &RedisMockIncrParams{ctx, key},
	}
	mmIncr.expectations = append(mmIncr.expectations, expectation)
	return expectation
}

// Then sets up IRedis.Incr return parameters for the expectation previously defined by the When method
func (e *RedisMockIncrExpectation) Then(ip1 *redis.IntCmd) *RedisMock {
	e.results = &RedisMockIncrResults{ip1}
	return e.mock
}

// Incr implements IRedis
func (mmIncr *RedisMock) Incr(ctx context.Context, key string) (ip1 *redis.IntCmd) {
	mm_atomic.AddUint64(&mmIncr.beforeIncrCounter, 1)
	defer mm_atomic.AddUint64(&mmIncr.afterIncrCounter, 1)

	if mmIncr.inspectFuncIncr != nil {
		mmIncr.inspectFuncIncr(ctx, key)
	}

	mm_params := &RedisMockIncrParams{ctx, key}

	// Record call args
	mmIncr.IncrMock.mutex.Lock()
	mmIncr.IncrMock.callArgs = append(mmIncr.IncrMock.callArgs, mm_params)
	mmIncr.IncrMock.mutex.Unlock()

	for _, e := range mmIncr.IncrMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ip1
		}
	}

	if mmIncr.IncrMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmIncr.IncrMock.defaultExpectation.Counter, 1)
		mm_want := mmIncr.IncrMock.defaultExpectation.params
		mm_got := RedisMockIncrParams{ctx, key}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmIncr.t.Errorf("RedisMock.Incr got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmIncr.IncrMock.defaultExpectation.results
		if mm_results == nil {
			mmIncr.t.Fatal("No results are set for the RedisMock.Incr")
		}
		return (*mm_results).ip1
	}
	if mmIncr.funcIncr != nil {
		return mmIncr.funcIncr(ctx, key)
	}
	mmIncr.t.Fatalf("Unexpected call to RedisMock.Incr. %v %v", ctx, key)
	return
}

// IncrAfterCounter returns a count of finished RedisMock.Incr invocations
func (mmIncr *RedisMock) IncrAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIncr.afterIncrCounter)
}

// IncrBeforeCounter returns a count of RedisMock.Incr invocations
func (mmIncr *RedisMock) IncrBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmIncr.beforeIncrCounter)
}

// Calls returns a list of arguments used in each call to RedisMock.Incr.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmIncr *mRedisMockIncr) Calls() []*RedisMockIncrParams {
	mmIncr.mutex.RLock()

	argCopy := make([]*RedisMockIncrParams, len(mmIncr.callArgs))
	copy(argCopy, mmIncr.callArgs)

	mmIncr.mutex.RUnlock()

	return argCopy
}

// MinimockIncrDone returns true if the count of the Incr invocations corresponds
// the number of defined expectations
func (m *RedisMock) MinimockIncrDone() bool {
	for _, e := range m.IncrMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IncrMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIncrCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIncr != nil && mm_atomic.LoadUint64(&m.afterIncrCounter) < 1 {
		return false
	}
	return true
}

// MinimockIncrInspect logs each unmet expectation
func (m *RedisMock) MinimockIncrInspect() {
	for _, e := range m.IncrMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RedisMock.Incr with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.IncrMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterIncrCounter) < 1 {
		if m.IncrMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RedisMock.Incr")
		} else {
			m.t.Errorf("Expected call to RedisMock.Incr with params: %#v", *m.IncrMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcIncr != nil && mm_atomic.LoadUint64(&m.afterIncrCounter) < 1 {
		m.t.Error("Expected call to RedisMock.Incr")
	}
}

type mRedisMockSet struct {
	mock               *RedisMock
	defaultExpectation *RedisMockSetExpectation
	expectations       []*RedisMockSetExpectation

	callArgs []*RedisMockSetParams
	mutex    sync.RWMutex
}

// RedisMockSetExpectation specifies expectation struct of the IRedis.Set
type RedisMockSetExpectation struct {
	mock    *RedisMock
	params  *RedisMockSetParams
	results *RedisMockSetResults
	Counter uint64
}

// RedisMockSetParams contains parameters of the IRedis.Set
type RedisMockSetParams struct {
	ctx   context.Context
	key   string
	value interface{}
}

// RedisMockSetResults contains results of the IRedis.Set
type RedisMockSetResults struct {
	sp1 *redis.StatusCmd
}

// Expect sets up expected params for IRedis.Set
func (mmSet *mRedisMockSet) Expect(ctx context.Context, key string, value interface{}) *mRedisMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RedisMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &RedisMockSetParams{ctx, key, value}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the IRedis.Set
func (mmSet *mRedisMockSet) Inspect(f func(ctx context.Context, key string, value interface{})) *mRedisMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for RedisMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by IRedis.Set
func (mmSet *mRedisMockSet) Return(sp1 *redis.StatusCmd) *RedisMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &RedisMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &RedisMockSetResults{sp1}
	return mmSet.mock
}

//Set uses given function f to mock the IRedis.Set method
func (mmSet *mRedisMockSet) Set(f func(ctx context.Context, key string, value interface{}) (sp1 *redis.StatusCmd)) *RedisMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the IRedis.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the IRedis.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// When sets expectation for the IRedis.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mRedisMockSet) When(ctx context.Context, key string, value interface{}) *RedisMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("RedisMock.Set mock is already set by Set")
	}

	expectation := &RedisMockSetExpectation{
		mock:   mmSet.mock,
		params: &RedisMockSetParams{ctx, key, value},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up IRedis.Set return parameters for the expectation previously defined by the When method
func (e *RedisMockSetExpectation) Then(sp1 *redis.StatusCmd) *RedisMock {
	e.results = &RedisMockSetResults{sp1}
	return e.mock
}

// Set implements IRedis
func (mmSet *RedisMock) Set(ctx context.Context, key string, value interface{}) (sp1 *redis.StatusCmd) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(ctx, key, value)
	}

	mm_params := &RedisMockSetParams{ctx, key, value}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, mm_params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.sp1
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		mm_want := mmSet.SetMock.defaultExpectation.params
		mm_got := RedisMockSetParams{ctx, key, value}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSet.t.Errorf("RedisMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSet.SetMock.defaultExpectation.results
		if mm_results == nil {
			mmSet.t.Fatal("No results are set for the RedisMock.Set")
		}
		return (*mm_results).sp1
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(ctx, key, value)
	}
	mmSet.t.Fatalf("Unexpected call to RedisMock.Set. %v %v %v", ctx, key, value)
	return
}

// SetAfterCounter returns a count of finished RedisMock.Set invocations
func (mmSet *RedisMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of RedisMock.Set invocations
func (mmSet *RedisMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to RedisMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mRedisMockSet) Calls() []*RedisMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*RedisMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *RedisMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *RedisMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RedisMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RedisMock.Set")
		} else {
			m.t.Errorf("Expected call to RedisMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to RedisMock.Set")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RedisMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDelInspect()

		m.MinimockGetInspect()

		m.MinimockIncrInspect()

		m.MinimockSetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RedisMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RedisMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDelDone() &&
		m.MinimockGetDone() &&
		m.MinimockIncrDone() &&
		m.MinimockSetDone()
}
