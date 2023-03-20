// package classification of product API
//
// Documentation for product API
//
// Schemes: http
// BasePath: /
// Version: 1.20.2
//
// Consumes:
// -applications/json
//
// Produces:
// -applications/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"github.com/areeb529/go-microservices/product-api/data"
)

//A list of products returns in the response
// swagger:response productsResponse
type productsResponseWrapper struct{
	// All products in the system
	// in: body
	Body []data.Product
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}


type KeyProduct struct{}

func (p *Products)MiddlewareValidateProduct(next http.Handler)(http.Handler){
	return http.HandlerFunc(func (rw http.ResponseWriter,r *http.Request){
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product",err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}
		// validate the product
		err = prod.Validate()
		if err != nil{
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error reading product: %s"),
				http.StatusBadRequest,
			)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(),KeyProduct{},prod)
		req := r.WithContext(ctx)

		// call the next handler which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, req)
	})
}
