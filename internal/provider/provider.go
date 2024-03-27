// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &emailProvider{}

func New() provider.Provider {
	return &emailProvider{}
}

type emailProvider struct {
}

func (p *emailProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "message"
}

func (p *emailProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
}

func (p *emailProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func (p *emailProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewEmailResource,
	}
}

func (p *emailProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
