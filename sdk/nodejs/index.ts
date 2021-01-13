// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./branchProtection";
export * from "./deployKey";
export * from "./deployKeyEnable";
export * from "./deployToken";
export * from "./getGroup";
export * from "./getGroupMembership";
export * from "./getProject";
export * from "./getProjects";
export * from "./getUser";
export * from "./getUsers";
export * from "./group";
export * from "./groupCluster";
export * from "./groupLabel";
export * from "./groupLdapLink";
export * from "./groupMembership";
export * from "./groupVariable";
export * from "./instanceCluster";
export * from "./instanceVariable";
export * from "./label";
export * from "./pipelineSchedule";
export * from "./pipelineScheduleVariable";
export * from "./pipelineTrigger";
export * from "./project";
export * from "./projectApprovalRule";
export * from "./projectCluster";
export * from "./projectHook";
export * from "./projectLevelMrApprovals";
export * from "./projectMembership";
export * from "./projectMirror";
export * from "./projectShareGroup";
export * from "./projectVariable";
export * from "./provider";
export * from "./serviceGithub";
export * from "./serviceJira";
export * from "./servicePipelinesEmail";
export * from "./serviceSlack";
export * from "./tagProtection";
export * from "./user";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { BranchProtection } from "./branchProtection";
import { DeployKey } from "./deployKey";
import { DeployKeyEnable } from "./deployKeyEnable";
import { DeployToken } from "./deployToken";
import { Group } from "./group";
import { GroupCluster } from "./groupCluster";
import { GroupLabel } from "./groupLabel";
import { GroupLdapLink } from "./groupLdapLink";
import { GroupMembership } from "./groupMembership";
import { GroupVariable } from "./groupVariable";
import { InstanceCluster } from "./instanceCluster";
import { InstanceVariable } from "./instanceVariable";
import { Label } from "./label";
import { PipelineSchedule } from "./pipelineSchedule";
import { PipelineScheduleVariable } from "./pipelineScheduleVariable";
import { PipelineTrigger } from "./pipelineTrigger";
import { Project } from "./project";
import { ProjectApprovalRule } from "./projectApprovalRule";
import { ProjectCluster } from "./projectCluster";
import { ProjectHook } from "./projectHook";
import { ProjectLevelMrApprovals } from "./projectLevelMrApprovals";
import { ProjectMembership } from "./projectMembership";
import { ProjectMirror } from "./projectMirror";
import { ProjectShareGroup } from "./projectShareGroup";
import { ProjectVariable } from "./projectVariable";
import { ServiceGithub } from "./serviceGithub";
import { ServiceJira } from "./serviceJira";
import { ServicePipelinesEmail } from "./servicePipelinesEmail";
import { ServiceSlack } from "./serviceSlack";
import { TagProtection } from "./tagProtection";
import { User } from "./user";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gitlab:index/branchProtection:BranchProtection":
                return new BranchProtection(name, <any>undefined, { urn })
            case "gitlab:index/deployKey:DeployKey":
                return new DeployKey(name, <any>undefined, { urn })
            case "gitlab:index/deployKeyEnable:DeployKeyEnable":
                return new DeployKeyEnable(name, <any>undefined, { urn })
            case "gitlab:index/deployToken:DeployToken":
                return new DeployToken(name, <any>undefined, { urn })
            case "gitlab:index/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "gitlab:index/groupCluster:GroupCluster":
                return new GroupCluster(name, <any>undefined, { urn })
            case "gitlab:index/groupLabel:GroupLabel":
                return new GroupLabel(name, <any>undefined, { urn })
            case "gitlab:index/groupLdapLink:GroupLdapLink":
                return new GroupLdapLink(name, <any>undefined, { urn })
            case "gitlab:index/groupMembership:GroupMembership":
                return new GroupMembership(name, <any>undefined, { urn })
            case "gitlab:index/groupVariable:GroupVariable":
                return new GroupVariable(name, <any>undefined, { urn })
            case "gitlab:index/instanceCluster:InstanceCluster":
                return new InstanceCluster(name, <any>undefined, { urn })
            case "gitlab:index/instanceVariable:InstanceVariable":
                return new InstanceVariable(name, <any>undefined, { urn })
            case "gitlab:index/label:Label":
                return new Label(name, <any>undefined, { urn })
            case "gitlab:index/pipelineSchedule:PipelineSchedule":
                return new PipelineSchedule(name, <any>undefined, { urn })
            case "gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable":
                return new PipelineScheduleVariable(name, <any>undefined, { urn })
            case "gitlab:index/pipelineTrigger:PipelineTrigger":
                return new PipelineTrigger(name, <any>undefined, { urn })
            case "gitlab:index/project:Project":
                return new Project(name, <any>undefined, { urn })
            case "gitlab:index/projectApprovalRule:ProjectApprovalRule":
                return new ProjectApprovalRule(name, <any>undefined, { urn })
            case "gitlab:index/projectCluster:ProjectCluster":
                return new ProjectCluster(name, <any>undefined, { urn })
            case "gitlab:index/projectHook:ProjectHook":
                return new ProjectHook(name, <any>undefined, { urn })
            case "gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals":
                return new ProjectLevelMrApprovals(name, <any>undefined, { urn })
            case "gitlab:index/projectMembership:ProjectMembership":
                return new ProjectMembership(name, <any>undefined, { urn })
            case "gitlab:index/projectMirror:ProjectMirror":
                return new ProjectMirror(name, <any>undefined, { urn })
            case "gitlab:index/projectShareGroup:ProjectShareGroup":
                return new ProjectShareGroup(name, <any>undefined, { urn })
            case "gitlab:index/projectVariable:ProjectVariable":
                return new ProjectVariable(name, <any>undefined, { urn })
            case "gitlab:index/serviceGithub:ServiceGithub":
                return new ServiceGithub(name, <any>undefined, { urn })
            case "gitlab:index/serviceJira:ServiceJira":
                return new ServiceJira(name, <any>undefined, { urn })
            case "gitlab:index/servicePipelinesEmail:ServicePipelinesEmail":
                return new ServicePipelinesEmail(name, <any>undefined, { urn })
            case "gitlab:index/serviceSlack:ServiceSlack":
                return new ServiceSlack(name, <any>undefined, { urn })
            case "gitlab:index/tagProtection:TagProtection":
                return new TagProtection(name, <any>undefined, { urn })
            case "gitlab:index/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gitlab", "index/branchProtection", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/deployKey", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/deployKeyEnable", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/deployToken", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/group", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupCluster", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupLabel", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupLdapLink", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupMembership", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupVariable", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/instanceCluster", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/instanceVariable", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/label", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/pipelineSchedule", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/pipelineScheduleVariable", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/pipelineTrigger", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/project", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectApprovalRule", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectCluster", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectHook", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectLevelMrApprovals", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectMembership", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectMirror", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectShareGroup", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectVariable", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/serviceGithub", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/serviceJira", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/servicePipelinesEmail", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/serviceSlack", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/tagProtection", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/user", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("gitlab", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:gitlab") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
