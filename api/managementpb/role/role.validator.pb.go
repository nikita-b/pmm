// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/role/role.proto

package rolev1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

func (this *CreateRoleRequest) Validate() error {
	if this.Title == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must not be an empty string`, this.Title))
	}
	return nil
}

func (this *CreateRoleResponse) Validate() error {
	return nil
}

func (this *UpdateRoleRequest) Validate() error {
	if !(this.RoleId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleId", fmt.Errorf(`value '%v' must be greater than '0'`, this.RoleId))
	}
	if this.Title == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must not be an empty string`, this.Title))
	}
	return nil
}

func (this *UpdateRoleResponse) Validate() error {
	return nil
}

func (this *DeleteRoleRequest) Validate() error {
	if !(this.RoleId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleId", fmt.Errorf(`value '%v' must be greater than '0'`, this.RoleId))
	}
	return nil
}

func (this *DeleteRoleResponse) Validate() error {
	return nil
}

func (this *GetRoleRequest) Validate() error {
	if !(this.RoleId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleId", fmt.Errorf(`value '%v' must be greater than '0'`, this.RoleId))
	}
	return nil
}

func (this *GetRoleResponse) Validate() error {
	return nil
}

func (this *SetDefaultRoleRequest) Validate() error {
	if !(this.RoleId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("RoleId", fmt.Errorf(`value '%v' must be greater than '0'`, this.RoleId))
	}
	return nil
}

func (this *SetDefaultRoleResponse) Validate() error {
	return nil
}

func (this *AssignRolesRequest) Validate() error {
	if !(this.UserId > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserId", fmt.Errorf(`value '%v' must be greater than '0'`, this.UserId))
	}
	return nil
}

func (this *AssignRolesResponse) Validate() error {
	return nil
}

func (this *ListRolesRequest) Validate() error {
	return nil
}

func (this *ListRolesResponse) Validate() error {
	for _, item := range this.Roles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Roles", err)
			}
		}
	}
	return nil
}

func (this *ListRolesResponse_RoleData) Validate() error {
	return nil
}
