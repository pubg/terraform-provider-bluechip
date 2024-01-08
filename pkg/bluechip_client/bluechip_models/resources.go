package bluechip_models

var _ ClusterApiResource[EmptySpec] = &Namespace{}

type Namespace struct {
	BaseResponse `json:"-"`

	*TypeMeta          `json:",inline"`
	*MetadataContainer `json:",inline"`
}

func (n Namespace) GetSpec() EmptySpec {
	return EmptySpec{}
}

func (n Namespace) SetSpec(_ EmptySpec) {
}

var _ ClusterApiResource[VendorSpec] = &Vendor{}

type Vendor struct {
	BaseResponse `json:"-"`

	*TypeMeta                  `json:",inline"`
	*MetadataContainer         `json:",inline"`
	*SpecContainer[VendorSpec] `json:",inline"`
}

type VendorSpec struct {
	BaseSpec `json:"-"`

	DisplayName string   `json:"displayName"`
	CodeName    string   `json:"codeName"`
	ShortName   string   `json:"shortName"`
	Regions     []string `json:"regions"`
}

var _ NamespacedApiResource[ClusterSpec] = &Cluster{}

type Cluster struct {
	BaseResponse `json:"-"`

	*TypeMeta                   `json:",inline"`
	*MetadataContainer          `json:",inline"`
	*SpecContainer[ClusterSpec] `json:",inline"`
}

type ClusterSpec struct {
	BaseSpec `json:"-"`

	Project          string                `json:"project"`
	Environment      string                `json:"environment"`
	OrganizationUnit string                `json:"organizationUnit"`
	Platform         string                `json:"platform"`
	Pubg             *ClusterSpecPubg      `json:"pubg"`
	Vendor           ClusterSpecVendor     `json:"vendor"`
	Kubernetes       ClusterSpecKubernetes `json:"kubernetes"`
}

type ClusterSpecPubg struct {
	Infra string `json:"infra"`
	Site  string `json:"site"`
}

type ClusterSpecVendor struct {
	Name      string `json:"name"`
	AccountId string `json:"accountId"`
	Engine    string `json:"engine"`
	Region    string `json:"region"`
}

type ClusterSpecKubernetes struct {
	Endpoint string `json:"endpoint"`
	CaCert   string `json:"caCert"`
	SaIssuer string `json:"saIssuer"`
	Version  string `json:"version"`
}

var _ NamespacedApiResource[CidrSpec] = &Cidr{}

type Cidr struct {
	BaseResponse `json:"-"`

	*TypeMeta                `json:",inline"`
	*MetadataContainer       `json:",inline"`
	*SpecContainer[CidrSpec] `json:",inline"`
}

type CidrSpec struct {
	BaseSpec `json:"-"`

	Ipv4Cidrs []string `json:"ipv4Cidrs"`
	Ipv6Cidrs []string `json:"ipv6Cidrs"`
}

var _ NamespacedApiResource[AccountSpec] = &Account{}

type Account struct {
	BaseResponse `json:"-"`

	*TypeMeta                   `json:",inline"`
	*MetadataContainer          `json:",inline"`
	*SpecContainer[AccountSpec] `json:",inline"`
}

type AccountSpec struct {
	BaseSpec `json:"-"`

	AccountId   string   `json:"accountId"`
	DisplayName string   `json:"displayName"`
	Description string   `json:"description"`
	Alias       string   `json:"alias"`
	Vendor      string   `json:"vendor"`
	Regions     []string `json:"regions"`
}

var _ NamespacedApiResource[ImageSpec] = &Image{}

type Image struct {
	BaseResponse `json:"-"`

	*TypeMeta                 `json:",inline"`
	*MetadataContainer        `json:",inline"`
	*SpecContainer[ImageSpec] `json:",inline"`
}

type ImageSpec struct {
	BaseSpec `json:"-"`

	App        string `json:"app"`
	Timestamp  int    `json:"timestamp"`
	CommitHash string `json:"commitHash"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Branch     string `json:"branch"`
}

var _ ClusterApiResource[UserSpec] = &User{}

type User struct {
	BaseResponse `json:"-"`

	*TypeMeta                `json:",inline"`
	*MetadataContainer       `json:",inline"`
	*SpecContainer[UserSpec] `json:",inline"`
}

type UserSpec struct {
	BaseSpec `json:"-"`

	Password   string            `json:"password"`
	Groups     []string          `json:"groups,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

var _ ClusterApiResource[OidcAuthSpec] = &OidcAuth{}

type OidcAuth struct {
	BaseResponse `json:"-"`

	*TypeMeta                    `json:",inline"`
	*MetadataContainer           `json:",inline"`
	*SpecContainer[OidcAuthSpec] `json:",inline"`
}

type OidcAuthSpec struct {
	BaseSpec `json:"-"`

	UsernameClaim    string             `json:"usernameClaim"`
	UsernamePrefix   *string            `json:"usernamePrefix,omitempty"`
	Issuer           string             `json:"issuer"`
	ClientId         string             `json:"clientId"`
	RequiredClaims   []string           `json:"requiredClaims,omitempty"`
	GroupsClaim      *string            `json:"groupsClaim,omitempty"`
	GroupsPrefix     *string            `json:"groupsPrefix,omitempty"`
	AttributeMapping []AttributeMapping `json:"attributeMapping,omitempty"`
}

type AttributeMapping struct {
	From string `json:"from"`
	To   string `json:"to"`
}

var _ ClusterApiResource[ClusterRoleBindingSpec] = &ClusterRoleBinding{}

type ClusterRoleBinding struct {
	BaseResponse `json:"-"`

	*TypeMeta                              `json:",inline"`
	*MetadataContainer                     `json:",inline"`
	*SpecContainer[ClusterRoleBindingSpec] `json:",inline"`
}

type ClusterRoleBindingSpec struct {
	BaseSpec `json:"-"`

	SubjectsRef  SubjectRef        `json:"subjectsRef"`
	PolicyInline []PolicyStatement `json:"policyInline,omitempty"`
	PolicyRef    *string           `json:"policyRef,omitempty"`
}

var _ NamespacedApiResource[RoleBindingSpec] = &RoleBinding{}

type RoleBinding struct {
	BaseResponse `json:"-"`

	*TypeMeta                       `json:",inline"`
	*MetadataContainer              `json:",inline"`
	*SpecContainer[RoleBindingSpec] `json:",inline"`
}

type RoleBindingSpec struct {
	BaseSpec `json:"-"`

	SubjectsRef  SubjectRef        `json:"subjectsRef"`
	PolicyInline []PolicyStatement `json:"policyInline,omitempty"`
	PolicyRef    *string           `json:"policyRef,omitempty"`
}

type SubjectRef struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type PolicyStatement struct {
	Actions   []string         `json:"actions"`
	Paths     []string         `json:"paths,omitempty"`
	Resources []PolicyResource `json:"resources,omitempty"`
}

type PolicyResource struct {
	ApiGroup string `json:"apiGroup"`
	Kind     string `json:"kind"`
}
