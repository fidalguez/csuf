// Code generated by protoc-gen-go.
// DO NOT EDIT!

/*
Package appengine is a generated protocol buffer package.

It is generated from these files:
	appengine_internal/image

It has these top-level messages:
	ImagesServiceError
	ImagesServiceTransform
	Transform
	ImageData
	InputSettings
	OutputSettings
	ImagesTransformRequest
	ImagesTransformResponse
	CompositeImageOptions
	ImagesCanvas
	ImagesCompositeRequest
	ImagesCompositeResponse
	ImagesHistogramRequest
	ImagesHistogram
	ImagesHistogramResponse
	ImagesGetUrlBaseRequest
	ImagesGetUrlBaseResponse
	ImagesDeleteUrlBaseRequest
	ImagesDeleteUrlBaseResponse
*/
package appengine

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ImagesServiceError_ErrorCode int32

const (
	ImagesServiceError_UNSPECIFIED_ERROR  ImagesServiceError_ErrorCode = 1
	ImagesServiceError_BAD_TRANSFORM_DATA ImagesServiceError_ErrorCode = 2
	ImagesServiceError_NOT_IMAGE          ImagesServiceError_ErrorCode = 3
	ImagesServiceError_BAD_IMAGE_DATA     ImagesServiceError_ErrorCode = 4
	ImagesServiceError_IMAGE_TOO_LARGE    ImagesServiceError_ErrorCode = 5
	ImagesServiceError_INVALID_BLOB_KEY   ImagesServiceError_ErrorCode = 6
	ImagesServiceError_ACCESS_DENIED      ImagesServiceError_ErrorCode = 7
	ImagesServiceError_OBJECT_NOT_FOUND   ImagesServiceError_ErrorCode = 8
)

var ImagesServiceError_ErrorCode_name = map[int32]string{
	1: "UNSPECIFIED_ERROR",
	2: "BAD_TRANSFORM_DATA",
	3: "NOT_IMAGE",
	4: "BAD_IMAGE_DATA",
	5: "IMAGE_TOO_LARGE",
	6: "INVALID_BLOB_KEY",
	7: "ACCESS_DENIED",
	8: "OBJECT_NOT_FOUND",
}
var ImagesServiceError_ErrorCode_value = map[string]int32{
	"UNSPECIFIED_ERROR":  1,
	"BAD_TRANSFORM_DATA": 2,
	"NOT_IMAGE":          3,
	"BAD_IMAGE_DATA":     4,
	"IMAGE_TOO_LARGE":    5,
	"INVALID_BLOB_KEY":   6,
	"ACCESS_DENIED":      7,
	"OBJECT_NOT_FOUND":   8,
}

func (x ImagesServiceError_ErrorCode) Enum() *ImagesServiceError_ErrorCode {
	p := new(ImagesServiceError_ErrorCode)
	*p = x
	return p
}
func (x ImagesServiceError_ErrorCode) String() string {
	return proto.EnumName(ImagesServiceError_ErrorCode_name, int32(x))
}
func (x *ImagesServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImagesServiceError_ErrorCode_value, data, "ImagesServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = ImagesServiceError_ErrorCode(value)
	return nil
}

type ImagesServiceTransform_Type int32

const (
	ImagesServiceTransform_RESIZE           ImagesServiceTransform_Type = 1
	ImagesServiceTransform_ROTATE           ImagesServiceTransform_Type = 2
	ImagesServiceTransform_HORIZONTAL_FLIP  ImagesServiceTransform_Type = 3
	ImagesServiceTransform_VERTICAL_FLIP    ImagesServiceTransform_Type = 4
	ImagesServiceTransform_CROP             ImagesServiceTransform_Type = 5
	ImagesServiceTransform_IM_FEELING_LUCKY ImagesServiceTransform_Type = 6
)

var ImagesServiceTransform_Type_name = map[int32]string{
	1: "RESIZE",
	2: "ROTATE",
	3: "HORIZONTAL_FLIP",
	4: "VERTICAL_FLIP",
	5: "CROP",
	6: "IM_FEELING_LUCKY",
}
var ImagesServiceTransform_Type_value = map[string]int32{
	"RESIZE":           1,
	"ROTATE":           2,
	"HORIZONTAL_FLIP":  3,
	"VERTICAL_FLIP":    4,
	"CROP":             5,
	"IM_FEELING_LUCKY": 6,
}

