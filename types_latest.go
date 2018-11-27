package tdlib

//
type TlObjectIdType uint32

type TlReader struct {
}

func (rd *TlReader) ReadUInt32() uint32 {
	return 0
}

func (rd *TlReader) ReadInt32() int32 {
	return 0
}

func (rd *TlReader) ReadObjectId() TlObjectIdType {
	return TlObjectIdType(0)
}

type TlWriter struct {
}

type TlObject interface {
	GetObjectTypeId() TlObjectIdType
}

type TlObjectSerializer interface {
	TlSerialize(wr *TlWriter)
	TlDeserialize(rd *TlReader)
}

//

const (
	EnumUserTypeBot     TlObjectIdType = 1
	EnumUserTypeRegular TlObjectIdType = 2
	EnumUserTypeDeleted TlObjectIdType = 3
	EnumUserTypeUnknown TlObjectIdType = 4
)

type UserTypeInterface interface {
	UserTypeObjectId() TlObjectIdType
}

type X struct {
	a int32
	b UserTypeInterface
	c int32
}

func Serialize_X(x *X, wr *TlWriter) {
	// wr.WriteUInt32(x.GetObejctTypeId())
	// x.Serialize(wr)
}

func Deserialize_X(rd *TlReader) *X {
	rd.ReadObjectId() // typeId
	x := X{}
	x.Desrialize(rd)
	return &x
}

func (x *X) Serialize(wr *TlWriter) {
	//	x.a = rd.ReadInt32()
	//	x.b = DesrializeUserType(rd)
	//	x.c = rd.ReadInt32()
}

func (x *X) Desrialize(rd *TlReader) {
	x.a = rd.ReadInt32()
	x.b = DesrializeUserType(rd)
	x.c = rd.ReadInt32()
}

func DesrializeUserType(rd *TlReader) UserTypeInterface {

	typeId := rd.ReadObjectId()

	switch typeId {
	case EnumUserTypeRegular:
		u := UserTypeRegular{}
		return &u
	case EnumUserTypeDeleted:
		u := UserTypeDeleted{}
		return &u
		// ...
	}

	return nil
}

type UserTypeRegular struct {
}

func (o *UserTypeRegular) GetObjectTypeId() TlObjectIdType {
	return EnumUserTypeRegular
}

func (o *UserTypeRegular) UserTypeObjectId() TlObjectIdType {
	return o.GetObjectTypeId()
}

func NewUserTypeRegular() *UserTypeRegular {
	return &UserTypeRegular{}
}

func (*UserTypeRegular) Serialize(wr *TlWriter) {
	// a = rd.ReadInt32()
	// b = rd.ReadString()
	// ...
}

func (*UserTypeRegular) Desrialize(rd *TlReader) {
	// wr.Write(a)
	// wr.Write(b)
	// ...
}

type UserTypeDeleted struct {
}

func (o *UserTypeDeleted) GetObjectTypeId() TlObjectIdType {
	return EnumUserTypeDeleted
}

func (o *UserTypeDeleted) UserTypeObjectId() TlObjectIdType {
	return o.GetObjectTypeId()
}

func NewUserTypeDeleted() *UserTypeDeleted {
	return &UserTypeDeleted{}
}

func (*UserTypeDeleted) Serialize(wr *TlWriter) {
	// a = rd.ReadInt32()
	// b = rd.ReadString()
	// ...
}

func (*UserTypeDeleted) Desrialize(rd *TlReader) {
	// wr.Write(a)
	// wr.Write(b)
	// ...
}

type UserTypeBot struct {
	CanJoinGroups          bool
	CanReadAllGroupMssages bool
	IsInline               bool
	InlineQueryPlaceholder string
	NeedLocation           bool
}

func NewUserTypeBot(can_join_groups bool,
	can_read_all_group_messages bool,
	is_inline bool,
	inline_query_placeholder string,
	need_location bool) *UserTypeBot {
	return &UserTypeBot{
		CanJoinGroups:          can_join_groups,
		CanReadAllGroupMssages: can_read_all_group_messages,
		IsInline:               is_inline,
		InlineQueryPlaceholder: inline_query_placeholder,
		NeedLocation:           need_location,
	}
}

func (*UserTypeBot) Serialize(wr *TlWriter) {
	// a = rd.ReadInt32()
	// b = rd.ReadString()
	// CanJoinGroups = rd.ReadBool()
	// ...
}

func (*UserTypeBot) Desrialize(rd *TlReader) {
	// wr.Write(a)
	// wr.Write(b)
	// ...
}

type UserTypeUnknown struct {
}
