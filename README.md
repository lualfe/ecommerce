# Ecommerce

## Pre-Requisites
- [Docker Desktop](https://hub.docker.com/?overlay=onboarding) 
- [Golang 1.13.6](https://golang.org/)
- [Visual Studio Code](https://code.visualstudio.com/)
- [Git](https://git-scm.com/downloads)
- [Insomnia](https://insomnia.rest/)

## Configure local environment (only for the fist time)
1. Create Containers
- Enter in project root and run 
````
. scripts/start.sh
````
  
2. Config environment
- Enter in project root and run
````
. scripts/config.sh
````

3. Enter in project root and run application and the migrations
````
go run server.go
````

4. Import the `insomnia_ecommerce.json` file in your insomnia

## note

- If creating containers failed you can run the ` clear.sh` script to remove and start again