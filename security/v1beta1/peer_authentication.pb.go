// Copyright 2020 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: security/v1beta1/peer_authentication.proto

// $schema: istio.security.v1beta1.PeerAuthentication
// $title: PeerAuthentication
// $description: Peer authentication configuration for workloads.
// $location: https://istio.io/docs/reference/config/security/peer_authentication.html
// $aliases: [/docs/reference/config/security/v1beta1/peer_authentication]

package v1beta1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1beta1 "istio.io/api/type/v1beta1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PeerAuthentication_MutualTLS_Mode int32

const (
	// Inherit from parent, if has one. Otherwise treated as PERMISSIVE.
	PeerAuthentication_MutualTLS_UNSET PeerAuthentication_MutualTLS_Mode = 0
	// Connection is not tunneled.
	PeerAuthentication_MutualTLS_DISABLE PeerAuthentication_MutualTLS_Mode = 1
	// Connection can be either plaintext or mTLS tunnel.
	PeerAuthentication_MutualTLS_PERMISSIVE PeerAuthentication_MutualTLS_Mode = 2
	// Connection is an mTLS tunnel (TLS with client cert must be presented).
	PeerAuthentication_MutualTLS_STRICT PeerAuthentication_MutualTLS_Mode = 3
)

// Enum value maps for PeerAuthentication_MutualTLS_Mode.
var (
	PeerAuthentication_MutualTLS_Mode_name = map[int32]string{
		0: "UNSET",
		1: "DISABLE",
		2: "PERMISSIVE",
		3: "STRICT",
	}
	PeerAuthentication_MutualTLS_Mode_value = map[string]int32{
		"UNSET":      0,
		"DISABLE":    1,
		"PERMISSIVE": 2,
		"STRICT":     3,
	}
)

func (x PeerAuthentication_MutualTLS_Mode) Enum() *PeerAuthentication_MutualTLS_Mode {
	p := new(PeerAuthentication_MutualTLS_Mode)
	*p = x
	return p
}

func (x PeerAuthentication_MutualTLS_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeerAuthentication_MutualTLS_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_security_v1beta1_peer_authentication_proto_enumTypes[0].Descriptor()
}

func (PeerAuthentication_MutualTLS_Mode) Type() protoreflect.EnumType {
	return &file_security_v1beta1_peer_authentication_proto_enumTypes[0]
}

