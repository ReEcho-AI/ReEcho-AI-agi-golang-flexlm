package model

import (
	"github.com/google/uuid"
)

type DocumentID uuid.UUID

func NewDocumentID() DocumentID {
	return DocumentID(uuid.New())
}

type Document struct {
	ID       DocumentID
	Text     string
	Metadata *DocumentMetadata
}

type Source string

type DocumentMetadata struct {
}

type DocumentChunkMetadata struct {
	*DocumentMetadata
	DocumentID DocumentID
}

type DocumentChunkID uuid.UUID

func NewDocumentChu