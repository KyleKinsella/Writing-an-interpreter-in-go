package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

// Any type that has both a Type() ObjectType method and an Inspect() string method automatically satisfies the Object interface. // 

// Now the *Integer type (a pointer to Integer) has both methods that the Object interface requires:

// Type() returns an ObjectType
// Inspect() returns a string

// So Go automatically considers *Integer to implement the Object interface.
// No explicit keyword (like implements) is needed.

func (i *Integer) Inspect() string {return fmt.Sprintf("%d", i.Value)}
func (i *Integer) Type() ObjectType {return INTEGER_OBJ}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {return BOOLEAN_OBJ}
func (b *Boolean) Inspect() string {return fmt.Sprintf("%t", b.Value)}

type Null struct {}

func (n *Null) Type() ObjectType {return NULL_OBJ}
func (n *Null) Inspect() string {return "null"}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {return RETURN_VALUE_OBJ}
func (rv *ReturnValue) Inspect() string {return rv.Value.Inspect()}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {return ERROR_OBJ}
func (e *Error) Inspect() string {return "ERROR: " + e.Message}