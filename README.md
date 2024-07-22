# atm-simulation
A program which simulate the functionalities of ATM
- There is two datasource on this program
  --
    - Default datasource contains two records with the following details:
    --
      - Name          : John Doe
     
        PIN           : 012108

        Balance       : $100

        Account Number: 112233

      - Name          : Jane Doe

        PIN           : 932012

        Balance       : $30

        Account Number: 112244

    - CSV datasource contains records with format -> Name;PIN;Balance;Account Number:
    --
      - Alice Johnson;230501;1247;135678
      - Name          : Alice Johnson
     
        PIN           : 230501

        Balance       : $1247

        Account Number: 135678

- This program includes:
  --
   - Login using Account Number and PIN
   - Withdraw the Balance
   - View Account Balance
   - Transfer to Other Account Number
   - Transactions History
