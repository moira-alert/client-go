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

type DtoTreeOfProblems struct {
	SyntaxOk bool `json:"syntax_ok,omitempty"`
	TreeOfProblems *DtoProblemOfTarget `json:"tree_of_problems,omitempty"`
}
