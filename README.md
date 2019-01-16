# Golang push notification with FCM and APN.

How to use:

For push notification for iOS use IOSCredential struct 
```go
type IOSCredential struct {
	Send_box         bool
	Key_id           string
	P8_file_path     string
	Issuer_claim_key string
	Apn_topics       string
	Device_token     string
	Title            string
	Subtitle         string
	Message          string
	Sound            string
}
```
Call APNPushNotification method and pass reference of IOSCredential struct
```go
func (reqData IOSCredential) APNPushNotification() (*Response, error) {
```

For push notification for android use AndroidCredential struct 
```go
type AndroidCredential struct {
	Server_key   string
	Device_token string
	Title        string
	Subtitle     string
	Message      string
	Sound        string
}
```

Call FCMPushNotification method and pass reference of AndroidCredential struct
```go
func (reqData AndroidCredential) FCMPushNotification() (*Response, error) {
```
