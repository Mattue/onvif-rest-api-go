package onvifGoDefinitions

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"github.com/metaleap/go-xsd/types"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetServiceCapabilitiesResponse"`

	// The capabilities for the PTZ service is returned in the Capabilities element.
	Capabilities *ServiceCapabilities `xml:"Capabilities,omitempty" json:"Capabilities,omitempty"`
}

type GetNodes struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetNodes"`
}

type GetNodesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetNodesResponse"`

	// A list of the existing PTZ Nodes on the device.
	PTZNode []*PTZNode `xml:"PTZNode,omitempty" json:"PTZNode,omitempty"`
}

type GetNode struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetNode"`

	// Token of the requested PTZNode.
	NodeToken *ReferenceToken `xml:"NodeToken,omitempty" json:"NodeToken,omitempty"`
}

type GetNodeResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetNodeResponse"`

	// A requested PTZNode.
	PTZNode *PTZNode `xml:"PTZNode,omitempty" json:"PTZNode,omitempty"`
}

type GetConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetConfigurations"`
}

type GetConfigurationsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetConfigurationsResponse"`

	// A list of all existing PTZConfigurations on the device.
	PTZConfiguration []*PTZConfiguration `xml:"PTZConfiguration,omitempty" json:"PTZConfiguration,omitempty"`
}

type GetConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetConfiguration"`

	// Token of the requested PTZConfiguration.
	PTZConfigurationToken *ReferenceToken `xml:"PTZConfigurationToken,omitempty" json:"PTZConfigurationToken,omitempty"`
}

type GetConfigurationResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetConfigurationResponse"`

	// A requested PTZConfiguration.
	PTZConfiguration *PTZConfiguration `xml:"PTZConfiguration,omitempty" json:"PTZConfiguration,omitempty"`
}

type SetConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl SetConfiguration"`

	PTZConfiguration *PTZConfiguration `xml:"PTZConfiguration,omitempty" json:"PTZConfiguration,omitempty"`

	// Flag that makes configuration persistent. Example: User wants the configuration to exist after reboot.
	ForcePersistence bool `xml:"ForcePersistence,omitempty" json:"ForcePersistence,omitempty"`
}

type SetConfigurationResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl SetConfigurationResponse"`
}

type GetConfigurationOptions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetConfigurationOptions"`

	// Token of an existing configuration that the options are intended for.
	ConfigurationToken *ReferenceToken `xml:"ConfigurationToken,omitempty" json:"ConfigurationToken,omitempty"`
}

type GetConfigurationOptionsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetConfigurationOptionsResponse"`

	// The requested PTZ configuration options.
	PTZConfigurationOptions *PTZConfigurationOptions `xml:"PTZConfigurationOptions,omitempty" json:"PTZConfigurationOptions,omitempty"`
}

type SendAuxiliaryCommand struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl SendAuxiliaryCommand"`

	// A reference to the MediaProfile where the operation should take place.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// The Auxiliary request data.
	AuxiliaryData *AuxiliaryData `xml:"AuxiliaryData,omitempty" json:"AuxiliaryData,omitempty"`
}

type SendAuxiliaryCommandResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl SendAuxiliaryCommandResponse"`

	// The response contains the auxiliary response.
	AuxiliaryResponse *AuxiliaryData `xml:"AuxiliaryResponse,omitempty" json:"AuxiliaryResponse,omitempty"`
}

type GetPresets struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresets"`

	// A reference to the MediaProfile where the operation should take place.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`
}

type GetPresetsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresetsResponse"`

	// A list of presets which are available for the requested MediaProfile.
	Preset []*PTZPreset `xml:"Preset,omitempty" json:"Preset,omitempty"`
}

type SetPreset struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl SetPreset"`

	// A reference to the MediaProfile where the operation should take place.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// A requested preset name.
	PresetName string `xml:"PresetName,omitempty" json:"PresetName,omitempty"`

	// A requested preset token.
	PresetToken *ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty"`
}

type SetPresetResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl SetPresetResponse"`

	// A token to the Preset which has been set.
	PresetToken *ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty"`
}

type RemovePreset struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl RemovePreset"`

	// A reference to the MediaProfile where the operation should take place.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// A requested preset token.
	PresetToken *ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty"`
}

type RemovePresetResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl RemovePresetResponse"`
}

type GotoPreset struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GotoPreset"`

	// A reference to the MediaProfile where the operation should take place.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// A requested preset token.
	PresetToken *ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty"`

	// A requested speed.The speed parameter can only be specified when Speed Spaces are available for the PTZ Node.
	Speed *PTZSpeed `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type GotoPresetResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GotoPresetResponse"`
}

type GetStatus struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetStatus"`

	// A reference to the MediaProfile where the PTZStatus should be requested.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`
}

type GetStatusResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetStatusResponse"`

	// The PTZStatus for the requested MediaProfile.
	PTZStatus *PTZStatus `xml:"PTZStatus,omitempty" json:"PTZStatus,omitempty"`
}

type GotoHomePosition struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GotoHomePosition"`

	// A reference to the MediaProfile where the operation should take place.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// A requested speed.The speed parameter can only be specified when Speed Spaces are available for the PTZ Node.
	Speed *PTZSpeed `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type GotoHomePositionResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GotoHomePositionResponse"`
}

type SetHomePosition struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl SetHomePosition"`

	// A reference to the MediaProfile where the home position should be set.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`
}

type SetHomePositionResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl SetHomePositionResponse"`
}

type ContinuousMove struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl ContinuousMove"`

	// A reference to the MediaProfile.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// A Velocity vector specifying the velocity of pan, tilt and zoom.
	Velocity *PTZSpeed `xml:"Velocity,omitempty" json:"Velocity,omitempty"`

	// An optional Timeout parameter.
	Timeout *time.Duration `xml:"Timeout,omitempty" json:"Timeout,omitempty"`
}

type ContinuousMoveResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl ContinuousMoveResponse"`
}

type RelativeMove struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl RelativeMove"`

	// A reference to the MediaProfile.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// A positional Translation relative to the current position
	Translation *PTZVector `xml:"Translation,omitempty" json:"Translation,omitempty"`

	// An optional Speed parameter.
	Speed *PTZSpeed `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type RelativeMoveResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl RelativeMoveResponse"`
}

type AbsoluteMove struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl AbsoluteMove"`

	// A reference to the MediaProfile.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// A Position vector specifying the absolute target position.
	Position *PTZVector `xml:"Position,omitempty" json:"Position,omitempty"`

	// An optional Speed.
	Speed *PTZSpeed `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type AbsoluteMoveResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl AbsoluteMoveResponse"`
}

type GeoMove struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GeoMove"`

	// A reference to the MediaProfile.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// The geolocation of the target position.
	Target *GeoLocation `xml:"Target,omitempty" json:"Target,omitempty"`

	// An optional Speed.
	Speed *PTZSpeed `xml:"Speed,omitempty" json:"Speed,omitempty"`

	// An optional indication of the height of the target/area.
	AreaHeight float32 `xml:"AreaHeight,omitempty" json:"AreaHeight,omitempty"`

	// An optional indication of the width of the target/area.
	AreaWidth float32 `xml:"AreaWidth,omitempty" json:"AreaWidth,omitempty"`
}

type GeoMoveResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GeoMoveResponse"`
}

type Stop struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl Stop"`

	// A reference to the MediaProfile that indicate what should be stopped.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// Set true when we want to stop ongoing pan and tilt movements.If PanTilt arguments are not present, this command stops these movements.
	PanTilt bool `xml:"PanTilt,omitempty" json:"PanTilt,omitempty"`

	// Set true when we want to stop ongoing zoom movement.If Zoom arguments are not present, this command stops ongoing zoom movement.
	Zoom bool `xml:"Zoom,omitempty" json:"Zoom,omitempty"`
}

type StopResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl StopResponse"`
}

type GetPresetTours struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresetTours"`

	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`
}

type GetPresetToursResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresetToursResponse"`

	PresetTour []*PresetTour `xml:"PresetTour,omitempty" json:"PresetTour,omitempty"`
}

type GetPresetTour struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresetTour"`

	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	PresetTourToken *ReferenceToken `xml:"PresetTourToken,omitempty" json:"PresetTourToken,omitempty"`
}

type GetPresetTourResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresetTourResponse"`

	PresetTour *PresetTour `xml:"PresetTour,omitempty" json:"PresetTour,omitempty"`
}

type GetPresetTourOptions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresetTourOptions"`

	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	PresetTourToken *ReferenceToken `xml:"PresetTourToken,omitempty" json:"PresetTourToken,omitempty"`
}

type GetPresetTourOptionsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetPresetTourOptionsResponse"`

	Options *PTZPresetTourOptions `xml:"Options,omitempty" json:"Options,omitempty"`
}

type CreatePresetTour struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl CreatePresetTour"`

	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`
}

type CreatePresetTourResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl CreatePresetTourResponse"`

	PresetTourToken *ReferenceToken `xml:"PresetTourToken,omitempty" json:"PresetTourToken,omitempty"`
}

type ModifyPresetTour struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl ModifyPresetTour"`

	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	PresetTour *PresetTour `xml:"PresetTour,omitempty" json:"PresetTour,omitempty"`
}

type ModifyPresetTourResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl ModifyPresetTourResponse"`
}

type OperatePresetTour struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl OperatePresetTour"`

	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	PresetTourToken *ReferenceToken `xml:"PresetTourToken,omitempty" json:"PresetTourToken,omitempty"`

	Operation *PTZPresetTourOperation `xml:"Operation,omitempty" json:"Operation,omitempty"`
}

type OperatePresetTourResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl OperatePresetTourResponse"`
}

type RemovePresetTour struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl RemovePresetTour"`

	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	PresetTourToken *ReferenceToken `xml:"PresetTourToken,omitempty" json:"PresetTourToken,omitempty"`
}

type RemovePresetTourResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl RemovePresetTourResponse"`
}

type GetCompatibleConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetCompatibleConfigurations"`

	// Contains the token of an existing media profile the configurations shall be compatible with.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`
}

type GetCompatibleConfigurationsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl GetCompatibleConfigurationsResponse"`

	// A list of all existing PTZConfigurations on the NVT that is suitable to be added to the addressed media profile.
	PTZConfiguration []*PTZConfiguration `xml:"PTZConfiguration,omitempty" json:"PTZConfiguration,omitempty"`
}

type MoveAndStartTracking struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl MoveAndStartTracking"`

	// A reference to the MediaProfile where the operation should take place.
	ProfileToken *ReferenceToken `xml:"ProfileToken,omitempty" json:"ProfileToken,omitempty"`

	// A preset token.
	PresetToken *ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty"`

	// The geolocation of the target position.
	GeoLocation *GeoLocation `xml:"GeoLocation,omitempty" json:"GeoLocation,omitempty"`

	// A Position vector specifying the absolute target position.
	TargetPosition *PTZVector `xml:"TargetPosition,omitempty" json:"TargetPosition,omitempty"`

	// Speed vector specifying the velocity of pan, tilt and zoom.
	Speed *PTZSpeed `xml:"Speed,omitempty" json:"Speed,omitempty"`

	// Object ID of the object to track.
	ObjectID int32 `xml:"ObjectID,omitempty" json:"ObjectID,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type MoveAndStartTrackingResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver20/ptz/wsdl MoveAndStartTrackingResponse"`
}

type ServiceCapabilities struct {
	Items []string `xml:",any" json:"items,omitempty"`

	// Indicates whether or not EFlip is supported.

	EFlip bool `xml:"EFlip,attr,omitempty" json:"EFlip,omitempty"`

	// Indicates whether or not reversing of PT control direction is supported.

	Reverse bool `xml:"Reverse,attr,omitempty" json:"Reverse,omitempty"`

	// Indicates support for the GetCompatibleConfigurations command.

	GetCompatibleConfigurations bool `xml:"GetCompatibleConfigurations,attr,omitempty" json:"GetCompatibleConfigurations,omitempty"`

	// Indicates that the PTZStatus includes MoveStatus information.

	MoveStatus bool `xml:"MoveStatus,attr,omitempty" json:"MoveStatus,omitempty"`

	// Indicates that the PTZStatus includes Position information.

	StatusPosition bool `xml:"StatusPosition,attr,omitempty" json:"StatusPosition,omitempty"`

	// Indication of the methods of MoveAndTrack that are supported, acceptable values are defined in tt:MoveAndTrackMethod.

	MoveAndTrack *StringList `xml:"MoveAndTrack,attr,omitempty" json:"MoveAndTrack,omitempty"`
}

type Base64Binary struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/05/xmlmime base64Binary"`

	Value []byte `xml:",chardata" json:"-,"`

	ContentType string `xml:"contentType,attr,omitempty" json:"contentType,omitempty"`
}

type HexBinary struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/05/xmlmime hexBinary"`

	Value []byte `xml:",chardata" json:"-,"`

	ContentType string `xml:"contentType,attr,omitempty" json:"contentType,omitempty"`
}

//type FaultcodeEnum *QName

//const (
//	FaultcodeEnumTnsDataEncodingUnknown FaultcodeEnum = "tns:DataEncodingUnknown"
//
//	FaultcodeEnumTnsMustUnderstand FaultcodeEnum = "tns:MustUnderstand"
//
//	FaultcodeEnumTnsReceiver FaultcodeEnum = "tns:Receiver"
//
//	FaultcodeEnumTnsSender FaultcodeEnum = "tns:Sender"
//
//	FaultcodeEnumTnsVersionMismatch FaultcodeEnum = "tns:VersionMismatch"
//)

type NotUnderstood NotUnderstoodType

type Upgrade UpgradeType

type Envelope struct {
	Header *Header `xml:"Header,omitempty" json:"Header,omitempty"`

	Body *Body `xml:"Body,omitempty" json:"Body,omitempty"`
}

type Header struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Body struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Fault struct {
	Code *Faultcode `xml:"Code,omitempty" json:"Code,omitempty"`

	Reason *Faultreason `xml:"Reason,omitempty" json:"Reason,omitempty"`

	Node AnyURI `xml:"Node,omitempty" json:"Node,omitempty"`

	Role AnyURI `xml:"Role,omitempty" json:"Role,omitempty"`

	Detail *Detail `xml:"Detail,omitempty" json:"Detail,omitempty"`
}

type Faultreason struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope faultreason"`

	Text []*Reasontext `xml:"Text,omitempty" json:"Text,omitempty"`
}

type Reasontext string

type Faultcode struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope faultcode"`

	//Value *FaultcodeEnum `xml:"Value,omitempty" json:"Value,omitempty"`

	Subcode *Subcode `xml:"Subcode,omitempty" json:"Subcode,omitempty"`
}

type Subcode struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope subcode"`

	//Value *QName `xml:"Value,omitempty" json:"Value,omitempty"`

	Subcode *Subcode `xml:"Subcode,omitempty" json:"Subcode,omitempty"`
}

type Detail struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope detail"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type NotUnderstoodType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope NotUnderstood"`

	//Qname *QName `xml:"qname,attr,omitempty" json:"qname,omitempty"`
}

type SupportedEnvType struct {
	//Qname *QName `xml:"qname,attr,omitempty" json:"qname,omitempty"`
}

type UpgradeType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Upgrade"`

	SupportedEnvelope []*SupportedEnvType `xml:"SupportedEnvelope,omitempty" json:"SupportedEnvelope,omitempty"`
}

type RelationshipTypeOpenEnum string

type RelationshipType AnyURI

const (
	RelationshipTypeHttpwww_w3_org200508addressingreply RelationshipType = "http://www.w3.org/2005/08/addressing/reply"
)

type FaultCodesOpenEnumType string

//type FaultCodesType *QName

//const (
//	FaultCodesTypeTnsInvalidAddressingHeader FaultCodesType = "tns:InvalidAddressingHeader"
//
//	FaultCodesTypeTnsInvalidAddress FaultCodesType = "tns:InvalidAddress"
//
//	FaultCodesTypeTnsInvalidEPR FaultCodesType = "tns:InvalidEPR"
//
//	FaultCodesTypeTnsInvalidCardinality FaultCodesType = "tns:InvalidCardinality"
//
//	FaultCodesTypeTnsMissingAddressInEPR FaultCodesType = "tns:MissingAddressInEPR"
//
//	FaultCodesTypeTnsDuplicateMessageID FaultCodesType = "tns:DuplicateMessageID"
//
//	FaultCodesTypeTnsActionMismatch FaultCodesType = "tns:ActionMismatch"
//
//	FaultCodesTypeTnsMessageAddressingHeaderRequired FaultCodesType = "tns:MessageAddressingHeaderRequired"
//
//	FaultCodesTypeTnsDestinationUnreachable FaultCodesType = "tns:DestinationUnreachable"
//
//	FaultCodesTypeTnsActionNotSupported FaultCodesType = "tns:ActionNotSupported"
//
//	FaultCodesTypeTnsEndpointUnavailable FaultCodesType = "tns:EndpointUnavailable"
//)

type EndpointReference EndpointReferenceType

type Metadata MetadataType

type MessageID AttributedURIType

type RelatesTo RelatesToType

type ReplyTo EndpointReferenceType

type From EndpointReferenceType

type FaultTo EndpointReferenceType

type To AttributedURIType

type Action AttributedURIType

type RetryAfter AttributedUnsignedLongType

type ProblemHeaderQName AttributedQNameType

type ProblemHeader AttributedAnyType

type ProblemIRI AttributedURIType

type ProblemAction ProblemActionType

type EndpointReferenceType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing EndpointReference"`

	Address *AttributedURIType `xml:"Address,omitempty" json:"Address,omitempty"`

	ReferenceParameters *ReferenceParametersType `xml:"ReferenceParameters,omitempty" json:"ReferenceParameters,omitempty"`

	Metadata *Metadata `xml:"Metadata,omitempty" json:"Metadata,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ReferenceParametersType struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type MetadataType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing Metadata"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RelatesToType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing RelatesTo"`

	Value AnyURI `xml:",chardata" json:"-,"`

	RelationshipType *RelationshipTypeOpenEnum `xml:"RelationshipType,attr,omitempty" json:"RelationshipType,omitempty"`
}

type AttributedURIType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing MessageID"`

	Value AnyURI `xml:",chardata" json:"-,"`
}

type AttributedUnsignedLongType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing RetryAfter"`

	Value uint64 `xml:",chardata" json:"-,"`
}

type AttributedQNameType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ProblemHeaderQName"`

	//Value *QName `xml:",chardata" json:"-,"`
}

type AttributedAnyType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ProblemHeader"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ProblemActionType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ProblemAction"`

	Action *Action `xml:"Action,omitempty" json:"Action,omitempty"`

	SoapAction AnyURI `xml:"SoapAction,omitempty" json:"SoapAction,omitempty"`
}

type BaseFault BaseFaultType

type BaseFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsrf/bf-2 BaseFault"`

	Timestamp soap.XSDDateTime `xml:"Timestamp,omitempty" json:"Timestamp,omitempty"`

	Originator *EndpointReferenceType `xml:"Originator,omitempty" json:"Originator,omitempty"`

	ErrorCode struct {
		AnyType

		Dialect AnyURI `xml:"dialect,attr,omitempty" json:"dialect,omitempty"`
	} `xml:"ErrorCode,omitempty" json:"ErrorCode,omitempty"`

	Description []struct {
		Value string `xml:",chardata" json:"-,"`

		string `xml:",attr,omitempty" json:",omitempty"`
	} `xml:"Description,omitempty" json:"Description,omitempty"`

	FaultCause struct {
	} `xml:"FaultCause,omitempty" json:"FaultCause,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type FullTopicExpression string

type ConcreteTopicExpression string

//type SimpleTopicExpression *QName

type TopicNamespace TopicNamespaceType

type TopicSet TopicSetType

type Documentation struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ExtensibleDocumented struct {
	Documentation *Documentation `xml:"documentation,omitempty" json:"documentation,omitempty"`
}

//type QueryExpressionType struct {
//	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/t-1 ProducerProperties"`
//
//	Items []string `xml:",any" json:"items,omitempty"`
//
//	Dialect AnyURI `xml:"Dialect,attr,omitempty" json:"Dialect,omitempty"`
//}

type TopicNamespaceType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/t-1 TopicNamespace"`

	*ExtensibleDocumented

	Topic []struct {
		*TopicType

		Parent *ConcreteTopicExpression `xml:"parent,attr,omitempty" json:"parent,omitempty"`
	} `xml:"Topic,omitempty" json:"Topic,omitempty"`

	Name NCName `xml:"name,attr,omitempty" json:"name,omitempty"`

	TargetNamespace AnyURI `xml:"targetNamespace,attr,omitempty" json:"targetNamespace,omitempty"`

	Final bool `xml:"final,attr,omitempty" json:"final,omitempty"`
}

type TopicType struct {
	*ExtensibleDocumented

	//MessagePattern *QueryExpressionType `xml:"MessagePattern,omitempty" json:"MessagePattern,omitempty"`

	Topic []*TopicType `xml:"Topic,omitempty" json:"Topic,omitempty"`

	Name NCName `xml:"name,attr,omitempty" json:"name,omitempty"`

	MessageTypes string `xml:"messageTypes,attr,omitempty" json:"messageTypes,omitempty"`

	Final bool `xml:"final,attr,omitempty" json:"final,omitempty"`
}

type TopicSetType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/t-1 TopicSet"`

	*ExtensibleDocumented
}

type AbsoluteOrRelativeTimeType string

type TopicExpression TopicExpressionType

type FixedTopicSet bool

type TopicExpressionDialect AnyURI

type NotificationProducerRP struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NotificationProducerRP"`

	TopicExpression []*TopicExpression `xml:"TopicExpression,omitempty" json:"TopicExpression,omitempty"`

	FixedTopicSet *FixedTopicSet `xml:"FixedTopicSet,omitempty" json:"FixedTopicSet,omitempty"`

	TopicExpressionDialect []*TopicExpressionDialect `xml:"TopicExpressionDialect,omitempty" json:"TopicExpressionDialect,omitempty"`

	TopicSet *TopicSet `xml:"TopicSet,omitempty" json:"TopicSet,omitempty"`
}

type ConsumerReference EndpointReferenceType

type Filter FilterType

type SubscriptionPolicy SubscriptionPolicyType

type CreationTime soap.XSDDateTime

type SubscriptionManagerRP struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscriptionManagerRP"`

	ConsumerReference *ConsumerReference `xml:"ConsumerReference,omitempty" json:"ConsumerReference,omitempty"`

	Filter *Filter `xml:"Filter,omitempty" json:"Filter,omitempty"`

	SubscriptionPolicy *SubscriptionPolicy `xml:"SubscriptionPolicy,omitempty" json:"SubscriptionPolicy,omitempty"`

	CreationTime *CreationTime `xml:"CreationTime,omitempty" json:"CreationTime,omitempty"`
}

type SubscriptionReference EndpointReferenceType

type Topic TopicExpressionType

type ProducerReference EndpointReferenceType

type NotificationMessage NotificationMessageHolderType

type Notify struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Notify"`

	NotificationMessage []*NotificationMessage `xml:"NotificationMessage,omitempty" json:"NotificationMessage,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type CurrentTime soap.XSDDateTime

type TerminationTime soap.XSDDateTime

//type ProducerProperties QueryExpressionType

//type MessageContent QueryExpressionType

type UseRaw struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UseRaw"`
}

type Subscribe struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Subscribe"`

	ConsumerReference *EndpointReferenceType `xml:"ConsumerReference,omitempty" json:"ConsumerReference,omitempty"`

	Filter *FilterType `xml:"Filter,omitempty" json:"Filter,omitempty"`

	InitialTerminationTime *AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime,omitempty" json:"InitialTerminationTime,omitempty"`

	SubscriptionPolicy struct {
	} `xml:"SubscriptionPolicy,omitempty" json:"SubscriptionPolicy,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type SubscribeResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscribeResponse"`

	SubscriptionReference *EndpointReferenceType `xml:"SubscriptionReference,omitempty" json:"SubscriptionReference,omitempty"`

	CurrentTime *CurrentTime `xml:"CurrentTime,omitempty" json:"CurrentTime,omitempty"`

	TerminationTime *TerminationTime `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type GetCurrentMessage struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 GetCurrentMessage"`

	Topic *TopicExpressionType `xml:"Topic,omitempty" json:"Topic,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type GetCurrentMessageResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 GetCurrentMessageResponse"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type SubscribeCreationFailedFault SubscribeCreationFailedFaultType

type InvalidFilterFault InvalidFilterFaultType

type TopicExpressionDialectUnknownFault TopicExpressionDialectUnknownFaultType

type InvalidTopicExpressionFault InvalidTopicExpressionFaultType

type TopicNotSupportedFault TopicNotSupportedFaultType

type MultipleTopicsSpecifiedFault MultipleTopicsSpecifiedFaultType

type InvalidProducerPropertiesExpressionFault InvalidProducerPropertiesExpressionFaultType

type InvalidMessageContentExpressionFault InvalidMessageContentExpressionFaultType

type UnrecognizedPolicyRequestFault UnrecognizedPolicyRequestFaultType

type UnsupportedPolicyRequestFault UnsupportedPolicyRequestFaultType

type NotifyMessageNotSupportedFault NotifyMessageNotSupportedFaultType

type UnacceptableInitialTerminationTimeFault UnacceptableInitialTerminationTimeFaultType

type NoCurrentMessageOnTopicFault NoCurrentMessageOnTopicFaultType

type GetMessages struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 GetMessages"`

	//MaximumNumber *NonNegativeInteger `xml:"MaximumNumber,omitempty" json:"MaximumNumber,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type GetMessagesResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 GetMessagesResponse"`

	NotificationMessage []*NotificationMessage `xml:"NotificationMessage,omitempty" json:"NotificationMessage,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type DestroyPullPoint struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 DestroyPullPoint"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type DestroyPullPointResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 DestroyPullPointResponse"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type UnableToGetMessagesFault UnableToGetMessagesFaultType

type UnableToDestroyPullPointFault UnableToDestroyPullPointFaultType

type CreatePullPoint struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 CreatePullPoint"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type CreatePullPointResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 CreatePullPointResponse"`

	PullPoint *EndpointReferenceType `xml:"PullPoint,omitempty" json:"PullPoint,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type UnableToCreatePullPointFault UnableToCreatePullPointFaultType

type Renew struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Renew"`

	TerminationTime *AbsoluteOrRelativeTimeType `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RenewResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 RenewResponse"`

	TerminationTime *TerminationTime `xml:"TerminationTime,omitempty" json:"TerminationTime,omitempty"`

	CurrentTime *CurrentTime `xml:"CurrentTime,omitempty" json:"CurrentTime,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type UnacceptableTerminationTimeFault UnacceptableTerminationTimeFaultType

type Unsubscribe struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Unsubscribe"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type UnsubscribeResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnsubscribeResponse"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type UnableToDestroySubscriptionFault UnableToDestroySubscriptionFaultType

type PauseSubscription struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 PauseSubscription"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PauseSubscriptionResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 PauseSubscriptionResponse"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ResumeSubscription struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 ResumeSubscription"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ResumeSubscriptionResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 ResumeSubscriptionResponse"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PauseFailedFault PauseFailedFaultType

type ResumeFailedFault ResumeFailedFaultType

//type QueryExpressionType struct {
//	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 ProducerProperties"`
//
//	Items []string `xml:",any" json:"items,omitempty"`
//
//	Dialect AnyURI `xml:"Dialect,attr,omitempty" json:"Dialect,omitempty"`
//}

type TopicExpressionType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 TopicExpression"`

	Items []string `xml:",any" json:"items,omitempty"`

	Dialect AnyURI `xml:"Dialect,attr,omitempty" json:"Dialect,omitempty"`
}

type FilterType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Filter"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type SubscriptionPolicyType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscriptionPolicy"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type NotificationMessageHolderType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NotificationMessage"`

	SubscriptionReference *SubscriptionReference `xml:"SubscriptionReference,omitempty" json:"SubscriptionReference,omitempty"`

	Topic *Topic `xml:"Topic,omitempty" json:"Topic,omitempty"`

	ProducerReference *ProducerReference `xml:"ProducerReference,omitempty" json:"ProducerReference,omitempty"`

	Message struct {
	} `xml:"Message,omitempty" json:"Message,omitempty"`
}

type SubscribeCreationFailedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscribeCreationFailedFault"`

	*BaseFaultType
}

type InvalidFilterFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidFilterFault"`

	*BaseFaultType

	//UnknownFilter []*QName `xml:"UnknownFilter,omitempty" json:"UnknownFilter,omitempty"`
}

type TopicExpressionDialectUnknownFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 TopicExpressionDialectUnknownFault"`

	*BaseFaultType
}

type InvalidTopicExpressionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidTopicExpressionFault"`

	*BaseFaultType
}

type TopicNotSupportedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 TopicNotSupportedFault"`

	*BaseFaultType
}

type MultipleTopicsSpecifiedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 MultipleTopicsSpecifiedFault"`

	*BaseFaultType
}

type InvalidProducerPropertiesExpressionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidProducerPropertiesExpressionFault"`

	*BaseFaultType
}

type InvalidMessageContentExpressionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidMessageContentExpressionFault"`

	*BaseFaultType
}

type UnrecognizedPolicyRequestFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnrecognizedPolicyRequestFault"`

	*BaseFaultType

	//UnrecognizedPolicy []*QName `xml:"UnrecognizedPolicy,omitempty" json:"UnrecognizedPolicy,omitempty"`
}

type UnsupportedPolicyRequestFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnsupportedPolicyRequestFault"`

	*BaseFaultType

	//UnsupportedPolicy []*QName `xml:"UnsupportedPolicy,omitempty" json:"UnsupportedPolicy,omitempty"`
}

type NotifyMessageNotSupportedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NotifyMessageNotSupportedFault"`

	*BaseFaultType
}

type UnacceptableInitialTerminationTimeFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnacceptableInitialTerminationTimeFault"`

	*BaseFaultType

	MinimumTime soap.XSDDateTime `xml:"MinimumTime,omitempty" json:"MinimumTime,omitempty"`

	MaximumTime soap.XSDDateTime `xml:"MaximumTime,omitempty" json:"MaximumTime,omitempty"`
}

type NoCurrentMessageOnTopicFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NoCurrentMessageOnTopicFault"`

	*BaseFaultType
}

type UnableToGetMessagesFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToGetMessagesFault"`

	*BaseFaultType
}

type UnableToDestroyPullPointFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToDestroyPullPointFault"`

	*BaseFaultType
}

type UnableToCreatePullPointFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToCreatePullPointFault"`

	*BaseFaultType
}

type UnacceptableTerminationTimeFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnacceptableTerminationTimeFault"`

	*BaseFaultType

	MinimumTime soap.XSDDateTime `xml:"MinimumTime,omitempty" json:"MinimumTime,omitempty"`

	MaximumTime soap.XSDDateTime `xml:"MaximumTime,omitempty" json:"MaximumTime,omitempty"`
}

type UnableToDestroySubscriptionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToDestroySubscriptionFault"`

	*BaseFaultType
}

type PauseFailedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 PauseFailedFault"`

	*BaseFaultType
}

type ResumeFailedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 ResumeFailedFault"`

	*BaseFaultType
}

type Include struct {
	Items []string `xml:",any" json:"items,omitempty"`

	Href AnyURI `xml:"href,attr,omitempty" json:"href,omitempty"`
}

// Unique identifier for a physical or logical resource.
// Tokens should be assigned such that they are unique within a device. Tokens must be at least unique within its class.
// Length up to 64 characters. Token may be extended by intermediate terminal with adding prefix to make it global unique.
// The length should be within 36 characters for generating at local device. See "Remote Token" section in Resource Query specification.

type ReferenceToken string

type MoveStatus string

const (
	MoveStatusIDLE MoveStatus = "IDLE"

	MoveStatusMOVING MoveStatus = "MOVING"

	MoveStatusUNKNOWN MoveStatus = "UNKNOWN"
)

type Entity string

const (
	EntityDevice Entity = "Device"

	EntityVideoSource Entity = "VideoSource"

	EntityAudioSource Entity = "AudioSource"
)

type IntRange struct {
	Min int32 `xml:"Min,omitempty" json:"Min,omitempty"`

	Max int32 `xml:"Max,omitempty" json:"Max,omitempty"`
}

type Vector2D struct {
	X float32 `xml:"x,attr,omitempty" json:"x,omitempty"`

	Y float32 `xml:"y,attr,omitempty" json:"y,omitempty"`

	//
	// Pan/tilt coordinate space selector. The following options are defined:
	//
	//

	Space AnyURI `xml:"space,attr,omitempty" json:"space,omitempty"`
}

type Vector1D struct {
	X float32 `xml:"x,attr,omitempty" json:"x,omitempty"`

	//
	// Zoom coordinate space selector. The following options are defined:
	//
	//

	Space AnyURI `xml:"space,attr,omitempty" json:"space,omitempty"`
}

type PTZVector struct {

	// Pan and tilt position. The x component corresponds to pan and the y component to tilt.
	PanTilt *Vector2D `xml:"PanTilt,omitempty" json:"PanTilt,omitempty"`

	// A zoom position.
	Zoom *Vector1D `xml:"Zoom,omitempty" json:"Zoom,omitempty"`
}

