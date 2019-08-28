//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/rocksolidlabs/gofish/common"
)

// EventFormatType is
type EventFormatType string

const (

	// EventEventFormatType The subscription destination will receive JSON
	// Bodies of the Resource Type Event.
	EventEventFormatType EventFormatType = "Event"
	// MetricReportEventFormatType The subscription destination will receive
	// JSON Bodies of the Resource Type MetricReport.
	MetricReportEventFormatType EventFormatType = "MetricReport"
)

// EventService is used to represent an event service for a Redfish
// implementation.
type EventService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// DeliveryRetryAttempts shall be the
	// number of retrys attempted for any given event to the subscription
	// destination before the subscription is terminated.  This retry is at
	// the service level, meaning the HTTP POST to the Event Destination was
	// returned by the HTTP operation as unsuccessful (4xx or 5xx return
	// code) or an HTTP timeout occurred this many times before the Event
	// Destination subscription is terminated.
	DeliveryRetryAttempts int
	// DeliveryRetryIntervalSeconds shall be the interval in seconds between the
	// retry attempts for any given event
	// to the subscription destination.
	DeliveryRetryIntervalSeconds int
	// Description provides a description of this resource.
	Description string
	// EventFormatTypes shall indicate the the
	// content types of the message that this service can send to the event
	// destination.  If this property is not present, the EventFormatType
	// shall be assumed to be Event.
	EventFormatTypes []EventFormatType
	// RegistryPrefixes is the array of the Prefixes of the Message Registries
	// that shall be allowed for an Event Subscription.
	RegistryPrefixes []string
	// ResourceTypes is used for an Event Subscription.
	ResourceTypes []string
	// SSEFilterPropertiesSupported shall contain a set of properties that
	// indicate which properties are supported in the $filter query parameter
	// for the URI indicated by the ServerSentEventUri property.
	SSEFilterPropertiesSupported SSEFilterPropertiesSupported
	// ServerSentEventURI shall be a URI that specifies an HTML5 Server-Sent
	// Event conformant endpoint.
	ServerSentEventURI string `json:"ServerSentEventUri"`
	// ServiceEnabled shall be a boolean indicating whether this service is enabled.
	ServiceEnabled bool
	// Status is This property shall contain any status or health properties of
	// the resource.
	Status common.Status
	// SubordinateResourcesSupported is When set to true, the service is
	// indicating that it supports the SubordinateResource property on Event
	// Subscriptions and on generated Events.
	SubordinateResourcesSupported bool
	// Subscriptions shall contain the link to a collection of type
	// EventDestinationCollection.
	subscriptions string
}

// UnmarshalJSON unmarshals a EventService object from the raw JSON.
func (eventservice *EventService) UnmarshalJSON(b []byte) error {
	type temp EventService
	var t struct {
		temp
		Subscriptions common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*eventservice = EventService(t.temp)
	eventservice.subscriptions = string(t.Subscriptions)

	return nil
}

// GetEventService will get a EventService instance from the service.
func GetEventService(c common.Client, uri string) (*EventService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var eventservice EventService
	err = json.NewDecoder(resp.Body).Decode(&eventservice)
	if err != nil {
		return nil, err
	}

	eventservice.SetClient(c)
	return &eventservice, nil
}

// ListReferencedEventServices gets the collection of EventService from
// a provided reference.
func ListReferencedEventServices(c common.Client, link string) ([]*EventService, error) {
	var result []*EventService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, eventserviceLink := range links.ItemLinks {
		eventservice, err := GetEventService(c, eventserviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, eventservice)
	}

	return result, nil
}

// SSEFilterPropertiesSupported shall contain a set of properties that indicate
// which properties are supported in the $filter query parameter for the URI
// indicated by the ServerSentEventUri property.
type SSEFilterPropertiesSupported struct {
	// EventFormatType shall be a boolean indicating if this service supports
	// the use of the EventFormatType property in the $filter query parameter as
	// described by the specification.
	EventFormatType bool
	// MessageID shall be a boolean indicating if this service supports the use
	// of the MessageId property in the $filter query parameter as described by
	// the specification.
	MessageID bool `json:"MessageId"`
	// MetricReportDefinition shall be a boolean indicating if this service
	// supports the use of the MetricReportDefinition property in the $filter
	// query parameter as described by the specification.
	MetricReportDefinition bool
	// OriginResource shall be a boolean indicating if this service supports the
	// use of the OriginResource property in the $filter query parameter as
	// described by the specification.
	OriginResource bool
	// RegistryPrefix shall be a boolean indicating if this service supports the
	// use of the RegistryPrefix property in the $filter query parameter as
	// described by the specification.
	RegistryPrefix bool
	// ResourceType shall be a boolean indicating if this service supports the
	// use of the ResourceType property in the $filter query parameter as
	// described by the specification.
	ResourceType bool
}
