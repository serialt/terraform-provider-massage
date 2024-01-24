// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/smtp"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ExampleResource{}
var _ resource.ResourceWithImportState = &ExampleResource{}

func NewExampleResource() resource.Resource {
	return &ExampleResource{}
}

// ExampleResource defines the resource implementation.
type ExampleResource struct {
}

// ExampleResourceModel describes the resource data model.
type ExampleResourceModel struct {
	To           types.String `tfsdk:"to"`
	From         types.String `tfsdk:"from"`
	ReplyTo      types.String `tfsdk:"reply_to"`
	Subject      types.String `tfsdk:"subject"`
	Preamble     types.String `tfsdk:"proeable"`
	Body         types.String `tfsdk:"body"`
	SmtpServer   types.String `tfsdk:"smtp_server"`
	SmtpPort     types.String `tfsdk:"smtp_port"`
	SmtpUsername types.String `tfsdk:"smtp_username"`
	SmtpPassword types.String `tfsdk:"smtp_password"`
}

func (r *ExampleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_email"
}

func (r *ExampleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"to": schema.StringAttribute{
				MarkdownDescription: "to user",
				Optional:            true,
			},
			"from": schema.StringAttribute{
				MarkdownDescription: "from ",
				Default:             stringdefault.StaticString("example value when not configured"),
			},
			"reply_to": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Example identifier",
				Default:             stringdefault.StaticString("user"),
			},
			"subject": schema.StringAttribute{

				MarkdownDescription: "subject",
				Default:             stringdefault.StaticString("user"),
			},
			"preamble": schema.StringAttribute{

				MarkdownDescription: "Example identifier",
				Default:             stringdefault.StaticString("user"),
			},
			"body": schema.StringAttribute{

				MarkdownDescription: "Example identifier",
				Default:             stringdefault.StaticString("user"),
			},
			"smtp_server": schema.StringAttribute{

				MarkdownDescription: "Example identifier",
				Default:             stringdefault.StaticString("user"),
			},
			"smtp_port": schema.StringAttribute{

				MarkdownDescription: "Example identifier",
				Default:             stringdefault.StaticString("user"),
			},
			"smtp_username": schema.StringAttribute{

				MarkdownDescription: "Example identifier",
				Default:             stringdefault.StaticString("user"),
			},
			"smtp_password": schema.StringAttribute{

				MarkdownDescription: "Example identifier",
				Default:             stringdefault.StaticString("password"),
			},
		},
	}
}

func (r *ExampleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

}

func (r *ExampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ExampleResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	data.To = types.StringValue("to")
	data.From = types.StringValue("from")
	data.ReplyTo = types.StringValue("reply_to")
	data.Subject = types.StringValue("subject")
	data.Preamble = types.StringValue("preamble")
	data.Body = types.StringValue("body")
	data.SmtpServer = types.StringValue("smtp_server")
	data.SmtpPort = types.StringValue("smtp_port")
	data.SmtpUsername = types.StringValue("smtp_username")
	data.SmtpPassword = types.StringValue("smtp_password")

	msg := "From: " + data.From.String() + "\n" +
		"To: " + data.To.String() + "\n" +
		"Reply-To: " + data.ReplyTo.String() + "\n" +
		"Subject: " + data.Subject.String() + "\n" +
		data.Preamble.String() + "\n\n" +
		data.Body.String()

	err := smtp.SendMail(data.SmtpServer.String()+":"+data.SmtpPort.String(),
		smtp.PlainAuth("", data.SmtpUsername.String(), data.SmtpPassword.String(), data.SmtpServer.String()),
		data.From.String(), []string{data.To.String()}, []byte(msg))

	if err != nil {
		tflog.Error(ctx, "sendmail failed")
		return
	}
	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "created a resource")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ExampleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *ExampleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *ExampleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *ExampleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
