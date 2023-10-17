package test2

type Test2 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetTest2() Test2 {
	return Test2{
		Id:   1,
		Name: "zs",
	}
}
