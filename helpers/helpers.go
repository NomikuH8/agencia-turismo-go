package helpers

func RemoveSliceAtIndex(slice []any, index int) []any {
	newSlice := append(slice, slice[:index], slice[index+1])
	return newSlice
}
