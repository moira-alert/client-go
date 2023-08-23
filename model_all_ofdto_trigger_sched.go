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

// Determines when Moira should monitor trigger
type AllOfdtoTriggerSched struct {
	Days []MoiraScheduleDataDay `json:"days,omitempty"`
	EndOffset int32 `json:"endOffset,omitempty"`
	StartOffset int32 `json:"startOffset,omitempty"`
	TzOffset int32 `json:"tzOffset,omitempty"`
}
