package helper

import (
	errors2 "errors"
	"reflect"
)

type MergeModuleImplm interface {
	Merge(dst, src interface{}) error
	deepMerge(dst, src reflect.Value, deepLevel int) error
	getSrcStructName() string
	getDstStructName() string
	ifStructHasMergeableFields(field reflect.Value) (canExport bool)
}

type MergeModule struct {
	dst interface{}
	src interface{}
	ref reflect.Kind
}

/*
	main function
	Convert struct,interface to reflect.Value first
	dst: Desination <must pass as pointer for Editable>
	src: Source <Source to implement coppy and set to dst>
*/
func (m *MergeModule) Merge(dst, src interface{}) error {
	var (
		dstConvert reflect.Value
		srcConvert reflect.Value
	)
	//force dst must be pointer
	dstValue := reflect.ValueOf(dst)
	if kindOfDst := dstValue.Kind(); kindOfDst != reflect.Ptr {
		return errors2.New("Destination must be pointer")
	}
	dstConvert = dstValue.Elem()
	srcValue := reflect.ValueOf(src)
	if kindOfSrcValue := srcValue.Kind(); kindOfSrcValue == reflect.Ptr {
		srcConvert = srcValue.Elem()
	}
	if checkErr := m.deepMerge(dstConvert, srcConvert, 0); checkErr != nil {
		return checkErr
	}
	return nil
}

/*
	Version1 merge struct
*/
func (m *MergeModule) deepMerge(dst, src reflect.Value, deepLevel int) (err error) {
	switch dst.Kind() {
	case reflect.Struct:
		if m.ifStructHasMergeableFields(dst) {
			for i, n := 0, dst.NumField(); i < n; i++ {
				if err = m.deepMerge(dst.Field(i), src, deepLevel+1); err != nil {
					return
				}
			}
		} else if dst.CanSet() && !isEmpty(src) {

		}
	default:
	}

	//check MergeAbleField
	return nil

}

/*
	Check if struct has exported field to merge
*/
func (m *MergeModule) ifStructHasMergeableFields(dst reflect.Value) (canExport bool) {
	for i, n := 0, dst.NumField(); i < n; i++ {
		exportedComponent := true
		field := dst.Type().Field(i)
		pkgPath := field.PkgPath
		if len(pkgPath) > 0 {
			exportedComponent = false
		}
		c := field.Name[0]
		if 'a' <= c && c <= 'z' || c == '_' {
			exportedComponent = false
		}
		if field.Anonymous && dst.Field(i).Kind() == reflect.Struct {
			canExport = canExport || m.ifStructHasMergeableFields(dst.Field(i))
		} else if exportedComponent {
			canExport = canExport || len(field.PkgPath) == 0
		}
	}
	return
}

/*
	check if value is not empty
*/
func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint32, reflect.Uint64, reflect.Uint16, reflect.Uint8:
		return v.Uint() == 0
	case reflect.Interface, reflect.Ptr:
		if v.IsNil() {
			return true
		}
	case reflect.Func:
		return v.IsNil()
	}
	return false

}
func getStructName(strct interface{}) (structName string) {
	if t := reflect.TypeOf(strct); t.Kind() != reflect.Ptr {
		structName = t.Name()
		return
	} else {
		structName = t.Elem().Name()
		return
	}
}

func (m *MergeModule) getSrcStructName() string {
	return getStructName(m.src)

}

func (m *MergeModule) getDstStructName() string {
	return getStructName(m.dst)
}

func MergeModuleInitialize(src, dst interface{}) MergeModuleImplm {
	return &MergeModule{
		dst: dst,
		src: src,
	}
}
