package event

import (
	"crypto/sha256"
	"fmt"
	"github.com/kurtosis-tech/stacktrace"
	"strings"
)

const (
	yesStr = "yes"
	noStr = "no"


	//We are following these naming conventions for event's data
	//https://segment.com/docs/getting-started/04-full-install/#event-naming-best-practices
	enclaveIDPropertyKey = "enclave_id"
	moduleIDPropertyKey = "module_id"
	userAcceptSendingMetricsKey = "user_accept_sending_metrics"
	shouldCleanAllPropertyKey = "should_clean_all"

	//Categories
	installCategory = "Install"
	enclaveCategory = "Enclave"
	moduleCategory = "Module"

	//Actions
	consentAction = "Consent"
	createAction = "Create"
	stopAction = "Stop"
	destroyAction = "Destroy"
	cleanAction = "Clean"
	loadAction = "Load"
	unloadAction = "Unload"
	executeAction = "Execute"

)

func NewUserAcceptSendingMetricsEvent(userAcceptSendingMetrics bool) (*Event, error) {
	var metricsValue string
	if userAcceptSendingMetrics{
		metricsValue = yesStr
	} else {
		metricsValue = noStr
	}

	event, err := newEvent(installCategory, consentAction, userAcceptSendingMetricsKey, metricsValue)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new user accept sending metrics event")
	}

	return event, nil
}

func NewCreateEnclaveEvent(enclaveId string) (*Event, error) {
	hashedEnclaveId, err := validateEnclaveIdAndGetHashedValue(enclaveId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred validating ang getting hashed enclave id")
	}

	event, err := newEvent(enclaveCategory, createAction, enclaveIDPropertyKey, hashedEnclaveId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new create enclave event")
	}

	return event, nil
}

func NewStopEnclaveEvent(enclaveId string) (*Event, error) {
	hashedEnclaveId, err := validateEnclaveIdAndGetHashedValue(enclaveId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred validating ang getting hashed enclave id")
	}

	event, err := newEvent(enclaveCategory, stopAction, enclaveIDPropertyKey, hashedEnclaveId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new stop enclave event")
	}

	return event, nil
}

func NewDestroyEnclaveEvent(enclaveId string) (*Event, error) {
	hashedEnclaveId, err := validateEnclaveIdAndGetHashedValue(enclaveId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred validating ang getting hashed enclave id")
	}

	event, err := newEvent(enclaveCategory, destroyAction, enclaveIDPropertyKey, hashedEnclaveId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new destroy enclave event")
	}

	return event, nil
}

func NewCleanEnclaveEvent(shouldCleanAll bool) (*Event, error) {
	var metricsValue string
	if shouldCleanAll{
		metricsValue = yesStr
	} else {
		metricsValue = noStr
	}

	event, err := newEvent(enclaveCategory, cleanAction, shouldCleanAllPropertyKey, metricsValue)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new clean enclave event")
	}
	return event, nil
}

func NewLoadModuleEvent(moduleId string) (*Event, error) {
	event, err := newEvent(moduleCategory, loadAction, moduleIDPropertyKey, moduleId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new load module event")
	}
	return event, nil
}

func NewUnloadModuleEvent(moduleId string) (*Event, error) {
	event, err := newEvent(moduleCategory, unloadAction, moduleIDPropertyKey, moduleId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new unload module event")
	}
	return event, nil
}

func NewExecuteModuleEvent(moduleId string) (*Event, error) {
	event, err := newEvent(moduleCategory, executeAction, moduleIDPropertyKey, moduleId)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred creating a new execute module event")
	}
	return event, nil
}

// ================================================================================================
//                                  Private Helper Functions
// ================================================================================================
func validateEnclaveIdAndGetHashedValue(enclaveId string) (string, error)  {
	enclaveId = strings.TrimSpace(enclaveId)

	if enclaveId == "" {
		return "", stacktrace.NewError("Enclave ID can not be empty string")
	}

	hashedEnclaveId := hashString(enclaveId)

	return hashedEnclaveId, nil
}

func hashString(value string) string {
	hash := sha256.New()

	hash.Write([]byte(value))

	hashedByteSlice := hash.Sum(nil)

	hexValue := fmt.Sprintf("%x", hashedByteSlice)

	return hexValue
}
