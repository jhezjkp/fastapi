package consts

const (
	SESSION_USER = "session_user"
	SESSION_APP  = "session_app"
	SESSION_KEY  = "session_key"

	HOST_KEY               = "host"
	USER_ID_KEY            = "user_id"
	APP_ID_KEY             = "app_id"
	SECRET_KEY             = "sk"
	APP_IS_LIMIT_QUOTA_KEY = "app_is_limit_quota"
	KEY_IS_LIMIT_QUOTA_KEY = "key_is_limit_quota"

	CORP_OPENAI     = "OpenAI"
	CORP_AZURE      = "Azure"
	CORP_BAIDU      = "Baidu"
	CORP_XFYUN      = "Xfyun"
	CORP_ALIYUN     = "Aliyun"
	CORP_ZHIPUAI    = "ZhipuAI"
	CORP_GOOGLE     = "Google"
	CORP_DEEPSEEK   = "DeepSeek"
	CORP_MIDJOURNEY = "Midjourney"
	CORP_GCP_CLAUDE = "GCPClaude"

	ROLE_SYSTEM    = "system"
	ROLE_USER      = "user"
	ROLE_ASSISTANT = "assistant"
	ROLE_FUNCTION  = "function"
	ROLE_TOOL      = "tool"

	GPT_PREFIX     = "gpt-"
	DEFAULT_MODEL  = "gpt-3.5-turbo"
	QUOTA_USD_UNIT = 500000.0 // $1 = 50万tokens
)
