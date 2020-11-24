package service_manage

import (
	"time"

	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_manage"
	"github.com/anden007/afocus-godf/src/repositories/repository_manage"
	"github.com/anden007/afocus-godf/src/services/service_core"

	"github.com/google/uuid"
)

type MessageService struct {
	lruCache             *service_core.LruCache
	message_repo         *repository_manage.MessageRepository
	message_send_service *MessageSendService
	user_repo            *repository_manage.UserRepository
}

func NewMessageService(lruCache *service_core.LruCache, message_repo *repository_manage.MessageRepository, message_send_service *MessageSendService, user_repo *repository_manage.UserRepository) *MessageService {
	instance := &MessageService{
		lruCache:             lruCache,
		message_repo:         message_repo,
		message_send_service: message_send_service,
		user_repo:            user_repo,
	}
	return instance
}

func (m *MessageService) GetById(id uuid.UUID) (result model_manage.Message, err error) {
	result, err = m.message_repo.GetById(id)
	return
}

func (m *MessageService) GetAll() (result []model_manage.Message, err error) {
	result, err = m.message_repo.GetAll()
	return
}

func (m *MessageService) Add(entity model_manage.Message) (err error) {
	var userList []model_manage.User
	if entity.Range == 0 {
		// 全体广播
		userList, err = m.user_repo.GetAll()
	} else {
		userList, err = m.user_repo.GetByIds(entity.UserIds)
	}
	err = m.message_repo.Add(entity)
	if err == nil {
		for _, user := range userList {
			m.message_send_service.Add(model_manage.MessageSend{
				Id:               lib.NewGuid(),
				MessageId:        entity.Id,
				ReceiverId:       user.Id,
				ReceiverNickName: user.NickName,
				Status:           0,
				CreateTime:       time.Now(),
			})
		}
	}
	return
}

func (m *MessageService) Edit(entity model_manage.Message) (err error) {
	err = m.message_repo.Edit(entity)
	return
}

func (m *MessageService) DelByIds(ids []uuid.UUID) (err error) {
	err = m.message_repo.DelByIds(ids)
	return
}

func (m *MessageService) Updates(id uuid.UUID, fieldValues map[string]interface{}) (err error) {
	err = m.message_repo.Updates(id, fieldValues)
	return
}

func (m *MessageService) GetByCondition(condition map[string]string) (result []model_manage.Message, total int64, err error) {
	result, total, err = m.message_repo.GetByCondition(condition)
	return
}
