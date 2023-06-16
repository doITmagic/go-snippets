The Factory Method defines a method, which should be used for creating objects instead of using a direct constructor call (new operator). 
Subclasses can override this method to change the class of objects that will be created.

It’s impossible to implement the classic Factory Method pattern in Go due to lack of OOP features such as classes and inheritance. 
However, we can still implement the basic version of the pattern, the Simple Factory.
In this example, we’re going to build various types of weapons using a factory struct.

First, we create the Pistoler interface, which defines all methods a Pistol should have. 
Second we create the Pistol struct, which implements the Pistoler interface.
Third we create 2 concrete types of Pistols, the Beretta and the Colt, both of which embed the Pistol struct and indirectly implement all Pistol method.

Finally, we create getGun() function, which acts as a factory for creating different types of Pistols by taking a string parameter 
and returning a Pistoler interface.


