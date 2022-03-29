## Delivery Apps

---

## Getting Started

### Prerequisite

This project required Golang programming language to run:
- [Golang](https://go.dev/doc/install)

You can use any IDE such as VSCode and GoLand. But it is recommended if you use IDE with debugger attached. If you are using VSCode you can follow [here](https://code.visualstudio.com/docs/languages/go) to set up GO programming language including IntelliSense, Navigation, etc.

It is recommended to use postgres as DBMS to manage the data, you can follow this [link](https://www.postgresql.org/download/) to download.

---

### Installation

To run this project locally follow these steps:
1. Clone the repository
    ```git clone https://github.com/hafizulrifkihawari/delivery.git```
2. Copy .env.example to .env file and fill out the variables needed.
3. After all variables needed filled, you are ready to use including develop and debug the code. If you are using VS Code go to main.go and simply hit F5 and debugging mode should started.

This project also can be run using docker, you can run the project by using these following commands:
1. Build the image
    ```docker build -t delivery .```
2. Run container on port, in this example we are using port 9000
    ```docker run -p 9000:9000 delivery```

Notes:
This repo include auto migration for the table schemes and seeding the extracted and transformed data into the corresponding tables. So there is no extra commands needed to run **ETL**. 

---

### Deployment
This API has been deployed on [Heroku](https://secret-ocean-69132.herokuapp.com/)

---

### Documentation
For a complete documentation you can import [this](https://www.getpostman.com/collections/352cac989e21937713cb) to your postman application.

### BASE URL
```
https://secret-ocean-69132.herokuapp.com/delivery
```

---

### Health Check
Check for service availability

**Endpoint**
```
GET /health-check
```

**Response**
```
{
    "status": 200,
    "message": "OK"
}
```

---

### List restaurant by datetime
Filter restaurants by specified datetime given on request parameter

**Endpoint**
```
GET /restaurant/:search-type?datetime=1648566026
```

**Path Variable**
- :search-type = `"date"` -> this search type to identify the endpoint to filter by date

**Request Query Parameter**
- datetime -> epoch datetime

**Response Body**
```
{
    "status": 200,
    "message": "OK",
    "data": [
        {
            "id": 5,
            "name": "12 Baltimore",
            "opening_hour": {
                "day": "Tues",
                "open_at": "13:30:00",
                "close_at": "15:45:00"
            }
        },
        ...
    ]
}
```

---

### List pagination restaurant by dish and price range
- Filter restaurants by given query parameter 

**Endpoint**
```
GET /restaurant/:search_type?price_start=1&price_end=15&limit=100&num_dishes_gt=10
```

**Path Variable**
- :search_type -> `dish` -> this search type to identify the endpoint to filter by dish

**Request Query Parameter**
- price_start -> filter by range start price
- price_end -> filter by range start end
- limit -> amount of item provided
- num_dishes_gt -> filter by number of dishes greater than x value
- num_dishes_lt -> filter by number of dishes less than x value

**Response**
```
{
    "status": 200,
    "message": "OK",
    "data": [
        {
            "id": 2,
            "name": "024 Grille"
        },
        {
            "id": 7,
            "name": "13 Coins"
        },
        ...
   ]
}
```

---


### Search Restaurants
- Search restaurants name and/or dishes

**Endpoint**
```
GET /restaurant/:search_type?term=French
```

**Path Variable**
- :search_type -> `search` -> this search type to identify the endpoint to search

**Request Query Parameter**
- term -> search value text

**Response**
```
{
    "status": 200,
    "message": "OK",
    "data": [
        {
            "id": 5746,
            "restaurant_name": "Gerard's Restaurant Maui",
            "dish_name": "French Vanilla Ice Cream with Crushed Fresh Pineapple"
        },
        {
            "id": 6218,
            "restaurant_name": "Hama Sushi",
            "dish_name": "French vanilla ice cream"
        },
        ...
    ]
}
```

---

### Purchase Dish

**Endpoint**
```
POST /purchase
```

**Request Body**

```
{
    "user_id": 1,
    "menu_id": 40
}
```

**Response**
```
{
    "status": 200,
    "message": "OK",
    "data": {
        "name": "Edward Gonzalez",
        "restaurant_name": "12 Baltimore",
        "dish_name": "Ugnstekt Apple",
        "transaction_amount": 10.43,
        "cash_balance": 227.18
    }
}
```
