package main

import (
	"fmt"
	"src/client"
	"src/client/products"
	"testing"
)

func TestOurClient(t *testing.T) {
	// cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.Default

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
	t.Fail()
}
