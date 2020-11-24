package service_manage

import (
	"fmt"

	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/repositories/repository_manage"
	"github.com/anden007/afocus-godf/src/services/service_core"

	"github.com/google/uuid"
)

type MessageSendService struct {
	lruCache      *service_core.LruCache
	repo          *repository_manage.MessageSendRepository
	messageCenter *service_core.MsgCenter
}

func NewMessageSendService(lruCache *service_core.LruCache, repo *repository_manage.MessageSendRepository, msgCenter *service_core.MsgCenter) *MessageSendService {
	instance := &MessageSendService{
		lruCache:      lruCache,
		repo:          repo,
		messageCenter: msgCenter,
	}
	return instance
}

func (m *MessageSendService) GetMsgCenterChannelId(userId uuid.UUID) (result string, err error) {
	result, err = m.messageCenter.CreateAppChannelId(fmt.Sprintf("%s_nats_cid", userId.String()))
	return
}

func (m *MessageSendService) GetById(id uuid.UUID) (result model_manage.MessageSend, err error) {
	result, err = m.repo.GetById(id)
	return
}

func (m *MessageSendService) GetAll() (result []model_manage.MessageSend, err error) {
	result, err = m.repo.GetAll()
	return
}

func (m *MessageSendService) Add(entity model_manage.MessageSend) (err error) {
	err = m.repo.Add(entity)
	if err == nil {
		channelId := fmt.Sprintf("%s_nats_cid", entity.ReceiverId.String())
		m.messageCenter.SendMsg(channelId, "new_msg", "1")
	}
	return
}

func (m *MessageSendService) Edit(entity model_manage.MessageSend) (err error) {
	err = m.repo.Edit(entity)
	return
}

func (m *MessageSendService) DelByIds(ids []uuid.UUID) (err error) {
	err = m.repo.DelByIds(ids)
	return
}

func (m *MessageSendService) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	err = m.repo.Updates(id, fieldValues)
	return
}

func (m *MessageSendService) GetByCondition(condition map[string]string) (result []model_manage.MessageSend, total int64, err error) {
	result, total, err = m.repo.GetByCondition(condition)
	return
}
