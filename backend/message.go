package app

// * parseTag parses "endpoints" field tag into endpointsTag struct.
//   <https://github.com/GoogleCloudPlatform/go-endpoints/blob/master/endpoints/apiconfig.go#L643>
//
//       type MyMessage struct {
//           SomeField int `endpoints:"req,min=0,max=100,desc="Int field"`
//           WithDefault string `endpoints:"d=Hello gopher"`
//       }
//
//   - req, required (boolean)
//   - d=val, default value
//   - min=val, min value
//   - max=val, max value
//   - desc=val, description

// ItemsRequestMessage ...
type ItemsRequestMessage struct {
}

// ItemsResponseMessage ...
type ItemsResponseMessage struct {
	Items []*ItemResponseMessage `json:"items"`
}

// ItemResponseMessage ...
type ItemResponseMessage struct {
}
