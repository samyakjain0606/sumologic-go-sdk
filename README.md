# Go API client for openapi

# Getting Started
Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on the collector and search APIs, see our [API home page](https://help.sumologic.com/docs/api).
## API Endpoints
Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment, see [API endpoints](https://help.sumologic.com/?cid=3011).

  <table>
    <tr>
      <td> <strong>Deployment</strong> </td>
      <td> <strong>Endpoint</strong> </td>
    </tr>
    <tr>
      <td> AU </td>
      <td> https://api.au.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> CA </td>
      <td> https://api.ca.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> DE </td>
      <td> https://api.de.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> EU </td>
      <td> https://api.eu.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> FED </td>
      <td> https://api.fed.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> IN </td>
      <td> https://api.in.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> JP </td>
      <td> https://api.jp.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> US1 </td>
      <td> https://api.sumologic.com/api/ </td>
    </tr>
    <tr>
      <td> US2 </td>
      <td> https://api.us2.sumologic.com/api/ </td>
    </tr>
  </table>

## Authentication
Sumo Logic supports the following options for API authentication:
- Access ID and Access Key
- Base64 encoded Access ID and Access Key

See [Access Keys](https://help.sumologic.com/docs/manage/security/access-keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once.
When you have an Access ID and Access Key you can execute requests such as the following:
  ```bash
  curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users
  ```

Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.

If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:
  ```bash
  curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users
  ```


Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.

## Status Codes
Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.
  <table>
    <tr>
      <td> <strong>HTTP Status Code</strong> </td>
      <td> <strong>Error Code</strong> </td>
      <td> <strong>Description</strong> </td>
    </tr>
    <tr>
      <td> 301 </td>
      <td> moved </td>
      <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---301-error---moved) for details.</td>
    </tr>
    <tr>
      <td> 401 </td>
      <td> unauthorized </td>
      <td> Credential could not be verified.</td>
    </tr>
    <tr>
      <td> 403 </td>
      <td> forbidden </td>
      <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/docs/api/troubleshooting/#api---401-error---credential-could-not-be-verified) for details.</td>
    </tr>
    <tr>
      <td> 404 </td>
      <td> notfound </td>
      <td> Requested resource could not be found. </td>
    </tr>
    <tr>
      <td> 405 </td>
      <td> method.unsupported </td>
      <td> Unsupported method for URL. </td>
    </tr>
    <tr>
      <td> 415 </td>
      <td> contenttype.invalid </td>
      <td> Invalid content type. </td>
    </tr>
    <tr>
      <td> 429 </td>
      <td> rate.limit.exceeded </td>
      <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>
    </tr>
    <tr>
      <td> 500 </td>
      <td> internal.error </td>
      <td> Internal server error. </td>
    </tr>
    <tr>
      <td> 503 </td>
      <td> service.unavailable </td>
      <td> Service is currently unavailable. </td>
    </tr>
  </table>

## Filtering
Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.

For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:
  ```bash
  api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe
  ```

## Sorting
Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.

For example, to get 20 users sorted by their `email` in descending order:
  ```bash
  api.sumologic.com/v1/users?limit=20&sort=-email
  ```

## Asynchronous Request
Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request.
1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies
  your asynchronous job.

2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous
  request will typically provide an endpoint to poll for the status of asynchronous job. A successful response
  from the status endpoint will have the following structure:
  ```json
  {
      \"status\": \"Status of asynchronous request\",
      \"statusMessage\": \"Optional message with additional information in case request succeeds\",
      \"error\": \"Error object in case request fails\"
  }
  ```
  The `status` field can have one of the following values:
    1. `Success`: The job succeeded. The `statusMessage` field might have additional information.
    2. `InProgress`: The job is still running.
    3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.

3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))
  to fetch the result of an asynchronous job.


### Example
Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`.
1. Start an export job for the folder
  ```bash
  curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export
  ```
  See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and
  `deployment`.
  On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.
  ```bash
  {
      \"id\": \"C03E086C137F38B4\"
  }
  ```

2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.
  ```bash
  curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status
  ```
  You may get a response like
  ```json
  {
      \"status\": \"InProgress\",
      \"statusMessage\": null,
      \"error\": null
  }
  ```
  It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.

3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the
  [export result](#operation/getAsyncExportResult) endpoint.
  ```bash
  curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result
  ```

  The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.
  ```json
  {
      \"status\": \"Failed\",
      \"errors\": {
          \"code\": \"content1:too_many_items\",
          \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"
      }
  }
  ```


## Rate Limiting
* A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user.
* A rate limit of 10 concurrent requests to any API endpoint applies to an access key.

If a rate is exceeded, a rate limit exceeded 429 status code is returned.

## Generating Clients
You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.

### Using [NPM](https://www.npmjs.com/get-npm)
1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI
  on the command line:
  ```bash
  npm install @openapitools/openapi-generator-cli -g
  ```
  You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).

2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`.
3. Use the following command to generate `python` client inside the `sumo/client/python` directory:
  ```bash
  openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python
  ```


### Using [Homebrew](https://brew.sh/)
1. Install OpenAPI Generator
  ```bash
  brew install openapi-generator
  ```

