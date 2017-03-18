package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	//"github.com/spf13/viper"
	//"strings"
	//"github.com/hyperledger/fabric/core/peer"
	"errors"
	"io"
)

const (
	CONFIGPATH = "peersafe.com/chaincode_db/invokeDB/config.js"
	cfgpath = "peersafe.com/helloworld/args.js"
)

type TrustObject struct {
	TrustCCid string
	AdminCert string
}

//期望成员列表
type ExpectInfo struct {
	ExpectList []string `json:"ExpectList"`
}

var trustObj TrustObject


func WriteFile(path, value string) error {
	fmt.Printf("writeFile --start-----path:%s\n", path)

	//if err := os.Remove(path); err != nil {
	//	return fmt.Errorf("os.Remove(path) err : %s", path)
	//}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("writeFile open file err : %s", err.Error())
	}

	fmt.Printf("writeFile -------value:%s\n", value)
	_, err = io.WriteString(f, value)
	if err != nil {
		return fmt.Errorf("writeFile write file err : %s", err.Error())
	}
	fmt.Printf("writeFile ----end---path:%s\n", path)

	return nil
}

func WriteTrustConfig() {
	newCfgJson, err := json.MarshalIndent(trustObj, "", "\t")
	if err != nil {
		fmt.Printf("json.Marshal ccid.js error %s ", err.Error())
	}
	if err := ioutil.WriteFile(CONFIGPATH, newCfgJson, os.ModeType); err != nil {
		fmt.Printf("ioutil.WriteFile cfgpath error %s ", err.Error())
	}
}

func ReadTrustConfig() {
	cfgjson, err := ioutil.ReadFile(CONFIGPATH)
	if os.IsNotExist(err) {
		//如果文件不存在
		fmt.Printf("bbbbbbbbbbbbbbbioutil.ReadFile %s error %s \n", CONFIGPATH, err.Error())
		WriteTrustConfig()
		return
	}

	if err != nil {
		fmt.Printf("ioutil.ReadFile %s error %s \n", CONFIGPATH, err.Error())
	}

	if cfgjson != nil {
		if err := json.Unmarshal(cfgjson, &trustObj); err != nil {
			fmt.Printf("json.Unmarshal %s error %s \n", CONFIGPATH, err.Error())
		}
	}
	fmt.Printf("cfgjson  TrustCCid  =  %s \n", trustObj.TrustCCid)
	fmt.Printf("cfgjson  AdminCert  =  %s \n", trustObj.AdminCert)
}

type PendingTxJson struct {
	PendingTx []PendingTxItem `json:"PendingTx"` //待处理交易项
}

type PendingTxItem struct {
	TxId       string `json:"TxId"`       //待处理交易流水id
	RootId     string `json:"RootId"`     //待处理交易根系列id
	StepNo     uint32 `json:"StepNo"`     //待处理交易步骤编号
	MemberType string `json:"MemberType"` //待处理交易交易的成员类型
}

func RemoveSlice(slice []PendingTxItem, start, end int) []PendingTxItem {
	//注意如果是移除一个时，end = start + 1
	return append(slice[:start], slice[end:]...)
}

