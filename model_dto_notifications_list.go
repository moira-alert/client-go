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

type DtoNotificationsList struct {
	List []MoiraScheduledNotification `json:"list,omitempty"`
	Total int64 `json:"total,omitempty"`
}
