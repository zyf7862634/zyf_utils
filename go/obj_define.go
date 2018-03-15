package define

//AuthorizationType 授权类型
type AuthorizationType int

const (
	//USER 本机构的当前操作人
	USER AuthorizationType = iota + 1

	//GROUP 本机构的当前操作角色
	GROUP

	//OTHER 指定的其他机构
	OTHER
)

//ResultStatus 调用结果状态
type ResultStatus int

const (
	//FAILED 调用失败
	FAILED ResultStatus = iota - 1

	//SUCCESS 调用成功
	SUCCESS
)

//AssertStatus 资产状态
type AssertStatus int

const (
	//NORMAL 正常
	NORMAL AssertStatus = iota - 1

	//FREEZE 冻结
	FREEZE

	//FINALITY 终结
	FINALITY
)

//AssetSpendStatus 调用结果状态
type AssetSpendStatus int

const (
	//SPENT 已花费
	SPENT AssetSpendStatus = iota - 1

	//UNSPENT 未花费
	UNSPENT
)

//error_code
const (
	//NO_AUTH_INVOKE 未授权资产操作
	NO_AUTH_INVOKE = "EC0000_001"

	//NO_AUTH_QUERY 未授权资产查询
	NO_AUTH_QUERY = "EC0000_002"

	//INVOKE_ASSET_NOT_MATCH_CC 操作资产类型与合约不匹配
	INVOKE_ASSET_NOT_MATCH_CC = "EC0000_003"

	//SCHEMA_CHECK_ERROR 资产Schema验证失败
	SCHEMA_CHECK_ERROR = "EC0000_004"

	//ASSERT_CREATE_FAILED 创建资产失败
	ASSERT_CREATE_FAILED = "ET0001_001"

	//ASSERT_TRANSFER_FAILED 资产转移失败
	ASSERT_TRANSFER_FAILED = "ET0002_001"

	//ASSERT_TRACE_FAILED 资产溯源失败
	ASSERT_TRACE_FAILED = "EQ0001_001"

	//ASSERT_QUERT_FAILED 资产交易数据失败
	ASSERT_QUERT_FAILED = "EQ0004_001"

	//ASSERT_UNSPENT_FAILED 资产未花失败
	ASSERT_UNSPENT_FAILED = "EQ0005_001"
)

//fee_rule_configuration
const (
	//YEARLY_USER 包年用户
	YEARLY_USER = "F1001"

	//MONTHLY_USER 包月用户
	MONTHLY_USER = "F1002"

	//COUNT_USER 计次用户
	COUNT_USER = "F2001"

	//PREPAID_USER 预付费用户
	PREPAID_USER = "F3001"

	//AFTER_PAYING_USER  后付费用户
	AFTER_PAYING_USER = "F4001"
)

//MetaData 资产对应的schema规则 等
type MetaData struct {
	CreateSchemaURL    string `json:"create_schema_url"`
	TransferSchemaURL  string `json:"transfer_schema_url"`
	CreateRuleScript   string `json:"create_rule_script"`
	TransferRuleScript string `json:"transfer_rule_script"`
}

//AssetInfo 资产价值对象
type AssetInfo struct {
	AssetType        string           `json:"asset_type"`         //资产类型
	MetaData         MetaData         `json:"meta_data"`          //资产元数据，比如schema url
	ACLList          []AssetAuthObj   `json:"acl_list"`           //资产授权列表
	BalanceList      []Balance        `json:"balance_list"`       //资产余额-数值属性
	RightList        []string         `json:"right_list"`         //权益列表
	ExtraAttrs       interface{}      `json:"extra_attrs"`        //扩展属性对象
	QueryCondition   interface{}      `json:"query_condition"`    //资产查询终止条件
	AssetStatus      AssertStatus     `json:"asset_status"`       //资产状态
	AssetSpendDtatus AssetSpendStatus `json:"asset_spend_status"` //资产花费状态
	AssertID         string           `json:"asset_id"`           //资产唯一标识(创建时产生)
	AssertAddress    string           `json:"asset_address"`      //资产地址
}

//Balance 资产余额对象
type Balance struct {
	ID     string `json:"id"`     //余额名称
	Value  uint32 `json:"value"`  //余额价值
	Remark string `json:"remark"` //余额描述信息
}

//AssetAuthObj 资产授权对象定义
type AssetAuthObj struct {
	AuthorizationType int    `json:"authorization_type"` //授权类型
	Authorizer        string `json:"authorizer"`         //授权人
	AuthorizedEntity  string `json:"authorized_entity"`  //被授权人
	PermissionList    []int  `json:"permission_list"`    //权限列表
}

//Sign 签名信息
type Sign struct {
	Signature string `json:"signature"`
	Cert      string `json:"cert"`
}

//TxParam invoke 交易params
type TxParam struct {
	TxIn  []TxIn  `json:"tx_in"`
	TxOut []TxOut `json:"tx_out"`
}

//TraceCondition 溯源条件
// type TraceCondition struct {
// 	TraceUpLevel   uint32 `json:"trace_up_level"`
// 	TraceDownLevel uint32 `json:"trace_down_level"`
// }

//TxOut UTXO资产交易输出
type TxOut struct {
	WalletAddress string    `json:"wallet_address"`
	Asset         AssetInfo `json:"assert_info"` //分配到该钱包资产价值
}

//TxIn UTXO资产交易输入
type TxIn struct {
	AssetAddr string `json:"assert_addr"` //资产地址的world state的key--AssetAddr
}

//AssetInherentProp 资产的固有属性
// type AssetInherentProp struct {
// 	ContractID          string //合约ID
// 	ContractDescription string //合约描述
// 	ContractAttachment  string //合约附加参数
// 	CreateTime          string //资产创建时间
// }

//AssetTrace 资产溯源对象
type AssetTrace struct {
	AssetAddress   string      `json:"asset_address"` //资产地址
	TraceCondition interface{} `json:"trace_level"`
}

//TxHeader 交易头
type TxHeader struct {
	TxID       string `json:"tx_id"`       //交易编号
	ContractID string `json:"contract_id"` //合约ID（chaincodeId）
	TimeStamp  int64  `json:"time_stamp"`  //时间戳
	Sign       []Sign `json:"sign"`        //txin对应的签名或者是调用方签名
}

//QueryCondition 查询条件
// type QueryCondition struct {
// 	StartTime int64 `json:"start_time"` //开始时间
// 	EndTime   int64 `json:"end_time"`   //结束时间
// }

