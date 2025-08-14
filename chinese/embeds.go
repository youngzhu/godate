package chinese

import "embed"

/*
项目结构：
project/
├── go.mod
├── pkg/
│   ├── utils.go
│   ├── embeds.go
│   ├── embeds_test.go
│   └── data/
│       └── config.json
└── cmd/
    └── main.go
*/
//go:embed data/timor/*
var LocalResources embed.FS
