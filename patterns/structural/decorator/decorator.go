package decorator

import (
	"fmt"
)

// Component represents the base component interface.
type Component interface {
	Operation() string
}

// ConcreteComponent represents a concrete implementation of the component interface.
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// Decorator represents the base decorator interface.
type Decorator interface {
	Component
}

// ConcreteDecoratorA represents a concrete decorator implementation.
type ConcreteDecoratorA struct {
	component Component
}

func (d *ConcreteDecoratorA) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorA(%s)", d.component.Operation())
}

// ConcreteDecoratorB represents another concrete decorator implementation.
type ConcreteDecoratorB struct {
	component Component
}

func (d *ConcreteDecoratorB) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorB(%s)", d.component.Operation())
}

func Example() {
	// Create a ConcreteComponent.
	component := &ConcreteComponent{}

	// Create decorators and wrap the component.
	decoratorA := &ConcreteDecoratorA{component: component}
	decoratorB := &ConcreteDecoratorB{component: decoratorA}

	// Call the operation on the final wrapped component.
	result := decoratorB.Operation()
	fmt.Println(result)

	// Output:
	// Received: ConcreteDecoratorB(ConcreteDecoratorA(ConcreteComponent))
}
