package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func testCreateRequest(handler http.HandlerFunc) *httptest.ResponseRecorder {
	request, _ := http.NewRequest("POST", "/task", nil)
	request.PostForm = url.Values{
		"description": {"Example task"},
	}
	response := httptest.NewRecorder()
	handler(response, request)
	return response
}

func testToogleRequest(id uint, handler http.HandlerFunc) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(
		"PUT",
		fmt.Sprintf("/task?id=%d", id),
		nil,
	)
	response := httptest.NewRecorder()
	handler(response, request)
	return response
}


func TestAPICreateSuccess(t *testing.T) {
	taskRepo := &TaskItemMemoryRepository{}
	api := TaskAPI{
		TaskItemRepository: taskRepo,
	}
	response := testCreateRequest(api.CreateHandler)
	assert.Equal(t, response.Code, http.StatusCreated)
	task := TaskItem{}
	err := json.Unmarshal(response.Body.Bytes(), &task)
	assert.NoError(t, err)
	_, err = taskRepo.FindOne(task.ID)
	assert.NoError(t, err)
}

func TestAPICreateFailure(t *testing.T) {
	taskRepo := &MockErrorTaskItemRepository{
		TaskItemRepository: &TaskItemMemoryRepository{},
		SaveError:          errors.New("Mocked save error"),
	}
	api := TaskAPI{
		TaskItemRepository: taskRepo,
	}
	response := testCreateRequest(api.CreateHandler)
	assert.Equal(t, response.Code, http.StatusBadRequest)
	assert.Equal(
		t,
		fmt.Sprintf("%s\n", taskRepo.SaveError),
		response.Body.String(),
	)
}

func TestAPIToogleSuccess(t *testing.T) {
	taskRepo := &TaskItemMemoryRepository{}
	api := TaskAPI{
		TaskItemRepository: taskRepo,
	}
	task := TaskItem{
		Description: "Example task item",
	}
	taskRepo.Save(&task)
	response := testToogleRequest(task.ID, api.ToogleHandler)
	assert.Equal(t, response.Code, http.StatusAccepted)
	updatedTask := TaskItem{}
	err := json.Unmarshal(response.Body.Bytes(), &updatedTask)
	assert.NoError(t, err)
	assert.True(t, updatedTask.IsComplete())
	assert.WithinDuration(
		t,
		time.Now(),
		updatedTask.CompletedAt,
		time.Duration(time.Second),
	)
	persistedTask, err := taskRepo.FindOne(updatedTask.ID)
	assert.NoError(t, err)
	assert.True(
		t,
		updatedTask.CompletedAt.Equal(persistedTask.CompletedAt),
	)
}
