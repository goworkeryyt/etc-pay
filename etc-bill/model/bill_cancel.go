/**
 * @Time: 2022/3/10 18:41
 * @Author: yt.yin
 */

package billmod

// EtcBillCancelRequest ETC账单撤销请求记录表
type EtcBillCancelRequest struct {

	ID                string       `json:"id"               gorm:"column:id;comment:主键;primary_key;type:varchar(36);"`

	/** 支付账单id */
	OutId             string       `json:"outId"            gorm:"column:out_id;comment:支付账单id;index;type:varchar(36);"`

	/** 商户编号 */
	MerchantNo        string       `json:"merchantNo"       gorm:"column:merchant_no;comment:商户编号;index;type:varchar(36);"`

	/** 车牌 */
	Vlp               string       `json:"vlp"              gorm:"column:vlp;comment:车牌;type:varchar(20);index;"`

	/** 撤销请求时间 */
	CreateTime        string       `json:"createTime"       gorm:"column:create_time;comment:撤销请求时间;type:varchar(20)"`

	/** 撤销处理状态  00 撤销成功 01 撤销失败 */
	Status            string       `json:"status"           gorm:"column:status;comment:撤销处理状态;type:varchar(2)"`

	/** 备注 */
	Remark            string       `json:"remark"           gorm:"column:remark;comment:备注;type:varchar(255)"`

	/** 撤销成功时间 */
	CancelTime        string       `json:"cancelTime"       gorm:"column:cancel_time;comment:撤销时间;type:varchar(32)"`
}

// TableName ETC账单撤销请求记录表
func (EtcBillCancelRequest) TableName() string {
	return "etc_bill_cancel_request"
}
