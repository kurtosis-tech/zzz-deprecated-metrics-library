package client

import "github.com/sirupsen/logrus"

// doNothingClient: This metrics client implementation has been created for instantiate when user rejects
// sending metrics, so it doesn't really track metrics the only logic that it contains is loging
// the traking methods calls. It also can be used for test purpose
type doNothingClient struct {
	callback Callback
}

func newDoNothingClient(callback Callback) *doNothingClient {
	return &doNothingClient{callback: callback}
}

func (client *doNothingClient) TrackShouldSendMetricsUserElection(didUserAcceptSendingMetrics bool) error {
	logrus.Debugf("Do-nothing metrics client TrackShouldSendMetricsUserElection called with argument didUserAcceptSendingMetrics '%v'; skipping sending event", didUserAcceptSendingMetrics)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackCreateEnclave(enclaveId string) error {
	logrus.Debugf("Do-nothing metrics client TrackCreateEnclave called with argument enclaveId '%v'; skipping sending event", enclaveId)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackStopEnclave(enclaveId string) error {
	logrus.Debugf("Do-nothing metrics client TrackStopEnclave called with argument enclaveId '%v'; skipping sending event", enclaveId)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackDestroyEnclave(enclaveId string) error {
	logrus.Debugf("Do-nothing metrics client TrackDestroyEnclave called with argument enclaveId '%v'; skipping sending event", enclaveId)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackKurtosisRun(packageId string, isRemote bool, isDryRun bool, isScript bool) error {
	logrus.Debugf("Do-nothing metrics client TrackKurtosisRun called with arguments packageId '%v', isRemote '%v', isDryRun '%v' and isScript '%v'; skipping sending event", packageId, isRemote, isDryRun, isScript)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) TrackKurtosisRunFinishedEvent(packageId string, numberOfServices int) error {
	logrus.Debugf("Do-nothing metrics client TrackKurtosisRunFinishedEvent called with arguments packageId '%v' and numberOfServices '%v'; skipping sending event", packageId, numberOfServices)
	client.callback.Success()
	return nil
}

func (client *doNothingClient) close() (err error) {
	logrus.Debugf("Do-nothing metrics client close method called")
	return nil
}
