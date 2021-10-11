// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdjustmentType string

// Enum values for AdjustmentType
const (
	AdjustmentTypeChangeInCapacity        AdjustmentType = "ChangeInCapacity"
	AdjustmentTypePercentChangeInCapacity AdjustmentType = "PercentChangeInCapacity"
	AdjustmentTypeExactCapacity           AdjustmentType = "ExactCapacity"
)

// Values returns all known values for AdjustmentType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AdjustmentType) Values() []AdjustmentType {
	return []AdjustmentType{
		"ChangeInCapacity",
		"PercentChangeInCapacity",
		"ExactCapacity",
	}
}

type MetricAggregationType string

// Enum values for MetricAggregationType
const (
	MetricAggregationTypeAverage MetricAggregationType = "Average"
	MetricAggregationTypeMinimum MetricAggregationType = "Minimum"
	MetricAggregationTypeMaximum MetricAggregationType = "Maximum"
)

// Values returns all known values for MetricAggregationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MetricAggregationType) Values() []MetricAggregationType {
	return []MetricAggregationType{
		"Average",
		"Minimum",
		"Maximum",
	}
}

type MetricStatistic string

// Enum values for MetricStatistic
const (
	MetricStatisticAverage     MetricStatistic = "Average"
	MetricStatisticMinimum     MetricStatistic = "Minimum"
	MetricStatisticMaximum     MetricStatistic = "Maximum"
	MetricStatisticSampleCount MetricStatistic = "SampleCount"
	MetricStatisticSum         MetricStatistic = "Sum"
)

// Values returns all known values for MetricStatistic. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MetricStatistic) Values() []MetricStatistic {
	return []MetricStatistic{
		"Average",
		"Minimum",
		"Maximum",
		"SampleCount",
		"Sum",
	}
}

type MetricType string

// Enum values for MetricType
const (
	MetricTypeDynamoDBReadCapacityUtilization                         MetricType = "DynamoDBReadCapacityUtilization"
	MetricTypeDynamoDBWriteCapacityUtilization                        MetricType = "DynamoDBWriteCapacityUtilization"
	MetricTypeALBRequestCountPerTarget                                MetricType = "ALBRequestCountPerTarget"
	MetricTypeRDSReaderAverageCPUUtilization                          MetricType = "RDSReaderAverageCPUUtilization"
	MetricTypeRDSReaderAverageDatabaseConnections                     MetricType = "RDSReaderAverageDatabaseConnections"
	MetricTypeEC2SpotFleetRequestAverageCPUUtilization                MetricType = "EC2SpotFleetRequestAverageCPUUtilization"
	MetricTypeEC2SpotFleetRequestAverageNetworkIn                     MetricType = "EC2SpotFleetRequestAverageNetworkIn"
	MetricTypeEC2SpotFleetRequestAverageNetworkOut                    MetricType = "EC2SpotFleetRequestAverageNetworkOut"
	MetricTypeSageMakerVariantInvocationsPerInstance                  MetricType = "SageMakerVariantInvocationsPerInstance"
	MetricTypeECSServiceAverageCPUUtilization                         MetricType = "ECSServiceAverageCPUUtilization"
	MetricTypeECSServiceAverageMemoryUtilization                      MetricType = "ECSServiceAverageMemoryUtilization"
	MetricTypeAppStreamAverageCapacityUtilization                     MetricType = "AppStreamAverageCapacityUtilization"
	MetricTypeComprehendInferenceUtilization                          MetricType = "ComprehendInferenceUtilization"
	MetricTypeLambdaProvisionedConcurrencyUtilization                 MetricType = "LambdaProvisionedConcurrencyUtilization"
	MetricTypeCassandraReadCapacityUtilization                        MetricType = "CassandraReadCapacityUtilization"
	MetricTypeCassandraWriteCapacityUtilization                       MetricType = "CassandraWriteCapacityUtilization"
	MetricTypeKafkaBrokerStorageUtilization                           MetricType = "KafkaBrokerStorageUtilization"
	MetricTypeElastiCachePrimaryEngineCPUUtilization                  MetricType = "ElastiCachePrimaryEngineCPUUtilization"
	MetricTypeElastiCacheReplicaEngineCPUUtilization                  MetricType = "ElastiCacheReplicaEngineCPUUtilization"
	MetricTypeElastiCacheDatabaseMemoryUsageCountedForEvictPercentage MetricType = "ElastiCacheDatabaseMemoryUsageCountedForEvictPercentage"
	MetricTypeNeptuneReaderAverageCPUUtilization                      MetricType = "NeptuneReaderAverageCPUUtilization"
)

