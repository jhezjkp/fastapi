package common

type PresetConfig struct {
	IsSupportSystemRole bool   `bson:"is_support_system_role,omitempty" json:"is_support_system_role,omitempty"` // 是否支持system角色
	SystemRolePrompt    string `bson:"system_role_prompt,omitempty"     json:"system_role_prompt,omitempty"`     // system角色预设提示词
	MinTokens           int    `bson:"min_tokens,omitempty"             json:"min_tokens,omitempty"`             // max_tokens取值的最小值
	MaxTokens           int    `bson:"max_tokens,omitempty"             json:"max_tokens,omitempty"`             // max_tokens取值的最大值
}

type TextQuota struct {
	BillingMethod   int     `bson:"billing_method,omitempty"   json:"billing_method,omitempty"`         // 计费方式[1:倍率, 2:固定额度]
	PromptRatio     float64 `bson:"prompt_ratio,omitempty"     json:"prompt_ratio,omitempty"     d:"1"` // 提示倍率(提问倍率)
	CompletionRatio float64 `bson:"completion_ratio,omitempty" json:"completion_ratio,omitempty" d:"1"` // 补全倍率(回答倍率)
	FixedQuota      int     `bson:"fixed_quota,omitempty"      json:"fixed_quota,omitempty"`            // 固定额度
}

type ImageQuota struct {
	Width      int    `bson:"width,omitempty"       json:"width,omitempty"`       // 宽度
	Height     int    `bson:"height,omitempty"      json:"height,omitempty"`      // 高度
	Mode       string `bson:"mode,omitempty"        json:"mode,omitempty"`        // 模式[low, high, auto]
	FixedQuota int    `bson:"fixed_quota,omitempty" json:"fixed_quota,omitempty"` // 固定额度
	IsDefault  bool   `bson:"is_default,omitempty"  json:"is_default,omitempty"`  // 是否默认选项
}

type AudioQuota struct {
	BillingMethod   int     `bson:"billing_method,omitempty"   json:"billing_method,omitempty"`         // 计费方式[1:倍率, 2:固定额度]
	PromptRatio     float64 `bson:"prompt_ratio,omitempty"     json:"prompt_ratio,omitempty"     d:"1"` // 提示倍率(提问倍率)
	CompletionRatio float64 `bson:"completion_ratio,omitempty" json:"completion_ratio,omitempty" d:"1"` // 补全倍率(回答倍率)
	FixedQuota      int     `bson:"fixed_quota,omitempty"      json:"fixed_quota,omitempty"`            // 固定额度
}

type MultimodalQuota struct {
	TextQuota   TextQuota    `bson:"text_quota,omitempty"   json:"text_quota,omitempty"`   // 文本额度
	ImageQuotas []ImageQuota `bson:"image_quotas,omitempty" json:"image_quotas,omitempty"` // 图像额度
}

type RealtimeQuota struct {
	TextQuota  TextQuota  `bson:"text_quota,omitempty"  json:"text_quota,omitempty"`  // 文本额度
	AudioQuota AudioQuota `bson:"audio_quota,omitempty" json:"audio_quota,omitempty"` // 音频额度
	FixedQuota int        `bson:"fixed_quota,omitempty" json:"fixed_quota,omitempty"` // 固定额度
}

type MidjourneyQuota struct {
	Name       string `bson:"name,omitempty"        json:"name,omitempty"`        // 名称
	Action     string `bson:"action,omitempty"      json:"action,omitempty"`      // 动作[IMAGINE, UPSCALE, VARIATION, ZOOM, PAN, DESCRIBE, BLEND, SHORTEN, SWAP_FACE]
	Path       string `bson:"path,omitempty"        json:"path,omitempty"`        // 路径
	FixedQuota int    `bson:"fixed_quota,omitempty" json:"fixed_quota,omitempty"` // 固定额度
}

type ForwardConfig struct {
	ForwardRule   int      `bson:"forward_rule,omitempty"   json:"forward_rule,omitempty"`   // 转发规则[1:全部转发, 2:按关键字, 3:内容长度]
	MatchRule     []int    `bson:"match_rule,omitempty"     json:"match_rule,omitempty"`     // 转发规则为2时的匹配规则[1:智能匹配, 2:正则匹配]
	TargetModel   string   `bson:"target_model,omitempty"   json:"target_model,omitempty"`   // 转发规则为1和3时的目标模型
	DecisionModel string   `bson:"decision_model,omitempty" json:"decision_model,omitempty"` // 转发规则为2时并且匹配规则为1时的判定模型
	Keywords      []string `bson:"keywords,omitempty"       json:"keywords,omitempty"`       // 转发规则为2时的关键字
	TargetModels  []string `bson:"target_models,omitempty"  json:"target_models,omitempty"`  // 转发规则为2时的目标模型
	ContentLength int      `bson:"content_length,omitempty" json:"content_length,omitempty"` // 转发规则为3时的内容长度
}

type FallbackConfig struct {
	FallbackModel     string `bson:"fallback_model,omitempty"      json:"fallback_model,omitempty"`      // 后备模型
	FallbackModelName string `bson:"fallback_model_name,omitempty" json:"fallback_model_name,omitempty"` // 后备模型名称
}

type Message struct {
	Role    string `bson:"role,omitempty"    json:"role,omitempty"`    // 角色
	Content string `bson:"content,omitempty" json:"content,omitempty"` // 内容
}

type Retry struct {
	IsRetry    bool   `bson:"is_retry,omitempty"    json:"is_retry,omitempty"`    // 是否重试
	RetryCount int    `bson:"retry_count,omitempty" json:"retry_count,omitempty"` // 重试次数
	ErrMsg     string `bson:"err_msg,omitempty"     json:"err_msg,omitempty"`     // 错误信息
}

type ImageData struct {
	URL           string `bson:"url,omitempty"`
	B64JSON       string `bson:"b64_json,omitempty"`
	RevisedPrompt string `bson:"revised_prompt,omitempty"`
}
