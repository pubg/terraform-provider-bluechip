package bluechip_models

var _ ListResponse[BaseResponse] = &ListResponseImpl[BaseResponse]{}

type ListResponseImpl[Item BaseResponse] struct {
	BaseResponse `json:"-"`

	*TypeMeta `json:",inline"`
	Metadata  ListMetadata `json:"metadata"`
	Items     []Item       `json:"items"`
}

func (r *ListResponseImpl[Item]) GetMetadata() ListMetadata {
	return r.Metadata
}

func (r *ListResponseImpl[Item]) GetItems() []Item {
	return r.Items
}

type ListRequest struct {
	Items     []QueryTerm `json:"items"`
	NextToken *string     `json:"nextToken,omitempty"`
}

const OperatorFuzzy = "fuzzy"
const OperatorEquals = "equal"
const OperatorNotEquals = "notEqual"
const OperatorWildcard = "wildcard"
const OperatorRegex = "regex"
const OperatorMatchPhrase = "matchPhrase"

type QueryTerm struct {
	Operator string `json:"operator"`
	Field    string `json:"field"`
	Value    string `json:"value"`
}