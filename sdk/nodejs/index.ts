// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./branch";
export * from "./branchProtection";
export * from "./clusterAgent";
export * from "./clusterAgentToken";
export * from "./deployKey";
export * from "./deployKeyEnable";
export * from "./deployToken";
export * from "./getBranch";
export * from "./getClusterAgent";
export * from "./getClusterAgents";
export * from "./getCurrentUser";
export * from "./getGroup";
export * from "./getGroupMembership";
export * from "./getGroupVariable";
export * from "./getGroupVariables";
export * from "./getInstanceDeployKeys";
export * from "./getInstanceVariable";
export * from "./getInstanceVariables";
export * from "./getProject";
export * from "./getProjectIssue";
export * from "./getProjectIssues";
export * from "./getProjectMilestone";
export * from "./getProjectMilestones";
export * from "./getProjectProtectedBranch";
export * from "./getProjectProtectedBranches";
export * from "./getProjectTag";
export * from "./getProjectTags";
export * from "./getProjectVariable";
export * from "./getProjectVariables";
export * from "./getProjects";
export * from "./getReleaseLink";
export * from "./getReleaseLinks";
export * from "./getRepositoryFile";
export * from "./getUser";
export * from "./getUsers";
export * from "./group";
export * from "./groupAccessToken";
export * from "./groupBadge";
export * from "./groupCluster";
export * from "./groupCustomAttribute";
export * from "./groupLabel";
export * from "./groupLdapLink";
export * from "./groupMembership";
export * from "./groupProjectFileTemplate";
export * from "./groupShareGroup";
export * from "./groupVariable";
export * from "./instanceCluster";
export * from "./instanceVariable";
export * from "./label";
export * from "./managedLicense";
export * from "./personalAccessToken";
export * from "./pipelineSchedule";
export * from "./pipelineScheduleVariable";
export * from "./pipelineTrigger";
export * from "./project";
export * from "./projectAccessToken";
export * from "./projectApprovalRule";
export * from "./projectBadge";
export * from "./projectCluster";
export * from "./projectCustomAttribute";
export * from "./projectEnvironment";
export * from "./projectFreezePeriod";
export * from "./projectHook";
export * from "./projectIssue";
export * from "./projectLevelMrApprovals";
export * from "./projectMembership";
export * from "./projectMilestone";
export * from "./projectMirror";
export * from "./projectProtectedEnvironment";
export * from "./projectRunnerEnablement";
export * from "./projectShareGroup";
export * from "./projectTag";
export * from "./projectVariable";
export * from "./provider";
export * from "./releaseLink";
export * from "./repositoryFile";
export * from "./runner";
export * from "./serviceExternalWiki";
export * from "./serviceGithub";
export * from "./serviceJira";
export * from "./serviceMicrosoftTeams";
export * from "./servicePipelinesEmail";
export * from "./serviceSlack";
export * from "./systemHook";
export * from "./tagProtection";
export * from "./topic";
export * from "./user";
export * from "./userCustomAttribute";
export * from "./userSshKey";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { Branch } from "./branch";
import { BranchProtection } from "./branchProtection";
import { ClusterAgent } from "./clusterAgent";
import { ClusterAgentToken } from "./clusterAgentToken";
import { DeployKey } from "./deployKey";
import { DeployKeyEnable } from "./deployKeyEnable";
import { DeployToken } from "./deployToken";
import { Group } from "./group";
import { GroupAccessToken } from "./groupAccessToken";
import { GroupBadge } from "./groupBadge";
import { GroupCluster } from "./groupCluster";
import { GroupCustomAttribute } from "./groupCustomAttribute";
import { GroupLabel } from "./groupLabel";
import { GroupLdapLink } from "./groupLdapLink";
import { GroupMembership } from "./groupMembership";
import { GroupProjectFileTemplate } from "./groupProjectFileTemplate";
import { GroupShareGroup } from "./groupShareGroup";
import { GroupVariable } from "./groupVariable";
import { InstanceCluster } from "./instanceCluster";
import { InstanceVariable } from "./instanceVariable";
import { Label } from "./label";
import { ManagedLicense } from "./managedLicense";
import { PersonalAccessToken } from "./personalAccessToken";
import { PipelineSchedule } from "./pipelineSchedule";
import { PipelineScheduleVariable } from "./pipelineScheduleVariable";
import { PipelineTrigger } from "./pipelineTrigger";
import { Project } from "./project";
import { ProjectAccessToken } from "./projectAccessToken";
import { ProjectApprovalRule } from "./projectApprovalRule";
import { ProjectBadge } from "./projectBadge";
import { ProjectCluster } from "./projectCluster";
import { ProjectCustomAttribute } from "./projectCustomAttribute";
import { ProjectEnvironment } from "./projectEnvironment";
import { ProjectFreezePeriod } from "./projectFreezePeriod";
import { ProjectHook } from "./projectHook";
import { ProjectIssue } from "./projectIssue";
import { ProjectLevelMrApprovals } from "./projectLevelMrApprovals";
import { ProjectMembership } from "./projectMembership";
import { ProjectMilestone } from "./projectMilestone";
import { ProjectMirror } from "./projectMirror";
import { ProjectProtectedEnvironment } from "./projectProtectedEnvironment";
import { ProjectRunnerEnablement } from "./projectRunnerEnablement";
import { ProjectShareGroup } from "./projectShareGroup";
import { ProjectTag } from "./projectTag";
import { ProjectVariable } from "./projectVariable";
import { ReleaseLink } from "./releaseLink";
import { RepositoryFile } from "./repositoryFile";
import { Runner } from "./runner";
import { ServiceExternalWiki } from "./serviceExternalWiki";
import { ServiceGithub } from "./serviceGithub";
import { ServiceJira } from "./serviceJira";
import { ServiceMicrosoftTeams } from "./serviceMicrosoftTeams";
import { ServicePipelinesEmail } from "./servicePipelinesEmail";
import { ServiceSlack } from "./serviceSlack";
import { SystemHook } from "./systemHook";
import { TagProtection } from "./tagProtection";
import { Topic } from "./topic";
import { User } from "./user";
import { UserCustomAttribute } from "./userCustomAttribute";
import { UserSshKey } from "./userSshKey";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gitlab:index/branch:Branch":
                return new Branch(name, <any>undefined, { urn })
            case "gitlab:index/branchProtection:BranchProtection":
                return new BranchProtection(name, <any>undefined, { urn })
            case "gitlab:index/clusterAgent:ClusterAgent":
                return new ClusterAgent(name, <any>undefined, { urn })
            case "gitlab:index/clusterAgentToken:ClusterAgentToken":
                return new ClusterAgentToken(name, <any>undefined, { urn })
            case "gitlab:index/deployKey:DeployKey":
                return new DeployKey(name, <any>undefined, { urn })
            case "gitlab:index/deployKeyEnable:DeployKeyEnable":
                return new DeployKeyEnable(name, <any>undefined, { urn })
            case "gitlab:index/deployToken:DeployToken":
                return new DeployToken(name, <any>undefined, { urn })
            case "gitlab:index/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "gitlab:index/groupAccessToken:GroupAccessToken":
                return new GroupAccessToken(name, <any>undefined, { urn })
            case "gitlab:index/groupBadge:GroupBadge":
                return new GroupBadge(name, <any>undefined, { urn })
            case "gitlab:index/groupCluster:GroupCluster":
                return new GroupCluster(name, <any>undefined, { urn })
            case "gitlab:index/groupCustomAttribute:GroupCustomAttribute":
                return new GroupCustomAttribute(name, <any>undefined, { urn })
            case "gitlab:index/groupLabel:GroupLabel":
                return new GroupLabel(name, <any>undefined, { urn })
            case "gitlab:index/groupLdapLink:GroupLdapLink":
                return new GroupLdapLink(name, <any>undefined, { urn })
            case "gitlab:index/groupMembership:GroupMembership":
                return new GroupMembership(name, <any>undefined, { urn })
            case "gitlab:index/groupProjectFileTemplate:GroupProjectFileTemplate":
                return new GroupProjectFileTemplate(name, <any>undefined, { urn })
            case "gitlab:index/groupShareGroup:GroupShareGroup":
                return new GroupShareGroup(name, <any>undefined, { urn })
            case "gitlab:index/groupVariable:GroupVariable":
                return new GroupVariable(name, <any>undefined, { urn })
            case "gitlab:index/instanceCluster:InstanceCluster":
                return new InstanceCluster(name, <any>undefined, { urn })
            case "gitlab:index/instanceVariable:InstanceVariable":
                return new InstanceVariable(name, <any>undefined, { urn })
            case "gitlab:index/label:Label":
                return new Label(name, <any>undefined, { urn })
            case "gitlab:index/managedLicense:ManagedLicense":
                return new ManagedLicense(name, <any>undefined, { urn })
            case "gitlab:index/personalAccessToken:PersonalAccessToken":
                return new PersonalAccessToken(name, <any>undefined, { urn })
            case "gitlab:index/pipelineSchedule:PipelineSchedule":
                return new PipelineSchedule(name, <any>undefined, { urn })
            case "gitlab:index/pipelineScheduleVariable:PipelineScheduleVariable":
                return new PipelineScheduleVariable(name, <any>undefined, { urn })
            case "gitlab:index/pipelineTrigger:PipelineTrigger":
                return new PipelineTrigger(name, <any>undefined, { urn })
            case "gitlab:index/project:Project":
                return new Project(name, <any>undefined, { urn })
            case "gitlab:index/projectAccessToken:ProjectAccessToken":
                return new ProjectAccessToken(name, <any>undefined, { urn })
            case "gitlab:index/projectApprovalRule:ProjectApprovalRule":
                return new ProjectApprovalRule(name, <any>undefined, { urn })
            case "gitlab:index/projectBadge:ProjectBadge":
                return new ProjectBadge(name, <any>undefined, { urn })
            case "gitlab:index/projectCluster:ProjectCluster":
                return new ProjectCluster(name, <any>undefined, { urn })
            case "gitlab:index/projectCustomAttribute:ProjectCustomAttribute":
                return new ProjectCustomAttribute(name, <any>undefined, { urn })
            case "gitlab:index/projectEnvironment:ProjectEnvironment":
                return new ProjectEnvironment(name, <any>undefined, { urn })
            case "gitlab:index/projectFreezePeriod:ProjectFreezePeriod":
                return new ProjectFreezePeriod(name, <any>undefined, { urn })
            case "gitlab:index/projectHook:ProjectHook":
                return new ProjectHook(name, <any>undefined, { urn })
            case "gitlab:index/projectIssue:ProjectIssue":
                return new ProjectIssue(name, <any>undefined, { urn })
            case "gitlab:index/projectLevelMrApprovals:ProjectLevelMrApprovals":
                return new ProjectLevelMrApprovals(name, <any>undefined, { urn })
            case "gitlab:index/projectMembership:ProjectMembership":
                return new ProjectMembership(name, <any>undefined, { urn })
            case "gitlab:index/projectMilestone:ProjectMilestone":
                return new ProjectMilestone(name, <any>undefined, { urn })
            case "gitlab:index/projectMirror:ProjectMirror":
                return new ProjectMirror(name, <any>undefined, { urn })
            case "gitlab:index/projectProtectedEnvironment:ProjectProtectedEnvironment":
                return new ProjectProtectedEnvironment(name, <any>undefined, { urn })
            case "gitlab:index/projectRunnerEnablement:ProjectRunnerEnablement":
                return new ProjectRunnerEnablement(name, <any>undefined, { urn })
            case "gitlab:index/projectShareGroup:ProjectShareGroup":
                return new ProjectShareGroup(name, <any>undefined, { urn })
            case "gitlab:index/projectTag:ProjectTag":
                return new ProjectTag(name, <any>undefined, { urn })
            case "gitlab:index/projectVariable:ProjectVariable":
                return new ProjectVariable(name, <any>undefined, { urn })
            case "gitlab:index/releaseLink:ReleaseLink":
                return new ReleaseLink(name, <any>undefined, { urn })
            case "gitlab:index/repositoryFile:RepositoryFile":
                return new RepositoryFile(name, <any>undefined, { urn })
            case "gitlab:index/runner:Runner":
                return new Runner(name, <any>undefined, { urn })
            case "gitlab:index/serviceExternalWiki:ServiceExternalWiki":
                return new ServiceExternalWiki(name, <any>undefined, { urn })
            case "gitlab:index/serviceGithub:ServiceGithub":
                return new ServiceGithub(name, <any>undefined, { urn })
            case "gitlab:index/serviceJira:ServiceJira":
                return new ServiceJira(name, <any>undefined, { urn })
            case "gitlab:index/serviceMicrosoftTeams:ServiceMicrosoftTeams":
                return new ServiceMicrosoftTeams(name, <any>undefined, { urn })
            case "gitlab:index/servicePipelinesEmail:ServicePipelinesEmail":
                return new ServicePipelinesEmail(name, <any>undefined, { urn })
            case "gitlab:index/serviceSlack:ServiceSlack":
                return new ServiceSlack(name, <any>undefined, { urn })
            case "gitlab:index/systemHook:SystemHook":
                return new SystemHook(name, <any>undefined, { urn })
            case "gitlab:index/tagProtection:TagProtection":
                return new TagProtection(name, <any>undefined, { urn })
            case "gitlab:index/topic:Topic":
                return new Topic(name, <any>undefined, { urn })
            case "gitlab:index/user:User":
                return new User(name, <any>undefined, { urn })
            case "gitlab:index/userCustomAttribute:UserCustomAttribute":
                return new UserCustomAttribute(name, <any>undefined, { urn })
            case "gitlab:index/userSshKey:UserSshKey":
                return new UserSshKey(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gitlab", "index/branch", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/branchProtection", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/clusterAgent", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/clusterAgentToken", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/deployKey", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/deployKeyEnable", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/deployToken", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/group", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupAccessToken", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupBadge", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupCluster", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupCustomAttribute", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupLabel", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupLdapLink", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupMembership", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupProjectFileTemplate", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupShareGroup", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/groupVariable", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/instanceCluster", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/instanceVariable", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/label", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/managedLicense", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/personalAccessToken", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/pipelineSchedule", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/pipelineScheduleVariable", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/pipelineTrigger", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/project", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectAccessToken", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectApprovalRule", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectBadge", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectCluster", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectCustomAttribute", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectEnvironment", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectFreezePeriod", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectHook", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectIssue", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectLevelMrApprovals", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectMembership", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectMilestone", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectMirror", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectProtectedEnvironment", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectRunnerEnablement", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectShareGroup", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectTag", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/projectVariable", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/releaseLink", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/repositoryFile", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/runner", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/serviceExternalWiki", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/serviceGithub", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/serviceJira", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/serviceMicrosoftTeams", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/servicePipelinesEmail", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/serviceSlack", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/systemHook", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/tagProtection", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/topic", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/user", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/userCustomAttribute", _module)
pulumi.runtime.registerResourceModule("gitlab", "index/userSshKey", _module)

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
