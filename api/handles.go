package api

import (
	"base-convertor/convertor"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//convert-base?num=12&curBase=10&newBase=2
func convertBaseHandler(c echo.Context) error{
	number := c.QueryParam("num")
	curBase:= c.QueryParam("curBase")
	//newBase:= c.QueryParam("newBase")

	curBaseNum , err := strconv.Atoi(curBase)
	if err != nil {
		return c.String(http.StatusBadRequest , err.Error())
	}

	numIsOk , err := convertor.CheckBase(number, curBaseNum)
	if err != nil {
		return c.String(http.StatusBadRequest , err.Error())
	}
	if !numIsOk {
		return c.String(http.StatusBadRequest , "base dosen't match with number")
	}

	resultNum , err := convertor.ToDecimal(number,curBaseNum)

	return c.String(http.StatusAccepted ,strconv.Itoa(resultNum))
}