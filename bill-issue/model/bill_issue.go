/**
 * @Time: 2022/3/11 00:23
 * @Author: yt.yin
 */

package issuemod

// BillIssueRcd ETC 账单下发记录表
type BillIssueRcd struct {

	ID                  string         `json:"id"                  gorm:"column:id;primary_key;type:varchar(36)"`

	/** 支付账单id */
	OutId               string         `json:"outId"               gorm:"column:out_id;index;type:varchar(36)"`

	/** 商户号 */
	MerchantNo 	        string         `json:"merchantNo"          gorm:"column:merchant_no;index;type:varchar(36)"`

	/** 动作类型 0 是新扣费  -1 是删除 */
	Action              int            `json:"action"              gorm:"column:action;comment:0-下发,-1-删除;"`

	/** 节点服务器编号 */
	DeviceId            string         `json:"deviceId"            gorm:"column:device_id;comment:节点服务器编号;index;type:varchar(36)"`

	/** 记录创建 */
	CreateTime          string         `json:"createTime"          gorm:"column:create_time;comment:记录创建;type:varchar(20)"`

	/** 下发日期 */
	IssueDate           string         `json:"issueDate"           gorm:"column:issue_date;comment:下发日期;type:varchar(20)"`

	/** 下发时间 */
	IssueTime   	    string         `json:"issueTime"           gorm:"column:issue_time;comment:下发时间;type:varchar(20)"`

	/** 下发次数 */
	IssueCount          int            `json:"issueCount"          gorm:"column:issue_count;comment:下发次数;"`

	/** 账单内容 */
	BillContent         string         `json:"billContent"         gorm:"column:bill_content;comment:账单内容;type:text"`

	/** 备注信息 */
	Remark        	    string         `json:"remark"              gorm:"column:remark;comment:备注信息 ;type:text"`
}

// BillIssueRcdTemp 下发临时记录表,收到ACK确认报文后挪到成功表
type BillIssueRcdTemp struct{
	BillIssueRcd
}

// TableName 下发临时记录表
func (BillIssueRcdTemp) TableName() string {
	return "bill_issue_rcd_temp"
}

// BillIssueRcdSuccess 下发成功记录表
type BillIssueRcdSuccess struct{
	BillIssueRcd
}

// TableName 下发成功记录表
func (BillIssueRcdSuccess) TableName() string {
	return "bill_issue_rcd_success"
}

// BillIssueRcdError 下发错误表
type BillIssueRcdError struct{
	BillIssueRcd
}

// TableName 下发成功记录表
func (BillIssueRcdError) TableName() string {
	return "bill_issue_rcd_error"
}


