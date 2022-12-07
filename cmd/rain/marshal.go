package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/ptr"
	"gopkg.in/yaml.v3"
)

type StackSetConfig struct {

	// The name to associate with the stack set. The name must be unique in the Region
	// where you create your stack set. A stack name can contain only alphanumeric
	// characters (case-sensitive) and hyphens. It must start with an alphabetic
	// character and can't be longer than 128 characters.
	//
	// This member is required.
	StackSetName *string

	// The Amazon Resource Number (ARN) of the IAM role to use to create this stack
	// set. Specify an IAM role only if you are using customized administrator roles to
	// control which users or groups can manage specific stack sets within the same
	// administrator account. For more information, see Prerequisites: Granting
	// Permissions for Stack Set Operations
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html)
	// in the CloudFormation User Guide.
	AdministrationRoleARN *string

	// Describes whether StackSets automatically deploys to Organizations accounts that
	// are added to the target organization or organizational unit (OU). Specify only
	// if PermissionModel is SERVICE_MANAGED.
	AutoDeployment *types.AutoDeployment

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account. By default, SELF is specified. Use SELF for
	// stack sets with self-managed permissions.
	//
	// * To create a stack set with
	// service-managed permissions while signed in to the management account, specify
	// SELF.
	//
	// * To create a stack set with service-managed permissions while signed in
	// to a delegated administrator account, specify DELEGATED_ADMIN. Your Amazon Web
	// Services account must be registered as a delegated admin in the management
	// account. For more information, see Register a delegated administrator
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html)
	// in the CloudFormation User Guide.
	//
	// Stack sets with service-managed permissions
	// are created in the management account, including stack sets that are created by
	// delegated administrators.
	CallAs types.CallAs

	// In some cases, you must explicitly acknowledge that your stack set template
	// contains certain capabilities in order for CloudFormation to create the stack
	// set and related stack instances.
	//
	// * CAPABILITY_IAM and CAPABILITY_NAMED_IAM Some
	// stack templates might include resources that can affect permissions in your
	// Amazon Web Services account; for example, by creating new Identity and Access
	// Management (IAM) users. For those stack sets, you must explicitly acknowledge
	// this by specifying one of these capabilities. The following IAM resources
	// require you to specify either the CAPABILITY_IAM or CAPABILITY_NAMED_IAM
	// capability.
	//
	// * If you have IAM resources, you can specify either capability.
	//
	// *
	// If you have IAM resources with custom names, you must specify
	// CAPABILITY_NAMED_IAM.
	//
	// * If you don't specify either of these capabilities,
	// CloudFormation returns an InsufficientCapabilities error.
	//
	// If your stack
	// template contains these resources, we recommend that you review all permissions
	// associated with them and edit their permissions if necessary.
	//
	// *
	// AWS::IAM::AccessKey
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html)
	//
	// *
	// AWS::IAM::Group
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html)
	//
	// *
	// AWS::IAM::InstanceProfile
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html)
	//
	// *
	// AWS::IAM::Policy
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html)
	//
	// *
	// AWS::IAM::Role
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html)
	//
	// *
	// AWS::IAM::User
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html)
	//
	// *
	// AWS::IAM::UserToGroupAddition
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html)
	//
	// For
	// more information, see Acknowledging IAM Resources in CloudFormation Templates
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities).
	//
	// *
	// CAPABILITY_AUTO_EXPAND Some templates reference macros. If your stack set
	// template references one or more macros, you must create the stack set directly
	// from the processed template, without first reviewing the resulting changes in a
	// change set. To create the stack set directly, you must acknowledge this
	// capability. For more information, see Using CloudFormation Macros to Perform
	// Custom Processing on Templates
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html).
	// Stack sets with service-managed permissions don't currently support the use of
	// macros in templates. (This includes the AWS::Include
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/create-reusable-transform-function-snippets-and-add-to-your-template-with-aws-include-transform.html)
	// and AWS::Serverless
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-aws-serverless.html)
	// transforms, which are macros hosted by CloudFormation.) Even if you specify this
	// capability for a stack set with service-managed permissions, if you reference a
	// macro in your template the stack set operation will fail.
	Capabilities []types.Capability

	// A description of the stack set. You can use the description to identify the
	// stack set's purpose or other important information.
	Description *string

	// The name of the IAM execution role to use to create the stack set. If you do not
	// specify an execution role, CloudFormation uses the
	// AWSCloudFormationStackSetExecutionRole role for the stack set operation. Specify
	// an IAM role only if you are using customized execution roles to control which
	// stack resources users and groups can include in their stack sets.
	ExecutionRoleName *string

	// Describes whether StackSets performs non-conflicting operations concurrently and
	// queues conflicting operations.
	ManagedExecution *types.ManagedExecution

	// The input parameters for the stack set template.
	Parameters []types.Parameter

	// Describes how the IAM roles required for stack set operations are created. By
	// default, SELF-MANAGED is specified.
	//
	// * With self-managed permissions, you must
	// create the administrator and execution roles required to deploy to target
	// accounts. For more information, see Grant Self-Managed Stack Set Permissions
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-self-managed.html).
	//
	// *
	// With service-managed permissions, StackSets automatically creates the IAM roles
	// required to deploy to accounts managed by Organizations. For more information,
	// see Grant Service-Managed Stack Set Permissions
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs-service-managed.html).
	PermissionModel types.PermissionModels

	// The stack ID you are importing into a new stack set. Specify the Amazon Resource
	// Number (ARN) of the stack.
	StackId *string

	// The key-value pairs to associate with this stack set and the stacks created from
	// it. CloudFormation also propagates these tags to supported resources that are
	// created in the stacks. A maximum number of 50 tags can be specified. If you
	// specify tags as part of a CreateStackSet action, CloudFormation checks to see if
	// you have the required IAM permission to tag resources. If you don't, the entire
	// CreateStackSet action fails with an access denied error, and the stack set is
	// not created.
	Tags []types.Tag
}

