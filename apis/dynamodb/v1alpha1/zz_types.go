/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ArchivalSummary struct {
	ArchivalBackupARN *string `json:"archivalBackupARN,omitempty"`

	ArchivalDateTime *metav1.Time `json:"archivalDateTime,omitempty"`

	ArchivalReason *string `json:"archivalReason,omitempty"`
}

type AttributeDefinition struct {
	AttributeName *string `json:"attributeName,omitempty"`

	AttributeType *string `json:"attributeType,omitempty"`
}

type AutoScalingSettingsDescription struct {
	AutoScalingRoleARN *string `json:"autoScalingRoleARN,omitempty"`

	MaximumUnits *int64 `json:"maximumUnits,omitempty"`

	MinimumUnits *int64 `json:"minimumUnits,omitempty"`
}

type AutoScalingSettingsUpdate struct {
	MaximumUnits *int64 `json:"maximumUnits,omitempty"`

	MinimumUnits *int64 `json:"minimumUnits,omitempty"`
}

type BackupDescription struct {
	// Contains the details of the backup created for the table.
	BackupDetails *BackupDetails `json:"backupDetails,omitempty"`
	// Contains the details of the table when the backup was created.
	SourceTableDetails *SourceTableDetails `json:"sourceTableDetails,omitempty"`
	// Contains the details of the features enabled on the table when the backup
	// was created. For example, LSIs, GSIs, streams, TTL.
	SourceTableFeatureDetails *SourceTableFeatureDetails `json:"sourceTableFeatureDetails,omitempty"`
}

type BackupDetails struct {
	BackupARN *string `json:"backupARN,omitempty"`

	BackupCreationDateTime *metav1.Time `json:"backupCreationDateTime,omitempty"`

	BackupExpiryDateTime *metav1.Time `json:"backupExpiryDateTime,omitempty"`

	BackupName *string `json:"backupName,omitempty"`

	BackupSizeBytes *int64 `json:"backupSizeBytes,omitempty"`

	BackupStatus *string `json:"backupStatus,omitempty"`

	BackupType *string `json:"backupType,omitempty"`
}

type BackupSummary struct {
	BackupARN *string `json:"backupARN,omitempty"`

	BackupCreationDateTime *metav1.Time `json:"backupCreationDateTime,omitempty"`

	BackupExpiryDateTime *metav1.Time `json:"backupExpiryDateTime,omitempty"`

	BackupName *string `json:"backupName,omitempty"`

	BackupSizeBytes *int64 `json:"backupSizeBytes,omitempty"`

	BackupStatus *string `json:"backupStatus,omitempty"`

	BackupType *string `json:"backupType,omitempty"`

	TableARN *string `json:"tableARN,omitempty"`

	TableID *string `json:"tableID,omitempty"`

	TableName *string `json:"tableName,omitempty"`
}

type BatchStatementError struct {
	Message *string `json:"message,omitempty"`
}

type BatchStatementResponse struct {
	TableName *string `json:"tableName,omitempty"`
}

type BillingModeSummary struct {
	BillingMode *string `json:"billingMode,omitempty"`

	LastUpdateToPayPerRequestDateTime *metav1.Time `json:"lastUpdateToPayPerRequestDateTime,omitempty"`
}

type ConditionCheck struct {
	TableName *string `json:"tableName,omitempty"`
}

type ConsumedCapacity struct {
	TableName *string `json:"tableName,omitempty"`
}

type ContributorInsightsSummary struct {
	IndexName *string `json:"indexName,omitempty"`

	TableName *string `json:"tableName,omitempty"`
}

type CreateGlobalSecondaryIndexAction struct {
	IndexName *string `json:"indexName,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	// Represents attributes that are copied (projected) from the table into an
	// index. These are in addition to the primary key attributes and index key
	// attributes, which are automatically projected.
	Projection *Projection `json:"projection,omitempty"`
	// Represents the provisioned throughput settings for a specified table or index.
	// The settings can be modified using the UpdateTable operation.
	//
	// For current minimum and maximum provisioned throughput values, see Service,
	// Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// in the Amazon DynamoDB Developer Guide.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`
}

type CreateReplicaAction struct {
	RegionName *string `json:"regionName,omitempty"`
}

type CreateReplicationGroupMemberAction struct {
	GlobalSecondaryIndexes []*ReplicaGlobalSecondaryIndex `json:"globalSecondaryIndexes,omitempty"`

	KMSMasterKeyID *string `json:"kmsMasterKeyID,omitempty"`
	// Replica-specific provisioned throughput settings. If not specified, uses
	// the source table's provisioned throughput settings.
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `json:"provisionedThroughputOverride,omitempty"`

	RegionName *string `json:"regionName,omitempty"`
}

