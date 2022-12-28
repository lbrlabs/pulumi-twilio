package twilio

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/lbrlabs/pulumi-twilio/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/twilio/terraform-provider-twilio/twilio"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "twilio"
	// modules:
	mainMod        = "index" // the twilio module
	accountsMod    = "accounts"
	apiAccountsMod = "apiaccounts"
	autopilotAssistantsMod = "autopilotassistants/v1"
	bulkexportsMod = "bulkexports/v1"
	chatMod = "chat/v1"
	chatv2Mod = "chat/v2"
	conversationsMod = "conversations/v1"
	eventsMod = "events/v1"
	flexMod = "flex/v1"
	ipMessagingMod = "ipmessaging/v1"
	ipMessagingv2Mod = "ipmessaging/v2"
)

// twilioMember manufactures a type token for the Twilio package and the given module and type.
func twilioMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// scalewayType manufactures a type token for the Twilio package and the given module and type.
func twilioType(mod string, typ string) tokens.Type {
	return tokens.Type(twilioMember(mod, typ))
}

// twilioDataSource manufactures a standard resource token given a module and resource name.
// It automatically uses the Twilio package and names the file by simply lower casing the data
// source's first character.
func twilioDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return twilioMember(mod+"/"+fn, res)
}

// twilioResource manufactures a standard resource token given a module and resource name.
// It automatically uses the Twilio package and names the file by simply lower casing the resource's
// first character.
func twilioResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return twilioType(mod+"/"+fn, res)
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(twilio.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "twilio",
		DisplayName:       "Twilio",
		Publisher:         "lbrlabs",
		LogoURL:           "",
		PluginDownloadURL: "github://api.github.com/lbrlabs",
		Description:       "A Pulumi package for creating and managing Twilio cloud resources.",
		Keywords:          []string{"pulumi", "twilio", "category/cloud"},
		License:           "Apache-2.0",
		Homepage:          "https://leebriggs.co.uk/projects#pulumi-twilio",
		Repository:        "https://github.com/lbrlabs/pulumi-twilio",
		GitHubOrg:         "twilio",
		Config:            map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"twilio_accounts_credentials_aws_v1":           {Tok: twilioResource(accountsMod, "CredentialsAwsV1")},
			"twilio_accounts_credentials_public_keys_v1":   {Tok: twilioResource(accountsMod, "CredentialsPublicKeysV1")},
			"twilio_api_accounts_addresses":                {Tok: twilioResource(apiAccountsMod, "Addresses")},
			"twilio_api_accounts_applications":             {Tok: twilioResource(apiAccountsMod, "Applications")},
			"twilio_api_accounts_calls":                    {Tok: twilioResource(apiAccountsMod, "Calls")},
			"twilio_api_accounts_calls_feedback_summary":   {Tok: twilioResource(apiAccountsMod, "CallsFeedbackSummary")},
			"twilio_api_accounts_calls_recordings":         {Tok: twilioResource(apiAccountsMod, "CallsRecordings")},
			"twilio_api_accounts_conferences_participants": {Tok: twilioResource(apiAccountsMod, "ConferencesParticipans")},
			"twilio_api_accounts_incoming_phone_numbers":   {Tok: twilioResource(apiAccountsMod, "IncomingPhoneNumbers")},
			"twilio_api_accounts_incoming_phone_numbers_assigned_add_ons": {
				Tok: twilioResource(apiAccountsMod, "IncomingPhoneNumbersAssignedAddons"),
			},
			"twilio_api_accounts_keys": {
				Tok: twilioResource(apiAccountsMod, "Keys"),
			},
			"twilio_api_accounts_messages": {
				Tok: twilioResource(apiAccountsMod, "Messages"),
			},
			"twilio_api_accounts_queues": {
				Tok: twilioResource(apiAccountsMod, "Queues"),
			},
			"twilio_api_accounts_signing_keys": {
				Tok: twilioResource(apiAccountsMod, "SigningKeys"),
			},
			"twilio_api_accounts_sip_credential_lists": {
				Tok: twilioResource(apiAccountsMod, "SipCredentialsLists"),
			},
			"twilio_api_accounts_sip_credential_lists_credentials": {
				Tok: twilioResource(apiAccountsMod, "SipCredentialsListsCredentials"),
			},
			"twilio_api_accounts_sip_domains": {
				Tok: twilioResource(apiAccountsMod, "SipDomains"),
			},
			"twilio_api_accounts_sip_domains_auth_calls_credential_list_mappings": {
				Tok: twilioResource(apiAccountsMod, "SipDomainsAuthCallsCredentialListMappings"),
			},
			"twilio_api_accounts_sip_domains_auth_calls_ip_access_control_list_mappings": {
				Tok: twilioResource(apiAccountsMod, "SipDomainsAuthCallsIpAccessControlListMappings"),
			},
			"twilio_api_accounts_sip_domains_auth_registrations_credential_list_mappings": {
				Tok: twilioResource(apiAccountsMod, "SipDomainsAuthRegistrationsCredentialListMappings"),
			},
			"twilio_api_accounts_sip_domains_credential_list_mappings": {
				Tok: twilioResource(apiAccountsMod, "SipDomainsCredentialListMappings"),
			},
			"twilio_api_accounts_sip_domains_ip_access_control_list_mappings": {
				Tok: twilioResource(apiAccountsMod, "SipDomainsIpAccessControlListMappings"),
			},
			"twilio_api_accounts_sip_ip_access_control_lists": {
				Tok: twilioResource(apiAccountsMod, "SipIpAccessControlLists"),
			},
			"twilio_api_accounts_sip_ip_access_control_lists_ip_addresses": {
				Tok: twilioResource(apiAccountsMod, "SipIpAccessControlListsIpAddresses"),
			},
			"twilio_api_accounts_usage_triggers": {
				Tok: twilioResource(apiAccountsMod, "UsageTriggers"),
			},
			"twilio_autopilot_assistants_field_types_field_values_v1": {
				Tok: twilioResource(autopilotAssistantsMod, "FieldTypesFieldValues"),
			},
			"twilio_autopilot_assistants_field_types_v1": {
				Tok: twilioResource(autopilotAssistantsMod, "FieldTypes"),
			},
			"twilio_autopilot_assistants_model_builds_v1": {
				Tok: twilioResource(autopilotAssistantsMod, "ModelBuilds"),
			},
			"twilio_autopilot_assistants_queries_v1": {
				Tok: twilioResource(autopilotAssistantsMod, "Queries"),
			},
			"twilio_autopilot_assistants_tasks_fields_v1": {
				Tok: twilioResource(autopilotAssistantsMod, "TasksFields"),
			},
			"twilio_autopilot_assistants_tasks_samples_v1": {
				Tok: twilioResource(autopilotAssistantsMod, "TasksSamples"),
			},
			"twilio_autopilot_assistants_tasks_v1": {
				Tok: twilioResource(autopilotAssistantsMod, "Tasks"),
			},
			"twilio_autopilot_assistants_v1": {
				Tok: twilioResource(autopilotAssistantsMod, "Assistants"),
			},
			"twilio_autopilot_assistants_webhooks_v1": {
				Tok: twilioResource(autopilotAssistantsMod, "Webhooks"),
			},
			"twilio_bulkexports_exports_jobs_v1": {
				Tok: twilioResource(bulkexportsMod, "ExportsJobs"),
			},
			"twilio_chat_credentials_v1": {
				Tok: twilioResource(chatMod, "Credentials"),
			},
			// "twilio_chat_credentials_v2": {
			// 	Tok: twilioResource(chatv2Mod, "Credentials"),
			// },
			"twilio_chat_services_channels_invites_v1": {
				Tok: twilioResource(chatMod, "ServiceChannelsInvites"),
			},
			"twilio_chat_services_channels_invites_v2": {
				Tok: twilioResource(chatv2Mod, "ServicesChannelsInvites"),
			},
			"twilio_chat_services_channels_members_v1": {
				Tok: twilioResource(chatMod, "ServicesChannelsMembers"),
			},
			"twilio_chat_services_channels_members_v2": {
				Tok: twilioResource(chatv2Mod, "ServicesChannelsMembers"),
			},
			"twilio_chat_services_channels_messages_v1": {
				Tok: twilioResource(chatMod, "ServicesChannelsMessages"),
			},
			"twilio_chat_services_channels_messages_v2": {
				Tok: twilioResource(chatv2Mod, "ServicesChannelsMessages"),
			},
			"twilio_chat_services_channels_v1": {
				Tok: twilioResource(chatMod, "ServicesChannels"),
			},
			"twilio_chat_services_channels_v2": {
				Tok: twilioResource(chatv2Mod, "ServicesChannels"),
			},
			"twilio_chat_services_channels_webhooks_v2": {
				Tok: twilioResource(chatv2Mod, "ServicesChannelsWebhooks"),
			},
			"twilio_chat_services_roles_v1": {
				Tok: twilioResource(chatMod, "ServicesRoles"),
			},
			"twilio_chat_services_roles_v2": {
				Tok: twilioResource(chatv2Mod, "ServicesRoles"),
			},
			"twilio_chat_services_users_v1": {
				Tok: twilioResource(chatMod, "ServicesUsers"),
			},
			"twilio_chat_services_users_v2": {
				Tok: twilioResource(chatv2Mod, "ServicesUsers"),
			},
			"twilio_chat_services_v1": {
				Tok: twilioResource(chatMod, "Services"),
			},
			"twilio_chat_services_v2": {
				Tok: twilioResource(chatv2Mod, "Services"),
			},
			"twilio_conversations_conversations_messages_v1": {
				Tok: twilioResource(conversationsMod, "ConversationsMessages"),
			},
			"twilio_conversations_conversations_participants_v1": {
				Tok: twilioResource(conversationsMod, "ConversationsParticipants"),
			},
			"twilio_conversations_conversations_v1": {
				Tok: twilioResource(conversationsMod, "Conversations"),
			},
			"twilio_conversations_conversations_webhooks_v1": {
				Tok: twilioResource(conversationsMod, "ConversationsWebhooks"),
			},
			"twilio_conversations_credentials_v1": {
				Tok: twilioResource(conversationsMod, "Credentials"),
			},
			"twilio_conversations_roles_v1": {
				Tok: twilioResource(conversationsMod, "Roles"),
			},
			"twilio_conversations_services_conversations_messages_v1": {
				Tok: twilioResource(conversationsMod, "ServicesConversationsMessages"),
			},
			"twilio_conversations_services_conversations_participants_v1": {
				Tok: twilioResource(conversationsMod, "ServicesConversationsParticipants"),
			},
			"twilio_conversations_services_conversations_v1": {
				Tok: twilioResource(conversationsMod, "ServicesConversations"),
			},
			"twilio_conversations_services_conversations_webhooks_v1": {
				Tok: twilioResource(conversationsMod, "ServicesConversationsWebhooks"),
			},
			"twilio_conversations_services_roles_v1": {
				Tok: twilioResource(conversationsMod, "ServicesRoles"),
			},
			"twilio_conversations_services_users_v1": {
				Tok: twilioResource(conversationsMod, "ServicesUsers"),
			},
			"twilio_conversations_services_v1": {
				Tok: twilioResource(conversationsMod, "Services"),
			},
			"twilio_conversations_users_v1": {
				Tok: twilioResource(conversationsMod, "Users"),
			},
			"twilio_events_sinks_v1": {
				Tok: twilioResource(eventsMod, "Sinks"),
			},
			"twilio_events_subscriptions_subscribed_events_v1": {
				Tok: twilioResource(eventsMod, "SubscriptionsSubscribedEvents"),
			},
			"twilio_events_subscriptions_v1": {
				Tok: twilioResource(eventsMod, "Subscriptions"),
			},
			"twilio_flex_channels_v1": {
				Tok: twilioResource(flexMod, "Channels"),
			},
			"twilio_flex_flex_flows_v1": {
				Tok: twilioResource(flexMod, "FlexFlows"),
			},
			"twilio_flex_web_channels_v1": {
				Tok: twilioResource(flexMod, "WebChannels"),
			},
			"twilio_ip_messaging_credentials_v1": {
				Tok: twilioResource(ipMessagingMod, "Credentials"),
			},
			"twilio_ip_messaging_credentials_v2": {
				Tok: twilioResource(ipMessagingv2Mod, "Credentials"),
			},
			"twilio_ip_messaging_services_channels_invites_v1": {
				Tok: twilioResource(ipMessagingMod, "ServicesChannelsInvites"),
			},
			"twilio_ip_messaging_services_channels_invites_v2": {
				Tok: twilioResource(ipMessagingv2Mod, "ServicesChannelsInvites"),
			},

		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
