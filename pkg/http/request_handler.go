package http

import "github.com/gin-gonic/gin"

func ReadRequest(ctx *gin.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request.Context(), request)
}

// func GetRequestID(c *gin.Context) string {

// }

// func GetRequestToken(c *gin.Context) string {

// }
