#include <atk/atk.h>
#include <glib-object.h>

extern void goPanic(const char*);

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 12))
AtkHyperlink* atk_hyperlink_impl_get_hyperlink(AtkHyperlinkImpl* v) __attribute__((weak)) {
	goPanic("atk_hyperlink_impl_get_hyperlink: library too old: needs at least 1.12");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 12))
const gchar* atk_document_get_attribute_value(AtkDocument* v, const gchar* _0) __attribute__((weak)) {
	goPanic("atk_document_get_attribute_value: library too old: needs at least 1.12");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 12))
const gchar* atk_image_get_image_locale(AtkImage* v) __attribute__((weak)) {
	goPanic("atk_image_get_image_locale: library too old: needs at least 1.12");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 12))
const gchar* atk_streamable_content_get_uri(AtkStreamableContent* v, const gchar* _0) __attribute__((weak)) {
	goPanic("atk_streamable_content_get_uri: library too old: needs at least 1.12");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 12))
gboolean atk_document_set_attribute_value(AtkDocument* v, const gchar* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("atk_document_set_attribute_value: library too old: needs at least 1.12");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 12))
gdouble atk_component_get_alpha(AtkComponent* v) __attribute__((weak)) {
	goPanic("atk_component_get_alpha: library too old: needs at least 1.12");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 12))
void atk_value_get_minimum_increment(AtkValue* v, GValue* _0) __attribute__((weak)) {
	goPanic("atk_value_get_minimum_increment: library too old: needs at least 1.12");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 13))
const AtkMisc* atk_misc_get_instance(void) __attribute__((weak)) {
	goPanic("atk_misc_get_instance: library too old: needs at least 1.13");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 13))
void atk_misc_threads_enter(AtkMisc* v) __attribute__((weak)) {
	goPanic("atk_misc_threads_enter: library too old: needs at least 1.13");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 13))
void atk_misc_threads_leave(AtkMisc* v) __attribute__((weak)) {
	goPanic("atk_misc_threads_leave: library too old: needs at least 1.13");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 20))
const gchar* atk_get_version(void) __attribute__((weak)) {
	goPanic("atk_get_version: library too old: needs at least 1.20");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 3))
AtkTextRange** atk_text_get_bounded_ranges(AtkText* v, AtkTextRectangle* _0, AtkCoordType _1, AtkTextClipType _2, AtkTextClipType _3) __attribute__((weak)) {
	goPanic("atk_text_get_bounded_ranges: library too old: needs at least 1.3");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 3))
void atk_text_get_range_extents(AtkText* v, gint _0, gint _1, AtkCoordType _2, AtkTextRectangle* _3) __attribute__((weak)) {
	goPanic("atk_text_get_range_extents: library too old: needs at least 1.3");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 30))
AtkObject* atk_plug_new(void) __attribute__((weak)) {
	goPanic("atk_plug_new: library too old: needs at least 1.30");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 30))
gboolean atk_socket_is_occupied(AtkSocket* v) __attribute__((weak)) {
	goPanic("atk_socket_is_occupied: library too old: needs at least 1.30");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 30))
gchar* atk_plug_get_id(AtkPlug* v) __attribute__((weak)) {
	goPanic("atk_plug_get_id: library too old: needs at least 1.30");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 30))
void atk_socket_embed(AtkSocket* v, const gchar* _0) __attribute__((weak)) {
	goPanic("atk_socket_embed: library too old: needs at least 1.30");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 4))
gboolean atk_hyperlink_is_selected_link(AtkHyperlink* v) __attribute__((weak)) {
	goPanic("atk_hyperlink_is_selected_link: library too old: needs at least 1.4");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 6))
AtkObject* atk_get_focus_object(void) __attribute__((weak)) {
	goPanic("atk_get_focus_object: library too old: needs at least 1.6");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 9))
void atk_relation_add_target(AtkRelation* v, AtkObject* _0) __attribute__((weak)) {
	goPanic("atk_relation_add_target: library too old: needs at least 1.9");
}
#endif

#if (ATK_MAJOR_VERSION < 1 || (ATK_MAJOR_VERSION == 1 && ATK_MINOR_VERSION < 9))
void atk_relation_set_add_relation_by_type(AtkRelationSet* v, AtkRelationType _0, AtkObject* _1) __attribute__((weak)) {
	goPanic("atk_relation_set_add_relation_by_type: library too old: needs at least 1.9");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 10))
