package utils

import (
	"fmt"
	"log"
	"reflect"
)

const (
	aSlice = "slice"
	aSliceOfSlice = "sliceSlice"
	aMap = "map"
)

type generation struct {
	// outputType is here to track what we detect the output should be
	// it should be one of
	// slice of map (struct)
	// slice of slice of map (struct)
	// or just a map (struct)
	outputType string
	
	// sliceOfSliceOutput will be populated if the outputType is an slice of slice
	sliceOfSliceOutput [][]map[string]interface{}

	// sliceOutput will be populated if the outputType is an slice
	sliceOutput []map[string]interface{}

	// mapOutput will be populated if the outputType is a map
	mapOutput map[string]interface{}
}

func GenerateDesiredOutput(desiredOutput interface{}) ([]byte, error) {
	ifaceType := reflect.TypeOf(desiredOutput)
	ifaceVals := reflect.ValueOf(desiredOutput)
	log.Printf("desired value: %#v\n", ifaceVals)
	var gen generation 
	if ifaceType.Kind() == reflect.Slice{
		log.Println("is a slice")
		gen.outputType = aSlice
		gen.sliceOutput = make([]map[string]interface{}, 0)
		gen.generateDesiredSlice(ifaceVals.Type().Elem())
	} else if ifaceType.Kind() == reflect.Struct {
		gen.outputType = aMap
		gen.mapOutput = make(map[string]interface{})
		gen.generateDesiredStruct(ifaceVals.Type())
	} else {
		gen.generateDesiredValue(ifaceVals.Type())
	}
	return gen.bytes()
}

func (g generation) generateDesiredSlice(ifaceType reflect.Type) error {
	// By now we know that its a slice -- []interface{}
	// and might be []someStruct{}
	// or it might be [][]someStruct{}
		if ifaceType.Kind() == reflect.Slice{
			log.Println("is a slice of slice")
			g.outputType = aSliceOfSlice
			g.sliceOfSliceOutput = make([][]map[string]interface{}, 0)
			g.generateDesiredSliceOfSlice(ifaceType)
		} else if ifaceType.Kind() == reflect.Struct {
			log.Println("is a struct")
			g.generateDesiredStruct(ifaceType)
		} else {
			err := g.generateDesiredValue(ifaceType)
			if err != nil {
				return err
			}
		}
	return nil
}

func (g generation) generateDesiredSliceOfSlice(ifaceType reflect.Type) error {
	return nil
}

func (g generation) generateDesiredStruct(ifaceType reflect.Type) error {
	var thisStruct = make(map[string]interface{})
	for i := 0; i < ifaceType.NumField(); i++ {
    key := ifaceType.Field(i).Tag.Get("json")
		thisStruct[key] = i
	}
	log.Printf("this struct: %#v", thisStruct)
	if g.outputType == aSlice {
		g.sliceOutput = append(g.sliceOutput, thisStruct)
	} else if g.outputType == aSliceOfSlice {
		g.sliceOfSliceOutput = append(g.sliceOfSliceOutput, []map[string]interface{}{thisStruct})
	} else {
		g.mapOutput = thisStruct
	}
	return nil
}

func (g generation) generateDesiredValue(ifaceVals reflect.Type) error {
	ifaceType := reflect.TypeOf(ifaceVals)
	switch ifaceType.Kind() {
	case reflect.String:
		// If it's a string, depending on what it's field name is, we generate a value
	case reflect.Int64:
		// If it's an int64 we generate a random int64
	case reflect.Int:
		// If it's an int we generate a random int
	case reflect.Float32:
		// if it's a float32 we generate a random float32
	case reflect.Bool:
		// if it's a bool "flip a coin"
	default:
		return fmt.Errorf("unknown type: %s", ifaceType.Kind().String())
	}
	return nil
}

func (g generation) bytes() ([]byte, error) {
	return nil, nil
}

// ContainsNamedValue determines if string name with given value is present in x
// fieldName string must be passed as it is used in x i.e. "FirstName"
// value must be same type expected in x i.e. string "Brandon"
// currently only supports a slice of exported structs i.e. []User
func ContainsNamedValue(x interface{}, fieldName string, value interface{}) bool {
	ifaceType := reflect.TypeOf(x)
	ifaceVals := reflect.ValueOf(x)
	if ifaceType.Kind() == reflect.Slice {
		if ifaceVals.Len() > 0 {
			if ifaceVals.Index(0).Kind() == reflect.Struct {
				for i := 0; i < ifaceVals.Len(); i++ {
					iface := ifaceVals.Index(i).FieldByName(fieldName)
					if iface.Interface() == value {
						return true
					}
				}
			}
		}
	}
	return false
}
