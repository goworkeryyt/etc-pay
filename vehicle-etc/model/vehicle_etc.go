/**
 * @Time: 2022/3/10 19:04
 * @Author: yt.yin
 */

package vetcmod

// VehicleEtcInfo 车辆ETC信息
type VehicleEtcInfo struct {

	/** 主键 */
	ID                string       `json:"id"               gorm:"column:id;primary_key;type:varchar(36);"`

	/** 创建时间 */
	CreateTime        string       `json:"createTime"       gorm:"column:create_time;comment:创建时间;index;type:varchar(20);"`

	/** 更新时间 */
	UpdateTime        string       `json:"updateTime"       gorm:"column:update_time;comment:更新时间;type:varchar(20);"`

	/** 车牌号码 */
	Vlp               string       `json:"vlp"              gorm:"column:vlp;comment:车牌;index;;type:varchar(20);"`

	/** 车牌颜色编码 */
	Vlpc              string       `json:"vlpc"             gorm:"column:vlpc;comment:车牌颜色;type:varchar(4);"`

	/** 车辆类型编码 */
	Vc                string       `json:"vc"               gorm:"column:vc;comment:车辆类型;type:varchar(4);"`

	/** 卡上余额 */
	Balance           int64        `json:"balance"          gorm:"column:balance;comment:余额;"`

	/** OBU序号编码 此处放的obu Mac地址 并非合同序号 */
	ObuMac            string       `json:"obuMac"           gorm:"column:obu_mac;comment:obu mac;index;type:varchar(16);"`

	/** ETC 卡号 */
	CardNo            string       `json:"cardNo"           gorm:"column:card_no;comment:ETC卡号;index;type:varchar(20);"`

	/** 卡片类型：22-储值卡，23-记账卡 */
	CardType          int          `json:"cardType"         gorm:"column:card_type;comment:卡片类型：22-储值卡，23-记账卡;type:varchar(4);"`

	/** 发行商代码 */
	Issuer            string       `json:"issuer"           gorm:"column:issuer;comment:发行商代码;type:varchar(20);"`

	/** 卡过期日期 */
	CardExpireDate    string       `json:"cardExpireDate"   gorm:"column:card_expire_date;comment:卡过期日期;type:varchar(10);"`

	/** Obu过期日期 */
	ObuExpireDate    string        `json:"obuExpireDate"    gorm:"column:obu_expire_date;comment:Obu过期日期;type:varchar(10);"`

	/** obu 合同序号 */
	ObuSerialNumber   string       `json:"obuSerialNumber"  gorm:"column:obu_serial_number;comment:obu合同序号;type:varchar(20);index"`

	/** etc状态 00 正常 01 黑名单 02 余额不足 05 未插卡 06 已过有效期 09 其他异常*/
	EtcStatus         string       `json:"etcStatus"        gorm:"column:etc_status;comment:etc状态 00 正常 01 黑名单 02 余额不足 05 未插卡 09 其他异常;type:varchar(4);"`
}

// TableName 车辆ETC信息
func (VehicleEtcInfo) TableName() string {
	return "vehicle_etc_info"
}