package gin_ctx

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"inititaryplanner/common/utils"
)

func TestGetCtxAndReqFromGinCtx(t *testing.T) {
	type Req struct {
		B1 string `json:"b1,omitempty"`
		B2 struct {
			B2_2 int `json:"b2_2,omitempty"`
		} `json:"b2,omitempty"`
		Q1 string   `json:"q1,omitempty"`
		Q2 []string `json:"q2,omitempty"`
		P1 string   `json:"p1,omitempty"`
	}

	tests := []struct {
		name    string
		ginCtx  *gin.Context
		wantReq Req
		wantCtx context.Context
		wantErr error
	}{
		{
			name: "param",
			ginCtx: &gin.Context{
				Keys: map[string]interface{}{
					"log_id": "log_id",
				},
				Request: &http.Request{
					Method: http.MethodGet,
				},
				Params: gin.Params{
					{Key: "p1", Value: "p1_value"},
				},
			},
			wantReq: Req{
				P1: "p1_value",
			},
			wantCtx: context.WithValue(context.Background(), "log_id", "log_id"),
		},
		{
			name: "query",
			ginCtx: &gin.Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: "q1=a",
					},
					Method: http.MethodGet,
				},
			},
			wantReq: Req{
				Q1: "a",
			},
			wantCtx: context.Background(),
		},
		{
			name: "query",
			ginCtx: &gin.Context{
				Request: &http.Request{
					URL: &url.URL{
						RawQuery: "q2=a&q2=b",
					},
					Method: http.MethodGet,
				},
				Params: gin.Params{
					{Key: "p1", Value: "p1_value"},
				},
			},
			wantReq: Req{
				P1: "p1_value",
				Q2: []string{"a", "b"},
			},
			wantCtx: context.Background(),
		},
		{
			name: "param",
			ginCtx: &gin.Context{
				Keys: map[string]interface{}{
					"log_id": "log_id",
				},
				Request: &http.Request{
					Body: io.NopCloser(bytes.NewBuffer([]byte(utils.SafeJson(Req{B1: "a", B2: struct {
						B2_2 int `json:"b2_2,omitempty"`
					}(struct{ B2_2 int }{B2_2: 1})})))),
					Header: map[string][]string{
						"Content-Type": []string{"application/json"},
					},
					Method: http.MethodPost,
				},
			},
			wantReq: Req{
				B1: "a",
				B2: struct {
					B2_2 int `json:"b2_2,omitempty"`
				}(struct{ B2_2 int }{B2_2: 1}),
			},
			wantCtx: context.WithValue(context.Background(), "log_id", "log_id"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, req, err := GetCtxAndReqFromGinCtx[Req](tt.ginCtx, tt.wantReq)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantCtx, ctx)
			assert.Equal(t, tt.wantReq, req)
		})
	}
}
