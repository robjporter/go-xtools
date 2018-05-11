package xhttp

type Definition struct {
	short string
	long  string
}

var (
	codes map[int]Definition
)

func init() {
	codes = make(map[int]Definition)
	addDefaultCodes()
}

func addDefaultCodes() {
	codes[100] = Definition{short: "Continue", long: "The server has received the request headers and the client should proceed to send the request body (in the case of a request for which a body needs to be sent; for example, a POST request). Sending a large request body to a server after a request has been rejected for inappropriate headers would be inefficient. To have a server check the request's headers, a client must send Expect: 100-continue as a header in its initial request and receive a 100 Continue status code in response before sending the body. The response 417 Expectation Failed indicates the request should not be continued."}
	codes[101] = Definition{short: "Switching Protocols", long: "The requester has asked the server to switch protocols and the server has agreed to do so."}
	codes[102] = Definition{short: "Processing (WebDAV; RFC 2518)", long: "A WebDAV request may contain many sub-requests involving file operations, requiring a long time to complete the request. This code indicates that the server has received and is processing the request, but no response is available yet.[6] This prevents the client from timing out and assuming the request was lost."}
	codes[200] = Definition{short: "OK", long: "Standard response for successful HTTP requests. The actual response will depend on the request method used. In a GET request, the response will contain an entity corresponding to the requested resource. In a POST request, the response will contain an entity describing or containing the result of the action."}
	codes[201] = Definition{short: "Created", long: "The request has been fulfilled, resulting in the creation of a new resource."}
	codes[202] = Definition{short: "Accepted", long: "The request has been accepted for processing, but the processing has not been completed. The request might or might not be eventually acted upon, and may be disallowed when processing occurs."}
	codes[203] = Definition{short: "Non-Authoritative Information", long: "The server is a transforming proxy (e.g. a Web accelerator) that received a 200 OK from its origin, but is returning a modified version of the origin's response."}
	codes[204] = Definition{short: "No Content", long: "The server successfully processed the request and is not returning any content."}
	codes[205] = Definition{short: "Reset Content", long: "The server successfully processed the request, but is not returning any content. Unlike a 204 response, this response requires that the requester reset the document view."}
	codes[206] = Definition{short: "Partial Content (RFC 7233)", long: "The server is delivering only part of the resource (byte serving) due to a range header sent by the client. The range header is used by HTTP clients to enable resuming of interrupted downloads, or split a download into multiple simultaneous streams."}
	codes[207] = Definition{short: "Multi-Status (WebDAV; RFC 4918)", long: "The message body that follows is an XML message and can contain a number of separate response codes, depending on how many sub-requests were made."}
	codes[208] = Definition{short: "Already Reported (WebDAV; RFC 5842)", long: "The members of a DAV binding have already been enumerated in a previous reply to this request, and are not being included again."}
	codes[226] = Definition{short: "IM Used (RFC 3229)", long: "The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance."}
	codes[300] = Definition{short: "Multiple Choices", long: "Indicates multiple options for the resource from which the client may choose (via agent-driven content negotiation). For example, this code could be used to present multiple video format options, to list files with different filename extensions, or to suggest word-sense disambiguation."}
	codes[301] = Definition{short: "Moved Permanently", long: "This and all future requests should be directed to the given URI."}
	codes[302] = Definition{short: "Found", long: "This is an example of industry practice contradicting the standard. The HTTP/1.0 specification (RFC 1945) required the client to perform a temporary redirect (the original describing phrase was 'Moved Temporarily'), but popular browsers implemented 302 with the functionality of a 303 See Other. Therefore, HTTP/1.1 added status codes 303 and 307 to distinguish between the two behaviours.[22] However, some Web applications and frameworks use the 302 status code as if it were the 303."}
	codes[303] = Definition{short: "See Other (since HTTP/1.1)", long: "The response to the request can be found under another URI using a GET method. When received in response to a POST (or PUT/DELETE), the client should presume that the server has received the data and should issue a redirect with a separate GET message."}
	codes[304] = Definition{short: "Not Modified (RFC 7232)", long: "Indicates that the resource has not been modified since the version specified by the request headers If-Modified-Since or If-None-Match. In such case, there is no need to retransmit the resource since the client still has a previously-downloaded copy."}
	codes[305] = Definition{short: "Use Proxy (since HTTP/1.1)", long: "The requested resource is available only through a proxy, the address for which is provided in the response. Many HTTP clients (such as Mozilla[26] and Internet Explorer) do not correctly handle responses with this status code, primarily for security reasons."}
	codes[306] = Definition{short: "Switch Proxy", long: "No longer used. Originally meant 'Subsequent requests should use the specified proxy.'"}
	codes[307] = Definition{short: "Temporary Redirect (since HTTP/1.1)", long: "In this case, the request should be repeated with another URI; however, future requests should still use the original URI. In contrast to how 302 was historically implemented, the request method is not allowed to be changed when reissuing the original request. For example, a POST request should be repeated using another POST request."}
	codes[308] = Definition{short: "Permanent Redirect (RFC 7538)", long: "The request and all future requests should be repeated using another URI. 307 and 308 parallel the behaviors of 302 and 301, but do not allow the HTTP method to change. So, for example, submitting a form to a permanently redirected resource may continue smoothly."}
	codes[400] = Definition{short: "Bad Request", long: "The server cannot or will not process the request due to an apparent client error (e.g., malformed request syntax, too large size, invalid request message framing, or deceptive request routing)."}
	codes[401] = Definition{short: "Unauthorized (RFC 7235)", long: "Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided. The response must include a WWW-Authenticate header field containing a challenge applicable to the requested resource. See Basic access authentication and Digest access authentication.[33] 401 semantically means 'unauthenticated', i.e. the user does not have the necessary credentials. Note: Some sites issue HTTP 401 when an IP address is banned from the website (usually the website domain) and that specific address is refused permission to access a website."}
	codes[402] = Definition{short: "Payment Required", long: "Reserved for future use. The original intention was that this code might be used as part of some form of digital cash or micropayment scheme, but that has not happened, and this code is not usually used. Google Developers API uses this status if a particular developer has exceeded the daily limit on requests."}
	codes[403] = Definition{short: "Forbidden", long: "The request was valid, but the server is refusing action. The user might not have the necessary permissions for a resource."}
	codes[404] = Definition{short: "Not Found", long: "The requested resource could not be found but may be available in the future. Subsequent requests by the client are permissible."}
	codes[405] = Definition{short: "Method Not Allowed", long: "A request method is not supported for the requested resource; for example, a GET request on a form that requires data to be presented via POST, or a PUT request on a read-only resource."}
	codes[406] = Definition{short: "Not Acceptable", long: "The requested resource is capable of generating only content not acceptable according to the Accept headers sent in the request."}
	codes[407] = Definition{short: "Proxy Authentication Required (RFC 7235)", long: "The client must first authenticate itself with the proxy."}
	codes[408] = Definition{short: "Request Time-out", long: "The server timed out waiting for the request. According to HTTP specifications: 'The client did not produce a request within the time that the server was prepared to wait. The client MAY repeat the request without modifications at any later time.'"}
	codes[409] = Definition{short: "Conflict", long: "Indicates that the request could not be processed because of conflict in the request, such as an edit conflict between multiple simultaneous updates."}
	codes[410] = Definition{short: "Gone", long: "Indicates that the resource requested is no longer available and will not be available again. This should be used when a resource has been intentionally removed and the resource should be purged. Upon receiving a 410 status code, the client should not request the resource in the future. Clients such as search engines should remove the resource from their indices. Most use cases do not require clients and search engines to purge the resource, and a '404 Not Found' may be used instead."}
	codes[411] = Definition{short: "Length Required", long: "The request did not specify the length of its content, which is required by the requested resource."}
	codes[412] = Definition{short: "Precondition Failed (RFC 7232)", long: "The server does not meet one of the preconditions that the requester put on the request."}
	codes[413] = Definition{short: "Payload Too Large (RFC 7231)", long: "The request is larger than the server is willing or able to process. Previously called 'Request Entity Too Large'."}
	codes[414] = Definition{short: "URI Too Long (RFC 7231)", long: "The URI provided was too long for the server to process. Often the result of too much data being encoded as a query-string of a GET request, in which case it should be converted to a POST request.[44] Called 'Request-URI Too Long' previously."}
	codes[415] = Definition{short: "Unsupported Media Type", long: "The request entity has a media type which the server or resource does not support. For example, the client uploads an image as image/svg+xml, but the server requires that images use a different format."}
	codes[416] = Definition{short: "Range Not Satisfiable (RFC 7233)", long: "The client has asked for a portion of the file (byte serving), but the server cannot supply that portion. For example, if the client asked for a part of the file that lies beyond the end of the file.[46] Called 'Requested Range Not Satisfiable' previously."}
	codes[417] = Definition{short: "Expectation Failed", long: "The server cannot meet the requirements of the Expect request-header field."}
	codes[418] = Definition{short: "I'm a teapot (RFC 2324)", long: "This code was defined in 1998 as one of the traditional IETF April Fools' jokes, in RFC 2324, Hyper Text Coffee Pot Control Protocol, and is not expected to be implemented by actual HTTP servers. The RFC specifies this code should be returned by teapots requested to brew coffee. This HTTP status is used as an Easter egg in some websites, including Google.com."}
	codes[421] = Definition{short: "Misdirected Request (RFC 7540)", long: "The request was directed at a server that is not able to produce a response (for example because a connection reuse)."}
	codes[422] = Definition{short: "Unprocessable Entity (WebDAV; RFC 4918)", long: "The request was well-formed but was unable to be followed due to semantic errors."}
	codes[423] = Definition{short: "Locked (WebDAV; RFC 4918)", long: "The resource that is being accessed is locked."}
	codes[424] = Definition{short: "Failed Dependency (WebDAV; RFC 4918)", long: "The request failed due to failure of a previous request (e.g., a PROPPATCH)."}
	codes[426] = Definition{short: "Upgrade Required", long: "The client should switch to a different protocol such as TLS/1.0, given in the Upgrade header field."}
	codes[428] = Definition{short: "Precondition Required (RFC 6585)", long: "The origin server requires the request to be conditional. Intended to prevent 'the 'lost update' problem, where a client GETs a resource's state, modifies it, and PUTs it back to the server, when meanwhile a third party has modified the state on the server, leading to a conflict.'"}
	codes[429] = Definition{short: "Too Many Requests (RFC 6585)", long: "The user has sent too many requests in a given amount of time. Intended for use with rate-limiting schemes."}
	codes[431] = Definition{short: "Request Header Fields Too Large (RFC 6585)", long: "The server is unwilling to process the request because either an individual header field, or all the header fields collectively, are too large."}
	codes[451] = Definition{short: "Unavailable For Legal Reasons (RFC 7725)", long: "A server operator has received a legal demand to deny access to a resource or to a set of resources that includes the requested resource. The code 451 was chosen as a reference to the novel Fahrenheit 451."}

	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
	codes[206] = Definition{short: "", long: ""}
}

