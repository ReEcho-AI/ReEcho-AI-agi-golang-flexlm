package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"

	"github.com/zawakin/lightweight-agi/datastore"
	"github.com/zawakin/lightweight-agi/datastore/providers/inmemory"
	"github.com/zawakin/lightweight-agi/examples/agi"
	"github.com/zawakin/lightweight-agi/examples/prompts"
	"github.com/zawakin/lightweight-agi/llmclient/provider"
	"github.com/zawakin/lightweight-agi/prompt"
)

var (
	// defaultOpenAICompletionModel = openai.GPT3Dot5Turbo
	defaultOpenAI