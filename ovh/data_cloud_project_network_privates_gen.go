// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package ovh

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	ovhtypes "github.com/ovh/terraform-provider-ovh/ovh/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CloudProjectNetworkPrivatesDataSourceSchema(ctx context.Context) schema.Schema {
	attrs := map[string]schema.Attribute{
		"networks": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Network id",
						MarkdownDescription: "Network id",
					},
					"name": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Network name",
						MarkdownDescription: "Network name",
					},
					"regions": schema.ListNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"openstack_id": schema.StringAttribute{
									CustomType:          ovhtypes.TfStringType{},
									Computed:            true,
									Description:         "Network id on openstack region",
									MarkdownDescription: "Network id on openstack region",
								},
								"region": schema.StringAttribute{
									CustomType:          ovhtypes.TfStringType{},
									Computed:            true,
									Description:         "Network region",
									MarkdownDescription: "Network region",
								},
								"status": schema.StringAttribute{
									CustomType:          ovhtypes.TfStringType{},
									Computed:            true,
									Description:         "Network region status",
									MarkdownDescription: "Network region status",
								},
							},
							CustomType: CloudProjectNetworkPrivatesRegionsType{
								ObjectType: types.ObjectType{
									AttrTypes: CloudProjectNetworkPrivatesRegionsValue{}.AttributeTypes(ctx),
								},
							},
						},
						CustomType:          ovhtypes.NewTfListNestedType[CloudProjectNetworkPrivatesRegionsValue](ctx),
						Computed:            true,
						Description:         "Details about private network in region",
						MarkdownDescription: "Details about private network in region",
					},
					"status": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Network status",
						MarkdownDescription: "Network status",
					},
					"type": schema.StringAttribute{
						CustomType:          ovhtypes.TfStringType{},
						Computed:            true,
						Description:         "Network type",
						MarkdownDescription: "Network type",
					},
					"vlan_id": schema.Int64Attribute{
						CustomType:          ovhtypes.TfInt64Type{},
						Computed:            true,
						Description:         "Network VLAN id",
						MarkdownDescription: "Network VLAN id",
					},
				},
				CustomType: CloudProjectNetworkPrivatesType{
					ObjectType: types.ObjectType{
						AttrTypes: CloudProjectNetworkPrivatesValue{}.AttributeTypes(ctx),
					},
				},
			},
			CustomType: ovhtypes.NewTfListNestedType[CloudProjectNetworkPrivatesValue](ctx),
			Computed:   true,
		},
		"service_name": schema.StringAttribute{
			CustomType:          ovhtypes.TfStringType{},
			Required:            true,
			Description:         "Service name",
			MarkdownDescription: "Service name",
		},
	}

	return schema.Schema{
		Attributes: attrs,
	}
}

type CloudProjectNetworkPrivatesModel struct {
	Networks    ovhtypes.TfListNestedValue[CloudProjectNetworkPrivatesValue] `tfsdk:"networks" json:"networks"`
	ServiceName ovhtypes.TfStringValue                                       `tfsdk:"service_name" json:"serviceName"`
}

func (v *CloudProjectNetworkPrivatesModel) MergeWith(other *CloudProjectNetworkPrivatesModel) {

	if (v.Networks.IsUnknown() || v.Networks.IsNull()) && !other.Networks.IsUnknown() {
		v.Networks = other.Networks
	}

	if (v.ServiceName.IsUnknown() || v.ServiceName.IsNull()) && !other.ServiceName.IsUnknown() {
		v.ServiceName = other.ServiceName
	}

}

var _ basetypes.ObjectTypable = CloudProjectNetworkPrivatesType{}

type CloudProjectNetworkPrivatesType struct {
	basetypes.ObjectType
}