func GetCodeInfo(code int) Definition {
	return codes[code]
}

func GetCodeShortInfo(code int) string {
	return codes[code].short
}

func GetCodeLongInfo(code int) string {
	return codes[code].long
}

/*
500 Internal Server Error
A generic error message, given when an unexpected condition was encountered and no more specific message is suitable.[57]
501 Not Implemented
The server either does not recognize the request method, or it lacks the ability to fulfill the request. Usually this implies future availability (e.g., a new feature of a web-service API).[58]
502 Bad Gateway
The server was acting as a gateway or proxy and received an invalid response from the upstream server.[59]
503 Service Unavailable
The server is currently unavailable (because it is overloaded or down for maintenance). Generally, this is a temporary state.[60]
504 Gateway Time-out
The server was acting as a gateway or proxy and did not receive a timely response from the upstream server.[61]
505 HTTP Version Not Supported
The server does not support the HTTP protocol version used in the request.[62]
506 Variant Also Negotiates (RFC 2295)
Transparent content negotiation for the request results in a circular reference.[63]
507 Insufficient Storage (WebDAV; RFC 4918)
The server is unable to store the representation needed to complete the request.[15]
508 Loop Detected (WebDAV; RFC 5842)
The server detected an infinite loop while processing the request (sent in lieu of 208 Already Reported).
510 Not Extended (RFC 2774)
Further extensions to the request are required for the server to fulfill it.[64]
511 Network Authentication Required (RFC 6585)
The client needs to authenticate to gain network access. Intended for use by intercepting proxies used to control access to the network (e.g., "captive portals" used to require agreement to Terms of Service before granting full Internet access via a Wi-Fi hotspot).[53]


103 Checkpoint
Used in the resumable requests proposal to resume aborted PUT or POST requests.[65]
103 Early Hints
Used to return some response headers before entire HTTP response.[66][67]
420 Method Failure (Spring Framework)
A deprecated response used by the Spring Framework when a method has failed.[68]
420 Enhance Your Calm (Twitter)
Returned by version 1 of the Twitter Search and Trends API when the client is being rate limited; versions 1.1 and later use the 429 Too Many Requests response code instead.[69]
450 Blocked by Windows Parental Controls (Microsoft)
The Microsoft extension code indicated when Windows Parental Controls are turned on and are blocking access to the given webpage.[70]
498 Invalid Token (Esri)
Returned by ArcGIS for Server. Code 498 indicates an expired or otherwise invalid token.[71]
499 Token Required (Esri)
Returned by ArcGIS for Server. Code 499 indicates that a token is required but was not submitted.[71]
509 Bandwidth Limit Exceeded (Apache Web Server/cPanel)
The server has exceeded the bandwidth specified by the server administrator; this is often used by shared hosting providers to limit the bandwidth of customers.[72]
530 Site is frozen
Used by the Pantheon web platform to indicate a site that has been frozen due to inactivity.[73]
598 (Informal convention) Network read timeout error
Used by some HTTP proxies to signal a network read timeout behind the proxy to a client in front of the proxy.[74][75]
599 (Informal convention) Network connect timeout error
Used to indicate when the connection to the network times out.[76][citation needed]

440 Login Time-out
The client's session has expired and must log in again.[77]
449 Retry With
The server cannot honour the request because the user has not provided the required information.[78]
451 Redirect
Used in Exchange ActiveSync when either a more efficient server is available or the server cannot access the users' mailbox.[79] The client is expected to re-run the HTTP AutoDiscover operation to find a more appropriate server.[80]

444 No Response
Used to indicate that the server has returned no information to the client and closed the connection.
495 SSL Certificate Error
An expansion of the 400 Bad Request response code, used when the client has provided an invalid client certificate.
496 SSL Certificate Required
An expansion of the 400 Bad Request response code, used when a client certificate is required but not provided.
497 HTTP Request Sent to HTTPS Port
An expansion of the 400 Bad Request response code, used when the client has made a HTTP request to a port listening for HTTPS requests.
499 Client Closed Request
Used when the client has closed the request before the server could send a response.


520 Unknown Error
The 520 error is used as a "catch-all response for when the origin server returns something unexpected", listing connection resets, large headers, and empty or invalid responses as common triggers.
521 Web Server Is Down
The origin server has refused the connection from Cloudflare.
522 Connection Timed Out
Cloudflare could not negotiate a TCP handshake with the origin server.
523 Origin Is Unreachable
Cloudflare could not reach the origin server; for example, if the DNS records for the origin server are incorrect.
524 A Timeout Occurred
Cloudflare was able to complete a TCP connection to the origin server, but did not receive a timely HTTP response.
525 SSL Handshake Failed
Cloudflare could not negotiate a SSL/TLS handshake with the origin server.
526 Invalid SSL Certificate
Cloudflare could not validate the SSL/TLS certificate that the origin server presented.
527 Railgun Error
Error 527 indicates that the requests timeout or failed after the WAN connection has been established.[84]




110	Restart marker replay . In this case, the text is exact and not left to the particular implementation; it must read: MARK yyyy = mmmm where yyyy is User-process data stream marker, and mmmm server's equivalent marker (note the spaces between markers and "=").
120	Service ready in nnn minutes.
125	Data connection already open; transfer starting.
150	File status okay; about to open data connection.
200 Series	The requested action has been successfully completed.
202	Command not implemented, superfluous at this site.
211	System status, or system help reply.
212	Directory status.
213	File status.
214	Help message. Explains how to use the server or the meaning of a particular non-standard command. This reply is useful only to the human user.
215	NAME system type. Where NAME is an official system name from the registry kept by IANA.
220	Service ready for new user.
221	Service closing control connection.
225	Data connection open; no transfer in progress.
226	Closing data connection. Requested file action successful (for example, file transfer or file abort).
227	Entering Passive Mode (h1,h2,h3,h4,p1,p2).
228	Entering Long Passive Mode (long address, port).
229	Entering Extended Passive Mode (|||port|).
230	User logged in, proceed. Logged out if appropriate.
231	User logged out; service terminated.
232	Logout command noted, will complete when transfer done.
234	Specifies that the server accepts the authentication mechanism specified by the client, and the exchange of security data is complete. A higher level nonstandard code created by Microsoft.
250	Requested file action okay, completed.
257	"PATHNAME" created.
300 Series	The command has been accepted, but the requested action is on hold, pending receipt of further information.
331	User name okay, need password.
332	Need account for login.
350	Requested file action pending further information
400 Series	The command was not accepted and the requested action did not take place, but the error condition is temporary and the action may be requested again.
421	Service not available, closing control connection. This may be a reply to any command if the service knows it must shut down.
425	Can't open data connection.
426	Connection closed; transfer aborted.
430	Invalid username or password
434	Requested host unavailable.
450	Requested file action not taken.
451	Requested action aborted. Local error in processing.
452	Requested action not taken. Insufficient storage space in system.File unavailable (e.g., file busy).
500 Series	Syntax error, command unrecognized and the requested action did not take place. This may include errors such as command line too long.
501	Syntax error in parameters or arguments.
502	Command not implemented.
503	Bad sequence of commands.
504	Command not implemented for that parameter.
530	Not logged in.
532	Need account for storing files.
534	Could Not Connect to Server - Policy Requires SSL
550	Requested action not taken. File unavailable (e.g., file not found, no access).
551	Requested action aborted. Page type unknown.
552	Requested file action aborted. Exceeded storage allocation (for current directory or dataset).
553	Requested action not taken. File name not allowed.
600 Series	Replies regarding confidentiality and integrity
631	Integrity protected reply.
632	Confidentiality and integrity protected reply.
633	Confidentiality protected reply.
10000 Series	Common Winsock Error Codes
10054	Connection reset by peer. The connection was forcibly closed by the remote host.
10060	Cannot connect to remote server.
10061	Cannot connect to remote server. The connection is actively refused by the server.
10066	Directory not empty.
10068	Too many users, server is full.
*/
