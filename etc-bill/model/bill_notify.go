/**
 * @Time: 2022/3/10 18:24
 * @Author: yt.yin
 */

package billmod

// EtcBillNotifyRcd ETC账单回调记录表
type EtcBillNotifyRcd struct{

	ID                string       `json:"id"               gorm:"column:id;comment:主键;primary_key;type:varchar(36);"`

	/** 支付账单id */
	OutId             string       `json:"outId"            gorm:"column:out_id;comment:支付账单id;index;type:varchar(36);"`

	/** 商户编号 */
	MerchantNo        string       `json:"merchantNo"       gorm:"column:merchant_no;comment:商户编号;index;type:varchar(36);"`

	/** 回调时间 */
	NotifyTime        string       `json:"notifyTime"       gorm:"column:notify_time;comment:回调时间;type:varchar(32);"`

	/** 请求方法 */
	Method            string       `json:"method"           gorm:"column:method;comment:请求方法;type:varchar(10);"`

	/** 请求路径 */
	Path              string       `json:"path"             gorm:"column:path;comment:请求路径;type:varchar(255);"`

	/** 请求参数 */
	Params            string       `json:"params"           gorm:"column:path;comment:请求参数;type:text;"`

	/** 请求状态 */
	Status            int          `json:"status"           gorm:"column:status;comment:请求状态"`

	/** 延迟即请求耗时 */
	Latency           int64        `json:"latency"          gorm:"column:latency;comment:延迟"`

	/** 响应结果 */
	Resp              string       `json:"resp"             gorm:"column:resp;comment:响应Body;type:longtext;"`

	/** 错误信息 */
	Error             string       `json:"error"            gorm:"column:error;comment:错误信息;type:text;"`
}

// TableName ETC账单回调记录表
func (EtcBillNotifyRcd) TableName() string {
	return "etc_bill_notify_rcd"
}
