package service

import (
	"context"
	"testing"

	"github.com/hellojqk/config/server/model"
	"github.com/stretchr/testify/assert"
)

func TestUserLoginParam(t *testing.T) {
	code, err := UserLoginParam(context.Background(), model.UserLoginParam{Key: "admin", Password: "123123"})
	assert.Equal(t, nil, err)
	t.Log(code)
}

func TestUserTokenValid(t *testing.T) {
	key, err := UserTokenValid(context.Background(), "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTU3MTkxNDIsImlhdCI6MTYxNTY0NzE0MiwiaXNzIjoiY29uZmlnMSIsInN1YiI6ImFkbWluIn0.qkTzb50L7x5TG2VPUCOKAvZVMtfNy4RyUwWeD90kbUCxCFYPGqMbcA_DRri1zhnMFvCA3bnZOU-va1Pfvj2zauIE-yLiEcAQ60jcdMvMKbWWP-nIgiPoTb-WQcmkNFmsMOAuDn9mopOGJY-631sH4E3cPrQK6z9uQfYSkGXtqverMVjTXb2E3vP9L8_J-4Q9ntbOA2XOdEQkPi_eouinTQLxo0ZWz5zfaeErY9h7COGVOQmW9vLS_pPnkCeNyHMAZQNEUhB4wPFZcmo1C4bCIlmq6-eHLgDkd8lzpTvxW8GI_BpyfBScv3lU9ccf0IGrKsxk_Rwblndd3P0c8N0gNg")
	assert.Equal(t, nil, err)
	t.Log("key", key)
}
