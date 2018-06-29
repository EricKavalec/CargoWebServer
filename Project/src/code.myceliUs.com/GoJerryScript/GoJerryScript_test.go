package GoJerryScript

import "testing"

func TestHelloJerry(t *testing.T) {
	// Init the script engine.
	Jerry_init(Jerry_init_flag_t(JERRY_INIT_EMPTY))

	/* Register 'print' function  */
	RegisterPrintHandler()

	//str := "function add(a, b){return a+b;}; add(1, 2);"
	str := "print ('Hello, World!');"
	var arg0 Uint8                     // nil pointer
	var arg1 = int64(0)                // 0 length
	var arg2 = NewUint8FromString(str) // The script.
	var arg3 = int64(len(str) + 1)
	var arg4 = NewUint32FromInt(int32(JERRY_PARSE_NO_OPTS))

	parsed_code := Jerry_parse(arg0, arg1, arg2, arg3, arg4)
	if !Jerry_value_is_error(parsed_code) {

		/* Execute the parsed source code in the Global scope */
		jerry_value_t := Jerry_run(parsed_code)

		/* Returned value must be freed */
		Jerry_release_value(jerry_value_t)
	}

	/* Parsed source code must be freed */
	Jerry_release_value(parsed_code)

	// Cleanup the script engine.
	Jerry_cleanup()
}

func TestEvalScript(t *testing.T) {
	engine := NewEngine(9696, JERRY_INIT_EMPTY)

	// Test eval string function.
	engine.AppendFunction("SayHelloTo", []string{"to"}, "function SayHelloTo(to){return 'Hello ' + to + '!';}", JERRY_PARSE_NO_OPTS)
	str, _ := engine.EvalScript("SayHelloTo(jerry);", []Property{{Name: "jerry", Value: "Jerry Script"}})

	if str != "Hello Jerry Script!" {
		t.Error("Expected 'Hello Jerry Script!', got ", str)
	}

	// Test numeric function
	engine.AppendFunction("Add", []string{"a", "b"}, "function Add(a, b){return a + b;}", JERRY_PARSE_NO_OPTS)

	number, _ := engine.EvalScript("Add(a, b);", []Property{{Name: "a", Value: 1}, {Name: "b", Value: 2.25}})
	if number != 3.25 {
		t.Error("Expected 3.25, got ", number)
	}

	// Test boolean function
	engine.AppendFunction("TestBool", []string{"val"}, "function TestBool(val){val>0;}", JERRY_PARSE_NO_OPTS)

	boolean, _ := engine.EvalScript("TestBool(val)", []Property{{Name: "val", Value: 1}})
	if boolean == false {
		t.Error("Expected true, got ", boolean)
	}

	// Test with array
	engine.AppendFunction("TestArray", []string{"arr, val"}, "function TestArray(arr, val){arr.push(val); return arr;}", JERRY_PARSE_NO_OPTS)

	arr, err := engine.EvalScript("TestArray(arr, val);", []Property{{Name: "arr", Value: []interface{}{1.0, 3.0, 4.0}}, {Name: "val", Value: 2.25}})
	if err == nil {
		t.Log(arr)
	}
	engine.Clear()
}