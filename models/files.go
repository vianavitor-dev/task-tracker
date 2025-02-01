package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Files struct {
	PathName string
}

func (f *Files) ReadFile() ([]byte, error) {

	file, err := os.Open(f.PathName)
	if err != nil {
		return nil, fmt.Errorf("ReadFile %v", err)
	}

	defer file.Close()

	bytesValues, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ReadFile %v", err)
	}

	return bytesValues, nil
}

func (f *Files) FileToTasks(tasks *[]Task) error {

	bytesValues, err := f.ReadFile()
	if err != nil {
		return fmt.Errorf("FileToTasks %v", err)
	}

	if err := json.Unmarshal(bytesValues, &tasks); err != nil {
		return fmt.Errorf("FileToTasks %v", err)
	}

	return nil
}

func (f *Files) AppendFile(data Task) (int, error) {

	var tasks = []Task{}

	file, err := os.OpenFile(f.PathName, os.O_CREATE, os.ModeAppend)
	if err != nil {
		return 0, fmt.Errorf("AppendFile %v", err)
	}

	defer file.Close()

	if err := f.FileToTasks(&tasks); err != nil && len(tasks) > 0 {
		return 0, fmt.Errorf("AppendFile %v", err)
	}

	if data.ID == 0 {
		data.ID = len(tasks) + 1
	}

	tasks = append(tasks, data)

	jsonTasks, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return 0, fmt.Errorf("AppendFile %v", err)
	}

	if _, err := file.Write(jsonTasks); err != nil {
		return 0, fmt.Errorf("AppendFile %v", err)
	}

	return data.ID, nil
}

func (f *Files) TruncadeTask(newTasks []Task) error {

	file, err := os.Create(f.PathName)
	if err != nil {
		return fmt.Errorf("TruncadeTask %v", err)
	}

	defer file.Close()

	jsonNewTasks, err := json.MarshalIndent(newTasks, "", " ")
	if err != nil {
		fmt.Print(fmt.Errorf("TruncadeTask : %v", err))
		os.Exit(0)
	}

	if len(newTasks) <= 0 {
		if _, err := file.Write(nil); err != nil {
			return fmt.Errorf("TruncadeTask : %v", err)
		}
	} else {
		if _, err := file.Write(jsonNewTasks); err != nil {
			return fmt.Errorf("TruncadeTask : %v", err)
		}
	}

	return nil
}