func (x ImagesServiceTransform_Type) Enum() *ImagesServiceTransform_Type {
	p := new(ImagesServiceTransform_Type)
	*p = x
	return p
}
func (x ImagesServiceTransform_Type) String() string {
	return proto.EnumName(ImagesServiceTransform_Type_name, int32(x))
}
func (x *ImagesServiceTransform_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ImagesServiceTransform_Type_value, data, "ImagesServiceTransform_Type")
	if err != nil {
		return err
	}
	*x = ImagesServiceTransform_Type(value)
	return nil
}

type InputSettings_ORIENTATION_CORRECTION_TYPE int32

const (
	InputSettings_UNCHANGED_ORIENTATION InputSettings_ORIENTATION_CORRECTION_TYPE = 0
	InputSettings_CORRECT_ORIENTATION   InputSettings_ORIENTATION_CORRECTION_TYPE = 1
)

var InputSettings_ORIENTATION_CORRECTION_TYPE_name = map[int32]string{
	0: "UNCHANGED_ORIENTATION",
	1: "CORRECT_ORIENTATION",
}
var InputSettings_ORIENTATION_CORRECTION_TYPE_value = map[string]int32{
	"UNCHANGED_ORIENTATION": 0,
	"CORRECT_ORIENTATION":   1,
}

func (x InputSettings_ORIENTATION_CORRECTION_TYPE) Enum() *InputSettings_ORIENTATION_CORRECTION_TYPE {
	p := new(InputSettings_ORIENTATION_CORRECTION_TYPE)
	*p = x
	return p
}
func (x InputSettings_ORIENTATION_CORRECTION_TYPE) String() string {
	return proto.EnumName(InputSettings_ORIENTATION_CORRECTION_TYPE_name, int32(x))
}
func (x *InputSettings_ORIENTATION_CORRECTION_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(InputSettings_ORIENTATION_CORRECTION_TYPE_value, data, "InputSettings_ORIENTATION_CORRECTION_TYPE")
	if err != nil {
		return err
	}
	*x = InputSettings_ORIENTATION_CORRECTION_TYPE(value)
	return nil
}

type OutputSettings_MIME_TYPE int32

const (
	OutputSettings_PNG  OutputSettings_MIME_TYPE = 0
	OutputSettings_JPEG OutputSettings_MIME_TYPE = 1
	OutputSettings_WEBP OutputSettings_MIME_TYPE = 2
)

var OutputSettings_MIME_TYPE_name = map[int32]string{
	0: "PNG",
	1: "JPEG",
	2: "WEBP",
}
var OutputSettings_MIME_TYPE_value = map[string]int32{
	"PNG":  0,
	"JPEG": 1,
	"WEBP": 2,
}

func (x OutputSettings_MIME_TYPE) Enum() *OutputSettings_MIME_TYPE {
	p := new(OutputSettings_MIME_TYPE)
	*p = x
	return p
}
func (x OutputSettings_MIME_TYPE) String() string {
	return proto.EnumName(OutputSettings_MIME_TYPE_name, int32(x))
}
func (x *OutputSettings_MIME_TYPE) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OutputSettings_MIME_TYPE_value, data, "OutputSettings_MIME_TYPE")
	if err != nil {
		return err
	}
	*x = OutputSettings_MIME_TYPE(value)
	return nil
}

type CompositeImageOptions_ANCHOR int32

const (
	CompositeImageOptions_TOP_LEFT     CompositeImageOptions_ANCHOR = 0
	CompositeImageOptions_TOP          CompositeImageOptions_ANCHOR = 1
	CompositeImageOptions_TOP_RIGHT    CompositeImageOptions_ANCHOR = 2
	CompositeImageOptions_LEFT         CompositeImageOptions_ANCHOR = 3
	CompositeImageOptions_CENTER       CompositeImageOptions_ANCHOR = 4
	CompositeImageOptions_RIGHT        CompositeImageOptions_ANCHOR = 5
	CompositeImageOptions_BOTTOM_LEFT  CompositeImageOptions_ANCHOR = 6
	CompositeImageOptions_BOTTOM       CompositeImageOptions_ANCHOR = 7
	CompositeImageOptions_BOTTOM_RIGHT CompositeImageOptions_ANCHOR = 8
)

