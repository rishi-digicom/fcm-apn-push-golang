# Golang push notification with FCM and APN.

How to use:

For push notification for iOS with support of p8 file, use IOSCredential struct 
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
	Badge 		 int
}
```
Call APNPushNotification method and pass reference of IOSCredential struct
```go
func (reqData IOSCredential) APNPushNotification() (*Response, error) {
```
For push notification for iOS with support of PEM file, use IOSCredentialWithCert struct 
```go
type IOSCredentialWithCert struct {
	Send_box      bool
	PEM_file_path string
	Apn_topics    string
	Badge         int
	Device_token  string
	Title         string
	Subtitle      string
	Message       string
	Sound         string
}
```
Call APNPushNotificationPEM method and pass reference of IOSCredentialWithCert struct
```go
func (reqData IOSCredentialWithCert) APNPushNotificationPEM() (*Response, error) {
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
