package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"

	"github.com/zawakin/lightweight-agi/datastore"
	"github.com/zawakin/lightweight-a