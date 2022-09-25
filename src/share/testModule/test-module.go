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

type PagingDto struct {
	Page   int
	Limit  int
	Offset int
}

type PagingDto2 struct {
	PagingDto
}

type CustomStructA struct {
	PagingDto
	Age         int
	Id          int
	Description string
}

func (t *TestModule) Run(moduleType ModuleType) {
	switch moduleType {
	case MERGE_STRUCT:
		type worker struct {
			PagingDto
			Name           string
			Age            int
			Id             int
			IdentityNumber string
			InterfaceTest  interface{}
		}

		pgDto := &PagingDto2{}

		pgDto.Page = 4
		pgDto.Limit = 40
		pgDto.Offset = 239

		workerA := &worker{
			Name: "Michael J.Viper",
		}

		workerB := &worker{
			Age:            36,
			Id:             34,
			IdentityNumber: "23232323",
			InterfaceTest: addictionAlInterFace{
				id: 34,
			},
		}
		//workerB.limit = 10
		//workerB.offset= 9
		fmt.Println(workerB)
		megerModule := helper.MergeModuleInitialize(workerA, pgDto)
		mergeErr := megerModule.MergeTwoStruct(workerA, pgDto, &helper.Config{})
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
