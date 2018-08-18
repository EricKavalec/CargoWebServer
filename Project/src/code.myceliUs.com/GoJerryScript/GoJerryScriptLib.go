package GoJerryScript

//#cgo LDFLAGS: -L/usr/local/lib -ljerry-core -ljerry-ext -ljerry-libm -ljerry-port-default
/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "jerryscript.h"
#include "jerryscript-ext/handler.h"
#include "jerryscript-debugger.h"

typedef jerry_value_t* jerry_value_p;
extern jerry_value_t call_function ( const jerry_value_t, const jerry_value_t, const jerry_value_p, jerry_size_t);
extern void setGoMethod(const char* name, jerry_value_t obj);
extern const char* get_object_reference_uuid(uintptr_t ref);
extern void delete_object_reference(uintptr_t ref);
extern jerry_value_t create_string (const char *str_p);
extern jerry_value_t eval (const char *source_p, size_t source_size, bool is_strict);
extern jerry_value_t parse_function (const char *resource_name_p, size_t resource_name_length,
                      const char *arg_list_p, size_t arg_list_size,
                      const char *source_p, size_t source_size, uint32_t parse_opts);
extern jerry_value_t create_error (jerry_error_t error_type, const char *message_p);
extern jerry_size_t string_to_utf8_char_buffer (const jerry_value_t value, char *buffer_p, size_t size);
extern jerry_size_t get_string_size (const jerry_value_t value);
extern jerry_value_t create_native_object(const char* uuid);
extern jerry_value_t create_array (uint32_t);
extern jerry_value_t set_property_by_index (const jerry_value_t, uint32_t, const jerry_value_t);
extern jerry_value_t get_property_by_index (const jerry_value_t, uint32_t);
extern uint32_t get_array_length (const jerry_value_t);

*/
import "C"

//import "reflect"
import "unsafe"
import "encoding/binary"
import "math"
import "code.myceliUs.com/Utility"
import "errors"
import "reflect"
import "fmt"
import "encoding/json"
import "strconv"

//import "strconv"
import "log"

// Global variable.
var (
	// Callback function used by dynamic type, it's call when an entity is set.
	// Can be use to store dynamic type in a cache.
	SetEntity func(interface{}) = func(val interface{}) {
		log.Println("---> set entity ", val)
	}

	// Channel to be use to transfert information from client and server
	Call_remote_actions_chan chan *Action
)

// Function to access remote action channels.

func evalScript(script string) (Value, error) {

	// Now I will evaluate the function...
	src := C.CString(script)
	r := C.eval(src, C.size_t(len(script)), false)

	// Free the allocated value.
	C.free(unsafe.Pointer(src))

	// Create a Uint_32 value from the result.
	ret := jerry_value_t_To_uint32_t(r)

	value := NewValue(ret)

	if Jerry_value_is_error(ret) {
		return *value, errors.New("Fail to run script " + script)
	}

	return *value, nil
}

/**
 *
 */
func appendJsFunction(name string, args []string, src string) error {

	// Parameters
	resource_name_p := C.CString(name)
	resource_name_length := C.size_t(len(name))
	source_p := C.CString(src)
	source_size := C.size_t(len(src))

	// A list of string value separated by ','
	args_ := ""
	for i := 0; i < len(args); i++ {
		args_ += args[i]
		if i < len(args)-1 {
			args_ += ", "
		}
	}

	// The list of arguments.
	arg_list_p := C.CString(args_)
	arg_list_size := C.size_t(len(args_))

	r := C.parse_function(resource_name_p, resource_name_length, arg_list_p, arg_list_size, source_p, source_size, C.uint32_t(JERRY_PARSE_NO_OPTS))

	parsed_code := jerry_value_t_To_uint32_t(r)

	// free memory used.
	C.free(unsafe.Pointer(resource_name_p))
	C.free(unsafe.Pointer(source_p))
	C.free(unsafe.Pointer(arg_list_p))

	// Keep the function object in a value.
	if !Jerry_value_is_error(parsed_code) {
		Jerry_release_value(parsed_code)
		// run the script once.
		evalScript(src)
		return nil
	} else {
		Jerry_release_value(parsed_code)
		log.Println("Fail to parse function " + name)
		return errors.New("Fail to parse function " + name)
	}

}

/**
 * Set a Go function as a method on a given object.
 */
