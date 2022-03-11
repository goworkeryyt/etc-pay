/**
 * @Time: 2022/3/11 10:42
 * @Author: yt.yin
 */

package issuemod

// BillAck 账单确认
type BillAck struct {

	/** 节点服务器编号 */
	DeviceId            string         `json:"deviceId"`

	/** 支付账单id */
	OutId               string         `json:"outId"`

	/** 动作类型 0 是新扣费  -1 是删除 */
	Action              int            `json:"action"`
}
