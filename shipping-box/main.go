package shipping_box

type Product struct {
	Name string
	Len  int
	Wid  int
	Hei  int
}

type Box struct {
	Len int
	Wid int
	Hei int
}

func getBestBox(avaliableBoxes []Box, products []Product) Box {

}
