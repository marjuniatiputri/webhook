package gcf

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/marjuniatiputri/webhook"
)

func init() {
	functions.HTTP("WebHook", webhook.PostBalasan)
}