var CompositeImageOptions_ANCHOR_name = map[int32]string{
	0: "TOP_LEFT",
	1: "TOP",
	2: "TOP_RIGHT",
	3: "LEFT",
	4: "CENTER",
	5: "RIGHT",
	6: "BOTTOM_LEFT",
	7: "BOTTOM",
	8: "BOTTOM_RIGHT",
}
var CompositeImageOptions_ANCHOR_value = map[string]int32{
	"TOP_LEFT":     0,
	"TOP":          1,
	"TOP_RIGHT":    2,
	"LEFT":         3,
	"CENTER":       4,
	"RIGHT":        5,
	"BOTTOM_LEFT":  6,
	"BOTTOM":       7,
	"BOTTOM_RIGHT": 8,
}

func (x CompositeImageOptions_ANCHOR) Enum() *CompositeImageOptions_ANCHOR {
	p := new(CompositeImageOptions_ANCHOR)
	*p = x
	return p
}
func (x CompositeImageOptions_ANCHOR) String() string {
	return proto.EnumName(CompositeImageOptions_ANCHOR_name, int32(x))
}
func (x *CompositeImageOptions_ANCHOR) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CompositeImageOptions_ANCHOR_value, data, "CompositeImageOptions_ANCHOR")
	if err != nil {
		return err
	}
	*x = CompositeImageOptions_ANCHOR(value)
	return nil
}

type ImagesServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ImagesServiceError) Reset()         { *m = ImagesServiceError{} }
func (m *ImagesServiceError) String() string { return proto.CompactTextString(m) }
func (*ImagesServiceError) ProtoMessage()    {}

type ImagesServiceTransform struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ImagesServiceTransform) Reset()         { *m = ImagesServiceTransform{} }
func (m *ImagesServiceTransform) String() string { return proto.CompactTextString(m) }
func (*ImagesServiceTransform) ProtoMessage()    {}

