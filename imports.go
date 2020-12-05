package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	"strconv"
)

// importPaths extracts all the imports of the proto files and assign them
// unique alias for imports
func (m *Module) importPaths(file pgs.File) (path2Alias, alias2Path map[string]string) {
	path2Alias = map[string]string{
		"context":                                       "context",
		"google.golang.org/grpc":                        "grpc",
		"google.golang.org/grpc/codes":                  "codes",
		"google.golang.org/grpc/status":                 "status",
		"github.com/Shivam010/protoc-gen-redact/redact": "redact",
	}
	alias2Path = map[string]string{
		"context": "context",
		"grpc":    "google.golang.org/grpc",
		"codes":   "google.golang.org/grpc/codes",
		"status":  "google.golang.org/grpc/status",
		"redact":  "github.com/Shivam010/protoc-gen-redact/redact",
	}

	self := m.ctx.ImportPath(file).String()
	for _, imp := range file.Imports() {
		path := m.ctx.ImportPath(imp).String()
		if self == path {
			continue
		}
		if _, ok := path2Alias[path]; ok {
			// already exist
			continue
		}
		alias := m.ctx.PackageName(imp).String()
		_, ok := alias2Path[alias]
		cnt := 0
		for ok {
			cnt++
			_, ok = alias2Path[alias+strconv.Itoa(cnt)]
		}
		if cnt > 0 {
			alias = alias + strconv.Itoa(cnt)
		}
		path2Alias[path] = alias
		alias2Path[alias] = path
	}
	return
}

// references lists all the import-references from different proto packages
// to suppress any unused import errors
func (m *Module) references(file pgs.File, nameWithAlias func(n pgs.Entity) string) []string {
	imports := file.Imports()
	list := make([]string, 0, len(imports)+5)
	list = append(list, "grpc.Server",
		"context.Context",
		"redact.Redactor",
		"codes.Code",
		"status.Status",
	)

	self := m.ctx.ImportPath(file)
	for _, imp := range imports {
		if m.ctx.ImportPath(imp) == self {
			continue
		}
		// messages
		msgL := imp.AllMessages()
		if len(msgL) > 0 {
			list = append(list, nameWithAlias(msgL[0]))
			continue
		}
		// or enums
		enmL := imp.AllEnums()
		if len(enmL) > 0 {
			list = append(list, nameWithAlias(enmL[0]))
			continue
		}
		// or services
		srvL := imp.Services()
		if len(srvL) > 0 {
			list = append(list, nameWithAlias(srvL[0]))
			continue
		}
	}
	return list
}
