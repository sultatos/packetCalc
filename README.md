This is a simple package calculator demo that calculates the minimum packages required to ship the order. The BE is written in go and exposes 2 REST endpoints 
* POST /packet :which accepts a form with a single parameter items as input and returns a json with the packets that the shipment will be send
* POST /packet-sizes : which accepts a json with a single parameter packetSizes as input that is an array of integers of the packages

* There is also a UI developed with htmx and templ just to showcase the functionality of the BE, nothing fancy.
## Getting Started
Clone the repository and run the ```make run``` command to build the application. The application will be available at http://localhost:8080
There is also a dockerfile that can be used to build the application in a container.

```bash
## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
