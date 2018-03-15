// locations.go
package wpt

import (
	"errors"
)

const defaultURL = "http://www.webpagetest.org"

const (
	wptQueryLocations   = "getLocations.php"
	wptQueryRunTest     = "runtest.php"
	wptQueryCancelTest  = "cancelTest.php"
	wptQueryTestStatus  = "testStatus.php"
	wptQueryTestResults = "jsonResult.php"
	wptQueryTestHistory = "testlog.php"
)

const (
	wptStatusTestRunning   = 100
	wptStatusTestQueued    = 101
	wptStatusTestSuccess   = 200
	wptStatusTestNotFound  = 400
	wptStatusTestCancelled = 402
)

var (
	errCreateRequest = errors.New("Error creating request")
	errQueryServer   = errors.New("Error querying WPT server")
	errBadResponse   = errors.New("Bad response code from WPT server")
	errReadBody      = errors.New("Error reading response body")
)

type Location struct {
	Label         string         `json:"Label"`
	Location      string         `json:"location"`
	Browsers      string         `json:"Browsers"`
	RelayServer   string         `json:"relayServer"`
	RelayLocation string         `json:"relayLocation"`
	LabelShort    string         `json:"labelShort"`
	Default       bool           `json:"default"`
	Group         string         `json:"group"`
	PendingTests  map[string]int `json:"PendingTests"`
}

type WPTLocations struct {
	StatusCode int                 `json:"statusCode"`
	StatusText string              `json:"statusText"`
	Data       map[string]Location `json:"data"`
}

type IterationPages struct {
	Breakdown  string `json:"breakdown"`
	Checklist  string `json:"checklist"`
	Details    string `json:"details"`
	Domains    string `json:"domains"`
	ScreenShot string `json:"screenShot"`
}

type IterationThumbnails struct {
	Waterfall  string `json:"waterfall"`
	Checklist  string `json:"checklist"`
	ScreenShot string `json:"screenShot"`
}

type IterationImages struct {
	Checklist      string `json:"checklist"`
	ConnectionView string `json:"connectionView"`
	ScreenShot     string `json:"screenShot"`
	ScreenShotPng  string `json:"screenShotPng"`
	Waterfall      string `json:"waterfall"`
}

type IterationRawData struct {
	Headers      string `json:"headers"`
	PageData     string `json:"pageData"`
	RequestsData string `json:"requestsData"`
	ScriptTiming string `json:"scriptTiming"`
	Trace        string `json:"trace"`
	Utilization  string `json:"utilization"`
}

type IterationVideoFrame struct {
	VisuallyComplete int64  `json:"VisuallyComplete"`
	Image            string `json:"image"`
	Time             int64  `json:"time"`
}

type IterationDomain struct {
	Bytes       int64  `json:"bytes"`
	CDNProvider string `json:"cdn_provider"`
	Connections int    `json:"connections"`
	Requests    int    `json:"requests"`
}

type IterationBreakdown struct {
	Bytes             int64 `json:"bytes"`
	BytesUncompressed int64 `json:"bytesUncompressed"`
	Color             []int `json:"color"`
	Requests          int   `json:"requests"`
}

type IterationConsoleLog struct {
	Column int    `json:"column"`
	Level  string `json:"level"`
	Line   int    `json:"line"`
	Source string `json:"source"`
	Text   string `json:"text"`
	URL    string `json:"url"`
}