func (x PeerAuthentication_MutualTLS_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeerAuthentication_MutualTLS_Mode.Descriptor instead.
func (PeerAuthentication_MutualTLS_Mode) EnumDescriptor() ([]byte, []int) {
	return file_security_v1beta1_peer_authentication_proto_rawDescGZIP(), []int{0, 0, 0}
}

// PeerAuthentication defines how traffic will be tunneled (or not) to the sidecar.
//
// Examples:
//
// Policy to allow mTLS traffic for all workloads under namespace `foo`:
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: PeerAuthentication
// metadata:
//   name: default
//   namespace: foo
// spec:
//   mtls:
//     mode: STRICT
// ```
// For mesh level, put the policy in root-namespace according to your Istio installation.
//
// Policies to allow both mTLS & plaintext traffic for all workloads under namespace `foo`, but
// require mTLS for workload `finance`.
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: PeerAuthentication
// metadata:
//   name: default
//   namespace: foo
// spec:
//   mtls:
//     mode: PERMISSIVE
// ---
// apiVersion: security.istio.io/v1beta1
// kind: PeerAuthentication
// metadata:
//   name: default
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: finance
//   mtls:
//     mode: STRICT
// ```
// Policy to allow mTLS strict for all workloads, but leave port 8080 to
// plaintext:
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: PeerAuthentication
// metadata:
//   name: default
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: finance
//   mtls:
//     mode: STRICT
//   portLevelMtls:
//     8080:
//       mode: DISABLE
// ```
// Policy to inherit mTLS mode from namespace (or mesh) settings, and overwrite
// settings for port 8080
// ```yaml
// apiVersion: security.istio.io/v1beta1
// kind: PeerAuthentication
// metadata:
//   name: default
//   namespace: foo
// spec:
//   selector:
//     matchLabels:
//       app: finance
//   mtls:
//     mode: UNSET
//   portLevelMtls:
//     8080:
//       mode: DISABLE
// ```
//
// <!-- crd generation tags
// +cue-gen:PeerAuthentication:groupName:security.istio.io
// +cue-gen:PeerAuthentication:version:v1beta1
// +cue-gen:PeerAuthentication:storageVersion
// +cue-gen:PeerAuthentication:annotations:helm.sh/resource-policy=keep
// +cue-gen:PeerAuthentication:labels:app=istio-pilot,chart=istio,istio=security,heritage=Tiller,release=istio
// +cue-gen:PeerAuthentication:subresource:status
// +cue-gen:PeerAuthentication:scope:Namespaced
// +cue-gen:PeerAuthentication:resource:categories=istio-io,security-istio-io,shortNames=pa
// +cue-gen:PeerAuthentication:preserveUnknownFields:false
// +cue-gen:PeerAuthentication:printerColumn:name=Mode,type=string,JSONPath=.spec.mtls.mode,description="Defines the mTLS mode used for peer authentication."
// +cue-gen:PeerAuthentication:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=security.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type PeerAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The selector determines the workloads to apply the ChannelAuthentication on.
	// If not set, the policy will be applied to all workloads in the same namespace as the policy.
	Selector *v1beta1.WorkloadSelector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Mutual TLS settings for workload. If not defined, inherit from parent.
	Mtls *PeerAuthentication_MutualTLS `protobuf:"bytes,2,opt,name=mtls,proto3" json:"mtls,omitempty"`
	// Port specific mutual TLS settings. These only apply when a workload selector
	// is specified.
	PortLevelMtls map[uint32]*PeerAuthentication_MutualTLS `protobuf:"bytes,3,rep,name=port_level_mtls,json=portLevelMtls,proto3" json:"port_level_mtls,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PeerAuthentication) Reset() {
	*x = PeerAuthentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_v1beta1_peer_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerAuthentication) ProtoMessage() {}

func (x *PeerAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1beta1_peer_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerAuthentication.ProtoReflect.Descriptor instead.
func (*PeerAuthentication) Descriptor() ([]byte, []int) {
	return file_security_v1beta1_peer_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *PeerAuthentication) GetSelector() *v1beta1.WorkloadSelector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *PeerAuthentication) GetMtls() *PeerAuthentication_MutualTLS {
	if x != nil {
		return x.Mtls
	}
	return nil
}

func (x *PeerAuthentication) GetPortLevelMtls() map[uint32]*PeerAuthentication_MutualTLS {
	if x != nil {
		return x.PortLevelMtls
	}
	return nil
}

// Mutual TLS settings.
type PeerAuthentication_MutualTLS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the mTLS mode used for peer authentication.
	Mode PeerAuthentication_MutualTLS_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=istio.security.v1beta1.PeerAuthentication_MutualTLS_Mode" json:"mode,omitempty"`
}

func (x *PeerAuthentication_MutualTLS) Reset() {
	*x = PeerAuthentication_MutualTLS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_v1beta1_peer_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerAuthentication_MutualTLS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerAuthentication_MutualTLS) ProtoMessage() {}

