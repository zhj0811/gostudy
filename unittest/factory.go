package unittest

//go:generate mockgen -destination mocks/mock_factory.go -source=factory.go
// Factory interface
type Factory interface {
	Println(args ...string)
}

// // Data struct
// type Data struct {
// 	Key   string `json:"key"`
// 	Value string `json:"value"`
// }

// // Println func
// func (data *Data) Println(args ...string) {
// 	fmt.Printf("%v", data)
// 	fmt.Println(args)
// }
