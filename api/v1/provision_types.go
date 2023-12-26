/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ProvisionSpec defines the desired state of Provision
type ProvisionSpec struct {
	AccessControlGroupNoListN                     string `json:"accessControlGroupNoList,omitempty"`
	AssociateWithPublicIp                         bool   `json:"associateWithPublicIp,omitempty"`
	BlockDevicePartitionMountPoint                string `json:"blockDevicePartitionMountPoint,omitempty"`
	BlockDevicePartitionSize                      string `json:"blockDevicePartitionSize,omitempty"`
	BlockStorageMappingBlockStorageName           string `json:"blockStorageMappingBlockStorageName,omitempty"`
	BlockStorageMappingBlockStorageSize           string `json:"blockStorageMappingBlockStorageSize,omitempty"`
	BlockStorageMappingBlockStorageVolumeTypeCode string `json:"blockStorageMappingBlockStorageVolumeTypeCode,omitempty"`
	BlockStorageMappingEncrypted                  string `json:"blockStorageMappingEncrypted,omitempty"`
	BlockStorageMappingOrder                      int    `json:"blockStorageMappingList,omitempty"`
	BlockStorageMappingSnapshotInstanceNo         string `json:"blockStorageMappingSnapshotInstanceNo,omitempty"`
	FeeSystemTypeCode                             string `json:"feeSystemTypeCode,omitempty"`
	InitScriptNo                                  string `json:"initScriptNo,omitempty"`
	IsEncryptedBaseBlockStorageVolume             bool   `json:"isEncryptedBaseBlockStorageVolume,omitempty"`
	IsProtectServerTermination                    bool   `json:"isProtectServerTermination,omitempty"`
	LoginKeyName                                  string `json:"loginKeyName,omitempty"`
	MemberServerImageInstanceNo                   string `json:"memberServerImageInstanceNo,omitempty"`
	NetworkInterfaceIp                            string `json:"networkInterfaceIp,omitempty"`
	NetworkInterfaceNo                            string `json:"networkInterfaceNo,omitempty"`
	NetworkInterfaceOrder                         int    `json:"networkInterfaceList,omitempty"`
	NetworkInterfaceSubnetNo                      string `json:"networkInterfaceSubnetNo,omitempty"`
	PlacementGroupNo                              string `json:"placementGroupNo,omitempty"`
	Verb                                          string `json:"verb,omitempty"`
	RAIDTypeName                                  string `json:"raidTypeName,omitempty"`
	ResponseFormatType                            string `json:"responseFormatType,omitempty"`
	ServerCreateCount                             int    `json:"serverCreateCount,omitempty"`
	ServerCreateStartNo                           int    `json:"serverCreateStartNo,omitempty"`
	ServerDescription                             string `json:"serverDescription,omitempty"`
	ServerImageNo                                 string `json:"serverImageNo,omitempty"`
	ServerImageProductCode                        string `json:"serverImageProductCode,omitempty"`
	ServerName                                    string `json:"serverName,omitempty"`
	ServerProductCode                             string `json:"serverProductCode,omitempty"`
	ServerSpecCode                                string `json:"serverSpecCode,omitempty"`
	SubnetNo                                      string `json:"subnetNo,omitempty"`
	VpcNo                                         string `json:"vpcNo,omitempty"`
}

// ProvisionStatus defines the observed state of Provision
type ProvisionStatus struct {
	Version       int    `json:"version,omitempty"`
	Errors        int    `json:"errors,omitempty"`
	Warnings      int    `json:"warnings,omitempty"`
	FailureReason string `json:"failureReason,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Provision is the Schema for the provisions API
type Provision struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProvisionSpec   `json:"spec,omitempty"`
	Status ProvisionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ProvisionList contains a list of Provision
type ProvisionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Provision `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Provision{}, &ProvisionList{})
}
