package notify

import (
	"github.com/appleboy/go-fcm"
	"github.com/appleboy/go-hms-push/push/core"
	"github.com/sideshow/apns2"
)

type AppClients struct {
	ID   string
	FCM  *fcm.Client
	APNS *apns2.Client
	HMS  *core.HMSClient
}
