// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20200217

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
)

type AcknowledgeMessageRequest struct {
	*tchttp.BaseRequest

	// Unique ID used to identify the message, which can be obtained from the returned value of `receiveMessage`.
	MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

	// Topic name, which can be obtained from the returned value of `receiveMessage` and is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	AckTopic *string `json:"AckTopic,omitempty" name:"AckTopic"`

	// Subscriber name, which can be obtained from the returned value of `receiveMessage`. Make sure that it is the same as the subscriber name identified in `receiveMessage`; otherwise, the received message cannot be correctly acknowledged.
	SubName *string `json:"SubName,omitempty" name:"SubName"`
}

func (r *AcknowledgeMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcknowledgeMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "MessageId")
	delete(f, "AckTopic")
	delete(f, "SubName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "AcknowledgeMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type AcknowledgeMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// If it is an empty string, no error occurred.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcknowledgeMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *AcknowledgeMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearCmqQueueRequest struct {
	*tchttp.BaseRequest

	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *ClearCmqQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearCmqQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearCmqQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ClearCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearCmqQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearCmqSubscriptionFilterTagsRequest struct {
	*tchttp.BaseRequest

	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *ClearCmqSubscriptionFilterTagsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearCmqSubscriptionFilterTagsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ClearCmqSubscriptionFilterTagsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ClearCmqSubscriptionFilterTagsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearCmqSubscriptionFilterTagsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ClearCmqSubscriptionFilterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Number of access points
	EndPointNum *int64 `json:"EndPointNum,omitempty" name:"EndPointNum"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Whether the cluster is healthy. 1: healthy; 0: exceptional
	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`

	// Cluster health information
	// Note: this field may return null, indicating that no valid values can be obtained.
	HealthyInfo *string `json:"HealthyInfo,omitempty" name:"HealthyInfo"`

	// Cluster status. 0: creating; 1: normal; 2: terminating; 3: deleted; 4. isolated; 5. creation failed; 6: deletion failed
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// Maximum number of namespaces
	MaxNamespaceNum *int64 `json:"MaxNamespaceNum,omitempty" name:"MaxNamespaceNum"`

	// Maximum number of topics
	MaxTopicNum *int64 `json:"MaxTopicNum,omitempty" name:"MaxTopicNum"`

	// Maximum QPS
	MaxQps *int64 `json:"MaxQps,omitempty" name:"MaxQps"`

	// Maximum message retention period in seconds
	MessageRetentionTime *int64 `json:"MessageRetentionTime,omitempty" name:"MessageRetentionTime"`

	// Maximum storage capacity
	MaxStorageCapacity *int64 `json:"MaxStorageCapacity,omitempty" name:"MaxStorageCapacity"`

	// Cluster version
	// Note: this field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitempty" name:"Version"`

	// Public network access point
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`

	// VPC access point
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`

	// Number of namespaces
	// Note: this field may return null, indicating that no valid values can be obtained.
	NamespaceNum *int64 `json:"NamespaceNum,omitempty" name:"NamespaceNum"`

	// Limit of used storage in MB
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsedStorageBudget *int64 `json:"UsedStorageBudget,omitempty" name:"UsedStorageBudget"`

	// Maximum message production rate in messages
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxPublishRateInMessages *int64 `json:"MaxPublishRateInMessages,omitempty" name:"MaxPublishRateInMessages"`

	// Maximum message push rate in messages
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxDispatchRateInMessages *int64 `json:"MaxDispatchRateInMessages,omitempty" name:"MaxDispatchRateInMessages"`

	// Maximum message production rate in bytes
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxPublishRateInBytes *int64 `json:"MaxPublishRateInBytes,omitempty" name:"MaxPublishRateInBytes"`

	// Maximum message push rate in bytes
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxDispatchRateInBytes *int64 `json:"MaxDispatchRateInBytes,omitempty" name:"MaxDispatchRateInBytes"`

	// Number of created topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`

	// Maximum message delay in seconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMessageDelayInSeconds *int64 `json:"MaxMessageDelayInSeconds,omitempty" name:"MaxMessageDelayInSeconds"`

	// Whether to enable public network access. If this parameter is left empty, the feature will be enabled by default
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" name:"PublicAccessEnabled"`

	// Tag
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Billing mode:
	// `0`: Pay-as-you-go
	// `1`: Monthly subscription
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
}

type CmqDeadLetterPolicy struct {

	// Dead letter queue.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterQueue *string `json:"DeadLetterQueue,omitempty" name:"DeadLetterQueue"`

	// Dead letter queue policy.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// Maximum period in seconds before an unconsumed message expires, which is required if `Policy` is 1. Value range: 300–43200. This value should be smaller than `MsgRetentionSeconds` (maximum message retention period)
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// Maximum number of receipts.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`
}

type CmqDeadLetterSource struct {

	// Message queue ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// Message queue name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type CmqQueue struct {

	// Message queue ID.
	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

	// Message queue name.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// Limit of the number of messages produced per second. The value for consumed messages is 1.1 times this value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// Bandwidth limit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bps *uint64 `json:"Bps,omitempty" name:"Bps"`

	// Maximum retention period for inflight messages.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxDelaySeconds *uint64 `json:"MaxDelaySeconds,omitempty" name:"MaxDelaySeconds"`

	// Maximum number of heaped messages. The value range is 1,000,000–10,000,000 during the beta test and can be 1,000,000–1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// Long polling wait time for message reception. Value range: 0–30 seconds. Default value: 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// Message retention period. Value range: 60–1296000 seconds (i.e., 1 minute–15 days). Default value: 345600 (i.e., 4 days).
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// Message visibility timeout period. Value range: 1–43200 seconds (i.e., 12 hours). Default value: 30.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// Maximum message length. Value range: 1024–1048576 bytes (i.e., 1–1024 KB). Default value: 65536.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// Maximum time range during which a message can be rewound in the queue, which ranges from 0 to 43,200 seconds. 0 indicates that message rewind is disabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// Queue creation time. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Time when the queue attribute is last modified. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// Total number of messages in `Active` status (i.e., unconsumed) in the queue, which is an approximate value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ActiveMsgNum *uint64 `json:"ActiveMsgNum,omitempty" name:"ActiveMsgNum"`

	// Total number of messages in `Inactive` status (i.e., being consumed) in the queue, which is an approximate value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InactiveMsgNum *uint64 `json:"InactiveMsgNum,omitempty" name:"InactiveMsgNum"`

	// Number of delayed messages.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DelayMsgNum *uint64 `json:"DelayMsgNum,omitempty" name:"DelayMsgNum"`

	// Number of retained messages which have been deleted by the `DelMsg` API but are still within their rewind time range.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RewindMsgNum *uint64 `json:"RewindMsgNum,omitempty" name:"RewindMsgNum"`

	// Minimum unconsumed time of message in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MinMsgTime *uint64 `json:"MinMsgTime,omitempty" name:"MinMsgTime"`

	// Transaction message queue. true: transaction message type; false: other message types.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Transaction *bool `json:"Transaction,omitempty" name:"Transaction"`

	// Dead letter queue.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterSource []*CmqDeadLetterSource `json:"DeadLetterSource,omitempty" name:"DeadLetterSource"`

	// Dead letter queue policy.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterPolicy *CmqDeadLetterPolicy `json:"DeadLetterPolicy,omitempty" name:"DeadLetterPolicy"`

	// Transaction message policy.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TransactionPolicy *CmqTransactionPolicy `json:"TransactionPolicy,omitempty" name:"TransactionPolicy"`

	// Creator `Uin`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Associated tag.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Message trace. true: enabled; false: not enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Trace *bool `json:"Trace,omitempty" name:"Trace"`

	// Tenant ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`

	// Namespace name
	// Note: this field may return null, indicating that no valid values can be obtained.
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Cluster status. 0: creating; 1: normal; 2: terminating; 3: deleted; 4. isolated; 5. creation failed; 6: deletion failed
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// The maximum number of unacknowledged messages.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MaxUnackedMsgNum *int64 `json:"MaxUnackedMsgNum,omitempty" name:"MaxUnackedMsgNum"`

	// Maximum size of heaped messages in bytes.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MaxMsgBacklogSize *int64 `json:"MaxMsgBacklogSize,omitempty" name:"MaxMsgBacklogSize"`

	// Queue storage space configured for message rewind. Value range: 1,024-10,240 MB (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitempty" name:"RetentionSizeInMB"`
}

type CmqSubscription struct {

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// Subscription ID, which will be used during monitoring data pull.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`

	// Subscription owner `APPID`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicOwner *uint64 `json:"TopicOwner,omitempty" name:"TopicOwner"`

	// Number of messages to be delivered in the subscription.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`

	// Time when the subscription attribute is last modified. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// Subscription creation time. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Filtering policy for subscribing to and receiving messages.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`

	// Endpoint that receives notifications, which varies by `protocol`: for HTTP, the endpoint must start with `http://`, and the `host` can be a domain or IP; for `queue`, `queueName` should be entered.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// Filtering policy selected when a subscription is created:
	// If `filterType` is 1, `filterTag` will be used for filtering.
	// If `filterType` is 2, `bindingKey` will be used for filtering.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags"`

	// Subscription protocol. Currently, two protocols are supported: HTTP and queue. To use the HTTP protocol, you need to build your own web server to receive messages. With the queue protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// CMQ push server retry policy in case an error occurs while pushing a message to the endpoint. Valid values:
	// (1) BACKOFF_RETRY: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message.
	// (2) EXPONENTIAL_DECAY_RETRY: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: EXPONENTIAL_DECAY_RETRY.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`

	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `protocol` is `queue`, this value must be `SIMPLIFIED`. If `protocol` is `HTTP`, both values are acceptable, and the default value is `JSON`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

type CmqTopic struct {

	// Topic ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

	// Topic name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Maximum lifecycle of message in topic. After the period specified by this parameter has elapsed since a message is sent to the topic, the message will be deleted no matter whether it has been successfully pushed to the user. This parameter is measured in seconds and defaulted to one day (86,400 seconds), which cannot be modified.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// Maximum message size, which ranges from 1,024 to 1,048,576 bytes (i.e., 1–1,024 KB). The default value is 65,536.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// Number of messages published per second.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`

	// Filtering policy selected when a subscription is created:
	// If `filterType` is 1, `FilterTag` will be used for filtering.
	// If `filterType` is 2, `BindingKey` will be used for filtering.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`

	// Topic creation time. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Time when the topic attribute is last modified. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`

	// Number of current messages in the topic (number of retained messages).
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`

	// Creator `Uin`. The `resource` field for CAM authentication is composed of this field.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`

	// Associated tag.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Message trace. true: enabled; false: not enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Trace *bool `json:"Trace,omitempty" name:"Trace"`

	// Tenant ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`

	// Namespace name
	// Note: this field may return null, indicating that no valid values can be obtained.
	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`

	// Cluster status. 0: creating; 1: normal; 2: terminating; 3: deleted; 4. isolated; 5. creation failed; 6: deletion failed
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type CmqTransactionPolicy struct {

	// First lookback time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`

	// Maximum number of queries.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// Cluster name, which can contain up to 16 letters, digits, hyphens, and underscores.
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// ID of your dedicated physical cluster. If it is not passed in, cluster resources will be created in a public cluster by default.
	BindClusterId *uint64 `json:"BindClusterId,omitempty" name:"BindClusterId"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Cluster tag list (deprecated).
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Whether to enable public network access. If this parameter is left empty, the feature will be enabled by default
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" name:"PublicAccessEnabled"`
}

func (r *CreateClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterName")
	delete(f, "BindClusterId")
	delete(f, "Remark")
	delete(f, "Tags")
	delete(f, "PublicAccessEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cluster ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqQueueRequest struct {
	*tchttp.BaseRequest

	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// Maximum number of heaped messages. The value range is 1,000,000–10,000,000 during the beta test and can be 1,000,000–1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// Long polling wait time for message reception. Value range: 0–30 seconds. Default value: 0.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// Message visibility timeout period. Value range: 1–43200 seconds (i.e., 12 hours). Default value: 30.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// 1: transaction queue; 0: general queue
	Transaction *uint64 `json:"Transaction,omitempty" name:"Transaction"`

	// First lookback interval
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`

	// Maximum number of lookbacks
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`

	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// Dead letter policy. 0: message has been consumed multiple times but not deleted; 1: `Time-To-Live` has elapsed
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// Maximum receipt times. Value range: 1–1000
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`

	// Maximum period in seconds before an unconsumed message expires, which is required if `policy` is 1. Value range: 300–43200. This value should be smaller than `msgRetentionSeconds` (maximum message retention period)
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// Whether to enable message trace. true: yes; false: no. If this field is not configured, the feature will not be enabled
	Trace *bool `json:"Trace,omitempty" name:"Trace"`

	// Tag array.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`

	// Queue storage space configured for message rewind. Value range: 1,024-10,240 MB (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitempty" name:"RetentionSizeInMB"`
}

func (r *CreateCmqQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "MaxMsgHeapNum")
	delete(f, "PollingWaitSeconds")
	delete(f, "VisibilityTimeout")
	delete(f, "MaxMsgSize")
	delete(f, "MsgRetentionSeconds")
	delete(f, "RewindSeconds")
	delete(f, "Transaction")
	delete(f, "FirstQueryInterval")
	delete(f, "MaxQueryCount")
	delete(f, "DeadLetterQueueName")
	delete(f, "Policy")
	delete(f, "MaxReceiveCount")
	delete(f, "MaxTimeToLive")
	delete(f, "Trace")
	delete(f, "Tags")
	delete(f, "RetentionSizeInMB")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCmqQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// `queueId` of a successfully created queue
		QueueId *string `json:"QueueId,omitempty" name:"QueueId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmqQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqSubscribeRequest struct {
	*tchttp.BaseRequest

	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// Subscription protocol. Currently, two protocols are supported: HTTP and queue. To use the HTTP protocol, you need to build your own web server to receive messages. With the queue protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages will not be received. Currently, messages can be pushed only to the public network and classic network.
	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`

	// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1. BACKOFF_RETRY: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message; 2. EXPONENTIAL_DECAY_RETRY: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: EXPONENTIAL_DECAY_RETRY.
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`

	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTag []*string `json:"FilterTag,omitempty" name:"FilterTag"`

	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`

	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

func (r *CreateCmqSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	delete(f, "Protocol")
	delete(f, "Endpoint")
	delete(f, "NotifyStrategy")
	delete(f, "FilterTag")
	delete(f, "BindingKey")
	delete(f, "NotifyContentFormat")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCmqSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Subscription ID
		SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmqSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqTopicRequest struct {
	*tchttp.BaseRequest

	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// Used to specify the message match policy for the topic. 1: tag match policy (default value); 2: routing match policy.
	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`

	// Message retention period. Value range: 60–86400 seconds (i.e., 1 minute–1 day). Default value: 86400.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitempty" name:"Trace"`

	// Tag array.
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateCmqTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "MaxMsgSize")
	delete(f, "FilterType")
	delete(f, "MsgRetentionSeconds")
	delete(f, "Trace")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCmqTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Topic ID
		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmqTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCmqTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest

	// Environment (namespace) name, which can contain up to 16 letters, digits, hyphens, and underscores.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Unconsumed message expiration time in seconds. Minimum value: 60; maximum value: 1296000 (15 days).
	MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Message retention policy
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
}

func (r *CreateEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "MsgTTL")
	delete(f, "Remark")
	delete(f, "ClusterId")
	delete(f, "RetentionPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Environment (namespace) name.
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// TTL for unconsumed messages in seconds.
		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

		// Remarks (up to 128 characters).
	// Note: this field may return null, indicating that no valid values can be obtained.
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// Namespace ID
		NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentRoleRequest struct {
	*tchttp.BaseRequest

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Role name.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateEnvironmentRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "RoleName")
	delete(f, "Permissions")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEnvironmentRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateEnvironmentRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// Cluster name, which can contain 3–64 letters, digits, hyphens, and underscores
	Name *string `json:"Name,omitempty" name:"Name"`

	// Cluster description (up to 128 characters)
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateRocketMQClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Name")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cluster ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQGroupRequest struct {
	*tchttp.BaseRequest

	// Group name (8–64 characters)
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Namespace. Currently, only one namespace is supported
	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`

	// Whether to enable consumption
	ReadEnable *bool `json:"ReadEnable,omitempty" name:"ReadEnable"`

	// Whether to enable broadcast consumption
	BroadcastEnable *bool `json:"BroadcastEnable,omitempty" name:"BroadcastEnable"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateRocketMQGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "GroupId")
	delete(f, "Namespaces")
	delete(f, "ReadEnable")
	delete(f, "BroadcastEnable")
	delete(f, "ClusterId")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace name, which can contain 3–64 letters, digits, hyphens, and underscores
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Retention time of unconsumed messages in milliseconds. Value range: 60 seconds–15 days
	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`

	// Retention time of persisted messages in milliseconds
	RetentionTime *uint64 `json:"RetentionTime,omitempty" name:"RetentionTime"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateRocketMQNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Ttl")
	delete(f, "RetentionTime")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQTopicRequest struct {
	*tchttp.BaseRequest

	// Topic name, which can contain 3–64 letters, digits, hyphens, and underscores
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Topic namespace. Currently, you can create topics only in one single namespace.
	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`

	// Topic type. Valid values: Normal, GlobalOrder, PartitionedOrder.
	Type *string `json:"Type,omitempty" name:"Type"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Topic remarks (up to 128 characters)
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Number of partitions, which doesn't take effect for globally sequential messages
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
}

func (r *CreateRocketMQTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topic")
	delete(f, "Namespaces")
	delete(f, "Type")
	delete(f, "ClusterId")
	delete(f, "Remark")
	delete(f, "PartitionNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRocketMQTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest

	// Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	delete(f, "Remark")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Role name
		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

		// Role token
		Token *string `json:"Token,omitempty" name:"Token"`

		// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionRequest struct {
	*tchttp.BaseRequest

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscriber name, which can contain up to 150 letters, digits, hyphens, and underscores.
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// Whether the creation is idempotent; if not, you cannot create subscriptions with the same name.
	IsIdempotent *bool `json:"IsIdempotent,omitempty" name:"IsIdempotent"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Whether to automatically create a dead letter topic and a retry letter topic. true: yes (default value); false: no.
	AutoCreatePolicyTopic *bool `json:"AutoCreatePolicyTopic,omitempty" name:"AutoCreatePolicyTopic"`

	// Naming convention for dead letter and retry letter topics. `LEGACY` indicates to use the legacy naming convention, and `COMMUNITY` indicates to use the naming convention in the Pulsar community.
	PostFixPattern *string `json:"PostFixPattern,omitempty" name:"PostFixPattern"`
}

func (r *CreateSubscriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscriptionRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	delete(f, "IsIdempotent")
	delete(f, "Remark")
	delete(f, "ClusterId")
	delete(f, "AutoCreatePolicyTopic")
	delete(f, "PostFixPattern")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateSubscriptionRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Creation result.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Topic name, which can contain up to 64 letters, digits, hyphens, and underscores.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 0: non-partitioned topic; other values: number of partitions in the partitioned topic (up to 128).
	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 0: general message;
	// 1: globally sequential message;
	// 2: partitionally sequential message;
	// 3: retry letter queue;
	// 4: dead letter queue.
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Pulsar topic type.
	// `0`: Non-persistent and non-partitioned
	// `1`: Non-persistent and partitioned
	// `2`: Persistent and non-partitioned
	// `3`: Persistent and partitioned
	PulsarTopicType *int64 `json:"PulsarTopicType,omitempty" name:"PulsarTopicType"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Partitions")
	delete(f, "Remark")
	delete(f, "TopicType")
	delete(f, "ClusterId")
	delete(f, "PulsarTopicType")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Environment (namespace) name.
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// Topic name.
		TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

		// 0: non-partitioned topic; other values: number of partitions in the partitioned topic.
		Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

		// Remarks (up to 128 characters).
	// Note: this field may return null, indicating that no valid values can be obtained.
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 0: general message;
	// 1: globally sequential message;
	// 2: partitionally sequential message;
	// 3: retry letter queue;
	// 4: dead letter queue;
	// 5: transaction message.
	// Note: this field may return null, indicating that no valid values can be obtained.
		TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster to be deleted.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cluster ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqQueueRequest struct {
	*tchttp.BaseRequest

	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DeleteCmqQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCmqQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmqQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqSubscribeRequest struct {
	*tchttp.BaseRequest

	// Topic name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *DeleteCmqSubscribeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqSubscribeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCmqSubscribeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmqSubscribeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqTopicRequest struct {
	*tchttp.BaseRequest

	// Topic name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DeleteCmqTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCmqTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmqTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCmqTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentRolesRequest struct {
	*tchttp.BaseRequest

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Array of role names.
	RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteEnvironmentRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "RoleNames")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEnvironmentRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentRolesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEnvironmentRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// Array of environments (namespaces). Up to 20 environments can be deleted at a time.
	EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentIds")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Array of environments (namespaces) successfully deleted.
		EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// ID of the cluster to be deleted.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteRocketMQClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQGroupRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Consumer group name
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteRocketMQGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DeleteRocketMQNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQTopicRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Topic name
	Topic *string `json:"Topic,omitempty" name:"Topic"`
}

func (r *DeleteRocketMQTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRocketMQTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRocketMQTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRolesRequest struct {
	*tchttp.BaseRequest

	// Array of role names.
	RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleNames")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRolesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Name array of roles successfully deleted.
		RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindVpcsRequest struct {
	*tchttp.BaseRequest

	// Offset. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeBindVpcsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindVpcsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindVpcsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindVpcsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of records.
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Set of VPCs.
		VpcSets []*VpcBindRecord `json:"VpcSets,omitempty" name:"VpcSets"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindVpcsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClusterDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cluster details
		ClusterSet *Cluster `json:"ClusterSet,omitempty" name:"ClusterSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqDeadLetterSourceQueuesRequest struct {
	*tchttp.BaseRequest

	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// Starting position of the list of topics to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Number of topics to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Filter by `SourceQueueName`
	SourceQueueName *string `json:"SourceQueueName,omitempty" name:"SourceQueueName"`
}

func (r *DescribeCmqDeadLetterSourceQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqDeadLetterSourceQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "DeadLetterQueueName")
	delete(f, "Limit")
	delete(f, "Offset")
	delete(f, "SourceQueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqDeadLetterSourceQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqDeadLetterSourceQueuesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of eligible queues
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Source queues of dead letter queue
		QueueSet []*CmqDeadLetterSource `json:"QueueSet,omitempty" name:"QueueSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqDeadLetterSourceQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqDeadLetterSourceQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqQueueDetailRequest struct {
	*tchttp.BaseRequest

	// Exact match by `QueueName`
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DescribeCmqQueueDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqQueueDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqQueueDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqQueueDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// List of queue details.
		QueueDescribe *CmqQueue `json:"QueueDescribe,omitempty" name:"QueueDescribe"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqQueueDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqQueueDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqSubscriptionDetailRequest struct {
	*tchttp.BaseRequest

	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Starting position of the list of topics to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// Number of topics to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// Fuzzy search by `SubscriptionName`
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *DescribeCmqSubscriptionDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqSubscriptionDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SubscriptionName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqSubscriptionDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqSubscriptionDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Set of subscription attributes
	// Note: this field may return null, indicating that no valid values can be obtained.
		SubscriptionSet []*CmqSubscription `json:"SubscriptionSet,omitempty" name:"SubscriptionSet"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqSubscriptionDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqSubscriptionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqTopicDetailRequest struct {
	*tchttp.BaseRequest

	// Exact match by `TopicName`.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DescribeCmqTopicDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqTopicDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqTopicDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Topic details
		TopicDescribe *CmqTopic `json:"TopicDescribe,omitempty" name:"TopicDescribe"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqTopicDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqTopicDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeEnvironmentAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// TTL for unconsumed messages in seconds. Maximum value: 1296000 seconds (i.e., 15 days).
		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

		// Consumption rate limit in bytes/second. 0: unlimited.
		RateInByte *uint64 `json:"RateInByte,omitempty" name:"RateInByte"`

		// Consumption rate limit in messages/second. 0: unlimited.
		RateInSize *uint64 `json:"RateInSize,omitempty" name:"RateInSize"`

		// Retention policy for consumed messages in hours. 0: deleted immediately after consumption.
		RetentionHours *uint64 `json:"RetentionHours,omitempty" name:"RetentionHours"`

		// Retention policy for consumed messages in GB. 0: deleted immediately after consumption.
		RetentionSize *uint64 `json:"RetentionSize,omitempty" name:"RetentionSize"`

		// Environment (namespace) name.
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// Number of replicas.
		Replicas *uint64 `json:"Replicas,omitempty" name:"Replicas"`

		// Remarks.
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvironmentAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublisherSummaryRequest struct {
	*tchttp.BaseRequest

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace name.
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Topic name.
	Topic *string `json:"Topic,omitempty" name:"Topic"`
}

func (r *DescribePublisherSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublisherSummaryRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Namespace")
	delete(f, "Topic")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublisherSummaryRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublisherSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Production rate (messages/sec).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		MsgRateIn *float64 `json:"MsgRateIn,omitempty" name:"MsgRateIn"`

		// Production rate (byte/sec).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		MsgThroughputIn *float64 `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`

		// The number of producers.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		PublisherCount *int64 `json:"PublisherCount,omitempty" name:"PublisherCount"`

		// Message storage size in bytes.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublisherSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublisherSummaryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublishersRequest struct {
	*tchttp.BaseRequest

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace name.
	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`

	// Topic name.
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Parameter filter. The `ProducerName` and `Address` fields are supported.
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`

	// Offset for query. Default value: `0`.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// The number of query results displayed per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Sort by field.
	Sort *Sort `json:"Sort,omitempty" name:"Sort"`
}

func (r *DescribePublishersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublishersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Namespace")
	delete(f, "Topic")
	delete(f, "Filters")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Sort")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePublishersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribePublishersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Total number of query results.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// List of producer information.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
		Publishers []*Publisher `json:"Publishers,omitempty" name:"Publishers"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePublishersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePublishersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeRocketMQClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Cluster information
		ClusterInfo *RocketMQClusterInfo `json:"ClusterInfo,omitempty" name:"ClusterInfo"`

		// Cluster configuration
		ClusterConfig *RocketMQClusterConfig `json:"ClusterConfig,omitempty" name:"ClusterConfig"`

		// Recent cluster usage
	// Note: this field may return null, indicating that no valid values can be obtained.
		ClusterStats *RocketMQClusterRecentStats `json:"ClusterStats,omitempty" name:"ClusterStats"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRolesRequest struct {
	*tchttp.BaseRequest

	// Fuzzy query by role name
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Offset. If this parameter is left empty, 0 will be used by default.
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20.
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// * RoleName
	// Filter by role name for exact query.
	// Type: String
	// Required: no
	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRolesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of records.
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// Array of roles.
		RoleSets []*Role `json:"RoleSets,omitempty" name:"RoleSets"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// Filter parameter name
	Name *string `json:"Name,omitempty" name:"Name"`

	// Value
	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ModifyClusterRequest struct {
	*tchttp.BaseRequest

	// ID of the Pulsar cluster to be updated.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Updated cluster name.
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Enables public network access, which can only be `true`.
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" name:"PublicAccessEnabled"`
}

func (r *ModifyClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	delete(f, "Remark")
	delete(f, "PublicAccessEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Pulsar cluster ID
		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqQueueAttributeRequest struct {
	*tchttp.BaseRequest

	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// Maximum number of heaped messages. The value range is 1,000,000–10,000,000 during the beta test and can be 1,000,000–1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`

	// Long polling wait time for message reception. Value range: 0–30 seconds. Default value: 0.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`

	// Message visibility timeout period. Value range: 1–43200 seconds (i.e., 12 hours). Default value: 30.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`

	// Max message size, which defaults to 1,024 KB for the queue of TDMQ for CMQ and cannot be modified.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`

	// First query time
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`

	// Maximum number of queries
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`

	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`

	// Maximum period in seconds before an unconsumed message expires, which is required if `MaxTimeToLivepolicy` is 1. Value range: 300–43200. This value should be smaller than `MsgRetentionSeconds` (maximum message retention period)
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`

	// Maximum number of receipts
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`

	// Dead letter queue policy
	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`

	// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitempty" name:"Trace"`

	// Whether to enable transaction. 1: yes; 0: no
	Transaction *uint64 `json:"Transaction,omitempty" name:"Transaction"`

	// Queue storage space configured for message rewind. Value range: 1,024-10,240 MB (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitempty" name:"RetentionSizeInMB"`
}

func (r *ModifyCmqQueueAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqQueueAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "MaxMsgHeapNum")
	delete(f, "PollingWaitSeconds")
	delete(f, "VisibilityTimeout")
	delete(f, "MaxMsgSize")
	delete(f, "MsgRetentionSeconds")
	delete(f, "RewindSeconds")
	delete(f, "FirstQueryInterval")
	delete(f, "MaxQueryCount")
	delete(f, "DeadLetterQueueName")
	delete(f, "MaxTimeToLive")
	delete(f, "MaxReceiveCount")
	delete(f, "Policy")
	delete(f, "Trace")
	delete(f, "Transaction")
	delete(f, "RetentionSizeInMB")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCmqQueueAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqQueueAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmqQueueAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqQueueAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqSubscriptionAttributeRequest struct {
	*tchttp.BaseRequest

	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// CMQ push server retry policy in case an error occurs while pushing a message to the endpoint. Valid values:
	// (1) BACKOFF_RETRY: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message.
	// (2) EXPONENTIAL_DECAY_RETRY: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: EXPONENTIAL_DECAY_RETRY.
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`

	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `HTTP`, both values are acceptable, and the default value is `JSON`.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`

	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags"`

	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`
}

func (r *ModifyCmqSubscriptionAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqSubscriptionAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "SubscriptionName")
	delete(f, "NotifyStrategy")
	delete(f, "NotifyContentFormat")
	delete(f, "FilterTags")
	delete(f, "BindingKey")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCmqSubscriptionAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqSubscriptionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmqSubscriptionAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqSubscriptionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqTopicAttributeRequest struct {
	*tchttp.BaseRequest

	// Topic name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`

	// Message retention period. Value range: 60–86400 seconds (i.e., 1 minute–1 day). Default value: 86400.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`

	// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *ModifyCmqTopicAttributeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqTopicAttributeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "MaxMsgSize")
	delete(f, "MsgRetentionSeconds")
	delete(f, "Trace")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCmqTopicAttributeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqTopicAttributeResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmqTopicAttributeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCmqTopicAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest

	// Namespace name.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// TTL for unconsumed messages in seconds. Maximum value: 1296000 seconds.
	MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Message retention policy
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
}

func (r *ModifyEnvironmentAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentAttributesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "MsgTTL")
	delete(f, "Remark")
	delete(f, "ClusterId")
	delete(f, "RetentionPolicy")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvironmentAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Namespace name.
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// TTL for unconsumed messages in seconds.
		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

		// Remarks (up to 128 characters).
	// Note: this field may return null, indicating that no valid values can be obtained.
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// Namespace ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnvironmentAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentRoleRequest struct {
	*tchttp.BaseRequest

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Role name.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyEnvironmentRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "RoleName")
	delete(f, "Permissions")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvironmentRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnvironmentRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyEnvironmentRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// RocketMQ cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// 3–64 letters, digits, hyphens, and underscores
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyRocketMQClusterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQClusterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "ClusterName")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQClusterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQGroupRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Consumer group name
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Whether to enable consumption
	ReadEnable *bool `json:"ReadEnable,omitempty" name:"ReadEnable"`

	// Whether to enable broadcast consumption
	BroadcastEnable *bool `json:"BroadcastEnable,omitempty" name:"BroadcastEnable"`
}

func (r *ModifyRocketMQGroupRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQGroupRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupId")
	delete(f, "Remark")
	delete(f, "ReadEnable")
	delete(f, "BroadcastEnable")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQGroupResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace name, which can contain 3–64 letters, digits, hyphens, and underscores
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Retention time of unconsumed messages in milliseconds. Value range: 60 seconds–15 days
	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`

	// Retention time for persisted messages in milliseconds
	RetentionTime *uint64 `json:"RetentionTime,omitempty" name:"RetentionTime"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyRocketMQNamespaceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQNamespaceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Ttl")
	delete(f, "RetentionTime")
	delete(f, "Remark")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQNamespaceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQTopicRequest struct {
	*tchttp.BaseRequest

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Topic name
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Number of partitions, which is invalid for globally sequential messages and cannot be less than the current number of partitions.
	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
}

func (r *ModifyRocketMQTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Topic")
	delete(f, "Remark")
	delete(f, "PartitionNum")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoleRequest struct {
	*tchttp.BaseRequest

	// Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyRoleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoleRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoleName")
	delete(f, "Remark")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRoleRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Role name
		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

		// Remarks
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Number of partitions, which must be equal to or greater than the original number of partitions. To maintain the original number of partitions, enter the original number. Modifying the number of partitions will take effect only for non-globally sequential messages. There can be up to 128 partitions.
	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Partitions")
	delete(f, "Remark")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyTopicRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Number of partitions
		Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

		// Remarks (up to 128 characters).
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishCmqMsgRequest struct {
	*tchttp.BaseRequest

	// Topic name
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Message content
	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`

	// Message tag
	MsgTag []*string `json:"MsgTag,omitempty" name:"MsgTag"`
}

func (r *PublishCmqMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishCmqMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicName")
	delete(f, "MsgContent")
	delete(f, "MsgTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "PublishCmqMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type PublishCmqMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// `true` indicates that the sending is successful
		Result *bool `json:"Result,omitempty" name:"Result"`

		// Message ID
		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishCmqMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *PublishCmqMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Publisher struct {

	// Producer ID.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProducerId *int64 `json:"ProducerId,omitempty" name:"ProducerId"`

	// Producer name.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// Producer address.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitempty" name:"Address"`

	// Client version.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`

	// Message production rate (message/sec).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MsgRateIn *float64 `json:"MsgRateIn,omitempty" name:"MsgRateIn"`

	// Message production throughput rate (byte/sec).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MsgThroughputIn *float64 `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`

	// Average message size in bytes.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AverageMsgSize *float64 `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`

	// Connection time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ConnectedSince *string `json:"ConnectedSince,omitempty" name:"ConnectedSince"`

	// Serial number of the topic partition connected to the producer.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Partition *int64 `json:"Partition,omitempty" name:"Partition"`
}

type ReceiveMessageRequest struct {
	*tchttp.BaseRequest

	// Name of the topic which receives the message. It is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Subscriber name
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// Default value: 1000. Messages received by the consumer will first be stored in the `receiverQueueSize` queue to tune the message receiving rate.
	ReceiverQueueSize *int64 `json:"ReceiverQueueSize,omitempty" name:"ReceiverQueueSize"`

	// Default value: Latest. It is used to determine the position where the consumer initially receives messages. Valid values: Earliest, Latest.
	SubInitialPosition *string `json:"SubInitialPosition,omitempty" name:"SubInitialPosition"`
}

func (r *ReceiveMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReceiveMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topic")
	delete(f, "SubscriptionName")
	delete(f, "ReceiverQueueSize")
	delete(f, "SubInitialPosition")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReceiveMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ReceiveMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique primary key used to identify the message
		MessageID *string `json:"MessageID,omitempty" name:"MessageID"`

		// Content of the received message
		MessagePayload *string `json:"MessagePayload,omitempty" name:"MessagePayload"`

		// Provided to the `Ack` API and used to acknowledge messages in the topic
		AckTopic *string `json:"AckTopic,omitempty" name:"AckTopic"`

		// Returned error message. If it is an empty string, no error occurred.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// Returned subscriber name, which will be used when an acknowledgment consumer is created.
	// Note: this field may return null, indicating that no valid values can be obtained.
		SubName *string `json:"SubName,omitempty" name:"SubName"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReceiveMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ReceiveMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetMsgSubOffsetByTimestampRequest struct {
	*tchttp.BaseRequest

	// Namespace name.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Subscriber name.
	Subscription *string `json:"Subscription,omitempty" name:"Subscription"`

	// Timestamp, accurate down to the millisecond.
	ToTimestamp *uint64 `json:"ToTimestamp,omitempty" name:"ToTimestamp"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ResetMsgSubOffsetByTimestampRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetMsgSubOffsetByTimestampRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Subscription")
	delete(f, "ToTimestamp")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetMsgSubOffsetByTimestampRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetMsgSubOffsetByTimestampResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Result.
	// Note: this field may return null, indicating that no valid values can be obtained.
		Result *bool `json:"Result,omitempty" name:"Result"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetMsgSubOffsetByTimestampResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetMsgSubOffsetByTimestampResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetRocketMQConsumerOffSetRequest struct {
	*tchttp.BaseRequest

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Namespace name.
	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`

	// Consumer group name.
	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`

	// Topic name.
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Reset method. 0: Start from the latest offset; 1: Start from specified time point.
	Type *uint64 `json:"Type,omitempty" name:"Type"`

	// The specified timestamp that has been reset, in milliseconds. This parameter only takes effect when the value of `Type` is `1`.
	ResetTimestamp *uint64 `json:"ResetTimestamp,omitempty" name:"ResetTimestamp"`
}

func (r *ResetRocketMQConsumerOffSetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRocketMQConsumerOffSetRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "GroupId")
	delete(f, "Topic")
	delete(f, "Type")
	delete(f, "ResetTimestamp")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ResetRocketMQConsumerOffSetRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ResetRocketMQConsumerOffSetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetRocketMQConsumerOffSetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ResetRocketMQConsumerOffSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetentionPolicy struct {

	// Message retention period
	TimeInMinutes *int64 `json:"TimeInMinutes,omitempty" name:"TimeInMinutes"`

	// Message retention size
	SizeInMB *int64 `json:"SizeInMB,omitempty" name:"SizeInMB"`
}

type RewindCmqQueueRequest struct {
	*tchttp.BaseRequest

	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// After this time is configured, the `(Batch)receiveMessage` API will consume the messages received after this timestamp in the order in which they are produced.
	StartConsumeTime *uint64 `json:"StartConsumeTime,omitempty" name:"StartConsumeTime"`
}

func (r *RewindCmqQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RewindCmqQueueRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "StartConsumeTime")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "RewindCmqQueueRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type RewindCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RewindCmqQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *RewindCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQClusterConfig struct {

	// Maximum TPS per namespace
	MaxTpsPerNamespace *uint64 `json:"MaxTpsPerNamespace,omitempty" name:"MaxTpsPerNamespace"`

	// Maximum number of namespaces
	MaxNamespaceNum *uint64 `json:"MaxNamespaceNum,omitempty" name:"MaxNamespaceNum"`

	// Number of used namespaces
	UsedNamespaceNum *uint64 `json:"UsedNamespaceNum,omitempty" name:"UsedNamespaceNum"`

	// Maximum number of topics
	MaxTopicNum *uint64 `json:"MaxTopicNum,omitempty" name:"MaxTopicNum"`

	// Number of used topics
	UsedTopicNum *uint64 `json:"UsedTopicNum,omitempty" name:"UsedTopicNum"`

	// Maximum number of groups
	MaxGroupNum *uint64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`

	// Number of used groups
	UsedGroupNum *uint64 `json:"UsedGroupNum,omitempty" name:"UsedGroupNum"`

	// Maximum message retention period in milliseconds
	MaxRetentionTime *uint64 `json:"MaxRetentionTime,omitempty" name:"MaxRetentionTime"`

	// Maximum message delay in milliseconds
	MaxLatencyTime *uint64 `json:"MaxLatencyTime,omitempty" name:"MaxLatencyTime"`
}

type RocketMQClusterInfo struct {

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`

	// Region information
	Region *string `json:"Region,omitempty" name:"Region"`

	// Creation time in milliseconds
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// Cluster remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Public network access address
	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`

	// VPC access address
	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`

	// Whether the namespace access point is supported.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SupportNamespaceEndpoint *bool `json:"SupportNamespaceEndpoint,omitempty" name:"SupportNamespaceEndpoint"`
}

type RocketMQClusterRecentStats struct {

	// Number of topics
	TopicNum *uint64 `json:"TopicNum,omitempty" name:"TopicNum"`

	// Number of produced messages
	ProducedMsgNum *uint64 `json:"ProducedMsgNum,omitempty" name:"ProducedMsgNum"`

	// Number of consumed messages
	ConsumedMsgNum *uint64 `json:"ConsumedMsgNum,omitempty" name:"ConsumedMsgNum"`

	// Number of retained messages
	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitempty" name:"AccumulativeMsgNum"`
}

type Role struct {

	// Role name.
	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`

	// Value of the role token.
	Token *string `json:"Token,omitempty" name:"Token"`

	// Remarks.
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type SendBatchMessagesRequest struct {
	*tchttp.BaseRequest

	// Name of the topic to which to send the message. It is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Content of the message to be sent
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// String-Type token, which is optional and will be automatically obtained by the system.
	StringToken *string `json:"StringToken,omitempty" name:"StringToken"`

	// Producer name, which must be globally unique. If it is not configured, the system will automatically generate one.
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// Message sending timeout period in seconds. Default value: 30s
	SendTimeout *int64 `json:"SendTimeout,omitempty" name:"SendTimeout"`

	// Maximum number of produced messages which can be cached in the memory. Default value: 1000
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitempty" name:"MaxPendingMessages"`

	// Maximum number of messages in each batch. Default value: 1000 messages/batch
	BatchingMaxMessages *int64 `json:"BatchingMaxMessages,omitempty" name:"BatchingMaxMessages"`

	// Maximum wait time for each batch, after which the batch will be sent no matter whether the specified number or size of messages in the batch is reached. Default value: 10 ms
	BatchingMaxPublishDelay *int64 `json:"BatchingMaxPublishDelay,omitempty" name:"BatchingMaxPublishDelay"`

	// Maximum allowed size of messages in each batch. Default value: 128 KB
	BatchingMaxBytes *int64 `json:"BatchingMaxBytes,omitempty" name:"BatchingMaxBytes"`
}

func (r *SendBatchMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendBatchMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topic")
	delete(f, "Payload")
	delete(f, "StringToken")
	delete(f, "ProducerName")
	delete(f, "SendTimeout")
	delete(f, "MaxPendingMessages")
	delete(f, "BatchingMaxMessages")
	delete(f, "BatchingMaxPublishDelay")
	delete(f, "BatchingMaxBytes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendBatchMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendBatchMessagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// Unique message ID
	// Note: this field may return null, indicating that no valid values can be obtained.
		MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

		// Error message. If an empty string is returned, no error occurred.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendBatchMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendBatchMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendCmqMsgRequest struct {
	*tchttp.BaseRequest

	// Queue name
	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`

	// Message content
	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`

	// Delay time
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" name:"DelaySeconds"`
}

func (r *SendCmqMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendCmqMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "QueueName")
	delete(f, "MsgContent")
	delete(f, "DelaySeconds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendCmqMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendCmqMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// `true` indicates that the sending is successful
		Result *bool `json:"Result,omitempty" name:"Result"`

		// Message ID
		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendCmqMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendCmqMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendMessagesRequest struct {
	*tchttp.BaseRequest

	// Name of the topic to which to send the message. It is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	Topic *string `json:"Topic,omitempty" name:"Topic"`

	// Content of the message to be sent
	Payload *string `json:"Payload,omitempty" name:"Payload"`

	// Token used for authentication, which is optional and will be automatically obtained by the system.
	StringToken *string `json:"StringToken,omitempty" name:"StringToken"`

	// Producer name, which must be globally unique. If it is not configured, the system will randomly generate one.
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// Message sending timeout period, which is 30s by default.
	SendTimeout *int64 `json:"SendTimeout,omitempty" name:"SendTimeout"`

	// Maximum number of produced messages which can be cached in the memory. Default value: 1000
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitempty" name:"MaxPendingMessages"`
}

func (r *SendMessagesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMessagesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Topic")
	delete(f, "Payload")
	delete(f, "StringToken")
	delete(f, "ProducerName")
	delete(f, "SendTimeout")
	delete(f, "MaxPendingMessages")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendMessagesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendMessagesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// messageID, which must be globally unique and is the metadata information used to identify the message.
	// Note: this field may return null, indicating that no valid values can be obtained.
		MessageId *string `json:"MessageId,omitempty" name:"MessageId"`

		// Returned error message. If an empty string is returned, no error occurred.
	// Note: this field may return null, indicating that no valid values can be obtained.
		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendMessagesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendMsgRequest struct {
	*tchttp.BaseRequest

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// Topic name. If the topic is a partitioned topic, you need to specify the partition; otherwise, messages will be sent to partition 0 by default, such as `my_topic-partition-0`.
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// Message content, which cannot be empty and can contain up to 5,242,880 bytes.
	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *SendMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "MsgContent")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type SendMsgResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Sort struct {

	// Sorting field.
	Name *string `json:"Name,omitempty" name:"Name"`

	// Ascending order: `ASC`; descending order: `DESC`.
	Order *string `json:"Order,omitempty" name:"Order"`
}

type Tag struct {

	// Value of the tag key
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// Value of the tag value
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type UnbindCmqDeadLetterRequest struct {
	*tchttp.BaseRequest

	// Source queue name of dead letter policy. Calling this API will clear the dead letter queue policy of this queue.
	SourceQueueName *string `json:"SourceQueueName,omitempty" name:"SourceQueueName"`
}

func (r *UnbindCmqDeadLetterRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCmqDeadLetterRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SourceQueueName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UnbindCmqDeadLetterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type UnbindCmqDeadLetterResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindCmqDeadLetterResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UnbindCmqDeadLetterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcBindRecord struct {

	// Tenant VPC ID
	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`

	// Tenant VPC subnet ID
	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`

	// Route ID
	RouterId *string `json:"RouterId,omitempty" name:"RouterId"`

	// VPC ID
	Ip *string `json:"Ip,omitempty" name:"Ip"`

	// VPC port
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// Remarks (up to 128 characters)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}
