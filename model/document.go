package model

import (
	"github.com/google/uuid"
)

type DocumentID uuid.UUID

func NewDocumentID() DocumentID {
	return DocumentID(uuid.New())
}

type 