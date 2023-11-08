package livereplacer

import "embed"

//go:generate npm run build
//go:embed all:lib/*
var Files embed.FS
