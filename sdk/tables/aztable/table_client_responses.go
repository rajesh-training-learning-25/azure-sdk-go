// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztable

import "time"

type TableAddEntityResponse struct {
	ClientRequestID                 string
	RequestID                       string
	Version                         string
	Date                            time.Time
	ETag                            string
	XMSContinuationNextPartitionKey string
	XMSContinuationNextRowKey       string
	PreferenceApplied               string
	ContentType                     string
}
