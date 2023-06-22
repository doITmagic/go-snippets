In this Decorator pattern example, we have a `Component` interface that defines the basic operations. `ConcreteComponent` is a concrete implementation of the `Component` interface.
The `Decorator` interface extends the `Component` interface and represents the base interface for all decorators. `ConcreteDecoratorA` and `ConcreteDecoratorB` are concrete implementations of the `Decorator` interface.
The decorators wrap an instance of the component and add additional behavior to its operation. Each decorator invokes the operation of the wrapped component and adds its own behavior before or after the invocation.
In the `Example` function, we create a `ConcreteComponent` instance. We then create decorators (`ConcreteDecoratorA` and `ConcreteDecoratorB`) and wrap the component instance with them. Finally, we call the `Operation` method on the final wrapped component, which triggers the entire chain of decorators and component.
The output of the program will be:
```go 
ConcreteDecoratorB(ConcreteDecoratorA(ConcreteComponent))
```

This demonstrates how the Decorator pattern allows us to dynamically add or modify behavior of an object by wrapping it with decorators. Each decorator adds its own functionality while still maintaining the original interface and behavior of the component.
