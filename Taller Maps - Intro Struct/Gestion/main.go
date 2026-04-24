package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func (p Producto) CalcularTotal() float64 {
	return p.Precio * float64(p.Cantidad)
}

func (p Producto) HayStock() bool {
	return p.Cantidad > 0
}

func imprimirProducto(p Producto) {
	estado := "disponible"
	if !p.HayStock() {
		estado = "sin stock"
	}

	fmt.Printf("Producto: %s | Precio: %.2f | Stock: %d | Total: %.2f | Estado: %s\n",
		p.Nombre, p.Precio, p.Cantidad, p.CalcularTotal(), estado)
}

func main() {
	p1 := Producto{"Celular", 500.0, 10}
	p2 := Producto{"Teclado", 45.50, 0}
	p3 := Producto{"Monitor", 250.0, 5}

	inventario := []Producto{p1, p2, p3}

	for _, p := range inventario {
		imprimirProducto(p)
	}
}