type Transform struct {
	Width            *int32   `protobuf:"varint,1,opt,name=width" json:"width,omitempty"`
	Height           *int32   `protobuf:"varint,2,opt,name=height" json:"height,omitempty"`
	CropToFit        *bool    `protobuf:"varint,11,opt,name=crop_to_fit,def=0" json:"crop_to_fit,omitempty"`
	CropOffsetX      *float32 `protobuf:"fixed32,12,opt,name=crop_offset_x,def=0.5" json:"crop_offset_x,omitempty"`
	CropOffsetY      *float32 `protobuf:"fixed32,13,opt,name=crop_offset_y,def=0.5" json:"crop_offset_y,omitempty"`
	Rotate           *int32   `protobuf:"varint,3,opt,name=rotate,def=0" json:"rotate,omitempty"`
	HorizontalFlip   *bool    `protobuf:"varint,4,opt,name=horizontal_flip,def=0" json:"horizontal_flip,omitempty"`
	VerticalFlip     *bool    `protobuf:"varint,5,opt,name=vertical_flip,def=0" json:"vertical_flip,omitempty"`
	CropLeftX        *float32 `protobuf:"fixed32,6,opt,name=crop_left_x,def=0" json:"crop_left_x,omitempty"`
	CropTopY         *float32 `protobuf:"fixed32,7,opt,name=crop_top_y,def=0" json:"crop_top_y,omitempty"`
	CropRightX       *float32 `protobuf:"fixed32,8,opt,name=crop_right_x,def=1" json:"crop_right_x,omitempty"`
	CropBottomY      *float32 `protobuf:"fixed32,9,opt,name=crop_bottom_y,def=1" json:"crop_bottom_y,omitempty"`
	Autolevels       *bool    `protobuf:"varint,10,opt,name=autolevels,def=0" json:"autolevels,omitempty"`
	AllowStretch     *bool    `protobuf:"varint,14,opt,name=allow_stretch,def=0" json:"allow_stretch,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Transform) Reset()         { *m = Transform{} }
func (m *Transform) String() string { return proto.CompactTextString(m) }
func (*Transform) ProtoMessage()    {}

const Default_Transform_CropToFit bool = false
const Default_Transform_CropOffsetX float32 = 0.5
const Default_Transform_CropOffsetY float32 = 0.5
const Default_Transform_Rotate int32 = 0
const Default_Transform_HorizontalFlip bool = false
const Default_Transform_VerticalFlip bool = false
const Default_Transform_CropLeftX float32 = 0
const Default_Transform_CropTopY float32 = 0
const Default_Transform_CropRightX float32 = 1
const Default_Transform_CropBottomY float32 = 1
const Default_Transform_Autolevels bool = false
const Default_Transform_AllowStretch bool = false

func (m *Transform) GetWidth() int32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *Transform) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *Transform) GetCropToFit() bool {
	if m != nil && m.CropToFit != nil {
		return *m.CropToFit
	}
	return Default_Transform_CropToFit
}

func (m *Transform) GetCropOffsetX() float32 {
	if m != nil && m.CropOffsetX != nil {
		return *m.CropOffsetX
	}
	return Default_Transform_CropOffsetX
}

func (m *Transform) GetCropOffsetY() float32 {
	if m != nil && m.CropOffsetY != nil {
		return *m.CropOffsetY
	}
	return Default_Transform_CropOffsetY
}

func (m *Transform) GetRotate() int32 {
	if m != nil && m.Rotate != nil {
		return *m.Rotate
	}
	return Default_Transform_Rotate
}

func (m *Transform) GetHorizontalFlip() bool {
	if m != nil && m.HorizontalFlip != nil {
		return *m.HorizontalFlip
	}
	return Default_Transform_HorizontalFlip
}

func (m *Transform) GetVerticalFlip() bool {
	if m != nil && m.VerticalFlip != nil {
		return *m.VerticalFlip
	}
	return Default_Transform_VerticalFlip
}

func (m *Transform) GetCropLeftX() float32 {
	if m != nil && m.CropLeftX != nil {
		return *m.CropLeftX
	}
	return Default_Transform_CropLeftX
}

func (m *Transform) GetCropTopY() float32 {
	if m != nil && m.CropTopY != nil {
		return *m.CropTopY
	}
	return Default_Transform_CropTopY
}

func (m *Transform) GetCropRightX() float32 {
	if m != nil && m.CropRightX != nil {
		return *m.CropRightX
	}
	return Default_Transform_CropRightX
}

func (m *Transform) GetCropBottomY() float32 {
	if m != nil && m.CropBottomY != nil {
		return *m.CropBottomY
	}
	return Default_Transform_CropBottomY
}

func (m *Transform) GetAutolevels() bool {
	if m != nil && m.Autolevels != nil {
		return *m.Autolevels
	}
	return Default_Transform_Autolevels
}

func (m *Transform) GetAllowStretch() bool {
	if m != nil && m.AllowStretch != nil {
		return *m.AllowStretch
	}
	return Default_Transform_AllowStretch
}

type ImageData struct {
	Content          []byte  `protobuf:"bytes,1,req,name=content" json:"content,omitempty"`
	BlobKey          *string `protobuf:"bytes,2,opt,name=blob_key" json:"blob_key,omitempty"`
	Width            *int32  `protobuf:"varint,3,opt,name=width" json:"width,omitempty"`
	Height           *int32  `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ImageData) Reset()         { *m = ImageData{} }
func (m *ImageData) String() string { return proto.CompactTextString(m) }
func (*ImageData) ProtoMessage()    {}

func (m *ImageData) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ImageData) GetBlobKey() string {
	if m != nil && m.BlobKey != nil {
		return *m.BlobKey
	}
	return ""
}

func (m *ImageData) GetWidth() int32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *ImageData) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type InputSettings struct {
	CorrectExifOrientation     *InputSettings_ORIENTATION_CORRECTION_TYPE `protobuf:"varint,1,opt,name=correct_exif_orientation,enum=appengine.InputSettings_ORIENTATION_CORRECTION_TYPE,def=0" json:"correct_exif_orientation,omitempty"`
	ParseMetadata              *bool                                      `protobuf:"varint,2,opt,name=parse_metadata,def=0" json:"parse_metadata,omitempty"`
	TransparentSubstitutionRgb *int32                                     `protobuf:"varint,3,opt,name=transparent_substitution_rgb" json:"transparent_substitution_rgb,omitempty"`
	XXX_unrecognized           []byte                                     `json:"-"`
}

func (m *InputSettings) Reset()         { *m = InputSettings{} }
func (m *InputSettings) String() string { return proto.CompactTextString(m) }
func (*InputSettings) ProtoMessage()    {}

const Default_InputSettings_CorrectExifOrientation InputSettings_ORIENTATION_CORRECTION_TYPE = InputSettings_UNCHANGED_ORIENTATION
const Default_InputSettings_ParseMetadata bool = false

