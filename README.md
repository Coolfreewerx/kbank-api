## How to Run
1. Run Docker

```
docker-compose up -d
```

2. Run program

```
go run main.go
```
or 
```
air
```

## ER Diagram
Database ER diagram for this project
- [ER_Diagram](https://dbdocs.io/nattapon.mee/Bill-Payment?view=relationships)

## API Service
### Routes
- POST `http://localhost:8080/api/billpayment/lookup`
- POST `http://localhost:8080/api/billpayment/payment`
- GET  `localhost:8080/api/billers`
- POST `localhost:8080/api/billers`
- GET `localhost:8080/api/customers`
- GET `localhost:8080/api/customers/:id`
- POST `localhost:8080/api/customers`
- GET `localhost:8080/api/bills`
- POST `localhost:8080/api/bills`
- PUT `localhost:8080/api/bills/:id`

### KBank API (bill-payment-service)
#### Inquiry Lookup
- METHOD POST `http://localhost:8080/api/billpayment/lookup`
- BODY
```json
{
  "functionName": "BillLookup",
  "transactionId": "98099310720200530183382100",
  "transactionDateTime": "2023-08-15T21:39:46+07:00",
  "billerType": "",
  "billerId": "98499",
  "terminalNo": "xxxxxxx8888",
  "channelCode": "MOB",
  "tranAmount": "120.00",
  "senderBankCode": "Kbank",
  "reference1": "300000025751",
  "reference2": "1733084",
  "language": "EN",
  "apiKey": "SecureKey123"
}

```

#### Response
- status : `200 OK`
- BODY
```json
{
    "functionName": "BillLookupResponse",
    "transactionId": "98099310720200530183382100",
    "transactionDateTime": "2023-08-17T21:17:04+07:00",
    "billerTransactionId": "98099310720200530183382100",
    "responseCode": "0000",
    "responseDescription": "Success",
    "billerType": "",
    "billerId": "98499",
    "terminalNo": "xxxxxxx8888",
    "promptPayTransactionId": "",
    "typeofReceiver": "",
    "reference1": "300000025751",
    "reference2": "1733084",
    "reference3": "",
    "tranAmount": "",
    "info1": "",
    "info2": "",
    "info3": "",
    "additional": {
        "toBillerAccountName": "",
        "toBillerServiceName": "",
        "payerFee": "",
        "settlementDate": "",
        "receiverTaxID": "",
        "dueDate": "",
        "rtpReference": "",
        "rsAppId": ""
    }
}
```

#### Payment Confirm
- METHOD POST `http://localhost:8080/api/billpayment/payment`
- BODY
```json
{
  "functionName": "BillPayment",
  "transactionId": "98099310720200530183382100",
  "transactionDateTime": "2023-08-06T10:33:18.450+07:00",
  "billerType": "",
  "billerId": "98499",
  "terminalNo": "xxxxxxx8888",
  "channelCode": "MOB",
  "tranAmount": "120.00",
  "senderBankCode": "Kbank",
  "isRetry": "0",
  "reference1": "300000025751",
  "reference2": "1733084",
  "language": "EN",
  "apiKey": "SecureKey123"
}

```
#### Response
- status : `200 OK`
- BODY
```json
{
    "functionName": "BillPaymentResponse",
    "transactionId": "98099310720200530183382100",
    "responseDateTime": "2023-08-17T21:09:08+07:00",
    "billerTransactionId": "98099310720200530183382100",
    "responseCode": "0000",
    "responseDescription": "Success",
    "terminalNo": "xxxxxxx8888",
    "promptPayTransactionId": "",
    "additional": {
        "settlementDate": "",
        "rsAppId": ""
    }
}
```
*** After you payment confirm that bill will update status to already_paid
### Backend API (Database)
#### Biller_account
- METHOD GET `localhost:8080/api/billers`
##### Response
```json
[
    {
        "id": 98499,
        "name": "ThaiTee",
        "service_name": "Abank"
    }
]
```
- METHOD POST `localhost:8080/api/billers`
- BODY
```json
{
    "name" : "Yum",
    "service_name" : "SBB"
}
```
##### Response
```json
{
    "id": 2,
    "name": "Yum",
    "service_name": "SBB"
}
```

#### Customer
- METHOD GET `localhost:8080/api/customers`
##### Response
```json
[
    {
        "id": 300000025751,
        "firstName": "Teerapat",
        "lastName": "Ku",
        "titleName": "",
        "mobileNumber": "0888888888",
        "createdAt": "2023-08-16T15:06:05.652791Z"
    },
    {
        "id": 1,
        "firstName": "Testers",
        "lastName": "Tu",
        "titleName": "",
        "mobileNumber": "0999999999",
        "createdAt": "2023-08-16T22:23:28.643741Z"
    }
]
```
- METHOD GET `localhost:8080/api/customers/:id`
##### Response
```json
{
    "id": 1,
    "firstName": "Testers",
    "lastName": "Tu",
    "titleName": "",
    "mobileNumber": "0999999999",
    "createdAt": "2023-08-16T22:23:28.643741Z"
}
```
- METHOD POST `localhost:8080/api/customers`
- BODY
```json
{
    "firstName": "Teera",
    "lastName" : "Pl",
    "titleName" : "",
    "mobileNumber" : "0999999999"
}
```
##### Response
```json
{
    "id": 2,
    "firstName": "Teera",
    "lastName": "Pl",
    "titleName": "",
    "mobileNumber": "0999999999",
    "createdAt": "2023-08-17T21:53:56+07:00"
}
```

#### Bill
- METHOD GET `localhost:8080/api/bills`
##### Response
```json
[
    {
        "id": 1,
        "billerId": 98499,
        "reference1": 300000025751,
        "reference2": 1733084,
        "trasactionId": "98099310720200530183382100",
        "channelCode": "MOB",
        "senderBankCode": "Kbank",
        "status": "already_paid",
        "tranAmount": 120,
        "createdAt": "2023-08-16T15:06:05Z",
        "updatedAt": "2023-08-17T15:39:57Z"
    },
    {
        "id": 2,
        "billerId": 98499,
        "reference1": 1,
        "reference2": 17003,
        "trasactionId": "98099310720200530183382100",
        "channelCode": "MOB",
        "senderBankCode": "Kbank",
        "status": "already_paid",
        "tranAmount": 210,
        "createdAt": "2023-08-17T13:12:19Z",
        "updatedAt": "2023-08-17T13:15:19Z"
    }
]
```
- METHOD POST `localhost:8080/api/bills`
- BODY
```json
{
    "billerId": 98499,
    "reference1": 1,
    "reference2": 170049,
    "status": "waiting",
    "tranAmount": 150
}
```
##### Response
```json
{
    "id": 1,
    "billerId": 98499,
    "reference1": 1,
    "reference2": 170049,
    "trasactionId": "",
    "channelCode": "",
    "senderBankCode": "",
    "status": "waiting",
    "tranAmount": 150,
    "createdAt": "2023-08-18T16:35:04+07:00",
    "updatedAt": "2023-08-18T16:35:04+07:00"
}
```
- METHOD PUT `localhost:8080/api/bills/:id`
- BODY
```json
{
    "billerId": 98499,
    "reference1": 1,
    "reference2": 170049,
    "status": "waiting",
    "tranAmount": 220
}
```
##### Response
```json
{
    "id": 1,
    "billerId": 98499,
    "reference1": 1,
    "reference2": 170049,
    "trasactionId": "",
    "channelCode": "",
    "senderBankCode": "",
    "status": "waiting",
    "tranAmount": 220,
    "createdAt": "2023-08-18T16:35:04+07:00",
    "updatedAt": "2023-08-18T16:36:23+07:00"
}
```