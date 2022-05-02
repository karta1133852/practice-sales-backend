package models

import (
	"math"
	"strconv"
	"strings"
)

// UsersModel
type usersModel struct{} // 方便閱讀 private
type Users struct {      // 包裝給外部使用
	*usersModel
}

type ProductItem struct {
	ProductNo int
	Quantity  int
}

type OrderData struct {
	OriginalPrice int
	PayedCoin     int
	PayedPoint    int
	Exchange      int
	Discount      int
	Products      []ProductItem
}

func (_ *usersModel) CheckTotal(orderData OrderData, vipType string) error {

	appropriateTotal := int(math.Round(float64(orderData.OriginalPrice) * (1.0 - float64(orderData.Discount)/100.0)))
	equivalentTotal := orderData.PayedCoin + int(math.Round(float64(orderData.PayedPoint*orderData.Exchange)/100.0))

	// 此處為第二次新增的折扣方式
	// 如果有VIP身份扣100點以上折抵，另外享再九折優惠
	if vipType != "Normal" && orderData.PayedPoint >= 100 {
		appropriateTotal = int(math.Round(float64(appropriateTotal) * 0.9))
	}

	if appropriateTotal != equivalentTotal {
		return &CustomError{StatusCode: 422, Title: "參數錯誤", Message: "付款金額錯誤"}
	} else {
		return nil
	}
}

func (_ *usersModel) FormatProductItems(products []ProductItem) (strNo string, strQuantity string) {

	length := len(products)
	productNo := make([]string, length)
	quantity := make([]string, length)
	for i, product := range products {
		productNo[i] = strconv.Itoa(product.ProductNo)
		quantity[i] = strconv.Itoa(product.Quantity)
	}

	strNo = strings.Join(productNo[:], ", ")
	strQuantity = strings.Join(quantity[:], ", ")

	return
}
