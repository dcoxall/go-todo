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
)

func testCreateRequest(
	handler http.HandlerFunc,
	tests func(*httptest.ResponseRecorder)) {
	request, _ := http.NewRequest("POST", "/task", nil)
	request.PostForm = url.Values{
		"description": {"Example task"},
	}
	response := httptest.NewRecorder()
	handler(response, request)
	tests(response)
}

func TestAPICreateSuccess(t *testing.T) {
	taskRepo := &TaskItemMemoryRepository{}
	api := TaskAPI{
		TaskItemRepository: taskRepo,
	}
	testCreateRequest(
		api.CreateHandler,
		func(response *httptest.ResponseRecorder) {
			assert.Equal(t, response.Code, http.StatusCreated)
			task := TaskItem{}
			err := json.Unmarshal(response.Body.Bytes(), &task)
			assert.NoError(t, err)
			_, err = taskRepo.FindOne(task.ID)
			assert.NoError(t, err)
		},
	)
}

func TestAPICreateFailure(t *testing.T) {
	taskRepo := &MockErrorTaskItemRepository{
		TaskItemRepository: &TaskItemMemoryRepository{},
		SaveError:          errors.New("Mocked save error"),
	}
	api := TaskAPI{
		TaskItemRepository: taskRepo,
	}
	testCreateRequest(
		api.CreateHandler,
		func(response *httptest.ResponseRecorder) {
			assert.Equal(t, response.Code, http.StatusBadRequest)
			assert.Equal(
				t,
				fmt.Sprintf("%s\n", taskRepo.SaveError),
				response.Body.String(),
			)
		},
	)
}
