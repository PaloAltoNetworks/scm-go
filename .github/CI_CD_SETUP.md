# CI/CD Setup Guide

## Overview

This repository uses GitHub Actions to automatically run tests on every push or pull request to the `develop` branch. The tests require Strata Cloud Manager (SCM) API credentials, which are stored securely as GitHub Secrets.

## GitHub Secrets Configuration

### Required Secrets

You need to configure the following 8 secrets in your GitHub repository:

| Secret Name | Description | Example Value |
|------------|-------------|---------------|
| `SCM_AUTH_URL` | SCM OAuth2 authentication URL | `https://auth.apps.paloaltonetworks.com/oauth2/access_token` |
| `SCM_CLIENT_ID` | OAuth2 client ID | Your client ID from SCM |
| `SCM_CLIENT_SECRET` | OAuth2 client secret | Your client secret from SCM |
| `SCM_HOST` | SCM API host | `api.strata.paloaltonetworks.com` |
| `SCM_SCOPE` | OAuth2 scope with TSG ID | `tsg_id:1234567890` |
| `SCM_LOGGING` | Logging level for API calls | `debug` or `quiet` |
| `SCM_PROTOCOL` | Protocol for API calls | `https` |
| `SCM_SKIP_VERIFY_CERTIFICATE` | Skip SSL certificate verification | `false` or `true` |

### How to Add Secrets to GitHub

1. Go to your GitHub repository
2. Click on **Settings** → **Secrets and variables** → **Actions**
3. Click **New repository secret**
4. Add each secret one by one:
   - **Name**: Enter the secret name (e.g., `SCM_AUTH_URL`)
   - **Secret**: Paste the actual value
   - Click **Add secret**

## Workflow Details

### Trigger Events

The workflow runs automatically on:
- **Push** to the `develop` branch
- **Pull Request** targeting the `develop` branch

### What the Workflow Does

1. **Checkout code**: Pulls the latest code from the repository
2. **Set up Go**: Installs Go 1.24.0 (required version for this project)
3. **Download dependencies**: Runs `go mod download`
4. **Create config file**: Generates `config/scm-config.json` from GitHub Secrets
5. **Run tests**: Executes all tests with `go test -v ./...`
6. **Report results**: Provides test summary in the GitHub Actions UI

### Timeout

Tests have a 30-minute timeout. If tests take longer, the workflow will fail.

## Viewing Test Results

1. Go to your repository on GitHub
2. Click the **Actions** tab
3. Click on any workflow run to see details
4. Expand the **Run all tests** step to see test output

## Troubleshooting

### Tests Failing Due to Authentication

**Error**: `401 Unauthorized` or authentication errors

**Solution**: Verify that all GitHub Secrets are correctly configured:
- Check for typos in secret names
- Ensure secret values don't have extra spaces or newlines
- Verify credentials are still valid in SCM console

### Tests Timing Out

**Error**: Workflow exceeds 30-minute timeout

**Solution**:
- Check if SCM API is experiencing slowness
- Consider splitting tests into multiple jobs if needed

### Config File Not Found

**Error**: `config/scm-config.json: no such file or directory`

**Solution**: The workflow creates this file automatically. Ensure the "Create SCM config file from secrets" step completes successfully.

## Future Enhancements

Planned improvements:
- Email notifications to commit author on test failure
- Slack channel notifications for test results
- Parallel test execution by service domain
- Test coverage reporting

## Local Testing

To run tests locally with the same configuration:

1. Ensure you have `config/scm-config.json` with valid credentials
2. Run: `go test -v ./...`

Note: Never commit `config/scm-config.json` with real credentials to the repository.
