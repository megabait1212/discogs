package discogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/mrjones/oauth"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "http://api.discogs.com/"
	userAgent      = "go-discogs/" + libraryVersion + " +https://github.com/dmikalova/go_discogs"

	// headerRateLimit     = "X-RateLimit-Limit"
	// headerRateRemaining = "X-RateLimit-Remaining"
	// headerRateReset     = "X-RateLimit-Reset"
)

var (
	accessToken   *oauth.AccessToken
	oauthConsumer *oauth.Consumer
)

// A Client manages communication with the Discogs API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests. Defaults to the public Discogs API, but can be set to a different endpoint. BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating with the Discogs API.
	UserAgent string

	// Rate specifies the current rate limit for the client as determined by the most recent API call.  If the client is used in a multi-user application, this rate may not always be up-to-date. Call RateLimit() to check the current rate.
	Rate Rate

	// Services used for talking to different parts of the GitHub API.
	Artist                       *ArtistService
	ArtistReleases               *ArtistReleasesService
	Identity                     *IdentityService
	Image                        *ImageService
	Label                        *LabelService
	LabelReleases                *LabelReleasesService
	MarketplaceFee               *MarketplaceFeeService
	MarketplaceListing           *MarketplaceListingService
	MarketplaceOrder             *MarketplaceOrderService
	MarketplacePriceSuggestions  *MarketplacePriceSuggestionsService
	Master                       *MasterService
	MasterVersions               *MasterVersionsService
	Release                      *ReleaseService
	User                         *UserService
	UserCollectionFields         *UserCollectionFieldsService
	UserCollectionFolder         *UserCollectionFolderService
	UserCollectionFolderReleases *UserCollectionFolderReleasesService
	UserCollectionFolders        *UserCollectionFoldersService
	UserContributions            *UserContributionsService
	UserInventory                *UserInventoryService
	UserSubmissions              *UserSubmissionsService
	UserWants                    *UserWantsService
}

// NewClient returns a new Discogs API client. If a nil httpClient is provided, http.DefaultClient will be used. To use API methods which require authentication, provide an http.Client that will perform the authentication for you,such as that provided by the goauth2 library.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}
	c.Artist = &ArtistService{client: c}
	c.ArtistReleases = &ArtistReleasesService{client: c}
	c.Label = &LabelService{client: c}
	c.LabelReleases = &LabelReleasesService{client: c}
	c.Identity = &IdentityService{client: c}
	c.Image = &ImageService{client: c}
	c.MarketplaceFee = &MarketplaceFeeService{client: c}
	c.MarketplaceListing = &MarketplaceListingService{client: c}
	c.MarketplaceOrder = &MarketplaceOrderService{client: c}
	c.MarketplacePriceSuggestions = &MarketplacePriceSuggestionsService{client: c}
	c.Master = &MasterService{client: c}
	c.MasterVersions = &MasterVersionsService{client: c}
	c.Release = &ReleaseService{client: c}
	c.User = &UserService{client: c}
	c.UserCollectionFields = &UserCollectionFieldsService{client: c}
	c.UserCollectionFolder = &UserCollectionFolderService{client: c}
	c.UserCollectionFolderReleases = &UserCollectionFolderReleasesService{client: c}
	c.UserCollectionFolders = &UserCollectionFoldersService{client: c}
	c.UserContributions = &UserContributionsService{client: c}
	c.UserInventory = &UserInventoryService{client: c}
	c.UserSubmissions = &UserSubmissionsService{client: c}
	c.UserWants = &UserWantsService{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr, in which case it is resolved relative to the BaseURL of the Client. Relative URLs should always be specified without a preceding slash.  If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// req.Header.Add("Accept", mediaTypeV3)
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response is a GitHub API response.  This wraps the standard http.Response
// returned from GitHub and provides convenient access to things like
// pagination links.
type Response struct {
	*http.Response

	// These fields provide the page values for paginating through a set of
	// results.  Any or all of these may be set to the zero value for
	// responses that are not part of a paginated set, or for which there
	// are no additional pages.

	// NextPage  int
	// PrevPage  int
	// FirstPage int
	// LastPage  int

	// Rate
}

// newResponse creats a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	// response.populatePageValues()
	// response.populateRate()
	return response
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	response := newResponse(resp)

	// c.Rate = response.Rate

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}
	return response, err
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	return errorResponse
}