type Delete struct {
	TableName *string `json:"tableName,omitempty"`
}

type DeleteGlobalSecondaryIndexAction struct {
	IndexName *string `json:"indexName,omitempty"`
}

type DeleteReplicaAction struct {
	RegionName *string `json:"regionName,omitempty"`
}

type DeleteReplicationGroupMemberAction struct {
	RegionName *string `json:"regionName,omitempty"`
}

type Endpoint struct {
	Address *string `json:"address,omitempty"`

	CachePeriodInMinutes *int64 `json:"cachePeriodInMinutes,omitempty"`
}

type ExportDescription struct {
	ItemCount *int64 `json:"itemCount,omitempty"`

	TableARN *string `json:"tableARN,omitempty"`

	TableID *string `json:"tableID,omitempty"`
}

type Get struct {
	TableName *string `json:"tableName,omitempty"`
}

type GlobalSecondaryIndex struct {
	IndexName *string `json:"indexName,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	// Represents attributes that are copied (projected) from the table into an
	// index. These are in addition to the primary key attributes and index key
	// attributes, which are automatically projected.
	Projection *Projection `json:"projection,omitempty"`
	// Represents the provisioned throughput settings for a specified table or index.
	// The settings can be modified using the UpdateTable operation.
	//
	// For current minimum and maximum provisioned throughput values, see Service,
	// Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// in the Amazon DynamoDB Developer Guide.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`
}

type GlobalSecondaryIndexAutoScalingUpdate struct {
	IndexName *string `json:"indexName,omitempty"`
}

type GlobalSecondaryIndexDescription struct {
	Backfilling *bool `json:"backfilling,omitempty"`

	IndexARN *string `json:"indexARN,omitempty"`

	IndexName *string `json:"indexName,omitempty"`

	IndexSizeBytes *int64 `json:"indexSizeBytes,omitempty"`

	IndexStatus *string `json:"indexStatus,omitempty"`

	ItemCount *int64 `json:"itemCount,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	// Represents attributes that are copied (projected) from the table into an
	// index. These are in addition to the primary key attributes and index key
	// attributes, which are automatically projected.
	Projection *Projection `json:"projection,omitempty"`
	// Represents the provisioned throughput settings for the table, consisting
	// of read and write capacity units, along with data about increases and decreases.
	ProvisionedThroughput *ProvisionedThroughputDescription `json:"provisionedThroughput,omitempty"`
}

type GlobalSecondaryIndexInfo struct {
	IndexName *string `json:"indexName,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	// Represents attributes that are copied (projected) from the table into an
	// index. These are in addition to the primary key attributes and index key
	// attributes, which are automatically projected.
	Projection *Projection `json:"projection,omitempty"`
	// Represents the provisioned throughput settings for a specified table or index.
	// The settings can be modified using the UpdateTable operation.
	//
	// For current minimum and maximum provisioned throughput values, see Service,
	// Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// in the Amazon DynamoDB Developer Guide.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`
}

type GlobalSecondaryIndexUpdate struct {
	// Represents a new global secondary index to be added to an existing table.
	Create *CreateGlobalSecondaryIndexAction `json:"create,omitempty"`
	// Represents a global secondary index to be deleted from an existing table.
	Delete *DeleteGlobalSecondaryIndexAction `json:"delete,omitempty"`
	// Represents the new provisioned throughput settings to be applied to a global
	// secondary index.
	Update *UpdateGlobalSecondaryIndexAction `json:"update,omitempty"`
}

type GlobalTableDescription struct {
	CreationDateTime *metav1.Time `json:"creationDateTime,omitempty"`

	GlobalTableARN *string `json:"globalTableARN,omitempty"`

	GlobalTableName *string `json:"globalTableName,omitempty"`

	GlobalTableStatus *string `json:"globalTableStatus,omitempty"`

	ReplicationGroup []*ReplicaDescription `json:"replicationGroup,omitempty"`
}

type GlobalTableGlobalSecondaryIndexSettingsUpdate struct {
	IndexName *string `json:"indexName,omitempty"`

	ProvisionedWriteCapacityUnits *int64 `json:"provisionedWriteCapacityUnits,omitempty"`
}

type GlobalTable_SDK struct {
	GlobalTableName *string `json:"globalTableName,omitempty"`

	ReplicationGroup []*Replica `json:"replicationGroup,omitempty"`
}