func (t CloudProjectNetworkPrivatesType) Equal(o attr.Type) bool {
	other, ok := o.(CloudProjectNetworkPrivatesType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t CloudProjectNetworkPrivatesType) String() string {
	return "CloudProjectNetworkPrivatesType"
}

func (t CloudProjectNetworkPrivatesType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return nil, diags
	}

	idVal, ok := idAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be ovhtypes.TfStringValue, was: %T`, idAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return nil, diags
	}

	nameVal, ok := nameAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be ovhtypes.TfStringValue, was: %T`, nameAttribute))
	}

	regionsAttribute, ok := attributes["regions"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`regions is missing from object`)

		return nil, diags
	}

	regionsVal, ok := regionsAttribute.(ovhtypes.TfListNestedValue[CloudProjectNetworkPrivatesRegionsValue])

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`regions expected to be ovhtypes.TfListNestedValue[CloudProjectNetworkPrivatesRegionsValue], was: %T`, regionsAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return nil, diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	typeAttribute, ok := attributes["type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`type is missing from object`)

		return nil, diags
	}

	typeVal, ok := typeAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`type expected to be ovhtypes.TfStringValue, was: %T`, typeAttribute))
	}

	vlanIdAttribute, ok := attributes["vlan_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`vlan_id is missing from object`)

		return nil, diags
	}

	vlanIdVal, ok := vlanIdAttribute.(ovhtypes.TfInt64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`vlan_id expected to be ovhtypes.TfInt64Value, was: %T`, vlanIdAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return CloudProjectNetworkPrivatesValue{
		Id:                              idVal,
		Name:                            nameVal,
		Regions:                         regionsVal,
		Status:                          statusVal,
		CloudProjectNetworkPrivatesType: typeVal,
		VlanId:                          vlanIdVal,
		state:                           attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectNetworkPrivatesValueNull() CloudProjectNetworkPrivatesValue {
	return CloudProjectNetworkPrivatesValue{
		state: attr.ValueStateNull,
	}
}

func NewCloudProjectNetworkPrivatesValueUnknown() CloudProjectNetworkPrivatesValue {
	return CloudProjectNetworkPrivatesValue{
		state: attr.ValueStateUnknown,
	}
}

func NewCloudProjectNetworkPrivatesValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (CloudProjectNetworkPrivatesValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing CloudProjectNetworkPrivatesValue Attribute Value",
				"While creating a CloudProjectNetworkPrivatesValue value, a missing attribute value was detected. "+
					"A CloudProjectNetworkPrivatesValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectNetworkPrivatesValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid CloudProjectNetworkPrivatesValue Attribute Type",
				"While creating a CloudProjectNetworkPrivatesValue value, an invalid attribute value was detected. "+
					"A CloudProjectNetworkPrivatesValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectNetworkPrivatesValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("CloudProjectNetworkPrivatesValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra CloudProjectNetworkPrivatesValue Attribute Value",
				"While creating a CloudProjectNetworkPrivatesValue value, an extra attribute value was detected. "+
					"A CloudProjectNetworkPrivatesValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra CloudProjectNetworkPrivatesValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewCloudProjectNetworkPrivatesValueUnknown(), diags
	}

	idAttribute, ok := attributes["id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`id is missing from object`)

		return NewCloudProjectNetworkPrivatesValueUnknown(), diags
	}

	idVal, ok := idAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`id expected to be ovhtypes.TfStringValue, was: %T`, idAttribute))
	}

	nameAttribute, ok := attributes["name"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`name is missing from object`)

		return NewCloudProjectNetworkPrivatesValueUnknown(), diags
	}

	nameVal, ok := nameAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`name expected to be ovhtypes.TfStringValue, was: %T`, nameAttribute))
	}

	regionsAttribute, ok := attributes["regions"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`regions is missing from object`)

		return NewCloudProjectNetworkPrivatesValueUnknown(), diags
	}

	regionsVal, ok := regionsAttribute.(ovhtypes.TfListNestedValue[CloudProjectNetworkPrivatesRegionsValue])

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`regions expected to be ovhtypes.TfListNestedValue[CloudProjectNetworkPrivatesRegionsValue], was: %T`, regionsAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return NewCloudProjectNetworkPrivatesValueUnknown(), diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	typeAttribute, ok := attributes["type"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`type is missing from object`)

		return NewCloudProjectNetworkPrivatesValueUnknown(), diags
	}

	typeVal, ok := typeAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`type expected to be ovhtypes.TfStringValue, was: %T`, typeAttribute))
	}

	vlanIdAttribute, ok := attributes["vlan_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`vlan_id is missing from object`)

		return NewCloudProjectNetworkPrivatesValueUnknown(), diags
	}

	vlanIdVal, ok := vlanIdAttribute.(ovhtypes.TfInt64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`vlan_id expected to be ovhtypes.TfInt64Value, was: %T`, vlanIdAttribute))
	}

	if diags.HasError() {
		return NewCloudProjectNetworkPrivatesValueUnknown(), diags
	}

	return CloudProjectNetworkPrivatesValue{
		Id:                              idVal,
		Name:                            nameVal,
		Regions:                         regionsVal,
		Status:                          statusVal,
		CloudProjectNetworkPrivatesType: typeVal,
		VlanId:                          vlanIdVal,
		state:                           attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectNetworkPrivatesValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) CloudProjectNetworkPrivatesValue {
	object, diags := NewCloudProjectNetworkPrivatesValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewCloudProjectNetworkPrivatesValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t CloudProjectNetworkPrivatesType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewCloudProjectNetworkPrivatesValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewCloudProjectNetworkPrivatesValueUnknown(), nil
	}

	if in.IsNull() {
		return NewCloudProjectNetworkPrivatesValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewCloudProjectNetworkPrivatesValueMust(CloudProjectNetworkPrivatesValue{}.AttributeTypes(ctx), attributes), nil
}