type IterationRequest struct {
	AllEnd           int64  `json:"all_end"`
	AllMs            int64  `json:"all_ms"`
	AllStart         int64  `json:"all_start"`
	BytesIn          int64  `json:"bytesIn"`
	BytesOut         int64  `json:"bytesOut"`
	CacheControl     string `json:"cacheControl"`
	CacheTime        int64  `json:"cache_time"`
	CdnProvider      string `json:"cdn_provider"`
	CertificateBytes int64  `json:"certificate_bytes"`
	ClientPort       int64  `json:"client_port"`
	ConnectEnd       int64  `json:"connect_end"`
	ConnectMs        int64  `json:"connect_ms"`
	ConnectStart     int64  `json:"connect_start"`
	ContentEncoding  string `json:"contentEncoding"`
	ContentType      string `json:"contentType"`
	DNSEnd           int64  `json:"dns_end"`
	DNSMs            int64  `json:"dns_ms"`
	DNSStart         int64  `json:"dns_start"`
	DownloadEnd      int64  `json:"download_end"`
	DownloadMs       int64  `json:"download_ms"`
	DownloadStart    int64  `json:"download_start"`
	Expires          string `json:"expires"`
	FullURL          string `json:"full_url"`
	GzipSave         int64  `json:"gzip_save"`
	GzipTotal        int64  `json:"gzip_total"`
	Headers          struct {
		Request  []string `json:"request"`
		Response []string `json:"response"`
	} `json:"headers"`
	Host                   string `json:"host"`
	HTTP2StreamDependency  int64  `json:"http2_stream_dependency"`
	HTTP2StreamExclusive   int64  `json:"http2_stream_exclusive"`
	HTTP2StreamID          int64  `json:"http2_stream_id"`
	HTTP2StreamWeight      int64  `json:"http2_stream_weight"`
	ImageSave              int64  `json:"image_save"`
	ImageTotal             int64  `json:"image_total"`
	Index                  int64  `json:"index"`
	Initiator              string `json:"initiator"`
	InitiatorColumn        int64  `json:"initiator_column"`
	InitiatorDetail        string `json:"initiator_detail"`
	InitiatorFunction      string `json:"initiator_function"`
	InitiatorLine          int64  `json:"initiator_line"`
	InitiatorType          string `json:"initiator_type"`
	IPAddr                 string `json:"ip_addr"`
	IsSecure               int64  `json:"is_secure"`
	JpegScanCount          int64  `json:"jpeg_scan_count"`
	LoadEnd                int64  `json:"load_end"`
	LoadMs                 int64  `json:"load_ms"`
	LoadStart              int64  `json:"load_start"`
	Method                 string `json:"method"`
	MinifySave             int64  `json:"minify_save"`
	MinifyTotal            int64  `json:"minify_total"`
	Number                 int64  `json:"number"`
	ObjectSize             int64  `json:"objectSize"`
	ObjectSizeUncompressed int64  `json:"objectSizeUncompressed"`
	Priority               string `json:"priority"`
	Protocol               string `json:"protocol"`
	RequestID              int64  `json:"request_id"`
	ResponseCode           int64  `json:"responseCode"`
	ScoreCache             int64  `json:"score_cache"`
	ScoreCdn               int64  `json:"score_cdn"`
	ScoreCombine           int64  `json:"score_combine"`
	ScoreCompress          int64  `json:"score_compress"`
	ScoreCookies           int64  `json:"score_cookies"`
	ScoreEtags             int64  `json:"score_etags"`
	ScoreGzip              int64  `json:"score_gzip"`
	ScoreKeep_alive        int64  `json:"score_keep-alive"`
	ScoreMinify            int64  `json:"score_minify"`
	ScoreProgressiveJpeg   int64  `json:"score_progressive_jpeg"`
	ServerCount            int64  `json:"server_count"`
	ServerRtt              int64  `json:"server_rtt"`
	Socket                 int64  `json:"socket"`
	SslEnd                 int64  `json:"ssl_end"`
	SslMs                  int64  `json:"ssl_ms"`
	SslStart               int64  `json:"ssl_start"`
	TtfbEnd                int64  `json:"ttfb_end"`
	TtfbMs                 int64  `json:"ttfb_ms"`
	TtfbStart              int64  `json:"ttfb_start"`
	Type                   int64  `json:"type"`
	URL                    string `json:"url"`
	WasPushed              int64  `json:"was_pushed"`
}

