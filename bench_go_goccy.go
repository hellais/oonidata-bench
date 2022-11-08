package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/goccy/go-json"
)

type BaseMeasurement struct {
	annotations            map[string]string      `json:"annotations"`                     // dict[str, str]
	input                  interface{}            `json:"input,omitempty"`                 // typing.Union[str, typing.List[str], NoneType]
	report_id              string                 `json:"report_id"`                       // <class 'str'>
	measurement_start_time string                 `json:"measurement_start_time"`          // <class 'str'>
	test_start_time        string                 `json:"test_start_time"`                 // <class 'str'>
	probe_asn              string                 `json:"probe_asn"`                       // <class 'str'>
	probe_network_name     string                 `json:"probe_network_name,omitempty"`    // typing.Optional[str]
	probe_cc               string                 `json:"probe_cc"`                        // <class 'str'>
	probe_ip               string                 `json:"probe_ip,omitempty"`              // typing.Optional[str]
	resolver_asn           string                 `json:"resolver_asn,omitempty"`          // typing.Optional[str]
	resolver_ip            string                 `json:"resolver_ip,omitempty"`           // typing.Optional[str]
	resolver_network_name  string                 `json:"resolver_network_name,omitempty"` // typing.Optional[str]
	test_name              string                 `json:"test_name"`                       // <class 'str'>
	test_version           string                 `json:"test_version"`                    // <class 'str'>
	test_runtime           float32                `json:"test_runtime"`                    // <class 'float'>
	software_name          string                 `json:"software_name"`                   // <class 'str'>
	software_version       string                 `json:"software_version"`                // <class 'str'>
	test_helpers           map[string]interface{} `json:"test_helpers,omitempty"`          // typing.Optional[dict]
	test_keys              BaseTestKeys           `json:"test_keys"`                       // <class 'oonidata.dataformat.BaseTestKeys'>
	data_format_version    string                 `json:"data_format_version,omitempty"`   // typing.Optional[str]
	measurement_uid        string                 `json:"measurement_uid,omitempty"`       // typing.Optional[str]
}

