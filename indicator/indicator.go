// Copyright 2015 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license
// that can be found in the LICENSE file in the root of the source
// tree.

package indicator

import (
	"github.com/pborman/uuid"
	"github.com/activeshadow/libcybox/observable"
	"github.com/activeshadow/libstix/common"
	"github.com/activeshadow/libstix/defs"
	"time"
)

// ----------------------------------------------------------------------
// Define Message Type
// ----------------------------------------------------------------------

type CompositeIndicatorExpressionType struct {
	Operator  string         `json:"operator,omitempty"`
	Indicator *IndicatorType `json:"indicator,omitempty"`
}

type RelatedCampaignReferencesType struct {
	Scope            string                              `json:"scope,omitempty"`
	Related_Campaign common.RelatedCampaignReferenceType `json:"relatedCampaign,omitempty"`
}

// Support for IndicatorType version 2.1.1

type IndicatorType struct {
	common.IndicatorBaseType
	Version                      string                               `json:"version,omitempty"`
	Negate                       bool                                 `json:"negate,omitempty"`
	Title                        string                               `json:"title,omitempty"`
	Types                        []string                             `json:"type,omitempty"`
	AlternativeIDs               []string                             `json:"alternative_ids,omitempty"`
	Descriptions                 []common.StructuredTextType          `json:"descriptions,omitempty"`
	ShortDescriptions            []common.StructuredTextType          `json:"short_descriptions,omitempty"`
	ValidTimePositions           []ValidTimeType                      `json:"valid_time_positions,omitempty"`
	Observable                   *observable.ObservableType           `json:"observable,omitempty"`
	CompositeIndicatorExpression *CompositeIndicatorExpressionType    `json:"composite_indicator_expression,omitempty"`
	IndicatedTTP                 []common.RelatedTTPType              `json:"indicated_ttps,omitempty"`
	KillChainPhases              []common.KillChainPhaseReferenceType `json:"kill_chain_phases,omitempty"`
	TestMechanisms               []TestMechanismType                  `json:"test_mechanisms,omitempty"`
	LikelyImpact                 *common.StatementType                `json:"likely_impact,omitempty"`
	SuggestedCOAs                []SuggestedCOAsType                  `json:"suggested_coas,omitempty"`
	Handling                     []common.MarkingSpecificationType    `json:"handling,omitempty"`
	Confidence                   *common.ConfidenceType               `json:"confidence,omitempty"`
	Sightings                    *SightingsType                       `json:"sightings,omitempty"`
	RelatedIndicators            *RelatedIndicatorsType               `json:"related_indicators,omitempty"`
	RelatedCampaigns             *RelatedCampaignReferencesType       `json:"related_campaigns,omitempty"`
	RelatedPackages              []common.RelatedPackageRefType       `json:"related_packages,omitempty"`
	Producer                     *common.InformationSourceType        `json:"producer,omitempty"`
}

// TODO need to test this, this may not work

type SuggestedCOAsType struct {
	Scope        string                            `json:"scope,omitempty"`
	SuggestedCOA *common.RelatedCourseOfActionType `json:"suggestedCOA,omitempty"`
}

type TestMechanismType struct {
	Id       string                        `json:"id,omitempty"`
	IdRef    string                        `json:"idref,omitempty"`
	Efficacy *common.StatementType         `json:"efficacy,omitempty"`
	Producer *common.InformationSourceType `json:"producer,omitempty"`
}

type ValidTimeType struct {
	StartTime      string `json:"start_time,omitempty"`
	StartPrecision string `json:"start_precision,omitempty"`
	EndTime        string `json:"end_time,omitempty"`
	EndPrecision   string `json:"end_precision,omitempty"`
}

// ----------------------------------------------------------------------
// Create Functions
// ----------------------------------------------------------------------

func New() IndicatorType {
	obj := CreateIndicator()
	obj.CreateId()
	return obj
}

func CreateIndicator() IndicatorType {
	var obj IndicatorType
	return obj
}

// ----------------------------------------------------------------------
// Methods IndicatorType
// ----------------------------------------------------------------------

func (this *IndicatorType) CreateId() {
	this.Id = defs.COMPANY + ":indicator-" + uuid.New()
}

func (this *IndicatorType) AddIdRef(idref string) {
	this.IdRef = idref
}

func (this *IndicatorType) SetTimestampToNow() {
	this.Timestamp = time.Now().Format(time.RFC3339)
}

func (this *IndicatorType) AddTimeStamp(t string) {
	// TODO Need to format the string in to ISO 8601 format or check that it is in the right format
	this.Timestamp = t
}

func (this *IndicatorType) AddVersion(ver string) {
	this.Version = ver
}

func (this *IndicatorType) SetNegate(b bool) {
	this.Negate = b
}

func (this *IndicatorType) AddTitle(t string) {
	this.Title = t
}

