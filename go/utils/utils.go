package utils

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/json"
	"strings"
	"errors"
	"math/rand"
	"fmt"
	"time"
	"reflect"
	"bytes"
	"bufio"
	"io"
	"github.com/op/go-logging"
)

type PendingTxItem struct {
	TxId       string `json:"TxId"`       //待处理交易流水id
	RootId     string `json:"RootId"`     //待处理交易根系列id
	StepNo     uint32 `json:"StepNo"`     //待处理交易步骤编号
	MemberType string `json:"MemberType"` //待处理交易交易的成员类型
}
//随机字符串
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var rnd = rand.NewSource(time.Now().UnixNano())
func randomString(n int) string {
	b := make([]byte, n)

	for i, cache, remain := n-1, rnd.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rnd.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

//设置log级别
func SetLogLevel(modelLevel, modelName string)  error{
	if modelLevel == "" {
		modelLevel = "DEBUG"
	}
	format := logging.MustStringFormatter("%{shortfile} %{time:2006-01-02 15:04:05.000} [%{module}] %{level:.4s} : %{message}")
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logLevel, err := logging.LogLevel(modelLevel)
	if err != nil{
		return err
	}
	//map[k]v; eg:  var logger = logging.MustGetLogger("event")
	logging.SetBackend(backendFormatter).SetLevel(logLevel, modelName)
	return nil
}

//从切片移除某项
func RemoveSlice(slice []PendingTxItem, start, end int) []PendingTxItem {
	//注意如果是移除一个时，end = start + 1
	return append(slice[:start], slice[end:]...)
}

//结构体切片 变成 []interface
//eg: []PendingTxItem --> []interface{}
// 不能直接转 interface 可以接受任何类型， []interface{}不可以
//[]interface{} = ToSlice([]PendingTxItem)
func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
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

//创建文件 如果目录不存在就创建目录再创建文件
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

//读文件 不存在是否创建
func ReadFile(fileName string, isCreat bool) ([]byte, error) {
	//创建多级目录文件
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			if isCreat {
				if err := CreatFile(fileName); err != nil {
					return nil,err
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
	return ioutil.ReadFile(fileName)
}

//写文件不存在会创建文件
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
	return ioutil.WriteFile(fileName, data, 0755)
}

//字符串解析成结构体，如果字符串为空 返回为true
func JsonStrUnmarshal(data string, v interface{}) (bool, error) {
	if data == "" {
		return true, nil
	}
	if err := json.Unmarshal([]byte(data), v); err != nil {
		return false, err
	}
	return false, nil
}

//去除字符串的所有空格， strings.TrimSpace(str) 只能去掉首尾空格
func StrRemoveSpace(source string) string {
	return strings.Replace(source, " ", "", -1)
}

//添加到字符串数组，不可重复添加
func AddStr(strList []string, str string) ([]string, error) {
	if ContainsStr(strList, str) {
		return strList, errors.New("The member exist in the list")
	}
	return append(strList, str), nil
}

//从字符串数组中移除，必须在数组中包含
func RemoveStr(strList []string, str string) ([]string, error) {
	var index = int(-1)
	for i, v := range strList {
		if str == v {
			index = i
			break
		}
	}
	if index < 0 {
		return strList, errors.New("There's no exist the member:" + str)
	}

	return append(strList[:index], strList[index+1:]...), nil
}

//字符串数组中包含对应字符串
func ContainsStr(strList []string, str string) bool {
	for _, v := range strList {
		if v == str {
			return true
		}
	}
	return false
}

//获取当前时间
func getCurTime() uint64 {
	return uint64(time.Now().UTC().Unix())
}

//彩色打印
func colorPrint(log string) {
	fmt.Printf("\033[1m\033[45;33m" + log + "\033[0m\n")
}


func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func ToChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

func ArrayToChaincodeArgs(args []string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

//JSONDecode json decode
func JSONDecode(data []byte, ret interface{}) error {
	if len(data) == 0 {
		return errors.New("JsonDecode failed data is nil")
	}
	return json.Unmarshal(data, ret)
}

var loggera = logging.MustGetLogger("aaa")

//JSONEncode json encode
func JSONEncode(data interface{}) ([]byte, error) {
	loggera.Debug("JSONEncodedd")
	if data == nil {
		return nil, errors.New("JSONEncode failed data is nil")
	}
	return json.Marshal(data)
}

//生成随机数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min) + min
	return randNum
}


func ModifyHosts(filePath, newIp, domain string) error {
	//调用该函数需要用户对/etc/hosts 可写， sudo chown ubuntu:ubuntu /etc/hosts
	Spacer := "       "//7 space
	newHosts, err := ioutil.ReadFile(filePath)//读取原有hosts 保存
	if err != nil {
		return err
	}
	fi, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	allMap := make(map[string]string)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		qq := strings.Split(string(a), Spacer)
		if len(qq) >= 2 {
			allMap[qq[1]] = qq[0]
		}
	}
	//如果域名存在
	tempNew := newIp+ Spacer + domain
	if ip, ok := allMap[domain]; ok {//如果存在之前的域名
		tempOld := ip + Spacer + domain
		newHosts = []byte(strings.Replace(string(newHosts),tempOld,tempNew,1))
	}else {//如果不存在之前的域名，则直接追加
		newHosts = []byte(string(newHosts) + "\n" + tempNew)
	}
	if err := ioutil.WriteFile(filePath,newHosts,0644); err != nil {
		return err
	}
	return nil
}

//比较大小
func min(a uint64, b uint64) uint64 {
	return b ^ ((a ^ b) & (-(uint64(a-b) >> 63)))
}
