#include <stdlib.h>
#include <glib.h>

struct ATContextClass {

};
struct AccessibleInterface {

};
struct ActionableInterface {
          g_iface;
    void* get_action_name;
    void* set_action_name;
    void* get_action_target_value;
    void* set_action_target_value;
};
struct ActivateActionClass {

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
struct AlternativeTriggerClass {

};
struct AnyFilterClass {

};
struct ApplicationClass {
          parent_class;
    void* window_added;
    void* window_removed;
    void  padding;
};
struct ApplicationWindowClass {
         parent_class;
    void padding;
};
struct BinLayoutClass {
     parent_class;
};
struct Bitset {

};
struct BitsetIter {
    void private_data;
};
struct BookmarkListClass {
     parent_class;
};
struct BoolFilterClass {
     parent_class;
};
struct Border {
    gint16 left;
    gint16 right;
    gint16 top;
    gint16 bottom;
};
struct BoxClass {
         parent_class;
    void padding;
};
struct BoxLayoutClass {
     parent_class;
};
struct BuildableIface {
          g_iface;
    void* set_id;
    void* get_id;
    void* add_child;
    void* set_buildable_property;
    void* construct_child;
    void* custom_tag_start;
    void* custom_tag_end;
    void* custom_finished;
    void* parser_finished;
    void* get_internal_child;
};
struct BuildableParseContext {

};
struct BuildableParser {
    void* start_element;
    void* end_element;
    void* text;
    void* error;
    void  padding;
};
struct BuilderCScopeClass {
     parent_class;
};
struct BuilderClass {

};
struct BuilderListItemFactoryClass {

};
struct BuilderScopeInterface {
          g_iface;
    void* get_type_from_name;
    void* get_type_from_function;
    void* create_closure;
};
struct ButtonClass {
          parent_class;
    void* clicked;
    void* activate;
    void  padding;
};
struct ButtonPrivate {

};
struct CallbackActionClass {

};
struct CellAreaClass {
          parent_class;
    void* add;
    void* remove;
    void* foreach;
    void* foreach_alloc;
    void* event;
    void* snapshot;
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
    void  padding;
};
struct CellAreaContextClass {
          parent_class;
    void* allocate;
    void* reset;
    void* get_preferred_height_for_width;
    void* get_preferred_width_for_height;
    void  padding;
};
struct CellAreaContextPrivate {

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
struct CellRendererClass {
          parent_class;
    void* get_request_mode;
    void* get_preferred_width;
    void* get_preferred_height_for_width;
    void* get_preferred_height;
    void* get_preferred_width_for_height;
    void* get_aligned_area;
    void* snapshot;
    void* activate;
    void* start_editing;
    void* editing_canceled;
    void* editing_started;
    void  padding;
};
struct CellRendererClassPrivate {

};
struct CellRendererPrivate {

};
struct CellRendererTextClass {
          parent_class;
    void* edited;
    void  padding;
};
struct CenterBoxClass {

};
struct CenterLayoutClass {
     parent_class;
};
struct CheckButtonClass {
          parent_class;
    void* toggled;
    void* activate;
    void  padding;
};
struct ColorChooserInterface {
          base_interface;
    void* get_rgba;
    void* set_rgba;
    void* add_palette;
    void* color_activated;
    void  padding;
};
struct ColumnViewClass {

};
struct ColumnViewColumnClass {

};
struct ComboBoxClass {
          parent_class;
    void* changed;
    void* format_entry_text;
    void  padding;
};
struct ConstraintClass {
     parent_class;
};
struct ConstraintGuideClass {
     parent_class;
};
struct ConstraintLayoutChildClass {
     parent_class;
};
struct ConstraintLayoutClass {
     parent_class;
};
struct ConstraintTargetInterface {

};
struct CssLocation {
    gsize bytes;
    gsize chars;
    gsize lines;
    gsize line_bytes;
    gsize line_chars;
};
struct CssProviderClass {

};
struct CssProviderPrivate {

};
struct CssSection {

};
struct CssStyleChange {

};
struct CustomFilterClass {
     parent_class;
};
struct CustomLayoutClass {
     parent_class;
};
struct CustomSorterClass {
     parent_class;
};
struct DialogClass {
          parent_class;
    void* response;
    void* close;
    void  padding;
};
struct DirectoryListClass {
     parent_class;
};
struct DragIconClass {
     parent_class;
};
struct DragSourceClass {

};
struct DrawingAreaClass {
          parent_class;
    void* resize;
    void  padding;
};
struct DropControllerMotionClass {

};
struct DropDownClass {
     parent_class;
};
struct DropTargetAsyncClass {

};
struct DropTargetClass {

};
struct EditableInterface {
          base_iface;
    void* insert_text;
    void* delete_text;
    void* changed;
    void* get_text;
    void* do_insert_text;
    void* do_delete_text;
    void* get_selection_bounds;
    void* set_selection_bounds;
    void* get_delegate;
};
struct EditableLabelClass {
     parent_class;
};
struct EmojiChooserClass {

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
struct EntryClass {
          parent_class;
    void* activate;
    void  padding;
};
struct EventControllerClass {

};
struct EventControllerFocusClass {

};
struct EventControllerKeyClass {

};
struct EventControllerLegacyClass {

};
struct EventControllerMotionClass {

};
struct EventControllerScrollClass {

};
struct EveryFilterClass {

};
struct ExpressionWatch {

};
struct FileChooserNativeClass {
     parent_class;
};
struct FilterClass {
          parent_class;
    void* match;
    void* get_strictness;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct FilterListModelClass {
     parent_class;
};
struct FixedClass {
         parent_class;
    void padding;
};
struct FixedLayoutChildClass {
     parent_class;
};
struct FixedLayoutClass {
     parent_class;
};
struct FlattenListModelClass {
     parent_class;
};
struct FlowBoxChildClass {
          parent_class;
    void* activate;
    void  padding;
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
struct FrameClass {
          parent_class;
    void* compute_child_allocation;
    void  padding;
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
struct GestureClickClass {

};
struct GestureDragClass {

};
struct GestureLongPressClass {

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
struct GridClass {
         parent_class;
    void padding;
};
struct GridLayoutChildClass {
     parent_class;
};
struct GridLayoutClass {
     parent_class;
};
struct GridViewClass {

};
struct IMContextClass {
          parent_class;
    void* preedit_start;
    void* preedit_end;
    void* preedit_changed;
    void* commit;
    void* retrieve_surrounding;
    void* delete_surrounding;
    void* set_client_widget;
    void* get_preedit_string;
    void* filter_keypress;
    void* focus_in;
    void* focus_out;
    void* reset;
    void* set_cursor_location;
    void* set_use_preedit;
    void* set_surrounding;
    void* get_surrounding;
    void* set_surrounding_with_selection;
    void* get_surrounding_with_selection;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
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
struct KeyvalTriggerClass {

};
struct LayoutChildClass {
     parent_class;
};
struct LayoutManagerClass {
          parent_class;
    void* get_request_mode;
    void* measure;
    void* allocate;
          layout_child_type;
    void* create_layout_child;
    void* root;
    void* unroot;
    void  _padding;
};
struct ListBaseClass {

};
struct ListBoxRowClass {
          parent_class;
    void* activate;
    void  padding;
};
struct ListItemClass {

};
struct ListItemFactoryClass {

};
struct ListStoreClass {
         parent_class;
    void padding;
};
struct ListStorePrivate {

};
struct ListViewClass {

};
struct MapListModelClass {
     parent_class;
};
struct MediaControlsClass {
     parent_class;
};
struct MediaFileClass {
          parent_class;
    void* open;
    void* close;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct MediaStreamClass {
          parent_class;
    void* play;
    void* pause;
    void* seek;
    void* update_audio;
    void* realize;
    void* unrealize;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct MessageDialogClass {

};
struct MnemonicActionClass {

};
struct MnemonicTriggerClass {

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
struct MultiFilterClass {

};
struct MultiSelectionClass {
     parent_class;
};
struct MultiSorterClass {
     parent_class;
};
struct NamedActionClass {

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
struct NativeInterface {

};
struct NeverTriggerClass {

};
struct NoSelectionClass {
     parent_class;
};
struct NothingActionClass {

};
struct NumericSorterClass {
     parent_class;
};
struct OrientableIface {
     base_iface;
};
struct OverlayLayoutChildClass {
     parent_class;
};
struct OverlayLayoutClass {
     parent_class;
};
struct PadActionEntry {
          type;
    int   index;
    int   mode;
    void* label;
    void* action_name;
};
struct PadControllerClass {

};
struct PageRange {
    int start;
    int end;
};
struct PaperSize {

};
struct PasswordEntryClass {

};
struct PictureClass {
     parent_class;
};
struct PopoverClass {
          parent_class;
    void* closed;
    void* activate_default;
    void  reserved;
};
struct PrintBackend {

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
    void  padding;
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
struct RangeClass {
          parent_class;
    void* value_changed;
    void* adjust_bounds;
    void* move_slider;
    void* get_range_border;
    void* change_value;
    void  padding;
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
struct RequestedSize {
    gpointer data;
    int      minimum_size;
    int      natural_size;
};
struct Requisition {
    int width;
    int height;
};
struct RootInterface {

};
struct ScaleButtonClass {
          parent_class;
    void* value_changed;
    void  padding;
};
struct ScaleClass {
          parent_class;
    void* get_layout_offsets;
    void  padding;
};
struct ScrollableInterface {
          base_iface;
    void* get_border;
};
struct SelectionFilterModelClass {
     parent_class;
};
struct SelectionModelInterface {
          g_iface;
    void* is_selected;
    void* get_selection_in_range;
    void* select_item;
    void* unselect_item;
    void* select_range;
    void* unselect_range;
    void* select_all;
    void* unselect_all;
    void* set_selection;
};
struct ShortcutActionClass {

};
struct ShortcutClass {
     parent_class;
};
struct ShortcutControllerClass {

};
struct ShortcutLabelClass {

};
struct ShortcutManagerInterface {
          g_iface;
    void* add_controller;
    void* remove_controller;
};
struct ShortcutTriggerClass {

};
struct ShortcutsGroupClass {

};
struct ShortcutsSectionClass {

};
struct ShortcutsShortcutClass {

};
struct SignalActionClass {

};
struct SignalListItemFactoryClass {

};
struct SingleSelectionClass {
     parent_class;
};
struct SliceListModelClass {
     parent_class;
};
struct SnapshotClass {

};
struct SortListModelClass {
     parent_class;
};
struct SorterClass {
          parent_class;
    void* compare;
    void* get_order;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
    void* _gtk_reserved5;
    void* _gtk_reserved6;
    void* _gtk_reserved7;
    void* _gtk_reserved8;
};
struct StringFilterClass {
     parent_class;
};
struct StringListClass {
     parent_class;
};
struct StringObjectClass {
     parent_class;
};
struct StringSorterClass {
     parent_class;
};
struct StyleContextClass {
          parent_class;
    void* changed;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TextBufferClass {
          parent_class;
    void* insert_text;
    void* insert_paintable;
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
    void* undo;
    void* redo;
    void* _gtk_reserved1;
    void* _gtk_reserved2;
    void* _gtk_reserved3;
    void* _gtk_reserved4;
};
struct TextBufferPrivate {

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
    int      dummy3;
    int      dummy4;
    int      dummy5;
    int      dummy6;
    int      dummy7;
    int      dummy8;
    gpointer dummy9;
    gpointer dummy10;
    int      dummy11;
    int      dummy12;
    int      dummy13;
    gpointer dummy14;
};
struct TextMarkClass {
         parent_class;
    void padding;
};
struct TextTagClass {
         parent_class;
    void padding;
};
struct TextTagPrivate {

};
struct TextViewClass {
          parent_class;
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
    void* snapshot_layer;
    void* extend_selection;
    void* insert_emoji;
    void  padding;
};
struct TextViewPrivate {

};
struct ToggleButtonClass {
          parent_class;
    void* toggled;
    void  padding;
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
struct TreeExpanderClass {
     parent_class;
};
struct TreeIter {
    int      stamp;
    gpointer user_data;
    gpointer user_data2;
    gpointer user_data3;
};
struct TreeListModelClass {
     parent_class;
};
struct TreeListRowClass {
     parent_class;
};
struct TreeListRowSorterClass {
     parent_class;
};
struct TreeModelFilterClass {
          parent_class;
    void* visible;
    void* modify;
    void  padding;
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
    void padding;
};
struct TreeModelSortPrivate {

};
struct TreePath {

};
struct TreeRowReference {

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
    void padding;
};
struct TreeStorePrivate {

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
    void  _reserved;
};
struct VideoClass {
     parent_class;
};
struct WidgetClass {
          parent_class;
    void* show;
    void* hide;
    void* map;
    void* unmap;
    void* realize;
    void* unrealize;
    void* root;
    void* unroot;
    void* size_allocate;
    void* state_flags_changed;
    void* direction_changed;
    void* get_request_mode;
    void* measure;
    void* mnemonic_activate;
    void* grab_focus;
    void* focus;
    void* set_focus_child;
    void* move_focus;
    void* keynav_failed;
    void* query_tooltip;
    void* compute_expand;
    void* css_changed;
    void* system_setting_changed;
    void* snapshot;
    void* contains;
    void* priv;
    void  padding;
};
struct WidgetClassPrivate {

};
struct WidgetPaintableClass {
     parent_class;
};
struct WidgetPrivate {

};
struct WindowClass {
          parent_class;
    void* activate_focus;
    void* activate_default;
    void* keys_changed;
    void* enable_debugging;
    void* close_request;
    void  padding;
};
struct WindowControlsClass {
     parent_class;
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
struct WindowHandleClass {
     parent_class;
};