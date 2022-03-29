## Delivery Apps

## Getting Started

### Prerequisite

This project required Golang programming language to run:
- [Golang](https://go.dev/doc/install)

You can use any IDE such as VSCode and GoLand. But it is recommended if you use IDE with debugger attached. If you are using VSCode you can follow [here](https://code.visualstudio.com/docs/languages/go) to set up GO programming language including IntelliSense, Navigation, etc.

It is recommended to use postgres as DBMS to manage the data, you can follow this [link](https://www.postgresql.org/download/) to download.

### Installation

To run this project locally follow these steps:
1. Clone the repository
    `git clone https://github.com/hafizulrifkihawari/delivery.git`
2. Copy .env.example to .env file and fill out the variables needed.
3. After all variables needed filled, you are ready to use including develop and debug the code. If you are using VS Code go to main.go and simply hit F5 and debugging mode should started.

This project also can be run using docker, you can run the project by using these following commands:
1. Build the image
    `docker build -t delivery .`
2. Run container on port, in this example we are using port 9000
    `docker run -p 9000:9000 delivery`

### Documentation

For a complete documentation you can import [this](https://www.getpostman.com/collections/352cac989e21937713cb) to your postman application.



### Deployment

This API has been deployed on [Heroku](https://secret-ocean-69132.herokuapp.com/), or you can copy link below
- `https://secret-ocean-69132.herokuapp.com/`
