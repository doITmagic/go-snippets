Publish-Subscribe is a messaging pattern used to communicate messages between different components without 
these components knowing anything about each other's identity.
It is similar to the Observer behavioral design pattern. The fundamental design principals of both Observer and Publish-Subscribe is the decoupling of those interested in being informed about Event Messages from the informer (Observers or Publishers). 
Meaning that you don't have to program the messages to be sent directly to specific receivers.
To accomplish this, an intermediary, called a "message broker" or "event bus", receives published messages, and then routes them on to subscribers.
There are three components messages, topics, users.