type KeySchemaElement struct {
	AttributeName *string `json:"attributeName,omitempty"`

	KeyType *string `json:"keyType,omitempty"`
}

type KinesisDataStreamDestination struct {
	DestinationStatusDescription *string `json:"destinationStatusDescription,omitempty"`

	StreamARN *string `json:"streamARN,omitempty"`
}

type LocalSecondaryIndex struct {
	IndexName *string `json:"indexName,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	// Represents attributes that are copied (projected) from the table into an
	// index. These are in addition to the primary key attributes and index key
	// attributes, which are automatically projected.
	Projection *Projection `json:"projection,omitempty"`
}

type LocalSecondaryIndexDescription struct {
	IndexARN *string `json:"indexARN,omitempty"`

	IndexName *string `json:"indexName,omitempty"`

	IndexSizeBytes *int64 `json:"indexSizeBytes,omitempty"`

	ItemCount *int64 `json:"itemCount,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	// Represents attributes that are copied (projected) from the table into an
	// index. These are in addition to the primary key attributes and index key
	// attributes, which are automatically projected.
	Projection *Projection `json:"projection,omitempty"`
}

type LocalSecondaryIndexInfo struct {
	IndexName *string `json:"indexName,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	// Represents attributes that are copied (projected) from the table into an
	// index. These are in addition to the primary key attributes and index key
	// attributes, which are automatically projected.
	Projection *Projection `json:"projection,omitempty"`
}

type PointInTimeRecoveryDescription struct {
	EarliestRestorableDateTime *metav1.Time `json:"earliestRestorableDateTime,omitempty"`

	LatestRestorableDateTime *metav1.Time `json:"latestRestorableDateTime,omitempty"`
}

type Projection struct {
	NonKeyAttributes []*string `json:"nonKeyAttributes,omitempty"`

	ProjectionType *string `json:"projectionType,omitempty"`
}

type ProvisionedThroughput struct {
	ReadCapacityUnits *int64 `json:"readCapacityUnits,omitempty"`

	WriteCapacityUnits *int64 `json:"writeCapacityUnits,omitempty"`
}

type ProvisionedThroughputDescription struct {
	LastDecreaseDateTime *metav1.Time `json:"lastDecreaseDateTime,omitempty"`

	LastIncreaseDateTime *metav1.Time `json:"lastIncreaseDateTime,omitempty"`

	NumberOfDecreasesToday *int64 `json:"numberOfDecreasesToday,omitempty"`

	ReadCapacityUnits *int64 `json:"readCapacityUnits,omitempty"`

	WriteCapacityUnits *int64 `json:"writeCapacityUnits,omitempty"`
}

type ProvisionedThroughputOverride struct {
	ReadCapacityUnits *int64 `json:"readCapacityUnits,omitempty"`
}

type Put struct {
	TableName *string `json:"tableName,omitempty"`
}

type Replica struct {
	RegionName *string `json:"regionName,omitempty"`
}

type ReplicaAutoScalingDescription struct {
	RegionName *string `json:"regionName,omitempty"`

	ReplicaStatus *string `json:"replicaStatus,omitempty"`
}

type ReplicaAutoScalingUpdate struct {
	RegionName *string `json:"regionName,omitempty"`
}

type ReplicaDescription struct {
	GlobalSecondaryIndexes []*ReplicaGlobalSecondaryIndexDescription `json:"globalSecondaryIndexes,omitempty"`

	KMSMasterKeyID *string `json:"kmsMasterKeyID,omitempty"`
	// Replica-specific provisioned throughput settings. If not specified, uses
	// the source table's provisioned throughput settings.
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `json:"provisionedThroughputOverride,omitempty"`

	RegionName *string `json:"regionName,omitempty"`

	ReplicaInaccessibleDateTime *metav1.Time `json:"replicaInaccessibleDateTime,omitempty"`

	ReplicaStatus *string `json:"replicaStatus,omitempty"`

	ReplicaStatusDescription *string `json:"replicaStatusDescription,omitempty"`

	ReplicaStatusPercentProgress *string `json:"replicaStatusPercentProgress,omitempty"`
}

type ReplicaGlobalSecondaryIndex struct {
	IndexName *string `json:"indexName,omitempty"`
	// Replica-specific provisioned throughput settings. If not specified, uses
	// the source table's provisioned throughput settings.
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `json:"provisionedThroughputOverride,omitempty"`
}

type ReplicaGlobalSecondaryIndexAutoScalingDescription struct {
	IndexName *string `json:"indexName,omitempty"`

	IndexStatus *string `json:"indexStatus,omitempty"`
}

type ReplicaGlobalSecondaryIndexAutoScalingUpdate struct {
	IndexName *string `json:"indexName,omitempty"`
}

type ReplicaGlobalSecondaryIndexDescription struct {
	IndexName *string `json:"indexName,omitempty"`
	// Replica-specific provisioned throughput settings. If not specified, uses
	// the source table's provisioned throughput settings.
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `json:"provisionedThroughputOverride,omitempty"`
}

type ReplicaGlobalSecondaryIndexSettingsDescription struct {
	IndexName *string `json:"indexName,omitempty"`

	IndexStatus *string `json:"indexStatus,omitempty"`

	ProvisionedReadCapacityUnits *int64 `json:"provisionedReadCapacityUnits,omitempty"`

	ProvisionedWriteCapacityUnits *int64 `json:"provisionedWriteCapacityUnits,omitempty"`
}

type ReplicaGlobalSecondaryIndexSettingsUpdate struct {
	IndexName *string `json:"indexName,omitempty"`

	ProvisionedReadCapacityUnits *int64 `json:"provisionedReadCapacityUnits,omitempty"`
}

type ReplicaSettingsDescription struct {
	RegionName *string `json:"regionName,omitempty"`
	// Contains the details for the read/write capacity mode.
	ReplicaBillingModeSummary *BillingModeSummary `json:"replicaBillingModeSummary,omitempty"`

	ReplicaProvisionedReadCapacityUnits *int64 `json:"replicaProvisionedReadCapacityUnits,omitempty"`

	ReplicaProvisionedWriteCapacityUnits *int64 `json:"replicaProvisionedWriteCapacityUnits,omitempty"`

	ReplicaStatus *string `json:"replicaStatus,omitempty"`
}

type ReplicaSettingsUpdate struct {
	RegionName *string `json:"regionName,omitempty"`

	ReplicaProvisionedReadCapacityUnits *int64 `json:"replicaProvisionedReadCapacityUnits,omitempty"`
}

type ReplicaUpdate struct {
	// Represents a replica to be added.
	Create *CreateReplicaAction `json:"create,omitempty"`
	// Represents a replica to be removed.
	Delete *DeleteReplicaAction `json:"delete,omitempty"`
}

type ReplicationGroupUpdate struct {
	// Represents a replica to be created.
	Create *CreateReplicationGroupMemberAction `json:"create,omitempty"`
	// Represents a replica to be deleted.
	Delete *DeleteReplicationGroupMemberAction `json:"delete,omitempty"`
	// Represents a replica to be modified.
	Update *UpdateReplicationGroupMemberAction `json:"update,omitempty"`
}

type RestoreSummary struct {
	RestoreDateTime *metav1.Time `json:"restoreDateTime,omitempty"`

	RestoreInProgress *bool `json:"restoreInProgress,omitempty"`

	SourceBackupARN *string `json:"sourceBackupARN,omitempty"`

	SourceTableARN *string `json:"sourceTableARN,omitempty"`
}

type SSEDescription struct {
	InaccessibleEncryptionDateTime *metav1.Time `json:"inaccessibleEncryptionDateTime,omitempty"`

	KMSMasterKeyARN *string `json:"kmsMasterKeyARN,omitempty"`

	SSEType *string `json:"sseType,omitempty"`

	Status *string `json:"status,omitempty"`
}

type SSESpecification struct {
	Enabled *bool `json:"enabled,omitempty"`

	KMSMasterKeyID *string `json:"kmsMasterKeyID,omitempty"`

	SSEType *string `json:"sseType,omitempty"`
}

type SourceTableDetails struct {
	BillingMode *string `json:"billingMode,omitempty"`

	ItemCount *int64 `json:"itemCount,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`
	// Represents the provisioned throughput settings for a specified table or index.
	// The settings can be modified using the UpdateTable operation.
	//
	// For current minimum and maximum provisioned throughput values, see Service,
	// Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// in the Amazon DynamoDB Developer Guide.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`

	TableARN *string `json:"tableARN,omitempty"`

	TableCreationDateTime *metav1.Time `json:"tableCreationDateTime,omitempty"`

	TableID *string `json:"tableID,omitempty"`

	TableName *string `json:"tableName,omitempty"`

	TableSizeBytes *int64 `json:"tableSizeBytes,omitempty"`
}

