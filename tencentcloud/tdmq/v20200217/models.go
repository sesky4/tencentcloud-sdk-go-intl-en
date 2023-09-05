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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go-intl-en/tencentcloud/common/json"
)

// Predefined struct for user
type AcknowledgeMessageRequestParams struct {
	// Unique ID used to identify the message, which can be obtained from the returned value of `receiveMessage`.
	MessageId *string `json:"MessageId,omitnil" name:"MessageId"`

	// Topic name, which can be obtained from the returned value of `receiveMessage` and is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	AckTopic *string `json:"AckTopic,omitnil" name:"AckTopic"`

	// Subscriber name, which can be obtained from the returned value of `receiveMessage`. Make sure that it is the same as the subscriber name identified in `receiveMessage`; otherwise, the received message cannot be correctly acknowledged.
	SubName *string `json:"SubName,omitnil" name:"SubName"`
}

type AcknowledgeMessageRequest struct {
	*tchttp.BaseRequest
	
	// Unique ID used to identify the message, which can be obtained from the returned value of `receiveMessage`.
	MessageId *string `json:"MessageId,omitnil" name:"MessageId"`

	// Topic name, which can be obtained from the returned value of `receiveMessage` and is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	AckTopic *string `json:"AckTopic,omitnil" name:"AckTopic"`

	// Subscriber name, which can be obtained from the returned value of `receiveMessage`. Make sure that it is the same as the subscriber name identified in `receiveMessage`; otherwise, the received message cannot be correctly acknowledged.
	SubName *string `json:"SubName,omitnil" name:"SubName"`
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

// Predefined struct for user
type AcknowledgeMessageResponseParams struct {
	// If it is an empty string, no error occurred.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type AcknowledgeMessageResponse struct {
	*tchttp.BaseResponse
	Response *AcknowledgeMessageResponseParams `json:"Response"`
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

type BindCluster struct {
	// Name of a physical cluster.
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`
}

// Predefined struct for user
type ClearCmqQueueRequestParams struct {
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
}

type ClearCmqQueueRequest struct {
	*tchttp.BaseRequest
	
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
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

// Predefined struct for user
type ClearCmqQueueResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ClearCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *ClearCmqQueueResponseParams `json:"Response"`
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

// Predefined struct for user
type ClearCmqSubscriptionFilterTagsRequestParams struct {
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`
}

type ClearCmqSubscriptionFilterTagsRequest struct {
	*tchttp.BaseRequest
	
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`
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

// Predefined struct for user
type ClearCmqSubscriptionFilterTagsResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ClearCmqSubscriptionFilterTagsResponse struct {
	*tchttp.BaseResponse
	Response *ClearCmqSubscriptionFilterTagsResponseParams `json:"Response"`
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
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Cluster name.
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Remarks.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Number of access points
	EndPointNum *int64 `json:"EndPointNum,omitnil" name:"EndPointNum"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Whether the cluster is healthy. 1: healthy; 0: exceptional
	Healthy *int64 `json:"Healthy,omitnil" name:"Healthy"`

	// Cluster health information
	// Note: this field may return null, indicating that no valid values can be obtained.
	HealthyInfo *string `json:"HealthyInfo,omitnil" name:"HealthyInfo"`

	// Cluster status. 0: creating; 1: normal; 2: terminating; 3: deleted; 4. isolated; 5. creation failed; 6: deletion failed
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Maximum number of namespaces
	MaxNamespaceNum *int64 `json:"MaxNamespaceNum,omitnil" name:"MaxNamespaceNum"`

	// Maximum number of topics
	MaxTopicNum *int64 `json:"MaxTopicNum,omitnil" name:"MaxTopicNum"`

	// Maximum QPS
	MaxQps *int64 `json:"MaxQps,omitnil" name:"MaxQps"`

	// Maximum message retention period in seconds
	MessageRetentionTime *int64 `json:"MessageRetentionTime,omitnil" name:"MessageRetentionTime"`

	// Maximum storage capacity
	MaxStorageCapacity *int64 `json:"MaxStorageCapacity,omitnil" name:"MaxStorageCapacity"`

	// Cluster version
	// Note: this field may return null, indicating that no valid values can be obtained.
	Version *string `json:"Version,omitnil" name:"Version"`

	// Public network access point
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicEndPoint *string `json:"PublicEndPoint,omitnil" name:"PublicEndPoint"`

	// VPC access point
	// Note: this field may return null, indicating that no valid values can be obtained.
	VpcEndPoint *string `json:"VpcEndPoint,omitnil" name:"VpcEndPoint"`

	// Number of namespaces
	// Note: this field may return null, indicating that no valid values can be obtained.
	NamespaceNum *int64 `json:"NamespaceNum,omitnil" name:"NamespaceNum"`

	// Limit of used storage in MB
	// Note: this field may return null, indicating that no valid values can be obtained.
	UsedStorageBudget *int64 `json:"UsedStorageBudget,omitnil" name:"UsedStorageBudget"`

	// Maximum message production rate in messages
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxPublishRateInMessages *int64 `json:"MaxPublishRateInMessages,omitnil" name:"MaxPublishRateInMessages"`

	// Maximum message push rate in messages
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxDispatchRateInMessages *int64 `json:"MaxDispatchRateInMessages,omitnil" name:"MaxDispatchRateInMessages"`

	// Maximum message production rate in bytes
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxPublishRateInBytes *int64 `json:"MaxPublishRateInBytes,omitnil" name:"MaxPublishRateInBytes"`

	// Maximum message push rate in bytes
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxDispatchRateInBytes *int64 `json:"MaxDispatchRateInBytes,omitnil" name:"MaxDispatchRateInBytes"`

	// Number of created topics
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicNum *int64 `json:"TopicNum,omitnil" name:"TopicNum"`

	// Maximum message delay in seconds
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMessageDelayInSeconds *int64 `json:"MaxMessageDelayInSeconds,omitnil" name:"MaxMessageDelayInSeconds"`

	// Whether to enable public network access. If this parameter is left empty, the feature will be enabled by default
	// Note: this field may return null, indicating that no valid values can be obtained.
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil" name:"PublicAccessEnabled"`

	// Tag
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Billing mode:
	// `0`: Pay-as-you-go
	// `1`: Monthly subscription
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	PayMode *int64 `json:"PayMode,omitnil" name:"PayMode"`
}

type CmqDeadLetterPolicy struct {
	// Dead letter queue.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterQueue *string `json:"DeadLetterQueue,omitnil" name:"DeadLetterQueue"`

	// Dead letter queue policy.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Policy *uint64 `json:"Policy,omitnil" name:"Policy"`

	// Maximum period in seconds before an unconsumed message expires, which is required if `Policy` is 1. Value range: 300–43200. This value should be smaller than `MsgRetentionSeconds` (maximum message retention period)
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil" name:"MaxTimeToLive"`

	// Maximum number of receipts.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil" name:"MaxReceiveCount"`
}

type CmqDeadLetterSource struct {
	// Message queue ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	QueueId *string `json:"QueueId,omitnil" name:"QueueId"`

	// Message queue name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
}

type CmqQueue struct {
	// Message queue ID.
	QueueId *string `json:"QueueId,omitnil" name:"QueueId"`

	// Message queue name.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// Limit of the number of messages produced per second. The value for consumed messages is 1.1 times this value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Qps *uint64 `json:"Qps,omitnil" name:"Qps"`

	// Bandwidth limit.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Bps *uint64 `json:"Bps,omitnil" name:"Bps"`

	// Maximum retention period for inflight messages.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxDelaySeconds *uint64 `json:"MaxDelaySeconds,omitnil" name:"MaxDelaySeconds"`

	// Maximum number of heaped messages. The value range is 1,000,000–10,000,000 during the beta test and can be 1,000,000–1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil" name:"MaxMsgHeapNum"`

	// Long polling wait time for message reception. Value range: 0–30 seconds. Default value: 0.
	// Note: this field may return null, indicating that no valid values can be obtained.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil" name:"PollingWaitSeconds"`

	// Message retention period. Value range: 60–1296000 seconds (i.e., 1 minute–15 days). Default value: 345600 (i.e., 4 days).
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil" name:"MsgRetentionSeconds"`

	// Message visibility timeout period. Value range: 1–43200 seconds (i.e., 12 hours). Default value: 30.
	// Note: this field may return null, indicating that no valid values can be obtained.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil" name:"VisibilityTimeout"`

	// Maximum message length. Value range: 1024–1048576 bytes (i.e., 1–1024 KB). Default value: 65536.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil" name:"MaxMsgSize"`

	// Maximum time range during which a message can be rewound in the queue, which ranges from 0 to 43,200 seconds. 0 indicates that message rewind is disabled.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil" name:"RewindSeconds"`

	// Queue creation time. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// Time when the queue attribute is last modified. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitnil" name:"LastModifyTime"`

	// Total number of messages in `Active` status (i.e., unconsumed) in the queue, which is an approximate value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ActiveMsgNum *uint64 `json:"ActiveMsgNum,omitnil" name:"ActiveMsgNum"`

	// Total number of messages in `Inactive` status (i.e., being consumed) in the queue, which is an approximate value.
	// Note: this field may return null, indicating that no valid values can be obtained.
	InactiveMsgNum *uint64 `json:"InactiveMsgNum,omitnil" name:"InactiveMsgNum"`

	// Number of delayed messages.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DelayMsgNum *uint64 `json:"DelayMsgNum,omitnil" name:"DelayMsgNum"`

	// Number of retained messages which have been deleted by the `DelMsg` API but are still within their rewind time range.
	// Note: this field may return null, indicating that no valid values can be obtained.
	RewindMsgNum *uint64 `json:"RewindMsgNum,omitnil" name:"RewindMsgNum"`

	// Minimum unconsumed time of message in seconds.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MinMsgTime *uint64 `json:"MinMsgTime,omitnil" name:"MinMsgTime"`

	// Transaction message queue. true: transaction message type; false: other message types.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Transaction *bool `json:"Transaction,omitnil" name:"Transaction"`

	// Dead letter queue.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterSource []*CmqDeadLetterSource `json:"DeadLetterSource,omitnil" name:"DeadLetterSource"`

	// Dead letter queue policy.
	// Note: this field may return null, indicating that no valid values can be obtained.
	DeadLetterPolicy *CmqDeadLetterPolicy `json:"DeadLetterPolicy,omitnil" name:"DeadLetterPolicy"`

	// Transaction message policy.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TransactionPolicy *CmqTransactionPolicy `json:"TransactionPolicy,omitnil" name:"TransactionPolicy"`

	// Creator `Uin`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateUin *uint64 `json:"CreateUin,omitnil" name:"CreateUin"`

	// Associated tag.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Message trace. true: enabled; false: not enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Trace *bool `json:"Trace,omitnil" name:"Trace"`

	// Tenant ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// Namespace name
	// Note: this field may return null, indicating that no valid values can be obtained.
	NamespaceName *string `json:"NamespaceName,omitnil" name:"NamespaceName"`

	// Cluster status. 0: creating; 1: normal; 2: terminating; 3: deleted; 4. isolated; 5. creation failed; 6: deletion failed
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// The maximum number of unacknowledged messages.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MaxUnackedMsgNum *int64 `json:"MaxUnackedMsgNum,omitnil" name:"MaxUnackedMsgNum"`

	// Maximum size of heaped messages in bytes.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MaxMsgBacklogSize *int64 `json:"MaxMsgBacklogSize,omitnil" name:"MaxMsgBacklogSize"`

	// Queue storage space configured for message rewind. Value range: 1,024-10,240 MB (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitnil" name:"RetentionSizeInMB"`
}

type CmqSubscription struct {
	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// Subscription ID, which will be used during monitoring data pull.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubscriptionId *string `json:"SubscriptionId,omitnil" name:"SubscriptionId"`

	// Subscription owner `APPID`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicOwner *uint64 `json:"TopicOwner,omitnil" name:"TopicOwner"`

	// Number of messages to be delivered in the subscription.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgCount *uint64 `json:"MsgCount,omitnil" name:"MsgCount"`

	// Time when the subscription attribute is last modified. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitnil" name:"LastModifyTime"`

	// Subscription creation time. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// Filtering policy for subscribing to and receiving messages.
	// Note: this field may return null, indicating that no valid values can be obtained.
	BindingKey []*string `json:"BindingKey,omitnil" name:"BindingKey"`

	// Endpoint that receives notifications, which varies by `protocol`: for HTTP, the endpoint must start with `http://`, and the `host` can be a domain or IP; for `queue`, `queueName` should be entered.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// Filtering policy selected when a subscription is created:
	// If `filterType` is 1, `filterTag` will be used for filtering.
	// If `filterType` is 2, `bindingKey` will be used for filtering.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FilterTags []*string `json:"FilterTags,omitnil" name:"FilterTags"`

	// Subscription protocol. Currently, two protocols are supported: HTTP and queue. To use the HTTP protocol, you need to build your own web server to receive messages. With the queue protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// CMQ push server retry policy in case an error occurs while pushing a message to the endpoint. Valid values:
	// (1) BACKOFF_RETRY: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message.
	// (2) EXPONENTIAL_DECAY_RETRY: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: EXPONENTIAL_DECAY_RETRY.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NotifyStrategy *string `json:"NotifyStrategy,omitnil" name:"NotifyStrategy"`

	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `protocol` is `queue`, this value must be `SIMPLIFIED`. If `protocol` is `HTTP`, both values are acceptable, and the default value is `JSON`.
	// Note: this field may return null, indicating that no valid values can be obtained.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitnil" name:"NotifyContentFormat"`
}

type CmqTopic struct {
	// Topic ID.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// Topic name.
	// Note: this field may return null, indicating that no valid values can be obtained.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Maximum lifecycle of message in topic. After the period specified by this parameter has elapsed since a message is sent to the topic, the message will be deleted no matter whether it has been successfully pushed to the user. This parameter is measured in seconds and defaulted to one day (86,400 seconds), which cannot be modified.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil" name:"MsgRetentionSeconds"`

	// Maximum message size, which ranges from 1,024 to 1,048,576 bytes (i.e., 1–1,024 KB). The default value is 65,536.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil" name:"MaxMsgSize"`

	// Number of messages published per second.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Qps *uint64 `json:"Qps,omitnil" name:"Qps"`

	// Filtering policy selected when a subscription is created:
	// If `filterType` is 1, `FilterTag` will be used for filtering.
	// If `filterType` is 2, `BindingKey` will be used for filtering.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FilterType *uint64 `json:"FilterType,omitnil" name:"FilterType"`

	// Topic creation time. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// Time when the topic attribute is last modified. A Unix timestamp accurate down to the millisecond will be returned.
	// Note: this field may return null, indicating that no valid values can be obtained.
	LastModifyTime *uint64 `json:"LastModifyTime,omitnil" name:"LastModifyTime"`

	// Number of current messages in the topic (number of retained messages).
	// Note: this field may return null, indicating that no valid values can be obtained.
	MsgCount *uint64 `json:"MsgCount,omitnil" name:"MsgCount"`

	// Creator `Uin`. The `resource` field for CAM authentication is composed of this field.
	// Note: this field may return null, indicating that no valid values can be obtained.
	CreateUin *uint64 `json:"CreateUin,omitnil" name:"CreateUin"`

	// Associated tag.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Message trace. true: enabled; false: not enabled
	// Note: this field may return null, indicating that no valid values can be obtained.
	Trace *bool `json:"Trace,omitnil" name:"Trace"`

	// Tenant ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	TenantId *string `json:"TenantId,omitnil" name:"TenantId"`

	// Namespace name
	// Note: this field may return null, indicating that no valid values can be obtained.
	NamespaceName *string `json:"NamespaceName,omitnil" name:"NamespaceName"`

	// Cluster status. 0: creating; 1: normal; 2: terminating; 3: deleted; 4. isolated; 5. creation failed; 6: deletion failed
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Valid values: `0` (Pulsar), `1` (RocketMQ).
	// Note: This field may return null, indicating that no valid values can be obtained.
	BrokerType *int64 `json:"BrokerType,omitnil" name:"BrokerType"`
}

type CmqTransactionPolicy struct {
	// First lookback time.
	// Note: this field may return null, indicating that no valid values can be obtained.
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil" name:"FirstQueryInterval"`

	// Maximum number of queries.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil" name:"MaxQueryCount"`
}

type Consumer struct {
	// The time when the consumer started connecting.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConnectedSince *string `json:"ConnectedSince,omitnil" name:"ConnectedSince"`

	// Consumer address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerAddr *string `json:"ConsumerAddr,omitnil" name:"ConsumerAddr"`

	// Consumer name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerName *string `json:"ConsumerName,omitnil" name:"ConsumerName"`

	// Consumer version.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClientVersion *string `json:"ClientVersion,omitnil" name:"ClientVersion"`

	// Serial number of the topic partition connected to the consumer.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Partition *int64 `json:"Partition,omitnil" name:"Partition"`
}

type ConsumersSchedule struct {
	// ID of the current partition.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Partitions *uint64 `json:"Partitions,omitnil" name:"Partitions"`

	// The number of messages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NumberOfEntries *uint64 `json:"NumberOfEntries,omitnil" name:"NumberOfEntries"`

	// The number of heaped messages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgBacklog *uint64 `json:"MsgBacklog,omitnil" name:"MsgBacklog"`

	// The total number of messages delivered by the consumer per second.
	MsgRateOut *string `json:"MsgRateOut,omitnil" name:"MsgRateOut"`

	// The size (in bytes) of messages consumed by the consumer per second.
	MsgThroughputOut *string `json:"MsgThroughputOut,omitnil" name:"MsgThroughputOut"`

	// Percentage of messages discarded due to timeout.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgRateExpired *string `json:"MsgRateExpired,omitnil" name:"MsgRateExpired"`
}

// Predefined struct for user
type CreateClusterRequestParams struct {
	// Cluster name, which can contain up to 16 letters, digits, hyphens, and underscores.
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// ID of your dedicated physical cluster. If it is not passed in, cluster resources will be created in a public cluster by default.
	BindClusterId *uint64 `json:"BindClusterId,omitnil" name:"BindClusterId"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Cluster tag list (deprecated).
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Whether to enable public network access. If this parameter is left empty, the feature will be enabled by default
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil" name:"PublicAccessEnabled"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster name, which can contain up to 16 letters, digits, hyphens, and underscores.
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// ID of your dedicated physical cluster. If it is not passed in, cluster resources will be created in a public cluster by default.
	BindClusterId *uint64 `json:"BindClusterId,omitnil" name:"BindClusterId"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Cluster tag list (deprecated).
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Whether to enable public network access. If this parameter is left empty, the feature will be enabled by default
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil" name:"PublicAccessEnabled"`
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

// Predefined struct for user
type CreateClusterResponseParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateCmqQueueRequestParams struct {
	// Queue name, which is unique under the same account in a single region. It is a string of up to 64 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// Maximum number of heaped messages. The value range is 1,000,000–10,000,000 during the beta test and can be 1,000,000–1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil" name:"MaxMsgHeapNum"`

	// Long polling wait time for message reception. Value range: 0–30 seconds. Default value: 0.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil" name:"PollingWaitSeconds"`

	// Message visibility timeout period. Value range: 1–43200 seconds (i.e., 12 hours). Default value: 30.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil" name:"VisibilityTimeout"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil" name:"MaxMsgSize"`

	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil" name:"MsgRetentionSeconds"`

	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil" name:"RewindSeconds"`

	// 1: transaction queue; 0: general queue
	Transaction *uint64 `json:"Transaction,omitnil" name:"Transaction"`

	// First lookback interval
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil" name:"FirstQueryInterval"`

	// Maximum number of lookbacks
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil" name:"MaxQueryCount"`

	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil" name:"DeadLetterQueueName"`

	// Dead letter policy. 0: message has been consumed multiple times but not deleted; 1: `Time-To-Live` has elapsed
	Policy *uint64 `json:"Policy,omitnil" name:"Policy"`

	// Maximum receipt times. Value range: 1–1000
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil" name:"MaxReceiveCount"`

	// Maximum period in seconds before an unconsumed message expires, which is required if `policy` is 1. Value range: 300–43200. This value should be smaller than `msgRetentionSeconds` (maximum message retention period)
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil" name:"MaxTimeToLive"`

	// Whether to enable message trace. true: yes; false: no. If this field is not configured, the feature will not be enabled
	Trace *bool `json:"Trace,omitnil" name:"Trace"`

	// Tag array.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Queue storage space configured for message rewind. Value range: 10,240-512,000 MB (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitnil" name:"RetentionSizeInMB"`
}

type CreateCmqQueueRequest struct {
	*tchttp.BaseRequest
	
	// Queue name, which is unique under the same account in a single region. It is a string of up to 64 characters. It can contain letters, digits, and hyphens (-) and must start with a letter.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// Maximum number of heaped messages. The value range is 1,000,000–10,000,000 during the beta test and can be 1,000,000–1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil" name:"MaxMsgHeapNum"`

	// Long polling wait time for message reception. Value range: 0–30 seconds. Default value: 0.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil" name:"PollingWaitSeconds"`

	// Message visibility timeout period. Value range: 1–43200 seconds (i.e., 12 hours). Default value: 30.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil" name:"VisibilityTimeout"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil" name:"MaxMsgSize"`

	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil" name:"MsgRetentionSeconds"`

	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil" name:"RewindSeconds"`

	// 1: transaction queue; 0: general queue
	Transaction *uint64 `json:"Transaction,omitnil" name:"Transaction"`

	// First lookback interval
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil" name:"FirstQueryInterval"`

	// Maximum number of lookbacks
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil" name:"MaxQueryCount"`

	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil" name:"DeadLetterQueueName"`

	// Dead letter policy. 0: message has been consumed multiple times but not deleted; 1: `Time-To-Live` has elapsed
	Policy *uint64 `json:"Policy,omitnil" name:"Policy"`

	// Maximum receipt times. Value range: 1–1000
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil" name:"MaxReceiveCount"`

	// Maximum period in seconds before an unconsumed message expires, which is required if `policy` is 1. Value range: 300–43200. This value should be smaller than `msgRetentionSeconds` (maximum message retention period)
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil" name:"MaxTimeToLive"`

	// Whether to enable message trace. true: yes; false: no. If this field is not configured, the feature will not be enabled
	Trace *bool `json:"Trace,omitnil" name:"Trace"`

	// Tag array.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`

	// Queue storage space configured for message rewind. Value range: 10,240-512,000 MB (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitnil" name:"RetentionSizeInMB"`
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

// Predefined struct for user
type CreateCmqQueueResponseParams struct {
	// `queueId` of a successfully created queue
	QueueId *string `json:"QueueId,omitnil" name:"QueueId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *CreateCmqQueueResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateCmqSubscribeRequestParams struct {
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// Subscription protocol. Currently, two protocols are supported: HTTP and queue. To use the HTTP protocol, you need to build your own web server to receive messages. With the queue protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages will not be received. Currently, messages can be pushed only to the public network and classic network.
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1. BACKOFF_RETRY: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message; 2. EXPONENTIAL_DECAY_RETRY: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: EXPONENTIAL_DECAY_RETRY.
	NotifyStrategy *string `json:"NotifyStrategy,omitnil" name:"NotifyStrategy"`

	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTag []*string `json:"FilterTag,omitnil" name:"FilterTag"`

	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKey []*string `json:"BindingKey,omitnil" name:"BindingKey"`

	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitnil" name:"NotifyContentFormat"`
}

type CreateCmqSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// Subscription protocol. Currently, two protocols are supported: HTTP and queue. To use the HTTP protocol, you need to build your own web server to receive messages. With the queue protocol, messages are automatically pushed to a CMQ queue and you can pull them concurrently.
	Protocol *string `json:"Protocol,omitnil" name:"Protocol"`

	// `Endpoint` for notification receipt, which is distinguished by `Protocol`. For `http`, `Endpoint` must begin with `http://` and `host` can be a domain name or IP. For `Queue`, enter `QueueName`. Note that currently the push service cannot push messages to a VPC; therefore, if a VPC domain name or address is entered for `Endpoint`, pushed messages will not be received. Currently, messages can be pushed only to the public network and classic network.
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// CMQ push server retry policy in case an error occurs while pushing a message to `Endpoint`. Valid values: 1. BACKOFF_RETRY: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message; 2. EXPONENTIAL_DECAY_RETRY: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: EXPONENTIAL_DECAY_RETRY.
	NotifyStrategy *string `json:"NotifyStrategy,omitnil" name:"NotifyStrategy"`

	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTag []*string `json:"FilterTag,omitnil" name:"FilterTag"`

	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKey []*string `json:"BindingKey,omitnil" name:"BindingKey"`

	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `http`, both options are acceptable, and the default value is `JSON`.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitnil" name:"NotifyContentFormat"`
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

// Predefined struct for user
type CreateCmqSubscribeResponseParams struct {
	// Subscription ID
	SubscriptionId *string `json:"SubscriptionId,omitnil" name:"SubscriptionId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCmqSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *CreateCmqSubscribeResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateCmqTopicRequestParams struct {
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil" name:"MaxMsgSize"`

	// Used to specify the message match policy for the topic. 1: tag match policy (default value); 2: routing match policy.
	FilterType *uint64 `json:"FilterType,omitnil" name:"FilterType"`

	// Message retention period. Value range: 60–86400 seconds (i.e., 1 minute–1 day). Default value: 86400.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil" name:"MsgRetentionSeconds"`

	// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitnil" name:"Trace"`

	// Tag array.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
}

type CreateCmqTopicRequest struct {
	*tchttp.BaseRequest
	
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil" name:"MaxMsgSize"`

	// Used to specify the message match policy for the topic. 1: tag match policy (default value); 2: routing match policy.
	FilterType *uint64 `json:"FilterType,omitnil" name:"FilterType"`

	// Message retention period. Value range: 60–86400 seconds (i.e., 1 minute–1 day). Default value: 86400.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil" name:"MsgRetentionSeconds"`

	// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitnil" name:"Trace"`

	// Tag array.
	Tags []*Tag `json:"Tags,omitnil" name:"Tags"`
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

// Predefined struct for user
type CreateCmqTopicResponseParams struct {
	// Topic ID
	TopicId *string `json:"TopicId,omitnil" name:"TopicId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateCmqTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateCmqTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateEnvironmentRequestParams struct {
	// Environment (namespace) name, which can contain up to 16 letters, digits, hyphens, and underscores.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Retention period for unconsumed messages in seconds. Value range: 60s to 1,296,000s (or 15 days).
	MsgTTL *uint64 `json:"MsgTTL,omitnil" name:"MsgTTL"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Message retention policy
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitnil" name:"RetentionPolicy"`

	// Whether to enable "Auto-Create Subscription"
	AutoSubscriptionCreation *bool `json:"AutoSubscriptionCreation,omitnil" name:"AutoSubscriptionCreation"`
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name, which can contain up to 16 letters, digits, hyphens, and underscores.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Retention period for unconsumed messages in seconds. Value range: 60s to 1,296,000s (or 15 days).
	MsgTTL *uint64 `json:"MsgTTL,omitnil" name:"MsgTTL"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Message retention policy
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitnil" name:"RetentionPolicy"`

	// Whether to enable "Auto-Create Subscription"
	AutoSubscriptionCreation *bool `json:"AutoSubscriptionCreation,omitnil" name:"AutoSubscriptionCreation"`
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
	delete(f, "AutoSubscriptionCreation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateEnvironmentRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateEnvironmentResponseParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// TTL for unconsumed messages in seconds.
	MsgTTL *uint64 `json:"MsgTTL,omitnil" name:"MsgTTL"`

	// Remarks (up to 128 characters).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Namespace ID
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvironmentResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateEnvironmentRoleRequestParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Role name.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions []*string `json:"Permissions,omitnil" name:"Permissions"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type CreateEnvironmentRoleRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Role name.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions []*string `json:"Permissions,omitnil" name:"Permissions"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type CreateEnvironmentRoleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateEnvironmentRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateEnvironmentRoleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRabbitMQUserRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username, which is used for login.
	User *string `json:"User,omitnil" name:"User"`

	// Password, which is used for login.
	Password *string `json:"Password,omitnil" name:"Password"`

	// Description
	Description *string `json:"Description,omitnil" name:"Description"`

	// User tag, which defines a user's permission scope for accessing RabbitMQ Managementu200d.
	// Valid values: `management` (Common console user), monitoring` (Console admin user), other values: Non-console user.
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// The maximum number of connections for the user. If this parameter is left empty, there's no limit for the number.
	MaxConnections *int64 `json:"MaxConnections,omitnil" name:"MaxConnections"`

	// The maximum number of channels for the user. If this parameter is left empty, there's no limit for the number.
	MaxChannels *int64 `json:"MaxChannels,omitnil" name:"MaxChannels"`
}

type CreateRabbitMQUserRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username, which is used for login.
	User *string `json:"User,omitnil" name:"User"`

	// Password, which is used for login.
	Password *string `json:"Password,omitnil" name:"Password"`

	// Description
	Description *string `json:"Description,omitnil" name:"Description"`

	// User tag, which defines a user's permission scope for accessing RabbitMQ Managementu200d.
	// Valid values: `management` (Common console user), monitoring` (Console admin user), other values: Non-console user.
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// The maximum number of connections for the user. If this parameter is left empty, there's no limit for the number.
	MaxConnections *int64 `json:"MaxConnections,omitnil" name:"MaxConnections"`

	// The maximum number of channels for the user. If this parameter is left empty, there's no limit for the number.
	MaxChannels *int64 `json:"MaxChannels,omitnil" name:"MaxChannels"`
}

func (r *CreateRabbitMQUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "Tags")
	delete(f, "MaxConnections")
	delete(f, "MaxChannels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQUserResponseParams struct {
	// Username, which is used for login.
	User *string `json:"User,omitnil" name:"User"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRabbitMQUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQUserResponseParams `json:"Response"`
}

func (r *CreateRabbitMQUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQVipInstanceRequestParams struct {
	// AZ
	ZoneIds []*int64 `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Node specification (`rabbit-vip-basic-1`: Basic; `rabbit-vip-basic-2`: Standard; `rabbit-vip-basic-3`: Advanced I; `rabbit-vip-basic-4`: Advanced II). If this parameter is left empty, the default value is `rabbit-vip-basic-1`.
	NodeSpec *string `json:"NodeSpec,omitnil" name:"NodeSpec"`

	// Number of nodes, which is at least three for multi-AZ deployment. If this parameter is left empty, the value will be set to 1 for single-AZ deployment and 3 for multi-AZ deployment by default.
	NodeNum *int64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// Storage capacity of a single node, which is 200 GB by default.
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// Whether to enable mirrored queue. Default value: `false`.
	EnableCreateDefaultHaMirrorQueue *bool `json:"EnableCreateDefaultHaMirrorQueue,omitnil" name:"EnableCreateDefaultHaMirrorQueue"`

	// Whether to enable auto-renewal. Default value: `true`.
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Validity period, which is one month by default.
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`
}

type CreateRabbitMQVipInstanceRequest struct {
	*tchttp.BaseRequest
	
	// AZ
	ZoneIds []*int64 `json:"ZoneIds,omitnil" name:"ZoneIds"`

	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// VPC subnet ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Node specification (`rabbit-vip-basic-1`: Basic; `rabbit-vip-basic-2`: Standard; `rabbit-vip-basic-3`: Advanced I; `rabbit-vip-basic-4`: Advanced II). If this parameter is left empty, the default value is `rabbit-vip-basic-1`.
	NodeSpec *string `json:"NodeSpec,omitnil" name:"NodeSpec"`

	// Number of nodes, which is at least three for multi-AZ deployment. If this parameter is left empty, the value will be set to 1 for single-AZ deployment and 3 for multi-AZ deployment by default.
	NodeNum *int64 `json:"NodeNum,omitnil" name:"NodeNum"`

	// Storage capacity of a single node, which is 200 GB by default.
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// Whether to enable mirrored queue. Default value: `false`.
	EnableCreateDefaultHaMirrorQueue *bool `json:"EnableCreateDefaultHaMirrorQueue,omitnil" name:"EnableCreateDefaultHaMirrorQueue"`

	// Whether to enable auto-renewal. Default value: `true`.
	AutoRenewFlag *bool `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Validity period, which is one month by default.
	TimeSpan *int64 `json:"TimeSpan,omitnil" name:"TimeSpan"`
}

func (r *CreateRabbitMQVipInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQVipInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ZoneIds")
	delete(f, "VpcId")
	delete(f, "SubnetId")
	delete(f, "ClusterName")
	delete(f, "NodeSpec")
	delete(f, "NodeNum")
	delete(f, "StorageSize")
	delete(f, "EnableCreateDefaultHaMirrorQueue")
	delete(f, "AutoRenewFlag")
	delete(f, "TimeSpan")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQVipInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQVipInstanceResponseParams struct {
	// Order ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	TranId *string `json:"TranId,omitnil" name:"TranId"`

	// Instance ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRabbitMQVipInstanceResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQVipInstanceResponseParams `json:"Response"`
}

func (r *CreateRabbitMQVipInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQVirtualHostRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Vhost name
	VirtualHost *string `json:"VirtualHost,omitnil" name:"VirtualHost"`

	// Description
	Description *string `json:"Description,omitnil" name:"Description"`

	// Message trace flag. Valid values: `true` (Enabled), `false` (Disabled, which is the default value).
	TraceFlag *bool `json:"TraceFlag,omitnil" name:"TraceFlag"`
}

type CreateRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Vhost name
	VirtualHost *string `json:"VirtualHost,omitnil" name:"VirtualHost"`

	// Description
	Description *string `json:"Description,omitnil" name:"Description"`

	// Message trace flag. Valid values: `true` (Enabled), `false` (Disabled, which is the default value).
	TraceFlag *bool `json:"TraceFlag,omitnil" name:"TraceFlag"`
}

func (r *CreateRabbitMQVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Description")
	delete(f, "TraceFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRabbitMQVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRabbitMQVirtualHostResponseParams struct {
	// Vhost name
	VirtualHost *string `json:"VirtualHost,omitnil" name:"VirtualHost"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *CreateRabbitMQVirtualHostResponseParams `json:"Response"`
}

func (r *CreateRabbitMQVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRocketMQClusterRequestParams struct {
	// Cluster name, which can contain 3–64 letters, digits, hyphens, and underscores
	Name *string `json:"Name,omitnil" name:"Name"`

	// Cluster description (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type CreateRocketMQClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster name, which can contain 3–64 letters, digits, hyphens, and underscores
	Name *string `json:"Name,omitnil" name:"Name"`

	// Cluster description (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`
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

// Predefined struct for user
type CreateRocketMQClusterResponseParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRocketMQGroupRequestParams struct {
	// Group name (8–64 characters)
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Namespace. Currently, only one namespace is supported
	Namespaces []*string `json:"Namespaces,omitnil" name:"Namespaces"`

	// Whether to enable consumption
	ReadEnable *bool `json:"ReadEnable,omitnil" name:"ReadEnable"`

	// Whether to enable broadcast consumption
	BroadcastEnable *bool `json:"BroadcastEnable,omitnil" name:"BroadcastEnable"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Group type (`TCP`, `HTTP`)
	GroupType *string `json:"GroupType,omitnil" name:"GroupType"`

	// The maximum number of retries for a group
	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitnil" name:"RetryMaxTimes"`
}

type CreateRocketMQGroupRequest struct {
	*tchttp.BaseRequest
	
	// Group name (8–64 characters)
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Namespace. Currently, only one namespace is supported
	Namespaces []*string `json:"Namespaces,omitnil" name:"Namespaces"`

	// Whether to enable consumption
	ReadEnable *bool `json:"ReadEnable,omitnil" name:"ReadEnable"`

	// Whether to enable broadcast consumption
	BroadcastEnable *bool `json:"BroadcastEnable,omitnil" name:"BroadcastEnable"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Group type (`TCP`, `HTTP`)
	GroupType *string `json:"GroupType,omitnil" name:"GroupType"`

	// The maximum number of retries for a group
	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitnil" name:"RetryMaxTimes"`
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
	delete(f, "GroupType")
	delete(f, "RetryMaxTimes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateRocketMQGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateRocketMQGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRocketMQNamespaceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name, which can contain 3–64 letters, digits, hyphens, and underscores
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// This parameter is disused.
	Ttl *uint64 `json:"Ttl,omitnil" name:"Ttl"`

	// This parameter is disused.
	RetentionTime *uint64 `json:"RetentionTime,omitnil" name:"RetentionTime"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type CreateRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name, which can contain 3–64 letters, digits, hyphens, and underscores
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// This parameter is disused.
	Ttl *uint64 `json:"Ttl,omitnil" name:"Ttl"`

	// This parameter is disused.
	RetentionTime *uint64 `json:"RetentionTime,omitnil" name:"RetentionTime"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`
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

// Predefined struct for user
type CreateRocketMQNamespaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQNamespaceResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRocketMQTopicRequestParams struct {
	// Topic name, which can contain 3–64 letters, digits, hyphens, and underscores
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Topic namespace. Currently, you can create topics only in one single namespace.
	Namespaces []*string `json:"Namespaces,omitnil" name:"Namespaces"`

	// Topic type. Valid values: `Normal`, `PartitionedOrder`, `Transaction`, `DelayScheduled`.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Topic remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Number of partitions, which doesn't take effect for globally sequential messages
	PartitionNum *int64 `json:"PartitionNum,omitnil" name:"PartitionNum"`
}

type CreateRocketMQTopicRequest struct {
	*tchttp.BaseRequest
	
	// Topic name, which can contain 3–64 letters, digits, hyphens, and underscores
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Topic namespace. Currently, you can create topics only in one single namespace.
	Namespaces []*string `json:"Namespaces,omitnil" name:"Namespaces"`

	// Topic type. Valid values: `Normal`, `PartitionedOrder`, `Transaction`, `DelayScheduled`.
	Type *string `json:"Type,omitnil" name:"Type"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Topic remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Number of partitions, which doesn't take effect for globally sequential messages
	PartitionNum *int64 `json:"PartitionNum,omitnil" name:"PartitionNum"`
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

// Predefined struct for user
type CreateRocketMQTopicResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateRocketMQTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateRoleRequestParams struct {
	// Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest
	
	// Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type CreateRoleResponseParams struct {
	// Role name
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Role token
	Token *string `json:"Token,omitnil" name:"Token"`

	// Remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Namespaces that are bound in batches
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	EnvironmentRoleSets []*EnvironmentRoleSet `json:"EnvironmentRoleSets,omitnil" name:"EnvironmentRoleSets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse
	Response *CreateRoleResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateSubscriptionRequestParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscriber name, which can contain up to 128 characters.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// Whether the creation is idempotent; if not, you cannot create subscriptions with the same name.
	IsIdempotent *bool `json:"IsIdempotent,omitnil" name:"IsIdempotent"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Whether to automatically create a dead letter topic and a retry letter topic. true: yes (default value); false: no.
	AutoCreatePolicyTopic *bool `json:"AutoCreatePolicyTopic,omitnil" name:"AutoCreatePolicyTopic"`

	// Naming convention for dead letter and retry letter topics. `LEGACY` indicates to use the legacy naming convention, and `COMMUNITY` indicates to use the naming convention in the Pulsar community.
	PostFixPattern *string `json:"PostFixPattern,omitnil" name:"PostFixPattern"`
}

type CreateSubscriptionRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscriber name, which can contain up to 128 characters.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// Whether the creation is idempotent; if not, you cannot create subscriptions with the same name.
	IsIdempotent *bool `json:"IsIdempotent,omitnil" name:"IsIdempotent"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Whether to automatically create a dead letter topic and a retry letter topic. true: yes (default value); false: no.
	AutoCreatePolicyTopic *bool `json:"AutoCreatePolicyTopic,omitnil" name:"AutoCreatePolicyTopic"`

	// Naming convention for dead letter and retry letter topics. `LEGACY` indicates to use the legacy naming convention, and `COMMUNITY` indicates to use the naming convention in the Pulsar community.
	PostFixPattern *string `json:"PostFixPattern,omitnil" name:"PostFixPattern"`
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

// Predefined struct for user
type CreateSubscriptionResponseParams struct {
	// Creation result.
	Result *bool `json:"Result,omitnil" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateSubscriptionResponse struct {
	*tchttp.BaseResponse
	Response *CreateSubscriptionResponseParams `json:"Response"`
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

// Predefined struct for user
type CreateTopicRequestParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name, which can contain up to 64 letters, digits, hyphens, and underscores.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// The value “1” indicates a non-partitioned topic (a topic with no partitions) will be created. A value between 1 (exclusive) and 128 (inclusive) indicates the partition count of a partitioned topic.
	Partitions *uint64 `json:"Partitions,omitnil" name:"Partitions"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// This input parameter will be disused soon. You can use `PulsarTopicType` instead.
	// 0: General message;
	// 1: Globally sequential message;
	// 2: Partitionally sequential message;
	// 3: Retry letter topic;
	// 4: Dead letter topic.
	TopicType *uint64 `json:"TopicType,omitnil" name:"TopicType"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Pulsar topic type.
	// `0`: Non-persistent and non-partitioned
	// `1`: Non-persistent and partitioned
	// `2`: Persistent and non-partitioned
	// `3`: Persistent and partitioned
	PulsarTopicType *int64 `json:"PulsarTopicType,omitnil" name:"PulsarTopicType"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name, which can contain up to 64 letters, digits, hyphens, and underscores.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// The value “1” indicates a non-partitioned topic (a topic with no partitions) will be created. A value between 1 (exclusive) and 128 (inclusive) indicates the partition count of a partitioned topic.
	Partitions *uint64 `json:"Partitions,omitnil" name:"Partitions"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// This input parameter will be disused soon. You can use `PulsarTopicType` instead.
	// 0: General message;
	// 1: Globally sequential message;
	// 2: Partitionally sequential message;
	// 3: Retry letter topic;
	// 4: Dead letter topic.
	TopicType *uint64 `json:"TopicType,omitnil" name:"TopicType"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Pulsar topic type.
	// `0`: Non-persistent and non-partitioned
	// `1`: Non-persistent and partitioned
	// `2`: Persistent and non-partitioned
	// `3`: Persistent and partitioned
	PulsarTopicType *int64 `json:"PulsarTopicType,omitnil" name:"PulsarTopicType"`
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

// Predefined struct for user
type CreateTopicResponseParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Valid value: 0 or 1. Non-partitioned topic: No partitions. A value greater than 1: The partition count of a partitioned topic. `0` is returned for existing non-partitioned topics, and `1` is returned for incremental non-partitioned topics.
	Partitions *uint64 `json:"Partitions,omitnil" name:"Partitions"`

	// Remarks (up to 128 characters).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// 0: General message;
	// 1: Globally sequential message;
	// 2: Partitionally sequential message;
	// 3: Retry letter topic;
	// 4: Dead letter topic.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicType *uint64 `json:"TopicType,omitnil" name:"TopicType"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *CreateTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteClusterRequestParams struct {
	// ID of the cluster to be deleted.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster to be deleted.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type DeleteClusterResponseParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteCmqQueueRequestParams struct {
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
}

type DeleteCmqQueueRequest struct {
	*tchttp.BaseRequest
	
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
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

// Predefined struct for user
type DeleteCmqQueueResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCmqQueueResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteCmqSubscribeRequestParams struct {
	// Topic name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`
}

type DeleteCmqSubscribeRequest struct {
	*tchttp.BaseRequest
	
	// Topic name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`
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

// Predefined struct for user
type DeleteCmqSubscribeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCmqSubscribeResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCmqSubscribeResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteCmqTopicRequestParams struct {
	// Topic name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`
}

type DeleteCmqTopicRequest struct {
	*tchttp.BaseRequest
	
	// Topic name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`
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

// Predefined struct for user
type DeleteCmqTopicResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteCmqTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCmqTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteEnvironmentRolesRequestParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Array of role names.
	RoleNames []*string `json:"RoleNames,omitnil" name:"RoleNames"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DeleteEnvironmentRolesRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Array of role names.
	RoleNames []*string `json:"RoleNames,omitnil" name:"RoleNames"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type DeleteEnvironmentRolesResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteEnvironmentRolesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEnvironmentRolesResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteEnvironmentsRequestParams struct {
	// Array of environments (namespaces). Up to 20 environments can be deleted at a time.
	EnvironmentIds []*string `json:"EnvironmentIds,omitnil" name:"EnvironmentIds"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DeleteEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// Array of environments (namespaces). Up to 20 environments can be deleted at a time.
	EnvironmentIds []*string `json:"EnvironmentIds,omitnil" name:"EnvironmentIds"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type DeleteEnvironmentsResponseParams struct {
	// Array of environments (namespaces) successfully deleted.
	EnvironmentIds []*string `json:"EnvironmentIds,omitnil" name:"EnvironmentIds"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteEnvironmentsResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRabbitMQUserRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username, which is used for login.
	User *string `json:"User,omitnil" name:"User"`
}

type DeleteRabbitMQUserRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username, which is used for login.
	User *string `json:"User,omitnil" name:"User"`
}

func (r *DeleteRabbitMQUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQUserResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRabbitMQUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQUserResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQVipInstanceRequestParams struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

type DeleteRabbitMQVipInstanceRequest struct {
	*tchttp.BaseRequest
	
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`
}

func (r *DeleteRabbitMQVipInstanceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQVipInstanceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQVipInstanceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQVipInstanceResponseParams struct {
	// Order ID
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	TranId *string `json:"TranId,omitnil" name:"TranId"`

	// Instance ID
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRabbitMQVipInstanceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQVipInstanceResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQVipInstanceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQVirtualHostRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Vhost name
	VirtualHost *string `json:"VirtualHost,omitnil" name:"VirtualHost"`
}

type DeleteRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Vhost name
	VirtualHost *string `json:"VirtualHost,omitnil" name:"VirtualHost"`
}

func (r *DeleteRabbitMQVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRabbitMQVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRabbitMQVirtualHostResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRabbitMQVirtualHostResponseParams `json:"Response"`
}

func (r *DeleteRabbitMQVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRocketMQClusterRequestParams struct {
	// ID of the cluster to be deleted.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DeleteRocketMQClusterRequest struct {
	*tchttp.BaseRequest
	
	// ID of the cluster to be deleted.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type DeleteRocketMQClusterResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRocketMQGroupRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Consumer group name
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
}

type DeleteRocketMQGroupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Consumer group name
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`
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

// Predefined struct for user
type DeleteRocketMQGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRocketMQNamespaceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`
}

type DeleteRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`
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

// Predefined struct for user
type DeleteRocketMQNamespaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQNamespaceResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRocketMQTopicRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Topic name
	Topic *string `json:"Topic,omitnil" name:"Topic"`
}

type DeleteRocketMQTopicRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Topic name
	Topic *string `json:"Topic,omitnil" name:"Topic"`
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

// Predefined struct for user
type DeleteRocketMQTopicResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRocketMQTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteRolesRequestParams struct {
	// Array of role names.
	RoleNames []*string `json:"RoleNames,omitnil" name:"RoleNames"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DeleteRolesRequest struct {
	*tchttp.BaseRequest
	
	// Array of role names.
	RoleNames []*string `json:"RoleNames,omitnil" name:"RoleNames"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type DeleteRolesResponseParams struct {
	// Name array of roles successfully deleted.
	RoleNames []*string `json:"RoleNames,omitnil" name:"RoleNames"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteRolesResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRolesResponseParams `json:"Response"`
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

// Predefined struct for user
type DeleteSubscriptionsRequestParams struct {
	// Subscription set. Up to 20 subscriptions can be deleted at a time.
	SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitnil" name:"SubscriptionTopicSets"`

	// Pulsar cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Whether to forcibly delete a subscription. Default value: `false`.
	Force *bool `json:"Force,omitnil" name:"Force"`
}

type DeleteSubscriptionsRequest struct {
	*tchttp.BaseRequest
	
	// Subscription set. Up to 20 subscriptions can be deleted at a time.
	SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitnil" name:"SubscriptionTopicSets"`

	// Pulsar cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Whether to forcibly delete a subscription. Default value: `false`.
	Force *bool `json:"Force,omitnil" name:"Force"`
}

func (r *DeleteSubscriptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubscriptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SubscriptionTopicSets")
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteSubscriptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteSubscriptionsResponseParams struct {
	// Array of successfully deleted subscriptions.
	SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitnil" name:"SubscriptionTopicSets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteSubscriptionsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteSubscriptionsResponseParams `json:"Response"`
}

func (r *DeleteSubscriptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteSubscriptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicsRequestParams struct {
	// Topic set. Up to 20 topics can be deleted at a time.
	TopicSets []*TopicRecord `json:"TopicSets,omitnil" name:"TopicSets"`

	// Pulsar cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Whether to forcibly delete a topic. Default value: `false`.
	Force *bool `json:"Force,omitnil" name:"Force"`
}

type DeleteTopicsRequest struct {
	*tchttp.BaseRequest
	
	// Topic set. Up to 20 topics can be deleted at a time.
	TopicSets []*TopicRecord `json:"TopicSets,omitnil" name:"TopicSets"`

	// Pulsar cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Whether to forcibly delete a topic. Default value: `false`.
	Force *bool `json:"Force,omitnil" name:"Force"`
}

func (r *DeleteTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "TopicSets")
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "Force")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteTopicsResponseParams struct {
	// Array of deleted topics.
	TopicSets []*TopicRecord `json:"TopicSets,omitnil" name:"TopicSets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DeleteTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DeleteTopicsResponseParams `json:"Response"`
}

func (r *DeleteTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindClustersRequestParams struct {

}

type DescribeBindClustersRequest struct {
	*tchttp.BaseRequest
	
}

func (r *DescribeBindClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeBindClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindClustersResponseParams struct {
	// The number of dedicated clusters
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of dedicated clusters
	ClusterSet []*BindCluster `json:"ClusterSet,omitnil" name:"ClusterSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBindClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindClustersResponseParams `json:"Response"`
}

func (r *DescribeBindClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeBindClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeBindVpcsRequestParams struct {
	// Offset. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribeBindVpcsRequest struct {
	*tchttp.BaseRequest
	
	// Offset. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type DescribeBindVpcsResponseParams struct {
	// Number of records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Set of VPCs.
	VpcSets []*VpcBindRecord `json:"VpcSets,omitnil" name:"VpcSets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeBindVpcsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeBindVpcsResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClusterDetailRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type DescribeClusterDetailResponseParams struct {
	// Cluster details
	ClusterSet *Cluster `json:"ClusterSet,omitnil" name:"ClusterSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClusterDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeClustersRequestParams struct {
	// Start offset, which defaults to 0 if left empty.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results to be returned, which defaults to 10 if left empty. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Filter by cluster ID.
	ClusterIdList []*string `json:"ClusterIdList,omitnil" name:"ClusterIdList"`

	// Whether to filter by tag.
	IsTagFilter *bool `json:"IsTagFilter,omitnil" name:"IsTagFilter"`

	// Filter. Currently, you can filter by tag.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest
	
	// Start offset, which defaults to 0 if left empty.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results to be returned, which defaults to 10 if left empty. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Filter by cluster ID.
	ClusterIdList []*string `json:"ClusterIdList,omitnil" name:"ClusterIdList"`

	// Whether to filter by tag.
	IsTagFilter *bool `json:"IsTagFilter,omitnil" name:"IsTagFilter"`

	// Filter. Currently, you can filter by tag.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterIdList")
	delete(f, "IsTagFilter")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeClustersResponseParams struct {
	// The number of clusters.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Cluster information list
	ClusterSet []*Cluster `json:"ClusterSet,omitnil" name:"ClusterSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeClustersResponseParams `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCmqDeadLetterSourceQueuesRequestParams struct {
	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil" name:"DeadLetterQueueName"`

	// Starting position of the list of topics to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Number of topics to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter by `SourceQueueName`
	SourceQueueName *string `json:"SourceQueueName,omitnil" name:"SourceQueueName"`
}

type DescribeCmqDeadLetterSourceQueuesRequest struct {
	*tchttp.BaseRequest
	
	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil" name:"DeadLetterQueueName"`

	// Starting position of the list of topics to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Number of topics to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Filter by `SourceQueueName`
	SourceQueueName *string `json:"SourceQueueName,omitnil" name:"SourceQueueName"`
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

// Predefined struct for user
type DescribeCmqDeadLetterSourceQueuesResponseParams struct {
	// Number of eligible queues
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Source queues of dead letter queue
	QueueSet []*CmqDeadLetterSource `json:"QueueSet,omitnil" name:"QueueSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCmqDeadLetterSourceQueuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqDeadLetterSourceQueuesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCmqQueueDetailRequestParams struct {
	// Exact match by `QueueName`
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
}

type DescribeCmqQueueDetailRequest struct {
	*tchttp.BaseRequest
	
	// Exact match by `QueueName`
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`
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

// Predefined struct for user
type DescribeCmqQueueDetailResponseParams struct {
	// List of queue details.
	QueueDescribe *CmqQueue `json:"QueueDescribe,omitnil" name:"QueueDescribe"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCmqQueueDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqQueueDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCmqQueuesRequestParams struct {
	// Starting position of a queue list to be returned on the current page in case of paginated return. If a value is entered, `limit` must be specified. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of queues to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Filter by `QueueName`
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// Filter by CMQ queue name.
	QueueNameList []*string `json:"QueueNameList,omitnil" name:"QueueNameList"`

	// For filtering by tag, this parameter must be set to `true`.
	IsTagFilter *bool `json:"IsTagFilter,omitnil" name:"IsTagFilter"`

	// Filter. Currently, you can filter by tag. The tag name must be prefixed with “tag:”, such as “tag: owner”, “tag: environment”, or “tag: business”.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeCmqQueuesRequest struct {
	*tchttp.BaseRequest
	
	// Starting position of a queue list to be returned on the current page in case of paginated return. If a value is entered, `limit` must be specified. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of queues to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Filter by `QueueName`
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// Filter by CMQ queue name.
	QueueNameList []*string `json:"QueueNameList,omitnil" name:"QueueNameList"`

	// For filtering by tag, this parameter must be set to `true`.
	IsTagFilter *bool `json:"IsTagFilter,omitnil" name:"IsTagFilter"`

	// Filter. Currently, you can filter by tag. The tag name must be prefixed with “tag:”, such as “tag: owner”, “tag: environment”, or “tag: business”.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeCmqQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqQueuesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "QueueName")
	delete(f, "QueueNameList")
	delete(f, "IsTagFilter")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqQueuesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCmqQueuesResponseParams struct {
	// The number of queues.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Queue list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	QueueList []*CmqQueue `json:"QueueList,omitnil" name:"QueueList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCmqQueuesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqQueuesResponseParams `json:"Response"`
}

func (r *DescribeCmqQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCmqSubscriptionDetailRequestParams struct {
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Starting position of the list of topics to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of topics to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Fuzzy search by `SubscriptionName`
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`
}

type DescribeCmqSubscriptionDetailRequest struct {
	*tchttp.BaseRequest
	
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Starting position of the list of topics to be returned on the current page in case of paginated return. If a value is entered, `limit` is required. If this parameter is left empty, 0 will be used by default
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Number of topics to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Fuzzy search by `SubscriptionName`
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`
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

// Predefined struct for user
type DescribeCmqSubscriptionDetailResponseParams struct {
	// Total number
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Set of subscription attributes
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubscriptionSet []*CmqSubscription `json:"SubscriptionSet,omitnil" name:"SubscriptionSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCmqSubscriptionDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqSubscriptionDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCmqTopicDetailRequestParams struct {
	// Exact match by `TopicName`.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`
}

type DescribeCmqTopicDetailRequest struct {
	*tchttp.BaseRequest
	
	// Exact match by `TopicName`.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`
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

// Predefined struct for user
type DescribeCmqTopicDetailResponseParams struct {
	// Topic details
	TopicDescribe *CmqTopic `json:"TopicDescribe,omitnil" name:"TopicDescribe"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCmqTopicDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqTopicDetailResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeCmqTopicsRequestParams struct {
	// Starting position of a queue list to be returned on the current page in case of paginated return. If a value is entered, `limit` must be specified. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of queues to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Fuzzy search by `TopicName`
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Filter by CMQ topic name.
	TopicNameList []*string `json:"TopicNameList,omitnil" name:"TopicNameList"`

	// For filtering by tag, this parameter must be set to `true`.
	IsTagFilter *bool `json:"IsTagFilter,omitnil" name:"IsTagFilter"`

	// Filter. Currently, you can filter by tag. The tag name must be prefixed with “tag:”, such as “tag: owner”, “tag: environment”, or “tag: business”.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeCmqTopicsRequest struct {
	*tchttp.BaseRequest
	
	// Starting position of a queue list to be returned on the current page in case of paginated return. If a value is entered, `limit` must be specified. If this parameter is left empty, 0 will be used by default.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of queues to be returned per page in case of paginated return. If this parameter is not passed in, 20 will be used by default. Maximum value: 50.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Fuzzy search by `TopicName`
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Filter by CMQ topic name.
	TopicNameList []*string `json:"TopicNameList,omitnil" name:"TopicNameList"`

	// For filtering by tag, this parameter must be set to `true`.
	IsTagFilter *bool `json:"IsTagFilter,omitnil" name:"IsTagFilter"`

	// Filter. Currently, you can filter by tag. The tag name must be prefixed with “tag:”, such as “tag: owner”, “tag: environment”, or “tag: business”.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeCmqTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TopicName")
	delete(f, "TopicNameList")
	delete(f, "IsTagFilter")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeCmqTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeCmqTopicsResponseParams struct {
	// Topic list.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicList []*CmqTopic `json:"TopicList,omitnil" name:"TopicList"`

	// The total number of topics.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeCmqTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeCmqTopicsResponseParams `json:"Response"`
}

func (r *DescribeCmqTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeCmqTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentAttributesRequestParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribeEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type DescribeEnvironmentAttributesResponseParams struct {
	// TTL for unconsumed messages in seconds. Maximum value: 1296000 seconds (i.e., 15 days).
	MsgTTL *uint64 `json:"MsgTTL,omitnil" name:"MsgTTL"`

	// Consumption rate limit in bytes/second. 0: unlimited.
	RateInByte *uint64 `json:"RateInByte,omitnil" name:"RateInByte"`

	// Consumption rate limit in messages/second. 0: unlimited.
	RateInSize *uint64 `json:"RateInSize,omitnil" name:"RateInSize"`

	// Retention policy for consumed messages in hours. 0: deleted immediately after consumption.
	RetentionHours *uint64 `json:"RetentionHours,omitnil" name:"RetentionHours"`

	// Retention policy for consumed messages in GB. 0: deleted immediately after consumption.
	RetentionSize *uint64 `json:"RetentionSize,omitnil" name:"RetentionSize"`

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Number of replicas.
	Replicas *uint64 `json:"Replicas,omitnil" name:"Replicas"`

	// Remarks.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentAttributesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeEnvironmentRolesRequestParams struct {
	// Environment/namespace name (required).
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Offset, which defaults to 0 if left empty.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results to be returned, which defaults to 10 if left empty. The maximum value is 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pulsar cluster ID (required).
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Role name.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// * RoleName
	// Filter by role name for exact query.
	// Type: String
	// Required: No
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeEnvironmentRolesRequest struct {
	*tchttp.BaseRequest
	
	// Environment/namespace name (required).
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Offset, which defaults to 0 if left empty.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results to be returned, which defaults to 10 if left empty. The maximum value is 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Pulsar cluster ID (required).
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Role name.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// * RoleName
	// Filter by role name for exact query.
	// Type: String
	// Required: No
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeEnvironmentRolesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentRolesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "RoleName")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentRolesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentRolesResponseParams struct {
	// The number of records.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Namespace role set.
	EnvironmentRoleSets []*EnvironmentRole `json:"EnvironmentRoleSets,omitnil" name:"EnvironmentRoleSets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEnvironmentRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentRolesResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentRolesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsRequestParams struct {
	// Fuzzy search by namespace name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Offset, which defaults to 0 if left empty.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results to be returned, which defaults to 10 if left empty. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// * EnvironmentId
	// Filter by namespace for exact query.
	// Type: String
	// Required: No
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest
	
	// Fuzzy search by namespace name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Offset, which defaults to 0 if left empty.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results to be returned, which defaults to 10 if left empty. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// * EnvironmentId
	// Filter by namespace for exact query.
	// Type: String
	// Required: No
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeEnvironmentsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeEnvironmentsResponseParams struct {
	// The number of namespaces.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Array of namespace sets.
	EnvironmentSet []*Environment `json:"EnvironmentSet,omitnil" name:"EnvironmentSet"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeEnvironmentsResponseParams `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePublisherSummaryRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name.
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// Topic name.
	Topic *string `json:"Topic,omitnil" name:"Topic"`
}

type DescribePublisherSummaryRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name.
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// Topic name.
	Topic *string `json:"Topic,omitnil" name:"Topic"`
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

// Predefined struct for user
type DescribePublisherSummaryResponseParams struct {
	// Production rate (messages/sec).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MsgRateIn *float64 `json:"MsgRateIn,omitnil" name:"MsgRateIn"`

	// Production rate (byte/sec).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MsgThroughputIn *float64 `json:"MsgThroughputIn,omitnil" name:"MsgThroughputIn"`

	// The number of producers.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	PublisherCount *int64 `json:"PublisherCount,omitnil" name:"PublisherCount"`

	// Message storage size in bytes.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	StorageSize *int64 `json:"StorageSize,omitnil" name:"StorageSize"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePublisherSummaryResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublisherSummaryResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePublishersRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name.
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// Topic name.
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Parameter filter. The `ProducerName` and `Address` fields are supported.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Offset for query. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// The number of query results displayed per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Sort by field.
	Sort *Sort `json:"Sort,omitnil" name:"Sort"`
}

type DescribePublishersRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name.
	Namespace *string `json:"Namespace,omitnil" name:"Namespace"`

	// Topic name.
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Parameter filter. The `ProducerName` and `Address` fields are supported.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Offset for query. Default value: `0`.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// The number of query results displayed per page. Default value: `20`.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Sort by field.
	Sort *Sort `json:"Sort,omitnil" name:"Sort"`
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

// Predefined struct for user
type DescribePublishersResponseParams struct {
	// Total number of query results.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of producer information.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Publishers []*Publisher `json:"Publishers,omitnil" name:"Publishers"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePublishersResponse struct {
	*tchttp.BaseResponse
	Response *DescribePublishersResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribePulsarProInstanceDetailRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribePulsarProInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *DescribePulsarProInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePulsarProInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePulsarProInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePulsarProInstanceDetailResponseParams struct {
	// Cluster information
	ClusterInfo *PulsarProClusterInfo `json:"ClusterInfo,omitnil" name:"ClusterInfo"`

	// Cluster network access point information
	// Note: This field may return null, indicating that no valid values can be obtained.
	NetworkAccessPointInfos []*PulsarNetworkAccessPointInfo `json:"NetworkAccessPointInfos,omitnil" name:"NetworkAccessPointInfos"`

	// Cluster specification information
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterSpecInfo *PulsarProClusterSpecInfo `json:"ClusterSpecInfo,omitnil" name:"ClusterSpecInfo"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePulsarProInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribePulsarProInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribePulsarProInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePulsarProInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePulsarProInstancesRequestParams struct {
	// Query condition filter
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// The maximum number of queried items, which defaults to `20`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Start offset for query
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribePulsarProInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Query condition filter
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// The maximum number of queried items, which defaults to `20`.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Start offset for query
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribePulsarProInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePulsarProInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribePulsarProInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribePulsarProInstancesResponseParams struct {
	// The total number of unpaginated items
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Instance information list
	Instances []*PulsarProInstance `json:"Instances,omitnil" name:"Instances"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribePulsarProInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribePulsarProInstancesResponseParams `json:"Response"`
}

func (r *DescribePulsarProInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribePulsarProInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQNodeListRequestParams struct {
	// TDMQ for RabbitMQ cluster ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The maximum entries per page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Node name for fuzzy search
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// Name and value of a filter.
	// Currently, only the `nodeStatus` filter is supported.
	// Valid values: `running`, `down`.
	// It is an array type and can contain multiple filters.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Sorting by a specified element.
	// Valid values: `cpuUsage`, `diskUsage`.
	SortElement *string `json:"SortElement,omitnil" name:"SortElement"`

	// Sorting order.
	// Valid values: `ascend`, `descend`.
	SortOrder *string `json:"SortOrder,omitnil" name:"SortOrder"`
}

type DescribeRabbitMQNodeListRequest struct {
	*tchttp.BaseRequest
	
	// TDMQ for RabbitMQ cluster ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The maximum entries per page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Node name for fuzzy search
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// Name and value of a filter.
	// Currently, only the `nodeStatus` filter is supported.
	// Valid values: `running`, `down`.
	// It is an array type and can contain multiple filters.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Sorting by a specified element.
	// Valid values: `cpuUsage`, `diskUsage`.
	SortElement *string `json:"SortElement,omitnil" name:"SortElement"`

	// Sorting order.
	// Valid values: `ascend`, `descend`.
	SortOrder *string `json:"SortOrder,omitnil" name:"SortOrder"`
}

func (r *DescribeRabbitMQNodeListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQNodeListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NodeName")
	delete(f, "Filters")
	delete(f, "SortElement")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQNodeListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQNodeListResponseParams struct {
	// The number of clusters
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Cluster list
	// Note: This field may return null, indicating that no valid value can be obtained.
	NodeList []*RabbitMQPrivateNode `json:"NodeList,omitnil" name:"NodeList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRabbitMQNodeListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQNodeListResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQNodeListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQNodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQUserRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username search by prefix or suffix
	SearchUser *string `json:"SearchUser,omitnil" name:"SearchUser"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Username, which is queried by exact match.
	User *string `json:"User,omitnil" name:"User"`

	// User tag, which is used to filter users.
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

type DescribeRabbitMQUserRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username search by prefix or suffix
	SearchUser *string `json:"SearchUser,omitnil" name:"SearchUser"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Username, which is queried by exact match.
	User *string `json:"User,omitnil" name:"User"`

	// User tag, which is used to filter users.
	Tags []*string `json:"Tags,omitnil" name:"Tags"`
}

func (r *DescribeRabbitMQUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "SearchUser")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "User")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQUserResponseParams struct {
	// Returned number of users
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The list of the created TDMQ for RabbitMQ users
	RabbitMQUserList []*RabbitMQUser `json:"RabbitMQUserList,omitnil" name:"RabbitMQUserList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRabbitMQUserResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQUserResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVipInstancesRequestParams struct {
	// Query condition filter
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// The maximum number of queried items, which defaults to 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Start offset for query
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeRabbitMQVipInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Query condition filter
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// The maximum number of queried items, which defaults to 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Start offset for query
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeRabbitMQVipInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVipInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQVipInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVipInstancesResponseParams struct {
	// The total number of unpaginated items
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Instance information list
	Instances []*RabbitMQVipInstance `json:"Instances,omitnil" name:"Instances"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRabbitMQVipInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQVipInstancesResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQVipInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVipInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVirtualHostListRequestParams struct {
	// A default parameter that won’t be used
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The maximum number of entries per page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

type DescribeRabbitMQVirtualHostListRequest struct {
	*tchttp.BaseRequest
	
	// A default parameter that won’t be used
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Offset
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The maximum number of entries per page
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`
}

func (r *DescribeRabbitMQVirtualHostListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVirtualHostListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Offset")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQVirtualHostListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVirtualHostListResponseParams struct {
	// The number of clusters
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Cluster list
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	VirtualHostList []*RabbitMQPrivateVirtualHost `json:"VirtualHostList,omitnil" name:"VirtualHostList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRabbitMQVirtualHostListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQVirtualHostListResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQVirtualHostListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVirtualHostListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVirtualHostRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Vhost name. If this parameter is not specified, all will be queried by default.
	VirtualHost *string `json:"VirtualHost,omitnil" name:"VirtualHost"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Fuzzy query by vhost name
	Filters *Filter `json:"Filters,omitnil" name:"Filters"`


	SortElement *string `json:"SortElement,omitnil" name:"SortElement"`


	SortOrder *string `json:"SortOrder,omitnil" name:"SortOrder"`
}

type DescribeRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Vhost name. If this parameter is not specified, all will be queried by default.
	VirtualHost *string `json:"VirtualHost,omitnil" name:"VirtualHost"`

	// Pagination offset
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Pagination limit
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Fuzzy query by vhost name
	Filters *Filter `json:"Filters,omitnil" name:"Filters"`

	SortElement *string `json:"SortElement,omitnil" name:"SortElement"`

	SortOrder *string `json:"SortOrder,omitnil" name:"SortOrder"`
}

func (r *DescribeRabbitMQVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "Filters")
	delete(f, "SortElement")
	delete(f, "SortOrder")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRabbitMQVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRabbitMQVirtualHostResponseParams struct {
	// Returned number of vhosts
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of vhost details
	VirtualHostList []*RabbitMQVirtualHostInfo `json:"VirtualHostList,omitnil" name:"VirtualHostList"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRabbitMQVirtualHostResponseParams `json:"Response"`
}

func (r *DescribeRabbitMQVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQClusterRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribeRocketMQClusterRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type DescribeRocketMQClusterResponseParams struct {
	// Cluster information
	ClusterInfo *RocketMQClusterInfo `json:"ClusterInfo,omitnil" name:"ClusterInfo"`

	// Cluster configuration
	ClusterConfig *RocketMQClusterConfig `json:"ClusterConfig,omitnil" name:"ClusterConfig"`

	// Recent cluster usage
	// Note: this field may return null, indicating that no valid values can be obtained.
	ClusterStats *RocketMQClusterRecentStats `json:"ClusterStats,omitnil" name:"ClusterStats"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeRocketMQClustersRequestParams struct {
	// Offset.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The max number of returned results.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Search by cluster ID.
	IdKeyword *string `json:"IdKeyword,omitnil" name:"IdKeyword"`

	// Search by cluster name.
	NameKeyword *string `json:"NameKeyword,omitnil" name:"NameKeyword"`

	// Filter by cluster ID.
	ClusterIdList []*string `json:"ClusterIdList,omitnil" name:"ClusterIdList"`

	// For filtering by tag, this parameter must be set to `true`.
	IsTagFilter *bool `json:"IsTagFilter,omitnil" name:"IsTagFilter"`

	// Filter. Currently, you can filter only by tag.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeRocketMQClustersRequest struct {
	*tchttp.BaseRequest
	
	// Offset.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The max number of returned results.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Search by cluster ID.
	IdKeyword *string `json:"IdKeyword,omitnil" name:"IdKeyword"`

	// Search by cluster name.
	NameKeyword *string `json:"NameKeyword,omitnil" name:"NameKeyword"`

	// Filter by cluster ID.
	ClusterIdList []*string `json:"ClusterIdList,omitnil" name:"ClusterIdList"`

	// For filtering by tag, this parameter must be set to `true`.
	IsTagFilter *bool `json:"IsTagFilter,omitnil" name:"IsTagFilter"`

	// Filter. Currently, you can filter only by tag.
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

func (r *DescribeRocketMQClustersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQClustersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "IdKeyword")
	delete(f, "NameKeyword")
	delete(f, "ClusterIdList")
	delete(f, "IsTagFilter")
	delete(f, "Filters")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQClustersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQClustersResponseParams struct {
	// Cluster information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ClusterList []*RocketMQClusterDetail `json:"ClusterList,omitnil" name:"ClusterList"`

	// The total number of returned results.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRocketMQClustersResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQClustersResponseParams `json:"Response"`
}

func (r *DescribeRocketMQClustersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQGroupsRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace.
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The max number of returned results.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Topic name, which can be used to query all subscription groups under the topic
	FilterTopic *string `json:"FilterTopic,omitnil" name:"FilterTopic"`

	// Consumer group query by consumer group name. Fuzzy query is supported
	FilterGroup *string `json:"FilterGroup,omitnil" name:"FilterGroup"`

	// Sort by specified field. Valid values: `tps`, `accumulative`.
	SortedBy *string `json:"SortedBy,omitnil" name:"SortedBy"`

	// Sort in ascending or descending order. Valid values: `asc`, `desc`.
	SortOrder *string `json:"SortOrder,omitnil" name:"SortOrder"`

	// Subscription group name. After it is specified, the information of only this subscription group will be returned.
	FilterOneGroup *string `json:"FilterOneGroup,omitnil" name:"FilterOneGroup"`

	// Group type
	Types []*string `json:"Types,omitnil" name:"Types"`
}

type DescribeRocketMQGroupsRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace.
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The max number of returned results.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Topic name, which can be used to query all subscription groups under the topic
	FilterTopic *string `json:"FilterTopic,omitnil" name:"FilterTopic"`

	// Consumer group query by consumer group name. Fuzzy query is supported
	FilterGroup *string `json:"FilterGroup,omitnil" name:"FilterGroup"`

	// Sort by specified field. Valid values: `tps`, `accumulative`.
	SortedBy *string `json:"SortedBy,omitnil" name:"SortedBy"`

	// Sort in ascending or descending order. Valid values: `asc`, `desc`.
	SortOrder *string `json:"SortOrder,omitnil" name:"SortOrder"`

	// Subscription group name. After it is specified, the information of only this subscription group will be returned.
	FilterOneGroup *string `json:"FilterOneGroup,omitnil" name:"FilterOneGroup"`

	// Group type
	Types []*string `json:"Types,omitnil" name:"Types"`
}

func (r *DescribeRocketMQGroupsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQGroupsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "FilterTopic")
	delete(f, "FilterGroup")
	delete(f, "SortedBy")
	delete(f, "SortOrder")
	delete(f, "FilterOneGroup")
	delete(f, "Types")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQGroupsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQGroupsResponseParams struct {
	// The total number of subscription groups.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of subscription groups
	Groups []*RocketMQGroup `json:"Groups,omitnil" name:"Groups"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRocketMQGroupsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQGroupsResponseParams `json:"Response"`
}

func (r *DescribeRocketMQGroupsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQMsgRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace ID
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name. Pass in the group ID when querying a dead letter queue.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Message ID
	MsgId *string `json:"MsgId,omitnil" name:"MsgId"`

	// ID of a TDMQ for Pulsar message
	PulsarMsgId *string `json:"PulsarMsgId,omitnil" name:"PulsarMsgId"`

	// The value of this parameter is `true` when you query a dead letter queue. It only applies to TDMQ for RocketMQ.
	QueryDlqMsg *bool `json:"QueryDlqMsg,omitnil" name:"QueryDlqMsg"`
}

type DescribeRocketMQMsgRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace ID
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name. Pass in the group ID when querying a dead letter queue.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Message ID
	MsgId *string `json:"MsgId,omitnil" name:"MsgId"`

	// ID of a TDMQ for Pulsar message
	PulsarMsgId *string `json:"PulsarMsgId,omitnil" name:"PulsarMsgId"`

	// The value of this parameter is `true` when you query a dead letter queue. It only applies to TDMQ for RocketMQ.
	QueryDlqMsg *bool `json:"QueryDlqMsg,omitnil" name:"QueryDlqMsg"`
}

func (r *DescribeRocketMQMsgRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQMsgRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "MsgId")
	delete(f, "PulsarMsgId")
	delete(f, "QueryDlqMsg")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQMsgRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQMsgResponseParams struct {
	// Message body
	Body *string `json:"Body,omitnil" name:"Body"`

	// Details parameter
	Properties *string `json:"Properties,omitnil" name:"Properties"`

	// Production time
	ProduceTime *string `json:"ProduceTime,omitnil" name:"ProduceTime"`

	// Message ID
	MsgId *string `json:"MsgId,omitnil" name:"MsgId"`

	// Producer address
	ProducerAddr *string `json:"ProducerAddr,omitnil" name:"ProducerAddr"`

	// Consumption details of a consumer group
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	MessageTracks []*RocketMQMessageTrack `json:"MessageTracks,omitnil" name:"MessageTracks"`

	// Topic name displayed on the details page
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ShowTopicName *string `json:"ShowTopicName,omitnil" name:"ShowTopicName"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRocketMQMsgResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQMsgResponseParams `json:"Response"`
}

func (r *DescribeRocketMQMsgResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQNamespacesRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The max number of returned results.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Search by name.
	NameKeyword *string `json:"NameKeyword,omitnil" name:"NameKeyword"`
}

type DescribeRocketMQNamespacesRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Offset.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The max number of returned results.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Search by name.
	NameKeyword *string `json:"NameKeyword,omitnil" name:"NameKeyword"`
}

func (r *DescribeRocketMQNamespacesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQNamespacesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "NameKeyword")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQNamespacesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQNamespacesResponseParams struct {
	// List of namespaces
	Namespaces []*RocketMQNamespace `json:"Namespaces,omitnil" name:"Namespaces"`

	// The total number of returned results.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRocketMQNamespacesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQNamespacesResponseParams `json:"Response"`
}

func (r *DescribeRocketMQNamespacesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopicsRequestParams struct {
	// Offset for query.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Query limit.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace.
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.
	FilterType []*string `json:"FilterType,omitnil" name:"FilterType"`

	// Search by topic name. Fuzzy query is supported.
	FilterName *string `json:"FilterName,omitnil" name:"FilterName"`
}

type DescribeRocketMQTopicsRequest struct {
	*tchttp.BaseRequest
	
	// Offset for query.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// Query limit.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace.
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.
	FilterType []*string `json:"FilterType,omitnil" name:"FilterType"`

	// Search by topic name. Fuzzy query is supported.
	FilterName *string `json:"FilterName,omitnil" name:"FilterName"`
}

func (r *DescribeRocketMQTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "FilterType")
	delete(f, "FilterName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQTopicsResponseParams struct {
	// The total number of query records.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// List of topic information
	Topics []*RocketMQTopic `json:"Topics,omitnil" name:"Topics"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRocketMQTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQTopicsResponseParams `json:"Response"`
}

func (r *DescribeRocketMQTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQVipInstanceDetailRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribeRocketMQVipInstanceDetailRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *DescribeRocketMQVipInstanceDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQVipInstanceDetailRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQVipInstanceDetailRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQVipInstanceDetailResponseParams struct {
	// Cluster information
	ClusterInfo *RocketMQClusterInfo `json:"ClusterInfo,omitnil" name:"ClusterInfo"`

	// Cluster configuration
	InstanceConfig *RocketMQInstanceConfig `json:"InstanceConfig,omitnil" name:"InstanceConfig"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRocketMQVipInstanceDetailResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQVipInstanceDetailResponseParams `json:"Response"`
}

func (r *DescribeRocketMQVipInstanceDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQVipInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQVipInstancesRequestParams struct {
	// Query condition filter
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// The maximum number of queried items, which defaults to 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Start offset for query
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

type DescribeRocketMQVipInstancesRequest struct {
	*tchttp.BaseRequest
	
	// Query condition filter
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// The maximum number of queried items, which defaults to 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Start offset for query
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`
}

func (r *DescribeRocketMQVipInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQVipInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "Filters")
	delete(f, "Limit")
	delete(f, "Offset")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRocketMQVipInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRocketMQVipInstancesResponseParams struct {
	// The total number of unpaginated items
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Instance information list
	Instances []*RocketMQVipInstance `json:"Instances,omitnil" name:"Instances"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRocketMQVipInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRocketMQVipInstancesResponseParams `json:"Response"`
}

func (r *DescribeRocketMQVipInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRocketMQVipInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRolesRequestParams struct {
	// Fuzzy query by role name
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Offset. If this parameter is left empty, 0 will be used by default.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// * RoleName
	// Filter by role name for exact query.
	// Type: String
	// Required: no
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
}

type DescribeRolesRequest struct {
	*tchttp.BaseRequest
	
	// Fuzzy query by role name
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Offset. If this parameter is left empty, 0 will be used by default.
	Offset *int64 `json:"Offset,omitnil" name:"Offset"`

	// Number of results to be returned. If this parameter is left empty, 10 will be used by default. The maximum value is 20.
	Limit *int64 `json:"Limit,omitnil" name:"Limit"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// * RoleName
	// Filter by role name for exact query.
	// Type: String
	// Required: no
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`
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

// Predefined struct for user
type DescribeRolesResponseParams struct {
	// Number of records.
	TotalCount *int64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// Array of roles.
	RoleSets []*Role `json:"RoleSets,omitnil" name:"RoleSets"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeRolesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRolesResponseParams `json:"Response"`
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

// Predefined struct for user
type DescribeSubscriptionsRequestParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Offset, which defaults to 0 if left empty.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results to be returned, which defaults to 10 if left empty. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Fuzzy match by subscriber name.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// Data filter.
	Filters []*FilterSubscription `json:"Filters,omitnil" name:"Filters"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type DescribeSubscriptionsRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Offset, which defaults to 0 if left empty.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results to be returned, which defaults to 10 if left empty. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Fuzzy match by subscriber name.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// Data filter.
	Filters []*FilterSubscription `json:"Filters,omitnil" name:"Filters"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

func (r *DescribeSubscriptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscriptionsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "SubscriptionName")
	delete(f, "Filters")
	delete(f, "ClusterId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeSubscriptionsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeSubscriptionsResponseParams struct {
	// Array of subscriber sets.
	SubscriptionSets []*Subscription `json:"SubscriptionSets,omitnil" name:"SubscriptionSets"`

	// The total number of returned results.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeSubscriptionsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeSubscriptionsResponseParams `json:"Response"`
}

func (r *DescribeSubscriptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeSubscriptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicsRequestParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Fuzzy match by topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Offset, which defaults to 0 if left empty.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results to be returned, which defaults to 10 if left empty. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Topic type description:
	// 0: Non-persistent and non-partitioned topic;
	// 1: Non-persistent and partitioned topic;
	// 2: Persistent and non-partitioned topic;
	// 3: Persistent and partitioned topic.
	TopicType *uint64 `json:"TopicType,omitnil" name:"TopicType"`

	// Pulsar cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// * TopicName
	// Query by topic name for exact search.
	// Type: String
	// Required: No
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Topic creator:
	// 1: User
	// 2: System
	TopicCreator *uint64 `json:"TopicCreator,omitnil" name:"TopicCreator"`
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Fuzzy match by topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Offset, which defaults to 0 if left empty.
	Offset *uint64 `json:"Offset,omitnil" name:"Offset"`

	// The number of results to be returned, which defaults to 10 if left empty. The maximum value is 20.
	Limit *uint64 `json:"Limit,omitnil" name:"Limit"`

	// Topic type description:
	// 0: Non-persistent and non-partitioned topic;
	// 1: Non-persistent and partitioned topic;
	// 2: Persistent and non-partitioned topic;
	// 3: Persistent and partitioned topic.
	TopicType *uint64 `json:"TopicType,omitnil" name:"TopicType"`

	// Pulsar cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// * TopicName
	// Query by topic name for exact search.
	// Type: String
	// Required: No
	Filters []*Filter `json:"Filters,omitnil" name:"Filters"`

	// Topic creator:
	// 1: User
	// 2: System
	TopicCreator *uint64 `json:"TopicCreator,omitnil" name:"TopicCreator"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "EnvironmentId")
	delete(f, "TopicName")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "TopicType")
	delete(f, "ClusterId")
	delete(f, "Filters")
	delete(f, "TopicCreator")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeTopicsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeTopicsResponseParams struct {
	// Array of topic sets.
	TopicSets []*Topic `json:"TopicSets,omitnil" name:"TopicSets"`

	// The number of topics.
	TotalCount *uint64 `json:"TotalCount,omitnil" name:"TotalCount"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeTopicsResponseParams `json:"Response"`
}

func (r *DescribeTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Environment struct {
	// Namespace name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Description.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Retention period for unconsumed messages in seconds. Maximum value: 1,296,000 seconds (15 days).
	MsgTTL *int64 `json:"MsgTTL,omitnil" name:"MsgTTL"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Last modified.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// Namespace ID.
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Namespace name.
	NamespaceName *string `json:"NamespaceName,omitnil" name:"NamespaceName"`

	// The number of topics.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicNum *int64 `json:"TopicNum,omitnil" name:"TopicNum"`

	// Message retention policy.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitnil" name:"RetentionPolicy"`

	// Whether to enable "Auto-Create Subscription"
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	AutoSubscriptionCreation *bool `json:"AutoSubscriptionCreation,omitnil" name:"AutoSubscriptionCreation"`
}

type EnvironmentRole struct {
	// Environment (namespace).
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Role name.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions []*string `json:"Permissions,omitnil" name:"Permissions"`

	// Role description.
	RoleDescribe *string `json:"RoleDescribe,omitnil" name:"RoleDescribe"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type EnvironmentRoleSet struct {
	// The IDs of the bound namespaces cannot be delicate and the namespaces must contain resources
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Permissions to be bound to a namespace. Enumerated values: `consume`, `produce`, and `consume, produce`. This parameter cannot be left empty.
	// 
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Permissions []*string `json:"Permissions,omitnil" name:"Permissions"`
}

type Filter struct {
	// Filter parameter name
	Name *string `json:"Name,omitnil" name:"Name"`

	// Value
	Values []*string `json:"Values,omitnil" name:"Values"`
}

type FilterSubscription struct {
	// Whether to only display subscriptions that include real consumers.
	ConsumerHasCount *bool `json:"ConsumerHasCount,omitnil" name:"ConsumerHasCount"`

	// Whether to only display subscriptions with heaped messages.
	ConsumerHasBacklog *bool `json:"ConsumerHasBacklog,omitnil" name:"ConsumerHasBacklog"`

	// Whether to only display subscriptions with messages discarded after expiration.
	ConsumerHasExpired *bool `json:"ConsumerHasExpired,omitnil" name:"ConsumerHasExpired"`

	// Filter by subscription name for exact query.
	SubscriptionNames []*string `json:"SubscriptionNames,omitnil" name:"SubscriptionNames"`
}

type InstanceNodeDistribution struct {
	// AZ
	ZoneName *string `json:"ZoneName,omitnil" name:"ZoneName"`

	// AZ ID
	ZoneId *string `json:"ZoneId,omitnil" name:"ZoneId"`

	// Number of nodes
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`
}

// Predefined struct for user
type ModifyClusterRequestParams struct {
	// ID of the Pulsar cluster to be updated.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Updated cluster name.
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Remarks.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Enables public network access, which can only be `true`.
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil" name:"PublicAccessEnabled"`
}

type ModifyClusterRequest struct {
	*tchttp.BaseRequest
	
	// ID of the Pulsar cluster to be updated.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Updated cluster name.
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Remarks.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Enables public network access, which can only be `true`.
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil" name:"PublicAccessEnabled"`
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

// Predefined struct for user
type ModifyClusterResponseParams struct {
	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyClusterResponse struct {
	*tchttp.BaseResponse
	Response *ModifyClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyCmqQueueAttributeRequestParams struct {
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// Maximum number of heaped messages. The value range is 1,000,000–10,000,000 during the beta test and can be 1,000,000–1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil" name:"MaxMsgHeapNum"`

	// Long polling wait time for message reception. Value range: 0–30 seconds. Default value: 0.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil" name:"PollingWaitSeconds"`

	// Message visibility timeout period. Value range: 1–43200 seconds (i.e., 12 hours). Default value: 30.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil" name:"VisibilityTimeout"`

	// Max message size, which defaults to 1,024 KB for the queue of TDMQ for CMQ and cannot be modified.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil" name:"MaxMsgSize"`

	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil" name:"MsgRetentionSeconds"`

	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil" name:"RewindSeconds"`

	// First query time
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil" name:"FirstQueryInterval"`

	// Maximum number of queries
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil" name:"MaxQueryCount"`

	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil" name:"DeadLetterQueueName"`

	// Maximum period in seconds before an unconsumed message expires, which is required if `MaxTimeToLivepolicy` is 1. Value range: 300–43200. This value should be smaller than `MsgRetentionSeconds` (maximum message retention period)
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil" name:"MaxTimeToLive"`

	// Maximum number of receipts
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil" name:"MaxReceiveCount"`

	// Dead letter queue policy
	Policy *uint64 `json:"Policy,omitnil" name:"Policy"`

	// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitnil" name:"Trace"`

	// Whether to enable transaction. 1: yes; 0: no
	Transaction *uint64 `json:"Transaction,omitnil" name:"Transaction"`

	// Queue storage space configured for message rewind. Value range: 10,240-512,000 MB (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitnil" name:"RetentionSizeInMB"`
}

type ModifyCmqQueueAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// Maximum number of heaped messages. The value range is 1,000,000–10,000,000 during the beta test and can be 1,000,000–1,000,000,000 after the product is officially released. The default value is 10,000,000 during the beta test and will be 100,000,000 after the product is officially released.
	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitnil" name:"MaxMsgHeapNum"`

	// Long polling wait time for message reception. Value range: 0–30 seconds. Default value: 0.
	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitnil" name:"PollingWaitSeconds"`

	// Message visibility timeout period. Value range: 1–43200 seconds (i.e., 12 hours). Default value: 30.
	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitnil" name:"VisibilityTimeout"`

	// Max message size, which defaults to 1,024 KB for the queue of TDMQ for CMQ and cannot be modified.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil" name:"MaxMsgSize"`

	// The max period during which a message is retained before it is automatically acknowledged. Value range: 30-43,200 seconds (30 seconds to 12 hours). Default value: 3600 seconds (1 hour).
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil" name:"MsgRetentionSeconds"`

	// Rewindable time of messages in the queue. Value range: 0-1,296,000s (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RewindSeconds *uint64 `json:"RewindSeconds,omitnil" name:"RewindSeconds"`

	// First query time
	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitnil" name:"FirstQueryInterval"`

	// Maximum number of queries
	MaxQueryCount *uint64 `json:"MaxQueryCount,omitnil" name:"MaxQueryCount"`

	// Dead letter queue name
	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitnil" name:"DeadLetterQueueName"`

	// Maximum period in seconds before an unconsumed message expires, which is required if `MaxTimeToLivepolicy` is 1. Value range: 300–43200. This value should be smaller than `MsgRetentionSeconds` (maximum message retention period)
	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitnil" name:"MaxTimeToLive"`

	// Maximum number of receipts
	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitnil" name:"MaxReceiveCount"`

	// Dead letter queue policy
	Policy *uint64 `json:"Policy,omitnil" name:"Policy"`

	// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitnil" name:"Trace"`

	// Whether to enable transaction. 1: yes; 0: no
	Transaction *uint64 `json:"Transaction,omitnil" name:"Transaction"`

	// Queue storage space configured for message rewind. Value range: 10,240-512,000 MB (if message rewind is enabled). The value “0” indicates that message rewind is not enabled.
	RetentionSizeInMB *uint64 `json:"RetentionSizeInMB,omitnil" name:"RetentionSizeInMB"`
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

// Predefined struct for user
type ModifyCmqQueueAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCmqQueueAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCmqQueueAttributeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyCmqSubscriptionAttributeRequestParams struct {
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// CMQ push server retry policy in case an error occurs while pushing a message to the endpoint. Valid values:
	// (1) BACKOFF_RETRY: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message.
	// (2) EXPONENTIAL_DECAY_RETRY: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: EXPONENTIAL_DECAY_RETRY.
	NotifyStrategy *string `json:"NotifyStrategy,omitnil" name:"NotifyStrategy"`

	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `HTTP`, both values are acceptable, and the default value is `JSON`.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitnil" name:"NotifyContentFormat"`

	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTags []*string `json:"FilterTags,omitnil" name:"FilterTags"`

	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKey []*string `json:"BindingKey,omitnil" name:"BindingKey"`
}

type ModifyCmqSubscriptionAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Topic name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscription name, which must be unique in the same topic under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// CMQ push server retry policy in case an error occurs while pushing a message to the endpoint. Valid values:
	// (1) BACKOFF_RETRY: backoff retry, which is to retry at a fixed interval, discard the message after a certain number of retries, and continue to push the next message.
	// (2) EXPONENTIAL_DECAY_RETRY: exponential decay retry, which is to retry at an exponentially increasing interval, such as 1s, 2s, 4s, 8s, and so on. As a message can be retained in a topic for one day, failed messages will be discarded at most after one day of retry. Default value: EXPONENTIAL_DECAY_RETRY.
	NotifyStrategy *string `json:"NotifyStrategy,omitnil" name:"NotifyStrategy"`

	// Push content format. Valid values: 1. JSON; 2. SIMPLIFIED, i.e., the raw format. If `Protocol` is `queue`, this value must be `SIMPLIFIED`. If `Protocol` is `HTTP`, both values are acceptable, and the default value is `JSON`.
	NotifyContentFormat *string `json:"NotifyContentFormat,omitnil" name:"NotifyContentFormat"`

	// Message body tag (used for message filtering). The number of tags cannot exceed 5, and each tag can contain up to 16 characters. It is used in conjunction with the `MsgTag` parameter of `(Batch)PublishMessage`. Rules: 1. If `FilterTag` is not configured, no matter whether `MsgTag` is configured, the subscription will receive all messages published to the topic; 2. If the array of `FilterTag` values has a value, only when at least one of the values in the array also exists in the array of `MsgTag` values (i.e., `FilterTag` and `MsgTag` have an intersection) can the subscription receive messages published to the topic; 3. If the array of `FilterTag` values has a value, but `MsgTag` is not configured, then no message published to the topic will be received, which can be considered as a special case of rule 2 as `FilterTag` and `MsgTag` do not intersect in this case. The overall design idea of rules is based on the intention of the subscriber.
	FilterTags []*string `json:"FilterTags,omitnil" name:"FilterTags"`

	// The number of `BindingKey` cannot exceed 5, and the length of each `BindingKey` cannot exceed 64 bytes. This field indicates the filtering policy for subscribing to and receiving messages. Each `BindingKey` includes up to 15 dots (namely up to 16 segments).
	BindingKey []*string `json:"BindingKey,omitnil" name:"BindingKey"`
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

// Predefined struct for user
type ModifyCmqSubscriptionAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCmqSubscriptionAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCmqSubscriptionAttributeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyCmqTopicAttributeRequestParams struct {
	// Topic name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil" name:"MaxMsgSize"`

	// Message retention period. Value range: 60–86400 seconds (i.e., 1 minute–1 day). Default value: 86400.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil" name:"MsgRetentionSeconds"`

	// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitnil" name:"Trace"`
}

type ModifyCmqTopicAttributeRequest struct {
	*tchttp.BaseRequest
	
	// Topic name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Maximum message length. Value range: 1024–65536 bytes (i.e., 1–64 KB). Default value: 65536.
	MaxMsgSize *uint64 `json:"MaxMsgSize,omitnil" name:"MaxMsgSize"`

	// Message retention period. Value range: 60–86400 seconds (i.e., 1 minute–1 day). Default value: 86400.
	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitnil" name:"MsgRetentionSeconds"`

	// Whether to enable message trace. true: yes; false: no. If this field is left empty, the feature will not be enabled.
	Trace *bool `json:"Trace,omitnil" name:"Trace"`
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

// Predefined struct for user
type ModifyCmqTopicAttributeResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyCmqTopicAttributeResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCmqTopicAttributeResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyEnvironmentAttributesRequestParams struct {
	// Namespace name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Retention period for unconsumed messages in seconds. Value range: 60s to 1,296,000s (or 15 days).
	MsgTTL *uint64 `json:"MsgTTL,omitnil" name:"MsgTTL"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Message retention policy
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitnil" name:"RetentionPolicy"`

	// Whether to enable "Auto-Create Subscription"
	AutoSubscriptionCreation *bool `json:"AutoSubscriptionCreation,omitnil" name:"AutoSubscriptionCreation"`
}

type ModifyEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest
	
	// Namespace name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Retention period for unconsumed messages in seconds. Value range: 60s to 1,296,000s (or 15 days).
	MsgTTL *uint64 `json:"MsgTTL,omitnil" name:"MsgTTL"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Message retention policy
	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitnil" name:"RetentionPolicy"`

	// Whether to enable "Auto-Create Subscription"
	AutoSubscriptionCreation *bool `json:"AutoSubscriptionCreation,omitnil" name:"AutoSubscriptionCreation"`
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
	delete(f, "AutoSubscriptionCreation")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyEnvironmentAttributesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyEnvironmentAttributesResponseParams struct {
	// Namespace name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// TTL for unconsumed messages in seconds.
	MsgTTL *uint64 `json:"MsgTTL,omitnil" name:"MsgTTL"`

	// Remarks (up to 128 characters).
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Namespace ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvironmentAttributesResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyEnvironmentRoleRequestParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Role name.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions []*string `json:"Permissions,omitnil" name:"Permissions"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type ModifyEnvironmentRoleRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Role name.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Permissions, which is a non-empty string array of `produce` and `consume` at the most.
	Permissions []*string `json:"Permissions,omitnil" name:"Permissions"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type ModifyEnvironmentRoleResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyEnvironmentRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyEnvironmentRoleResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRabbitMQUserRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username, which is used for login.
	User *string `json:"User,omitnil" name:"User"`

	// Password, which is used for login.
	Password *string `json:"Password,omitnil" name:"Password"`

	// Description. If this parameter is not passed in, it won't be modified.
	Description *string `json:"Description,omitnil" name:"Description"`

	// User tag, which defines a user's permission scope for accessing RabbitMQ Management. If this parameter is not passed in, it won't be modified.
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// The maximum number of connections for the user. If this parameter is not passed in, it won't be modified.
	MaxConnections *int64 `json:"MaxConnections,omitnil" name:"MaxConnections"`

	// The maximum number of channels for the user. If this parameter is not passed in, it won't be modified.
	MaxChannels *int64 `json:"MaxChannels,omitnil" name:"MaxChannels"`
}

type ModifyRabbitMQUserRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username, which is used for login.
	User *string `json:"User,omitnil" name:"User"`

	// Password, which is used for login.
	Password *string `json:"Password,omitnil" name:"Password"`

	// Description. If this parameter is not passed in, it won't be modified.
	Description *string `json:"Description,omitnil" name:"Description"`

	// User tag, which defines a user's permission scope for accessing RabbitMQ Management. If this parameter is not passed in, it won't be modified.
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// The maximum number of connections for the user. If this parameter is not passed in, it won't be modified.
	MaxConnections *int64 `json:"MaxConnections,omitnil" name:"MaxConnections"`

	// The maximum number of channels for the user. If this parameter is not passed in, it won't be modified.
	MaxChannels *int64 `json:"MaxChannels,omitnil" name:"MaxChannels"`
}

func (r *ModifyRabbitMQUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "User")
	delete(f, "Password")
	delete(f, "Description")
	delete(f, "Tags")
	delete(f, "MaxConnections")
	delete(f, "MaxChannels")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQUserResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRabbitMQUserResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQUserResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQVirtualHostRequestParams struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Vhost name
	VirtualHost *string `json:"VirtualHost,omitnil" name:"VirtualHost"`

	// Description
	Description *string `json:"Description,omitnil" name:"Description"`

	// Message trace flag. Valid values: `true` (Enabled), `false` (Disabled).
	TraceFlag *bool `json:"TraceFlag,omitnil" name:"TraceFlag"`
}

type ModifyRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest
	
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Vhost name
	VirtualHost *string `json:"VirtualHost,omitnil" name:"VirtualHost"`

	// Description
	Description *string `json:"Description,omitnil" name:"Description"`

	// Message trace flag. Valid values: `true` (Enabled), `false` (Disabled).
	TraceFlag *bool `json:"TraceFlag,omitnil" name:"TraceFlag"`
}

func (r *ModifyRabbitMQVirtualHostRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "VirtualHost")
	delete(f, "Description")
	delete(f, "TraceFlag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRabbitMQVirtualHostRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRabbitMQVirtualHostResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRabbitMQVirtualHostResponseParams `json:"Response"`
}

func (r *ModifyRabbitMQVirtualHostResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQClusterRequestParams struct {
	// RocketMQ cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 3–64 letters, digits, hyphens, and underscores
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Whether to enable the HTTP access over the public network
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil" name:"PublicAccessEnabled"`
}

type ModifyRocketMQClusterRequest struct {
	*tchttp.BaseRequest
	
	// RocketMQ cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// 3–64 letters, digits, hyphens, and underscores
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Whether to enable the HTTP access over the public network
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil" name:"PublicAccessEnabled"`
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
	delete(f, "PublicAccessEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQClusterRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQClusterResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRocketMQClusterResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQClusterResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRocketMQGroupRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Consumer group name
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Whether to enable consumption
	ReadEnable *bool `json:"ReadEnable,omitnil" name:"ReadEnable"`

	// Whether to enable broadcast consumption
	BroadcastEnable *bool `json:"BroadcastEnable,omitnil" name:"BroadcastEnable"`

	// The maximum number of retries
	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitnil" name:"RetryMaxTimes"`
}

type ModifyRocketMQGroupRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Consumer group name
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Whether to enable consumption
	ReadEnable *bool `json:"ReadEnable,omitnil" name:"ReadEnable"`

	// Whether to enable broadcast consumption
	BroadcastEnable *bool `json:"BroadcastEnable,omitnil" name:"BroadcastEnable"`

	// The maximum number of retries
	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitnil" name:"RetryMaxTimes"`
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
	delete(f, "RetryMaxTimes")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQGroupRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQGroupResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRocketMQGroupResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQGroupResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRocketMQInstanceSpecRequestParams struct {
	// ID of the exclusive instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance specification.
	// Valid values: `rocket-vip-basic-1` (Basic),
	// `rocket-vip-basic-2` (Standard),
	// `rocket-vip-basic-3` (Advanced I),
	// `rocket-vip-basic-4` (Advanced II).
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// Node count
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Storage space in GB
	StorageSize *uint64 `json:"StorageSize,omitnil" name:"StorageSize"`
}

type ModifyRocketMQInstanceSpecRequest struct {
	*tchttp.BaseRequest
	
	// ID of the exclusive instance
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance specification.
	// Valid values: `rocket-vip-basic-1` (Basic),
	// `rocket-vip-basic-2` (Standard),
	// `rocket-vip-basic-3` (Advanced I),
	// `rocket-vip-basic-4` (Advanced II).
	Specification *string `json:"Specification,omitnil" name:"Specification"`

	// Node count
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Storage space in GB
	StorageSize *uint64 `json:"StorageSize,omitnil" name:"StorageSize"`
}

func (r *ModifyRocketMQInstanceSpecRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQInstanceSpecRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "InstanceId")
	delete(f, "Specification")
	delete(f, "NodeCount")
	delete(f, "StorageSize")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQInstanceSpecRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQInstanceSpecResponseParams struct {
	// Order ID
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	OrderId *string `json:"OrderId,omitnil" name:"OrderId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRocketMQInstanceSpecResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQInstanceSpecResponseParams `json:"Response"`
}

func (r *ModifyRocketMQInstanceSpecResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyRocketMQInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQNamespaceRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name, which can contain 3–64 letters, digits, hyphens, and underscores
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// This parameter is disused.
	Ttl *uint64 `json:"Ttl,omitnil" name:"Ttl"`

	// This parameter is disused.
	RetentionTime *uint64 `json:"RetentionTime,omitnil" name:"RetentionTime"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Whether to enable the public network access
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil" name:"PublicAccessEnabled"`
}

type ModifyRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name, which can contain 3–64 letters, digits, hyphens, and underscores
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// This parameter is disused.
	Ttl *uint64 `json:"Ttl,omitnil" name:"Ttl"`

	// This parameter is disused.
	RetentionTime *uint64 `json:"RetentionTime,omitnil" name:"RetentionTime"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Whether to enable the public network access
	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitnil" name:"PublicAccessEnabled"`
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
	delete(f, "PublicAccessEnabled")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyRocketMQNamespaceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyRocketMQNamespaceResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQNamespaceResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRocketMQTopicRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Topic name
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Number of partitions, which is invalid for globally sequential messages and cannot be less than the current number of partitions.
	PartitionNum *int64 `json:"PartitionNum,omitnil" name:"PartitionNum"`
}

type ModifyRocketMQTopicRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Topic name
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Remarks (up to 128 characters)
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Number of partitions, which is invalid for globally sequential messages and cannot be less than the current number of partitions.
	PartitionNum *int64 `json:"PartitionNum,omitnil" name:"PartitionNum"`
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

// Predefined struct for user
type ModifyRocketMQTopicResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRocketMQTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRocketMQTopicResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyRoleRequestParams struct {
	// Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type ModifyRoleRequest struct {
	*tchttp.BaseRequest
	
	// Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Cluster ID (required)
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type ModifyRoleResponseParams struct {
	// Role name
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Remarks
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyRoleResponse struct {
	*tchttp.BaseResponse
	Response *ModifyRoleResponseParams `json:"Response"`
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

// Predefined struct for user
type ModifyTopicRequestParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Number of partitions, which must be equal to or greater than the original number of partitions. To maintain the original number of partitions, enter the original number. Modifying the number of partitions will take effect only for non-globally sequential messages. There can be up to 128 partitions.
	Partitions *uint64 `json:"Partitions,omitnil" name:"Partitions"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Number of partitions, which must be equal to or greater than the original number of partitions. To maintain the original number of partitions, enter the original number. Modifying the number of partitions will take effect only for non-globally sequential messages. There can be up to 128 partitions.
	Partitions *uint64 `json:"Partitions,omitnil" name:"Partitions"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type ModifyTopicResponseParams struct {
	// Number of partitions
	Partitions *uint64 `json:"Partitions,omitnil" name:"Partitions"`

	// Remarks (up to 128 characters).
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *ModifyTopicResponseParams `json:"Response"`
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

type PartitionsTopic struct {
	// Average size of the messages published in the last interval in bytes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	AverageMsgSize *string `json:"AverageMsgSize,omitnil" name:"AverageMsgSize"`

	// The number of consumers.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerCount *string `json:"ConsumerCount,omitnil" name:"ConsumerCount"`

	// The total number of recorded messages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitnil" name:"LastConfirmedEntry"`

	// Time when the last ledger was created.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitnil" name:"LastLedgerCreatedTimestamp"`

	// The number of messages published by local and replicated publishers per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgRateIn *string `json:"MsgRateIn,omitnil" name:"MsgRateIn"`

	// The total number of messages delivered by local and replicated consumers per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgRateOut *string `json:"MsgRateOut,omitnil" name:"MsgRateOut"`

	// The size (in bytes) of messages published by local and replicated publishers per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgThroughputIn *string `json:"MsgThroughputIn,omitnil" name:"MsgThroughputIn"`

	// The size (in bytes) of messages delivered by local and replicated consumers per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgThroughputOut *string `json:"MsgThroughputOut,omitnil" name:"MsgThroughputOut"`

	// The total number of recorded messages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NumberOfEntries *string `json:"NumberOfEntries,omitnil" name:"NumberOfEntries"`

	// Subpartition ID.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Partitions *int64 `json:"Partitions,omitnil" name:"Partitions"`

	// The number of producers.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProducerCount *string `json:"ProducerCount,omitnil" name:"ProducerCount"`

	// Total size of all stored messages in bytes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *string `json:"TotalSize,omitnil" name:"TotalSize"`

	// Topic type description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicType *uint64 `json:"TopicType,omitnil" name:"TopicType"`
}

// Predefined struct for user
type PublishCmqMsgRequestParams struct {
	// Topic name
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Message content. The total message size is up to 1,024 KB.
	MsgContent *string `json:"MsgContent,omitnil" name:"MsgContent"`

	// Message tag. You can pass in multiple tags or a single route. Each tag or route can contain up to 64 characters.
	MsgTag []*string `json:"MsgTag,omitnil" name:"MsgTag"`
}

type PublishCmqMsgRequest struct {
	*tchttp.BaseRequest
	
	// Topic name
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Message content. The total message size is up to 1,024 KB.
	MsgContent *string `json:"MsgContent,omitnil" name:"MsgContent"`

	// Message tag. You can pass in multiple tags or a single route. Each tag or route can contain up to 64 characters.
	MsgTag []*string `json:"MsgTag,omitnil" name:"MsgTag"`
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

// Predefined struct for user
type PublishCmqMsgResponseParams struct {
	// `true` indicates that the sending is successful
	Result *bool `json:"Result,omitnil" name:"Result"`

	// Message ID
	MsgId *string `json:"MsgId,omitnil" name:"MsgId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type PublishCmqMsgResponse struct {
	*tchttp.BaseResponse
	Response *PublishCmqMsgResponseParams `json:"Response"`
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
	ProducerId *int64 `json:"ProducerId,omitnil" name:"ProducerId"`

	// Producer name.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ProducerName *string `json:"ProducerName,omitnil" name:"ProducerName"`

	// Producer address.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Address *string `json:"Address,omitnil" name:"Address"`

	// Client version.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ClientVersion *string `json:"ClientVersion,omitnil" name:"ClientVersion"`

	// Message production rate (message/sec).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MsgRateIn *float64 `json:"MsgRateIn,omitnil" name:"MsgRateIn"`

	// Message production throughput rate (byte/sec).
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	MsgThroughputIn *float64 `json:"MsgThroughputIn,omitnil" name:"MsgThroughputIn"`

	// Average message size in bytes.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	AverageMsgSize *float64 `json:"AverageMsgSize,omitnil" name:"AverageMsgSize"`

	// Connection time.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	ConnectedSince *string `json:"ConnectedSince,omitnil" name:"ConnectedSince"`

	// Serial number of the topic partition connected to the producer.
	// Note: this field may return `null`, indicating that no valid values can be obtained.
	Partition *int64 `json:"Partition,omitnil" name:"Partition"`
}

type PulsarNetworkAccessPointInfo struct {
	// VPC ID. This field is left empty for supporting network and public network access points.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID. This field is left empty for supporting network and public network access points.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Access address
	Endpoint *string `json:"Endpoint,omitnil" name:"Endpoint"`

	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Access point type: 
	// `0`: Supporting network access point 
	// `1`: VPC access point 
	// `2`: Public network access point
	RouteType *uint64 `json:"RouteType,omitnil" name:"RouteType"`
}

type PulsarProClusterInfo struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Description
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Cluster status. Valid values: `0` (Creating), `1` (Normal), `2` (Isolated).
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Cluster version
	Version *string `json:"Version,omitnil" name:"Version"`

	// Node distribution
	// Note: This field may return null, indicating that no valid values can be obtained.
	NodeDistribution []*InstanceNodeDistribution `json:"NodeDistribution,omitnil" name:"NodeDistribution"`

	// Max storage capacity in MB
	MaxStorage *uint64 `json:"MaxStorage,omitnil" name:"MaxStorage"`

	// Whether the route can be modified
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	CanEditRoute *bool `json:"CanEditRoute,omitnil" name:"CanEditRoute"`
}

type PulsarProClusterSpecInfo struct {
	// Cluster specification name
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// Peak TPS
	MaxTps *uint64 `json:"MaxTps,omitnil" name:"MaxTps"`

	// Peak bandwidth in Mbps
	MaxBandWidth *uint64 `json:"MaxBandWidth,omitnil" name:"MaxBandWidth"`

	// Maximum number of namespaces
	MaxNamespaces *uint64 `json:"MaxNamespaces,omitnil" name:"MaxNamespaces"`

	// Maximum number of topic partitions
	MaxTopics *uint64 `json:"MaxTopics,omitnil" name:"MaxTopics"`

	// Elastic TPS beyond the specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScalableTps *uint64 `json:"ScalableTps,omitnil" name:"ScalableTps"`
}

type PulsarProInstance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance version
	InstanceVersion *string `json:"InstanceVersion,omitnil" name:"InstanceVersion"`

	// Instance status. Valid values: `0` (Creating), `1` (Normal), `2` (Isolated), `3` (Terminated), `4` (Abnormal), `5` (Delivery failed), `6` (Adjusting configuration), `7` (Configuration adjustment failed).
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Instance specification name
	ConfigDisplay *string `json:"ConfigDisplay,omitnil" name:"ConfigDisplay"`

	// Peak TPS
	MaxTps *uint64 `json:"MaxTps,omitnil" name:"MaxTps"`

	// Storage capacity in GB
	MaxStorage *uint64 `json:"MaxStorage,omitnil" name:"MaxStorage"`

	// Instance expiration time in milliseconds
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// Renewal mode. Valid values: `0` (Manual renewal, which is the default mode), `1` (Auto-renewal), `2` (Manual renewal, which is specified by users).
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Payment mode. Valid values: `0` (Pay-as-you-go), `1` (Monthly subscription).
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Instance specification ID
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// Elastic TPS beyond the specification
	// Note: This field may return null, indicating that no valid values can be obtained.
	ScalableTps *uint64 `json:"ScalableTps,omitnil" name:"ScalableTps"`

	// VPC ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`

	// Peak bandwidth in Mbps
	MaxBandWidth *uint64 `json:"MaxBandWidth,omitnil" name:"MaxBandWidth"`
}

type RabbitMQPrivateNode struct {
	// Node name
	// Note: This field may return null, indicating that no valid value can be obtained.
	NodeName *string `json:"NodeName,omitnil" name:"NodeName"`

	// Node status
	// Note: This field may return null, indicating that no valid value can be obtained.
	NodeStatus *string `json:"NodeStatus,omitnil" name:"NodeStatus"`

	// CPU utilization
	// Note: This field may return null, indicating that no valid values can be obtained.
	CPUUsage *string `json:"CPUUsage,omitnil" name:"CPUUsage"`

	// Memory usage in MB
	// Note: This field may return null, indicating that no valid values can be obtained.
	Memory *uint64 `json:"Memory,omitnil" name:"Memory"`

	// Disk utilization
	// Note: This field may return null, indicating that no valid values can be obtained.
	DiskUsage *string `json:"DiskUsage,omitnil" name:"DiskUsage"`

	// The number of RabbitMQ Erlang processes
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProcessNumber *uint64 `json:"ProcessNumber,omitnil" name:"ProcessNumber"`
}

type RabbitMQPrivateVirtualHost struct {
	// Vhost name
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	VirtualHostName *string `json:"VirtualHostName,omitnil" name:"VirtualHostName"`

	// Vhost description
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil" name:"Description"`
}

type RabbitMQUser struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Username, which is used for login.
	User *string `json:"User,omitnil" name:"User"`

	// Password, which is used for login.
	Password *string `json:"Password,omitnil" name:"Password"`

	// User description
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil" name:"Description"`

	// User tag, which defines a user's permission scope for accessing RabbitMQ Managementu200d.
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// User creation time
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Last user modification time
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// User type. Valid values: `System` (Created by system), `User` (Created by user).
	Type *string `json:"Type,omitnil" name:"Type"`
}

type RabbitMQVipInstance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance version
	// Note: This field may return null, indicating that no valid value can be obtained.
	InstanceVersion *string `json:"InstanceVersion,omitnil" name:"InstanceVersion"`

	// Instance status. Valid values: `0` (Creating), `1` (Normal), `2` (Isolated), `3` (Terminated), `4` (Abnormal), `5` (Delivery failed).
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Number of nodes
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Instance specification name
	ConfigDisplay *string `json:"ConfigDisplay,omitnil" name:"ConfigDisplay"`

	// Peak TPS
	MaxTps *uint64 `json:"MaxTps,omitnil" name:"MaxTps"`

	// Peak bandwidth in Mbps
	MaxBandWidth *uint64 `json:"MaxBandWidth,omitnil" name:"MaxBandWidth"`

	// Storage capacity in GB
	MaxStorage *uint64 `json:"MaxStorage,omitnil" name:"MaxStorage"`

	// Instance expiration time in milliseconds
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// Renewal mode. Valid values: `0` (Manual renewal, which is the default mode), `1` (Auto-renewal), `2` (Manual renewal, which is specified by users).
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Payment mode. `0`: Postpaid; `1`: Prepaid.
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Remarks
	// Note: This field may return null, indicating that no valid value can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Instance specification ID
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// Cluster exception
	// Note: This field may return null, indicating that no valid values can be obtained.
	ExceptionInformation *string `json:"ExceptionInformation,omitnil" name:"ExceptionInformation"`

	// Instance status. Valid values: `0` (Creating), `1` (Normal), `2` (Isolated), `3` (Terminated), `4` (Abnormal), `5` (Delivery failed).
	// This parameter is used to display the instance status additionally and distinguish from the `Status` parameter.
	ClusterStatus *int64 `json:"ClusterStatus,omitnil" name:"ClusterStatus"`
}

type RabbitMQVirtualHostInfo struct {
	// Cluster instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Vhost name
	VirtualHost *string `json:"VirtualHost,omitnil" name:"VirtualHost"`

	// Vhost description
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Description *string `json:"Description,omitnil" name:"Description"`

	// Vhost tag
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Tags []*string `json:"Tags,omitnil" name:"Tags"`

	// Creation time
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Modification time
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ModifyTime *string `json:"ModifyTime,omitnil" name:"ModifyTime"`

	// Statistics of vhost overview
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	VirtualHostStatistics *RabbitMQVirtualHostStatistics `json:"VirtualHostStatistics,omitnil" name:"VirtualHostStatistics"`


	Status *string `json:"Status,omitnil" name:"Status"`


	MessageHeapCount *int64 `json:"MessageHeapCount,omitnil" name:"MessageHeapCount"`


	MessageRateIn *float64 `json:"MessageRateIn,omitnil" name:"MessageRateIn"`


	MessageRateOut *float64 `json:"MessageRateOut,omitnil" name:"MessageRateOut"`
}

type RabbitMQVirtualHostStatistics struct {
	// The number of queues in the current vhost
	CurrentQueues *int64 `json:"CurrentQueues,omitnil" name:"CurrentQueues"`

	// The number of exchanges in the current vhost
	CurrentExchanges *int64 `json:"CurrentExchanges,omitnil" name:"CurrentExchanges"`

	// The number of connections in the current vhost
	CurrentConnections *int64 `json:"CurrentConnections,omitnil" name:"CurrentConnections"`

	// The number of channels in the current vhost
	CurrentChannels *int64 `json:"CurrentChannels,omitnil" name:"CurrentChannels"`

	// The number of users in the current vhost
	CurrentUsers *int64 `json:"CurrentUsers,omitnil" name:"CurrentUsers"`
}

// Predefined struct for user
type ReceiveMessageRequestParams struct {
	// Name of the topic which receives the message. It is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Subscriber name
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// Default value: 1000. Messages received by the consumer will first be stored in the `receiverQueueSize` queue to tune the message receiving rate.
	ReceiverQueueSize *int64 `json:"ReceiverQueueSize,omitnil" name:"ReceiverQueueSize"`

	// A parameter used to determine the position where the consumer initially receives messages. Valid values: `Earliest` (default), `Latest`.
	SubInitialPosition *string `json:"SubInitialPosition,omitnil" name:"SubInitialPosition"`

	// This parameter is used to specify the maximum number of received messages in a batch for `BatchReceivePolicy`. The default value is 0, indicating that `BatchReceivePolicy` is disabled.
	MaxNumMessages *int64 `json:"MaxNumMessages,omitnil" name:"MaxNumMessages"`

	// This parameter is used to specify the maximum body size (in bytes) of received messages in a batch for `BatchReceivePolicy`. The default value is 0, indicating that `BatchReceivePolicy` is disabled.
	MaxNumBytes *int64 `json:"MaxNumBytes,omitnil" name:"MaxNumBytes"`

	// This parameter is used to specify the maximum wait timeout (in milliseconds) for receiving a batch of messages for `BatchReceivePolicy`. The default value is 0, indicating that `BatchReceivePolicy` is disabled.
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`
}

type ReceiveMessageRequest struct {
	*tchttp.BaseRequest
	
	// Name of the topic which receives the message. It is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Subscriber name
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// Default value: 1000. Messages received by the consumer will first be stored in the `receiverQueueSize` queue to tune the message receiving rate.
	ReceiverQueueSize *int64 `json:"ReceiverQueueSize,omitnil" name:"ReceiverQueueSize"`

	// A parameter used to determine the position where the consumer initially receives messages. Valid values: `Earliest` (default), `Latest`.
	SubInitialPosition *string `json:"SubInitialPosition,omitnil" name:"SubInitialPosition"`

	// This parameter is used to specify the maximum number of received messages in a batch for `BatchReceivePolicy`. The default value is 0, indicating that `BatchReceivePolicy` is disabled.
	MaxNumMessages *int64 `json:"MaxNumMessages,omitnil" name:"MaxNumMessages"`

	// This parameter is used to specify the maximum body size (in bytes) of received messages in a batch for `BatchReceivePolicy`. The default value is 0, indicating that `BatchReceivePolicy` is disabled.
	MaxNumBytes *int64 `json:"MaxNumBytes,omitnil" name:"MaxNumBytes"`

	// This parameter is used to specify the maximum wait timeout (in milliseconds) for receiving a batch of messages for `BatchReceivePolicy`. The default value is 0, indicating that `BatchReceivePolicy` is disabled.
	Timeout *int64 `json:"Timeout,omitnil" name:"Timeout"`
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
	delete(f, "MaxNumMessages")
	delete(f, "MaxNumBytes")
	delete(f, "Timeout")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ReceiveMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ReceiveMessageResponseParams struct {
	// Unique primary key used to identify the message
	MessageID *string `json:"MessageID,omitnil" name:"MessageID"`

	// Content of the received message
	MessagePayload *string `json:"MessagePayload,omitnil" name:"MessagePayload"`

	// Provided to the `Ack` API and used to acknowledge messages in the topic
	AckTopic *string `json:"AckTopic,omitnil" name:"AckTopic"`

	// Returned error message. If it is an empty string, no error occurred.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// Returned subscriber name, which will be used when an acknowledgment consumer is created.
	// Note: this field may return null, indicating that no valid values can be obtained.
	SubName *string `json:"SubName,omitnil" name:"SubName"`

	// MessageIDs returned by `BatchReceivePolicy` at a time, which are separated by “###”.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MessageIDList *string `json:"MessageIDList,omitnil" name:"MessageIDList"`

	// Message contents returned by `BatchReceivePolicy` at a time, which are separated by “###”.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MessagesPayload *string `json:"MessagesPayload,omitnil" name:"MessagesPayload"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ReceiveMessageResponse struct {
	*tchttp.BaseResponse
	Response *ReceiveMessageResponseParams `json:"Response"`
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

// Predefined struct for user
type ResetMsgSubOffsetByTimestampRequestParams struct {
	// Namespace name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscriber name.
	Subscription *string `json:"Subscription,omitnil" name:"Subscription"`

	// Timestamp, accurate down to the millisecond.
	ToTimestamp *uint64 `json:"ToTimestamp,omitnil" name:"ToTimestamp"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type ResetMsgSubOffsetByTimestampRequest struct {
	*tchttp.BaseRequest
	
	// Namespace name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscriber name.
	Subscription *string `json:"Subscription,omitnil" name:"Subscription"`

	// Timestamp, accurate down to the millisecond.
	ToTimestamp *uint64 `json:"ToTimestamp,omitnil" name:"ToTimestamp"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type ResetMsgSubOffsetByTimestampResponseParams struct {
	// Result.
	// Note: this field may return null, indicating that no valid values can be obtained.
	Result *bool `json:"Result,omitnil" name:"Result"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetMsgSubOffsetByTimestampResponse struct {
	*tchttp.BaseResponse
	Response *ResetMsgSubOffsetByTimestampResponseParams `json:"Response"`
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

// Predefined struct for user
type ResetRocketMQConsumerOffSetRequestParams struct {
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name.
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Consumer group name.
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Topic name.
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Reset method. 0: Start from the latest offset; 1: Start from specified time point.
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// The specified timestamp that has been reset, in milliseconds. This parameter only takes effect when the value of `Type` is `1`.
	ResetTimestamp *uint64 `json:"ResetTimestamp,omitnil" name:"ResetTimestamp"`
}

type ResetRocketMQConsumerOffSetRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID.
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace name.
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Consumer group name.
	GroupId *string `json:"GroupId,omitnil" name:"GroupId"`

	// Topic name.
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Reset method. 0: Start from the latest offset; 1: Start from specified time point.
	Type *uint64 `json:"Type,omitnil" name:"Type"`

	// The specified timestamp that has been reset, in milliseconds. This parameter only takes effect when the value of `Type` is `1`.
	ResetTimestamp *uint64 `json:"ResetTimestamp,omitnil" name:"ResetTimestamp"`
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

// Predefined struct for user
type ResetRocketMQConsumerOffSetResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type ResetRocketMQConsumerOffSetResponse struct {
	*tchttp.BaseResponse
	Response *ResetRocketMQConsumerOffSetResponseParams `json:"Response"`
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
	TimeInMinutes *int64 `json:"TimeInMinutes,omitnil" name:"TimeInMinutes"`

	// Message retention size
	SizeInMB *int64 `json:"SizeInMB,omitnil" name:"SizeInMB"`
}

// Predefined struct for user
type RewindCmqQueueRequestParams struct {
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// After this time is configured, the `(Batch)receiveMessage` API will consume the messages received after this timestamp in the order in which they are produced.
	StartConsumeTime *uint64 `json:"StartConsumeTime,omitnil" name:"StartConsumeTime"`
}

type RewindCmqQueueRequest struct {
	*tchttp.BaseRequest
	
	// Queue name, which must be unique under the same account in the same region. It can contain up to 64 letters, digits, and hyphens and must begin with a letter.
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// After this time is configured, the `(Batch)receiveMessage` API will consume the messages received after this timestamp in the order in which they are produced.
	StartConsumeTime *uint64 `json:"StartConsumeTime,omitnil" name:"StartConsumeTime"`
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

// Predefined struct for user
type RewindCmqQueueResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type RewindCmqQueueResponse struct {
	*tchttp.BaseResponse
	Response *RewindCmqQueueResponseParams `json:"Response"`
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
	MaxTpsPerNamespace *uint64 `json:"MaxTpsPerNamespace,omitnil" name:"MaxTpsPerNamespace"`

	// Maximum number of namespaces
	MaxNamespaceNum *uint64 `json:"MaxNamespaceNum,omitnil" name:"MaxNamespaceNum"`

	// Number of used namespaces
	UsedNamespaceNum *uint64 `json:"UsedNamespaceNum,omitnil" name:"UsedNamespaceNum"`

	// Maximum number of topics
	MaxTopicNum *uint64 `json:"MaxTopicNum,omitnil" name:"MaxTopicNum"`

	// Number of used topics
	UsedTopicNum *uint64 `json:"UsedTopicNum,omitnil" name:"UsedTopicNum"`

	// Maximum number of groups
	MaxGroupNum *uint64 `json:"MaxGroupNum,omitnil" name:"MaxGroupNum"`

	// Number of used groups
	UsedGroupNum *uint64 `json:"UsedGroupNum,omitnil" name:"UsedGroupNum"`

	// Maximum message retention period in milliseconds
	MaxRetentionTime *uint64 `json:"MaxRetentionTime,omitnil" name:"MaxRetentionTime"`

	// Maximum message delay in milliseconds
	MaxLatencyTime *uint64 `json:"MaxLatencyTime,omitnil" name:"MaxLatencyTime"`

	// The maximum number of queues in a single topic
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxQueuesPerTopic *uint64 `json:"MaxQueuesPerTopic,omitnil" name:"MaxQueuesPerTopic"`
}

type RocketMQClusterDetail struct {
	// Basic cluster information.
	Info *RocketMQClusterInfo `json:"Info,omitnil" name:"Info"`

	// Cluster configuration information.
	Config *RocketMQClusterConfig `json:"Config,omitnil" name:"Config"`

	// Cluster status. 0: Creating; 1: Normal; 2: Terminating; 3: Deleted; 4. Isolated; 5. Creation failed; 6: Deletion failed.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil" name:"Status"`
}

type RocketMQClusterInfo struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Cluster name
	ClusterName *string `json:"ClusterName,omitnil" name:"ClusterName"`

	// Region information
	Region *string `json:"Region,omitnil" name:"Region"`

	// Creation time in milliseconds
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// Cluster remarks
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Public network access address
	PublicEndPoint *string `json:"PublicEndPoint,omitnil" name:"PublicEndPoint"`

	// VPC access address
	VpcEndPoint *string `json:"VpcEndPoint,omitnil" name:"VpcEndPoint"`

	// Whether the namespace access point is supported.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	SupportNamespaceEndpoint *bool `json:"SupportNamespaceEndpoint,omitnil" name:"SupportNamespaceEndpoint"`

	// VPC Information
	// Note: This field may return null, indicating that no valid values can be obtained.
	Vpcs []*VpcConfig `json:"Vpcs,omitnil" name:"Vpcs"`

	// Whether it is an exclusive instance
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsVip *bool `json:"IsVip,omitnil" name:"IsVip"`

	// TDMQ for RocketMQ cluster type flag
	// Note: This field may return null, indicating that no valid values can be obtained.
	RocketMQFlag *bool `json:"RocketMQFlag,omitnil" name:"RocketMQFlag"`

	// Billing status (`1`: Normal; `2`: Service suspended; `3`: Terminated)
	// Note: This field may return null, indicating that no valid values can be obtained.
	Status *int64 `json:"Status,omitnil" name:"Status"`

	// Service suspension time in milliseconds
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsolateTime *int64 `json:"IsolateTime,omitnil" name:"IsolateTime"`

	// HTTP-based public network access address
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpPublicEndpoint *string `json:"HttpPublicEndpoint,omitnil" name:"HttpPublicEndpoint"`

	// HTTP-based VPC access address
	// Note: This field may return null, indicating that no valid values can be obtained.
	HttpVpcEndpoint *string `json:"HttpVpcEndpoint,omitnil" name:"HttpVpcEndpoint"`

	// Internal TCP access address
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	InternalEndpoint *string `json:"InternalEndpoint,omitnil" name:"InternalEndpoint"`

	// Internal HTTP access address
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	HttpInternalEndpoint *string `json:"HttpInternalEndpoint,omitnil" name:"HttpInternalEndpoint"`
}

type RocketMQClusterRecentStats struct {
	// Number of topics
	TopicNum *uint64 `json:"TopicNum,omitnil" name:"TopicNum"`

	// Number of produced messages
	ProducedMsgNum *uint64 `json:"ProducedMsgNum,omitnil" name:"ProducedMsgNum"`

	// Number of consumed messages
	ConsumedMsgNum *uint64 `json:"ConsumedMsgNum,omitnil" name:"ConsumedMsgNum"`

	// Number of retained messages
	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitnil" name:"AccumulativeMsgNum"`
}

type RocketMQGroup struct {
	// Consumer group name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// The number of online consumers.
	ConsumerNum *uint64 `json:"ConsumerNum,omitnil" name:"ConsumerNum"`

	// Consumption TPS.
	TPS *uint64 `json:"TPS,omitnil" name:"TPS"`

	// The total number of heaped messages.
	TotalAccumulative *int64 `json:"TotalAccumulative,omitnil" name:"TotalAccumulative"`

	// 0: Cluster consumption mode; 1: Broadcast consumption mode; -1: Unknown.
	ConsumptionMode *int64 `json:"ConsumptionMode,omitnil" name:"ConsumptionMode"`

	// Whether to allow consumption.
	ReadEnabled *bool `json:"ReadEnabled,omitnil" name:"ReadEnabled"`

	// The number of partitions in a retry topic.
	// Note: This field may return null, indicating that no valid values can be obtained.
	RetryPartitionNum *uint64 `json:"RetryPartitionNum,omitnil" name:"RetryPartitionNum"`

	// Creation time in milliseconds.
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// Modification time in milliseconds.
	UpdateTime *uint64 `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// Client protocol.
	ClientProtocol *string `json:"ClientProtocol,omitnil" name:"ClientProtocol"`

	// Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Consumer type. Enumerated values: `ACTIVELY` or `PASSIVELY`.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerType *string `json:"ConsumerType,omitnil" name:"ConsumerType"`

	// Whether to enable broadcast consumption.
	BroadcastEnabled *bool `json:"BroadcastEnabled,omitnil" name:"BroadcastEnabled"`

	// Group type
	// Note: This field may return null, indicating that no valid values can be obtained.
	GroupType *string `json:"GroupType,omitnil" name:"GroupType"`

	// The number of retries
	// Note: This field may return null, indicating that no valid values can be obtained.
	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitnil" name:"RetryMaxTimes"`
}

type RocketMQInstanceConfig struct {
	// Maximum TPS per namespace
	MaxTpsPerNamespace *uint64 `json:"MaxTpsPerNamespace,omitnil" name:"MaxTpsPerNamespace"`

	// Maximum number of namespaces
	MaxNamespaceNum *uint64 `json:"MaxNamespaceNum,omitnil" name:"MaxNamespaceNum"`

	// Number of used namespaces
	UsedNamespaceNum *uint64 `json:"UsedNamespaceNum,omitnil" name:"UsedNamespaceNum"`

	// Maximum number of topics
	MaxTopicNum *uint64 `json:"MaxTopicNum,omitnil" name:"MaxTopicNum"`

	// Number of used topics
	UsedTopicNum *uint64 `json:"UsedTopicNum,omitnil" name:"UsedTopicNum"`

	// Maximum number of groups
	MaxGroupNum *uint64 `json:"MaxGroupNum,omitnil" name:"MaxGroupNum"`

	// Number of used groups
	UsedGroupNum *uint64 `json:"UsedGroupNum,omitnil" name:"UsedGroupNum"`

	// Cluster type
	ConfigDisplay *string `json:"ConfigDisplay,omitnil" name:"ConfigDisplay"`

	// Number of nodes in the cluster
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Node distribution
	NodeDistribution []*InstanceNodeDistribution `json:"NodeDistribution,omitnil" name:"NodeDistribution"`

	// Topic distribution
	TopicDistribution []*RocketMQTopicDistribution `json:"TopicDistribution,omitnil" name:"TopicDistribution"`


	MaxQueuesPerTopic *uint64 `json:"MaxQueuesPerTopic,omitnil" name:"MaxQueuesPerTopic"`
}

type RocketMQMessageTrack struct {
	// Consumer group
	Group *string `json:"Group,omitnil" name:"Group"`

	// Consumption status
	ConsumeStatus *string `json:"ConsumeStatus,omitnil" name:"ConsumeStatus"`

	// Message trace type
	TrackType *string `json:"TrackType,omitnil" name:"TrackType"`

	// Exception information
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	ExceptionDesc *string `json:"ExceptionDesc,omitnil" name:"ExceptionDesc"`
}

type RocketMQNamespace struct {
	// Namespace name, which can contain 3–64 letters, digits, hyphens, and underscores.
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Retention period for unconsumed messages in milliseconds. Valid range: 60 seconds–15 days. This parameter is disused.
	Ttl *uint64 `json:"Ttl,omitnil" name:"Ttl"`

	// Retention period for persistently stored messages in milliseconds.
	RetentionTime *uint64 `json:"RetentionTime,omitnil" name:"RetentionTime"`

	// Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Public network access point address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	PublicEndpoint *string `json:"PublicEndpoint,omitnil" name:"PublicEndpoint"`

	// VPC access point address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	VpcEndpoint *string `json:"VpcEndpoint,omitnil" name:"VpcEndpoint"`

	// Internal access point address
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	InternalEndpoint *string `json:"InternalEndpoint,omitnil" name:"InternalEndpoint"`
}

type RocketMQTopic struct {
	// Topic name.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Topic type. Enumerated values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`, `Retry`, and `DeadLetter`.
	Type *string `json:"Type,omitnil" name:"Type"`

	// The number of subscription groups
	GroupNum *uint64 `json:"GroupNum,omitnil" name:"GroupNum"`

	// Description.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// The number of read/write partitions.
	PartitionNum *uint64 `json:"PartitionNum,omitnil" name:"PartitionNum"`

	// Creation time in milliseconds.
	CreateTime *uint64 `json:"CreateTime,omitnil" name:"CreateTime"`

	// Creation time in milliseconds.
	UpdateTime *uint64 `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

type RocketMQTopicDistribution struct {
	// Topic type
	TopicType *string `json:"TopicType,omitnil" name:"TopicType"`

	// Number of topics
	Count *uint64 `json:"Count,omitnil" name:"Count"`
}

type RocketMQVipInstance struct {
	// Instance ID
	InstanceId *string `json:"InstanceId,omitnil" name:"InstanceId"`

	// Instance name
	InstanceName *string `json:"InstanceName,omitnil" name:"InstanceName"`

	// Instance version
	// Note: This field may return null, indicating that no valid values can be obtained.
	InstanceVersion *string `json:"InstanceVersion,omitnil" name:"InstanceVersion"`

	// Instance status. Valid values: `0` (Creating), `1` (Normal), `2` (Isolated), `3` (Terminated), `4` (Abnormal), `5` (Delivery failed).
	Status *uint64 `json:"Status,omitnil" name:"Status"`

	// Number of nodes
	NodeCount *uint64 `json:"NodeCount,omitnil" name:"NodeCount"`

	// Instance specification name
	ConfigDisplay *string `json:"ConfigDisplay,omitnil" name:"ConfigDisplay"`

	// Peak TPS
	MaxTps *uint64 `json:"MaxTps,omitnil" name:"MaxTps"`

	// Peak bandwidth in Mbps
	MaxBandWidth *uint64 `json:"MaxBandWidth,omitnil" name:"MaxBandWidth"`

	// Storage capacity in GB
	MaxStorage *uint64 `json:"MaxStorage,omitnil" name:"MaxStorage"`

	// Instance expiration time in milliseconds
	ExpireTime *uint64 `json:"ExpireTime,omitnil" name:"ExpireTime"`

	// Renewal mode. Valid values: `0` (Manual renewal, which is the default mode), `1` (Auto-renewal), `2` (Manual renewal, which is specified by users).
	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitnil" name:"AutoRenewFlag"`

	// Payment mode. 0: Postpaid; 1: Prepaid.
	PayMode *uint64 `json:"PayMode,omitnil" name:"PayMode"`

	// Remarks
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Instance specification ID
	SpecName *string `json:"SpecName,omitnil" name:"SpecName"`

	// The maximum message retention period in hours
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	MaxRetention *int64 `json:"MaxRetention,omitnil" name:"MaxRetention"`

	// The minimum message retention period in hours
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	MinRetention *int64 `json:"MinRetention,omitnil" name:"MinRetention"`

	// Instance message retention period in hours
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	Retention *int64 `json:"Retention,omitnil" name:"Retention"`
}

type Role struct {
	// Role name.
	RoleName *string `json:"RoleName,omitnil" name:"RoleName"`

	// Value of the role token.
	Token *string `json:"Token,omitnil" name:"Token"`

	// Remarks.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Creation time.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Update time.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`
}

// Predefined struct for user
type SendBatchMessagesRequestParams struct {
	// Name of the topic to which to send the message. It is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Content of the message to be sent
	Payload *string `json:"Payload,omitnil" name:"Payload"`

	// String-Type token, which is optional and will be automatically obtained by the system.
	StringToken *string `json:"StringToken,omitnil" name:"StringToken"`

	// Producer name, which must be globally unique. If it is not configured, the system will automatically generate one.
	ProducerName *string `json:"ProducerName,omitnil" name:"ProducerName"`

	// Message sending timeout period in seconds. Default value: 30s
	SendTimeout *int64 `json:"SendTimeout,omitnil" name:"SendTimeout"`

	// Maximum number of produced messages which can be cached in the memory. Default value: 1000
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitnil" name:"MaxPendingMessages"`

	// Maximum number of messages in each batch. Default value: 1000 messages/batch
	BatchingMaxMessages *int64 `json:"BatchingMaxMessages,omitnil" name:"BatchingMaxMessages"`

	// Maximum wait time for each batch, after which the batch will be sent no matter whether the specified number or size of messages in the batch is reached. Default value: 10 ms
	BatchingMaxPublishDelay *int64 `json:"BatchingMaxPublishDelay,omitnil" name:"BatchingMaxPublishDelay"`

	// Maximum allowed size of messages in each batch. Default value: 128 KB
	BatchingMaxBytes *int64 `json:"BatchingMaxBytes,omitnil" name:"BatchingMaxBytes"`
}

type SendBatchMessagesRequest struct {
	*tchttp.BaseRequest
	
	// Name of the topic to which to send the message. It is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Content of the message to be sent
	Payload *string `json:"Payload,omitnil" name:"Payload"`

	// String-Type token, which is optional and will be automatically obtained by the system.
	StringToken *string `json:"StringToken,omitnil" name:"StringToken"`

	// Producer name, which must be globally unique. If it is not configured, the system will automatically generate one.
	ProducerName *string `json:"ProducerName,omitnil" name:"ProducerName"`

	// Message sending timeout period in seconds. Default value: 30s
	SendTimeout *int64 `json:"SendTimeout,omitnil" name:"SendTimeout"`

	// Maximum number of produced messages which can be cached in the memory. Default value: 1000
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitnil" name:"MaxPendingMessages"`

	// Maximum number of messages in each batch. Default value: 1000 messages/batch
	BatchingMaxMessages *int64 `json:"BatchingMaxMessages,omitnil" name:"BatchingMaxMessages"`

	// Maximum wait time for each batch, after which the batch will be sent no matter whether the specified number or size of messages in the batch is reached. Default value: 10 ms
	BatchingMaxPublishDelay *int64 `json:"BatchingMaxPublishDelay,omitnil" name:"BatchingMaxPublishDelay"`

	// Maximum allowed size of messages in each batch. Default value: 128 KB
	BatchingMaxBytes *int64 `json:"BatchingMaxBytes,omitnil" name:"BatchingMaxBytes"`
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

// Predefined struct for user
type SendBatchMessagesResponseParams struct {
	// Unique message ID
	// Note: this field may return null, indicating that no valid values can be obtained.
	MessageId *string `json:"MessageId,omitnil" name:"MessageId"`

	// Error message. If an empty string is returned, no error occurred.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SendBatchMessagesResponse struct {
	*tchttp.BaseResponse
	Response *SendBatchMessagesResponseParams `json:"Response"`
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

// Predefined struct for user
type SendCmqMsgRequestParams struct {
	// Queue name
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// Message content
	MsgContent *string `json:"MsgContent,omitnil" name:"MsgContent"`

	// Delay time
	DelaySeconds *int64 `json:"DelaySeconds,omitnil" name:"DelaySeconds"`
}

type SendCmqMsgRequest struct {
	*tchttp.BaseRequest
	
	// Queue name
	QueueName *string `json:"QueueName,omitnil" name:"QueueName"`

	// Message content
	MsgContent *string `json:"MsgContent,omitnil" name:"MsgContent"`

	// Delay time
	DelaySeconds *int64 `json:"DelaySeconds,omitnil" name:"DelaySeconds"`
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

// Predefined struct for user
type SendCmqMsgResponseParams struct {
	// `true` indicates that the sending is successful
	Result *bool `json:"Result,omitnil" name:"Result"`

	// Message ID
	MsgId *string `json:"MsgId,omitnil" name:"MsgId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SendCmqMsgResponse struct {
	*tchttp.BaseResponse
	Response *SendCmqMsgResponseParams `json:"Response"`
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

// Predefined struct for user
type SendMessagesRequestParams struct {
	// Name of the topic to which to send the message. It is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Content of the message to be sent
	Payload *string `json:"Payload,omitnil" name:"Payload"`

	// Token used for authentication, which is optional and will be automatically obtained by the system.
	StringToken *string `json:"StringToken,omitnil" name:"StringToken"`

	// Producer name, which is randomly generated and must be globally unique. If you set the producer name manually, the producer may fail to be created, causing message sending failure.
	// This parameter is used only when a specific producer is allowed to produce messages. It won’t be used in most cases.
	ProducerName *string `json:"ProducerName,omitnil" name:"ProducerName"`

	// Message sending timeout period, which is 30s by default.
	SendTimeout *int64 `json:"SendTimeout,omitnil" name:"SendTimeout"`

	// Maximum number of produced messages which can be cached in the memory. Default value: 1000
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitnil" name:"MaxPendingMessages"`
}

type SendMessagesRequest struct {
	*tchttp.BaseRequest
	
	// Name of the topic to which to send the message. It is better to be the full path of the topic, such as `tenant/namespace/topic`. If it is not specified, `public/default` will be used by default.
	Topic *string `json:"Topic,omitnil" name:"Topic"`

	// Content of the message to be sent
	Payload *string `json:"Payload,omitnil" name:"Payload"`

	// Token used for authentication, which is optional and will be automatically obtained by the system.
	StringToken *string `json:"StringToken,omitnil" name:"StringToken"`

	// Producer name, which is randomly generated and must be globally unique. If you set the producer name manually, the producer may fail to be created, causing message sending failure.
	// This parameter is used only when a specific producer is allowed to produce messages. It won’t be used in most cases.
	ProducerName *string `json:"ProducerName,omitnil" name:"ProducerName"`

	// Message sending timeout period, which is 30s by default.
	SendTimeout *int64 `json:"SendTimeout,omitnil" name:"SendTimeout"`

	// Maximum number of produced messages which can be cached in the memory. Default value: 1000
	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitnil" name:"MaxPendingMessages"`
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

// Predefined struct for user
type SendMessagesResponseParams struct {
	// messageID, which must be globally unique and is the metadata information used to identify the message.
	// Note: this field may return null, indicating that no valid values can be obtained.
	MessageId *string `json:"MessageId,omitnil" name:"MessageId"`

	// Returned error message. If an empty string is returned, no error occurred.
	// Note: this field may return null, indicating that no valid values can be obtained.
	ErrorMsg *string `json:"ErrorMsg,omitnil" name:"ErrorMsg"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SendMessagesResponse struct {
	*tchttp.BaseResponse
	Response *SendMessagesResponseParams `json:"Response"`
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

// Predefined struct for user
type SendMsgRequestParams struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name. If the topic is a partitioned topic, you need to specify the partition; otherwise, messages will be sent to partition 0 by default, such as `my_topic-partition-0`.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Message content, which cannot be empty and can contain up to 5,242,880 bytes.
	MsgContent *string `json:"MsgContent,omitnil" name:"MsgContent"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
}

type SendMsgRequest struct {
	*tchttp.BaseRequest
	
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name. If the topic is a partitioned topic, you need to specify the partition; otherwise, messages will be sent to partition 0 by default, such as `my_topic-partition-0`.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Message content, which cannot be empty and can contain up to 5,242,880 bytes.
	MsgContent *string `json:"MsgContent,omitnil" name:"MsgContent"`

	// Pulsar cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`
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

// Predefined struct for user
type SendMsgResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SendMsgResponse struct {
	*tchttp.BaseResponse
	Response *SendMsgResponseParams `json:"Response"`
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

// Predefined struct for user
type SendRocketMQMessageRequestParams struct {
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace ID
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Message content
	MsgBody *string `json:"MsgBody,omitnil" name:"MsgBody"`

	// Message key
	MsgKey *string `json:"MsgKey,omitnil" name:"MsgKey"`

	// Message tag
	MsgTag *string `json:"MsgTag,omitnil" name:"MsgTag"`
}

type SendRocketMQMessageRequest struct {
	*tchttp.BaseRequest
	
	// Cluster ID
	ClusterId *string `json:"ClusterId,omitnil" name:"ClusterId"`

	// Namespace ID
	NamespaceId *string `json:"NamespaceId,omitnil" name:"NamespaceId"`

	// Topic name
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Message content
	MsgBody *string `json:"MsgBody,omitnil" name:"MsgBody"`

	// Message key
	MsgKey *string `json:"MsgKey,omitnil" name:"MsgKey"`

	// Message tag
	MsgTag *string `json:"MsgTag,omitnil" name:"MsgTag"`
}

func (r *SendRocketMQMessageRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendRocketMQMessageRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ClusterId")
	delete(f, "NamespaceId")
	delete(f, "TopicName")
	delete(f, "MsgBody")
	delete(f, "MsgKey")
	delete(f, "MsgTag")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "SendRocketMQMessageRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type SendRocketMQMessageResponseParams struct {
	// Message sending result
	Result *bool `json:"Result,omitnil" name:"Result"`

	// Message ID
	// Note: u200dThis field may return null, indicating that no valid values can be obtained.
	MsgId *string `json:"MsgId,omitnil" name:"MsgId"`

	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type SendRocketMQMessageResponse struct {
	*tchttp.BaseResponse
	Response *SendRocketMQMessageResponseParams `json:"Response"`
}

func (r *SendRocketMQMessageResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *SendRocketMQMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Sort struct {
	// Sorting field.
	Name *string `json:"Name,omitnil" name:"Name"`

	// Ascending order: `ASC`; descending order: `DESC`.
	Order *string `json:"Order,omitnil" name:"Order"`
}

type Subscription struct {
	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// The time when the consumer started connecting.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	ConnectedSince *string `json:"ConnectedSince,omitnil" name:"ConnectedSince"`

	// Consumer address.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerAddr *string `json:"ConsumerAddr,omitnil" name:"ConsumerAddr"`

	// The number of consumers.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerCount *string `json:"ConsumerCount,omitnil" name:"ConsumerCount"`

	// Consumer name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerName *string `json:"ConsumerName,omitnil" name:"ConsumerName"`

	// The number of heaped messages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgBacklog *string `json:"MsgBacklog,omitnil" name:"MsgBacklog"`

	// Percentage of messages under this subscription that were discarded but not sent after TTL.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgRateExpired *string `json:"MsgRateExpired,omitnil" name:"MsgRateExpired"`

	// The total number of messages delivered by the consumer per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgRateOut *string `json:"MsgRateOut,omitnil" name:"MsgRateOut"`

	// The size (in bytes) of messages consumed by the consumer per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgThroughputOut *string `json:"MsgThroughputOut,omitnil" name:"MsgThroughputOut"`

	// Subscription name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`

	// Set of consumers.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerSets []*Consumer `json:"ConsumerSets,omitnil" name:"ConsumerSets"`

	// Whether the consumer is online.
	// Note: This field may return null, indicating that no valid values can be obtained.
	IsOnline *bool `json:"IsOnline,omitnil" name:"IsOnline"`

	// Set of consumption progress information.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumersScheduleSets []*ConsumersSchedule `json:"ConsumersScheduleSets,omitnil" name:"ConsumersScheduleSets"`

	// Remarks.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Last modified.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// Subscription type. Valid values: `Exclusive`, `Shared`, `Failover`, and `Key_Shared`. An empty string or `NULL`: Unknown.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubType *string `json:"SubType,omitnil" name:"SubType"`

	// Whether messages are blocked as the limit of unacknowledged messages has been reached.
	// Note: This field may return null, indicating that no valid values can be obtained.
	BlockedSubscriptionOnUnackedMsgs *bool `json:"BlockedSubscriptionOnUnackedMsgs,omitnil" name:"BlockedSubscriptionOnUnackedMsgs"`

	// The maximum number of unacknowledged messages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MaxUnackedMsgNum *int64 `json:"MaxUnackedMsgNum,omitnil" name:"MaxUnackedMsgNum"`
}

type SubscriptionTopic struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Subscription name.
	SubscriptionName *string `json:"SubscriptionName,omitnil" name:"SubscriptionName"`
}

type Tag struct {
	// Value of the tag key
	TagKey *string `json:"TagKey,omitnil" name:"TagKey"`

	// Value of the tag value
	TagValue *string `json:"TagValue,omitnil" name:"TagValue"`
}

type Topic struct {
	// Average size of the messages published in the last interval in bytes.
	// Note: This field may return `null`, indicating that no valid values can be obtained.
	AverageMsgSize *string `json:"AverageMsgSize,omitnil" name:"AverageMsgSize"`

	// The number of consumers.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerCount *string `json:"ConsumerCount,omitnil" name:"ConsumerCount"`

	// The total number of recorded messages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitnil" name:"LastConfirmedEntry"`

	// Time when the last ledger was created.
	// Note: This field may return null, indicating that no valid values can be obtained.
	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitnil" name:"LastLedgerCreatedTimestamp"`

	// The number of messages published by local and replicated publishers per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgRateIn *string `json:"MsgRateIn,omitnil" name:"MsgRateIn"`

	// The total number of messages delivered by local and replicated consumers per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgRateOut *string `json:"MsgRateOut,omitnil" name:"MsgRateOut"`

	// The size (in bytes) of messages published by local and replicated publishers per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgThroughputIn *string `json:"MsgThroughputIn,omitnil" name:"MsgThroughputIn"`

	// The size (in bytes) of messages delivered by local and replicated consumers per second.
	// Note: This field may return null, indicating that no valid values can be obtained.
	MsgThroughputOut *string `json:"MsgThroughputOut,omitnil" name:"MsgThroughputOut"`

	// The total number of recorded messages.
	// Note: This field may return null, indicating that no valid values can be obtained.
	NumberOfEntries *string `json:"NumberOfEntries,omitnil" name:"NumberOfEntries"`

	// Partition count ≤ 0: there are no subpartitions in the topic.
	// Note: This field may return null, indicating that no valid values can be obtained.
	Partitions *int64 `json:"Partitions,omitnil" name:"Partitions"`

	// The number of producers.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProducerCount *string `json:"ProducerCount,omitnil" name:"ProducerCount"`

	// The size of all stored messages in bytes.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TotalSize *string `json:"TotalSize,omitnil" name:"TotalSize"`

	// Subpartitions in a partitioned topic.
	// Note: This field may return null, indicating that no valid values can be obtained.
	SubTopicSets []*PartitionsTopic `json:"SubTopicSets,omitnil" name:"SubTopicSets"`

	// Topic type description:
	// 0: General message;
	// 1: Globally sequential message;
	// 2: Partitionally sequential message;
	// 3: Retry letter topic;
	// 4: Dead letter topic;
	// 5: Transaction message.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicType *uint64 `json:"TopicType,omitnil" name:"TopicType"`

	// Environment (namespace) name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	// Note: This field may return null, indicating that no valid values can be obtained.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`

	// Remarks (up to 128 characters).
	// Note: This field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`

	// Creation time.
	// Note: This field may return null, indicating that no valid values can be obtained.
	CreateTime *string `json:"CreateTime,omitnil" name:"CreateTime"`

	// Last modified.
	// Note: This field may return null, indicating that no valid values can be obtained.
	UpdateTime *string `json:"UpdateTime,omitnil" name:"UpdateTime"`

	// The maximum number of producers.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ProducerLimit *string `json:"ProducerLimit,omitnil" name:"ProducerLimit"`

	// The maximum number of consumers.
	// Note: This field may return null, indicating that no valid values can be obtained.
	ConsumerLimit *string `json:"ConsumerLimit,omitnil" name:"ConsumerLimit"`

	// `0`: Non-persistent and non-partitioned
	// `1`: Non-persistent and partitioned
	// `2`: Persistent and non-partitioned
	// `3`: Persistent and partitioned
	// Note: This field may return null, indicating that no valid values can be obtained.
	PulsarTopicType *int64 `json:"PulsarTopicType,omitnil" name:"PulsarTopicType"`
}

type TopicRecord struct {
	// Environment (namespace) name.
	EnvironmentId *string `json:"EnvironmentId,omitnil" name:"EnvironmentId"`

	// Topic name.
	TopicName *string `json:"TopicName,omitnil" name:"TopicName"`
}

// Predefined struct for user
type UnbindCmqDeadLetterRequestParams struct {
	// Source queue name of dead letter policy. Calling this API will clear the dead letter queue policy of this queue.
	SourceQueueName *string `json:"SourceQueueName,omitnil" name:"SourceQueueName"`
}

type UnbindCmqDeadLetterRequest struct {
	*tchttp.BaseRequest
	
	// Source queue name of dead letter policy. Calling this API will clear the dead letter queue policy of this queue.
	SourceQueueName *string `json:"SourceQueueName,omitnil" name:"SourceQueueName"`
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

// Predefined struct for user
type UnbindCmqDeadLetterResponseParams struct {
	// The unique request ID, which is returned for each request. RequestId is required for locating a problem.
	RequestId *string `json:"RequestId,omitnil" name:"RequestId"`
}

type UnbindCmqDeadLetterResponse struct {
	*tchttp.BaseResponse
	Response *UnbindCmqDeadLetterResponseParams `json:"Response"`
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
	UniqueVpcId *string `json:"UniqueVpcId,omitnil" name:"UniqueVpcId"`

	// Tenant VPC subnet ID
	UniqueSubnetId *string `json:"UniqueSubnetId,omitnil" name:"UniqueSubnetId"`

	// Route ID
	RouterId *string `json:"RouterId,omitnil" name:"RouterId"`

	// VPC ID
	Ip *string `json:"Ip,omitnil" name:"Ip"`

	// VPC port
	Port *uint64 `json:"Port,omitnil" name:"Port"`

	// Remarks (up to 128 characters)
	// Note: this field may return null, indicating that no valid values can be obtained.
	Remark *string `json:"Remark,omitnil" name:"Remark"`
}

type VpcConfig struct {
	// VPC ID
	VpcId *string `json:"VpcId,omitnil" name:"VpcId"`

	// Subnet ID
	SubnetId *string `json:"SubnetId,omitnil" name:"SubnetId"`
}