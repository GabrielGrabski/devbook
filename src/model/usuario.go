package model

import "time"

type Usuario struct {
	Id        int       `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
