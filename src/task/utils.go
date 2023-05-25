package task

import (
	"encoding/json"
	"io/ioutil"
)

func LoadTasks() []Task {
	file, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return []Task{}
	}

	return tasks
}

func SaveTasks(tasks []Task) {
	file, _ := json.MarshalIndent(tasks, "", "    ")
	_ = ioutil.WriteFile("tasks.json", file, 0644)
}

func Atoi(s string) int {
	n := 0
	for _, ch := range []byte(s) {
		n = n*10 + int(ch-'0')
	}
	return n
}
