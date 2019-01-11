# fcm-and-apn-push-notification

<pre>
POST http://localhost:4002/sendPushNotification
{
	"device_type":"ios", // required | "ios" or "android"
	"device_token":"YOUR DEVICE TOKEN", // required
	"title":"test notification", // required
	"sub_title":"this is sub title", // optional
	"message":"this is test message", // optional
	"sound":"default" // optional
}
</pre>

## Required Dependencies
- github.com/julienschmidt/httprouter
- crypto/ecdsa
- crypto/x509
- encoding/pem
- github.com/dgrijalva/jwt-go
- encoding/json

## install go build
```go
go install fcm-and-apn-push-notification
```

## go run
```go
go run main.go -http=":4001"
```
