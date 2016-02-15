package relsymderivsecgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp1/secondarypricelimits"
)

//NoRelatedSym is a repeating group in RelSymDerivSecGrp
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
	//SecondaryPriceLimits Component
	SecondaryPriceLimits secondarypricelimits.Component
	//CorporateAction is a non-required field for NoRelatedSym.
	CorporateAction *string `fix:"292"`
}

//Component is a fix50sp1 RelSymDerivSecGrp Component
type Component struct {
	//NoRelatedSym is a non-required field for RelSymDerivSecGrp.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
}

func New() *Component { return new(Component) }