func (t CloudProjectNetworkPrivatesType) ValueType(ctx context.Context) attr.Value {
	return CloudProjectNetworkPrivatesValue{}
}

var _ basetypes.ObjectValuable = CloudProjectNetworkPrivatesValue{}

type CloudProjectNetworkPrivatesValue struct {
	Id                              ovhtypes.TfStringValue                                              `tfsdk:"id" json:"id"`
	Name                            ovhtypes.TfStringValue                                              `tfsdk:"name" json:"name"`
	Regions                         ovhtypes.TfListNestedValue[CloudProjectNetworkPrivatesRegionsValue] `tfsdk:"regions" json:"regions"`
	Status                          ovhtypes.TfStringValue                                              `tfsdk:"status" json:"status"`
	CloudProjectNetworkPrivatesType ovhtypes.TfStringValue                                              `tfsdk:"type" json:"type"`
	VlanId                          ovhtypes.TfInt64Value                                               `tfsdk:"vlan_id" json:"vlanId"`
	state                           attr.ValueState
}

func (v *CloudProjectNetworkPrivatesValue) UnmarshalJSON(data []byte) error {
	type JsonCloudProjectNetworkPrivatesValue CloudProjectNetworkPrivatesValue

	var tmp JsonCloudProjectNetworkPrivatesValue
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v.Id = tmp.Id
	v.Name = tmp.Name
	v.Regions = tmp.Regions
	v.Status = tmp.Status
	v.CloudProjectNetworkPrivatesType = tmp.CloudProjectNetworkPrivatesType
	v.VlanId = tmp.VlanId

	v.state = attr.ValueStateKnown

	return nil
}

func (v *CloudProjectNetworkPrivatesValue) MergeWith(other *CloudProjectNetworkPrivatesValue) {

	if (v.Id.IsUnknown() || v.Id.IsNull()) && !other.Id.IsUnknown() {
		v.Id = other.Id
	}

	if (v.Name.IsUnknown() || v.Name.IsNull()) && !other.Name.IsUnknown() {
		v.Name = other.Name
	}

	if (v.Regions.IsUnknown() || v.Regions.IsNull()) && !other.Regions.IsUnknown() {
		v.Regions = other.Regions
	}

	if (v.Status.IsUnknown() || v.Status.IsNull()) && !other.Status.IsUnknown() {
		v.Status = other.Status
	}

	if (v.CloudProjectNetworkPrivatesType.IsUnknown() || v.CloudProjectNetworkPrivatesType.IsNull()) && !other.CloudProjectNetworkPrivatesType.IsUnknown() {
		v.CloudProjectNetworkPrivatesType = other.CloudProjectNetworkPrivatesType
	}

	if (v.VlanId.IsUnknown() || v.VlanId.IsNull()) && !other.VlanId.IsUnknown() {
		v.VlanId = other.VlanId
	}

	if (v.state == attr.ValueStateUnknown || v.state == attr.ValueStateNull) && other.state != attr.ValueStateUnknown {
		v.state = other.state
	}
}

