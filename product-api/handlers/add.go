package handlers

import (
	"net/http"
	"github.com/areeb529/go-microservices/product-api/data"
)


func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}