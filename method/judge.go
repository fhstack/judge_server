package method

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/l-f-h/common"
	"github.com/l-f-h/judge_server/rpc"
	"github.com/l-f-h/judge_server/rpc/generated"
	"log"
	"net/http"
	"strconv"
)

func constructJudgeRequest(c *gin.Context) (*pb_gen.JudgeRequest, error) {
	req := pb_gen.JudgeRequest{}

	req.SourceCode = c.PostForm("source_code")
	req.ProblemId = c.PostForm("problem_id")
	req.SubmitId = c.PostForm("submit_id")
	ml, err := strconv.ParseFloat(c.PostForm("memory_limit"), 32)
	if err != nil {
		log.Print("parse memory_limit filed failed")
		return nil, err
	}
	tl, err := strconv.ParseFloat(c.PostForm("time_limit"), 32)
	if err != nil {
		log.Print("parse time_limit filed failed")
		return nil, err
	}
	req.MemoryLimit = float32(ml)
	req.TimeLimit = float32(tl)
	req.Language = c.PostForm("language")
	req.CallBackUrl = c.PostForm("callback_url")
	return &req, nil
}

func Judge(c *gin.Context) {
	req, err := constructJudgeRequest(c)
	if err != nil {
		c.String(http.StatusInternalServerError, common.Message[common.ServerInternalError])
		return
	}
	if req.CallBackUrl != "" {
		c.String(http.StatusOK, "Already CallBack")
		rpc.Judge(context.Background(), req)
		return
	}
	resp := rpc.Judge(context.Background(), req)
	if resp.ResCode == common.RpcFailed {
		c.String(http.StatusInternalServerError, common.Message[common.RpcFailed])
		return
	}
	log.Println("Rpc successful")
	c.JSON(http.StatusOK, gin.H{
		"ResCode":     resp.ResCode,
		"Message":     common.Message[resp.ResCode],
		"Time":        resp.Time,
		"Memory":      resp.Memory,
		"CompileInfo": resp.CompileInfo,
	})
}
