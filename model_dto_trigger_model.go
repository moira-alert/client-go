/*
 * Moira Alert
 *
 * This is an API description for [Moira Alert API](https://moira.readthedocs.io/en/latest/overview.html) Check us out on [Github](https://github.com/moira-alert) or look up our [guide on getting started with Moira](https://moira.readthedocs.io)
 *
 * API version: master
 * Contact: opensource@skbkontur.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DtoTriggerModel struct {
	// A list of targets that have only alone metrics
	AloneMetrics map[string]bool `json:"alone_metrics,omitempty"`
	// Datetime when the trigger was created
	CreatedAt string `json:"created_at,omitempty"`
	// Username who created trigger
	CreatedBy string `json:"created_by,omitempty"`
	// Description string
	Desc string `json:"desc,omitempty"`
	// ERROR threshold
	ErrorValue float64 `json:"error_value,omitempty"`
	// Used if you need more complex logic than provided by WARN/ERROR values
	Expression string `json:"expression,omitempty"`
	// Trigger unique ID
	Id string `json:"id,omitempty"`
	// Shows if trigger is remote (graphite-backend) based or stored inside Moira-Redis DB
	IsRemote bool `json:"is_remote,omitempty"`
	// If true, first event NODATA → OK will be omitted
	MuteNewMetrics bool `json:"mute_new_metrics,omitempty"`
	// Trigger name
	Name string `json:"name,omitempty"`
	// Graphite patterns for trigger
	Patterns []string `json:"patterns,omitempty"`
	// Determines when Moira should monitor trigger
	Sched *AllOfdtoTriggerModelSched `json:"sched,omitempty"`
	// Set of tags to manipulate subscriptions
	Tags []string `json:"tags,omitempty"`
	// Graphite-like targets: t1, t2, ...
	Targets []string `json:"targets,omitempty"`
	// Could be: rising, falling, expression
	TriggerType string `json:"trigger_type,omitempty"`
	// When there are no metrics for trigger, Moira will switch metric to TTLState state after TTL seconds
	Ttl int32 `json:"ttl,omitempty"`
	// When there are no metrics for trigger, Moira will switch metric to TTLState state after TTL seconds
	TtlState string `json:"ttl_state,omitempty"`
	// Datetime  when the trigger was updated
	UpdatedAt string `json:"updated_at,omitempty"`
	// Username who updated trigger
	UpdatedBy string `json:"updated_by,omitempty"`
	// WARN threshold
	WarnValue float64 `json:"warn_value,omitempty"`
}
