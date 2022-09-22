package helper

import (
	"encoding/json"
	"errors"
	"go/constant"
)

type InterFace interface {
	Decode(dataJSONString string) error
	DecodeDefault() error
	Encode() error
	GetData() any
	Setter(field string, value any, valType constant.Kind)
	GetDataInJSON() map[string]interface{}
}

type GenericStructUtilities[T map[string]interface{}, V any] struct {
	DataInJSON     T
	Data           V
	DataJsonString string
	AcceptKind     []constant.Kind
	InterFace
}

func InitializeGenericUtilities[T map[string]interface{}, V any](obj V) (InterFace, error) {
	GenericStructUtilities := &GenericStructUtilities[T, V]{}
	GenericStructUtilities.Data = obj
	if err := GenericStructUtilities.Encode(); err != nil {
		return nil, err
	}
	if err := GenericStructUtilities.Decode(GenericStructUtilities.DataJsonString); err != nil {
		return nil, err
	}
	return GenericStructUtilities, nil
}

//func CastingAcceptKindDetect(g *GenericStructUtilities) bool {
//
//}

func (g *GenericStructUtilities) Encode() error {
	b, err := json.Marshal(g.Data)
	if err != nil {
		return errors.New(err.Error())
	}
	g.DataJsonString = string(b)
	return nil
}

func (g *GenericStructUtilities) Decode(dataJSONString string) error {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(dataJSONString), &m); err != nil {
		return errors.New(err.Error())
	}
	g.DataInJSON = m
	return nil
}

func (g *GenericStructUtilities) DecodeDefault() error {
	var m any
	if err := json.Unmarshal([]byte(g.DataJsonString), &m); err != nil {
		return errors.New(err.Error())
	}
	g.Data = m
	return nil
}

func (g *GenericStructUtilities) GetData() any {
	return g.Data
}

func (g *GenericStructUtilities) GetDataInJSON() map[string]interface{} {
	return g.DataInJSON
}

//func (g *GenericStructUtilities) GetData() any {
//	return g.Data
//}

func (g *GenericStructUtilities) Setter(field string, value any, valType constant.Kind) {
	if valType == constant.Int {
		g.DataInJSON[field] = value.(int)
	} else if valType == constant.String {
		g.DataInJSON[field] = value.(string)
	}
}