func (v CloudProjectNetworkPrivatesValue) Attributes() map[string]attr.Value {
	return map[string]attr.Value{
		"id":      v.Id,
		"name":    v.Name,
		"regions": v.Regions,
		"status":  v.Status,
		"type":    v.CloudProjectNetworkPrivatesType,
		"vlanId":  v.VlanId,
	}
}
func (v CloudProjectNetworkPrivatesValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 6)

	var val tftypes.Value
	var err error

	attrTypes["id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["name"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["regions"] = basetypes.ListType{
		ElemType: CloudProjectNetworkPrivatesRegionsValue{}.Type(ctx),
	}.TerraformType(ctx)
	attrTypes["status"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["type"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["vlan_id"] = basetypes.Int64Type{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 6)

		val, err = v.Id.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["id"] = val

		val, err = v.Name.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["name"] = val

		val, err = v.Regions.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["regions"] = val

		val, err = v.Status.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["status"] = val

		val, err = v.CloudProjectNetworkPrivatesType.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["type"] = val

		val, err = v.VlanId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["vlan_id"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v CloudProjectNetworkPrivatesValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v CloudProjectNetworkPrivatesValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v CloudProjectNetworkPrivatesValue) String() string {
	return "CloudProjectNetworkPrivatesValue"
}

func (v CloudProjectNetworkPrivatesValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"id":      ovhtypes.TfStringType{},
			"name":    ovhtypes.TfStringType{},
			"regions": ovhtypes.NewTfListNestedType[CloudProjectNetworkPrivatesRegionsValue](ctx),
			"status":  ovhtypes.TfStringType{},
			"type":    ovhtypes.TfStringType{},
			"vlan_id": ovhtypes.TfInt64Type{},
		},
		map[string]attr.Value{
			"id":      v.Id,
			"name":    v.Name,
			"regions": v.Regions,
			"status":  v.Status,
			"type":    v.CloudProjectNetworkPrivatesType,
			"vlan_id": v.VlanId,
		})

	return objVal, diags
}

func (v CloudProjectNetworkPrivatesValue) Equal(o attr.Value) bool {
	other, ok := o.(CloudProjectNetworkPrivatesValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Id.Equal(other.Id) {
		return false
	}

	if !v.Name.Equal(other.Name) {
		return false
	}

	if !v.Regions.Equal(other.Regions) {
		return false
	}

	if !v.Status.Equal(other.Status) {
		return false
	}

	if !v.CloudProjectNetworkPrivatesType.Equal(other.CloudProjectNetworkPrivatesType) {
		return false
	}

	if !v.VlanId.Equal(other.VlanId) {
		return false
	}

	return true
}

func (v CloudProjectNetworkPrivatesValue) Type(ctx context.Context) attr.Type {
	return CloudProjectNetworkPrivatesType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v CloudProjectNetworkPrivatesValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"id":      ovhtypes.TfStringType{},
		"name":    ovhtypes.TfStringType{},
		"regions": ovhtypes.NewTfListNestedType[CloudProjectNetworkPrivatesRegionsValue](ctx),
		"status":  ovhtypes.TfStringType{},
		"type":    ovhtypes.TfStringType{},
		"vlan_id": ovhtypes.TfInt64Type{},
	}
}

var _ basetypes.ObjectTypable = CloudProjectNetworkPrivatesRegionsType{}

type CloudProjectNetworkPrivatesRegionsType struct {
	basetypes.ObjectType
}