// func (this *IndicatorType) AddType(vocab, value string) {
// 	if this.Types == nil {
// 		m := make(map[string][]string)
// 		this.Types = m
// 	}
// 	this.Types[vocab] = append(this.Types[vocab], value)
// }

// func (this *IndicatorType) AddStandardType(value string) {
// 	this.AddType(defs.INDICATOR_TYPE_VOCAB, value)
// }

func (this *IndicatorType) AddType(value string) {
	if this.Types == nil {
		a := make([]string, 0)
		this.Types = a
	}
	this.Types = append(this.Types, value)
}

func (this *IndicatorType) AddAlternativeID(value string) {
	if this.AlternativeIDs == nil {
		a := make([]string, 0)
		this.AlternativeIDs = a
	}
	this.AlternativeIDs = append(this.AlternativeIDs, value)
}

func (this *IndicatorType) AddDescriptionText(markingidref, value string) {
	this.AddDescription(defs.RFC6838TEXTPLAIN, markingidref, value)
}

func (this *IndicatorType) AddDescription(format, markingidref, value string) {
	if this.Descriptions == nil {
		a := make([]common.StructuredTextType, 0)
		this.Descriptions = a
	}

	var data common.StructuredTextType
	data.CreateId()
	data.AddMarkingIdRef(markingidref)
	data.AddFormat(format)
	data.AddValue(value)

	this.Descriptions = append(this.Descriptions, data)
}

func (this *IndicatorType) AddShortDescriptionText(markingidref, value string) {
	this.AddShortDescription(defs.RFC6838TEXTPLAIN, markingidref, value)
}

func (this *IndicatorType) AddShortDescription(format, markingidref, value string) {
	if this.ShortDescriptions == nil {
		a := make([]common.StructuredTextType, 0)
		this.ShortDescriptions = a
	}

	var data common.StructuredTextType
	data.CreateId()
	data.AddMarkingIdRef(markingidref)
	data.AddFormat(format)
	data.AddValue(value)

	this.ShortDescriptions = append(this.ShortDescriptions, data)
}

func (this *IndicatorType) initValidTimePosition() {
	if this.ValidTimePositions == nil {
		a := make([]ValidTimeType, 0)
		this.ValidTimePositions = a
	}
}

func (this *IndicatorType) AddValidTimePosition(start, end string) {
	this.initValidTimePosition()

	tp := ValidTimeType{
		StartTime:      start,
		StartPrecision: "Second",
		EndTime:        end,
		EndPrecision:   "Second",
	}
	this.ValidTimePositions = append(this.ValidTimePositions, tp)
}

func (this *IndicatorType) NewObservable() *observable.ObservableType {
	o := observable.New()
	this.AddObservable(o)
	return this.Observable
}

func (this *IndicatorType) AddObservable(o observable.ObservableType) {
	this.Observable = &o
}

// TODO Composite_Indicator_Expression

// TODO Indicated_TTP

func (this *IndicatorType) AddKillChain(k common.KillChainPhaseReferenceType) {
	if this.KillChainPhases == nil {
		a := make([]common.KillChainPhaseReferenceType, 0)
		this.KillChainPhases = a
	}
	this.KillChainPhases = append(this.KillChainPhases, k)
}

func (this *IndicatorType) AddKillChainPhaseAndChain(phase, chain string) {
	if this.KillChainPhases == nil {
		a := make([]common.KillChainPhaseReferenceType, 0)
		this.KillChainPhases = a
	}
	data := common.KillChainPhaseReferenceType{
		PhaseId:     phase,
		KillChainId: chain,
	}
	this.KillChainPhases = append(this.KillChainPhases, data)
}

// TODO Test_Mechanisms

func (this *IndicatorType) AddLikelyImpact(s common.StatementType) {
	this.LikelyImpact = &s
}

func (this *IndicatorType) AddHandling(m common.MarkingSpecificationType) {
	if this.Handling == nil {
		a := make([]common.MarkingSpecificationType, 0)
		this.Handling = a
	}
	this.Handling = append(this.Handling, m)
}

func (this *IndicatorType) AddConfidence(c common.ConfidenceType) {
	this.Confidence = &c
}

func (this *IndicatorType) AddSightings(s SightingsType) {
	this.Sightings = &s
}

func (this *IndicatorType) AddRelatedIndicators(r RelatedIndicatorsType) {
	this.RelatedIndicators = &r
}

// TODO Related_Campaigns

func (this *IndicatorType) AddRelatedPackage(p common.RelatedPackageRefType) {
	if this.RelatedPackages == nil {
		a := make([]common.RelatedPackageRefType, 0)
		this.RelatedPackages = a
	}
	this.RelatedPackages = append(this.RelatedPackages, p)
}

func (this *IndicatorType) AddProducer(source common.InformationSourceType) {
	this.Producer = &source
}