func (m *InputSettings) GetCorrectExifOrientation() InputSettings_ORIENTATION_CORRECTION_TYPE {
	if m != nil && m.CorrectExifOrientation != nil {
		return *m.CorrectExifOrientation
	}
	return Default_InputSettings_CorrectExifOrientation
}

func (m *InputSettings) GetParseMetadata() bool {
	if m != nil && m.ParseMetadata != nil {
		return *m.ParseMetadata
	}
	return Default_InputSettings_ParseMetadata
}

func (m *InputSettings) GetTransparentSubstitutionRgb() int32 {
	if m != nil && m.TransparentSubstitutionRgb != nil {
		return *m.TransparentSubstitutionRgb
	}
	return 0
}

type OutputSettings struct {
	MimeType         *OutputSettings_MIME_TYPE `protobuf:"varint,1,opt,name=mime_type,enum=appengine.OutputSettings_MIME_TYPE,def=0" json:"mime_type,omitempty"`
	Quality          *int32                    `protobuf:"varint,2,opt,name=quality" json:"quality,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *OutputSettings) Reset()         { *m = OutputSettings{} }
func (m *OutputSettings) String() string { return proto.CompactTextString(m) }
func (*OutputSettings) ProtoMessage()    {}

const Default_OutputSettings_MimeType OutputSettings_MIME_TYPE = OutputSettings_PNG

func (m *OutputSettings) GetMimeType() OutputSettings_MIME_TYPE {
	if m != nil && m.MimeType != nil {
		return *m.MimeType
	}
	return Default_OutputSettings_MimeType
}

func (m *OutputSettings) GetQuality() int32 {
	if m != nil && m.Quality != nil {
		return *m.Quality
	}
	return 0
}

type ImagesTransformRequest struct {
	Image            *ImageData      `protobuf:"bytes,1,req,name=image" json:"image,omitempty"`
	Transform        []*Transform    `protobuf:"bytes,2,rep,name=transform" json:"transform,omitempty"`
	Output           *OutputSettings `protobuf:"bytes,3,req,name=output" json:"output,omitempty"`
	Input            *InputSettings  `protobuf:"bytes,4,opt,name=input" json:"input,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ImagesTransformRequest) Reset()         { *m = ImagesTransformRequest{} }
func (m *ImagesTransformRequest) String() string { return proto.CompactTextString(m) }
func (*ImagesTransformRequest) ProtoMessage()    {}

func (m *ImagesTransformRequest) GetImage() *ImageData {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImagesTransformRequest) GetTransform() []*Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *ImagesTransformRequest) GetOutput() *OutputSettings {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *ImagesTransformRequest) GetInput() *InputSettings {
	if m != nil {
		return m.Input
	}
	return nil
}

