package main

import (
	"fmt"
	"github.com/Shivam010/protoc-gen-redact/redact"
	pgs "github.com/lyft/protoc-gen-star"
)

// processFields extracts each fields information
func (m *Module) processFields(field pgs.Field, nameWithAlias func(n pgs.Entity) string) *FieldData {
	typ := field.Type()
	flData := &FieldData{
		Name:       m.ctx.Name(field).String(),
		IsMap:      typ.IsMap(),
		IsRepeated: typ.IsRepeated(),
		IsMessage:  typ.IsEmbed(),
	}
	em := typ.Embed()
	if em == nil {
		if ele := typ.Element(); ele != nil {
			em = ele.Embed()
		}
	}
	// embed message
	if em != nil {
		flData.EmbedMessageName = m.ctx.Name(em).String()
		flData.EmbedMessageNameWithAlias = nameWithAlias(em)
	}

	_redact, fieldRules := false, &redact.FieldRules{}
	ok := m.must(field.Extension(redact.E_Redact, &_redact))
	ok = ok || m.must(field.Extension(redact.E_Custom, &fieldRules))

	// safe field: no option is defined
	if !ok {
		return flData
	}

	// check for custom field rules
	if fieldRules == nil || fieldRules.Values == nil {
		// no field rules
		if !_redact {
			// and redaction is also denied
			return flData
		}
		// default rules will be used
		flData.Redact = true
		flData.RedactionValue = RedactionDefaults(typ.ProtoType(), typ.IsRepeated() || typ.IsMap())
		if typ.IsEmbed() {
			flData.NestedEmbedCall = true
		}
		return flData
	}

	// custom field rules are defined, hence prefill defaults
	flData.Redact = true
	flData.RedactionValue = RedactionDefaults(typ.ProtoType(), typ.IsRepeated() || typ.IsMap())
	// custom values
	m.redactedCustomValue(flData, field, fieldRules)
	return flData
}

func (m *Module) redactedCustomValue(flData *FieldData, field pgs.Field, fieldRules *redact.FieldRules) {
	typ := field.Type()
	// extract rule information
	info := m.RuleInformation(fieldRules)
	// match field types & rule types
	if info.ProtoType != 0 && info.ProtoType != typ.ProtoType() {
		m.failWithInvalidType(field)
		return // unreachable
	}
	if typ.ProtoLabel() == pgs.Repeated && info.ProtoLabel != pgs.Repeated {
		m.failWithInvalidType(field)
		return // unreachable
	}
	if info.ProtoType != pgs.MessageT && info.ProtoLabel != pgs.Repeated {
		// simple type fields
		flData.RedactionValue = fmt.Sprintf("%v", info.RedactionValue)
		return
	}

	// if message type
	if info.ProtoType == pgs.MessageT {
		rule := fieldRules.Values.(*redact.FieldRules_Message).Message
		// default value is nil
		flData.RedactionValue = `nil`
		if rule.Empty {
			//flData.RedactionValue = m.ctx.Type(field).String() + "{}"
			flData.RedactionValue = fmt.Sprintf("&%s{}", flData.EmbedMessageNameWithAlias)
			return
		}
		if rule.Nil {
			flData.RedactionValue = "nil"
			return
		}
		if rule.Skip {
			flData.EmbedSkip = true
			return
		}
		flData.NestedEmbedCall = true
		return
	}

	// else info.ProtoLabel == pgs.Repeated
	rule := fieldRules.Values.(*redact.FieldRules_Element).Element
	if rule.Empty {
		if flData.EmbedMessageNameWithAlias == "" {
			flData.RedactionValue = m.ctx.Type(field).String() + "{}"
			return
		}
		if flData.IsRepeated {
			flData.RedactionValue = fmt.Sprintf("[]*%s{}", flData.EmbedMessageNameWithAlias)
			return
		}
		// map type
		key := m.ctx.Type(field).Key().String()
		flData.RedactionValue = fmt.Sprintf("map[%s]*%s{}", key, flData.EmbedMessageNameWithAlias)
		return
	}
	if rule.Nested {
		// iterate over all items and redact with defaults
		flData.Iterate = true
		flData.RedactionValue = RedactionDefaults(typ.Element().ProtoType(), false)
		if typ.Element().IsEmbed() {
			flData.NestedEmbedCall = true
		}
		return
	}
	if rules := rule.Item; rules != nil && rules.Values != nil {
		if rules.GetElement() != nil {
			m.Failf("Option provided to %s is invalid, '(redact.custom).element.item.element...' is not possible",
				field.FullyQualifiedName())
		}
		info := m.RuleInformation(rules)
		// match types
		if info.ProtoType != typ.Element().ProtoType() {
			m.failWithInvalidType(field)
			return // unreachable
		}
		// default value is nil
		flData.Iterate = true
		flData.RedactionValue = "nil"
		if info.ProtoType != pgs.MessageT {
			// simple type fields
			flData.RedactionValue = fmt.Sprintf("%v", info.RedactionValue)
		} else {
			// message type embedded field
			rule := rules.Values.(*redact.FieldRules_Message).Message
			flData.RedactionValue = `nil`
			if rule.Empty {
				//flData.RedactionValue = m.ctx.Type(field).String() + "{}"
				flData.RedactionValue = fmt.Sprintf("&%s{}", flData.EmbedMessageNameWithAlias)
				return
			}
			if rule.Nil {
				flData.RedactionValue = "nil"
				return
			}
			if rule.Skip {
				flData.EmbedSkip = true
				return
			}
			flData.NestedEmbedCall = true
		}
	}
}

