package builtins

import (
	. "github.com/prologic/monkey-lang/object"
)

// First ...
func First(args ...Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1",
			len(args))
	}
	if args[0].Type() != ARRAY {
		return newError("argument to `first` must be array, got %s",
			args[0].Type())
	}

	arr := args[0].(*Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return nil
}
