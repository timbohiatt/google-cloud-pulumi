
## Variables

| name | description | type | required | default |
|---|---|:---:|:---:|:---:|
| [name](variables.tf#L140) | Project name and id suffix. | <code>string</code> | ✓ |  |
| [auto_create_network](variables.tf#L17) | Whether to create the default network for the project. | <code>bool</code> |  | <code>false</code> |
| [billing_account](variables.tf#L23) | Billing account id. | <code>string</code> |  | <code>null</code> |
| [contacts](variables.tf#L29) | List of essential contacts for this resource. Must be in the form EMAIL -> [NOTIFICATION_TYPES]. Valid notification types are ALL, SUSPENSION, SECURITY, TECHNICAL, BILLING, LEGAL, PRODUCT_UPDATES. | <code>map&#40;list&#40;string&#41;&#41;</code> |  | <code>&#123;&#125;</code> |
| [custom_roles](variables.tf#L36) | Map of role name => list of permissions to create in this project. | <code>map&#40;list&#40;string&#41;&#41;</code> |  | <code>&#123;&#125;</code> |
| [default_service_account](variables.tf#L43) | Project default service account setting: can be one of `delete`, `deprivilege`, `disable`, or `keep`. | <code>string</code> |  | <code>&#34;keep&#34;</code> |
| [descriptive_name](variables.tf#L49) | Name of the project name. Used for project name instead of `name` variable. | <code>string</code> |  | <code>null</code> |
| [group_iam](variables.tf#L55) | Authoritative IAM binding for organization groups, in {GROUP_EMAIL => [ROLES]} format. Group emails need to be static. Can be used in combination with the `iam` variable. | <code>map&#40;list&#40;string&#41;&#41;</code> |  | <code>&#123;&#125;</code> |
| [iam](variables.tf#L62) | IAM bindings in {ROLE => [MEMBERS]} format. | <code>map&#40;list&#40;string&#41;&#41;</code> |  | <code>&#123;&#125;</code> |
| [iam_additive](variables.tf#L69) | IAM additive bindings in {ROLE => [MEMBERS]} format. | <code>map&#40;list&#40;string&#41;&#41;</code> |  | <code>&#123;&#125;</code> |
| [iam_additive_members](variables.tf#L76) | IAM additive bindings in {MEMBERS => [ROLE]} format. This might break if members are dynamic values. | <code>map&#40;list&#40;string&#41;&#41;</code> |  | <code>&#123;&#125;</code> |
| [labels](variables.tf#L82) | Resource labels. | <code>map&#40;string&#41;</code> |  | <code>&#123;&#125;</code> |
| [lien_reason](variables.tf#L89) | If non-empty, creates a project lien with this description. | <code>string</code> |  | <code>&#34;&#34;</code> |
| [logging_exclusions](variables.tf#L95) | Logging exclusions for this project in the form {NAME -> FILTER}. | <code>map&#40;string&#41;</code> |  | <code>&#123;&#125;</code> |
| [logging_sinks](variables.tf#L102) | Logging sinks to create for this project. | <code title="map&#40;object&#40;&#123;&#10;  bq_partitioned_table &#61; optional&#40;bool&#41;&#10;  description          &#61; optional&#40;string&#41;&#10;  destination          &#61; string&#10;  disabled             &#61; optional&#40;bool, false&#41;&#10;  exclusions           &#61; optional&#40;map&#40;string&#41;, &#123;&#125;&#41;&#10;  filter               &#61; string&#10;  iam                  &#61; optional&#40;bool, true&#41;&#10;  type                 &#61; string&#10;  unique_writer        &#61; optional&#40;bool&#41;&#10;&#125;&#41;&#41;">map&#40;object&#40;&#123;&#8230;&#125;&#41;&#41;</code> |  | <code>&#123;&#125;</code> |
| [metric_scopes](variables.tf#L133) | List of projects that will act as metric scopes for this project. | <code>list&#40;string&#41;</code> |  | <code>&#91;&#93;</code> |
| [org_policies](variables.tf#L145) | Organization policies applied to this project keyed by policy name. | <code title="map&#40;object&#40;&#123;&#10;  inherit_from_parent &#61; optional&#40;bool&#41; &#35; for list policies only.&#10;  reset               &#61; optional&#40;bool&#41;&#10;  rules &#61; optional&#40;list&#40;object&#40;&#123;&#10;    allow &#61; optional&#40;object&#40;&#123;&#10;      all    &#61; optional&#40;bool&#41;&#10;      values &#61; optional&#40;list&#40;string&#41;&#41;&#10;    &#125;&#41;&#41;&#10;    deny &#61; optional&#40;object&#40;&#123;&#10;      all    &#61; optional&#40;bool&#41;&#10;      values &#61; optional&#40;list&#40;string&#41;&#41;&#10;    &#125;&#41;&#41;&#10;    enforce &#61; optional&#40;bool&#41; &#35; for boolean policies only.&#10;    condition &#61; optional&#40;object&#40;&#123;&#10;      description &#61; optional&#40;string&#41;&#10;      expression  &#61; optional&#40;string&#41;&#10;      location    &#61; optional&#40;string&#41;&#10;      title       &#61; optional&#40;string&#41;&#10;    &#125;&#41;, &#123;&#125;&#41;&#10;  &#125;&#41;&#41;, &#91;&#93;&#41;&#10;&#125;&#41;&#41;">map&#40;object&#40;&#123;&#8230;&#125;&#41;&#41;</code> |  | <code>&#123;&#125;</code> |
| [org_policies_data_path](variables.tf#L172) | Path containing org policies in YAML format. | <code>string</code> |  | <code>null</code> |
| [oslogin](variables.tf#L178) | Enable OS Login. | <code>bool</code> |  | <code>false</code> |
| [oslogin_admins](variables.tf#L184) | List of IAM-style identities that will be granted roles necessary for OS Login administrators. | <code>list&#40;string&#41;</code> |  | <code>&#91;&#93;</code> |
| [oslogin_users](variables.tf#L192) | List of IAM-style identities that will be granted roles necessary for OS Login users. | <code>list&#40;string&#41;</code> |  | <code>&#91;&#93;</code> |
| [parent](variables.tf#L199) | Parent folder or organization in 'folders/folder_id' or 'organizations/org_id' format. | <code>string</code> |  | <code>null</code> |
| [prefix](variables.tf#L209) | Optional prefix used to generate project id and name. | <code>string</code> |  | <code>null</code> |
| [project_create](variables.tf#L219) | Create project. When set to false, uses a data source to reference existing project. | <code>bool</code> |  | <code>true</code> |
| [service_config](variables.tf#L225) | Configure service API activation. | <code title="object&#40;&#123;&#10;  disable_on_destroy         &#61; bool&#10;  disable_dependent_services &#61; bool&#10;&#125;&#41;">object&#40;&#123;&#8230;&#125;&#41;</code> |  | <code title="&#123;&#10;  disable_on_destroy         &#61; false&#10;  disable_dependent_services &#61; false&#10;&#125;">&#123;&#8230;&#125;</code> |
| [service_encryption_key_ids](variables.tf#L237) | Cloud KMS encryption key in {SERVICE => [KEY_URL]} format. | <code>map&#40;list&#40;string&#41;&#41;</code> |  | <code>&#123;&#125;</code> |
| [service_perimeter_bridges](variables.tf#L244) | Name of VPC-SC Bridge perimeters to add project into. See comment in the variables file for format. | <code>list&#40;string&#41;</code> |  | <code>null</code> |
| [service_perimeter_standard](variables.tf#L251) | Name of VPC-SC Standard perimeter to add project into. See comment in the variables file for format. | <code>string</code> |  | <code>null</code> |
| [services](variables.tf#L257) | Service APIs to enable. | <code>list&#40;string&#41;</code> |  | <code>&#91;&#93;</code> |
| [shared_vpc_host_config](variables.tf#L263) | Configures this project as a Shared VPC host project (mutually exclusive with shared_vpc_service_project). | <code title="object&#40;&#123;&#10;  enabled          &#61; bool&#10;  service_projects &#61; optional&#40;list&#40;string&#41;, &#91;&#93;&#41;&#10;&#125;&#41;">object&#40;&#123;&#8230;&#125;&#41;</code> |  | <code>null</code> |
| [shared_vpc_service_config](variables.tf#L272) | Configures this project as a Shared VPC service project (mutually exclusive with shared_vpc_host_config). | <code title="object&#40;&#123;&#10;  host_project         &#61; string&#10;  service_identity_iam &#61; optional&#40;map&#40;list&#40;string&#41;&#41;&#41;&#10;&#125;&#41;">object&#40;&#123;&#8230;&#125;&#41;</code> |  | <code>null</code> |
| [skip_delete](variables.tf#L282) | Allows the underlying resources to be destroyed without destroying the project itself. | <code>bool</code> |  | <code>false</code> |
| [tag_bindings](variables.tf#L288) | Tag bindings for this project, in key => tag value id format. | <code>map&#40;string&#41;</code> |  | <code>null</code> |

## Outputs

| name | description | sensitive |
|---|---|:---:|
| [custom_roles](outputs.tf#L17) | Ids of the created custom roles. |  |
| [name](outputs.tf#L25) | Project name. |  |
| [number](outputs.tf#L37) | Project number. |  |
| [project_id](outputs.tf#L54) | Project id. |  |
| [service_accounts](outputs.tf#L73) | Product robot service accounts in project. |  |
| [sink_writer_identities](outputs.tf#L89) | Writer identities created for each sink. |  |
