package helper

import (
	errors2 "errors"
	"fmt"
	"reflect"
)

var (
	mustPointer = "Destination must be pointer"
	isPointer   = "is pointer"
)

var (
	ErrNilArguments                = errors2.New("src and dst must not be nil")
	ErrDifferentArgumentsTypes     = errors2.New("src and dst must be of same type")
	ErrNotSupported                = errors2.New("only structs, maps, and slices are supported")
	ErrExpectedMapAsDestination    = errors2.New("dst was expected to be a map")
	ErrExpectedStructAsDestination = errors2.New("dst was expected to be a struct")
	ErrNonPointerAgument           = errors2.New("dst must be a pointer")
)

type MergeModuleImplm interface {
	MergeTwoStruct(dst, src interface{}, config *Config) error
	deepMerge(dst, src reflect.Value, deepLevel int, config *Config) error
	getSrcStructName() string
	getDstStructName() string
	ifStructHasMergeableFields(field reflect.Value) (canExport bool)
	Merge() (err error)
}

type MergeModule struct {
	dst interface{}
	src interface{}
	ref reflect.Kind
}

type Config struct {
	Override              bool
	AppendSlice           bool
	TypeCheck             bool
	OverwriteWithEmptySrc bool
}

/*
	main function
	Convert struct,interface to reflect.Value first
	dst: Desination <must pass as pointer for Editable>
	src: Source <Source to implement coppy and set to dst>
*/
func (m *MergeModule) MergeTwoStruct(dst, src interface{}, config *Config) error {
	//logService.LogWithMsg("Start Debugging",logger.INFO)
	fmt.Println("===========Start Debugging======================")
	var (
		dstConvert reflect.Value
		srcConvert reflect.Value
	)
	//force dst must be pointer
	dstValue := reflect.ValueOf(dst)
	if kindOfDst := dstValue.Kind(); kindOfDst != reflect.Ptr {
		fmt.Println(mustPointer)
		return errors2.New(mustPointer)
	} else if kindOfDst == reflect.Ptr {
		dstConvert = dstValue.Elem()
	}
	srcValue := reflect.ValueOf(src)
	if kindOfSrcValue := srcValue.Kind(); kindOfSrcValue == reflect.Ptr {
		fmt.Println("kindOfSrcValue Is Pointer")
		srcConvert = srcValue.Elem()
	}
	fmt.Printf("dst Value after convert: %v\nsrc Value after convert: %v\n", dstConvert, srcConvert)
	if checkErr := m.deepMerge(dstConvert, srcConvert, 0, config); checkErr != nil {
		fmt.Println(checkErr)
		return checkErr
	}
	return nil
}

func (m *MergeModule) Merge() (err error) {
	err = m.MergeTwoStruct(&m.dst, m.src, &Config{})
	if err != nil {
		return
	}
	return nil
}

/*
	Version1 merge struct
*/
func (m *MergeModule) deepMerge(dst, src reflect.Value, deepLevel int, config *Config) (err error) {
	fmt.Printf("DeepLevel: %v\n", deepLevel)

	Override := config.Override
	//AppendSlice := config.AppendSlice
	//TypeCheck := config.TypeCheck
	OverwriteWithEmptySrc := config.OverwriteWithEmptySrc

	fmt.Println("Entry DeepMerge ======================")
	switch dst.Kind() {
	case reflect.Struct:
		fmt.Printf("Kind of destination is %v :::Value: %v\n", dst.Kind(), dst)
		fmt.Printf("Kind of src is %v :::Value: %v\n", src.Kind(), src)
		fmt.Printf("Amount Of Dst Field is %v\n", dst.NumField())
		fmt.Printf("Amount Of Src Field is %v\n", src.NumField())

		if m.ifStructHasMergeableFields(dst) {
			for i, n := 0, dst.NumField(); i < n; i++ {
				if err = m.deepMerge(dst.Field(i), src.Field(i), deepLevel+1, config); err != nil {
					return
				}
			}
		} else if dst.CanSet() && (isReflectNil(src) || Override) && (!isEmpty(src) || OverwriteWithEmptySrc) {
			fmt.Printf("Set Source\n")
			fmt.Printf("[CASE][reflect.Struct]dst before set :%v\n", dst)
			dst.Set(src)
			fmt.Printf("[CASE][reflect.Struct]dst after set :%v\n", dst)

		}
	case reflect.Ptr:
		fallthrough
	case reflect.Interface:
		fmt.Println("dst is interface")
		if isReflectNil(src) {
			fmt.Println("src  is Nil")
			if dst.CanSet() && src.Type().AssignableTo(dst.Type()) {
				dst.Set(src)
			}
			break
		}
		if src.Kind() != reflect.Interface {
			if dst.IsNil() || (src.Kind() != reflect.Ptr && Override) {
				if dst.CanSet() && (Override || isEmpty(dst)) {
					dst.Set(src)
				}
			} else if src.Kind() == reflect.Ptr {
				if err = m.deepMerge(dst, src, deepLevel, config); err != nil {
					return
				}
			} else if dst.Elem().Type() == src.Type() {
				if err = m.deepMerge(dst, src, deepLevel, config); err != nil {
					return
				}
			} else {
				return ErrDifferentArgumentsTypes
			}
		}

		if dst.IsNil() || Override {
			fmt.Println("DST IS NIL")
			if dst.CanSet() && (Override || isEmpty(dst)) {
				dst.Set(src)
				break
			}
		}
		if dst.Elem().Kind() == src.Kind() {
			fmt.Println("DST AND SRC ARE THE SAME TYPE")
			if err = m.deepMerge(dst, src, deepLevel, config); err != nil {
				return
			}
		}
	//setValue for field have base type
	default:
		fmt.Printf("[CASE][default]dst before set :%v\n", dst)
		if mustSet := (isEmpty(dst) || Override) && (!isEmpty(src) || OverwriteWithEmptySrc); mustSet {
			fmt.Println("Must set")
			fmt.Printf("[CASE][default]dst before set :%v\n", dst)
			dst.Set(src)
			fmt.Printf("[CASE][default]dst after set :%v\n", dst)

		}
	}
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
	case reflect.Float64, reflect.Float32:
		return v.Float() == 0
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
func MergeModuleInitialize(dst, src interface{}) MergeModuleImplm {
	return &MergeModule{
		dst: dst,
		src: src,
	}
}

func isReflectNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func, reflect.Map, reflect.Slice:
		if v.IsNil() {
			fmt.Printf("%v is Nil\n", v.Kind())

		}
		return v.IsNil()
	default:
		return false
	}
}
