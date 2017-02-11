package Utility

import (
	b64 "encoding/base64"
	"encoding/gob"
	"errors"
	"log"
	"reflect"
	"strconv"
	"strings"
)

/**
 * That map will contain the list of all type to be created dynamicaly
 */
var typeRegistry = make(map[string]reflect.Type)

func GetTypeOf(typeName string) reflect.Type {
	if t, ok := typeRegistry[typeName]; ok {
		return reflect.New(t).Type()
	}
	return nil
}

/**
 * Register an instance of the type.
 */
func RegisterType(typedNil interface{}) {

	t := reflect.TypeOf(typedNil).Elem()
	index := strings.LastIndex(t.PkgPath(), "/")
	var typeName = t.Name()
	if strings.HasSuffix(typeName, "_impl") == true {
		typeName = strings.Replace(typeName, "_impl", "", -1)
	}
	if _, ok := typeRegistry[t.PkgPath()[index+1:]+"."+typeName]; !ok {
		if index > 0 {

			typeRegistry[t.PkgPath()[index+1:]+"."+typeName] = t
			gob.RegisterName(t.PkgPath()[index+1:]+"."+typeName, typedNil)
			//log.Println("------> type: ", t.PkgPath()[index+1:]+"."+typeName, " was register as dynamic type.")
		} else {
			typeRegistry[t.PkgPath()+"."+typeName] = t
			gob.RegisterName(t.PkgPath()+"."+typeName, typedNil)
			//log.Println("------> type: ", t.PkgPath()+"."+typeName, " was register as dynamic type.")
		}
	}
}

func toInt(value interface{}) int {
	switch v := value.(type) {
	case string:
		result, err := strconv.Atoi(v)
		if err == nil {
			return int(result)
		}
		return 0
	case int8:
		return int(value.(int8))
	case int32:
		return int(value.(int32))
	case int64:
		return int(value.(int64))
	case float64:
		return int(value.(float64))
	case float32:
		return int(value.(float32))
	default:
		log.Println("type not found with type name ", reflect.TypeOf(value).String())
		return 0
	}
}

/**
 * Initialyse base type value.
 */
func initializeBaseTypeValue(t reflect.Type, value interface{}) reflect.Value {

	var v reflect.Value

	switch t.Kind() {
	case reflect.String:
		// Here it's possible that the value contain the map of values...
		// I that case I will
		v = reflect.ValueOf(value.(string))
	case reflect.Bool:
		v = reflect.ValueOf(value.(bool))
	case reflect.Int:
		v = reflect.ValueOf(toInt(value))
	case reflect.Int8:
		v = reflect.ValueOf(toInt(value))
	case reflect.Int32:
		v = reflect.ValueOf(toInt(value))
	case reflect.Int64:
		v = reflect.ValueOf(toInt(value))
	case reflect.Uint:
		v = reflect.ValueOf(value.(uint64))
	case reflect.Uint8:
		v = reflect.ValueOf(value.(uint64))
	case reflect.Uint32:
		v = reflect.ValueOf(value.(uint64))
	case reflect.Uint64:
		v = reflect.ValueOf(value.(uint64))
	case reflect.Float32:
		v = reflect.ValueOf(value.(float64))
	case reflect.Float64:
		v = reflect.ValueOf(value.(float64))
	default:
		log.Println("unexpected type %T\n", t)
	}

	return v
}

/**
 * Create an instance of the type with it name.
 */
func MakeInstance(typeName string, data map[string]interface{}) reflect.Value {
	value := initializeStructureValue(typeName, data)
	return value
}

/**
 * Intialyse the struct fields with the values contain in the map.
 */
