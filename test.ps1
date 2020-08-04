#!/usr/bin/env pwsh

Set-StrictMode -Version latest
$ErrorActionPreference = "Stop"

# Generate an image name using the data in the "component.json" file
$component = Get-Content -Path "component.json" | ConvertFrom-Json
$image="$($component.registry)/$($component.name):$($component.version)-test"

# Set environment variables
$env:IMAGE = $image

try {
    # Workaround to remove dangling images
    docker-compose -f ./docker/docker-compose.test.yml down

    docker-compose -f ./docker/docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from test
} finally {
    # Save the "try" result to avoid overwriting it with the "down" command below
    $exitCode = $LastExitCode 

    docker-compose -f ./docker/docker-compose.test.yml down
}

# Return the exit code of the "docker-compose.test.yml up" command
exit $exitCode 