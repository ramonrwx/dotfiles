# Building Terraform and Local Providers in Termux

This guide will help you set up Terraform and its
providers in Termux on Android.

## Clone the Terraform Repository

- Clone the Terraform repository:

      git clone --depth=1 git@github.com:hashicorp/terraform

- Navigate to the Terraform directory:

      cd terraform

## Update the `pq` Package

> There is a bug in the `pq` package that doesn't
> recognize Termux. It's recommended to update to the
> latest version.

- Update the `pq` package to the latest version:

      go get -u github.com/lib/pq@latest

## Install Terraform

- Install Terraform:

      go install

# Building Terraform Providers

## Linode Provider

### Clone the Linode Provider

- Clone the Linode provider repository:

      git clone --depth=1 -b main git@github.com:linode/terraform-provider-linode

- Navigate to the Linode provider directory:

      cd terraform-provider-linode

### Build the Linode Provider

- Checkout to the latest version (e.g., v2.9.2):

      git checkout v2.9.2

- Build the Linode provider:

      go build

- Rename the provider executable:

      mv terraform-provider-linode terraform-provider-linode_v2.9.2

- Create a local provider directory and move the
executable:

      mkdir -p ~/.terraform.d/plugins/terraform.local/local/linode/2.9.2/android_arm64
      mv terraform-provider-linode_v2.9.2 ~/.terraform.d/plugins/terraform.local/local/linode/2.9.2/android_arm64

## Cloudflare Provider

### Clone the Cloudflare Provider

- Clone the Cloudflare provider repository:

      git clone git@github.com:cloudflare/terraform-provider-cloudflare

- Navigate to the Cloudflare provider directory:

      cd terraform-provider-cloudflare

### Build the Cloudflare Provider

- Checkout to the latest version (e.g., v4.17.0):

      git checkout v4.17.0

- Build the Cloudflare provider:

      go build

- Rename the provider executable:

      mv terraform-provider-cloudflare terraform-provider-cloudflare_v4.17.0

- Create a local provider directory and move the
executable:

      mkdir -p ~/.terraform.d/plugins/terraform.local/local/cloudflare/4.17.0/android_arm64
      mv terraform-provider-cloudflare_v4.17.0 ~/.terraform.d/plugins/terraform.local/local/cloudflare/4.17.0/android_arm64

# Using Local Terraform Providers

You can now use these local providers in your Terraform
configuration. For example, in your `.tf` file, define the
required provider:

```hcl
terraform {
    required_providers {
        cloudflare = {
            source  = "terraform.local/local/cloudflare"
            version = ">=4.17.0"
        }
    }
}
```
