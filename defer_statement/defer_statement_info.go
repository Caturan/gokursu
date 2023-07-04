package deferstatement

/*
	The 'defer' statement in Go is used to ensure that a function call is performed later in a program's execution
	 usually for purposes of cleanup.

	The 'defer' keyword is used to postpone the execution of a function until the surrounding function returns,
	 whether that function returns normally at the end of its body,
	 or returns early because of an error or 'return' statement.

	The deferred calls are executed in "last in, first out" orede, so if multiple functions are deferred
	 they will be executed in the reverse order they were deferred.
*/
