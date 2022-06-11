#include <stdlib.h>
#include <glib.h>

struct ActionIface {
          parent;
    void* do_action;
    void* get_n_actions;
    void* get_description;
    void* get_name;
    void* get_keybinding;
    void* set_description;
    void* get_localized_name;
};
struct Attribute {
    void* name;
    void* value;
};
struct ComponentIface {
          parent;
    void* add_focus_handler;
    void* contains;
    void* ref_accessible_at_point;
    void* get_extents;
    void* get_position;
    void* get_size;
    void* grab_focus;
    void* remove_focus_handler;
    void* set_extents;
    void* set_position;
    void* set_size;
    void* get_layer;
    void* get_mdi_zorder;
    void* bounds_changed;
    void* get_alpha;
    void* scroll_to;
    void* scroll_to_point;
};
struct DocumentIface {
          parent;
    void* get_document_type;
    void* get_document;
    void* get_document_locale;
    void* get_document_attributes;
    void* get_document_attribute_value;
    void* set_document_attribute;
    void* get_current_page_number;
    void* get_page_count;
};
struct EditableTextIface {
          parent_interface;
    void* set_run_attributes;
    void* set_text_contents;
    void* insert_text;
    void* copy_text;
    void* cut_text;
    void* delete_text;
    void* paste_text;
};
struct GObjectAccessibleClass {
             parent_class;
    gpointer pad1;
    gpointer pad2;
};
struct HyperlinkClass {
             parent;
    void*    get_uri;
    void*    get_object;
    void*    get_end_index;
    void*    get_start_index;
    void*    is_valid;
    void*    get_n_anchors;
    void*    link_state;
    void*    is_selected_link;
    void*    link_activated;
    gpointer pad1;
};
struct HyperlinkImplIface {
          parent;
    void* get_hyperlink;
};
struct HypertextIface {
          parent;
    void* get_link;
    void* get_n_links;
    void* get_link_index;
    void* link_selected;
};
struct ImageIface {
          parent;
    void* get_image_position;
    void* get_image_description;
    void* get_image_size;
    void* set_image_description;
    void* get_image_locale;
};
struct Implementor {

};
struct KeyEventStruct {
    gint    type;
    guint   state;
    guint   keyval;
    gint    length;
    void*   string;
    guint16 keycode;
    guint32 timestamp;
};
struct MiscClass {
          parent;
    void* threads_enter;
    void* threads_leave;
    void  vfuncs;
};
struct NoOpObjectClass {
     parent_class;
};
struct NoOpObjectFactoryClass {
     parent_class;
};
struct ObjectClass {
             parent;
    void*    get_name;
    void*    get_description;
    void*    get_parent;
    void*    get_n_children;
    void*    ref_child;
    void*    get_index_in_parent;
    void*    ref_relation_set;
    void*    get_role;
    void*    get_layer;
    void*    get_mdi_zorder;
    void*    ref_state_set;
    void*    set_name;
    void*    set_description;
    void*    set_parent;
    void*    set_role;
    void*    connect_property_change_handler;
    void*    remove_property_change_handler;
    void*    initialize;
    void*    children_changed;
    void*    focus_event;
    void*    property_change;
    void*    state_change;
    void*    visible_data_changed;
    void*    active_descendant_changed;
    void*    get_attributes;
    void*    get_object_locale;
    gpointer pad1;
};
struct ObjectFactoryClass {
             parent_class;
    void*    create_accessible;
    void*    invalidate;
    void*    get_accessible_type;
    gpointer pad1;
    gpointer pad2;
};
struct PlugClass {
          parent_class;
    void* get_object_id;
};
struct PropertyValues {
    void* property_name;
          old_value;
          new_value;
};
struct Range {

};
struct Rectangle {
    gint x;
    gint y;
    gint width;
    gint height;
};
struct RegistryClass {
     parent_class;
};
struct RelationClass {
     parent;
};
struct RelationSetClass {
             parent;
    gpointer pad1;
    gpointer pad2;
};
struct SelectionIface {
          parent;
    void* add_selection;
    void* clear_selection;
    void* ref_selection;
    void* get_selection_count;
    void* is_child_selected;
    void* remove_selection;
    void* select_all_selection;
    void* selection_changed;
};
struct SocketClass {
          parent_class;
    void* embed;
};
struct StateSetClass {
     parent;
};
struct StreamableContentIface {
             parent;
    void*    get_n_mime_types;
    void*    get_mime_type;
    void*    get_stream;
    void*    get_uri;
    gpointer pad1;
    gpointer pad2;
    gpointer pad3;
};
struct TableCellIface {
          parent;
    void* get_column_span;
    void* get_column_header_cells;
    void* get_position;
    void* get_row_span;
    void* get_row_header_cells;
    void* get_row_column_span;
    void* get_table;
};
struct TableIface {
          parent;
    void* ref_at;
    void* get_index_at;
    void* get_column_at_index;
    void* get_row_at_index;
    void* get_n_columns;
    void* get_n_rows;
    void* get_column_extent_at;
    void* get_row_extent_at;
    void* get_caption;
    void* get_column_description;
    void* get_column_header;
    void* get_row_description;
    void* get_row_header;
    void* get_summary;
    void* set_caption;
    void* set_column_description;
    void* set_column_header;
    void* set_row_description;
    void* set_row_header;
    void* set_summary;
    void* get_selected_columns;
    void* get_selected_rows;
    void* is_column_selected;
    void* is_row_selected;
    void* is_selected;
    void* add_row_selection;
    void* remove_row_selection;
    void* add_column_selection;
    void* remove_column_selection;
    void* row_inserted;
    void* column_inserted;
    void* row_deleted;
    void* column_deleted;
    void* row_reordered;
    void* column_reordered;
    void* model_changed;
};
struct TextIface {
          parent;
    void* get_text;
    void* get_text_after_offset;
    void* get_text_at_offset;
    void* get_character_at_offset;
    void* get_text_before_offset;
    void* get_caret_offset;
    void* get_run_attributes;
    void* get_default_attributes;
    void* get_character_extents;
    void* get_character_count;
    void* get_offset_at_point;
    void* get_n_selections;
    void* get_selection;
    void* add_selection;
    void* remove_selection;
    void* set_selection;
    void* set_caret_offset;
    void* text_changed;
    void* text_caret_moved;
    void* text_selection_changed;
    void* text_attributes_changed;
    void* get_range_extents;
    void* get_bounded_ranges;
    void* get_string_at_offset;
    void* scroll_substring_to;
    void* scroll_substring_to_point;
};
struct TextRange {
          bounds;
    gint  start_offset;
    gint  end_offset;
    void* content;
};
struct TextRectangle {
    gint x;
    gint y;
    gint width;
    gint height;
};
struct UtilClass {
          parent;
    void* add_global_event_listener;
    void* remove_global_event_listener;
    void* add_key_event_listener;
    void* remove_key_event_listener;
    void* get_root;
    void* get_toolkit_name;
    void* get_toolkit_version;
};
struct ValueIface {
          parent;
    void* get_current_value;
    void* get_maximum_value;
    void* get_minimum_value;
    void* set_current_value;
    void* get_minimum_increment;
    void* get_value_and_text;
    void* get_range;
    void* get_increment;
    void* get_sub_ranges;
    void* set_value;
};
struct WindowIface {
     parent;
};