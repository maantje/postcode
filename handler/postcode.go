package handler

import (
	"database/sql"
	"net/http"

	"strings"

	. "github.com/go-jet/jet/v2/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/maantje/postcode/.gen/model"
	. "github.com/maantje/postcode/.gen/table"
	"github.com/maantje/postcode/resource"
	"github.com/maantje/postcode/util"
)

type PostcodeHandler struct {
	Db *sql.DB
}

func (h *PostcodeHandler) Search(c echo.Context) error {
	address := &model.Addresses{}

	postcode := strings.ToUpper(
		strings.ReplaceAll(c.Param("postcode"), " ", ""),
	)

	houseNumber := util.ParamInt64(c, "houseNumber", 0)
	lAddition := strings.ToLower(c.QueryParam("addition"))
	uAddition := strings.ToUpper(c.QueryParam("addition"))

	stmt := SELECT(
		Addresses.AllColumns,
	).FROM(
		Addresses,
	).WHERE(
		Addresses.Postcode.EQ(String(postcode)).
			AND(Addresses.HouseNumber.EQ(Int(houseNumber))).
			AND(Addresses.Addition.EQ(String(lAddition)).OR(Addresses.Addition.EQ(String(uAddition)))),
	)

	if err := stmt.Query(h.Db, address); err != nil {
		return c.NoContent(404)
	}

	return c.JSON(http.StatusOK, resource.NewAddressResource(address))
}
