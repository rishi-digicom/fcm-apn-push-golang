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
}​
3
```go
4
func readPrivateKey(filepath string) (*ecdsa.PrivateKey, error) {
5
        file, err := ioutil.ReadFile(filepath)
6
        if err != nil {
7
                return nil, err
8
        }
9
        block, _ := pem.Decode(file)
10
​
11
        key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
12
        if err != nil {
13
                return nil, err
14
        }
15
        if privateKey, ok := key.(*ecdsa.PrivateKey); ok {
16
                return privateKey, nil
17
        }
18
        return nil, errors.New("")
19
}
20
```
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

