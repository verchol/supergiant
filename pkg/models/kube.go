package models

type Kube struct {
	BaseModel

	// belongs_to CloudAccount
	CloudAccount   *CloudAccount `json:"cloud_account,omitempty"`
	CloudAccountID *int64        `json:"cloud_account_id" gorm:"not null;index"`

	// has_many Nodes
	Nodes []*Node `json:"nodes,omitempty"`

	// has_many Apps
	Apps []*App `json:"apps,omitempty"`

	// has_many Entrypoints
	Entrypoints []*Entrypoint `json:"entrypoints,omitempty"`

	// has_many Volumes
	Volumes []*Volume `json:"volumes,omitempty"`

	Name string `json:"name" validate:"nonzero,max=12,regexp=^[a-z]([-a-z0-9]*[a-z0-9])?$" gorm:"not null;unique_index"`

	Username string `json:"username" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`

	Config     *AWSKubeConfig `json:"config" gorm:"-" validate:"nonzero" sg:"store_as_json_in=ConfigJSON"`
	ConfigJSON []byte         `json:"-" gorm:"not null"`

	Ready bool `json:"ready" sg:"readonly" gorm:"index"`
}

type AWSKubeConfig struct {
	Region              string   `json:"region" validate:"nonzero,regexp=^[a-z]{2}-[a-z]+-[0-9]$"`
	AvailabilityZone    string   `json:"availability_zone" validate:"nonzero,regexp=^[a-z]{2}-[a-z]+-[0-9][a-z]$"`
	InstanceTypes       []string `json:"instance_types" validate:"min=1"`
	MasterInstanceType  string   `json:"master_instance_type" validate:"nonzero" sg:"default=m4.large"`
	VPCIPRange          string   `json:"vpc_ip_range" validate:"nonzero" sg:"default=172.20.0.0/16"`
	PublicSubnetIPRange string   `json:"public_subnet_ip_range" validate:"nonzero" sg:"default=172.20.0.0/24"`
	MasterPrivateIP     string   `json:"master_private_ip" validate:"nonzero" sg:"default=172.20.0.9"`

	PrivateKey                    string `json:"private_key,omitempty" sg:"readonly,private"`
	VPCID                         string `json:"vpc_id" sg:"readonly"`
	InternetGatewayID             string `json:"internet_gateway_id" sg:"readonly"`
	PublicSubnetID                string `json:"public_subnet_id" sg:"readonly"`
	RouteTableID                  string `json:"route_table_id" sg:"readonly"`
	RouteTableSubnetAssociationID string `json:"route_table_subnet_association_id" sg:"readonly"`
	ELBSecurityGroupID            string `json:"elb_security_group_id" sg:"readonly"`
	NodeSecurityGroupID           string `json:"node_security_group_id" sg:"readonly"`
	MasterID                      string `json:"master_id" sg:"readonly"`
	MasterPublicIP                string `json:"master_public_ip" sg:"readonly"`
}

func (m *Kube) AutoScalingGroupName(instanceType string) string {
	return m.Name + "-" + instanceType
}

func (m *Kube) AutoScalingGroupNames() (names []string) {
	for _, instanceType := range m.Config.InstanceTypes {
		names = append(names, m.AutoScalingGroupName(instanceType))
	}
	return
}