func (t CloudProjectNetworkPrivatesRegionsType) Equal(o attr.Type) bool {
	other, ok := o.(CloudProjectNetworkPrivatesRegionsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t CloudProjectNetworkPrivatesRegionsType) String() string {
	return "CloudProjectNetworkPrivatesRegionsType"
}

func (t CloudProjectNetworkPrivatesRegionsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	openstackIdAttribute, ok := attributes["openstack_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`openstack_id is missing from object`)

		return nil, diags
	}

	openstackIdVal, ok := openstackIdAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`openstack_id expected to be ovhtypes.TfStringValue, was: %T`, openstackIdAttribute))
	}

	regionAttribute, ok := attributes["region"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`region is missing from object`)

		return nil, diags
	}

	regionVal, ok := regionAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`region expected to be ovhtypes.TfStringValue, was: %T`, regionAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return nil, diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return CloudProjectNetworkPrivatesRegionsValue{
		OpenstackId: openstackIdVal,
		Region:      regionVal,
		Status:      statusVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectNetworkPrivatesRegionsValueNull() CloudProjectNetworkPrivatesRegionsValue {
	return CloudProjectNetworkPrivatesRegionsValue{
		state: attr.ValueStateNull,
	}
}

func NewCloudProjectNetworkPrivatesRegionsValueUnknown() CloudProjectNetworkPrivatesRegionsValue {
	return CloudProjectNetworkPrivatesRegionsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewCloudProjectNetworkPrivatesRegionsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (CloudProjectNetworkPrivatesRegionsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing CloudProjectNetworkPrivatesRegionsValue Attribute Value",
				"While creating a CloudProjectNetworkPrivatesRegionsValue value, a missing attribute value was detected. "+
					"A CloudProjectNetworkPrivatesRegionsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectNetworkPrivatesRegionsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid CloudProjectNetworkPrivatesRegionsValue Attribute Type",
				"While creating a CloudProjectNetworkPrivatesRegionsValue value, an invalid attribute value was detected. "+
					"A CloudProjectNetworkPrivatesRegionsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("CloudProjectNetworkPrivatesRegionsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("CloudProjectNetworkPrivatesRegionsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra CloudProjectNetworkPrivatesRegionsValue Attribute Value",
				"While creating a CloudProjectNetworkPrivatesRegionsValue value, an extra attribute value was detected. "+
					"A CloudProjectNetworkPrivatesRegionsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra CloudProjectNetworkPrivatesRegionsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewCloudProjectNetworkPrivatesRegionsValueUnknown(), diags
	}

	openstackIdAttribute, ok := attributes["openstack_id"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`openstack_id is missing from object`)

		return NewCloudProjectNetworkPrivatesRegionsValueUnknown(), diags
	}

	openstackIdVal, ok := openstackIdAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`openstack_id expected to be ovhtypes.TfStringValue, was: %T`, openstackIdAttribute))
	}

	regionAttribute, ok := attributes["region"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`region is missing from object`)

		return NewCloudProjectNetworkPrivatesRegionsValueUnknown(), diags
	}

	regionVal, ok := regionAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`region expected to be ovhtypes.TfStringValue, was: %T`, regionAttribute))
	}

	statusAttribute, ok := attributes["status"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`status is missing from object`)

		return NewCloudProjectNetworkPrivatesRegionsValueUnknown(), diags
	}

	statusVal, ok := statusAttribute.(ovhtypes.TfStringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`status expected to be ovhtypes.TfStringValue, was: %T`, statusAttribute))
	}

	if diags.HasError() {
		return NewCloudProjectNetworkPrivatesRegionsValueUnknown(), diags
	}

	return CloudProjectNetworkPrivatesRegionsValue{
		OpenstackId: openstackIdVal,
		Region:      regionVal,
		Status:      statusVal,
		state:       attr.ValueStateKnown,
	}, diags
}

