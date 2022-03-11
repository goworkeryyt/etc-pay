/**
 * @Time: 2022/3/11 10:33
 * @Author: yt.yin
 */

package issuemod

// BillContent 节点服务器账单内容
type BillContent struct {

	/** 出场记录Id */
	OutId               string       `json:"outId"`

	/** 动作类型 0 是新扣费  -1 是删除 */
	Action              int          `json:"action"`

	/** 车牌号码 */
	Vlp                 string       `json:"vlp"`

	/** 入场时间 yyyyMMddHHmmss */
	EnterTime           string       `json:"enterTime"`

	/** 出场时间 yyyyMMddHHmmss*/
	ExitTime            string       `json:"exitTime"`

	/** 此处放的obu Mac地址 并非合同序号 */
	ObuMac              string       `json:"obuMac"`

	/** 实扣金额 */
	ConsumeMoney        int64        `json:"consumeMoney"`
}
