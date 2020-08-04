#!/usr/bin/env pwsh

# Recreate image names using the data in the "component.json" file
$component = Get-Content -Path "component.json" | ConvertFrom-Json
$buildImage="$($component.registry)/$($component.name):$($component.version)-build"
$testImage="$($component.registry)/$($component.name):$($component.version)-test"
$rcImage="$($component.registry)/$($component.name):$($component.version)-$($component.build)-rc"

# Remove docker images
docker rmi $buildImage --force
docker rmi $testImage --force
docker rmi $rcImage --force
docker image prune --force

# Remove existed containers
docker ps -a | Select-String -Pattern "Exit" | foreach($_) { docker rm $_.ToString().Split(" ")[0] }
