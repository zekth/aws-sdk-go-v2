// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package opsworkscm provides the client and types for making API
// requests to OpsWorksCM.
//
// AWS OpsWorks for configuration management (CM) is a service that runs and
// manages configuration management servers. You can use AWS OpsWorks CM to
// create and manage AWS OpsWorks for Chef Automate and AWS OpsWorks for Puppet
// Enterprise servers, and add or remove nodes for the servers to manage.
//
// Glossary of terms
//
//    * Server: A configuration management server that can be highly-available.
//    The configuration management server runs on an Amazon Elastic Compute
//    Cloud (EC2) instance, and may use various other AWS services, such as
//    Amazon Relational Database Service (RDS) and Elastic Load Balancing. A
//    server is a generic abstraction over the configuration manager that you
//    want to use, much like Amazon RDS. In AWS OpsWorks CM, you do not start
//    or stop servers. After you create servers, they continue to run until
//    they are deleted.
//
//    * Engine: The engine is the specific configuration manager that you want
//    to use. Valid values in this release include ChefAutomate and Puppet.
//
//    * Backup: This is an application-level backup of the data that the configuration
//    manager stores. AWS OpsWorks CM creates an S3 bucket for backups when
//    you launch the first server. A backup maintains a snapshot of a server's
//    configuration-related attributes at the time the backup starts.
//
//    * Events: Events are always related to a server. Events are written during
//    server creation, when health checks run, when backups are created, when
//    system maintenance is performed, etc. When you delete a server, the server's
//    events are also deleted.
//
//    * Account attributes: Every account has attributes that are assigned in
//    the AWS OpsWorks CM database. These attributes store information about
//    configuration limits (servers, backups, etc.) and your customer account.
//
// Endpoints
//
// AWS OpsWorks CM supports the following endpoints, all HTTPS. You must connect
// to one of the following endpoints. Your servers can only be accessed or managed
// within the endpoint in which they are created.
//
//    * opsworks-cm.us-east-1.amazonaws.com
//
//    * opsworks-cm.us-east-2.amazonaws.com
//
//    * opsworks-cm.us-west-1.amazonaws.com
//
//    * opsworks-cm.us-west-2.amazonaws.com
//
//    * opsworks-cm.ap-northeast-1.amazonaws.com
//
//    * opsworks-cm.ap-southeast-1.amazonaws.com
//
//    * opsworks-cm.ap-southeast-2.amazonaws.com
//
//    * opsworks-cm.eu-central-1.amazonaws.com
//
//    * opsworks-cm.eu-west-1.amazonaws.com
//
// For more information, see AWS OpsWorks endpoints and quotas (https://docs.aws.amazon.com/general/latest/gr/opsworks-service.html)
// in the AWS General Reference.
//
// Throttling limits
//
// All API operations allow for five requests per second with a burst of 10
// requests per second.
//
// See https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01 for more information on this service.
//
// See opsworkscm package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/opsworkscm/
//
// Using the Client
//
// To use OpsWorksCM with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the OpsWorksCM client for more information on
// creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/opsworkscm/#New
package opsworkscm
