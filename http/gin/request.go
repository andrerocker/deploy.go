package gin

import (
	"github.com/gin-gonic/gin"
	"io"
)

type GinRequest struct {
	context *gin.Context
}

func NewRequest(context *gin.Context) GinRequest {
	return GinRequest{context}
}

func (self GinRequest) Reader() io.Reader {
	return self.context.Request.Body
}

func (self GinRequest) Writer() io.Writer {
	return Flushed(self.context.Writer)
}

func (self GinRequest) Parameter(name string) string {
	return self.context.Params.ByName(name)
}

func (self GinRequest) Abort(code int) {
	self.context.String(code, "unauthorized")
	self.context.Abort()
}
