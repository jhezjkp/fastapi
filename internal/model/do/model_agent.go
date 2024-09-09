package do

import "github.com/gogf/gf/v2/util/gmeta"

const (
	MODEL_AGENT_COLLECTION = "model_agent"
)

type ModelAgent struct {
	gmeta.Meta         `collection:"model_agent" bson:"-"`
	Corp               string `bson:"corp,omitempty"`                 // 公司
	Name               string `bson:"name,omitempty"`                 // 模型代理名称
	BaseUrl            string `bson:"base_url,omitempty"`             // 模型代理地址
	Path               string `bson:"path,omitempty"`                 // 模型代理地址路径
	Weight             int    `bson:"weight,omitempty"`               // 权重
	Remark             string `bson:"remark,omitempty"`               // 备注
	Status             int    `bson:"status,omitempty"`               // 状态[1:正常, 2:禁用, -1:删除]
	IsAutoDisabled     bool   `bson:"is_auto_disabled,omitempty"`     // 是否自动禁用
	AutoDisabledReason string `bson:"auto_disabled_reason,omitempty"` // 自动禁用原因
	Creator            string `bson:"creator,omitempty"`              // 创建人
	Updater            string `bson:"updater,omitempty"`              // 更新人
	CreatedAt          int64  `bson:"created_at,omitempty"`           // 创建时间
	UpdatedAt          int64  `bson:"updated_at,omitempty"`           // 更新时间
}