type ImagesTransformResponse struct {
	Image            *ImageData `protobuf:"bytes,1,req,name=image" json:"image,omitempty"`
	SourceMetadata   *string    `protobuf:"bytes,2,opt,name=source_metadata" json:"source_metadata,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ImagesTransformResponse) Reset()         { *m = ImagesTransformResponse{} }
func (m *ImagesTransformResponse) String() string { return proto.CompactTextString(m) }
func (*ImagesTransformResponse) ProtoMessage()    {}

func (m *ImagesTransformResponse) GetImage() *ImageData {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImagesTransformResponse) GetSourceMetadata() string {
	if m != nil && m.SourceMetadata != nil {
		return *m.SourceMetadata
	}
	return ""
}

type CompositeImageOptions struct {
	SourceIndex      *int32                        `protobuf:"varint,1,req,name=source_index" json:"source_index,omitempty"`
	XOffset          *int32                        `protobuf:"varint,2,req,name=x_offset" json:"x_offset,omitempty"`
	YOffset          *int32                        `protobuf:"varint,3,req,name=y_offset" json:"y_offset,omitempty"`
	Opacity          *float32                      `protobuf:"fixed32,4,req,name=opacity" json:"opacity,omitempty"`
	Anchor           *CompositeImageOptions_ANCHOR `protobuf:"varint,5,req,name=anchor,enum=appengine.CompositeImageOptions_ANCHOR" json:"anchor,omitempty"`
	XXX_unrecognized []byte                        `json:"-"`
}

func (m *CompositeImageOptions) Reset()         { *m = CompositeImageOptions{} }
func (m *CompositeImageOptions) String() string { return proto.CompactTextString(m) }
func (*CompositeImageOptions) ProtoMessage()    {}

func (m *CompositeImageOptions) GetSourceIndex() int32 {
	if m != nil && m.SourceIndex != nil {
		return *m.SourceIndex
	}
	return 0
}

func (m *CompositeImageOptions) GetXOffset() int32 {
	if m != nil && m.XOffset != nil {
		return *m.XOffset
	}
	return 0
}

func (m *CompositeImageOptions) GetYOffset() int32 {
	if m != nil && m.YOffset != nil {
		return *m.YOffset
	}
	return 0
}

func (m *CompositeImageOptions) GetOpacity() float32 {
	if m != nil && m.Opacity != nil {
		return *m.Opacity
	}
	return 0
}

func (m *CompositeImageOptions) GetAnchor() CompositeImageOptions_ANCHOR {
	if m != nil && m.Anchor != nil {
		return *m.Anchor
	}
	return CompositeImageOptions_TOP_LEFT
}

type ImagesCanvas struct {
	Width            *int32          `protobuf:"varint,1,req,name=width" json:"width,omitempty"`
	Height           *int32          `protobuf:"varint,2,req,name=height" json:"height,omitempty"`
	Output           *OutputSettings `protobuf:"bytes,3,req,name=output" json:"output,omitempty"`
	Color            *int32          `protobuf:"varint,4,opt,name=color,def=-1" json:"color,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ImagesCanvas) Reset()         { *m = ImagesCanvas{} }
func (m *ImagesCanvas) String() string { return proto.CompactTextString(m) }
func (*ImagesCanvas) ProtoMessage()    {}

const Default_ImagesCanvas_Color int32 = -1

func (m *ImagesCanvas) GetWidth() int32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *ImagesCanvas) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *ImagesCanvas) GetOutput() *OutputSettings {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *ImagesCanvas) GetColor() int32 {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return Default_ImagesCanvas_Color
}

