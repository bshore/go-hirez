package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// TODO: Slice of Slice isn't working

const (
	aValue = "value"
	aMap = "map"
	aSlice = "slice"
	aSliceMap = "sliceMap"
	aSliceOfSliceMap = "sliceSliceMap"
)

type generation struct {
	// outputType is here to track what we detect the output should be
	// it should be one of:
	// a map
	// a slice
	// slice of map
	// slice of slice of map
	outputType string

	// valueOutput will be popualte if the outputType is a generic type
	valueOutput interface{}
	
	// mapOutput will be populated if the outputType is a map
	mapOutput map[string]interface{}

	// sliceOutput will be populated if the outputType is a slice
	sliceOutput []interface{}
	
	// sliceMapOutput will be populated if the outputType is a slice of map
	sliceMapOutput []map[string]interface{}
	
	// sliceOfSliceMapOutput will be populated if the outputType is an slice of slice of map
	sliceOfSliceMapOutput [][]map[string]interface{}
}

func GenerateDesiredOutput(desiredOutput interface{}) ([]byte, error) {
	ifaceType := reflect.TypeOf(desiredOutput)
	ifaceVals := reflect.ValueOf(desiredOutput)
	var gen generation

	// If it's a slice, we need to check what it's a slice of
	if ifaceType.Kind() == reflect.Slice{
		// get the element of the slice
		ifaceElement := ifaceVals.Type().Elem()
		switch ifaceElement.Kind() {
		case reflect.Slice:
			// handle if it's a slice of slice
			gen.outputType = aSliceOfSliceMap
			gen.sliceOfSliceMapOutput = make([][]map[string]interface{}, 0)
			theSlice, err := gen.generateDesiredSlice(ifaceElement)
			if err != nil {
				return nil, err
			}
			gen.sliceOfSliceMapOutput = append(gen.sliceOfSliceMapOutput, theSlice)
		case reflect.Struct:
			// handle if it's a slice of struct
			gen.outputType = aSliceMap
			gen.sliceMapOutput = make([]map[string]interface{}, 0)
			theStruct, err := gen.generateDesiredStruct(ifaceElement)
			if err != nil {
				return nil, err
			}
			gen.sliceMapOutput = append(gen.sliceMapOutput, theStruct)
		default:
			// handle of it's a slice of a generic type
			gen.outputType = aSlice
			gen.sliceOutput = make([]interface{}, 0)
			theValue, err := gen.generateDesiredValue(ifaceElement)
			if err != nil {
				return nil, err
			}
			gen.sliceOutput = append(gen.sliceOutput, theValue)
		}
	}	else if ifaceType.Kind() == reflect.Struct {
		// handle if it's just a struct
		gen.outputType = aMap
		theStruct, err := gen.generateDesiredStruct(ifaceVals.Type())
		if err != nil {
			return nil, err
		}
		gen.mapOutput = theStruct
	} else {
		// handle if it's just a single value
		gen.outputType = aValue
		value, err := gen.generateDesiredValue(ifaceVals.Type())
		if err != nil {
			return nil, err
		}
		gen.valueOutput = value
	}
	return gen.bytes()
}

func (g generation) generateDesiredSlice(ifaceType reflect.Type) ([]map[string]interface{}, error) {
	var thisSlice = make([]map[string]interface{}, 0)
	ifaceElement := ifaceType.Elem()
	switch ifaceElement.Kind() {
	case reflect.Slice:
		return nil, fmt.Errorf("was not expecting a slice of slice of slice")
	case reflect.Struct:
		theStruct, err := g.generateDesiredStruct(ifaceElement)
		if err != nil {
			return nil, err
		}
		thisSlice = append(thisSlice, theStruct)
	default:
		return nil, fmt.Errorf("was not expecting a slice of slice of non-struct")
	}
	return thisSlice, nil
}

func (g generation) generateDesiredStruct(ifaceType reflect.Type) (map[string]interface{}, error) {
	var thisStruct = make(map[string]interface{})
	for i := 0; i < ifaceType.NumField(); i++ {
    key := ifaceType.Field(i).Tag.Get("json")
		valType := ifaceType.Field(i).Type
		value, err := g.generateDesiredValue(valType)
		if err != nil {
			return nil, err
		}
		thisStruct[key] = value
	}
	return thisStruct, nil
}

func (g generation) generateDesiredValue(ifaceType reflect.Type) (interface{}, error) {
	switch ifaceType.Kind() {
	case reflect.Slice:
		return g.generateDesiredSlice(ifaceType)
	case reflect.Struct:
		return g.generateDesiredStruct(ifaceType)
	case reflect.String:
		// If it's a string, depending on what it's field name is, we generate a value
		return "random string", nil
	case reflect.Int64:
		// If it's an int64 we generate a random int64
		return int64(1), nil
	case reflect.Int:
		// If it's an int we generate a random int
		return 2, nil
	case reflect.Float32:
		// if it's a float32 we generate a random float32
		return 3.4, nil
	case reflect.Bool:
		// if it's a bool "flip a coin"
		return true,nil
	default:
		return nil, fmt.Errorf("unknown type: %s", ifaceType.Kind().String())
	}
}

func (g generation) bytes() ([]byte, error) {
	switch g.outputType {
	case aMap:
		return json.Marshal(g.mapOutput)
	case aValue:
		return json.Marshal(g.valueOutput)
	case aSlice:
		return json.Marshal(g.sliceOutput)
	case aSliceMap:
		bs, err := json.MarshalIndent(g.sliceMapOutput, "", "\t")
		if err != nil {
			return nil, err
		}
		log.Println(string(bs))
		return bs, nil
	case aSliceOfSliceMap:
		return json.Marshal(g.sliceOfSliceMapOutput)
	default:
		return nil, fmt.Errorf("could not determine output type: %v", g.outputType)
	}
}
