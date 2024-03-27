// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/serialt/smail"
)

// emailResource defines the resource implementation.
type emailResource struct {
}

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &emailResource{}
	_ resource.ResourceWithImportState = &emailResource{}
)

func NewEmailResource() resource.Resource {
	return &emailResource{}
}

// emailResourceModel describes the resource data model.
type emailResourceModel struct {
	To           types.String `tfsdk:"to"`
	Subject      types.String `tfsdk:"subject"`
	Body         types.String `tfsdk:"body"`
	SmtpServer   types.String `tfsdk:"smtp_server"`
	SmtpPort     types.String `tfsdk:"smtp_port"`
	SmtpUsername types.String `tfsdk:"smtp_username"`
	SmtpPassword types.String `tfsdk:"smtp_password"`
}

func (r *emailResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_email"
}

func (r *emailResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "email resource",

		Attributes: map[string]schema.Attribute{
			"to": schema.StringAttribute{
				MarkdownDescription: "to user",
				Optional:            false,
			},
			"subject": schema.StringAttribute{

				MarkdownDescription: "subject",
				Optional:            false,
			},
			"body": schema.StringAttribute{
				MarkdownDescription: "email body",
				Optional:            false,
			},
			"smtp_server": schema.StringAttribute{
				MarkdownDescription: "smtp server",
				Optional:            false,
			},
			"smtp_port": schema.StringAttribute{
				MarkdownDescription: "smtp server port",
				Optional:            false,
			},
			"smtp_username": schema.StringAttribute{
				MarkdownDescription: "email username",
				Optional:            false,
			},
			"smtp_password": schema.StringAttribute{
				MarkdownDescription: "email password",
				Optional:            false,
			},
		},
	}
}

func (r *emailResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	tflog.Debug(ctx, "Create email resource")

	var data emailResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	msmail := &smail.Mailer{}
	msmail.User = data.SmtpUsername.ValueString()
	msmail.Pass = data.SmtpPassword.ValueString()
	msmail.Smtp = data.SmtpServer.ValueString()
	port, _ := strconv.Atoi(data.SmtpPort.ValueString())
	msmail.Port = port
	msmail.MailTo = []string{data.To.ValueString()}
	msmail.Subject = data.Subject.ValueString()
	msmail.Body = data.Body.ValueString()

	err := msmail.SendMail()

	if err != nil {
		tflog.Error(ctx, "sendmail failed")
		return
	}
	tflog.Debug(ctx, "created a resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *emailResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *emailResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *emailResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *emailResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
