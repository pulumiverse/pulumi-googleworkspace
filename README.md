# GoogleWorkspace Resource Provider

The GoogleWorkspace Resource Provider lets you manage [GoogleWorkspace](https://workspace.google.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/googleworkspace
```

or `yarn`:

```bash
yarn add @pulumi/googleworkspace
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_googleworkspace
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumi/pulumi-googleworkspace/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.GoogleWorkspace
```

## Configuration

The following configuration points are available for the `GoogleWorkspace` provider:

- `googleworkspace:accessToken` -  A temporary OAuth 2.0 access token obtained from the Google Authorization server, i.e. the `Authorization: Bearer` token used to authenticate HTTP requests to Google Admin SDK APIs. This is an alternative to `credentials`, and ignores the `scopes` field. If both are specified, `access_token` will be used over the `credentials` field.
- `googleworkspace:credentials` (environment: `GOOGLEWORKSPACE_CREDENTIALS`) - Either the path to or the contents of a service account key file in JSON format you can manage key files using the Cloud Console). If not provided, the application default credentials will be used.
- `googleworkspace:customerId` (environment: `GOOGLEWORKSPACE_CUSTOMER_ID`) - The customer id provided with your Google Workspace subscription. It is found in the admin console under Account Settings.
- `googleworkspace:impersonatedUserEmail` (environment: `GOOGLEWORKSPACE_IMPERSONATED_USER_EMAIL`) - The impersonated user's email with access to the Admin APIs can access the Admin SDK Directory API. `impersonated_user_email` is required for all services except group and user management.
- `googleworkspace:oauthScopes` - The list of the scopes required for your application (for a list of possible scopes, see [Authorize requests](https://developers.google.com/admin-sdk/directory/v1/guides/authorizing))
- `googleworkspace:serviceAccount` - The service account used to create the provided `access_token` if authenticating using the `access_token` method and needing to impersonate a user. This service account will require the GCP role `Service Account Token Creator` if needing to impersonate a user.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/googleworkspace/api-docs/).