func setGoMethod(object Uint32_t, name string, fct interface{}) {
	if fct != nil {
		Utility.RegisterFunction(name, fct)
	}
	cs := C.CString(name)
	if Jerry_value_is_object(object) {
		C.setGoMethod(cs, uint32_t_To_Jerry_value_t(object))
	}
	defer C.free(unsafe.Pointer(cs))
}

/**
 * Call a Js function / method
 */
func callJsFunction(obj Uint32_t, name string, params []interface{}) (Value, error) {
	var thisPtr C.jerry_value_t
	var fctPtr C.jerry_value_t
	var fct Uint32_t

	thisPtr = uint32_t_To_Jerry_value_t(obj)
	fctName := goToJs(name)
	defer goToJs(fctName)
	fct = Jerry_get_property(obj, fctName)
	defer Jerry_release_value(fct)

	fctPtr = uint32_t_To_Jerry_value_t(fct)
	var r Uint32_t
	var err error

	// if the function is define...
	if Jerry_value_is_function(fct) {
		// Now I will set the arguments...
		args := make([]C.jerry_value_t, len(params))
		for i := 0; i < len(params); i++ {
			if params[i] == nil {
				null := Jerry_create_null()
				defer Jerry_release_value(null)
				args[i] = uint32_t_To_Jerry_value_t(null)
			} else {
				p := goToJs(params[i])
				defer Jerry_release_value(p)
				args[i] = uint32_t_To_Jerry_value_t(p)
			}
		}

		var r_ C.jerry_value_t
		if len(args) > 0 {
			r_ = C.call_function(fctPtr, thisPtr, (C.jerry_value_p)(unsafe.Pointer(&args[0])), C.jerry_value_t(len(params)))
		} else {
			var args_ C.jerry_value_p
			r_ = C.call_function(fctPtr, thisPtr, args_, C.jerry_value_t(0))
		}

		r = jerry_value_t_To_uint32_t(r_)
	} else {
		err = errors.New("Function " + name + " dosent exist")
	}

	if Jerry_value_is_error(r) {
		err = errors.New("Fail to call function " + name)
	}

	result := NewValue(r)

	return *result, err
}

// Go function reside in the client, a remote call is made here.
func callGoFunction(name string, params ...interface{}) (interface{}, error) {
	action := new(Action)
	action.Name = name
	action.UUID = Utility.RandomUUID()

	// Set the list of parameters.
	for i := 0; i < len(params); i++ {
		action.AppendParam("arg"+strconv.Itoa(i), params[i])
	}

	// Send the action to the client side.
	Call_remote_actions_chan <- action

	// Set back the action with it results in it.
	action = <-Call_remote_actions_chan

	var err error
	if action.Results[1] != nil {
		err = action.Results[1].(error)
	}

	return action.Results[0], err
}

//export object_native_free_callback
func object_native_free_callback(native_p C.uintptr_t) {
	uuid := C.GoString(C.get_object_reference_uuid(native_p))
	C.delete_object_reference(native_p)
	GetCache().removeObject(uuid)
}

// The handler is call directly from Jerry script and is use to connect JS and GO
//export handler
func handler(fct C.jerry_value_t, this C.jerry_value_t, args C.uintptr_t, length int) C.jerry_value_t {
	// The function pointer.
	fctPtr := (Uint32_t)(SwigcptrUint32_t(C.uintptr_t((uintptr)(unsafe.Pointer(&fct)))))
	if Jerry_value_is_function(fctPtr) {
		propName := goToJs("name")
		defer Jerry_release_value(propName)
		proValue := Jerry_get_property(fctPtr, propName)
		defer Jerry_release_value(proValue)
		name, err := jsToGo(proValue)
		if err == nil {
			params := make([]interface{}, 0)
			for i := 0; i < length; i++ {
				val, err := jsToGo((Uint32_t)(SwigcptrUint32_t(C.uintptr_t(args))))
				if err == nil {
					params = append(params, val)
				} else {
					log.Panicln(err)
					jsError := createError(JERRY_ERROR_COMMON, err.Error())
					return uint32_t_To_Jerry_value_t(jsError)
				}
				args += 4
			}

			// This is the owner of the function.
			thisPtr := (Uint32_t)(SwigcptrUint32_t(C.uintptr_t((uintptr)(unsafe.Pointer(&this)))))
			if Jerry_value_is_object(thisPtr) {
				propUuid_ := Jerry_get_property(thisPtr, goToJs("uuid_"))
				defer Jerry_release_value(propUuid_)
				uuid, err := jsToGo(propUuid_)
				if err == nil {
					object := GetCache().getObject(uuid.(string))

					// Call object method.
					result, err := Utility.CallMethod(object, name.(string), params)
					if err == nil {
						if result == nil {
							return uint32_t_To_Jerry_value_t(Jerry_create_null())
						}
						jsVal := goToJs(result)
						return uint32_t_To_Jerry_value_t(jsVal)
					} else {
						jsError := createError(JERRY_ERROR_COMMON, err.(error).Error())
						return uint32_t_To_Jerry_value_t(jsError)
					}
				}

			} else {
				// There is no function owner I will simply call go function.
				result, err := callGoFunction(name.(string), params...)
				if err == nil && result != nil {
					jsVal := goToJs(result)
					val := (*uintptr)(unsafe.Pointer(jsVal.Swigcptr()))
					return C.jerry_value_t(*val)
				} else if err != nil {
					log.Panicln(err)
					jsError := createError(JERRY_ERROR_COMMON, err.Error())
					return uint32_t_To_Jerry_value_t(jsError)
				}
			}

		} else if err != nil {
			log.Panicln(err)
			jsError := createError(JERRY_ERROR_COMMON, err.Error())
			return uint32_t_To_Jerry_value_t(jsError)
		}
	}

	// here i will retrun a null value
	return uint32_t_To_Jerry_value_t(Jerry_create_undefined())
}

