package maps

// Dictionary custom type to abstract the map logic
type Dictionary map[string]string

// Search : a search function that takes in a dictionary and a string and returns the value form a key
func (d Dictionary) Search(word string) string {
	return d[word]
}