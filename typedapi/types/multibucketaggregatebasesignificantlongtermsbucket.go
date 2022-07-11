// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// MultiBucketAggregateBaseSignificantLongTermsBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/Aggregate.ts#L314-L316
type MultiBucketAggregateBaseSignificantLongTermsBucket struct {
	Buckets BucketsSignificantLongTermsBucket `json:"buckets"`
	Meta    *Metadata                         `json:"meta,omitempty"`
}

// MultiBucketAggregateBaseSignificantLongTermsBucketBuilder holds MultiBucketAggregateBaseSignificantLongTermsBucket struct and provides a builder API.
type MultiBucketAggregateBaseSignificantLongTermsBucketBuilder struct {
	v *MultiBucketAggregateBaseSignificantLongTermsBucket
}

// NewMultiBucketAggregateBaseSignificantLongTermsBucket provides a builder for the MultiBucketAggregateBaseSignificantLongTermsBucket struct.
func NewMultiBucketAggregateBaseSignificantLongTermsBucketBuilder() *MultiBucketAggregateBaseSignificantLongTermsBucketBuilder {
	r := MultiBucketAggregateBaseSignificantLongTermsBucketBuilder{
		&MultiBucketAggregateBaseSignificantLongTermsBucket{},
	}

	return &r
}

// Build finalize the chain and returns the MultiBucketAggregateBaseSignificantLongTermsBucket struct
func (rb *MultiBucketAggregateBaseSignificantLongTermsBucketBuilder) Build() MultiBucketAggregateBaseSignificantLongTermsBucket {
	return *rb.v
}

func (rb *MultiBucketAggregateBaseSignificantLongTermsBucketBuilder) Buckets(buckets *BucketsSignificantLongTermsBucketBuilder) *MultiBucketAggregateBaseSignificantLongTermsBucketBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *MultiBucketAggregateBaseSignificantLongTermsBucketBuilder) Meta(meta *MetadataBuilder) *MultiBucketAggregateBaseSignificantLongTermsBucketBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}