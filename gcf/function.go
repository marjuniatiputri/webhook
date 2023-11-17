package gcf

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/jxxzy/webhook"
)

func init() {
	functions.HTTP("WebHook", webhook.PostBalasan)
}
