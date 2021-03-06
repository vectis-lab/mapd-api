// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "mapd"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  TSessionId connect(string user, string passwd, string dbname)")
  fmt.Fprintln(os.Stderr, "  void disconnect(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  TServerStatus get_server_status(TSessionId session)")
  fmt.Fprintln(os.Stderr, "   get_tables(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  TTableDetails get_table_details(TSessionId session, string table_name)")
  fmt.Fprintln(os.Stderr, "   get_users(TSessionId session)")
  fmt.Fprintln(os.Stderr, "   get_databases(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  string get_version()")
  fmt.Fprintln(os.Stderr, "  void start_heap_profile(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  void stop_heap_profile(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  string get_heap_profile(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  string get_memory_gpu(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  TMemorySummary get_memory_summary(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  void clear_cpu_memory(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  void clear_gpu_memory(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  TQueryResult sql_execute(TSessionId session, string query, bool column_format, string nonce, i32 first_n)")
  fmt.Fprintln(os.Stderr, "  TGpuDataFrame sql_execute_df(TSessionId session, string query, i32 first_n)")
  fmt.Fprintln(os.Stderr, "  TGpuDataFrame sql_execute_gpudf(TSessionId session, string query, i32 device_id, i32 first_n)")
  fmt.Fprintln(os.Stderr, "  void interrupt(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  TTableDescriptor sql_validate(TSessionId session, string query)")
  fmt.Fprintln(os.Stderr, "  void set_execution_mode(TSessionId session, TExecuteMode mode)")
  fmt.Fprintln(os.Stderr, "  TRenderResult render_vega(TSessionId session, i64 widget_id, string vega_json, i32 compression_level, string nonce)")
  fmt.Fprintln(os.Stderr, "  TPixelTableRowResult get_result_row_for_pixel(TSessionId session, i64 widget_id, TPixel pixel,  table_col_names, bool column_format, i32 pixelRadius, string nonce)")
  fmt.Fprintln(os.Stderr, "  TFrontendView get_frontend_view(TSessionId session, string view_name)")
  fmt.Fprintln(os.Stderr, "   get_frontend_views(TSessionId session)")
  fmt.Fprintln(os.Stderr, "  void create_frontend_view(TSessionId session, string view_name, string view_state, string image_hash, string view_metadata)")
  fmt.Fprintln(os.Stderr, "  void delete_frontend_view(TSessionId session, string view_name)")
  fmt.Fprintln(os.Stderr, "  TFrontendView get_link_view(TSessionId session, string link)")
  fmt.Fprintln(os.Stderr, "  string create_link(TSessionId session, string view_state, string view_metadata)")
  fmt.Fprintln(os.Stderr, "  void load_table_binary(TSessionId session, string table_name,  rows)")
  fmt.Fprintln(os.Stderr, "  void load_table(TSessionId session, string table_name,  rows)")
  fmt.Fprintln(os.Stderr, "  TDetectResult detect_column_types(TSessionId session, string file_name, TCopyParams copy_params)")
  fmt.Fprintln(os.Stderr, "  void create_table(TSessionId session, string table_name, TRowDescriptor row_desc, TTableType table_type)")
  fmt.Fprintln(os.Stderr, "  void import_table(TSessionId session, string table_name, string file_name, TCopyParams copy_params)")
  fmt.Fprintln(os.Stderr, "  void import_geo_table(TSessionId session, string table_name, string file_name, TCopyParams copy_params, TRowDescriptor row_desc)")
  fmt.Fprintln(os.Stderr, "  TImportStatus import_table_status(TSessionId session, string import_id)")
  fmt.Fprintln(os.Stderr, "  TPendingQuery start_query(TSessionId session, string query_ra, bool just_explain)")
  fmt.Fprintln(os.Stderr, "  TStepResult execute_first_step(TPendingQuery pending_query)")
  fmt.Fprintln(os.Stderr, "  void broadcast_serialized_rows(string serialized_rows, TRowDescriptor row_desc, TQueryId query_id)")
  fmt.Fprintln(os.Stderr, "  TRawPixelDataResult render_vega_raw_pixels(TSessionId session, i64 widget_id, i16 node_idx, string vega_json)")
  fmt.Fprintln(os.Stderr, "  void insert_data(TSessionId session, TInsertData insert_data)")
  fmt.Fprintln(os.Stderr, "  TTableDescriptor get_table_descriptor(TSessionId session, string table_name)")
  fmt.Fprintln(os.Stderr, "  TRowDescriptor get_row_descriptor(TSessionId session, string table_name)")
  fmt.Fprintln(os.Stderr, "  TRenderResult render(TSessionId session, string query, string render_type, string nonce)")
  fmt.Fprintln(os.Stderr, "  TPixelResult get_rows_for_pixels(TSessionId session, i64 widget_id,  pixels, string table_name,  col_names, bool column_format, string nonce)")
  fmt.Fprintln(os.Stderr, "  TPixelRowResult get_row_for_pixel(TSessionId session, i64 widget_id, TPixel pixel, string table_name,  col_names, bool column_format, i32 pixelRadius, string nonce)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := mapd.NewMapDClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "connect":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "Connect requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.Connect(value0, value1, value2))
    fmt.Print("\n")
    break
  case "disconnect":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Disconnect requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.Disconnect(value0))
    fmt.Print("\n")
    break
  case "get_server_status":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetServerStatus requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.GetServerStatus(value0))
    fmt.Print("\n")
    break
  case "get_tables":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetTables requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.GetTables(value0))
    fmt.Print("\n")
    break
  case "get_table_details":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetTableDetails requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.GetTableDetails(value0, value1))
    fmt.Print("\n")
    break
  case "get_users":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetUsers requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.GetUsers(value0))
    fmt.Print("\n")
    break
  case "get_databases":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetDatabases requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.GetDatabases(value0))
    fmt.Print("\n")
    break
  case "get_version":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetVersion requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetVersion())
    fmt.Print("\n")
    break
  case "start_heap_profile":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "StartHeapProfile requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.StartHeapProfile(value0))
    fmt.Print("\n")
    break
  case "stop_heap_profile":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "StopHeapProfile requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.StopHeapProfile(value0))
    fmt.Print("\n")
    break
  case "get_heap_profile":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetHeapProfile requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.GetHeapProfile(value0))
    fmt.Print("\n")
    break
  case "get_memory_gpu":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetMemoryGpu requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.GetMemoryGpu(value0))
    fmt.Print("\n")
    break
  case "get_memory_summary":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetMemorySummary requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.GetMemorySummary(value0))
    fmt.Print("\n")
    break
  case "clear_cpu_memory":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ClearCPUMemory requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.ClearCPUMemory(value0))
    fmt.Print("\n")
    break
  case "clear_gpu_memory":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ClearGpuMemory requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.ClearGpuMemory(value0))
    fmt.Print("\n")
    break
  case "sql_execute":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "SqlExecute requires 5 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    tmp4, err158 := (strconv.Atoi(flag.Arg(5)))
    if err158 != nil {
      Usage()
      return
    }
    argvalue4 := int32(tmp4)
    value4 := argvalue4
    fmt.Print(client.SqlExecute(value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "sql_execute_df":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "SqlExecuteDf requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err161 := (strconv.Atoi(flag.Arg(3)))
    if err161 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    fmt.Print(client.SqlExecuteDf(value0, value1, value2))
    fmt.Print("\n")
    break
  case "sql_execute_gpudf":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "SqlExecuteGpudf requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    tmp2, err164 := (strconv.Atoi(flag.Arg(3)))
    if err164 != nil {
      Usage()
      return
    }
    argvalue2 := int32(tmp2)
    value2 := argvalue2
    tmp3, err165 := (strconv.Atoi(flag.Arg(4)))
    if err165 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    fmt.Print(client.SqlExecuteGpudf(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "interrupt":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Interrupt requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.Interrupt(value0))
    fmt.Print("\n")
    break
  case "sql_validate":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SqlValidate requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.SqlValidate(value0, value1))
    fmt.Print("\n")
    break
  case "set_execution_mode":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "SetExecutionMode requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    tmp1, err := (strconv.Atoi(flag.Arg(2)))
    if err != nil {
      Usage()
     return
    }
    argvalue1 := mapd.TExecuteMode(tmp1)
    value1 := argvalue1
    fmt.Print(client.SetExecutionMode(value0, value1))
    fmt.Print("\n")
    break
  case "render_vega":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "RenderVega requires 5 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1, err171 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err171 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    tmp3, err173 := (strconv.Atoi(flag.Arg(4)))
    if err173 != nil {
      Usage()
      return
    }
    argvalue3 := int32(tmp3)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    fmt.Print(client.RenderVega(value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "get_result_row_for_pixel":
    if flag.NArg() - 1 != 7 {
      fmt.Fprintln(os.Stderr, "GetResultRowForPixel requires 7 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1, err176 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err176 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg177 := flag.Arg(3)
    mbTrans178 := thrift.NewTMemoryBufferLen(len(arg177))
    defer mbTrans178.Close()
    _, err179 := mbTrans178.WriteString(arg177)
    if err179 != nil {
      Usage()
      return
    }
    factory180 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt181 := factory180.GetProtocol(mbTrans178)
    argvalue2 := mapd.NewTPixel()
    err182 := argvalue2.Read(jsProt181)
    if err182 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    arg183 := flag.Arg(4)
    mbTrans184 := thrift.NewTMemoryBufferLen(len(arg183))
    defer mbTrans184.Close()
    _, err185 := mbTrans184.WriteString(arg183)
    if err185 != nil { 
      Usage()
      return
    }
    factory186 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt187 := factory186.GetProtocol(mbTrans184)
    containerStruct3 := mapd.NewMapDGetResultRowForPixelArgs()
    err188 := containerStruct3.ReadField4(jsProt187)
    if err188 != nil {
      Usage()
      return
    }
    argvalue3 := containerStruct3.TableColNames
    value3 := argvalue3
    argvalue4 := flag.Arg(5) == "true"
    value4 := argvalue4
    tmp5, err190 := (strconv.Atoi(flag.Arg(6)))
    if err190 != nil {
      Usage()
      return
    }
    argvalue5 := int32(tmp5)
    value5 := argvalue5
    argvalue6 := flag.Arg(7)
    value6 := argvalue6
    fmt.Print(client.GetResultRowForPixel(value0, value1, value2, value3, value4, value5, value6))
    fmt.Print("\n")
    break
  case "get_frontend_view":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetFrontendView requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.GetFrontendView(value0, value1))
    fmt.Print("\n")
    break
  case "get_frontend_views":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GetFrontendViews requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    fmt.Print(client.GetFrontendViews(value0))
    fmt.Print("\n")
    break
  case "create_frontend_view":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "CreateFrontendView requires 5 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    argvalue4 := flag.Arg(5)
    value4 := argvalue4
    fmt.Print(client.CreateFrontendView(value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "delete_frontend_view":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "DeleteFrontendView requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.DeleteFrontendView(value0, value1))
    fmt.Print("\n")
    break
  case "get_link_view":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetLinkView requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.GetLinkView(value0, value1))
    fmt.Print("\n")
    break
  case "create_link":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "CreateLink requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.CreateLink(value0, value1, value2))
    fmt.Print("\n")
    break
  case "load_table_binary":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "LoadTableBinary requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    arg209 := flag.Arg(3)
    mbTrans210 := thrift.NewTMemoryBufferLen(len(arg209))
    defer mbTrans210.Close()
    _, err211 := mbTrans210.WriteString(arg209)
    if err211 != nil { 
      Usage()
      return
    }
    factory212 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt213 := factory212.GetProtocol(mbTrans210)
    containerStruct2 := mapd.NewMapDLoadTableBinaryArgs()
    err214 := containerStruct2.ReadField3(jsProt213)
    if err214 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.Rows
    value2 := argvalue2
    fmt.Print(client.LoadTableBinary(value0, value1, value2))
    fmt.Print("\n")
    break
  case "load_table":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "LoadTable requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    arg217 := flag.Arg(3)
    mbTrans218 := thrift.NewTMemoryBufferLen(len(arg217))
    defer mbTrans218.Close()
    _, err219 := mbTrans218.WriteString(arg217)
    if err219 != nil { 
      Usage()
      return
    }
    factory220 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt221 := factory220.GetProtocol(mbTrans218)
    containerStruct2 := mapd.NewMapDLoadTableArgs()
    err222 := containerStruct2.ReadField3(jsProt221)
    if err222 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.Rows
    value2 := argvalue2
    fmt.Print(client.LoadTable(value0, value1, value2))
    fmt.Print("\n")
    break
  case "detect_column_types":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "DetectColumnTypes requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    arg225 := flag.Arg(3)
    mbTrans226 := thrift.NewTMemoryBufferLen(len(arg225))
    defer mbTrans226.Close()
    _, err227 := mbTrans226.WriteString(arg225)
    if err227 != nil {
      Usage()
      return
    }
    factory228 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt229 := factory228.GetProtocol(mbTrans226)
    argvalue2 := mapd.NewTCopyParams()
    err230 := argvalue2.Read(jsProt229)
    if err230 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    fmt.Print(client.DetectColumnTypes(value0, value1, value2))
    fmt.Print("\n")
    break
  case "create_table":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "CreateTable requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    arg233 := flag.Arg(3)
    mbTrans234 := thrift.NewTMemoryBufferLen(len(arg233))
    defer mbTrans234.Close()
    _, err235 := mbTrans234.WriteString(arg233)
    if err235 != nil { 
      Usage()
      return
    }
    factory236 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt237 := factory236.GetProtocol(mbTrans234)
    containerStruct2 := mapd.NewMapDCreateTableArgs()
    err238 := containerStruct2.ReadField3(jsProt237)
    if err238 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.RowDesc
    value2 := mapd.TRowDescriptor(argvalue2)
    tmp3, err := (strconv.Atoi(flag.Arg(4)))
    if err != nil {
      Usage()
     return
    }
    argvalue3 := mapd.TTableType(tmp3)
    value3 := argvalue3
    fmt.Print(client.CreateTable(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "import_table":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "ImportTable requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    arg242 := flag.Arg(4)
    mbTrans243 := thrift.NewTMemoryBufferLen(len(arg242))
    defer mbTrans243.Close()
    _, err244 := mbTrans243.WriteString(arg242)
    if err244 != nil {
      Usage()
      return
    }
    factory245 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt246 := factory245.GetProtocol(mbTrans243)
    argvalue3 := mapd.NewTCopyParams()
    err247 := argvalue3.Read(jsProt246)
    if err247 != nil {
      Usage()
      return
    }
    value3 := argvalue3
    fmt.Print(client.ImportTable(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "import_geo_table":
    if flag.NArg() - 1 != 5 {
      fmt.Fprintln(os.Stderr, "ImportGeoTable requires 5 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    arg251 := flag.Arg(4)
    mbTrans252 := thrift.NewTMemoryBufferLen(len(arg251))
    defer mbTrans252.Close()
    _, err253 := mbTrans252.WriteString(arg251)
    if err253 != nil {
      Usage()
      return
    }
    factory254 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt255 := factory254.GetProtocol(mbTrans252)
    argvalue3 := mapd.NewTCopyParams()
    err256 := argvalue3.Read(jsProt255)
    if err256 != nil {
      Usage()
      return
    }
    value3 := argvalue3
    arg257 := flag.Arg(5)
    mbTrans258 := thrift.NewTMemoryBufferLen(len(arg257))
    defer mbTrans258.Close()
    _, err259 := mbTrans258.WriteString(arg257)
    if err259 != nil { 
      Usage()
      return
    }
    factory260 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt261 := factory260.GetProtocol(mbTrans258)
    containerStruct4 := mapd.NewMapDImportGeoTableArgs()
    err262 := containerStruct4.ReadField5(jsProt261)
    if err262 != nil {
      Usage()
      return
    }
    argvalue4 := containerStruct4.RowDesc
    value4 := mapd.TRowDescriptor(argvalue4)
    fmt.Print(client.ImportGeoTable(value0, value1, value2, value3, value4))
    fmt.Print("\n")
    break
  case "import_table_status":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ImportTableStatus requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.ImportTableStatus(value0, value1))
    fmt.Print("\n")
    break
  case "start_query":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "StartQuery requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3) == "true"
    value2 := argvalue2
    fmt.Print(client.StartQuery(value0, value1, value2))
    fmt.Print("\n")
    break
  case "execute_first_step":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ExecuteFirstStep requires 1 args")
      flag.Usage()
    }
    arg268 := flag.Arg(1)
    mbTrans269 := thrift.NewTMemoryBufferLen(len(arg268))
    defer mbTrans269.Close()
    _, err270 := mbTrans269.WriteString(arg268)
    if err270 != nil {
      Usage()
      return
    }
    factory271 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt272 := factory271.GetProtocol(mbTrans269)
    argvalue0 := mapd.NewTPendingQuery()
    err273 := argvalue0.Read(jsProt272)
    if err273 != nil {
      Usage()
      return
    }
    value0 := argvalue0
    fmt.Print(client.ExecuteFirstStep(value0))
    fmt.Print("\n")
    break
  case "broadcast_serialized_rows":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "BroadcastSerializedRows requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    arg275 := flag.Arg(2)
    mbTrans276 := thrift.NewTMemoryBufferLen(len(arg275))
    defer mbTrans276.Close()
    _, err277 := mbTrans276.WriteString(arg275)
    if err277 != nil { 
      Usage()
      return
    }
    factory278 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt279 := factory278.GetProtocol(mbTrans276)
    containerStruct1 := mapd.NewMapDBroadcastSerializedRowsArgs()
    err280 := containerStruct1.ReadField2(jsProt279)
    if err280 != nil {
      Usage()
      return
    }
    argvalue1 := containerStruct1.RowDesc
    value1 := mapd.TRowDescriptor(argvalue1)
    argvalue2, err281 := (strconv.ParseInt(flag.Arg(3), 10, 64))
    if err281 != nil {
      Usage()
      return
    }
    value2 := mapd.TQueryId(argvalue2)
    fmt.Print(client.BroadcastSerializedRows(value0, value1, value2))
    fmt.Print("\n")
    break
  case "render_vega_raw_pixels":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "RenderVegaRawPixels requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1, err283 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err283 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    tmp2, err284 := (strconv.Atoi(flag.Arg(3)))
    if err284 != nil {
      Usage()
      return
    }
    argvalue2 := int16(tmp2)
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    fmt.Print(client.RenderVegaRawPixels(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "insert_data":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "InsertData requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    arg287 := flag.Arg(2)
    mbTrans288 := thrift.NewTMemoryBufferLen(len(arg287))
    defer mbTrans288.Close()
    _, err289 := mbTrans288.WriteString(arg287)
    if err289 != nil {
      Usage()
      return
    }
    factory290 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt291 := factory290.GetProtocol(mbTrans288)
    argvalue1 := mapd.NewTInsertData()
    err292 := argvalue1.Read(jsProt291)
    if err292 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    fmt.Print(client.InsertData(value0, value1))
    fmt.Print("\n")
    break
  case "get_table_descriptor":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetTableDescriptor requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.GetTableDescriptor(value0, value1))
    fmt.Print("\n")
    break
  case "get_row_descriptor":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "GetRowDescriptor requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.GetRowDescriptor(value0, value1))
    fmt.Print("\n")
    break
  case "render":
    if flag.NArg() - 1 != 4 {
      fmt.Fprintln(os.Stderr, "Render requires 4 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    fmt.Print(client.Render(value0, value1, value2, value3))
    fmt.Print("\n")
    break
  case "get_rows_for_pixels":
    if flag.NArg() - 1 != 7 {
      fmt.Fprintln(os.Stderr, "GetRowsForPixels requires 7 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1, err302 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err302 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg303 := flag.Arg(3)
    mbTrans304 := thrift.NewTMemoryBufferLen(len(arg303))
    defer mbTrans304.Close()
    _, err305 := mbTrans304.WriteString(arg303)
    if err305 != nil { 
      Usage()
      return
    }
    factory306 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt307 := factory306.GetProtocol(mbTrans304)
    containerStruct2 := mapd.NewMapDGetRowsForPixelsArgs()
    err308 := containerStruct2.ReadField3(jsProt307)
    if err308 != nil {
      Usage()
      return
    }
    argvalue2 := containerStruct2.Pixels
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    arg310 := flag.Arg(5)
    mbTrans311 := thrift.NewTMemoryBufferLen(len(arg310))
    defer mbTrans311.Close()
    _, err312 := mbTrans311.WriteString(arg310)
    if err312 != nil { 
      Usage()
      return
    }
    factory313 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt314 := factory313.GetProtocol(mbTrans311)
    containerStruct4 := mapd.NewMapDGetRowsForPixelsArgs()
    err315 := containerStruct4.ReadField5(jsProt314)
    if err315 != nil {
      Usage()
      return
    }
    argvalue4 := containerStruct4.ColNames
    value4 := argvalue4
    argvalue5 := flag.Arg(6) == "true"
    value5 := argvalue5
    argvalue6 := flag.Arg(7)
    value6 := argvalue6
    fmt.Print(client.GetRowsForPixels(value0, value1, value2, value3, value4, value5, value6))
    fmt.Print("\n")
    break
  case "get_row_for_pixel":
    if flag.NArg() - 1 != 8 {
      fmt.Fprintln(os.Stderr, "GetRowForPixel requires 8 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := mapd.TSessionId(argvalue0)
    argvalue1, err319 := (strconv.ParseInt(flag.Arg(2), 10, 64))
    if err319 != nil {
      Usage()
      return
    }
    value1 := argvalue1
    arg320 := flag.Arg(3)
    mbTrans321 := thrift.NewTMemoryBufferLen(len(arg320))
    defer mbTrans321.Close()
    _, err322 := mbTrans321.WriteString(arg320)
    if err322 != nil {
      Usage()
      return
    }
    factory323 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt324 := factory323.GetProtocol(mbTrans321)
    argvalue2 := mapd.NewTPixel()
    err325 := argvalue2.Read(jsProt324)
    if err325 != nil {
      Usage()
      return
    }
    value2 := argvalue2
    argvalue3 := flag.Arg(4)
    value3 := argvalue3
    arg327 := flag.Arg(5)
    mbTrans328 := thrift.NewTMemoryBufferLen(len(arg327))
    defer mbTrans328.Close()
    _, err329 := mbTrans328.WriteString(arg327)
    if err329 != nil { 
      Usage()
      return
    }
    factory330 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt331 := factory330.GetProtocol(mbTrans328)
    containerStruct4 := mapd.NewMapDGetRowForPixelArgs()
    err332 := containerStruct4.ReadField5(jsProt331)
    if err332 != nil {
      Usage()
      return
    }
    argvalue4 := containerStruct4.ColNames
    value4 := argvalue4
    argvalue5 := flag.Arg(6) == "true"
    value5 := argvalue5
    tmp6, err334 := (strconv.Atoi(flag.Arg(7)))
    if err334 != nil {
      Usage()
      return
    }
    argvalue6 := int32(tmp6)
    value6 := argvalue6
    argvalue7 := flag.Arg(8)
    value7 := argvalue7
    fmt.Print(client.GetRowForPixel(value0, value1, value2, value3, value4, value5, value6, value7))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
