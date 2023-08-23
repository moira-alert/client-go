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

type MoiraScheduledNotification struct {
	Contact *MoiraContactData `json:"contact,omitempty"`
	Event *MoiraNotificationEvent `json:"event,omitempty"`
	Plotting *MoiraPlottingData `json:"plotting,omitempty"`
	SendFail int32 `json:"send_fail,omitempty"`
	Throttled bool `json:"throttled,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
	Trigger *MoiraTriggerData `json:"trigger,omitempty"`
}
