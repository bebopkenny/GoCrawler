package main

import (
	"net/http"

	"github.com/bebopkenny/GoCrawler/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)