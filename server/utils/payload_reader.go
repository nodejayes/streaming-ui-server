package utils

import "github.com/nodejayes/streaming-ui-server/server/ui/types"

func ReadPayload[T any](action types.Action) T {
	var convertedPayload T
	convertedPayload, _ = action.GetPayload().(T)
	return convertedPayload
}
