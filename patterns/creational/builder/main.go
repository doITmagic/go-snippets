package main

import "fmt"

// Product represents the final object being constructed.
type Product struct {
	part1 string
	part2 int
	part3 bool
}

// Builder is an interface that defines the steps to build the product.
type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetResult() Product
}

// ConcreteBuilder implements the Builder interface.
type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) BuildPart1() {
	b.product.part1 = "Part 1"
}

func (b *ConcreteBuilder) BuildPart2() {
	b.product.part2 = 42
}

func (b *ConcreteBuilder) BuildPart3() {
	b.product.part3 = true
}

func (b *ConcreteBuilder) GetResult() Product {
	return b.product
}

// Director controls the construction process using a builder.
type Director struct {
	builder Builder
}

func (d *Director) Construct() Product {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
	return d.builder.GetResult()
}

func main() {
	builder := &ConcreteBuilder{}
	director := &Director{builder: builder}

	product := director.Construct()
	fmt.Println(product)
}
