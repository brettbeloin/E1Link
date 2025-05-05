package main

type gen interface {
	~int | ~string | ~float64
}

type node[T gen] struct {
	value T
	next  *node[T]
}

type link[T gen] struct {
	head node[T]
	tail node[T]

	// Returns the length of the list
	count int
}

func main() {

}

// Puts a new value at the Tail end of the list
func (l *link[T]) add() {

}

// Inserts a new value at a given index, pushing the existing value at that index to the next index spot, and so on. Insert may ONLY target indices that are currently in use. In other words, if you have 5 elements in your list, you may insert at any index between 0 and 4 inclusive. Index 5 would be considered out of bounds as it is not currently in use during the insertion process. Any index less than zero or equal to or greater than Count should throw an index out of bounds exception.
func (l *link[T]) insert(val T, idx int) {

}

// Returns the value at the given index. Any index less than zero or equal to or greater than Count should throw an index out of bounds exception.
func (l *link[T]) get(idx int) {

}

// Removes and returns the first value in the list
func (l *link[T]) remove() {

}

// Removes and returns the last value in the list
func (l *link[T]) removeLast() {

}

// Removes and returns the value at a given index. Any index less than zero or equal to or greater than Count should throw an index out of bounds exception.
func (l *link[T]) removeAt(idx int) {

}

// Removes all values in the list.
func (l *link[T]) clear() {

}

// Searches for a value in the list and returns the first index of that value when found. If the key is not found in the list, the method returns -1.
func (l *link[T]) search(val T) {

}

// An override method that creates and returns a string representation of all the values in the list. The string must be in the format of “v0, v1, v2, .., vn-1” where n-1 is the last index in the list. An empty list should return an empty string (but not null). While every value in the string is separated by a comma and space, the string must NOT have any unnecessary commas or spaces at the beginning or end.
func (l *link[T]) toString() {

}

func singleLinkList() {

}

func doubleLinkList() {

}
