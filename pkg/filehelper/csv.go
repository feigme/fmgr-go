package filehelper

import (
	"encoding/csv"
	"errors"
	"fmt"

	"os"

	"github.com/axgle/mahonia"
)

func ReadCsv(path string, callback func(rows [][]string) error) error {
	if path == "" {
		return errors.New("文件路径为空! ")
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("读取csv失败！%+v", err)
		return err
	}
	defer file.Close()

	// 编码转换
	dec := mahonia.NewDecoder("utf-16")

	csvReader := csv.NewReader(dec.NewReader(file))
	csvReader.Comma = '\t'
	csvReader.FieldsPerRecord = -1

	readAll, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("读取csv内容失败！%+v", err)
	}

	return callback(readAll)
}
