package handlers

import (
	"net/http"

	"github.com/Sahas001/go-micro/data"
)

//  swagger:route POST /products products addProducts
//  Create a new Product
//
//  responses:
//    200: productResponse
//    422: errorValidation
//    501: errorResponse

func (p *Products) AddProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
