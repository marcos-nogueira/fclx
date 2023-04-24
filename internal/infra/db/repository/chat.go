package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/marcos-nogueira/fclx/chatservice/internal/domain/entity"
	"github.com/marcos-nogueira/fclx/chatservice/internal/infra/db"
)

type ChatRepositoryMySQL struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewChatRepositoryMySQL(dbt *sql.DB) *ChatRepositoryMySQL {
	return &ChatRepositoryMySQL{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (r *ChatRepositoryMySQL) CreateChat(ctx context.Context, chat *entity.Chat) error {

	err := r.Queries.CreateChat(
		ctx,
		db.CreateChatParams{
			ID:               chat.ID,
			UserID:           chat.UserID,
			InitialMessageID: chat.InitialSystemMessage.Contenct,
			Status:           chat.Status,
			TokenUsage:       int32(chat.TokenUsage),
			Model:            chat.Config.Model.Name,
			ModelMaxTokens:   int32(chat.Config.MaxTokens),
			Temperature:      float64(chat.Config.Temperature),
			TopP:             float64(chat.Config.TopP),
			N:                int32(chat.Config.N),
			Stop:             chat.Config.Stop[0],
			MaxTokens:        int32(chat.Config.MaxTokens),
			PresencePenalty:  float64(chat.Config.PresencePanalty),
			FrequencyPenalty: float64(chat.Config.FrequencyPanalty),
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	)
	if err != nil {
		return err
	}

	err = r.Queries.AddMessage(
		ctx,
		db.AddMessageParams{
			ID:        chat.InitialSystemMessage.ID,
			ChatID:    chat.ID,
			Content:   chat.InitialSystemMessage.Contenct,
			Role:      chat.InitialSystemMessage.Role,
			Tokens:    int32(chat.InitialSystemMessage.Tokens),
			CreatedAt: chat.InitialSystemMessage.CreateAt,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
