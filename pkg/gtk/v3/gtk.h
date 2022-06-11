#include <stdlib.h>
#include <glib.h>

struct AboutDialogClass {
          parent_class;
    void* activate_link;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct AboutDialogPrivate {

};
struct AccelGroupClass {
          parent_class;
    void* accel_changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct AccelGroupEntry {
            key;
    void*   closure;
    guint32 accel_path_quark;
};
struct AccelGroupPrivate {

};
struct AccelKey {
    guint accel_key;
          accel_mods;
    guint accel_flags  : 16;
};
struct AccelLabelClass {
          parent_class;
    void* signal_quote1;
    void* signal_quote2;
    void* mod_name_shift;
    void* mod_name_control;
    void* mod_name_alt;
    void* mod_separator;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct AccelLabelPrivate {

};
struct AccelMapClass {

};
struct AccessibleClass {
          parent_class;
    void* connect_widget_destroyed;
    void* widget_set;
    void* widget_unset;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct AccessiblePrivate {

};
struct ActionBarClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ActionBarPrivate {

};
struct ActionClass {
          parent_class;
    void* activate;
          menu_item_type;
          toolbar_item_type;
    void* create_menu_item;
    void* create_tool_item;
    void* connect_proxy;
    void* disconnect_proxy;
    void* create_menu;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ActionEntry {
    void* name;
    void* stock_id;
    void* label;
    void* accelerator;
    void* tooltip;
          callback;
};
struct ActionGroupClass {
          parent_class;
    void* get_action;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ActionGroupPrivate {

};
struct ActionPrivate {

};
struct ActionableInterface {
          g_iface;
    void* get_action_name;
    void* set_action_name;
    void* get_action_target_value;
    void* set_action_target_value;
};
struct ActivatableIface {
          g_iface;
    void* update;
    void* sync_action_properties;
};
struct AdjustmentClass {
          parent_class;
    void* changed;
    void* value_changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct AdjustmentPrivate {

};
struct AlignmentClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct AlignmentPrivate {

};
struct AppChooserButtonClass {
          parent_class;
    void* custom_item_activated;
    void  padding;
};
struct AppChooserButtonPrivate {

};
struct AppChooserDialogClass {
         parent_class;
    void padding;
};
struct AppChooserDialogPrivate {

};
struct AppChooserWidgetClass {
          parent_class;
    void* application_selected;
    void* application_activated;
    void* populate_popup;
    void  padding;
};
struct AppChooserWidgetPrivate {

};
struct ApplicationClass {
          parent_class;
    void* window_added;
    void* window_removed;
    void  padding;
};
struct ApplicationPrivate {

};
struct ApplicationWindowClass {
         parent_class;
    void padding;
};
struct ApplicationWindowPrivate {

};
struct ArrowAccessibleClass {
     parent_class;
};
struct ArrowAccessiblePrivate {

};
struct ArrowClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ArrowPrivate {

};
struct AspectFrameClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct AspectFramePrivate {

};
struct AssistantClass {
          parent_class;
    void* prepare;
    void* apply;
    void* close;
    void* cancel;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
};
struct AssistantPrivate {

};
struct BinClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct BinPrivate {

};
struct BindingArg {
     arg_type;
};
struct BindingEntry {
    guint keyval;
          modifiers;
    void* binding_set;
    guint destroyed      : 1;
    guint in_emission    : 1;
    guint marks_unbound  : 1;
    void* set_next;
    void* hash_next;
    void* signals;
};
struct BindingSet {
    void* set_name;
    gint  priority;
    void* widget_path_pspecs;
    void* widget_class_pspecs;
    void* class_branch_pspecs;
    void* entries;
    void* current;
    guint parsed  : 1;
};
struct BindingSignal {
    void* next;
    void* signal_name;
    guint n_args;
    void* args;
};
struct BooleanCellAccessibleClass {
     parent_class;
};
struct BooleanCellAccessiblePrivate {

};
struct Border {
    gint16 left;
    gint16 right;
    gint16 top;
    gint16 bottom;
};
struct BoxClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct BoxPrivate {

};
struct BuildableIface {
          g_iface;
    void* set_name;
    void* get_name;
    void* add_child;
    void* set_buildable_property;
    void* construct_child;
    void* custom_tag_start;
    void* custom_tag_end;
    void* custom_finished;
    void* parser_finished;
    void* get_internal_child;
};
struct BuilderClass {
          parent_class;
    void* get_type_from_name;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct BuilderPrivate {

};
struct ButtonAccessibleClass {
     parent_class;
};
struct ButtonAccessiblePrivate {

};
struct ButtonBoxClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ButtonBoxPrivate {

};
struct ButtonClass {
          parent_class;
    void* pressed;
    void* released;
    void* clicked;
    void* enter;
    void* leave;
    void* activate;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ButtonPrivate {

};
struct CalendarClass {
          parent_class;
    void* month_changed;
    void* day_selected;
    void* day_selected_double_click;
    void* prev_month;
    void* next_month;
    void* prev_year;
    void* next_year;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CalendarPrivate {

};
struct CellAccessibleClass {
          parent_class;
    void* update_cache;
};
struct CellAccessibleParentIface {
          parent;
    void* get_cell_extents;
    void* get_cell_area;
    void* grab_focus;
    void* get_child_index;
    void* get_renderer_state;
    void* expand_collapse;
    void* activate;
    void* edit;
    void* update_relationset;
    void* get_cell_position;
    void* get_column_header_cells;
    void* get_row_header_cells;
};
struct CellAccessiblePrivate {

};
struct CellAreaBoxClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellAreaBoxPrivate {

};
struct CellAreaClass {
          parent_class;
    void* add;
    void* remove;
    void* foreach;
    void* foreach_alloc;
    void* event;
    void* render;
    void* apply_attributes;
    void* create_context;
    void* copy_context;
    void* get_request_mode;
    void* get_preferred_width;
    void* get_preferred_height_for_width;
    void* get_preferred_height;
    void* get_preferred_width_for_height;
    void* set_cell_property;
    void* get_cell_property;
    void* focus;
    void* is_activatable;
    void* activate;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct CellAreaContextClass {
          parent_class;
    void* allocate;
    void* reset;
    void* get_preferred_height_for_width;
    void* get_preferred_width_for_height;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
};
struct CellAreaContextPrivate {

};
struct CellAreaPrivate {

};
struct CellEditableIface {
          g_iface;
    void* editing_done;
    void* remove_widget;
    void* start_editing;
};
struct CellLayoutIface {
          g_iface;
    void* pack_start;
    void* pack_end;
    void* clear;
    void* add_attribute;
    void* set_cell_data_func;
    void* clear_attributes;
    void* reorder;
    void* get_cells;
    void* get_area;
};
struct CellRendererAccelClass {
          parent_class;
    void* accel_edited;
    void* accel_cleared;
    void* _gtk_reserved0;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellRendererAccelPrivate {

};
struct CellRendererClass {
          parent_class;
    void* get_request_mode;
    void* get_preferred_width;
    void* get_preferred_height_for_width;
    void* get_preferred_height;
    void* get_preferred_width_for_height;
    void* get_aligned_area;
    void* get_size;
    void* render;
    void* activate;
    void* start_editing;
    void* editing_canceled;
    void* editing_started;
    void* priv;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellRendererClassPrivate {

};
struct CellRendererComboClass {
          parent;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellRendererComboPrivate {

};
struct CellRendererPixbufClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellRendererPixbufPrivate {

};
struct CellRendererPrivate {

};
struct CellRendererProgressClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellRendererProgressPrivate {

};
struct CellRendererSpinClass {
          parent;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellRendererSpinPrivate {

};
struct CellRendererSpinnerClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellRendererSpinnerPrivate {

};
struct CellRendererTextClass {
          parent_class;
    void* edited;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellRendererTextPrivate {

};
struct CellRendererToggleClass {
          parent_class;
    void* toggled;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellRendererTogglePrivate {

};
struct CellViewClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CellViewPrivate {

};
struct CheckButtonClass {
          parent_class;
    void* draw_indicator;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CheckMenuItemAccessibleClass {
     parent_class;
};
struct CheckMenuItemAccessiblePrivate {

};
struct CheckMenuItemClass {
          parent_class;
    void* toggled;
    void* draw_indicator;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CheckMenuItemPrivate {

};
struct ColorButtonClass {
          parent_class;
    void* color_set;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ColorButtonPrivate {

};
struct ColorChooserDialogClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ColorChooserDialogPrivate {

};
struct ColorChooserInterface {
          base_interface;
    void* get_rgba;
    void* set_rgba;
    void* add_palette;
    void* color_activated;
    void  padding;
};
struct ColorChooserWidgetClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct ColorChooserWidgetPrivate {

};
struct ColorSelectionClass {
          parent_class;
    void* color_changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ColorSelectionDialogClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ColorSelectionDialogPrivate {

};
struct ColorSelectionPrivate {

};
struct ComboBoxAccessibleClass {
     parent_class;
};
struct ComboBoxAccessiblePrivate {

};
struct ComboBoxClass {
          parent_class;
    void* changed;
    void* format_entry_text;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
};
struct ComboBoxPrivate {

};
struct ComboBoxTextClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ComboBoxTextPrivate {

};
struct ContainerAccessibleClass {
          parent_class;
    void* add_gtk;
    void* remove_gtk;
};
struct ContainerAccessiblePrivate {

};
struct ContainerCellAccessibleClass {
     parent_class;
};
struct ContainerCellAccessiblePrivate {

};
struct ContainerClass {
          parent_class;
    void* add;
    void* remove;
    void* check_resize;
    void* forall;
    void* set_focus_child;
    void* child_type;
    void* composite_name;
    void* set_child_property;
    void* get_child_property;
    void* get_path_for_child;
          _handle_border_width  : 1;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct ContainerPrivate {

};
struct CssProviderClass {
          parent_class;
    void* parsing_error;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct CssProviderPrivate {

};
struct CssSection {

};
struct DialogClass {
          parent_class;
    void* response;
    void* close;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct DialogPrivate {

};
struct DrawingAreaClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct EditableInterface {
          base_iface;
    void* insert_text;
    void* delete_text;
    void* changed;
    void* do_insert_text;
    void* do_delete_text;
    void* get_chars;
    void* set_selection_bounds;
    void* get_selection_bounds;
    void* set_position;
    void* get_position;
};
struct EntryAccessibleClass {
     parent_class;
};
struct EntryAccessiblePrivate {

};
struct EntryBufferClass {
          parent_class;
    void* inserted_text;
    void* deleted_text;
    void* get_text;
    void* get_length;
    void* insert_text;
    void* delete_text;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct EntryBufferPrivate {

};
struct EntryClass {
          parent_class;
    void* populate_popup;
    void* activate;
    void* move_cursor;
    void* insert_at_cursor;
    void* delete_from_cursor;
    void* backspace;
    void* cut_clipboard;
    void* copy_clipboard;
    void* paste_clipboard;
    void* toggle_overwrite;
    void* get_text_area_size;
    void* get_frame_size;
    void* insert_emoji;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
};
struct EntryCompletionClass {
          parent_class;
    void* match_selected;
    void* action_activated;
    void* insert_prefix;
    void* cursor_on_match;
    void* no_matches;
    void* _gtk_reserved0;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
};
struct EntryCompletionPrivate {

};
struct EntryPrivate {

};
struct EventBoxClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct EventBoxPrivate {

};
struct EventControllerClass {

};
struct EventControllerKeyClass {

};
struct EventControllerMotionClass {

};
struct EventControllerScrollClass {

};
struct ExpanderAccessibleClass {
     parent_class;
};
struct ExpanderAccessiblePrivate {

};
struct ExpanderClass {
          parent_class;
    void* activate;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ExpanderPrivate {

};
struct FileChooserButtonClass {
          parent_class;
    void* file_set;
    void* __gtk_reserved1;
    void* __gtk_reserved2;
    void* __gtk_reserved3;
    void* __gtk_reserved4;
};
struct FileChooserButtonPrivate {

};
struct FileChooserDialogClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct FileChooserDialogPrivate {

};
struct FileChooserNativeClass {
     parent_class;
};
struct FileChooserWidgetClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct FileChooserWidgetPrivate {

};
struct FileFilterInfo {
          contains;
    void* filename;
    void* uri;
    void* display_name;
    void* mime_type;
};
struct FixedChild {
    void* widget;
    gint  x;
    gint  y;
};
struct FixedClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct FixedPrivate {

};
struct FlowBoxAccessibleClass {
     parent_class;
};
struct FlowBoxAccessiblePrivate {

};
struct FlowBoxChildAccessibleClass {
     parent_class;
};
struct FlowBoxChildClass {
          parent_class;
    void* activate;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
};
struct FlowBoxClass {
          parent_class;
    void* child_activated;
    void* selected_children_changed;
    void* activate_cursor_child;
    void* toggle_cursor_child;
    void* move_cursor;
    void* select_all;
    void* unselect_all;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
};
struct FontButtonClass {
          parent_class;
    void* font_set;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct FontButtonPrivate {

};
struct FontChooserDialogClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct FontChooserDialogPrivate {

};
struct FontChooserIface {
          base_iface;
    void* get_font_family;
    void* get_font_face;
    void* get_font_size;
    void* set_filter_func;
    void* font_activated;
    void* set_font_map;
    void* get_font_map;
    void  padding;
};
struct FontChooserWidgetClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct FontChooserWidgetPrivate {

};
struct FontSelectionClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct FontSelectionDialogClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct FontSelectionDialogPrivate {

};
struct FontSelectionPrivate {

};
struct FrameAccessibleClass {
     parent_class;
};
struct FrameAccessiblePrivate {

};
struct FrameClass {
          parent_class;
    void* compute_child_allocation;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct FramePrivate {

};
struct GLAreaClass {
          parent_class;
    void* render;
    void* resize;
    void* create_context;
    void  _padding;
};
struct GestureClass {

};
struct GestureDragClass {

};
struct GestureLongPressClass {

};
struct GestureMultiPressClass {

};
struct GesturePanClass {

};
struct GestureRotateClass {

};
struct GestureSingleClass {

};
struct GestureStylusClass {

};
struct GestureSwipeClass {

};
struct GestureZoomClass {

};
struct Gradient {

};
struct GridClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct GridPrivate {

};
struct HBoxClass {
     parent_class;
};
struct HButtonBoxClass {
     parent_class;
};
struct HPanedClass {
     parent_class;
};
struct HSVClass {
          parent_class;
    void* changed;
    void* move;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct HSVPrivate {

};
struct HScaleClass {
     parent_class;
};
struct HScrollbarClass {
     parent_class;
};
struct HSeparatorClass {
     parent_class;
};
struct HandleBoxClass {
          parent_class;
    void* child_attached;
    void* child_detached;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct HandleBoxPrivate {

};
struct HeaderBarAccessibleClass {
     parent_class;
};
struct HeaderBarAccessiblePrivate {

};
struct HeaderBarClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct HeaderBarPrivate {

};
struct IMContextClass {
          parent_class;
    void* preedit_start;
    void* preedit_end;
    void* preedit_changed;
    void* commit;
    void* retrieve_surrounding;
    void* delete_surrounding;
    void* set_client_window;
    void* get_preedit_string;
    void* filter_keypress;
    void* focus_in;
    void* focus_out;
    void* reset;
    void* set_cursor_location;
    void* set_use_preedit;
    void* set_surrounding;
    void* get_surrounding;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
};
struct IMContextInfo {
    void* context_id;
    void* context_name;
    void* domain;
    void* domain_dirname;
    void* default_locales;
};
struct IMContextSimpleClass {
     parent_class;
};
struct IMContextSimplePrivate {

};
struct IMMulticontextClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct IMMulticontextPrivate {

};
struct IconFactoryClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct IconFactoryPrivate {

};
struct IconInfoClass {

};
struct IconSet {

};
struct IconSource {

};
struct IconThemeClass {
          parent_class;
    void* changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct IconThemePrivate {

};
struct IconViewAccessibleClass {
     parent_class;
};
struct IconViewAccessiblePrivate {

};
struct IconViewClass {
          parent_class;
    void* item_activated;
    void* selection_changed;
    void* select_all;
    void* unselect_all;
    void* select_cursor_item;
    void* toggle_cursor_item;
    void* move_cursor;
    void* activate_cursor_item;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct IconViewPrivate {

};
struct ImageAccessibleClass {
     parent_class;
};
struct ImageAccessiblePrivate {

};
struct ImageCellAccessibleClass {
     parent_class;
};
struct ImageCellAccessiblePrivate {

};
struct ImageClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ImageMenuItemClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ImageMenuItemPrivate {

};
struct ImagePrivate {

};
struct InfoBarClass {
          parent_class;
    void* response;
    void* close;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct InfoBarPrivate {

};
struct InvisibleClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct InvisiblePrivate {

};
struct LabelAccessibleClass {
     parent_class;
};
struct LabelAccessiblePrivate {

};
struct LabelClass {
          parent_class;
    void* move_cursor;
    void* copy_clipboard;
    void* populate_popup;
    void* activate_link;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct LabelPrivate {

};
struct LabelSelectionInfo {

};
struct LayoutClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct LayoutPrivate {

};
struct LevelBarAccessibleClass {
     parent_class;
};
struct LevelBarAccessiblePrivate {

};
struct LevelBarClass {
          parent_class;
    void* offset_changed;
    void  padding;
};
struct LevelBarPrivate {

};
struct LinkButtonAccessibleClass {
     parent_class;
};
struct LinkButtonAccessiblePrivate {

};
struct LinkButtonClass {
          parent_class;
    void* activate_link;
    void* _gtk_padding1;
    void* _gtk_padding2;
    void* _gtk_padding3;
    void* _gtk_padding4;
};
struct LinkButtonPrivate {

};
struct ListBoxAccessibleClass {
     parent_class;
};
struct ListBoxAccessiblePrivate {

};
struct ListBoxClass {
          parent_class;
    void* row_selected;
    void* row_activated;
    void* activate_cursor_row;
    void* toggle_cursor_row;
    void* move_cursor;
    void* selected_rows_changed;
    void* select_all;
    void* unselect_all;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
};
struct ListBoxRowAccessibleClass {
     parent_class;
};
struct ListBoxRowClass {
          parent_class;
    void* activate;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
};
struct ListStoreClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ListStorePrivate {

};
struct LockButtonAccessibleClass {
     parent_class;
};
struct LockButtonAccessiblePrivate {

};
struct LockButtonClass {
          parent_class;
    void* reserved0;
    void* reserved1;
    void* reserved2;
    void* reserved3;
    void* reserved4;
    void* reserved5;
    void* reserved6;
    void* reserved7;
};
struct LockButtonPrivate {

};
struct MenuAccessibleClass {
     parent_class;
};
struct MenuAccessiblePrivate {

};
struct MenuBarClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct MenuBarPrivate {

};
struct MenuButtonAccessibleClass {
     parent_class;
};
struct MenuButtonAccessiblePrivate {

};
struct MenuButtonClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct MenuButtonPrivate {

};
struct MenuClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct MenuItemAccessibleClass {
     parent_class;
};
struct MenuItemAccessiblePrivate {

};
struct MenuItemClass {
          parent_class;
    guint hide_on_activate  : 1;
    void* activate;
    void* activate_item;
    void* toggle_size_request;
    void* toggle_size_allocate;
    void* set_label;
    void* get_label;
    void* select;
    void* deselect;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct MenuItemPrivate {

};
struct MenuPrivate {

};
struct MenuShellAccessibleClass {
     parent_class;
};
struct MenuShellAccessiblePrivate {

};
struct MenuShellClass {
          parent_class;
    guint submenu_placement  : 1;
    void* deactivate;
    void* selection_done;
    void* move_current;
    void* activate_current;
    void* cancel;
    void* select_item;
    void* insert;
    void* get_popup_delay;
    void* move_selected;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct MenuShellPrivate {

};
struct MenuToolButtonClass {
          parent_class;
    void* show_menu;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct MenuToolButtonPrivate {

};
struct MessageDialogClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct MessageDialogPrivate {

};
struct MiscClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct MiscPrivate {

};
struct MountOperationClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct MountOperationPrivate {

};
struct NativeDialogClass {
          parent_class;
    void* response;
    void* show;
    void* hide;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct NotebookAccessibleClass {
     parent_class;
};
struct NotebookAccessiblePrivate {

};
struct NotebookClass {
          parent_class;
    void* switch_page;
    void* select_page;
    void* focus_tab;
    void* change_current_page;
    void* move_focus_out;
    void* reorder_tab;
    void* insert_page;
    void* create_window;
    void* page_reordered;
    void* page_removed;
    void* page_added;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct NotebookPageAccessibleClass {
     parent_class;
};
struct NotebookPageAccessiblePrivate {

};
struct NotebookPrivate {

};
struct NumerableIconClass {
         parent_class;
    void padding;
};
struct NumerableIconPrivate {

};
struct OffscreenWindowClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct OrientableIface {
     base_iface;
};
struct OverlayClass {
          parent_class;
    void* get_child_position;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct OverlayPrivate {

};
struct PadActionEntry {
          type;
    gint  index;
    gint  mode;
    void* label;
    void* action_name;
};
struct PadControllerClass {

};
struct PageRange {
    gint start;
    gint end;
};
struct PanedAccessibleClass {
     parent_class;
};
struct PanedAccessiblePrivate {

};
struct PanedClass {
          parent_class;
    void* cycle_child_focus;
    void* toggle_handle_focus;
    void* move_handle;
    void* cycle_handle_focus;
    void* accept_position;
    void* cancel_position;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct PanedPrivate {

};
struct PaperSize {

};
struct PlacesSidebarClass {

};
struct PlugAccessibleClass {
     parent_class;
};
struct PlugAccessiblePrivate {

};
struct PlugClass {
          parent_class;
    void* embedded;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct PlugPrivate {

};
struct PopoverAccessibleClass {
     parent_class;
};
struct PopoverClass {
          parent_class;
    void* closed;
    void  reserved;
};
struct PopoverMenuClass {
         parent_class;
    void reserved;
};
struct PopoverPrivate {

};
struct PrintOperationClass {
          parent_class;
    void* done;
    void* begin_print;
    void* paginate;
    void* request_page_setup;
    void* draw_page;
    void* end_print;
    void* status_changed;
    void* create_custom_widget;
    void* custom_widget_apply;
    void* preview;
    void* update_custom_widget;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct PrintOperationPreviewIface {
          g_iface;
    void* ready;
    void* got_page_size;
    void* render_page;
    void* is_selected;
    void* end_preview;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct PrintOperationPrivate {

};
struct ProgressBarAccessibleClass {
     parent_class;
};
struct ProgressBarAccessiblePrivate {

};
struct ProgressBarClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ProgressBarPrivate {

};
struct RadioActionClass {
          parent_class;
    void* changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct RadioActionEntry {
    void* name;
    void* stock_id;
    void* label;
    void* accelerator;
    void* tooltip;
    gint  value;
};
struct RadioActionPrivate {

};
struct RadioButtonAccessibleClass {
     parent_class;
};
struct RadioButtonAccessiblePrivate {

};
struct RadioButtonClass {
          parent_class;
    void* group_changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct RadioButtonPrivate {

};
struct RadioMenuItemAccessibleClass {
     parent_class;
};
struct RadioMenuItemAccessiblePrivate {

};
struct RadioMenuItemClass {
          parent_class;
    void* group_changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct RadioMenuItemPrivate {

};
struct RadioToolButtonClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct RangeAccessibleClass {
     parent_class;
};
struct RangeAccessiblePrivate {

};
struct RangeClass {
          parent_class;
    void* slider_detail;
    void* stepper_detail;
    void* value_changed;
    void* adjust_bounds;
    void* move_slider;
    void* get_range_border;
    void* change_value;
    void* get_range_size_request;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
};
struct RangePrivate {

};
struct RcContext {

};
struct RcProperty {
    guint32 type_name;
    guint32 property_name;
    void*   origin;
            value;
};
struct RcStyleClass {
          parent_class;
    void* create_rc_style;
    void* parse;
    void* merge;
    void* create_style;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct RecentActionClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct RecentActionPrivate {

};
struct RecentChooserDialogClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct RecentChooserDialogPrivate {

};
struct RecentChooserIface {
          base_iface;
    void* set_current_uri;
    void* get_current_uri;
    void* select_uri;
    void* unselect_uri;
    void* select_all;
    void* unselect_all;
    void* get_items;
    void* get_recent_manager;
    void* add_filter;
    void* remove_filter;
    void* list_filters;
    void* set_sort_func;
    void* item_activated;
    void* selection_changed;
};
struct RecentChooserMenuClass {
          parent_class;
    void* gtk_recent1;
    void* gtk_recent2;
    void* gtk_recent3;
    void* gtk_recent4;
};
struct RecentChooserMenuPrivate {

};
struct RecentChooserWidgetClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct RecentChooserWidgetPrivate {

};
struct RecentData {
    void*    display_name;
    void*    description;
    void*    mime_type;
    void*    app_name;
    void*    app_exec;
    void**   groups;
    gboolean is_private;
};
struct RecentFilterInfo {
           contains;
    void*  uri;
    void*  display_name;
    void*  mime_type;
    void** applications;
    void** groups;
    gint   age;
};
struct RecentInfo {

};
struct RecentManagerClass {
          parent_class;
    void* changed;
    void* _gtk_recent1;
    void* _gtk_recent2;
    void* _gtk_recent3;
    void* _gtk_recent4;
};
struct RecentManagerPrivate {

};
struct RendererCellAccessibleClass {
     parent_class;
};
struct RendererCellAccessiblePrivate {

};
struct RequestedSize {
    gpointer data;
    gint     minimum_size;
    gint     natural_size;
};
struct Requisition {
    gint width;
    gint height;
};
struct RevealerClass {
     parent_class;
};
struct ScaleAccessibleClass {
     parent_class;
};
struct ScaleAccessiblePrivate {

};
struct ScaleButtonAccessibleClass {
     parent_class;
};
struct ScaleButtonAccessiblePrivate {

};
struct ScaleButtonClass {
          parent_class;
    void* value_changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ScaleButtonPrivate {

};
struct ScaleClass {
          parent_class;
    void* format_value;
    void* draw_value;
    void* get_layout_offsets;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ScalePrivate {

};
struct ScrollableInterface {
          base_iface;
    void* get_border;
};
struct ScrollbarClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ScrolledWindowAccessibleClass {
     parent_class;
};
struct ScrolledWindowAccessiblePrivate {

};
struct ScrolledWindowClass {
          parent_class;
    gint  scrollbar_spacing;
    void* scroll_child;
    void* move_focus_out;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ScrolledWindowPrivate {

};
struct SearchBarClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct SearchEntryClass {
          parent_class;
    void* search_changed;
    void* next_match;
    void* previous_match;
    void* stop_search;
};
struct SelectionData {

};
struct SeparatorClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct SeparatorMenuItemClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct SeparatorPrivate {

};
struct SeparatorToolItemClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct SeparatorToolItemPrivate {

};
struct SettingsClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct SettingsPrivate {

};
struct SettingsValue {
    void* origin;
          value;
};
struct ShortcutLabelClass {

};
struct ShortcutsGroupClass {

};
struct ShortcutsSectionClass {

};
struct ShortcutsShortcutClass {

};
struct ShortcutsWindowClass {
          parent_class;
    void* close;
    void* search;
};
struct SizeGroupClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct SizeGroupPrivate {

};
struct SocketAccessibleClass {
     parent_class;
};
struct SocketAccessiblePrivate {

};
struct SocketClass {
          parent_class;
    void* plug_added;
    void* plug_removed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct SocketPrivate {

};
struct SpinButtonAccessibleClass {
     parent_class;
};
struct SpinButtonAccessiblePrivate {

};
struct SpinButtonClass {
          parent_class;
    void* input;
    void* output;
    void* value_changed;
    void* change_value;
    void* wrapped;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct SpinButtonPrivate {

};
struct SpinnerAccessibleClass {
     parent_class;
};
struct SpinnerAccessiblePrivate {

};
struct SpinnerClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct SpinnerPrivate {

};
struct StackAccessibleClass {
     parent_class;
};
struct StackClass {
     parent_class;
};
struct StackSidebarClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct StackSidebarPrivate {

};
struct StackSwitcherClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct StatusIconClass {
          parent_class;
    void* activate;
    void* popup_menu;
    void* size_changed;
    void* button_press_event;
    void* button_release_event;
    void* scroll_event;
    void* query_tooltip;
    void* __gtk_reserved1;
    void* __gtk_reserved2;
    void* __gtk_reserved3;
    void* __gtk_reserved4;
};
struct StatusIconPrivate {

};
struct StatusbarAccessibleClass {
     parent_class;
};
struct StatusbarAccessiblePrivate {

};
struct StatusbarClass {
             parent_class;
    gpointer reserved;
    void*    text_pushed;
    void*    text_popped;
    void*    _gtk_reserved1;
    void*    _gtk_reserved2;
    void*    _gtk_reserved3;
    void*    _gtk_reserved4;
};
struct StatusbarPrivate {

};
struct StockItem {
    void* stock_id;
    void* label;
          modifier;
    guint keyval;
    void* translation_domain;
};
struct StyleClass {
          parent_class;
    void* realize;
    void* unrealize;
    void* copy;
    void* clone;
    void* init_from_rc;
    void* set_background;
    void* render_icon;
    void* draw_hline;
    void* draw_vline;
    void* draw_shadow;
    void* draw_arrow;
    void* draw_diamond;
    void* draw_box;
    void* draw_flat_box;
    void* draw_check;
    void* draw_option;
    void* draw_tab;
    void* draw_shadow_gap;
    void* draw_box_gap;
    void* draw_extension;
    void* draw_focus;
    void* draw_slider;
    void* draw_handle;
    void* draw_expander;
    void* draw_layout;
    void* draw_resize_grip;
    void* draw_spinner;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
    void* _gtk_reserved9;
    void* _gtk_reserved10;
    void* _gtk_reserved11;
};
struct StyleContextClass {
          parent_class;
    void* changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct StyleContextPrivate {

};
struct StylePropertiesClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct StylePropertiesPrivate {

};
struct StyleProviderIface {
          g_iface;
    void* get_style;
    void* get_style_property;
    void* get_icon_factory;
};
struct SwitchAccessibleClass {
     parent_class;
};
struct SwitchAccessiblePrivate {

};
struct SwitchClass {
          parent_class;
    void* activate;
    void* state_set;
    void* _switch_padding_1;
    void* _switch_padding_2;
    void* _switch_padding_3;
    void* _switch_padding_4;
    void* _switch_padding_5;
};
struct SwitchPrivate {

};
struct SymbolicColor {

};
struct TableChild {
    void*   widget;
    guint16 left_attach;
    guint16 right_attach;
    guint16 top_attach;
    guint16 bottom_attach;
    guint16 xpadding;
    guint16 ypadding;
    guint   xexpand  : 1;
    guint   yexpand  : 1;
    guint   xshrink  : 1;
    guint   yshrink  : 1;
    guint   xfill    : 1;
    guint   yfill    : 1;
};
struct TableClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TablePrivate {

};
struct TableRowCol {
    guint16 requisition;
    guint16 allocation;
    guint16 spacing;
    guint   need_expand  : 1;
    guint   need_shrink  : 1;
    guint   expand       : 1;
    guint   shrink       : 1;
    guint   empty        : 1;
};
struct TargetEntry {
    void* target;
    guint flags;
    guint info;
};
struct TargetList {

};
struct TargetPair {
          target;
    guint flags;
    guint info;
};
struct TearoffMenuItemClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TearoffMenuItemPrivate {

};
struct TextAppearance {
          bg_color;
          fg_color;
    gint  rise;
    guint underline         : 4;
    guint strikethrough     : 1;
    guint draw_bg           : 1;
    guint inside_selection  : 1;
    guint is_text           : 1;
};
struct TextAttributes {
    guint   refcount;
            appearance;
            justification;
            direction;
    void*   font;
    gdouble font_scale;
    gint    left_margin;
    gint    right_margin;
    gint    indent;
    gint    pixels_above_lines;
    gint    pixels_below_lines;
    gint    pixels_inside_wrap;
    void*   tabs;
            wrap_mode;
    void*   language;
    void*   pg_bg_color;
    guint   invisible       : 1;
    guint   bg_full_height  : 1;
    guint   editable        : 1;
    guint   no_fallback     : 1;
    void*   pg_bg_rgba;
    gint    letter_spacing;
};
struct TextBTree {

};
struct TextBufferClass {
          parent_class;
    void* insert_text;
    void* insert_pixbuf;
    void* insert_child_anchor;
    void* delete_range;
    void* changed;
    void* modified_changed;
    void* mark_set;
    void* mark_deleted;
    void* apply_tag;
    void* remove_tag;
    void* begin_user_action;
    void* end_user_action;
    void* paste_done;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TextBufferPrivate {

};
struct TextCellAccessibleClass {
     parent_class;
};
struct TextCellAccessiblePrivate {

};
struct TextChildAnchorClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TextIter {
    gpointer dummy1;
    gpointer dummy2;
    gint     dummy3;
    gint     dummy4;
    gint     dummy5;
    gint     dummy6;
    gint     dummy7;
    gint     dummy8;
    gpointer dummy9;
    gpointer dummy10;
    gint     dummy11;
    gint     dummy12;
    gint     dummy13;
    gpointer dummy14;
};
struct TextMarkClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TextTagClass {
          parent_class;
    void* event;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TextTagPrivate {

};
struct TextTagTableClass {
          parent_class;
    void* tag_changed;
    void* tag_added;
    void* tag_removed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TextTagTablePrivate {

};
struct TextViewAccessibleClass {
     parent_class;
};
struct TextViewAccessiblePrivate {

};
struct TextViewClass {
          parent_class;
    void* populate_popup;
    void* move_cursor;
    void* set_anchor;
    void* insert_at_cursor;
    void* delete_from_cursor;
    void* backspace;
    void* cut_clipboard;
    void* copy_clipboard;
    void* paste_clipboard;
    void* toggle_overwrite;
    void* create_buffer;
    void* draw_layer;
    void* extend_selection;
    void* insert_emoji;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TextViewPrivate {

};
struct ThemeEngine {

};
struct ThemingEngineClass {
          parent_class;
    void* render_line;
    void* render_background;
    void* render_frame;
    void* render_frame_gap;
    void* render_extension;
    void* render_check;
    void* render_option;
    void* render_arrow;
    void* render_expander;
    void* render_focus;
    void* render_layout;
    void* render_slider;
    void* render_handle;
    void* render_activity;
    void* render_icon_pixbuf;
    void* render_icon;
    void* render_icon_surface;
    void  padding;
};
struct ThemingEnginePrivate {

};
struct ToggleActionClass {
          parent_class;
    void* toggled;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ToggleActionEntry {
    void*    name;
    void*    stock_id;
    void*    label;
    void*    accelerator;
    void*    tooltip;
             callback;
    gboolean is_active;
};
struct ToggleActionPrivate {

};
struct ToggleButtonAccessibleClass {
     parent_class;
};
struct ToggleButtonAccessiblePrivate {

};
struct ToggleButtonClass {
          parent_class;
    void* toggled;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ToggleButtonPrivate {

};
struct ToggleToolButtonClass {
          parent_class;
    void* toggled;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ToggleToolButtonPrivate {

};
struct ToolButtonClass {
          parent_class;
          button_type;
    void* clicked;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ToolButtonPrivate {

};
struct ToolItemClass {
          parent_class;
    void* create_menu_proxy;
    void* toolbar_reconfigured;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ToolItemGroupClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ToolItemGroupPrivate {

};
struct ToolItemPrivate {

};
struct ToolPaletteClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ToolPalettePrivate {

};
struct ToolShellIface {
          g_iface;
    void* get_icon_size;
    void* get_orientation;
    void* get_style;
    void* get_relief_style;
    void* rebuild_menu;
    void* get_text_orientation;
    void* get_text_alignment;
    void* get_ellipsize_mode;
    void* get_text_size_group;
};
struct ToolbarClass {
          parent_class;
    void* orientation_changed;
    void* style_changed;
    void* popup_context_menu;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ToolbarPrivate {

};
struct ToplevelAccessibleClass {
     parent_class;
};
struct ToplevelAccessiblePrivate {

};
struct TreeDragDestIface {
          g_iface;
    void* drag_data_received;
    void* row_drop_possible;
};
struct TreeDragSourceIface {
          g_iface;
    void* row_draggable;
    void* drag_data_get;
    void* drag_data_delete;
};
struct TreeIter {
    gint     stamp;
    gpointer user_data;
    gpointer user_data2;
    gpointer user_data3;
};
struct TreeModelFilterClass {
          parent_class;
    void* visible;
    void* modify;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TreeModelFilterPrivate {

};
struct TreeModelIface {
          g_iface;
    void* row_changed;
    void* row_inserted;
    void* row_has_child_toggled;
    void* row_deleted;
    void* rows_reordered;
    void* get_flags;
    void* get_n_columns;
    void* get_column_type;
    void* get_iter;
    void* get_path;
    void* get_value;
    void* iter_next;
    void* iter_previous;
    void* iter_children;
    void* iter_has_child;
    void* iter_n_children;
    void* iter_nth_child;
    void* iter_parent;
    void* ref_node;
    void* unref_node;
};
struct TreeModelSortClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TreeModelSortPrivate {

};
struct TreePath {

};
struct TreeRowReference {

};
struct TreeSelectionClass {
          parent_class;
    void* changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TreeSelectionPrivate {

};
struct TreeSortableIface {
          g_iface;
    void* sort_column_changed;
    void* get_sort_column_id;
    void* set_sort_column_id;
    void* set_sort_func;
    void* set_default_sort_func;
    void* has_default_sort_func;
};
struct TreeStoreClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TreeStorePrivate {

};
struct TreeViewAccessibleClass {
     parent_class;
};
struct TreeViewAccessiblePrivate {

};
struct TreeViewClass {
          parent_class;
    void* row_activated;
    void* test_expand_row;
    void* test_collapse_row;
    void* row_expanded;
    void* row_collapsed;
    void* columns_changed;
    void* cursor_changed;
    void* move_cursor;
    void* select_all;
    void* unselect_all;
    void* select_cursor_row;
    void* toggle_cursor_row;
    void* expand_collapse_cursor_row;
    void* select_cursor_parent;
    void* start_interactive_search;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct TreeViewColumnClass {
          parent_class;
    void* clicked;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TreeViewColumnPrivate {

};
struct TreeViewPrivate {

};
struct UIManagerClass {
          parent_class;
    void* add_widget;
    void* actions_changed;
    void* connect_proxy;
    void* disconnect_proxy;
    void* pre_activate;
    void* post_activate;
    void* get_widget;
    void* get_action;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct UIManagerPrivate {

};
struct VBoxClass {
     parent_class;
};
struct VButtonBoxClass {
     parent_class;
};
struct VPanedClass {
     parent_class;
};
struct VScaleClass {
     parent_class;
};
struct VScrollbarClass {
     parent_class;
};
struct VSeparatorClass {
     parent_class;
};
struct ViewportClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct ViewportPrivate {

};
struct VolumeButtonClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct WidgetAccessibleClass {
          parent_class;
    void* notify_gtk;
};
struct WidgetAccessiblePrivate {

};
struct WidgetClass {
          parent_class;
    guint activate_signal;
    void* dispatch_child_properties_changed;
    void* destroy;
    void* show;
    void* show_all;
    void* hide;
    void* map;
    void* unmap;
    void* realize;
    void* unrealize;
    void* size_allocate;
    void* state_changed;
    void* state_flags_changed;
    void* parent_set;
    void* hierarchy_changed;
    void* style_set;
    void* direction_changed;
    void* grab_notify;
    void* child_notify;
    void* draw;
    void* get_request_mode;
    void* get_preferred_height;
    void* get_preferred_width_for_height;
    void* get_preferred_width;
    void* get_preferred_height_for_width;
    void* mnemonic_activate;
    void* grab_focus;
    void* focus;
    void* move_focus;
    void* keynav_failed;
    void* event;
    void* button_press_event;
    void* button_release_event;
    void* scroll_event;
    void* motion_notify_event;
    void* delete_event;
    void* destroy_event;
    void* key_press_event;
    void* key_release_event;
    void* enter_notify_event;
    void* leave_notify_event;
    void* configure_event;
    void* focus_in_event;
    void* focus_out_event;
    void* map_event;
    void* unmap_event;
    void* property_notify_event;
    void* selection_clear_event;
    void* selection_request_event;
    void* selection_notify_event;
    void* proximity_in_event;
    void* proximity_out_event;
    void* visibility_notify_event;
    void* window_state_event;
    void* damage_event;
    void* grab_broken_event;
    void* selection_get;
    void* selection_received;
    void* drag_begin;
    void* drag_end;
    void* drag_data_get;
    void* drag_data_delete;
    void* drag_leave;
    void* drag_motion;
    void* drag_drop;
    void* drag_data_received;
    void* drag_failed;
    void* popup_menu;
    void* show_help;
    void* get_accessible;
    void* screen_changed;
    void* can_activate_accel;
    void* composited_changed;
    void* query_tooltip;
    void* compute_expand;
    void* adjust_size_request;
    void* adjust_size_allocation;
    void* style_updated;
    void* touch_event;
    void* get_preferred_height_and_baseline_for_width;
    void* adjust_baseline_request;
    void* adjust_baseline_allocation;
    void* queue_draw_region;
    void* priv;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
};
struct WidgetClassPrivate {

};
struct WidgetPath {

};
struct WidgetPrivate {

};
struct WindowAccessibleClass {
     parent_class;
};
struct WindowAccessiblePrivate {

};
struct WindowClass {
          parent_class;
    void* set_focus;
    void* activate_focus;
    void* activate_default;
    void* keys_changed;
    void* enable_debugging;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
};
struct WindowGeometryInfo {

};
struct WindowGroupClass {
          parent_class;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct WindowGroupPrivate {

};
struct WindowPrivate {

};
struct _MountOperationHandler {

};
struct _MountOperationHandlerIface {
          parent_iface;
    void* handle_ask_password;
    void* handle_ask_question;
    void* handle_close;
    void* handle_show_processes;
};
struct _MountOperationHandlerProxy {
    gpointer parent_instance;
    void*    priv;
};
struct _MountOperationHandlerProxyClass {
     parent_class;
};
struct _MountOperationHandlerProxyPrivate {

};
struct _MountOperationHandlerSkeleton {
    gpointer parent_instance;
    void*    priv;
};
struct _MountOperationHandlerSkeletonClass {
     parent_class;
};
struct _MountOperationHandlerSkeletonPrivate {

};