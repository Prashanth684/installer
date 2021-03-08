package aws

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/sets"

	awsconfig "github.com/openshift/installer/pkg/asset/installconfig/aws"
)

// PreferredInstanceType returns a preferred instance type from the list of instance types provided in descending order of preference
// based on filters like the list of required availability zones.
func PreferredInstanceType(ctx context.Context, meta *awsconfig.Metadata, types []string, zones []string) (string, []string, error) {
	if len(types) == 0 {
		return "", zones, errors.New("at least one instance type required, empty instance types given")
	}

	sess, err := meta.Session(ctx)
	if err != nil {
		return types[0], zones, err
	}

	client := ec2.New(sess, aws.NewConfig().WithRegion(meta.Region))
	resp, err := client.DescribeInstanceTypeOfferingsWithContext(ctx, &ec2.DescribeInstanceTypeOfferingsInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("location"),
				Values: aws.StringSlice(zones),
			},
			{
				Name:   aws.String("instance-type"),
				Values: aws.StringSlice(types),
			},
		},
		LocationType: aws.String("availability-zone"),
	})
	if err != nil {
		return types[0], zones, err
	}
	reqZones := sets.NewString(zones...)
	found := map[string][]string{}
	for _, offering := range resp.InstanceTypeOfferings {
		found[aws.StringValue(offering.InstanceType)] = append(found[aws.StringValue(offering.InstanceType)], aws.StringValue(offering.Location))
	}
	for _, t := range types {
		if reqZones.Difference(sets.NewString(found[t]...)).Len() == 0 {
			return t, zones, nil
		}
	}
	// HACK: in the ARM case, the m6g and m6gd instances are not available in all the zones. This might be OK as is, but the
	// manifests generated list all the zones as available which is incorrect. Instead return only the available zones
	// Delete the zones in which the instance is not available and return the rest
	for k := range reqZones.Difference(sets.NewString(found[types[0]]...)) {
		reqZones.Delete(k)
	}
	return types[0], reqZones.List(), errors.New("no instance type found for the zone constraint")
}
