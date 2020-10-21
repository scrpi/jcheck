package jcheck

import (
	"testing"
)

func TestRules(t *testing.T) {
	testcases := []struct {
		js              string
		path            string
		checks          []CheckFunc
		expectedResultN int
	}{
		{`{"test": "string"}`, "test", []CheckFunc{StringEquals("string")}, 0},
		{`{"test": "string"}`, "test", []CheckFunc{StringHasPrefix("str")}, 0},
		{`{"test": "string"}`, "test", []CheckFunc{StringHasSuffix("ing")}, 0},
		{`{"test": "string"}`, "test", []CheckFunc{StringEquals("string")}, 0},
		{`{"test": {"nested": "string0"}}`, "test", []CheckFunc{IsObject()}, 0},
		{`{"test": {"nested": "string"}}`, "test.nested", []CheckFunc{StringEquals("string")}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{IsArray()}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenEquals(3)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenLT(4)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenLTE(4)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenLTE(3)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenGT(2)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenGTE(2)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenGTE(3)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array.1", []CheckFunc{NumEquals(20.4)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array.0", []CheckFunc{NumGT(10.1)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array.0", []CheckFunc{NumGTE(10.1)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array.0", []CheckFunc{NumGTE(10.2)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array.0", []CheckFunc{NumLT(10.3)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array.0", []CheckFunc{NumLTE(10.3)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array.0", []CheckFunc{NumLTE(10.2)}, 0},
		{`{"array": [10.2, 20.4, 30.6]}`, "array.0", []CheckFunc{NumGT(10), NumLT(11)}, 0},
		{`{"number": "750m"}`, "number", []CheckFunc{NumEquals(0.75)}, 0},
		{`{"number": "750m"}`, "number", []CheckFunc{NumGT(0.7)}, 0},
		{`{"number": "750m"}`, "number", []CheckFunc{NumGTE(0.7)}, 0},
		{`{"number": "750m"}`, "number", []CheckFunc{NumGTE(0.75)}, 0},
		{`{"number": "750m"}`, "number", []CheckFunc{NumLT(0.8)}, 0},
		{`{"number": "750m"}`, "number", []CheckFunc{NumLTE(0.8)}, 0},
		{`{"number": "750m"}`, "number", []CheckFunc{NumLTE(0.75)}, 0},

		{`{"test": "200m"}`, "test", []CheckFunc{IsString()}, 0},
		{`{"test": "200"}`, "test", []CheckFunc{IsString()}, 0},
		{`{"test": 200}`, "test", []CheckFunc{IsNumber()}, 0},

		{`{"test": "200u"}`, "test", []CheckFunc{NumEquals(0.0002)}, 0},
		{`{"test": "200m"}`, "test", []CheckFunc{NumEquals(0.2)}, 0},
		{`{"test": "200K"}`, "test", []CheckFunc{NumEquals(200000)}, 0},
		{`{"test": "200M"}`, "test", []CheckFunc{NumEquals(200000000)}, 0},

		{`{"test": true}`, "test", []CheckFunc{IsBoolean()}, 0},
		{`{"test": false}`, "test", []CheckFunc{IsBoolean()}, 0},
		{`{"test": true}`, "test", []CheckFunc{IsTrue()}, 0},
		{`{"test": false}`, "test", []CheckFunc{IsFalse()}, 0},
		{`{"test": null}`, "test", []CheckFunc{IsNull()}, 0},
		{`{"test": "true"}`, "test", []CheckFunc{IsString()}, 0},
		{`{"test": 0}`, "test", []CheckFunc{IsNumber()}, 0},

		{`{"test": "true"}`, "test", []CheckFunc{IsNumber()}, 1},
		{`{"test": "true"}`, "test", []CheckFunc{IsBoolean()}, 1},
		{`{"test": "string"}`, "test", []CheckFunc{StringHasSuffix("str")}, 1},
		{`{"test": "string"}`, "test", []CheckFunc{StringHasPrefix("ing")}, 1},
		{`{"test": null }`, "test", []CheckFunc{StringEquals("str")}, 1},
		{`{"test": null }`, "test", []CheckFunc{StringHasPrefix("str")}, 1},
		{`{"test": null }`, "test", []CheckFunc{StringHasSuffix("ing")}, 1},
		{`{"array": [10.2, 20.4, 30.6]}`, "array.1", []CheckFunc{NumEquals(10.2)}, 1},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenEquals(2)}, 1},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenLT(2)}, 1},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenLTE(2)}, 1},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenGT(4)}, 1},
		{`{"array": [10.2, 20.4, 30.6]}`, "array", []CheckFunc{ArrayLenGTE(4)}, 1},
		{`{"array": "notarray" }`, "array", []CheckFunc{ArrayLenEquals(2)}, 1},
		{`{"array": "notarray" }`, "array", []CheckFunc{ArrayLenLT(2)}, 1},
		{`{"array": "notarray" }`, "array", []CheckFunc{ArrayLenLTE(2)}, 1},
		{`{"array": "notarray" }`, "array", []CheckFunc{ArrayLenGT(2)}, 1},
		{`{"array": "notarray" }`, "array", []CheckFunc{ArrayLenGTE(2)}, 1},
		{`{"test": "string"}`, "test", []CheckFunc{StringEquals("notstring")}, 1},

		{`{"number": "200X"}`, "number", []CheckFunc{NumEquals(2)}, 1},
		{`{"number": "200X"}`, "number", []CheckFunc{NumGT(2)}, 1},
		{`{"number": "200X"}`, "number", []CheckFunc{NumGTE(2)}, 1},
		{`{"number": "200X"}`, "number", []CheckFunc{NumGTE(2)}, 1},
		{`{"number": "200X"}`, "number", []CheckFunc{NumLT(2)}, 1},
		{`{"number": "200X"}`, "number", []CheckFunc{NumLTE(2)}, 1},
		{`{"number": "200X"}`, "number", []CheckFunc{NumLTE(2)}, 1},
		{`{"array": [1] }`, "array", []CheckFunc{NumEquals(2)}, 1},
		{`{"array": [1] }`, "array", []CheckFunc{NumGT(2)}, 1},
		{`{"array": [1] }`, "array", []CheckFunc{NumGTE(2)}, 1},
		{`{"array": [1] }`, "array", []CheckFunc{NumGTE(2)}, 1},
		{`{"array": [1] }`, "array", []CheckFunc{NumLT(2)}, 1},
		{`{"array": [1] }`, "array", []CheckFunc{NumLTE(2)}, 1},
		{`{"array": [1] }`, "array", []CheckFunc{NumLTE(2)}, 1},
		{`{"number": "750M"}`, "number", []CheckFunc{NumEquals(0.75)}, 1},
		{`{"number": "75m"}`, "number", []CheckFunc{NumGT(0.7)}, 1},
		{`{"number": "75m"}`, "number", []CheckFunc{NumGTE(0.7)}, 1},
		{`{"number": "75m"}`, "number", []CheckFunc{NumGTE(0.75)}, 1},
		{`{"number": "750K"}`, "number", []CheckFunc{NumLT(0.8)}, 1},
		{`{"number": "750K"}`, "number", []CheckFunc{NumLTE(0.8)}, 1},
		{`{"number": "750K"}`, "number", []CheckFunc{NumLTE(0.75)}, 1},
	}

	for _, tc := range testcases {
		jc, err := NewJSONChecker(tc.js, DefaultPermitted())
		if err != nil {
			t.Error(err)
			continue
		}
		jc.AddRule(tc.path, tc.checks...)
		results, _ := jc.Check()
		if len(results) != tc.expectedResultN {
			for _, r := range results {
				t.Log(r)
			}
			t.Errorf("%s - %q - Failed", tc.js, tc.path)
		}
	}
}

