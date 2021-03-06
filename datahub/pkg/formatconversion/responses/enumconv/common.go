package enumconv

import (
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb/schemas"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
)

var MetricTypeNameMap = map[schemas.MetricType]common.MetricType{
	schemas.MetricTypeUndefined:       common.MetricType_METRICS_TYPE_UNDEFINED,
	schemas.CPUUsageSecondsPercentage: common.MetricType_CPU_USAGE_SECONDS_PERCENTAGE,
	schemas.MemoryUsageBytes:          common.MetricType_MEMORY_USAGE_BYTES,
	schemas.PowerUsageWatts:           common.MetricType_POWER_USAGE_WATTS,
	schemas.TemperatureCelsius:        common.MetricType_TEMPERATURE_CELSIUS,
	schemas.DutyCycle:                 common.MetricType_DUTY_CYCLE,
	schemas.CurrentOffset:             common.MetricType_CURRENT_OFFSET,
	schemas.Lag:                       common.MetricType_LAG,
}

var ResourceBoundaryNameMap = map[schemas.ResourceBoundary]common.ResourceBoundary{
	schemas.ResourceBoundaryUndefined: common.ResourceBoundary_RESOURCE_BOUNDARY_UNDEFINED,
	schemas.ResourceRaw:               common.ResourceBoundary_RESOURCE_RAW,
	schemas.ResourceUpperBound:        common.ResourceBoundary_RESOURCE_UPPER_BOUND,
	schemas.ResourceLowerBound:        common.ResourceBoundary_RESOURCE_LOWER_BOUND,
}

var ResourceQuotaNameMap = map[schemas.ResourceQuota]common.ResourceQuota{
	schemas.ResourceQuotaUndefined: common.ResourceQuota_RESOURCE_QUOTA_UNDEFINED,
	schemas.ResourceLimit:          common.ResourceQuota_RESOURCE_LIMIT,
	schemas.ResourceRequest:        common.ResourceQuota_RESOURCE_REQUEST,
	schemas.ResourceInitialLimit:   common.ResourceQuota_RESOURCE_INITIAL_LIMIT,
	schemas.ResourceInitialRequest: common.ResourceQuota_RESOURCE_INITIAL_REQUEST,
}