type SourceTableFeatureDetails struct {
	GlobalSecondaryIndexes []*GlobalSecondaryIndexInfo `json:"globalSecondaryIndexes,omitempty"`

	LocalSecondaryIndexes []*LocalSecondaryIndexInfo `json:"localSecondaryIndexes,omitempty"`
	// The description of the server-side encryption status on the specified table.
	SSEDescription *SSEDescription `json:"sseDescription,omitempty"`
	// Represents the DynamoDB Streams configuration for a table in DynamoDB.
	StreamDescription *StreamSpecification `json:"streamDescription,omitempty"`
	// The description of the Time to Live (TTL) status on the specified table.
	TimeToLiveDescription *TimeToLiveDescription `json:"timeToLiveDescription,omitempty"`
}

type StreamSpecification struct {
	StreamEnabled *bool `json:"streamEnabled,omitempty"`

	StreamViewType *string `json:"streamViewType,omitempty"`
}

type TableAutoScalingDescription struct {
	TableName *string `json:"tableName,omitempty"`

	TableStatus *string `json:"tableStatus,omitempty"`
}

type TableDescription struct {
	// Contains details of a table archival operation.
	ArchivalSummary *ArchivalSummary `json:"archivalSummary,omitempty"`

	AttributeDefinitions []*AttributeDefinition `json:"attributeDefinitions,omitempty"`
	// Contains the details for the read/write capacity mode.
	BillingModeSummary *BillingModeSummary `json:"billingModeSummary,omitempty"`

	CreationDateTime *metav1.Time `json:"creationDateTime,omitempty"`

	GlobalSecondaryIndexes []*GlobalSecondaryIndexDescription `json:"globalSecondaryIndexes,omitempty"`

	GlobalTableVersion *string `json:"globalTableVersion,omitempty"`

	ItemCount *int64 `json:"itemCount,omitempty"`

	KeySchema []*KeySchemaElement `json:"keySchema,omitempty"`

	LatestStreamARN *string `json:"latestStreamARN,omitempty"`

	LatestStreamLabel *string `json:"latestStreamLabel,omitempty"`

	LocalSecondaryIndexes []*LocalSecondaryIndexDescription `json:"localSecondaryIndexes,omitempty"`
	// Represents the provisioned throughput settings for the table, consisting
	// of read and write capacity units, along with data about increases and decreases.
	ProvisionedThroughput *ProvisionedThroughputDescription `json:"provisionedThroughput,omitempty"`

	Replicas []*ReplicaDescription `json:"replicas,omitempty"`
	// Contains details for the restore.
	RestoreSummary *RestoreSummary `json:"restoreSummary,omitempty"`
	// The description of the server-side encryption status on the specified table.
	SSEDescription *SSEDescription `json:"sseDescription,omitempty"`
	// Represents the DynamoDB Streams configuration for a table in DynamoDB.
	StreamSpecification *StreamSpecification `json:"streamSpecification,omitempty"`

	TableARN *string `json:"tableARN,omitempty"`

	TableID *string `json:"tableID,omitempty"`

	TableName *string `json:"tableName,omitempty"`

	TableSizeBytes *int64 `json:"tableSizeBytes,omitempty"`

	TableStatus *string `json:"tableStatus,omitempty"`
}

