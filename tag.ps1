#!/usr/bin/env pwsh

Set-StrictMode -Version latest
$ErrorActionPreference = "Stop"

# Generate a tag name using the data in the "component.json" file
$component = Get-Content -Path "component.json" | ConvertFrom-Json
$tag="v$($component.version)-$($component.build)"

# Apply the tag and push to git
git tag $tag
git push --tags