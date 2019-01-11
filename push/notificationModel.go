package push

import (
	"bytes"
	"encoding/json"
	"fcm-apn-push-golang/service"
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
	Alert Alert `json:"alert"`
}

type Alert struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Body     string `json:"body"`
	Sound    string `json:"sound"`
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

	alert := Alert{reqData.Title, reqData.Subtitle, reqData.Message, reqData.Sound}
	aps := Aps{alert}
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

	response.Code = res.StatusCode
	response.Message = res.Status
	return &response, nil
}
