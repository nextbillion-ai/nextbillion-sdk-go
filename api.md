# Fleetify

## Routes

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Routing">Routing</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteNewResponse">FleetifyRouteNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteRedispatchResponse">FleetifyRouteRedispatchResponse</a>

Methods:

- <code title="post /fleetify/routes">client.Fleetify.Routes.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteNewParams">FleetifyRouteNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteNewResponse">FleetifyRouteNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /fleetify/routes/{routeID}/redispatch">client.Fleetify.Routes.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteService.Redispatch">Redispatch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, routeID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteRedispatchParams">FleetifyRouteRedispatchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteRedispatchResponse">FleetifyRouteRedispatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Steps

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DocumentSubmission">DocumentSubmission</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RouteStepCompletionMode">RouteStepCompletionMode</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RouteStepGeofenceConfigParam">RouteStepGeofenceConfigParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RouteStepsRequestParam">RouteStepsRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DocumentSubmission">DocumentSubmission</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RouteStepCompletionMode">RouteStepCompletionMode</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RouteStepGeofenceConfig">RouteStepGeofenceConfig</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RouteStepsResponse">RouteStepsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepNewResponse">FleetifyRouteStepNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepUpdateResponse">FleetifyRouteStepUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepDeleteResponse">FleetifyRouteStepDeleteResponse</a>

Methods:

- <code title="post /fleetify/routes/{routeID}/steps">client.Fleetify.Routes.Steps.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, routeID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepNewParams">FleetifyRouteStepNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepNewResponse">FleetifyRouteStepNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /fleetify/routes/{routeID}/steps/{stepID}">client.Fleetify.Routes.Steps.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepUpdateParams">FleetifyRouteStepUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepUpdateResponse">FleetifyRouteStepUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /fleetify/routes/{routeID}/steps/{stepID}">client.Fleetify.Routes.Steps.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepDeleteParams">FleetifyRouteStepDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepDeleteResponse">FleetifyRouteStepDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /fleetify/routes/{routeID}/steps/{stepID}">client.Fleetify.Routes.Steps.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, stepID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyRouteStepCompleteParams">FleetifyRouteStepCompleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## DocumentTemplates

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DocumentTemplateContentRequestParam">DocumentTemplateContentRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DocumentTemplateContentResponse">DocumentTemplateContentResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateNewResponse">FleetifyDocumentTemplateNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateGetResponse">FleetifyDocumentTemplateGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateUpdateResponse">FleetifyDocumentTemplateUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateListResponse">FleetifyDocumentTemplateListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateDeleteResponse">FleetifyDocumentTemplateDeleteResponse</a>

Methods:

- <code title="post /fleetify/document_templates">client.Fleetify.DocumentTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateNewParams">FleetifyDocumentTemplateNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateNewResponse">FleetifyDocumentTemplateNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fleetify/document_templates/{id}">client.Fleetify.DocumentTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateGetParams">FleetifyDocumentTemplateGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateGetResponse">FleetifyDocumentTemplateGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /fleetify/document_templates/{id}">client.Fleetify.DocumentTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateUpdateParams">FleetifyDocumentTemplateUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateUpdateResponse">FleetifyDocumentTemplateUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /fleetify/document_templates">client.Fleetify.DocumentTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateListParams">FleetifyDocumentTemplateListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateListResponse">FleetifyDocumentTemplateListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /fleetify/document_templates/{id}">client.Fleetify.DocumentTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateDeleteParams">FleetifyDocumentTemplateDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#FleetifyDocumentTemplateDeleteResponse">FleetifyDocumentTemplateDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Skynet

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSubscribeResponse">SkynetSubscribeResponse</a>

Methods:

- <code title="post /skynet/subscribe">client.Skynet.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetService.Subscribe">Subscribe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSubscribeParams">SkynetSubscribeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSubscribeResponse">SkynetSubscribeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Asset

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MetaData">MetaData</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MetaData">MetaData</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetNewResponse">SkynetAssetNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetGetResponse">SkynetAssetGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetGetListResponse">SkynetAssetGetListResponse</a>