type StackSetInstancesConfig struct {

	// The names of one or more Amazon Web Services Regions where you want to create
	// stack instances using the specified Amazon Web Services accounts.
	//
	// This member is required.
	Regions []string

	// The name or unique ID of the stack set that you want to create stack instances
	// from.
	//
	// This member is required.
	StackSetName *string

	// [Self-managed permissions] The names of one or more Amazon Web Services accounts
	// that you want to create stack instances in the specified Region(s) for. You can
	// specify Accounts or DeploymentTargets, but not both.
	Accounts []string

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account. By default, SELF is specified. Use SELF for
	// stack sets with self-managed permissions.
	//
	// * If you are signed in to the
	// management account, specify SELF.
	//
	// * If you are signed in to a delegated
	// administrator account, specify DELEGATED_ADMIN. Your Amazon Web Services account
	// must be registered as a delegated administrator in the management account. For
	// more information, see Register a delegated administrator
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html)
	// in the CloudFormation User Guide.
	CallAs types.CallAs

	// [Service-managed permissions] The Organizations accounts for which to create
	// stack instances in the specified Amazon Web Services Regions. You can specify
	// Accounts or DeploymentTargets, but not both.
	DeploymentTargets *types.DeploymentTargets

	// Preferences for how CloudFormation performs this stack set operation.
	OperationPreferences *types.StackSetOperationPreferences
}

type configFileFormat struct {
	Parameters        map[string]string `yaml:"Parameters"`
	Tags              map[string]string `yaml:"Tags"`
	StackSet          StackSetConfig
	StackSetInstanses StackSetInstancesConfig
}

func main() {
	t := configFileFormat{}
	t.Parameters = make(map[string]string)
	t.Tags = make(map[string]string)
	t.Parameters["queueName"] = "testQueue"
	t.Tags["stackset"] = "testStackSet"
	t.StackSet.StackSetName = ptr.String("test")
	t.StackSet.Description = ptr.String("test description")
	t.StackSet.Capabilities = []types.Capability{
		"CAPABILITY_NAMED_IAM",
		"CAPABILITY_AUTO_EXPAND",
	}
	t.StackSetInstanses.Accounts = append(t.StackSetInstanses.Accounts, "577613639135")
	t.StackSetInstanses.Regions = append(t.StackSetInstanses.Regions, "us-east-1", "us-east-2")

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

}
