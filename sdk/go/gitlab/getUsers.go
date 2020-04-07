// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package gitlab

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides details about a list of users in the gitlab provider. The results include id, username, email, name and more about the requested users. Users can also be sorted and filtered using several options.
//
// **NOTE**: Some of the available options require administrator privileges. Please visit [Gitlab API documentation][usersForAdmins] for more information.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-gitlab/blob/master/website/docs/d/users.html.markdown.
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	var rv GetUsersResult
	err := ctx.Invoke("gitlab:index/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	// Filter users that are active.
	Active *bool `pulumi:"active"`
	// Filter users that are blocked.
	Blocked *bool `pulumi:"blocked"`
	// Search for users created after a specific date. (Requires administrator privileges)
	CreatedAfter *string `pulumi:"createdAfter"`
	// Search for users created before a specific date. (Requires administrator privileges)
	CreatedBefore *string `pulumi:"createdBefore"`
	// Lookup users by external provider. (Requires administrator privileges)
	ExternProvider *string `pulumi:"externProvider"`
	// Lookup users by external UID. (Requires administrator privileges)
	ExternUid *string `pulumi:"externUid"`
	// Order the users' list by `id`, `name`, `username`, `createdAt` or `updatedAt`. (Requires administrator privileges)
	OrderBy *string `pulumi:"orderBy"`
	// Search users by username, name or email.
	Search *string `pulumi:"search"`
	// Sort users' list in asc or desc order. (Requires administrator privileges)
	Sort *string `pulumi:"sort"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	Active         *bool   `pulumi:"active"`
	Blocked        *bool   `pulumi:"blocked"`
	CreatedAfter   *string `pulumi:"createdAfter"`
	CreatedBefore  *string `pulumi:"createdBefore"`
	ExternProvider *string `pulumi:"externProvider"`
	// The external UID of the user.
	ExternUid *string `pulumi:"externUid"`
	// id is the provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	OrderBy *string `pulumi:"orderBy"`
	Search  *string `pulumi:"search"`
	Sort    *string `pulumi:"sort"`
	// The list of users.
	Users []GetUsersUser `pulumi:"users"`
}
