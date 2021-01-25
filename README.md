# GOLANG UPDATE POINT
## Install dependencies
```
$ go mod download 
```


## Develop environment at localhost:7000
```
$ CompileDaemon -build="go build simpleGo" -command="./simpleGo" 
```

# API

## Update Point

` https://${url}/updatePoint`

> ### Request Body

```
  {
    "memberId": "",
    "productCode": "",
    "terminalId": "",
    "volumn": "",
    "price": ""
  }
```
<br />

> ### Method

    POST

> ### Success Response
<br />

 * **Code:** 200 OK <br />
  * **Content:** 
```
  {
    "memberId": "",
    "productName": "",
    "receivePoint": number 
  }
```
<br />

> ### Error Response
  * **Code:** 404 NOT FOUND 



