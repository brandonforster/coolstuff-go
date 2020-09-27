# coolstuff

## How to use
The easiest way to run the application is to use the Makefile.

To verify functionality, run `make test`. This will run the unit tests locally.

To run the application locally, run `make run`. This will start the server locally on your machine.

To run the application in Docker, run `make docker_build` and `make docker_run`. This will start the server in a
 Docker container. To remove the container and the image, run `make docker_clean`.
 
 To run end to end tests, launch the app either locally or in Docker and run Postman. Make the requests in the
  Postman collection. Each of the requests contain tests that verify that the request executed successfully.
