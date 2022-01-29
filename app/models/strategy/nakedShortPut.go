package strategy

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/feigme/fmgr-go/app/models"
)

// 裸卖看跌
type NakedShortPut struct {
	Code  string // 期权code
	Count int    // 数量
	Price string // 价格
}

// 人机交互
func (o *NakedShortPut) Hci() error {
	blue.Println("请输入期权code: ")
	fmt.Scanln(&o.Code)
	_, err := models.CheckOptionCode(o.Code)
	if err != nil {
		return err
	}

	blue.Println("请输入数量: ")
	fmt.Scanln(&o.Count)
	if o.Count < 1 {
		return errors.New("数量必须大于0! ")
	}

	blue.Println("请输入价格: ")
	fmt.Scanln(&o.Price)
	_, err = strconv.ParseFloat(o.Price, 64)
	if err != nil {
		return errors.New("价格格式错误! ")
	}

	return nil
}
