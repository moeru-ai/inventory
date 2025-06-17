package models

import (
	"sync"

	v1 "github.com/moeru-ai/inventory/apis/inventoryapi/v1"
)

var (
	commonTasks = &ModelList{
		textsFeatureExtraction: make([]*v1.GetModelsModelItem, 0),
	}
)

func Models() []*v1.GetModelsModelItem {
	return commonTasks.Models()
}

type ModelList struct {
	sync.RWMutex

	textsFeatureExtraction []*v1.GetModelsModelItem
}

func (l *ModelList) Models() []*v1.GetModelsModelItem {
	l.RLock()
	defer l.RUnlock()

	return l.textsFeatureExtraction
}

func (l *ModelList) Add(model *v1.GetModelsModelItem) {
	l.Lock()
	defer l.Unlock()

	l.textsFeatureExtraction = append(l.textsFeatureExtraction, model)
}

func (l *ModelList) Find(modelID string) (*v1.GetModelsModelItem, bool) {
	l.RLock()
	defer l.RUnlock()

	for _, model := range l.textsFeatureExtraction {
		if model.Id == modelID {
			return model, true
		}
	}

	return nil, false
}
