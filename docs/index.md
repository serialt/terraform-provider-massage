# Message provider 

Provides a resource which allows you to send message by terraform

```hcl
terraform {
  required_providers {
    message = {
      source = "serialt/message"
    }
  }
}

resource "message_email" "example" {
  to      = "to@local.com"
  subject = "Hello from Terraform"
  body    = "This is a test email sent from Terraform using a custom email provider."
  smtp_server   = "smtp.email.com"
  smtp_port     = "465"
  smtp_username = "xxxxxxxx"
  smtp_password = "XVphxxxxxxxxxx"
}
```