func initializeStructureValue(typeName string, data map[string]interface{}) reflect.Value {

	// Here I will create the value...
	t := typeRegistry[typeName]
	if t == nil {
		return reflect.ValueOf(data)
	}
	v := reflect.New(t)
	for name, value := range data {
		ft, exist := t.FieldByName(name)
		if exist && value != nil {
			switch ft.Type.Kind() {
			case reflect.Slice:
				// That's mean the value contain an array...
				if reflect.TypeOf(value).String() == "[]interface {}" {
					values := value.([]interface{})
					for i := 0; i < len(values); i++ {
						var fv reflect.Value
						switch v_ := values[i].(type) {
						// Here i have a sub-value.
						case map[string]interface{}:
							fv = initializeStructureValue(v_["TYPENAME"].(string), v_)

						default:
							// A base type...
							// Here I will try to convert the base type to the one I have in
							// the structure.
							if v.Elem().FieldByName(name).Type().String() == "[]string" {
								var string_ string
								fv = initializeBaseTypeValue(reflect.TypeOf(string_), v_)
							} else if v.Elem().FieldByName(name).Type().String() == "[]int" {
								var int_ int
								fv = initializeBaseTypeValue(reflect.TypeOf(int_), v_)
							} else if v.Elem().FieldByName(name).Type().String() == "[]float" {
								var float_ float64
								fv = initializeBaseTypeValue(reflect.TypeOf(float_), v_)
							} else if v.Elem().FieldByName(name).Type().String() == "[]bool" {
								var bool_ bool
								fv = initializeBaseTypeValue(reflect.TypeOf(bool_), v_)
							} else {
								fv = initializeBaseTypeValue(reflect.TypeOf(v_), v_)
							}

						}
						v.Elem().FieldByName(name).Set(reflect.Append(v.Elem().FieldByName(name), fv))
					}
				} else {
					// Here the value is a base type...
					fv := initializeBaseTypeValue(reflect.TypeOf(value), value)
					if ft.Type.String() != fv.Type().String() {
						// So here a conversion is necessary...
						if ft.Type.String() == "[]uint8" || ft.Type.String() == "[]byte" || fv.Type().String() == "string" {
							val := fv.String()
							val_, err := b64.StdEncoding.DecodeString(val)
							if err == nil {
								val = string(val_)
							}
							// Set the value...
							v.Elem().FieldByName(name).Set(reflect.ValueOf([]byte(val)))
						}
					} else {
						v.Elem().FieldByName(name).Set(fv)
					}
				}
			case reflect.Struct:
				fv, _ := InitializeStructure(value.(map[string]interface{}))
				v.Elem().FieldByName(name).Set(fv.Elem())
			case reflect.Ptr:
				fv, _ := InitializeStructure(value.(map[string]interface{}))
				v.Elem().FieldByName(name).Set(fv)
			case reflect.Interface:
				// To recurse is divine!-)
				if reflect.TypeOf(value).String() == "map[string]interface {}" {
					if typeName_, ok := value.(map[string]interface{})["TYPENAME"]; ok {
						if _, ok := typeRegistry[typeName_.(string)]; ok {
							fv, _ := InitializeStructure(value.(map[string]interface{}))
							v.Elem().FieldByName(name).Set(fv.Elem())
						} else {
							// Here it's a dynamic entity...
							v.Elem().FieldByName(name).Set(reflect.ValueOf(value))
						}
					} else {
						v.Elem().FieldByName(name).Set(reflect.ValueOf(value))
					}
				}
			default:
				// Convert is use to enumeration type who are int and must be convert to
				// it const type representation.
				fv := initializeBaseTypeValue(ft.Type, value).Convert(ft.Type)
				v.Elem().FieldByName(name).Set(fv)
			}
		}
	}

	// Return the initialysed value...
	return v
}

/**
 * Initialyse an array of structures, return it as interface (array of the actual
 * objects)
 */
func InitializeStructures(data []interface{}, typeName string) (reflect.Value, error) {
	// Here I will get the type name, only dynamic type can be use here...
	var values reflect.Value
	if len(data) > 0 {
		// Structure data must be a map[string]interface{}
		if _, ok := data[0].(map[string]interface{}); ok {
			if typeName_, ok := data[0].(map[string]interface{})["TYPENAME"]; ok {
				// Now I will create empty structure and initialyse it with the value found in the map values.
				for i := 0; i < len(data); i++ {
					obj := MakeInstance(typeName_.(string), data[i].(map[string]interface{}))
					if i == 0 {
						if len(typeName) == 0 {
							values = reflect.MakeSlice(reflect.SliceOf(obj.Type()), 0, 0)
						} else if t, ok := typeRegistry[typeName]; ok {
							values = reflect.MakeSlice(reflect.SliceOf(reflect.New(t).Type()), 0, 0)
						} else {
							emptyInterfaceArray := make([]interface{}, 0, 0)
							values = reflect.ValueOf(emptyInterfaceArray)
						}
					}
					values = reflect.Append(values, obj)
				}
				return values, nil
			} else {
				return reflect.ValueOf(data), nil
			}
		} else {
			return values, errors.New("NotDynamicObject")
		}
	} else {
		// Here there is no value in the array.
		if t, ok := typeRegistry[typeName]; ok {
			values = reflect.MakeSlice(reflect.SliceOf(reflect.New(t).Type()), 0, 0)
		} else {
			emptyInterfaceArray := make([]interface{}, 0, 0)
			values = reflect.ValueOf(emptyInterfaceArray)
		}
	}
	return values, nil
}

/**
 * Initialyse a single object from it value.
 */
func InitializeStructure(data map[string]interface{}) (reflect.Value, error) {
	// Here I will get the type name, only dynamic type can be use here...
	var value reflect.Value
	if typeName, ok := data["TYPENAME"]; ok {
		if _, ok := typeRegistry[typeName.(string)]; ok {
			value = MakeInstance(typeName.(string), data)
			return value, nil
		} else {
			// Return the value itself...
			return reflect.ValueOf(data), nil
		}
	} else {
		return value, errors.New("NotDynamicObject")
	}
}

/**
 * Initialyse an array of values other than structure...
 */
func InitializeArray(data []interface{}, typeName string) (reflect.Value, error) {

	var values reflect.Value

	if strings.HasPrefix(typeName, "[]") {
		emptyInterfaceArray := make([]interface{}, 0, 0)
		values = reflect.ValueOf(emptyInterfaceArray)
	}

	for i := 0; i < len(data); i++ {
		if i == 0 {
			if len(typeName) == 0 {
				values = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(data[i])), 0, 0)
			}
		}
		values = reflect.Append(values, reflect.ValueOf(data[i]))
	}
	return values, nil
}