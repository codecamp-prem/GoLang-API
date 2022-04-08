package models

// Product struct
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float32 `json:"price"`
}

//CreateProduct : Insert a new product
func CreateProduct(products *[]Product, newProduct *Product) (err error) {
	*products = append(*products, *newProduct)
	return nil
}

//ReadProductByID : get product by ID
func ReadProductByID(products *[]Product, id int) (product Product) {
	for _, item := range *products {
		if item.ID == id {
			return item
		}
	}
	return Product{}
}

// UpdateProductByID : updatedProduct by id
func UpdateProductByID(products *[]Product, updatedProduct *Product) (product Product) {
	for index, item := range *products {
		if item.ID == (*updatedProduct).ID {
			(*products)[index].Name = (*updatedProduct).Name
			(*products)[index].Stock = (*updatedProduct).Stock
			(*products)[index].Price = (*updatedProduct).Price

			return (*products)[index]
		}
	}
	return Product{}
}

//DeleteProductByID : delete product by ID
func DeleteProductByID(products *[]Product, id int) (product Product) {
	var deletedProduct Product
	for index, item := range *products {
		if item.ID == id {
			n := len(*products)
			deletedProduct = item
			(*products)[index] = (*products)[n-1]
			*products = (*products)[:n-1]
			return deletedProduct
		}
	}
	return Product{}
}