// Values returns all known values for MetricType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MetricType) Values() []MetricType {
	return []MetricType{
		"DynamoDBReadCapacityUtilization",
		"DynamoDBWriteCapacityUtilization",
		"ALBRequestCountPerTarget",
		"RDSReaderAverageCPUUtilization",
		"RDSReaderAverageDatabaseConnections",
		"EC2SpotFleetRequestAverageCPUUtilization",
		"EC2SpotFleetRequestAverageNetworkIn",
		"EC2SpotFleetRequestAverageNetworkOut",
		"SageMakerVariantInvocationsPerInstance",
		"ECSServiceAverageCPUUtilization",
		"ECSServiceAverageMemoryUtilization",
		"AppStreamAverageCapacityUtilization",
		"ComprehendInferenceUtilization",
		"LambdaProvisionedConcurrencyUtilization",
		"CassandraReadCapacityUtilization",
		"CassandraWriteCapacityUtilization",
		"KafkaBrokerStorageUtilization",
		"ElastiCachePrimaryEngineCPUUtilization",
		"ElastiCacheReplicaEngineCPUUtilization",
		"ElastiCacheDatabaseMemoryUsageCountedForEvictPercentage",
		"NeptuneReaderAverageCPUUtilization",
	}
}

type PolicyType string

// Enum values for PolicyType
const (
	PolicyTypeStepScaling           PolicyType = "StepScaling"
	PolicyTypeTargetTrackingScaling PolicyType = "TargetTrackingScaling"
)

// Values returns all known values for PolicyType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PolicyType) Values() []PolicyType {
	return []PolicyType{
		"StepScaling",
		"TargetTrackingScaling",
	}
}

type ScalableDimension string

// Enum values for ScalableDimension
const (
	ScalableDimensionECSServiceDesiredCount                           ScalableDimension = "ecs:service:DesiredCount"
	ScalableDimensionEC2SpotFleetRequestTargetCapacity                ScalableDimension = "ec2:spot-fleet-request:TargetCapacity"
	ScalableDimensionEMRInstanceGroupInstanceCount                    ScalableDimension = "elasticmapreduce:instancegroup:InstanceCount"
	ScalableDimensionAppstreamFleetDesiredCapacity                    ScalableDimension = "appstream:fleet:DesiredCapacity"
	ScalableDimensionDynamoDBTableReadCapacityUnits                   ScalableDimension = "dynamodb:table:ReadCapacityUnits"
	ScalableDimensionDynamoDBTableWriteCapacityUnits                  ScalableDimension = "dynamodb:table:WriteCapacityUnits"
	ScalableDimensionDynamoDBIndexReadCapacityUnits                   ScalableDimension = "dynamodb:index:ReadCapacityUnits"
	ScalableDimensionDynamoDBIndexWriteCapacityUnits                  ScalableDimension = "dynamodb:index:WriteCapacityUnits"
	ScalableDimensionRDSClusterReadReplicaCount                       ScalableDimension = "rds:cluster:ReadReplicaCount"
	ScalableDimensionSageMakerVariantDesiredInstanceCount             ScalableDimension = "sagemaker:variant:DesiredInstanceCount"
	ScalableDimensionCustomResourceScalableDimension                  ScalableDimension = "custom-resource:ResourceType:Property"
	ScalableDimensionComprehendDocClassifierEndpointInferenceUnits    ScalableDimension = "comprehend:document-classifier-endpoint:DesiredInferenceUnits"
	ScalableDimensionComprehendEntityRecognizerEndpointInferenceUnits ScalableDimension = "comprehend:entity-recognizer-endpoint:DesiredInferenceUnits"
	ScalableDimensionLambdaFunctionProvisionedConcurrency             ScalableDimension = "lambda:function:ProvisionedConcurrency"
	ScalableDimensionCassandraTableReadCapacityUnits                  ScalableDimension = "cassandra:table:ReadCapacityUnits"
	ScalableDimensionCassandraTableWriteCapacityUnits                 ScalableDimension = "cassandra:table:WriteCapacityUnits"
	ScalableDimensionKafkaBrokerStorageVolumeSize                     ScalableDimension = "kafka:broker-storage:VolumeSize"
	ScalableDimensionElastiCacheReplicationGroupNodeGroups            ScalableDimension = "elasticache:replication-group:NodeGroups"
	ScalableDimensionElastiCacheReplicationGroupReplicas              ScalableDimension = "elasticache:replication-group:Replicas"
	ScalableDimensionNeptuneClusterReadReplicaCount                   ScalableDimension = "neptune:cluster:ReadReplicaCount"
)

