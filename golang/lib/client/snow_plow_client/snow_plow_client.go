package snow_plow_client

import (
	"github.com/kurtosis-tech/metrics-library/golang/lib/event"
	metrics_source "github.com/kurtosis-tech/metrics-library/golang/lib/source"
	"github.com/kurtosis-tech/stacktrace"
	"github.com/sirupsen/logrus"
	sp "github.com/snowplow/snowplow-golang-tracker/v2/tracker"
)

const (
	spCollectorURI = "8e280f93-12b7-4610-bd19-a5d7bc9e41dd.app.try-snowplow.com"
	spProtocol     = "https"
	spNamespace    = "kurtosistech"
	//Now we are using "pc" as default, but in the future we could use "srv"
	//for Kurt-Engine and Kurt-API sources is they run in KurtosisSAS
	//Available values https://github.com/snowplow/enrich/issues/450
	spDefaultPlatform = "pc"

	yesStr = "yes"
	noStr = "no"
)

var spOptionCallback = func(successCount []sp.CallbackResult, failureCount []sp.CallbackResult) {
	for _, result := range successCount {
		logrus.Debugf("SnowPlow emitter succes count: %v", result.Count)
		logrus.Debugf("SnowPlow emitter succes status: %v", result.Status)
	}
	for _, result := range failureCount {
		logrus.Debugf("SnowPlow emitter failure count: %v", result.Count)
		logrus.Debugf("SnowPlow emitter failure status: %v", result.Status)
	}
}

type SnowPlowClient struct {
	tracker *sp.Tracker
}

func NewSnowPlowClient(source metrics_source.Source, userId string) (*SnowPlowClient, error) {
	if err := source.IsValid(); err != nil {
		return nil, stacktrace.Propagate(err, "Invalid source")
	}

	subject := sp.InitSubject()
	subject.SetUserId(userId)
	emitter := sp.InitEmitter(
		sp.RequireCollectorUri(spCollectorURI),
		sp.OptionCallback(spOptionCallback),
		sp.OptionProtocol(spProtocol),
	)
	tracker := sp.InitTracker(
		sp.RequireEmitter(emitter),
		sp.OptionSubject(subject),
		sp.OptionNamespace(spNamespace),
		//Now we are using "pc" as default, but in the future we could use "srv"
		//for Kurt-Engine and Kurt-API sources is they run in KurtosisSAS
		sp.OptionPlatform(spDefaultPlatform),
		sp.OptionAppId(string(source)),
	)

	return &SnowPlowClient{tracker: tracker}, nil
}

func (snowPlow *SnowPlowClient) TrackUserAcceptSendingMetrics(userAcceptSendingMetrics bool) error {

	newEvent, err := event.NewUserAcceptSendingMetricsEvent(userAcceptSendingMetrics)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new user accept sending metrics event")
	}

	snowPlow.track(newEvent)

	return nil
}

func (snowPlow *SnowPlowClient) TrackCreateEnclave(enclaveId string) error {

	newEvent, err := event.NewCreateEnclaveEvent(enclaveId)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new create enclave event")
	}

	snowPlow.track(newEvent)

	return nil
}

func (snowPlow *SnowPlowClient) TrackStopEnclave(enclaveId string) error {

	newEvent, err := event.NewStopEnclaveEvent(enclaveId)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new stop enclave event")
	}

	snowPlow.track(newEvent)

	return nil
}


func (snowPlow *SnowPlowClient) TrackDestroyEnclave(enclaveId string) error {

	newEvent, err := event.NewDestroyEnclaveEvent(enclaveId)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new destroy enclave event")
	}

	snowPlow.track(newEvent)

	return nil
}

func (snowPlow *SnowPlowClient) TrackCleanEnclave(shouldCleanAll bool) error {

	newEvent, err := event.NewCleanEnclaveEvent(shouldCleanAll)
	if err != nil {
		return stacktrace.Propagate(err, "An error occurred creating a new clean enclave event")
	}

	snowPlow.track(newEvent)

	return nil
}

func (snowPlow *SnowPlowClient) track(event *event.Event) error {

	//We are using StructuredEvent types because we can match current Kurtosis Events with this type
	//of events, we also could use SelfDescribing events type if the Structured is not enough
	//More about SnowPlow events:
	//https://docs.snowplowanalytics.com/docs/understanding-tracking-design/out-of-the-box-vs-custom-events-and-entities/
	//https://docs.snowplowanalytics.com/docs/collecting-data/collecting-from-own-applications/golang-tracker/tracking-specific-events/#struct-event
	//https://docs.snowplowanalytics.com/docs/collecting-data/collecting-from-own-applications/javascript-trackers/javascript-tracker/javascript-tracker-v2/tracking-specific-events/#tracking-custom-structured-events
	snowPlow.tracker.TrackStructEvent(sp.StructuredEvent{
		Category: sp.NewString(event.GetCategory()),
		Action:   sp.NewString(event.GetAction()),
		Label:    sp.NewString(event.GetPropertyKey()),
		Property: sp.NewString(event.GetPropertyValue()),
	})
	return nil
}
