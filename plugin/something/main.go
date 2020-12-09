package main

type Igor struct {
	Name string
}

func Something(v ...interface{}) (interface{}, error) {
	igor := new(Igor)
	igor.Name = v[0].(string)
	return igor, nil
}
