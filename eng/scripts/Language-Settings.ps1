$Language = "go"
$packagePattern = "go.mod"
# rewrite from ChangeLog-Operations.ps1 used in Get-ChangeLogEntriesFromContent for go uses vx.x.x as version number
$RELEASE_TITLE_REGEX = "(?<releaseNoteTitle>^\#+\s+(?<version>v$([AzureEngSemanticVersion]::SEMVER_REGEX))(\s+(?<releaseStatus>\(.+\))))"

# rewrite from artifact-metadata-parsing.ps1 used in RetrievePackages for fetch go single module info
function Get-go-PackageInfoFromPackageFile ($pkg, $workingDirectory)
{
    $workFolder = $pkg.Directory
    $releaseNotes = ""

    if ($workFolder -match "sdk[\/|\\]((?<repopath>(?<service>.*?)[\/|\\])?(?<arm>arm)?(?<pkgname>.*))")
    {
        if ($matches["arm"])
        {
            $packageName = "arm" + $matches["pkgname"]
            $rpName = $matches["service"]
            $pkgId = "sdk/$rpName/$packageName"
        }
        else
        {
            $packageName = $matches["pkgname"]
            $pkgId = "sdk/$packageName"
        }
    }

    $pkgVersion = Get-Version $workFolder

    $changeLogLoc = @(Get-ChildItem -Path $workFolder -Recurse -Include "CHANGELOG.md")[0]
    if ($changeLogLoc)
    {
        $releaseNotes = Get-ChangeLogEntryAsString -ChangeLogLocation $changeLogLoc -VersionString v$pkgVersion
    }

    $resultObj = New-Object PSObject -Property @{
        PackageId      = $pkgId
        PackageVersion = $pkgVersion
        ReleaseTag     = "$pkgId/v$pkgVersion"
        Deployable     = $true
        ReleaseNotes   = $releaseNotes
    }

    return $resultObj
}

# get version from specific files (*constants.go, *version.go)
function Get-Version ($pkgPath)
{
    # find any file with suffix
    $versionFiles = @()
    $version_file_suffixs = "*constants.go", "*version.go"
    foreach ($versionFileSuffix in $version_file_suffixs)
    {
        Get-ChildItem -Recurse -Path $pkgPath -Filter $versionFileSuffix | ForEach-Object {
            Write-Host "Adding $_ to list of version files"
            $versionFiles += $_
        }
    }
    
    # for each version file, use regex to search go version num
    $go_version_regex = ".+\s*=\s*`".*v?(?<version>$([AzureEngSemanticVersion]::SEMVER_REGEX))`""
    foreach ($versionFile in $versionFiles)
    {
        try
        {
            $content = Get-Content $versionFile -Raw
            # finding where the version number are
            if ($content -match $go_version_regex)
            {
                return $matches["version"]
            }
        }
        catch
        {
            Write-Error "Error parsing version."
            Write-Error $_
        }
    }

    Write-Host "Cannot find release version."
    exit 1
}

# rewrite from artifact-metadata-parsing.ps1 as all 0.x.x version should be considered prerelease in go sdk
function CreateReleases($pkgList, $releaseApiUrl, $releaseSha)
{
    foreach ($pkgInfo in $pkgList)
    {
        Write-Host "Creating release $($pkgInfo.Tag)"
  
        $releaseNotes = ""
        if ($null -ne $pkgInfo.ReleaseNotes)
        {
            $releaseNotes = $pkgInfo.ReleaseNotes
        }
  
        $isPrerelease = $False
  
        $parsedSemver = [AzureEngSemanticVersion]::ParseVersionString($pkgInfo.PackageVersion)
  
        if ($parsedSemver -and ($parsedSemver.IsPrerelease -or ($parsedSemver.Major -eq 0)))
        {
            $isPrerelease = $true
        }
  
        $url = $releaseApiUrl
        $body = ConvertTo-Json @{
            tag_name         = $pkgInfo.Tag
            target_commitish = $releaseSha
            name             = $pkgInfo.Tag
            draft            = $False
            prerelease       = $isPrerelease
            body             = $releaseNotes
        }
  
        $headers = @{
            "Content-Type"  = "application/json"
            "Authorization" = "token $env:GH_TOKEN"
        }

        Invoke-RestMethod -Uri $url -Body $body -Headers $headers -Method "Post" -MaximumRetryCount 3 -RetryIntervalSec 10
    }
}