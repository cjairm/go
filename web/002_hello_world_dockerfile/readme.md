#  Web Test - Description

First testing for web development w/docker file

# How to use it

* Download main.go and Dockerfile
* Save it in some folder
* Execute `docker build -t [docker-image-name] .`. Now that we have defined everything we need for our Go application to run in our Dockerfile we can now build an image using this file.
* Execute `docker run -p 8080:8081 -it [docker-image-name]`. In order to run this newly created image, we can use the docker run command and pass in the ports we want to map to and the image we wish to run.

## Author(s)

* Carlos Mendez

### Created

Jun 14, 2020