Methods:

- <code title="post /skynet/asset">client.Skynet.Asset.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetNewParams">SkynetAssetNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetNewResponse">SkynetAssetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /skynet/asset/{id}">client.Skynet.Asset.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetGetParams">SkynetAssetGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetGetResponse">SkynetAssetGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /skynet/asset/{id}">client.Skynet.Asset.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetUpdateParams">SkynetAssetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /skynet/asset/{id}">client.Skynet.Asset.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetDeleteParams">SkynetAssetDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /skynet/asset/list">client.Skynet.Asset.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetService.GetList">GetList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetGetListParams">SkynetAssetGetListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetGetListResponse">SkynetAssetGetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /skynet/asset/{id}/track">client.Skynet.Asset.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetService.Track">Track</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetTrackParams">SkynetAssetTrackParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /skynet/asset/{id}/attributes">client.Skynet.Asset.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetService.UpdateAttributes">UpdateAttributes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetUpdateAttributesParams">SkynetAssetUpdateAttributesParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Event

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetEventGetListResponse">SkynetAssetEventGetListResponse</a>

Methods:

- <code title="get /skynet/asset/{id}/event/list">client.Skynet.Asset.Event.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetEventService.GetList">GetList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetEventGetListParams">SkynetAssetEventGetListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetEventGetListResponse">SkynetAssetEventGetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Location

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#TrackLocation">TrackLocation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetLocationGetLastResponse">SkynetAssetLocationGetLastResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetLocationGetListResponse">SkynetAssetLocationGetListResponse</a>

Methods:

- <code title="get /skynet/asset/{id}/location/last">client.Skynet.Asset.Location.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetLocationService.GetLast">GetLast</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetLocationGetLastParams">SkynetAssetLocationGetLastParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetLocationGetLastResponse">SkynetAssetLocationGetLastResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /skynet/asset/{id}/location/list">client.Skynet.Asset.Location.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetLocationService.GetList">GetList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetLocationGetListParams">SkynetAssetLocationGetListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetAssetLocationGetListResponse">SkynetAssetLocationGetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Monitor

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Metadata">Metadata</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Metadata">Metadata</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Monitor">Monitor</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Pagination">Pagination</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorNewResponse">SkynetMonitorNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorGetResponse">SkynetMonitorGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorGetListResponse">SkynetMonitorGetListResponse</a>

Methods:

- <code title="post /skynet/monitor">client.Skynet.Monitor.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorNewParams">SkynetMonitorNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorNewResponse">SkynetMonitorNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /skynet/monitor/{id}">client.Skynet.Monitor.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorGetParams">SkynetMonitorGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorGetResponse">SkynetMonitorGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /skynet/monitor/{id}">client.Skynet.Monitor.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorUpdateParams">SkynetMonitorUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /skynet/monitor/{id}">client.Skynet.Monitor.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorDeleteParams">SkynetMonitorDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /skynet/monitor/list">client.Skynet.Monitor.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorService.GetList">GetList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorGetListParams">SkynetMonitorGetListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetMonitorGetListResponse">SkynetMonitorGetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Trip

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Asset">Asset</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#TripStop">TripStop</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripGetResponse">SkynetTripGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripGetSummaryResponse">SkynetTripGetSummaryResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripStartResponse">SkynetTripStartResponse</a>

Methods:

- <code title="get /skynet/trip/{id}">client.Skynet.Trip.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripGetParams">SkynetTripGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripGetResponse">SkynetTripGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /skynet/trip/{id}">client.Skynet.Trip.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripUpdateParams">SkynetTripUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /skynet/trip/{id}">client.Skynet.Trip.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripDeleteParams">SkynetTripDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /skynet/trip/end">client.Skynet.Trip.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripService.End">End</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripEndParams">SkynetTripEndParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /skynet/trip/{id}/summary">client.Skynet.Trip.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripService.GetSummary">GetSummary</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripGetSummaryParams">SkynetTripGetSummaryParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripGetSummaryResponse">SkynetTripGetSummaryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /skynet/trip/start">client.Skynet.Trip.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripStartParams">SkynetTripStartParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetTripStartResponse">SkynetTripStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## NamespacedApikeys

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetNamespacedApikeyDeleteNamespacedApikeysResponse">SkynetNamespacedApikeyDeleteNamespacedApikeysResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetNamespacedApikeyNamespacedApikeysResponse">SkynetNamespacedApikeyNamespacedApikeysResponse</a>

