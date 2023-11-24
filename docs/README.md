# Problem

- If you have a F&B business, you need the way to manage ordering, request, provider and inventory. It will take very long time and manual work to do those kind of stuff. It will be easier to do it systematically
- Keep track everything will be easier by using system instead of excel file, it will also help to look back with statistic.

# Business Context

## Business structure

- a F&B company which have series of coffee and bar around Viet Nam. Currently, they are having stores in some biggest city in Viet Nam. Each store will have coffee and possible bar at night.
- each service(coffee or bar) at one store will have one store manager
- store manager manage stuff in the store and request to buy items if needed
- purchasing manager is responsible for receiving order request, finding suitable provider for each item, contact and work with provider to buy and deliver to store.

## Users

- Store manager:
    - request items what they want to buy
    - keep track status of request
- **Purchasing manager:** main role in the system
    - receive ordering request
    - find suitable provider and order all the items
    - export request to receive approval from financial accounting and manager
    - manage provider information to order:
        - phone number
        - zalo
        - address,…
    - update status of ordering request
    - export report by filtering for manager
        - time range, store, items
    - grant access to new users
- inventory manager:
    - export information
- manager
    - grant access to new users

# Solution

- having multiple users around Viet Nam, so we need to allow anyone with network can access
- nowadays, using web browser is the most effective way to access data through network
- so building **web application** will be the goal.

# Expectation

- keep track information
- statistic with filtering, searching