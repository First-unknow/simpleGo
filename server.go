package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"simpleGo/src"
)

type User struct {
	MemberId    string `json:"memberId" form:"memberId" query:"memberId"`
	ProductCode string `json:"productCode" form:"productCode" query:"productCode"`
	TerminalId  string `json:"terminalId" form:"terminalId" query:"terminalId"`
	Volumn      string `json:"volumn" form:"volumn" query:"volumn"`
	Price       string `json:"price" form:"price" query:"price"`
}

type Map map[string]interface{}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.POST("/updatePoint", updatePoint)

	// Start server
	e.Logger.Fatal(e.Start(":7000"))
}

// Handler
func updatePoint(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	currentMember := src.FindMember(string(u.MemberId))
	if _, haveMemberId := currentMember["cardNumber"]; !haveMemberId {
		return c.JSON(http.StatusNotFound, "Not Found")
	}
	currentProduct := src.FindProduct(string(u.ProductCode))
	if _, haveProductCode := currentProduct["productCode"]; !haveProductCode {
		return c.JSON(http.StatusNotFound, "Not Found")
	}

	mockupRes := Map{
		"memberId":     u.MemberId,
		"productName":  currentProduct["productName"],
		"receivePoint": 0.0,
	}

	if currentProduct["productType"] == "Oil" {
		maxCapabilityOil := src.FindCapabilityOil(currentMember["memberClass"], currentProduct["productGroup"])
		if maxCapabilityOil == -1 {
			return c.JSON(http.StatusNotFound, "Not Found")
		}
		volumnTypeInt, err := strconv.Atoi(u.Volumn)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusNotFound, "Not Found")
		}
		receivePoint := src.PointOil(volumnTypeInt, maxCapabilityOil, currentProduct, currentMember["cardType"])
		if receivePoint == -1 {
			return c.JSON(http.StatusNotFound, "Not Found")
		}

		mockupRes["receivePoint"] = receivePoint
		return c.JSON(http.StatusOK, mockupRes)

	} else if currentProduct["productType"] == "Non-Oil" {
		buSize := src.FindBuSize(u.TerminalId)
		maxCapabilityNonOil := src.FindCapabilityNonOil(buSize)
		if maxCapabilityNonOil == -1 {
			return c.JSON(http.StatusNotFound, "Not Found2")
		}
		priceTypeInt, err := strconv.Atoi(u.Price)
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusNotFound, "Not Found3")
		}

		receivePoint := src.PointNonOil(priceTypeInt, maxCapabilityNonOil, currentProduct)
		if receivePoint == -1 {
			return c.JSON(http.StatusNotFound, "Not Found4")
		}

		mockupRes["receivePoint"] = receivePoint
		return c.JSON(http.StatusOK, mockupRes)
	}

	return c.JSON(http.StatusOK, mockupRes)
}
