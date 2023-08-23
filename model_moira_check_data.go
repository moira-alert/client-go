/*
 * Moira Alert
 *
 * This is an API description for [Moira Alert API](https://moira.readthedocs.io/en/latest/overview.html) Check us out on [Github](https://github.com/moira-alert) or look up our [guide](https://moira.readthedocs.io) on getting started with Moira
 *
 * API version: master
 * Contact: opensource@skbkontur.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MoiraCheckData struct {
	EventTimestamp int64 `json:"event_timestamp,omitempty"`
	LastSuccessfulCheckTimestamp int64 `json:"last_successful_check_timestamp,omitempty"`
	Maintenance int64 `json:"maintenance,omitempty"`
	MaintenanceInfo *MoiraMaintenanceInfo `json:"maintenance_info,omitempty"`
	Metrics map[string]MoiraMetricState `json:"metrics,omitempty"`
	// MetricsToTargetRelation is a map that holds relation between metric names that was alone during last check and targets that fetched this metric  {\"t1\": \"metric.name.1\", \"t2\": \"metric.name.2\"}
	MetricsToTargetRelation map[string]string `json:"metrics_to_target_relation,omitempty"`
	Msg string `json:"msg,omitempty"`
	Score int64 `json:"score,omitempty"`
	State string `json:"state,omitempty"`
	Suppressed bool `json:"suppressed,omitempty"`
	SuppressedState string `json:"suppressed_state,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
}
