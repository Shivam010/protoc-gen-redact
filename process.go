package main

import (
	"github.com/Shivam010/protoc-gen-redact/redact"
	pgs "github.com/lyft/protoc-gen-star"
	"google.golang.org/grpc/codes"
	"strings"
)

const (
	// invalidCodeErrMsg: 0 <= grpc status code <= 16
	// Reference: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
	invalidCodeErrMsg = `InternalStatusCode must be a valid GRPC code(0 <= code <= 16)`
	// defaultErrMsg: for the service method/rpc redaction
	defaultErrMsg = `Permission Denied. Method: "%service%.%method%" has been redacted`
	// error message format specifiers
	specifierMethod  = "%method%"
	specifierService = "%service%"
)

// Process processes the file and adds its generated code into Module.Artifacts
func (m *Module) Process(file pgs.File) {
	// check file option: FileSkip
	fileSkip := false
	m.must(file.Extension(redact.E_FileSkip, &fileSkip))
	if fileSkip {
		return
	}

	// imports and their aliases
	path2Alias, alias2Path := m.importPaths(file)
	nameWithAlias := func(n pgs.Entity) string {
		imp := m.ctx.ImportPath(n).String()
		name := m.ctx.Name(n).String()
		if alias := path2Alias[imp]; alias != "" {
			name = alias + "." + name
		}
		return name
	}

	data := &ProtoFileData{
		Source:     file.Name().String(),
		Package:    m.ctx.PackageName(file).String(),
		Imports:    alias2Path,
		References: m.references(file, nameWithAlias),
		Services:   make([]*ServiceData, 0, len(file.Services())),
		Messages:   make([]*MessageData, 0, len(file.AllMessages())),
	}

	// all services
	for _, srv := range file.Services() {
		data.Services = append(data.Services, m.processService(srv, nameWithAlias))
	}

	// all messages
	for _, msg := range file.AllMessages() {
		data.Messages = append(data.Messages, m.processMessage(msg, nameWithAlias, true))
	}

	// render file in the template
	name := m.ctx.OutputPath(file).SetExt(".redact.go")
	m.AddGeneratorTemplateFile(name.String(), m.tmpl, data)
}

// processService extracts all pgs.Service and their pgs.Method(s) information and
// structures them into ServiceData
func (m *Module) processService(srv pgs.Service, nameWithAlias func(n pgs.Entity) string) *ServiceData {
	srvData := &ServiceData{
		Name:    m.ctx.Name(srv).String(),
		Methods: make([]*MethodData, 0, len(srv.Methods())),
	}

	// check service option: ServiceSkip
	srvSkip := false
	m.must(srv.Extension(redact.E_ServiceSkip, &srvSkip))
	if srvSkip {
		srvData.Skip = true
		// continue
	}

	// check internal service options
	srvInternal := false
	m.must(srv.Extension(redact.E_InternalService, &srvInternal))
	srvCode := uint32(codes.PermissionDenied) // default code
	if !m.must(srv.Extension(redact.E_InternalServiceCode, &srvCode)) {
		srvCode = uint32(codes.PermissionDenied)
	}
	if srvCode > uint32(codes.Unauthenticated) { // 16
		m.Fail(invalidCodeErrMsg)
	}
	srvErrMsg := ""
	if !m.must(srv.Extension(redact.E_InternalServiceErrMessage, &srvErrMsg)) {
		srvErrMsg = defaultErrMsg
	}

	// methods
	for _, meth := range srv.Methods() {
		in := meth.Input()
		out := meth.Output()
		methData := &MethodData{
			Name:   m.ctx.Name(meth).String(),
			Input:  nameWithAlias(in),
			Output: m.processMessage(out, nameWithAlias),
		}
		srvData.Methods = append(srvData.Methods, methData)

		// check method skip options
		methSkip := false
		m.must(meth.Extension(redact.E_MethodSkip, &methSkip))
		if methSkip || srvSkip {
			methData.Skip = true
			continue
		}

		methInternal := false
		m.must(meth.Extension(redact.E_InternalMethod, &methInternal))
		methCode := srvCode // serviceCode
		if !m.must(meth.Extension(redact.E_InternalMethodCode, &methCode)) {
			methCode = srvCode
		}
		if methCode > uint32(codes.Unauthenticated) { // 16
			m.Fail(invalidCodeErrMsg)
		}
		methErrMsg := srvErrMsg
		if !m.must(meth.Extension(redact.E_InternalMethodErrMessage, &methErrMsg)) {
			methErrMsg = srvErrMsg
		}

		// apply format specifiers
		methErrMsg = strings.ReplaceAll(methErrMsg, specifierMethod, methData.Name)
		methErrMsg = strings.ReplaceAll(methErrMsg, specifierService, srvData.Name)

		methData.ErrMessage = "`" + methErrMsg + "`"
		methData.StatusCode = codes.Code(methCode).String()
		methData.Internal = srvInternal || methInternal
	}
	return srvData
}

// processMessage extracts all pgs.Message and their pgs.Field(s) information and
// structures them into MessageData
func (m *Module) processMessage(msg pgs.Message, nameWithAlias func(n pgs.Entity) string, wantFields ...bool) *MessageData {
	msgData := &MessageData{
		Name:      m.ctx.Name(msg).String(),
		WithAlias: nameWithAlias(msg),
		Fields:    make([]*FieldData, 0, len(msg.Fields())*2),
	}

	// check message ignore options
	msgData.Ignore = false
	m.must(msg.Extension(redact.E_Ignored, &msgData.Ignore))
	if msgData.Ignore {
		return msgData
	}

	// check message nil options
	msgData.ToNil = false
	m.must(msg.Extension(redact.E_Nil, &msgData.ToNil))

	// check message empty options
	msgData.ToEmpty = false
	m.must(msg.Extension(redact.E_Empty, &msgData.ToEmpty))

	if len(wantFields) > 0 {
		for _, field := range msg.Fields() {
			msgData.Fields = append(msgData.Fields, m.processFields(field, nameWithAlias))
		}
	}
	return msgData
}
