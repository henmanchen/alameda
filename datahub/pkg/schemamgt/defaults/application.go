package defaults

import (
	"github.com/containers-ai/alameda/internal/pkg/database/common"
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
)

func SchemaApplicationKafkaTopic() *schemas.Schema {
	// Kafka topic
	schema := schemas.NewSchema(schemas.Application, "kafka", "topic")
	measurement := schemas.NewMeasurement("kafka_topic", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("dummy", true, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}

func SchemaApplicationKafkaCG() *schemas.Schema {
	// Kafka consumer group
	schema := schemas.NewSchema(schemas.Application, "kafka", "consumer_group")
	measurement := schemas.NewMeasurement("kafka_consumer_group", schemas.MetricTypeUndefined, schemas.ResourceBoundaryUndefined, schemas.ResourceQuotaUndefined)
	measurement.AddColumn("name", true, schemas.Tag, common.String)
	measurement.AddColumn("namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("cluster_name", true, schemas.Tag, common.String)
	measurement.AddColumn("topic_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_name", true, schemas.Tag, common.String)
	measurement.AddColumn("alameda_scaler_namespace", true, schemas.Tag, common.String)
	measurement.AddColumn("resource_k8s_namespace", false, schemas.Field, common.String)
	measurement.AddColumn("resource_k8s_name", false, schemas.Field, common.String)
	measurement.AddColumn("resource_k8s_kind", false, schemas.Field, common.String)
	measurement.AddColumn("resource_k8s_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("resource_k8s_spec_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("resource_k8s_min_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("resource_k8s_max_replicas", false, schemas.Field, common.Int32)
	measurement.AddColumn("resource_custom_name", false, schemas.Field, common.String)
	measurement.AddColumn("policy", false, schemas.Field, common.String)
	measurement.AddColumn("enable_execution", false, schemas.Field, common.Bool)
	measurement.AddColumn("resource_cpu_limit", false, schemas.Field, common.String)
	measurement.AddColumn("resource_cpu_request", false, schemas.Field, common.String)
	measurement.AddColumn("resource_memory_limit", false, schemas.Field, common.String)
	measurement.AddColumn("resource_memory_request", false, schemas.Field, common.String)
	measurement.AddColumn("volumes_size", false, schemas.Field, common.String)
	measurement.AddColumn("volumes_pvc_size", false, schemas.Field, common.String)
	schema.Measurements = append(schema.Measurements, measurement)
	return schema
}