type IterationTestTiming struct {
	AllRunsDuration       int64 `json:"AllRunsDuration"`
	ExtensionBlank        int64 `json:"ExtensionBlank"`
	ExtensionStart        int64 `json:"ExtensionStart"`
	MeasureStep           int64 `json:"MeasureStep"`
	ProcessRequests       int64 `json:"ProcessRequests"`
	ProcessTrace          int64 `json:"ProcessTrace"`
	ProcessVideo          int64 `json:"ProcessVideo"`
	RunOptimizationChecks int64 `json:"RunOptimizationChecks"`
	RunTest               int64 `json:"RunTest"`
	SaveResult            int64 `json:"SaveResult"`
	UploadImages          int64 `json:"UploadImages"`
	WaitForIdle           int64 `json:"WaitForIdle"`
}

type Iteration struct {
	SpeedIndex                    int64                         `json:"SpeedIndex"`
	TTFB                          int64                         `json:"TTFB"`
	URL                           string                        `json:"URL"`
	AdultSite                     int64                         `json:"adult_site"`
	Aft                           int64                         `json:"aft"`
	BasePageCDN                   string                        `json:"base_page_cdn"`
	BasePageRedirects             int64                         `json:"base_page_redirects"`
	BasePageTTFB                  int64                         `json:"base_page_ttfb"`
	BigImageCount                 int64                         `json:"bigImageCount"`
	Breakdown                     map[string]IterationBreakdown `json:"breakdown"`
	BrowserMainMemoryInKB         int64                         `json:"browser_main_memory_kb"`
	BrowserName                   string                        `json:"browser_name"`
	BrowserOtherPrivateMemoryInKB int64                         `json:"browser_other_private_memory_kb"`
	BrowserProcessCount           int64                         `json:"browser_process_count"`
	BrowserVersion                string                        `json:"browser_version"`
	BrowserWorkingSetInKB         int64                         `json:"browser_working_set_kb"`
	BytesIn                       int64                         `json:"bytesIn"`
	BytesInDoc                    int64                         `json:"bytesInDoc"`
	BytesOut                      int64                         `json:"bytesOut"`
	BytesOutDoc                   int64                         `json:"bytesOutDoc"`
	Cached                        int64                         `json:"cached"`
	CertificateBytes              int64                         `json:"certificate_bytes"`
	Connections                   int64                         `json:"connections"`
	ConsoleLog                    []IterationConsoleLog         `json:"consoleLog"`
	Date                          int64                         `json:"date"`
	DocCPUTime                    float64                       `json:"docCPUms"`
	DocCPUPercent                 int64                         `json:"docCPUpct"`
	DocTime                       int64                         `json:"docTime"`
	DomContentLoadedEventEnd      int64                         `json:"domContentLoadedEventEnd"`
	DomContentLoadedEventStart    int64                         `json:"domContentLoadedEventStart"`
	DomElements                   int64                         `json:"domElements"`
	DomInteractive                int64                         `json:"domInteractive"`
	DomLoading                    int64                         `json:"domLoading"`
	DomTime                       int64                         `json:"domTime"`
	Domains                       map[string]IterationDomain    `json:"domains"`
	EffectiveBps                  int64                         `json:"effectiveBps"`
	EffectiveBpsDoc               int64                         `json:"effectiveBpsDoc"`
	EventName                     string                        `json:"eventName"`
	FirstPaint                    int64                         `json:"firstPaint"`
	FixedViewport                 int64                         `json:"fixed_viewport"`
	FullyLoaded                   int64                         `json:"fullyLoaded"`
	FullyLoadedCPUTime            float64                       `json:"fullyLoadedCPUms"`
	FullyLoadedCPUPercent         int64                         `json:"fullyLoadedCPUpct"`
	GzipSavings                   int64                         `json:"gzip_savings"`
	GzipTotal                     int64                         `json:"gzip_total"`
	ImageSavings                  int64                         `json:"image_savings"`
	ImageTotal                    int64                         `json:"image_total"`
	Images                        IterationImages               `json:"images"`
	IsResponsive                  int                           `json:"isResponsive"`
	LastVisualChange              int64                         `json:"lastVisualChange"`
	LoadEventEnd                  int64                         `json:"loadEventEnd"`
	LoadEventStart                int64                         `json:"loadEventStart"`
	LoadTime                      int64                         `json:"loadTime"`
	MaybeCaptcha                  int                           `json:"maybeCaptcha"`
	MinifySavings                 int64                         `json:"minify_savings"`
	MinifyTotal                   int64                         `json:"minify_total"`
	NumSteps                      int64                         `json:"numSteps"`
	OptimizationChecked           int64                         `json:"optimization_checked"`
	PageSpeedVersion              string                        `json:"pageSpeedVersion"`
	Pages                         IterationPages                `json:"pages"`
	RawData                       IterationRawData              `json:"rawData"`
	Render                        int64                         `json:"render"`
	Requests                      []IterationRequest            `json:"requests"`
	RequestsDoc                   int64                         `json:"requestsDoc"`
	RequestsFull                  int64                         `json:"requestsFull"`
	Responses200                  int64                         `json:"responses_200"`
	Responses404                  int64                         `json:"responses_404"`
	ResponsesOther                int64                         `json:"responses_other"`
	Result                        int64                         `json:"result"`
	Run                           int64                         `json:"run"`
	ScoreCache                    int64                         `json:"score_cache"`
	ScoreCDN                      int64                         `json:"score_cdn"`
	ScoreCombine                  int64                         `json:"score_combine"`
	ScoreCompress                 int64                         `json:"score_compress"`
	ScoreCookies                  int64                         `json:"score_cookies"`
	ScoreEtags                    int64                         `json:"score_etags"`
	ScoreGzip                     int64                         `json:"score_gzip"`
	ScoreKeepAlive                int64                         `json:"score_keep-alive"`
	ScoreMinify                   int64                         `json:"score_minify"`
	ScoreProgressiveJpeg          int64                         `json:"score_progressive_jpeg"`
	ServerCount                   int64                         `json:"server_count"`
	ServerRTT                     int64                         `json:"server_rtt"`
	SmallImageCount               int64                         `json:"smallImageCount"`
	Step                          int64                         `json:"step"`
	TestTiming                    IterationTestTiming           `json:"testTiming"`
	Tester                        string                        `json:"tester"`
	Thumbnails                    IterationThumbnails           `json:"thumbnails"`
	Title                         string                        `json:"title"`
	TitleTime                     int64                         `json:"titleTime"`
	VideoFrames                   []IterationVideoFrame         `json:"videoFrames"`
	VisualComplete                int64                         `json:"visualComplete"`
	VisualComplete85              int64                         `json:"visualComplete85"`
	VisualComplete90              int64                         `json:"visualComplete90"`
	VisualComplete95              int64                         `json:"visualComplete95"`
	VisualComplete99              int64                         `json:"visualComplete99"`
}

