Terraform Provider for CHECK POINT SOFTWARE TECHNOLOGIES
=========================

- Website: https://www.terraform.io
- Documentation: https://www.terraform.io/docs/providers/checkpoint/index.html
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

Using the provider
----------------------
To use a released provider in your Terraform environment, run [`terraform init`](https://www.terraform.io/docs/commands/init.html) and Terraform will automatically install the provider. To specify a particular provider version when installing released providers, see the [Terraform documentation on provider versioning](https://www.terraform.io/docs/configuration/providers.html#version-provider-versions).

To instead use a custom-built provider in your Terraform environment (e.g. the provider binary from the build instructions below), follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory,  run `terraform init` to initialize it.

For either installation method, documentation about the provider specific configuration options can be found on the [provider's website](https://www.terraform.io/docs/providers/checkpoint/index.html).

Requirements
------------
-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

Building the provider
---------------------
1. Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint`

```sh
$ git clone git@github.com:terraform-providers/terraform-provider-checkpoint $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
```

2. To build the provider locally, run the following command from the repository root directory.
```sh
# Windows
go build -o terraform-provider-checkpoint.exe
```

3. (Optional) Enter the provider directory and build the provider. This will put the provider binary in `$GOPATH/bin` directory.
```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ make build
```

4. For local development, update `dev_overrides` configuration file. See the next section below.

Developing the provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*).

You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

From Terraform v0.14 and later, [development overrides for provider developers](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers) can be used to tell Terraform where the local provider is located.

To do this, create Terraform CLI configuration [file](https://developer.hashicorp.com/terraform/cli/config/config-file#locations). For windows create `terraform.rc` in `%APPDATA%` or `~/.terraformrc` for all other platforms, and add the following block:

```hcl
provider_installation {
  dev_overrides {
    "checkpoint" = "<full path to local provider binary directory>",
  }
}
```

Create terraform file (`*.tf`) with provider configuration and run `terraform apply`, Terraform will use the local provider you set in `dev_overrides`.
```hcl
# Local provider configuration
provider "checkpoint" {
  server = "192.0.2.1"
  username = "aa"
  password = "aaaa"
  context = "web_api"
  session_name = "Terraform session"
}

# Add any resource you want to run ...
```

Running local tests
---------------------------
1. Run specific test from IDE. Go to test file ends with `*_test.go` and click 'Run Test'.
   This requires to define the following environment variables: `CHECKPOINT_SERVER`, `CHECKPOINT_USERNAME`, `CHECKPOINT_PASSWORD` and `CHECKPOINT_CONTEXT`

2. (Optional) In order to test the provider, you can simply run `make test`.
```sh
$ make test
```

3. (Optional) In order to run the full suite of Acceptance tests, run `make testacc`.
*Note:* Acceptance tests create real resources, and often cost money to run.
```sh
$ make testacc
```
