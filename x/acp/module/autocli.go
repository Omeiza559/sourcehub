package acp

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/sourcenetwork/sourcehub/api/sourcehub/acp"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "Policy",
					Use:            "policy [id]",
					Short:          "Queries for a Policy by its ID",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "PolicyIds",
					Use:            "policy-i-ds",
					Short:          "Lists Registered Policies IDs",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "FilterRelationships",
					Use:            "filter-relationships [policy-id]",
					Short:          "Filters Relationships within a Policy",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "policyId"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePolicy",
					Use:            "create-policy [policy]",
					Short:          "Send a CreatePolicy tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "policy"}},
				},
				{
					RpcMethod:      "SetRelationship",
					Use:            "set-relationship [policy]",
					Short:          "Send a SetRelationship tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "policy"}},
				},
				{
					RpcMethod:      "DeleteRelationship",
					Use:            "delete-relationship [policy-id]",
					Short:          "Send a DeleteRelationship tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "policyId"}},
				},
				{
					RpcMethod:      "RegisterObject",
					Use:            "register-object [policy-id]",
					Short:          "Send a RegisterObject tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "policyId"}},
				},
				{
					RpcMethod:      "UnregisterObject",
					Use:            "unregister-object [policy-id]",
					Short:          "Send a UnregisterObject tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "policyId"}},
				},
				{
					RpcMethod:      "CheckAccess",
					Use:            "check-access [policy-id]",
					Short:          "Send a CheckAccess tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "policyId"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
