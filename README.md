# assignment-golang-backend


### Done Requirements


- Authentication and Authorization
  - Register
  - Login

- List of transactions
  - default 10, with query params combined

- Topup
  - Credit to Wallet
  - Amount validation
  - Description formatting
  - Record to transactions

- Transfer
  - Negative case
  - Body validation
  - Record to transactions  

- Users detail With Auth Middleware

- Unit Test
  - Usecase
    - Register and login 
    - Get Transactions
    - Topup & Transfer (95.7% on transfer)
  - Handler
    - No Route
    - user handler (login, get detail, register)
    - wallet handler (topup, transfer)


### Not Done
- Pagination optionals
- ERD
- Seeding
- Swagger Documentation