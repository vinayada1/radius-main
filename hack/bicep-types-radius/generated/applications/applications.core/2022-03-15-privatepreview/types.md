# Applications.Core @ 2022-03-15-privatepreview

## Resource Applications.Core/applications@2022-03-15-privatepreview
* **Valid Scope(s)**: ResourceGroup
### Properties
* **apiVersion**: '2022-03-15-privatepreview' (ReadOnly, DeployTimeConstant): The resource api version
* **id**: string (ReadOnly, DeployTimeConstant): The resource id
* **location**: string (Required): The geo-location where the resource lives
* **name**: string (Required, DeployTimeConstant): The resource name
* **properties**: [ApplicationProperties](#applicationproperties) (Required): Application properties
* **systemData**: [SystemData](#systemdata) (ReadOnly): Metadata pertaining to creation and last modification of the resource.
* **tags**: [TrackedResourceTags](#trackedresourcetags): Resource tags.
* **type**: 'Applications.Core/applications' (ReadOnly, DeployTimeConstant): The resource type

## Resource Applications.Core/environments@2022-03-15-privatepreview
* **Valid Scope(s)**: ResourceGroup
### Properties
* **apiVersion**: '2022-03-15-privatepreview' (ReadOnly, DeployTimeConstant): The resource api version
* **id**: string (ReadOnly, DeployTimeConstant): The resource id
* **location**: string (Required): The geo-location where the resource lives
* **name**: string (Required, DeployTimeConstant): The resource name
* **properties**: [EnvironmentProperties](#environmentproperties) (Required): Application environment properties
* **systemData**: [SystemData](#systemdata) (ReadOnly): Metadata pertaining to creation and last modification of the resource.
* **tags**: [TrackedResourceTags](#trackedresourcetags): Resource tags.
* **type**: 'Applications.Core/environments' (ReadOnly, DeployTimeConstant): The resource type

## Resource Applications.Core/httpRoutes@2022-03-15-privatepreview
* **Valid Scope(s)**: ResourceGroup
### Properties
* **apiVersion**: '2022-03-15-privatepreview' (ReadOnly, DeployTimeConstant): The resource api version
* **id**: string (ReadOnly, DeployTimeConstant): The resource id
* **location**: string (Required): The geo-location where the resource lives
* **name**: string (Required, DeployTimeConstant): The resource name
* **properties**: [HttpRouteProperties](#httprouteproperties) (Required): HTTP Route properties
* **systemData**: [SystemData](#systemdata) (ReadOnly): Metadata pertaining to creation and last modification of the resource.
* **tags**: [TrackedResourceTags](#trackedresourcetags): Resource tags.
* **type**: 'Applications.Core/httpRoutes' (ReadOnly, DeployTimeConstant): The resource type

## ApplicationProperties
### Properties
* **environment**: string (Required): The resource id of the environment linked to application.
* **provisioningState**: 'Accepted' | 'Canceled' | 'Deleting' | 'Failed' | 'Provisioning' | 'Succeeded' | 'Updating' (ReadOnly): Provisioning state of the resource at the time the operation was called.

## SystemData
### Properties
* **createdAt**: string: The timestamp of resource creation (UTC).
* **createdBy**: string: The identity that created the resource.
* **createdByType**: 'Application' | 'Key' | 'ManagedIdentity' | 'User': The type of identity that created the resource.
* **lastModifiedAt**: string: The timestamp of resource last modification (UTC)
* **lastModifiedBy**: string: The identity that last modified the resource.
* **lastModifiedByType**: 'Application' | 'Key' | 'ManagedIdentity' | 'User': The type of identity that created the resource.

## TrackedResourceTags
### Properties
### Additional Properties
* **Additional Properties Type**: string

## EnvironmentProperties
### Properties
* **compute**: [EnvironmentCompute](#environmentcompute) (Required): Compute resource used by application environment resource.
* **provisioningState**: 'Accepted' | 'Canceled' | 'Deleting' | 'Failed' | 'Provisioning' | 'Succeeded' | 'Updating' (ReadOnly): Provisioning state of the resource at the time the operation was called.

## EnvironmentCompute
### Properties
* **kind**: 'kubernetes' (Required): Type of compute resource.
* **resourceId**: string (Required): The resource id of the compute resource for application environment.

## TrackedResourceTags
### Properties
### Additional Properties
* **Additional Properties Type**: string

## HttpRouteProperties
### Properties
* **application**: string (Required): The resource id of the application linked to HTTP Route resource.
* **hostname**: string: The internal hostname accepting traffic for the HTTP Route. Readonly.
* **port**: int: The port number for the HTTP Route. Defaults to 80. Readonly.
* **provisioningState**: 'Accepted' | 'Canceled' | 'Deleting' | 'Failed' | 'Provisioning' | 'Succeeded' | 'Updating' (ReadOnly): Provisioning state of the resource at the time the operation was called.
* **scheme**: string: The scheme used for traffic. Readonly.
* **status**: [ResourceStatus](#resourcestatus): Status of a resource.
* **url**: string: A stable URL that that can be used to route traffic to a resource. Readonly.

## ResourceStatus
### Properties
* **outputResources**: any[]: Array of AnyObject

## TrackedResourceTags
### Properties
### Additional Properties
* **Additional Properties Type**: string
