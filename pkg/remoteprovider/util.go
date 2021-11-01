package remoteprovider

import (
	"github.com/gin-gonic/gin"
	"github.com/iver-wharf/wharf-core/pkg/ginutil"
)

func SetupClientFromContext(client Client, c *gin.Context) bool {
	remoteProviderURL, ok := ginutil.RequireQueryString(c, "remoteProviderUrl")
	if !ok {
		return false
	}
	_, token, _ := c.Request.BasicAuth()
	setupClient(client, c, remoteProviderURL, token)
	return true
}

func setupClient(client Client, c *gin.Context, remoteProviderURL, token string) {
	client.SetContext(c)
	client.SetToken(token)
	client.SetRemoteProviderURL(remoteProviderURL)
}