func uint32_t_To_Jerry_value_t(val Uint32_t) C.jerry_value_t {
	val_ := (*uintptr)(unsafe.Pointer(val.Swigcptr()))
	return C.jerry_value_t(*val_)
}

func jerry_value_t_To_uint32_t(val C.jerry_value_t) Uint32_t {
	return (Uint32_t)(SwigcptrUint32_t(C.uintptr_t((uintptr)(unsafe.Pointer(&val)))))

}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

////////////// Uint 8 //////////////
// The Uint8 Type represent a 8 bit char.
type Uint8 struct {
	// The pointer that old the data.
	ptr unsafe.Pointer
}

/**
 * Free the values.
 */
func (self Uint8) Free() {
	C.free(unsafe.Pointer(self.ptr))
}

/**
 * Access the undelying memeory values pointer.
 */
func (self Uint8) Swigcptr() uintptr {
	return uintptr(self.ptr)
}

/**
 * Create an error message.
 */
func createError(errorType int, errorMsg string) Uint32_t {
	msg := C.CString(errorMsg)
	err := C.create_error(C.jerry_error_t(errorType), msg)
	C.free(unsafe.Pointer(msg))
	return jerry_value_t_To_uint32_t(err)
}

/**
 * Create a new JerryScript String from go string
 */
func newJsString(val string) Uint32_t {
	str := C.create_string(C.CString(val))
	return jerry_value_t_To_uint32_t(str)
}

/**
 * Create a go string from a JS string pointer.
 */
func jsStrToGoStr(str Uint32_t) string {

	// Size info, ptr and it value
	size := C.size_t(C.get_string_size(uint32_t_To_Jerry_value_t(str)))

	buffer := (*C.char)(unsafe.Pointer(C.malloc(size)))

	// Test if the string is a valid utf8 string...
	C.string_to_utf8_char_buffer(uint32_t_To_Jerry_value_t(str), buffer, size)

	// Copy the value to a string.
	value := C.GoStringN(buffer, C.int(size))

	// free the buffer.
	C.free(unsafe.Pointer(buffer))

	return value
}