func NewCloudProjectNetworkPrivatesRegionsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) CloudProjectNetworkPrivatesRegionsValue {
	object, diags := NewCloudProjectNetworkPrivatesRegionsValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewCloudProjectNetworkPrivatesRegionsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t CloudProjectNetworkPrivatesRegionsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewCloudProjectNetworkPrivatesRegionsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewCloudProjectNetworkPrivatesRegionsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewCloudProjectNetworkPrivatesRegionsValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewCloudProjectNetworkPrivatesRegionsValueMust(CloudProjectNetworkPrivatesRegionsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t CloudProjectNetworkPrivatesRegionsType) ValueType(ctx context.Context) attr.Value {
	return CloudProjectNetworkPrivatesRegionsValue{}
}

var _ basetypes.ObjectValuable = CloudProjectNetworkPrivatesRegionsValue{}

type CloudProjectNetworkPrivatesRegionsValue struct {
	OpenstackId ovhtypes.TfStringValue `tfsdk:"openstack_id" json:"openstackId"`
	Region      ovhtypes.TfStringValue `tfsdk:"region" json:"region"`
	Status      ovhtypes.TfStringValue `tfsdk:"status" json:"status"`
	state       attr.ValueState
}

func (v *CloudProjectNetworkPrivatesRegionsValue) UnmarshalJSON(data []byte) error {
	type JsonCloudProjectNetworkPrivatesRegionsValue CloudProjectNetworkPrivatesRegionsValue

	var tmp JsonCloudProjectNetworkPrivatesRegionsValue
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	v.OpenstackId = tmp.OpenstackId
	v.Region = tmp.Region
	v.Status = tmp.Status

	v.state = attr.ValueStateKnown

	return nil
}

func (v *CloudProjectNetworkPrivatesRegionsValue) MergeWith(other *CloudProjectNetworkPrivatesRegionsValue) {

	if (v.OpenstackId.IsUnknown() || v.OpenstackId.IsNull()) && !other.OpenstackId.IsUnknown() {
		v.OpenstackId = other.OpenstackId
	}

	if (v.Region.IsUnknown() || v.Region.IsNull()) && !other.Region.IsUnknown() {
		v.Region = other.Region
	}

	if (v.Status.IsUnknown() || v.Status.IsNull()) && !other.Status.IsUnknown() {
		v.Status = other.Status
	}

	if (v.state == attr.ValueStateUnknown || v.state == attr.ValueStateNull) && other.state != attr.ValueStateUnknown {
		v.state = other.state
	}
}

func (v CloudProjectNetworkPrivatesRegionsValue) Attributes() map[string]attr.Value {
	return map[string]attr.Value{
		"openstackId": v.OpenstackId,
		"region":      v.Region,
		"status":      v.Status,
	}
}
func (v CloudProjectNetworkPrivatesRegionsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 3)

	var val tftypes.Value
	var err error

	attrTypes["openstack_id"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["region"] = basetypes.StringType{}.TerraformType(ctx)
	attrTypes["status"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 3)

		val, err = v.OpenstackId.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["openstack_id"] = val

		val, err = v.Region.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["region"] = val

		val, err = v.Status.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["status"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v CloudProjectNetworkPrivatesRegionsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v CloudProjectNetworkPrivatesRegionsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v CloudProjectNetworkPrivatesRegionsValue) String() string {
	return "CloudProjectNetworkPrivatesRegionsValue"
}

func (v CloudProjectNetworkPrivatesRegionsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	objVal, diags := types.ObjectValue(
		map[string]attr.Type{
			"openstack_id": ovhtypes.TfStringType{},
			"region":       ovhtypes.TfStringType{},
			"status":       ovhtypes.TfStringType{},
		},
		map[string]attr.Value{
			"openstack_id": v.OpenstackId,
			"region":       v.Region,
			"status":       v.Status,
		})

	return objVal, diags
}

func (v CloudProjectNetworkPrivatesRegionsValue) Equal(o attr.Value) bool {
	other, ok := o.(CloudProjectNetworkPrivatesRegionsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.OpenstackId.Equal(other.OpenstackId) {
		return false
	}

	if !v.Region.Equal(other.Region) {
		return false
	}

	if !v.Status.Equal(other.Status) {
		return false
	}

	return true
}

func (v CloudProjectNetworkPrivatesRegionsValue) Type(ctx context.Context) attr.Type {
	return CloudProjectNetworkPrivatesRegionsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v CloudProjectNetworkPrivatesRegionsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"openstack_id": ovhtypes.TfStringType{},
		"region":       ovhtypes.TfStringType{},
		"status":       ovhtypes.TfStringType{},
	}
}
