package package1

var someVar = 4545454545454545

// SomeExportedVar - some exported variable to other packages
var SomeExportedVar = 4545454545454545

//GetData - some testing function
func GetData() []int {
	return []int{1, 2, 3, 4}
}

func somePrivateMethodGetData() []int {
	return []int{1, 2, 3, 4}
}
