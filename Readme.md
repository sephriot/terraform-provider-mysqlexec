# MySQLExec


### Motivation
During my everyday work my team and I had to execute custom SQL query on MySql server.
MySQL provider does not allow executing custom SQL commands on the server.

I've decided to create custom provider for both enterprise use and for fun.

I know well that Terraform is a tool for infrastructure provisioning, not creating data structures.
This custom provider was created mostly out of curiosity and necessity for single use case.

### Usage

This provider is mostly based on MySQL provider.

First compile the provider by running:
```
go build .
```

Then copy it to example directory, and fill up missing values.
```$xslt
cp terraform-provider-mysqlexec example/
cd example/
```

Now initialize providers in example directory and execute apply
```$xslt
terraform init
terraform apply
```

That's all, you should see the changes done by your custom script.

Both direct query and one stored in file can be used with this provider.
See example for more details.