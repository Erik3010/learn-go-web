package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (product Product) CekStock() (status string) {
	if product.Stock < 3 {
		status = "Stock hampir habis"
	} else if product.Stock < 10 {
		status = "Stock terbatas"
	}
	return
}