type BaseTestKeys struct {
	client_resolver string `json:"client_resolver,omitempty"` // typing.Optional[str]
}
type BinaryData struct {
	format string `json:"format"` // <class 'str'>
	data   string `json:"data"`   // <class 'str'>
}
type DNSAnswer struct {
	answer_type      string `json:"answer_type"`                // <class 'str'>
	asn              int    `json:"asn,omitempty"`              // typing.Optional[int]
	as_org_name      string `json:"as_org_name,omitempty"`      // typing.Optional[str]
	expiration_limit string `json:"expiration_limit,omitempty"` // typing.Optional[str]
	hostname         string `json:"hostname,omitempty"`         // typing.Optional[str]
	ipv4             string `json:"ipv4,omitempty"`             // typing.Optional[str]
	ipv6             string `json:"ipv6,omitempty"`             // typing.Optional[str]
	minimum_ttl      string `json:"minimum_ttl,omitempty"`      // typing.Optional[str]
	refresh_interval string `json:"refresh_interval,omitempty"` // typing.Optional[str]
	responsible_name string `json:"responsible_name,omitempty"` // typing.Optional[str]
	retry_interval   string `json:"retry_interval,omitempty"`   // typing.Optional[str]
	serial_number    string `json:"serial_number,omitempty"`    // typing.Optional[str]
	ttl              int    `json:"ttl,omitempty"`              // typing.Optional[int]
}
type DNSCheck struct {
	annotations            map[string]string      `json:"annotations"`                     // dict[str, str]
	input                  interface{}            `json:"input,omitempty"`                 // typing.Union[str, typing.List[str], NoneType]
	report_id              string                 `json:"report_id"`                       // <class 'str'>
	measurement_start_time string                 `json:"measurement_start_time"`          // <class 'str'>
	test_start_time        string                 `json:"test_start_time"`                 // <class 'str'>
	probe_asn              string                 `json:"probe_asn"`                       // <class 'str'>
	probe_network_name     string                 `json:"probe_network_name,omitempty"`    // typing.Optional[str]
	probe_cc               string                 `json:"probe_cc"`                        // <class 'str'>
	probe_ip               string                 `json:"probe_ip,omitempty"`              // typing.Optional[str]
	resolver_asn           string                 `json:"resolver_asn,omitempty"`          // typing.Optional[str]
	resolver_ip            string                 `json:"resolver_ip,omitempty"`           // typing.Optional[str]
	resolver_network_name  string                 `json:"resolver_network_name,omitempty"` // typing.Optional[str]
	test_name              string                 `json:"test_name"`                       // <class 'str'>
	test_version           string                 `json:"test_version"`                    // <class 'str'>
	test_runtime           float32                `json:"test_runtime"`                    // <class 'float'>
	software_name          string                 `json:"software_name"`                   // <class 'str'>
	software_version       string                 `json:"software_version"`                // <class 'str'>
	test_helpers           map[string]interface{} `json:"test_helpers,omitempty"`          // typing.Optional[dict]
	test_keys              DNSCheckTestKeys       `json:"test_keys"`                       // <class 'oonidata.dataformat.DNSCheckTestKeys'>
	data_format_version    string                 `json:"data_format_version,omitempty"`   // typing.Optional[str]
	measurement_uid        string                 `json:"measurement_uid,omitempty"`       // typing.Optional[str]
}
type DNSCheckTestKeys struct {
	client_resolver   string                       `json:"client_resolver,omitempty"`   // typing.Optional[str]
	bootstrap         URLGetterTestKeys            `json:"bootstrap,omitempty"`         // typing.Optional[oonidata.dataformat.URLGetterTestKeys]
	bootstrap_failure string                       `json:"bootstrap_failure,omitempty"` // typing.Optional[str]
	lookups           map[string]URLGetterTestKeys `json:"lookups"`                     // dict[str, oonidata.dataformat.URLGetterTestKeys]
}
type DNSQuery struct {
	failure          string      `json:"failure,omitempty"`          // typing.Optional[str]
	hostname         string      `json:"hostname"`                   // <class 'str'>
	query_type       string      `json:"query_type"`                 // <class 'str'>
	dial_id          int         `json:"dial_id,omitempty"`          // typing.Optional[int]
	engine           string      `json:"engine,omitempty"`           // typing.Optional[str]
	resolver_address string      `json:"resolver_address,omitempty"` // typing.Optional[str]
	t                float32     `json:"t,omitempty"`                // typing.Optional[float]
	transaction_id   int         `json:"transaction_id,omitempty"`   // typing.Optional[int]
	answers          []DNSAnswer `json:"answers,omitempty"`          // typing.Optional[typing.List[oonidata.dataformat.DNSAnswer]]
}
type HTTPBase struct {
	body               interface{}     `json:"body,omitempty"`               // typing.Union[str, oonidata.dataformat.BinaryData, NoneType]
	body_is_truncated  bool            `json:"body_is_truncated,omitempty"`  // typing.Optional[bool]
	headers            interface{}     `json:"headers,omitempty"`            // typing.Optional[dict[str, typing.Union[str, oonidata.dataformat.BinaryData, NoneType]]]
	headers_list       [][]interface{} `json:"headers_list,omitempty"`       // typing.Optional[typing.List[typing.List[typing.Union[str, oonidata.dataformat.BinaryData, NoneType]]]]
	headers_list_bytes []interface{}   `json:"headers_list_bytes,omitempty"` // typing.Optional[typing.List[typing.Tuple[str, bytes]]]
	body_bytes         []byte          `json:"body_bytes,omitempty"`         // typing.Optional[bytes]
}
type HTTPRequest struct {
	body               interface{}     `json:"body,omitempty"`               // typing.Union[str, oonidata.dataformat.BinaryData, NoneType]
	body_is_truncated  bool            `json:"body_is_truncated,omitempty"`  // typing.Optional[bool]
	headers            interface{}     `json:"headers,omitempty"`            // typing.Optional[dict[str, typing.Union[str, oonidata.dataformat.BinaryData, NoneType]]]
	headers_list       [][]interface{} `json:"headers_list,omitempty"`       // typing.Optional[typing.List[typing.List[typing.Union[str, oonidata.dataformat.BinaryData, NoneType]]]]
	headers_list_bytes []interface{}   `json:"headers_list_bytes,omitempty"` // typing.Optional[typing.List[typing.Tuple[str, bytes]]]
	body_bytes         []byte          `json:"body_bytes,omitempty"`         // typing.Optional[bytes]
	url                string          `json:"url"`                          // <class 'str'>
	method             string          `json:"method,omitempty"`             // typing.Optional[str]
	tor                TorInfo         `json:"tor,omitempty"`                // typing.Optional[oonidata.dataformat.TorInfo]
	x_transport        string          `json:"x_transport,omitempty"`        // typing.Optional[str]
}
type HTTPResponse struct {
	body               interface{}     `json:"body,omitempty"`               // typing.Union[str, oonidata.dataformat.BinaryData, NoneType]
	body_is_truncated  bool            `json:"body_is_truncated,omitempty"`  // typing.Optional[bool]
	headers            interface{}     `json:"headers,omitempty"`            // typing.Optional[dict[str, typing.Union[str, oonidata.dataformat.BinaryData, NoneType]]]
	headers_list       [][]interface{} `json:"headers_list,omitempty"`       // typing.Optional[typing.List[typing.List[typing.Union[str, oonidata.dataformat.BinaryData, NoneType]]]]
	headers_list_bytes []interface{}   `json:"headers_list_bytes,omitempty"` // typing.Optional[typing.List[typing.Tuple[str, bytes]]]
	body_bytes         []byte          `json:"body_bytes,omitempty"`         // typing.Optional[bytes]
	code               int             `json:"code,omitempty"`               // typing.Optional[int]
}
type HTTPTransaction struct {
	failure        string       `json:"failure,omitempty"`        // typing.Optional[str]
	request        HTTPRequest  `json:"request,omitempty"`        // typing.Optional[oonidata.dataformat.HTTPRequest]
	response       HTTPResponse `json:"response,omitempty"`       // typing.Optional[oonidata.dataformat.HTTPResponse]
	t              float32      `json:"t,omitempty"`              // typing.Optional[float]
	transaction_id int          `json:"transaction_id,omitempty"` // typing.Optional[int]
}
type NetworkEvent struct {
	failure        string   `json:"failure,omitempty"`        // typing.Optional[str]
	operation      string   `json:"operation"`                // <class 'str'>
	t              float32  `json:"t"`                        // <class 'float'>
	address        string   `json:"address,omitempty"`        // typing.Optional[str]
	dial_id        int      `json:"dial_id,omitempty"`        // typing.Optional[int]
	num_bytes      int      `json:"num_bytes,omitempty"`      // typing.Optional[int]
	proto          string   `json:"proto,omitempty"`          // typing.Optional[str]
	tags           []string `json:"tags,omitempty"`           // typing.Optional[typing.List[str]]
	transaction_id string   `json:"transaction_id,omitempty"` // typing.Optional[str]
	conn_id        int      `json:"conn_id,omitempty"`        // typing.Optional[int]
}
type TCPConnect struct {
	ip     string           `json:"ip"`          // <class 'str'>
	port   int              `json:"port"`        // <class 'int'>
	status TCPConnectStatus `json:"status"`      // <class 'oonidata.dataformat.TCPConnectStatus'>
	t      float32          `json:"t,omitempty"` // typing.Optional[float]
}
type TCPConnectStatus struct {
	blocked bool   `json:"blocked,omitempty"` // typing.Optional[bool]
	success bool   `json:"success"`           // <class 'bool'>
	failure string `json:"failure,omitempty"` // typing.Optional[str]
}
type TLSHandshake struct {
	failure             string       `json:"failure,omitempty"`             // typing.Optional[str]
	peer_certificates   []BinaryData `json:"peer_certificates,omitempty"`   // typing.Optional[typing.List[oonidata.dataformat.BinaryData]]
	address             string       `json:"address,omitempty"`             // typing.Optional[str]
	cipher_suite        string       `json:"cipher_suite,omitempty"`        // typing.Optional[str]
	negotiated_protocol string       `json:"negotiated_protocol,omitempty"` // typing.Optional[str]
	no_tls_verify       bool         `json:"no_tls_verify,omitempty"`       // typing.Optional[bool]
	server_name         string       `json:"server_name,omitempty"`         // typing.Optional[str]
	t                   float32      `json:"t,omitempty"`                   // typing.Optional[float]
	tags                []string     `json:"tags,omitempty"`                // typing.Optional[typing.List[str]]
	tls_version         string       `json:"tls_version,omitempty"`         // typing.Optional[str]
	transaction_id      int          `json:"transaction_id,omitempty"`      // typing.Optional[int]
}
type Tor struct {
	annotations            map[string]string      `json:"annotations"`                     // dict[str, str]
	input                  interface{}            `json:"input,omitempty"`                 // typing.Union[str, typing.List[str], NoneType]
	report_id              string                 `json:"report_id"`                       // <class 'str'>
	measurement_start_time string                 `json:"measurement_start_time"`          // <class 'str'>
	test_start_time        string                 `json:"test_start_time"`                 // <class 'str'>
	probe_asn              string                 `json:"probe_asn"`                       // <class 'str'>
	probe_network_name     string                 `json:"probe_network_name,omitempty"`    // typing.Optional[str]
	probe_cc               string                 `json:"probe_cc"`                        // <class 'str'>
	probe_ip               string                 `json:"probe_ip,omitempty"`              // typing.Optional[str]
	resolver_asn           string                 `json:"resolver_asn,omitempty"`          // typing.Optional[str]
	resolver_ip            string                 `json:"resolver_ip,omitempty"`           // typing.Optional[str]
	resolver_network_name  string                 `json:"resolver_network_name,omitempty"` // typing.Optional[str]
	test_name              string                 `json:"test_name"`                       // <class 'str'>
	test_version           string                 `json:"test_version"`                    // <class 'str'>
	test_runtime           float32                `json:"test_runtime"`                    // <class 'float'>
	software_name          string                 `json:"software_name"`                   // <class 'str'>
	software_version       string                 `json:"software_version"`                // <class 'str'>
	test_helpers           map[string]interface{} `json:"test_helpers,omitempty"`          // typing.Optional[dict]
	test_keys              TorTestKeys            `json:"test_keys"`                       // <class 'oonidata.dataformat.TorTestKeys'>
	data_format_version    string                 `json:"data_format_version,omitempty"`   // typing.Optional[str]
	measurement_uid        string                 `json:"measurement_uid,omitempty"`       // typing.Optional[str]
}
type TorInfo struct {
	is_tor    bool   `json:"is_tor"`              // <class 'bool'>
	exit_ip   string `json:"exit_ip,omitempty"`   // typing.Optional[str]
	exit_name string `json:"exit_name,omitempty"` // typing.Optional[str]
}
type TorTestKeys struct {
	targets map[string]TorTestTarget `json:"targets"` // dict[str, oonidata.dataformat.TorTestTarget]
}
type TorTestTarget struct {
	failure         string            `json:"failure,omitempty"`        // typing.Optional[str]
	network_events  []NetworkEvent    `json:"network_events,omitempty"` // typing.Optional[typing.List[oonidata.dataformat.NetworkEvent]]
	queries         []DNSQuery        `json:"queries,omitempty"`        // typing.Optional[typing.List[oonidata.dataformat.DNSQuery]]
	requests        []HTTPTransaction `json:"requests,omitempty"`       // typing.Optional[typing.List[oonidata.dataformat.HTTPTransaction]]
	tls_handshakes  []TLSHandshake    `json:"tls_handshakes,omitempty"` // typing.Optional[typing.List[oonidata.dataformat.TLSHandshake]]
	tcp_connect     []TCPConnect      `json:"tcp_connect,omitempty"`    // typing.Optional[typing.List[oonidata.dataformat.TCPConnect]]
	target_address  string            `json:"target_address"`           // <class 'str'>
	target_name     string            `json:"target_name,omitempty"`    // typing.Optional[str]
	target_protocol string            `json:"target_protocol"`          // <class 'str'>
}
type URLGetterTestKeys struct {
	client_resolver string            `json:"client_resolver,omitempty"` // typing.Optional[str]
	failure         string            `json:"failure,omitempty"`         // typing.Optional[str]
	socksproxy      string            `json:"socksproxy,omitempty"`      // typing.Optional[str]
	tls_handshakes  []TLSHandshake    `json:"tls_handshakes,omitempty"`  // typing.Optional[typing.List[oonidata.dataformat.TLSHandshake]]
	network_events  []NetworkEvent    `json:"network_events,omitempty"`  // typing.Optional[typing.List[oonidata.dataformat.NetworkEvent]]
	queries         []DNSQuery        `json:"queries,omitempty"`         // typing.Optional[typing.List[oonidata.dataformat.DNSQuery]]
	tcp_connect     []TCPConnect      `json:"tcp_connect,omitempty"`     // typing.Optional[typing.List[oonidata.dataformat.TCPConnect]]
	requests        []HTTPTransaction `json:"requests,omitempty"`        // typing.Optional[typing.List[oonidata.dataformat.HTTPTransaction]]
}
type WebConnectivity struct {
	Annotations            map[string]string       `json:"annotations"`                     // dict[str, str]
	input                  interface{}             `json:"input,omitempty"`                 // typing.Union[str, typing.List[str], NoneType]
	report_id              string                  `json:"report_id"`                       // <class 'str'>
	measurement_start_time string                  `json:"measurement_start_time"`          // <class 'str'>
	test_start_time        string                  `json:"test_start_time"`                 // <class 'str'>
	probe_asn              string                  `json:"probe_asn"`                       // <class 'str'>
	probe_network_name     string                  `json:"probe_network_name,omitempty"`    // typing.Optional[str]
	probe_cc               string                  `json:"probe_cc"`                        // <class 'str'>
	probe_ip               string                  `json:"probe_ip,omitempty"`              // typing.Optional[str]
	resolver_asn           string                  `json:"resolver_asn,omitempty"`          // typing.Optional[str]
	resolver_ip            string                  `json:"resolver_ip,omitempty"`           // typing.Optional[str]
	resolver_network_name  string                  `json:"resolver_network_name,omitempty"` // typing.Optional[str]
	test_name              string                  `json:"test_name"`                       // <class 'str'>
	test_version           string                  `json:"test_version"`                    // <class 'str'>
	test_runtime           float32                 `json:"test_runtime"`                    // <class 'float'>
	SoftwareName           string                  `json:"software_name"`                   // <class 'str'>
	software_version       string                  `json:"software_version"`                // <class 'str'>
	test_helpers           map[string]interface{}  `json:"test_helpers,omitempty"`          // typing.Optional[dict]
	test_keys              WebConnectivityTestKeys `json:"test_keys"`                       // <class 'oonidata.dataformat.WebConnectivityTestKeys'>
	data_format_version    string                  `json:"data_format_version,omitempty"`   // typing.Optional[str]
	measurement_uid        string                  `json:"measurement_uid,omitempty"`       // typing.Optional[str]
}
type WebConnectivityControl struct {
	tcp_connect  interface{}                       `json:"tcp_connect,omitempty"`  // typing.Optional[dict[str, oonidata.dataformat.WebConnectivityControlTCPConnectStatus]]
	http_request WebConnectivityControlHTTPRequest `json:"http_request,omitempty"` // typing.Optional[oonidata.dataformat.WebConnectivityControlHTTPRequest]
	dns          WebConnectivityControlDNS         `json:"dns,omitempty"`          // typing.Optional[oonidata.dataformat.WebConnectivityControlDNS]
}
type WebConnectivityControlDNS struct {
	failure string   `json:"failure,omitempty"` // typing.Optional[str]
	addrs   []string `json:"addrs,omitempty"`   // typing.Optional[typing.List[str]]
}
type WebConnectivityControlHTTPRequest struct {
	body_length int               `json:"body_length,omitempty"` // typing.Optional[int]
	failure     string            `json:"failure,omitempty"`     // typing.Optional[str]
	title       string            `json:"title,omitempty"`       // typing.Optional[str]
	headers     map[string]string `json:"headers,omitempty"`     // typing.Optional[dict[str, str]]
	status_code int               `json:"status_code,omitempty"` // typing.Optional[int]
}
type WebConnectivityControlTCPConnectStatus struct {
	status  bool   `json:"status,omitempty"`  // typing.Optional[bool]
	failure string `json:"failure,omitempty"` // typing.Optional[str]
}
type WebConnectivityTestKeys struct {
	dns_experiment_failure  string                 `json:"dns_experiment_failure,omitempty"`  // typing.Optional[str]
	control_failure         string                 `json:"control_failure,omitempty"`         // typing.Optional[str]
	http_experiment_failure string                 `json:"http_experiment_failure,omitempty"` // typing.Optional[str]
	dns_consistency         string                 `json:"dns_consistency,omitempty"`         // typing.Optional[str]
	body_length_match       bool                   `json:"body_length_match,omitempty"`       // typing.Optional[bool]
	body_proportion         float32                `json:"body_proportion,omitempty"`         // typing.Optional[float]
	status_code_match       bool                   `json:"status_code_match,omitempty"`       // typing.Optional[bool]
	headers_match           bool                   `json:"headers_match,omitempty"`           // typing.Optional[bool]
	title_match             bool                   `json:"title_match,omitempty"`             // typing.Optional[bool]
	accessible              bool                   `json:"accessible,omitempty"`              // typing.Optional[bool]
	blocking                interface{}            `json:"blocking,omitempty"`                // typing.Union[str, bool, NoneType]
	control                 WebConnectivityControl `json:"control,omitempty"`                 // typing.Optional[oonidata.dataformat.WebConnectivityControl]
	tls_handshakes          []TLSHandshake         `json:"tls_handshakes,omitempty"`          // typing.Optional[typing.List[oonidata.dataformat.TLSHandshake]]
	network_events          []NetworkEvent         `json:"network_events,omitempty"`          // typing.Optional[typing.List[oonidata.dataformat.NetworkEvent]]
	queries                 []DNSQuery             `json:"queries,omitempty"`                 // typing.Optional[typing.List[oonidata.dataformat.DNSQuery]]
	tcp_connect             []TCPConnect           `json:"tcp_connect,omitempty"`             // typing.Optional[typing.List[oonidata.dataformat.TCPConnect]]
	requests                []HTTPTransaction      `json:"requests,omitempty"`                // typing.Optional[typing.List[oonidata.dataformat.HTTPTransaction]]
	x_status                int                    `json:"x_status,omitempty"`                // typing.Optional[int]
	x_dns_runtime           int                    `json:"x_dns_runtime,omitempty"`           // typing.Optional[int]
	x_th_runtime            int                    `json:"x_th_runtime,omitempty"`            // typing.Optional[int]
	x_tcptls_runtime        int                    `json:"x_tcptls_runtime,omitempty"`        // typing.Optional[int]
	x_http_runtime          int                    `json:"x_http_runtime,omitempty"`          // typing.Optional[int]
	client_resolver         string                 `json:"client_resolver,omitempty"`         // typing.Optional[str]
	agent                   string                 `json:"agent,omitempty"`                   // typing.Optional[str]
	retries                 int                    `json:"retries,omitempty"`                 // typing.Optional[int]
	socksproxy              string                 `json:"socksproxy,omitempty"`              // typing.Optional[str]
}

func main() {
	file, err := os.Open("sample-file.jsonl")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines = make([]string, 0)
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	start := time.Now()

	iters := 0
	for i := 1; i <= 10; i++ {
		for _, line := range lines {
			data := WebConnectivity{}
			json.Unmarshal([]byte(line), &data)
			iters += 1
		}

	}
	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println("# go_goccy_json")
	fmt.Printf("  runtime: %f\n", elapsed.Seconds())
	fmt.Printf("  iters/s: %f\n", float64(iters)/elapsed.Seconds())
	fmt.Printf("  iters: %d\n", iters)
	fmt.Printf("  errs: 0\n")
}