Methods:

- <code title="delete /skynet/namespaced-apikeys">client.Skynet.NamespacedApikeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetNamespacedApikeyService.DeleteNamespacedApikeys">DeleteNamespacedApikeys</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetNamespacedApikeyDeleteNamespacedApikeysParams">SkynetNamespacedApikeyDeleteNamespacedApikeysParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetNamespacedApikeyDeleteNamespacedApikeysResponse">SkynetNamespacedApikeyDeleteNamespacedApikeysResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /skynet/namespaced-apikeys">client.Skynet.NamespacedApikeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetNamespacedApikeyService.NamespacedApikeys">NamespacedApikeys</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetNamespacedApikeyNamespacedApikeysParams">SkynetNamespacedApikeyNamespacedApikeysParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetNamespacedApikeyNamespacedApikeysResponse">SkynetNamespacedApikeyNamespacedApikeysResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Config

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetConfigListResponse">SkynetConfigListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetConfigTestwebhookResponse">SkynetConfigTestwebhookResponse</a>

Methods:

- <code title="put /skynet/config">client.Skynet.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetConfigService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetConfigNewParams">SkynetConfigNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /skynet/config">client.Skynet.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetConfigService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetConfigListParams">SkynetConfigListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetConfigListResponse">SkynetConfigListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /skynet/config/testwebhook">client.Skynet.Config.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetConfigService.Testwebhook">Testwebhook</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetConfigTestwebhookParams">SkynetConfigTestwebhookParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetConfigTestwebhookResponse">SkynetConfigTestwebhookResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Search

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SearchResponse">SearchResponse</a>

Methods:

- <code title="get /skynet/search/around">client.Skynet.Search.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSearchService.GetAround">GetAround</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSearchGetAroundParams">SkynetSearchGetAroundParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /skynet/search/bound">client.Skynet.Search.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSearchService.GetBound">GetBound</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSearchGetBoundParams">SkynetSearchGetBoundParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Polygon

Methods:

- <code title="post /skynet/search/polygon">client.Skynet.Search.Polygon.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSearchPolygonService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSearchPolygonNewParams">SkynetSearchPolygonNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /skynet/search/polygon">client.Skynet.Search.Polygon.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSearchPolygonService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSearchPolygonListParams">SkynetSearchPolygonListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SearchResponse">SearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Skynet

### Asset

Methods:

- <code title="post /skynet/skynet/asset/{id}/bind">client.Skynet.Skynet.Asset.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSkynetAssetService.Bind">Bind</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SkynetSkynetAssetBindParams">SkynetSkynetAssetBindParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Geocode

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Access">Access</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Categories">Categories</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#ContactObject">ContactObject</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Contacts">Contacts</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MapView">MapView</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Position">Position</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeGetResponse">GeocodeGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeBatchNewResponse">GeocodeBatchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeStructuredGetResponse">GeocodeStructuredGetResponse</a>

Methods:

- <code title="get /geocode">client.Geocode.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeGetParams">GeocodeGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeGetResponse">GeocodeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /geocode/batch">client.Geocode.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeService.BatchNew">BatchNew</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeBatchNewParams">GeocodeBatchNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeBatchNewResponse">GeocodeBatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /geocode/structured">client.Geocode.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeService.StructuredGet">StructuredGet</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeStructuredGetParams">GeocodeStructuredGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeocodeStructuredGetResponse">GeocodeStructuredGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Optimization

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#PostResponse">PostResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationComputeResponse">OptimizationComputeResponse</a>

Methods:

- <code title="get /optimization/json">client.Optimization.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationService.Compute">Compute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationComputeParams">OptimizationComputeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationComputeResponse">OptimizationComputeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /optimization/re_optimization">client.Optimization.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationService.ReOptimize">ReOptimize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationReOptimizeParams">OptimizationReOptimizeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#PostResponse">PostResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## DriverAssignment

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#LocationParam">LocationParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#VehicleParam">VehicleParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Location">Location</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationDriverAssignmentAssignResponse">OptimizationDriverAssignmentAssignResponse</a>

Methods:

- <code title="post /optimization/driver-assignment/v1">client.Optimization.DriverAssignment.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationDriverAssignmentService.Assign">Assign</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationDriverAssignmentAssignParams">OptimizationDriverAssignmentAssignParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationDriverAssignmentAssignResponse">OptimizationDriverAssignmentAssignResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## V2

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#JobParam">JobParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#ShipmentParam">ShipmentParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationV2GetResultResponse">OptimizationV2GetResultResponse</a>

Methods:

- <code title="get /optimization/v2/result">client.Optimization.V2.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationV2Service.GetResult">GetResult</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationV2GetResultParams">OptimizationV2GetResultParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationV2GetResultResponse">OptimizationV2GetResultResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /optimization/v2">client.Optimization.V2.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationV2Service.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#OptimizationV2SubmitParams">OptimizationV2SubmitParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#PostResponse">PostResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Geofence

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceEntityCreateParam">GeofenceEntityCreateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#Geofence">Geofence</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceNewResponse">GeofenceNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceGetResponse">GeofenceGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceListResponse">GeofenceListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceContainsResponse">GeofenceContainsResponse</a>

Methods:

- <code title="post /geofence">client.Geofence.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceNewParams">GeofenceNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceNewResponse">GeofenceNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /geofence/{id}">client.Geofence.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceGetParams">GeofenceGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceGetResponse">GeofenceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /geofence/{id}">client.Geofence.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceUpdateParams">GeofenceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /geofence/list">client.Geofence.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceListParams">GeofenceListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceListResponse">GeofenceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /geofence/{id}">client.Geofence.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceDeleteParams">GeofenceDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /geofence/contain">client.Geofence.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceService.Contains">Contains</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceContainsParams">GeofenceContainsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceContainsResponse">GeofenceContainsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Console

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#PolygonGeojson">PolygonGeojson</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceConsolePreviewResponse">GeofenceConsolePreviewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceConsoleSearchResponse">GeofenceConsoleSearchResponse</a>

Methods:

- <code title="post /geofence/console/preview">client.Geofence.Console.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceConsoleService.Preview">Preview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceConsolePreviewParams">GeofenceConsolePreviewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceConsolePreviewResponse">GeofenceConsolePreviewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /geofence/console/search">client.Geofence.Console.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceConsoleService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceConsoleSearchParams">GeofenceConsoleSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceConsoleSearchResponse">GeofenceConsoleSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Batch

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceBatchNewResponse">GeofenceBatchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceBatchQueryResponse">GeofenceBatchQueryResponse</a>

Methods:

- <code title="post /geofence/batch">client.Geofence.Batch.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceBatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceBatchNewParams">GeofenceBatchNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceBatchNewResponse">GeofenceBatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /geofence/batch">client.Geofence.Batch.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceBatchService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceBatchDeleteParams">GeofenceBatchDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SimpleResp">SimpleResp</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /geofence/batch">client.Geofence.Batch.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceBatchService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceBatchQueryParams">GeofenceBatchQueryParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#GeofenceBatchQueryResponse">GeofenceBatchQueryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Discover

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DiscoverListResponse">DiscoverListResponse</a>

Methods:

- <code title="get /discover">client.Discover.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DiscoverService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DiscoverListParams">DiscoverListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DiscoverListResponse">DiscoverListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Browse

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BrowseSearchResponse">BrowseSearchResponse</a>

Methods:

- <code title="get /browse">client.Browse.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BrowseService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BrowseSearchParams">BrowseSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BrowseSearchResponse">BrowseSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Mdm

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MdmNewDistanceMatrixResponse">MdmNewDistanceMatrixResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MdmGetDistanceMatrixStatusResponse">MdmGetDistanceMatrixStatusResponse</a>

Methods:

- <code title="post /mdm/create">client.Mdm.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MdmService.NewDistanceMatrix">NewDistanceMatrix</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MdmNewDistanceMatrixParams">MdmNewDistanceMatrixParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MdmNewDistanceMatrixResponse">MdmNewDistanceMatrixResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /mdm/status">client.Mdm.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MdmService.GetDistanceMatrixStatus">GetDistanceMatrixStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MdmGetDistanceMatrixStatusParams">MdmGetDistanceMatrixStatusParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MdmGetDistanceMatrixStatusResponse">MdmGetDistanceMatrixStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Isochrone

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#IsochroneComputeResponse">IsochroneComputeResponse</a>

Methods:

- <code title="get /isochrone/json">client.Isochrone.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#IsochroneService.Compute">Compute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#IsochroneComputeParams">IsochroneComputeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#IsochroneComputeResponse">IsochroneComputeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Restrictions

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RichGroupDtoRequestParam">RichGroupDtoRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RichGroupDtoResponse">RichGroupDtoResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionDeleteResponse">RestrictionDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionListPaginatedResponse">RestrictionListPaginatedResponse</a>

Methods:

- <code title="post /restrictions/{restriction_type}">client.Restrictions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, restrictionType <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionNewParamsRestrictionType">RestrictionNewParamsRestrictionType</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionNewParams">RestrictionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RichGroupDtoResponse">RichGroupDtoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /restrictions/{id}">client.Restrictions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionGetParams">RestrictionGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RichGroupDtoResponse">RichGroupDtoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /restrictions/{id}">client.Restrictions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionUpdateParams">RestrictionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RichGroupDtoResponse">RichGroupDtoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /restrictions">client.Restrictions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionListParams">RestrictionListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RichGroupDtoResponse">RichGroupDtoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /restrictions/{id}">client.Restrictions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionDeleteParams">RestrictionDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionDeleteResponse">RestrictionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /restrictions/list">client.Restrictions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionService.ListPaginated">ListPaginated</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionListPaginatedParams">RestrictionListPaginatedParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionListPaginatedResponse">RestrictionListPaginatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /restrictions/{id}/state">client.Restrictions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionService.SetState">SetState</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionSetStateParams">RestrictionSetStateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RichGroupDtoResponse">RichGroupDtoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RestrictionsItems

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionsItemListResponse">RestrictionsItemListResponse</a>

Methods:

- <code title="get /restrictions_items">client.RestrictionsItems.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionsItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionsItemListParams">RestrictionsItemListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RestrictionsItemListResponse">RestrictionsItemListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Distancematrix

## Json

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DistancematrixJsonGetResponse">DistancematrixJsonGetResponse</a>

Methods:

- <code title="post /distancematrix/json">client.Distancematrix.Json.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DistancematrixJsonService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /distancematrix/json">client.Distancematrix.Json.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DistancematrixJsonService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DistancematrixJsonGetParams">DistancematrixJsonGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DistancematrixJsonGetResponse">DistancematrixJsonGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Autocomplete

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AutocompleteSuggestResponse">AutocompleteSuggestResponse</a>

Methods:

- <code title="get /autocomplete">client.Autocomplete.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AutocompleteService.Suggest">Suggest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AutocompleteSuggestParams">AutocompleteSuggestParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AutocompleteSuggestResponse">AutocompleteSuggestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Navigation

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#NavigationGetRouteResponse">NavigationGetRouteResponse</a>

Methods:

- <code title="get /navigation/json">client.Navigation.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#NavigationService.GetRoute">GetRoute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#NavigationGetRouteParams">NavigationGetRouteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#NavigationGetRouteResponse">NavigationGetRouteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Map

Methods:

- <code title="post /map/segments">client.Map.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MapService.NewSegment">NewSegment</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Autosuggest

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AutosuggestSuggestResponse">AutosuggestSuggestResponse</a>

