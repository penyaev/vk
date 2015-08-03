package vk

import (
	"net/url"
	"strconv"
	"time"
)

const (
	paramCode         = "code"
	paramToken        = "access_token"
	paramVersion      = "v"
	paramAppID        = "client_id"
	paramScope        = "scope"
	paramRedirectURI  = "redirect_uri"
	paramDisplay      = "display"
	paramHttps        = "https"
	paramResponseType = "response_type"

	oauthHost         = "oauth.vk.com"
	oauthDisplay      = "page"
	oauthPath         = "/authorize/"
	oauthResponseType = "token"
	oauthRedirectURI  = "https://oauth.vk.com/blank.html"
	oauthScheme       = "https"

	defaultHost    = "api.vk.com"
	defaultPath    = "/method/"
	defaultScheme  = "https"
	defaultVersion = "5.35"
	defaultMethod  = "GET"
	defaultHttps   = "1"

	maxRequestsPerSecond = 3
	minimumRate          = time.Second / maxRequestsPerSecond
	methodExecute        = "execute"
	maxRequestRepeat     = 10
)

// int64s formats int64 as base10 string
func int64s(v int64) string {
	return strconv.FormatInt(v, 10)
}

type Client struct {
	httpClient HttpClient
}

type Request struct {
	Method string     `json:"method"`
	Token  string     `json:"token"`
	Values url.Values `json:"values"`
}

func (c *Client) SetHttpClient(httpClient HttpClient) {
	c.httpClient = httpClient
}

func (c *Client) addParams(values url.Values) {
	values.Add(paramVersion, defaultVersion)
	values.Add(paramHttps, defaultHttps)
}

type Auth struct {
	ID           int64
	Scope        Scope
	RedirectURI  string
	ResponseType string
	Display      string
}

func (a Auth) URL() string {
	u := url.URL{}
	u.Host = oauthHost
	u.Scheme = oauthScheme
	u.Path = oauthPath

	if len(a.RedirectURI) == 0 {
		a.RedirectURI = oauthRedirectURI
	}
	if len(a.ResponseType) == 0 {
		a.ResponseType = oauthResponseType
	}
	if len(a.Display) == 0 {
		a.Display = oauthDisplay
	}

	values := u.Query()
	values.Add(paramResponseType, a.ResponseType)
	values.Add(paramScope, a.Scope.String())
	values.Add(paramAppID, int64s(a.ID))
	values.Add(paramRedirectURI, a.RedirectURI)
	values.Add(paramVersion, defaultVersion)
	values.Add(paramDisplay, a.Display)

	u.RawQuery = values.Encode()
	return u.String()
}

func New() *Client {
	c := new(Client)
	c.SetHttpClient(defaultHttpClient)
	return c
}