/*
An ErrorResponse reports one or more errors caused by an API request.

GitHub API docs: http://developer.github.com/v3/#client-errors
*/
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
	Errors   []Error        `json:"errors"`  // more detail on individual errors
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message, r.Errors)
}

/*
An Error reports more details on an individual error in an ErrorResponse.
These are the possible validation error codes:

    missing:
        resource does not exist
    missing_field:
        a required field on a resource has not been set
    invalid:
        the formatting of a field is invalid
    already_exists:
        another resource has the same valid as this field

GitHub API docs: http://developer.github.com/v3/#client-errors
*/
type Error struct {
	Resource string `json:"resource"` // resource on which the error occurred
	Field    string `json:"field"`    // field on which the error occurred
	Code     string `json:"code"`     // validation error code
}

// Stub to fill out later.
type Rate struct {
	// The number of requests per hour the client is currently limited to.
	Limit int `json:"limit"`

	// The number of remaining requests the client can make this hour.
	Remaining int `json:"remaining"`

	// The time at which the current rate limit will reset.
	Reset Timestamp `json:"reset"`
}

// Stub to fill out later.
type RateLimits struct {
	// The rate limit for API requests. The default is 1 per second.
	Core *Rate `json:"core"`

	// The rate limit for search API requests.  Unauthenticated requests
	// are limited to 5 requests per minutes.  Authenticated requests are
	// limited to 20 per minute.
	//
	// GitHub API docs: https://developer.github.com/v3/search/#rate-limit
	Search *Rate `json:"search"`
}

// Stub to fill out later.
type Timestamp struct {
	time.Time
}

type AccessToken struct {
	Token          string
	Secret         string
	AdditionalData map[string]string
}

func GetAccessToken() (*oauth.AccessToken, *oauth.Consumer) {
	if accessToken == nil {
		getFromFile()
	}

	if oauthConsumer == nil {
		consumerKey := "rPJwszCjnGXfRIjFquYD"
		consumerSecret := "kTIjZmEnGCNiQqqRgyhgUfOJFHnSAVdI"

		oauthConsumer = oauth.NewConsumer(
			consumerKey,
			consumerSecret,
			oauth.ServiceProvider{
				RequestTokenUrl:   "http://api.discogs.com/oauth/request_token",
				AuthorizeTokenUrl: "http://www.discogs.com/oauth/authorize",
				AccessTokenUrl:    "http://api.discogs.com/oauth/access_token",
			},
		)
	}

	return accessToken, oauthConsumer

}

func getFromFile() {
	file := "/Users/dfoltin/.config/discogs/token.json"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("ReadFile Error")
		OAuth1()
		return
	}

	err = json.Unmarshal(data, &accessToken)
	if err != nil {
		fmt.Println("Unmarshall Error")
		OAuth1()
		return
	}

}

func OAuth1() {
	consumerKey := "rPJwszCjnGXfRIjFquYD"
	consumerSecret := "kTIjZmEnGCNiQqqRgyhgUfOJFHnSAVdI"

	oauthConsumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "http://api.discogs.com/oauth/request_token",
			AuthorizeTokenUrl: "http://www.discogs.com/oauth/authorize",
			AccessTokenUrl:    "http://api.discogs.com/oauth/access_token",
		},
	)

	requestToken, url, err := oauthConsumer.GetRequestTokenAndUrl("oob")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("(1) Go to: " + url)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here: ")

	verificationCode := ""
	fmt.Scanln(&verificationCode)

	accessToken, err = oauthConsumer.AuthorizeToken(requestToken, verificationCode)
	if err != nil {
		log.Fatal(err)
	}

	writeToken()
}

func writeToken() {
	file := "/Users/dfoltin/.config/discogs/token.json"
	tokenJSON, _ := json.MarshalIndent(accessToken, "", "	")

	err := ioutil.WriteFile(file, tokenJSON, 0644)
	if err != nil {
		panic(err)
	}
}
