# Mealwhile - Golang Implementation

## About Mealwhile
Mealwhile is an application for planning and executing the event kitchen for different occasions, e. g. christian summer camps. It supports you when planning and making the meals for the event.

For more information check the <a href="https://github.com/JohannesBreitling/MealwhileDocumentation">documentation repository</a>.

## Technical Information

### OAPI Codegeneration
The ```server.gen.go``` and the ```server-model.gen.go``` are generated automatically from the api specification by the <a href="https://github.com/deepmap/oapi-codegen">oapi-codegen package</a>.
To regenerate the files use the following commands:
``````
oapi-codegen -package controller -generate types <relative-path-to-api-specification> > server-model.gen.go
``````
``````
oapi-codegen -package controller -generate server <relative-path-to-api-specification> > server-model.gen.go
``````