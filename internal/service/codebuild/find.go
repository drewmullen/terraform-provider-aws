package codebuild

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// FindReportGroupByARN returns the Report Group corresponding to the specified Arn.
func FindReportGroupByARN(conn *codebuild.CodeBuild, arn string) (*codebuild.ReportGroup, error) {

	output, err := conn.BatchGetReportGroups(&codebuild.BatchGetReportGroupsInput{
		ReportGroupArns: aws.StringSlice([]string{arn}),
	})
	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, nil
	}

	if len(output.ReportGroups) == 0 {
		return nil, nil
	}

	reportGroup := output.ReportGroups[0]
	if reportGroup == nil {
		return nil, nil
	}

	return reportGroup, nil
}

func FindProjectByARN(conn *codebuild.CodeBuild, arn string) (*codebuild.Project, error) {
	input := &codebuild.BatchGetProjectsInput{
		Names: []*string{aws.String(arn)},
	}

	output, err := conn.BatchGetProjects(input)
	if err != nil {
		return nil, err
	}

	if output == nil || len(output.Projects) == 0 || output.Projects[0] == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	if count := len(output.Projects); count > 1 {
		return nil, tfresource.NewTooManyResultsError(count, input)
	}

	return output.Projects[0], nil
}