type TestRun struct {
	FirstView  Iteration `json:"firstView,omitempty"`
	RepeatView Iteration `json:"repeatView,omitempty"`
}

type WPTResults struct {
	Data struct {
		Average           interface{}        `json:"average"`
		BwDown            int64              `json:"bwDown"`
		BwUp              int64              `json:"bwUp"`
		Completed         int64              `json:"completed"`
		Connectivity      string             `json:"connectivity"`
		From              string             `json:"from"`
		FVOnly            bool               `json:"fvonly"`
		ID                string             `json:"id"`
		Label             string             `json:"label"`
		Latency           int64              `json:"latency"`
		Location          string             `json:"location"`
		Median            TestRun            `json:"median"`
		Mobile            int64              `json:"mobile"`
		PackLossRate      string             `json:"plr"`
		Runs              map[string]TestRun `json:"runs"`
		StandardDeviation interface{}        `json:"standardDeviation"`
		SuccessfulFVRuns  int64              `json:"successfulFVRuns"`
		SuccessfulRVRuns  int64              `json:"successfulRVRuns"`
		Summary           string             `json:"summary"`
		TestURL           string             `json:"testUrl"`
		Tester            string             `json:"tester"`
		TesterDNS         string             `json:"testerDNS"`
		URL               string             `json:"url"`
	} `json:"data"`
	StatusCode         int    `json:"statusCode"`
	StatusText         string `json:"statusText"`
	WebPagetestVersion string `json:"webPagetestVersion"`
}

