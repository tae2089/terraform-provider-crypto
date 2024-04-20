package datasource

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/tae2089/terraform-provider-crypto/internal/util"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &RsaDataSource{}

func NewRsaDataSource() datasource.DataSource {
	return &RsaDataSource{}
}

// RsaDataSource defines the data source implementation.
type RsaDataSource struct {
}

// RsaDataSourceModel describes the data source data model.
type RsaDataSourceModel struct {
	PrivateKey types.String `tfsdk:"private_key"`
	PublicKey  types.String `tfsdk:"public_key"`
}

func (d *RsaDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rsa"
}

func (d *RsaDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"private_key": schema.StringAttribute{
				Description: "Private key used to decrypt data.",
				Computed:    true,
			},
			"public_key": schema.StringAttribute{
				Description: "Public key used to encrypt data.",
				Computed:    true,
			},
		},
	}
}

func (d *RsaDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {

}

func (d *RsaDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Read Terraform configuration data into the model
	// var diags diag.Diagnostics
	var state RsaDataSourceModel
	privateKey, publicKey, err := util.CreateRSA()
	if err != nil {
		resp.Diagnostics.AddError("failed to generate RSA key pair", fmt.Sprintf("Unable to read example, got error: %s", err))
	}
	tflog.Trace(ctx, "read a data source")
	state = RsaDataSourceModel{
		PrivateKey: types.StringValue(privateKey),
		PublicKey:  types.StringValue(publicKey),
	}
	//object, _ := types.ObjectValueFrom(ctx, rsaOutputObjectAttrTypes(), o)
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
