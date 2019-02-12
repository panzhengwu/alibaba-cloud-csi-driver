package ess

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ScalingConfiguration is a nested struct in ess response
type ScalingConfiguration struct {
	ScalingConfigurationId      string         `json:"ScalingConfigurationId" xml:"ScalingConfigurationId"`
	ScalingConfigurationName    string         `json:"ScalingConfigurationName" xml:"ScalingConfigurationName"`
	ScalingGroupId              string         `json:"ScalingGroupId" xml:"ScalingGroupId"`
	InstanceName                string         `json:"InstanceName" xml:"InstanceName"`
	ImageId                     string         `json:"ImageId" xml:"ImageId"`
	ImageName                   string         `json:"ImageName" xml:"ImageName"`
	HostName                    string         `json:"HostName" xml:"HostName"`
	InstanceType                string         `json:"InstanceType" xml:"InstanceType"`
	Cpu                         int            `json:"Cpu" xml:"Cpu"`
	Memory                      int            `json:"Memory" xml:"Memory"`
	InstanceGeneration          string         `json:"InstanceGeneration" xml:"InstanceGeneration"`
	SecurityGroupId             string         `json:"SecurityGroupId" xml:"SecurityGroupId"`
	IoOptimized                 string         `json:"IoOptimized" xml:"IoOptimized"`
	InternetChargeType          string         `json:"InternetChargeType" xml:"InternetChargeType"`
	InternetMaxBandwidthIn      int            `json:"InternetMaxBandwidthIn" xml:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut     int            `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	SystemDiskCategory          string         `json:"SystemDiskCategory" xml:"SystemDiskCategory"`
	SystemDiskSize              int            `json:"SystemDiskSize" xml:"SystemDiskSize"`
	LifecycleState              string         `json:"LifecycleState" xml:"LifecycleState"`
	CreationTime                string         `json:"CreationTime" xml:"CreationTime"`
	LoadBalancerWeight          int            `json:"LoadBalancerWeight" xml:"LoadBalancerWeight"`
	UserData                    string         `json:"UserData" xml:"UserData"`
	KeyPairName                 string         `json:"KeyPairName" xml:"KeyPairName"`
	RamRoleName                 string         `json:"RamRoleName" xml:"RamRoleName"`
	DeploymentSetId             string         `json:"DeploymentSetId" xml:"DeploymentSetId"`
	SecurityEnhancementStrategy string         `json:"SecurityEnhancementStrategy" xml:"SecurityEnhancementStrategy"`
	SpotStrategy                string         `json:"SpotStrategy" xml:"SpotStrategy"`
	PasswordInherit             bool           `json:"PasswordInherit" xml:"PasswordInherit"`
	InstanceTypes               InstanceTypes  `json:"InstanceTypes" xml:"InstanceTypes"`
	DataDisks                   DataDisks      `json:"DataDisks" xml:"DataDisks"`
	Tags                        Tags           `json:"Tags" xml:"Tags"`
	SpotPriceLimit              SpotPriceLimit `json:"SpotPriceLimit" xml:"SpotPriceLimit"`
}
