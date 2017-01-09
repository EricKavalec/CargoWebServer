package Server

import "fmt"
import "strings"
import "log"
import "code.myceliUs.com/CargoWebServer/Cargo/Utility"

/**
 * Restrictions for Datatypes
 */
type RestrictionType int

const (
	// Defines a list of acceptable values
	RestrictionType_Enumeration RestrictionType = 1 + iota

	// Specifies the maximum number of decimal places allowed. Must be equal to
	// or greater than zero
	RestrictionType_FractionDigits

	// Specifies the exact number of characters or list items allowed. Must be
	// equal to or greater than zero
	RestrictionType_Length

	// Specifies the upper bounds for numeric values (the value must be less
	// than this value)
	RestrictionType_MaxExclusive

	// Specifies the upper bounds for numeric values (the value must be less than
	// or equal to this value)
	RestrictionType_MaxInclusive

	// Specifies the maximum number of characters or list items allowed. Must be
	// equal to or greater than zero
	RestrictionType_MaxLength

	// Specifies the lower bounds for numeric values (the value must be greater
	// than this value)
	RestrictionType_MinExclusive

	// Specifies the lower bounds for numeric values (the value must be greater than or equal to this value)
	RestrictionType_MinInclusive

	// Specifies the minimum number of characters or list items allowed. Must be equal to or greater than zero
	RestrictionType_MinLength

	// Defines the exact sequence of characters that are acceptable
	RestrictionType_Pattern

	// Specifies the exact number of digits allowed. Must be greater than zero
	RestrictionType_TotalDigits

	// Specifies how white space (line feeds, tabs, spaces, and carriage returns) is handled
	RestrictionType_WhiteSpace
)

// Put constaint in a field to reduce the range of
// possibles values of a given type.
// For example an email is a string with a pattern to respect.
// so the range of string restrict by a pattern became the range of
// email.
type Restriction struct {
	// The the of the restriction (Facet)
	Type RestrictionType
	// The value...
	Value string
}

/**
 * This structure is use to make query over the key value data store.
 */
type EntityPrototype struct {

	// The name of the entity
	// The type name is compose of the package name, a comma and
	// the type name itself.
	TypeName string

	// The documentation for that entity
	Documentation string

	// True if the entity prototype is an abstrac class...
	IsAbstract bool

	// In that case the prototype define a list of given item type.
	ListOf string

	// The class derived from this entity.
	SubstitutionGroup []string

	// The list of super type, equivalent to extension.
	SuperTypeNames []string

	// Restriction of the range of possible value.
	Restrictions []*Restriction

	// The ids that compose the entity...
	Ids []string

	// The indexation of this entity
	Indexs []string

	// The list of fields of the entity
	Fields []string

	// That contain the field documentation if there is so...
	FieldsDocumentation []string

	// The list of fields type of the entity
	// ex. []string:Test.Item:Ref
	// [] means the field is an array
	// string is the format of the reference in the case
	// of type other than xsd base type.
	// Type is written like PacakageName.TypeName
	// If the field is a reference to other element, (an aggregation)
	// Ref is needed at the end. Otherwise it considere at composition.
	FieldsType []string

	// Fields visibility
	FieldsVisibility []bool

	// If the field can be nil value...
	FieldsNillable []bool

	// The order of the field, use to display in tabular form...
	FieldsOrder []int

	// The prototype version.
	Version string
}

func NewEntityPrototype() *EntityPrototype {

	prototype := new(EntityPrototype)

	prototype.Fields = make([]string, 0)
	prototype.FieldsOrder = make([]int, 0)
	prototype.FieldsType = make([]string, 0)
	prototype.FieldsDocumentation = make([]string, 0)
	prototype.FieldsNillable = make([]bool, 0)
	prototype.FieldsVisibility = make([]bool, 0)
	prototype.SuperTypeNames = make([]string, 0)
	prototype.Restrictions = make([]*Restriction, 0)
	prototype.Indexs = make([]string, 0)
	prototype.Ids = make([]string, 0)

	// Append the default fields at begin...
	prototype.Fields = append(prototype.Fields, "uuid")
	prototype.Ids = append(prototype.Ids, "uuid")
	prototype.FieldsOrder = append(prototype.FieldsOrder, 0)
	prototype.FieldsType = append(prototype.FieldsType, "xs.string")
	prototype.FieldsVisibility = append(prototype.FieldsVisibility, false)

	prototype.Fields = append(prototype.Fields, "parentUuid")
	prototype.Indexs = append(prototype.Indexs, "parentUuid")
	prototype.FieldsOrder = append(prototype.FieldsOrder, 1)
	prototype.FieldsType = append(prototype.FieldsType, "xs.string")
	prototype.FieldsVisibility = append(prototype.FieldsVisibility, false)

	return prototype
}

