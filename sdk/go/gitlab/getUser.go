// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gitlab

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # gitlab\_user
//
// Provides details about a specific user in the gitlab provider. Especially the ability to lookup the id for linking to other resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gitlab/sdk/v4/go/gitlab"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "myuser"
// 		_, err := gitlab.LookupUser(ctx, &gitlab.LookupUserArgs{
// 			Username: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("gitlab:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// The e-mail address of the user. (Requires administrator privileges)
	Email *string `pulumi:"email"`
	// The ID of the user.
	UserId *int `pulumi:"userId"`
	// The username of the user.
	Username *string `pulumi:"username"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// The avatar URL of the user.
	AvatarUrl string `pulumi:"avatarUrl"`
	// The bio of the user.
	Bio string `pulumi:"bio"`
	// Whether the user can create groups.
	CanCreateGroup bool `pulumi:"canCreateGroup"`
	// Whether the user can create projects.
	CanCreateProject bool `pulumi:"canCreateProject"`
	// User's color scheme ID.
	ColorSchemeId int `pulumi:"colorSchemeId"`
	// Date the user was created at.
	CreatedAt string `pulumi:"createdAt"`
	// Current user's sign-in date.
	CurrentSignInAt string `pulumi:"currentSignInAt"`
	// The e-mail address of the user.
	Email string `pulumi:"email"`
	// The external UID of the user.
	ExternUid string `pulumi:"externUid"`
	// Whether the user is external.
	External bool `pulumi:"external"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Whether the user is an admin.
	IsAdmin bool `pulumi:"isAdmin"`
	// Last user's sign-in date.
	LastSignInAt string `pulumi:"lastSignInAt"`
	// Linkedin profile of the user.
	Linkedin string `pulumi:"linkedin"`
	// The location of the user.
	Location string `pulumi:"location"`
	// The name of the user.
	Name string `pulumi:"name"`
	Note string `pulumi:"note"`
	// The organization of the user.
	Organization string `pulumi:"organization"`
	// Number of projects the user can create.
	ProjectsLimit int `pulumi:"projectsLimit"`
	// Skype username of the user.
	Skype string `pulumi:"skype"`
	// Whether the user is active or blocked.
	State string `pulumi:"state"`
	// User's theme ID.
	ThemeId int `pulumi:"themeId"`
	// Twitter username of the user.
	Twitter string `pulumi:"twitter"`
	// Whether user's two factor auth is enabled.
	TwoFactorEnabled bool `pulumi:"twoFactorEnabled"`
	UserId           int  `pulumi:"userId"`
	// The UID provider of the user.
	UserProvider string `pulumi:"userProvider"`
	// The username of the user.
	Username string `pulumi:"username"`
	// User's website URL.
	WebsiteUrl string `pulumi:"websiteUrl"`
}