// Values returns all known values for ScalableDimension. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScalableDimension) Values() []ScalableDimension {
	return []ScalableDimension{
		"ecs:service:DesiredCount",
		"ec2:spot-fleet-request:TargetCapacity",
		"elasticmapreduce:instancegroup:InstanceCount",
		"appstream:fleet:DesiredCapacity",
		"dynamodb:table:ReadCapacityUnits",
		"dynamodb:table:WriteCapacityUnits",
		"dynamodb:index:ReadCapacityUnits",
		"dynamodb:index:WriteCapacityUnits",
		"rds:cluster:ReadReplicaCount",
		"sagemaker:variant:DesiredInstanceCount",
		"custom-resource:ResourceType:Property",
		"comprehend:document-classifier-endpoint:DesiredInferenceUnits",
		"comprehend:entity-recognizer-endpoint:DesiredInferenceUnits",
		"lambda:function:ProvisionedConcurrency",
		"cassandra:table:ReadCapacityUnits",
		"cassandra:table:WriteCapacityUnits",
		"kafka:broker-storage:VolumeSize",
		"elasticache:replication-group:NodeGroups",
		"elasticache:replication-group:Replicas",
		"neptune:cluster:ReadReplicaCount",
	}
}

type ScalingActivityStatusCode string

// Enum values for ScalingActivityStatusCode
const (
	ScalingActivityStatusCodePending     ScalingActivityStatusCode = "Pending"
	ScalingActivityStatusCodeInProgress  ScalingActivityStatusCode = "InProgress"
	ScalingActivityStatusCodeSuccessful  ScalingActivityStatusCode = "Successful"
	ScalingActivityStatusCodeOverridden  ScalingActivityStatusCode = "Overridden"
	ScalingActivityStatusCodeUnfulfilled ScalingActivityStatusCode = "Unfulfilled"
	ScalingActivityStatusCodeFailed      ScalingActivityStatusCode = "Failed"
)

// Values returns all known values for ScalingActivityStatusCode. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ScalingActivityStatusCode) Values() []ScalingActivityStatusCode {
	return []ScalingActivityStatusCode{
		"Pending",
		"InProgress",
		"Successful",
		"Overridden",
		"Unfulfilled",
		"Failed",
	}
}

type ServiceNamespace string

// Enum values for ServiceNamespace
const (
	ServiceNamespaceEcs            ServiceNamespace = "ecs"
	ServiceNamespaceEmr            ServiceNamespace = "elasticmapreduce"
	ServiceNamespaceEc2            ServiceNamespace = "ec2"
	ServiceNamespaceAppstream      ServiceNamespace = "appstream"
	ServiceNamespaceDynamodb       ServiceNamespace = "dynamodb"
	ServiceNamespaceRds            ServiceNamespace = "rds"
	ServiceNamespaceSagemaker      ServiceNamespace = "sagemaker"
	ServiceNamespaceCustomResource ServiceNamespace = "custom-resource"
	ServiceNamespaceComprehend     ServiceNamespace = "comprehend"
	ServiceNamespaceLambda         ServiceNamespace = "lambda"
	ServiceNamespaceCassandra      ServiceNamespace = "cassandra"
	ServiceNamespaceKafka          ServiceNamespace = "kafka"
	ServiceNamespaceElasticache    ServiceNamespace = "elasticache"
	ServiceNamespaceNeptune        ServiceNamespace = "neptune"
)

// Values returns all known values for ServiceNamespace. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServiceNamespace) Values() []ServiceNamespace {
	return []ServiceNamespace{
		"ecs",
		"elasticmapreduce",
		"ec2",
		"appstream",
		"dynamodb",
		"rds",
		"sagemaker",
		"custom-resource",
		"comprehend",
		"lambda",
		"cassandra",
		"kafka",
		"elasticache",
		"neptune",
	}
}
