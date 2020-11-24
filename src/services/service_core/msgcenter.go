package service_core

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/types"

	jsoniter "github.com/json-iterator/go"
	jsontime "github.com/liamylian/jsontime/v2/v2"
	"github.com/nats-io/nats.go"
)

type Channel struct {
	id           string
	appName      string
	onReceiveMsg func(content *types.MsgContent)
}

type MsgCenter struct {
	enable         bool
	appName        string
	natsConnection *nats.Conn
	channels       map[string]*Channel
	json           jsoniter.API
}

func NewMsgCenter() *MsgCenter {
	var err error
	instance := new(MsgCenter)
	loadTime := time.Now()
	appName := os.Getenv("app-name")
	server := os.Getenv("nats-server")
	token := os.Getenv("nats-token")
	enable := os.Getenv("nats-enable")
	if strings.EqualFold("true", enable) {
		instance.enable = true
		instance.appName = appName
		instance.natsConnection, err = nats.Connect(server, nats.Token(token))
		if err != nil {
			panic(fmt.Sprintf("Connect Nats Server has error:%s", err.Error()))
		} else {
			fmt.Print("MsgCenter initialization success.\n")
		}
		instance.channels = make(map[string]*Channel)
		instance.json = jsontime.ConfigWithCustomTimeFormat
		if lib.IS_DEV_MODE {
			fmt.Println("> Service: MsgCenter loaded, ", time.Since(loadTime))
		}
	} else {
		fmt.Printf("> Service: MsgCenter is disabled. if you need enable it,please set 'nats-enable' to 'true' in .env file.\n")
	}
	return instance
}

func (m *MsgCenter) onReceiveMsg(msg *nats.Msg) {
	var content *types.MsgContent
	err := m.json.Unmarshal(msg.Data, &content)
	if err == nil && content != nil {
		m.channels[msg.Subject].onReceiveMsg(content)
	}
}

func (m *MsgCenter) OpenChannel(channelIds []string, onReceiveMsg func(*types.MsgContent)) (err error) {
	if m.enable {
		for _, channelId := range channelIds {
			appChannelId, _ := m.CreateAppChannelId(channelId)
			if _, exists := m.channels[appChannelId]; !exists {
				channel := new(Channel)
				channel.appName = m.appName
				channel.id = channelId
				channel.onReceiveMsg = onReceiveMsg
				m.channels[appChannelId] = channel
				_, err = m.natsConnection.Subscribe(appChannelId, m.onReceiveMsg)
			}
		}
	}
	return
}

func (m *MsgCenter) CloseChannel(channelIds []string) (err error) {
	err = nil
	if m.enable {
		for _, channelId := range channelIds {
			appChannelId, _ := m.CreateAppChannelId(channelId)
			delete(m.channels, appChannelId)
		}
	}
	return
}

func (m *MsgCenter) SendMsg(channelId, msgType, msg string) (err error) {
	if m.enable {
		appChannelId, _ := m.CreateAppChannelId(channelId)
		msgContent := types.MsgContent{
			ChannelId: appChannelId,
			MsgType:   msgType,
			Data:      msg,
		}
		data, _ := m.json.Marshal(msgContent)
		err = m.natsConnection.Publish(appChannelId, data)
	} else {
		lib.LogCenter().Warn("警告:nats组件尚未启用，客户端将无法收到消息，如需启用，请在env配置文件中将'nats-enable'设置为'true'")
	}
	return
}

func (m *MsgCenter) CreateAppChannelId(channelId string) (result string, err error) {
	//此方法可结合SHA等加密算法提供更安全的频道Id生成逻辑
	err = nil
	result = lib.SHA1(fmt.Sprintf("%s_%s", m.appName, channelId))
	return
}
