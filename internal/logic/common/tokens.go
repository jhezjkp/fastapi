package common

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	sdkm "github.com/iimeta/fastapi-sdk/model"
	"github.com/iimeta/fastapi-sdk/tiktoken"
	"github.com/iimeta/fastapi/internal/consts"
	"github.com/iimeta/fastapi/internal/model"
	mcommon "github.com/iimeta/fastapi/internal/model/common"
	"github.com/iimeta/fastapi/utility/logger"
)

func GetPromptTokens(ctx context.Context, model string, messages []sdkm.ChatCompletionMessage) int {

	promptTime := gtime.TimestampMilli()

	promptTokens, err := tiktoken.NumTokensFromMessages(model, messages)
	if err != nil {
		logger.Errorf(ctx, "GetPromptTokens NumTokensFromMessages model: %s, messages: %s, error: %v", model, gjson.MustEncodeString(messages), err)
		if promptTokens, err = tiktoken.NumTokensFromMessages(consts.DEFAULT_MODEL, messages); err != nil {
			logger.Errorf(ctx, "GetPromptTokens NumTokensFromMessages model: %s, messages: %s, error: %v", consts.DEFAULT_MODEL, gjson.MustEncodeString(messages), err)
		}
	}
	logger.Debugf(ctx, "GetPromptTokens NumTokensFromMessages model: %s, len(messages): %d, promptTokens: %d, time: %d", model, len(gjson.MustEncodeString(messages)), promptTokens, gtime.TimestampMilli()-promptTime)

	return promptTokens
}

func GetCompletionTokens(ctx context.Context, model, completion string) int {

	completionTime := gtime.TimestampMilli()
	completionTokens, err := tiktoken.NumTokensFromString(model, completion)
	if err != nil {
		logger.Errorf(ctx, "GetCompletionTokens NumTokensFromString model: %s, completion: %s, error: %v", model, completion, err)
		if completionTokens, err = tiktoken.NumTokensFromString(consts.DEFAULT_MODEL, completion); err != nil {
			logger.Errorf(ctx, "GetCompletionTokens NumTokensFromString model: %s, completion: %s, error: %v", consts.DEFAULT_MODEL, completion, err)
		}
	}
	logger.Debugf(ctx, "GetCompletionTokens NumTokensFromString model: %s, len(completion): %d, completionTokens: %d, time: %d", model, len(completion), completionTokens, gtime.TimestampMilli()-completionTime)

	return completionTokens
}

func GetMultimodalTokens(ctx context.Context, model string, multiContent []interface{}, reqModel *model.Model) (textTokens, imageTokens int) {

	for _, value := range multiContent {

		content := value.(map[string]interface{})

		if content["type"] == "image_url" {

			imageUrl := content["image_url"].(map[string]interface{})
			detail := imageUrl["detail"]

			var imageQuota mcommon.ImageQuota
			for _, quota := range reqModel.MultimodalQuota.ImageQuotas {

				if quota.Mode == detail {
					imageQuota = quota
					break
				}

				if quota.IsDefault {
					imageQuota = quota
				}
			}

			imageTokens += imageQuota.FixedQuota

		} else {
			contentTime := gtime.TimestampMilli()
			tokens, err := tiktoken.NumTokensFromString(model, gconv.String(content))
			if err != nil {
				logger.Errorf(ctx, "GetMultimodalQuota NumTokensFromString model: %s, content: %s, error: %v", model, gconv.String(content), err)
				if tokens, err = tiktoken.NumTokensFromString(consts.DEFAULT_MODEL, gconv.String(content)); err != nil {
					logger.Errorf(ctx, "GetMultimodalQuota NumTokensFromString model: %s, content: %s, error: %v", consts.DEFAULT_MODEL, gconv.String(content), err)
				}
			}
			textTokens += tokens
			logger.Debugf(ctx, "GetMultimodalQuota NumTokensFromString model: %s, len(content): %d, tokens: %d, time: %d", model, len(gconv.String(content)), tokens, gtime.TimestampMilli()-contentTime)
		}
	}

	return textTokens, imageTokens
}
