package api

import (
	"A-module/log"
	iBoFOSV1 "A-module/routers/mtool/api/ibofos/v1"
	"A-module/routers/mtool/model"
	"A-module/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpResponse(c *gin.Context, res model.Response, err error) {
	switch err {
	case iBoFOSV1.ErrBadReq:
		BadRequest(c, 12000)
	case iBoFOSV1.ErrSending:
		makeResponse(c, 400, error.Error(err), 19002)
	case iBoFOSV1.ErrJson:
		makeResponse(c, 400, error.Error(err), 12310)
	case iBoFOSV1.ErrRes:
		BadRequest(c, res.Result.Status.Code)
	default:
		SuccessWithRes(c, res)
	}
}
func Unauthorized(ctx *gin.Context, code int) {
	makeResponse(ctx, http.StatusUnauthorized, "", code)
}

func BadRequest(ctx *gin.Context, code int) {
	makeResponse(ctx, http.StatusBadRequest, "", code)
}

func makeResponse(ctx *gin.Context, httpStatus int, description string, code int) {
	res := model.Response{}
	makeResponseWithRes(ctx, httpStatus, description, code, res)
}

func makeResponseWithRes(ctx *gin.Context, httpStatus int, description string, code int, res model.Response) {
	res.Result.Status.Code = code
	if description == "" {
		res.Result.Status.Description = StatusDescription(code)
	}
	log.Printf("makeResponse : %+v", res)
	ctx.AbortWithStatusJSON(httpStatus, &res)
}

func Success(ctx *gin.Context) {
	res := model.Response{}
	SuccessWithRes(ctx, res)
}

func SuccessWithRes(ctx *gin.Context, res model.Response) {
	makeResponseWithRes(ctx, http.StatusOK, "", 0, res)
}

// https://golang.org/doc/effective_go.html#Getters
// it's neither idiomatic nor necessary to put Get into the getter's name.
func StatusDescription(code int) string {
	return setting.StatusMap[code]
}
