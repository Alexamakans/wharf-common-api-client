package remoteprovider

import (
	"github.com/gin-gonic/gin"
)

func SetupClientFromContext(client Client, remoteProviderURL string, c *gin.Context) bool {
	_, token, _ := c.Request.BasicAuth()
	setupClient(client, c, remoteProviderURL, token)
	return true
}

func setupClient(client Client, c *gin.Context, remoteProviderURL, token string) {
	client.SetContext(c)
	client.SetToken(token)
	client.SetRemoteProviderURL(remoteProviderURL)
}
