#!/usr/bin/env pwsh

Set-StrictMode -Version latest
$ErrorActionPreference = "Stop"

# Generate an image name and set the registry using the data in the "component.json" file
$component = Get-Content -Path "component.json" | ConvertFrom-Json
$image="$($component.registry)/$($component.name):$($component.version)-$($component.build)-rc"
$registry=$component.registry

# Automatically login to the server
if ($env:DOCKER_USER -ne $null -and $env:DOCKER_PASS -ne $null) {
    docker login $registry -u $env:DOCKER_USER -p $env:DOCKER_PASS
}

# Publish the image
docker push $image

