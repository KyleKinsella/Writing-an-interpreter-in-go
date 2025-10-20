package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// the below code is for: "extending the environment", this means that we create a new instance of object.Environment with a pointer
// to the environment it should extend. By doing this, we enclose a fresh and empty environment with an existing one.

// we have also done some refactoring to the following functions: "Get(...), NewEnvironment(), added: "	outer *Environment " to our 
// Environment struct, and wrote the below function: NewEnclosedEnvironment(...) "

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}