type WPTHistory struct {
	Items []HistoryItem
}

type HistoryItem struct {
	DataTime string `csv:"Date/Time"`
	Location string `csv:"Location"`
	TestID   string `csv:"Test ID"`
	URL      string `csv:"URL"`
	Label    string `csv:"Label"`
}

type TestOptions struct {
	URL            string   `url:"url"`
	Label          string   `url:"label,omitempty"`
	Location       string   `url:"-"`
	Connectivity   string   `url:"-"`
	LocationString string   `url:"location,omitempty"`
	Block          []string `url:"-"`
	BlockString    string   `url:"block,omitempty"`
	Login          string   `url:"login,omitempty"`
	Password       string   `url:"password,omitempty"`
	Notify         string   `url:"notify,omitempty"`
	Pingback       string   `url:"pingback,omitempty"`
	CMDLine        string   `url:"cmdline,omitempty"`
	TSViewID       string   `url:"tsview_id,omitempty"`
	Custom         string   `url:"custom,omitempty"`
	Tester         string   `url:"tester,omitempty"`
	Affinity       string   `url:"affinity,omitempty"`
	DomElement     string   `url:"domelement,omitempty"`
	Script         string   `url:"script,omitempty"`
	TimelineStack  int      `url:"timelineStack,omitempty"`
	Runs           int      `url:"runs,omitempty"`
	Connections    int      `url:"connections,omitempty"`
	Authtype       int      `url:"authType,omitempty"`
	BWDown         int      `url:"bwDown,omitempty"`
	BWUp           int      `url:"bwUp,omitempty"`
	Latency        int      `url:"latency,omitempty"`
	PackLossRate   int      `url:"plr,omitempty"`
	IQ             int      `url:"iq,omitempty"`
	FVOnly         int      `url:"fvonly,omitempty"`
	Private        int      `url:"private,omitempty"`
	StopOnLoad     int      `url:"web10,omitempty"`
	Video          int      `url:"video,omitempty"`
	TCPDump        int      `url:"tcpdump,omitempty"`
	NoOpt          int      `url:"noopt,omitempty"`
	NoImages       int      `url:"noimages,omitempty"`
	NoHeaders      int      `url:"noheaders,omitempty"`
	PNGSS          int      `url:"pngss,omitempty"`
	NoScript       int      `url:"noscript,omitempty"`
	ClearCerts     int      `url:"clearcerts,omitempty"`
	Mobile         int      `url:"mobile,omitempty"`
	MV             int      `url:"mv,omitempty"`
	HTMLBody       int      `url:"htmlbody,omitempty"`
	Timeline       int      `url:"timeline,omitempty"`
	IgnoreSSL      int      `url:"ignoreSSL,omitempty"`
	Lighthouse     int      `url:"lighthouse,omitempty"`
	MedianMetric   string   `url:"medianMetric,omitempty"`
}

type TestResponse struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Data       struct {
		TestId     string `json:"testId"`
		OwnerKey   string `json:"ownerKey"`
		JSONUrl    string `json:"jsonUrl"`
		XMLUrl     string `json:"xmlUrl"`
		UserUrl    string `json:"userUrl"`
		SummaryCSV string `json:"summaryCSV"`
		DetailCSV  string `json:"detailCSV"`
	} `json:"data"`
}

type TestStatus struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Id         string `json:"id"`
	Data       struct {
		StatusCode int         `json:"statusCode"`
		StatusText string      `json:"statusText"`
		Id         string      `json:"id"`
		TestInfo   TestOptions `json:"testInfo"`
	} `json:"data"`
	TestId          string `json:"testId"`
	Runs            int    `json:"runs"`
	FVOnly          int    `json:"fvonly"`
	Remote          bool   `json:"remote"`
	TestsExpected   int    `json:"testsExpected"`
	Location        string `json:"location"`
	Elapsed         int    `json:"elapsed"`
	BehindCount     int    `json:"behindCount"`
	FVRunsCompleted int    `json:"fvRunsCompleted"`
	RVRunsCompleted int    `json:"rvRunsCompleted"`
}
