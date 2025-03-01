// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Represents the event action configuration for an element of a Component or
// ComponentChild . Use for the workflow feature in Amplify Studio that allows you
// to bind events and actions to components. ActionParameters defines the action
// that is performed when an event occurs on the component.
type ActionParameters struct {

	// The HTML anchor link to the location to open. Specify this value for a
	// navigation action.
	Anchor *ComponentProperty

	// A dictionary of key-value pairs mapping Amplify Studio properties to fields in
	// a data model. Use when the action performs an operation on an Amplify DataStore
	// model.
	Fields map[string]ComponentProperty

	// Specifies whether the user should be signed out globally. Specify this value
	// for an auth sign out action.
	Global *ComponentProperty

	// The unique ID of the component that the ActionParameters apply to.
	Id *ComponentProperty

	// The name of the data model. Use when the action performs an operation on an
	// Amplify DataStore model.
	Model *string

	// A key-value pair that specifies the state property name and its initial value.
	State *MutationActionSetStateParameter

	// The element within the same component to modify when the action occurs.
	Target *ComponentProperty

	// The type of navigation action. Valid values are url and anchor . This value is
	// required for a navigation action.
	Type *ComponentProperty

	// The URL to the location to open. Specify this value for a navigation action.
	Url *ComponentProperty

	noSmithyDocumentSerde
}

