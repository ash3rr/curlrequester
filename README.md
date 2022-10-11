# Curl Requester

Send a CURL request to create a webhook in BitBucket Server.

Requires Bit Bucket Secret & token deployed into Kubernetes which stores the secret & token, the app_name which comes from an Argo Workflow template, a configmap
which stores the API URL as well as the webhook URL. 
