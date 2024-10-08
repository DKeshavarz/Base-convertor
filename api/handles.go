package api

import (
	"base-convertor/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//convert-base?num=12&curBase=10&newBase=2
func convertBaseHandler(c echo.Context) error{
	number := c.QueryParam("num")
	curBase:= c.QueryParam("curBase")
	newBase:= c.QueryParam("newBase")

	curBaseNum , _ := strconv.Atoi(curBase)
	err := util.CheckBase(number, curBaseNum)
	fmt.Println(curBaseNum , err)
	return c.String(http.StatusAccepted , number + " " + curBase + "->" + newBase)
}