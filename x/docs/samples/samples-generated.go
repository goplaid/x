package samples

var PageFuncAndEventFuncDefinition = string([]byte{0x74, 0x79, 0x70, 0x65, 0x20, 0x50, 0x61, 0x67, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x28, 0x63, 0x74, 0x78, 0x20, 0x2a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x29, 0x20, 0x28, 0x72, 0x20, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x65, 0x72, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x29, 0xa, 0x74, 0x79, 0x70, 0x65, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x28, 0x63, 0x74, 0x78, 0x20, 0x2a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x29, 0x20, 0x28, 0x72, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x65, 0x72, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x29, 0xa})
var EventFuncHubDefinition = string([]byte{0x74, 0x79, 0x70, 0x65, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x48, 0x75, 0x62, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x20, 0x7b, 0xa, 0x9, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x28, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x49, 0x64, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2c, 0x20, 0x65, 0x66, 0x20, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x29, 0x20, 0x28, 0x6b, 0x65, 0x79, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x29, 0xa, 0x7d, 0xa})
var HelloWorldMainSample = string([]byte{0x66, 0x6d, 0x74, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x6c, 0x6e, 0x28, 0x22, 0x53, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x64, 0x6f, 0x63, 0x73, 0x20, 0x61, 0x74, 0x20, 0x3a, 0x39, 0x30, 0x30, 0x30, 0x22, 0x29, 0xa, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x28, 0x22, 0x2f, 0x22, 0x2c, 0x20, 0x6d, 0x75, 0x78, 0x29, 0xa, 0x65, 0x72, 0x72, 0x20, 0x3a, 0x3d, 0x20, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x28, 0x22, 0x3a, 0x39, 0x30, 0x30, 0x30, 0x22, 0x2c, 0x20, 0x6e, 0x69, 0x6c, 0x29, 0xa, 0x69, 0x66, 0x20, 0x65, 0x72, 0x72, 0x20, 0x21, 0x3d, 0x20, 0x6e, 0x69, 0x6c, 0x20, 0x7b, 0xa, 0x9, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x28, 0x65, 0x72, 0x72, 0x29, 0xa, 0x7d})
var HelloWorldReloadSample = string([]byte{0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x28, 0xa, 0x9, 0x22, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xa, 0xa, 0x9, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x70, 0x6c, 0x61, 0x69, 0x64, 0x2f, 0x77, 0x65, 0x62, 0x22, 0xa, 0x9, 0x2e, 0x20, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x67, 0x6f, 0x22, 0xa, 0x29, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x20, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x28, 0x63, 0x74, 0x78, 0x20, 0x2a, 0x77, 0x65, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x29, 0x20, 0x28, 0x70, 0x72, 0x20, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x65, 0x72, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x63, 0x74, 0x78, 0x2e, 0x48, 0x75, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x63, 0x28, 0x22, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2c, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x29, 0xa, 0x9, 0x70, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x20, 0x3d, 0x20, 0x44, 0x69, 0x76, 0x28, 0xa, 0x9, 0x9, 0x48, 0x31, 0x28, 0x22, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x54, 0x65, 0x78, 0x74, 0x28, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x4e, 0x6f, 0x77, 0x28, 0x29, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x28, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x52, 0x46, 0x43, 0x33, 0x33, 0x33, 0x39, 0x4e, 0x61, 0x6e, 0x6f, 0x29, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x77, 0x65, 0x62, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x28, 0xa, 0x9, 0x9, 0x9, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x28, 0x22, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x50, 0x61, 0x67, 0x65, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x29, 0x2e, 0x4f, 0x6e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x28, 0x22, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x29, 0xa, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x28, 0x63, 0x74, 0x78, 0x20, 0x2a, 0x77, 0x65, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x29, 0x20, 0x28, 0x65, 0x72, 0x20, 0x77, 0x65, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x65, 0x72, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x3d, 0x20, 0x74, 0x72, 0x75, 0x65, 0xa, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0xa, 0x7d, 0xa})
var HelloWorldSample = string([]byte{0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x28, 0xa, 0x9, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x70, 0x6c, 0x61, 0x69, 0x64, 0x2f, 0x77, 0x65, 0x62, 0x22, 0xa, 0x9, 0x2e, 0x20, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x67, 0x6f, 0x22, 0xa, 0x29, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x20, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x28, 0x63, 0x74, 0x78, 0x20, 0x2a, 0x77, 0x65, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x29, 0x20, 0x28, 0x70, 0x72, 0x20, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x65, 0x72, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x70, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x20, 0x3d, 0x20, 0x48, 0x31, 0x28, 0x22, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x22, 0x29, 0xa, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0xa, 0x7d})
var TypeSafeBuilderSample = string([]byte{0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x28, 0xa, 0x9, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x70, 0x6c, 0x61, 0x69, 0x64, 0x2f, 0x77, 0x65, 0x62, 0x22, 0xa, 0x9, 0x2e, 0x20, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x67, 0x6f, 0x22, 0xa, 0x29, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x20, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x28, 0x61, 0x72, 0x67, 0x73, 0x20, 0x2e, 0x2e, 0x2e, 0x48, 0x54, 0x4d, 0x4c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x29, 0x20, 0x48, 0x54, 0x4d, 0x4c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x20, 0x7b, 0xa, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x20, 0x5b, 0x5d, 0x48, 0x54, 0x4d, 0x4c, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0xa, 0x9, 0x66, 0x6f, 0x72, 0x20, 0x5f, 0x2c, 0x20, 0x61, 0x72, 0x67, 0x20, 0x3a, 0x3d, 0x20, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x61, 0x72, 0x67, 0x73, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x20, 0x3d, 0x20, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x28, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x2c, 0x20, 0x44, 0x69, 0x76, 0x28, 0x61, 0x72, 0x67, 0x29, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x28, 0x22, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x22, 0x29, 0x29, 0xa, 0x9, 0x7d, 0xa, 0xa, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x48, 0x54, 0x4d, 0x4c, 0x28, 0xa, 0x9, 0x9, 0x48, 0x65, 0x61, 0x64, 0x28, 0xa, 0x9, 0x9, 0x9, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x28, 0x22, 0x58, 0x4d, 0x4c, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x47, 0x6f, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x42, 0x6f, 0x64, 0x79, 0x28, 0xa, 0x9, 0x9, 0x9, 0x48, 0x31, 0x28, 0x22, 0x58, 0x4d, 0x4c, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x47, 0x6f, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x9, 0x50, 0x28, 0x29, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x28, 0x22, 0x74, 0x68, 0x69, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x62, 0x65, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x61, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x20, 0x6d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x20, 0x74, 0x6f, 0x20, 0x58, 0x4d, 0x4c, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x9, 0x41, 0x28, 0x29, 0x2e, 0x48, 0x72, 0x65, 0x66, 0x28, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x22, 0x29, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x28, 0x22, 0x47, 0x6f, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x9, 0x50, 0x28, 0xa, 0x9, 0x9, 0x9, 0x9, 0x54, 0x65, 0x78, 0x74, 0x28, 0x22, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x73, 0x6f, 0x6d, 0x65, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x9, 0x9, 0x42, 0x28, 0x22, 0x6d, 0x69, 0x78, 0x65, 0x64, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x9, 0x9, 0x54, 0x65, 0x78, 0x74, 0x28, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x20, 0x46, 0x6f, 0x72, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x73, 0x65, 0x65, 0x20, 0x74, 0x68, 0x65, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x9, 0x9, 0x41, 0x28, 0x29, 0x2e, 0x48, 0x72, 0x65, 0x66, 0x28, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x22, 0x29, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x28, 0x22, 0x47, 0x6f, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x9, 0x9, 0x54, 0x65, 0x78, 0x74, 0x28, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x9, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x9, 0x50, 0x28, 0x29, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x28, 0x22, 0x73, 0x6f, 0x6d, 0x65, 0x20, 0x74, 0x65, 0x78, 0x74, 0x22, 0x29, 0x2c, 0xa, 0xa, 0x9, 0x9, 0x9, 0x50, 0x28, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x2e, 0x2e, 0x2e, 0x29, 0x2c, 0xa, 0x9, 0x9, 0x29, 0x2c, 0xa, 0x9, 0x29, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x20, 0x54, 0x79, 0x70, 0x65, 0x53, 0x61, 0x66, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x46, 0x28, 0x63, 0x74, 0x78, 0x20, 0x2a, 0x77, 0x65, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x29, 0x20, 0x28, 0x70, 0x72, 0x20, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x65, 0x72, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x70, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x20, 0x3d, 0x20, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x28, 0x48, 0x35, 0x28, 0x22, 0x31, 0x22, 0x29, 0x2c, 0x20, 0x42, 0x28, 0x22, 0x32, 0x22, 0x29, 0x2c, 0x20, 0x53, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x28, 0x22, 0x33, 0x22, 0x29, 0x29, 0xa, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0xa, 0x7d, 0xa})
var DemoLayoutSample = string([]byte{0x66, 0x75, 0x6e, 0x63, 0x20, 0x64, 0x65, 0x6d, 0x6f, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x28, 0x69, 0x6e, 0x20, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x29, 0x20, 0x28, 0x6f, 0x75, 0x74, 0x20, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x46, 0x75, 0x6e, 0x63, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x66, 0x75, 0x6e, 0x63, 0x28, 0x63, 0x74, 0x78, 0x20, 0x2a, 0x77, 0x65, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x29, 0x20, 0x28, 0x70, 0x72, 0x20, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2c, 0x20, 0x65, 0x72, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x29, 0x20, 0x7b, 0xa, 0xa, 0x9, 0x9, 0x63, 0x74, 0x78, 0x2e, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x48, 0x54, 0x4d, 0x4c, 0x28, 0x60, 0xa, 0x9, 0x9, 0x9, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x27, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x76, 0x75, 0x65, 0x2e, 0x6a, 0x73, 0x27, 0x3e, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0xa, 0x9, 0x9, 0x60, 0x29, 0xa, 0xa, 0x9, 0x9, 0x63, 0x74, 0x78, 0x2e, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x61, 0x69, 0x6c, 0x48, 0x54, 0x4d, 0x4c, 0x28, 0x60, 0x3c, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x20, 0x73, 0x72, 0x63, 0x3d, 0x27, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6a, 0x73, 0x27, 0x3e, 0x3c, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3e, 0x60, 0x29, 0xa, 0x9, 0x9, 0x63, 0x74, 0x78, 0x2e, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x48, 0x54, 0x4d, 0x4c, 0x28, 0x60, 0xa, 0x9, 0x9, 0x3c, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3e, 0xa, 0x9, 0x9, 0x9, 0x5b, 0x76, 0x2d, 0x63, 0x6c, 0x6f, 0x61, 0x6b, 0x5d, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x9, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x6e, 0x6f, 0x6e, 0x65, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x9, 0x3c, 0x2f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3e, 0xa, 0x9, 0x9, 0x60, 0x29, 0xa, 0xa, 0x9, 0x9, 0x76, 0x61, 0x72, 0x20, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x20, 0x77, 0x65, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0xa, 0x9, 0x9, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x2c, 0x20, 0x65, 0x72, 0x72, 0x20, 0x3d, 0x20, 0x69, 0x6e, 0x28, 0x63, 0x74, 0x78, 0x29, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x65, 0x72, 0x72, 0x20, 0x21, 0x3d, 0x20, 0x6e, 0x69, 0x6c, 0x20, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x28, 0x65, 0x72, 0x72, 0x29, 0xa, 0x9, 0x9, 0x7d, 0xa, 0xa, 0x9, 0x9, 0x70, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x20, 0x3d, 0x20, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xa, 0xa, 0x9, 0x9, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa})
var HelloWorldMuxSample1 = string([]byte{0x6d, 0x75, 0x78, 0x20, 0x3a, 0x3d, 0x20, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x4d, 0x75, 0x78, 0x28, 0x29})
var HelloWorldMuxSample2 = string([]byte{0x77, 0x62, 0x20, 0x3a, 0x3d, 0x20, 0x77, 0x65, 0x62, 0x2e, 0x4e, 0x65, 0x77, 0x28, 0x29, 0xa, 0x6d, 0x75, 0x78, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x28, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x68, 0x2c, 0x20, 0x77, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x28, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x29, 0x29})
var HelloWorldReloadMuxSample1 = string([]byte{0x6d, 0x75, 0x78, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x28, 0xa, 0x9, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x74, 0x68, 0x2c, 0xa, 0x9, 0x77, 0x62, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x28, 0xa, 0x9, 0x9, 0x64, 0x65, 0x6d, 0x6f, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x28, 0xa, 0x9, 0x9, 0x9, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x2c, 0xa, 0x9, 0x9, 0x29, 0x2c, 0xa, 0x9, 0x29, 0x2c, 0xa, 0x29})
