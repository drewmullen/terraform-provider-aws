---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_user"
description: |-
  Manages a Resource QuickSight User.
---

# Resource: aws_quicksight_user

Resource for managing QuickSight User

## Example Usage

### Basic quicksight user

```terraform
resource "aws_quicksight_user" "example" {
  user_name     = "an-author"
  email         = "author@example.com"
  identity_type = "QUICKSIGHT"
  user_role     = "AUTHOR"
}
```

### IAM based user

```terraform
resource "aws_iam_user" "admin" {
  name          = "quicksightAdmin"
  path          = "/"
  force_destroy = true
}

resource "aws_iam_user_login_profile" "admin" {
  user    = aws_iam_user.admin.name
  pgp_key = "keybase:username"
}

resource "aws_quicksight_user" "admin" {
  email         = "admin@example.com"
  identity_type = "IAM"
  user_role     = "ADMIN"
  iam_arn       = aws_iam_user.admin.arn
  depends_on    = [aws_iam_user_login_profile.admin]
}
```

## Argument Reference

The following arguments are supported:


* `email` - (Required) The email address of the user that you want to register.
* `identity_type` - (Required) Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
* `user_role` - (Required) The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`.
* `user_name` - (Optional) The Amazon QuickSight user name that you want to create for the user you are registering. Not applicable if `identity_type` is `IAM`
* `aws_account_id` - (Optional) The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
* `iam_arn` - (Optional) The ARN of the IAM user or role that you are registering with Amazon QuickSight. Requied if `identity_type` is `IAM`
* `namespace`  - (Optional) The namespace. Currently, you should set this to `default`.
* `session_name` - (Optional) The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `arn` - Amazon Resource Name (ARN) of the user

## Import

Importing is currently not supported on this resource.