2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`.
3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:
  ```bash
  openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python
  ```


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.au.sumologic.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessKeyManagementAPI* | [**CreateAccessKey**](docs/AccessKeyManagementAPI.md#createaccesskey) | **Post** /v1/accessKeys | Create an access key.
*AccessKeyManagementAPI* | [**DeleteAccessKey**](docs/AccessKeyManagementAPI.md#deleteaccesskey) | **Delete** /v1/accessKeys/{id} | Delete an access key.
*AccessKeyManagementAPI* | [**ListAccessKeys**](docs/AccessKeyManagementAPI.md#listaccesskeys) | **Get** /v1/accessKeys | List all access keys.
*AccessKeyManagementAPI* | [**ListPersonalAccessKeys**](docs/AccessKeyManagementAPI.md#listpersonalaccesskeys) | **Get** /v1/accessKeys/personal | List personal keys.
*AccessKeyManagementAPI* | [**UpdateAccessKey**](docs/AccessKeyManagementAPI.md#updateaccesskey) | **Put** /v1/accessKeys/{id} | Update an access key.
*AccountManagementAPI* | [**CreateSubdomain**](docs/AccountManagementAPI.md#createsubdomain) | **Post** /v1/account/subdomain | Create account subdomain.
*AccountManagementAPI* | [**DeletePendingUpdateRequest**](docs/AccountManagementAPI.md#deletependingupdaterequest) | **Delete** /v1/plan/pendingUpdateRequest | Delete the pending plan update request, if any.
*AccountManagementAPI* | [**DeleteSubdomain**](docs/AccountManagementAPI.md#deletesubdomain) | **Delete** /v1/account/subdomain | Delete the configured subdomain.
*AccountManagementAPI* | [**ExportUsageReport**](docs/AccountManagementAPI.md#exportusagereport) | **Post** /v1/account/usage/report | Export credits usage details as CSV.
*AccountManagementAPI* | [**GetAccountOwner**](docs/AccountManagementAPI.md#getaccountowner) | **Get** /v1/account/accountOwner | Get the owner of an account.
*AccountManagementAPI* | [**GetPendingUpdateRequest**](docs/AccountManagementAPI.md#getpendingupdaterequest) | **Get** /v1/plan/pendingUpdateRequest | Get the pending plan update request, if any.
*AccountManagementAPI* | [**GetStatus**](docs/AccountManagementAPI.md#getstatus) | **Get** /v1/account/status | Get overview of the account status.
*AccountManagementAPI* | [**GetStatusForReport**](docs/AccountManagementAPI.md#getstatusforreport) | **Get** /v1/account/usage/report/{jobId}/status | Get report generation status.
*AccountManagementAPI* | [**GetSubdomain**](docs/AccountManagementAPI.md#getsubdomain) | **Get** /v1/account/subdomain | Get the configured subdomain.
*AccountManagementAPI* | [**RecoverSubdomains**](docs/AccountManagementAPI.md#recoversubdomains) | **Post** /v1/account/subdomain/recover | Recover subdomains for a user.
*AccountManagementAPI* | [**UpdateSubdomain**](docs/AccountManagementAPI.md#updatesubdomain) | **Put** /v1/account/subdomain | Update account subdomain.
*AppManagementAPI* | [**GetApp**](docs/AppManagementAPI.md#getapp) | **Get** /v1/apps/{uuid} | Get an app by UUID.
*AppManagementAPI* | [**GetAsyncInstallStatus**](docs/AppManagementAPI.md#getasyncinstallstatus) | **Get** /v1/apps/install/{jobId}/status | App install job status.
*AppManagementAPI* | [**InstallApp**](docs/AppManagementAPI.md#installapp) | **Post** /v1/apps/{uuid}/install | Install an app by UUID.
*AppManagementAPI* | [**ListApps**](docs/AppManagementAPI.md#listapps) | **Get** /v1/apps | List available apps.
*ArchiveManagementAPI* | [**CreateArchiveJob**](docs/ArchiveManagementAPI.md#createarchivejob) | **Post** /v1/archive/{sourceId}/jobs | Create an ingestion job.
*ArchiveManagementAPI* | [**DeleteArchiveJob**](docs/ArchiveManagementAPI.md#deletearchivejob) | **Delete** /v1/archive/{sourceId}/jobs/{id} | Delete an ingestion job.
*ArchiveManagementAPI* | [**ListArchiveJobsBySourceId**](docs/ArchiveManagementAPI.md#listarchivejobsbysourceid) | **Get** /v1/archive/{sourceId}/jobs | Get ingestion jobs for an Archive Source.
*ArchiveManagementAPI* | [**ListArchiveJobsCountPerSource**](docs/ArchiveManagementAPI.md#listarchivejobscountpersource) | **Get** /v1/archive/jobs/count | List ingestion jobs for all Archive Sources.
*ConnectionManagementAPI* | [**CreateConnection**](docs/ConnectionManagementAPI.md#createconnection) | **Post** /v1/connections | Create a new connection.
*ConnectionManagementAPI* | [**DeleteConnection**](docs/ConnectionManagementAPI.md#deleteconnection) | **Delete** /v1/connections/{id} | Delete a connection.
*ConnectionManagementAPI* | [**GetConnection**](docs/ConnectionManagementAPI.md#getconnection) | **Get** /v1/connections/{id} | Get a connection.
*ConnectionManagementAPI* | [**GetIncidentTemplates**](docs/ConnectionManagementAPI.md#getincidenttemplates) | **Post** /v1/connections/incidentTemplates | Get incident templates for CloudSOAR connections.
*ConnectionManagementAPI* | [**ListConnections**](docs/ConnectionManagementAPI.md#listconnections) | **Get** /v1/connections | Get a list of connections.
*ConnectionManagementAPI* | [**TestConnection**](docs/ConnectionManagementAPI.md#testconnection) | **Post** /v1/connections/test | Test a new connection url.
*ConnectionManagementAPI* | [**UpdateConnection**](docs/ConnectionManagementAPI.md#updateconnection) | **Put** /v1/connections/{id} | Update a connection.
*ContentManagementAPI* | [**AsyncCopyStatus**](docs/ContentManagementAPI.md#asynccopystatus) | **Get** /v2/content/{id}/copy/{jobId}/status | Content copy job status.
*ContentManagementAPI* | [**BeginAsyncCopy**](docs/ContentManagementAPI.md#beginasynccopy) | **Post** /v2/content/{id}/copy | Start a content copy job.
*ContentManagementAPI* | [**BeginAsyncDelete**](docs/ContentManagementAPI.md#beginasyncdelete) | **Delete** /v2/content/{id}/delete | Start a content deletion job.
*ContentManagementAPI* | [**BeginAsyncExport**](docs/ContentManagementAPI.md#beginasyncexport) | **Post** /v2/content/{id}/export | Start a content export job.
*ContentManagementAPI* | [**BeginAsyncImport**](docs/ContentManagementAPI.md#beginasyncimport) | **Post** /v2/content/folders/{folderId}/import | Start a content import job.
*ContentManagementAPI* | [**GetAsyncDeleteStatus**](docs/ContentManagementAPI.md#getasyncdeletestatus) | **Get** /v2/content/{id}/delete/{jobId}/status | Content deletion job status.
*ContentManagementAPI* | [**GetAsyncExportResult**](docs/ContentManagementAPI.md#getasyncexportresult) | **Get** /v2/content/{contentId}/export/{jobId}/result | Content export job result.
*ContentManagementAPI* | [**GetAsyncExportStatus**](docs/ContentManagementAPI.md#getasyncexportstatus) | **Get** /v2/content/{contentId}/export/{jobId}/status | Content export job status.
*ContentManagementAPI* | [**GetAsyncImportStatus**](docs/ContentManagementAPI.md#getasyncimportstatus) | **Get** /v2/content/folders/{folderId}/import/{jobId}/status | Content import job status.
*ContentManagementAPI* | [**GetItemByPath**](docs/ContentManagementAPI.md#getitembypath) | **Get** /v2/content/path | Get content item by path.
*ContentManagementAPI* | [**GetPathById**](docs/ContentManagementAPI.md#getpathbyid) | **Get** /v2/content/{contentId}/path | Get path of an item.
*ContentManagementAPI* | [**MoveItem**](docs/ContentManagementAPI.md#moveitem) | **Post** /v2/content/{id}/move | Move an item.
*ContentPermissionsAPI* | [**AddContentPermissions**](docs/ContentPermissionsAPI.md#addcontentpermissions) | **Put** /v2/content/{id}/permissions/add | Add permissions to a content item.
*ContentPermissionsAPI* | [**GetContentPermissions**](docs/ContentPermissionsAPI.md#getcontentpermissions) | **Get** /v2/content/{id}/permissions | Get permissions of a content item
*ContentPermissionsAPI* | [**RemoveContentPermissions**](docs/ContentPermissionsAPI.md#removecontentpermissions) | **Put** /v2/content/{id}/permissions/remove | Remove permissions from a content item.
*DashboardManagementAPI* | [**CreateDashboard**](docs/DashboardManagementAPI.md#createdashboard) | **Post** /v2/dashboards | Create a new dashboard.
*DashboardManagementAPI* | [**DeleteDashboard**](docs/DashboardManagementAPI.md#deletedashboard) | **Delete** /v2/dashboards/{id} | Delete a dashboard.
*DashboardManagementAPI* | [**GenerateDashboardReport**](docs/DashboardManagementAPI.md#generatedashboardreport) | **Post** /v2/dashboards/reportJobs | Start a report job
*DashboardManagementAPI* | [**GetAsyncReportGenerationResult**](docs/DashboardManagementAPI.md#getasyncreportgenerationresult) | **Get** /v2/dashboards/reportJobs/{jobId}/result | Get report generation job result
*DashboardManagementAPI* | [**GetAsyncReportGenerationStatus**](docs/DashboardManagementAPI.md#getasyncreportgenerationstatus) | **Get** /v2/dashboards/reportJobs/{jobId}/status | Get report generation job status
*DashboardManagementAPI* | [**GetDashboard**](docs/DashboardManagementAPI.md#getdashboard) | **Get** /v2/dashboards/{id} | Get a dashboard.
*DashboardManagementAPI* | [**GetDashboardMigrationResult**](docs/DashboardManagementAPI.md#getdashboardmigrationresult) | **Get** /v2/dashboards/migrate/{jobId}/result | Get dashboard migration result.
*DashboardManagementAPI* | [**GetDashboardMigrationStatus**](docs/DashboardManagementAPI.md#getdashboardmigrationstatus) | **Get** /v2/dashboards/migrate/{jobId}/status | Get dashboard migration status.
*DashboardManagementAPI* | [**ListDashboards**](docs/DashboardManagementAPI.md#listdashboards) | **Get** /v2/dashboards | List all dashboards.
*DashboardManagementAPI* | [**MigrateReportToDashboard**](docs/DashboardManagementAPI.md#migratereporttodashboard) | **Post** /v2/dashboards/migrate | Migrate Legacy Dashboards to Dashboards(New)
*DashboardManagementAPI* | [**PreviewMigrateReportToDashboard**](docs/DashboardManagementAPI.md#previewmigratereporttodashboard) | **Post** /v2/dashboards/migrate/preview | Preview of Migrating Legacy Dashboards to Dashboards(New)
*DashboardManagementAPI* | [**UpdateDashboard**](docs/DashboardManagementAPI.md#updatedashboard) | **Put** /v2/dashboards/{id} | Update a dashboard.
*DynamicParsingRuleManagementAPI* | [**CreateDynamicParsingRule**](docs/DynamicParsingRuleManagementAPI.md#createdynamicparsingrule) | **Post** /v1/dynamicParsingRules | Create a new dynamic parsing rule.
*DynamicParsingRuleManagementAPI* | [**DeleteDynamicParsingRule**](docs/DynamicParsingRuleManagementAPI.md#deletedynamicparsingrule) | **Delete** /v1/dynamicParsingRules/{id} | Delete a dynamic parsing rule.
*DynamicParsingRuleManagementAPI* | [**GetDynamicParsingRule**](docs/DynamicParsingRuleManagementAPI.md#getdynamicparsingrule) | **Get** /v1/dynamicParsingRules/{id} | Get a dynamic parsing rule.
*DynamicParsingRuleManagementAPI* | [**ListDynamicParsingRules**](docs/DynamicParsingRuleManagementAPI.md#listdynamicparsingrules) | **Get** /v1/dynamicParsingRules | Get a list of dynamic parsing rules.
*DynamicParsingRuleManagementAPI* | [**UpdateDynamicParsingRule**](docs/DynamicParsingRuleManagementAPI.md#updatedynamicparsingrule) | **Put** /v1/dynamicParsingRules/{id} | Update a dynamic parsing rule.
*ExtractionRuleManagementAPI* | [**CreateExtractionRule**](docs/ExtractionRuleManagementAPI.md#createextractionrule) | **Post** /v1/extractionRules | Create a new field extraction rule.
*ExtractionRuleManagementAPI* | [**DeleteExtractionRule**](docs/ExtractionRuleManagementAPI.md#deleteextractionrule) | **Delete** /v1/extractionRules/{id} | Delete a field extraction rule.
*ExtractionRuleManagementAPI* | [**GetExtractionRule**](docs/ExtractionRuleManagementAPI.md#getextractionrule) | **Get** /v1/extractionRules/{id} | Get a field extraction rule.
*ExtractionRuleManagementAPI* | [**ListExtractionRules**](docs/ExtractionRuleManagementAPI.md#listextractionrules) | **Get** /v1/extractionRules | Get a list of field extraction rules.
*ExtractionRuleManagementAPI* | [**UpdateExtractionRule**](docs/ExtractionRuleManagementAPI.md#updateextractionrule) | **Put** /v1/extractionRules/{id} | Update a field extraction rule.
*FieldManagementV1API* | [**CreateField**](docs/FieldManagementV1API.md#createfield) | **Post** /v1/fields | Create a new field.
*FieldManagementV1API* | [**DeleteField**](docs/FieldManagementV1API.md#deletefield) | **Delete** /v1/fields/{id} | Delete a custom field.
*FieldManagementV1API* | [**DisableField**](docs/FieldManagementV1API.md#disablefield) | **Delete** /v1/fields/{id}/disable | Disable a custom field.
*FieldManagementV1API* | [**EnableField**](docs/FieldManagementV1API.md#enablefield) | **Put** /v1/fields/{id}/enable | Enable custom field with a specified identifier.
*FieldManagementV1API* | [**GetBuiltInField**](docs/FieldManagementV1API.md#getbuiltinfield) | **Get** /v1/fields/builtin/{id} | Get a built-in field.
*FieldManagementV1API* | [**GetCustomField**](docs/FieldManagementV1API.md#getcustomfield) | **Get** /v1/fields/{id} | Get a custom field.
*FieldManagementV1API* | [**GetFieldQuota**](docs/FieldManagementV1API.md#getfieldquota) | **Get** /v1/fields/quota | Get capacity information.
*FieldManagementV1API* | [**ListBuiltInFields**](docs/FieldManagementV1API.md#listbuiltinfields) | **Get** /v1/fields/builtin | Get a list of built-in fields.
*FieldManagementV1API* | [**ListCustomFields**](docs/FieldManagementV1API.md#listcustomfields) | **Get** /v1/fields | Get a list of all custom fields.
*FieldManagementV1API* | [**ListDroppedFields**](docs/FieldManagementV1API.md#listdroppedfields) | **Get** /v1/fields/dropped | Get a list of dropped fields.
*FolderManagementAPI* | [**CreateFolder**](docs/FolderManagementAPI.md#createfolder) | **Post** /v2/content/folders | Create a new folder.
*FolderManagementAPI* | [**GetAdminRecommendedFolderAsync**](docs/FolderManagementAPI.md#getadminrecommendedfolderasync) | **Get** /v2/content/folders/adminRecommended | Schedule Admin Recommended folder job
*FolderManagementAPI* | [**GetAdminRecommendedFolderAsyncResult**](docs/FolderManagementAPI.md#getadminrecommendedfolderasyncresult) | **Get** /v2/content/folders/adminRecommended/{jobId}/result | Get Admin Recommended folder job result
*FolderManagementAPI* | [**GetAdminRecommendedFolderAsyncStatus**](docs/FolderManagementAPI.md#getadminrecommendedfolderasyncstatus) | **Get** /v2/content/folders/adminRecommended/{jobId}/status | Get Admin Recommended folder job status
*FolderManagementAPI* | [**GetFolder**](docs/FolderManagementAPI.md#getfolder) | **Get** /v2/content/folders/{id} | Get a folder.
*FolderManagementAPI* | [**GetGlobalFolderAsync**](docs/FolderManagementAPI.md#getglobalfolderasync) | **Get** /v2/content/folders/global | Schedule Global View job
*FolderManagementAPI* | [**GetGlobalFolderAsyncResult**](docs/FolderManagementAPI.md#getglobalfolderasyncresult) | **Get** /v2/content/folders/global/{jobId}/result | Get Global View job result
*FolderManagementAPI* | [**GetGlobalFolderAsyncStatus**](docs/FolderManagementAPI.md#getglobalfolderasyncstatus) | **Get** /v2/content/folders/global/{jobId}/status | Get Global View job status
*FolderManagementAPI* | [**GetPersonalFolder**](docs/FolderManagementAPI.md#getpersonalfolder) | **Get** /v2/content/folders/personal | Get personal folder.
*FolderManagementAPI* | [**UpdateFolder**](docs/FolderManagementAPI.md#updatefolder) | **Put** /v2/content/folders/{id} | Update a folder.
*HealthEventsAPI* | [**ListAllHealthEvents**](docs/HealthEventsAPI.md#listallhealthevents) | **Get** /v1/healthEvents | Get a list of health events.
*HealthEventsAPI* | [**ListAllHealthEventsForResources**](docs/HealthEventsAPI.md#listallhealtheventsforresources) | **Post** /v1/healthEvents/resources | Health events for specific resources.
*IngestBudgetManagementV1API* | [**AssignCollectorToBudget**](docs/IngestBudgetManagementV1API.md#assigncollectortobudget) | **Put** /v1/ingestBudgets/{id}/collectors/{collectorId} | Assign a Collector to a budget.
*IngestBudgetManagementV1API* | [**CreateIngestBudget**](docs/IngestBudgetManagementV1API.md#createingestbudget) | **Post** /v1/ingestBudgets | Create a new ingest budget.
*IngestBudgetManagementV1API* | [**DeleteIngestBudget**](docs/IngestBudgetManagementV1API.md#deleteingestbudget) | **Delete** /v1/ingestBudgets/{id} | Delete an ingest budget.
*IngestBudgetManagementV1API* | [**GetAssignedCollectors**](docs/IngestBudgetManagementV1API.md#getassignedcollectors) | **Get** /v1/ingestBudgets/{id}/collectors | Get a list of Collectors.
*IngestBudgetManagementV1API* | [**GetIngestBudget**](docs/IngestBudgetManagementV1API.md#getingestbudget) | **Get** /v1/ingestBudgets/{id} | Get an ingest budget.
*IngestBudgetManagementV1API* | [**ListIngestBudgets**](docs/IngestBudgetManagementV1API.md#listingestbudgets) | **Get** /v1/ingestBudgets | Get a list of ingest budgets.
*IngestBudgetManagementV1API* | [**RemoveCollectorFromBudget**](docs/IngestBudgetManagementV1API.md#removecollectorfrombudget) | **Delete** /v1/ingestBudgets/{id}/collectors/{collectorId} | Remove Collector from a budget.
*IngestBudgetManagementV1API* | [**ResetUsage**](docs/IngestBudgetManagementV1API.md#resetusage) | **Post** /v1/ingestBudgets/{id}/usage/reset | Reset usage.
*IngestBudgetManagementV1API* | [**UpdateIngestBudget**](docs/IngestBudgetManagementV1API.md#updateingestbudget) | **Put** /v1/ingestBudgets/{id} | Update an ingest budget.
*IngestBudgetManagementV2API* | [**CreateIngestBudgetV2**](docs/IngestBudgetManagementV2API.md#createingestbudgetv2) | **Post** /v2/ingestBudgets | Create a new ingest budget.
*IngestBudgetManagementV2API* | [**DeleteIngestBudgetV2**](docs/IngestBudgetManagementV2API.md#deleteingestbudgetv2) | **Delete** /v2/ingestBudgets/{id} | Delete an ingest budget.
*IngestBudgetManagementV2API* | [**GetIngestBudgetV2**](docs/IngestBudgetManagementV2API.md#getingestbudgetv2) | **Get** /v2/ingestBudgets/{id} | Get an ingest budget.
*IngestBudgetManagementV2API* | [**ListIngestBudgetsV2**](docs/IngestBudgetManagementV2API.md#listingestbudgetsv2) | **Get** /v2/ingestBudgets | Get a list of ingest budgets.
*IngestBudgetManagementV2API* | [**ResetUsageV2**](docs/IngestBudgetManagementV2API.md#resetusagev2) | **Post** /v2/ingestBudgets/{id}/usage/reset | Reset usage.
*IngestBudgetManagementV2API* | [**UpdateIngestBudgetV2**](docs/IngestBudgetManagementV2API.md#updateingestbudgetv2) | **Put** /v2/ingestBudgets/{id} | Update an ingest budget.
*LogSearchesEstimatedUsageAPI* | [**GetLogSearchEstimatedUsage**](docs/LogSearchesEstimatedUsageAPI.md#getlogsearchestimatedusage) | **Post** /v1/logSearches/estimatedUsage | Gets estimated usage details.
*LogSearchesEstimatedUsageAPI* | [**GetLogSearchEstimatedUsageByTier**](docs/LogSearchesEstimatedUsageAPI.md#getlogsearchestimatedusagebytier) | **Post** /v1/logSearches/estimatedUsageByTier | Gets Tier Wise estimated usage details.
*LogSearchesManagementAPI* | [**CreateLogSearch**](docs/LogSearchesManagementAPI.md#createlogsearch) | **Post** /v1/logSearches | Save a log search.
*LogSearchesManagementAPI* | [**DeleteLogSearch**](docs/LogSearchesManagementAPI.md#deletelogsearch) | **Delete** /v1/logSearches/{id} | Delete the saved log search.
*LogSearchesManagementAPI* | [**GetLogSearch**](docs/LogSearchesManagementAPI.md#getlogsearch) | **Get** /v1/logSearches/{id} | Get the saved log search.
*LogSearchesManagementAPI* | [**ListLogSearches**](docs/LogSearchesManagementAPI.md#listlogsearches) | **Get** /v1/logSearches | List all saved log searches.
*LogSearchesManagementAPI* | [**UpdateLogSearch**](docs/LogSearchesManagementAPI.md#updatelogsearch) | **Put** /v1/logSearches/{id} | Update the saved log Search.
*LogsDataForwardingManagementAPI* | [**CreateDataForwardingBucket**](docs/LogsDataForwardingManagementAPI.md#createdataforwardingbucket) | **Post** /v1/logsDataForwarding/destinations | Create an S3 data forwarding destination.
*LogsDataForwardingManagementAPI* | [**CreateDataForwardingRule**](docs/LogsDataForwardingManagementAPI.md#createdataforwardingrule) | **Post** /v1/logsDataForwarding/rules | Create an S3 data forwarding rule.
*LogsDataForwardingManagementAPI* | [**DeleteDataForwardingBucket**](docs/LogsDataForwardingManagementAPI.md#deletedataforwardingbucket) | **Delete** /v1/logsDataForwarding/destinations/{id} | Delete an S3 data forwarding destination.
*LogsDataForwardingManagementAPI* | [**DeleteDataForwardingRule**](docs/LogsDataForwardingManagementAPI.md#deletedataforwardingrule) | **Delete** /v1/logsDataForwarding/rules/{indexId} | Delete an S3 data forwarding rule by its index.
*LogsDataForwardingManagementAPI* | [**GetDataForwardingBuckets**](docs/LogsDataForwardingManagementAPI.md#getdataforwardingbuckets) | **Get** /v1/logsDataForwarding/destinations | Get Amazon S3 data forwarding destinations.
*LogsDataForwardingManagementAPI* | [**GetDataForwardingDestination**](docs/LogsDataForwardingManagementAPI.md#getdataforwardingdestination) | **Get** /v1/logsDataForwarding/destinations/{id} | Get an S3 data forwarding destination.
*LogsDataForwardingManagementAPI* | [**GetDataForwardingRule**](docs/LogsDataForwardingManagementAPI.md#getdataforwardingrule) | **Get** /v1/logsDataForwarding/rules/{indexId} | Get an S3 data forwarding rule by its index.
*LogsDataForwardingManagementAPI* | [**GetRulesAndBuckets**](docs/LogsDataForwardingManagementAPI.md#getrulesandbuckets) | **Get** /v1/logsDataForwarding/rules | Get all S3 data forwarding rules.
*LogsDataForwardingManagementAPI* | [**UpdateDataForwardingBucket**](docs/LogsDataForwardingManagementAPI.md#updatedataforwardingbucket) | **Put** /v1/logsDataForwarding/destinations/{id} | Update an S3 data forwarding destination.
*LogsDataForwardingManagementAPI* | [**UpdateDataForwardingRule**](docs/LogsDataForwardingManagementAPI.md#updatedataforwardingrule) | **Put** /v1/logsDataForwarding/rules/{indexId} | Update an S3 data forwarding rule by its index.
*LookupManagementAPI* | [**CreateTable**](docs/LookupManagementAPI.md#createtable) | **Post** /v1/lookupTables | Create a lookup table.
*LookupManagementAPI* | [**DeleteTable**](docs/LookupManagementAPI.md#deletetable) | **Delete** /v1/lookupTables/{id} | Delete a lookup table.
*LookupManagementAPI* | [**DeleteTableRow**](docs/LookupManagementAPI.md#deletetablerow) | **Put** /v1/lookupTables/{id}/deleteTableRow | Delete a lookup table row.
*LookupManagementAPI* | [**LookupTableById**](docs/LookupManagementAPI.md#lookuptablebyid) | **Get** /v1/lookupTables/{id} | Get a lookup table.
*LookupManagementAPI* | [**RequestJobStatus**](docs/LookupManagementAPI.md#requestjobstatus) | **Get** /v1/lookupTables/jobs/{jobId}/status | Get the status of an async job.
*LookupManagementAPI* | [**TruncateTable**](docs/LookupManagementAPI.md#truncatetable) | **Post** /v1/lookupTables/{id}/truncate | Empty a lookup table.
*LookupManagementAPI* | [**UpdateTable**](docs/LookupManagementAPI.md#updatetable) | **Put** /v1/lookupTables/{id} | Edit a lookup table.
*LookupManagementAPI* | [**UpdateTableRow**](docs/LookupManagementAPI.md#updatetablerow) | **Put** /v1/lookupTables/{id}/row | Insert or Update a lookup table row.
*LookupManagementAPI* | [**UploadFile**](docs/LookupManagementAPI.md#uploadfile) | **Post** /v1/lookupTables/{id}/upload | Upload a CSV file.
*MetricsQueryAPI* | [**RunMetricsQueries**](docs/MetricsQueryAPI.md#runmetricsqueries) | **Post** /v1/metricsQueries | Run metrics queries
*MetricsSearchesManagementAPI* | [**CreateMetricsSearch**](docs/MetricsSearchesManagementAPI.md#createmetricssearch) | **Post** /v1/metricsSearches | Save a metrics search.
*MetricsSearchesManagementAPI* | [**DeleteMetricsSearch**](docs/MetricsSearchesManagementAPI.md#deletemetricssearch) | **Delete** /v1/metricsSearches/{id} | Deletes a metrics search.
*MetricsSearchesManagementAPI* | [**GetMetricsSearch**](docs/MetricsSearchesManagementAPI.md#getmetricssearch) | **Get** /v1/metricsSearches/{id} | Get a metrics search.
*MetricsSearchesManagementAPI* | [**UpdateMetricsSearch**](docs/MetricsSearchesManagementAPI.md#updatemetricssearch) | **Put** /v1/metricsSearches/{id} | Updates a metrics search.
*MonitorsLibraryManagementAPI* | [**DisableMonitorByIds**](docs/MonitorsLibraryManagementAPI.md#disablemonitorbyids) | **Put** /v1/monitors/disable | Disable monitors.
*MonitorsLibraryManagementAPI* | [**GetMonitorUsageInfo**](docs/MonitorsLibraryManagementAPI.md#getmonitorusageinfo) | **Get** /v1/monitors/usageInfo | Usage info of monitors.
*MonitorsLibraryManagementAPI* | [**GetMonitorsFullPath**](docs/MonitorsLibraryManagementAPI.md#getmonitorsfullpath) | **Get** /v1/monitors/{id}/path | Get the path of a monitor or folder.
*MonitorsLibraryManagementAPI* | [**GetMonitorsLibraryRoot**](docs/MonitorsLibraryManagementAPI.md#getmonitorslibraryroot) | **Get** /v1/monitors/root | Get the root monitors folder.
*MonitorsLibraryManagementAPI* | [**MonitorsCopy**](docs/MonitorsLibraryManagementAPI.md#monitorscopy) | **Post** /v1/monitors/{id}/copy | Copy a monitor or folder.
*MonitorsLibraryManagementAPI* | [**MonitorsCreate**](docs/MonitorsLibraryManagementAPI.md#monitorscreate) | **Post** /v1/monitors | Create a monitor or folder. 
*MonitorsLibraryManagementAPI* | [**MonitorsDeleteById**](docs/MonitorsLibraryManagementAPI.md#monitorsdeletebyid) | **Delete** /v1/monitors/{id} | Delete a monitor or folder. 
*MonitorsLibraryManagementAPI* | [**MonitorsDeleteByIds**](docs/MonitorsLibraryManagementAPI.md#monitorsdeletebyids) | **Delete** /v1/monitors | Bulk delete a monitor or folder. 
*MonitorsLibraryManagementAPI* | [**MonitorsExportItem**](docs/MonitorsLibraryManagementAPI.md#monitorsexportitem) | **Get** /v1/monitors/{id}/export | Export a monitor or folder.
*MonitorsLibraryManagementAPI* | [**MonitorsGetByPath**](docs/MonitorsLibraryManagementAPI.md#monitorsgetbypath) | **Get** /v1/monitors/path | Read a monitor or folder by its path.
*MonitorsLibraryManagementAPI* | [**MonitorsImportItem**](docs/MonitorsLibraryManagementAPI.md#monitorsimportitem) | **Post** /v1/monitors/{parentId}/import | Import a monitor or folder.
*MonitorsLibraryManagementAPI* | [**MonitorsMove**](docs/MonitorsLibraryManagementAPI.md#monitorsmove) | **Post** /v1/monitors/{id}/move | Move a monitor or folder.
*MonitorsLibraryManagementAPI* | [**MonitorsReadById**](docs/MonitorsLibraryManagementAPI.md#monitorsreadbyid) | **Get** /v1/monitors/{id} | Get a monitor or folder.
*MonitorsLibraryManagementAPI* | [**MonitorsReadByIds**](docs/MonitorsLibraryManagementAPI.md#monitorsreadbyids) | **Get** /v1/monitors | Bulk read a monitor or folder.
*MonitorsLibraryManagementAPI* | [**MonitorsReadPermissionSummariesByIdGroupBySubjects**](docs/MonitorsLibraryManagementAPI.md#monitorsreadpermissionsummariesbyidgroupbysubjects) | **Get** /v1/monitors/{id}/permissionSummariesBySubjects | List permission summaries for a monitor or folder. 
*MonitorsLibraryManagementAPI* | [**MonitorsReadPermissionsById**](docs/MonitorsLibraryManagementAPI.md#monitorsreadpermissionsbyid) | **Get** /v1/monitors/{id}/permissions | List explicit permissions on monitor or folder. 
*MonitorsLibraryManagementAPI* | [**MonitorsRevokePermissions**](docs/MonitorsLibraryManagementAPI.md#monitorsrevokepermissions) | **Put** /v1/monitors/permissions/revoke | Revoke all permissions on monitor or folder. 
*MonitorsLibraryManagementAPI* | [**MonitorsSearch**](docs/MonitorsLibraryManagementAPI.md#monitorssearch) | **Get** /v1/monitors/search | Search for a monitor or folder.
*MonitorsLibraryManagementAPI* | [**MonitorsSetPermissions**](docs/MonitorsLibraryManagementAPI.md#monitorssetpermissions) | **Put** /v1/monitors/permissions/set | Set permissions on monitor or folder. 
*MonitorsLibraryManagementAPI* | [**MonitorsUpdateById**](docs/MonitorsLibraryManagementAPI.md#monitorsupdatebyid) | **Put** /v1/monitors/{id} | Update a monitor or folder. 
*MutingSchedulesLibraryManagementAPI* | [**GetMutingSchedulesFullPath**](docs/MutingSchedulesLibraryManagementAPI.md#getmutingschedulesfullpath) | **Get** /v1/mutingSchedules/{id}/path | Get the path of a mutingschedule or folder.
*MutingSchedulesLibraryManagementAPI* | [**GetMutingSchedulesLibraryRoot**](docs/MutingSchedulesLibraryManagementAPI.md#getmutingscheduleslibraryroot) | **Get** /v1/mutingSchedules/root | Get the root mutingSchedules folder.
*MutingSchedulesLibraryManagementAPI* | [**MutingSchedulesCopy**](docs/MutingSchedulesLibraryManagementAPI.md#mutingschedulescopy) | **Post** /v1/mutingSchedules/{id}/copy | Copy a mutingschedule or folder.
*MutingSchedulesLibraryManagementAPI* | [**MutingSchedulesCreate**](docs/MutingSchedulesLibraryManagementAPI.md#mutingschedulescreate) | **Post** /v1/mutingSchedules | Create a mutingschedule or folder. 
*MutingSchedulesLibraryManagementAPI* | [**MutingSchedulesDeleteById**](docs/MutingSchedulesLibraryManagementAPI.md#mutingschedulesdeletebyid) | **Delete** /v1/mutingSchedules/{id} | Delete a mutingschedule or folder. 
*MutingSchedulesLibraryManagementAPI* | [**MutingSchedulesDeleteByIds**](docs/MutingSchedulesLibraryManagementAPI.md#mutingschedulesdeletebyids) | **Delete** /v1/mutingSchedules | Bulk delete a mutingschedule or folder. 
*MutingSchedulesLibraryManagementAPI* | [**MutingSchedulesExportItem**](docs/MutingSchedulesLibraryManagementAPI.md#mutingschedulesexportitem) | **Get** /v1/mutingSchedules/{id}/export | Export a mutingschedule or folder.
*MutingSchedulesLibraryManagementAPI* | [**MutingSchedulesImportItem**](docs/MutingSchedulesLibraryManagementAPI.md#mutingschedulesimportitem) | **Post** /v1/mutingSchedules/{parentId}/import | Import a mutingschedule or folder.
*MutingSchedulesLibraryManagementAPI* | [**MutingSchedulesReadById**](docs/MutingSchedulesLibraryManagementAPI.md#mutingschedulesreadbyid) | **Get** /v1/mutingSchedules/{id} | Get a mutingschedule or folder.
*MutingSchedulesLibraryManagementAPI* | [**MutingSchedulesReadByIds**](docs/MutingSchedulesLibraryManagementAPI.md#mutingschedulesreadbyids) | **Get** /v1/mutingSchedules | Bulk read a mutingschedule or folder.
*MutingSchedulesLibraryManagementAPI* | [**MutingSchedulesSearch**](docs/MutingSchedulesLibraryManagementAPI.md#mutingschedulessearch) | **Get** /v1/mutingSchedules/search | Search for a mutingschedule or folder.
*MutingSchedulesLibraryManagementAPI* | [**MutingSchedulesUpdateById**](docs/MutingSchedulesLibraryManagementAPI.md#mutingschedulesupdatebyid) | **Put** /v1/mutingSchedules/{id} | Update a mutingschedule or folder. 
*OrgsManagementAPI* | [**GetChildUsages**](docs/OrgsManagementAPI.md#getchildusages) | **Post** /v1/organizations/usages | Get usages for child orgs.
*OtCollectorManagementExternalAPI* | [**DeleteOTCollector**](docs/OtCollectorManagementExternalAPI.md#deleteotcollector) | **Delete** /v1/otCollectors/{id} | Delete an OT Collector.
*OtCollectorManagementExternalAPI* | [**DeleteOfflineOTCollectors**](docs/OtCollectorManagementExternalAPI.md#deleteofflineotcollectors) | **Delete** /v1/otCollectors/offline | Delete all Offline OT Collectors
*OtCollectorManagementExternalAPI* | [**GetOTCollector**](docs/OtCollectorManagementExternalAPI.md#getotcollector) | **Get** /v1/otCollectors/{id} | Get OT Collector by ID.
*OtCollectorManagementExternalAPI* | [**GetOTCollectorsByNames**](docs/OtCollectorManagementExternalAPI.md#getotcollectorsbynames) | **Get** /v1/otCollectors/otCollectorsByName | Get OT Collectors by name.
*OtCollectorManagementExternalAPI* | [**GetOTCollectorsCount**](docs/OtCollectorManagementExternalAPI.md#getotcollectorscount) | **Get** /v1/otCollectors/totalCount | Get a count of OT Collectors.
*OtCollectorManagementExternalAPI* | [**GetPaginatedOTCollectors**](docs/OtCollectorManagementExternalAPI.md#getpaginatedotcollectors) | **Post** /v1/otCollectors | Get paginated list of OT Collectors
*PartitionManagementAPI* | [**CancelRetentionUpdate**](docs/PartitionManagementAPI.md#cancelretentionupdate) | **Post** /v1/partitions/{id}/cancelRetentionUpdate | Cancel a retention update for a partition
*PartitionManagementAPI* | [**CreatePartition**](docs/PartitionManagementAPI.md#createpartition) | **Post** /v1/partitions | Create a new partition.
*PartitionManagementAPI* | [**DecommissionPartition**](docs/PartitionManagementAPI.md#decommissionpartition) | **Post** /v1/partitions/{id}/decommission | Decommission a partition.
*PartitionManagementAPI* | [**GetPartition**](docs/PartitionManagementAPI.md#getpartition) | **Get** /v1/partitions/{id} | Get a partition.
*PartitionManagementAPI* | [**ListPartitions**](docs/PartitionManagementAPI.md#listpartitions) | **Get** /v1/partitions | Get a list of partitions.
*PartitionManagementAPI* | [**UpdatePartition**](docs/PartitionManagementAPI.md#updatepartition) | **Put** /v1/partitions/{id} | Update a partition.
*PasswordPolicyAPI* | [**GetPasswordPolicy**](docs/PasswordPolicyAPI.md#getpasswordpolicy) | **Get** /v1/passwordPolicy | Get the current password policy.
*PasswordPolicyAPI* | [**SetPasswordPolicy**](docs/PasswordPolicyAPI.md#setpasswordpolicy) | **Put** /v1/passwordPolicy | Update password policy.
*PoliciesManagementAPI* | [**GetAuditPolicy**](docs/PoliciesManagementAPI.md#getauditpolicy) | **Get** /v1/policies/audit | Get Audit policy.
*PoliciesManagementAPI* | [**GetDataAccessLevelPolicy**](docs/PoliciesManagementAPI.md#getdataaccesslevelpolicy) | **Get** /v1/policies/dataAccessLevel | Get Data Access Level policy.
*PoliciesManagementAPI* | [**GetMaxUserSessionTimeoutPolicy**](docs/PoliciesManagementAPI.md#getmaxusersessiontimeoutpolicy) | **Get** /v1/policies/maxUserSessionTimeout | Get Max User Session Timeout policy.
*PoliciesManagementAPI* | [**GetSearchAuditPolicy**](docs/PoliciesManagementAPI.md#getsearchauditpolicy) | **Get** /v1/policies/searchAudit | Get Search Audit policy.
*PoliciesManagementAPI* | [**GetShareDashboardsOutsideOrganizationPolicy**](docs/PoliciesManagementAPI.md#getsharedashboardsoutsideorganizationpolicy) | **Get** /v1/policies/shareDashboardsOutsideOrganization | Get Share Dashboards Outside Organization policy.
*PoliciesManagementAPI* | [**GetUserConcurrentSessionsLimitPolicy**](docs/PoliciesManagementAPI.md#getuserconcurrentsessionslimitpolicy) | **Get** /v1/policies/userConcurrentSessionsLimit | Get User Concurrent Sessions Limit policy.
*PoliciesManagementAPI* | [**SetAuditPolicy**](docs/PoliciesManagementAPI.md#setauditpolicy) | **Put** /v1/policies/audit | Set Audit policy.
*PoliciesManagementAPI* | [**SetDataAccessLevelPolicy**](docs/PoliciesManagementAPI.md#setdataaccesslevelpolicy) | **Put** /v1/policies/dataAccessLevel | Set Data Access Level policy.
*PoliciesManagementAPI* | [**SetMaxUserSessionTimeoutPolicy**](docs/PoliciesManagementAPI.md#setmaxusersessiontimeoutpolicy) | **Put** /v1/policies/maxUserSessionTimeout | Set Max User Session Timeout policy.
*PoliciesManagementAPI* | [**SetSearchAuditPolicy**](docs/PoliciesManagementAPI.md#setsearchauditpolicy) | **Put** /v1/policies/searchAudit | Set Search Audit policy.
*PoliciesManagementAPI* | [**SetShareDashboardsOutsideOrganizationPolicy**](docs/PoliciesManagementAPI.md#setsharedashboardsoutsideorganizationpolicy) | **Put** /v1/policies/shareDashboardsOutsideOrganization | Set Share Dashboards Outside Organization policy.
*PoliciesManagementAPI* | [**SetUserConcurrentSessionsLimitPolicy**](docs/PoliciesManagementAPI.md#setuserconcurrentsessionslimitpolicy) | **Put** /v1/policies/userConcurrentSessionsLimit | Set User Concurrent Sessions Limit policy.
*RoleManagementAPI* | [**AssignRoleToUser**](docs/RoleManagementAPI.md#assignroletouser) | **Put** /v1/roles/{roleId}/users/{userId} | Assign a role to a user.
*RoleManagementAPI* | [**CreateRole**](docs/RoleManagementAPI.md#createrole) | **Post** /v1/roles | Create a new role.
*RoleManagementAPI* | [**DeleteRole**](docs/RoleManagementAPI.md#deleterole) | **Delete** /v1/roles/{id} | Delete a role.
*RoleManagementAPI* | [**GetRole**](docs/RoleManagementAPI.md#getrole) | **Get** /v1/roles/{id} | Get a role.
*RoleManagementAPI* | [**ListRoles**](docs/RoleManagementAPI.md#listroles) | **Get** /v1/roles | Get a list of roles.
*RoleManagementAPI* | [**RemoveRoleFromUser**](docs/RoleManagementAPI.md#removerolefromuser) | **Delete** /v1/roles/{roleId}/users/{userId} | Remove role from a user.
*RoleManagementAPI* | [**UpdateRole**](docs/RoleManagementAPI.md#updaterole) | **Put** /v1/roles/{id} | Update a role.
*SamlConfigurationManagementAPI* | [**CreateAllowlistedUser**](docs/SamlConfigurationManagementAPI.md#createallowlisteduser) | **Post** /v1/saml/allowlistedUsers/{userId} | Allowlist a user.
*SamlConfigurationManagementAPI* | [**CreateIdentityProvider**](docs/SamlConfigurationManagementAPI.md#createidentityprovider) | **Post** /v1/saml/identityProviders | Create a new SAML configuration.
*SamlConfigurationManagementAPI* | [**DeleteAllowlistedUser**](docs/SamlConfigurationManagementAPI.md#deleteallowlisteduser) | **Delete** /v1/saml/allowlistedUsers/{userId} | Remove an allowlisted user.
*SamlConfigurationManagementAPI* | [**DeleteIdentityProvider**](docs/SamlConfigurationManagementAPI.md#deleteidentityprovider) | **Delete** /v1/saml/identityProviders/{id} | Delete a SAML configuration.
*SamlConfigurationManagementAPI* | [**DisableSamlLockdown**](docs/SamlConfigurationManagementAPI.md#disablesamllockdown) | **Post** /v1/saml/lockdown/disable | Disable SAML lockdown.
*SamlConfigurationManagementAPI* | [**EnableSamlLockdown**](docs/SamlConfigurationManagementAPI.md#enablesamllockdown) | **Post** /v1/saml/lockdown/enable | Require SAML for sign-in.
*SamlConfigurationManagementAPI* | [**GetAllowlistedUsers**](docs/SamlConfigurationManagementAPI.md#getallowlistedusers) | **Get** /v1/saml/allowlistedUsers | Get list of allowlisted users.
*SamlConfigurationManagementAPI* | [**GetIdentityProviders**](docs/SamlConfigurationManagementAPI.md#getidentityproviders) | **Get** /v1/saml/identityProviders | Get a list of SAML configurations.
*SamlConfigurationManagementAPI* | [**GetSamlMetadata**](docs/SamlConfigurationManagementAPI.md#getsamlmetadata) | **Get** /v1/saml/identityProviders/{id}/metadata | Get SAML configuration metadata XML.
*SamlConfigurationManagementAPI* | [**UpdateIdentityProvider**](docs/SamlConfigurationManagementAPI.md#updateidentityprovider) | **Put** /v1/saml/identityProviders/{id} | Update a SAML configuration.
*ScheduledViewManagementAPI* | [**CreateScheduledView**](docs/ScheduledViewManagementAPI.md#createscheduledview) | **Post** /v1/scheduledViews | Create a new scheduled view.
*ScheduledViewManagementAPI* | [**DisableScheduledView**](docs/ScheduledViewManagementAPI.md#disablescheduledview) | **Delete** /v1/scheduledViews/{id}/disable | Disable a scheduled view.
*ScheduledViewManagementAPI* | [**GetScheduledView**](docs/ScheduledViewManagementAPI.md#getscheduledview) | **Get** /v1/scheduledViews/{id} | Get a scheduled view.
*ScheduledViewManagementAPI* | [**ListScheduledViews**](docs/ScheduledViewManagementAPI.md#listscheduledviews) | **Get** /v1/scheduledViews | Get a list of scheduled views.
*ScheduledViewManagementAPI* | [**PauseScheduledView**](docs/ScheduledViewManagementAPI.md#pausescheduledview) | **Post** /v1/scheduledViews/{id}/pause | Pause a scheduled view.
*ScheduledViewManagementAPI* | [**StartScheduledView**](docs/ScheduledViewManagementAPI.md#startscheduledview) | **Post** /v1/scheduledViews/{id}/start | Start a scheduled view.
*ScheduledViewManagementAPI* | [**UpdateScheduledView**](docs/ScheduledViewManagementAPI.md#updatescheduledview) | **Put** /v1/scheduledViews/{id} | Update a scheduled view.
*ServiceAllowlistManagementAPI* | [**AddAllowlistedCidrs**](docs/ServiceAllowlistManagementAPI.md#addallowlistedcidrs) | **Post** /v1/serviceAllowlist/addresses/add | Allowlist CIDRs/IP addresses.
*ServiceAllowlistManagementAPI* | [**DeleteAllowlistedCidrs**](docs/ServiceAllowlistManagementAPI.md#deleteallowlistedcidrs) | **Post** /v1/serviceAllowlist/addresses/remove | Remove allowlisted CIDRs/IP addresses.
*ServiceAllowlistManagementAPI* | [**DisableAllowlisting**](docs/ServiceAllowlistManagementAPI.md#disableallowlisting) | **Post** /v1/serviceAllowlist/disable | Disable service allowlisting.
*ServiceAllowlistManagementAPI* | [**EnableAllowlisting**](docs/ServiceAllowlistManagementAPI.md#enableallowlisting) | **Post** /v1/serviceAllowlist/enable | Enable service allowlisting.
*ServiceAllowlistManagementAPI* | [**GetAllowlistingStatus**](docs/ServiceAllowlistManagementAPI.md#getallowlistingstatus) | **Get** /v1/serviceAllowlist/status | Get the allowlisting status.
*ServiceAllowlistManagementAPI* | [**ListAllowlistedCidrs**](docs/ServiceAllowlistManagementAPI.md#listallowlistedcidrs) | **Get** /v1/serviceAllowlist/addresses | List all allowlisted CIDRs/IP addresses.
*ServiceMapAPI* | [**GetServiceMap**](docs/ServiceMapAPI.md#getservicemap) | **Get** /v1/tracing/serviceMap | Get a service map.
*SlosLibraryManagementAPI* | [**GetSloUsageInfo**](docs/SlosLibraryManagementAPI.md#getslousageinfo) | **Get** /v1/slos/usageInfo | Usage info of SLOs.
*SlosLibraryManagementAPI* | [**GetSlosFullPath**](docs/SlosLibraryManagementAPI.md#getslosfullpath) | **Get** /v1/slos/{id}/path | Get the path of a slo or folder.
*SlosLibraryManagementAPI* | [**GetSlosLibraryRoot**](docs/SlosLibraryManagementAPI.md#getsloslibraryroot) | **Get** /v1/slos/root | Get the root slos folder.
*SlosLibraryManagementAPI* | [**Sli**](docs/SlosLibraryManagementAPI.md#sli) | **Get** /v1/slos/sli | Bulk fetch SLI values, error budget remaining and SLI computation status for the current compliance period.
*SlosLibraryManagementAPI* | [**SlosCopy**](docs/SlosLibraryManagementAPI.md#sloscopy) | **Post** /v1/slos/{id}/copy | Copy a slo or folder.
*SlosLibraryManagementAPI* | [**SlosCreate**](docs/SlosLibraryManagementAPI.md#sloscreate) | **Post** /v1/slos | Create a slo or folder. 
*SlosLibraryManagementAPI* | [**SlosDeleteById**](docs/SlosLibraryManagementAPI.md#slosdeletebyid) | **Delete** /v1/slos/{id} | Delete a slo or folder. 
*SlosLibraryManagementAPI* | [**SlosDeleteByIds**](docs/SlosLibraryManagementAPI.md#slosdeletebyids) | **Delete** /v1/slos | Bulk delete a slo or folder. 
*SlosLibraryManagementAPI* | [**SlosExportItem**](docs/SlosLibraryManagementAPI.md#slosexportitem) | **Get** /v1/slos/{id}/export | Export a slo or folder.
*SlosLibraryManagementAPI* | [**SlosGetByPath**](docs/SlosLibraryManagementAPI.md#slosgetbypath) | **Get** /v1/slos/path | Read a slo or folder by its path.
*SlosLibraryManagementAPI* | [**SlosImportItem**](docs/SlosLibraryManagementAPI.md#slosimportitem) | **Post** /v1/slos/{parentId}/import | Import a slo or folder.
*SlosLibraryManagementAPI* | [**SlosMove**](docs/SlosLibraryManagementAPI.md#slosmove) | **Post** /v1/slos/{id}/move | Move a slo or folder.
*SlosLibraryManagementAPI* | [**SlosReadById**](docs/SlosLibraryManagementAPI.md#slosreadbyid) | **Get** /v1/slos/{id} | Get a slo or folder.
*SlosLibraryManagementAPI* | [**SlosReadByIds**](docs/SlosLibraryManagementAPI.md#slosreadbyids) | **Get** /v1/slos | Bulk read a slo or folder.
*SlosLibraryManagementAPI* | [**SlosSearch**](docs/SlosLibraryManagementAPI.md#slossearch) | **Get** /v1/slos/search | Search for a slo or folder.
*SlosLibraryManagementAPI* | [**SlosUpdateById**](docs/SlosLibraryManagementAPI.md#slosupdatebyid) | **Put** /v1/slos/{id} | Update a slo or folder. 
*SpanAnalyticsAPI* | [**CancelSpanQuery**](docs/SpanAnalyticsAPI.md#cancelspanquery) | **Delete** /v1/tracing/spanquery/{queryId} | Cancel a span analytics query.
*SpanAnalyticsAPI* | [**CreateSpanQuery**](docs/SpanAnalyticsAPI.md#createspanquery) | **Post** /v1/tracing/spanquery | Run a span analytics query asynchronously.
*SpanAnalyticsAPI* | [**GetSpanQueryAggregates**](docs/SpanAnalyticsAPI.md#getspanqueryaggregates) | **Get** /v1/tracing/spanquery/{queryId}/aggregates | Get span analytics query aggregated results.
*SpanAnalyticsAPI* | [**GetSpanQueryFacets**](docs/SpanAnalyticsAPI.md#getspanqueryfacets) | **Get** /v1/tracing/spanquery/{queryId}/rows/{rowId}/facets | Get a list of facets of a span analytics query.
*SpanAnalyticsAPI* | [**GetSpanQueryFieldValues**](docs/SpanAnalyticsAPI.md#getspanqueryfieldvalues) | **Get** /v1/tracing/spanquery/fields/{field}/values | Get span analytics query filter field values.
*SpanAnalyticsAPI* | [**GetSpanQueryFields**](docs/SpanAnalyticsAPI.md#getspanqueryfields) | **Get** /v1/tracing/spanquery/fields | Get filter fields for span analytics queries.
*SpanAnalyticsAPI* | [**GetSpanQueryResult**](docs/SpanAnalyticsAPI.md#getspanqueryresult) | **Get** /v1/tracing/spanquery/{queryId}/rows/{rowId}/spans | Get results of a span analytics query.
*SpanAnalyticsAPI* | [**GetSpanQueryStatus**](docs/SpanAnalyticsAPI.md#getspanquerystatus) | **Get** /v1/tracing/spanquery/{queryId}/status | Get a span analytics query status.
*SpanAnalyticsAPI* | [**PauseSpanQuery**](docs/SpanAnalyticsAPI.md#pausespanquery) | **Put** /v1/tracing/spanquery/{queryId}/pause | Pause a span analytics query.
*SpanAnalyticsAPI* | [**ResumeSpanQuery**](docs/SpanAnalyticsAPI.md#resumespanquery) | **Put** /v1/tracing/spanquery/{queryId}/resume | Resume a span analytics query.
*TokensLibraryManagementAPI* | [**CreateToken**](docs/TokensLibraryManagementAPI.md#createtoken) | **Post** /v1/tokens | Create a token.
*TokensLibraryManagementAPI* | [**DeleteToken**](docs/TokensLibraryManagementAPI.md#deletetoken) | **Delete** /v1/tokens/{id} | Delete a token.
*TokensLibraryManagementAPI* | [**GetToken**](docs/TokensLibraryManagementAPI.md#gettoken) | **Get** /v1/tokens/{id} | Get a token.
*TokensLibraryManagementAPI* | [**ListTokens**](docs/TokensLibraryManagementAPI.md#listtokens) | **Get** /v1/tokens | Get a list of tokens.
*TokensLibraryManagementAPI* | [**UpdateToken**](docs/TokensLibraryManagementAPI.md#updatetoken) | **Put** /v1/tokens/{id} | Update a token.
*TracesAPI* | [**CancelTraceQuery**](docs/TracesAPI.md#canceltracequery) | **Delete** /v1/tracing/tracequery/{queryId} | Cancel a trace search query.
*TracesAPI* | [**CreateTraceQuery**](docs/TracesAPI.md#createtracequery) | **Post** /v1/tracing/tracequery | Run a trace search query asynchronously.
*TracesAPI* | [**GetCriticalPath**](docs/TracesAPI.md#getcriticalpath) | **Get** /v1/tracing/traces/{traceId}/criticalPath | Get a critical path of a trace.
*TracesAPI* | [**GetCriticalPathServiceBreakdown**](docs/TracesAPI.md#getcriticalpathservicebreakdown) | **Get** /v1/tracing/traces/{traceId}/criticalPath/breakdown/service | Get a critical path service breakdown of a trace.
*TracesAPI* | [**GetMetrics**](docs/TracesAPI.md#getmetrics) | **Get** /v1/tracing/metrics | Get trace search query metrics.
*TracesAPI* | [**GetSpan**](docs/TracesAPI.md#getspan) | **Get** /v1/tracing/traces/{traceId}/spans/{spanId} | Get span details.
*TracesAPI* | [**GetSpanBillingInfo**](docs/TracesAPI.md#getspanbillinginfo) | **Get** /v1/tracing/traces/{traceId}/spans/{spanId}/billingInfo | Get span billing details.
*TracesAPI* | [**GetSpans**](docs/TracesAPI.md#getspans) | **Get** /v1/tracing/traces/{traceId}/spans | Get a list of trace spans.
*TracesAPI* | [**GetTrace**](docs/TracesAPI.md#gettrace) | **Get** /v1/tracing/traces/{traceId} | Get trace details.
*TracesAPI* | [**GetTraceLightEvents**](docs/TracesAPI.md#gettracelightevents) | **Get** /v1/tracing/traces/{traceId}/traceEvents | Get a list of events (without their attributes) per span for a trace.
*TracesAPI* | [**GetTraceQueryFieldValues**](docs/TracesAPI.md#gettracequeryfieldvalues) | **Get** /v1/tracing/tracequery/fields/{field}/values | Get trace search query filter field values.
*TracesAPI* | [**GetTraceQueryFields**](docs/TracesAPI.md#gettracequeryfields) | **Get** /v1/tracing/tracequery/fields | Get filter fields for trace search queries.
*TracesAPI* | [**GetTraceQueryResult**](docs/TracesAPI.md#gettracequeryresult) | **Get** /v1/tracing/tracequery/{queryId}/rows/{rowId}/traces | Get results of a trace search query.
*TracesAPI* | [**GetTraceQueryStatus**](docs/TracesAPI.md#gettracequerystatus) | **Get** /v1/tracing/tracequery/{queryId}/status | Get a trace search query status.
*TracesAPI* | [**TraceExists**](docs/TracesAPI.md#traceexists) | **Get** /v1/tracing/traces/{traceId}/exists | Check if the trace exists.
*TransformationRuleManagementAPI* | [**CreateRule**](docs/TransformationRuleManagementAPI.md#createrule) | **Post** /v1/transformationRules | Create a new transformation rule.
*TransformationRuleManagementAPI* | [**DeleteRule**](docs/TransformationRuleManagementAPI.md#deleterule) | **Delete** /v1/transformationRules/{id} | Delete a transformation rule.
*TransformationRuleManagementAPI* | [**GetTransformationRule**](docs/TransformationRuleManagementAPI.md#gettransformationrule) | **Get** /v1/transformationRules/{id} | Get a transformation rule.
*TransformationRuleManagementAPI* | [**GetTransformationRules**](docs/TransformationRuleManagementAPI.md#gettransformationrules) | **Get** /v1/transformationRules | Get a list of transformation rules.
*TransformationRuleManagementAPI* | [**UpdateTransformationRule**](docs/TransformationRuleManagementAPI.md#updatetransformationrule) | **Put** /v1/transformationRules/{id} | Update a transformation rule.
*UserManagementAPI* | [**CreateUser**](docs/UserManagementAPI.md#createuser) | **Post** /v1/users | Create a new user.
*UserManagementAPI* | [**DeleteUser**](docs/UserManagementAPI.md#deleteuser) | **Delete** /v1/users/{id} | Delete a user.
*UserManagementAPI* | [**DisableMfa**](docs/UserManagementAPI.md#disablemfa) | **Put** /v1/users/{id}/mfa/disable | Disable MFA for user.
*UserManagementAPI* | [**GetUser**](docs/UserManagementAPI.md#getuser) | **Get** /v1/users/{id} | Get a user.
*UserManagementAPI* | [**ListUsers**](docs/UserManagementAPI.md#listusers) | **Get** /v1/users | Get a list of users.
*UserManagementAPI* | [**RequestChangeEmail**](docs/UserManagementAPI.md#requestchangeemail) | **Post** /v1/users/{id}/email/requestChange | Change email address.
*UserManagementAPI* | [**ResetPassword**](docs/UserManagementAPI.md#resetpassword) | **Post** /v1/users/{id}/password/reset | Reset password.
*UserManagementAPI* | [**UnlockUser**](docs/UserManagementAPI.md#unlockuser) | **Post** /v1/users/{id}/unlock | Unlock a user.
*UserManagementAPI* | [**UpdateUser**](docs/UserManagementAPI.md#updateuser) | **Put** /v1/users/{id} | Update a user.


## Documentation For Models

 - [AWSLambda](docs/AWSLambda.md)
 - [AccessKey](docs/AccessKey.md)
 - [AccessKeyCreateRequest](docs/AccessKeyCreateRequest.md)
 - [AccessKeyPublic](docs/AccessKeyPublic.md)
 - [AccessKeyUpdateRequest](docs/AccessKeyUpdateRequest.md)
 - [AccountStatusResponse](docs/AccountStatusResponse.md)
 - [Action](docs/Action.md)
 - [AddOrReplaceTransformation](docs/AddOrReplaceTransformation.md)
 - [AdhocMutingResponse](docs/AdhocMutingResponse.md)
 - [AgentRemoteConfigStatusTracker](docs/AgentRemoteConfigStatusTracker.md)
 - [AggregateOnTransformation](docs/AggregateOnTransformation.md)
 - [AggregationGroupByAttribute](docs/AggregationGroupByAttribute.md)
 - [AggregationQueryBucketResult](docs/AggregationQueryBucketResult.md)
 - [AggregationQueryRequest](docs/AggregationQueryRequest.md)
 - [AggregationQueryResultResponse](docs/AggregationQueryResultResponse.md)
 - [AggregationQueryRow](docs/AggregationQueryRow.md)
 - [AggregationQueryRowStatus](docs/AggregationQueryRowStatus.md)
 - [AggregationQueryStatusResponse](docs/AggregationQueryStatusResponse.md)
 - [AlertChartDataResult](docs/AlertChartDataResult.md)
 - [AlertChartMetadata](docs/AlertChartMetadata.md)
 - [AlertEntityInfo](docs/AlertEntityInfo.md)
 - [AlertMonitorQuery](docs/AlertMonitorQuery.md)
 - [AlertSearchNotificationSyncDefinition](docs/AlertSearchNotificationSyncDefinition.md)
 - [AlertSignalContext](docs/AlertSignalContext.md)
 - [AlertsLibraryAlert](docs/AlertsLibraryAlert.md)
 - [AlertsLibraryAlertExport](docs/AlertsLibraryAlertExport.md)
 - [AlertsLibraryAlertResponse](docs/AlertsLibraryAlertResponse.md)
 - [AlertsLibraryAlertUpdate](docs/AlertsLibraryAlertUpdate.md)
 - [AlertsLibraryBase](docs/AlertsLibraryBase.md)
 - [AlertsLibraryBaseExport](docs/AlertsLibraryBaseExport.md)
 - [AlertsLibraryBaseResponse](docs/AlertsLibraryBaseResponse.md)
 - [AlertsLibraryBaseUpdate](docs/AlertsLibraryBaseUpdate.md)
 - [AlertsLibraryFolder](docs/AlertsLibraryFolder.md)
 - [AlertsLibraryFolderExport](docs/AlertsLibraryFolderExport.md)
 - [AlertsLibraryFolderResponse](docs/AlertsLibraryFolderResponse.md)
 - [AlertsLibraryFolderUpdate](docs/AlertsLibraryFolderUpdate.md)
 - [AlertsLibraryItemWithPath](docs/AlertsLibraryItemWithPath.md)
 - [AlertsListPageObject](docs/AlertsListPageObject.md)
 - [AlertsListPageResponse](docs/AlertsListPageResponse.md)
 - [AllowlistedUserResult](docs/AllowlistedUserResult.md)
 - [AllowlistingStatus](docs/AllowlistingStatus.md)
 - [AndTracingExpression](docs/AndTracingExpression.md)
 - [AnomalyDataPointsCountRequest](docs/AnomalyDataPointsCountRequest.md)
 - [App](docs/App.md)
 - [AppDefinition](docs/AppDefinition.md)
 - [AppInstallRequest](docs/AppInstallRequest.md)
 - [AppItemsList](docs/AppItemsList.md)
 - [AppListItem](docs/AppListItem.md)
 - [AppManifest](docs/AppManifest.md)
 - [AppRecommendation](docs/AppRecommendation.md)
 - [ArchiveJob](docs/ArchiveJob.md)
 - [ArchiveJobsCount](docs/ArchiveJobsCount.md)
 - [ArrayTracingValue](docs/ArrayTracingValue.md)
 - [AsyncJobStatus](docs/AsyncJobStatus.md)
 - [AsyncTraceQueryRequest](docs/AsyncTraceQueryRequest.md)
 - [AsyncTraceQueryRow](docs/AsyncTraceQueryRow.md)
 - [AttributeReversedIndex](docs/AttributeReversedIndex.md)
 - [AttributeValueReversedIndex](docs/AttributeValueReversedIndex.md)
 - [AuditPolicy](docs/AuditPolicy.md)
 - [AuthnCertificateResult](docs/AuthnCertificateResult.md)
 - [AutoCompleteValueSyncDefinition](docs/AutoCompleteValueSyncDefinition.md)
 - [AwsCloudWatchCollectionErrorTracker](docs/AwsCloudWatchCollectionErrorTracker.md)
 - [AwsInventoryCollectionErrorTracker](docs/AwsInventoryCollectionErrorTracker.md)
 - [AxisRange](docs/AxisRange.md)
 - [AzureEventHubConnectionErrorTracker](docs/AzureEventHubConnectionErrorTracker.md)
 - [AzureFunctions](docs/AzureFunctions.md)
 - [BaseExtractionRuleDefinition](docs/BaseExtractionRuleDefinition.md)
 - [Baselines](docs/Baselines.md)
 - [BeginAsyncJobResponse](docs/BeginAsyncJobResponse.md)
 - [BeginAsyncJobResponseV2](docs/BeginAsyncJobResponseV2.md)
 - [BeginBoundedTimeRange](docs/BeginBoundedTimeRange.md)
 - [BooleanArrayEventAttributeValue](docs/BooleanArrayEventAttributeValue.md)
 - [BooleanEventAttributeValue](docs/BooleanEventAttributeValue.md)
 - [BucketDefinition](docs/BucketDefinition.md)
 - [BucketKey](docs/BucketKey.md)
 - [BucketValue](docs/BucketValue.md)
 - [BuiltinField](docs/BuiltinField.md)
 - [BuiltinFieldUsage](docs/BuiltinFieldUsage.md)
 - [BulkAsyncStatusResponse](docs/BulkAsyncStatusResponse.md)
 - [BulkBeginAsyncJobResponse](docs/BulkBeginAsyncJobResponse.md)
 - [BulkErrorResponse](docs/BulkErrorResponse.md)
 - [BurnRate](docs/BurnRate.md)
 - [CSEWindowsAccessErrorTracker](docs/CSEWindowsAccessErrorTracker.md)
 - [CSEWindowsErrorAppendingToQueueFilesTracker](docs/CSEWindowsErrorAppendingToQueueFilesTracker.md)
 - [CSEWindowsErrorParsingRecordsTracker](docs/CSEWindowsErrorParsingRecordsTracker.md)
 - [CSEWindowsErrorTracker](docs/CSEWindowsErrorTracker.md)
 - [CSEWindowsExcessiveBacklogTracker](docs/CSEWindowsExcessiveBacklogTracker.md)
 - [CSEWindowsExcessiveEventLogMonitorsTracker](docs/CSEWindowsExcessiveEventLogMonitorsTracker.md)
 - [CSEWindowsExcessiveFilesPendingUploadTracker](docs/CSEWindowsExcessiveFilesPendingUploadTracker.md)
 - [CSEWindowsInvalidConfigurationTracker](docs/CSEWindowsInvalidConfigurationTracker.md)
 - [CSEWindowsInvalidUserPermissionsTracker](docs/CSEWindowsInvalidUserPermissionsTracker.md)
 - [CSEWindowsOldestRecordTimestampExceedsThresholdTracker](docs/CSEWindowsOldestRecordTimestampExceedsThresholdTracker.md)
 - [CSEWindowsParsingErrorTracker](docs/CSEWindowsParsingErrorTracker.md)
 - [CSEWindowsRuntimeErrorTracker](docs/CSEWindowsRuntimeErrorTracker.md)
 - [CSEWindowsRuntimeWarningTracker](docs/CSEWindowsRuntimeWarningTracker.md)
 - [CSEWindowsSensorOfflineTracker](docs/CSEWindowsSensorOfflineTracker.md)
 - [CSEWindowsSensorOutOfStorageTracker](docs/CSEWindowsSensorOutOfStorageTracker.md)
 - [CSEWindowsStorageLimitApproachingTracker](docs/CSEWindowsStorageLimitApproachingTracker.md)
 - [CSEWindowsStorageLimitExceededTracker](docs/CSEWindowsStorageLimitExceededTracker.md)
 - [CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker](docs/CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker.md)
 - [CalculatorRequest](docs/CalculatorRequest.md)
 - [CalendarCompliance](docs/CalendarCompliance.md)
 - [CapabilityDefinition](docs/CapabilityDefinition.md)
 - [CapabilityDefinitionGroup](docs/CapabilityDefinitionGroup.md)
 - [CapabilityList](docs/CapabilityList.md)
 - [CapabilityMap](docs/CapabilityMap.md)
 - [Capacity](docs/Capacity.md)
 - [ChangeEmailRequest](docs/ChangeEmailRequest.md)
 - [ChartDataRequest](docs/ChartDataRequest.md)
 - [ChartDataResult](docs/ChartDataResult.md)
 - [ChildUsage](docs/ChildUsage.md)
 - [ChildUsageDetail](docs/ChildUsageDetail.md)
 - [ChildUsageDetailsRequest](docs/ChildUsageDetailsRequest.md)
 - [ChildUsageDetailsResponse](docs/ChildUsageDetailsResponse.md)
 - [Cidr](docs/Cidr.md)
 - [CidrList](docs/CidrList.md)
 - [CollectionAffectedDueToIngestBudgetTracker](docs/CollectionAffectedDueToIngestBudgetTracker.md)
 - [CollectionAwsInventoryThrottledTracker](docs/CollectionAwsInventoryThrottledTracker.md)
 - [CollectionAwsInventoryUnauthorizedTracker](docs/CollectionAwsInventoryUnauthorizedTracker.md)
 - [CollectionAwsMetadataTagsFetchDeniedTracker](docs/CollectionAwsMetadataTagsFetchDeniedTracker.md)
 - [CollectionCloudWatchGetStatisticsDeniedTracker](docs/CollectionCloudWatchGetStatisticsDeniedTracker.md)
 - [CollectionCloudWatchGetStatisticsThrottledTracker](docs/CollectionCloudWatchGetStatisticsThrottledTracker.md)
 - [CollectionCloudWatchListMetricsDeniedTracker](docs/CollectionCloudWatchListMetricsDeniedTracker.md)
 - [CollectionCloudWatchTagsFetchDeniedTracker](docs/CollectionCloudWatchTagsFetchDeniedTracker.md)
 - [CollectionDockerClientBuildingFailedTracker](docs/CollectionDockerClientBuildingFailedTracker.md)
 - [CollectionInvalidFilePathTracker](docs/CollectionInvalidFilePathTracker.md)
 - [CollectionPathAccessDeniedTracker](docs/CollectionPathAccessDeniedTracker.md)
 - [CollectionRemoteConnectionFailedTracker](docs/CollectionRemoteConnectionFailedTracker.md)
 - [CollectionS3AccessDeniedTracker](docs/CollectionS3AccessDeniedTracker.md)
 - [CollectionS3GetObjectAccessDeniedTracker](docs/CollectionS3GetObjectAccessDeniedTracker.md)
 - [CollectionS3InvalidKeyTracker](docs/CollectionS3InvalidKeyTracker.md)
 - [CollectionS3ListingFailedTracker](docs/CollectionS3ListingFailedTracker.md)
 - [CollectionS3SlowListingTracker](docs/CollectionS3SlowListingTracker.md)
 - [CollectionWindowsEventChannelConnectionFailedTracker](docs/CollectionWindowsEventChannelConnectionFailedTracker.md)
 - [CollectionWindowsHostConnectionFailedTracker](docs/CollectionWindowsHostConnectionFailedTracker.md)
 - [Collector](docs/Collector.md)
 - [CollectorIdentity](docs/CollectorIdentity.md)
 - [CollectorLimitApproachingTracker](docs/CollectorLimitApproachingTracker.md)
 - [CollectorRegistrationTokenResponse](docs/CollectorRegistrationTokenResponse.md)
 - [CollectorResourceIdentity](docs/CollectorResourceIdentity.md)
 - [ColoringRule](docs/ColoringRule.md)
 - [ColoringThreshold](docs/ColoringThreshold.md)
 - [CompleteLiteralTimeRange](docs/CompleteLiteralTimeRange.md)
 - [Compliance](docs/Compliance.md)
 - [CompliancePeriodProgress](docs/CompliancePeriodProgress.md)
 - [CompliancePeriodRef](docs/CompliancePeriodRef.md)
 - [CompliancePeriods](docs/CompliancePeriods.md)
 - [ConfidenceScoreResponse](docs/ConfidenceScoreResponse.md)
 - [ConfigureSubdomainRequest](docs/ConfigureSubdomainRequest.md)
 - [Connection](docs/Connection.md)
 - [ConnectionDefinition](docs/ConnectionDefinition.md)
 - [Consumable](docs/Consumable.md)
 - [ConsumptionDetails](docs/ConsumptionDetails.md)
 - [Content](docs/Content.md)
 - [ContentCopyParams](docs/ContentCopyParams.md)
 - [ContentList](docs/ContentList.md)
 - [ContentPath](docs/ContentPath.md)
 - [ContentPermissionAssignment](docs/ContentPermissionAssignment.md)
 - [ContentPermissionResult](docs/ContentPermissionResult.md)
 - [ContentPermissionUpdateRequest](docs/ContentPermissionUpdateRequest.md)
 - [ContentSyncDefinition](docs/ContentSyncDefinition.md)
 - [ContractDetails](docs/ContractDetails.md)
 - [ContractPeriod](docs/ContractPeriod.md)
 - [CorrelatedEvent](docs/CorrelatedEvent.md)
 - [CorrelatedEvents](docs/CorrelatedEvents.md)
 - [CpcQueryBucketResult](docs/CpcQueryBucketResult.md)
 - [CpcQueryBucketStatus](docs/CpcQueryBucketStatus.md)
 - [CpcQueryRequest](docs/CpcQueryRequest.md)
 - [CpcQueryResultRequest](docs/CpcQueryResultRequest.md)
 - [CpcQueryResultResponse](docs/CpcQueryResultResponse.md)
 - [CpcQueryRow](docs/CpcQueryRow.md)
 - [CpcQueryRowStatus](docs/CpcQueryRowStatus.md)
 - [CpcQueryStatusResponse](docs/CpcQueryStatusResponse.md)
 - [CpcServiceSummary](docs/CpcServiceSummary.md)
 - [CpcSummary](docs/CpcSummary.md)
 - [CreateAggregationQueryResponse](docs/CreateAggregationQueryResponse.md)
 - [CreateArchiveJobRequest](docs/CreateArchiveJobRequest.md)
 - [CreateBucketDefinition](docs/CreateBucketDefinition.md)
 - [CreateBucketDefinitionItems](docs/CreateBucketDefinitionItems.md)
 - [CreateCpcQueryResponse](docs/CreateCpcQueryResponse.md)
 - [CreateDataForwardingRule](docs/CreateDataForwardingRule.md)
 - [CreatePartitionDefinition](docs/CreatePartitionDefinition.md)
 - [CreateRoleDefinition](docs/CreateRoleDefinition.md)
 - [CreateRoleDefinitionV2](docs/CreateRoleDefinitionV2.md)
 - [CreateScheduledViewDefinition](docs/CreateScheduledViewDefinition.md)
 - [CreateTraceQueryResponse](docs/CreateTraceQueryResponse.md)
 - [CreateUserDefinition](docs/CreateUserDefinition.md)
 - [CreditsBreakdown](docs/CreditsBreakdown.md)
 - [CriticalPathResponse](docs/CriticalPathResponse.md)
 - [CriticalPathServiceBreakdownElementBase](docs/CriticalPathServiceBreakdownElementBase.md)
 - [CriticalPathServiceBreakdownElementDetail](docs/CriticalPathServiceBreakdownElementDetail.md)
 - [CriticalPathServiceBreakdownResponse](docs/CriticalPathServiceBreakdownResponse.md)
 - [CriticalPathServiceBreakdownSummary](docs/CriticalPathServiceBreakdownSummary.md)
 - [CseInsightConfidenceRequest](docs/CseInsightConfidenceRequest.md)
 - [CseSignalNotificationSyncDefinition](docs/CseSignalNotificationSyncDefinition.md)
 - [CsvVariableSourceDefinition](docs/CsvVariableSourceDefinition.md)
 - [CurrentBillingPeriod](docs/CurrentBillingPeriod.md)
 - [CurrentPlan](docs/CurrentPlan.md)
 - [CustomField](docs/CustomField.md)
 - [CustomFieldUsage](docs/CustomFieldUsage.md)
 - [Dashboard](docs/Dashboard.md)
 - [DashboardMigrationRequest](docs/DashboardMigrationRequest.md)
 - [DashboardMigrationResult](docs/DashboardMigrationResult.md)
 - [DashboardMigrationStatus](docs/DashboardMigrationStatus.md)
 - [DashboardReportModeTemplate](docs/DashboardReportModeTemplate.md)
 - [DashboardRequest](docs/DashboardRequest.md)
 - [DashboardSearchResult](docs/DashboardSearchResult.md)
 - [DashboardSearchSessionIds](docs/DashboardSearchSessionIds.md)
 - [DashboardSearchStatus](docs/DashboardSearchStatus.md)
 - [DashboardSyncDefinition](docs/DashboardSyncDefinition.md)
 - [DashboardTemplate](docs/DashboardTemplate.md)
 - [DashboardV2SyncDefinition](docs/DashboardV2SyncDefinition.md)
 - [DataAccessLevelPolicy](docs/DataAccessLevelPolicy.md)
 - [DataForwardingRule](docs/DataForwardingRule.md)
 - [DataIngestAffectedTracker](docs/DataIngestAffectedTracker.md)
 - [DataPoint](docs/DataPoint.md)
 - [DataPointCount](docs/DataPointCount.md)
 - [DataPoints](docs/DataPoints.md)
 - [DataValue](docs/DataValue.md)
 - [Datadog](docs/Datadog.md)
 - [DatastoreRetentionPeriod](docs/DatastoreRetentionPeriod.md)
 - [DatastoreSourceStatusResponse](docs/DatastoreSourceStatusResponse.md)
 - [DatastoreStatusResponse](docs/DatastoreStatusResponse.md)
 - [DateTimeTracingValue](docs/DateTimeTracingValue.md)
 - [DimensionKeyValue](docs/DimensionKeyValue.md)
 - [DimensionTransformation](docs/DimensionTransformation.md)
 - [DirectDownloadReportAction](docs/DirectDownloadReportAction.md)
 - [DisableMfaRequest](docs/DisableMfaRequest.md)
 - [DisableMonitorResponse](docs/DisableMonitorResponse.md)
 - [DisableMonitorWarning](docs/DisableMonitorWarning.md)
 - [DisableUnusedAccessKeysPolicy](docs/DisableUnusedAccessKeysPolicy.md)
 - [DoubleArrayEventAttributeValue](docs/DoubleArrayEventAttributeValue.md)
 - [DoubleEventAttributeValue](docs/DoubleEventAttributeValue.md)
 - [DoubleTracingValue](docs/DoubleTracingValue.md)
 - [DroppedField](docs/DroppedField.md)
 - [DynamicRule](docs/DynamicRule.md)
 - [DynamicRuleDefinition](docs/DynamicRuleDefinition.md)
 - [Email](docs/Email.md)
 - [EmailSearchNotificationSyncDefinition](docs/EmailSearchNotificationSyncDefinition.md)
 - [EndpointDefinition](docs/EndpointDefinition.md)
 - [EndpointResponse](docs/EndpointResponse.md)
 - [EntitlementConsumption](docs/EntitlementConsumption.md)
 - [EntitlementUsage](docs/EntitlementUsage.md)
 - [Entitlements](docs/Entitlements.md)
 - [EpochTimeRangeBoundary](docs/EpochTimeRangeBoundary.md)
 - [ErrorDescription](docs/ErrorDescription.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [EstimatedUsageDetails](docs/EstimatedUsageDetails.md)
 - [EstimatedUsageDetailsWithTier](docs/EstimatedUsageDetailsWithTier.md)
 - [EventAttributeValue](docs/EventAttributeValue.md)
 - [EventContext](docs/EventContext.md)
 - [EventsOfInterestScatterPanel](docs/EventsOfInterestScatterPanel.md)
 - [ExportableLookupTableInfo](docs/ExportableLookupTableInfo.md)
 - [Extension](docs/Extension.md)
 - [ExternalReference](docs/ExternalReference.md)
 - [ExtraDetails](docs/ExtraDetails.md)
 - [ExtractionRule](docs/ExtractionRule.md)
 - [ExtractionRuleDefinition](docs/ExtractionRuleDefinition.md)
 - [FieldName](docs/FieldName.md)
 - [FieldQuotaUsage](docs/FieldQuotaUsage.md)
 - [FieldTracingFilter](docs/FieldTracingFilter.md)
 - [FileCollectionErrorTracker](docs/FileCollectionErrorTracker.md)
 - [FilledRange](docs/FilledRange.md)
 - [Folder](docs/Folder.md)
 - [FolderDefinition](docs/FolderDefinition.md)
 - [FolderSyncDefinition](docs/FolderSyncDefinition.md)
 - [GcpMetricsCollectionBrokenTracker](docs/GcpMetricsCollectionBrokenTracker.md)
 - [GenerateReportRequest](docs/GenerateReportRequest.md)
 - [GetCollectorsUsageResponse](docs/GetCollectorsUsageResponse.md)
 - [GetDataForwardingDestinations](docs/GetDataForwardingDestinations.md)
 - [GetIncidentTemplatesRequest](docs/GetIncidentTemplatesRequest.md)
 - [GetIncidentTemplatesResponse](docs/GetIncidentTemplatesResponse.md)
 - [GetRoleDefinitionV2](docs/GetRoleDefinitionV2.md)
 - [GetRulesAndBucketsResult](docs/GetRulesAndBucketsResult.md)
 - [GetSourcesUsageResponse](docs/GetSourcesUsageResponse.md)
 - [GetViewFilterDefinition](docs/GetViewFilterDefinition.md)
 - [GranularMarkingType](docs/GranularMarkingType.md)
 - [Grid](docs/Grid.md)
 - [GroupDefinition](docs/GroupDefinition.md)
 - [GroupFieldsRequest](docs/GroupFieldsRequest.md)
 - [GroupFieldsResponse](docs/GroupFieldsResponse.md)
 - [Header](docs/Header.md)
 - [HealthEvent](docs/HealthEvent.md)
 - [HighCardinalityDimensionDroppedTracker](docs/HighCardinalityDimensionDroppedTracker.md)
 - [HipChat](docs/HipChat.md)
 - [IncidentTemplate](docs/IncidentTemplate.md)
 - [IngestBudget](docs/IngestBudget.md)
 - [IngestBudgetDefinition](docs/IngestBudgetDefinition.md)
 - [IngestBudgetDefinitionV2](docs/IngestBudgetDefinitionV2.md)
 - [IngestBudgetExceededTracker](docs/IngestBudgetExceededTracker.md)
 - [IngestBudgetResourceIdentity](docs/IngestBudgetResourceIdentity.md)
 - [IngestBudgetV2](docs/IngestBudgetV2.md)
 - [IngestThrottlingTracker](docs/IngestThrottlingTracker.md)
 - [InstalledCollectorOfflineTracker](docs/InstalledCollectorOfflineTracker.md)
 - [IntegerArrayEventAttributeValue](docs/IntegerArrayEventAttributeValue.md)
 - [IntegerEventAttributeValue](docs/IntegerEventAttributeValue.md)
 - [IntegerTracingValue](docs/IntegerTracingValue.md)
 - [Iso8601TimeRange](docs/Iso8601TimeRange.md)
 - [Iso8601TimeRangeBoundary](docs/Iso8601TimeRangeBoundary.md)
 - [Jira](docs/Jira.md)
 - [KeyValuePair](docs/KeyValuePair.md)
 - [KillChainPhase](docs/KillChainPhase.md)
 - [LabelValueLookupAutoCompleteSyncDefinition](docs/LabelValueLookupAutoCompleteSyncDefinition.md)
 - [LabelValuePairsAutoCompleteSyncDefinition](docs/LabelValuePairsAutoCompleteSyncDefinition.md)
 - [Layout](docs/Layout.md)
 - [LayoutStructure](docs/LayoutStructure.md)
 - [LightSpanEvent](docs/LightSpanEvent.md)
 - [LinkedDashboard](docs/LinkedDashboard.md)
 - [ListAccessKeysResult](docs/ListAccessKeysResult.md)
 - [ListAppsResult](docs/ListAppsResult.md)
 - [ListArchiveJobsCount](docs/ListArchiveJobsCount.md)
 - [ListArchiveJobsResponse](docs/ListArchiveJobsResponse.md)
 - [ListBuiltinFieldsResponse](docs/ListBuiltinFieldsResponse.md)
 - [ListBuiltinFieldsUsageResponse](docs/ListBuiltinFieldsUsageResponse.md)
 - [ListCollectorIdentitiesResponse](docs/ListCollectorIdentitiesResponse.md)
 - [ListConnectionsResponse](docs/ListConnectionsResponse.md)
 - [ListCustomFieldsResponse](docs/ListCustomFieldsResponse.md)
 - [ListCustomFieldsUsageResponse](docs/ListCustomFieldsUsageResponse.md)
 - [ListDroppedFieldsResponse](docs/ListDroppedFieldsResponse.md)
 - [ListDynamicRulesResponse](docs/ListDynamicRulesResponse.md)
 - [ListExtractionRulesResponse](docs/ListExtractionRulesResponse.md)
 - [ListFieldNamesResponse](docs/ListFieldNamesResponse.md)
 - [ListHealthEventResponse](docs/ListHealthEventResponse.md)
 - [ListIngestBudgetsResponse](docs/ListIngestBudgetsResponse.md)
 - [ListIngestBudgetsResponseV2](docs/ListIngestBudgetsResponseV2.md)
 - [ListPartitionsInfoResponse](docs/ListPartitionsInfoResponse.md)
 - [ListPartitionsResponse](docs/ListPartitionsResponse.md)
 - [ListPermissionsResponse](docs/ListPermissionsResponse.md)
 - [ListRoleModelsResponse](docs/ListRoleModelsResponse.md)
 - [ListRoleModelsResponseV2](docs/ListRoleModelsResponseV2.md)
 - [ListScheduledViewsResponse](docs/ListScheduledViewsResponse.md)
 - [ListTagResult](docs/ListTagResult.md)
 - [ListTokensBaseResponse](docs/ListTokensBaseResponse.md)
 - [ListUserId](docs/ListUserId.md)
 - [ListUserModelsResponse](docs/ListUserModelsResponse.md)
 - [LiteralTimeRangeBoundary](docs/LiteralTimeRangeBoundary.md)
 - [LogQueryVariableSourceDefinition](docs/LogQueryVariableSourceDefinition.md)
 - [LogSearch](docs/LogSearch.md)
 - [LogSearchDefinition](docs/LogSearchDefinition.md)
 - [LogSearchEstimatedUsageByTierDefinition](docs/LogSearchEstimatedUsageByTierDefinition.md)
 - [LogSearchEstimatedUsageDefinition](docs/LogSearchEstimatedUsageDefinition.md)
 - [LogSearchEstimatedUsageRequest](docs/LogSearchEstimatedUsageRequest.md)
 - [LogSearchEstimatedUsageRequestV2](docs/LogSearchEstimatedUsageRequestV2.md)
 - [LogSearchNotificationThresholdSyncDefinition](docs/LogSearchNotificationThresholdSyncDefinition.md)
 - [LogSearchParameterAutoCompleteSyncDefinition](docs/LogSearchParameterAutoCompleteSyncDefinition.md)
 - [LogSearchQuery](docs/LogSearchQuery.md)
 - [LogSearchQueryParameterSyncDefinition](docs/LogSearchQueryParameterSyncDefinition.md)
 - [LogSearchQueryParameterSyncDefinitionBase](docs/LogSearchQueryParameterSyncDefinitionBase.md)
 - [LogSearchQueryTimeRangeBase](docs/LogSearchQueryTimeRangeBase.md)
 - [LogSearchQueryTimeRangeBaseExceptParsingMode](docs/LogSearchQueryTimeRangeBaseExceptParsingMode.md)
 - [LogSearchScheduleSyncDefinition](docs/LogSearchScheduleSyncDefinition.md)
 - [LogsMissingDataCondition](docs/LogsMissingDataCondition.md)
 - [LogsOutlier](docs/LogsOutlier.md)
 - [LogsOutlierCondition](docs/LogsOutlierCondition.md)
 - [LogsStaticCondition](docs/LogsStaticCondition.md)
 - [LogsToMetricsRuleDisabledTracker](docs/LogsToMetricsRuleDisabledTracker.md)
 - [LogsToMetricsRuleIdentity](docs/LogsToMetricsRuleIdentity.md)
 - [LookupAsyncJobStatus](docs/LookupAsyncJobStatus.md)
 - [LookupPreviewData](docs/LookupPreviewData.md)
 - [LookupRequestToken](docs/LookupRequestToken.md)
 - [LookupTable](docs/LookupTable.md)
 - [LookupTableDefinition](docs/LookupTableDefinition.md)
 - [LookupTableField](docs/LookupTableField.md)
 - [LookupTableSyncDefinition](docs/LookupTableSyncDefinition.md)
 - [LookupTablesLimits](docs/LookupTablesLimits.md)
 - [LookupUpdateDefinition](docs/LookupUpdateDefinition.md)
 - [MaxUserSessionTimeoutPolicy](docs/MaxUserSessionTimeoutPolicy.md)
 - [Metadata](docs/Metadata.md)
 - [MetadataModel](docs/MetadataModel.md)
 - [MetadataVariableSourceDefinition](docs/MetadataVariableSourceDefinition.md)
 - [MetadataWithUserInfo](docs/MetadataWithUserInfo.md)
 - [MetricDefinition](docs/MetricDefinition.md)
 - [MetricNameAsMetatagTracker](docs/MetricNameAsMetatagTracker.md)
 - [MetricNameErrorTracker](docs/MetricNameErrorTracker.md)
 - [MetricNameMissingTracker](docs/MetricNameMissingTracker.md)
 - [MetricTracingFilter](docs/MetricTracingFilter.md)
 - [MetricsCardinalityLimitExceededTracker](docs/MetricsCardinalityLimitExceededTracker.md)
 - [MetricsFilter](docs/MetricsFilter.md)
 - [MetricsHighCardinalityDetectedTracker](docs/MetricsHighCardinalityDetectedTracker.md)
 - [MetricsMetadataKeyLengthLimitExceeded](docs/MetricsMetadataKeyLengthLimitExceeded.md)
 - [MetricsMetadataKeyLengthLimitExceededTracker](docs/MetricsMetadataKeyLengthLimitExceededTracker.md)
 - [MetricsMetadataKeyValuePairsLimitExceeded](docs/MetricsMetadataKeyValuePairsLimitExceeded.md)
 - [MetricsMetadataKeyValuePairsLimitExceededTracker](docs/MetricsMetadataKeyValuePairsLimitExceededTracker.md)
 - [MetricsMetadataLimitsExceededTracker](docs/MetricsMetadataLimitsExceededTracker.md)
 - [MetricsMetadataTotalMetadataSizeLimitExceeded](docs/MetricsMetadataTotalMetadataSizeLimitExceeded.md)
 - [MetricsMetadataTotalMetadataSizeLimitExceededTracker](docs/MetricsMetadataTotalMetadataSizeLimitExceededTracker.md)
 - [MetricsMetadataValueLengthLimitExceeded](docs/MetricsMetadataValueLengthLimitExceeded.md)
 - [MetricsMetadataValueLengthLimitExceededTracker](docs/MetricsMetadataValueLengthLimitExceededTracker.md)
 - [MetricsMissingDataCondition](docs/MetricsMissingDataCondition.md)
 - [MetricsOutlier](docs/MetricsOutlier.md)
 - [MetricsOutlierCondition](docs/MetricsOutlierCondition.md)
 - [MetricsQueryData](docs/MetricsQueryData.md)
 - [MetricsQueryRequest](docs/MetricsQueryRequest.md)
 - [MetricsQueryResponse](docs/MetricsQueryResponse.md)
 - [MetricsQueryResultContext](docs/MetricsQueryResultContext.md)
 - [MetricsQueryResultInfo](docs/MetricsQueryResultInfo.md)
 - [MetricsQueryRow](docs/MetricsQueryRow.md)
 - [MetricsQuerySyncDefinition](docs/MetricsQuerySyncDefinition.md)
 - [MetricsSavedSearchQuerySyncDefinition](docs/MetricsSavedSearchQuerySyncDefinition.md)
 - [MetricsSavedSearchSyncDefinition](docs/MetricsSavedSearchSyncDefinition.md)
 - [MetricsSearchInstance](docs/MetricsSearchInstance.md)
 - [MetricsSearchQuery](docs/MetricsSearchQuery.md)
 - [MetricsSearchSyncDefinition](docs/MetricsSearchSyncDefinition.md)
 - [MetricsSearchV1](docs/MetricsSearchV1.md)
 - [MetricsStaticCondition](docs/MetricsStaticCondition.md)
 - [MewboardSyncDefinition](docs/MewboardSyncDefinition.md)
 - [MicrosoftTeams](docs/MicrosoftTeams.md)
 - [MigratedDashboardInfo](docs/MigratedDashboardInfo.md)
 - [MigrationPreviewResponse](docs/MigrationPreviewResponse.md)
 - [Monitor](docs/Monitor.md)
 - [MonitorGroupInfo](docs/MonitorGroupInfo.md)
 - [MonitorNotification](docs/MonitorNotification.md)
 - [MonitorPlaybook](docs/MonitorPlaybook.md)
 - [MonitorQueries](docs/MonitorQueries.md)
 - [MonitorQuery](docs/MonitorQuery.md)
 - [MonitorScope](docs/MonitorScope.md)
 - [MonitorSubscription](docs/MonitorSubscription.md)
 - [MonitorSubscriptionsListResponse](docs/MonitorSubscriptionsListResponse.md)
 - [MonitorSubscriptionsStatus](docs/MonitorSubscriptionsStatus.md)
 - [MonitorTemplatesLibraryBase](docs/MonitorTemplatesLibraryBase.md)
 - [MonitorTemplatesLibraryBaseExport](docs/MonitorTemplatesLibraryBaseExport.md)
 - [MonitorTemplatesLibraryBaseResponse](docs/MonitorTemplatesLibraryBaseResponse.md)
 - [MonitorTemplatesLibraryBaseUpdate](docs/MonitorTemplatesLibraryBaseUpdate.md)
 - [MonitorTemplatesLibraryFolder](docs/MonitorTemplatesLibraryFolder.md)
 - [MonitorTemplatesLibraryFolderExport](docs/MonitorTemplatesLibraryFolderExport.md)
 - [MonitorTemplatesLibraryFolderResponse](docs/MonitorTemplatesLibraryFolderResponse.md)
 - [MonitorTemplatesLibraryFolderUpdate](docs/MonitorTemplatesLibraryFolderUpdate.md)
 - [MonitorTemplatesLibraryItemWithPath](docs/MonitorTemplatesLibraryItemWithPath.md)
 - [MonitorTemplatesLibraryMonitorTemplate](docs/MonitorTemplatesLibraryMonitorTemplate.md)
 - [MonitorTemplatesLibraryMonitorTemplateExport](docs/MonitorTemplatesLibraryMonitorTemplateExport.md)
 - [MonitorTemplatesLibraryMonitorTemplateResponse](docs/MonitorTemplatesLibraryMonitorTemplateResponse.md)
 - [MonitorTemplatesLibraryMonitorTemplateUpdate](docs/MonitorTemplatesLibraryMonitorTemplateUpdate.md)
 - [MonitorTrigger](docs/MonitorTrigger.md)
 - [MonitorUsage](docs/MonitorUsage.md)
 - [MonitorsLibraryBase](docs/MonitorsLibraryBase.md)
 - [MonitorsLibraryBaseExport](docs/MonitorsLibraryBaseExport.md)
 - [MonitorsLibraryBaseResponse](docs/MonitorsLibraryBaseResponse.md)
 - [MonitorsLibraryBaseUpdate](docs/MonitorsLibraryBaseUpdate.md)
 - [MonitorsLibraryFolder](docs/MonitorsLibraryFolder.md)
 - [MonitorsLibraryFolderExport](docs/MonitorsLibraryFolderExport.md)
 - [MonitorsLibraryFolderResponse](docs/MonitorsLibraryFolderResponse.md)
 - [MonitorsLibraryFolderUpdate](docs/MonitorsLibraryFolderUpdate.md)
 - [MonitorsLibraryItemWithPath](docs/MonitorsLibraryItemWithPath.md)
 - [MonitorsLibraryMonitor](docs/MonitorsLibraryMonitor.md)
 - [MonitorsLibraryMonitorExport](docs/MonitorsLibraryMonitorExport.md)
 - [MonitorsLibraryMonitorResponse](docs/MonitorsLibraryMonitorResponse.md)
 - [MonitorsLibraryMonitorUpdate](docs/MonitorsLibraryMonitorUpdate.md)
 - [MutingInformationResponse](docs/MutingInformationResponse.md)
 - [MutingScheduleResponse](docs/MutingScheduleResponse.md)
 - [MutingSchedulesLibraryBase](docs/MutingSchedulesLibraryBase.md)
 - [MutingSchedulesLibraryBaseExport](docs/MutingSchedulesLibraryBaseExport.md)
 - [MutingSchedulesLibraryBaseResponse](docs/MutingSchedulesLibraryBaseResponse.md)
 - [MutingSchedulesLibraryBaseUpdate](docs/MutingSchedulesLibraryBaseUpdate.md)
 - [MutingSchedulesLibraryFolder](docs/MutingSchedulesLibraryFolder.md)
 - [MutingSchedulesLibraryFolderExport](docs/MutingSchedulesLibraryFolderExport.md)
 - [MutingSchedulesLibraryFolderResponse](docs/MutingSchedulesLibraryFolderResponse.md)
 - [MutingSchedulesLibraryFolderUpdate](docs/MutingSchedulesLibraryFolderUpdate.md)
 - [MutingSchedulesLibraryItemWithPath](docs/MutingSchedulesLibraryItemWithPath.md)
 - [MutingSchedulesLibraryMutingSchedule](docs/MutingSchedulesLibraryMutingSchedule.md)
 - [MutingSchedulesLibraryMutingScheduleExport](docs/MutingSchedulesLibraryMutingScheduleExport.md)
 - [MutingSchedulesLibraryMutingScheduleResponse](docs/MutingSchedulesLibraryMutingScheduleResponse.md)
 - [MutingSchedulesLibraryMutingScheduleUpdate](docs/MutingSchedulesLibraryMutingScheduleUpdate.md)
 - [NewRelic](docs/NewRelic.md)
 - [NextInstancesRequest](docs/NextInstancesRequest.md)
 - [NextInstancesResponse](docs/NextInstancesResponse.md)
 - [NoTraceFieldValuesReason](docs/NoTraceFieldValuesReason.md)
 - [NoneAutoCompleteSyncDefinition](docs/NoneAutoCompleteSyncDefinition.md)
 - [NormalizedIndicator](docs/NormalizedIndicator.md)
 - [NotificationThresholdSyncDefinition](docs/NotificationThresholdSyncDefinition.md)
 - [OAuthRefreshFailedTracker](docs/OAuthRefreshFailedTracker.md)
 - [OTCErrorProcessingSpansTracker](docs/OTCErrorProcessingSpansTracker.md)
 - [OTCExporterErrorTracker](docs/OTCExporterErrorTracker.md)
 - [OTCExporterHighFailuresExportingSpansTracker](docs/OTCExporterHighFailuresExportingSpansTracker.md)
 - [OTCExporterLargeTraceBatchesTracker](docs/OTCExporterLargeTraceBatchesTracker.md)
 - [OTCProcessErrorTracker](docs/OTCProcessErrorTracker.md)
 - [OTCProcessHighMemoryUsageTracker](docs/OTCProcessHighMemoryUsageTracker.md)
 - [OTCProcessSpansDroppedTracker](docs/OTCProcessSpansDroppedTracker.md)
 - [OTCProcessSpansRefusedTracker](docs/OTCProcessSpansRefusedTracker.md)
 - [OTCReceiverErrorTracker](docs/OTCReceiverErrorTracker.md)
 - [OTCReceiverNoSpansObservedTracker](docs/OTCReceiverNoSpansObservedTracker.md)
 - [OTCReceiverSpansDroppedTracker](docs/OTCReceiverSpansDroppedTracker.md)
 - [OTCReceiverSpansRefusedTracker](docs/OTCReceiverSpansRefusedTracker.md)
 - [OTCWarningProcessingSpansTracker](docs/OTCWarningProcessingSpansTracker.md)
 - [OTCollector](docs/OTCollector.md)
 - [OTCollectorCountResponse](docs/OTCollectorCountResponse.md)
 - [OTCollectorHealthIncidentsTracker](docs/OTCollectorHealthIncidentsTracker.md)
 - [OTCollectorLimitApproachingTracker](docs/OTCollectorLimitApproachingTracker.md)
 - [OTCollectorListResponse](docs/OTCollectorListResponse.md)
 - [OTCollectorSystemInfo](docs/OTCollectorSystemInfo.md)
 - [OTCollectorVersion](docs/OTCollectorVersion.md)
 - [OnDemandProvisioningInfo](docs/OnDemandProvisioningInfo.md)
 - [OpenInQuery](docs/OpenInQuery.md)
 - [Operator](docs/Operator.md)
 - [OperatorData](docs/OperatorData.md)
 - [OperatorParameter](docs/OperatorParameter.md)
 - [Opsgenie](docs/Opsgenie.md)
 - [OrTracingExpression](docs/OrTracingExpression.md)
 - [OrderBy](docs/OrderBy.md)
 - [OrgIdentity](docs/OrgIdentity.md)
 - [OtTag](docs/OtTag.md)
 - [OutlierBound](docs/OutlierBound.md)
 - [OutlierDataValue](docs/OutlierDataValue.md)
 - [OutlierSeriesDataPoint](docs/OutlierSeriesDataPoint.md)
 - [PagerDuty](docs/PagerDuty.md)
 - [PaginatedDashboards](docs/PaginatedDashboards.md)
 - [PaginatedListAccessKeysResult](docs/PaginatedListAccessKeysResult.md)
 - [PaginatedListEndpoints](docs/PaginatedListEndpoints.md)
 - [PaginatedLogSearches](docs/PaginatedLogSearches.md)
 - [PaginatedOTCollectorsRequest](docs/PaginatedOTCollectorsRequest.md)
 - [PaginatedOTCollectorsRequestFilters](docs/PaginatedOTCollectorsRequestFilters.md)
 - [PaginatedOTCollectorsResponse](docs/PaginatedOTCollectorsResponse.md)
 - [Panel](docs/Panel.md)
 - [PanelItem](docs/PanelItem.md)
 - [ParameterAutoCompleteSyncDefinition](docs/ParameterAutoCompleteSyncDefinition.md)
 - [Partition](docs/Partition.md)
 - [PartitionInfo](docs/PartitionInfo.md)
 - [PartitionsResponse](docs/PartitionsResponse.md)
 - [PasswordPolicy](docs/PasswordPolicy.md)
 - [Path](docs/Path.md)
 - [PathItem](docs/PathItem.md)
 - [PendingUpdateRequest](docs/PendingUpdateRequest.md)
 - [PermissionIdentifier](docs/PermissionIdentifier.md)
 - [PermissionIdentifiers](docs/PermissionIdentifiers.md)
 - [PermissionStatement](docs/PermissionStatement.md)
 - [PermissionStatementDefinition](docs/PermissionStatementDefinition.md)
 - [PermissionStatementDefinitions](docs/PermissionStatementDefinitions.md)
 - [PermissionStatements](docs/PermissionStatements.md)
 - [PermissionSubject](docs/PermissionSubject.md)
 - [PermissionSummariesBySubjects](docs/PermissionSummariesBySubjects.md)
 - [PermissionSummaryBySubjects](docs/PermissionSummaryBySubjects.md)
 - [PermissionSummaryMeta](docs/PermissionSummaryMeta.md)
 - [Permissions](docs/Permissions.md)
 - [Plan](docs/Plan.md)
 - [PlanUpdateEmail](docs/PlanUpdateEmail.md)
 - [PlansCatalog](docs/PlansCatalog.md)
 - [PlaybookExecutionParameters](docs/PlaybookExecutionParameters.md)
 - [PlaybookExecutionResponse](docs/PlaybookExecutionResponse.md)
 - [PlaybookRunningListRequest](docs/PlaybookRunningListRequest.md)
 - [PlaybookRunningResult](docs/PlaybookRunningResult.md)
 - [Points](docs/Points.md)
 - [PreviewLookupTableField](docs/PreviewLookupTableField.md)
 - [ProductGroup](docs/ProductGroup.md)
 - [ProductSubscriptionOption](docs/ProductSubscriptionOption.md)
 - [ProductVariable](docs/ProductVariable.md)
 - [ProrationDetails](docs/ProrationDetails.md)
 - [Quantity](docs/Quantity.md)
 - [QueriesParametersResult](docs/QueriesParametersResult.md)
 - [Query](docs/Query.md)
 - [QueryBasedSli](docs/QueryBasedSli.md)
 - [QueryParameterSyncDefinition](docs/QueryParameterSyncDefinition.md)
 - [RangeTracingValue](docs/RangeTracingValue.md)
 - [RelatedAlert](docs/RelatedAlert.md)
 - [RelatedAlertsLibraryAlertResponse](docs/RelatedAlertsLibraryAlertResponse.md)
 - [RelativeTimeRangeBoundary](docs/RelativeTimeRangeBoundary.md)
 - [RemoveIndicatorsRequest](docs/RemoveIndicatorsRequest.md)
 - [ReportAction](docs/ReportAction.md)
 - [ReportAutoParsingInfo](docs/ReportAutoParsingInfo.md)
 - [ReportFilterSyncDefinition](docs/ReportFilterSyncDefinition.md)
 - [ReportPanelSyncDefinition](docs/ReportPanelSyncDefinition.md)
 - [Request](docs/Request.md)
 - [ResolvableTimeRange](docs/ResolvableTimeRange.md)
 - [ResourceIdentities](docs/ResourceIdentities.md)
 - [ResourceIdentity](docs/ResourceIdentity.md)
 - [RoleDefinition](docs/RoleDefinition.md)
 - [RoleModel](docs/RoleModel.md)
 - [RoleModelV2](docs/RoleModelV2.md)
 - [RollingCompliance](docs/RollingCompliance.md)
 - [RootSpanTracingFilter](docs/RootSpanTracingFilter.md)
 - [RowDeleteDefinition](docs/RowDeleteDefinition.md)
 - [RowUpdateDefinition](docs/RowUpdateDefinition.md)
 - [RuleAndBucketDetail](docs/RuleAndBucketDetail.md)
 - [RunAs](docs/RunAs.md)
 - [S3CollectionErrorTracker](docs/S3CollectionErrorTracker.md)
 - [SamlIdentityProvider](docs/SamlIdentityProvider.md)
 - [SamlIdentityProviderRequest](docs/SamlIdentityProviderRequest.md)
 - [SaveLogSearchRequest](docs/SaveLogSearchRequest.md)
 - [SaveMetricsSearchRequest](docs/SaveMetricsSearchRequest.md)
 - [SaveToLookupNotificationSyncDefinition](docs/SaveToLookupNotificationSyncDefinition.md)
 - [SaveToViewNotificationSyncDefinition](docs/SaveToViewNotificationSyncDefinition.md)
 - [SavedSearchSyncDefinition](docs/SavedSearchSyncDefinition.md)
 - [SavedSearchSyncDefinitionBase](docs/SavedSearchSyncDefinitionBase.md)
 - [SavedSearchWithScheduleSyncDefinition](docs/SavedSearchWithScheduleSyncDefinition.md)
 - [ScannedBytes](docs/ScannedBytes.md)
 - [ScheduleDefinition](docs/ScheduleDefinition.md)
 - [ScheduleNotificationSyncDefinition](docs/ScheduleNotificationSyncDefinition.md)
 - [ScheduleSearchParameterSyncDefinition](docs/ScheduleSearchParameterSyncDefinition.md)
 - [ScheduledView](docs/ScheduledView.md)
 - [SearchAuditPolicy](docs/SearchAuditPolicy.md)
 - [SearchQueryContext](docs/SearchQueryContext.md)
 - [SearchQueryFieldAndType](docs/SearchQueryFieldAndType.md)
 - [SearchScheduleSyncDefinition](docs/SearchScheduleSyncDefinition.md)
 - [SelfServiceCreditsBaselines](docs/SelfServiceCreditsBaselines.md)
 - [SelfServicePlan](docs/SelfServicePlan.md)
 - [SeriesAxisRange](docs/SeriesAxisRange.md)
 - [SeriesData](docs/SeriesData.md)
 - [SeriesMetadata](docs/SeriesMetadata.md)
 - [ServiceManifestDataSourceParameter](docs/ServiceManifestDataSourceParameter.md)
 - [ServiceMapEdge](docs/ServiceMapEdge.md)
 - [ServiceMapNode](docs/ServiceMapNode.md)
 - [ServiceMapPanel](docs/ServiceMapPanel.md)
 - [ServiceMapResponse](docs/ServiceMapResponse.md)
 - [ServiceNow](docs/ServiceNow.md)
 - [ServiceNowConnection](docs/ServiceNowConnection.md)
 - [ServiceNowDefinition](docs/ServiceNowDefinition.md)
 - [ServiceNowFieldsSyncDefinition](docs/ServiceNowFieldsSyncDefinition.md)
 - [ServiceNowSearchNotificationSyncDefinition](docs/ServiceNowSearchNotificationSyncDefinition.md)
 - [ShareDashboardsOutsideOrganizationPolicy](docs/ShareDashboardsOutsideOrganizationPolicy.md)
 - [SharedBucket](docs/SharedBucket.md)
 - [SignalContext](docs/SignalContext.md)
 - [SignalsJobResult](docs/SignalsJobResult.md)
 - [SignalsRequest](docs/SignalsRequest.md)
 - [SignalsResponse](docs/SignalsResponse.md)
 - [Slack](docs/Slack.md)
 - [Sli](docs/Sli.md)
 - [SliQueries](docs/SliQueries.md)
 - [SliQueriesValidationResult](docs/SliQueriesValidationResult.md)
 - [SliQuery](docs/SliQuery.md)
 - [SliQueryGroup](docs/SliQueryGroup.md)
 - [SliStatus](docs/SliStatus.md)
 - [SloBurnRateCondition](docs/SloBurnRateCondition.md)
 - [SloSliCondition](docs/SloSliCondition.md)
 - [SloUsage](docs/SloUsage.md)
 - [SlosLibraryBase](docs/SlosLibraryBase.md)
 - [SlosLibraryBaseExport](docs/SlosLibraryBaseExport.md)
 - [SlosLibraryBaseResponse](docs/SlosLibraryBaseResponse.md)
 - [SlosLibraryBaseUpdate](docs/SlosLibraryBaseUpdate.md)
 - [SlosLibraryFolder](docs/SlosLibraryFolder.md)
 - [SlosLibraryFolderExport](docs/SlosLibraryFolderExport.md)
 - [SlosLibraryFolderResponse](docs/SlosLibraryFolderResponse.md)
 - [SlosLibraryFolderUpdate](docs/SlosLibraryFolderUpdate.md)
 - [SlosLibraryItemWithPath](docs/SlosLibraryItemWithPath.md)
 - [SlosLibrarySlo](docs/SlosLibrarySlo.md)
 - [SlosLibrarySloExport](docs/SlosLibrarySloExport.md)
 - [SlosLibrarySloResponse](docs/SlosLibrarySloResponse.md)
 - [SlosLibrarySloUpdate](docs/SlosLibrarySloUpdate.md)
 - [Source](docs/Source.md)
 - [SourceResourceIdentity](docs/SourceResourceIdentity.md)
 - [SpanCalculationAggregator](docs/SpanCalculationAggregator.md)
 - [SpanCalculationAvgAggregator](docs/SpanCalculationAvgAggregator.md)
 - [SpanCalculationMaxAggregator](docs/SpanCalculationMaxAggregator.md)
 - [SpanCalculationMinAggregator](docs/SpanCalculationMinAggregator.md)
 - [SpanCalculationPctAggregator](docs/SpanCalculationPctAggregator.md)
 - [SpanCalculationSumAggregator](docs/SpanCalculationSumAggregator.md)
 - [SpanEvent](docs/SpanEvent.md)
 - [SpanEventAttribute](docs/SpanEventAttribute.md)
 - [SpanIngestLimitExceededTracker](docs/SpanIngestLimitExceededTracker.md)
 - [SpanLink](docs/SpanLink.md)
 - [SpanPathSegment](docs/SpanPathSegment.md)
 - [SpanQueryAggregateAggregateData](docs/SpanQueryAggregateAggregateData.md)
 - [SpanQueryAggregateDataSeries](docs/SpanQueryAggregateDataSeries.md)
 - [SpanQueryAggregateMetaData](docs/SpanQueryAggregateMetaData.md)
 - [SpanQueryAggregatePointData](docs/SpanQueryAggregatePointData.md)
 - [SpanQueryAggregateResponse](docs/SpanQueryAggregateResponse.md)
 - [SpanQueryAggregateResult](docs/SpanQueryAggregateResult.md)
 - [SpanQueryFieldDetail](docs/SpanQueryFieldDetail.md)
 - [SpanQueryFieldsResponse](docs/SpanQueryFieldsResponse.md)
 - [SpanQueryRequest](docs/SpanQueryRequest.md)
 - [SpanQueryResponse](docs/SpanQueryResponse.md)
 - [SpanQueryResultFacetsResponse](docs/SpanQueryResultFacetsResponse.md)
 - [SpanQueryResultSpansResponse](docs/SpanQueryResultSpansResponse.md)
 - [SpanQueryRow](docs/SpanQueryRow.md)
 - [SpanQueryRowError](docs/SpanQueryRowError.md)
 - [SpanQueryRowFacet](docs/SpanQueryRowFacet.md)
 - [SpanQueryRowResponse](docs/SpanQueryRowResponse.md)
 - [SpanQueryRowStatus](docs/SpanQueryRowStatus.md)
 - [SpanQuerySpanData](docs/SpanQuerySpanData.md)
 - [SpanQueryStatusResponse](docs/SpanQueryStatusResponse.md)
 - [SpansCalculationVisualization](docs/SpansCalculationVisualization.md)
 - [SpansCountVisualization](docs/SpansCountVisualization.md)
 - [SpansFieldGroupBy](docs/SpansFieldGroupBy.md)
 - [SpansFilter](docs/SpansFilter.md)
 - [SpansFilterKeyValuePair](docs/SpansFilterKeyValuePair.md)
 - [SpansFilterStandaloneKey](docs/SpansFilterStandaloneKey.md)
 - [SpansGroupBy](docs/SpansGroupBy.md)
 - [SpansLimitItem](docs/SpansLimitItem.md)
 - [SpansQueryData](docs/SpansQueryData.md)
 - [SpansTimeGroupBy](docs/SpansTimeGroupBy.md)
 - [SpansVisualization](docs/SpansVisualization.md)
 - [StaticCondition](docs/StaticCondition.md)
 - [StaticSeriesDataPoint](docs/StaticSeriesDataPoint.md)
 - [StixIndicator](docs/StixIndicator.md)
 - [StringArrayEventAttributeValue](docs/StringArrayEventAttributeValue.md)
 - [StringEventAttributeValue](docs/StringEventAttributeValue.md)
 - [StringTracingValue](docs/StringTracingValue.md)
 - [SubdomainAvailabilityResponse](docs/SubdomainAvailabilityResponse.md)
 - [SubdomainDefinitionResponse](docs/SubdomainDefinitionResponse.md)
 - [SubdomainUrlResponse](docs/SubdomainUrlResponse.md)
 - [SumoCloudSOAR](docs/SumoCloudSOAR.md)
 - [SumoOrgsUsageBackfillRequest](docs/SumoOrgsUsageBackfillRequest.md)
 - [SumoSearchPanel](docs/SumoSearchPanel.md)
 - [TableRow](docs/TableRow.md)
 - [TagReversedIndex](docs/TagReversedIndex.md)
 - [TagValueReversedIndex](docs/TagValueReversedIndex.md)
 - [TagsReversedIndexResponse](docs/TagsReversedIndexResponse.md)
 - [Template](docs/Template.md)
 - [TestConnectionResponse](docs/TestConnectionResponse.md)
 - [TextEntriesAutoCompleteSyncDefinition](docs/TextEntriesAutoCompleteSyncDefinition.md)
 - [TextPanel](docs/TextPanel.md)
 - [TimeRangeBoundary](docs/TimeRangeBoundary.md)
 - [TimeSeries](docs/TimeSeries.md)
 - [TimeSeriesList](docs/TimeSeriesList.md)
 - [TimeSeriesRow](docs/TimeSeriesRow.md)
 - [TokenBaseDefinition](docs/TokenBaseDefinition.md)
 - [TokenBaseDefinitionUpdate](docs/TokenBaseDefinitionUpdate.md)
 - [TokenBaseResponse](docs/TokenBaseResponse.md)
 - [TopologyLabelMap](docs/TopologyLabelMap.md)
 - [TopologySearchLabel](docs/TopologySearchLabel.md)
 - [TotalCredits](docs/TotalCredits.md)
 - [TraceDbSpanInfo](docs/TraceDbSpanInfo.md)
 - [TraceDetail](docs/TraceDetail.md)
 - [TraceExistsResponse](docs/TraceExistsResponse.md)
 - [TraceFieldDetail](docs/TraceFieldDetail.md)
 - [TraceFieldValuesResponse](docs/TraceFieldValuesResponse.md)
 - [TraceFieldsResponse](docs/TraceFieldsResponse.md)
 - [TraceHttpSpanInfo](docs/TraceHttpSpanInfo.md)
 - [TraceLightEventsResponse](docs/TraceLightEventsResponse.md)
 - [TraceMessageBusSpanInfo](docs/TraceMessageBusSpanInfo.md)
 - [TraceMetricDetail](docs/TraceMetricDetail.md)
 - [TraceMetricsResponse](docs/TraceMetricsResponse.md)
 - [TraceQueryExpression](docs/TraceQueryExpression.md)
 - [TraceQueryResultResponse](docs/TraceQueryResultResponse.md)
 - [TraceQueryRowStatus](docs/TraceQueryRowStatus.md)
 - [TraceQueryStatusResponse](docs/TraceQueryStatusResponse.md)
 - [TraceSpan](docs/TraceSpan.md)
 - [TraceSpanBillingInfo](docs/TraceSpanBillingInfo.md)
 - [TraceSpanCriticalPathContribution](docs/TraceSpanCriticalPathContribution.md)
 - [TraceSpanDetail](docs/TraceSpanDetail.md)
 - [TraceSpanInfo](docs/TraceSpanInfo.md)
 - [TraceSpanStatus](docs/TraceSpanStatus.md)
 - [TraceSpansResponse](docs/TraceSpansResponse.md)
 - [TracesFilter](docs/TracesFilter.md)
 - [TracesListPanel](docs/TracesListPanel.md)
 - [TracesQueryData](docs/TracesQueryData.md)
 - [TracingValue](docs/TracingValue.md)
 - [TrackerIdentity](docs/TrackerIdentity.md)
 - [TransformationRuleDefinition](docs/TransformationRuleDefinition.md)
 - [TransformationRuleRequest](docs/TransformationRuleRequest.md)
 - [TransformationRuleResponse](docs/TransformationRuleResponse.md)
 - [TransformationRulesResponse](docs/TransformationRulesResponse.md)
 - [TriggerCondition](docs/TriggerCondition.md)
 - [UnvalidatedMonitorQuery](docs/UnvalidatedMonitorQuery.md)
 - [UpdateBucketDefinition](docs/UpdateBucketDefinition.md)
 - [UpdateDataForwardingRule](docs/UpdateDataForwardingRule.md)
 - [UpdateExtractionRuleDefinition](docs/UpdateExtractionRuleDefinition.md)
 - [UpdateFolderRequest](docs/UpdateFolderRequest.md)
 - [UpdatePartitionDefinition](docs/UpdatePartitionDefinition.md)
 - [UpdateRequest](docs/UpdateRequest.md)
 - [UpdateRoleDefinition](docs/UpdateRoleDefinition.md)
 - [UpdateRoleDefinitionV2](docs/UpdateRoleDefinitionV2.md)
 - [UpdateScheduledViewDefinition](docs/UpdateScheduledViewDefinition.md)
 - [UpdateUserDefinition](docs/UpdateUserDefinition.md)
 - [UpgradePlans](docs/UpgradePlans.md)
 - [UploadBlobIndicatorsRequest](docs/UploadBlobIndicatorsRequest.md)
 - [UploadCsvIndicatorsRequest](docs/UploadCsvIndicatorsRequest.md)
 - [UploadNormalizedIndicatorRequest](docs/UploadNormalizedIndicatorRequest.md)
 - [UploadStixIndicatorsRequest](docs/UploadStixIndicatorsRequest.md)
 - [UsageDetails](docs/UsageDetails.md)
 - [UsageReportRequest](docs/UsageReportRequest.md)
 - [UsageReportResponse](docs/UsageReportResponse.md)
 - [UsageReportStatusResponse](docs/UsageReportStatusResponse.md)
 - [UserConcurrentSessionsLimitPolicy](docs/UserConcurrentSessionsLimitPolicy.md)
 - [UserInfo](docs/UserInfo.md)
 - [UserInterests](docs/UserInterests.md)
 - [UserModel](docs/UserModel.md)
 - [ValueOnlyLookupAutoCompleteSyncDefinition](docs/ValueOnlyLookupAutoCompleteSyncDefinition.md)
 - [Variable](docs/Variable.md)
 - [VariableSourceDefinition](docs/VariableSourceDefinition.md)
 - [VariableValuesData](docs/VariableValuesData.md)
 - [VariableValuesLogQueryRequest](docs/VariableValuesLogQueryRequest.md)
 - [VariablesValuesData](docs/VariablesValuesData.md)
 - [ViewFilterDefinition](docs/ViewFilterDefinition.md)
 - [ViewRetentionProperties](docs/ViewRetentionProperties.md)
 - [VisualAggregateData](docs/VisualAggregateData.md)
 - [VisualAxisData](docs/VisualAxisData.md)
 - [VisualDataAxes](docs/VisualDataAxes.md)
 - [VisualDataSeries](docs/VisualDataSeries.md)
 - [VisualMetaData](docs/VisualMetaData.md)
 - [VisualOutlierData](docs/VisualOutlierData.md)
 - [VisualPointData](docs/VisualPointData.md)
 - [WarningDescription](docs/WarningDescription.md)
 - [WarningDetails](docs/WarningDetails.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookConnection](docs/WebhookConnection.md)
 - [WebhookDefinition](docs/WebhookDefinition.md)
 - [WebhookSearchNotificationSyncDefinition](docs/WebhookSearchNotificationSyncDefinition.md)
 - [Window](docs/Window.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### basicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), openapi.ContextBasicAuth, openapi.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



