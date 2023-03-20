package handlers

import (
	"net/http"
	"github.com/areeb529/go-microservices/product-api/data"
)

// swagger:route GET /products products listProducts
//returns a list of products
//responses:
//	200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	//fetch the products from datastore
	lp := data.GetProducts()
	//serialize the list to JSON
	err := lp.ToJSON(rw)         
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}