package common

import (
	"encoding/json"
	"fmt"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
)

var Client *apns2.Client

func PushNotifcation(deviceToken string, payload map[string]interface{}) {
	res, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}
	notification := &apns2.Notification{
		DeviceToken: deviceToken,
		Topic:       ENV.APPID,
		Payload:     res,
	}

	response, err := Client.Push(notification)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

func init() {
	// put p12 in project root path
	cert, err := certificate.FromP12File("./cert.p12", "123")

	if err != nil {
		fmt.Println(err)
	}
	// Development Production
	Client = apns2.NewClient(cert).Development()
}
