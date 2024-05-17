package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IssueFromFile struct {
	Title            string
	DescriptionPath  string
	Assignees        []string
	AdditionalLabels []string
}

type Issue struct {
	Title            string
	Description      string
	Assignees        []string
	URL              pulumi.StringOutput
	NoCreateIssue    bool
	AdditionalText   string
	Done             bool
	AdditionalLabels []string
}

type Section struct {
	Title       string
	Issues      []*Issue
	SubSections []*Section
}

type EpicIssue struct {
	Title         string
	Description   string
	Assignees     []string
	RelatedIssues []int
	Sections      []*Section
}

func createFlakyTestIssue(ctx *pulumi.Context, repo *github.Repository, issue *IssueFromFile) {
	description, err := os.ReadFile(filepath.Join("issue-texts", issue.DescriptionPath))
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	labels := pulumi.ToStringArray(
		append(issue.AdditionalLabels, "kind/engineering", "impact/flaky-test"))
	_, err = github.NewIssue(ctx, issue.Title, &github.IssueArgs{
		Repository: repo.Name,
		Title:      pulumi.String(fmt.Sprintf("`%s` is flaky", issue.Title)),
		Body:       pulumi.String(description),
		Assignees:  pulumi.ToStringArray(issue.Assignees),
		Labels:     labels,
	})
	if err != nil {
		fmt.Println("Error creating issue", err)
		return
	}
}

func createIssue(ctx *pulumi.Context, repo *github.Repository, issue *Issue) {
	i, err := github.NewIssue(ctx, issue.Title, &github.IssueArgs{
		Repository: repo.Name,
		Title:      pulumi.String(issue.Title),
		Body:       pulumi.String(issue.Description),
		Assignees:  pulumi.ToStringArray(issue.Assignees),
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

func createWorkItem(ctx *pulumi.Context, repo *github.Repository, issue *Issue) {
	labels := pulumi.ToStringArray(
		append(issue.AdditionalLabels, "kind/engineering"),
	)
	_, err := github.NewIssue(ctx, issue.Title, &github.IssueArgs{
		Repository: repo.Name,
		Title:      pulumi.String(issue.Title),
		Body:       pulumi.String(issue.Description),
		Assignees:  pulumi.ToStringArray(issue.Assignees),
		Labels:     labels,
	})
	if err != nil {
		fmt.Println("Error creating issue", err)
		return
	}
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
			check := " "
			if issue.Done {
				check = "x"
			}
			if issue.NoCreateIssue {
				description = pulumi.Sprintf("%s\n- [%s] %s", description, check, issue.Title)
			} else {
				description = pulumi.Sprintf("%s\n- [%s] %s", description, check, issue.URL)
			}
			if issue.AdditionalText != "" {
				description = pulumi.Sprintf("%s (%s)", description, issue.AdditionalText)
			}
		}

		for _, subSection := range section.SubSections {
			description = pulumi.Sprintf("%s\n\n### %s", description, subSection.Title)
			for _, issue := range subSection.Issues {
				check := " "
				if issue.Done {
					check = "x"
				}
				if issue.NoCreateIssue {
					description = pulumi.Sprintf("%s\n- [%s] %s", description, check, issue.Title)
				} else {
					description = pulumi.Sprintf("%s\n- [%s] %s", description, check, issue.URL)
				}
				if issue.AdditionalText != "" {
					description = pulumi.Sprintf("%s (%s)", description, issue.AdditionalText)
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
		Assignees:  pulumi.ToStringArray(epic.Assignees),
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
			Assignees:     []string{"tgummerer"},
			RelatedIssues: []int{2059},
			Sections: []*Section{
				{
					Title: "Design (M103)",
					Issues: []*Issue{
						{
							Title:          "Write design doc for programmatic default providers",
							NoCreateIssue:  true,
							AdditionalText: "https://docs.google.com/document/d/12AhuLGpK-hV5f0Zv0PqO9JwxuiRJm5xjQk2YkWYJKuk/edit",
							Done:           true,
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
									Title:          "Engine support for programmatic default providers",
									AdditionalText: "[WIP PR](https://github.com/pulumi/pulumi/pull/16105)",
								},
								{
									Title:          "Go SDK support for programmatic default providers",
									AdditionalText: "[WIP PR](https://github.com/pulumi/pulumi/pull/16105)",
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

		pulumiDotnetIssue := &Issue{
			Title:            "Stop hardcoding the dotnet SDK version number in tests",
			Description:      "We currently hardcode the version of the dotnet SDK in tests.  This breaks everytime we forget to bump this after a new dotnet SDK release.  We should try to stop hardcoding this, so we don't have issues with the merge queue being blocked every time we do a release and one of the providers requires the new version.",
			AdditionalLabels: []string{"area/testing", "area/codegen"},
		}
		createWorkItem(ctx, repo, pulumiDotnetIssue)

		goTransformationFlaky := &IssueFromFile{
			Title:           "TestGoTransformations/go/simple",
			DescriptionPath: "TestGoTransformations.md",
		}
		createFlakyTestIssue(ctx, repo, goTransformationFlaky)

		return nil
	})
}
