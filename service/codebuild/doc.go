// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codebuild provides the client and types for making API
// requests to AWS CodeBuild.
//
// AWS CodeBuild is a fully managed build service in the cloud. AWS CodeBuild
// compiles your source code, runs unit tests, and produces artifacts that are
// ready to deploy. AWS CodeBuild eliminates the need to provision, manage,
// and scale your own build servers. It provides prepackaged build environments
// for the most popular programming languages and build tools, such as Apache
// Maven, Gradle, and more. You can also fully customize build environments
// in AWS CodeBuild to use your own build tools. AWS CodeBuild scales automatically
// to meet peak build requests. You pay only for the build time you consume.
// For more information about AWS CodeBuild, see the AWS CodeBuild User Guide.
//
// AWS CodeBuild supports these operations:
//
//    * BatchDeleteBuilds: Deletes one or more builds.
//
//    * BatchGetProjects: Gets information about one or more build projects.
//    A build project defines how AWS CodeBuild runs a build. This includes
//    information such as where to get the source code to build, the build environment
//    to use, the build commands to run, and where to store the build output.
//    A build environment is a representation of operating system, programming
//    language runtime, and tools that AWS CodeBuild uses to run a build. You
//    can add tags to build projects to help manage your resources and costs.
//
//    * CreateProject: Creates a build project.
//
//    * CreateWebhook: For an existing AWS CodeBuild build project that has
//    its source code stored in a GitHub or Bitbucket repository, enables AWS
//    CodeBuild to start rebuilding the source code every time a code change
//    is pushed to the repository.
//
//    * UpdateWebhook: Changes the settings of an existing webhook.
//
//    * DeleteProject: Deletes a build project.
//
//    * DeleteWebhook: For an existing AWS CodeBuild build project that has
//    its source code stored in a GitHub or Bitbucket repository, stops AWS
//    CodeBuild from rebuilding the source code every time a code change is
//    pushed to the repository.
//
//    * ListProjects: Gets a list of build project names, with each build project
//    name representing a single build project.
//
//    * UpdateProject: Changes the settings of an existing build project.
//
//    * BatchGetBuilds: Gets information about one or more builds.
//
//    * ListBuilds: Gets a list of build IDs, with each build ID representing
//    a single build.
//
//    * ListBuildsForProject: Gets a list of build IDs for the specified build
//    project, with each build ID representing a single build.
//
//    * StartBuild: Starts running a build.
//
//    * StopBuild: Attempts to stop running a build.
//
//    * ListCuratedEnvironmentImages: Gets information about Docker images that
//    are managed by AWS CodeBuild.
//
// See https://docs.aws.amazon.com/goto/WebAPI/codebuild-2016-10-06 for more information on this service.
//
// See codebuild package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codebuild/
//
// Using the Client
//
// To AWS CodeBuild with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS CodeBuild client CodeBuild for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codebuild/#New
package codebuild
