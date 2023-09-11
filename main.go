package main

import (
    "github.com/aws/constructs-go/constructs/v10"
    "github.com/aws/jsii-runtime-go"
    "github.com/hashicorp/terraform-cdk-go/cdktf"

    "github.com/hashicorp/cdktf-provider-docker-go/docker/v3/container"
    "github.com/hashicorp/cdktf-provider-docker-go/docker/v3/image"
    dockerprovider "github.com/hashicorp/cdktf-provider-docker-go/docker/v3/provider"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
    stack := cdktf.NewTerraformStack(scope, &id)

    dockerprovider.NewDockerProvider(stack, jsii.String("docker"), &dockerprovider.DockerProviderConfig{})

    dockerImage := image.NewImage(stack, jsii.String("nginxImage"), &image.ImageConfig{
        Name:        jsii.String("nginx:latest"),
        KeepLocally: jsii.Bool(false),
    })

    container.NewContainer(stack, jsii.String("nginxContainer"), &container.ContainerConfig{
        Image: dockerImage.Name(),
        Name:  jsii.String("tutorial"),
        Ports: &[]*container.ContainerPorts{{
            Internal: jsii.Number(80), External: jsii.Number(8000),
        }},
    })

    return stack
}

func main() {
    app := cdktf.NewApp(nil)

    NewMyStack(app, "learn-cdktf-docker")

    app.Synth()
}
