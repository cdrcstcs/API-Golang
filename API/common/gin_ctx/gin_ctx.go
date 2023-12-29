package gin_ctx

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// GetCtxAndReqFromGinCtx T is value not pointer
func GetCtxAndReqFromGinCtx[T any](c *gin.Context, req T) (context.Context, T, error) {
	req, err := decodeReq[T](c, req)
	if err != nil {
		return nil, req, errors.Wrap(err, "decode req fails")
	}
	ctx := decodeCtx(c)

	return ctx, req, nil
}

func decodeCtx(c *gin.Context) context.Context {
	ctx := context.Background()
	if c.Keys == nil {
		return ctx
	}
	for k, v := range c.Keys {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}

func decodeReq[T any](c *gin.Context, req T) (T, error) {
	if err := c.ShouldBind(&req); err != nil { // should bind already handle c.Request.URL for GET request. POST should put this in body
		return req, errors.Wrap(err, "cannot bind request")
	}

	value := map[string]interface{}{}
	for _, v := range c.Params {
		value[v.Key] = v.Value
	}

	cfg := &mapstructure.DecoderConfig{
		Result:  &req,
		TagName: "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	err := decoder.Decode(req)
	if err != nil {
		return req, err
	}

	return req, nil
}
