package push

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fcm-apn-push-golang/service"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"
)

/* request data struct */

type AndroidCredential struct {
	Server_key   string
	Device_token string
	Title        string
	Subtitle     string
	Message      string
	Sound        string
}

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
	Badge            int
}

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

/* fcm struct */

type FcmData struct {
	Notification Notification `json:"notification"`
	Data         Data         `json:"data"`
	To           string       `json:"to"`
}

type Notification struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

type Data struct {
	Message string `json:"message"`
}

/* apn struct */

type ApnData struct {
	Aps Aps `json:"aps"`
}

type Aps struct {
	Alert Alert  `json:"alert"`
	Badge int    `json:"badge"`
	Sound string `json:"sound"`
}

type Alert struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
}

/* response data struct */

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (reqData AndroidCredential) FCMPushNotification() (*Response, error) {
	response := Response{}
	// set payload data

	notification := Notification{reqData.Message, reqData.Title}
	data := Data{reqData.Message}
	payload := FcmData{notification, data, reqData.Device_token}
	fj, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := "https://fcm.googleapis.com/fcm/send"

	req, err := http.NewRequest("POST", url, bytes.NewReader(fj))
	if err != nil {
		return nil, err
	}

	// set header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "key="+reqData.Server_key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, _ := ioutil.ReadAll(res.Body)
	result := string(resBody)

	defer req.Body.Close()
	response.Code = res.StatusCode
	response.Message = result

	return &response, nil
}

func (reqData IOSCredential) APNPushNotification() (*Response, error) {
	response := Response{}
	// generate JWT token

	token, err := service.GenerateJwtToken(reqData.Key_id, reqData.Issuer_claim_key, reqData.P8_file_path)
	if err != nil {
		return nil, err
	}

	Authorization := "bearer " + token

	// set payload data
	url := ""
	if reqData.Send_box == false {
		url = "https://api.development.push.apple.com:443/3/device/" + reqData.Device_token
	} else {
		url = "https://api.push.apple.com:443/3/device/"
	}

	alert := Alert{reqData.Title, reqData.Subtitle, reqData.Message}
	aps := Aps{alert, reqData.Badge, reqData.Sound}
	payload := ApnData{aps}

	aj, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(aj))
	if err != nil {
		return nil, err
	}

	// set header

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("authorization", Authorization)
	req.Header.Add("apns-topic", reqData.Apn_topics)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)
	result := string(resBody)

	response.Code = res.StatusCode
	response.Message = result
	return &response, nil
}

func (reqData IOSCredentialWithCert) APNPushNotificationPEM() (*Response, error) {
	response := Response{}

	url := ""
	if reqData.Send_box == false {
		url = "https://api.development.push.apple.com:443/3/device/" + reqData.Device_token
	} else {
		url = "https://api.push.apple.com:443/3/device/"
	}

	// set payload data
	alert := Alert{reqData.Title, reqData.Subtitle, reqData.Message}
	aps := Aps{alert, reqData.Badge, reqData.Sound}
	payload := ApnData{aps}

	aj, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(aj))
	if err != nil {
		return nil, err
	}

	// set header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("apns-topic", reqData.Apn_topics)

	var client *http.Client
	var transport *http.Transport
	certificate, err := tls.LoadX509KeyPair(reqData.PEM_file_path, reqData.PEM_file_path)
	if err != nil {
		return nil, err
	}

	configuration := &tls.Config{
		Certificates: []tls.Certificate{certificate},
	}

	configuration.BuildNameToCertificate()
	transport = &http.Transport{TLSClientConfig: configuration}
	client = &http.Client{Transport: transport}

	err = http2.ConfigureTransport(transport)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, _ := ioutil.ReadAll(res.Body)
	result := string(resBody)

	defer req.Body.Close()

	response.Code = res.StatusCode
	response.Message = result
	return &response, nil
}
