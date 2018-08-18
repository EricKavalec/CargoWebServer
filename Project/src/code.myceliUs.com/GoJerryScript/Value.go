package GoJerryScript

import "reflect"
import "code.myceliUs.com/Utility"
import "errors"
import "log"

/**
 * Create a new value and set it finalyse methode.
 */
func NewValue(ptr Uint32_t) *Value {

	v := new(Value)
	var err error

	// Export the value.
	v.Val, err = jsToGo(ptr)

	if err != nil {
		log.Println("---> error: ", err)
		return nil
	}

	return v
}

/**
 * Interface Js value as Go value.
 */
type Value struct {
	// The go value that reflect the js value.
	Val interface{}
}

/**
 * If the value is an object it can be use as so.
 */
func (self Value) Object() *Object {
	if reflect.TypeOf(self.Val).String() == "*GoJerryScript.Object" {
		return self.Val.(*Object)
	}

	// not an object.
	return nil
}

/**
 * Export the Javascript value in Go.
 */
func (self Value) Export() (interface{}, error) {
	return self.Val, nil
}

/**
 * Type validation function.
 */
func (self *Value) IsString() bool {
	if reflect.TypeOf(self.Val).Kind() == reflect.String {
		return true
	}
	return false
}

func (self *Value) ToString() (string, error) {
	if !self.IsString() {
		return "", errors.New("The value is not a string!")
	}
	return self.Val.(string), nil
}

func (self *Value) IsBoolean() bool {
	if reflect.TypeOf(self.Val).Kind() == reflect.Bool {
		return true
	}
	return false
}

func (self *Value) ToBoolean() (bool, error) {
	if !self.IsBoolean() {
		return false, errors.New("The value is not a boolean!")
	}
	return self.Val.(bool), nil
}

func (self *Value) IsNull() bool {
	return self.Val == nil
}

func (self *Value) IsUndefined() bool {
	return self.Val == nil
}

func (self *Value) IsNumber() bool {
	if reflect.TypeOf(self.Val).Kind() == reflect.Float32 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Float64 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Int ||
		reflect.TypeOf(self.Val).Kind() == reflect.Int8 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Int16 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Int32 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Int64 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Uint ||
		reflect.TypeOf(self.Val).Kind() == reflect.Uint8 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Uint16 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Uint32 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Uint64 {
		return true
	}
	return false
}

func (self *Value) ToFloat() (float64, error) {
	if !self.IsNumber() {
		return 0.0, errors.New("The value is not a numeric!")
	}

	return self.Val.(float64), nil
}

func (self *Value) ToInteger() (int64, error) {
	if !(reflect.TypeOf(self.Val).Kind() == reflect.Int ||
		reflect.TypeOf(self.Val).Kind() == reflect.Int8 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Int16 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Int32 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Int64 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Uint ||
		reflect.TypeOf(self.Val).Kind() == reflect.Uint8 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Uint16 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Uint32 ||
		reflect.TypeOf(self.Val).Kind() == reflect.Uint64) {
		return 0, errors.New("The value is not an Interger!")
	}

	return int64(Utility.ToInt(self.Val)), nil
}