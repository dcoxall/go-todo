package todo

import (
	"encoding/json"
	"net/http"
	"time"
	"strconv"
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

func (api *TaskAPI) ToogleHandler(w http.ResponseWriter, req *http.Request) {
	var id uint
	if rawID, err := strconv.Atoi(req.URL.Query().Get("id")); err != nil {
		http.NotFound(w, req)
	} else {
		id = uint(rawID)
	}
	if task, err := api.TaskItemRepository.FindOne(id); err == nil {
		task.CompletedAt = time.Now()
		if err = api.TaskItemRepository.Save(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			outs, _ := json.Marshal(&task)
			w.WriteHeader(http.StatusAccepted)
			w.Write(outs)
		}
	} else {
		http.NotFound(w, req)
	}
}
