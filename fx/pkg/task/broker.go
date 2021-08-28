package task

import (
	"encoding/json"
	"fmt"
	"fx/core"

	"fx/models"
	"fx/pkg/e"
	"fx/pkg/logging"

	"github.com/spf13/cast"
)

func Size() int64 {
	return models.Size()
}

func Pop() []TaskDetail {
	queues := models.Pop()
	details := []TaskDetail{}
	for _, queue := range queues {
		config := map[string]string{}
		if err := json.Unmarshal([]byte(queue.Content), &config); err == nil {
			config["type"] = queue.Category
			config["id"] = cast.ToString(queue.ID)
			fmt.Println("config:", config)
			details = append(details, TaskDetail{fn: call, params: config})
		} else {
			logging.Error(err.Error())
		}
	}
	return details
}

func call(config map[string]string) {

	report := core.XFind(config["type"], config)
	if content, err := json.Marshal(report); err == nil {
		models.Report(config["id"], string(content), e.GetMsg(e.SUCCESS), 3)
	} else {
		models.Report(config["id"], "", e.GetMsg(e.ERROR_FX), 4)
	}
}
