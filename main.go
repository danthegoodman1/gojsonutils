package gojsonutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func getType(item any) string {
	if _, isObject := item.(map[string]any); isObject {
		return "object"
	} else if _, isArr := item.([]any); isArr {
		return "array"
	}
	return reflect.TypeOf(item).String()
}

// Flatten takes in a JSON object, and returns the flattened variant.
// Optionally can specify a separator, the default is `__` (double underscore).
func Flatten(object any, separator *string) (any, error) {
	sep := "__"
	if separator != nil {
		sep = *separator
	}

	flat := map[string]any{}
	arr := []any{}

	hasObj := false
	hasArr := false

	if objectArray, isArray := object.([]any); isArray {
		// Is an array
		for i, item := range objectArray {
			if itemObject, isObject := item.(map[string]any); isObject {
				// We have an object
				if hasArr {
					str, err := json.Marshal(item)
					if err != nil {
						return nil, fmt.Errorf("error in json.Marshal array,object,hasArr: %w", err)
					}
					return string(str), nil
				}
				hasObj = true
				flattened, err := Flatten(itemObject, &sep)
				if err != nil {
					return nil, fmt.Errorf("error in Flatten array,object: %w", err)
				}
				flattenedObj, isObj := flattened.(map[string]any)
				if !isObj {
					return nil, errors.New("array,object flattened object is not object")
				}
				for key, val := range flattenedObj {
					if _, exists := flat[key]; !exists {
						flat[key] = []any{}
						if i != 0 {
							// back-fill nulls
							bfn := []any{}
							for ind := 0; ind < i; ind++ {
								bfn = append(bfn, nil)
							}
							flat[key] = bfn
						}
					}
					a := flat[key].([]any)
					flat[key] = append(a, val)
				}
				// Forward-fill nulls
				for key := range flat {
					if _, exists := flattenedObj[key]; !exists {
						a := flat[key].([]any)
						flat[key] = append(a, nil)
					}
				}
			} else if _, isArr := item.([]any); isArr {
				// We have an array
				if hasObj {
					str, err := json.Marshal(object)
					if err != nil {
						return nil, fmt.Errorf("error in json.Marshal array,array,hasObj: %w", err)
					}
					return string(str), nil
				}
				flattened, err := Flatten(item, &sep)
				if err != nil {
					return nil, fmt.Errorf("error in Flatten array,array: %w", err)
				}
				// array of arrays support
				hasArr = true
				arr = append(arr, flattened)
			} else if item == nil && hasArr {
				// if we have an array of arrays, but one item is null instead of an array
				arr = append(arr, item)
			} else {
				// Just a value
				if hasObj || hasArr {
					str, err := json.Marshal(object)
					if err != nil {
						return nil, fmt.Errorf("error in json.Marshal array,else: %w", err)
					}
					return string(str), nil
				}
				// verify that we don't have mixed results
				var nonNulls []any
				for _, val := range objectArray {
					if val != nil {
						nonNulls = append(nonNulls, val)
					}
				}
				if len(nonNulls) <= 1 {
					// Either array of nulls, empty, or we have one non-null item
					return object, nil
				}
				firstNonNullType := getType(nonNulls[0])
				//const mixedResults = nonNulls.slice(1).filter((objItem) => firstNonNullType !== getType(objItem) && firstNonNullType).length > 0
				//if (mixedResults) {
				//	return JSON.stringify(obj)
				//}
				//return obj
				mixedResults := false
				for _, item := range nonNulls[1:] {
					if getType(item) != firstNonNullType {
						mixedResults = true
						break
					}
				}
				if mixedResults {
					str, err := json.Marshal(item)
					if err != nil {
						return nil, fmt.Errorf("error in json.Marshal array,else,mixedresults: %w", err)
					}
					return string(str), nil
				}
				fmt.Println("returning object", object)
				return object, nil
			}
		}
	} else if object, isObject := object.(map[string]any); isObject {
		// Is an object
		hasObj = true
		for key, val := range object {
			if valObj, isObject := val.(map[string]any); isObject {
				flattened, err := Flatten(valObj, &sep)
				if err != nil {
					return nil, fmt.Errorf("error in object,object: %w", err)
				}
				for nkey, nval := range flattened.(map[string]any) {
					flat[key+sep+nkey] = nval
				}
			} else if valArr, isArr := val.([]any); isArr {
				flattened, err := Flatten(valArr, &sep)
				if err != nil {
					return nil, fmt.Errorf("error in object,array: %w", err)
				}
				if flattenedObj, isObj := flattened.(map[string]any); isObj {
					// merge
					for nkey, nval := range flattenedObj {
						flat[key+sep+nkey] = nval
					}
				} else if _, isArr := flattened.([]any); isArr {
					// set it
					flat[key] = flattened
				} else {
					// stringified or mixed results
					flat[key] = flattened
				}
			} else {
				flat[key] = val
			}
		}
	}

	if hasObj {
		return flat, nil
	} else {
		return arr, nil
	}
}
