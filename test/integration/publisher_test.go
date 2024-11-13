package integration_test

import (
	"base-gin/app/domain/dto"
	"base-gin/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Success(t *testing.T) {
	req := dto.PublisherCreateReq{
		Name: util.RandomStringAlpha(8),
		City: util.RandomStringAlpha(10),
	}

	w := doTest("POST", "/v1/publishers", req,
		createAuthAccessToken(dummyAdmin.Account.Username))
	assert.Equal(t, 201, w.Code)
}