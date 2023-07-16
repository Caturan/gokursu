package restful

/*
	A RESTful API (Representational State Transfer API) is a type of web service that follows the principle of REST architecture.
	It is designed to provide a standardized way for clients (such as web or mobile applications) to interact with server-side resource over the HTTP protocol.

	The key characteristics of a RESTful API are:

	1. Resources : Resources are the central abstractions in a RESTful API. Each resource represent an object or entity (eg., user,product,post)
	that the API can perform operations on. Resources are identified by unique URLs (Uniform Resource Identifiers)

	2. HTTP Methods : RESTful APIs use standard HTTP methods to perform CRUD (Create, Read, Update, Delete) operations on resources. The common HTTP methods used in
	RESTful APIs are:
	* GET : Retrieve data from the server (read)
	* POST : Create a new resource on the server.
	* PUT : Update an existing resource on the server.
	* DELETE : Remove a resource from the server.

	3. Statelessness : RESTful APIs are stateless, meaning that each request from the client to the server must contain all the information needed to understand and
	process the request. The server doesn't store any client session information between request.

	4. Representation : Resources in a RESTful API are represented in a specific format, often JSON or XML. The client can request data in a particular representation,
	and the server will respond with request format.

	5. Uniform Interface : RESTful APIs follow a uniform and consistent interface, which simplifies commnunication between clients and servers.
	This is achieved by adhering to standad conventions for resource naming, HTTP methods and response formats.

	6. Hypermedia (HATEOAS - Hypermedia as the Engine of Application State): This is an optiona constraint in RESTful APIs.
	It means that API responses should include hyperlinks that the client can follow to discover and interact with other resources.
	HATEOAS enablesa more dynmaic and self-descriptive API.

	RESTful APIs are widely used for buildings web services that allow different software systems to communicate with each other over the internet.
	They provide a scalable, stateless, and easy-to-understand approach to data exchange, making them popular for building modern web and mobile applications.

*/