// RuleInfo response type for Module.RuleInformation
type RuleInfo struct {
	RedactionValue interface{}
	// equivalent field type information
	ProtoType  pgs.ProtoType
	ProtoLabel pgs.ProtoLabel
}

// RuleInformation returns required information from the redact.FieldRules
func (m *Module) RuleInformation(rules *redact.FieldRules) (res RuleInfo) {
	// custom rules validation and values
	switch rule := rules.Values.(type) {
	case *redact.FieldRules_Float:
		res.ProtoType = pgs.FloatT
		res.RedactionValue = rule.Float
	case *redact.FieldRules_Double:
		res.ProtoType = pgs.DoubleT
		res.RedactionValue = rule.Double
	case *redact.FieldRules_Int32:
		res.ProtoType = pgs.Int32T
		res.RedactionValue = rule.Int32
	case *redact.FieldRules_Int64:
		res.ProtoType = pgs.Int64T
		res.RedactionValue = rule.Int64
	case *redact.FieldRules_Uint32:
		res.ProtoType = pgs.UInt32T
		res.RedactionValue = rule.Uint32
	case *redact.FieldRules_Uint64:
		res.ProtoType = pgs.UInt64T
		res.RedactionValue = rule.Uint64
	case *redact.FieldRules_Sint32:
		res.ProtoType = pgs.SInt32
		res.RedactionValue = rule.Sint32
	case *redact.FieldRules_Sint64:
		res.ProtoType = pgs.SInt64
		res.RedactionValue = rule.Sint64
	case *redact.FieldRules_Fixed32:
		res.ProtoType = pgs.Fixed32T
		res.RedactionValue = rule.Fixed32
	case *redact.FieldRules_Fixed64:
		res.ProtoType = pgs.Fixed64T
		res.RedactionValue = rule.Fixed64
	case *redact.FieldRules_Sfixed32:
		res.ProtoType = pgs.SFixed32
		res.RedactionValue = rule.Sfixed32
	case *redact.FieldRules_Sfixed64:
		res.ProtoType = pgs.SFixed64
		res.RedactionValue = rule.Sfixed64
	case *redact.FieldRules_Bool:
		res.ProtoType = pgs.BoolT
		res.RedactionValue = rule.Bool
	case *redact.FieldRules_String_:
		res.ProtoType = pgs.StringT
		res.RedactionValue = fmt.Sprintf("`%v`", rule.String_)
	case *redact.FieldRules_Bytes:
		res.ProtoType = pgs.BytesT
		res.RedactionValue = fmt.Sprintf("[]byte(`%v`)", string(rule.Bytes))
	case *redact.FieldRules_Enum:
		res.ProtoType = pgs.EnumT
		res.RedactionValue = rule.Enum
	case *redact.FieldRules_Message:
		res.ProtoType = pgs.MessageT
		if rule == nil || rule.Message == nil {
			m.Fail("(redact.custom).message is nil, no option defined")
			return // unreachable
		}
	case *redact.FieldRules_Element:
		res.ProtoLabel = pgs.Repeated
		if rule == nil || rule.Element == nil {
			m.Fail("(redact.custom).element is nil, no option defined")
			return // unreachable
		}
	default:
		m.Fail("Something went wrong")
	}
	return res
}
