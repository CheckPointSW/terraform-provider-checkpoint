---
layout: "checkpoint"
page_title: "checkpoint_management_resource_uri"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-resource-uri"
description: |-
This resource allows you to execute Check Point Resource Uri.
---

# checkpoint_management_resource_uri

This resource allows you to execute Check Point Resource Uri.

## Example Usage


```hcl
resource "checkpoint_management_resource_uri" "uri" {
  name = "newUriResource"
  use_this_resource_to = "optimize_url_logging"
  connection_methods = {
    transparent = "false"
    tunneling = "true"
    proxy = false
  }
  uri_match_specification_type = "wildcards"
  match_wildcards {
    host = "hostName"
    path = "pathName"
    query =  "query"
    schemes {
      gopher = true
      other = "string2"
    }
    methods {
      get = true
      post = true
      head = true
      put = true
      other = "done7"
    }
  }
  action {
    strip_activex_tags =  true
    strip_applet_tags = true
    strip_ftp_links = true
    strip_port_strings = true
    strip_script_tags = true

  }
  soap = {
    inspection = "allow_all_soap_requests"
    file_id = "scheme1"
    track_connections = "mail_alert"
  }
  cvp {
    allowed_to_modify_content = true
    enable_cvp =  false
    reply_order = "return_data_after_content_is_approved"
    server = "serverName"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `use_this_resource_to` - (Optional) Select the use of the URI resource. 
* `connection_methods` - (Optional) Connection methods.connection_methods blocks are documented below.
* `uri_match_specification_type` - (Optional) The type can be Wild Cards or UFP, where a UFP server holds categories of forbidden web sites. 
* `exception_track` - (Optional) Configures how to track connections that match this rule but fail the content security checks. An example of an exception is a connection with an unsupported scheme or method. 
* `match_ufp` - (Optional) Match - UFP settings.match_ufp blocks are documented below.
* `match_wildcards` - (Optional) Match - Wildcards settings.match_wildcards blocks are documented below.
* `action` - (Optional) Action settings.action blocks are documented below.
* `cvp` - (Optional) CVP settings.cvp blocks are documented below.
* `soap` - (Optional) SOAP settings.soap blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`connection_methods` supports the following:

* `transparent` - (Optional) The security server is invisible to the client that originates the connection, and to the server. The Transparent connection method is the most secure. 
* `proxy` - (Optional) The Resource is applied when people specify the Check Point Security Gateway as a proxy in their browser. 
* `tunneling` - (Optional) The Resource is applied when people specify the Security Gateway as a proxy in their browser, and is used for connections where Security Gateway cannot examine the contents of the packets, not even the URL. 


`match_ufp` supports the following:

* `server` - (Optional) The UID or Name of the UFP server that is an OPSEC certified third party application that checks URLs against a list of permitted categories. 
* `caching_control` - (Optional) Specifies if and how caching is to be enabled. 
* `ignore_ufp_server_after_failure` - (Optional) The UFP server will be ignored after numerous UFP server connections were unsuccessful. 
* `number_of_failures_before_ignore` - (Optional) Signifies at what point the UFP server should be ignored. 
* `timeout_before_reconnecting` - (Optional) The amount of time that must pass before a UFP server connection should be attempted. 


`match_wildcards` supports the following:

* `schemes` - (Optional) Select the URI Schemes to which this resource applies.schemes blocks are documented below.
* `methods` - (Optional) Select the URI Schemes to which this resource applies.methods blocks are documented below.
* `host` - (Optional) The functionality of the Host parameter depends on the DNS setup of the addressed server. For the host, only the IP address or the full DNS name should be used. 
* `path` - (Optional) Name matching is based on appending the file name in the request to the current working directory (unless the file name is already a full path name) and comparing the result to the path specified in the Resource definition. 
* `query` - (Optional) The parameters that are sent to the URI when it is accessed. 


`action` supports the following:

* `replacement_uri` - (Optional) If the Action in a rule which uses this resource is Drop or Reject, then the Replacement URI is displayed instead of the one requested by the user. 
* `strip_script_tags` - (Optional) Strip JAVA scripts. 
* `strip_applet_tags` - (Optional) Strip JAVA applets. 
* `strip_activex_tags` - (Optional) Strip activeX tags. 
* `strip_ftp_links` - (Optional) Strip ftp links. 
* `strip_port_strings` - (Optional) Strip ports. 


`cvp` supports the following:

* `enable_cvp` - (Optional) Select to enable the Content Vectoring Protocol. 
* `server` - (Optional) The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application. 
* `allowed_to_modify_content` - (Optional) Configures the CVP server to inspect but not modify content. 
* `send_http_headers_to_cvp` - (Optional) Select, if you would like the CVP server to check the HTTP headers of the message packets. 
* `reply_order` - (Optional) Designates when the CVP server returns data to the Security Gateway security server. 
* `send_http_request_to_cvp` - (Optional) Used to protect against undesirable content in the HTTP request, for example, when inspecting peer-to-peer connections. 
* `send_only_unsafe_file_types` - (Optional) Improves the performance of the CVP server. This option does not send to the CVP server traffic that is considered safe. 


`soap` supports the following:

* `inspection` - (Optional) Allow all SOAP Requests, or Allow only SOAP requests specified in the following file-id. 
* `file_id` - (Optional) A file containing SOAP requests. 
* `track_connections` - (Optional) The method of tracking SOAP connections. 


`schemes` supports the following:

* `http` - (Optional) Http scheme. 
* `ftp` - (Optional) Ftp scheme. 
* `gopher` - (Optional) Gopher scheme. 
* `mailto` - (Optional) Mailto scheme. 
* `news` - (Optional) News scheme. 
* `wais` - (Optional) Wais scheme. 
* `other` - (Optional) You can specify another scheme in the Other field. You can use wildcards. 


`methods` supports the following:

* `get` - (Optional) GET method. 
* `post` - (Optional) POST method. 
* `head` - (Optional) HEAD method. 
* `put` - (Optional) PUT method. 
* `other` - (Optional) You can specify another method in the Other field. You can use wildcards. 
