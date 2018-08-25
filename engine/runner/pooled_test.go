/*
Sniperkit-Bot
- Status: analyzed
*/

package runner

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/action"
	"github.com/sniperkit/snk.fork.tibcosoftware-flogo-lib/core/data"
)

type MockFullAction struct {
	mock.Mock
}

func (m *MockFullAction) Config() *action.Config {
	return nil
}

func (m *MockFullAction) Metadata() *action.Metadata {
	return nil
}

func (m *MockFullAction) IOMetadata() *data.IOMetadata {
	return nil
}

func (m *MockFullAction) Run(context context.Context, inputs map[string]*data.Attribute, handler action.ResultHandler) error {
	args := m.Called(context, inputs, handler)
	return args.Error(0)
}

// This mock action will handle the result and mark it done
type MockResultAction struct {
	mock.Mock
}

func (m *MockResultAction) Config() *action.Config {
	return nil
}

func (m *MockResultAction) Metadata() *action.Metadata {
	return nil
}

func (m *MockResultAction) IOMetadata() *data.IOMetadata {
	return nil
}

func (m *MockResultAction) Run(context context.Context, inputs map[string]*data.Attribute, handler action.ResultHandler) error {
	args := m.Called(context, inputs, handler)
	go func() {
		dataAttr, _ := data.NewAttribute("data", data.TypeString, "mock")
		codeAttr, _ := data.NewAttribute("code", data.TypeInteger, 200)
		resultData := map[string]*data.Attribute{
			"data": dataAttr,
			"code": codeAttr,
		}
		handler.HandleResult(resultData, nil)
		handler.Done()
	}()
	return args.Error(0)
}

// TestNewPooledOk test creation of new Pooled runner
func TestNewPooledOk(t *testing.T) {
	config := &PooledConfig{NumWorkers: 1, WorkQueueSize: 1}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
}

// TestStartOk test that Start method is fine
func TestStartOk(t *testing.T) {
	config := &PooledConfig{NumWorkers: 3, WorkQueueSize: 3}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
	// It should have a worker queue of the size expected
	assert.Equal(t, 3, cap(runner.workerQueue))
	// It should have a workers of the expected size
	assert.Equal(t, 3, len(runner.workers))
	// Runner should be active
	assert.True(t, runner.active)
}

// TestRunNilError test that running a nil action trows and error
func TestRunNilError(t *testing.T) {
	config := &PooledConfig{NumWorkers: 5, WorkQueueSize: 5}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
	_, err = runner.Execute(nil, nil, nil)
	assert.NotNil(t, err)
}

// TestRunInnactiveError test that running an innactive runner trows and error
func TestRunInnactiveError(t *testing.T) {
	config := &PooledConfig{NumWorkers: 5, WorkQueueSize: 5}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	a := new(MockFullAction)
	_, err := runner.Execute(nil, a, nil)
	assert.NotNil(t, err)
}

// TestRunErrorInAction test that running an action returns an error
func TestRunErrorInAction(t *testing.T) {
	config := &PooledConfig{NumWorkers: 5, WorkQueueSize: 5}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
	a := new(MockFullAction)
	a.On("Run", nil, mock.AnythingOfType("map[string]*data.Attribute"), mock.AnythingOfType("*runner.AsyncResultHandler")).Return(errors.New("Error in action"))
	_, err = runner.Execute(nil, a, nil)
	assert.NotNil(t, err)
	assert.Equal(t, "Error in action", err.Error())
}

// TestRunOk test that running an action is ok
func TestRunOk(t *testing.T) {
	config := &PooledConfig{NumWorkers: 5, WorkQueueSize: 5}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
	a := new(MockResultAction)
	a.On("Run", nil, mock.AnythingOfType("map[string]*data.Attribute"), mock.AnythingOfType("*runner.AsyncResultHandler")).Return(nil)
	results, err := runner.Execute(nil, a, nil)
	assert.Nil(t, err)
	codeAttr := results["code"]
	assert.NotNil(t, codeAttr)
	assert.Equal(t, 200, codeAttr.Value())
	dataAttr := results["data"]
	assert.NotNil(t, dataAttr)
	assert.Equal(t, "mock", dataAttr.Value())
}

// TestStopOk test that Stop method is fine
func TestStopOk(t *testing.T) {
	config := &PooledConfig{NumWorkers: 3, WorkQueueSize: 3}
	runner := NewPooled(config)
	assert.NotNil(t, runner)
	err := runner.Start()
	assert.Nil(t, err)
	err = runner.Stop()
	assert.Nil(t, err)
	assert.False(t, runner.active)

}
