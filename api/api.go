package api

import (
	"context"
	"fmt"
	"goproject_port/datastruct"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type serv interface {
	GetIn(nIn int) (int, error)
	GetChIn(nIn int) (chan datastruct.Reqest, error)
	PostOut(nOut int, valOut int) error
	PostChOut(nOut int) (chan datastruct.Reqest, error)

	GetAll()
}

type api struct {
	Serv serv
}

func New(serv serv, numIn int, numOut int, ctx context.Context) *api {
	a := &api{Serv: serv}

	for i := 0; i < numIn; i++ {
		go a.GetIn(i, ctx)
	}
	for i := 0; i < numOut; i++ {
		go a.PostOut(i, ctx)
	}

	return a
}

func (a *api) Run(Host string) {
	r := gin.Default()
	r.GET("/port/READ", a.GetChIn)
	r.POST("/port/WRITE", a.PostChOut)
	//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	r.GET("/port/all", a.GetAll)

	go r.Run(Host)

}

func (a *api) GetChIn(g *gin.Context) {
	GetContext(g)
	respChan := make(chan datastruct.Response, 1)
	nInS := g.Request.URL.Query().Get("n")
	nIn, err := strconv.Atoi(nInS)
	if err != nil {
		logrus.Error("Failed to convert string to int ", err.Error())
		return
	}
	ch, err := a.Serv.GetChIn(nIn)
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		logrus.Error(err)
		return
	}

	select {
	case ch <- datastruct.Reqest{Resp: respChan}:
		logrus.Debug("Reqest start in chan № ", nIn)
	case <-g.Done():
		logrus.Debug("Context done")
		return
	}
	var response datastruct.Response
	select {
	case response = <-respChan:
		logrus.Debug("Response end in chan № ", nIn)
	case <-g.Done():
		logrus.Debug("Context done")
		return
	}

	if response.Error != nil {
		g.JSON(http.StatusBadRequest, response.Error.Error())
		logrus.Error(response.Error.Error())
		return
	}
	logrus.Info(fmt.Sprintf("The value of port number %v is = %v", nIn, response.Value))
	g.JSON(http.StatusOK, response.Value)
}

func (a *api) GetIn(nIn int, ctx context.Context) {
	ch, _ := a.Serv.GetChIn(nIn)
	defer close(ch)
	for {
		var reqest datastruct.Reqest
		select {
		case reqest = <-ch:
			logrus.Debug("Reqest end in chan № ", nIn)
		case <-ctx.Done():
			logrus.Debug("Context done")
			return
		}
		valIn, err := a.Serv.GetIn(nIn)
		response := datastruct.Response{Error: err, Value: valIn}
		select {
		case reqest.Resp <- response:
			logrus.Debug("Response start in chan № ", nIn)
		case <-ctx.Done():
			logrus.Debug("Context done")
			return
		}
	}
}

func (a *api) PostChOut(g *gin.Context) {
	respChan := make(chan datastruct.Response, 1)
	param := struct {
		NOut   int
		Value  int
		TransN int
	}{}
	err := g.ShouldBindJSON(&param)
	if err != nil {
		logrus.Error("Error create JSON", err.Error())
		return
	}
	ch, err := a.Serv.PostChOut(param.NOut)
	if err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
		logrus.Error(err)
		return
	}
	select {
	case ch <- datastruct.Reqest{Resp: respChan, Value: param.Value}:
		logrus.Debug("Reqest start in chan № ", param.NOut)
	case <-g.Done():
		logrus.Debug("Context done")
		return
	}
	var response datastruct.Response
	select {
	case response = <-respChan:
		logrus.Debug("Response end in chan № ", param.NOut)
	case <-g.Done():
		logrus.Debug("Context done")
		return
	}

	if response.Error != nil {
		g.JSON(http.StatusBadRequest, response.Error.Error())
		logrus.Error(response.Error.Error())
		return
	}
	logrus.Info(fmt.Sprintf("Success writing transaction %v with value %v to port number %v Out", param.TransN, param.Value, param.NOut))
	g.JSON(http.StatusOK, param)
}

func (a *api) PostOut(nOut int, ctx context.Context) {
	ch, _ := a.Serv.PostChOut(nOut)
	defer close(ch)
	for {
		var reqest datastruct.Reqest
		select {
		case reqest = <-ch:
			logrus.Debug("Reqest end in chan № ", nOut)
		case <-ctx.Done():
			logrus.Debug("Context done")
			return
		}
		err := a.Serv.PostOut(nOut, reqest.Value)
		response := datastruct.Response{Error: err, Value: 0}
		select {
		case reqest.Resp <- response:
			logrus.Debug("Response start in chan № ", nOut)
		case <-ctx.Done():
			logrus.Debug("Context done")
			return
		}
	}
}

func (a *api) GetAll(g *gin.Context) {
	a.Serv.GetAll()
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func GetContext(g *gin.Context) {
	query := g.Request.URL.Query()
	for key, values := range query {
		fmt.Printf("Parameter: %s, Values: %v\n", key, values)
	}
}
