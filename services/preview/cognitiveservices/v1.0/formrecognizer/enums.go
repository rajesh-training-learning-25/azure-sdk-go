package formrecognizer

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Status enumerates the values for status.
type Status string

const (
	// Failure ...
	Failure Status = "failure"
	// PartialSuccess ...
	PartialSuccess Status = "partialSuccess"
	// Success ...
	Success Status = "success"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Failure, PartialSuccess, Success}
}

// Status1 enumerates the values for status 1.
type Status1 string

const (
	// Created ...
	Created Status1 = "created"
	// Invalid ...
	Invalid Status1 = "invalid"
	// Ready ...
	Ready Status1 = "ready"
)

// PossibleStatus1Values returns an array of possible values for the Status1 const type.
func PossibleStatus1Values() []Status1 {
	return []Status1{Created, Invalid, Ready}
}

// Status2 enumerates the values for status 2.
type Status2 string

const (
	// Status2Failure ...
	Status2Failure Status2 = "failure"
	// Status2PartialSuccess ...
	Status2PartialSuccess Status2 = "partialSuccess"
	// Status2Success ...
	Status2Success Status2 = "success"
)

// PossibleStatus2Values returns an array of possible values for the Status2 const type.
func PossibleStatus2Values() []Status2 {
	return []Status2{Status2Failure, Status2PartialSuccess, Status2Success}
}

// TextOperationStatusCodes enumerates the values for text operation status codes.
type TextOperationStatusCodes string

const (
	// Failed ...
	Failed TextOperationStatusCodes = "Failed"
	// NotStarted ...
	NotStarted TextOperationStatusCodes = "Not Started"
	// Running ...
	Running TextOperationStatusCodes = "Running"
	// Succeeded ...
	Succeeded TextOperationStatusCodes = "Succeeded"
)

// PossibleTextOperationStatusCodesValues returns an array of possible values for the TextOperationStatusCodes const type.
func PossibleTextOperationStatusCodesValues() []TextOperationStatusCodes {
	return []TextOperationStatusCodes{Failed, NotStarted, Running, Succeeded}
}

// TextRecognitionResultConfidenceClass enumerates the values for text recognition result confidence class.
type TextRecognitionResultConfidenceClass string

const (
	// High ...
	High TextRecognitionResultConfidenceClass = "High"
	// Low ...
	Low TextRecognitionResultConfidenceClass = "Low"
)

// PossibleTextRecognitionResultConfidenceClassValues returns an array of possible values for the TextRecognitionResultConfidenceClass const type.
func PossibleTextRecognitionResultConfidenceClassValues() []TextRecognitionResultConfidenceClass {
	return []TextRecognitionResultConfidenceClass{High, Low}
}

// TextRecognitionResultDimensionUnit enumerates the values for text recognition result dimension unit.
type TextRecognitionResultDimensionUnit string

const (
	// Inch ...
	Inch TextRecognitionResultDimensionUnit = "inch"
	// Pixel ...
	Pixel TextRecognitionResultDimensionUnit = "pixel"
)

// PossibleTextRecognitionResultDimensionUnitValues returns an array of possible values for the TextRecognitionResultDimensionUnit const type.
func PossibleTextRecognitionResultDimensionUnitValues() []TextRecognitionResultDimensionUnit {
	return []TextRecognitionResultDimensionUnit{Inch, Pixel}
}

// ValueType enumerates the values for value type.
type ValueType string

const (
	// ValueTypeFieldValue ...
	ValueTypeFieldValue ValueType = "fieldValue"
	// ValueTypeNumberValue ...
	ValueTypeNumberValue ValueType = "numberValue"
	// ValueTypeStringValue ...
	ValueTypeStringValue ValueType = "stringValue"
)

// PossibleValueTypeValues returns an array of possible values for the ValueType const type.
func PossibleValueTypeValues() []ValueType {
	return []ValueType{ValueTypeFieldValue, ValueTypeNumberValue, ValueTypeStringValue}
}
