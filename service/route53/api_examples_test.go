// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53_test

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

var _ aws.Config

// To associate a VPC with a hosted zone
//
// The following example associates the VPC with ID vpc-1a2b3c4d with the hosted zone
// with ID Z3M3LMPEXAMPLE.
func ExampleClient_AssociateVPCWithHostedZoneRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.AssociateVPCWithHostedZoneInput{
		Comment:      aws.String(""),
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
		VPC: &route53.VPC{
			VPCId:     aws.String("vpc-1a2b3c4d"),
			VPCRegion: route53.VPCRegionUsEast2,
		},
	}

	req := svc.AssociateVPCWithHostedZoneRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNotAuthorizedException:
				fmt.Println(route53.ErrCodeNotAuthorizedException, aerr.Error())
			case route53.ErrCodeInvalidVPCId:
				fmt.Println(route53.ErrCodeInvalidVPCId, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePublicZoneVPCAssociation:
				fmt.Println(route53.ErrCodePublicZoneVPCAssociation, aerr.Error())
			case route53.ErrCodeConflictingDomainExists:
				fmt.Println(route53.ErrCodeConflictingDomainExists, aerr.Error())
			case route53.ErrCodeLimitsExceeded:
				fmt.Println(route53.ErrCodeLimitsExceeded, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a basic resource record set
//
// The following example creates a resource record set that routes Internet traffic
// to a resource with an IP address of 192.0.2.44.
func ExampleClient_ChangeResourceRecordSetsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String("example.com"),
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.44"),
							},
						},
						TTL:  aws.Int64(60),
						Type: route53.RRTypeA,
					},
				},
			},
			Comment: aws.String("Web server for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.ChangeResourceRecordSetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create weighted resource record sets
//
// The following example creates two weighted resource record sets. The resource with
// a Weight of 100 will get 1/3rd of traffic (100/100+200), and the other resource will
// get the rest of the traffic for example.com.
func ExampleClient_ChangeResourceRecordSetsRequest_shared01() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						HealthCheckId: aws.String("abcdef11-2222-3333-4444-555555fedcba"),
						Name:          aws.String("example.com"),
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.44"),
							},
						},
						SetIdentifier: aws.String("Seattle data center"),
						TTL:           aws.Int64(60),
						Type:          route53.RRTypeA,
						Weight:        aws.Int64(100),
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						HealthCheckId: aws.String("abcdef66-7777-8888-9999-000000fedcba"),
						Name:          aws.String("example.com"),
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.45"),
							},
						},
						SetIdentifier: aws.String("Portland data center"),
						TTL:           aws.Int64(60),
						Type:          route53.RRTypeA,
						Weight:        aws.Int64(200),
					},
				},
			},
			Comment: aws.String("Web servers for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.ChangeResourceRecordSetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create an alias resource record set
//
// The following example creates an alias resource record set that routes traffic to
// a CloudFront distribution.
func ExampleClient_ChangeResourceRecordSetsRequest_shared02() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("d123rk29d0stfj.cloudfront.net"),
							EvaluateTargetHealth: aws.Bool(false),
							HostedZoneId:         aws.String("Z2FDTNDATAQYW2"),
						},
						Name: aws.String("example.com"),
						Type: route53.RRTypeA,
					},
				},
			},
			Comment: aws.String("CloudFront distribution for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.ChangeResourceRecordSetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create weighted alias resource record sets
//
// The following example creates two weighted alias resource record sets that route
// traffic to ELB load balancers. The resource with a Weight of 100 will get 1/3rd of
// traffic (100/100+200), and the other resource will get the rest of the traffic for
// example.com.
func ExampleClient_ChangeResourceRecordSetsRequest_shared03() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("example-com-123456789.us-east-2.elb.amazonaws.com "),
							EvaluateTargetHealth: aws.Bool(true),
							HostedZoneId:         aws.String("Z3AADJGX6KTTL2"),
						},
						Name:          aws.String("example.com"),
						SetIdentifier: aws.String("Ohio region"),
						Type:          route53.RRTypeA,
						Weight:        aws.Int64(100),
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("example-com-987654321.us-west-2.elb.amazonaws.com "),
							EvaluateTargetHealth: aws.Bool(true),
							HostedZoneId:         aws.String("Z1H1FL5HABSF5"),
						},
						Name:          aws.String("example.com"),
						SetIdentifier: aws.String("Oregon region"),
						Type:          route53.RRTypeA,
						Weight:        aws.Int64(200),
					},
				},
			},
			Comment: aws.String("ELB load balancers for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.ChangeResourceRecordSetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create latency resource record sets
//
// The following example creates two latency resource record sets that route traffic
// to EC2 instances. Traffic for example.com is routed either to the Ohio region or
// the Oregon region, depending on the latency between the user and those regions.
func ExampleClient_ChangeResourceRecordSetsRequest_shared04() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						HealthCheckId: aws.String("abcdef11-2222-3333-4444-555555fedcba"),
						Name:          aws.String("example.com"),
						Region:        route53.ResourceRecordSetRegionUsEast2,
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.44"),
							},
						},
						SetIdentifier: aws.String("Ohio region"),
						TTL:           aws.Int64(60),
						Type:          route53.RRTypeA,
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						HealthCheckId: aws.String("abcdef66-7777-8888-9999-000000fedcba"),
						Name:          aws.String("example.com"),
						Region:        route53.ResourceRecordSetRegionUsWest2,
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.45"),
							},
						},
						SetIdentifier: aws.String("Oregon region"),
						TTL:           aws.Int64(60),
						Type:          route53.RRTypeA,
					},
				},
			},
			Comment: aws.String("EC2 instances for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.ChangeResourceRecordSetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create latency alias resource record sets
//
// The following example creates two latency alias resource record sets that route traffic
// for example.com to ELB load balancers. Requests are routed either to the Ohio region
// or the Oregon region, depending on the latency between the user and those regions.
func ExampleClient_ChangeResourceRecordSetsRequest_shared05() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("example-com-123456789.us-east-2.elb.amazonaws.com "),
							EvaluateTargetHealth: aws.Bool(true),
							HostedZoneId:         aws.String("Z3AADJGX6KTTL2"),
						},
						Name:          aws.String("example.com"),
						Region:        route53.ResourceRecordSetRegionUsEast2,
						SetIdentifier: aws.String("Ohio region"),
						Type:          route53.RRTypeA,
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("example-com-987654321.us-west-2.elb.amazonaws.com "),
							EvaluateTargetHealth: aws.Bool(true),
							HostedZoneId:         aws.String("Z1H1FL5HABSF5"),
						},
						Name:          aws.String("example.com"),
						Region:        route53.ResourceRecordSetRegionUsWest2,
						SetIdentifier: aws.String("Oregon region"),
						Type:          route53.RRTypeA,
					},
				},
			},
			Comment: aws.String("ELB load balancers for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.ChangeResourceRecordSetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create failover resource record sets
//
// The following example creates primary and secondary failover resource record sets
// that route traffic to EC2 instances. Traffic is generally routed to the primary resource,
// in the Ohio region. If that resource is unavailable, traffic is routed to the secondary
// resource, in the Oregon region.
func ExampleClient_ChangeResourceRecordSetsRequest_shared06() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						Failover:      route53.ResourceRecordSetFailoverPrimary,
						HealthCheckId: aws.String("abcdef11-2222-3333-4444-555555fedcba"),
						Name:          aws.String("example.com"),
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.44"),
							},
						},
						SetIdentifier: aws.String("Ohio region"),
						TTL:           aws.Int64(60),
						Type:          route53.RRTypeA,
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						Failover:      route53.ResourceRecordSetFailoverSecondary,
						HealthCheckId: aws.String("abcdef66-7777-8888-9999-000000fedcba"),
						Name:          aws.String("example.com"),
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.45"),
							},
						},
						SetIdentifier: aws.String("Oregon region"),
						TTL:           aws.Int64(60),
						Type:          route53.RRTypeA,
					},
				},
			},
			Comment: aws.String("Failover configuration for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.ChangeResourceRecordSetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create failover alias resource record sets
//
// The following example creates primary and secondary failover alias resource record
// sets that route traffic to ELB load balancers. Traffic is generally routed to the
// primary resource, in the Ohio region. If that resource is unavailable, traffic is
// routed to the secondary resource, in the Oregon region.
func ExampleClient_ChangeResourceRecordSetsRequest_shared07() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("example-com-123456789.us-east-2.elb.amazonaws.com "),
							EvaluateTargetHealth: aws.Bool(true),
							HostedZoneId:         aws.String("Z3AADJGX6KTTL2"),
						},
						Failover:      route53.ResourceRecordSetFailoverPrimary,
						Name:          aws.String("example.com"),
						SetIdentifier: aws.String("Ohio region"),
						Type:          route53.RRTypeA,
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("example-com-987654321.us-west-2.elb.amazonaws.com "),
							EvaluateTargetHealth: aws.Bool(true),
							HostedZoneId:         aws.String("Z1H1FL5HABSF5"),
						},
						Failover:      route53.ResourceRecordSetFailoverSecondary,
						Name:          aws.String("example.com"),
						SetIdentifier: aws.String("Oregon region"),
						Type:          route53.RRTypeA,
					},
				},
			},
			Comment: aws.String("Failover alias configuration for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.ChangeResourceRecordSetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create geolocation resource record sets
//
// The following example creates four geolocation resource record sets that use IPv4
// addresses to route traffic to resources such as web servers running on EC2 instances.
// Traffic is routed to one of four IP addresses, for North America (NA), for South
// America (SA), for Europe (EU), and for all other locations (*).
func ExampleClient_ChangeResourceRecordSetsRequest_shared08() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						GeoLocation: &route53.GeoLocation{
							ContinentCode: aws.String("NA"),
						},
						Name: aws.String("example.com"),
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.44"),
							},
						},
						SetIdentifier: aws.String("North America"),
						TTL:           aws.Int64(60),
						Type:          route53.RRTypeA,
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						GeoLocation: &route53.GeoLocation{
							ContinentCode: aws.String("SA"),
						},
						Name: aws.String("example.com"),
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.45"),
							},
						},
						SetIdentifier: aws.String("South America"),
						TTL:           aws.Int64(60),
						Type:          route53.RRTypeA,
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						GeoLocation: &route53.GeoLocation{
							ContinentCode: aws.String("EU"),
						},
						Name: aws.String("example.com"),
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.46"),
							},
						},
						SetIdentifier: aws.String("Europe"),
						TTL:           aws.Int64(60),
						Type:          route53.RRTypeA,
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						GeoLocation: &route53.GeoLocation{
							CountryCode: aws.String("*"),
						},
						Name: aws.String("example.com"),
						ResourceRecords: []route53.ResourceRecord{
							{
								Value: aws.String("192.0.2.47"),
							},
						},
						SetIdentifier: aws.String("Other locations"),
						TTL:           aws.Int64(60),
						Type:          route53.RRTypeA,
					},
				},
			},
			Comment: aws.String("Geolocation configuration for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.ChangeResourceRecordSetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create geolocation alias resource record sets
//
// The following example creates four geolocation alias resource record sets that route
// traffic to ELB load balancers. Traffic is routed to one of four IP addresses, for
// North America (NA), for South America (SA), for Europe (EU), and for all other locations
// (*).
func ExampleClient_ChangeResourceRecordSetsRequest_shared09() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("example-com-123456789.us-east-2.elb.amazonaws.com "),
							EvaluateTargetHealth: aws.Bool(true),
							HostedZoneId:         aws.String("Z3AADJGX6KTTL2"),
						},
						GeoLocation: &route53.GeoLocation{
							ContinentCode: aws.String("NA"),
						},
						Name:          aws.String("example.com"),
						SetIdentifier: aws.String("North America"),
						Type:          route53.RRTypeA,
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("example-com-234567890.sa-east-1.elb.amazonaws.com "),
							EvaluateTargetHealth: aws.Bool(true),
							HostedZoneId:         aws.String("Z2P70J7HTTTPLU"),
						},
						GeoLocation: &route53.GeoLocation{
							ContinentCode: aws.String("SA"),
						},
						Name:          aws.String("example.com"),
						SetIdentifier: aws.String("South America"),
						Type:          route53.RRTypeA,
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("example-com-234567890.eu-central-1.elb.amazonaws.com "),
							EvaluateTargetHealth: aws.Bool(true),
							HostedZoneId:         aws.String("Z215JYRZR1TBD5"),
						},
						GeoLocation: &route53.GeoLocation{
							ContinentCode: aws.String("EU"),
						},
						Name:          aws.String("example.com"),
						SetIdentifier: aws.String("Europe"),
						Type:          route53.RRTypeA,
					},
				},
				{
					Action: route53.ChangeActionCreate,
					ResourceRecordSet: &route53.ResourceRecordSet{
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("example-com-234567890.ap-southeast-1.elb.amazonaws.com "),
							EvaluateTargetHealth: aws.Bool(true),
							HostedZoneId:         aws.String("Z1LMS91P8CMLE5"),
						},
						GeoLocation: &route53.GeoLocation{
							CountryCode: aws.String("*"),
						},
						Name:          aws.String("example.com"),
						SetIdentifier: aws.String("Other locations"),
						Type:          route53.RRTypeA,
					},
				},
			},
			Comment: aws.String("Geolocation alias configuration for example.com"),
		},
		HostedZoneId: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.ChangeResourceRecordSetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeInvalidChangeBatch:
				fmt.Println(route53.ErrCodeInvalidChangeBatch, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To add or remove tags from a hosted zone or health check
//
// The following example adds two tags and removes one tag from the hosted zone with
// ID Z3M3LMPEXAMPLE.
func ExampleClient_ChangeTagsForResourceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.ChangeTagsForResourceInput{
		AddTags: []route53.Tag{
			{
				Key:   aws.String("apex"),
				Value: aws.String("3874"),
			},
			{
				Key:   aws.String("acme"),
				Value: aws.String("4938"),
			},
		},
		RemoveTagKeys: []string{
			"Nadir",
		},
		ResourceId:   aws.String("Z3M3LMPEXAMPLE"),
		ResourceType: route53.TagResourceTypeHostedzone,
	}

	req := svc.ChangeTagsForResourceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			case route53.ErrCodeNoSuchHealthCheck:
				fmt.Println(route53.ErrCodeNoSuchHealthCheck, aerr.Error())
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodePriorRequestNotComplete:
				fmt.Println(route53.ErrCodePriorRequestNotComplete, aerr.Error())
			case route53.ErrCodeThrottlingException:
				fmt.Println(route53.ErrCodeThrottlingException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get information about a hosted zone
//
// The following example gets information about the Z3M3LMPEXAMPLE hosted zone.
func ExampleClient_GetHostedZoneRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := route53.New(cfg)
	input := &route53.GetHostedZoneInput{
		Id: aws.String("Z3M3LMPEXAMPLE"),
	}

	req := svc.GetHostedZoneRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