type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

type TimeToLiveDescription struct {
	AttributeName *string `json:"attributeName,omitempty"`

	TimeToLiveStatus *string `json:"timeToLiveStatus,omitempty"`
}

type TimeToLiveSpecification struct {
	AttributeName *string `json:"attributeName,omitempty"`
}

type Update struct {
	TableName *string `json:"tableName,omitempty"`
}

type UpdateGlobalSecondaryIndexAction struct {
	IndexName *string `json:"indexName,omitempty"`
	// Represents the provisioned throughput settings for a specified table or index.
	// The settings can be modified using the UpdateTable operation.
	//
	// For current minimum and maximum provisioned throughput values, see Service,
	// Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
	// in the Amazon DynamoDB Developer Guide.
	ProvisionedThroughput *ProvisionedThroughput `json:"provisionedThroughput,omitempty"`
}

type UpdateReplicationGroupMemberAction struct {
	GlobalSecondaryIndexes []*ReplicaGlobalSecondaryIndex `json:"globalSecondaryIndexes,omitempty"`

	KMSMasterKeyID *string `json:"kmsMasterKeyID,omitempty"`
	// Replica-specific provisioned throughput settings. If not specified, uses
	// the source table's provisioned throughput settings.
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `json:"provisionedThroughputOverride,omitempty"`

	RegionName *string `json:"regionName,omitempty"`
}
