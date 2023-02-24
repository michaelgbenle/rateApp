# Rates App
This a simple Backend application that allow users to exchange USD for Naira and vice versa at the current market rates.

#### Rate EndPoint : https://staging-biz.coinprofile.co/v3/currency/rate

### Explicit Tasks/Features
- [x] Create an authentication endpoint that allows users to sign up and log in to their account using email and password. No email validation is required
- [x] Each user should have a balance of 100 USD in their account after signing up.
- [x] Users should be able to convert between USD and NGN using the USDCNGN and USDCNGN_ conversion rate in the above rate endpoint. Use the rates provided so that your        application makes a profit on each conversion between the currencies
- [x]  Users should be able to get their balances and transaction history


### Endpoints
- [x] /ping - This endpoint is used to check if the application is running [GET]
- [x] /register - This endpoint is used to create a new user account [POST]
   ```json
  Sample Request Body
   {
       "email": "test@example.com",
        "password": "passworD5$"
  }
- [x] /login - This endpoint is used to log in a user [POST]
   ```json
    Sample Request Body
     {
         "email": "test@example.com",
          "password": "passworD5$"
    }

- [x] /user/ngnusd - This endpoint is used to convert USD to Naira [PATCH]
    ```json
  This route is a protected route, you need to be logged in to access it
   `
    Sample Request Body
     {
         "currency": "USD",
          "amount": "50"
    }

- [x] /user/usdngn - This endpoint is used to convert Naira to USD [PATCH]
  ```json
    This route is a protected route, you need to be logged in to access it
    Sample Request Body
     {
         "currency": "NGN",
          "amount": "50000"
    }

- [x] /user/transactions_history - This endpoint is used to get the user balances and transaction history [GET] .
it is a protected route, you need to be logged in to access it

### Environment Variables
1. `PORT` - This is the port the application will run on
2. `MONGO_URL` - This is the url of the mongo database
3. `BASE_URL` - This is the rate endpoint url above
4. `JWT_SECRET` - This is the secret used to generate jwt token

### How to run and test the application
1. Clone the repository
2. Run `go mod download` to download all the dependencies
3. Set the environment variables
4. Run `make run` to start the application
5. Run `make tests` to run the test cases


### Application Requirement Stack Used
1. Golang as the main programming language
2. Mongo database as the core database
3. Gin-Gonic, main golang web framework used