type PTZStatus struct {

	// Specifies the absolute position of the PTZ unit together with the Space references. The default absolute spaces of the corresponding PTZ configuration MUST be referenced within the Position element.
	Position *PTZVector `xml:"Position,omitempty" json:"Position,omitempty"`

	// Indicates if the Pan/Tilt/Zoom device unit is currently moving, idle or in an unknown state.
	MoveStatus *PTZMoveStatus `xml:"MoveStatus,omitempty" json:"MoveStatus,omitempty"`

	// States a current PTZ error.
	Error string `xml:"Error,omitempty" json:"Error,omitempty"`

	// Specifies the UTC time when this status was generated.
	UtcTime soap.XSDDateTime `xml:"UtcTime,omitempty" json:"UtcTime,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZMoveStatus struct {
	PanTilt *MoveStatus `xml:"PanTilt,omitempty" json:"PanTilt,omitempty"`

	Zoom *MoveStatus `xml:"Zoom,omitempty" json:"Zoom,omitempty"`
}

type Vector struct {
	X float32 `xml:"x,attr,omitempty" json:"x,omitempty"`

	Y float32 `xml:"y,attr,omitempty" json:"y,omitempty"`
}

type Rectangle struct {
	Bottom float32 `xml:"bottom,attr,omitempty" json:"bottom,omitempty"`

	Top float32 `xml:"top,attr,omitempty" json:"top,omitempty"`

	Right float32 `xml:"right,attr,omitempty" json:"right,omitempty"`

	Left float32 `xml:"left,attr,omitempty" json:"left,omitempty"`
}

type Polygon struct {
	Point []*Vector `xml:"Point,omitempty" json:"Point,omitempty"`
}

type Color struct {
	X float32 `xml:"X,attr,omitempty" json:"X,omitempty"`

	Y float32 `xml:"Y,attr,omitempty" json:"Y,omitempty"`

	Z float32 `xml:"Z,attr,omitempty" json:"Z,omitempty"`

	//
	// Acceptable values:
	//
	// If the Colorspace attribute is absent and not defined on higher level, YCbCr is implied.
	//
	// Deprecated values:
	//
	//

	Colorspace AnyURI `xml:"Colorspace,attr,omitempty" json:"Colorspace,omitempty"`

	// Likelihood that the color is correct.

	Likelihood float32 `xml:"Likelihood,attr,omitempty" json:"Likelihood,omitempty"`
}

type ColorCovariance struct {
	XX float32 `xml:"XX,attr,omitempty" json:"XX,omitempty"`

	YY float32 `xml:"YY,attr,omitempty" json:"YY,omitempty"`

	ZZ float32 `xml:"ZZ,attr,omitempty" json:"ZZ,omitempty"`

	XY float32 `xml:"XY,attr,omitempty" json:"XY,omitempty"`

	XZ float32 `xml:"XZ,attr,omitempty" json:"XZ,omitempty"`

	YZ float32 `xml:"YZ,attr,omitempty" json:"YZ,omitempty"`

	// Acceptable values are the same as in tt:Color.

	Colorspace AnyURI `xml:"Colorspace,attr,omitempty" json:"Colorspace,omitempty"`
}

type ColorDescriptor struct {
	ColorCluster []struct {
		Color *Color `xml:"Color,omitempty" json:"Color,omitempty"`

		Weight float32 `xml:"Weight,omitempty" json:"Weight,omitempty"`

		Covariance *ColorCovariance `xml:"Covariance,omitempty" json:"Covariance,omitempty"`
	} `xml:"ColorCluster,omitempty" json:"ColorCluster,omitempty"`

	Extension AnyType `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Transformation struct {
	Translate *Vector `xml:"Translate,omitempty" json:"Translate,omitempty"`

	Scale *Vector `xml:"Scale,omitempty" json:"Scale,omitempty"`

	Extension *TransformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type TransformationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type GeoLocation struct {
	Items []string `xml:",any" json:"items,omitempty"`

	// East west location as angle.

	Lon float64 `xml:"lon,attr,omitempty" json:"lon,omitempty"`

	// North south location as angle.

	Lat float64 `xml:"lat,attr,omitempty" json:"lat,omitempty"`

	// Hight in meters above sea level.

	Elevation float32 `xml:"elevation,attr,omitempty" json:"elevation,omitempty"`
}

type GeoOrientation struct {
	Items []string `xml:",any" json:"items,omitempty"`

	// Rotation around the x axis.

	Roll float32 `xml:"roll,attr,omitempty" json:"roll,omitempty"`

	// Rotation around the y axis.

	Pitch float32 `xml:"pitch,attr,omitempty" json:"pitch,omitempty"`

	// Rotation around the z axis.

	Yaw float32 `xml:"yaw,attr,omitempty" json:"yaw,omitempty"`
}

type LocalLocation struct {
	Items []string `xml:",any" json:"items,omitempty"`

	// East west location as angle.

	X float32 `xml:"x,attr,omitempty" json:"x,omitempty"`

	// North south location as angle.

	Y float32 `xml:"y,attr,omitempty" json:"y,omitempty"`

	// Offset in meters from the sea level.

	Z float32 `xml:"z,attr,omitempty" json:"z,omitempty"`
}

type LocalOrientation struct {
	Items []string `xml:",any" json:"items,omitempty"`

	// Rotation around the y axis.

	Pan float32 `xml:"pan,attr,omitempty" json:"pan,omitempty"`

	// Rotation around the z axis.

	Tilt float32 `xml:"tilt,attr,omitempty" json:"tilt,omitempty"`

	// Rotation around the x axis.

	Roll float32 `xml:"roll,attr,omitempty" json:"roll,omitempty"`
}

type SphericalCoordinate struct {
	Items []string `xml:",any" json:"items,omitempty"`

	// Distance in meters to the object.

	Distance float32 `xml:"Distance,attr,omitempty" json:"Distance,omitempty"`

	// Elevation angle in the range -90 to 90 degrees, where 0 is in level with the x-y plane.

	ElevationAngle float32 `xml:"ElevationAngle,attr,omitempty" json:"ElevationAngle,omitempty"`

	// Azimuth angle in the range -180 to 180 degrees counter clockwise, where 0 is rightwards.

	AzimuthAngle float32 `xml:"AzimuthAngle,attr,omitempty" json:"AzimuthAngle,omitempty"`
}

type LocationEntity struct {

	// Location on earth.
	GeoLocation *GeoLocation `xml:"GeoLocation,omitempty" json:"GeoLocation,omitempty"`

	// Orientation relative to earth.
	GeoOrientation *GeoOrientation `xml:"GeoOrientation,omitempty" json:"GeoOrientation,omitempty"`

	// Indoor location offset.
	LocalLocation *LocalLocation `xml:"LocalLocation,omitempty" json:"LocalLocation,omitempty"`

	// Indoor orientation offset.
	LocalOrientation *LocalOrientation `xml:"LocalOrientation,omitempty" json:"LocalOrientation,omitempty"`

	// Entity type the entry refers to, use a value from the tt:Entity enumeration.

	Entity string `xml:"Entity,attr,omitempty" json:"Entity,omitempty"`

	// Optional entity token.

	Token *ReferenceToken `xml:"Token,attr,omitempty" json:"Token,omitempty"`

	// If this value is true the entity cannot be deleted.

	Fixed bool `xml:"Fixed,attr,omitempty" json:"Fixed,omitempty"`

	// Optional reference to the XAddr of another devices DeviceManagement service.

	GeoSource AnyURI `xml:"GeoSource,attr,omitempty" json:"GeoSource,omitempty"`

	// If set the geo location is obtained internally.

	AutoGeo bool `xml:"AutoGeo,attr,omitempty" json:"AutoGeo,omitempty"`
}

// User readable name. Length up to 64 characters.

type Name string

type IntList []int32

type FloatList []float32

type StringAttrList []string

type StringList []string

type ReferenceTokenList []*ReferenceToken

type RotateMode string

const (

	// Enable the Rotate feature. Degree of rotation is specified Degree parameter.
	RotateModeOFF RotateMode = "OFF"

	// Disable the Rotate feature.
	RotateModeON RotateMode = "ON"

	// Rotate feature is automatically activated by the device.
	RotateModeAUTO RotateMode = "AUTO"
)

type SceneOrientationMode string

const (
	SceneOrientationModeMANUAL SceneOrientationMode = "MANUAL"

	SceneOrientationModeAUTO SceneOrientationMode = "AUTO"
)

type SceneOrientationOption string

const (
	SceneOrientationOptionBelow SceneOrientationOption = "Below"

	SceneOrientationOptionHorizon SceneOrientationOption = "Horizon"

	SceneOrientationOptionAbove SceneOrientationOption = "Above"
)

// Source view modes supported by device.

type ViewModes string

const (

	// Undewarped viewmode from device supporting fisheye lens.
	ViewModesTtFisheye ViewModes = "tt:Fisheye"

	// 360 degree panoramic view.
	ViewModesTt360Panorama ViewModes = "tt:360Panorama"

	// 180 degree panoramic view.
	ViewModesTt180Panorama ViewModes = "tt:180Panorama"

	// View mode combining four streams in single Quad, eg., applicable for devices supporting four heads.
	ViewModesTtQuad ViewModes = "tt:Quad"

	// Unaltered view from the sensor.
	ViewModesTtOriginal ViewModes = "tt:Original"

	// Viewmode combining the left side sensors, applicable for devices supporting multiple sensors.
	ViewModesTtLeftHalf ViewModes = "tt:LeftHalf"

	// Viewmode combining the right side sensors, applicable for devices supporting multiple sensors.
	ViewModesTtRightHalf ViewModes = "tt:RightHalf"

	// Dewarped view mode for device supporting fisheye lens.
	ViewModesTtDewarp ViewModes = "tt:Dewarp"
)

type VideoEncoding string

const (
	VideoEncodingJPEG VideoEncoding = "JPEG"

	VideoEncodingMPEG4 VideoEncoding = "MPEG4"

	VideoEncodingH264 VideoEncoding = "H264"
)

type Mpeg4Profile string

const (
	Mpeg4ProfileSP Mpeg4Profile = "SP"

	Mpeg4ProfileASP Mpeg4Profile = "ASP"
)

type H264Profile string

const (
	H264ProfileBaseline H264Profile = "Baseline"

	H264ProfileMain H264Profile = "Main"

	H264ProfileExtended H264Profile = "Extended"

	H264ProfileHigh H264Profile = "High"
)

//
// Video Media Subtypes as referenced by IANA (without the leading "video/" Video Media Type).  See also
//
// .
//

type VideoEncodingMimeNames string

const (
	VideoEncodingMimeNamesJPEG VideoEncodingMimeNames = "JPEG"

	VideoEncodingMimeNamesMPV4ES VideoEncodingMimeNames = "MPV4-ES"

	VideoEncodingMimeNamesH264 VideoEncodingMimeNames = "H264"

	VideoEncodingMimeNamesH265 VideoEncodingMimeNames = "H265"
)

type VideoEncodingProfiles string

const (
	VideoEncodingProfilesSimple VideoEncodingProfiles = "Simple"

	VideoEncodingProfilesAdvancedSimple VideoEncodingProfiles = "AdvancedSimple"

	VideoEncodingProfilesBaseline VideoEncodingProfiles = "Baseline"

	VideoEncodingProfilesMain VideoEncodingProfiles = "Main"

	VideoEncodingProfilesMain10 VideoEncodingProfiles = "Main10"

	VideoEncodingProfilesExtended VideoEncodingProfiles = "Extended"

	VideoEncodingProfilesHigh VideoEncodingProfiles = "High"
)

type AudioEncoding string

const (
	AudioEncodingG711 AudioEncoding = "G711"

	AudioEncodingG726 AudioEncoding = "G726"

	AudioEncodingAAC AudioEncoding = "AAC"
)

//
// Audio Media Subtypes as referenced by IANA (without the leading "audio/" Audio Media Type and except for the audio types defined in the restriction).  See also
//
// .
//

type AudioEncodingMimeNames string

const (
	AudioEncodingMimeNamesPCMU AudioEncodingMimeNames = "PCMU"

	// AudioEncodingMimeName G726 is used to represent G726-16,G726-24,G726-32 and G726-40 defined in the IANA Media Types
	AudioEncodingMimeNamesG726 AudioEncodingMimeNames = "G726"

	AudioEncodingMimeNamesMP4ALATM AudioEncodingMimeNames = "MP4A-LATM"

	AudioEncodingMimeNamesMpeg4generic AudioEncodingMimeNames = "mpeg4-generic"
)

type MetadataCompressionType string

const (
	MetadataCompressionTypeNone MetadataCompressionType = "None"

	MetadataCompressionTypeGZIP MetadataCompressionType = "GZIP"

	MetadataCompressionTypeEXI MetadataCompressionType = "EXI"
)

type StreamType string

const (
	StreamTypeRTPUnicast StreamType = "RTP-Unicast"

	StreamTypeRTPMulticast StreamType = "RTP-Multicast"
)

type TransportProtocol string

const (
	TransportProtocolUDP TransportProtocol = "UDP"

	// This value is deprecated.
	TransportProtocolTCP TransportProtocol = "TCP"

	TransportProtocolRTSP TransportProtocol = "RTSP"

	TransportProtocolHTTP TransportProtocol = "HTTP"
)

type ScopeDefinition string

const (
	ScopeDefinitionFixed ScopeDefinition = "Fixed"

	ScopeDefinitionConfigurable ScopeDefinition = "Configurable"
)

type DiscoveryMode string

const (
	DiscoveryModeDiscoverable DiscoveryMode = "Discoverable"

	DiscoveryModeNonDiscoverable DiscoveryMode = "NonDiscoverable"
)

type NetworkInterfaceConfigPriority int32

type Duplex string

const (
	DuplexFull Duplex = "Full"

	DuplexHalf Duplex = "Half"
)

type IANAIfTypes int32

type IPv6DHCPConfiguration string

const (
	IPv6DHCPConfigurationAuto IPv6DHCPConfiguration = "Auto"

	IPv6DHCPConfigurationStateful IPv6DHCPConfiguration = "Stateful"

	IPv6DHCPConfigurationStateless IPv6DHCPConfiguration = "Stateless"

	IPv6DHCPConfigurationOff IPv6DHCPConfiguration = "Off"
)

type NetworkProtocolType string

const (
	NetworkProtocolTypeHTTP NetworkProtocolType = "HTTP"

	NetworkProtocolTypeHTTPS NetworkProtocolType = "HTTPS"

	NetworkProtocolTypeRTSP NetworkProtocolType = "RTSP"
)

type NetworkHostType string

const (
	NetworkHostTypeIPv4 NetworkHostType = "IPv4"

	NetworkHostTypeIPv6 NetworkHostType = "IPv6"

	NetworkHostTypeDNS NetworkHostType = "DNS"
)

type IPv4Address string

type IPv6Address string

type HwAddress string

type IPType string

const (
	IPTypeIPv4 IPType = "IPv4"

	IPTypeIPv6 IPType = "IPv6"
)

type DNSName string

type Domain string

type IPAddressFilterType string

const (
	IPAddressFilterTypeAllow IPAddressFilterType = "Allow"

	IPAddressFilterTypeDeny IPAddressFilterType = "Deny"
)

type DynamicDNSType string

const (
	DynamicDNSTypeNoUpdate DynamicDNSType = "NoUpdate"

	DynamicDNSTypeClientUpdates DynamicDNSType = "ClientUpdates"

	DynamicDNSTypeServerUpdates DynamicDNSType = "ServerUpdates"
)

type Dot11SSIDType []byte

type Dot11StationMode string

const (
	Dot11StationModeAdhoc Dot11StationMode = "Ad-hoc"

	Dot11StationModeInfrastructure Dot11StationMode = "Infrastructure"

	Dot11StationModeExtended Dot11StationMode = "Extended"
)

type Dot11SecurityMode string

const (
	Dot11SecurityModeNone Dot11SecurityMode = "None"

	Dot11SecurityModeWEP Dot11SecurityMode = "WEP"

	Dot11SecurityModePSK Dot11SecurityMode = "PSK"

	Dot11SecurityModeDot1X Dot11SecurityMode = "Dot1X"

	Dot11SecurityModeExtended Dot11SecurityMode = "Extended"
)

type Dot11Cipher string

const (
	Dot11CipherCCMP Dot11Cipher = "CCMP"

	Dot11CipherTKIP Dot11Cipher = "TKIP"

	Dot11CipherAny Dot11Cipher = "Any"

	Dot11CipherExtended Dot11Cipher = "Extended"
)

type Dot11PSK []byte

type Dot11PSKPassphrase string

type Dot11SignalStrength string

const (
	Dot11SignalStrengthNone Dot11SignalStrength = "None"

	Dot11SignalStrengthVeryBad Dot11SignalStrength = "Very Bad"

	Dot11SignalStrengthBad Dot11SignalStrength = "Bad"

	Dot11SignalStrengthGood Dot11SignalStrength = "Good"

	Dot11SignalStrengthVeryGood Dot11SignalStrength = "Very Good"

	Dot11SignalStrengthExtended Dot11SignalStrength = "Extended"
)

type Dot11AuthAndMangementSuite string

const (
	Dot11AuthAndMangementSuiteNone Dot11AuthAndMangementSuite = "None"

	Dot11AuthAndMangementSuiteDot1X Dot11AuthAndMangementSuite = "Dot1X"

	Dot11AuthAndMangementSuitePSK Dot11AuthAndMangementSuite = "PSK"

	Dot11AuthAndMangementSuiteExtended Dot11AuthAndMangementSuite = "Extended"
)

type CapabilityCategory string

const (
	CapabilityCategoryAll CapabilityCategory = "All"

	CapabilityCategoryAnalytics CapabilityCategory = "Analytics"

	CapabilityCategoryDevice CapabilityCategory = "Device"

	CapabilityCategoryEvents CapabilityCategory = "Events"

	CapabilityCategoryImaging CapabilityCategory = "Imaging"

	CapabilityCategoryMedia CapabilityCategory = "Media"

	CapabilityCategoryPTZ CapabilityCategory = "PTZ"
)

// Enumeration describing the available system log modes.

type SystemLogType string

const (

	// Indicates that a system log is requested.
	SystemLogTypeSystem SystemLogType = "System"

	// Indicates that a access log is requested.
	SystemLogTypeAccess SystemLogType = "Access"
)

// Enumeration describing the available factory default modes.

type FactoryDefaultType string

const (

	// Indicates that a hard factory default is requested.
	FactoryDefaultTypeHard FactoryDefaultType = "Hard"

	// Indicates that a soft factory default is requested.
	FactoryDefaultTypeSoft FactoryDefaultType = "Soft"
)

type SetDateTimeType string

const (

	// Indicates that the date and time are set manually.
	SetDateTimeTypeManual SetDateTimeType = "Manual"

	// Indicates that the date and time are set through NTP
	SetDateTimeTypeNTP SetDateTimeType = "NTP"
)

type UserLevel string

const (
	UserLevelAdministrator UserLevel = "Administrator"

	UserLevelOperator UserLevel = "Operator"

	UserLevelUser UserLevel = "User"

	UserLevelAnonymous UserLevel = "Anonymous"

	UserLevelExtended UserLevel = "Extended"
)

type RelayLogicalState string

const (
	RelayLogicalStateActive RelayLogicalState = "active"

	RelayLogicalStateInactive RelayLogicalState = "inactive"
)

type RelayIdleState string

const (
	RelayIdleStateClosed RelayIdleState = "closed"

	RelayIdleStateOpen RelayIdleState = "open"
)

type RelayMode string

const (
	RelayModeMonostable RelayMode = "Monostable"

	RelayModeBistable RelayMode = "Bistable"
)

type DigitalIdleState string

const (
	DigitalIdleStateClosed DigitalIdleState = "closed"

	DigitalIdleStateOpen DigitalIdleState = "open"
)

type EFlipMode string

const (
	EFlipModeOFF EFlipMode = "OFF"

	EFlipModeON EFlipMode = "ON"

	EFlipModeExtended EFlipMode = "Extended"
)

type ReverseMode string

const (
	ReverseModeOFF ReverseMode = "OFF"

	ReverseModeON ReverseMode = "ON"

	ReverseModeAUTO ReverseMode = "AUTO"

	ReverseModeExtended ReverseMode = "Extended"
)

type AuxiliaryData string

type PTZPresetTourState string

const (
	PTZPresetTourStateIdle PTZPresetTourState = "Idle"

	PTZPresetTourStateTouring PTZPresetTourState = "Touring"

	PTZPresetTourStatePaused PTZPresetTourState = "Paused"

	PTZPresetTourStateExtended PTZPresetTourState = "Extended"
)

type PTZPresetTourDirection string

const (
	PTZPresetTourDirectionForward PTZPresetTourDirection = "Forward"

	PTZPresetTourDirectionBackward PTZPresetTourDirection = "Backward"

	PTZPresetTourDirectionExtended PTZPresetTourDirection = "Extended"
)

type PTZPresetTourOperation string

const (
	PTZPresetTourOperationStart PTZPresetTourOperation = "Start"

	PTZPresetTourOperationStop PTZPresetTourOperation = "Stop"

	PTZPresetTourOperationPause PTZPresetTourOperation = "Pause"

	PTZPresetTourOperationExtended PTZPresetTourOperation = "Extended"
)

type MoveAndTrackMethod string

const (
	MoveAndTrackMethodPresetToken MoveAndTrackMethod = "PresetToken"

	MoveAndTrackMethodGeoLocation MoveAndTrackMethod = "GeoLocation"

	MoveAndTrackMethodPTZVector MoveAndTrackMethod = "PTZVector"

	MoveAndTrackMethodObjectID MoveAndTrackMethod = "ObjectID"
)

type AutoFocusMode string

const (
	AutoFocusModeAUTO AutoFocusMode = "AUTO"

	AutoFocusModeMANUAL AutoFocusMode = "MANUAL"
)

type AFModes string

const (

	// Focus of a moving camera is updated only once after stopping a pan, tilt or zoom movement.
	AFModesOnceAfterMove AFModes = "OnceAfterMove"
)

type WideDynamicMode string

const (
	WideDynamicModeOFF WideDynamicMode = "OFF"

	WideDynamicModeON WideDynamicMode = "ON"
)

// Enumeration describing the available backlight compenstation modes.

type BacklightCompensationMode string

const (

	// Backlight compensation is disabled.
	BacklightCompensationModeOFF BacklightCompensationMode = "OFF"

	// Backlight compensation is enabled.
	BacklightCompensationModeON BacklightCompensationMode = "ON"
)

type ExposurePriority string

const (
	ExposurePriorityLowNoise ExposurePriority = "LowNoise"

	ExposurePriorityFrameRate ExposurePriority = "FrameRate"
)

type ExposureMode string

const (
	ExposureModeAUTO ExposureMode = "AUTO"

	ExposureModeMANUAL ExposureMode = "MANUAL"
)

type Enabled string

const (
	EnabledENABLED Enabled = "ENABLED"

	EnabledDISABLED Enabled = "DISABLED"
)

type WhiteBalanceMode string

const (
	WhiteBalanceModeAUTO WhiteBalanceMode = "AUTO"

	WhiteBalanceModeMANUAL WhiteBalanceMode = "MANUAL"
)

type IrCutFilterMode string

const (
	IrCutFilterModeON IrCutFilterMode = "ON"

	IrCutFilterModeOFF IrCutFilterMode = "OFF"

	IrCutFilterModeAUTO IrCutFilterMode = "AUTO"
)

type ImageStabilizationMode string

const (
	ImageStabilizationModeOFF ImageStabilizationMode = "OFF"

	ImageStabilizationModeON ImageStabilizationMode = "ON"

	ImageStabilizationModeAUTO ImageStabilizationMode = "AUTO"

	ImageStabilizationModeExtended ImageStabilizationMode = "Extended"
)

type IrCutFilterAutoBoundaryType string

const (
	IrCutFilterAutoBoundaryTypeCommon IrCutFilterAutoBoundaryType = "Common"

	IrCutFilterAutoBoundaryTypeToOn IrCutFilterAutoBoundaryType = "ToOn"

	IrCutFilterAutoBoundaryTypeToOff IrCutFilterAutoBoundaryType = "ToOff"

	IrCutFilterAutoBoundaryTypeExtended IrCutFilterAutoBoundaryType = "Extended"
)

type ToneCompensationMode string

const (
	ToneCompensationModeOFF ToneCompensationMode = "OFF"

	ToneCompensationModeON ToneCompensationMode = "ON"

	ToneCompensationModeAUTO ToneCompensationMode = "AUTO"
)

type DefoggingMode string

const (
	DefoggingModeOFF DefoggingMode = "OFF"

	DefoggingModeON DefoggingMode = "ON"

	DefoggingModeAUTO DefoggingMode = "AUTO"
)

type ImageSendingType string

const (
	ImageSendingTypeEmbedded ImageSendingType = "Embedded"

	ImageSendingTypeLocalStorage ImageSendingType = "LocalStorage"

	ImageSendingTypeRemoteStorage ImageSendingType = "RemoteStorage"
)

type PropertyOperation string

const (
	PropertyOperationInitialized PropertyOperation = "Initialized"

	PropertyOperationDeleted PropertyOperation = "Deleted"

	PropertyOperationChanged PropertyOperation = "Changed"
)

type Direction string

const (
	DirectionLeft Direction = "Left"

	DirectionRight Direction = "Right"

	DirectionAny Direction = "Any"
)

// Specifies a receiver connection mode.

type ReceiverMode string

const (

	// The receiver connects on demand, as required by consumers of the media streams.
	ReceiverModeAutoConnect ReceiverMode = "AutoConnect"

	// The receiver attempts to maintain a persistent connection to the configured endpoint.
	ReceiverModeAlwaysConnect ReceiverMode = "AlwaysConnect"

	// The receiver does not attempt to connect.
	ReceiverModeNeverConnect ReceiverMode = "NeverConnect"

	// This case should never happen.
	ReceiverModeUnknown ReceiverMode = "Unknown"
)

// Specifies the current connection state of the receiver.

type ReceiverState string

const (

	// The receiver is not connected.
	ReceiverStateNotConnected ReceiverState = "NotConnected"

	// The receiver is attempting to connect.
	ReceiverStateConnecting ReceiverState = "Connecting"

	// The receiver is connected.
	ReceiverStateConnected ReceiverState = "Connected"

	// This case should never happen.
	ReceiverStateUnknown ReceiverState = "Unknown"
)

type ReceiverReference *ReferenceToken

type RecordingReference *ReferenceToken

type TrackReference *ReferenceToken

type Description string

type XPathExpression string

type SearchState string

const (

	// The search is queued and not yet started.
	SearchStateQueued SearchState = "Queued"

	// The search is underway and not yet completed.
	SearchStateSearching SearchState = "Searching"

	// The search has been completed and no new results will be found.
	SearchStateCompleted SearchState = "Completed"

	// The state of the search is unknown. (This is not a valid response from GetSearchState.)
	SearchStateUnknown SearchState = "Unknown"
)

type JobToken *ReferenceToken

type TargetFormat string

const (

	// MP4 files with all tracks in a single file.
	TargetFormatMP4 TargetFormat = "MP4"

	// CMAF compliant MP4 files with 1 track per file.
	TargetFormatCMAF TargetFormat = "CMAF"
)

type EncryptionMode string

const (

	// AES-CTR mode full sample and video NAL Subsample encryption, defined in ISO/IEC 23001-7.
	EncryptionModeCENC EncryptionMode = "CENC"

	// AES-CBC mode partial video NAL pattern encryption, defined in ISO/IEC 23001-7.
	EncryptionModeCBCS EncryptionMode = "CBCS"
)

type RecordingStatus string

const (
	RecordingStatusInitiated RecordingStatus = "Initiated"

	RecordingStatusRecording RecordingStatus = "Recording"

	RecordingStatusStopped RecordingStatus = "Stopped"

	RecordingStatusRemoving RecordingStatus = "Removing"

	RecordingStatusRemoved RecordingStatus = "Removed"

	// This case should never happen.
	RecordingStatusUnknown RecordingStatus = "Unknown"
)

type TrackType string

const (
	TrackTypeVideo TrackType = "Video"

	TrackTypeAudio TrackType = "Audio"

	TrackTypeMetadata TrackType = "Metadata"

	// Placeholder for future extension.
	TrackTypeExtended TrackType = "Extended"
)

type RecordingJobReference *ReferenceToken

type RecordingJobMode string

type RecordingJobState string

type ModeOfOperation string

const (
	ModeOfOperationIdle ModeOfOperation = "Idle"

	ModeOfOperationActive ModeOfOperation = "Active"

	// This case should never happen.
	ModeOfOperationUnknown ModeOfOperation = "Unknown"
)

// AudioClassType acceptable values are;
// gun_shot, scream, glass_breaking, tire_screech

type AudioClassType string

const (
	AudioClassTypeGun_shot AudioClassType = "gun_shot"

	AudioClassTypeScream AudioClassType = "scream"

	AudioClassTypeGlass_breaking AudioClassType = "glass_breaking"

	AudioClassTypeTire_screech AudioClassType = "tire_screech"
)

type AudioClassification string

const (
	AudioClassificationGunShot AudioClassification = "GunShot"

	AudioClassificationScream AudioClassification = "Scream"

	AudioClassificationGlassBreaking AudioClassification = "GlassBreaking"

	AudioClassificationTireScreech AudioClassification = "TireScreech"

	AudioClassificationAlarm AudioClassification = "Alarm"
)

type OSDType string

const (
	OSDTypeText OSDType = "Text"

	OSDTypeImage OSDType = "Image"

	OSDTypeExtended OSDType = "Extended"
)

type StringItems struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema StringItems"`

	Item []string `xml:"Item,omitempty" json:"Item,omitempty"`
}

type Message struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/schema Message"`

	// Token value pairs that triggered this message. Typically only one item is present.
	Source *ItemList `xml:"Source,omitempty" json:"Source,omitempty"`

	Key *ItemList `xml:"Key,omitempty" json:"Key,omitempty"`

	Data *ItemList `xml:"Data,omitempty" json:"Data,omitempty"`

	Extension *MessageExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	UtcTime soap.XSDDateTime `xml:"UtcTime,attr,omitempty" json:"UtcTime,omitempty"`

	PropertyOperation *PropertyOperation `xml:"PropertyOperation,attr,omitempty" json:"PropertyOperation,omitempty"`
}

type DeviceEntity struct {

	// Unique identifier referencing the physical entity.

	Token *ReferenceToken `xml:"token,attr,omitempty" json:"token,omitempty"`
}

type IntRectangle struct {
	X int32 `xml:"x,attr,omitempty" json:"x,omitempty"`

	Y int32 `xml:"y,attr,omitempty" json:"y,omitempty"`

	Width int32 `xml:"width,attr,omitempty" json:"width,omitempty"`

	Height int32 `xml:"height,attr,omitempty" json:"height,omitempty"`
}

type IntRectangleRange struct {

	// Range of X-axis.
	XRange *IntRange `xml:"XRange,omitempty" json:"XRange,omitempty"`

	// Range of Y-axis.
	YRange *IntRange `xml:"YRange,omitempty" json:"YRange,omitempty"`

	// Range of width.
	WidthRange *IntRange `xml:"WidthRange,omitempty" json:"WidthRange,omitempty"`

	// Range of height.
	HeightRange *IntRange `xml:"HeightRange,omitempty" json:"HeightRange,omitempty"`
}

type FloatRange struct {
	Min float32 `xml:"Min,omitempty" json:"Min,omitempty"`

	Max float32 `xml:"Max,omitempty" json:"Max,omitempty"`
}

type DurationRange struct {
	Min *time.Duration `xml:"Min,omitempty" json:"Min,omitempty"`

	Max *time.Duration `xml:"Max,omitempty" json:"Max,omitempty"`
}

type IntItems struct {
	Items []int32 `xml:"Items,omitempty" json:"Items,omitempty"`
}

type FloatItems struct {
	Items []float32 `xml:"Items,omitempty" json:"Items,omitempty"`
}

