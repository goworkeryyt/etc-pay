/**
 * @Time: 2022/3/10 17:03
 * @Author: yt.yin
 */

package model

// EtcBillBase ETC支付账单基础表
type EtcBillBase struct {
	ID string `json:"id"               gorm:"column:id;comment:主键;primary_key;type:varchar(36);"`

	/** 支付账单id */
	OutId string `json:"outId"            gorm:"column:out_id;comment:支付账单id;index;type:varchar(36);"`

	/** 商户编号 */
	MerchantNo string `json:"merchantNo"       gorm:"column:merchant_no;comment:商户编号;index;type:varchar(36);"`

	/** 区域编号 */
	RegionNo string `json:"regionNo"         gorm:"column:region_no;comment:区域编号;type:varchar(36);"`

	/** 路段编号 */
	RoadNo string `json:"roadNo"           gorm:"column:road_no;comment:路段编号;type:varchar(36);index;"`

	/** 停车场编号 */
	ParkNo string `json:"parkNo"           gorm:"column:park_no;comment:停车场编号;type:varchar(36);index;"`

	/** 车牌 */
	Vlp string `json:"vlp"              gorm:"column:vlp;comment:车牌;type:varchar(20);index;"`

	/** 入场时间 */
	EnterTime string `json:"enterTime"        gorm:"column:enter_time;comment:入场时间;type:varchar(20);"`

	/** 出场时间 */
	ExitTime string `json:"exitTime"         gorm:"column:exit_time;comment:出场时间;type:varchar(20);"`

	/** obuMac */
	ObuMac string `json:"obuMac"           gorm:"column:obu_mac;comment:obu mac;type:varchar(16);"`

	/** 扣费金额，单位：分 */
	ConsumeMoney string `json:"consumeMoney"     gorm:"column:consume_money;comment:扣费金额单位:分;type:varchar(10);"`

	/** 记录接收时间 */
	CreateTime string `json:"createTime"       gorm:"column:create_time;comment:记录接收时间;type:varchar(32);"`
}

// EtcBillTemp ETC支付账单请求临时表
type EtcBillTemp struct {
	EtcBillBase
}

// TableName 支付账单请求临时表
func (EtcBillTemp) TableName() string {
	return "etc_bill_temp"
}

// EtcBillCancel 已撤销成功账单
type EtcBillCancel struct {
	EtcBillBase

	/** 撤销成功时间 */
	CancelTime string `json:"cancelTime"       gorm:"column:cancel_time;comment:撤销成功时间;type:varchar(32)"`
}

// TableName 已撤销成功账单表
func (EtcBillCancel) TableName() string {
	return "etc_bill_cancel"
}

// EtcBillIssue 已下发账单
type EtcBillIssue struct {
	EtcBillBase

	/** 下发给设备时间 */
	IssueTime string `json:"issueTime"        gorm:"column:issue_time;comment:下发给设备时间;type:varchar(20)"`
}

// TableName 已下发账单表
func (EtcBillIssue) TableName() string {
	return "etc_bill_issue"
}

// EtcBillTradeSuccess ETC账单交易成功记录表
type EtcBillTradeSuccess struct {
	EtcBillBase

	/** 交易时间 */
	TradeTime string `json:"tradeTime"        gorm:"column:trade_time;comment:交易时间;type:varchar(20)"`

	/** 交易成功上送时间 */
	UploadTime string `json:"uploadTime"       gorm:"column:upload_time;comment:交易成功上送时间;type:varchar(32)"`
}

// TableName ETC账单交易成功记录表
func (EtcBillTradeSuccess) TableName() string {
	return "etc_bill_trade_success"
}

// EtcBillNotifyFail 回调失败表
type EtcBillNotifyFail struct {
	EtcBillTradeSuccess

	/** 最后一次回调时间 */
	LastNotifyTime string `json:"lastNotifyTime"   gorm:"column:last_notify_time;comment:最后一次回调时间;type:varchar(32);"`

	/** 回调次数 */
	NotifyCount string `json:"notifyCount"      gorm:"column:notify_count;comment:回调次数;"`

	/** 回调失败原因 */
	NotifyFailErr string `json:"notifyFailErr"    gorm:"column:notify_fail_err;comment:回调失败明细;type:text"`
}

// TableName 回调失败表
func (EtcBillNotifyFail) TableName() string {
	return "etc_bill_notify_fail"
}

// EtcBillNotifySuccess 最终回调成功表，此单结束
type EtcBillNotifySuccess struct {
	EtcBillTradeSuccess

	/** 最后一次回调时间 */
	LastNotifyTime string `json:"lastNotifyTime"   gorm:"column:last_notify_time;comment:最后一次回调时间;type:varchar(32);"`

	/** 回调次数 */
	NotifyCount string `json:"notifyCount"      gorm:"column:notify_count;comment:回调次数;"`
}

// TableName 最终回调成功表
func (EtcBillNotifySuccess) TableName() string {
	return "etc_bill_notify_success"
}
