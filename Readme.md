# Mealwhile - Golang Implementation

## About Mealwhile
Mealwhile is an application for planning and executing the event kitchen for different occasions, e. g. christian summer camps. It supports you when planning and making the meals for the event.

For more information check the <a href="https://github.com/JohannesBreitling/MealwhileDocumentation">documentation repository</a>.

## Technical Information

### OAPI Codegeneration
The ```server.gen.go``` and the ```server-model.gen.go``` are generated automatically from the api specification by the <a href="https://github.com/deepmap/oapi-codegen">oapi-codegen package</a>.
To regenerate the files use the following commands:
``````
oapi-codegen -package controller -generate types ../spec/mealwhile-api.yaml > server-model.gen.go
``````
``````
oapi-codegen -package controller -generate server ../spec/mealwhile-api.yaml > server.gen.go
``````

### Build the image
#### Local
To build the image for the local platform (arm64) execute the following command in the parten folder:
``````
docker build --target local -t mealwhile:local .
``````
#### Server
To build the image for the server platform (linux/amd64) execute the following command in the parten folder:
``````
docker build --platform linux/amd64 --target server -t johannesbreitling/server:mealwhile .
``````