func goToJs(value interface{}) Uint32_t {
	var propValue Uint32_t
	var typeOf = reflect.TypeOf(value)

	if typeOf.Kind() == reflect.String {
		// String value
		propValue = newJsString(value.(string))

	} else if typeOf.Kind() == reflect.Bool {
		// Boolean value
		propValue = Jerry_create_boolean(value.(bool))
	} else if typeOf.Kind() == reflect.Int {
		propValue = Jerry_create_number(float64(value.(int)))
	} else if typeOf.Kind() == reflect.Int8 {
		propValue = Jerry_create_number(float64(value.(int8)))
	} else if typeOf.Kind() == reflect.Int16 {
		propValue = Jerry_create_number(float64(value.(int16)))
	} else if typeOf.Kind() == reflect.Int32 {
		propValue = Jerry_create_number(float64(value.(int32)))
	} else if typeOf.Kind() == reflect.Int64 {
		propValue = Jerry_create_number(float64(value.(int64)))
	} else if typeOf.Kind() == reflect.Uint {
		propValue = Jerry_create_number(float64(value.(uint)))
	} else if typeOf.Kind() == reflect.Uint8 {
		propValue = Jerry_create_number(float64(value.(uint8)))
	} else if typeOf.Kind() == reflect.Uint16 {
		propValue = Jerry_create_number(float64(value.(uint16)))
	} else if typeOf.Kind() == reflect.Uint32 {
		propValue = Jerry_create_number(float64(value.(uint32)))
	} else if reflect.TypeOf(value).Kind() == reflect.Uint64 {
		propValue = Jerry_create_number(float64(value.(uint64)))
	} else if typeOf.Kind() == reflect.Float32 {
		propValue = Jerry_create_number(float64(value.(float32)))
	} else if typeOf.Kind() == reflect.Float64 {
		propValue = Jerry_create_number(value.(float64))
	} else if typeOf.Kind() == reflect.Slice {

		// So here I will create a array and put value in it.
		s := reflect.ValueOf(value)
		l := uint32(s.Len())
		array := C.create_array(C.uint32_t(l))
		propValue = jerry_value_t_To_uint32_t(array)

		var i uint32
		for i = 0; i < l; i++ {
			v := goToJs(s.Index(int(i)).Interface())
			r := C.set_property_by_index(uint32_t_To_Jerry_value_t(propValue), C.uint32_t(i), uint32_t_To_Jerry_value_t(v))
			Jerry_release_value(jerry_value_t_To_uint32_t(r))
		}

	} else if typeOf.Kind() == reflect.Struct || typeOf.Kind() == reflect.Ptr {
		// So here I will use the object pointer address to generate it uuid value.
		ptrString := fmt.Sprintf("%d", value)
		uuid := Utility.GenerateUUID(ptrString)

		if GetCache().getObject(uuid) != nil {
			return GetCache().getJsObject(uuid)
		}

		// The JS object will be an handle to the Go object,
		// no property will be set in the Js object but getter and setter function
		// will be create instead.
		propValue = jerry_value_t_To_uint32_t(C.create_native_object(C.CString(uuid)))

		// First of all object properties.
		// Here I will make use of Go reflexion to create getter and setter.
		element := reflect.ValueOf(value).Elem()
		for i := 0; i < element.NumField(); i++ {
			valueField := element.Field(i)
			typeField := element.Type().Field(i)

			// Here I will set the property
			// Bust the number of field handler here.
			if valueField.CanInterface() {
				// So here is the field.
				fieldValue := goToJs(valueField.Interface())
				fieldName := goToJs(typeField.Name)
				defer Jerry_release_value(fieldValue)
				defer Jerry_release_value(fieldName)
				Jerry_release_value(Jerry_set_property(propValue, fieldName, fieldValue))
			}
		}

		// Register object method.
		for i := 0; i < element.Addr().NumMethod(); i++ {
			typeMethod := element.Addr().Type().Method(i)
			methodName := typeMethod.Name
			// Now I will create the method call
			setGoMethod(propValue, methodName, nil)
		}

		// Set the object in the cache.
		GetCache().setObject(uuid, value, propValue)

	} else if typeOf.String() == "GoJerryScript.SwigcptrUint32_t" {
		// already a Uint32_t
		propValue = value.(Uint32_t)
	} else {
		log.Panicln("---> type not found ", value, typeOf.String())
	}

	return propValue
}

/**
 * Return equivalent value of a 32 bit c pointer.
 */