type ImagesCompositeRequest struct {
	Image            []*ImageData             `protobuf:"bytes,1,rep,name=image" json:"image,omitempty"`
	Options          []*CompositeImageOptions `protobuf:"bytes,2,rep,name=options" json:"options,omitempty"`
	Canvas           *ImagesCanvas            `protobuf:"bytes,3,req,name=canvas" json:"canvas,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ImagesCompositeRequest) Reset()         { *m = ImagesCompositeRequest{} }
func (m *ImagesCompositeRequest) String() string { return proto.CompactTextString(m) }
func (*ImagesCompositeRequest) ProtoMessage()    {}

func (m *ImagesCompositeRequest) GetImage() []*ImageData {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *ImagesCompositeRequest) GetOptions() []*CompositeImageOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *ImagesCompositeRequest) GetCanvas() *ImagesCanvas {
	if m != nil {
		return m.Canvas
	}
	return nil
}

type ImagesCompositeResponse struct {
	Image            *ImageData `protobuf:"bytes,1,req,name=image" json:"image,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ImagesCompositeResponse) Reset()         { *m = ImagesCompositeResponse{} }
func (m *ImagesCompositeResponse) String() string { return proto.CompactTextString(m) }
func (*ImagesCompositeResponse) ProtoMessage()    {}

func (m *ImagesCompositeResponse) GetImage() *ImageData {
	if m != nil {
		return m.Image
	}
	return nil
}

type ImagesHistogramRequest struct {
	Image            *ImageData `protobuf:"bytes,1,req,name=image" json:"image,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ImagesHistogramRequest) Reset()         { *m = ImagesHistogramRequest{} }
func (m *ImagesHistogramRequest) String() string { return proto.CompactTextString(m) }
func (*ImagesHistogramRequest) ProtoMessage()    {}

func (m *ImagesHistogramRequest) GetImage() *ImageData {
	if m != nil {
		return m.Image
	}
	return nil
}

type ImagesHistogram struct {
	Red              []int32 `protobuf:"varint,1,rep,name=red" json:"red,omitempty"`
	Green            []int32 `protobuf:"varint,2,rep,name=green" json:"green,omitempty"`
	Blue             []int32 `protobuf:"varint,3,rep,name=blue" json:"blue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ImagesHistogram) Reset()         { *m = ImagesHistogram{} }
func (m *ImagesHistogram) String() string { return proto.CompactTextString(m) }
func (*ImagesHistogram) ProtoMessage()    {}

func (m *ImagesHistogram) GetRed() []int32 {
	if m != nil {
		return m.Red
	}
	return nil
}

func (m *ImagesHistogram) GetGreen() []int32 {
	if m != nil {
		return m.Green
	}
	return nil
}

func (m *ImagesHistogram) GetBlue() []int32 {
	if m != nil {
		return m.Blue
	}
	return nil
}

type ImagesHistogramResponse struct {
	Histogram        *ImagesHistogram `protobuf:"bytes,1,req,name=histogram" json:"histogram,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *ImagesHistogramResponse) Reset()         { *m = ImagesHistogramResponse{} }
func (m *ImagesHistogramResponse) String() string { return proto.CompactTextString(m) }
func (*ImagesHistogramResponse) ProtoMessage()    {}

func (m *ImagesHistogramResponse) GetHistogram() *ImagesHistogram {
	if m != nil {
		return m.Histogram
	}
	return nil
}

type ImagesGetUrlBaseRequest struct {
	BlobKey          *string `protobuf:"bytes,1,req,name=blob_key" json:"blob_key,omitempty"`
	CreateSecureUrl  *bool   `protobuf:"varint,2,opt,name=create_secure_url,def=0" json:"create_secure_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ImagesGetUrlBaseRequest) Reset()         { *m = ImagesGetUrlBaseRequest{} }
func (m *ImagesGetUrlBaseRequest) String() string { return proto.CompactTextString(m) }
func (*ImagesGetUrlBaseRequest) ProtoMessage()    {}

const Default_ImagesGetUrlBaseRequest_CreateSecureUrl bool = false

func (m *ImagesGetUrlBaseRequest) GetBlobKey() string {
	if m != nil && m.BlobKey != nil {
		return *m.BlobKey
	}
	return ""
}

func (m *ImagesGetUrlBaseRequest) GetCreateSecureUrl() bool {
	if m != nil && m.CreateSecureUrl != nil {
		return *m.CreateSecureUrl
	}
	return Default_ImagesGetUrlBaseRequest_CreateSecureUrl
}

type ImagesGetUrlBaseResponse struct {
	Url              *string `protobuf:"bytes,1,req,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ImagesGetUrlBaseResponse) Reset()         { *m = ImagesGetUrlBaseResponse{} }
func (m *ImagesGetUrlBaseResponse) String() string { return proto.CompactTextString(m) }
func (*ImagesGetUrlBaseResponse) ProtoMessage()    {}

func (m *ImagesGetUrlBaseResponse) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

type ImagesDeleteUrlBaseRequest struct {
	BlobKey          *string `protobuf:"bytes,1,req,name=blob_key" json:"blob_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ImagesDeleteUrlBaseRequest) Reset()         { *m = ImagesDeleteUrlBaseRequest{} }
func (m *ImagesDeleteUrlBaseRequest) String() string { return proto.CompactTextString(m) }
func (*ImagesDeleteUrlBaseRequest) ProtoMessage()    {}

func (m *ImagesDeleteUrlBaseRequest) GetBlobKey() string {
	if m != nil && m.BlobKey != nil {
		return *m.BlobKey
	}
	return ""
}

type ImagesDeleteUrlBaseResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ImagesDeleteUrlBaseResponse) Reset()         { *m = ImagesDeleteUrlBaseResponse{} }
func (m *ImagesDeleteUrlBaseResponse) String() string { return proto.CompactTextString(m) }
func (*ImagesDeleteUrlBaseResponse) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("appengine.ImagesServiceError_ErrorCode", ImagesServiceError_ErrorCode_name, ImagesServiceError_ErrorCode_value)
	proto.RegisterEnum("appengine.ImagesServiceTransform_Type", ImagesServiceTransform_Type_name, ImagesServiceTransform_Type_value)
	proto.RegisterEnum("appengine.InputSettings_ORIENTATION_CORRECTION_TYPE", InputSettings_ORIENTATION_CORRECTION_TYPE_name, InputSettings_ORIENTATION_CORRECTION_TYPE_value)
	proto.RegisterEnum("appengine.OutputSettings_MIME_TYPE", OutputSettings_MIME_TYPE_name, OutputSettings_MIME_TYPE_value)
	proto.RegisterEnum("appengine.CompositeImageOptions_ANCHOR", CompositeImageOptions_ANCHOR_name, CompositeImageOptions_ANCHOR_value)
}
