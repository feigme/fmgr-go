package filehelper

import (
	"errors"

	"github.com/xuri/excelize/v2"
)

// 读取excel
func ReadExcel(filePath string, callback func(rows [][]string) error) error {
	if filePath == "" {
		return errors.New("文件路径为空! ")
	}
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return err
	}

	defer f.Close()
	sheetList := f.GetSheetList()
	for _, sheet := range sheetList {

		// 获取 Sheet 上所有单元格
		rows, err := f.GetRows(sheet)
		if err != nil {
			return err
		}

		err = callback(rows)
		if err != nil {
			return err
		}
	}

	return nil
}
