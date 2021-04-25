/*
 * eodhistoricaldata
 *
 * eodhistoricaldata API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FundamentalsESGScores struct for FundamentalsESGScores
type FundamentalsESGScores struct {
	Disclaimer                 *string                                                `json:"Disclaimer,omitempty"`
	RatingDate                 *string                                                `json:"RatingDate,omitempty"`
	TotalEsg                   *float32                                               `json:"TotalEsg,omitempty"`
	TotalEsgPercentile         *float32                                               `json:"TotalEsgPercentile,omitempty"`
	EnvironmentScore           *float32                                               `json:"EnvironmentScore,omitempty"`
	EnvironmentScorePercentile *float32                                               `json:"EnvironmentScorePercentile,omitempty"`
	SocialScore                *float32                                               `json:"SocialScore,omitempty"`
	SocialScorePercentile      *float32                                               `json:"SocialScorePercentile,omitempty"`
	GovernanceScore            *float32                                               `json:"GovernanceScore,omitempty"`
	GovernanceScorePercentile  *float32                                               `json:"GovernanceScorePercentile,omitempty"`
	ControversyLevel           *float32                                               `json:"ControversyLevel,omitempty"`
	ActivitiesInvolvement      *map[string]FundamentalsESGScoresActivitiesInvolvement `json:"ActivitiesInvolvement,omitempty"`
}

// NewFundamentalsESGScores instantiates a new FundamentalsESGScores object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundamentalsESGScores() *FundamentalsESGScores {
	this := FundamentalsESGScores{}
	return &this
}

// NewFundamentalsESGScoresWithDefaults instantiates a new FundamentalsESGScores object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundamentalsESGScoresWithDefaults() *FundamentalsESGScores {
	this := FundamentalsESGScores{}
	return &this
}

// GetDisclaimer returns the Disclaimer field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetDisclaimer() string {
	if o == nil || o.Disclaimer == nil {
		var ret string
		return ret
	}
	return *o.Disclaimer
}

// GetDisclaimerOk returns a tuple with the Disclaimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetDisclaimerOk() (*string, bool) {
	if o == nil || o.Disclaimer == nil {
		return nil, false
	}
	return o.Disclaimer, true
}

// HasDisclaimer returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasDisclaimer() bool {
	if o != nil && o.Disclaimer != nil {
		return true
	}

	return false
}

// SetDisclaimer gets a reference to the given string and assigns it to the Disclaimer field.
func (o *FundamentalsESGScores) SetDisclaimer(v string) {
	o.Disclaimer = &v
}

// GetRatingDate returns the RatingDate field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetRatingDate() string {
	if o == nil || o.RatingDate == nil {
		var ret string
		return ret
	}
	return *o.RatingDate
}

// GetRatingDateOk returns a tuple with the RatingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetRatingDateOk() (*string, bool) {
	if o == nil || o.RatingDate == nil {
		return nil, false
	}
	return o.RatingDate, true
}

// HasRatingDate returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasRatingDate() bool {
	if o != nil && o.RatingDate != nil {
		return true
	}

	return false
}

// SetRatingDate gets a reference to the given string and assigns it to the RatingDate field.
func (o *FundamentalsESGScores) SetRatingDate(v string) {
	o.RatingDate = &v
}

// GetTotalEsg returns the TotalEsg field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetTotalEsg() float32 {
	if o == nil || o.TotalEsg == nil {
		var ret float32
		return ret
	}
	return *o.TotalEsg
}

// GetTotalEsgOk returns a tuple with the TotalEsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetTotalEsgOk() (*float32, bool) {
	if o == nil || o.TotalEsg == nil {
		return nil, false
	}
	return o.TotalEsg, true
}

// HasTotalEsg returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasTotalEsg() bool {
	if o != nil && o.TotalEsg != nil {
		return true
	}

	return false
}

// SetTotalEsg gets a reference to the given float32 and assigns it to the TotalEsg field.
func (o *FundamentalsESGScores) SetTotalEsg(v float32) {
	o.TotalEsg = &v
}

// GetTotalEsgPercentile returns the TotalEsgPercentile field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetTotalEsgPercentile() float32 {
	if o == nil || o.TotalEsgPercentile == nil {
		var ret float32
		return ret
	}
	return *o.TotalEsgPercentile
}

// GetTotalEsgPercentileOk returns a tuple with the TotalEsgPercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetTotalEsgPercentileOk() (*float32, bool) {
	if o == nil || o.TotalEsgPercentile == nil {
		return nil, false
	}
	return o.TotalEsgPercentile, true
}

// HasTotalEsgPercentile returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasTotalEsgPercentile() bool {
	if o != nil && o.TotalEsgPercentile != nil {
		return true
	}

	return false
}

// SetTotalEsgPercentile gets a reference to the given float32 and assigns it to the TotalEsgPercentile field.
func (o *FundamentalsESGScores) SetTotalEsgPercentile(v float32) {
	o.TotalEsgPercentile = &v
}

// GetEnvironmentScore returns the EnvironmentScore field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetEnvironmentScore() float32 {
	if o == nil || o.EnvironmentScore == nil {
		var ret float32
		return ret
	}
	return *o.EnvironmentScore
}

// GetEnvironmentScoreOk returns a tuple with the EnvironmentScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetEnvironmentScoreOk() (*float32, bool) {
	if o == nil || o.EnvironmentScore == nil {
		return nil, false
	}
	return o.EnvironmentScore, true
}

// HasEnvironmentScore returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasEnvironmentScore() bool {
	if o != nil && o.EnvironmentScore != nil {
		return true
	}

	return false
}

// SetEnvironmentScore gets a reference to the given float32 and assigns it to the EnvironmentScore field.
func (o *FundamentalsESGScores) SetEnvironmentScore(v float32) {
	o.EnvironmentScore = &v
}

// GetEnvironmentScorePercentile returns the EnvironmentScorePercentile field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetEnvironmentScorePercentile() float32 {
	if o == nil || o.EnvironmentScorePercentile == nil {
		var ret float32
		return ret
	}
	return *o.EnvironmentScorePercentile
}

// GetEnvironmentScorePercentileOk returns a tuple with the EnvironmentScorePercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetEnvironmentScorePercentileOk() (*float32, bool) {
	if o == nil || o.EnvironmentScorePercentile == nil {
		return nil, false
	}
	return o.EnvironmentScorePercentile, true
}

// HasEnvironmentScorePercentile returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasEnvironmentScorePercentile() bool {
	if o != nil && o.EnvironmentScorePercentile != nil {
		return true
	}

	return false
}

// SetEnvironmentScorePercentile gets a reference to the given float32 and assigns it to the EnvironmentScorePercentile field.
func (o *FundamentalsESGScores) SetEnvironmentScorePercentile(v float32) {
	o.EnvironmentScorePercentile = &v
}

// GetSocialScore returns the SocialScore field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetSocialScore() float32 {
	if o == nil || o.SocialScore == nil {
		var ret float32
		return ret
	}
	return *o.SocialScore
}

// GetSocialScoreOk returns a tuple with the SocialScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetSocialScoreOk() (*float32, bool) {
	if o == nil || o.SocialScore == nil {
		return nil, false
	}
	return o.SocialScore, true
}

// HasSocialScore returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasSocialScore() bool {
	if o != nil && o.SocialScore != nil {
		return true
	}

	return false
}

// SetSocialScore gets a reference to the given float32 and assigns it to the SocialScore field.
func (o *FundamentalsESGScores) SetSocialScore(v float32) {
	o.SocialScore = &v
}

// GetSocialScorePercentile returns the SocialScorePercentile field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetSocialScorePercentile() float32 {
	if o == nil || o.SocialScorePercentile == nil {
		var ret float32
		return ret
	}
	return *o.SocialScorePercentile
}

// GetSocialScorePercentileOk returns a tuple with the SocialScorePercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetSocialScorePercentileOk() (*float32, bool) {
	if o == nil || o.SocialScorePercentile == nil {
		return nil, false
	}
	return o.SocialScorePercentile, true
}

// HasSocialScorePercentile returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasSocialScorePercentile() bool {
	if o != nil && o.SocialScorePercentile != nil {
		return true
	}

	return false
}

// SetSocialScorePercentile gets a reference to the given float32 and assigns it to the SocialScorePercentile field.
func (o *FundamentalsESGScores) SetSocialScorePercentile(v float32) {
	o.SocialScorePercentile = &v
}

// GetGovernanceScore returns the GovernanceScore field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetGovernanceScore() float32 {
	if o == nil || o.GovernanceScore == nil {
		var ret float32
		return ret
	}
	return *o.GovernanceScore
}

// GetGovernanceScoreOk returns a tuple with the GovernanceScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetGovernanceScoreOk() (*float32, bool) {
	if o == nil || o.GovernanceScore == nil {
		return nil, false
	}
	return o.GovernanceScore, true
}

// HasGovernanceScore returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasGovernanceScore() bool {
	if o != nil && o.GovernanceScore != nil {
		return true
	}

	return false
}

// SetGovernanceScore gets a reference to the given float32 and assigns it to the GovernanceScore field.
func (o *FundamentalsESGScores) SetGovernanceScore(v float32) {
	o.GovernanceScore = &v
}

// GetGovernanceScorePercentile returns the GovernanceScorePercentile field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetGovernanceScorePercentile() float32 {
	if o == nil || o.GovernanceScorePercentile == nil {
		var ret float32
		return ret
	}
	return *o.GovernanceScorePercentile
}

// GetGovernanceScorePercentileOk returns a tuple with the GovernanceScorePercentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetGovernanceScorePercentileOk() (*float32, bool) {
	if o == nil || o.GovernanceScorePercentile == nil {
		return nil, false
	}
	return o.GovernanceScorePercentile, true
}

// HasGovernanceScorePercentile returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasGovernanceScorePercentile() bool {
	if o != nil && o.GovernanceScorePercentile != nil {
		return true
	}

	return false
}

// SetGovernanceScorePercentile gets a reference to the given float32 and assigns it to the GovernanceScorePercentile field.
func (o *FundamentalsESGScores) SetGovernanceScorePercentile(v float32) {
	o.GovernanceScorePercentile = &v
}

// GetControversyLevel returns the ControversyLevel field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetControversyLevel() float32 {
	if o == nil || o.ControversyLevel == nil {
		var ret float32
		return ret
	}
	return *o.ControversyLevel
}

// GetControversyLevelOk returns a tuple with the ControversyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetControversyLevelOk() (*float32, bool) {
	if o == nil || o.ControversyLevel == nil {
		return nil, false
	}
	return o.ControversyLevel, true
}

// HasControversyLevel returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasControversyLevel() bool {
	if o != nil && o.ControversyLevel != nil {
		return true
	}

	return false
}

// SetControversyLevel gets a reference to the given float32 and assigns it to the ControversyLevel field.
func (o *FundamentalsESGScores) SetControversyLevel(v float32) {
	o.ControversyLevel = &v
}

// GetActivitiesInvolvement returns the ActivitiesInvolvement field value if set, zero value otherwise.
func (o *FundamentalsESGScores) GetActivitiesInvolvement() map[string]FundamentalsESGScoresActivitiesInvolvement {
	if o == nil || o.ActivitiesInvolvement == nil {
		var ret map[string]FundamentalsESGScoresActivitiesInvolvement
		return ret
	}
	return *o.ActivitiesInvolvement
}

// GetActivitiesInvolvementOk returns a tuple with the ActivitiesInvolvement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundamentalsESGScores) GetActivitiesInvolvementOk() (*map[string]FundamentalsESGScoresActivitiesInvolvement, bool) {
	if o == nil || o.ActivitiesInvolvement == nil {
		return nil, false
	}
	return o.ActivitiesInvolvement, true
}

// HasActivitiesInvolvement returns a boolean if a field has been set.
func (o *FundamentalsESGScores) HasActivitiesInvolvement() bool {
	if o != nil && o.ActivitiesInvolvement != nil {
		return true
	}

	return false
}

// SetActivitiesInvolvement gets a reference to the given map[string]FundamentalsESGScoresActivitiesInvolvement and assigns it to the ActivitiesInvolvement field.
func (o *FundamentalsESGScores) SetActivitiesInvolvement(v map[string]FundamentalsESGScoresActivitiesInvolvement) {
	o.ActivitiesInvolvement = &v
}

func (o FundamentalsESGScores) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Disclaimer != nil {
		toSerialize["Disclaimer"] = o.Disclaimer
	}
	if o.RatingDate != nil {
		toSerialize["RatingDate"] = o.RatingDate
	}
	if o.TotalEsg != nil {
		toSerialize["TotalEsg"] = o.TotalEsg
	}
	if o.TotalEsgPercentile != nil {
		toSerialize["TotalEsgPercentile"] = o.TotalEsgPercentile
	}
	if o.EnvironmentScore != nil {
		toSerialize["EnvironmentScore"] = o.EnvironmentScore
	}
	if o.EnvironmentScorePercentile != nil {
		toSerialize["EnvironmentScorePercentile"] = o.EnvironmentScorePercentile
	}
	if o.SocialScore != nil {
		toSerialize["SocialScore"] = o.SocialScore
	}
	if o.SocialScorePercentile != nil {
		toSerialize["SocialScorePercentile"] = o.SocialScorePercentile
	}
	if o.GovernanceScore != nil {
		toSerialize["GovernanceScore"] = o.GovernanceScore
	}
	if o.GovernanceScorePercentile != nil {
		toSerialize["GovernanceScorePercentile"] = o.GovernanceScorePercentile
	}
	if o.ControversyLevel != nil {
		toSerialize["ControversyLevel"] = o.ControversyLevel
	}
	if o.ActivitiesInvolvement != nil {
		toSerialize["ActivitiesInvolvement"] = o.ActivitiesInvolvement
	}
	return json.Marshal(toSerialize)
}

type NullableFundamentalsESGScores struct {
	value *FundamentalsESGScores
	isSet bool
}

func (v NullableFundamentalsESGScores) Get() *FundamentalsESGScores {
	return v.value
}

func (v *NullableFundamentalsESGScores) Set(val *FundamentalsESGScores) {
	v.value = val
	v.isSet = true
}

func (v NullableFundamentalsESGScores) IsSet() bool {
	return v.isSet
}

func (v *NullableFundamentalsESGScores) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundamentalsESGScores(val *FundamentalsESGScores) *NullableFundamentalsESGScores {
	return &NullableFundamentalsESGScores{value: val, isSet: true}
}

func (v NullableFundamentalsESGScores) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundamentalsESGScores) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
