package api_go_sdk

import "time"

// Api Client arguments to init a new instance
type ApiClientArgs struct {
	Port                    int
	Fingerprint             string
	Sid                     string
	Server                  string
	HttpDebugLevel          string
	ProxyHost               string
	ProxyPort               int
	ApiVersion              string
	IgnoreServerCertificate bool
	AcceptServerCertificate bool
	DebugFile               string
	Context                 string
	AutoPublish				bool
	Timeout 				time.Duration
	Sleep 				time.Duration
	UserAgent			string
}

/*
Init a new APIClientArgs
Construct a new ApiClientArgs instance with the given parameters.

Port: the port that is being used
Fingerprint: server's fingerprint
Sid: session id
Server: server's ip
ProxyHost: proxy's ip
ProxyPort: proxy port
ApiVersion: the version of the api
IgnoreServerCertificate: indicates that the client should not check the server's certificate
AcceptServerCertificate: indicates that the client should automatically accept and save the server's certificate
DebugFile: name of debug file
Context: which API to use - Management API = web_api (default) or GAIA API = gaia_api
Timeout: HTTP Client timeout value
*/
func APIClientArgs(port int, fingerprint string, sid string, server string, proxyHost string, proxyPort int, apiVersion string, ignoreServerCertificate bool, acceptServerCertificate bool, debugFile string, context string, timeout time.Duration, sleep time.Duration, userAgent string) ApiClientArgs {

	return ApiClientArgs{
		Port: port,
		Fingerprint: fingerprint,
		Sid: sid,
		Server: server,
		ProxyHost: proxyHost,
		ProxyPort: proxyPort,
		ApiVersion: apiVersion,
		IgnoreServerCertificate: ignoreServerCertificate,
		AcceptServerCertificate: acceptServerCertificate,
		DebugFile: debugFile,
		Context: context,
		Timeout: timeout,
		Sleep: sleep,
		UserAgent: userAgent,
	}
}
