package slices

/*
& A slice is a dynamically-sized, flexible view into the elements of an array.
& Slices are reference types. If you assign one slice to another, both refer to the same array.
 If you change an element in one slice, you'll see the change in the other.
& You don't need to define the size of a slice at the time of declaration.
& Slices can be resized using the built-in append function.
& Slice type doesn't include its size, which means that []int is a type that includes all slices of ints, regardless of their length.
*/