// func TestPrintTree(t *testing.T) {
// 	obj := map[string]interface{}{}

// 	if err := json.Unmarshal([]byte(js), &obj); err != nil {
// 		t.Error(err)
// 	}

// 	nodes := scanNodeMap(nil, "", obj)

// 	for _, n := range nodes {
// 		n.forEachNode(func(n *node) { fmt.Println(n) })
// 	}
// }

func TestPermitted(t *testing.T) {
	js := `
{
	"object": {
		"permitted": null,
		"notpermitted": null 
	}
}
`
	jc, err := NewJSONChecker(js, DefaultNotPermitted())
	if err != nil {
		t.Error(err)
		return
	}

	jc.AddRule("object.permitted", Permitted())

	results, ok := jc.Check()
	if ok || len(results) != 1 {
		t.Error("object.notpermitted was not caught")
		for _, r := range results {
			t.Error(r)
		}
	}

	jc.AddRule("ob??ct.*", Permitted())
	results, ok = jc.Check()
	if !ok {
		for _, r := range results {
			t.Error(r)
		}
	}
}

func TestNotPermitted(t *testing.T) {
	js := `
{
	"object": {
		"permitted": null,
		"notpermitted": null 
	}
}
`
	jc, err := NewJSONChecker(js, DefaultPermitted())
	if err != nil {
		t.Error(err)
		return
	}

	// Everything should pass with no rules - default permitted
	results, ok := jc.Check()
	if !ok {
		for _, r := range results {
			t.Error(r)
		}
	}

	jc.AddRule("object.*", NotPermitted())
	results, ok = jc.Check()
	if ok {
		t.Error("NotPermitted rule did not catch anything")
	}

	if len(results) != 2 {
		t.Error("Unexpected result count. Expected 2.")
		for _, r := range results {
			t.Error(r)
		}
	}
}
