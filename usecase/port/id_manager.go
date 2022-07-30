package port

import "github.com/mkaiho/go-chat-sample-api/entity"

type IDManager interface {
	Generate() entity.ID
}
