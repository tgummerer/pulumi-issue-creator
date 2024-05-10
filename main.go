package main

import (
	"fmt"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Issue struct {
	Title         string
	Description   string
	URL           pulumi.StringOutput
	NoCreateIssue bool
}

type Section struct {
	Title       string
	Issues      []*Issue
	SubSections []*Section
}

type EpicIssue struct {
	Title         string
	Description   string
	RelatedIssues []int
	Sections      []*Section
}

func createIssue(ctx *pulumi.Context, repo *github.Repository, issue *Issue) {
	i, err := github.NewIssue(ctx, issue.Title, &github.IssueArgs{
		Repository: repo.Name,
		Title:      pulumi.String(issue.Title),
		Body:       pulumi.String(issue.Description),
		Labels: pulumi.StringArray{
			pulumi.String("kind/task"),
		},
	})
	if err != nil {
		fmt.Println("Error creating issue", err)
		return
	}
	issue.URL = pulumi.Sprintf("https://github.com/pulumi/pulumi/issues/%d", i.Number)
}

func generateDescription(epic *EpicIssue) pulumi.StringOutput {
	description := pulumi.Sprintf(epic.Description)
	relatedIssues := ""
	for _, issue := range epic.RelatedIssues {
		relatedIssues += fmt.Sprintf("#%d ", issue)
	}
	if relatedIssues != "" {
		description = pulumi.Sprintf("%s\n\nRelated issues: %s", description, relatedIssues)
	}
	for _, section := range epic.Sections {
		description = pulumi.Sprintf("%s\n\n## %s", description, section.Title)
		for _, issue := range section.Issues {
			if issue.NoCreateIssue {
				description = pulumi.Sprintf("%s\n- [ ] %s", description, issue.Title)
			} else {
				description = pulumi.Sprintf("%s\n- [ ] %s", description, issue.URL)
			}
		}

		for _, subSection := range section.SubSections {
			description = pulumi.Sprintf("%s\n\n### %s", description, subSection.Title)
			for _, issue := range subSection.Issues {
				if issue.NoCreateIssue {
					description = pulumi.Sprintf("%s\n- [ ] %s", description, issue.Title)
				} else {
					description = pulumi.Sprintf("%s\n- [ ] %s", description, issue.URL)
				}
			}
		}
	}
	return description
}

func createEpicIssue(ctx *pulumi.Context, repo *github.Repository, epic *EpicIssue) {
	// Create all the sub-issues first
	for _, section := range epic.Sections {
		for _, issue := range section.Issues {
			if issue.NoCreateIssue {
				continue
			}

			createIssue(ctx, repo, issue)
		}

		for _, subSection := range section.SubSections {
			for _, issue := range subSection.Issues {
				if issue.NoCreateIssue {
					continue
				}
				createIssue(ctx, repo, issue)
			}
		}
	}
	description := generateDescription(epic)
	_, err := github.NewIssue(ctx, epic.Title, &github.IssueArgs{
		Repository: repo.Name,
		Title:      pulumi.String("[Epic] " + epic.Title),
		Body:       description,
		Labels: pulumi.StringArray{
			pulumi.String("kind/epic"),
		},
	})
	if err != nil {
		fmt.Println("Error creating issue", err)
		return
	}
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		repo, err := github.NewRepository(ctx, "pulumi", &github.RepositoryArgs{
			DefaultBranch:       pulumi.String("master"),
			DeleteBranchOnMerge: pulumi.Bool(true),
			Description:         pulumi.String("Pulumi - Infrastructure as Code in any programming language. Build infrastructure intuitively on any cloud using familiar languages 🚀"),
			HasDiscussions:      pulumi.Bool(true),
			HasDownloads:        pulumi.Bool(true),
			HasIssues:           pulumi.Bool(true),
			HasProjects:         pulumi.Bool(true),
			HasWiki:             pulumi.Bool(true),
			HomepageUrl:         pulumi.String("https://www.pulumi.com"),
			Name:                pulumi.String("pulumi"),
			SecurityAndAnalysis: &github.RepositorySecurityAndAnalysisArgs{
				SecretScanning: &github.RepositorySecurityAndAnalysisSecretScanningArgs{
					Status: pulumi.String("disabled"),
				},
				SecretScanningPushProtection: &github.RepositorySecurityAndAnalysisSecretScanningPushProtectionArgs{
					Status: pulumi.String("disabled"),
				},
			},
			SquashMergeCommitMessage: pulumi.String("PR_BODY"),
			SquashMergeCommitTitle:   pulumi.String("PR_TITLE"),
			Topics: pulumi.StringArray{
				pulumi.String("javascript"),
				pulumi.String("aws"),
				pulumi.String("csharp"),
				pulumi.String("gcp"),
				pulumi.String("azure"),
				pulumi.String("dotnet"),
				pulumi.String("typescript"),
				pulumi.String("serverless"),
				pulumi.String("python"),
				pulumi.String("iac"),
				pulumi.String("go"),
				pulumi.String("cloud"),
				pulumi.String("kubernetes"),
				pulumi.String("infrastructure-as-code"),
				pulumi.String("containers"),
				pulumi.String("fsharp"),
				pulumi.String("golang"),
				pulumi.String("cloud-computing"),
			},
			Visibility:          pulumi.String("public"),
			VulnerabilityAlerts: pulumi.Bool(true),
		}, pulumi.Protect(true))
		if err != nil {
			return err
		}

		epicIssue := &EpicIssue{
			Title: "Programmatic Default Providers",
			Description: `Providers in pulumi can be instantiated in two different ways.  One is implicitly through the program.  If the user doesn’t set an explicit provider for a resource, the resource automatically gets a default provider assigned.  On the other hand, providers that have been created beforehand can be passed along to the resource registration request and that provider will be used to create the resource.

There is currently no way to set up a default provider for all the resources in a particular pulumi program.  This especially in combination with default providers can lead to confusion, as it’s easy to forget to pass in a provider to some resource, and then there could be multiple resources with different provider settings in the same program.  For example for the kubernetes provider, someone could carefully set one up, and then forget to pass that provider to a kubernetes resource. That would then create a new kubernetes default provider based on a local kubeconfig, which is probably not what the user wanted.

This issue is to track the work to allow users to programmatically set up a default provider for all the resources a provider supports.
`,
			RelatedIssues: []int{2059},
			Sections: []*Section{
				{
					Title: "Design (M103)",
					Issues: []*Issue{
						{
							Title:         "Write design doc for programmatic default providers",
							NoCreateIssue: true,
						},
					},
				},
				{
					Title: "Implementation Plan",
					SubSections: []*Section{
						{
							Title: "M104",
							Issues: []*Issue{
								{
									Title: "Engine support for programmatic default providers",
								},
								{
									Title: "Go SDK support for programmatic default providers",
								},
								{
									Title: "Node SDK support for programmatic default providers",
								},
								{
									Title: "Python SDK support for programmatic default providers",
								},
							},
						},
						{
							Title: "M105",
							Issues: []*Issue{
								{
									Title: ".NET SDK support for programmatic default providers",
								},
								{
									Title: "YAML SDK support for programmatic default providers",
								},
							},
						},
						{
							Title: "Future",
							Issues: []*Issue{
								{
									Title: "Java SDK support for programmatic default providers",
								},
							},
						},
					},
				},
				{
					Title: "Announce",
					Issues: []*Issue{
						{
							Title: "Write Blog post for programmatic default providers",
						},
						{
							Title:         ":icecream:",
							NoCreateIssue: true,
						},
					},
				},
			},
		}

		createEpicIssue(ctx, repo, epicIssue)
		return nil
	})
}