func Remove(slice []PendingTxItem, start, end int) []PendingTxItem {
	result := make([]PendingTxItem, len(slice) - (end - start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}

func checkViper() {
	//var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl"}
	//configPath := "D:\\Work\\gocode\\src\\peersafe.com\\chaincode_loc\\citic_client_api\\test\\api_test\\client_sdk.yaml"
	//viper.SetConfigFile(configPath)
	//
	//err := viper.ReadInConfig()
	//if err != nil {
	//	fmt.Errorf("Fatal error when reading %s config file: %s", "client_sdk", err.Error())
	//}
	//fmt.Printf("ccid :%s\n",viper.Get("chaincode.id.name"))
	//return
}

type SaveTxData struct {
	PermTable []PermInfo   `json:"PermTable"`
}
type PermInfo struct {
	StepNo   uint32    `json:"stepNo"`
	PermList []MemberHashToPerm `json:"permList"`
}

type MemberHashToPerm struct {
	MemberHash string `json:"memberHash"`
	Perm       string `json:"perm"`
}

func checkSliceNum() error {
	permTable := "{\"PermTable\":[" +
		"{\"stepNo\":1,\"permList\":[{ \"memberHash\":\"\",\"perm\":\" r\"},{\"memberHash\":\"\",\"perm\":\"r\"},{\"memberHash\":\"\",\"perm\":\"r\"}]}," +
		"{\"stepNo\":2,\"permList\":[{\"memberHash\" :\"\",\"perm\":\"r\"}]}," +
		"{\"stepNo\":3,\"permList\":[{\"memberHash\" :\"\",\"perm\":\"r\"}]}," +
		"{\"stepNo\":4,\"permList\":[{\"memberHash\" :\"\",\"perm\":\"r\"}]}" +
		"]}"
	var data SaveTxData
	json.Unmarshal([]byte(permTable), &data)
	for i, k := range data.PermTable {
		//Rule_SetAll 规则权限表必须步骤号必须 1，2,3..顺序递增
		fmt.Printf("stepNo ：%d  i :%d\n", k.StepNo, i)
		if k.StepNo != uint32(i + 1) {
			return errors.New("k.StepNo != i + 1")
		}
		for _, m := range k.PermList {
			if m.Perm == "" {
				return errors.New("param perm is empty  error")
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("----------------helloworld start-------------------")
	//err := checkSliceNum()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}


	_, err := ioutil.ReadFile("D:\\Work\\gocode\\src\\helloworld\\mm1.txt")
	if err != nil{
		if os.IsNotExist(err){
			os.Create("D:\\Work\\gocode\\src\\helloworld\\mm1.txt")
		}
	}
	ioutil.WriteFile()
	//WriteFile("D:\\Work\\gocode\\src\\helloworld\\mm.txt","55")
	////ReadTrustConfig()
	////WriteTrustConfig()
	//list :=[] string {"user1","user2","user3" }
	//obj := ExpectInfo{
	//	ExpectList: list,
	//}
	//args , err := json.Marshal(obj)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//var	bb  []PendingTxItem
	//mm := PendingTxItem{
	//	TxId: "111",
	//	RootId: "222",
	//	StepNo: 0,
	//	MemberType: "memtype",
	//}
	//mm1 := PendingTxItem{
	//	TxId: "111b",
	//	RootId: "222",
	//	StepNo: 0,
	//	MemberType: "memtype",
	//}
	//bb = append(bb,mm)
	//bb = append(bb,mm1)
	//fmt.Printf("%v\n",bb)
	//fmt.Printf("------------------------------------\n")
	//bb = RemoveSlice(bb,0,1)
	//fmt.Printf("%v\n",bb)
	//ss := "{\"notifyMemberList\":[{\"member\":\"user3\"},{\"member\":\"user4\"}]}"
	//bb := strings.TrimSpace(ss)

	//ss1 := "{\"notifyMemberList\":[{\"member\":\"user3\"},{\"member\":\"user4\"}]}"

	//ss3 := " 1 2 3 4 "
	//bb1 := strings.TrimSpace(ss1)
	//bb2 :=strings.Trim(ss1,":")
	//bb3 := strings.Trim(ss3," ")

	//
	//bb3 := strings.Replace(ss1, " ", "", -1)
	//fmt.Printf("%s\n",ss1)
	//fmt.Printf("%s\n",bb1)
	//fmt.Printf("%s\n",bb2)
	//fmt.Printf("%s\n",bb3)
	//
	//mm := ""
	//mm1 := "w"
	////mm2 := "W"
	//mm3 := "RW"
	//fmt.Printf("%v\n",strings.Contains(mm,"w"))
	//fmt.Printf("%v\n",strings.Contains(mm,"W"))
	//fmt.Printf("%v\n",strings.Contains(mm1,"w"))
	//fmt.Printf("%v\n",strings.Contains(mm3,"W"))


}
