/**
 * @Time: 2022/3/10 22:54
 * @Author: yt.yin
 */

package vetcmod

// EtcStatusErr etc状态异常
type EtcStatusErr struct {

	ID                string       `json:"id"               gorm:"column:id;primary_key;type:varchar(36)"`

	/** 创建时间 */
	CreateTime        string       `json:"createTime"       gorm:"column:create_time;index;type:varchar(20)"`

	/** 修改时间 */
	UpdateTime        string       `json:"updateTime"       gorm:"column:update_time;type:varchar(20)"`

	/** 设备编号 */
	DeviceId          string       `json:"deviceId"         gorm:"column:device_id;comment:设备编号;type:varchar(36);index;"`

	/** 入场时间 */
	EnterTime         string       `json:"enterTime"        gorm:"column:enter_time;comment:入场时间;type:varchar(20);"`

	/** OBU序号编码 此处放的obu Mac地址 并非合同序号 */
	ObuMac            string       `json:"obuMac"           gorm:"comment:obu mac;type:varchar(16)"`

	/** 发行商代码 */
	Issuer            string       `json:"issuer"           gorm:"column:issuer;comment:发行商代码;type:varchar(20);"`

	/** obu 合同序号 */
	ObuSerialNumber   string       `json:"obuSerialNumber"  gorm:"comment:obu合同序号;index;type:varchar(20)"`

	/** 错误代码 05 未插卡 06 已过有效期 */
	ErrCode           string       `json:"errCode"          gorm:"comment:错误代码 05 未插卡 06 已过有效期;type:varchar(4)"`
}
