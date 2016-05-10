// Code generated by protoc-gen-go.
// source: contracts.proto
// DO NOT EDIT!

/*
Package contracts is a generated protocol buffer package.

It is generated from these files:
	contracts.proto

It has these top-level messages:
	RicardianContract
	Listing
	Order
	OrderConfirmation
	Rating
	Dispute
	DisputeResolution
	Refund
	Timestamp
	ID
	Signatures
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Listing_Metadata_CategorySub int32

const (
	Listing_Metadata_CATEGORYSUB_NOT_SET Listing_Metadata_CategorySub = 0
	Listing_Metadata_FIXED_PRICE         Listing_Metadata_CategorySub = 1
	Listing_Metadata_AUCTION             Listing_Metadata_CategorySub = 2
)

var Listing_Metadata_CategorySub_name = map[int32]string{
	0: "CATEGORYSUB_NOT_SET",
	1: "FIXED_PRICE",
	2: "AUCTION",
}
var Listing_Metadata_CategorySub_value = map[string]int32{
	"CATEGORYSUB_NOT_SET": 0,
	"FIXED_PRICE":         1,
	"AUCTION":             2,
}

func (x Listing_Metadata_CategorySub) String() string {
	return proto.EnumName(Listing_Metadata_CategorySub_name, int32(x))
}
func (Listing_Metadata_CategorySub) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 0, 0}
}

type Listing_Metadata_Category int32

const (
	Listing_Metadata_CATEGORY_NOT_SET Listing_Metadata_Category = 0
	Listing_Metadata_PHYSICAL_GOOD    Listing_Metadata_Category = 1
	Listing_Metadata_DIGITAL_GOOD     Listing_Metadata_Category = 2
	Listing_Metadata_SERVICE          Listing_Metadata_Category = 3
)

var Listing_Metadata_Category_name = map[int32]string{
	0: "CATEGORY_NOT_SET",
	1: "PHYSICAL_GOOD",
	2: "DIGITAL_GOOD",
	3: "SERVICE",
}
var Listing_Metadata_Category_value = map[string]int32{
	"CATEGORY_NOT_SET": 0,
	"PHYSICAL_GOOD":    1,
	"DIGITAL_GOOD":     2,
	"SERVICE":          3,
}

func (x Listing_Metadata_Category) String() string {
	return proto.EnumName(Listing_Metadata_Category_name, int32(x))
}
func (Listing_Metadata_Category) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 0, 1}
}

type Signatures_Section int32

const (
	Signatures_SECTION_NOT_SET    Signatures_Section = 0
	Signatures_LISTING            Signatures_Section = 1
	Signatures_ORDER              Signatures_Section = 2
	Signatures_ORDER_CONFIRMATION Signatures_Section = 3
	Signatures_RATING             Signatures_Section = 4
	Signatures_DISPUTE            Signatures_Section = 5
	Signatures_DISPUTE_RESOLUTION Signatures_Section = 6
	Signatures_REFUND             Signatures_Section = 7
)

var Signatures_Section_name = map[int32]string{
	0: "SECTION_NOT_SET",
	1: "LISTING",
	2: "ORDER",
	3: "ORDER_CONFIRMATION",
	4: "RATING",
	5: "DISPUTE",
	6: "DISPUTE_RESOLUTION",
	7: "REFUND",
}
var Signatures_Section_value = map[string]int32{
	"SECTION_NOT_SET":    0,
	"LISTING":            1,
	"ORDER":              2,
	"ORDER_CONFIRMATION": 3,
	"RATING":             4,
	"DISPUTE":            5,
	"DISPUTE_RESOLUTION": 6,
	"REFUND":             7,
}

func (x Signatures_Section) String() string {
	return proto.EnumName(Signatures_Section_name, int32(x))
}
func (Signatures_Section) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{10, 0} }

type RicardianContract struct {
	VendorListing           []*Listing         `protobuf:"bytes,1,rep,name=vendorListing" json:"vendorListing,omitempty"`
	BuyerOrder              *Order             `protobuf:"bytes,2,opt,name=buyerOrder" json:"buyerOrder,omitempty"`
	VendorOrderConfirmation *OrderConfirmation `protobuf:"bytes,3,opt,name=vendorOrderConfirmation" json:"vendorOrderConfirmation,omitempty"`
	BuyerRating             *Rating            `protobuf:"bytes,4,opt,name=buyerRating" json:"buyerRating,omitempty"`
	Dispute                 *Dispute           `protobuf:"bytes,5,opt,name=dispute" json:"dispute,omitempty"`
	DisputeResolution       *DisputeResolution `protobuf:"bytes,6,opt,name=disputeResolution" json:"disputeResolution,omitempty"`
	Refund                  *Refund            `protobuf:"bytes,7,opt,name=refund" json:"refund,omitempty"`
	Signatures              []*Signatures      `protobuf:"bytes,8,rep,name=signatures" json:"signatures,omitempty"`
}

func (m *RicardianContract) Reset()                    { *m = RicardianContract{} }
func (m *RicardianContract) String() string            { return proto.CompactTextString(m) }
func (*RicardianContract) ProtoMessage()               {}
func (*RicardianContract) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *RicardianContract) GetVendorListing() []*Listing {
	if m != nil {
		return m.VendorListing
	}
	return nil
}

func (m *RicardianContract) GetBuyerOrder() *Order {
	if m != nil {
		return m.BuyerOrder
	}
	return nil
}

func (m *RicardianContract) GetVendorOrderConfirmation() *OrderConfirmation {
	if m != nil {
		return m.VendorOrderConfirmation
	}
	return nil
}

func (m *RicardianContract) GetBuyerRating() *Rating {
	if m != nil {
		return m.BuyerRating
	}
	return nil
}

func (m *RicardianContract) GetDispute() *Dispute {
	if m != nil {
		return m.Dispute
	}
	return nil
}

func (m *RicardianContract) GetDisputeResolution() *DisputeResolution {
	if m != nil {
		return m.DisputeResolution
	}
	return nil
}

func (m *RicardianContract) GetRefund() *Refund {
	if m != nil {
		return m.Refund
	}
	return nil
}

func (m *RicardianContract) GetSignatures() []*Signatures {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type Listing struct {
	ListingName        string            `protobuf:"bytes,1,opt,name=listingName" json:"listingName,omitempty"`
	VendorID           *ID               `protobuf:"bytes,2,opt,name=vendorID" json:"vendorID,omitempty"`
	Metadata           *Listing_Metadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	Item               *Listing_Item     `protobuf:"bytes,4,opt,name=item" json:"item,omitempty"`
	Shipping           *Listing_Shipping `protobuf:"bytes,5,opt,name=shipping" json:"shipping,omitempty"`
	Moderators         []string          `protobuf:"bytes,6,rep,name=moderators" json:"moderators,omitempty"`
	TermsAndConditions string            `protobuf:"bytes,7,opt,name=termsAndConditions" json:"termsAndConditions,omitempty"`
	RefundPolicy       string            `protobuf:"bytes,8,opt,name=refundPolicy" json:"refundPolicy,omitempty"`
}

func (m *Listing) Reset()                    { *m = Listing{} }
func (m *Listing) String() string            { return proto.CompactTextString(m) }
func (*Listing) ProtoMessage()               {}
func (*Listing) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Listing) GetVendorID() *ID {
	if m != nil {
		return m.VendorID
	}
	return nil
}

func (m *Listing) GetMetadata() *Listing_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listing) GetItem() *Listing_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Listing) GetShipping() *Listing_Shipping {
	if m != nil {
		return m.Shipping
	}
	return nil
}

type Listing_Metadata struct {
	Version     uint32                       `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Category    Listing_Metadata_Category    `protobuf:"varint,2,opt,name=category,enum=Listing_Metadata_Category" json:"category,omitempty"`
	CategorySub Listing_Metadata_CategorySub `protobuf:"varint,3,opt,name=categorySub,enum=Listing_Metadata_CategorySub" json:"categorySub,omitempty"`
	Expiry      *Timestamp                   `protobuf:"bytes,4,opt,name=expiry" json:"expiry,omitempty"`
}

func (m *Listing_Metadata) Reset()                    { *m = Listing_Metadata{} }
func (m *Listing_Metadata) String() string            { return proto.CompactTextString(m) }
func (*Listing_Metadata) ProtoMessage()               {}
func (*Listing_Metadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *Listing_Metadata) GetExpiry() *Timestamp {
	if m != nil {
		return m.Expiry
	}
	return nil
}

type Listing_Item struct {
	Title          string                  `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Description    string                  `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	ProcessingTime string                  `protobuf:"bytes,3,opt,name=processingTime" json:"processingTime,omitempty"`
	PricePerUnit   *Listing_Price          `protobuf:"bytes,4,opt,name=pricePerUnit" json:"pricePerUnit,omitempty"`
	Nsfw           bool                    `protobuf:"varint,5,opt,name=nsfw" json:"nsfw,omitempty"`
	Tags           []string                `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	ImageHashes    []string                `protobuf:"bytes,7,rep,name=imageHashes" json:"imageHashes,omitempty"`
	SKU            string                  `protobuf:"bytes,8,opt,name=SKU,json=sKU" json:"SKU,omitempty"`
	Condition      string                  `protobuf:"bytes,9,opt,name=condition" json:"condition,omitempty"`
	Options        []*Listing_Item_Options `protobuf:"bytes,10,rep,name=options" json:"options,omitempty"`
}

func (m *Listing_Item) Reset()                    { *m = Listing_Item{} }
func (m *Listing_Item) String() string            { return proto.CompactTextString(m) }
func (*Listing_Item) ProtoMessage()               {}
func (*Listing_Item) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 1} }

func (m *Listing_Item) GetPricePerUnit() *Listing_Price {
	if m != nil {
		return m.PricePerUnit
	}
	return nil
}

func (m *Listing_Item) GetOptions() []*Listing_Item_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

type Listing_Item_Options struct {
	Name            string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Values          []string       `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
	PriceWithOption *Listing_Price `protobuf:"bytes,3,opt,name=priceWithOption" json:"priceWithOption,omitempty"`
}

func (m *Listing_Item_Options) Reset()                    { *m = Listing_Item_Options{} }
func (m *Listing_Item_Options) String() string            { return proto.CompactTextString(m) }
func (*Listing_Item_Options) ProtoMessage()               {}
func (*Listing_Item_Options) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 1, 0} }

func (m *Listing_Item_Options) GetPriceWithOption() *Listing_Price {
	if m != nil {
		return m.PriceWithOption
	}
	return nil
}

type Listing_Shipping struct {
	FreeShipping      bool                                `protobuf:"varint,1,opt,name=freeShipping" json:"freeShipping,omitempty"`
	Domestic          *Listing_Price                      `protobuf:"bytes,2,opt,name=domestic" json:"domestic,omitempty"`
	International     *Listing_Price                      `protobuf:"bytes,3,opt,name=international" json:"international,omitempty"`
	ShippingRegions   []CountryCode                       `protobuf:"varint,4,rep,name=shippingRegions,enum=CountryCode" json:"shippingRegions,omitempty"`
	EstimatedDelivery *Listing_Shipping_EstimatedDelivery `protobuf:"bytes,5,opt,name=estimatedDelivery" json:"estimatedDelivery,omitempty"`
	ShippingOrigin    CountryCode                         `protobuf:"varint,6,opt,name=shippingOrigin,enum=CountryCode" json:"shippingOrigin,omitempty"`
}

func (m *Listing_Shipping) Reset()                    { *m = Listing_Shipping{} }
func (m *Listing_Shipping) String() string            { return proto.CompactTextString(m) }
func (*Listing_Shipping) ProtoMessage()               {}
func (*Listing_Shipping) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 2} }

func (m *Listing_Shipping) GetDomestic() *Listing_Price {
	if m != nil {
		return m.Domestic
	}
	return nil
}

func (m *Listing_Shipping) GetInternational() *Listing_Price {
	if m != nil {
		return m.International
	}
	return nil
}

func (m *Listing_Shipping) GetEstimatedDelivery() *Listing_Shipping_EstimatedDelivery {
	if m != nil {
		return m.EstimatedDelivery
	}
	return nil
}

type Listing_Shipping_EstimatedDelivery struct {
	Domestic      string `protobuf:"bytes,1,opt,name=domestic" json:"domestic,omitempty"`
	International string `protobuf:"bytes,2,opt,name=international" json:"international,omitempty"`
}

func (m *Listing_Shipping_EstimatedDelivery) Reset()         { *m = Listing_Shipping_EstimatedDelivery{} }
func (m *Listing_Shipping_EstimatedDelivery) String() string { return proto.CompactTextString(m) }
func (*Listing_Shipping_EstimatedDelivery) ProtoMessage()    {}
func (*Listing_Shipping_EstimatedDelivery) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1, 2, 0}
}

type Listing_Price struct {
	Bitcoin uint32              `protobuf:"varint,1,opt,name=bitcoin" json:"bitcoin,omitempty"`
	Fiat    *Listing_Price_Fiat `protobuf:"bytes,2,opt,name=fiat" json:"fiat,omitempty"`
}

func (m *Listing_Price) Reset()                    { *m = Listing_Price{} }
func (m *Listing_Price) String() string            { return proto.CompactTextString(m) }
func (*Listing_Price) ProtoMessage()               {}
func (*Listing_Price) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 3} }

func (m *Listing_Price) GetFiat() *Listing_Price_Fiat {
	if m != nil {
		return m.Fiat
	}
	return nil
}

type Listing_Price_Fiat struct {
	CurrencyCode string  `protobuf:"bytes,1,opt,name=currencyCode" json:"currencyCode,omitempty"`
	Price        float32 `protobuf:"fixed32,2,opt,name=price" json:"price,omitempty"`
}

func (m *Listing_Price_Fiat) Reset()                    { *m = Listing_Price_Fiat{} }
func (m *Listing_Price_Fiat) String() string            { return proto.CompactTextString(m) }
func (*Listing_Price_Fiat) ProtoMessage()               {}
func (*Listing_Price_Fiat) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 3, 0} }

// TODO: complete other messages
type Order struct {
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type OrderConfirmation struct {
}

func (m *OrderConfirmation) Reset()                    { *m = OrderConfirmation{} }
func (m *OrderConfirmation) String() string            { return proto.CompactTextString(m) }
func (*OrderConfirmation) ProtoMessage()               {}
func (*OrderConfirmation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type Rating struct {
}

func (m *Rating) Reset()                    { *m = Rating{} }
func (m *Rating) String() string            { return proto.CompactTextString(m) }
func (*Rating) ProtoMessage()               {}
func (*Rating) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type Dispute struct {
}

func (m *Dispute) Reset()                    { *m = Dispute{} }
func (m *Dispute) String() string            { return proto.CompactTextString(m) }
func (*Dispute) ProtoMessage()               {}
func (*Dispute) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type DisputeResolution struct {
}

func (m *DisputeResolution) Reset()                    { *m = DisputeResolution{} }
func (m *DisputeResolution) String() string            { return proto.CompactTextString(m) }
func (*DisputeResolution) ProtoMessage()               {}
func (*DisputeResolution) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type Refund struct {
}

func (m *Refund) Reset()                    { *m = Refund{} }
func (m *Refund) String() string            { return proto.CompactTextString(m) }
func (*Refund) ProtoMessage()               {}
func (*Refund) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type Timestamp struct {
	SecondsFromUnixEpoch int64 `protobuf:"varint,1,opt,name=secondsFromUnixEpoch" json:"secondsFromUnixEpoch,omitempty"`
}

func (m *Timestamp) Reset()                    { *m = Timestamp{} }
func (m *Timestamp) String() string            { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()               {}
func (*Timestamp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type ID struct {
	Guid         string      `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	BlockchainID string      `protobuf:"bytes,2,opt,name=blockchainID" json:"blockchainID,omitempty"`
	Pubkeys      *ID_Pubkeys `protobuf:"bytes,3,opt,name=pubkeys" json:"pubkeys,omitempty"`
}

func (m *ID) Reset()                    { *m = ID{} }
func (m *ID) String() string            { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()               {}
func (*ID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *ID) GetPubkeys() *ID_Pubkeys {
	if m != nil {
		return m.Pubkeys
	}
	return nil
}

type ID_Pubkeys struct {
	Guid    []byte `protobuf:"bytes,1,opt,name=guid,proto3" json:"guid,omitempty"`
	Bitcoin []byte `protobuf:"bytes,2,opt,name=bitcoin,proto3" json:"bitcoin,omitempty"`
}

func (m *ID_Pubkeys) Reset()                    { *m = ID_Pubkeys{} }
func (m *ID_Pubkeys) String() string            { return proto.CompactTextString(m) }
func (*ID_Pubkeys) ProtoMessage()               {}
func (*ID_Pubkeys) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9, 0} }

type Signatures struct {
	Section Signatures_Section `protobuf:"varint,1,opt,name=section,enum=Signatures_Section" json:"section,omitempty"`
	Guid    []byte             `protobuf:"bytes,2,opt,name=guid,proto3" json:"guid,omitempty"`
	Bitcoin []byte             `protobuf:"bytes,3,opt,name=bitcoin,proto3" json:"bitcoin,omitempty"`
}

func (m *Signatures) Reset()                    { *m = Signatures{} }
func (m *Signatures) String() string            { return proto.CompactTextString(m) }
func (*Signatures) ProtoMessage()               {}
func (*Signatures) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func init() {
	proto.RegisterType((*RicardianContract)(nil), "RicardianContract")
	proto.RegisterType((*Listing)(nil), "Listing")
	proto.RegisterType((*Listing_Metadata)(nil), "Listing.Metadata")
	proto.RegisterType((*Listing_Item)(nil), "Listing.Item")
	proto.RegisterType((*Listing_Item_Options)(nil), "Listing.Item.Options")
	proto.RegisterType((*Listing_Shipping)(nil), "Listing.Shipping")
	proto.RegisterType((*Listing_Shipping_EstimatedDelivery)(nil), "Listing.Shipping.EstimatedDelivery")
	proto.RegisterType((*Listing_Price)(nil), "Listing.Price")
	proto.RegisterType((*Listing_Price_Fiat)(nil), "Listing.Price.Fiat")
	proto.RegisterType((*Order)(nil), "Order")
	proto.RegisterType((*OrderConfirmation)(nil), "OrderConfirmation")
	proto.RegisterType((*Rating)(nil), "Rating")
	proto.RegisterType((*Dispute)(nil), "Dispute")
	proto.RegisterType((*DisputeResolution)(nil), "DisputeResolution")
	proto.RegisterType((*Refund)(nil), "Refund")
	proto.RegisterType((*Timestamp)(nil), "Timestamp")
	proto.RegisterType((*ID)(nil), "ID")
	proto.RegisterType((*ID_Pubkeys)(nil), "ID.Pubkeys")
	proto.RegisterType((*Signatures)(nil), "Signatures")
	proto.RegisterEnum("Listing_Metadata_CategorySub", Listing_Metadata_CategorySub_name, Listing_Metadata_CategorySub_value)
	proto.RegisterEnum("Listing_Metadata_Category", Listing_Metadata_Category_name, Listing_Metadata_Category_value)
	proto.RegisterEnum("Signatures_Section", Signatures_Section_name, Signatures_Section_value)
}

var fileDescriptor1 = []byte{
	// 1238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x56, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0xc6, 0x96, 0x63, 0xc9, 0xc7, 0x89, 0x63, 0x6f, 0x4a, 0xeb, 0xf1, 0xf0, 0x53, 0xc4, 0x5f,
	0x81, 0x41, 0xcc, 0x04, 0x06, 0xb8, 0x6b, 0x83, 0xa5, 0xa4, 0x9a, 0xa6, 0xb1, 0x59, 0xd9, 0x40,
	0xaf, 0x3c, 0xb2, 0xb4, 0x71, 0x76, 0x6a, 0x4b, 0x1e, 0x49, 0x2e, 0xf5, 0x33, 0xc0, 0x15, 0x57,
	0x5c, 0xf0, 0x00, 0x3c, 0x01, 0x2f, 0xc3, 0x83, 0x70, 0xcb, 0xd9, 0xd5, 0xca, 0x91, 0xed, 0xe6,
	0x4e, 0xe7, 0x3b, 0xe7, 0xec, 0xd9, 0xfd, 0xce, 0x9f, 0xe0, 0x38, 0x88, 0xa3, 0x2c, 0xf1, 0x83,
	0x2c, 0xb5, 0x96, 0x49, 0x9c, 0xc5, 0x3d, 0x12, 0xc4, 0x2b, 0x44, 0xd6, 0x41, 0x1c, 0x32, 0x85,
	0x99, 0x7f, 0x6a, 0xd0, 0xa1, 0x3c, 0xf0, 0x93, 0x90, 0xfb, 0x51, 0x5f, 0x39, 0x10, 0x0b, 0x8e,
	0x5e, 0xb1, 0x28, 0x8c, 0x93, 0x4b, 0x9e, 0x66, 0x3c, 0x9a, 0x75, 0x2b, 0x0f, 0xb5, 0x47, 0xcd,
	0x53, 0xc3, 0x52, 0x32, 0xdd, 0x56, 0x93, 0x4f, 0x00, 0xa6, 0xab, 0x35, 0x4b, 0x06, 0x49, 0xc8,
	0x92, 0x6e, 0xf5, 0x61, 0x05, 0x8d, 0xeb, 0x96, 0x94, 0x68, 0x49, 0x43, 0x2e, 0xe1, 0x41, 0xee,
	0x28, 0x45, 0x0c, 0x77, 0xcd, 0x93, 0x85, 0x9f, 0xf1, 0x38, 0xea, 0x6a, 0xd2, 0x89, 0x58, 0x7b,
	0x1a, 0x7a, 0x97, 0x0b, 0xf9, 0x0c, 0x9a, 0xf2, 0x6c, 0xea, 0xcb, 0x3b, 0xd6, 0xe4, 0x09, 0xba,
	0x95, 0x8b, 0xb4, 0xac, 0x23, 0x26, 0xe8, 0x21, 0x4f, 0x97, 0xab, 0x8c, 0x75, 0x0f, 0xa4, 0x99,
	0x61, 0xd9, 0xb9, 0x4c, 0x0b, 0x05, 0x79, 0x02, 0x1d, 0xf5, 0x49, 0x59, 0x1a, 0xcf, 0x57, 0xf2,
	0x5a, 0x75, 0x75, 0x2d, 0x7b, 0x57, 0x43, 0xf7, 0x8d, 0xc9, 0xfb, 0x50, 0x4f, 0xd8, 0xf5, 0x2a,
	0x0a, 0xbb, 0x7a, 0x71, 0x17, 0x29, 0x52, 0x05, 0x93, 0x2f, 0x00, 0x52, 0x3e, 0x8b, 0xfc, 0x6c,
	0x95, 0xb0, 0xb4, 0x6b, 0x48, 0x52, 0x9b, 0x96, 0xb7, 0x81, 0x68, 0x49, 0x6d, 0xfe, 0x71, 0x04,
	0x7a, 0x41, 0xf0, 0x43, 0x68, 0xce, 0xf3, 0xcf, 0x2b, 0x7f, 0xc1, 0x30, 0x1d, 0x95, 0x47, 0x0d,
	0x5a, 0x86, 0x30, 0xb6, 0x91, 0xf3, 0xe4, 0xda, 0x2a, 0x01, 0x9a, 0xe5, 0xda, 0x74, 0x03, 0x92,
	0x2f, 0xc1, 0x58, 0xb0, 0xcc, 0x0f, 0xfd, 0xcc, 0x57, 0x64, 0x77, 0x8a, 0x74, 0x5a, 0xcf, 0x95,
	0x82, 0x6e, 0x4c, 0xc8, 0x07, 0x50, 0xe3, 0x19, 0x5b, 0x28, 0x56, 0x8f, 0x36, 0xa6, 0x2e, 0x82,
	0x54, 0xaa, 0xc4, 0x89, 0xe9, 0x0d, 0x5f, 0x2e, 0x05, 0xf9, 0x07, 0x3b, 0x27, 0x7a, 0x4a, 0x41,
	0x37, 0x26, 0xe4, 0x3d, 0x80, 0x05, 0x56, 0x5e, 0xe2, 0x67, 0x71, 0x92, 0x22, 0xb1, 0x1a, 0x3e,
	0xa1, 0x84, 0x60, 0xd1, 0x91, 0x8c, 0x25, 0x8b, 0xf4, 0x2c, 0x0a, 0x31, 0xcd, 0x21, 0x17, 0x94,
	0xa6, 0x92, 0xc9, 0x06, 0x7d, 0x83, 0x06, 0x73, 0x7a, 0x98, 0xd3, 0x3a, 0x8c, 0xe7, 0x3c, 0x58,
	0x23, 0x9d, 0xc2, 0x72, 0x0b, 0xeb, 0xfd, 0x5b, 0x05, 0xa3, 0x78, 0x1c, 0xe9, 0x82, 0xfe, 0x8a,
	0x25, 0xa9, 0x48, 0xab, 0x20, 0xf0, 0x88, 0x16, 0x22, 0xf9, 0x16, 0x8c, 0xc0, 0xcf, 0xd8, 0x2c,
	0x4e, 0xd6, 0x92, 0xbc, 0xd6, 0x69, 0x6f, 0x8f, 0x1b, 0xab, 0xaf, 0x2c, 0xe8, 0xc6, 0x96, 0x3c,
	0x86, 0x66, 0xf1, 0xed, 0xad, 0xa6, 0x92, 0xd6, 0xd6, 0xe9, 0xbb, 0x77, 0xbb, 0xa2, 0x11, 0x2d,
	0x7b, 0xe0, 0x1b, 0xea, 0xec, 0xf5, 0x92, 0x63, 0xd8, 0x9c, 0x67, 0xb0, 0x46, 0x7c, 0xc1, 0xd2,
	0xcc, 0x5f, 0x2c, 0xa9, 0xd2, 0x98, 0x36, 0x34, 0x4b, 0xfe, 0xe4, 0x01, 0x9c, 0xf4, 0xcf, 0x46,
	0xce, 0xc5, 0x80, 0xbe, 0xf0, 0xc6, 0x3f, 0x4c, 0xae, 0x06, 0xa3, 0x89, 0xe7, 0x8c, 0xda, 0x6f,
	0x91, 0x63, 0x68, 0x9e, 0xbb, 0xbf, 0x38, 0xf6, 0x64, 0x48, 0xdd, 0xbe, 0xd3, 0xae, 0x90, 0x26,
	0xe8, 0x67, 0xe3, 0xfe, 0xc8, 0x1d, 0x5c, 0xb5, 0xab, 0x26, 0x05, 0xa3, 0x38, 0x85, 0xdc, 0x83,
	0x76, 0x71, 0x44, 0xc9, 0xbf, 0x03, 0x47, 0xc3, 0xa7, 0x2f, 0x3c, 0xb7, 0x7f, 0x76, 0x39, 0xb9,
	0x18, 0x0c, 0x6c, 0x3c, 0xa1, 0x0d, 0x87, 0xb6, 0x7b, 0xe1, 0x8e, 0x0a, 0xa4, 0x2a, 0xce, 0xf4,
	0x1c, 0xfa, 0x93, 0x08, 0xa0, 0xf5, 0xfe, 0xd1, 0xa0, 0x26, 0xea, 0x01, 0x0f, 0x3c, 0xc8, 0x78,
	0x36, 0x2f, 0x0a, 0x33, 0x17, 0x44, 0xd1, 0xe2, 0xa0, 0x09, 0x12, 0xbe, 0x94, 0xad, 0x54, 0xcd,
	0x8b, 0xb6, 0x04, 0xe1, 0xdc, 0x68, 0xe1, 0x18, 0x0a, 0x58, 0x9a, 0x22, 0x5d, 0xe2, 0xe5, 0x92,
	0xc2, 0x06, 0xdd, 0x41, 0xc9, 0x29, 0x1c, 0x2e, 0x13, 0x1e, 0xb0, 0x21, 0x4b, 0xc6, 0x11, 0xcf,
	0x14, 0x59, 0xad, 0x0d, 0xd1, 0x43, 0xa1, 0xa4, 0x5b, 0x36, 0x84, 0x40, 0x2d, 0x4a, 0xaf, 0x7f,
	0x95, 0x95, 0x69, 0x50, 0xf9, 0x2d, 0xb0, 0xcc, 0x9f, 0x15, 0xc5, 0x27, 0xbf, 0xc5, 0x2d, 0xf9,
	0xc2, 0x9f, 0xb1, 0xa7, 0x7e, 0x7a, 0xc3, 0x44, 0xbd, 0x09, 0x55, 0x19, 0x42, 0x16, 0x34, 0xef,
	0xd9, 0x58, 0xd5, 0x97, 0x96, 0x3e, 0x1b, 0x93, 0x77, 0xa0, 0x11, 0x14, 0x85, 0xd8, 0x6d, 0x48,
	0xfc, 0x16, 0x20, 0x5f, 0x81, 0x1e, 0x2f, 0xf3, 0xea, 0x05, 0xd9, 0xe2, 0x6f, 0x6f, 0x75, 0x8f,
	0x35, 0xc8, 0x95, 0xb4, 0xb0, 0xea, 0xc5, 0xa0, 0x2b, 0x4c, 0xde, 0xfa, 0xb6, 0xc3, 0xe5, 0x37,
	0xb9, 0x0f, 0xf5, 0x57, 0xfe, 0x7c, 0x85, 0x97, 0xab, 0xca, 0xcb, 0x29, 0x89, 0x7c, 0x0f, 0xc7,
	0xf2, 0xc5, 0x3f, 0xf3, 0xec, 0x26, 0xf7, 0x57, 0x8d, 0xbd, 0x4b, 0xcc, 0xae, 0x59, 0xef, 0x6f,
	0x0d, 0x8c, 0xa2, 0x43, 0x45, 0x1f, 0x5d, 0x27, 0x8c, 0x15, 0xb2, 0x0c, 0x6d, 0xd0, 0x2d, 0x8c,
	0x7c, 0x0e, 0x46, 0x18, 0x8b, 0xc2, 0xe4, 0x81, 0x9a, 0x2e, 0xbb, 0x31, 0x36, 0x7a, 0xf2, 0x0d,
	0x1c, 0xf1, 0x08, 0xfb, 0x35, 0x92, 0x53, 0xda, 0x9f, 0xdf, 0x71, 0xa9, 0x6d, 0x23, 0x6c, 0xc1,
	0xe3, 0x62, 0x52, 0x50, 0x36, 0x93, 0xe4, 0xd5, 0xf0, 0xb5, 0xad, 0xd3, 0x43, 0xab, 0x9f, 0xaf,
	0xad, 0x3e, 0x8e, 0x0a, 0xba, 0x6b, 0x44, 0x7e, 0x84, 0x8e, 0x08, 0x8b, 0x2b, 0x81, 0x85, 0x36,
	0x9b, 0x73, 0x6c, 0xe9, 0xb5, 0x9a, 0x46, 0x1f, 0xee, 0x4d, 0x23, 0xcb, 0xd9, 0x35, 0xa5, 0xfb,
	0xde, 0xf8, 0x80, 0x56, 0x11, 0x65, 0x90, 0xf0, 0x19, 0xcf, 0xb7, 0xc0, 0xee, 0x4d, 0x76, 0x6c,
	0x7a, 0x63, 0xe8, 0xec, 0x9d, 0x4e, 0x7a, 0x25, 0xde, 0xf2, 0x94, 0xde, 0xf2, 0xf4, 0xd1, 0x2e,
	0x4f, 0x79, 0x83, 0x6c, 0x83, 0xbd, 0xdf, 0x2b, 0x70, 0x20, 0x09, 0x13, 0xe3, 0x6b, 0xca, 0xb3,
	0x20, 0xe6, 0x9b, 0xf1, 0xa5, 0x44, 0xf2, 0x29, 0xd4, 0xae, 0xb9, 0x9f, 0xa9, 0xcc, 0x9c, 0x6c,
	0x13, 0x6d, 0x9d, 0xa3, 0x8a, 0x4a, 0x83, 0xde, 0x13, 0xa8, 0x09, 0x49, 0xa4, 0x3c, 0x58, 0x25,
	0x09, 0x8b, 0x02, 0xf9, 0x16, 0x75, 0xb5, 0x2d, 0x4c, 0xf4, 0xb4, 0x2c, 0x1b, 0x79, 0x6a, 0x95,
	0xe6, 0x82, 0xa9, 0xc3, 0x81, 0x5c, 0xc4, 0xe6, 0x09, 0x74, 0xf6, 0x36, 0xb2, 0x69, 0x40, 0x3d,
	0x5f, 0xb8, 0x66, 0x03, 0x74, 0xb5, 0x32, 0x85, 0xe5, 0xde, 0xf6, 0x94, 0x96, 0x72, 0x50, 0x9b,
	0x8f, 0xa1, 0xb1, 0x99, 0x79, 0xd8, 0xe8, 0xf7, 0x52, 0x26, 0x3a, 0x29, 0x3d, 0x4f, 0xe2, 0x05,
	0xf6, 0xf1, 0x6b, 0x67, 0x19, 0x07, 0x37, 0xf2, 0x82, 0x1a, 0x7d, 0xa3, 0xce, 0xfc, 0xab, 0x02,
	0x55, 0xdc, 0x6f, 0xd8, 0x39, 0xb3, 0x15, 0x0f, 0x8b, 0xce, 0x11, 0xdf, 0xe2, 0x9d, 0xd3, 0x79,
	0x1c, 0xbc, 0x0c, 0x6e, 0x7c, 0x1e, 0xa9, 0xc5, 0x88, 0xef, 0x2c, 0x63, 0xe4, 0x63, 0xd0, 0x97,
	0xab, 0xe9, 0x4b, 0xb6, 0x4e, 0x55, 0xa1, 0x36, 0x71, 0x6f, 0x5a, 0xc3, 0x1c, 0xa2, 0x85, 0xae,
	0xf7, 0x1d, 0xe8, 0x0a, 0xdb, 0x8a, 0x74, 0xa8, 0x22, 0x95, 0x92, 0x53, 0x95, 0x70, 0x21, 0x9a,
	0xff, 0x55, 0x00, 0x6e, 0x37, 0x3c, 0x2e, 0x4d, 0x1d, 0x5f, 0x91, 0x15, 0x4b, 0xa8, 0x85, 0xe9,
	0xba, 0xd5, 0x5a, 0x5e, 0xae, 0xa2, 0x85, 0xcd, 0x26, 0x56, 0xf5, 0xcd, 0xb1, 0xb4, 0xed, 0x58,
	0xbf, 0x55, 0x70, 0x3c, 0x2b, 0xcf, 0x13, 0x38, 0xf6, 0x1c, 0x39, 0xfd, 0x4b, 0x33, 0x1e, 0xc7,
	0xf7, 0xa5, 0xeb, 0x8d, 0xdc, 0xab, 0x0b, 0x9c, 0xee, 0x0d, 0xcc, 0x25, 0xb5, 0x1d, 0x8a, 0x63,
	0xfd, 0x3e, 0x10, 0xf9, 0x39, 0xe9, 0x0f, 0xae, 0xce, 0x5d, 0xfa, 0xfc, 0x4c, 0x6e, 0x0d, 0x8d,
	0x00, 0xa6, 0xe9, 0x4c, 0x9a, 0xd7, 0x84, 0xaf, 0xed, 0x7a, 0xc3, 0xf1, 0xc8, 0x69, 0x1f, 0x08,
	0x07, 0x25, 0x4c, 0xa8, 0xe3, 0x0d, 0x2e, 0xc7, 0xd2, 0xa1, 0x2e, 0x1d, 0x9c, 0xf3, 0xf1, 0x95,
	0xdd, 0xd6, 0xa7, 0x75, 0xf9, 0x8b, 0xf9, 0xf5, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0x2f,
	0x2c, 0x2e, 0x89, 0x0a, 0x00, 0x00,
}