package third_party

import (
	"embed"
)

//go:embed swagger/*
var OpenAPI embed.FS
