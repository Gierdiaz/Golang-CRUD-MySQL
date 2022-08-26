package models

type Product struct {
	Id           int
	Description  string
	Quantity     int
	Price        float64
	Amount       float64
	CreatedAt    string
}

func GetAll() ([]Product, error) {
	con := Connect()

	sql := "select * from product"

	rs, erro := con.Query(sql)

	if erro != nil {
		return nil, erro
	}

	var products []Product

	for rs.Next() {
		var product Product
		
		erro := rs.Scan(
			&product.Id,
			&product.Description,
			&product.Quantity,
			&product.Price,
			&product.Amount,
			&product.CreatedAt)

		if erro != nil {
			return nil, erro
		}

		products = append(products, product)

	}

	defer rs.Close()
	defer con.Close()

	return products, nil

}