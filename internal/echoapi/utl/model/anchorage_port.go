package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AnchoragePort represents AnchoragePort model
type AnchoragePort struct {
	ID                       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PortID                   string             `json:"port_id" bson:"port_id"`
	DateActual               string             `json:"date_actual" bson:"date_actual"`
	Segment                  string             `json:"segment" bson:"segment"`
	NCongestedVessels        int                `json:"n_congested_vessels" bson:"n_congested_vessels"`
	AvgWaitingTime           int                `json:"avg_waiting_time" bson:"avg_waiting_time"`
	FirstQuartileWaitingTime int                `json:"first_quartile_waiting_time" bson:"first_quartile_waiting_time"`
	MedianWaitingTime        int                `json:"median_waiting_time" bson:"median_waiting_time"`
	ThirdQuartileWaitingTime int                `json:"third_quartile_waiting_time" bson:"third_quartile_waiting_time"`
	MaxWaitingTime           int                `json:"max_waiting_time" bson:"max_waiting_time"`
	DayOfYear                int                `json:"day_of_year" bson:"day_of_year"`
	Year                     int                `json:"year" bson:"year"`
	UnifiedDate              string             `json:"unified_date" bson:"unified_date"`
}

// AnchoragePortArray represents AnchorageDataArray model
type AnchoragePortArray struct {
	PortID                   string   `json:"port_id" bson:"port_id"`
	DateActual               []string `json:"date_actual" bson:"date_actual"`
	Segment                  string   `json:"segment" bson:"segment"`
	NCongestedVessels        []int    `json:"n_congested_vessels" bson:"n_congested_vessels"`
	AvgWaitingTime           []int    `json:"avg_waiting_time" bson:"avg_waiting_time"`
	FirstQuartileWaitingTime []int    `json:"first_quartile_waiting_time" bson:"first_quartile_waiting_time"`
	MedianWaitingTime        []int    `json:"median_waiting_time" bson:"median_waiting_time"`
	ThirdQuartileWaitingTime []int    `json:"third_quartile_waiting_time" bson:"third_quartile_waiting_time"`
	MaxWaitingTime           []int    `json:"max_waiting_time" bson:"max_waiting_time"`
	DayOfYear                []int    `json:"day_of_year" bson:"day_of_year"`
	Year                     []int    `json:"year" bson:"year"`
	UnifiedDate              []string `json:"unified_date" bson:"unified_date"`
}

func (a AnchoragePort) String() string {
	return fmt.Sprintf("AnchoragePort<%s \t| %s \t| %d\t>\n", a.DateActual, a.Segment, a.NCongestedVessels)
}

// ConvertAnchorageOutput converts anchorage output
func ConvertAnchorageOutput(a []AnchoragePort) *AnchoragePortArray {
	if len(a) == 0 {
		return nil
	}
	var portID = a[0].PortID
	var dateActual = make([]string, len(a))
	var segment = a[0].Segment
	var nCongestedVessels = make([]int, len(a))
	var avgWaitingTime = make([]int, len(a))
	var firstQuartileWaitingTime = make([]int, len(a))
	var medianWaitingTime = make([]int, len(a))
	var thirdQuartileWaitingTime = make([]int, len(a))
	var maxWaitingTime = make([]int, len(a))
	var dayOfYear = make([]int, len(a))
	var year = make([]int, len(a))
	var unifiedDate = make([]string, len(a))

	for k, v := range a {
		dateActual[k] = v.DateActual
		nCongestedVessels[k] = v.NCongestedVessels
		avgWaitingTime[k] = v.AvgWaitingTime
		firstQuartileWaitingTime[k] = v.FirstQuartileWaitingTime
		medianWaitingTime[k] = v.MedianWaitingTime
		thirdQuartileWaitingTime[k] = v.ThirdQuartileWaitingTime
		maxWaitingTime[k] = v.MaxWaitingTime
		dayOfYear[k] = v.DayOfYear
		year[k] = v.Year
		unifiedDate[k] = v.UnifiedDate
	}

	arrAncPort := AnchoragePortArray{
		PortID:                   portID,
		DateActual:               dateActual,
		Segment:                  segment,
		NCongestedVessels:        nCongestedVessels,
		AvgWaitingTime:           avgWaitingTime,
		FirstQuartileWaitingTime: firstQuartileWaitingTime,
		MedianWaitingTime:        medianWaitingTime,
		ThirdQuartileWaitingTime: thirdQuartileWaitingTime,
		MaxWaitingTime:           maxWaitingTime,
		DayOfYear:                dayOfYear,
		Year:                     year,
		UnifiedDate:              unifiedDate,
	}
	return &arrAncPort
}
