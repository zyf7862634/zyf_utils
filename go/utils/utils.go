package utils

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
)

type PendingTxItem struct {
	TxId       string `json:"TxId"`       //待处理交易流水id
	RootId     string `json:"RootId"`     //待处理交易根系列id
	StepNo     uint32 `json:"StepNo"`     //待处理交易步骤编号
	MemberType string `json:"MemberType"` //待处理交易交易的成员类型
}

//从切片移除某项
func RemoveSlice(slice []PendingTxItem, start, end int) []PendingTxItem {
	//注意如果是移除一个时，end = start + 1
	return append(slice[:start], slice[end:]...)
}

//判断文件或文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreatFile(fileName string) error {
	dir := filepath.Dir(fileName)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			// Directory does not exist, create it //注意只能创建目录 filePath.Dir
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
			if _, err := os.Create(fileName); err != nil {
				return err
			}
		} else if os.IsExist(err) {
			if _, err := os.Create(fileName); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func ReadFile(fileName string, isCreat bool) ([]byte, error) {
	//创建多级目录文件
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			if isCreat {
				if err := CreatFile(fileName); err != nil {
					return err
				}
				return ioutil.ReadFile(fileName)
			} else {
				return nil, nil
			}
		} else if os.IsExist(err) {
			return ioutil.ReadFile(fileName)
		} else {
			return nil, nil
		}
	}
	return nil, nil
}

func WriteFile(fileName string, data []byte) error {
	//创建多级目录文件
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			if err := CreatFile(fileName); err != nil {
				return err
			}
			return ioutil.WriteFile(fileName, data, 0755)
		} else if os.IsExist(err) {
			return ioutil.WriteFile(fileName, data, 0755)
		} else {
			return nil
		}
	}
	return nil
}

func JsonStrUnmarshal(data string, v interface{}) (bool, error) {
	if data == "" {
		return true, nil
	}
	if err := json.Unmarshal([]byte(data), v); err != nil {
		return false, err
	}
	return false, nil
}