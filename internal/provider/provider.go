// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	mdatasource "github.com/tae2089/terraform-provider-crypto/internal/datasource"
	"github.com/tae2089/terraform-provider-crypto/internal/functions"
)

// Ensure cryptoProvider satisfies various provider interfaces.
var _ provider.Provider = &cryptoProvider{}
var _ provider.ProviderWithFunctions = &cryptoProvider{}

// cryptoProvider defines the provider implementation.
type cryptoProvider struct{}

// cryptoProviderModel describes the provider data model.
type CryptoProviderModel struct{}

func (p *cryptoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "crypto"
}

func (p *cryptoProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
}

func (p *cryptoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *cryptoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *cryptoProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		mdatasource.NewRsaDataSource,
	}
}

func (p *cryptoProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		functions.NewDecryptFunction,
	}
}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &cryptoProvider{}
	}
}
