In this example, we have a Product struct representing the final object to be built. 
The Builder interface defines the steps required to build the product, including BuildPart1(), BuildPart2(), BuildPart3(), and GetResult().
The ConcreteBuilder struct implements the Builder interface. 
It has its own instance of the product and implements each step of the building process.
The Director struct controls the construction process using a builder. It calls the builder's methods in a specific order to construct the product.
In the main function, we create a ConcreteBuilder instance and a Director instance. 
We then use the director to construct the product by calling the builder's methods. 
Finally, we retrieve the constructed product using GetResult() and print it.
The output of the program will be:
```sh
{Part 1 42 true}
```
This demonstrates how the Builder pattern separates the construction of a complex object from its representation. 
The director controls the building process, while the concrete builder implements the steps to construct the object. 
This allows for more flexibility and easier extension of the construction process, as well as the ability to construct different variations of the final object.