type AnyHolder struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoSource struct {
	*DeviceEntity

	// Frame rate in frames per second.
	Framerate float32 `xml:"Framerate,omitempty" json:"Framerate,omitempty"`

	// Horizontal and vertical resolution
	Resolution *VideoResolution `xml:"Resolution,omitempty" json:"Resolution,omitempty"`

	// Optional configuration of the image sensor.
	Imaging *ImagingSettings `xml:"Imaging,omitempty" json:"Imaging,omitempty"`

	Extension *VideoSourceExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type VideoSourceExtension struct {

	// Optional configuration of the image sensor. To be used if imaging service 2.00 is supported.
	Imaging *ImagingSettings20 `xml:"Imaging,omitempty" json:"Imaging,omitempty"`

	Extension *VideoSourceExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoSourceExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioSource struct {
	*DeviceEntity

	// number of available audio channels. (1: mono, 2: stereo)
	Channels int32 `xml:"Channels,omitempty" json:"Channels,omitempty"`
}

type Profile struct {

	// User readable name of the profile.
	Name *Name `xml:"Name,omitempty" json:"Name,omitempty"`

	// Optional configuration of the Video input.
	VideoSourceConfiguration *VideoSourceConfiguration `xml:"VideoSourceConfiguration,omitempty" json:"VideoSourceConfiguration,omitempty"`

	// Optional configuration of the Audio input.
	AudioSourceConfiguration *AudioSourceConfiguration `xml:"AudioSourceConfiguration,omitempty" json:"AudioSourceConfiguration,omitempty"`

	// Optional configuration of the Video encoder.
	VideoEncoderConfiguration *VideoEncoderConfiguration `xml:"VideoEncoderConfiguration,omitempty" json:"VideoEncoderConfiguration,omitempty"`

	// Optional configuration of the Audio encoder.
	AudioEncoderConfiguration *AudioEncoderConfiguration `xml:"AudioEncoderConfiguration,omitempty" json:"AudioEncoderConfiguration,omitempty"`

	// Optional configuration of the video analytics module and rule engine.
	VideoAnalyticsConfiguration *VideoAnalyticsConfiguration `xml:"VideoAnalyticsConfiguration,omitempty" json:"VideoAnalyticsConfiguration,omitempty"`

	// Optional configuration of the pan tilt zoom unit.
	PTZConfiguration *PTZConfiguration `xml:"PTZConfiguration,omitempty" json:"PTZConfiguration,omitempty"`

	// Optional configuration of the metadata stream.
	MetadataConfiguration *MetadataConfiguration `xml:"MetadataConfiguration,omitempty" json:"MetadataConfiguration,omitempty"`

	// Extensions defined in ONVIF 2.0
	Extension *ProfileExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Unique identifier of the profile.

	Token *ReferenceToken `xml:"token,attr,omitempty" json:"token,omitempty"`

	// A value of true signals that the profile cannot be deleted. Default is false.

	Fixed bool `xml:"fixed,attr,omitempty" json:"fixed,omitempty"`
}

type ProfileExtension struct {

	// Optional configuration of the Audio output.
	AudioOutputConfiguration *AudioOutputConfiguration `xml:"AudioOutputConfiguration,omitempty" json:"AudioOutputConfiguration,omitempty"`

	// Optional configuration of the Audio decoder.
	AudioDecoderConfiguration *AudioDecoderConfiguration `xml:"AudioDecoderConfiguration,omitempty" json:"AudioDecoderConfiguration,omitempty"`

	Extension *ProfileExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ProfileExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ConfigurationEntity struct {

	// User readable name. Length up to 64 characters.
	Name *Name `xml:"Name,omitempty" json:"Name,omitempty"`

	//
	// Number of internal references currently using this configuration.
	//
	//
	UseCount int32 `xml:"UseCount,omitempty" json:"UseCount,omitempty"`

	// Token that uniquely references this configuration. Length up to 64 characters.

	Token *ReferenceToken `xml:"token,attr,omitempty" json:"token,omitempty"`
}

type VideoSourceConfiguration struct {
	*ConfigurationEntity

	// Reference to the physical input.
	SourceToken *ReferenceToken `xml:"SourceToken,omitempty" json:"SourceToken,omitempty"`

	// Rectangle specifying the Video capturing area. The capturing area shall not be larger than the whole Video source area.
	Bounds *IntRectangle `xml:"Bounds,omitempty" json:"Bounds,omitempty"`

	Extension *VideoSourceConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Readonly parameter signalling Source configuration's view mode, for devices supporting different view modes as defined in tt:viewModes.

	ViewMode string `xml:"ViewMode,attr,omitempty" json:"ViewMode,omitempty"`
}

type VideoSourceConfigurationExtension struct {

	//
	// Optional element to configure rotation of captured image.
	// What resolutions a device supports shall be unaffected by the Rotate parameters.
	//
	// If a device is configured with Rotate=AUTO, the device shall take control over the Degree parameter and automatically update it so that a client can query current rotation.
	//
	// The device shall automatically apply the same rotation to its pan/tilt control direction depending on the following condition:
	// if Reverse=AUTO in PTControlDirection or if the device doesn’t support Reverse in PTControlDirection
	//
	Rotate *Rotate `xml:"Rotate,omitempty" json:"Rotate,omitempty"`

	Extension *VideoSourceConfigurationExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type VideoSourceConfigurationExtension2 struct {

	// Optional element describing the geometric lens distortion. Multiple instances for future variable lens support.
	LensDescription []*LensDescription `xml:"LensDescription,omitempty" json:"LensDescription,omitempty"`

	// Optional element describing the scene orientation in the camera’s field of view.
	SceneOrientation *SceneOrientation `xml:"SceneOrientation,omitempty" json:"SceneOrientation,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Rotate struct {

	// Parameter to enable/disable Rotation feature.
	Mode *RotateMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Optional parameter to configure how much degree of clockwise rotation of image  for On mode. Omitting this parameter for On mode means 180 degree rotation.
	Degree int32 `xml:"Degree,omitempty" json:"Degree,omitempty"`

	Extension *RotateExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type RotateExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type LensProjection struct {

	// Angle of incidence.
	Angle float32 `xml:"Angle,omitempty" json:"Angle,omitempty"`

	// Mapping radius as a consequence of the emergent angle.
	Radius float32 `xml:"Radius,omitempty" json:"Radius,omitempty"`

	// Optional ray absorption at the given angle due to vignetting. A value of one means no absorption.
	Transmittance float32 `xml:"Transmittance,omitempty" json:"Transmittance,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type LensOffset struct {

	// Optional horizontal offset of the lens center in normalized coordinates.

	X float32 `xml:"x,attr,omitempty" json:"x,omitempty"`

	// Optional vertical offset of the lens center in normalized coordinates.

	Y float32 `xml:"y,attr,omitempty" json:"y,omitempty"`
}

type LensDescription struct {

	// Offset of the lens center to the imager center in normalized coordinates.
	Offset *LensOffset `xml:"Offset,omitempty" json:"Offset,omitempty"`

	// Radial description of the projection characteristics. The resulting curve is defined by the B-Spline interpolation
	// over the given elements. The element for Radius zero shall not be provided. The projection points shall be ordered with ascending Radius.
	// Items outside the last projection Radius shall be assumed to be invisible (black).
	Projection []*LensProjection `xml:"Projection,omitempty" json:"Projection,omitempty"`

	// Compensation of the x coordinate needed for the ONVIF normalized coordinate system.
	XFactor float32 `xml:"XFactor,omitempty" json:"XFactor,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	// Optional focal length of the optical system.

	FocalLength float32 `xml:"FocalLength,attr,omitempty" json:"FocalLength,omitempty"`
}

type VideoSourceConfigurationOptions struct {

	// Supported range for the capturing area.
	// Device that does not support cropped streaming shall express BoundsRange option as mentioned below
	// BoundsRange->XRange and BoundsRange->YRange with same Min/Max values HeightRange and WidthRange Min/Max values same as VideoSource Height and Width Limits.
	BoundsRange *IntRectangleRange `xml:"BoundsRange,omitempty" json:"BoundsRange,omitempty"`

	// List of physical inputs.
	VideoSourceTokensAvailable []*ReferenceToken `xml:"VideoSourceTokensAvailable,omitempty" json:"VideoSourceTokensAvailable,omitempty"`

	Extension *VideoSourceConfigurationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Maximum number of profiles.

	MaximumNumberOfProfiles int32 `xml:"MaximumNumberOfProfiles,attr,omitempty" json:"MaximumNumberOfProfiles,omitempty"`
}

type VideoSourceConfigurationOptionsExtension struct {

	// Options of parameters for Rotation feature.
	Rotate *RotateOptions `xml:"Rotate,omitempty" json:"Rotate,omitempty"`

	Extension *VideoSourceConfigurationOptionsExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoSourceConfigurationOptionsExtension2 struct {

	// Scene orientation modes supported by the device for this configuration.
	SceneOrientationMode []*SceneOrientationMode `xml:"SceneOrientationMode,omitempty" json:"SceneOrientationMode,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RotateOptions struct {

	// Supported options of Rotate mode parameter.
	Mode []*RotateMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// List of supported degree value for rotation.
	DegreeList *IntItems `xml:"DegreeList,omitempty" json:"DegreeList,omitempty"`

	Extension *RotateOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Signals if a device requires a reboot after changing the rotation.
	// If a device can handle rotation changes without rebooting this value shall be set to false.

	Reboot bool `xml:"Reboot,attr,omitempty" json:"Reboot,omitempty"`
}

type RotateOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type SceneOrientation struct {

	// Parameter to assign the way the camera determines the scene orientation.
	Mode *SceneOrientationMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Assigned or determined scene orientation based on the Mode. When assigning the Mode to AUTO, this field
	// is optional and will be ignored by the device. When assigning the Mode to MANUAL, this field is required
	// and the device will return an InvalidArgs fault if missing.
	Orientation string `xml:"Orientation,omitempty" json:"Orientation,omitempty"`
}

type VideoEncoderConfiguration struct {
	*ConfigurationEntity

	// Used video codec, either Jpeg, H.264 or Mpeg4
	Encoding *VideoEncoding `xml:"Encoding,omitempty" json:"Encoding,omitempty"`

	// Configured video resolution
	Resolution *VideoResolution `xml:"Resolution,omitempty" json:"Resolution,omitempty"`

	// Relative value for the video quantizers and the quality of the video. A high value within supported quality range means higher quality
	Quality float32 `xml:"Quality,omitempty" json:"Quality,omitempty"`

	// Optional element to configure rate control related parameters.
	RateControl *VideoRateControl `xml:"RateControl,omitempty" json:"RateControl,omitempty"`

	// Optional element to configure Mpeg4 related parameters.
	MPEG4 *Mpeg4Configuration `xml:"MPEG4,omitempty" json:"MPEG4,omitempty"`

	// Optional element to configure H.264 related parameters.
	H264 *H264Configuration `xml:"H264,omitempty" json:"H264,omitempty"`

	// Defines the multicast settings that could be used for video streaming.
	Multicast *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty"`

	// The rtsp session timeout for the related video stream
	SessionTimeout *time.Duration `xml:"SessionTimeout,omitempty" json:"SessionTimeout,omitempty"`

	// A value of true indicates that frame rate is a fixed value rather than an upper limit,
	// and that the video encoder shall prioritize frame rate over all other adaptable
	// configuration values such as bitrate.  Default is false.

	GuaranteedFrameRate bool `xml:"GuaranteedFrameRate,attr,omitempty" json:"GuaranteedFrameRate,omitempty"`
}

type VideoResolution struct {

	// Number of the columns of the Video image.
	Width int32 `xml:"Width,omitempty" json:"Width,omitempty"`

	// Number of the lines of the Video image.
	Height int32 `xml:"Height,omitempty" json:"Height,omitempty"`
}

type VideoRateControl struct {

	// Maximum output framerate in fps. If an EncodingInterval is provided the resulting encoded framerate will be reduced by the given factor.
	FrameRateLimit int32 `xml:"FrameRateLimit,omitempty" json:"FrameRateLimit,omitempty"`

	// Interval at which images are encoded and transmitted. (A value of 1 means that every frame is encoded, a value of 2 means that every 2nd frame is encoded ...)
	EncodingInterval int32 `xml:"EncodingInterval,omitempty" json:"EncodingInterval,omitempty"`

	// the maximum output bitrate in kbps
	BitrateLimit int32 `xml:"BitrateLimit,omitempty" json:"BitrateLimit,omitempty"`
}

type Mpeg4Configuration struct {

	// Determines the interval in which the I-Frames will be coded. An entry of 1 indicates I-Frames are continuously generated. An entry of 2 indicates that every 2nd image is an I-Frame, and 3 only every 3rd frame, etc. The frames in between are coded as P or B Frames.
	GovLength int32 `xml:"GovLength,omitempty" json:"GovLength,omitempty"`

	// the Mpeg4 profile, either simple profile (SP) or advanced simple profile (ASP)
	Mpeg4Profile *Mpeg4Profile `xml:"Mpeg4Profile,omitempty" json:"Mpeg4Profile,omitempty"`
}

type H264Configuration struct {

	// Group of Video frames length. Determines typically the interval in which the I-Frames will be coded. An entry of 1 indicates I-Frames are continuously generated. An entry of 2 indicates that every 2nd image is an I-Frame, and 3 only every 3rd frame, etc. The frames in between are coded as P or B Frames.
	GovLength int32 `xml:"GovLength,omitempty" json:"GovLength,omitempty"`

	// the H.264 profile, either baseline, main, extended or high
	H264Profile *H264Profile `xml:"H264Profile,omitempty" json:"H264Profile,omitempty"`
}

type VideoEncoderConfigurationOptions struct {

	// Range of the quality values. A high value means higher quality.
	QualityRange *IntRange `xml:"QualityRange,omitempty" json:"QualityRange,omitempty"`

	// Optional JPEG encoder settings ranges (See also Extension element).
	JPEG *JpegOptions `xml:"JPEG,omitempty" json:"JPEG,omitempty"`

	// Optional MPEG-4 encoder settings ranges (See also Extension element).
	MPEG4 *Mpeg4Options `xml:"MPEG4,omitempty" json:"MPEG4,omitempty"`

	// Optional H.264 encoder settings ranges (See also Extension element).
	H264 *H264Options `xml:"H264,omitempty" json:"H264,omitempty"`

	Extension *VideoEncoderOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Indicates the support for the GuaranteedFrameRate attribute on the VideoEncoderConfiguration element.

	GuaranteedFrameRateSupported bool `xml:"GuaranteedFrameRateSupported,attr,omitempty" json:"GuaranteedFrameRateSupported,omitempty"`
}

type VideoEncoderOptionsExtension struct {

	// Optional JPEG encoder settings ranges.
	JPEG *JpegOptions2 `xml:"JPEG,omitempty" json:"JPEG,omitempty"`

	// Optional MPEG-4 encoder settings ranges.
	MPEG4 *Mpeg4Options2 `xml:"MPEG4,omitempty" json:"MPEG4,omitempty"`

	// Optional H.264 encoder settings ranges.
	H264 *H264Options2 `xml:"H264,omitempty" json:"H264,omitempty"`

	Extension *VideoEncoderOptionsExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoEncoderOptionsExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type JpegOptions struct {

	// List of supported image sizes.
	ResolutionsAvailable []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty"`

	// Supported frame rate in fps (frames per second).
	FrameRateRange *IntRange `xml:"FrameRateRange,omitempty" json:"FrameRateRange,omitempty"`

	// Supported encoding interval range. The encoding interval corresponds to the number of frames devided by the encoded frames. An encoding interval value of "1" means that all frames are encoded.
	EncodingIntervalRange *IntRange `xml:"EncodingIntervalRange,omitempty" json:"EncodingIntervalRange,omitempty"`
}

type JpegOptions2 struct {
	*JpegOptions

	// Supported range of encoded bitrate in kbps.
	BitrateRange *IntRange `xml:"BitrateRange,omitempty" json:"BitrateRange,omitempty"`
}

type Mpeg4Options struct {

	// List of supported image sizes.
	ResolutionsAvailable []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty"`

	// Supported group of Video frames length. This value typically corresponds to the I-Frame distance.
	GovLengthRange *IntRange `xml:"GovLengthRange,omitempty" json:"GovLengthRange,omitempty"`

	// Supported frame rate in fps (frames per second).
	FrameRateRange *IntRange `xml:"FrameRateRange,omitempty" json:"FrameRateRange,omitempty"`

	// Supported encoding interval range. The encoding interval corresponds to the number of frames devided by the encoded frames. An encoding interval value of "1" means that all frames are encoded.
	EncodingIntervalRange *IntRange `xml:"EncodingIntervalRange,omitempty" json:"EncodingIntervalRange,omitempty"`

	// List of supported MPEG-4 profiles.
	Mpeg4ProfilesSupported []*Mpeg4Profile `xml:"Mpeg4ProfilesSupported,omitempty" json:"Mpeg4ProfilesSupported,omitempty"`
}

type Mpeg4Options2 struct {
	*Mpeg4Options

	// Supported range of encoded bitrate in kbps.
	BitrateRange *IntRange `xml:"BitrateRange,omitempty" json:"BitrateRange,omitempty"`
}

type H264Options struct {

	// List of supported image sizes.
	ResolutionsAvailable []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty"`

	// Supported group of Video frames length. This value typically corresponds to the I-Frame distance.
	GovLengthRange *IntRange `xml:"GovLengthRange,omitempty" json:"GovLengthRange,omitempty"`

	// Supported frame rate in fps (frames per second).
	FrameRateRange *IntRange `xml:"FrameRateRange,omitempty" json:"FrameRateRange,omitempty"`

	// Supported encoding interval range. The encoding interval corresponds to the number of frames devided by the encoded frames. An encoding interval value of "1" means that all frames are encoded.
	EncodingIntervalRange *IntRange `xml:"EncodingIntervalRange,omitempty" json:"EncodingIntervalRange,omitempty"`

	// List of supported H.264 profiles.
	H264ProfilesSupported []*H264Profile `xml:"H264ProfilesSupported,omitempty" json:"H264ProfilesSupported,omitempty"`
}

type H264Options2 struct {
	*H264Options

	// Supported range of encoded bitrate in kbps.
	BitrateRange *IntRange `xml:"BitrateRange,omitempty" json:"BitrateRange,omitempty"`
}

type VideoEncoder2Configuration struct {
	*ConfigurationEntity

	//
	// Video Media Subtype for the video format. For definitions see tt:VideoEncodingMimeNames and
	//
	// .
	//
	Encoding string `xml:"Encoding,omitempty" json:"Encoding,omitempty"`

	// Configured video resolution
	Resolution *VideoResolution2 `xml:"Resolution,omitempty" json:"Resolution,omitempty"`

	// Optional element to configure rate control related parameters.
	RateControl *VideoRateControl2 `xml:"RateControl,omitempty" json:"RateControl,omitempty"`

	// Defines the multicast settings that could be used for video streaming.
	Multicast *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty"`

	// Relative value for the video quantizers and the quality of the video. A high value within supported quality range means higher quality
	Quality float32 `xml:"Quality,omitempty" json:"Quality,omitempty"`

	// Group of Video frames length. Determines typically the interval in which the I-Frames will be coded. An entry of 1 indicates I-Frames are continuously generated. An entry of 2 indicates that every 2nd image is an I-Frame, and 3 only every 3rd frame, etc. The frames in between are coded as P or B Frames.

	GovLength int32 `xml:"GovLength,attr,omitempty" json:"GovLength,omitempty"`

	// Distance between anchor frames of type I-Frame and P-Frame. '1' indicates no B-Frames, '2' indicates that every 2nd frame is encoded as B-Frame, '3' indicates a structure like IBBPBBP..., etc.

	AnchorFrameDistance int32 `xml:"AnchorFrameDistance,attr,omitempty" json:"AnchorFrameDistance,omitempty"`

	// The encoder profile as defined in tt:VideoEncodingProfiles.

	Profile string `xml:"Profile,attr,omitempty" json:"Profile,omitempty"`

	// A value of true indicates that frame rate is a fixed value rather than an upper limit,
	// and that the video encoder shall prioritize frame rate over all other adaptable
	// configuration values such as bitrate.  Default is false.

	GuaranteedFrameRate bool `xml:"GuaranteedFrameRate,attr,omitempty" json:"GuaranteedFrameRate,omitempty"`
}

type VideoResolution2 struct {

	// Number of the columns of the Video image.
	Width int32 `xml:"Width,omitempty" json:"Width,omitempty"`

	// Number of the lines of the Video image.
	Height int32 `xml:"Height,omitempty" json:"Height,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoRateControl2 struct {

	// Desired frame rate in fps. The actual rate may be lower due to e.g. performance limitations.
	FrameRateLimit float32 `xml:"FrameRateLimit,omitempty" json:"FrameRateLimit,omitempty"`

	// the maximum output bitrate in kbps
	BitrateLimit int32 `xml:"BitrateLimit,omitempty" json:"BitrateLimit,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	// Enforce constant bitrate.

	ConstantBitRate bool `xml:"ConstantBitRate,attr,omitempty" json:"ConstantBitRate,omitempty"`
}

type VideoEncoder2ConfigurationOptions struct {

	//
	// Video Media Subtype for the video format. For definitions see tt:VideoEncodingMimeNames and
	//
	// .
	//
	Encoding string `xml:"Encoding,omitempty" json:"Encoding,omitempty"`

	// Range of the quality values. A high value means higher quality.
	QualityRange *FloatRange `xml:"QualityRange,omitempty" json:"QualityRange,omitempty"`

	// List of supported image sizes.
	ResolutionsAvailable []*VideoResolution2 `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty"`

	// Supported range of encoded bitrate in kbps.
	BitrateRange *IntRange `xml:"BitrateRange,omitempty" json:"BitrateRange,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	// Exactly two values, which define the Lower and Upper bounds for the supported group of Video frames length. These values typically correspond to the I-Frame distance.

	GovLengthRange *IntList `xml:"GovLengthRange,attr,omitempty" json:"GovLengthRange,omitempty"`

	// Signals support for B-Frames. Upper bound for the supported anchor frame distance (must be larger than one).

	MaxAnchorFrameDistance int32 `xml:"MaxAnchorFrameDistance,attr,omitempty" json:"MaxAnchorFrameDistance,omitempty"`

	// List of supported target frame rates in fps (frames per second). The list shall be sorted with highest values first.

	FrameRatesSupported *FloatList `xml:"FrameRatesSupported,attr,omitempty" json:"FrameRatesSupported,omitempty"`

	// List of supported encoder profiles as defined in tt::VideoEncodingProfiles.

	ProfilesSupported *StringAttrList `xml:"ProfilesSupported,attr,omitempty" json:"ProfilesSupported,omitempty"`

	// Signal whether enforcing constant bitrate is supported.

	ConstantBitRateSupported bool `xml:"ConstantBitRateSupported,attr,omitempty" json:"ConstantBitRateSupported,omitempty"`

	// Indicates the support for the GuaranteedFrameRate attribute on the VideoEncoder2Configuration element.

	GuaranteedFrameRateSupported bool `xml:"GuaranteedFrameRateSupported,attr,omitempty" json:"GuaranteedFrameRateSupported,omitempty"`
}

type AudioSourceConfiguration struct {
	*ConfigurationEntity

	// Token of the Audio Source the configuration applies to
	SourceToken *ReferenceToken `xml:"SourceToken,omitempty" json:"SourceToken,omitempty"`
}

type AudioSourceConfigurationOptions struct {

	// Tokens of the audio source the configuration can be used for.
	InputTokensAvailable []*ReferenceToken `xml:"InputTokensAvailable,omitempty" json:"InputTokensAvailable,omitempty"`

	Extension *AudioSourceOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type AudioSourceOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioEncoderConfiguration struct {
	*ConfigurationEntity

	// Audio codec used for encoding the audio input (either G.711, G.726 or AAC)
	Encoding *AudioEncoding `xml:"Encoding,omitempty" json:"Encoding,omitempty"`

	// The output bitrate in kbps.
	Bitrate int32 `xml:"Bitrate,omitempty" json:"Bitrate,omitempty"`

	// The output sample rate in kHz.
	SampleRate int32 `xml:"SampleRate,omitempty" json:"SampleRate,omitempty"`

	// Defines the multicast settings that could be used for video streaming.
	Multicast *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty"`

	// The rtsp session timeout for the related audio stream
	SessionTimeout *time.Duration `xml:"SessionTimeout,omitempty" json:"SessionTimeout,omitempty"`
}

type AudioEncoderConfigurationOptions struct {

	// list of supported AudioEncoderConfigurations
	Options []*AudioEncoderConfigurationOption `xml:"Options,omitempty" json:"Options,omitempty"`
}

type AudioEncoderConfigurationOption struct {

	// The enoding used for audio data (either G.711, G.726 or AAC)
	Encoding *AudioEncoding `xml:"Encoding,omitempty" json:"Encoding,omitempty"`

	// List of supported bitrates in kbps for the specified Encoding
	BitrateList *IntItems `xml:"BitrateList,omitempty" json:"BitrateList,omitempty"`

	// List of supported Sample Rates in kHz for the specified Encoding
	SampleRateList *IntItems `xml:"SampleRateList,omitempty" json:"SampleRateList,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioEncoder2Configuration struct {
	*ConfigurationEntity

	//
	// Audio Media Subtype for the audio format. For definitions see tt:AudioEncodingMimeNames and
	//
	// .
	//
	Encoding string `xml:"Encoding,omitempty" json:"Encoding,omitempty"`

	// Optional multicast configuration of the audio stream.
	Multicast *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty"`

	// The output bitrate in kbps.
	Bitrate int32 `xml:"Bitrate,omitempty" json:"Bitrate,omitempty"`

	// The output sample rate in kHz.
	SampleRate int32 `xml:"SampleRate,omitempty" json:"SampleRate,omitempty"`
}

type AudioEncoder2ConfigurationOptions struct {

	//
	// Audio Media Subtype for the audio format. For definitions see tt:AudioEncodingMimeNames and
	//
	// .
	//
	Encoding string `xml:"Encoding,omitempty" json:"Encoding,omitempty"`

	// List of supported bitrates in kbps for the specified Encoding
	BitrateList *IntItems `xml:"BitrateList,omitempty" json:"BitrateList,omitempty"`

	// List of supported Sample Rates in kHz for the specified Encoding
	SampleRateList *IntItems `xml:"SampleRateList,omitempty" json:"SampleRateList,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoAnalyticsConfiguration struct {
	*ConfigurationEntity

	AnalyticsEngineConfiguration *AnalyticsEngineConfiguration `xml:"AnalyticsEngineConfiguration,omitempty" json:"AnalyticsEngineConfiguration,omitempty"`

	RuleEngineConfiguration *RuleEngineConfiguration `xml:"RuleEngineConfiguration,omitempty" json:"RuleEngineConfiguration,omitempty"`
}

type MetadataConfiguration struct {
	*ConfigurationEntity

	// optional element to configure which PTZ related data is to include in the metadata stream
	PTZStatus *PTZFilter `xml:"PTZStatus,omitempty" json:"PTZStatus,omitempty"`

	//
	// Optional element to configure the streaming of events. A client might be interested in receiving all,
	// none or some of the events produced by the device:
	//
	//
	Events *EventSubscription `xml:"Events,omitempty" json:"Events,omitempty"`

	// Defines whether the streamed metadata will include metadata from the analytics engines (video, cell motion, audio etc.)
	Analytics bool `xml:"Analytics,omitempty" json:"Analytics,omitempty"`

	// Defines the multicast settings that could be used for video streaming.
	Multicast *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty"`

	// The rtsp session timeout for the related audio stream (when using Media2 Service, this value is deprecated and ignored)
	SessionTimeout *time.Duration `xml:"SessionTimeout,omitempty" json:"SessionTimeout,omitempty"`

	// Indication which AnalyticsModules shall output metadata.
	// Note that the streaming behavior is undefined if the list includes items that are not part of the associated AnalyticsConfiguration.
	AnalyticsEngineConfiguration *AnalyticsEngineConfiguration `xml:"AnalyticsEngineConfiguration,omitempty" json:"AnalyticsEngineConfiguration,omitempty"`

	Extension *MetadataConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Optional parameter to configure compression type of Metadata payload. Use values from enumeration MetadataCompressionType.

	CompressionType string `xml:"CompressionType,attr,omitempty" json:"CompressionType,omitempty"`

	// Optional parameter to configure if the metadata stream shall contain the Geo Location coordinates of each target.

	GeoLocation bool `xml:"GeoLocation,attr,omitempty" json:"GeoLocation,omitempty"`

	// Optional parameter to configure if the generated metadata stream should contain shape information as polygon.

	ShapePolygon bool `xml:"ShapePolygon,attr,omitempty" json:"ShapePolygon,omitempty"`
}

type MetadataConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZFilter struct {

	// True if the metadata stream shall contain the PTZ status (IDLE, MOVING or UNKNOWN)
	Status bool `xml:"Status,omitempty" json:"Status,omitempty"`

	// True if the metadata stream shall contain the PTZ position
	Position bool `xml:"Position,omitempty" json:"Position,omitempty"`
}

type EventSubscription struct {
	Filter *FilterType `xml:"Filter,omitempty" json:"Filter,omitempty"`

	SubscriptionPolicy struct {
	} `xml:"SubscriptionPolicy,omitempty" json:"SubscriptionPolicy,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type MetadataConfigurationOptions struct {
	PTZStatusFilterOptions *PTZStatusFilterOptions `xml:"PTZStatusFilterOptions,omitempty" json:"PTZStatusFilterOptions,omitempty"`

	Extension *MetadataConfigurationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	// True if the device is able to stream the Geo Located positions of each target.

	GeoLocation bool `xml:"GeoLocation,attr,omitempty" json:"GeoLocation,omitempty"`

	// A device signalling support for content filtering shall support expressions with the provided expression size.

	MaxContentFilterSize int32 `xml:"MaxContentFilterSize,attr,omitempty" json:"MaxContentFilterSize,omitempty"`
}

type MetadataConfigurationOptionsExtension struct {

	// List of supported metadata compression type. Its options shall be chosen from tt:MetadataCompressionType.
	CompressionType []string `xml:"CompressionType,omitempty" json:"CompressionType,omitempty"`

	Extension *MetadataConfigurationOptionsExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type MetadataConfigurationOptionsExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZStatusFilterOptions struct {

	// True if the device is able to stream pan or tilt status information.
	PanTiltStatusSupported bool `xml:"PanTiltStatusSupported,omitempty" json:"PanTiltStatusSupported,omitempty"`

	// True if the device is able to stream zoom status inforamtion.
	ZoomStatusSupported bool `xml:"ZoomStatusSupported,omitempty" json:"ZoomStatusSupported,omitempty"`

	// True if the device is able to stream the pan or tilt position.
	PanTiltPositionSupported bool `xml:"PanTiltPositionSupported,omitempty" json:"PanTiltPositionSupported,omitempty"`

	// True if the device is able to stream zoom position information.
	ZoomPositionSupported bool `xml:"ZoomPositionSupported,omitempty" json:"ZoomPositionSupported,omitempty"`

	Extension *PTZStatusFilterOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZStatusFilterOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoOutput struct {
	*DeviceEntity

	Layout *Layout `xml:"Layout,omitempty" json:"Layout,omitempty"`

	// Resolution of the display in Pixel.
	Resolution *VideoResolution `xml:"Resolution,omitempty" json:"Resolution,omitempty"`

	// Refresh rate of the display in Hertz.
	RefreshRate float32 `xml:"RefreshRate,omitempty" json:"RefreshRate,omitempty"`

	// Aspect ratio of the display as physical extent of width divided by height.
	AspectRatio float32 `xml:"AspectRatio,omitempty" json:"AspectRatio,omitempty"`

	Extension *VideoOutputExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type VideoOutputExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoOutputConfiguration struct {
	*ConfigurationEntity

	// Token of the Video Output the configuration applies to
	OutputToken *ReferenceToken `xml:"OutputToken,omitempty" json:"OutputToken,omitempty"`
}

type VideoOutputConfigurationOptions struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoDecoderConfigurationOptions struct {

	// If the device is able to decode Jpeg streams this element describes the supported codecs and configurations
	JpegDecOptions *JpegDecOptions `xml:"JpegDecOptions,omitempty" json:"JpegDecOptions,omitempty"`

	// If the device is able to decode H.264 streams this element describes the supported codecs and configurations
	H264DecOptions *H264DecOptions `xml:"H264DecOptions,omitempty" json:"H264DecOptions,omitempty"`

	// If the device is able to decode Mpeg4 streams this element describes the supported codecs and configurations
	Mpeg4DecOptions *Mpeg4DecOptions `xml:"Mpeg4DecOptions,omitempty" json:"Mpeg4DecOptions,omitempty"`

	Extension *VideoDecoderConfigurationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type H264DecOptions struct {

	// List of supported H.264 Video Resolutions
	ResolutionsAvailable []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty"`

	// List of supported H264 Profiles (either baseline, main, extended or high)
	SupportedH264Profiles []*H264Profile `xml:"SupportedH264Profiles,omitempty" json:"SupportedH264Profiles,omitempty"`

	// Supported H.264 bitrate range in kbps
	SupportedInputBitrate *IntRange `xml:"SupportedInputBitrate,omitempty" json:"SupportedInputBitrate,omitempty"`

	// Supported H.264 framerate range in fps
	SupportedFrameRate *IntRange `xml:"SupportedFrameRate,omitempty" json:"SupportedFrameRate,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type JpegDecOptions struct {

	// List of supported Jpeg Video Resolutions
	ResolutionsAvailable []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty"`

	// Supported Jpeg bitrate range in kbps
	SupportedInputBitrate *IntRange `xml:"SupportedInputBitrate,omitempty" json:"SupportedInputBitrate,omitempty"`

	// Supported Jpeg framerate range in fps
	SupportedFrameRate *IntRange `xml:"SupportedFrameRate,omitempty" json:"SupportedFrameRate,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Mpeg4DecOptions struct {

	// List of supported Mpeg4 Video Resolutions
	ResolutionsAvailable []*VideoResolution `xml:"ResolutionsAvailable,omitempty" json:"ResolutionsAvailable,omitempty"`

	// List of supported Mpeg4 Profiles (either SP or ASP)
	SupportedMpeg4Profiles []*Mpeg4Profile `xml:"SupportedMpeg4Profiles,omitempty" json:"SupportedMpeg4Profiles,omitempty"`

	// Supported Mpeg4 bitrate range in kbps
	SupportedInputBitrate *IntRange `xml:"SupportedInputBitrate,omitempty" json:"SupportedInputBitrate,omitempty"`

	// Supported Mpeg4 framerate range in fps
	SupportedFrameRate *IntRange `xml:"SupportedFrameRate,omitempty" json:"SupportedFrameRate,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoDecoderConfigurationOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioOutput struct {
	*DeviceEntity
}

type AudioOutputConfiguration struct {
	*ConfigurationEntity

	// Token of the phsycial Audio output.
	OutputToken *ReferenceToken `xml:"OutputToken,omitempty" json:"OutputToken,omitempty"`

	//
	// An audio channel MAY support different types of audio transmission. While for full duplex
	// operation no special handling is required, in half duplex operation the transmission direction
	// needs to be switched.
	// The optional SendPrimacy parameter inside the AudioOutputConfiguration indicates which
	// direction is currently active. An NVC can switch between different modes by setting the
	// AudioOutputConfiguration.
	//
	// The following modes for the Send-Primacy are defined:
	//
	// Acoustic echo cancellation is out of ONVIF scope.
	//
	SendPrimacy AnyURI `xml:"SendPrimacy,omitempty" json:"SendPrimacy,omitempty"`

	// Volume setting of the output. The applicable range is defined via the option AudioOutputOptions.OutputLevelRange.
	OutputLevel int32 `xml:"OutputLevel,omitempty" json:"OutputLevel,omitempty"`
}

type AudioOutputConfigurationOptions struct {

	// Tokens of the physical Audio outputs (typically one).
	OutputTokensAvailable []*ReferenceToken `xml:"OutputTokensAvailable,omitempty" json:"OutputTokensAvailable,omitempty"`

	//
	// An
	//
	// channel MAY support different types of audio transmission. While for full duplex
	// operation no special handling is required, in half duplex operation the transmission direction
	// needs to be switched.
	// The optional SendPrimacy parameter inside the AudioOutputConfiguration indicates which
	// direction is currently active. An NVC can switch between different modes by setting the
	// AudioOutputConfiguration.
	//
	// The following modes for the Send-Primacy are defined:
	//
	// Acoustic echo cancellation is out of ONVIF scope.
	//
	SendPrimacyOptions []AnyURI `xml:"SendPrimacyOptions,omitempty" json:"SendPrimacyOptions,omitempty"`

	// Minimum and maximum level range supported for this Output.
	OutputLevelRange *IntRange `xml:"OutputLevelRange,omitempty" json:"OutputLevelRange,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioDecoderConfiguration struct {
	*ConfigurationEntity
}

type AudioDecoderConfigurationOptions struct {

	// If the device is able to decode AAC encoded audio this section describes the supported configurations
	AACDecOptions *AACDecOptions `xml:"AACDecOptions,omitempty" json:"AACDecOptions,omitempty"`

	// If the device is able to decode G711 encoded audio this section describes the supported configurations
	G711DecOptions *G711DecOptions `xml:"G711DecOptions,omitempty" json:"G711DecOptions,omitempty"`

	// If the device is able to decode G726 encoded audio this section describes the supported configurations
	G726DecOptions *G726DecOptions `xml:"G726DecOptions,omitempty" json:"G726DecOptions,omitempty"`

	Extension *AudioDecoderConfigurationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type G711DecOptions struct {

	// List of supported bitrates in kbps
	Bitrate *IntItems `xml:"Bitrate,omitempty" json:"Bitrate,omitempty"`

	// List of supported sample rates in kHz
	SampleRateRange *IntItems `xml:"SampleRateRange,omitempty" json:"SampleRateRange,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AACDecOptions struct {

	// List of supported bitrates in kbps
	Bitrate *IntItems `xml:"Bitrate,omitempty" json:"Bitrate,omitempty"`

	// List of supported sample rates in kHz
	SampleRateRange *IntItems `xml:"SampleRateRange,omitempty" json:"SampleRateRange,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type G726DecOptions struct {

	// List of supported bitrates in kbps
	Bitrate *IntItems `xml:"Bitrate,omitempty" json:"Bitrate,omitempty"`

	// List of supported sample rates in kHz
	SampleRateRange *IntItems `xml:"SampleRateRange,omitempty" json:"SampleRateRange,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioDecoderConfigurationOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type MulticastConfiguration struct {

	// The multicast address (if this address is set to 0 no multicast streaming is enaled)
	Address *IPAddress `xml:"Address,omitempty" json:"Address,omitempty"`

	// The RTP mutlicast destination port. A device may support RTCP. In this case the port value shall be even to allow the corresponding RTCP stream to be mapped to the next higher (odd) destination port number as defined in the RTSP specification.
	Port int32 `xml:"Port,omitempty" json:"Port,omitempty"`

	// In case of IPv6 the TTL value is assumed as the hop limit. Note that for IPV6 and administratively scoped IPv4 multicast the primary use for hop limit / TTL is to prevent packets from (endlessly) circulating and not limiting scope. In these cases the address contains the scope.
	TTL int32 `xml:"TTL,omitempty" json:"TTL,omitempty"`

	// Read only property signalling that streaming is persistant. Use the methods StartMulticastStreaming and StopMulticastStreaming to switch its state.
	AutoStart bool `xml:"AutoStart,omitempty" json:"AutoStart,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type StreamSetup struct {

	// Defines if a multicast or unicast stream is requested
	Stream *StreamType `xml:"Stream,omitempty" json:"Stream,omitempty"`

	Transport *Transport `xml:"Transport,omitempty" json:"Transport,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Transport struct {

	// Defines the network protocol for streaming, either UDP=RTP/UDP, RTSP=RTP/RTSP/TCP or HTTP=RTP/RTSP/HTTP/TCP
	Protocol *TransportProtocol `xml:"Protocol,omitempty" json:"Protocol,omitempty"`

	// Optional element to describe further tunnel options. This element is normally not needed
	Tunnel *Transport `xml:"Tunnel,omitempty" json:"Tunnel,omitempty"`
}

type MediaUri struct {

	// Stable Uri to be used for requesting the media stream
	Uri AnyURI `xml:"Uri,omitempty" json:"Uri,omitempty"`

	// Indicates if the Uri is only valid until the connection is established. The value shall be set to "false".
	InvalidAfterConnect bool `xml:"InvalidAfterConnect,omitempty" json:"InvalidAfterConnect,omitempty"`

	// Indicates if the Uri is invalid after a reboot of the device. The value shall be set to "false".
	InvalidAfterReboot bool `xml:"InvalidAfterReboot,omitempty" json:"InvalidAfterReboot,omitempty"`

	// Duration how long the Uri is valid. This parameter shall be set to PT0S to indicate that this stream URI is indefinitely valid even if the profile changes
	Timeout *time.Duration `xml:"Timeout,omitempty" json:"Timeout,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Scope struct {

	// Indicates if the scope is fixed or configurable.
	ScopeDef *ScopeDefinition `xml:"ScopeDef,omitempty" json:"ScopeDef,omitempty"`

	// Scope item URI.
	ScopeItem AnyURI `xml:"ScopeItem,omitempty" json:"ScopeItem,omitempty"`
}

type NetworkInterface struct {
	*DeviceEntity

	// Indicates whether or not an interface is enabled.
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty"`

	// Network interface information
	Info *NetworkInterfaceInfo `xml:"Info,omitempty" json:"Info,omitempty"`

	// Link configuration.
	Link *NetworkInterfaceLink `xml:"Link,omitempty" json:"Link,omitempty"`

	// IPv4 network interface configuration.
	IPv4 *IPv4NetworkInterface `xml:"IPv4,omitempty" json:"IPv4,omitempty"`

	// IPv6 network interface configuration.
	IPv6 *IPv6NetworkInterface `xml:"IPv6,omitempty" json:"IPv6,omitempty"`

	Extension *NetworkInterfaceExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type NetworkInterfaceExtension struct {
	InterfaceType *IANAIfTypes `xml:"InterfaceType,omitempty" json:"InterfaceType,omitempty"`

	// Extension point prepared for future 802.3 configuration.
	Dot3 []*Dot3Configuration `xml:"Dot3,omitempty" json:"Dot3,omitempty"`

	Dot11 []*Dot11Configuration `xml:"Dot11,omitempty" json:"Dot11,omitempty"`

	Extension *NetworkInterfaceExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Dot3Configuration struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type NetworkInterfaceExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type NetworkInterfaceLink struct {

	// Configured link settings.
	AdminSettings *NetworkInterfaceConnectionSetting `xml:"AdminSettings,omitempty" json:"AdminSettings,omitempty"`

	// Current active link settings.
	OperSettings *NetworkInterfaceConnectionSetting `xml:"OperSettings,omitempty" json:"OperSettings,omitempty"`

	// Integer indicating interface type, for example: 6 is ethernet.
	InterfaceType *IANAIfTypes `xml:"InterfaceType,omitempty" json:"InterfaceType,omitempty"`
}

type NetworkInterfaceConnectionSetting struct {

	// Auto negotiation on/off.
	AutoNegotiation bool `xml:"AutoNegotiation,omitempty" json:"AutoNegotiation,omitempty"`

	// Speed.
	Speed int32 `xml:"Speed,omitempty" json:"Speed,omitempty"`

	// Duplex type, Half or Full.
	Duplex *Duplex `xml:"Duplex,omitempty" json:"Duplex,omitempty"`
}

type NetworkInterfaceInfo struct {

	// Network interface name, for example eth0.
	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	// Network interface MAC address.
	HwAddress *HwAddress `xml:"HwAddress,omitempty" json:"HwAddress,omitempty"`

	// Maximum transmission unit.
	MTU int32 `xml:"MTU,omitempty" json:"MTU,omitempty"`
}

type IPv6NetworkInterface struct {

	// Indicates whether or not IPv6 is enabled.
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty"`

	// IPv6 configuration.
	Config *IPv6Configuration `xml:"Config,omitempty" json:"Config,omitempty"`
}

type IPv4NetworkInterface struct {

	// Indicates whether or not IPv4 is enabled.
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty"`

	// IPv4 configuration.
	Config *IPv4Configuration `xml:"Config,omitempty" json:"Config,omitempty"`
}

type IPv4Configuration struct {

	// List of manually added IPv4 addresses.
	Manual []*PrefixedIPv4Address `xml:"Manual,omitempty" json:"Manual,omitempty"`

	// Link local address.
	LinkLocal *PrefixedIPv4Address `xml:"LinkLocal,omitempty" json:"LinkLocal,omitempty"`

	// IPv4 address configured by using DHCP.
	FromDHCP *PrefixedIPv4Address `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty"`

	// Indicates whether or not DHCP is used.
	DHCP bool `xml:"DHCP,omitempty" json:"DHCP,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type IPv6Configuration struct {

	// Indicates whether router advertisment is used.
	AcceptRouterAdvert bool `xml:"AcceptRouterAdvert,omitempty" json:"AcceptRouterAdvert,omitempty"`

	// DHCP configuration.
	DHCP *IPv6DHCPConfiguration `xml:"DHCP,omitempty" json:"DHCP,omitempty"`

	// List of manually entered IPv6 addresses.
	Manual []*PrefixedIPv6Address `xml:"Manual,omitempty" json:"Manual,omitempty"`

	// List of link local IPv6 addresses.
	LinkLocal []*PrefixedIPv6Address `xml:"LinkLocal,omitempty" json:"LinkLocal,omitempty"`

	// List of IPv6 addresses configured by using DHCP.
	FromDHCP []*PrefixedIPv6Address `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty"`

	// List of IPv6 addresses configured by using router advertisment.
	FromRA []*PrefixedIPv6Address `xml:"FromRA,omitempty" json:"FromRA,omitempty"`

	Extension *IPv6ConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type IPv6ConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type NetworkProtocol struct {

	// Network protocol type string.
	Name *NetworkProtocolType `xml:"Name,omitempty" json:"Name,omitempty"`

	// Indicates if the protocol is enabled or not.
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty"`

	// The port that is used by the protocol.
	Port []int32 `xml:"Port,omitempty" json:"Port,omitempty"`

	Extension *NetworkProtocolExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type NetworkProtocolExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type NetworkHost struct {

	// Network host type: IPv4, IPv6 or DNS.
	Type *NetworkHostType `xml:"Type,omitempty" json:"Type,omitempty"`

	// IPv4 address.
	IPv4Address *IPv4Address `xml:"IPv4Address,omitempty" json:"IPv4Address,omitempty"`

	// IPv6 address.
	IPv6Address *IPv6Address `xml:"IPv6Address,omitempty" json:"IPv6Address,omitempty"`

	// DNS name.
	DNSname *DNSName `xml:"DNSname,omitempty" json:"DNSname,omitempty"`

	Extension *NetworkHostExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type NetworkHostExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type IPAddress struct {

	// Indicates if the address is an IPv4 or IPv6 address.
	Type *IPType `xml:"Type,omitempty" json:"Type,omitempty"`

	// IPv4 address.
	IPv4Address *IPv4Address `xml:"IPv4Address,omitempty" json:"IPv4Address,omitempty"`

	// IPv6 address
	IPv6Address *IPv6Address `xml:"IPv6Address,omitempty" json:"IPv6Address,omitempty"`
}

type PrefixedIPv4Address struct {

	// IPv4 address
	Address *IPv4Address `xml:"Address,omitempty" json:"Address,omitempty"`

	// Prefix/submask length
	PrefixLength int32 `xml:"PrefixLength,omitempty" json:"PrefixLength,omitempty"`
}

type PrefixedIPv6Address struct {

	// IPv6 address
	Address *IPv6Address `xml:"Address,omitempty" json:"Address,omitempty"`

	// Prefix/submask length
	PrefixLength int32 `xml:"PrefixLength,omitempty" json:"PrefixLength,omitempty"`
}

type HostnameInformation struct {

	// Indicates whether the hostname has been obtained from DHCP or not.
	FromDHCP bool `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty"`

	// Indicates the device hostname or an empty string if no hostname has been assigned.
	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Extension *HostnameInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type HostnameInformationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type DNSInformation struct {

	// Indicates whether or not DNS information is retrieved from DHCP.
	FromDHCP bool `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty"`

	// Search domain.
	SearchDomain []string `xml:"SearchDomain,omitempty" json:"SearchDomain,omitempty"`

	// List of DNS addresses received from DHCP.
	DNSFromDHCP []*IPAddress `xml:"DNSFromDHCP,omitempty" json:"DNSFromDHCP,omitempty"`

	// List of manually entered DNS addresses.
	DNSManual []*IPAddress `xml:"DNSManual,omitempty" json:"DNSManual,omitempty"`

	Extension *DNSInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type DNSInformationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type NTPInformation struct {

	// Indicates if NTP information is to be retrieved by using DHCP.
	FromDHCP bool `xml:"FromDHCP,omitempty" json:"FromDHCP,omitempty"`

	// List of NTP addresses retrieved by using DHCP.
	NTPFromDHCP []*NetworkHost `xml:"NTPFromDHCP,omitempty" json:"NTPFromDHCP,omitempty"`

	// List of manually entered NTP addresses.
	NTPManual []*NetworkHost `xml:"NTPManual,omitempty" json:"NTPManual,omitempty"`

	Extension *NTPInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type NTPInformationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type DynamicDNSInformation struct {

	// Dynamic DNS type.
	Type *DynamicDNSType `xml:"Type,omitempty" json:"Type,omitempty"`

	// DNS name.
	Name *DNSName `xml:"Name,omitempty" json:"Name,omitempty"`

	// Time to live.
	TTL *time.Duration `xml:"TTL,omitempty" json:"TTL,omitempty"`

	Extension *DynamicDNSInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type DynamicDNSInformationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type NetworkInterfaceSetConfiguration struct {

	// Indicates whether or not an interface is enabled.
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty"`

	// Link configuration.
	Link *NetworkInterfaceConnectionSetting `xml:"Link,omitempty" json:"Link,omitempty"`

	// Maximum transmission unit.
	MTU int32 `xml:"MTU,omitempty" json:"MTU,omitempty"`

	// IPv4 network interface configuration.
	IPv4 *IPv4NetworkInterfaceSetConfiguration `xml:"IPv4,omitempty" json:"IPv4,omitempty"`

	// IPv6 network interface configuration.
	IPv6 *IPv6NetworkInterfaceSetConfiguration `xml:"IPv6,omitempty" json:"IPv6,omitempty"`

	Extension *NetworkInterfaceSetConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type NetworkInterfaceSetConfigurationExtension struct {
	Dot3 []*Dot3Configuration `xml:"Dot3,omitempty" json:"Dot3,omitempty"`

	Dot11 []*Dot11Configuration `xml:"Dot11,omitempty" json:"Dot11,omitempty"`

	Extension *NetworkInterfaceSetConfigurationExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type IPv6NetworkInterfaceSetConfiguration struct {

	// Indicates whether or not IPv6 is enabled.
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty"`

	// Indicates whether router advertisment is used.
	AcceptRouterAdvert bool `xml:"AcceptRouterAdvert,omitempty" json:"AcceptRouterAdvert,omitempty"`

	// List of manually added IPv6 addresses.
	Manual []*PrefixedIPv6Address `xml:"Manual,omitempty" json:"Manual,omitempty"`

	// DHCP configuration.
	DHCP *IPv6DHCPConfiguration `xml:"DHCP,omitempty" json:"DHCP,omitempty"`
}

type IPv4NetworkInterfaceSetConfiguration struct {

	// Indicates whether or not IPv4 is enabled.
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty"`

	// List of manually added IPv4 addresses.
	Manual []*PrefixedIPv4Address `xml:"Manual,omitempty" json:"Manual,omitempty"`

	// Indicates whether or not DHCP is used.
	DHCP bool `xml:"DHCP,omitempty" json:"DHCP,omitempty"`
}

type NetworkGateway struct {

	// IPv4 address string.
	IPv4Address []*IPv4Address `xml:"IPv4Address,omitempty" json:"IPv4Address,omitempty"`

	// IPv6 address string.
	IPv6Address []*IPv6Address `xml:"IPv6Address,omitempty" json:"IPv6Address,omitempty"`
}

type NetworkZeroConfiguration struct {

	// Unique identifier of network interface.
	InterfaceToken *ReferenceToken `xml:"InterfaceToken,omitempty" json:"InterfaceToken,omitempty"`

	// Indicates whether the zero-configuration is enabled or not.
	Enabled bool `xml:"Enabled,omitempty" json:"Enabled,omitempty"`

	// The zero-configuration IPv4 address(es)
	Addresses []*IPv4Address `xml:"Addresses,omitempty" json:"Addresses,omitempty"`

	Extension *NetworkZeroConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type NetworkZeroConfigurationExtension struct {

	// Optional array holding the configuration for the second and possibly further interfaces.
	Additional []*NetworkZeroConfiguration `xml:"Additional,omitempty" json:"Additional,omitempty"`

	Extension *NetworkZeroConfigurationExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type NetworkZeroConfigurationExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type IPAddressFilter struct {
	Type *IPAddressFilterType `xml:"Type,omitempty" json:"Type,omitempty"`

	IPv4Address []*PrefixedIPv4Address `xml:"IPv4Address,omitempty" json:"IPv4Address,omitempty"`

	IPv6Address []*PrefixedIPv6Address `xml:"IPv6Address,omitempty" json:"IPv6Address,omitempty"`

	Extension *IPAddressFilterExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type IPAddressFilterExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Dot11Configuration struct {
	SSID *Dot11SSIDType `xml:"SSID,omitempty" json:"SSID,omitempty"`

	Mode *Dot11StationMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	Alias *Name `xml:"Alias,omitempty" json:"Alias,omitempty"`

	Priority *NetworkInterfaceConfigPriority `xml:"Priority,omitempty" json:"Priority,omitempty"`

	Security *Dot11SecurityConfiguration `xml:"Security,omitempty" json:"Security,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Dot11SecurityConfiguration struct {
	Mode *Dot11SecurityMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	Algorithm *Dot11Cipher `xml:"Algorithm,omitempty" json:"Algorithm,omitempty"`

	PSK *Dot11PSKSet `xml:"PSK,omitempty" json:"PSK,omitempty"`

	Dot1X *ReferenceToken `xml:"Dot1X,omitempty" json:"Dot1X,omitempty"`

	Extension *Dot11SecurityConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type Dot11SecurityConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Dot11PSKSet struct {

	//
	// According to IEEE802.11-2007 H.4.1 the RSNA PSK consists of 256 bits, or 64 octets when represented in hex
	//
	// Either Key or Passphrase shall be given, if both are supplied Key shall be used by the device and Passphrase ignored.
	//
	Key *Dot11PSK `xml:"Key,omitempty" json:"Key,omitempty"`

	//
	// According to IEEE802.11-2007 H.4.1 a pass-phrase is a sequence of between 8 and 63 ASCII-encoded characters and
	// each character in the pass-phrase must have an encoding in the range of 32 to 126 (decimal),inclusive.
	//
	// If only Passpharse is supplied the Key shall be derived using the algorithm described in IEEE802.11-2007 section H.4
	//
	Passphrase *Dot11PSKPassphrase `xml:"Passphrase,omitempty" json:"Passphrase,omitempty"`

	Extension *Dot11PSKSetExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type Dot11PSKSetExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type NetworkInterfaceSetConfigurationExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Dot11Capabilities struct {
	TKIP bool `xml:"TKIP,omitempty" json:"TKIP,omitempty"`

	ScanAvailableNetworks bool `xml:"ScanAvailableNetworks,omitempty" json:"ScanAvailableNetworks,omitempty"`

	MultipleConfiguration bool `xml:"MultipleConfiguration,omitempty" json:"MultipleConfiguration,omitempty"`

	AdHocStationMode bool `xml:"AdHocStationMode,omitempty" json:"AdHocStationMode,omitempty"`

	WEP bool `xml:"WEP,omitempty" json:"WEP,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Dot11Status struct {
	SSID *Dot11SSIDType `xml:"SSID,omitempty" json:"SSID,omitempty"`

	BSSID string `xml:"BSSID,omitempty" json:"BSSID,omitempty"`

	PairCipher *Dot11Cipher `xml:"PairCipher,omitempty" json:"PairCipher,omitempty"`

	GroupCipher *Dot11Cipher `xml:"GroupCipher,omitempty" json:"GroupCipher,omitempty"`

	SignalStrength *Dot11SignalStrength `xml:"SignalStrength,omitempty" json:"SignalStrength,omitempty"`

	ActiveConfigAlias *ReferenceToken `xml:"ActiveConfigAlias,omitempty" json:"ActiveConfigAlias,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Dot11AvailableNetworks struct {
	SSID *Dot11SSIDType `xml:"SSID,omitempty" json:"SSID,omitempty"`

	BSSID string `xml:"BSSID,omitempty" json:"BSSID,omitempty"`

	// See IEEE802.11 7.3.2.25.2 for details.
	AuthAndMangementSuite []*Dot11AuthAndMangementSuite `xml:"AuthAndMangementSuite,omitempty" json:"AuthAndMangementSuite,omitempty"`

	PairCipher []*Dot11Cipher `xml:"PairCipher,omitempty" json:"PairCipher,omitempty"`

	GroupCipher []*Dot11Cipher `xml:"GroupCipher,omitempty" json:"GroupCipher,omitempty"`

	SignalStrength *Dot11SignalStrength `xml:"SignalStrength,omitempty" json:"SignalStrength,omitempty"`

	Extension *Dot11AvailableNetworksExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type Dot11AvailableNetworksExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

//type Capabilities struct {
//
//	// Analytics capabilities
//	Analytics *AnalyticsCapabilities `xml:"Analytics,omitempty" json:"Analytics,omitempty"`
//
//	// Device capabilities
//	Device *DeviceCapabilities `xml:"Device,omitempty" json:"Device,omitempty"`
//
//	// Event capabilities
//	Events *EventCapabilities `xml:"Events,omitempty" json:"Events,omitempty"`
//
//	// Imaging capabilities
//	Imaging *ImagingCapabilities `xml:"Imaging,omitempty" json:"Imaging,omitempty"`
//
//	// Media capabilities
//	Media *MediaCapabilities `xml:"Media,omitempty" json:"Media,omitempty"`
//
//	// PTZ capabilities
//	PTZ *PTZCapabilities `xml:"PTZ,omitempty" json:"PTZ,omitempty"`
//
//	Extension *CapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
//}

type CapabilitiesExtension struct {
	DeviceIO *DeviceIOCapabilities `xml:"DeviceIO,omitempty" json:"DeviceIO,omitempty"`

	Display *DisplayCapabilities `xml:"Display,omitempty" json:"Display,omitempty"`

	Recording *RecordingCapabilities `xml:"Recording,omitempty" json:"Recording,omitempty"`

	Search *SearchCapabilities `xml:"Search,omitempty" json:"Search,omitempty"`

	Replay *ReplayCapabilities `xml:"Replay,omitempty" json:"Replay,omitempty"`

	Receiver *ReceiverCapabilities `xml:"Receiver,omitempty" json:"Receiver,omitempty"`

	AnalyticsDevice *AnalyticsDeviceCapabilities `xml:"AnalyticsDevice,omitempty" json:"AnalyticsDevice,omitempty"`

	Extensions *CapabilitiesExtension2 `xml:"Extensions,omitempty" json:"Extensions,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type CapabilitiesExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type AnalyticsCapabilities struct {

	// Analytics service URI.
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	// Indicates whether or not rules are supported.
	RuleSupport bool `xml:"RuleSupport,omitempty" json:"RuleSupport,omitempty"`

	// Indicates whether or not modules are supported.
	AnalyticsModuleSupport bool `xml:"AnalyticsModuleSupport,omitempty" json:"AnalyticsModuleSupport,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type DeviceCapabilities struct {

	// Device service URI.
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	// Network capabilities.
	Network *NetworkCapabilities `xml:"Network,omitempty" json:"Network,omitempty"`

	// System capabilities.
	System *SystemCapabilities `xml:"System,omitempty" json:"System,omitempty"`

	// I/O capabilities.
	IO *IOCapabilities `xml:"IO,omitempty" json:"IO,omitempty"`

	// Security capabilities.
	Security *SecurityCapabilities `xml:"Security,omitempty" json:"Security,omitempty"`

	Extension *DeviceCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type DeviceCapabilitiesExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type EventCapabilities struct {

	// Event service URI.
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	// Indicates whether or not WS Subscription policy is supported.
	WSSubscriptionPolicySupport bool `xml:"WSSubscriptionPolicySupport,omitempty" json:"WSSubscriptionPolicySupport,omitempty"`

	// Indicates whether or not WS Pull Point is supported.
	WSPullPointSupport bool `xml:"WSPullPointSupport,omitempty" json:"WSPullPointSupport,omitempty"`

	// Indicates whether or not WS Pausable Subscription Manager Interface is supported.
	WSPausableSubscriptionManagerInterfaceSupport bool `xml:"WSPausableSubscriptionManagerInterfaceSupport,omitempty" json:"WSPausableSubscriptionManagerInterfaceSupport,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type IOCapabilities struct {

	// Number of input connectors.
	InputConnectors int32 `xml:"InputConnectors,omitempty" json:"InputConnectors,omitempty"`

	// Number of relay outputs.
	RelayOutputs int32 `xml:"RelayOutputs,omitempty" json:"RelayOutputs,omitempty"`

	Extension *IOCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type IOCapabilitiesExtension struct {
	Auxiliary bool `xml:"Auxiliary,omitempty" json:"Auxiliary,omitempty"`

	AuxiliaryCommands []*AuxiliaryData `xml:"AuxiliaryCommands,omitempty" json:"AuxiliaryCommands,omitempty"`

	Extension *IOCapabilitiesExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type IOCapabilitiesExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type MediaCapabilities struct {

	// Media service URI.
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	// Streaming capabilities.
	StreamingCapabilities *RealTimeStreamingCapabilities `xml:"StreamingCapabilities,omitempty" json:"StreamingCapabilities,omitempty"`

	Extension *MediaCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type MediaCapabilitiesExtension struct {
	ProfileCapabilities *ProfileCapabilities `xml:"ProfileCapabilities,omitempty" json:"ProfileCapabilities,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RealTimeStreamingCapabilities struct {

	// Indicates whether or not RTP multicast is supported.
	RTPMulticast bool `xml:"RTPMulticast,omitempty" json:"RTPMulticast,omitempty"`

	// Indicates whether or not RTP over TCP is supported.
	RTP_TCP bool `xml:"RTP_TCP,omitempty" json:"RTP_TCP,omitempty"`

	// Indicates whether or not RTP/RTSP/TCP is supported.
	RTP_RTSP_TCP bool `xml:"RTP_RTSP_TCP,omitempty" json:"RTP_RTSP_TCP,omitempty"`

	Extension *RealTimeStreamingCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type RealTimeStreamingCapabilitiesExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ProfileCapabilities struct {

	// Maximum number of profiles.
	MaximumNumberOfProfiles int32 `xml:"MaximumNumberOfProfiles,omitempty" json:"MaximumNumberOfProfiles,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type NetworkCapabilities struct {

	// Indicates whether or not IP filtering is supported.
	IPFilter bool `xml:"IPFilter,omitempty" json:"IPFilter,omitempty"`

	// Indicates whether or not zeroconf is supported.
	ZeroConfiguration bool `xml:"ZeroConfiguration,omitempty" json:"ZeroConfiguration,omitempty"`

	// Indicates whether or not IPv6 is supported.
	IPVersion6 bool `xml:"IPVersion6,omitempty" json:"IPVersion6,omitempty"`

	// Indicates whether or not  is supported.
	DynDNS bool `xml:"DynDNS,omitempty" json:"DynDNS,omitempty"`

	Extension *NetworkCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type NetworkCapabilitiesExtension struct {
	Dot11Configuration bool `xml:"Dot11Configuration,omitempty" json:"Dot11Configuration,omitempty"`

	Extension *NetworkCapabilitiesExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type NetworkCapabilitiesExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type SecurityCapabilities struct {

	// Indicates whether or not TLS 1.1 is supported.
	TLS1_1 bool `xml:"TLS1.1,omitempty" json:"TLS1.1,omitempty"`

	// Indicates whether or not TLS 1.2 is supported.
	TLS1_2 bool `xml:"TLS1.2,omitempty" json:"TLS1.2,omitempty"`

	// Indicates whether or not onboard key generation is supported.
	OnboardKeyGeneration bool `xml:"OnboardKeyGeneration,omitempty" json:"OnboardKeyGeneration,omitempty"`

	// Indicates whether or not access policy configuration is supported.
	AccessPolicyConfig bool `xml:"AccessPolicyConfig,omitempty" json:"AccessPolicyConfig,omitempty"`

	// Indicates whether or not WS-Security X.509 token is supported.
	X_509Token bool `xml:"X.509Token,omitempty" json:"X.509Token,omitempty"`

	// Indicates whether or not WS-Security SAML token is supported.
	SAMLToken bool `xml:"SAMLToken,omitempty" json:"SAMLToken,omitempty"`

	// Indicates whether or not WS-Security Kerberos token is supported.
	KerberosToken bool `xml:"KerberosToken,omitempty" json:"KerberosToken,omitempty"`

	// Indicates whether or not WS-Security REL token is supported.
	RELToken bool `xml:"RELToken,omitempty" json:"RELToken,omitempty"`

	Extension *SecurityCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type SecurityCapabilitiesExtension struct {
	TLS1_0 bool `xml:"TLS1.0,omitempty" json:"TLS1.0,omitempty"`

	Extension *SecurityCapabilitiesExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type SecurityCapabilitiesExtension2 struct {
	Dot1X bool `xml:"Dot1X,omitempty" json:"Dot1X,omitempty"`

	//
	// EAP Methods supported by the device. The int values refer to the
	//
	// .
	//
	SupportedEAPMethod []int32 `xml:"SupportedEAPMethod,omitempty" json:"SupportedEAPMethod,omitempty"`

	RemoteUserHandling bool `xml:"RemoteUserHandling,omitempty" json:"RemoteUserHandling,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type SystemCapabilities struct {

	// Indicates whether or not WS Discovery resolve requests are supported.
	DiscoveryResolve bool `xml:"DiscoveryResolve,omitempty" json:"DiscoveryResolve,omitempty"`

	// Indicates whether or not WS-Discovery Bye is supported.
	DiscoveryBye bool `xml:"DiscoveryBye,omitempty" json:"DiscoveryBye,omitempty"`

	// Indicates whether or not remote discovery is supported.
	RemoteDiscovery bool `xml:"RemoteDiscovery,omitempty" json:"RemoteDiscovery,omitempty"`

	// Indicates whether or not system backup is supported.
	SystemBackup bool `xml:"SystemBackup,omitempty" json:"SystemBackup,omitempty"`

	// Indicates whether or not system logging is supported.
	SystemLogging bool `xml:"SystemLogging,omitempty" json:"SystemLogging,omitempty"`

	// Indicates whether or not firmware upgrade is supported.
	FirmwareUpgrade bool `xml:"FirmwareUpgrade,omitempty" json:"FirmwareUpgrade,omitempty"`

	// Indicates supported ONVIF version(s).
	SupportedVersions []*OnvifVersion `xml:"SupportedVersions,omitempty" json:"SupportedVersions,omitempty"`

	Extension *SystemCapabilitiesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade bool `xml:"HttpFirmwareUpgrade,omitempty" json:"HttpFirmwareUpgrade,omitempty"`

	HttpSystemBackup bool `xml:"HttpSystemBackup,omitempty" json:"HttpSystemBackup,omitempty"`

	HttpSystemLogging bool `xml:"HttpSystemLogging,omitempty" json:"HttpSystemLogging,omitempty"`

	HttpSupportInformation bool `xml:"HttpSupportInformation,omitempty" json:"HttpSupportInformation,omitempty"`

	Extension *SystemCapabilitiesExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type SystemCapabilitiesExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type OnvifVersion struct {

	// Major version number.
	Major int32 `xml:"Major,omitempty" json:"Major,omitempty"`

	// Two digit minor version number.
	// If major version number is less than "16", X.0.1 maps to "01" and X.2.1 maps to "21" where X stands for Major version number.
	// Otherwise, minor number is month of release, such as "06" for June.
	Minor int32 `xml:"Minor,omitempty" json:"Minor,omitempty"`
}

type ImagingCapabilities struct {

	// Imaging service URI.
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`
}

type PTZCapabilities struct {

	// PTZ service URI.
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type DeviceIOCapabilities struct {
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	VideoSources int32 `xml:"VideoSources,omitempty" json:"VideoSources,omitempty"`

	VideoOutputs int32 `xml:"VideoOutputs,omitempty" json:"VideoOutputs,omitempty"`

	AudioSources int32 `xml:"AudioSources,omitempty" json:"AudioSources,omitempty"`

	AudioOutputs int32 `xml:"AudioOutputs,omitempty" json:"AudioOutputs,omitempty"`

	RelayOutputs int32 `xml:"RelayOutputs,omitempty" json:"RelayOutputs,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type DisplayCapabilities struct {
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	// Indication that the SetLayout command supports only predefined layouts.
	FixedLayout bool `xml:"FixedLayout,omitempty" json:"FixedLayout,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingCapabilities struct {
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	ReceiverSource bool `xml:"ReceiverSource,omitempty" json:"ReceiverSource,omitempty"`

	MediaProfileSource bool `xml:"MediaProfileSource,omitempty" json:"MediaProfileSource,omitempty"`

	DynamicRecordings bool `xml:"DynamicRecordings,omitempty" json:"DynamicRecordings,omitempty"`

	DynamicTracks bool `xml:"DynamicTracks,omitempty" json:"DynamicTracks,omitempty"`

	MaxStringLength int32 `xml:"MaxStringLength,omitempty" json:"MaxStringLength,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type SearchCapabilities struct {
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	MetadataSearch bool `xml:"MetadataSearch,omitempty" json:"MetadataSearch,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ReplayCapabilities struct {

	// The address of the replay service.
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ReceiverCapabilities struct {

	// The address of the receiver service.
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	// Indicates whether the device can receive RTP multicast streams.
	RTP_Multicast bool `xml:"RTP_Multicast,omitempty" json:"RTP_Multicast,omitempty"`

	// Indicates whether the device can receive RTP/TCP streams
	RTP_TCP bool `xml:"RTP_TCP,omitempty" json:"RTP_TCP,omitempty"`

	// Indicates whether the device can receive RTP/RTSP/TCP streams.
	RTP_RTSP_TCP bool `xml:"RTP_RTSP_TCP,omitempty" json:"RTP_RTSP_TCP,omitempty"`

	// The maximum number of receivers supported by the device.
	SupportedReceivers int32 `xml:"SupportedReceivers,omitempty" json:"SupportedReceivers,omitempty"`

	// The maximum allowed length for RTSP URIs.
	MaximumRTSPURILength int32 `xml:"MaximumRTSPURILength,omitempty" json:"MaximumRTSPURILength,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AnalyticsDeviceCapabilities struct {
	XAddr AnyURI `xml:"XAddr,omitempty" json:"XAddr,omitempty"`

	// Obsolete property.
	RuleSupport bool `xml:"RuleSupport,omitempty" json:"RuleSupport,omitempty"`

	Extension *AnalyticsDeviceExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type AnalyticsDeviceExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type SystemLog struct {

	// The log information as attachment data.
	Binary *AttachmentData `xml:"Binary,omitempty" json:"Binary,omitempty"`

	// The log information as character data.
	String string `xml:"String,omitempty" json:"String,omitempty"`
}

type SupportInformation struct {

	// The support information as attachment data.
	Binary *AttachmentData `xml:"Binary,omitempty" json:"Binary,omitempty"`

	// The support information as character data.
	String string `xml:"String,omitempty" json:"String,omitempty"`
}

type BinaryData struct {

	// base64 encoded binary data.
	Data []byte `xml:"Data,omitempty" json:"Data,omitempty"`

	ContentType string `xml:"contentType,attr,omitempty" json:"contentType,omitempty"`
}

type AttachmentData struct {
	Include *Include `xml:"Include,omitempty" json:"Include,omitempty"`

	ContentType string `xml:"contentType,attr,omitempty" json:"contentType,omitempty"`
}

type BackupFile struct {
	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Data *AttachmentData `xml:"Data,omitempty" json:"Data,omitempty"`
}

type SystemLogUriList struct {
	SystemLog []*SystemLogUri `xml:"SystemLog,omitempty" json:"SystemLog,omitempty"`
}

type SystemLogUri struct {
	Type *SystemLogType `xml:"Type,omitempty" json:"Type,omitempty"`

	Uri AnyURI `xml:"Uri,omitempty" json:"Uri,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type SystemDateTime struct {

	// Indicates if the time is set manully or through NTP.
	DateTimeType *SetDateTimeType `xml:"DateTimeType,omitempty" json:"DateTimeType,omitempty"`

	// Informative indicator whether daylight savings is currently on/off.
	DaylightSavings bool `xml:"DaylightSavings,omitempty" json:"DaylightSavings,omitempty"`

	// Timezone information in Posix format.
	TimeZone *TimeZone `xml:"TimeZone,omitempty" json:"TimeZone,omitempty"`

	// Current system date and time in UTC format. This field is mandatory since version 2.0.
	UTCDateTime soap.XSDDateTime `xml:"UTCDateTime,omitempty" json:"UTCDateTime,omitempty"`

	// Date and time in local format.
	LocalDateTime soap.XSDDateTime `xml:"LocalDateTime,omitempty" json:"LocalDateTime,omitempty"`

	Extension *SystemDateTimeExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type SystemDateTimeExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type DateTime struct {
	Time soap.XSDTime `xml:"Time,omitempty" json:"Time,omitempty"`

	Date soap.XSDDate `xml:"Date,omitempty" json:"Date,omitempty"`
}

type Date struct {
	Year int32 `xml:"Year,omitempty" json:"Year,omitempty"`

	// Range is 1 to 12.
	Month int32 `xml:"Month,omitempty" json:"Month,omitempty"`

	// Range is 1 to 31.
	Day int32 `xml:"Day,omitempty" json:"Day,omitempty"`
}

type Time struct {

	// Range is 0 to 23.
	Hour int32 `xml:"Hour,omitempty" json:"Hour,omitempty"`

	// Range is 0 to 59.
	Minute int32 `xml:"Minute,omitempty" json:"Minute,omitempty"`

	// Range is 0 to 61 (typically 59).
	Second int32 `xml:"Second,omitempty" json:"Second,omitempty"`
}

type TimeZone struct {

	// Posix timezone string.
	TZ string `xml:"TZ,omitempty" json:"TZ,omitempty"`
}

type RemoteUser struct {
	Username string `xml:"Username,omitempty" json:"Username,omitempty"`

	Password string `xml:"Password,omitempty" json:"Password,omitempty"`

	UseDerivedPassword bool `xml:"UseDerivedPassword,omitempty" json:"UseDerivedPassword,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type User struct {

	// Username string.
	Username string `xml:"Username,omitempty" json:"Username,omitempty"`

	// Password string.
	Password string `xml:"Password,omitempty" json:"Password,omitempty"`

	// User level string.
	UserLevel *UserLevel `xml:"UserLevel,omitempty" json:"UserLevel,omitempty"`

	Extension *UserExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type UserExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type CertificateGenerationParameters struct {
	CertificateID string `xml:"CertificateID,omitempty" json:"CertificateID,omitempty"`

	Subject string `xml:"Subject,omitempty" json:"Subject,omitempty"`

	ValidNotBefore string `xml:"ValidNotBefore,omitempty" json:"ValidNotBefore,omitempty"`

	ValidNotAfter string `xml:"ValidNotAfter,omitempty" json:"ValidNotAfter,omitempty"`

	Extension *CertificateGenerationParametersExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type CertificateGenerationParametersExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Certificate struct {

	// Certificate id.
	CertificateID string `xml:"CertificateID,omitempty" json:"CertificateID,omitempty"`

	// base64 encoded DER representation of certificate.
	Certificate *BinaryData `xml:"Certificate,omitempty" json:"Certificate,omitempty"`
}

type CertificateStatus struct {

	// Certificate id.
	CertificateID string `xml:"CertificateID,omitempty" json:"CertificateID,omitempty"`

	// Indicates whether or not a certificate is used in a HTTPS configuration.
	Status bool `xml:"Status,omitempty" json:"Status,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type CertificateWithPrivateKey struct {
	CertificateID string `xml:"CertificateID,omitempty" json:"CertificateID,omitempty"`

	Certificate *BinaryData `xml:"Certificate,omitempty" json:"Certificate,omitempty"`

	PrivateKey *BinaryData `xml:"PrivateKey,omitempty" json:"PrivateKey,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type CertificateInformation struct {
	CertificateID string `xml:"CertificateID,omitempty" json:"CertificateID,omitempty"`

	IssuerDN string `xml:"IssuerDN,omitempty" json:"IssuerDN,omitempty"`

	SubjectDN string `xml:"SubjectDN,omitempty" json:"SubjectDN,omitempty"`

	KeyUsage *CertificateUsage `xml:"KeyUsage,omitempty" json:"KeyUsage,omitempty"`

	ExtendedKeyUsage *CertificateUsage `xml:"ExtendedKeyUsage,omitempty" json:"ExtendedKeyUsage,omitempty"`

	KeyLength int32 `xml:"KeyLength,omitempty" json:"KeyLength,omitempty"`

	Version string `xml:"Version,omitempty" json:"Version,omitempty"`

	SerialNum string `xml:"SerialNum,omitempty" json:"SerialNum,omitempty"`

	// Validity Range is from "NotBefore" to "NotAfter"; the corresponding DateTimeRange is from "From" to "Until"
	SignatureAlgorithm string `xml:"SignatureAlgorithm,omitempty" json:"SignatureAlgorithm,omitempty"`

	Validity *DateTimeRange `xml:"Validity,omitempty" json:"Validity,omitempty"`

	Extension *CertificateInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type CertificateUsage string

type CertificateInformationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Dot1XConfiguration struct {
	Dot1XConfigurationToken *ReferenceToken `xml:"Dot1XConfigurationToken,omitempty" json:"Dot1XConfigurationToken,omitempty"`

	Identity string `xml:"Identity,omitempty" json:"Identity,omitempty"`

	AnonymousID string `xml:"AnonymousID,omitempty" json:"AnonymousID,omitempty"`

	//
	// EAP Method type as defined in
	//
	// .
	//
	EAPMethod int32 `xml:"EAPMethod,omitempty" json:"EAPMethod,omitempty"`

	CACertificateID []string `xml:"CACertificateID,omitempty" json:"CACertificateID,omitempty"`

	EAPMethodConfiguration *EAPMethodConfiguration `xml:"EAPMethodConfiguration,omitempty" json:"EAPMethodConfiguration,omitempty"`

	Extension *Dot1XConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type Dot1XConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type EAPMethodConfiguration struct {

	// Confgiuration information for TLS Method.
	TLSConfiguration *TLSConfiguration `xml:"TLSConfiguration,omitempty" json:"TLSConfiguration,omitempty"`

	// Password for those EAP Methods that require a password. The password shall never be returned on a get method.
	Password string `xml:"Password,omitempty" json:"Password,omitempty"`

	Extension *EapMethodExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type EapMethodExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type TLSConfiguration struct {
	CertificateID string `xml:"CertificateID,omitempty" json:"CertificateID,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type GenericEapPwdConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type RelayOutputSettings struct {

	//
	// 'Bistable' or 'Monostable'
	//
	//
	Mode *RelayMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Time after which the relay returns to its idle state if it is in monostable mode. If the Mode field is set to bistable mode the value of the parameter can be ignored.
	DelayTime *time.Duration `xml:"DelayTime,omitempty" json:"DelayTime,omitempty"`

	//
	// 'open' or 'closed'
	//
	//
	IdleState *RelayIdleState `xml:"IdleState,omitempty" json:"IdleState,omitempty"`
}

type RelayOutput struct {
	*DeviceEntity

	Properties *RelayOutputSettings `xml:"Properties,omitempty" json:"Properties,omitempty"`
}

type DigitalInput struct {
	*DeviceEntity

	// Indicate the Digital IdleState status.

	IdleState *DigitalIdleState `xml:"IdleState,attr,omitempty" json:"IdleState,omitempty"`
}

type PTZNode struct {
	*DeviceEntity

	// A unique identifier that is used to reference PTZ Nodes.
	Name *Name `xml:"Name,omitempty" json:"Name,omitempty"`

	// A list of Coordinate Systems available for the PTZ Node. For each Coordinate System, the PTZ Node MUST specify its allowed range.
	SupportedPTZSpaces *PTZSpaces `xml:"SupportedPTZSpaces,omitempty" json:"SupportedPTZSpaces,omitempty"`

	// All preset operations MUST be available for this PTZ Node if one preset is supported.
	MaximumNumberOfPresets int32 `xml:"MaximumNumberOfPresets,omitempty" json:"MaximumNumberOfPresets,omitempty"`

	// A boolean operator specifying the availability of a home position. If set to true, the Home Position Operations MUST be available for this PTZ Node.
	HomeSupported bool `xml:"HomeSupported,omitempty" json:"HomeSupported,omitempty"`

	// A list of supported Auxiliary commands. If the list is not empty, the Auxiliary Operations MUST be available for this PTZ Node.
	AuxiliaryCommands []*AuxiliaryData `xml:"AuxiliaryCommands,omitempty" json:"AuxiliaryCommands,omitempty"`

	Extension *PTZNodeExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Indication whether the HomePosition of a Node is fixed or it can be changed via the SetHomePosition command.

	FixedHomePosition bool `xml:"FixedHomePosition,attr,omitempty" json:"FixedHomePosition,omitempty"`

	// Indication whether the Node supports the geo-referenced move command.

	GeoMove bool `xml:"GeoMove,attr,omitempty" json:"GeoMove,omitempty"`
}

type PTZNodeExtension struct {

	// Detail of supported Preset Tour feature.
	SupportedPresetTour *PTZPresetTourSupported `xml:"SupportedPresetTour,omitempty" json:"SupportedPresetTour,omitempty"`

	Extension *PTZNodeExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZNodeExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZPresetTourSupported struct {

	// Indicates number of preset tours that can be created. Required preset tour operations shall be available for this PTZ Node if one or more preset tour is supported.
	MaximumNumberOfPresetTours int32 `xml:"MaximumNumberOfPresetTours,omitempty" json:"MaximumNumberOfPresetTours,omitempty"`

	// Indicates which preset tour operations are available for this PTZ Node.
	PTZPresetTourOperation []*PTZPresetTourOperation `xml:"PTZPresetTourOperation,omitempty" json:"PTZPresetTourOperation,omitempty"`

	Extension *PTZPresetTourSupportedExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PTZPresetTourSupportedExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZConfiguration struct {
	*ConfigurationEntity

	// A mandatory reference to the PTZ Node that the PTZ Configuration belongs to.
	NodeToken *ReferenceToken `xml:"NodeToken,omitempty" json:"NodeToken,omitempty"`

	// If the PTZ Node supports absolute Pan/Tilt movements, it shall specify one Absolute Pan/Tilt Position Space as default.
	DefaultAbsolutePantTiltPositionSpace AnyURI `xml:"DefaultAbsolutePantTiltPositionSpace,omitempty" json:"DefaultAbsolutePantTiltPositionSpace,omitempty"`

	// If the PTZ Node supports absolute zoom movements, it shall specify one Absolute Zoom Position Space as default.
	DefaultAbsoluteZoomPositionSpace AnyURI `xml:"DefaultAbsoluteZoomPositionSpace,omitempty" json:"DefaultAbsoluteZoomPositionSpace,omitempty"`

	// If the PTZ Node supports relative Pan/Tilt movements, it shall specify one RelativePan/Tilt Translation Space as default.
	DefaultRelativePanTiltTranslationSpace AnyURI `xml:"DefaultRelativePanTiltTranslationSpace,omitempty" json:"DefaultRelativePanTiltTranslationSpace,omitempty"`

	// If the PTZ Node supports relative zoom movements, it shall specify one Relative Zoom Translation Space as default.
	DefaultRelativeZoomTranslationSpace AnyURI `xml:"DefaultRelativeZoomTranslationSpace,omitempty" json:"DefaultRelativeZoomTranslationSpace,omitempty"`

	// If the PTZ Node supports continuous Pan/Tilt movements, it shall specify one Continuous Pan/Tilt Velocity Space as default.
	DefaultContinuousPanTiltVelocitySpace AnyURI `xml:"DefaultContinuousPanTiltVelocitySpace,omitempty" json:"DefaultContinuousPanTiltVelocitySpace,omitempty"`

	// If the PTZ Node supports continuous zoom movements, it shall specify one Continuous Zoom Velocity Space as default.
	DefaultContinuousZoomVelocitySpace AnyURI `xml:"DefaultContinuousZoomVelocitySpace,omitempty" json:"DefaultContinuousZoomVelocitySpace,omitempty"`

	// If the PTZ Node supports absolute or relative PTZ movements, it shall specify corresponding default Pan/Tilt and Zoom speeds.
	DefaultPTZSpeed *PTZSpeed `xml:"DefaultPTZSpeed,omitempty" json:"DefaultPTZSpeed,omitempty"`

	// If the PTZ Node supports continuous movements, it shall specify a default timeout, after which the movement stops.
	//DefaultPTZTimeout *time.Duration `xml:"DefaultPTZTimeout,omitempty" json:"DefaultPTZTimeout,omitempty"`

	// The Pan/Tilt limits element should be present for a PTZ Node that supports an absolute Pan/Tilt. If the element is present it signals the support for configurable Pan/Tilt limits. If limits are enabled, the Pan/Tilt movements shall always stay within the specified range. The Pan/Tilt limits are disabled by setting the limits to –INF or +INF.
	PanTiltLimits *PanTiltLimits `xml:"PanTiltLimits,omitempty" json:"PanTiltLimits,omitempty"`

	// The Zoom limits element should be present for a PTZ Node that supports absolute zoom. If the element is present it signals the supports for configurable Zoom limits. If limits are enabled the zoom movements shall always stay within the specified range. The Zoom limits are disabled by settings the limits to -INF and +INF.
	ZoomLimits *ZoomLimits `xml:"ZoomLimits,omitempty" json:"ZoomLimits,omitempty"`

	Extension *PTZConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// The optional acceleration ramp used by the device when moving.

	MoveRamp int32 `xml:"MoveRamp,attr,omitempty" json:"MoveRamp,omitempty"`

	// The optional acceleration ramp used by the device when recalling presets.

	PresetRamp int32 `xml:"PresetRamp,attr,omitempty" json:"PresetRamp,omitempty"`

	// The optional acceleration ramp used by the device when executing PresetTours.

	PresetTourRamp int32 `xml:"PresetTourRamp,attr,omitempty" json:"PresetTourRamp,omitempty"`
}

type PTZConfigurationExtension struct {

	// Optional element to configure PT Control Direction related features.
	PTControlDirection *PTControlDirection `xml:"PTControlDirection,omitempty" json:"PTControlDirection,omitempty"`

	Extension *PTZConfigurationExtension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZConfigurationExtension2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTControlDirection struct {

	// Optional element to configure related parameters for E-Flip.
	EFlip *EFlip `xml:"EFlip,omitempty" json:"EFlip,omitempty"`

	// Optional element to configure related parameters for reversing of PT Control Direction.
	Reverse *Reverse `xml:"Reverse,omitempty" json:"Reverse,omitempty"`

	Extension *PTControlDirectionExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PTControlDirectionExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type EFlip struct {

	// Parameter to enable/disable E-Flip feature.
	Mode *EFlipMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Reverse struct {

	// Parameter to enable/disable Reverse feature.
	Mode *ReverseMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZConfigurationOptions struct {

	// A list of supported coordinate systems including their range limitations.
	Spaces *PTZSpaces `xml:"Spaces,omitempty" json:"Spaces,omitempty"`

	// A timeout Range within which Timeouts are accepted by the PTZ Node.
	PTZTimeout *DurationRange `xml:"PTZTimeout,omitempty" json:"PTZTimeout,omitempty"`

	// Supported options for PT Direction Control.
	PTControlDirection *PTControlDirectionOptions `xml:"PTControlDirection,omitempty" json:"PTControlDirection,omitempty"`

	Extension *PTZConfigurationOptions2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	// The list of acceleration ramps supported by the device. The
	// smallest acceleration value corresponds to the minimal index, the
	// highest acceleration corresponds to the maximum index.

	PTZRamps *IntList `xml:"PTZRamps,attr,omitempty" json:"PTZRamps,omitempty"`
}

type PTZConfigurationOptions2 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTControlDirectionOptions struct {

	// Supported options for EFlip feature.
	EFlip *EFlipOptions `xml:"EFlip,omitempty" json:"EFlip,omitempty"`

	// Supported options for Reverse feature.
	Reverse *ReverseOptions `xml:"Reverse,omitempty" json:"Reverse,omitempty"`

	Extension *PTControlDirectionOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PTControlDirectionOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type EFlipOptions struct {

	// Options of EFlip mode parameter.
	Mode []*EFlipMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	Extension *EFlipOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type EFlipOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ReverseOptions struct {

	// Options of Reverse mode parameter.
	Mode []*ReverseMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	Extension *ReverseOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ReverseOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PanTiltLimits struct {

	// A range of pan tilt limits.
	Range *Space2DDescription `xml:"Range,omitempty" json:"Range,omitempty"`
}

type ZoomLimits struct {

	// A range of zoom limit
	Range *Space1DDescription `xml:"Range,omitempty" json:"Range,omitempty"`
}

type PTZSpaces struct {

	// The Generic Pan/Tilt Position space is provided by every PTZ node that supports absolute Pan/Tilt, since it does not relate to a specific physical range.
	// Instead, the range should be defined as the full range of the PTZ unit normalized to the range -1 to 1 resulting in the following space description.
	AbsolutePanTiltPositionSpace []*Space2DDescription `xml:"AbsolutePanTiltPositionSpace,omitempty" json:"AbsolutePanTiltPositionSpace,omitempty"`

	// The Generic Zoom Position Space is provided by every PTZ node that supports absolute Zoom, since it does not relate to a specific physical range.
	// Instead, the range should be defined as the full range of the Zoom normalized to the range 0 (wide) to 1 (tele).
	// There is no assumption about how the generic zoom range is mapped to magnification, FOV or other physical zoom dimension.
	AbsoluteZoomPositionSpace []*Space1DDescription `xml:"AbsoluteZoomPositionSpace,omitempty" json:"AbsoluteZoomPositionSpace,omitempty"`

	// The Generic Pan/Tilt translation space is provided by every PTZ node that supports relative Pan/Tilt, since it does not relate to a specific physical range.
	// Instead, the range should be defined as the full positive and negative translation range of the PTZ unit normalized to the range -1 to 1,
	// where positive translation would mean clockwise rotation or movement in right/up direction resulting in the following space description.
	RelativePanTiltTranslationSpace []*Space2DDescription `xml:"RelativePanTiltTranslationSpace,omitempty" json:"RelativePanTiltTranslationSpace,omitempty"`

	// The Generic Zoom Translation Space is provided by every PTZ node that supports relative Zoom, since it does not relate to a specific physical range.
	// Instead, the corresponding absolute range should be defined as the full positive and negative translation range of the Zoom normalized to the range -1 to1,
	// where a positive translation maps to a movement in TELE direction. The translation is signed to indicate direction (negative is to wide, positive is to tele).
	// There is no assumption about how the generic zoom range is mapped to magnification, FOV or other physical zoom dimension. This results in the following space description.
	RelativeZoomTranslationSpace []*Space1DDescription `xml:"RelativeZoomTranslationSpace,omitempty" json:"RelativeZoomTranslationSpace,omitempty"`

	// The generic Pan/Tilt velocity space shall be provided by every PTZ node, since it does not relate to a specific physical range.
	// Instead, the range should be defined as a range of the PTZ unit’s speed normalized to the range -1 to 1, where a positive velocity would map to clockwise
	// rotation or movement in the right/up direction. A signed speed can be independently specified for the pan and tilt component resulting in the following space description.
	ContinuousPanTiltVelocitySpace []*Space2DDescription `xml:"ContinuousPanTiltVelocitySpace,omitempty" json:"ContinuousPanTiltVelocitySpace,omitempty"`

	// The generic zoom velocity space specifies a zoom factor velocity without knowing the underlying physical model. The range should be normalized from -1 to 1,
	// where a positive velocity would map to TELE direction. A generic zoom velocity space description resembles the following.
	ContinuousZoomVelocitySpace []*Space1DDescription `xml:"ContinuousZoomVelocitySpace,omitempty" json:"ContinuousZoomVelocitySpace,omitempty"`

	// The speed space specifies the speed for a Pan/Tilt movement when moving to an absolute position or to a relative translation.
	// In contrast to the velocity spaces, speed spaces do not contain any directional information. The speed of a combined Pan/Tilt
	// movement is represented by a single non-negative scalar value.
	PanTiltSpeedSpace []*Space1DDescription `xml:"PanTiltSpeedSpace,omitempty" json:"PanTiltSpeedSpace,omitempty"`

	// The speed space specifies the speed for a Zoom movement when moving to an absolute position or to a relative translation.
	// In contrast to the velocity spaces, speed spaces do not contain any directional information.
	ZoomSpeedSpace []*Space1DDescription `xml:"ZoomSpeedSpace,omitempty" json:"ZoomSpeedSpace,omitempty"`

	Extension *PTZSpacesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PTZSpacesExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Space2DDescription struct {

	// A URI of coordinate systems.
	URI AnyURI `xml:"URI,omitempty" json:"URI,omitempty"`

	// A range of x-axis.
	XRange *FloatRange `xml:"XRange,omitempty" json:"XRange,omitempty"`

	// A range of y-axis.
	YRange *FloatRange `xml:"YRange,omitempty" json:"YRange,omitempty"`
}

type Space1DDescription struct {

	// A URI of coordinate systems.
	URI AnyURI `xml:"URI,omitempty" json:"URI,omitempty"`

	// A range of x-axis.
	XRange *FloatRange `xml:"XRange,omitempty" json:"XRange,omitempty"`
}

type PTZSpeed struct {

	// Pan and tilt speed. The x component corresponds to pan and the y component to tilt. If omitted in a request, the current (if any) PanTilt movement should not be affected.
	PanTilt *Vector2D `xml:"PanTilt,omitempty" json:"PanTilt,omitempty"`

	// A zoom speed. If omitted in a request, the current (if any) Zoom movement should not be affected.
	Zoom *Vector1D `xml:"Zoom,omitempty" json:"Zoom,omitempty"`
}

type PTZPreset struct {

	// A list of preset position name.
	Name *Name `xml:"Name,omitempty" json:"Name,omitempty"`

	// A list of preset position.
	PTZPosition *PTZVector `xml:"PTZPosition,omitempty" json:"PTZPosition,omitempty"`

	Token *ReferenceToken `xml:"token,attr,omitempty" json:"token,omitempty"`
}

type PresetTour struct {

	// Readable name of the preset tour.
	Name *Name `xml:"Name,omitempty" json:"Name,omitempty"`

	// Read only parameters to indicate the status of the preset tour.
	Status *PTZPresetTourStatus `xml:"Status,omitempty" json:"Status,omitempty"`

	// Auto Start flag of the preset tour. True allows the preset tour to be activated always.
	AutoStart bool `xml:"AutoStart,omitempty" json:"AutoStart,omitempty"`

	// Parameters to specify the detail behavior of the preset tour.
	StartingCondition *PTZPresetTourStartingCondition `xml:"StartingCondition,omitempty" json:"StartingCondition,omitempty"`

	// A list of detail of touring spots including preset positions.
	TourSpot []*PTZPresetTourSpot `xml:"TourSpot,omitempty" json:"TourSpot,omitempty"`

	Extension *PTZPresetTourExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Unique identifier of this preset tour.

	Token *ReferenceToken `xml:"token,attr,omitempty" json:"token,omitempty"`
}

type PTZPresetTourExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZPresetTourSpot struct {

	// Detail definition of preset position of the tour spot.
	PresetDetail *PTZPresetTourPresetDetail `xml:"PresetDetail,omitempty" json:"PresetDetail,omitempty"`

	// Optional parameter to specify Pan/Tilt and Zoom speed on moving toward this tour spot.
	Speed *PTZSpeed `xml:"Speed,omitempty" json:"Speed,omitempty"`

	// Optional parameter to specify time duration of staying on this tour sport.
	StayTime *time.Duration `xml:"StayTime,omitempty" json:"StayTime,omitempty"`

	Extension *PTZPresetTourSpotExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PTZPresetTourSpotExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZPresetTourPresetDetail struct {
	Items []string `xml:",any" json:"items,omitempty"`

	// Option to specify the preset position with Preset Token defined in advance.
	PresetToken *ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty"`

	// Option to specify the preset position with the home position of this PTZ Node. "False" to this parameter shall be treated as an invalid argument.
	Home bool `xml:"Home,omitempty" json:"Home,omitempty"`

	// Option to specify the preset position with vector of PTZ node directly.
	PTZPosition *PTZVector `xml:"PTZPosition,omitempty" json:"PTZPosition,omitempty"`

	TypeExtension *PTZPresetTourTypeExtension `xml:"TypeExtension,omitempty" json:"TypeExtension,omitempty"`
}

type PTZPresetTourTypeExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZPresetTourStatus struct {

	// Indicates state of this preset tour by Idle/Touring/Paused.
	State *PTZPresetTourState `xml:"State,omitempty" json:"State,omitempty"`

	// Indicates a tour spot currently staying.
	CurrentTourSpot *PTZPresetTourSpot `xml:"CurrentTourSpot,omitempty" json:"CurrentTourSpot,omitempty"`

	Extension *PTZPresetTourStatusExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PTZPresetTourStatusExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZPresetTourStartingCondition struct {

	// Optional parameter to specify how many times the preset tour is recurred.
	RecurringTime int32 `xml:"RecurringTime,omitempty" json:"RecurringTime,omitempty"`

	// Optional parameter to specify how long time duration the preset tour is recurred.
	RecurringDuration *time.Duration `xml:"RecurringDuration,omitempty" json:"RecurringDuration,omitempty"`

	// Optional parameter to choose which direction the preset tour goes. Forward shall be chosen in case it is omitted.
	Direction *PTZPresetTourDirection `xml:"Direction,omitempty" json:"Direction,omitempty"`

	Extension *PTZPresetTourStartingConditionExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Execute presets in random order. If set to true and Direction is also present, Direction will be ignored and presets of the Tour will be recalled randomly.

	RandomPresetOrder bool `xml:"RandomPresetOrder,attr,omitempty" json:"RandomPresetOrder,omitempty"`
}

type PTZPresetTourStartingConditionExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZPresetTourOptions struct {

	// Indicates whether or not the AutoStart is supported.
	AutoStart bool `xml:"AutoStart,omitempty" json:"AutoStart,omitempty"`

	// Supported options for Preset Tour Starting Condition.
	StartingCondition *PTZPresetTourStartingConditionOptions `xml:"StartingCondition,omitempty" json:"StartingCondition,omitempty"`

	// Supported options for Preset Tour Spot.
	TourSpot *PTZPresetTourSpotOptions `xml:"TourSpot,omitempty" json:"TourSpot,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZPresetTourSpotOptions struct {

	// Supported options for detail definition of preset position of the tour spot.
	PresetDetail *PTZPresetTourPresetDetailOptions `xml:"PresetDetail,omitempty" json:"PresetDetail,omitempty"`

	// Supported range of stay time for a tour spot.
	StayTime *DurationRange `xml:"StayTime,omitempty" json:"StayTime,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZPresetTourPresetDetailOptions struct {

	// A list of available Preset Tokens for tour spots.
	PresetToken []*ReferenceToken `xml:"PresetToken,omitempty" json:"PresetToken,omitempty"`

	// An option to indicate Home postion for tour spots.
	Home bool `xml:"Home,omitempty" json:"Home,omitempty"`

	// Supported range of Pan and Tilt for tour spots.
	PanTiltPositionSpace *Space2DDescription `xml:"PanTiltPositionSpace,omitempty" json:"PanTiltPositionSpace,omitempty"`

	// Supported range of Zoom for a tour spot.
	ZoomPositionSpace *Space1DDescription `xml:"ZoomPositionSpace,omitempty" json:"ZoomPositionSpace,omitempty"`

	Extension *PTZPresetTourPresetDetailOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PTZPresetTourPresetDetailOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZPresetTourStartingConditionOptions struct {

	// Supported range of Recurring Time.
	RecurringTime *IntRange `xml:"RecurringTime,omitempty" json:"RecurringTime,omitempty"`

	// Supported range of Recurring Duration.
	RecurringDuration *DurationRange `xml:"RecurringDuration,omitempty" json:"RecurringDuration,omitempty"`

	// Supported options for Direction of Preset Tour.
	Direction []*PTZPresetTourDirection `xml:"Direction,omitempty" json:"Direction,omitempty"`

	Extension *PTZPresetTourStartingConditionOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PTZPresetTourStartingConditionOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ImagingStatus struct {
	FocusStatus *FocusStatus `xml:"FocusStatus,omitempty" json:"FocusStatus,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type FocusStatus struct {

	// Status of focus position.
	Position float32 `xml:"Position,omitempty" json:"Position,omitempty"`

	// Status of focus MoveStatus.
	MoveStatus *MoveStatus `xml:"MoveStatus,omitempty" json:"MoveStatus,omitempty"`

	// Error status of focus.
	Error string `xml:"Error,omitempty" json:"Error,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type FocusConfiguration struct {
	AutoFocusMode *AutoFocusMode `xml:"AutoFocusMode,omitempty" json:"AutoFocusMode,omitempty"`

	DefaultSpeed float32 `xml:"DefaultSpeed,omitempty" json:"DefaultSpeed,omitempty"`

	// Parameter to set autofocus near limit (unit: meter).
	NearLimit float32 `xml:"NearLimit,omitempty" json:"NearLimit,omitempty"`

	// Parameter to set autofocus far limit (unit: meter).
	// If set to 0.0, infinity will be used.
	FarLimit float32 `xml:"FarLimit,omitempty" json:"FarLimit,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ImagingSettings struct {

	// Enabled/disabled BLC mode (on/off).
	BacklightCompensation *BacklightCompensation `xml:"BacklightCompensation,omitempty" json:"BacklightCompensation,omitempty"`

	// Image brightness (unit unspecified).
	Brightness float32 `xml:"Brightness,omitempty" json:"Brightness,omitempty"`

	// Color saturation of the image (unit unspecified).
	ColorSaturation float32 `xml:"ColorSaturation,omitempty" json:"ColorSaturation,omitempty"`

	// Contrast of the image (unit unspecified).
	Contrast float32 `xml:"Contrast,omitempty" json:"Contrast,omitempty"`

	// Exposure mode of the device.
	Exposure *Exposure `xml:"Exposure,omitempty" json:"Exposure,omitempty"`

	// Focus configuration.
	Focus *FocusConfiguration `xml:"Focus,omitempty" json:"Focus,omitempty"`

	// Infrared Cutoff Filter settings.
	IrCutFilter *IrCutFilterMode `xml:"IrCutFilter,omitempty" json:"IrCutFilter,omitempty"`

	// Sharpness of the Video image.
	Sharpness float32 `xml:"Sharpness,omitempty" json:"Sharpness,omitempty"`

	// WDR settings.
	WideDynamicRange *WideDynamicRange `xml:"WideDynamicRange,omitempty" json:"WideDynamicRange,omitempty"`

	// White balance settings.
	WhiteBalance *WhiteBalance `xml:"WhiteBalance,omitempty" json:"WhiteBalance,omitempty"`

	Extension *ImagingSettingsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ImagingSettingsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Exposure struct {

	//
	// Exposure Mode
	//
	//
	Mode *ExposureMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// The exposure priority mode (low noise/framerate).
	Priority *ExposurePriority `xml:"Priority,omitempty" json:"Priority,omitempty"`

	// Rectangular exposure mask.
	Window *Rectangle `xml:"Window,omitempty" json:"Window,omitempty"`

	// Minimum value of exposure time range allowed to be used by the algorithm.
	MinExposureTime float32 `xml:"MinExposureTime,omitempty" json:"MinExposureTime,omitempty"`

	// Maximum value of exposure time range allowed to be used by the algorithm.
	MaxExposureTime float32 `xml:"MaxExposureTime,omitempty" json:"MaxExposureTime,omitempty"`

	// Minimum value of the sensor gain range that is allowed to be used by the algorithm.
	MinGain float32 `xml:"MinGain,omitempty" json:"MinGain,omitempty"`

	// Maximum value of the sensor gain range that is allowed to be used by the algorithm.
	MaxGain float32 `xml:"MaxGain,omitempty" json:"MaxGain,omitempty"`

	// Minimum value of the iris range allowed to be used by the algorithm.
	MinIris float32 `xml:"MinIris,omitempty" json:"MinIris,omitempty"`

	// Maximum value of the iris range allowed to be used by the algorithm.
	MaxIris float32 `xml:"MaxIris,omitempty" json:"MaxIris,omitempty"`

	// The fixed exposure time used by the image sensor (μs).
	ExposureTime float32 `xml:"ExposureTime,omitempty" json:"ExposureTime,omitempty"`

	// The fixed gain used by the image sensor (dB).
	Gain float32 `xml:"Gain,omitempty" json:"Gain,omitempty"`

	// The fixed attenuation of input light affected by the iris (dB). 0dB maps to a fully opened iris.
	Iris float32 `xml:"Iris,omitempty" json:"Iris,omitempty"`
}

type WideDynamicRange struct {

	// White dynamic range (on/off)
	Mode *WideDynamicMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Optional level parameter (unitless)
	Level float32 `xml:"Level,omitempty" json:"Level,omitempty"`
}

type BacklightCompensation struct {

	// Backlight compensation mode (on/off).
	Mode *BacklightCompensationMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Optional level parameter (unit unspecified).
	Level float32 `xml:"Level,omitempty" json:"Level,omitempty"`
}

type ImagingOptions struct {
	BacklightCompensation *BacklightCompensationOptions `xml:"BacklightCompensation,omitempty" json:"BacklightCompensation,omitempty"`

	Brightness *FloatRange `xml:"Brightness,omitempty" json:"Brightness,omitempty"`

	ColorSaturation *FloatRange `xml:"ColorSaturation,omitempty" json:"ColorSaturation,omitempty"`

	Contrast *FloatRange `xml:"Contrast,omitempty" json:"Contrast,omitempty"`

	Exposure *ExposureOptions `xml:"Exposure,omitempty" json:"Exposure,omitempty"`

	Focus *FocusOptions `xml:"Focus,omitempty" json:"Focus,omitempty"`

	IrCutFilterModes []*IrCutFilterMode `xml:"IrCutFilterModes,omitempty" json:"IrCutFilterModes,omitempty"`

	Sharpness *FloatRange `xml:"Sharpness,omitempty" json:"Sharpness,omitempty"`

	WideDynamicRange *WideDynamicRangeOptions `xml:"WideDynamicRange,omitempty" json:"WideDynamicRange,omitempty"`

	WhiteBalance *WhiteBalanceOptions `xml:"WhiteBalance,omitempty" json:"WhiteBalance,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type WideDynamicRangeOptions struct {
	Mode []*WideDynamicMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	Level *FloatRange `xml:"Level,omitempty" json:"Level,omitempty"`
}

type BacklightCompensationOptions struct {
	Mode []*WideDynamicMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	Level *FloatRange `xml:"Level,omitempty" json:"Level,omitempty"`
}

type FocusOptions struct {
	AutoFocusModes []*AutoFocusMode `xml:"AutoFocusModes,omitempty" json:"AutoFocusModes,omitempty"`

	DefaultSpeed *FloatRange `xml:"DefaultSpeed,omitempty" json:"DefaultSpeed,omitempty"`

	NearLimit *FloatRange `xml:"NearLimit,omitempty" json:"NearLimit,omitempty"`

	FarLimit *FloatRange `xml:"FarLimit,omitempty" json:"FarLimit,omitempty"`
}

type ExposureOptions struct {
	Mode []*ExposureMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	Priority []*ExposurePriority `xml:"Priority,omitempty" json:"Priority,omitempty"`

	MinExposureTime *FloatRange `xml:"MinExposureTime,omitempty" json:"MinExposureTime,omitempty"`

	MaxExposureTime *FloatRange `xml:"MaxExposureTime,omitempty" json:"MaxExposureTime,omitempty"`

	MinGain *FloatRange `xml:"MinGain,omitempty" json:"MinGain,omitempty"`

	MaxGain *FloatRange `xml:"MaxGain,omitempty" json:"MaxGain,omitempty"`

	MinIris *FloatRange `xml:"MinIris,omitempty" json:"MinIris,omitempty"`

	MaxIris *FloatRange `xml:"MaxIris,omitempty" json:"MaxIris,omitempty"`

	ExposureTime *FloatRange `xml:"ExposureTime,omitempty" json:"ExposureTime,omitempty"`

	Gain *FloatRange `xml:"Gain,omitempty" json:"Gain,omitempty"`

	Iris *FloatRange `xml:"Iris,omitempty" json:"Iris,omitempty"`
}

type WhiteBalanceOptions struct {
	Mode []*WhiteBalanceMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	YrGain *FloatRange `xml:"YrGain,omitempty" json:"YrGain,omitempty"`

	YbGain *FloatRange `xml:"YbGain,omitempty" json:"YbGain,omitempty"`
}

type FocusMove struct {

	// Parameters for the absolute focus control.
	Absolute *AbsoluteFocus `xml:"Absolute,omitempty" json:"Absolute,omitempty"`

	// Parameters for the relative focus control.
	Relative *RelativeFocus `xml:"Relative,omitempty" json:"Relative,omitempty"`

	// Parameter for the continuous focus control.
	Continuous *ContinuousFocus `xml:"Continuous,omitempty" json:"Continuous,omitempty"`
}

type AbsoluteFocus struct {

	// Position parameter for the absolute focus control.
	Position float32 `xml:"Position,omitempty" json:"Position,omitempty"`

	// Speed parameter for the absolute focus control.
	Speed float32 `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type RelativeFocus struct {

	// Distance parameter for the relative focus control.
	Distance float32 `xml:"Distance,omitempty" json:"Distance,omitempty"`

	// Speed parameter for the relative focus control.
	Speed float32 `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type ContinuousFocus struct {

	// Speed parameter for the Continuous focus control.
	Speed float32 `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type MoveOptions struct {
	Absolute *AbsoluteFocusOptions `xml:"Absolute,omitempty" json:"Absolute,omitempty"`

	Relative *RelativeFocusOptions `xml:"Relative,omitempty" json:"Relative,omitempty"`

	Continuous *ContinuousFocusOptions `xml:"Continuous,omitempty" json:"Continuous,omitempty"`
}

type AbsoluteFocusOptions struct {

	// Valid ranges of the position.
	Position *FloatRange `xml:"Position,omitempty" json:"Position,omitempty"`

	// Valid ranges of the speed.
	Speed *FloatRange `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type RelativeFocusOptions struct {

	// Valid ranges of the distance.
	Distance *FloatRange `xml:"Distance,omitempty" json:"Distance,omitempty"`

	// Valid ranges of the speed.
	Speed *FloatRange `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type ContinuousFocusOptions struct {

	// Valid ranges of the speed.
	Speed *FloatRange `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type WhiteBalance struct {

	// Auto whitebalancing mode (auto/manual).
	Mode *WhiteBalanceMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Rgain (unitless).
	CrGain float32 `xml:"CrGain,omitempty" json:"CrGain,omitempty"`

	// Bgain (unitless).
	CbGain float32 `xml:"CbGain,omitempty" json:"CbGain,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ImagingStatus20 struct {

	// Status of focus.
	FocusStatus20 *FocusStatus20 `xml:"FocusStatus20,omitempty" json:"FocusStatus20,omitempty"`

	Extension *ImagingStatus20Extension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ImagingStatus20Extension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type FocusStatus20 struct {

	// Status of focus position.
	Position float32 `xml:"Position,omitempty" json:"Position,omitempty"`

	// Status of focus MoveStatus.
	MoveStatus *MoveStatus `xml:"MoveStatus,omitempty" json:"MoveStatus,omitempty"`

	// Error status of focus.
	Error string `xml:"Error,omitempty" json:"Error,omitempty"`

	Extension *FocusStatus20Extension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type FocusStatus20Extension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ImagingSettings20 struct {

	// Enabled/disabled BLC mode (on/off).
	BacklightCompensation *BacklightCompensation20 `xml:"BacklightCompensation,omitempty" json:"BacklightCompensation,omitempty"`

	// Image brightness (unit unspecified).
	Brightness float32 `xml:"Brightness,omitempty" json:"Brightness,omitempty"`

	// Color saturation of the image (unit unspecified).
	ColorSaturation float32 `xml:"ColorSaturation,omitempty" json:"ColorSaturation,omitempty"`

	// Contrast of the image (unit unspecified).
	Contrast float32 `xml:"Contrast,omitempty" json:"Contrast,omitempty"`

	// Exposure mode of the device.
	Exposure *Exposure20 `xml:"Exposure,omitempty" json:"Exposure,omitempty"`

	// Focus configuration.
	Focus *FocusConfiguration20 `xml:"Focus,omitempty" json:"Focus,omitempty"`

	// Infrared Cutoff Filter settings.
	IrCutFilter *IrCutFilterMode `xml:"IrCutFilter,omitempty" json:"IrCutFilter,omitempty"`

	// Sharpness of the Video image.
	Sharpness float32 `xml:"Sharpness,omitempty" json:"Sharpness,omitempty"`

	// WDR settings.
	WideDynamicRange *WideDynamicRange20 `xml:"WideDynamicRange,omitempty" json:"WideDynamicRange,omitempty"`

	// White balance settings.
	WhiteBalance *WhiteBalance20 `xml:"WhiteBalance,omitempty" json:"WhiteBalance,omitempty"`

	Extension *ImagingSettingsExtension20 `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ImagingSettingsExtension20 struct {

	// Optional element to configure Image Stabilization feature.
	ImageStabilization *ImageStabilization `xml:"ImageStabilization,omitempty" json:"ImageStabilization,omitempty"`

	Extension *ImagingSettingsExtension202 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ImagingSettingsExtension202 struct {

	// An optional parameter applied to only auto mode to adjust timing of toggling Ir cut filter.
	IrCutFilterAutoAdjustment []*IrCutFilterAutoAdjustment `xml:"IrCutFilterAutoAdjustment,omitempty" json:"IrCutFilterAutoAdjustment,omitempty"`

	Extension *ImagingSettingsExtension203 `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ImagingSettingsExtension203 struct {

	// Optional element to configure Image Contrast Compensation.
	ToneCompensation *ToneCompensation `xml:"ToneCompensation,omitempty" json:"ToneCompensation,omitempty"`

	// Optional element to configure Image Defogging.
	Defogging *Defogging `xml:"Defogging,omitempty" json:"Defogging,omitempty"`

	// Optional element to configure Image Noise Reduction.
	NoiseReduction *NoiseReduction `xml:"NoiseReduction,omitempty" json:"NoiseReduction,omitempty"`

	Extension *ImagingSettingsExtension204 `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ImagingSettingsExtension204 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ImageStabilization struct {

	// Parameter to enable/disable Image Stabilization feature.
	Mode *ImageStabilizationMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Optional level parameter (unit unspecified)
	Level float32 `xml:"Level,omitempty" json:"Level,omitempty"`

	Extension *ImageStabilizationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ImageStabilizationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type IrCutFilterAutoAdjustment struct {

	// Specifies which boundaries to automatically toggle Ir cut filter following parameters are applied to. Its options shall be chosen from tt:IrCutFilterAutoBoundaryType.
	BoundaryType string `xml:"BoundaryType,omitempty" json:"BoundaryType,omitempty"`

	// Adjusts boundary exposure level for toggling Ir cut filter to on/off specified with unitless normalized value from +1.0 to -1.0. Zero is default and -1.0 is the darkest adjustment (Unitless).
	BoundaryOffset float32 `xml:"BoundaryOffset,omitempty" json:"BoundaryOffset,omitempty"`

	// Delay time of toggling Ir cut filter to on/off after crossing of the boundary exposure levels.
	ResponseTime *time.Duration `xml:"ResponseTime,omitempty" json:"ResponseTime,omitempty"`

	Extension *IrCutFilterAutoAdjustmentExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type IrCutFilterAutoAdjustmentExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type WideDynamicRange20 struct {

	// Wide dynamic range mode (on/off).
	Mode *WideDynamicMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Optional level parameter (unit unspecified).
	Level float32 `xml:"Level,omitempty" json:"Level,omitempty"`
}

type BacklightCompensation20 struct {

	// Backlight compensation mode (on/off).
	Mode *BacklightCompensationMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Optional level parameter (unit unspecified).
	Level float32 `xml:"Level,omitempty" json:"Level,omitempty"`
}

type Exposure20 struct {

	//
	// Exposure Mode
	//
	//
	Mode *ExposureMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// The exposure priority mode (low noise/framerate).
	Priority *ExposurePriority `xml:"Priority,omitempty" json:"Priority,omitempty"`

	// Rectangular exposure mask.
	Window *Rectangle `xml:"Window,omitempty" json:"Window,omitempty"`

	// Minimum value of exposure time range allowed to be used by the algorithm.
	MinExposureTime float32 `xml:"MinExposureTime,omitempty" json:"MinExposureTime,omitempty"`

	// Maximum value of exposure time range allowed to be used by the algorithm.
	MaxExposureTime float32 `xml:"MaxExposureTime,omitempty" json:"MaxExposureTime,omitempty"`

	// Minimum value of the sensor gain range that is allowed to be used by the algorithm.
	MinGain float32 `xml:"MinGain,omitempty" json:"MinGain,omitempty"`

	// Maximum value of the sensor gain range that is allowed to be used by the algorithm.
	MaxGain float32 `xml:"MaxGain,omitempty" json:"MaxGain,omitempty"`

	// Minimum value of the iris range allowed to be used by the algorithm.  0dB maps to a fully opened iris and positive values map to higher attenuation.
	MinIris float32 `xml:"MinIris,omitempty" json:"MinIris,omitempty"`

	// Maximum value of the iris range allowed to be used by the algorithm. 0dB maps to a fully opened iris and positive values map to higher attenuation.
	MaxIris float32 `xml:"MaxIris,omitempty" json:"MaxIris,omitempty"`

	// The fixed exposure time used by the image sensor (μs).
	ExposureTime float32 `xml:"ExposureTime,omitempty" json:"ExposureTime,omitempty"`

	// The fixed gain used by the image sensor (dB).
	Gain float32 `xml:"Gain,omitempty" json:"Gain,omitempty"`

	// The fixed attenuation of input light affected by the iris (dB). 0dB maps to a fully opened iris and positive values map to higher attenuation.
	Iris float32 `xml:"Iris,omitempty" json:"Iris,omitempty"`
}

type ToneCompensation struct {

	// Parameter to enable/disable or automatic ToneCompensation feature. Its options shall be chosen from tt:ToneCompensationMode Type.
	Mode string `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Optional level parameter specified with unitless normalized value from 0.0 to +1.0.
	Level float32 `xml:"Level,omitempty" json:"Level,omitempty"`

	Extension *ToneCompensationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ToneCompensationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Defogging struct {

	// Parameter to enable/disable or automatic Defogging feature. Its options shall be chosen from tt:DefoggingMode Type.
	Mode string `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Optional level parameter specified with unitless normalized value from 0.0 to +1.0.
	Level float32 `xml:"Level,omitempty" json:"Level,omitempty"`

	Extension *DefoggingExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type DefoggingExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type NoiseReduction struct {

	// Level parameter specified with unitless normalized value from 0.0 to +1.0. Level=0 means no noise reduction or minimal noise reduction.
	Level float32 `xml:"Level,omitempty" json:"Level,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ImagingOptions20 struct {

	// Valid range of Backlight Compensation.
	BacklightCompensation *BacklightCompensationOptions20 `xml:"BacklightCompensation,omitempty" json:"BacklightCompensation,omitempty"`

	// Valid range of Brightness.
	Brightness *FloatRange `xml:"Brightness,omitempty" json:"Brightness,omitempty"`

	// Valid range of Color Saturation.
	ColorSaturation *FloatRange `xml:"ColorSaturation,omitempty" json:"ColorSaturation,omitempty"`

	// Valid range of Contrast.
	Contrast *FloatRange `xml:"Contrast,omitempty" json:"Contrast,omitempty"`

	// Valid range of Exposure.
	Exposure *ExposureOptions20 `xml:"Exposure,omitempty" json:"Exposure,omitempty"`

	// Valid range of Focus.
	Focus *FocusOptions20 `xml:"Focus,omitempty" json:"Focus,omitempty"`

	// Valid range of IrCutFilterModes.
	IrCutFilterModes []*IrCutFilterMode `xml:"IrCutFilterModes,omitempty" json:"IrCutFilterModes,omitempty"`

	// Valid range of Sharpness.
	Sharpness *FloatRange `xml:"Sharpness,omitempty" json:"Sharpness,omitempty"`

	// Valid range of WideDynamicRange.
	WideDynamicRange *WideDynamicRangeOptions20 `xml:"WideDynamicRange,omitempty" json:"WideDynamicRange,omitempty"`

	// Valid range of WhiteBalance.
	WhiteBalance *WhiteBalanceOptions20 `xml:"WhiteBalance,omitempty" json:"WhiteBalance,omitempty"`

	Extension *ImagingOptions20Extension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ImagingOptions20Extension struct {

	// Options of parameters for Image Stabilization feature.
	ImageStabilization *ImageStabilizationOptions `xml:"ImageStabilization,omitempty" json:"ImageStabilization,omitempty"`

	Extension *ImagingOptions20Extension2 `xml:"Extension,omitempty" json:"Extension,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ImagingOptions20Extension2 struct {

	// Options of parameters for adjustment of Ir cut filter auto mode.
	IrCutFilterAutoAdjustment *IrCutFilterAutoAdjustmentOptions `xml:"IrCutFilterAutoAdjustment,omitempty" json:"IrCutFilterAutoAdjustment,omitempty"`

	Extension *ImagingOptions20Extension3 `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ImagingOptions20Extension3 struct {

	// Options of parameters for Tone Compensation feature.
	ToneCompensationOptions *ToneCompensationOptions `xml:"ToneCompensationOptions,omitempty" json:"ToneCompensationOptions,omitempty"`

	// Options of parameters for Defogging feature.
	DefoggingOptions *DefoggingOptions `xml:"DefoggingOptions,omitempty" json:"DefoggingOptions,omitempty"`

	// Options of parameter for Noise Reduction feature.
	NoiseReductionOptions *NoiseReductionOptions `xml:"NoiseReductionOptions,omitempty" json:"NoiseReductionOptions,omitempty"`

	Extension *ImagingOptions20Extension4 `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ImagingOptions20Extension4 struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ImageStabilizationOptions struct {

	// Supported options of Image Stabilization mode parameter.
	Mode []*ImageStabilizationMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Valid range of the Image Stabilization.
	Level *FloatRange `xml:"Level,omitempty" json:"Level,omitempty"`

	Extension *ImageStabilizationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ImageStabilizationOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type IrCutFilterAutoAdjustmentOptions struct {

	// Supported options of boundary types for adjustment of Ir cut filter auto mode. The opptions shall be chosen from tt:IrCutFilterAutoBoundaryType.
	BoundaryType []string `xml:"BoundaryType,omitempty" json:"BoundaryType,omitempty"`

	// Indicates whether or not boundary offset for toggling Ir cut filter is supported.
	BoundaryOffset bool `xml:"BoundaryOffset,omitempty" json:"BoundaryOffset,omitempty"`

	// Supported range of delay time for toggling Ir cut filter.
	ResponseTimeRange *DurationRange `xml:"ResponseTimeRange,omitempty" json:"ResponseTimeRange,omitempty"`

	Extension *IrCutFilterAutoAdjustmentOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type IrCutFilterAutoAdjustmentOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type WideDynamicRangeOptions20 struct {
	Mode []*WideDynamicMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	Level *FloatRange `xml:"Level,omitempty" json:"Level,omitempty"`
}

type BacklightCompensationOptions20 struct {

	// 'ON' or 'OFF'
	Mode []*BacklightCompensationMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Level range of BacklightCompensation.
	Level *FloatRange `xml:"Level,omitempty" json:"Level,omitempty"`
}

type ExposureOptions20 struct {

	//
	// Exposure Mode
	//
	//
	Mode []*ExposureMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// The exposure priority mode (low noise/framerate).
	Priority []*ExposurePriority `xml:"Priority,omitempty" json:"Priority,omitempty"`

	// Valid range of the Minimum ExposureTime.
	MinExposureTime *FloatRange `xml:"MinExposureTime,omitempty" json:"MinExposureTime,omitempty"`

	// Valid range of the Maximum ExposureTime.
	MaxExposureTime *FloatRange `xml:"MaxExposureTime,omitempty" json:"MaxExposureTime,omitempty"`

	// Valid range of the Minimum Gain.
	MinGain *FloatRange `xml:"MinGain,omitempty" json:"MinGain,omitempty"`

	// Valid range of the Maximum Gain.
	MaxGain *FloatRange `xml:"MaxGain,omitempty" json:"MaxGain,omitempty"`

	// Valid range of the Minimum Iris.
	MinIris *FloatRange `xml:"MinIris,omitempty" json:"MinIris,omitempty"`

	// Valid range of the Maximum Iris.
	MaxIris *FloatRange `xml:"MaxIris,omitempty" json:"MaxIris,omitempty"`

	// Valid range of the ExposureTime.
	ExposureTime *FloatRange `xml:"ExposureTime,omitempty" json:"ExposureTime,omitempty"`

	// Valid range of the Gain.
	Gain *FloatRange `xml:"Gain,omitempty" json:"Gain,omitempty"`

	// Valid range of the Iris.
	Iris *FloatRange `xml:"Iris,omitempty" json:"Iris,omitempty"`
}

type MoveOptions20 struct {

	// Valid ranges for the absolute control.
	Absolute *AbsoluteFocusOptions `xml:"Absolute,omitempty" json:"Absolute,omitempty"`

	// Valid ranges for the relative control.
	Relative *RelativeFocusOptions20 `xml:"Relative,omitempty" json:"Relative,omitempty"`

	// Valid ranges for the continuous control.
	Continuous *ContinuousFocusOptions `xml:"Continuous,omitempty" json:"Continuous,omitempty"`
}

type RelativeFocusOptions20 struct {

	// Valid ranges of the distance.
	Distance *FloatRange `xml:"Distance,omitempty" json:"Distance,omitempty"`

	// Valid ranges of the speed.
	Speed *FloatRange `xml:"Speed,omitempty" json:"Speed,omitempty"`
}

type WhiteBalance20 struct {

	// 'AUTO' or 'MANUAL'
	Mode *WhiteBalanceMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Rgain (unitless).
	CrGain float32 `xml:"CrGain,omitempty" json:"CrGain,omitempty"`

	// Bgain (unitless).
	CbGain float32 `xml:"CbGain,omitempty" json:"CbGain,omitempty"`

	Extension *WhiteBalance20Extension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type WhiteBalance20Extension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type FocusConfiguration20 struct {

	//
	// Mode of auto focus.
	//
	// Note: for devices supporting both manual and auto operation at the same time manual operation may be supported even if the Mode parameter is set to Auto.
	//
	AutoFocusMode *AutoFocusMode `xml:"AutoFocusMode,omitempty" json:"AutoFocusMode,omitempty"`

	DefaultSpeed float32 `xml:"DefaultSpeed,omitempty" json:"DefaultSpeed,omitempty"`

	// Parameter to set autofocus near limit (unit: meter).
	NearLimit float32 `xml:"NearLimit,omitempty" json:"NearLimit,omitempty"`

	// Parameter to set autofocus far limit (unit: meter).
	FarLimit float32 `xml:"FarLimit,omitempty" json:"FarLimit,omitempty"`

	Extension *FocusConfiguration20Extension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Zero or more modes as defined in enumeration tt:AFModes.

	AFMode *StringAttrList `xml:"AFMode,attr,omitempty" json:"AFMode,omitempty"`
}

type FocusConfiguration20Extension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type WhiteBalanceOptions20 struct {

	//
	// Mode of WhiteBalance.
	//
	//
	Mode []*WhiteBalanceMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	YrGain *FloatRange `xml:"YrGain,omitempty" json:"YrGain,omitempty"`

	YbGain *FloatRange `xml:"YbGain,omitempty" json:"YbGain,omitempty"`

	Extension *WhiteBalanceOptions20Extension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type WhiteBalanceOptions20Extension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type FocusOptions20 struct {

	//
	// Supported modes for auto focus.
	//
	//
	AutoFocusModes []*AutoFocusMode `xml:"AutoFocusModes,omitempty" json:"AutoFocusModes,omitempty"`

	// Valid range of DefaultSpeed.
	DefaultSpeed *FloatRange `xml:"DefaultSpeed,omitempty" json:"DefaultSpeed,omitempty"`

	// Valid range of NearLimit.
	NearLimit *FloatRange `xml:"NearLimit,omitempty" json:"NearLimit,omitempty"`

	// Valid range of FarLimit.
	FarLimit *FloatRange `xml:"FarLimit,omitempty" json:"FarLimit,omitempty"`

	Extension *FocusOptions20Extension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type FocusOptions20Extension struct {

	// Supported options for auto focus. Options shall be chosen from tt:AFModes.
	AFModes *StringAttrList `xml:"AFModes,omitempty" json:"AFModes,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ToneCompensationOptions struct {

	// Supported options for Tone Compensation mode. Its options shall be chosen from tt:ToneCompensationMode Type.
	Mode []string `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Indicates whether or not support Level parameter for Tone Compensation.
	Level bool `xml:"Level,omitempty" json:"Level,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type DefoggingOptions struct {

	// Supported options for Defogging mode. Its options shall be chosen from tt:DefoggingMode Type.
	Mode []string `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Indicates whether or not support Level parameter for Defogging.
	Level bool `xml:"Level,omitempty" json:"Level,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type NoiseReductionOptions struct {

	// Indicates whether or not support Level parameter for NoiseReduction.
	Level bool `xml:"Level,omitempty" json:"Level,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type MessageExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ItemList struct {
	SimpleItem []struct {

		// Item name.

		Name string `xml:"Name,attr,omitempty" json:"Name,omitempty"`

		// Item value. The type is defined in the corresponding description.

		Value *xsdt.AnySimpleType `xml:"Value,attr,omitempty" json:"Value,omitempty"`
	} `xml:"SimpleItem,omitempty" json:"SimpleItem,omitempty"`

	ElementItem []struct {

		// Item name.

		Name string `xml:"Name,attr,omitempty" json:"Name,omitempty"`
	} `xml:"ElementItem,omitempty" json:"ElementItem,omitempty"`

	Extension *ItemListExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ItemListExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type MessageDescription struct {

	//
	// Set of tokens producing this message. The list may only contain SimpleItemDescription items.
	// The set of tokens identify the component within the WS-Endpoint, which is responsible for the producing the message.
	//
	// For analytics events the token set shall include the VideoSourceConfigurationToken, the VideoAnalyticsConfigurationToken
	// and the name of the analytics module or rule.
	//
	Source *ItemListDescription `xml:"Source,omitempty" json:"Source,omitempty"`

	// Describes optional message payload parameters that may be used as key. E.g. object IDs of tracked objects are conveyed as key.
	Key *ItemListDescription `xml:"Key,omitempty" json:"Key,omitempty"`

	// Describes the payload of the message.
	Data *ItemListDescription `xml:"Data,omitempty" json:"Data,omitempty"`

	Extension *MessageDescriptionExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	//
	// Must be set to true when the described Message relates to a property. An alternative term of "property" is a "state" in contrast to a pure event, which contains relevant information for only a single point in time.
	//
	// Default is false.
	//

	IsProperty bool `xml:"IsProperty,attr,omitempty" json:"IsProperty,omitempty"`
}

type MessageDescriptionExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ItemListDescription struct {
	SimpleItemDescription []struct {

		// Item name. Must be unique within a list.

		Name string `xml:"Name,attr,omitempty" json:"Name,omitempty"`

		//Type *QName `xml:"Type,attr,omitempty" json:"Type,omitempty"`
	} `xml:"SimpleItemDescription,omitempty" json:"SimpleItemDescription,omitempty"`

	ElementItemDescription []struct {

		// Item name. Must be unique within a list.

		Name string `xml:"Name,attr,omitempty" json:"Name,omitempty"`

		// The type of the item. The Type must reference a defined type.

		//Type *QName `xml:"Type,attr,omitempty" json:"Type,omitempty"`
	} `xml:"ElementItemDescription,omitempty" json:"ElementItemDescription,omitempty"`

	Extension *ItemListDescriptionExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ItemListDescriptionExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Polyline struct {
	Point []*Vector `xml:"Point,omitempty" json:"Point,omitempty"`
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule []*Config `xml:"AnalyticsModule,omitempty" json:"AnalyticsModule,omitempty"`

	Extension *AnalyticsEngineConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type AnalyticsEngineConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type RuleEngineConfiguration struct {
	Rule []*Config `xml:"Rule,omitempty" json:"Rule,omitempty"`

	Extension *RuleEngineConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type RuleEngineConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Config struct {

	// List of configuration parameters as defined in the corresponding description.
	Parameters *ItemList `xml:"Parameters,omitempty" json:"Parameters,omitempty"`

	// Name of the configuration.

	Name string `xml:"Name,attr,omitempty" json:"Name,omitempty"`

	// The Type attribute specifies the type of rule and shall be equal to value of one of Name attributes of ConfigDescription elements returned by GetSupportedRules and GetSupportedAnalyticsModules command.

	//Type *QName `xml:"Type,attr,omitempty" json:"Type,omitempty"`
}

type ConfigDescription struct {

	// List describing the configuration parameters. The names of the parameters must be unique. If possible SimpleItems
	// should be used to transport the information to ease parsing of dynamically defined messages by a client
	// application.
	Parameters *ItemListDescription `xml:"Parameters,omitempty" json:"Parameters,omitempty"`

	Messages []struct {
		*MessageDescription

		// The topic of the message. For historical reason the element is named ParentTopic, but the full topic is expected.
		ParentTopic string `xml:"ParentTopic,omitempty" json:"ParentTopic,omitempty"`
	} `xml:"Messages,omitempty" json:"Messages,omitempty"`

	Extension *ConfigDescriptionExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// The Name attribute (e.g. "tt::LineDetector") uniquely identifies the type of rule, not a type definition in a schema.

	//Name *QName `xml:"Name,attr,omitempty" json:"Name,omitempty"`

	// The fixed attribute signals that it is not allowed to add or remove this type of configuration.

	Fixed bool `xml:"fixed,attr,omitempty" json:"fixed,omitempty"`

	// The maxInstances attribute signals the maximum number of instances per configuration.

	MaxInstances int32 `xml:"maxInstances,attr,omitempty" json:"maxInstances,omitempty"`
}

type ConfigDescriptionExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type SupportedRules struct {

	// Lists the location of all schemas that are referenced in the rules.
	RuleContentSchemaLocation []AnyURI `xml:"RuleContentSchemaLocation,omitempty" json:"RuleContentSchemaLocation,omitempty"`

	// List of rules supported by the Video Analytics configuration..
	RuleDescription []*ConfigDescription `xml:"RuleDescription,omitempty" json:"RuleDescription,omitempty"`

	Extension *SupportedRulesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Maximum number of concurrent instances.

	Limit int32 `xml:"Limit,attr,omitempty" json:"Limit,omitempty"`
}

type SupportedRulesExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type SupportedAnalyticsModules struct {

	// It optionally contains a list of URLs that provide the location of schema files.
	// These schema files describe the types and elements used in the analytics module descriptions.
	// Analytics module descriptions that reference types or elements imported from any ONVIF defined schema files
	// need not explicitly list those schema files.
	AnalyticsModuleContentSchemaLocation []AnyURI `xml:"AnalyticsModuleContentSchemaLocation,omitempty" json:"AnalyticsModuleContentSchemaLocation,omitempty"`

	AnalyticsModuleDescription []*ConfigDescription `xml:"AnalyticsModuleDescription,omitempty" json:"AnalyticsModuleDescription,omitempty"`

	Extension *SupportedAnalyticsModulesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Maximum number of concurrent instances.

	Limit int32 `xml:"Limit,attr,omitempty" json:"Limit,omitempty"`
}

type SupportedAnalyticsModulesExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PolylineArray struct {

	// Contains array of Polyline
	Segment []*Polyline `xml:"Segment,omitempty" json:"Segment,omitempty"`

	Extension *PolylineArrayExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PolylineArrayExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PolylineArrayConfiguration struct {

	// Contains PolylineArray configuration data
	PolylineArray *PolylineArray `xml:"PolylineArray,omitempty" json:"PolylineArray,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type MotionExpression struct {

	// Motion Expression data structure contains motion expression which is based on Scene Descriptor schema with XPATH syntax. The Type argument could allow introduction of different dialects
	Expression string `xml:"Expression,omitempty" json:"Expression,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	Type string `xml:"Type,attr,omitempty" json:"Type,omitempty"`
}

type MotionExpressionConfiguration struct {

	// Contains Rule MotionExpression configuration
	MotionExpression *MotionExpression `xml:"MotionExpression,omitempty" json:"MotionExpression,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type CellLayout struct {

	// Mapping of the cell grid to the Video frame. The cell grid is starting from the upper left corner and x dimension is going from left to right and the y dimension from up to down.
	Transformation *Transformation `xml:"Transformation,omitempty" json:"Transformation,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	// Number of columns of the cell grid (x dimension)

	Columns int32 `xml:"Columns,attr,omitempty" json:"Columns,omitempty"`

	// Number of rows of the cell grid (y dimension)

	Rows int32 `xml:"Rows,attr,omitempty" json:"Rows,omitempty"`
}

type PaneConfiguration struct {

	// Optional name of the pane configuration.
	PaneName string `xml:"PaneName,omitempty" json:"PaneName,omitempty"`

	// If the device has audio outputs, this element contains a pointer to the audio output that is associated with the pane. A client
	// can retrieve the available audio outputs of a device using the GetAudioOutputs command of the DeviceIO service.
	AudioOutputToken *ReferenceToken `xml:"AudioOutputToken,omitempty" json:"AudioOutputToken,omitempty"`

	// If the device has audio sources, this element contains a pointer to the audio source that is associated with this pane.
	// The audio connection from a decoder device to the NVT is established using the backchannel mechanism. A client can retrieve the available audio sources of a device using the GetAudioSources command of the
	// DeviceIO service.
	AudioSourceToken *ReferenceToken `xml:"AudioSourceToken,omitempty" json:"AudioSourceToken,omitempty"`

	// The configuration of the audio encoder including codec, bitrate
	// and sample rate.
	AudioEncoderConfiguration *AudioEncoderConfiguration `xml:"AudioEncoderConfiguration,omitempty" json:"AudioEncoderConfiguration,omitempty"`

	// A pointer to a Receiver that has the necessary information to receive
	// data from a Transmitter. This Receiver can be connected and the network video decoder displays the received data on the specified outputs. A client can retrieve the available Receivers using the
	// GetReceivers command of the Receiver Service.
	ReceiverToken *ReferenceToken `xml:"ReceiverToken,omitempty" json:"ReceiverToken,omitempty"`

	// A unique identifier in the display device.
	Token *ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PaneLayout struct {

	// Reference to the configuration of the streaming and coding parameters.
	Pane *ReferenceToken `xml:"Pane,omitempty" json:"Pane,omitempty"`

	// Describes the location and size of the area on the monitor. The area coordinate values are espressed in normalized units [-1.0, 1.0].
	Area *Rectangle `xml:"Area,omitempty" json:"Area,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type Layout struct {

	// List of panes assembling the display layout.
	PaneLayout []*PaneLayout `xml:"PaneLayout,omitempty" json:"PaneLayout,omitempty"`

	Extension *LayoutExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type LayoutExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type CodingCapabilities struct {

	// If the device supports audio encoding this section describes the supported codecs and their configuration.
	AudioEncodingCapabilities *AudioEncoderConfigurationOptions `xml:"AudioEncodingCapabilities,omitempty" json:"AudioEncodingCapabilities,omitempty"`

	// If the device supports audio decoding this section describes the supported codecs and their settings.
	AudioDecodingCapabilities *AudioDecoderConfigurationOptions `xml:"AudioDecodingCapabilities,omitempty" json:"AudioDecodingCapabilities,omitempty"`

	// This section describes the supported video codesc and their configuration.
	VideoDecodingCapabilities *VideoDecoderConfigurationOptions `xml:"VideoDecodingCapabilities,omitempty" json:"VideoDecodingCapabilities,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type LayoutOptions struct {

	// Lists the possible Pane Layouts of the Video Output
	PaneLayoutOptions []*PaneLayoutOptions `xml:"PaneLayoutOptions,omitempty" json:"PaneLayoutOptions,omitempty"`

	Extension *LayoutOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type LayoutOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PaneLayoutOptions struct {

	// List of areas assembling a layout. Coordinate values are in the range [-1.0, 1.0].
	Area []*Rectangle `xml:"Area,omitempty" json:"Area,omitempty"`

	Extension *PaneOptionExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type PaneOptionExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type Receiver struct {

	// Unique identifier of the receiver.
	Token *ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty"`

	// Describes the configuration of the receiver.
	Configuration *ReceiverConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ReceiverConfiguration struct {

	// The following connection modes are defined:
	Mode *ReceiverMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// Details of the URI to which the receiver should connect.
	MediaUri AnyURI `xml:"MediaUri,omitempty" json:"MediaUri,omitempty"`

	// Stream connection parameters.
	StreamSetup *StreamSetup `xml:"StreamSetup,omitempty" json:"StreamSetup,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ReceiverStateInformation struct {

	// The connection state of the receiver may have one of the following states:
	State *ReceiverState `xml:"State,omitempty" json:"State,omitempty"`

	// Indicates whether or not the receiver was created automatically.
	AutoCreated bool `xml:"AutoCreated,omitempty" json:"AutoCreated,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type SourceReference struct {
	Token *ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	Type AnyURI `xml:"Type,attr,omitempty" json:"Type,omitempty"`
}

type DateTimeRange struct {
	From soap.XSDDateTime `xml:"From,omitempty" json:"From,omitempty"`

	Until soap.XSDDateTime `xml:"Until,omitempty" json:"Until,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingSummary struct {

	// The earliest point in time where there is recorded data on the device.
	DataFrom soap.XSDDateTime `xml:"DataFrom,omitempty" json:"DataFrom,omitempty"`

	// The most recent point in time where there is recorded data on the device.
	DataUntil soap.XSDDateTime `xml:"DataUntil,omitempty" json:"DataUntil,omitempty"`

	// The device contains this many recordings.
	NumberRecordings int32 `xml:"NumberRecordings,omitempty" json:"NumberRecordings,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type SearchScope struct {

	// A list of sources that are included in the scope. If this list is included, only data from one of these sources shall be searched.
	IncludedSources []*SourceReference `xml:"IncludedSources,omitempty" json:"IncludedSources,omitempty"`

	// A list of recordings that are included in the scope. If this list is included, only data from one of these recordings shall be searched.
	IncludedRecordings []*RecordingReference `xml:"IncludedRecordings,omitempty" json:"IncludedRecordings,omitempty"`

	// An xpath expression used to specify what recordings to search. Only those recordings with an RecordingInformation structure that matches the filter shall be searched.
	RecordingInformationFilter *XPathExpression `xml:"RecordingInformationFilter,omitempty" json:"RecordingInformationFilter,omitempty"`

	// Extension point
	Extension *SearchScopeExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type SearchScopeExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type EventFilter struct {
	*FilterType
}

type PTZPositionFilter struct {

	// The lower boundary of the PTZ volume to look for.
	MinPosition *PTZVector `xml:"MinPosition,omitempty" json:"MinPosition,omitempty"`

	// The upper boundary of the PTZ volume to look for.
	MaxPosition *PTZVector `xml:"MaxPosition,omitempty" json:"MaxPosition,omitempty"`

	// If true, search for when entering the specified PTZ volume.
	EnterOrExit bool `xml:"EnterOrExit,omitempty" json:"EnterOrExit,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type MetadataFilter struct {
	MetadataStreamFilter *XPathExpression `xml:"MetadataStreamFilter,omitempty" json:"MetadataStreamFilter,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type FindRecordingResultList struct {

	// The state of the search when the result is returned. Indicates if there can be more results, or if the search is completed.
	SearchState *SearchState `xml:"SearchState,omitempty" json:"SearchState,omitempty"`

	// A RecordingInformation structure for each found recording matching the search.
	RecordingInformation []*RecordingInformation `xml:"RecordingInformation,omitempty" json:"RecordingInformation,omitempty"`
}

type FindEventResultList struct {

	// The state of the search when the result is returned. Indicates if there can be more results, or if the search is completed.
	SearchState *SearchState `xml:"SearchState,omitempty" json:"SearchState,omitempty"`

	// A FindEventResult structure for each found event matching the search.
	Result []*FindEventResult `xml:"Result,omitempty" json:"Result,omitempty"`
}

type FindEventResult struct {

	// The recording where this event was found. Empty string if no recording is associated with this event.
	RecordingToken *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty"`

	// A reference to the track where this event was found. Empty string if no track is associated with this event.
	TrackToken *TrackReference `xml:"TrackToken,omitempty" json:"TrackToken,omitempty"`

	// The time when the event occured.
	Time soap.XSDDateTime `xml:"Time,omitempty" json:"Time,omitempty"`

	// The description of the event.
	Event *NotificationMessageHolderType `xml:"Event,omitempty" json:"Event,omitempty"`

	// If true, indicates that the event is a virtual event generated for this particular search session to give the state of a property at the start time of the search.
	StartStateEvent bool `xml:"StartStateEvent,omitempty" json:"StartStateEvent,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type FindPTZPositionResultList struct {

	// The state of the search when the result is returned. Indicates if there can be more results, or if the search is completed.
	SearchState *SearchState `xml:"SearchState,omitempty" json:"SearchState,omitempty"`

	// A FindPTZPositionResult structure for each found PTZ position matching the search.
	Result []*FindPTZPositionResult `xml:"Result,omitempty" json:"Result,omitempty"`
}

type FindPTZPositionResult struct {

	// A reference to the recording containing the PTZ position.
	RecordingToken *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty"`

	// A reference to the metadata track containing the PTZ position.
	TrackToken *TrackReference `xml:"TrackToken,omitempty" json:"TrackToken,omitempty"`

	// The time when the PTZ position was valid.
	Time soap.XSDDateTime `xml:"Time,omitempty" json:"Time,omitempty"`

	// The PTZ position.
	Position *PTZVector `xml:"Position,omitempty" json:"Position,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type FindMetadataResultList struct {

	// The state of the search when the result is returned. Indicates if there can be more results, or if the search is completed.
	SearchState *SearchState `xml:"SearchState,omitempty" json:"SearchState,omitempty"`

	// A FindMetadataResult structure for each found set of Metadata matching the search.
	Result []*FindMetadataResult `xml:"Result,omitempty" json:"Result,omitempty"`
}

type FindMetadataResult struct {

	// A reference to the recording containing the metadata.
	RecordingToken *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty"`

	// A reference to the metadata track containing the matching metadata.
	TrackToken *TrackReference `xml:"TrackToken,omitempty" json:"TrackToken,omitempty"`

	// The point in time when the matching metadata occurs in the metadata track.
	Time soap.XSDDateTime `xml:"Time,omitempty" json:"Time,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingInformation struct {
	RecordingToken *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty"`

	// Information about the source of the recording. This gives a description of where the data in the recording comes from. Since a single
	// recording is intended to record related material, there is just one source. It is indicates the physical location or the
	// major data source for the recording. Currently the recordingconfiguration cannot describe each individual data source.
	Source *RecordingSourceInformation `xml:"Source,omitempty" json:"Source,omitempty"`

	EarliestRecording soap.XSDDateTime `xml:"EarliestRecording,omitempty" json:"EarliestRecording,omitempty"`

	LatestRecording soap.XSDDateTime `xml:"LatestRecording,omitempty" json:"LatestRecording,omitempty"`

	Content *Description `xml:"Content,omitempty" json:"Content,omitempty"`

	// Basic information about the track. Note that a track may represent a single contiguous time span or consist of multiple slices.
	Track []*TrackInformation `xml:"Track,omitempty" json:"Track,omitempty"`

	RecordingStatus *RecordingStatus `xml:"RecordingStatus,omitempty" json:"RecordingStatus,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingSourceInformation struct {

	// Identifier for the source chosen by the client that creates the structure.
	// This identifier is opaque to the device. Clients may use any type of URI for this field. A device shall support at least 128 characters.
	SourceId AnyURI `xml:"SourceId,omitempty" json:"SourceId,omitempty"`

	// Informative user readable name of the source, e.g. "Camera23". A device shall support at least 20 characters.
	Name *Name `xml:"Name,omitempty" json:"Name,omitempty"`

	// Informative description of the physical location of the source, e.g. the coordinates on a map.
	Location *Description `xml:"Location,omitempty" json:"Location,omitempty"`

	// Informative description of the source.
	Description *Description `xml:"Description,omitempty" json:"Description,omitempty"`

	// URI provided by the service supplying data to be recorded. A device shall support at least 128 characters.
	Address AnyURI `xml:"Address,omitempty" json:"Address,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingEncryption struct {

	// Key ID of the associated key for encryption.
	KID string `xml:"KID,omitempty" json:"KID,omitempty"`

	// Key for encrypting content.
	// The device shall not include this parameter when reading.
	Key []byte `xml:"Key,omitempty" json:"Key,omitempty"`

	// Optional list of track tokens to be encrypted.
	// If no track tokens are specified, all tracks are encrypted and no other encryption configurations shall exist for the recording.
	// Each track shall only be contained in one encryption configuration.
	Track []string `xml:"Track,omitempty" json:"Track,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	// Mode of encryption.
	// See tt:EncryptionMode for a list of definitions and capability trc:SupportedEncryptionModes for the supported encryption modes.

	Mode string `xml:"Mode,attr,omitempty" json:"Mode,omitempty"`
}

type RecordingTargetConfiguration struct {

	// Token of a storage configuration.
	Storage *ReferenceToken `xml:"Storage,omitempty" json:"Storage,omitempty"`

	// Format of the recording.
	// See tt:TargetFormat for a list of definitions and capability trc:SupportedTargetFormats for the supported formats.
	Format string `xml:"Format,omitempty" json:"Format,omitempty"`

	// Path prefix to be inserted in the object key.
	Prefix string `xml:"Prefix,omitempty" json:"Prefix,omitempty"`

	// Path postfix to be inserted in the object key.
	Postfix string `xml:"Postfix,omitempty" json:"Postfix,omitempty"`

	// Maximum duration of a span.
	SpanDuration *time.Duration `xml:"SpanDuration,omitempty" json:"SpanDuration,omitempty"`

	// Maximum duration of a segment.
	SegmentDuration *time.Duration `xml:"SegmentDuration,omitempty" json:"SegmentDuration,omitempty"`

	// Optional encryption configuration.
	// See capability trc:EncryptionEntryLimit for the number of supported entries.
	// By specifying multiple encryption entries per recording, different tracks can be encrypted with different configurations.
	// Each track shall only be contained in one encryption configuration.
	Encryption []*RecordingEncryption `xml:"Encryption,omitempty" json:"Encryption,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type TrackInformation struct {
	TrackToken *TrackReference `xml:"TrackToken,omitempty" json:"TrackToken,omitempty"`

	// Type of the track: "Video", "Audio" or "Metadata".
	// The track shall only be able to hold data of that type.
	TrackType *TrackType `xml:"TrackType,omitempty" json:"TrackType,omitempty"`

	// Informative description of the contents of the track.
	Description *Description `xml:"Description,omitempty" json:"Description,omitempty"`

	// The start date and time of the oldest recorded data in the track.
	DataFrom soap.XSDDateTime `xml:"DataFrom,omitempty" json:"DataFrom,omitempty"`

	// The stop date and time of the newest recorded data in the track.
	DataTo soap.XSDDateTime `xml:"DataTo,omitempty" json:"DataTo,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type MediaAttributes struct {

	// A reference to the recording that has these attributes.
	RecordingToken *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty"`

	// A set of attributes for each track.
	TrackAttributes []*TrackAttributes `xml:"TrackAttributes,omitempty" json:"TrackAttributes,omitempty"`

	// The attributes are valid from this point in time in the recording.
	From soap.XSDDateTime `xml:"From,omitempty" json:"From,omitempty"`

	// The attributes are valid until this point in time in the recording. Can be equal to 'From' to indicate that the attributes are only known to be valid for this particular point in time.
	Until soap.XSDDateTime `xml:"Until,omitempty" json:"Until,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type TrackAttributes struct {

	// The basic information about the track. Note that a track may represent a single contiguous time span or consist of multiple slices.
	TrackInformation *TrackInformation `xml:"TrackInformation,omitempty" json:"TrackInformation,omitempty"`

	// If the track is a video track, exactly one of this structure shall be present and contain the video attributes.
	VideoAttributes *VideoAttributes `xml:"VideoAttributes,omitempty" json:"VideoAttributes,omitempty"`

	// If the track is an audio track, exactly one of this structure shall be present and contain the audio attributes.
	AudioAttributes *AudioAttributes `xml:"AudioAttributes,omitempty" json:"AudioAttributes,omitempty"`

	// If the track is an metadata track, exactly one of this structure shall be present and contain the metadata attributes.
	MetadataAttributes *MetadataAttributes `xml:"MetadataAttributes,omitempty" json:"MetadataAttributes,omitempty"`

	Extension *TrackAttributesExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type TrackAttributesExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type VideoAttributes struct {

	// Average bitrate in kbps.
	Bitrate int32 `xml:"Bitrate,omitempty" json:"Bitrate,omitempty"`

	// The width of the video in pixels.
	Width int32 `xml:"Width,omitempty" json:"Width,omitempty"`

	// The height of the video in pixels.
	Height int32 `xml:"Height,omitempty" json:"Height,omitempty"`

	//
	// Video encoding of the track.  Use value from tt:VideoEncoding for MPEG4. Otherwise use values from tt:VideoEncodingMimeNames and
	//
	// .
	//
	Encoding string `xml:"Encoding,omitempty" json:"Encoding,omitempty"`

	// Average framerate in frames per second.
	Framerate float32 `xml:"Framerate,omitempty" json:"Framerate,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioAttributes struct {

	// The bitrate in kbps.
	Bitrate int32 `xml:"Bitrate,omitempty" json:"Bitrate,omitempty"`

	//
	// Audio encoding of the track.  Use values from tt:AudioEncoding for G711 and AAC. Otherwise use values from tt:AudioEncodingMimeNames and
	//
	// .
	//
	Encoding string `xml:"Encoding,omitempty" json:"Encoding,omitempty"`

	// The sample rate in kHz.
	Samplerate int32 `xml:"Samplerate,omitempty" json:"Samplerate,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type MetadataAttributes struct {

	// Indicates that there can be PTZ data in the metadata track in the specified time interval.
	CanContainPTZ bool `xml:"CanContainPTZ,omitempty" json:"CanContainPTZ,omitempty"`

	// Indicates that there can be analytics data in the metadata track in the specified time interval.
	CanContainAnalytics bool `xml:"CanContainAnalytics,omitempty" json:"CanContainAnalytics,omitempty"`

	// Indicates that there can be notifications in the metadata track in the specified time interval.
	CanContainNotifications bool `xml:"CanContainNotifications,omitempty" json:"CanContainNotifications,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	// List of all PTZ spaces active for recording. Note that events are only recorded on position changes and the actual point of recording may not necessarily contain an event of the specified type.

	PtzSpaces *StringAttrList `xml:"PtzSpaces,attr,omitempty" json:"PtzSpaces,omitempty"`
}

type RecordingConfiguration struct {

	// Information about the source of the recording.
	Source *RecordingSourceInformation `xml:"Source,omitempty" json:"Source,omitempty"`

	// Informative description of the source.
	Content *Description `xml:"Content,omitempty" json:"Content,omitempty"`

	// Sspecifies the maximum time that data in any track within the
	// recording shall be stored. The device shall delete any data older than the maximum retention
	// time. Such data shall not be accessible anymore. If the MaximumRetentionPeriod is set to 0,
	// the device shall not limit the retention time of stored data, except by resource constraints.
	// Whatever the value of MaximumRetentionTime, the device may automatically delete
	// recordings to free up storage space for new recordings.
	MaximumRetentionTime *time.Duration `xml:"MaximumRetentionTime,omitempty" json:"MaximumRetentionTime,omitempty"`

	// Optional external storage target configuration.
	Target *RecordingTargetConfiguration `xml:"Target,omitempty" json:"Target,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type TrackConfiguration struct {

	// Type of the track. It shall be equal to the strings “Video”,
	// “Audio” or “Metadata”. The track shall only be able to hold data of that type.
	TrackType *TrackType `xml:"TrackType,omitempty" json:"TrackType,omitempty"`

	// Informative description of the track.
	Description *Description `xml:"Description,omitempty" json:"Description,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type GetRecordingsResponseItem struct {

	// Token of the recording.
	RecordingToken *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty"`

	// Configuration of the recording.
	Configuration *RecordingConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty"`

	// List of tracks.
	Tracks *GetTracksResponseList `xml:"Tracks,omitempty" json:"Tracks,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type GetTracksResponseList struct {

	// Configuration of a track.
	Track []*GetTracksResponseItem `xml:"Track,omitempty" json:"Track,omitempty"`
}

type GetTracksResponseItem struct {

	// Token of the track.
	TrackToken *TrackReference `xml:"TrackToken,omitempty" json:"TrackToken,omitempty"`

	// Configuration of the track.
	Configuration *TrackConfiguration `xml:"Configuration,omitempty" json:"Configuration,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingJobConfiguration struct {

	// Identifies the recording to which this job shall store the received data.
	RecordingToken *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty"`

	//
	// The mode of the job. If it is idle, nothing shall happen. If it is active, the device shall try
	// to obtain data from the receivers. A client shall use GetRecordingJobState to determine if data transfer is really taking place.
	//
	// The only valid values for Mode shall be “Idle” and “Active”.
	//
	Mode *RecordingJobMode `xml:"Mode,omitempty" json:"Mode,omitempty"`

	// This shall be a non-negative number. If there are multiple recording jobs that store data to
	// the same track, the device will only store the data for the recording job with the highest
	// priority. The priority is specified per recording job, but the device shall determine the priority
	// of each track individually. If there are two recording jobs with the same priority, the device
	// shall record the data corresponding to the recording job that was activated the latest.
	Priority int32 `xml:"Priority,omitempty" json:"Priority,omitempty"`

	// Source of the recording.
	Source []*RecordingJobSource `xml:"Source,omitempty" json:"Source,omitempty"`

	Extension *RecordingJobConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// Optional filter defining on which event condition a recording job gets active.
	EventFilter *RecordingEventFilter `xml:"EventFilter,omitempty" json:"EventFilter,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`

	// This attribute adds an additional requirement for activating the recording job.
	// If this optional field is provided the job shall only record if the schedule exists and is active.

	ScheduleToken string `xml:"ScheduleToken,attr,omitempty" json:"ScheduleToken,omitempty"`
}

type RecordingEventFilter struct {
	Filter []struct {

		// Topic filter as defined in section 9.6.3 of the ONVIF Core Specification.
		Topic string `xml:"Topic,omitempty" json:"Topic,omitempty"`

		// Optional message source content filter as defined in section 9.4.4 of the ONVIF Core Specification.
		Source string `xml:"Source,omitempty" json:"Source,omitempty"`
	} `xml:"Filter,omitempty" json:"Filter,omitempty"`

	// Optional timespan to record before the actual event condition became active.
	Before *time.Duration `xml:"Before,omitempty" json:"Before,omitempty"`

	// Optional timespan to record after the actual event condition becomes inactive.
	After *time.Duration `xml:"After,omitempty" json:"After,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingJobConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingJobSource struct {

	// This field shall be a reference to the source of the data. The type of the source
	// is determined by the attribute Type in the SourceToken structure. If Type is
	// http://www.onvif.org/ver10/schema/Receiver, the token is a ReceiverReference. In this case
	// the device shall receive the data over the network. If Type is
	// http://www.onvif.org/ver10/schema/Profile, the token identifies a media profile, instructing the
	// device to obtain data from a profile that exists on the local device.
	SourceToken *SourceReference `xml:"SourceToken,omitempty" json:"SourceToken,omitempty"`

	// If this field is TRUE, and if the SourceToken is omitted, the device
	// shall create a receiver object (through the receiver service) and assign the
	// ReceiverReference to the SourceToken field. When retrieving the RecordingJobConfiguration
	// from the device, the AutoCreateReceiver field shall never be present.
	AutoCreateReceiver bool `xml:"AutoCreateReceiver,omitempty" json:"AutoCreateReceiver,omitempty"`

	// List of tracks associated with the recording.
	Tracks []*RecordingJobTrack `xml:"Tracks,omitempty" json:"Tracks,omitempty"`

	Extension *RecordingJobSourceExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type RecordingJobSourceExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingJobTrack struct {

	// If the received RTSP stream contains multiple tracks of the same type, the
	// SourceTag differentiates between those Tracks. This field can be ignored in case of recording a local source.
	SourceTag string `xml:"SourceTag,omitempty" json:"SourceTag,omitempty"`

	// The destination is the tracktoken of the track to which the device shall store the
	// received data.
	Destination *TrackReference `xml:"Destination,omitempty" json:"Destination,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingJobStateInformation struct {

	// Identification of the recording that the recording job records to.
	RecordingToken *RecordingReference `xml:"RecordingToken,omitempty" json:"RecordingToken,omitempty"`

	// Holds the aggregated state over the whole RecordingJobInformation structure.
	State *RecordingJobState `xml:"State,omitempty" json:"State,omitempty"`

	// Identifies the data source of the recording job.
	Sources []*RecordingJobStateSource `xml:"Sources,omitempty" json:"Sources,omitempty"`

	Extension *RecordingJobStateInformationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type RecordingJobStateInformationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingJobStateSource struct {

	// Identifies the data source of the recording job.
	SourceToken *SourceReference `xml:"SourceToken,omitempty" json:"SourceToken,omitempty"`

	// Holds the aggregated state over all substructures of RecordingJobStateSource.
	State *RecordingJobState `xml:"State,omitempty" json:"State,omitempty"`

	// List of track items.
	Tracks *RecordingJobStateTracks `xml:"Tracks,omitempty" json:"Tracks,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type RecordingJobStateTracks struct {
	Track []*RecordingJobStateTrack `xml:"Track,omitempty" json:"Track,omitempty"`
}

type RecordingJobStateTrack struct {

	// Identifies the track of the data source that provides the data.
	SourceTag string `xml:"SourceTag,omitempty" json:"SourceTag,omitempty"`

	// Indicates the destination track.
	Destination *TrackReference `xml:"Destination,omitempty" json:"Destination,omitempty"`

	// Optionally holds an implementation defined string value that describes the error.
	// The string should be in the English language.
	Error string `xml:"Error,omitempty" json:"Error,omitempty"`

	// Provides the job state of the track. The valid
	// values of state shall be “Idle”, “Active” and “Error”. If state equals “Error”, the Error field may be filled in with an implementation defined value.
	State *RecordingJobState `xml:"State,omitempty" json:"State,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type GetRecordingJobsResponseItem struct {
	JobToken *RecordingJobReference `xml:"JobToken,omitempty" json:"JobToken,omitempty"`

	JobConfiguration *RecordingJobConfiguration `xml:"JobConfiguration,omitempty" json:"JobConfiguration,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ReplayConfiguration struct {

	// The RTSP session timeout.
	SessionTimeout *time.Duration `xml:"SessionTimeout,omitempty" json:"SessionTimeout,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AnalyticsEngine struct {
	*ConfigurationEntity

	AnalyticsEngineConfiguration *AnalyticsDeviceEngineConfiguration `xml:"AnalyticsEngineConfiguration,omitempty" json:"AnalyticsEngineConfiguration,omitempty"`
}

type AnalyticsDeviceEngineConfiguration struct {
	EngineConfiguration []*EngineConfiguration `xml:"EngineConfiguration,omitempty" json:"EngineConfiguration,omitempty"`

	Extension *AnalyticsDeviceEngineConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type AnalyticsDeviceEngineConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type EngineConfiguration struct {
	VideoAnalyticsConfiguration *VideoAnalyticsConfiguration `xml:"VideoAnalyticsConfiguration,omitempty" json:"VideoAnalyticsConfiguration,omitempty"`

	AnalyticsEngineInputInfo *AnalyticsEngineInputInfo `xml:"AnalyticsEngineInputInfo,omitempty" json:"AnalyticsEngineInputInfo,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AnalyticsEngineInputInfo struct {
	InputInfo *Config `xml:"InputInfo,omitempty" json:"InputInfo,omitempty"`

	Extension *AnalyticsEngineInputInfoExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type AnalyticsEngineInputInfoExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type AnalyticsEngineInput struct {
	*ConfigurationEntity

	SourceIdentification *SourceIdentification `xml:"SourceIdentification,omitempty" json:"SourceIdentification,omitempty"`

	VideoInput *VideoEncoderConfiguration `xml:"VideoInput,omitempty" json:"VideoInput,omitempty"`

	MetadataInput *MetadataInput `xml:"MetadataInput,omitempty" json:"MetadataInput,omitempty"`
}

type SourceIdentification struct {
	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Token []*ReferenceToken `xml:"Token,omitempty" json:"Token,omitempty"`

	Extension *SourceIdentificationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type SourceIdentificationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type MetadataInput struct {
	MetadataConfig []*Config `xml:"MetadataConfig,omitempty" json:"MetadataConfig,omitempty"`

	Extension *MetadataInputExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type MetadataInputExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type AnalyticsEngineControl struct {
	*ConfigurationEntity

	// Token of the analytics engine (AnalyticsEngine) being controlled.
	EngineToken *ReferenceToken `xml:"EngineToken,omitempty" json:"EngineToken,omitempty"`

	// Token of the analytics engine configuration (VideoAnalyticsConfiguration) in effect.
	EngineConfigToken *ReferenceToken `xml:"EngineConfigToken,omitempty" json:"EngineConfigToken,omitempty"`

	// Tokens of the input (AnalyticsEngineInput) configuration applied.
	InputToken []*ReferenceToken `xml:"InputToken,omitempty" json:"InputToken,omitempty"`

	// Tokens of the receiver providing media input data. The order of ReceiverToken shall exactly match the order of InputToken.
	ReceiverToken []*ReferenceToken `xml:"ReceiverToken,omitempty" json:"ReceiverToken,omitempty"`

	Multicast *MulticastConfiguration `xml:"Multicast,omitempty" json:"Multicast,omitempty"`

	Subscription *Config `xml:"Subscription,omitempty" json:"Subscription,omitempty"`

	Mode *ModeOfOperation `xml:"Mode,omitempty" json:"Mode,omitempty"`
}

type AnalyticsStateInformation struct {

	// Token of the control object whose status is requested.
	AnalyticsEngineControlToken *ReferenceToken `xml:"AnalyticsEngineControlToken,omitempty" json:"AnalyticsEngineControlToken,omitempty"`

	State *AnalyticsState `xml:"State,omitempty" json:"State,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AnalyticsState struct {
	Error string `xml:"Error,omitempty" json:"Error,omitempty"`

	State string `xml:"State,omitempty" json:"State,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ActionEngineEventPayload struct {

	// Request Message
	RequestInfo *Envelope `xml:"RequestInfo,omitempty" json:"RequestInfo,omitempty"`

	// Response Message
	ResponseInfo *Envelope `xml:"ResponseInfo,omitempty" json:"ResponseInfo,omitempty"`

	// Fault Message
	Fault *Fault `xml:"Fault,omitempty" json:"Fault,omitempty"`

	Extension *ActionEngineEventPayloadExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ActionEngineEventPayloadExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioClassCandidate struct {

	// Indicates audio class label
	Type *AudioClassType `xml:"Type,omitempty" json:"Type,omitempty"`

	// A likelihood/probability that the corresponding audio event belongs to this class. The sum of the likelihoods shall NOT exceed 1
	Likelihood float32 `xml:"Likelihood,omitempty" json:"Likelihood,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type AudioClassDescriptor struct {

	// Array of audio class label and class probability
	ClassCandidate []*AudioClassCandidate `xml:"ClassCandidate,omitempty" json:"ClassCandidate,omitempty"`

	Extension *AudioClassDescriptorExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type AudioClassDescriptorExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ActiveConnection struct {
	CurrentBitrate float32 `xml:"CurrentBitrate,omitempty" json:"CurrentBitrate,omitempty"`

	CurrentFps float32 `xml:"CurrentFps,omitempty" json:"CurrentFps,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ProfileStatus struct {
	ActiveConnections []*ActiveConnection `xml:"ActiveConnections,omitempty" json:"ActiveConnections,omitempty"`

	Extension *ProfileStatusExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ProfileStatusExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type OSDReference struct {
	Value *ReferenceToken `xml:",chardata" json:"-,"`
}

type OSDPosConfiguration struct {

	//
	// For OSD position type, following are the pre-defined:
	//
	//
	Type string `xml:"Type,omitempty" json:"Type,omitempty"`

	Pos *Vector `xml:"Pos,omitempty" json:"Pos,omitempty"`

	Extension *OSDPosConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type OSDPosConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type OSDColor struct {
	Color *Color `xml:"Color,omitempty" json:"Color,omitempty"`

	Transparent int32 `xml:"Transparent,attr,omitempty" json:"Transparent,omitempty"`
}

type OSDTextConfiguration struct {

	//
	// The following OSD Text Type are defined:
	//
	//
	Type string `xml:"Type,omitempty" json:"Type,omitempty"`

	//
	// List of supported OSD date formats. This element shall be present when the value of Type field has Date or DateAndTime. The following DateFormat are defined:
	//
	//
	DateFormat string `xml:"DateFormat,omitempty" json:"DateFormat,omitempty"`

	//
	// List of supported OSD time formats. This element shall be present when the value of Type field has Time or DateAndTime. The following TimeFormat are defined:
	//
	//
	TimeFormat string `xml:"TimeFormat,omitempty" json:"TimeFormat,omitempty"`

	// Font size of the text in pt.
	FontSize int32 `xml:"FontSize,omitempty" json:"FontSize,omitempty"`

	// Font color of the text.
	FontColor *OSDColor `xml:"FontColor,omitempty" json:"FontColor,omitempty"`

	// Background color of the text.
	BackgroundColor *OSDColor `xml:"BackgroundColor,omitempty" json:"BackgroundColor,omitempty"`

	// The content of text to be displayed.
	PlainText string `xml:"PlainText,omitempty" json:"PlainText,omitempty"`

	Extension *OSDTextConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// This flag is applicable for Type Plain and defaults to true. When set to false the PlainText content will not be persistent across device reboots.

	IsPersistentText bool `xml:"IsPersistentText,attr,omitempty" json:"IsPersistentText,omitempty"`
}

type OSDTextConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type OSDImgConfiguration struct {

	// The URI of the image which to be displayed.
	ImgPath AnyURI `xml:"ImgPath,omitempty" json:"ImgPath,omitempty"`

	Extension *OSDImgConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type OSDImgConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type ColorspaceRange struct {
	X *FloatRange `xml:"X,omitempty" json:"X,omitempty"`

	Y *FloatRange `xml:"Y,omitempty" json:"Y,omitempty"`

	Z *FloatRange `xml:"Z,omitempty" json:"Z,omitempty"`

	// Acceptable values are the same as in tt:Color.
	Colorspace AnyURI `xml:"Colorspace,omitempty" json:"Colorspace,omitempty"`
}

type ColorOptions struct {

	// List the supported color.
	ColorList []*Color `xml:"ColorList,omitempty" json:"ColorList,omitempty"`

	// Define the range of color supported.
	ColorspaceRange []*ColorspaceRange `xml:"ColorspaceRange,omitempty" json:"ColorspaceRange,omitempty"`
}

type OSDColorOptions struct {

	// Optional list of supported colors.
	Color *ColorOptions `xml:"Color,omitempty" json:"Color,omitempty"`

	// Range of the transparent level. Larger means more tranparent.
	Transparent *IntRange `xml:"Transparent,omitempty" json:"Transparent,omitempty"`

	Extension *OSDColorOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type OSDColorOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type OSDTextOptions struct {

	// List of supported OSD text type. When a device indicates the supported number relating to Text type in MaximumNumberOfOSDs, the type shall be presented.
	Type []string `xml:"Type,omitempty" json:"Type,omitempty"`

	// Range of the font size value.
	FontSizeRange *IntRange `xml:"FontSizeRange,omitempty" json:"FontSizeRange,omitempty"`

	// List of supported date format.
	DateFormat []string `xml:"DateFormat,omitempty" json:"DateFormat,omitempty"`

	// List of supported time format.
	TimeFormat []string `xml:"TimeFormat,omitempty" json:"TimeFormat,omitempty"`

	// List of supported font color.
	FontColor *OSDColorOptions `xml:"FontColor,omitempty" json:"FontColor,omitempty"`

	// List of supported background color.
	BackgroundColor *OSDColorOptions `xml:"BackgroundColor,omitempty" json:"BackgroundColor,omitempty"`

	Extension *OSDTextOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type OSDTextOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type OSDImgOptions struct {

	// List of available image URIs.
	ImagePath []AnyURI `xml:"ImagePath,omitempty" json:"ImagePath,omitempty"`

	Extension *OSDImgOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`

	// List of supported image MIME types, such as "image/png".

	FormatsSupported *StringAttrList `xml:"FormatsSupported,attr,omitempty" json:"FormatsSupported,omitempty"`

	// The maximum size (in bytes) of the image that can be uploaded.

	MaxSize int32 `xml:"MaxSize,attr,omitempty" json:"MaxSize,omitempty"`

	// The maximum width (in pixels) of the image that can be uploaded.

	MaxWidth int32 `xml:"MaxWidth,attr,omitempty" json:"MaxWidth,omitempty"`

	// The maximum height (in pixels) of the image that can be uploaded.

	MaxHeight int32 `xml:"MaxHeight,attr,omitempty" json:"MaxHeight,omitempty"`
}

type OSDImgOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type OSDConfiguration struct {
	*DeviceEntity

	// Reference to the video source configuration.
	VideoSourceConfigurationToken *OSDReference `xml:"VideoSourceConfigurationToken,omitempty" json:"VideoSourceConfigurationToken,omitempty"`

	// Type of OSD.
	Type *OSDType `xml:"Type,omitempty" json:"Type,omitempty"`

	// Position configuration of OSD.
	Position *OSDPosConfiguration `xml:"Position,omitempty" json:"Position,omitempty"`

	// Text configuration of OSD. It shall be present when the value of Type field is Text.
	TextString *OSDTextConfiguration `xml:"TextString,omitempty" json:"TextString,omitempty"`

	// Image configuration of OSD. It shall be present when the value of Type field is Image
	Image *OSDImgConfiguration `xml:"Image,omitempty" json:"Image,omitempty"`

	Extension *OSDConfigurationExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type OSDConfigurationExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type MaximumNumberOfOSDs struct {
	Total int32 `xml:"Total,attr,omitempty" json:"Total,omitempty"`

	Image int32 `xml:"Image,attr,omitempty" json:"Image,omitempty"`

	PlainText int32 `xml:"PlainText,attr,omitempty" json:"PlainText,omitempty"`

	Date int32 `xml:"Date,attr,omitempty" json:"Date,omitempty"`

	Time int32 `xml:"Time,attr,omitempty" json:"Time,omitempty"`

	DateAndTime int32 `xml:"DateAndTime,attr,omitempty" json:"DateAndTime,omitempty"`
}

type OSDConfigurationOptions struct {

	// The maximum number of OSD configurations supported for the specified video source configuration. If the configuration does not support OSDs, this value shall be zero and the Type and PositionOption elements are ignored. If a device limits the number of instances by OSDType, it shall indicate the supported number for each type via the related attribute.
	MaximumNumberOfOSDs *MaximumNumberOfOSDs `xml:"MaximumNumberOfOSDs,omitempty" json:"MaximumNumberOfOSDs,omitempty"`

	// List supported type of OSD configuration. When a device indicates the supported number for each types in MaximumNumberOfOSDs, related type shall be presented. A device shall return Option element relating to listed type.
	Type []*OSDType `xml:"Type,omitempty" json:"Type,omitempty"`

	//
	// List available OSD position type. Following are the pre-defined:
	//
	//
	PositionOption []string `xml:"PositionOption,omitempty" json:"PositionOption,omitempty"`

	// Option of the OSD text configuration. This element shall be returned if the device is signaling the support for Text.
	TextOption *OSDTextOptions `xml:"TextOption,omitempty" json:"TextOption,omitempty"`

	// Option of the OSD image configuration. This element shall be returned if the device is signaling the support for Image.
	ImageOption *OSDImgOptions `xml:"ImageOption,omitempty" json:"ImageOption,omitempty"`

	Extension *OSDConfigurationOptionsExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type OSDConfigurationOptionsExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type FileProgress struct {

	// Exported file name
	FileName string `xml:"FileName,omitempty" json:"FileName,omitempty"`

	// Normalized percentage completion for uploading the exported file
	Progress float32 `xml:"Progress,omitempty" json:"Progress,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type ArrayOfFileProgress struct {

	// Exported file name and export progress information
	FileProgress []*FileProgress `xml:"FileProgress,omitempty" json:"FileProgress,omitempty"`

	Extension *ArrayOfFileProgressExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type ArrayOfFileProgressExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type StorageReferencePath struct {

	// identifier of an existing Storage Configuration.
	StorageToken *ReferenceToken `xml:"StorageToken,omitempty" json:"StorageToken,omitempty"`

	// gives the relative directory path on the storage
	RelativePath string `xml:"RelativePath,omitempty" json:"RelativePath,omitempty"`

	Extension *StorageReferencePathExtension `xml:"Extension,omitempty" json:"Extension,omitempty"`
}

type StorageReferencePathExtension struct {
	Items []string `xml:",any" json:"items,omitempty"`
}

type PolygonOptions struct {

	// True if the device supports defining a region only using Rectangle.
	// The rectangle points are still passed using a Polygon element if the device does not support polygon regions. In this case, the points provided in the Polygon element shall represent a rectangle.
	RectangleOnly bool `xml:"RectangleOnly,omitempty" json:"RectangleOnly,omitempty"`

	// Provides the minimum and maximum number of points that can be defined in the Polygon.
	// If RectangleOnly is not set to true, this parameter is required.
	VertexLimits *IntRange `xml:"VertexLimits,omitempty" json:"VertexLimits,omitempty"`

	Items []string `xml:",any" json:"items,omitempty"`
}

type PTZ interface {

	/* Returns the capabilities of the PTZ service. The result is returned in a typed answer. */
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	/*
		        Get the descriptions of the available PTZ Nodes.

		        A PTZ-capable device may have multiple PTZ Nodes. The PTZ Nodes may represent
						mechanical PTZ drivers, uploaded PTZ drivers or digital PTZ drivers. PTZ Nodes are the
						lowest level entities in the PTZ control API and reflect the supported PTZ capabilities. The
						PTZ Node is referenced either by its name or by its reference token.
	*/
	GetNodes(request *GetNodes) (*GetNodesResponse, error)

	GetNodesContext(ctx context.Context, request *GetNodes) (*GetNodesResponse, error)

	/* Get a specific PTZ Node identified by a reference
	   token or a name. */
	GetNode(request *GetNode) (*GetNodeResponse, error)

	GetNodeContext(ctx context.Context, request *GetNode) (*GetNodeResponse, error)

	/*
		        Get a specific PTZconfiguration from the device, identified by its reference token or name.

		        The default Position/Translation/Velocity Spaces are introduced to allow NVCs sending move
						requests without the need to specify a certain coordinate system. The default Speeds are
						introduced to control the speed of move requests (absolute, relative, preset), where no
						explicit speed has been set.

		        The allowed pan and tilt range for Pan/Tilt Limits is defined by a two-dimensional space range
						that is mapped to a specific Absolute Pan/Tilt Position Space. At least one Pan/Tilt Position
						Space is required by the PTZNode to support Pan/Tilt limits. The limits apply to all supported
						absolute, relative and continuous Pan/Tilt movements. The limits shall be checked within the
						coordinate system for which the limits have been specified. That means that even if
						movements are specified in a different coordinate system, the requested movements shall be
						transformed to the coordinate system of the limits where the limits can be checked. When a
						relative or continuous movements is specified, which would leave the specified limits, the PTZ
						unit has to move along the specified limits. The Zoom Limits have to be interpreted
						accordingly.
	*/
	GetConfiguration(request *GetConfiguration) (*GetConfigurationResponse, error)

	GetConfigurationContext(ctx context.Context, request *GetConfiguration) (*GetConfigurationResponse, error)

	/*
		        Get all the existing PTZConfigurations from the device.

		        The default Position/Translation/Velocity Spaces are introduced to allow NVCs sending move
						requests without the need to specify a certain coordinate system. The default Speeds are
						introduced to control the speed of move requests (absolute, relative, preset), where no
						explicit speed has been set.

		        The allowed pan and tilt range for Pan/Tilt Limits is defined by a two-dimensional space range
						that is mapped to a specific Absolute Pan/Tilt Position Space. At least one Pan/Tilt Position
						Space is required by the PTZNode to support Pan/Tilt limits. The limits apply to all supported
						absolute, relative and continuous Pan/Tilt movements. The limits shall be checked within the
						coordinate system for which the limits have been specified. That means that even if
						movements are specified in a different coordinate system, the requested movements shall be
						transformed to the coordinate system of the limits where the limits can be checked. When a
						relative or continuous movements is specified, which would leave the specified limits, the PTZ
						unit has to move along the specified limits. The Zoom Limits have to be interpreted
						accordingly.
	*/
	GetConfigurations(request *GetConfigurations) (*GetConfigurationsResponse, error)

	GetConfigurationsContext(ctx context.Context, request *GetConfigurations) (*GetConfigurationsResponse, error)

	/* Set/update a existing PTZConfiguration on the device. */
	SetConfiguration(request *SetConfiguration) (*SetConfigurationResponse, error)

	SetConfigurationContext(ctx context.Context, request *SetConfiguration) (*SetConfigurationResponse, error)

	/* List supported coordinate systems including their range limitations. Therefore, the options
	MAY differ depending on whether the PTZ Configuration is assigned to a Profile containing a
	Video Source Configuration. In that case, the options may additionally contain coordinate
	systems referring to the image coordinate system described by the Video Source
	Configuration. If the PTZ Node supports continuous movements, it shall return a Timeout Range within
	which Timeouts are accepted by the PTZ Node. */
	GetConfigurationOptions(request *GetConfigurationOptions) (*GetConfigurationOptionsResponse, error)

	GetConfigurationOptionsContext(ctx context.Context, request *GetConfigurationOptions) (*GetConfigurationOptionsResponse, error)

	/* Operation to send auxiliary commands to the PTZ device
	   mapped by the PTZNode in the selected profile. The
	   operation is supported
	   if the AuxiliarySupported element of the PTZNode is true */
	SendAuxiliaryCommand(request *SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, error)

	SendAuxiliaryCommandContext(ctx context.Context, request *SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, error)

	/* Operation to request all PTZ presets for the PTZNode
	   in the selected profile. The operation is supported if there is support
	   for at least on PTZ preset by the PTZNode. */
	GetPresets(request *GetPresets) (*GetPresetsResponse, error)

	GetPresetsContext(ctx context.Context, request *GetPresets) (*GetPresetsResponse, error)

	/* The SetPreset command saves the current device position parameters so that the device can
	move to the saved preset position through the GotoPreset operation.
	In order to create a new preset, the SetPresetRequest contains no PresetToken. If creation is
	successful, the Response contains the PresetToken which uniquely identifies the Preset. An
	existing Preset can be overwritten by specifying the PresetToken of the corresponding Preset.
	In both cases (overwriting or creation) an optional PresetName can be specified. The
	operation fails if the PTZ device is moving during the SetPreset operation.
	The device MAY internally save additional states such as imaging properties in the PTZ
	Preset which then should be recalled in the GotoPreset operation. */
	SetPreset(request *SetPreset) (*SetPresetResponse, error)

	SetPresetContext(ctx context.Context, request *SetPreset) (*SetPresetResponse, error)

	/* Operation to remove a PTZ preset for the Node in
	   the
	   selected profile. The operation is supported if the
	   PresetPosition
	   capability exists for teh Node in the
	   selected profile. */
	RemovePreset(request *RemovePreset) (*RemovePresetResponse, error)

	RemovePresetContext(ctx context.Context, request *RemovePreset) (*RemovePresetResponse, error)

	/* Operation to go to a saved preset position for the
	   PTZNode in the selected profile. The operation is supported if there is
	   support for at least on PTZ preset by the PTZNode. */
	GotoPreset(request *GotoPreset) (*GotoPresetResponse, error)

	GotoPresetContext(ctx context.Context, request *GotoPreset) (*GotoPresetResponse, error)

	/* Operation to move the PTZ device to it's "home" position. The operation is supported if the HomeSupported element in the PTZNode is true. */
	GotoHomePosition(request *GotoHomePosition) (*GotoHomePositionResponse, error)

	GotoHomePositionContext(ctx context.Context, request *GotoHomePosition) (*GotoHomePositionResponse, error)

	/* Operation to save current position as the home position.
	The SetHomePosition command returns with a failure if the “home” position is fixed and
	cannot be overwritten. If the SetHomePosition is successful, it is possible to recall the
	Home Position with the GotoHomePosition command. */
	SetHomePosition(request *SetHomePosition) (*SetHomePositionResponse, error)

	SetHomePositionContext(ctx context.Context, request *SetHomePosition) (*SetHomePositionResponse, error)

	/* Operation for continuous Pan/Tilt and Zoom movements. The operation is supported if the PTZNode supports at least one continuous Pan/Tilt or Zoom space. If the space argument is omitted, the default space set by the PTZConfiguration will be used. */
	ContinuousMove(request *ContinuousMove) (*ContinuousMoveResponse, error)

	ContinuousMoveContext(ctx context.Context, request *ContinuousMove) (*ContinuousMoveResponse, error)

	/*
		        Operation for Relative Pan/Tilt and Zoom Move. The operation is supported if the PTZNode supports at least one relative Pan/Tilt or Zoom space.

		        The speed argument is optional. If an x/y speed value is given it is up to the device to either use
						the x value as absolute resoluting speed vector or to map x and y to the component speed.
						If the speed argument is omitted, the default speed set by the PTZConfiguration will be used.
	*/
	RelativeMove(request *RelativeMove) (*RelativeMoveResponse, error)

	RelativeMoveContext(ctx context.Context, request *RelativeMove) (*RelativeMoveResponse, error)

	/* Operation to request PTZ status for the Node in the
	selected profile. */
	GetStatus(request *GetStatus) (*GetStatusResponse, error)

	GetStatusContext(ctx context.Context, request *GetStatus) (*GetStatusResponse, error)

	/*
		        Operation to move pan,tilt or zoom to a absolute destination.

		        The speed argument is optional. If an x/y speed value is given it is up to the device to either use
						the x value as absolute resoluting speed vector or to map x and y to the component speed.
						If the speed argument is omitted, the default speed set by the PTZConfiguration will be used.
	*/
	AbsoluteMove(request *AbsoluteMove) (*AbsoluteMoveResponse, error)

	AbsoluteMoveContext(ctx context.Context, request *AbsoluteMove) (*AbsoluteMoveResponse, error)

	/*
		        Operation to move pan,tilt or zoom to point to a destination based on the geolocation of the target.

		        The speed argument is optional. If an x/y speed value is given it is up to the device to either use
						the x value as absolute resoluting speed vector or to map x and y to the component speed.
						If the speed argument is omitted, the default speed set by the PTZConfiguration will be used.
						The area height and area dwidth parameters are optional, they can be used independently and may be used
						by the device to automatically determine the best zoom level to show the target.
	*/
	GeoMove(request *GeoMove) (*GeoMoveResponse, error)

	GeoMoveContext(ctx context.Context, request *GeoMove) (*GeoMoveResponse, error)

	/* Operation to stop ongoing pan, tilt and zoom movements of absolute relative and continuous type.
	If no stop argument for pan, tilt or zoom is set, the device will stop all ongoing pan, tilt and zoom movements. */
	Stop(request *Stop) (*StopResponse, error)

	StopContext(ctx context.Context, request *Stop) (*StopResponse, error)

	/* Operation to request PTZ preset tours in the selected media profiles. */
	GetPresetTours(request *GetPresetTours) (*GetPresetToursResponse, error)

	GetPresetToursContext(ctx context.Context, request *GetPresetTours) (*GetPresetToursResponse, error)

	/* Operation to request a specific PTZ preset tour in the selected media profile. */
	GetPresetTour(request *GetPresetTour) (*GetPresetTourResponse, error)

	GetPresetTourContext(ctx context.Context, request *GetPresetTour) (*GetPresetTourResponse, error)

	/* Operation to request available options to configure PTZ preset tour. */
	GetPresetTourOptions(request *GetPresetTourOptions) (*GetPresetTourOptionsResponse, error)

	GetPresetTourOptionsContext(ctx context.Context, request *GetPresetTourOptions) (*GetPresetTourOptionsResponse, error)

	/* Operation to create a preset tour for the selected media profile. */
	CreatePresetTour(request *CreatePresetTour) (*CreatePresetTourResponse, error)

	CreatePresetTourContext(ctx context.Context, request *CreatePresetTour) (*CreatePresetTourResponse, error)

	/* Operation to modify a preset tour for the selected media profile. */
	ModifyPresetTour(request *ModifyPresetTour) (*ModifyPresetTourResponse, error)

	ModifyPresetTourContext(ctx context.Context, request *ModifyPresetTour) (*ModifyPresetTourResponse, error)

	/* Operation to perform specific operation on the preset tour in selected media profile. */
	OperatePresetTour(request *OperatePresetTour) (*OperatePresetTourResponse, error)

	OperatePresetTourContext(ctx context.Context, request *OperatePresetTour) (*OperatePresetTourResponse, error)

	/* Operation to delete a specific preset tour from the media profile. */
	RemovePresetTour(request *RemovePresetTour) (*RemovePresetTourResponse, error)

	RemovePresetTourContext(ctx context.Context, request *RemovePresetTour) (*RemovePresetTourResponse, error)

	/*
		        Operation to get all available PTZConfigurations that can be added to the referenced media profile.

		        A device providing more than one PTZConfiguration or more than one VideoSourceConfiguration or which has any other resource
						interdependency between PTZConfiguration entities and other resources listable in a media profile should implement this operation.
						PTZConfiguration entities returned by this operation shall not fail on adding them to the referenced media profile.
	*/
	GetCompatibleConfigurations(request *GetCompatibleConfigurations) (*GetCompatibleConfigurationsResponse, error)

	GetCompatibleConfigurationsContext(ctx context.Context, request *GetCompatibleConfigurations) (*GetCompatibleConfigurationsResponse, error)

	/*
	   Operation to send an an atomic command to the device: move the camera to a wanted position and then delegate the PTZ control to the tracking algorithm.

	   An existing Speed argument overrides DefaultSpeed of the corresponding PTZ configuration during movement to the requested position.
	               If spaces are referenced within the Speed argument, they shall be speed spaces supported by the PTZ node.

	   If the detection and the tracking are done in the same device, an ObjectID reference can be passed as an argument, in order to specify which object should be tracked.

	   The operation shall fail if the requested absolute position is not reachable.
	*/
	MoveAndStartTracking(request *MoveAndStartTracking) (*MoveAndStartTrackingResponse, error)

	MoveAndStartTrackingContext(ctx context.Context, request *MoveAndStartTracking) (*MoveAndStartTrackingResponse, error)
}

type pTZ struct {
	client *soap.Client
}

func NewPTZ(client *soap.Client) PTZ {
	return &pTZ{
		client: client,
	}
}

func (service *pTZ) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetNodesContext(ctx context.Context, request *GetNodes) (*GetNodesResponse, error) {
	response := new(GetNodesResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetNodes(request *GetNodes) (*GetNodesResponse, error) {
	return service.GetNodesContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetNodeContext(ctx context.Context, request *GetNode) (*GetNodeResponse, error) {
	response := new(GetNodeResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetNode(request *GetNode) (*GetNodeResponse, error) {
	return service.GetNodeContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetConfigurationContext(ctx context.Context, request *GetConfiguration) (*GetConfigurationResponse, error) {
	response := new(GetConfigurationResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetConfiguration(request *GetConfiguration) (*GetConfigurationResponse, error) {
	return service.GetConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetConfigurationsContext(ctx context.Context, request *GetConfigurations) (*GetConfigurationsResponse, error) {
	response := new(GetConfigurationsResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetConfigurations(request *GetConfigurations) (*GetConfigurationsResponse, error) {
	return service.GetConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) SetConfigurationContext(ctx context.Context, request *SetConfiguration) (*SetConfigurationResponse, error) {
	response := new(SetConfigurationResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) SetConfiguration(request *SetConfiguration) (*SetConfigurationResponse, error) {
	return service.SetConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetConfigurationOptionsContext(ctx context.Context, request *GetConfigurationOptions) (*GetConfigurationOptionsResponse, error) {
	response := new(GetConfigurationOptionsResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetConfigurationOptions(request *GetConfigurationOptions) (*GetConfigurationOptionsResponse, error) {
	return service.GetConfigurationOptionsContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) SendAuxiliaryCommandContext(ctx context.Context, request *SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, error) {
	response := new(SendAuxiliaryCommandResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) SendAuxiliaryCommand(request *SendAuxiliaryCommand) (*SendAuxiliaryCommandResponse, error) {
	return service.SendAuxiliaryCommandContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetPresetsContext(ctx context.Context, request *GetPresets) (*GetPresetsResponse, error) {
	response := new(GetPresetsResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetPresets(request *GetPresets) (*GetPresetsResponse, error) {
	return service.GetPresetsContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) SetPresetContext(ctx context.Context, request *SetPreset) (*SetPresetResponse, error) {
	response := new(SetPresetResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) SetPreset(request *SetPreset) (*SetPresetResponse, error) {
	return service.SetPresetContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) RemovePresetContext(ctx context.Context, request *RemovePreset) (*RemovePresetResponse, error) {
	response := new(RemovePresetResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) RemovePreset(request *RemovePreset) (*RemovePresetResponse, error) {
	return service.RemovePresetContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GotoPresetContext(ctx context.Context, request *GotoPreset) (*GotoPresetResponse, error) {
	response := new(GotoPresetResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GotoPreset(request *GotoPreset) (*GotoPresetResponse, error) {
	return service.GotoPresetContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GotoHomePositionContext(ctx context.Context, request *GotoHomePosition) (*GotoHomePositionResponse, error) {
	response := new(GotoHomePositionResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GotoHomePosition(request *GotoHomePosition) (*GotoHomePositionResponse, error) {
	return service.GotoHomePositionContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) SetHomePositionContext(ctx context.Context, request *SetHomePosition) (*SetHomePositionResponse, error) {
	response := new(SetHomePositionResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) SetHomePosition(request *SetHomePosition) (*SetHomePositionResponse, error) {
	return service.SetHomePositionContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) ContinuousMoveContext(ctx context.Context, request *ContinuousMove) (*ContinuousMoveResponse, error) {
	response := new(ContinuousMoveResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) ContinuousMove(request *ContinuousMove) (*ContinuousMoveResponse, error) {
	return service.ContinuousMoveContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) RelativeMoveContext(ctx context.Context, request *RelativeMove) (*RelativeMoveResponse, error) {
	response := new(RelativeMoveResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) RelativeMove(request *RelativeMove) (*RelativeMoveResponse, error) {
	return service.RelativeMoveContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetStatusContext(ctx context.Context, request *GetStatus) (*GetStatusResponse, error) {
	response := new(GetStatusResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetStatus(request *GetStatus) (*GetStatusResponse, error) {
	return service.GetStatusContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) AbsoluteMoveContext(ctx context.Context, request *AbsoluteMove) (*AbsoluteMoveResponse, error) {
	response := new(AbsoluteMoveResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) AbsoluteMove(request *AbsoluteMove) (*AbsoluteMoveResponse, error) {
	return service.AbsoluteMoveContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GeoMoveContext(ctx context.Context, request *GeoMove) (*GeoMoveResponse, error) {
	response := new(GeoMoveResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GeoMove(request *GeoMove) (*GeoMoveResponse, error) {
	return service.GeoMoveContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) StopContext(ctx context.Context, request *Stop) (*StopResponse, error) {
	response := new(StopResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) Stop(request *Stop) (*StopResponse, error) {
	return service.StopContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetPresetToursContext(ctx context.Context, request *GetPresetTours) (*GetPresetToursResponse, error) {
	response := new(GetPresetToursResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetPresetTours(request *GetPresetTours) (*GetPresetToursResponse, error) {
	return service.GetPresetToursContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetPresetTourContext(ctx context.Context, request *GetPresetTour) (*GetPresetTourResponse, error) {
	response := new(GetPresetTourResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetPresetTour(request *GetPresetTour) (*GetPresetTourResponse, error) {
	return service.GetPresetTourContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetPresetTourOptionsContext(ctx context.Context, request *GetPresetTourOptions) (*GetPresetTourOptionsResponse, error) {
	response := new(GetPresetTourOptionsResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetPresetTourOptions(request *GetPresetTourOptions) (*GetPresetTourOptionsResponse, error) {
	return service.GetPresetTourOptionsContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) CreatePresetTourContext(ctx context.Context, request *CreatePresetTour) (*CreatePresetTourResponse, error) {
	response := new(CreatePresetTourResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) CreatePresetTour(request *CreatePresetTour) (*CreatePresetTourResponse, error) {
	return service.CreatePresetTourContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) ModifyPresetTourContext(ctx context.Context, request *ModifyPresetTour) (*ModifyPresetTourResponse, error) {
	response := new(ModifyPresetTourResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) ModifyPresetTour(request *ModifyPresetTour) (*ModifyPresetTourResponse, error) {
	return service.ModifyPresetTourContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) OperatePresetTourContext(ctx context.Context, request *OperatePresetTour) (*OperatePresetTourResponse, error) {
	response := new(OperatePresetTourResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) OperatePresetTour(request *OperatePresetTour) (*OperatePresetTourResponse, error) {
	return service.OperatePresetTourContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) RemovePresetTourContext(ctx context.Context, request *RemovePresetTour) (*RemovePresetTourResponse, error) {
	response := new(RemovePresetTourResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) RemovePresetTour(request *RemovePresetTour) (*RemovePresetTourResponse, error) {
	return service.RemovePresetTourContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) GetCompatibleConfigurationsContext(ctx context.Context, request *GetCompatibleConfigurations) (*GetCompatibleConfigurationsResponse, error) {
	response := new(GetCompatibleConfigurationsResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) GetCompatibleConfigurations(request *GetCompatibleConfigurations) (*GetCompatibleConfigurationsResponse, error) {
	return service.GetCompatibleConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *pTZ) MoveAndStartTrackingContext(ctx context.Context, request *MoveAndStartTracking) (*MoveAndStartTrackingResponse, error) {
	response := new(MoveAndStartTrackingResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pTZ) MoveAndStartTracking(request *MoveAndStartTracking) (*MoveAndStartTrackingResponse, error) {
	return service.MoveAndStartTrackingContext(
		context.Background(),
		request,
	)
}