// Contains the configuration settings for a user interface (UI) element for an
// Amplify app. A component is configured as a primary, stand-alone UI element. Use
// ComponentChild to configure an instance of a Component . A ComponentChild
// instance inherits the configuration of the main Component .
type Component struct {

	// The unique ID of the Amplify app associated with the component.
	//
	// This member is required.
	AppId *string

	// The information to connect a component's properties to data at runtime. You
	// can't specify tags as a valid property for bindingProperties .
	//
	// This member is required.
	BindingProperties map[string]ComponentBindingPropertiesValue

	// The type of the component. This can be an Amplify custom UI component or
	// another custom component.
	//
	// This member is required.
	ComponentType *string

	// The time that the component was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The unique ID of the component.
	//
	// This member is required.
	Id *string

	// The name of the component.
	//
	// This member is required.
	Name *string

	// Describes the component's properties that can be overriden in a customized
	// instance of the component. You can't specify tags as a valid property for
	// overrides .
	//
	// This member is required.
	Overrides map[string]map[string]string

	// Describes the component's properties. You can't specify tags as a valid
	// property for properties .
	//
	// This member is required.
	Properties map[string]ComponentProperty

	// A list of the component's variants. A variant is a unique style configuration
	// of a main component.
	//
	// This member is required.
	Variants []ComponentVariant

	// A list of the component's ComponentChild instances.
	Children []ComponentChild

	// The data binding configuration for the component's properties. Use this for a
	// collection component. You can't specify tags as a valid property for
	// collectionProperties .
	CollectionProperties map[string]ComponentDataConfiguration

	// Describes the events that can be raised on the component. Use for the workflow
	// feature in Amplify Studio that allows you to bind events and actions to
	// components.
	Events map[string]ComponentEvent

	// The time that the component was modified.
	ModifiedAt *time.Time

	// The schema version of the component when it was imported.
	SchemaVersion *string

	// The unique ID of the component in its original source system, such as Figma.
	SourceId *string

	// One or more key-value pairs to use when tagging the component.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents the data binding configuration for a component at runtime. You can
// use ComponentBindingPropertiesValue to add exposed properties to a component to
// allow different values to be entered when a component is reused in different
// places in an app.
type ComponentBindingPropertiesValue struct {

	// Describes the properties to customize with data at runtime.
	BindingProperties *ComponentBindingPropertiesValueProperties

	// The default value of the property.
	DefaultValue *string

	// The property type.
	Type *string

	noSmithyDocumentSerde
}

// Represents the data binding configuration for a specific property using data
// stored in Amazon Web Services. For Amazon Web Services connected properties, you
// can bind a property to data stored in an Amazon S3 bucket, an Amplify DataStore
// model or an authenticated user attribute.
type ComponentBindingPropertiesValueProperties struct {

	// An Amazon S3 bucket.
	Bucket *string

	// The default value to assign to the property.
	DefaultValue *string

	// The field to bind the data to.
	Field *string

	// The storage key for an Amazon S3 bucket.
	Key *string

	// An Amplify DataStore model.
	Model *string

	// A list of predicates for binding a component's properties to data.
	Predicates []Predicate

	// The name of a component slot.
	SlotName *string

	// An authenticated user attribute.
	UserAttribute *string

	noSmithyDocumentSerde
}

// A nested UI configuration within a parent Component .
type ComponentChild struct {

	// The type of the child component.
	//
	// This member is required.
	ComponentType *string

	// The name of the child component.
	//
	// This member is required.
	Name *string

	// Describes the properties of the child component. You can't specify tags as a
	// valid property for properties .
	//
	// This member is required.
	Properties map[string]ComponentProperty

	// The list of ComponentChild instances for this component.
	Children []ComponentChild

	// Describes the events that can be raised on the child component. Use for the
	// workflow feature in Amplify Studio that allows you to bind events and actions to
	// components.
	Events map[string]ComponentEvent

	// The unique ID of the child component in its original source system, such as
	// Figma.
	SourceId *string

	noSmithyDocumentSerde
}

// Represents a conditional expression to set a component property. Use
// ComponentConditionProperty to set a property to different values conditionally,
// based on the value of another property.
type ComponentConditionProperty struct {

	// The value to assign to the property if the condition is not met.
	Else *ComponentProperty

	// The name of a field. Specify this when the property is a data model.
	Field *string

	// The value of the property to evaluate.
	Operand *string

	// The type of the property to evaluate.
	OperandType *string

	// The operator to use to perform the evaluation, such as eq to represent equals.
	Operator *string

	// The name of the conditional property.
	Property *string

	// The value to assign to the property if the condition is met.
	Then *ComponentProperty

	noSmithyDocumentSerde
}

// Describes the configuration for binding a component's properties to data.
type ComponentDataConfiguration struct {

	// The name of the data model to use to bind data to a component.
	//
	// This member is required.
	Model *string

	// A list of IDs to use to bind data to a component. Use this property to bind
	// specifically chosen data, rather than data retrieved from a query.
	Identifiers []string

	// Represents the conditional logic to use when binding data to a component. Use
	// this property to retrieve only a subset of the data in a collection.
	Predicate *Predicate

	// Describes how to sort the component's properties.
	Sort []SortProperty

	noSmithyDocumentSerde
}

// Describes the configuration of an event. You can bind an event and a
// corresponding action to a Component or a ComponentChild . A button click is an
// example of an event.
type ComponentEvent struct {

	// The action to perform when a specific event is raised.
	Action *string

	// Binds an event to an action on a component. When you specify a bindingEvent ,
	// the event is called when the action is performed.
	BindingEvent *string

	// Describes information about the action.
	Parameters *ActionParameters

	noSmithyDocumentSerde
}

// Describes the configuration for all of a component's properties. Use
// ComponentProperty to specify the values to render or bind by default.
type ComponentProperty struct {

	// The information to bind the component property to data at runtime.
	BindingProperties *ComponentPropertyBindingProperties

	// The information to bind the component property to form data.
	Bindings map[string]FormBindingElement

	// The information to bind the component property to data at runtime. Use this for
	// collection components.
	CollectionBindingProperties *ComponentPropertyBindingProperties

	// The name of the component that is affected by an event.
	ComponentName *string

	// A list of component properties to concatenate to create the value to assign to
	// this component property.
	Concat []ComponentProperty

	// The conditional expression to use to assign a value to the component property.
	Condition *ComponentConditionProperty

	// Specifies whether the user configured the property in Amplify Studio after
	// importing it.
	Configured *bool

	// The default value to assign to the component property.
	DefaultValue *string

	// An event that occurs in your app. Use this for workflow data binding.
	Event *string

	// The default value assigned to the property when the component is imported into
	// an app.
	ImportedValue *string

	// The data model to use to assign a value to the component property.
	Model *string

	// The name of the component's property that is affected by an event.
	Property *string

	// The component type.
	Type *string

	// An authenticated user attribute to use to assign a value to the component
	// property.
	UserAttribute *string

	// The value to assign to the component property.
	Value *string

	noSmithyDocumentSerde
}

// Associates a component property to a binding property. This enables exposed
// properties on the top level component to propagate data to the component's
// property values.
type ComponentPropertyBindingProperties struct {

	// The component property to bind to the data field.
	//
	// This member is required.
	Property *string

	// The data field to bind the property to.
	Field *string

	noSmithyDocumentSerde
}

// Contains a summary of a component. This is a read-only data type that is
// returned by ListComponents .
type ComponentSummary struct {

	// The unique ID of the Amplify app associated with the component.
	//
	// This member is required.
	AppId *string

	// The component type.
	//
	// This member is required.
	ComponentType *string

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The unique ID of the component.
	//
	// This member is required.
	Id *string

	// The name of the component.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Describes the style configuration of a unique variation of a main component.
type ComponentVariant struct {

	// The properties of the component variant that can be overriden when customizing
	// an instance of the component. You can't specify tags as a valid property for
	// overrides .
	Overrides map[string]map[string]string

	// The combination of variants that comprise this variant. You can't specify tags
	// as a valid property for variantValues .
	VariantValues map[string]string

	noSmithyDocumentSerde
}

// Represents all of the information that is required to create a component.
type CreateComponentData struct {

	// The data binding information for the component's properties.
	//
	// This member is required.
	BindingProperties map[string]ComponentBindingPropertiesValue

	// The component type. This can be an Amplify custom UI component or another
	// custom component.
	//
	// This member is required.
	ComponentType *string

	// The name of the component
	//
	// This member is required.
	Name *string

	// Describes the component properties that can be overriden to customize an
	// instance of the component.
	//
	// This member is required.
	Overrides map[string]map[string]string

	// Describes the component's properties.
	//
	// This member is required.
	Properties map[string]ComponentProperty

	// A list of the unique variants of this component.
	//
	// This member is required.
	Variants []ComponentVariant

	// A list of child components that are instances of the main component.
	Children []ComponentChild

	// The data binding configuration for customizing a component's properties. Use
	// this for a collection component.
	CollectionProperties map[string]ComponentDataConfiguration

	// The event configuration for the component. Use for the workflow feature in
	// Amplify Studio that allows you to bind events and actions to components.
	Events map[string]ComponentEvent

	// The schema version of the component when it was imported.
	SchemaVersion *string

	// The unique ID of the component in its original source system, such as Figma.
	SourceId *string

	// One or more key-value pairs to use when tagging the component data.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents all of the information that is required to create a form.
type CreateFormData struct {

	// The type of data source to use to create the form.
	//
	// This member is required.
	DataType *FormDataTypeConfig

	// The configuration information for the form's fields.
	//
	// This member is required.
	Fields map[string]FieldConfig

	// Specifies whether to perform a create or update action on the form.
	//
	// This member is required.
	FormActionType FormActionType

	// The name of the form.
	//
	// This member is required.
	Name *string

	// The schema version of the form.
	//
	// This member is required.
	SchemaVersion *string

	// The configuration information for the visual helper elements for the form.
	// These elements are not associated with any data.
	//
	// This member is required.
	SectionalElements map[string]SectionalElement

	// The configuration for the form's style.
	//
	// This member is required.
	Style *FormStyle

	// The FormCTA object that stores the call to action configuration for the form.
	Cta *FormCTA

	// Specifies an icon or decoration to display on the form.
	LabelDecorator LabelDecorator

	// One or more key-value pairs to use when tagging the form data.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Represents all of the information that is required to create a theme.
type CreateThemeData struct {

	// The name of the theme.
	//
	// This member is required.
	Name *string

	// A list of key-value pairs that deﬁnes the properties of the theme.
	//
	// This member is required.
	Values []ThemeValues

	// Describes the properties that can be overriden to customize an instance of the
	// theme.
	Overrides []ThemeValues

	// One or more key-value pairs to use when tagging the theme data.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Describes the configuration of a request to exchange an access code for a token.
type ExchangeCodeForTokenRequestBody struct {

	// The access code to send in the request.
	//
	// This member is required.
	Code *string

	// The location of the application that will receive the access code.
	//
	// This member is required.
	RedirectUri *string

	// The ID of the client to request the token from.
	ClientId *string

	noSmithyDocumentSerde
}

// Describes the configuration information for a field in a table.
type FieldConfig struct {

	// Specifies whether to hide a field.
	Excluded *bool

	// Describes the configuration for the default input value to display for a field.
	InputType *FieldInputConfig

	// The label for the field.
	Label *string

	// Specifies the field position.
	Position FieldPosition

	// The validations to perform on the value in the field.
	Validations []FieldValidationConfiguration

	noSmithyDocumentSerde
}

// Describes the configuration for the default input values to display for a field.
type FieldInputConfig struct {

	// The input type for the field.
	//
	// This member is required.
	Type *string

	// Specifies whether a field has a default value.
	DefaultChecked *bool

	// The default country code for a phone number.
	DefaultCountryCode *string

	// The default value for the field.
	DefaultValue *string

	// The text to display to describe the field.
	DescriptiveText *string

	// The configuration for the file uploader field.
	FileUploaderConfig *FileUploaderFieldConfig

	// Specifies whether to render the field as an array. This property is ignored if
	// the dataSourceType for the form is a Data Store.
	IsArray *bool

	// The maximum value to display for the field.
	MaxValue *float32

	// The minimum value to display for the field.
	MinValue *float32

	// The name of the field.
	Name *string

	// The text to display as a placeholder for the field.
	Placeholder *string

	// Specifies a read only field.
	ReadOnly *bool

	// Specifies a field that requires input.
	Required *bool

	// The stepping increment for a numeric value in a field.
	Step *float32

	// The value for the field.
	Value *string

	// The information to use to customize the input fields with data at runtime.
	ValueMappings *ValueMappings

	noSmithyDocumentSerde
}

// Describes the field position.
//
// The following types satisfy this interface:
//
//	FieldPositionMemberBelow
//	FieldPositionMemberFixed
//	FieldPositionMemberRightOf
type FieldPosition interface {
	isFieldPosition()
}

// The field position is below the field specified by the string.
type FieldPositionMemberBelow struct {
	Value string

	noSmithyDocumentSerde
}

func (*FieldPositionMemberBelow) isFieldPosition() {}

// The field position is fixed and doesn't change in relation to other fields.
type FieldPositionMemberFixed struct {
	Value FixedPosition

	noSmithyDocumentSerde
}

func (*FieldPositionMemberFixed) isFieldPosition() {}

// The field position is to the right of the field specified by the string.
type FieldPositionMemberRightOf struct {
	Value string

	noSmithyDocumentSerde
}

func (*FieldPositionMemberRightOf) isFieldPosition() {}

// Describes the validation configuration for a field.
type FieldValidationConfiguration struct {

	// The validation to perform on an object type.
	//
	// This member is required.
	Type *string

	// The validation to perform on a number value.
	NumValues []int32

	// The validation to perform on a string value.
	StrValues []string

	// The validation message to display.
	ValidationMessage *string

	noSmithyDocumentSerde
}

// Describes the configuration for the file uploader field.
type FileUploaderFieldConfig struct {

	// The file types that are allowed to be uploaded by the file uploader. Provide
	// this information in an array of strings specifying the valid file extensions.
	//
	// This member is required.
	AcceptedFileTypes []string

	// The access level to assign to the uploaded files in the Amazon S3 bucket where
	// they are stored. The valid values for this property are private , protected , or
	// public . For detailed information about the permissions associated with each
	// access level, see File access levels (https://docs.amplify.aws/lib/storage/configureaccess/q/platform/js/)
	// in the Amplify documentation.
	//
	// This member is required.
	AccessLevel StorageAccessLevel

	// Allows the file upload operation to be paused and resumed. The default value is
	// false . When isResumable is set to true , the file uploader uses a multipart
	// upload to break the files into chunks before upload. The progress of the upload
	// isn't continuous, because the file uploader uploads a chunk at a time.
	IsResumable *bool

	// Specifies the maximum number of files that can be selected to upload. The
	// default value is an unlimited number of files.
	MaxFileCount *int32

	// The maximum file size in bytes that the file uploader will accept. The default
	// value is an unlimited file size.
	MaxSize *int32

	// Specifies whether to display or hide the image preview after selecting a file
	// for upload. The default value is true to display the image preview.
	ShowThumbnails *bool

	noSmithyDocumentSerde
}

// Contains the configuration settings for a Form user interface (UI) element for
// an Amplify app. A form is a component you can add to your project by specifying
// a data source as the default configuration for the form.
type Form struct {

	// The unique ID of the Amplify app associated with the form.
	//
	// This member is required.
	AppId *string

	// The type of data source to use to create the form.
	//
	// This member is required.
	DataType *FormDataTypeConfig

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// Stores the information about the form's fields.
	//
	// This member is required.
	Fields map[string]FieldConfig

	// The operation to perform on the specified form.
	//
	// This member is required.
	FormActionType FormActionType

	// The unique ID of the form.
	//
	// This member is required.
	Id *string

	// The name of the form.
	//
	// This member is required.
	Name *string

	// The schema version of the form when it was imported.
	//
	// This member is required.
	SchemaVersion *string

	// Stores the visual helper elements for the form that are not associated with any
	// data.
	//
	// This member is required.
	SectionalElements map[string]SectionalElement

	// Stores the configuration for the form's style.
	//
	// This member is required.
	Style *FormStyle

	// Stores the call to action configuration for the form.
	Cta *FormCTA

	// Specifies an icon or decoration to display on the form.
	LabelDecorator LabelDecorator

	// One or more key-value pairs to use when tagging the form.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Describes how to bind a component property to form data.
type FormBindingElement struct {

	// The name of the component to retrieve a value from.
	//
	// This member is required.
	Element *string

	// The property to retrieve a value from.
	//
	// This member is required.
	Property *string

	noSmithyDocumentSerde
}

// Describes the configuration for a button UI element that is a part of a form.
type FormButton struct {

	// Describes the button's properties.
	Children *string

	// Specifies whether the button is visible on the form.
	Excluded *bool

	// The position of the button.
	Position FieldPosition

	noSmithyDocumentSerde
}

// Describes the call to action button configuration for the form.
type FormCTA struct {

	// Displays a cancel button.
	Cancel *FormButton

	// Displays a clear button.
	Clear *FormButton

	// The position of the button.
	Position FormButtonsPosition

	// Displays a submit button.
	Submit *FormButton

	noSmithyDocumentSerde
}

// Describes the data type configuration for the data source associated with a
// form.
type FormDataTypeConfig struct {

	// The data source type, either an Amplify DataStore model or a custom data type.
	//
	// This member is required.
	DataSourceType FormDataSourceType

	// The unique name of the data type you are using as the data source for the form.
	//
	// This member is required.
	DataTypeName *string

	noSmithyDocumentSerde
}

// Represents the data binding configuration for a form's input fields at
// runtime.You can use FormInputBindingPropertiesValue to add exposed properties
// to a form to allow different values to be entered when a form is reused in
// different places in an app.
type FormInputBindingPropertiesValue struct {

	// Describes the properties to customize with data at runtime.
	BindingProperties *FormInputBindingPropertiesValueProperties

	// The property type.
	Type *string

	noSmithyDocumentSerde
}

// Represents the data binding configuration for a specific property using data
// stored in Amazon Web Services. For Amazon Web Services connected properties, you
// can bind a property to data stored in an Amplify DataStore model.
type FormInputBindingPropertiesValueProperties struct {

	// An Amplify DataStore model.
	Model *string

	noSmithyDocumentSerde
}

// Describes the configuration for an input field on a form. Use
// FormInputValueProperty to specify the values to render or bind by default.
type FormInputValueProperty struct {

	// The information to bind fields to data at runtime.
	BindingProperties *FormInputValuePropertyBindingProperties

	// A list of form properties to concatenate to create the value to assign to this
	// field property.
	Concat []FormInputValueProperty

	// The value to assign to the input field.
	Value *string

	noSmithyDocumentSerde
}

// Associates a form property to a binding property. This enables exposed
// properties on the top level form to propagate data to the form's property
// values.
type FormInputValuePropertyBindingProperties struct {

	// The form property to bind to the data field.
	//
	// This member is required.
	Property *string

	// The data field to bind the property to.
	Field *string

	noSmithyDocumentSerde
}

// Describes the configuration for the form's style.
type FormStyle struct {

	// The spacing for the horizontal gap.
	HorizontalGap FormStyleConfig

	// The size of the outer padding for the form.
	OuterPadding FormStyleConfig

	// The spacing for the vertical gap.
	VerticalGap FormStyleConfig

	noSmithyDocumentSerde
}

// Describes the configuration settings for the form's style properties.
//
// The following types satisfy this interface:
//
//	FormStyleConfigMemberTokenReference
//	FormStyleConfigMemberValue
type FormStyleConfig interface {
	isFormStyleConfig()
}

// A reference to a design token to use to bind the form's style properties to an
// existing theme.
type FormStyleConfigMemberTokenReference struct {
	Value string

	noSmithyDocumentSerde
}

func (*FormStyleConfigMemberTokenReference) isFormStyleConfig() {}

// The value of the style setting.
type FormStyleConfigMemberValue struct {
	Value string

	noSmithyDocumentSerde
}

func (*FormStyleConfigMemberValue) isFormStyleConfig() {}

// Describes the basic information about a form.
type FormSummary struct {

	// The unique ID for the app associated with the form summary.
	//
	// This member is required.
	AppId *string

	// The form's data source type.
	//
	// This member is required.
	DataType *FormDataTypeConfig

	// The name of the backend environment that is part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The type of operation to perform on the form.
	//
	// This member is required.
	FormActionType FormActionType

	// The ID of the form.
	//
	// This member is required.
	Id *string

	// The name of the form.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Represents the state configuration when an action modifies a property of
// another element within the same component.
type MutationActionSetStateParameter struct {

	// The name of the component that is being modified.
	//
	// This member is required.
	ComponentName *string

	// The name of the component property to apply the state configuration to.
	//
	// This member is required.
	Property *string

	// The state configuration to assign to the property.
	//
	// This member is required.
	Set *ComponentProperty

	noSmithyDocumentSerde
}

// Stores information for generating Amplify DataStore queries. Use a Predicate to
// retrieve a subset of the data in a collection.
type Predicate struct {

	// A list of predicates to combine logically.
	And []Predicate

	// The field to query.
	Field *string

	// The value to use when performing the evaluation.
	Operand *string

	// The type of value to use when performing the evaluation.
	OperandType *string

	// The operator to use to perform the evaluation.
	Operator *string

	// A list of predicates to combine logically.
	Or []Predicate

	noSmithyDocumentSerde
}

// Stores the metadata information about a feature on a form.
type PutMetadataFlagBody struct {

	// The new information to store.
	//
	// This member is required.
	NewValue *string

	noSmithyDocumentSerde
}

// Describes a refresh token.
type RefreshTokenRequestBody struct {

	// The token to use to refresh a previously issued access token that might have
	// expired.
	//
	// This member is required.
	Token *string

	// The ID of the client to request the token from.
	ClientId *string

	noSmithyDocumentSerde
}

// Stores the configuration information for a visual helper element for a form. A
// sectional element can be a header, a text block, or a divider. These elements
// are static and not associated with any data.
type SectionalElement struct {

	// The type of sectional element. Valid values are Heading , Text , and Divider .
	//
	// This member is required.
	Type *string

	// Excludes a sectional element that was generated by default for a specified data
	// model.
	Excluded *bool

	// Specifies the size of the font for a Heading sectional element. Valid values
	// are 1 | 2 | 3 | 4 | 5 | 6 .
	Level *int32

	// Specifies the orientation for a Divider sectional element. Valid values are
	// horizontal or vertical .
	Orientation *string

	// Specifies the position of the text in a field for a Text sectional element.
	Position FieldPosition

	// The text for a Text sectional element.
	Text *string

	noSmithyDocumentSerde
}

// Describes how to sort the data that you bind to a component.
type SortProperty struct {

	// The direction of the sort, either ascending or descending.
	//
	// This member is required.
	Direction SortDirection

	// The field to perform the sort on.
	//
	// This member is required.
	Field *string

	noSmithyDocumentSerde
}

// A theme is a collection of style settings that apply globally to the components
// associated with an Amplify application.
type Theme struct {

	// The unique ID for the Amplify app associated with the theme.
	//
	// This member is required.
	AppId *string

	// The time that the theme was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The name of the backend environment that is a part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The ID for the theme.
	//
	// This member is required.
	Id *string

	// The name of the theme.
	//
	// This member is required.
	Name *string

	// A list of key-value pairs that defines the properties of the theme.
	//
	// This member is required.
	Values []ThemeValues

	// The time that the theme was modified.
	ModifiedAt *time.Time

	// Describes the properties that can be overriden to customize a theme.
	Overrides []ThemeValues

	// One or more key-value pairs to use when tagging the theme.
	Tags map[string]string

	noSmithyDocumentSerde
}

// Describes the basic information about a theme.
type ThemeSummary struct {

	// The unique ID for the app associated with the theme summary.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment that is part of the Amplify app.
	//
	// This member is required.
	EnvironmentName *string

	// The ID of the theme.
	//
	// This member is required.
	Id *string

	// The name of the theme.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Describes the configuration of a theme's properties.
type ThemeValue struct {

	// A list of key-value pairs that define the theme's properties.
	Children []ThemeValues

	// The value of a theme property.
	Value *string

	noSmithyDocumentSerde
}

// A key-value pair that defines a property of a theme.
type ThemeValues struct {

	// The name of the property.
	Key *string

	// The value of the property.
	Value *ThemeValue

	noSmithyDocumentSerde
}

// Updates and saves all of the information about a component, based on component
// ID.
type UpdateComponentData struct {

	// The data binding information for the component's properties.
	BindingProperties map[string]ComponentBindingPropertiesValue

	// The components that are instances of the main component.
	Children []ComponentChild

	// The configuration for binding a component's properties to a data model. Use
	// this for a collection component.
	CollectionProperties map[string]ComponentDataConfiguration

	// The type of the component. This can be an Amplify custom UI component or
	// another custom component.
	ComponentType *string

	// The event configuration for the component. Use for the workflow feature in
	// Amplify Studio that allows you to bind events and actions to components.
	Events map[string]ComponentEvent

	// The unique ID of the component to update.
	Id *string

	// The name of the component to update.
	Name *string

	// Describes the properties that can be overriden to customize the component.
	Overrides map[string]map[string]string

	// Describes the component's properties.
	Properties map[string]ComponentProperty

	// The schema version of the component when it was imported.
	SchemaVersion *string

	// The unique ID of the component in its original source system, such as Figma.
	SourceId *string

	// A list of the unique variants of the main component being updated.
	Variants []ComponentVariant

	noSmithyDocumentSerde
}

// Updates and saves all of the information about a form, based on form ID.
type UpdateFormData struct {

	// The FormCTA object that stores the call to action configuration for the form.
	Cta *FormCTA

	// The type of data source to use to create the form.
	DataType *FormDataTypeConfig

	// The configuration information for the form's fields.
	Fields map[string]FieldConfig

	// Specifies whether to perform a create or update action on the form.
	FormActionType FormActionType

	// Specifies an icon or decoration to display on the form.
	LabelDecorator LabelDecorator

	// The name of the form.
	Name *string

	// The schema version of the form.
	SchemaVersion *string

	// The configuration information for the visual helper elements for the form.
	// These elements are not associated with any data.
	SectionalElements map[string]SectionalElement

	// The configuration for the form's style.
	Style *FormStyle

	noSmithyDocumentSerde
}

// Saves the data binding information for a theme.
type UpdateThemeData struct {

	// A list of key-value pairs that define the theme's properties.
	//
	// This member is required.
	Values []ThemeValues

	// The unique ID of the theme to update.
	Id *string

	// The name of the theme to update.
	Name *string

	// Describes the properties that can be overriden to customize the theme.
	Overrides []ThemeValues

	noSmithyDocumentSerde
}

// Associates a complex object with a display value. Use ValueMapping to store how
// to represent complex objects when they are displayed.
type ValueMapping struct {

	// The complex object.
	//
	// This member is required.
	Value *FormInputValueProperty

	// The value to display for the complex object.
	DisplayValue *FormInputValueProperty

	noSmithyDocumentSerde
}

// Represents the data binding configuration for a value map.
type ValueMappings struct {

	// The value and display value pairs.
	//
	// This member is required.
	Values []ValueMapping

	// The information to bind fields to data at runtime.
	BindingProperties map[string]FormInputBindingPropertiesValue

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isFieldPosition()   {}
func (*UnknownUnionMember) isFormStyleConfig() {}