/**
 * This function is use to retreive the position in the array of a given field.
 */
func (this *EntityPrototype) getFieldIndex(fieldName string) int {

	if this.Fields != nil {
		for i := 0; i < len(this.Fields); i++ {
			if this.Fields[i] == fieldName {
				return i
			}
		}
	}
	return -1
}

/**
 * Save the new entity prototype in the data store.
 */
func (this *EntityPrototype) Create() {

	// Append the default fields at end...

	// The list of childs uuid use by this entity
	if Utility.Contains(this.Fields, "childsUuid") == false {
		this.Fields = append(this.Fields, "childsUuid")
		this.FieldsOrder = append(this.FieldsOrder, len(this.FieldsOrder))
		this.FieldsType = append(this.FieldsType, "[]xs.string")
		this.FieldsVisibility = append(this.FieldsVisibility, false)
	}

	// The list of entity referenced by this entity
	if Utility.Contains(this.Fields, "referenced") == false {
		this.Fields = append(this.Fields, "referenced")
		this.FieldsOrder = append(this.FieldsOrder, len(this.FieldsOrder))
		this.FieldsType = append(this.FieldsType, "[]EntityRef")
		this.FieldsVisibility = append(this.FieldsVisibility, false)
	}

	storeId := this.TypeName[:strings.Index(this.TypeName, ".")]
	store := GetServer().GetDataManager().getDataStore(storeId).(*KeyValueDataStore)
	if store != nil {
		err := store.SetEntityPrototype(this)
		if err != nil {
			log.Println("Fail to save entity prototype ", this.TypeName, " in store id ", storeId)
		}
	}
}

/**
 * For debug purpose only...
 */
func (this *EntityPrototype) Print() {
	// The prototype Type Name...
	fmt.Println("\nTypeName:", this.TypeName)
	if len(this.SuperTypeNames) > 0 {
		fmt.Println("	Super Types:", this.SuperTypeNames)
	}

	if this.SubstitutionGroup != nil {
		fmt.Println("	Substitution Groups:", this.SubstitutionGroup)
	}

	if len(this.ListOf) > 0 {
		fmt.Println("	List of:", this.ListOf)
	}

	// Now the restrictions...
	if this.Restrictions != nil {
		for j := 0; j < len(this.Restrictions); j++ {
			if this.Restrictions[j].Type == RestrictionType_Enumeration {
				fmt.Println("	----> Enumration Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_FractionDigits {
				fmt.Println("	----> Fraction Digits Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_Length {
				fmt.Println("	----> Length Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MaxExclusive {
				fmt.Println("	----> Max Exclusive Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MaxInclusive {
				fmt.Println("	----> Max Inclusive Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MaxLength {
				fmt.Println("	----> Max Length Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MinExclusive {
				fmt.Println("	----> Min Exclusive Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MinInclusive {
				fmt.Println("	----> Min Inclusive Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_MinLength {
				fmt.Println("	----> Min Length Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_TotalDigits {
				fmt.Println("	----> Total Digits Restriction:", this.Restrictions[j].Value)
			} else if this.Restrictions[j].Type == RestrictionType_WhiteSpace {
				fmt.Println("	----> White Space Restriction:", this.Restrictions[j].Value)
			}
		}
	}

	// Now the fields...
	fmt.Println("	Fields:")
	for i := 0; i < len(this.Fields); i++ {
		if this.FieldsVisibility[i] == true {
			fmt.Println("	-->", this.Fields[i], ":", this.FieldsType[i])
		}
	}
}