func jsToGo(input Uint32_t) (interface{}, error) {

	// the Go value...
	var value interface{}

	// Now I will get the result if any...
	if Jerry_value_is_null(input) {
		return nil, nil
	} else if Jerry_value_is_undefined(input) {
		return nil, nil
	} else if Jerry_value_is_error(input) {
		// In that case I will return the error.
		log.Println("----> error found!")
	} else if Jerry_value_is_number(input) {
		value = Jerry_get_number_value(input)
	} else if Jerry_value_is_string(input) {
		value = jsStrToGoStr(input)
	} else if Jerry_value_is_boolean(input) {
		value = Jerry_get_boolean_value(input)
	} else if Jerry_value_is_typedarray(input) {
		/** Not made use of typed array **/
	} else if Jerry_value_is_array(input) {
		count := (uint32)(C.get_array_length(uint32_t_To_Jerry_value_t(input)))
		// So here I got a array without type so I will get it property by index
		// and interpret each result.
		value = make([]interface{}, 0)
		var i uint32
		for i = 0; i < count; i++ {
			e := jerry_value_t_To_uint32_t(C.get_property_by_index(uint32_t_To_Jerry_value_t(input), C.uint32_t(i)))
			//defer Jerry_release_value(e)
			v, err := jsToGo(e)
			if err == nil {
				value = append(value.([]interface{}), v)
			}
		}
		//Jerry_release_value(input)
	} else if Jerry_value_is_object(input) {
		// The go object will be a copy of the Js object.
		propUuid_ := goToJs("uuid_")
		defer Jerry_release_value(propUuid_)
		hasUuid_ := Jerry_has_own_property(input, propUuid_)
		defer Jerry_release_value(hasUuid_)
		if Jerry_get_boolean_value(hasUuid_) {
			uuid_ := Jerry_get_property(input, propUuid_)
			defer Jerry_release_value(uuid_)
			uuid, _ := jsToGo(uuid_)
			value = GetCache().getObject(uuid.(string))
		} else {
			stringified := Jerry_json_stringfy(input)
			// if there is no error
			if !Jerry_value_is_error(stringified) {
				jsonStr := jsStrToGoStr(stringified)
				data := make(map[string]interface{}, 0)
				err := json.Unmarshal([]byte(jsonStr), &data)
				if err == nil {
					if data["TYPENAME"] != nil {
						relfectValue := Utility.MakeInstance(data["TYPENAME"].(string), data, SetEntity)
						value = relfectValue.Interface()
					} else {
						// Here map[string]interface{} will be use.
						value = data
					}
				} else {
					return nil, err
				}
			} else {
				// Continue any way with nil object instead of an error...
				return nil, nil //errors.New("fail to stringfy object!")
			}
		}
	} else if Jerry_value_is_function(input) {
		// Here a function is found
		log.Println("---> function found!", input)
	} else if Jerry_value_is_abort(input) {
		// Here a function is found
		log.Println("--->abort!", input)
	} else if Jerry_value_is_arraybuffer(input) {
		// Here a function is found
		log.Println("--->array buffer!", input)
	} else if Jerry_value_is_constructor(input) {
		// Here a function is found
		log.Println("--->constructor!", input)
	} else if Jerry_value_is_promise(input) {
		// Here a function is found
		log.Println("--->promise!", input)
	} else {
		log.Panicln("---> not implemented Jerry value type.")
	}

	return value, nil
}

////////////// Uint 16 //////////////
// The Uint16 Type represent a 16 bit char.
type Uint16 struct {
	// The pointer that old the data.
	ptr unsafe.Pointer
}

/**
 * Free the values.
 */
func (self Uint16) Free() {
	C.free(unsafe.Pointer(self.ptr))
}

/**
 * Access the undelying memeory values pointer.
 */
func (self Uint16) Swigcptr() uintptr {
	return uintptr(self.ptr)
}

////////////// Uint 32 //////////////
// The Uint32 Type represent a 32 bit char.
type Uint32 struct {
	// The pointer that old the data.
	ptr unsafe.Pointer
}

func NewUint32FromInt(i int32) Uint32 {
	var val Uint32
	val.ptr = unsafe.Pointer(&i)
	return val
}

/**
 * Free the values.
 */
func (self Uint32) Free() {
	C.free(unsafe.Pointer(self.ptr))
}

/**
 * Access the undelying memeory values pointer.
 */
func (self Uint32) Swigcptr() uintptr {
	return uintptr(self.ptr)
}

////////////// Instance //////////////

// Reference to an object.
type Instance struct {
	// The pointer that old the data.
	ptr unsafe.Pointer
}

func NewInstance(obj interface{}) Instance {
	var instance Instance
	return instance
}

/**
 * Free the values.
 */
func (self Instance) Free() {
	C.free(unsafe.Pointer(self.ptr))
}

/**
 * Access the undelying memeory values pointer.
 */
func (self Instance) Swigcptr() uintptr {
	return uintptr(self.ptr)
}

/**
 * Variable with name and value.
 */
type Variable struct {
	Name  string
	Value interface{}
}

type Variables []Variable

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

func uintptrToBytes(u *uintptr) []byte {
	return (*[sizeOfUintPtr]byte)(unsafe.Pointer(u))[:]
}