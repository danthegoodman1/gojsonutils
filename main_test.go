package gojsonutils

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	flattenedMap = `{"arr":[1,2],"hey":"ho","lets":1,"nestedar":[[1],[2]],"obj__this":"key","obj__val":[1,2],"objar__a":["val","val2"],"objar__b":[2,null],"objar__c":[[1],null],"objar__d__e":[[1,2],null],"objar__darr":["[{\"e\":1},{\"e\":2},[44]]",null],"objar__dstr":["[{\"e\":1},{\"e\":2},44]",null],"objar__f__g":[1,null],"objar__f__h":[[2],null],"objar__new":[null,"nvale"]}`
)

var (
	mapToFlat = map[string]any{
		"hey":  "ho",
		"lets": 1,
		"arr":  []any{1, 2},
		"obj": map[string]any{
			"this": "key",
			"val":  []any{1, 2},
		},
		"objar": []any{
			map[string]any{
				"a":    "val",
				"b":    2,
				"c":    []any{1},
				"d":    []any{map[string]any{"e": 1}, map[string]any{"e": 2}},
				"dstr": []any{map[string]any{"e": 1}, map[string]any{"e": 2}, 44},
				"darr": []any{map[string]any{"e": 1}, map[string]any{"e": 2}, []any{44}},
				"f": map[string]any{
					"g": 1,
					"h": []any{2},
				},
			},
			map[string]any{
				"a":   "val2",
				"new": "nvale",
			},
		},
		"nestedar": []any{[]any{1}, []any{2}},
	}
)

func TestFlatten(t *testing.T) {
	flattened, err := Flatten(mapToFlat, nil)
	if err != nil {
		t.Fatal(err)
	}
	b, err := json.Marshal(flattened)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
	if string(b) != flattenedMap {
		t.Fatal("did not match flattenedMap!")
	}
}
