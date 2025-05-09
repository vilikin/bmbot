terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.96.0"
    }
  }

  backend "s3" {
    bucket = "vilikin-bmbot-tfstate"
    key = "state"
    use_lockfile = true
    region = "eu-north-1"
  }
}

provider "aws" {
  region = "eu-north-1"
}

module "hello_world_lambda" {
  source  = "terraform-aws-modules/lambda/aws"

  function_name = "hello-world-golang"

  cloudwatch_logs_retention_in_days = 3

  handler       = "bootstrap"
  runtime       = "provided.al2023"
  architectures = ["arm64"]

  trigger_on_package_timestamp = false

  logging_log_format = "JSON"

  source_path = [
    {
      path = "${path.module}/../lambda/"
      commands = [
        "GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -tags lambda.norpc -o ./.build/bootstrap ./handler/helloworld/main.go",
        ":zip .build",
      ]
    }
  ]
}