Methods:

- <code title="get /autosuggest">client.Autosuggest.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AutosuggestService.Suggest">Suggest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AutosuggestSuggestParams">AutosuggestSuggestParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AutosuggestSuggestResponse">AutosuggestSuggestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Directions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DirectionComputeRouteResponse">DirectionComputeRouteResponse</a>

Methods:

- <code title="post /directions/json">client.Directions.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DirectionService.ComputeRoute">ComputeRoute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DirectionComputeRouteParams">DirectionComputeRouteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#DirectionComputeRouteResponse">DirectionComputeRouteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Batch

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BatchNewResponse">BatchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BatchGetResponse">BatchGetResponse</a>

Methods:

- <code title="post /batch">client.Batch.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BatchNewParams">BatchNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BatchNewResponse">BatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /batch">client.Batch.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BatchGetParams">BatchGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#BatchGetResponse">BatchGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Multigeocode

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodeSearchResponse">MultigeocodeSearchResponse</a>

Methods:

- <code title="post /multigeocode/search">client.Multigeocode.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodeService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodeSearchParams">MultigeocodeSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodeSearchResponse">MultigeocodeSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Place

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#PlaceItemParam">PlaceItemParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#PlaceItem">PlaceItem</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceNewResponse">MultigeocodePlaceNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceGetResponse">MultigeocodePlaceGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceUpdateResponse">MultigeocodePlaceUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceDeleteResponse">MultigeocodePlaceDeleteResponse</a>

Methods:

- <code title="post /multigeocode/place">client.Multigeocode.Place.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceNewParams">MultigeocodePlaceNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceNewResponse">MultigeocodePlaceNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /multigeocode/place/{docId}">client.Multigeocode.Place.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, docID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceGetParams">MultigeocodePlaceGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceGetResponse">MultigeocodePlaceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /multigeocode/place/{docId}">client.Multigeocode.Place.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, docID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceUpdateParams">MultigeocodePlaceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceUpdateResponse">MultigeocodePlaceUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /multigeocode/place/{docId}">client.Multigeocode.Place.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, docID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceDeleteParams">MultigeocodePlaceDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#MultigeocodePlaceDeleteResponse">MultigeocodePlaceDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Revgeocode

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RevgeocodeGetResponse">RevgeocodeGetResponse</a>

Methods:

- <code title="get /revgeocode">client.Revgeocode.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RevgeocodeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RevgeocodeGetParams">RevgeocodeGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RevgeocodeGetResponse">RevgeocodeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# RouteReport

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RouteReportNewResponse">RouteReportNewResponse</a>

Methods:

- <code title="post /route_report">client.RouteReport.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RouteReportService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RouteReportNewParams">RouteReportNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#RouteReportNewResponse">RouteReportNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SnapToRoads

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SnapToRoadSnapResponse">SnapToRoadSnapResponse</a>

Methods:

- <code title="get /snapToRoads/json">client.SnapToRoads.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SnapToRoadService.Snap">Snap</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SnapToRoadSnapParams">SnapToRoadSnapParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#SnapToRoadSnapResponse">SnapToRoadSnapResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Postalcode

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#PostalcodeGetCoordinatesResponse">PostalcodeGetCoordinatesResponse</a>

Methods:

- <code title="post /postalcode">client.Postalcode.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#PostalcodeService.GetCoordinates">GetCoordinates</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#PostalcodeGetCoordinatesParams">PostalcodeGetCoordinatesParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#PostalcodeGetCoordinatesResponse">PostalcodeGetCoordinatesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Areas

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AreaListResponse">AreaListResponse</a>

Methods:

- <code title="get /areas">client.Areas.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AreaService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AreaListParams">AreaListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#AreaListResponse">AreaListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Lookup

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#LookupGetResponse">LookupGetResponse</a>

Methods:

- <code title="get /lookup">client.Lookup.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#LookupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#LookupGetParams">LookupGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go">nextbillionsdk</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/nextbillion-sdk-go#LookupGetResponse">LookupGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
