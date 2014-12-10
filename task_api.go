package todo

import (
	"encoding/json"
	"net/http"
)

type TaskAPI struct {
	TaskItemRepository TaskItemRepository
}

func (api *TaskAPI) CreateHandler(w http.ResponseWriter, req *http.Request) {
	task := TaskItem{
		Description: req.PostFormValue("description"),
	}
	if err := api.TaskItemRepository.Save(&task); err == nil {
		w.WriteHeader(http.StatusCreated)
		outs, _ := json.Marshal(&task)
		w.Write(outs)
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
