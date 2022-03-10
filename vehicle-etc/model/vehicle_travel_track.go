/**
 * @Time: 2022/3/10 23:26
 * @Author: yt.yin
 */

package vetcmod

// VehicleTravelTrack 车辆行驶轨迹
type VehicleTravelTrack struct {

	ID                string       `json:"id"               gorm:"column:id;comment:主键;primary_key;type:varchar(36)" `

	/** 入库时间 */
	CreateTime        string       `json:"createTime"       gorm:"column:create_time;comment:入库时间;type:varchar(20)"`

	/** 设备编号 */
	DeviceId          string       `json:"deviceId"         gorm:"column:device_id;comment:设备编号;type:varchar(36);index;"`

	/** 位置location */
	Location          string       `json:"location"         gorm:"-"`

	/** 车辆经过时间 */
	PassTime          string       `json:"passTime"         gorm:"column:pass_time;comment:入场时间;type:varchar(20);"`

	/** 车辆经过日期 -- 主要是数据分表用 */
	PassDate          string       `json:"passDate"         gorm:"column:pass_date;comment:车辆经过日期;index;type:varchar(10)"`

	/** 车牌号码 */
	Vlp               string       `json:"vlp"              gorm:"column:vlp;comment:车牌;index;type:varchar(36)"`

	/** 车牌颜色编码 */
	Vlpc              string       `json:"vlpc"             gorm:"column:vlpc;comment:车牌颜色;type:varchar(4)"`

	/** 车辆类型编码 */
	Vc                string       `json:"vc"               gorm:"column:vc;comment:车辆类型;type:varchar(4)"`

	/** 卡上余额 */
	Balance           int64        `json:"balance"          gorm:"column:balance;comment:余额;"`

	/** OBU序号编码 此处放的obu Mac地址 并非合同序号 */
	ObuMac            string       `json:"obuMac"           gorm:"column:obu_mac;comment:obu mac;index;type:varchar(16);"`

	/** ETC 卡号 */
	CardNo            string       `json:"cardNo"           gorm:"column:card_no;comment:ETC卡号;index;type:varchar(20);"`

	/** 卡片类型：22-储值卡，23-记账卡 */
	CardType          int          `json:"cardType"         gorm:"column:card_type;comment:卡片类型：22-储值卡，23-记账卡;type:varchar(4)"`

	/** 发行商代码 */
	Issuer            string       `json:"issuer"           gorm:"column:issuer;comment:发行商代码;type:varchar(20);"`

	/** obu 合同序号 */
	ObuSerialNumber   string       `json:"obuSerialNumber"  gorm:"column:obu_serial_number;comment:obu合同序号;type:varchar(20);index"`
}

// TableName 车辆ETC信息
func (VehicleTravelTrack) TableName() string {
	return "vehicle_travel_track"
}