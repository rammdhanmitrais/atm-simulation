# atm-simulation
A program which simulate the functionalities of ATM
- This program contains two records with the following details:
  --
   - Name          : John Doe
     
     PIN           : 012108

     Balance       : $100

     Account Number: 112233

   - Name          : Jane Doe

     PIN           : 932012

     Balance       : $30

     Account Number: 112244

- This program includes:
  --
   - Login using Account Number and PIN
   - Withdraw the Balance
   - View Account Balance
   - Transfer to Other Account Number

===========

# ATM System CLI
ATM System CLI is a tool for simulating an interaction of an ATM with a retail bank from the terminal.

## Pre-requisite
- installed <a href="https://nodejs.org/en/download/package-manager" target="_blank">node js</a>
- installed <a href="https://docs.npmjs.com/downloading-and-installing-node-js-and-npm" target="_blank">npm</a>

## Installation

```bash
npm install
```

## Usage
To start using ATM System CLI, run:

```bash
npm start
```

### Commands
- `login [name]`
example: `login Alice` -> will log in as Alice, or create user with name Alice if not exist
rules: 
    - need a parameter as user name to find or create user
    - if there are other parameters, it will not be processed (only name)
    - if there is another parameter `login John Doe`, the App will create user with name John
    - this command cannot be worked when there is a user still logged in
    
- `logout`: 
example: `logout` -> will log out of the current user
rules:
    - no need parameter
    - this command will be worked whether there is a user or not

- `deposit [amount]`: 
example: "deposit 400" -> will deposit 400 to the current user
rules:
    - need a parameter as amount that will be deposited to user account
    - if there are other parameters, it will not be processed (only amount)
    - the amount should be an integer and in the range 1 to 1500000 (default)
    - if the user has owe to other user, will automatically transfer the amount to the user it owes money to

- `withdraw [amount]`: 
example: "withdraw 100" -> will withdraw 100 from the current user
rules:
    - need a parameter as amount that will be withdraw from user balance
    - if there are other parameters, it will not be processed (only amount)
    - the amount should be an integer and in the range 1 to 1500000 (default)
    - this command will be worked, when the balance is sufficient

- `transfer [target name] [amount]`: 
example: "transfer Alice 50" -> will transfers 50 from the current user to the target user (Alice)
rules:
    - need parameters as user target name and amount
    - if there are other parameters, it will not be processed (only target name and amount)
    - the current user will transfer the amount to the target user name
    - the amount should be an integer, greater than 0 or in the range 1 to 1500000 (default)
    - if the balance is unsufficient, the deficiency will be considered an owe and charged to the current user

- `help`: 
example: "help" -> will show the list of available command

You can read the available list commands with this: `help`.

## Run testing
To check the unit test, run:

```bash
npm test
```
