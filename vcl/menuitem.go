
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TMenuItem struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewMenuItem(owner IComponent) *TMenuItem {
    m := new(TMenuItem)
    m.instance = MenuItem_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(m, (*TMenuItem).Free)
    return m
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsMenuItem(obj interface{}) *TMenuItem {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TMenuItem{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsMenuItem.
func MenuItemFromInst(inst uintptr) *TMenuItem {
    return AsMenuItem(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsMenuItem.
func MenuItemFromObj(obj IObject) *TMenuItem {
    return AsMenuItem(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsMenuItem.
func MenuItemFromUnsafePointer(ptr unsafe.Pointer) *TMenuItem {
    return AsMenuItem(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (m *TMenuItem) Free() {
    if m.instance != 0 {
        MenuItem_Free(m.instance)
        m.instance, m.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (m *TMenuItem) Instance() uintptr {
    return m.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (m *TMenuItem) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (m *TMenuItem) IsValid() bool {
    return m.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (m *TMenuItem) Is() TIs {
    return TIs(m.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (m *TMenuItem) As() TAs {
//    return TAs(m.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TMenuItemClass() TClass {
    return MenuItem_StaticClassType()
}

func (m *TMenuItem) Insert(Index int32, Item IComponent) {
    MenuItem_Insert(m.instance, Index , CheckPtr(Item))
}

func (m *TMenuItem) Delete(Index int32) {
    MenuItem_Delete(m.instance, Index)
}

// CN: 清除。
// EN: .
func (m *TMenuItem) Clear() {
    MenuItem_Clear(m.instance)
}

// CN: 单击。
// EN: .
func (m *TMenuItem) Click() {
    MenuItem_Click(m.instance)
}

func (m *TMenuItem) IndexOf(Item IComponent) int32 {
    return MenuItem_IndexOf(m.instance, CheckPtr(Item))
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (m *TMenuItem) HasParent() bool {
    return MenuItem_HasParent(m.instance)
}

func (m *TMenuItem) Add(Item IComponent) {
    MenuItem_Add(m.instance, CheckPtr(Item))
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (m *TMenuItem) FindComponent(AName string) *TComponent {
    return AsComponent(MenuItem_FindComponent(m.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (m *TMenuItem) GetNamePath() string {
    return MenuItem_GetNamePath(m.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (m *TMenuItem) Assign(Source IObject) {
    MenuItem_Assign(m.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (m *TMenuItem) ClassType() TClass {
    return MenuItem_ClassType(m.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (m *TMenuItem) ClassName() string {
    return MenuItem_ClassName(m.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (m *TMenuItem) InstanceSize() int32 {
    return MenuItem_InstanceSize(m.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (m *TMenuItem) InheritsFrom(AClass TClass) bool {
    return MenuItem_InheritsFrom(m.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (m *TMenuItem) Equals(Obj IObject) bool {
    return MenuItem_Equals(m.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (m *TMenuItem) GetHashCode() int32 {
    return MenuItem_GetHashCode(m.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (m *TMenuItem) ToString() string {
    return MenuItem_ToString(m.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (m *TMenuItem) Handle() HMENU {
    return MenuItem_GetHandle(m.instance)
}

func (m *TMenuItem) Count() int32 {
    return MenuItem_GetCount(m.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (m *TMenuItem) Parent() *TMenuItem {
    return AsMenuItem(MenuItem_GetParent(m.instance))
}

func (m *TMenuItem) Action() *TAction {
    return AsAction(MenuItem_GetAction(m.instance))
}

func (m *TMenuItem) SetAction(value IComponent) {
    MenuItem_SetAction(m.instance, CheckPtr(value))
}

func (m *TMenuItem) AutoCheck() bool {
    return MenuItem_GetAutoCheck(m.instance)
}

func (m *TMenuItem) SetAutoCheck(value bool) {
    MenuItem_SetAutoCheck(m.instance, value)
}

func (m *TMenuItem) Bitmap() *TBitmap {
    return AsBitmap(MenuItem_GetBitmap(m.instance))
}

func (m *TMenuItem) SetBitmap(value *TBitmap) {
    MenuItem_SetBitmap(m.instance, CheckPtr(value))
}

// CN: 获取控件标题。
// EN: Get the control title.
func (m *TMenuItem) Caption() string {
    return MenuItem_GetCaption(m.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (m *TMenuItem) SetCaption(value string) {
    MenuItem_SetCaption(m.instance, value)
}

// CN: 获取是否选中。
// EN: .
func (m *TMenuItem) Checked() bool {
    return MenuItem_GetChecked(m.instance)
}

// CN: 设置是否选中。
// EN: .
func (m *TMenuItem) SetChecked(value bool) {
    MenuItem_SetChecked(m.instance, value)
}

func (m *TMenuItem) Default() bool {
    return MenuItem_GetDefault(m.instance)
}

func (m *TMenuItem) SetDefault(value bool) {
    MenuItem_SetDefault(m.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (m *TMenuItem) Enabled() bool {
    return MenuItem_GetEnabled(m.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (m *TMenuItem) SetEnabled(value bool) {
    MenuItem_SetEnabled(m.instance, value)
}

// CN: 获取团组索引。
// EN: .
func (m *TMenuItem) GroupIndex() uint8 {
    return MenuItem_GetGroupIndex(m.instance)
}

// CN: 设置团组索引。
// EN: .
func (m *TMenuItem) SetGroupIndex(value uint8) {
    MenuItem_SetGroupIndex(m.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (m *TMenuItem) Hint() string {
    return MenuItem_GetHint(m.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (m *TMenuItem) SetHint(value string) {
    MenuItem_SetHint(m.instance, value)
}

// CN: 获取图像在images中的索引。
// EN: .
func (m *TMenuItem) ImageIndex() int32 {
    return MenuItem_GetImageIndex(m.instance)
}

// CN: 设置图像在images中的索引。
// EN: .
func (m *TMenuItem) SetImageIndex(value int32) {
    MenuItem_SetImageIndex(m.instance, value)
}

func (m *TMenuItem) RadioItem() bool {
    return MenuItem_GetRadioItem(m.instance)
}

func (m *TMenuItem) SetRadioItem(value bool) {
    MenuItem_SetRadioItem(m.instance, value)
}

// CN: 获取快捷键。
// EN: .
func (m *TMenuItem) ShortCut() TShortCut {
    return MenuItem_GetShortCut(m.instance)
}

// CN: 设置快捷键。
// EN: .
func (m *TMenuItem) SetShortCut(value TShortCut) {
    MenuItem_SetShortCut(m.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (m *TMenuItem) Visible() bool {
    return MenuItem_GetVisible(m.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (m *TMenuItem) SetVisible(value bool) {
    MenuItem_SetVisible(m.instance, value)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (m *TMenuItem) SetOnClick(fn TNotifyEvent) {
    MenuItem_SetOnClick(m.instance, fn)
}

func (m *TMenuItem) SetOnMeasureItem(fn TMenuMeasureItemEvent) {
    MenuItem_SetOnMeasureItem(m.instance, fn)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (m *TMenuItem) ComponentCount() int32 {
    return MenuItem_GetComponentCount(m.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (m *TMenuItem) ComponentIndex() int32 {
    return MenuItem_GetComponentIndex(m.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (m *TMenuItem) SetComponentIndex(value int32) {
    MenuItem_SetComponentIndex(m.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (m *TMenuItem) Owner() *TComponent {
    return AsComponent(MenuItem_GetOwner(m.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (m *TMenuItem) Name() string {
    return MenuItem_GetName(m.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (m *TMenuItem) SetName(value string) {
    MenuItem_SetName(m.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (m *TMenuItem) Tag() int {
    return MenuItem_GetTag(m.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (m *TMenuItem) SetTag(value int) {
    MenuItem_SetTag(m.instance, value)
}

func (m *TMenuItem) Items(Index int32) *TMenuItem {
    return AsMenuItem(MenuItem_GetItems(m.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (m *TMenuItem) Components(AIndex int32) *TComponent {
    return AsComponent(MenuItem_GetComponents(m.instance, AIndex))
}

