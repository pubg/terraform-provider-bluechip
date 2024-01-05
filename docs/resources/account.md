---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bluechip_account Resource - terraform-provider-bluechip"
subcategory: ""
description: |-
  
---

# bluechip_account (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Block List, Min: 1) (see [below for nested schema](#nestedblock--metadata))
- `spec` (Block List, Min: 1) (see [below for nested schema](#nestedblock--spec))

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Name is the name of the resource.
- `namespace` (String) Namespace is the namespace of the resource.

Optional:

- `annotations` (Map of String) Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.
- `labels` (Map of String) Labels are key value pairs that may be used to scope and select individual resources. They are not queryable and should be preserved when modifying objects.

Read-Only:

- `creation_timestamp` (String) CreationTimestamp is a timestamp representing the server time when this object was created.
- `update_timestamp` (String) UpdateTimestamp is a timestamp representing the server time when this object was last updated.


<a id="nestedblock--spec"></a>
### Nested Schema for `spec`

Required:

- `account_id` (String)
- `alias` (String)
- `description` (String)
- `display_name` (String)
- `regions` (Set of String)
- `vendor` (String)


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `default` (String)