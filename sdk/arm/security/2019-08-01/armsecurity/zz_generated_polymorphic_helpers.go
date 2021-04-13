// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import "encoding/json"

func unmarshalAllowlistCustomAlertRuleClassification(rawMsg *json.RawMessage) (AllowlistCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(*rawMsg, &m); err != nil {
		return nil, err
	}
	var b AllowlistCustomAlertRuleClassification
	switch m["ruleType"] {
	case "ConnectionFromIpNotAllowed":
		b = &ConnectionFromIPNotAllowed{}
	case "ConnectionToIpNotAllowed":
		b = &ConnectionToIPNotAllowed{}
	case "LocalUserNotAllowed":
		b = &LocalUserNotAllowed{}
	case "ProcessNotAllowed":
		b = &ProcessNotAllowed{}
	default:
		b = &AllowlistCustomAlertRule{}
	}
	return b, json.Unmarshal(*rawMsg, &b)
}

func unmarshalAllowlistCustomAlertRuleClassificationArray(rawMsg *json.RawMessage) (*[]AllowlistCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(*rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]AllowlistCustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalAllowlistCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalCustomAlertRuleClassification(rawMsg *json.RawMessage) (CustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(*rawMsg, &m); err != nil {
		return nil, err
	}
	var b CustomAlertRuleClassification
	switch m["ruleType"] {
	case "ActiveConnectionsNotInAllowedRange":
		b = &ActiveConnectionsNotInAllowedRange{}
	case "AllowlistCustomAlertRule":
		b = &AllowlistCustomAlertRule{}
	case "AmqpC2DMessagesNotInAllowedRange":
		b = &AmqpC2DMessagesNotInAllowedRange{}
	case "AmqpC2DRejectedMessagesNotInAllowedRange":
		b = &AmqpC2DRejectedMessagesNotInAllowedRange{}
	case "AmqpD2CMessagesNotInAllowedRange":
		b = &AmqpD2CMessagesNotInAllowedRange{}
	case "ConnectionFromIpNotAllowed":
		b = &ConnectionFromIPNotAllowed{}
	case "ConnectionToIpNotAllowed":
		b = &ConnectionToIPNotAllowed{}
	case "DenylistCustomAlertRule":
		b = &DenylistCustomAlertRule{}
	case "DirectMethodInvokesNotInAllowedRange":
		b = &DirectMethodInvokesNotInAllowedRange{}
	case "FailedLocalLoginsNotInAllowedRange":
		b = &FailedLocalLoginsNotInAllowedRange{}
	case "FileUploadsNotInAllowedRange":
		b = &FileUploadsNotInAllowedRange{}
	case "HttpC2DMessagesNotInAllowedRange":
		b = &HTTPC2DMessagesNotInAllowedRange{}
	case "HttpC2DRejectedMessagesNotInAllowedRange":
		b = &HTTPC2DRejectedMessagesNotInAllowedRange{}
	case "HttpD2CMessagesNotInAllowedRange":
		b = &HTTPD2CMessagesNotInAllowedRange{}
	case "ListCustomAlertRule":
		b = &ListCustomAlertRule{}
	case "LocalUserNotAllowed":
		b = &LocalUserNotAllowed{}
	case "MqttC2DMessagesNotInAllowedRange":
		b = &MqttC2DMessagesNotInAllowedRange{}
	case "MqttC2DRejectedMessagesNotInAllowedRange":
		b = &MqttC2DRejectedMessagesNotInAllowedRange{}
	case "MqttD2CMessagesNotInAllowedRange":
		b = &MqttD2CMessagesNotInAllowedRange{}
	case "ProcessNotAllowed":
		b = &ProcessNotAllowed{}
	case "QueuePurgesNotInAllowedRange":
		b = &QueuePurgesNotInAllowedRange{}
	case "ThresholdCustomAlertRule":
		b = &ThresholdCustomAlertRule{}
	case "TimeWindowCustomAlertRule":
		b = &TimeWindowCustomAlertRule{}
	case "TwinUpdatesNotInAllowedRange":
		b = &TwinUpdatesNotInAllowedRange{}
	case "UnauthorizedOperationsNotInAllowedRange":
		b = &UnauthorizedOperationsNotInAllowedRange{}
	default:
		b = &CustomAlertRule{}
	}
	return b, json.Unmarshal(*rawMsg, &b)
}

func unmarshalCustomAlertRuleClassificationArray(rawMsg *json.RawMessage) (*[]CustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(*rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]CustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalListCustomAlertRuleClassification(rawMsg *json.RawMessage) (ListCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(*rawMsg, &m); err != nil {
		return nil, err
	}
	var b ListCustomAlertRuleClassification
	switch m["ruleType"] {
	case "AllowlistCustomAlertRule":
		b = &AllowlistCustomAlertRule{}
	case "ConnectionFromIpNotAllowed":
		b = &ConnectionFromIPNotAllowed{}
	case "ConnectionToIpNotAllowed":
		b = &ConnectionToIPNotAllowed{}
	case "DenylistCustomAlertRule":
		b = &DenylistCustomAlertRule{}
	case "LocalUserNotAllowed":
		b = &LocalUserNotAllowed{}
	case "ProcessNotAllowed":
		b = &ProcessNotAllowed{}
	default:
		b = &ListCustomAlertRule{}
	}
	return b, json.Unmarshal(*rawMsg, &b)
}

func unmarshalListCustomAlertRuleClassificationArray(rawMsg *json.RawMessage) (*[]ListCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(*rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ListCustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalListCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalThresholdCustomAlertRuleClassification(rawMsg *json.RawMessage) (ThresholdCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(*rawMsg, &m); err != nil {
		return nil, err
	}
	var b ThresholdCustomAlertRuleClassification
	switch m["ruleType"] {
	case "ActiveConnectionsNotInAllowedRange":
		b = &ActiveConnectionsNotInAllowedRange{}
	case "AmqpC2DMessagesNotInAllowedRange":
		b = &AmqpC2DMessagesNotInAllowedRange{}
	case "AmqpC2DRejectedMessagesNotInAllowedRange":
		b = &AmqpC2DRejectedMessagesNotInAllowedRange{}
	case "AmqpD2CMessagesNotInAllowedRange":
		b = &AmqpD2CMessagesNotInAllowedRange{}
	case "DirectMethodInvokesNotInAllowedRange":
		b = &DirectMethodInvokesNotInAllowedRange{}
	case "FailedLocalLoginsNotInAllowedRange":
		b = &FailedLocalLoginsNotInAllowedRange{}
	case "FileUploadsNotInAllowedRange":
		b = &FileUploadsNotInAllowedRange{}
	case "HttpC2DMessagesNotInAllowedRange":
		b = &HTTPC2DMessagesNotInAllowedRange{}
	case "HttpC2DRejectedMessagesNotInAllowedRange":
		b = &HTTPC2DRejectedMessagesNotInAllowedRange{}
	case "HttpD2CMessagesNotInAllowedRange":
		b = &HTTPD2CMessagesNotInAllowedRange{}
	case "MqttC2DMessagesNotInAllowedRange":
		b = &MqttC2DMessagesNotInAllowedRange{}
	case "MqttC2DRejectedMessagesNotInAllowedRange":
		b = &MqttC2DRejectedMessagesNotInAllowedRange{}
	case "MqttD2CMessagesNotInAllowedRange":
		b = &MqttD2CMessagesNotInAllowedRange{}
	case "QueuePurgesNotInAllowedRange":
		b = &QueuePurgesNotInAllowedRange{}
	case "TimeWindowCustomAlertRule":
		b = &TimeWindowCustomAlertRule{}
	case "TwinUpdatesNotInAllowedRange":
		b = &TwinUpdatesNotInAllowedRange{}
	case "UnauthorizedOperationsNotInAllowedRange":
		b = &UnauthorizedOperationsNotInAllowedRange{}
	default:
		b = &ThresholdCustomAlertRule{}
	}
	return b, json.Unmarshal(*rawMsg, &b)
}

func unmarshalThresholdCustomAlertRuleClassificationArray(rawMsg *json.RawMessage) (*[]ThresholdCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(*rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ThresholdCustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalThresholdCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalTimeWindowCustomAlertRuleClassification(rawMsg *json.RawMessage) (TimeWindowCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(*rawMsg, &m); err != nil {
		return nil, err
	}
	var b TimeWindowCustomAlertRuleClassification
	switch m["ruleType"] {
	case "ActiveConnectionsNotInAllowedRange":
		b = &ActiveConnectionsNotInAllowedRange{}
	case "AmqpC2DMessagesNotInAllowedRange":
		b = &AmqpC2DMessagesNotInAllowedRange{}
	case "AmqpC2DRejectedMessagesNotInAllowedRange":
		b = &AmqpC2DRejectedMessagesNotInAllowedRange{}
	case "AmqpD2CMessagesNotInAllowedRange":
		b = &AmqpD2CMessagesNotInAllowedRange{}
	case "DirectMethodInvokesNotInAllowedRange":
		b = &DirectMethodInvokesNotInAllowedRange{}
	case "FailedLocalLoginsNotInAllowedRange":
		b = &FailedLocalLoginsNotInAllowedRange{}
	case "FileUploadsNotInAllowedRange":
		b = &FileUploadsNotInAllowedRange{}
	case "HttpC2DMessagesNotInAllowedRange":
		b = &HTTPC2DMessagesNotInAllowedRange{}
	case "HttpC2DRejectedMessagesNotInAllowedRange":
		b = &HTTPC2DRejectedMessagesNotInAllowedRange{}
	case "HttpD2CMessagesNotInAllowedRange":
		b = &HTTPD2CMessagesNotInAllowedRange{}
	case "MqttC2DMessagesNotInAllowedRange":
		b = &MqttC2DMessagesNotInAllowedRange{}
	case "MqttC2DRejectedMessagesNotInAllowedRange":
		b = &MqttC2DRejectedMessagesNotInAllowedRange{}
	case "MqttD2CMessagesNotInAllowedRange":
		b = &MqttD2CMessagesNotInAllowedRange{}
	case "QueuePurgesNotInAllowedRange":
		b = &QueuePurgesNotInAllowedRange{}
	case "TwinUpdatesNotInAllowedRange":
		b = &TwinUpdatesNotInAllowedRange{}
	case "UnauthorizedOperationsNotInAllowedRange":
		b = &UnauthorizedOperationsNotInAllowedRange{}
	default:
		b = &TimeWindowCustomAlertRule{}
	}
	return b, json.Unmarshal(*rawMsg, &b)
}

func unmarshalTimeWindowCustomAlertRuleClassificationArray(rawMsg *json.RawMessage) (*[]TimeWindowCustomAlertRuleClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(*rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]TimeWindowCustomAlertRuleClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalTimeWindowCustomAlertRuleClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func strptr(s string) *string {
	return &s
}
