package testModule

import (
	"fmt"
	"sykros-pro/gopro/src/utils/helper"
)

type ModuleType int64

const (
	MERGE_STRUCT ModuleType = 0
)

type TestModuleImpl interface {
	Run(moduleType ModuleType)
}

type TestModule struct{ TestModuleImpl }

type addictionAlInterFace struct {
	id int
}

func (t *TestModule) Run(moduleType ModuleType) {
	switch moduleType {
	case MERGE_STRUCT:
		type worker struct {
			Name           string
			Age            int
			Id             int
			IdentityNumber string
			InterfaceTest  interface{}
		}

		workerA := &worker{
			Name: "Michael J.Viper",
		}

		workerB := &worker{
			Age:            36,
			Id:             34,
			IdentityNumber: "23232323",
			InterfaceTest:  addictionAlInterFace{
				id: 34,
			},
		}

		megerModule := helper.MergeModuleInitialize(workerA, workerB)
		mergeErr := megerModule.MergeTwoStruct(workerA, workerB,&helper.Config{})
		if mergeErr != nil {
			println(mergeErr)
		}
		//fmt.Println("FDSF")
		fmt.Println(workerA)
	}
}

func NewTest() TestModuleImpl {
	return &TestModule{}
}