gchar* atk_text_get_string_at_offset(AtkText* v, gint _0, AtkTextGranularity _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("atk_text_get_string_at_offset: library too old: needs at least 2.10");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
AtkObject* atk_table_cell_get_table(AtkTableCell* v) __attribute__((weak)) {
	goPanic("atk_table_cell_get_table: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
AtkRange* atk_range_copy(AtkRange* v) __attribute__((weak)) {
	goPanic("atk_range_copy: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
AtkRange* atk_range_new(gdouble _0, gdouble _1, const gchar* _2) __attribute__((weak)) {
	goPanic("atk_range_new: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
AtkRange* atk_value_get_range(AtkValue* v) __attribute__((weak)) {
	goPanic("atk_value_get_range: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
GSList* atk_value_get_sub_ranges(AtkValue* v) __attribute__((weak)) {
	goPanic("atk_value_get_sub_ranges: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
const gchar* atk_range_get_description(AtkRange* v) __attribute__((weak)) {
	goPanic("atk_range_get_description: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
gboolean atk_table_cell_get_position(AtkTableCell* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("atk_table_cell_get_position: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
gboolean atk_table_cell_get_row_column_span(AtkTableCell* v, gint* _0, gint* _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("atk_table_cell_get_row_column_span: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
gdouble atk_range_get_lower_limit(AtkRange* v) __attribute__((weak)) {
	goPanic("atk_range_get_lower_limit: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
gdouble atk_range_get_upper_limit(AtkRange* v) __attribute__((weak)) {
	goPanic("atk_range_get_upper_limit: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
gdouble atk_value_get_increment(AtkValue* v) __attribute__((weak)) {
	goPanic("atk_value_get_increment: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
gint atk_document_get_current_page_number(AtkDocument* v) __attribute__((weak)) {
	goPanic("atk_document_get_current_page_number: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
gint atk_document_get_page_count(AtkDocument* v) __attribute__((weak)) {
	goPanic("atk_document_get_page_count: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
gint atk_table_cell_get_column_span(AtkTableCell* v) __attribute__((weak)) {
	goPanic("atk_table_cell_get_column_span: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
gint atk_table_cell_get_row_span(AtkTableCell* v) __attribute__((weak)) {
	goPanic("atk_table_cell_get_row_span: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
void atk_value_get_value_and_text(AtkValue* v, gdouble* _0, gchar** _1) __attribute__((weak)) {
	goPanic("atk_value_get_value_and_text: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 12))
void atk_value_set_value(AtkValue* v, const gdouble _0) __attribute__((weak)) {
	goPanic("atk_value_set_value: library too old: needs at least 2.12");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 30))
gboolean atk_component_scroll_to(AtkComponent* v, AtkScrollType _0) __attribute__((weak)) {
	goPanic("atk_component_scroll_to: library too old: needs at least 2.30");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 30))
gboolean atk_component_scroll_to_point(AtkComponent* v, AtkCoordType _0, gint _1, gint _2) __attribute__((weak)) {
	goPanic("atk_component_scroll_to_point: library too old: needs at least 2.30");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 32))
gboolean atk_text_scroll_substring_to(AtkText* v, gint _0, gint _1, AtkScrollType _2) __attribute__((weak)) {
	goPanic("atk_text_scroll_substring_to: library too old: needs at least 2.32");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 32))
gboolean atk_text_scroll_substring_to_point(AtkText* v, gint _0, gint _1, AtkCoordType _2, gint _3, gint _4) __attribute__((weak)) {
	goPanic("atk_text_scroll_substring_to_point: library too old: needs at least 2.32");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 34))
const gchar* atk_object_get_accessible_id(AtkObject* v) __attribute__((weak)) {
	goPanic("atk_object_get_accessible_id: library too old: needs at least 2.34");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 34))
void atk_object_set_accessible_id(AtkObject* v, const gchar* _0) __attribute__((weak)) {
	goPanic("atk_object_set_accessible_id: library too old: needs at least 2.34");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 35))
void atk_plug_set_child(AtkPlug* v, AtkObject* _0) __attribute__((weak)) {
	goPanic("atk_plug_set_child: library too old: needs at least 2.35");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 8))
const gchar* atk_object_get_object_locale(AtkObject* v) __attribute__((weak)) {
	goPanic("atk_object_get_object_locale: library too old: needs at least 2.8");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 8))
guint atk_get_binary_age(void) __attribute__((weak)) {
	goPanic("atk_get_binary_age: library too old: needs at least 2.8");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 8))
guint atk_get_interface_age(void) __attribute__((weak)) {
	goPanic("atk_get_interface_age: library too old: needs at least 2.8");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 8))
guint atk_get_major_version(void) __attribute__((weak)) {
	goPanic("atk_get_major_version: library too old: needs at least 2.8");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 8))
guint atk_get_micro_version(void) __attribute__((weak)) {
	goPanic("atk_get_micro_version: library too old: needs at least 2.8");
}
#endif

#if (ATK_MAJOR_VERSION < 2 || (ATK_MAJOR_VERSION == 2 && ATK_MINOR_VERSION < 8))
guint atk_get_minor_version(void) __attribute__((weak)) {
	goPanic("atk_get_minor_version: library too old: needs at least 2.8");
}
#endif
