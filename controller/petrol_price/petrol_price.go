package petrol_price

import (
	"backend/controller"
	"backend/dal"
	"backend/dal/dao"
	"backend/dal/model"
	"gorm.io/gorm"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(group *echo.Group) {
	group.GET("", query)
}

func query(c echo.Context) error {
	region := c.QueryParam("region")
	if len(region) == 0 {
		region = "北京"
	}
	d1 := dao.NewDao[model.PetrolPrice](dal.DB)
	recordList, err := d1.QueryList(func(db *gorm.DB) *gorm.DB {
		return db.Where(`region like ?`, "%"+region+"%")
	})
	if err != nil {
		return err
	}
	type petrolPrice struct {
		ReleaseDate string `json:"release_date"`
		Price0      string `json:"price_0"`
		Price92     string `json:"price_92"`
		Price95     string `json:"price_95"`
		Price98     string `json:"price_98"`
	}
	petrolPriceList := make([]petrolPrice, 0, len(recordList))
	for _, v := range recordList {
		petrolPriceList = append(petrolPriceList, petrolPrice{
			Price0:      v.Price0,
			Price92:     v.Price92,
			Price95:     v.Price95,
			Price98:     v.Price98,
			ReleaseDate: v.ReleaseDate,
		})
	}
	return c.JSON(http.StatusOK, controller.Output{
		Code:    0,
		Message: "",
		Data: map[string]interface{}{
			"petrol_price_list": petrolPriceList,
		},
	})
}
