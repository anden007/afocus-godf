package interface_core

import "github.com/anden007/afocus-godf/src/types"

type IMsgCenter interface {
	CreateAppChannelId(channelId string) (result string, err error)
	OpenChannel(channelIds []string, onReceiveMsg func(*types.MsgContent)) (err error)
	CloseChannel(channelIds []string) (err error)
	SendMsg(channelId, msgType, msg string) (err error)
}