func (x *PeerAuthentication_MutualTLS) ProtoReflect() protoreflect.Message {
	mi := &file_security_v1beta1_peer_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerAuthentication_MutualTLS.ProtoReflect.Descriptor instead.
func (*PeerAuthentication_MutualTLS) Descriptor() ([]byte, []int) {
	return file_security_v1beta1_peer_authentication_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PeerAuthentication_MutualTLS) GetMode() PeerAuthentication_MutualTLS_Mode {
	if x != nil {
		return x.Mode
	}
	return PeerAuthentication_MutualTLS_UNSET
}

var File_security_v1beta1_peer_authentication_proto protoreflect.FileDescriptor

var file_security_v1beta1_peer_authentication_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1b, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x98, 0x04, 0x0a, 0x12, 0x50, 0x65, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x73, 0x74,
	0x69, 0x6f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x48, 0x0a, 0x04, 0x6d, 0x74,
	0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x54, 0x4c, 0x53, 0x52, 0x04,
	0x6d, 0x74, 0x6c, 0x73, 0x12, 0x65, 0x0a, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x5f, 0x6d, 0x74, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x4d, 0x74, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x70, 0x6f,
	0x72, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x74, 0x6c, 0x73, 0x1a, 0x96, 0x01, 0x0a, 0x09,
	0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x54, 0x4c, 0x53, 0x12, 0x4d, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x65, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x54, 0x4c, 0x53, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x45, 0x52, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49,
	0x43, 0x54, 0x10, 0x03, 0x1a, 0x76, 0x0a, 0x12, 0x50, 0x6f, 0x72, 0x74, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x4d, 0x74, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4a, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x54, 0x4c,
	0x53, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1f, 0x5a, 0x1d,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_security_v1beta1_peer_authentication_proto_rawDescOnce sync.Once
	file_security_v1beta1_peer_authentication_proto_rawDescData = file_security_v1beta1_peer_authentication_proto_rawDesc
)

func file_security_v1beta1_peer_authentication_proto_rawDescGZIP() []byte {
	file_security_v1beta1_peer_authentication_proto_rawDescOnce.Do(func() {
		file_security_v1beta1_peer_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_v1beta1_peer_authentication_proto_rawDescData)
	})
	return file_security_v1beta1_peer_authentication_proto_rawDescData
}

var file_security_v1beta1_peer_authentication_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_security_v1beta1_peer_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_security_v1beta1_peer_authentication_proto_goTypes = []interface{}{
	(PeerAuthentication_MutualTLS_Mode)(0), // 0: istio.security.v1beta1.PeerAuthentication.MutualTLS.Mode
	(*PeerAuthentication)(nil),             // 1: istio.security.v1beta1.PeerAuthentication
	(*PeerAuthentication_MutualTLS)(nil),   // 2: istio.security.v1beta1.PeerAuthentication.MutualTLS
	nil,                                    // 3: istio.security.v1beta1.PeerAuthentication.PortLevelMtlsEntry
	(*v1beta1.WorkloadSelector)(nil),       // 4: istio.type.v1beta1.WorkloadSelector
}
var file_security_v1beta1_peer_authentication_proto_depIdxs = []int32{
	4, // 0: istio.security.v1beta1.PeerAuthentication.selector:type_name -> istio.type.v1beta1.WorkloadSelector
	2, // 1: istio.security.v1beta1.PeerAuthentication.mtls:type_name -> istio.security.v1beta1.PeerAuthentication.MutualTLS
	3, // 2: istio.security.v1beta1.PeerAuthentication.port_level_mtls:type_name -> istio.security.v1beta1.PeerAuthentication.PortLevelMtlsEntry
	0, // 3: istio.security.v1beta1.PeerAuthentication.MutualTLS.mode:type_name -> istio.security.v1beta1.PeerAuthentication.MutualTLS.Mode
	2, // 4: istio.security.v1beta1.PeerAuthentication.PortLevelMtlsEntry.value:type_name -> istio.security.v1beta1.PeerAuthentication.MutualTLS
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_security_v1beta1_peer_authentication_proto_init() }
func file_security_v1beta1_peer_authentication_proto_init() {
	if File_security_v1beta1_peer_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_security_v1beta1_peer_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerAuthentication); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_security_v1beta1_peer_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerAuthentication_MutualTLS); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_security_v1beta1_peer_authentication_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_v1beta1_peer_authentication_proto_goTypes,
		DependencyIndexes: file_security_v1beta1_peer_authentication_proto_depIdxs,
		EnumInfos:         file_security_v1beta1_peer_authentication_proto_enumTypes,
		MessageInfos:      file_security_v1beta1_peer_authentication_proto_msgTypes,
	}.Build()
	File_security_v1beta1_peer_authentication_proto = out.File
	file_security_v1beta1_peer_authentication_proto_rawDesc = nil
	file_security_v1beta1_peer_authentication_proto_goTypes = nil
	file_security_v1beta1_peer_authentication_proto_depIdxs = nil
}
