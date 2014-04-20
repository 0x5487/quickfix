package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityTypes msg type = w.
type SecurityTypes struct {
	message.Message
}

//SecurityTypesBuilder builds SecurityTypes messages.
type SecurityTypesBuilder struct {
	message.MessageBuilder
}

//CreateSecurityTypesBuilder returns an initialized SecurityTypesBuilder with specified required fields.
func CreateSecurityTypesBuilder(
	securityreqid field.SecurityReqID,
	securityresponseid field.SecurityResponseID,
	securityresponsetype field.SecurityResponseType) SecurityTypesBuilder {
	var builder SecurityTypesBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(securityresponsetype)
	return builder
}

//SecurityReqID is a required field for SecurityTypes.
func (m SecurityTypes) SecurityReqID() (*field.SecurityReqID, errors.MessageRejectError) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityTypes.
func (m SecurityTypes) GetSecurityReqID(f *field.SecurityReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a required field for SecurityTypes.
func (m SecurityTypes) SecurityResponseID() (*field.SecurityResponseID, errors.MessageRejectError) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityTypes.
func (m SecurityTypes) GetSecurityResponseID(f *field.SecurityResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseType is a required field for SecurityTypes.
func (m SecurityTypes) SecurityResponseType() (*field.SecurityResponseType, errors.MessageRejectError) {
	f := new(field.SecurityResponseType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseType reads a SecurityResponseType from SecurityTypes.
func (m SecurityTypes) GetSecurityResponseType(f *field.SecurityResponseType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoSecurityTypes is a non-required field for SecurityTypes.
func (m SecurityTypes) TotNoSecurityTypes() (*field.TotNoSecurityTypes, errors.MessageRejectError) {
	f := new(field.TotNoSecurityTypes)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoSecurityTypes reads a TotNoSecurityTypes from SecurityTypes.
func (m SecurityTypes) GetTotNoSecurityTypes(f *field.TotNoSecurityTypes) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for SecurityTypes.
func (m SecurityTypes) LastFragment() (*field.LastFragment, errors.MessageRejectError) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from SecurityTypes.
func (m SecurityTypes) GetLastFragment(f *field.LastFragment) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityTypes is a non-required field for SecurityTypes.
func (m SecurityTypes) NoSecurityTypes() (*field.NoSecurityTypes, errors.MessageRejectError) {
	f := new(field.NoSecurityTypes)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityTypes reads a NoSecurityTypes from SecurityTypes.
func (m SecurityTypes) GetNoSecurityTypes(f *field.NoSecurityTypes) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SecurityTypes.
func (m SecurityTypes) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SecurityTypes.
func (m SecurityTypes) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SecurityTypes.
func (m SecurityTypes) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SecurityTypes.
func (m SecurityTypes) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SecurityTypes.
func (m SecurityTypes) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SecurityTypes.
func (m SecurityTypes) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SecurityTypes.
func (m SecurityTypes) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SecurityTypes.
func (m SecurityTypes) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for SecurityTypes.
func (m SecurityTypes) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from SecurityTypes.
func (m SecurityTypes) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for SecurityTypes.
func (m SecurityTypes) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from SecurityTypes.
func (m SecurityTypes) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}
