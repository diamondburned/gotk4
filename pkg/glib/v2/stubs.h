#include <glib-object.h>
#include <glib.h>

extern void goPanic(const char*);

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 10))
const gchar* g_intern_static_string(const gchar* _0) {
	goPanic("g_intern_static_string: library too old: needs at least 2.10");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 10))
const gchar* g_intern_string(const gchar* _0) {
	goPanic("g_intern_string: library too old: needs at least 2.10");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 10))
gboolean g_main_context_is_owner(GMainContext* v) {
	goPanic("g_main_context_is_owner: library too old: needs at least 2.10");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
GSource* g_main_current_source(void) {
	goPanic("g_main_current_source: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
gboolean g_source_is_destroyed(GSource* v) {
	goPanic("g_source_is_destroyed: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
gboolean g_time_val_from_iso8601(const gchar* _0, GTimeVal* _1) {
	goPanic("g_time_val_from_iso8601: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
gboolean g_unichar_iswide_cjk(gunichar _0) {
	goPanic("g_unichar_iswide_cjk: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
gchar* g_time_val_to_iso8601(GTimeVal* v) {
	goPanic("g_time_val_to_iso8601: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
gdouble g_key_file_get_double(GKeyFile* v, const gchar* _0, const gchar* _1) {
	goPanic("g_key_file_get_double: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
gdouble* g_key_file_get_double_list(GKeyFile* v, const gchar* _0, const gchar* _1, gsize* _2) {
	goPanic("g_key_file_get_double_list: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
void g_hash_table_remove_all(GHashTable* _0) {
	goPanic("g_hash_table_remove_all: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
void g_hash_table_steal_all(GHashTable* _0) {
	goPanic("g_hash_table_steal_all: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
void g_key_file_set_double(GKeyFile* v, const gchar* _0, const gchar* _1, gdouble _2) {
	goPanic("g_key_file_set_double: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
void g_key_file_set_double_list(GKeyFile* v, const gchar* _0, const gchar* _1, gdouble* _2, gsize _3) {
	goPanic("g_key_file_set_double_list: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 12))
void g_source_set_funcs(GSource* v, GSourceFuncs* _0) {
	goPanic("g_source_set_funcs: library too old: needs at least 2.12");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
GRegex* g_match_info_get_regex(const GMatchInfo* v) {
	goPanic("g_match_info_get_regex: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
GRegex* g_regex_new(const gchar* _0, GRegexCompileFlags _1, GRegexMatchFlags _2) {
	goPanic("g_regex_new: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
GSource* g_timeout_source_new_seconds(guint _0) {
	goPanic("g_timeout_source_new_seconds: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
GUnicodeScript g_unichar_get_script(gunichar _0) {
	goPanic("g_unichar_get_script: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
const gchar* g_get_user_special_dir(GUserDirectory _0) {
	goPanic("g_get_user_special_dir: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
const gchar* g_match_info_get_string(const GMatchInfo* v) {
	goPanic("g_match_info_get_string: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
const gchar* g_regex_get_pattern(const GRegex* v) {
	goPanic("g_regex_get_pattern: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_key_file_load_from_dirs(GKeyFile* v, const gchar* _0, const gchar** _1, gchar** _2, GKeyFileFlags _3) {
	goPanic("g_key_file_load_from_dirs: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_match_info_fetch_named_pos(const GMatchInfo* v, const gchar* _0, gint* _1, gint* _2) {
	goPanic("g_match_info_fetch_named_pos: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_match_info_fetch_pos(const GMatchInfo* v, gint _0, gint* _1, gint* _2) {
	goPanic("g_match_info_fetch_pos: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_match_info_is_partial_match(const GMatchInfo* v) {
	goPanic("g_match_info_is_partial_match: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_match_info_matches(const GMatchInfo* v) {
	goPanic("g_match_info_matches: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_match_info_next(GMatchInfo* v) {
	goPanic("g_match_info_next: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_regex_check_replacement(const gchar* _0, gboolean* _1) {
	goPanic("g_regex_check_replacement: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_regex_match(const GRegex* v, const gchar* _0, GRegexMatchFlags _1, GMatchInfo** _2) {
	goPanic("g_regex_match: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_regex_match_all(const GRegex* v, const gchar* _0, GRegexMatchFlags _1, GMatchInfo** _2) {
	goPanic("g_regex_match_all: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_regex_match_simple(const gchar* _0, const gchar* _1, GRegexCompileFlags _2, GRegexMatchFlags _3) {
	goPanic("g_regex_match_simple: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_unichar_ismark(gunichar _0) {
	goPanic("g_unichar_ismark: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gboolean g_unichar_iszerowidth(gunichar _0) {
	goPanic("g_unichar_iszerowidth: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gchar* g_match_info_expand_references(const GMatchInfo* v, const gchar* _0) {
	goPanic("g_match_info_expand_references: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gchar* g_match_info_fetch(const GMatchInfo* v, gint _0) {
	goPanic("g_match_info_fetch: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gchar* g_match_info_fetch_named(const GMatchInfo* v, const gchar* _0) {
	goPanic("g_match_info_fetch_named: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gchar** g_match_info_fetch_all(const GMatchInfo* v) {
	goPanic("g_match_info_fetch_all: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gchar** g_regex_split(const GRegex* v, const gchar* _0, GRegexMatchFlags _1) {
	goPanic("g_regex_split: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gchar** g_regex_split_simple(const gchar* _0, const gchar* _1, GRegexCompileFlags _2, GRegexMatchFlags _3) {
	goPanic("g_regex_split_simple: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gint g_match_info_get_match_count(const GMatchInfo* v) {
	goPanic("g_match_info_get_match_count: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gint g_regex_get_capture_count(const GRegex* v) {
	goPanic("g_regex_get_capture_count: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gint g_regex_get_max_backref(const GRegex* v) {
	goPanic("g_regex_get_max_backref: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gint g_regex_get_string_number(const GRegex* v, const gchar* _0) {
	goPanic("g_regex_get_string_number: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
gint g_unichar_combining_class(gunichar _0) {
	goPanic("g_unichar_combining_class: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
void g_queue_clear(GQueue* v) {
	goPanic("g_queue_clear: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 14))
void g_queue_init(GQueue* v) {
	goPanic("g_queue_init: library too old: needs at least 2.14");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
GChecksum* g_checksum_copy(const GChecksum* v) {
	goPanic("g_checksum_copy: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
GChecksum* g_checksum_new(GChecksumType _0) {
	goPanic("g_checksum_new: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
char* g_uri_escape_string(const char* _0, const char* _1, gboolean _2) {
	goPanic("g_uri_escape_string: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
char* g_uri_parse_scheme(const char* _0) {
	goPanic("g_uri_parse_scheme: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
char* g_uri_unescape_segment(const char* _0, const char* _1, const char* _2) {
	goPanic("g_uri_unescape_segment: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
char* g_uri_unescape_string(const char* _0, const char* _1) {
	goPanic("g_uri_unescape_string: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
const gchar* g_checksum_get_string(GChecksum* v) {
	goPanic("g_checksum_get_string: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
const gchar* g_dpgettext(const gchar* _0, const gchar* _1, gsize _2) {
	goPanic("g_dpgettext: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
gboolean g_hash_table_iter_next(GHashTableIter* v, gpointer* _0, gpointer* _1) {
	goPanic("g_hash_table_iter_next: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
gchar* g_compute_checksum_for_data(GChecksumType _0, const guchar* _1, gsize _2) {
	goPanic("g_compute_checksum_for_data: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
gchar* g_compute_checksum_for_string(GChecksumType _0, const gchar* _1, gssize _2) {
	goPanic("g_compute_checksum_for_string: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
gchar* g_format_size_for_display(goffset _0) {
	goPanic("g_format_size_for_display: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
gssize g_checksum_type_get_length(GChecksumType _0) {
	goPanic("g_checksum_type_get_length: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
void g_checksum_update(GChecksum* v, const guchar* _0, gssize _1) {
	goPanic("g_checksum_update: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
void g_hash_table_iter_init(GHashTableIter* v, GHashTable* _0) {
	goPanic("g_hash_table_iter_init: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
void g_hash_table_iter_remove(GHashTableIter* v) {
	goPanic("g_hash_table_iter_remove: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 16))
void g_hash_table_iter_steal(GHashTableIter* v) {
	goPanic("g_hash_table_iter_steal: library too old: needs at least 2.16");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 18))
const gchar* g_dgettext(const gchar* _0, const gchar* _1) {
	goPanic("g_dgettext: library too old: needs at least 2.18");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 18))
const gchar* g_dngettext(const gchar* _0, const gchar* _1, const gchar* _2, gulong _3) {
	goPanic("g_dngettext: library too old: needs at least 2.18");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 18))
const gchar* g_dpgettext2(const gchar* _0, const gchar* _1, const gchar* _2) {
	goPanic("g_dpgettext2: library too old: needs at least 2.18");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 18))
gpointer g_markup_parse_context_get_user_data(GMarkupParseContext* v) {
	goPanic("g_markup_parse_context_get_user_data: library too old: needs at least 2.18");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 18))
gpointer g_markup_parse_context_pop(GMarkupParseContext* v) {
	goPanic("g_markup_parse_context_pop: library too old: needs at least 2.18");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 18))
void g_checksum_reset(GChecksum* v) {
	goPanic("g_checksum_reset: library too old: needs at least 2.18");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 18))
void g_markup_parse_context_push(GMarkupParseContext* v, const GMarkupParser* _0, gpointer _1) {
	goPanic("g_markup_parse_context_push: library too old: needs at least 2.18");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 2))
const gchar* g_get_application_name(void) {
	goPanic("g_get_application_name: library too old: needs at least 2.2");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 2))
const gchar* g_markup_parse_context_get_element(GMarkupParseContext* v) {
	goPanic("g_markup_parse_context_get_element: library too old: needs at least 2.2");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 2))
gchar* g_utf8_strreverse(const gchar* _0, gssize _1) {
	goPanic("g_utf8_strreverse: library too old: needs at least 2.2");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 2))
void g_set_application_name(const gchar* _0) {
	goPanic("g_set_application_name: library too old: needs at least 2.2");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 20))
gint g_poll(GPollFD* _0, guint _1, gint _2) {
	goPanic("g_poll: library too old: needs at least 2.20");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
GMainContext* g_main_context_get_thread_default(void) {
	goPanic("g_main_context_get_thread_default: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
gboolean g_double_equal(gconstpointer _0, gconstpointer _1) {
	goPanic("g_double_equal: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
gboolean g_hostname_is_ascii_encoded(const gchar* _0) {
	goPanic("g_hostname_is_ascii_encoded: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
gboolean g_hostname_is_ip_address(const gchar* _0) {
	goPanic("g_hostname_is_ip_address: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
gboolean g_hostname_is_non_ascii(const gchar* _0) {
	goPanic("g_hostname_is_non_ascii: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
gboolean g_int64_equal(gconstpointer _0, gconstpointer _1) {
	goPanic("g_int64_equal: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
gchar* g_hostname_to_ascii(const gchar* _0) {
	goPanic("g_hostname_to_ascii: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
gchar* g_hostname_to_unicode(const gchar* _0) {
	goPanic("g_hostname_to_unicode: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
guint g_double_hash(gconstpointer _0) {
	goPanic("g_double_hash: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
guint g_int64_hash(gconstpointer _0) {
	goPanic("g_int64_hash: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
void g_main_context_pop_thread_default(GMainContext* v) {
	goPanic("g_main_context_pop_thread_default: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
void g_main_context_push_thread_default(GMainContext* v) {
	goPanic("g_main_context_push_thread_default: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 22))
void g_reload_user_special_dirs_cache(void) {
	goPanic("g_reload_user_special_dirs_cache: library too old: needs at least 2.22");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_builder_end(GVariantBuilder* v) {
	goPanic("g_variant_builder_end: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_byteswap(GVariant* v) {
	goPanic("g_variant_byteswap: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_get_child_value(GVariant* v, gsize _0) {
	goPanic("g_variant_get_child_value: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_get_maybe(GVariant* v) {
	goPanic("g_variant_get_maybe: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_get_normal_form(GVariant* v) {
	goPanic("g_variant_get_normal_form: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_get_variant(GVariant* v) {
	goPanic("g_variant_get_variant: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_array(const GVariantType* _0, GVariant* const* _1, gsize _2) {
	goPanic("g_variant_new_array: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_boolean(gboolean _0) {
	goPanic("g_variant_new_boolean: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_byte(guint8 _0) {
	goPanic("g_variant_new_byte: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_dict_entry(GVariant* _0, GVariant* _1) {
	goPanic("g_variant_new_dict_entry: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_double(gdouble _0) {
	goPanic("g_variant_new_double: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_handle(gint32 _0) {
	goPanic("g_variant_new_handle: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_int16(gint16 _0) {
	goPanic("g_variant_new_int16: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_int32(gint32 _0) {
	goPanic("g_variant_new_int32: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_int64(gint64 _0) {
	goPanic("g_variant_new_int64: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_maybe(const GVariantType* _0, GVariant* _1) {
	goPanic("g_variant_new_maybe: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_object_path(const gchar* _0) {
	goPanic("g_variant_new_object_path: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_signature(const gchar* _0) {
	goPanic("g_variant_new_signature: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_string(const gchar* _0) {
	goPanic("g_variant_new_string: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_strv(const gchar* const* _0, gssize _1) {
	goPanic("g_variant_new_strv: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_tuple(GVariant* const* _0, gsize _1) {
	goPanic("g_variant_new_tuple: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_uint16(guint16 _0) {
	goPanic("g_variant_new_uint16: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_uint32(guint32 _0) {
	goPanic("g_variant_new_uint32: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_uint64(guint64 _0) {
	goPanic("g_variant_new_uint64: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_new_variant(GVariant* _0) {
	goPanic("g_variant_new_variant: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariant* g_variant_ref_sink(GVariant* v) {
	goPanic("g_variant_ref_sink: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariantBuilder* g_variant_builder_new(const GVariantType* _0) {
	goPanic("g_variant_builder_new: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariantClass g_variant_classify(GVariant* v) {
	goPanic("g_variant_classify: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
GVariantType* g_variant_type_new(const gchar* _0) {
	goPanic("g_variant_type_new: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
const GVariantType* g_variant_get_type(GVariant* v) {
	goPanic("g_variant_get_type: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
const gchar* g_variant_get_string(GVariant* v, gsize* _0) {
	goPanic("g_variant_get_string: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
const gchar* g_variant_get_type_string(GVariant* v) {
	goPanic("g_variant_get_type_string: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
const gchar** g_variant_get_strv(GVariant* v, gsize* _0) {
	goPanic("g_variant_get_strv: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gboolean g_variant_equal(gconstpointer v, gconstpointer _0) {
	goPanic("g_variant_equal: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gboolean g_variant_get_boolean(GVariant* v) {
	goPanic("g_variant_get_boolean: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gboolean g_variant_is_container(GVariant* v) {
	goPanic("g_variant_is_container: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gboolean g_variant_is_normal_form(GVariant* v) {
	goPanic("g_variant_is_normal_form: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gboolean g_variant_is_object_path(const gchar* _0) {
	goPanic("g_variant_is_object_path: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gboolean g_variant_is_of_type(GVariant* v, const GVariantType* _0) {
	goPanic("g_variant_is_of_type: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gboolean g_variant_is_signature(const gchar* _0) {
	goPanic("g_variant_is_signature: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gboolean g_variant_type_string_scan(const gchar* _0, const gchar* _1, const gchar** _2) {
	goPanic("g_variant_type_string_scan: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gchar* g_variant_dup_string(GVariant* v, gsize* _0) {
	goPanic("g_variant_dup_string: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gchar* g_variant_print(GVariant* v, gboolean _0) {
	goPanic("g_variant_print: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gchar** g_variant_dup_strv(GVariant* v, gsize* _0) {
	goPanic("g_variant_dup_strv: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gconstpointer g_variant_get_data(GVariant* v) {
	goPanic("g_variant_get_data: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gdouble g_variant_get_double(GVariant* v) {
	goPanic("g_variant_get_double: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gint16 g_variant_get_int16(GVariant* v) {
	goPanic("g_variant_get_int16: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gint32 g_variant_get_handle(GVariant* v) {
	goPanic("g_variant_get_handle: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gint32 g_variant_get_int32(GVariant* v) {
	goPanic("g_variant_get_int32: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gint64 g_variant_get_int64(GVariant* v) {
	goPanic("g_variant_get_int64: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gsize g_variant_get_size(GVariant* v) {
	goPanic("g_variant_get_size: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
gsize g_variant_n_children(GVariant* v) {
	goPanic("g_variant_n_children: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
guint g_variant_hash(gconstpointer v) {
	goPanic("g_variant_hash: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
guint16 g_variant_get_uint16(GVariant* v) {
	goPanic("g_variant_get_uint16: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
guint32 g_variant_get_uint32(GVariant* v) {
	goPanic("g_variant_get_uint32: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
guint64 g_variant_get_uint64(GVariant* v) {
	goPanic("g_variant_get_uint64: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
guint8 g_variant_get_byte(GVariant* v) {
	goPanic("g_variant_get_byte: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
void g_variant_builder_add_value(GVariantBuilder* v, GVariant* _0) {
	goPanic("g_variant_builder_add_value: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
void g_variant_builder_close(GVariantBuilder* v) {
	goPanic("g_variant_builder_close: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
void g_variant_builder_open(GVariantBuilder* v, const GVariantType* _0) {
	goPanic("g_variant_builder_open: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 24))
void g_variant_store(GVariant* v, gpointer _0) {
	goPanic("g_variant_store: library too old: needs at least 2.24");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
GRegexCompileFlags g_regex_get_compile_flags(const GRegex* v) {
	goPanic("g_regex_get_compile_flags: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
GRegexMatchFlags g_regex_get_match_flags(const GRegex* v) {
	goPanic("g_regex_get_match_flags: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
GTimeZone* g_time_zone_new(const gchar* _0) {
	goPanic("g_time_zone_new: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
GTimeZone* g_time_zone_new_local(void) {
	goPanic("g_time_zone_new_local: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
GTimeZone* g_time_zone_new_utc(void) {
	goPanic("g_time_zone_new_utc: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
GVariant* g_variant_new_bytestring(const gchar* _0) {
	goPanic("g_variant_new_bytestring: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
GVariant* g_variant_new_bytestring_array(const gchar* const* _0, gssize _1) {
	goPanic("g_variant_new_bytestring_array: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
const char* g_source_get_name(GSource* v) {
	goPanic("g_source_get_name: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
const gchar* g_dcgettext(const gchar* _0, const gchar* _1, gint _2) {
	goPanic("g_dcgettext: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
const gchar* g_time_zone_get_abbreviation(GTimeZone* v, gint _0) {
	goPanic("g_time_zone_get_abbreviation: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
const gchar* g_variant_get_bytestring(GVariant* v) {
	goPanic("g_variant_get_bytestring: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
const gchar** g_variant_get_bytestring_array(GVariant* v, gsize* _0) {
	goPanic("g_variant_get_bytestring_array: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
gboolean g_time_zone_is_dst(GTimeZone* v, gint _0) {
	goPanic("g_time_zone_is_dst: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
gboolean g_variant_is_floating(GVariant* v) {
	goPanic("g_variant_is_floating: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
gchar* g_variant_dup_bytestring(GVariant* v, gsize* _0) {
	goPanic("g_variant_dup_bytestring: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
gchar** g_variant_dup_bytestring_array(GVariant* v, gsize* _0) {
	goPanic("g_variant_dup_bytestring_array: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
gint g_time_zone_adjust_time(GTimeZone* v, GTimeType _0, gint64* _1) {
	goPanic("g_time_zone_adjust_time: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
gint g_time_zone_find_interval(GTimeZone* v, GTimeType _0, gint64 _1) {
	goPanic("g_time_zone_find_interval: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
gint g_variant_compare(gconstpointer v, gconstpointer _0) {
	goPanic("g_variant_compare: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
gint32 g_time_zone_get_offset(GTimeZone* v, gint _0) {
	goPanic("g_time_zone_get_offset: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
gint64 g_key_file_get_int64(GKeyFile* v, const gchar* _0, const gchar* _1) {
	goPanic("g_key_file_get_int64: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
guint64 g_key_file_get_uint64(GKeyFile* v, const gchar* _0, const gchar* _1) {
	goPanic("g_key_file_get_uint64: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
void g_key_file_set_int64(GKeyFile* v, const gchar* _0, const gchar* _1, gint64 _2) {
	goPanic("g_key_file_set_int64: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
void g_key_file_set_uint64(GKeyFile* v, const gchar* _0, const gchar* _1, guint64 _2) {
	goPanic("g_key_file_set_uint64: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
void g_source_set_name(GSource* v, const char* _0) {
	goPanic("g_source_set_name: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 26))
void g_source_set_name_by_id(guint _0, const char* _1) {
	goPanic("g_source_set_name_by_id: library too old: needs at least 2.26");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 28))
GVariant* g_variant_lookup_value(GVariant* v, const gchar* _0, const GVariantType* _1) {
	goPanic("g_variant_lookup_value: library too old: needs at least 2.28");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 28))
const gchar* g_get_user_runtime_dir(void) {
	goPanic("g_get_user_runtime_dir: library too old: needs at least 2.28");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 28))
gchar** g_get_environ(void) {
	goPanic("g_get_environ: library too old: needs at least 2.28");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 28))
gchar** g_get_locale_variants(const gchar* _0) {
	goPanic("g_get_locale_variants: library too old: needs at least 2.28");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 28))
gint64 g_get_monotonic_time(void) {
	goPanic("g_get_monotonic_time: library too old: needs at least 2.28");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 28))
gint64 g_get_real_time(void) {
	goPanic("g_get_real_time: library too old: needs at least 2.28");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 28))
gint64 g_source_get_time(GSource* v) {
	goPanic("g_source_get_time: library too old: needs at least 2.28");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 28))
void g_source_add_child_source(GSource* v, GSource* _0) {
	goPanic("g_source_add_child_source: library too old: needs at least 2.28");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 28))
void g_source_remove_child_source(GSource* v, GSource* _0) {
	goPanic("g_source_remove_child_source: library too old: needs at least 2.28");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
GUnicodeScript g_unicode_script_from_iso15924(guint32 _0) {
	goPanic("g_unicode_script_from_iso15924: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
GVariant* g_variant_new_objv(const gchar* const* _0, gssize _1) {
	goPanic("g_variant_new_objv: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
const gchar** g_variant_get_objv(GVariant* v, gsize* _0) {
	goPanic("g_variant_get_objv: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
gboolean g_unichar_compose(gunichar _0, gunichar _1, gunichar* _2) {
	goPanic("g_unichar_compose: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
gboolean g_unichar_decompose(gunichar _0, gunichar* _1, gunichar* _2) {
	goPanic("g_unichar_decompose: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
gchar* g_compute_hmac_for_data(GChecksumType _0, const guchar* _1, gsize _2, const guchar* _3, gsize _4) {
	goPanic("g_compute_hmac_for_data: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
gchar* g_compute_hmac_for_string(GChecksumType _0, const guchar* _1, gsize _2, const gchar* _3, gssize _4) {
	goPanic("g_compute_hmac_for_string: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
gchar* g_format_size(guint64 _0) {
	goPanic("g_format_size: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
gchar* g_format_size_full(guint64 _0, GFormatSizeFlags _1) {
	goPanic("g_format_size_full: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
gchar* g_regex_escape_nul(const gchar* _0, gint _1) {
	goPanic("g_regex_escape_nul: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
gchar* g_utf8_substring(const gchar* _0, glong _1, glong _2) {
	goPanic("g_utf8_substring: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
gchar** g_variant_dup_objv(GVariant* v, gsize* _0) {
	goPanic("g_variant_dup_objv: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
gsize g_unichar_fully_decompose(gunichar _0, gboolean _1, gunichar* _2, gsize _3) {
	goPanic("g_unichar_fully_decompose: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
guint32 g_unicode_script_to_iso15924(GUnicodeScript _0) {
	goPanic("g_unicode_script_to_iso15924: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 30))
void g_hash_table_iter_replace(GHashTableIter* v, gpointer _0) {
	goPanic("g_hash_table_iter_replace: library too old: needs at least 2.30");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
GBytes* g_bytes_new(gconstpointer _0, gsize _1) {
	goPanic("g_bytes_new: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
GBytes* g_bytes_new_from_bytes(GBytes* v, gsize _0, gsize _1) {
	goPanic("g_bytes_new_from_bytes: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
GMainContext* g_main_context_ref_thread_default(void) {
	goPanic("g_main_context_ref_thread_default: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
GMappedFile* g_mapped_file_new_from_fd(gint _0, gboolean _1) {
	goPanic("g_mapped_file_new_from_fd: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
GVariant* g_variant_new_fixed_array(const GVariantType* _0, gconstpointer _1, gsize _2, gsize _3) {
	goPanic("g_variant_new_fixed_array: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
const gchar* g_environ_getenv(gchar** _0, const gchar* _1) {
	goPanic("g_environ_getenv: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
gboolean g_bytes_equal(gconstpointer v, gconstpointer _0) {
	goPanic("g_bytes_equal: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
gboolean g_hash_table_add(GHashTable* _0, gpointer _1) {
	goPanic("g_hash_table_add: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
gboolean g_hash_table_contains(GHashTable* _0, gconstpointer _1) {
	goPanic("g_hash_table_contains: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
gchar** g_environ_setenv(gchar** _0, const gchar* _1, const gchar* _2, gboolean _3) {
	goPanic("g_environ_setenv: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
gchar** g_environ_unsetenv(gchar** _0, const gchar* _1) {
	goPanic("g_environ_unsetenv: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
gconstpointer g_bytes_get_data(GBytes* v, gsize* _0) {
	goPanic("g_bytes_get_data: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
gint g_bytes_compare(gconstpointer v, gconstpointer _0) {
	goPanic("g_bytes_compare: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
gsize g_bytes_get_size(GBytes* v) {
	goPanic("g_bytes_get_size: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 32))
guint g_bytes_hash(gconstpointer v) {
	goPanic("g_bytes_hash: library too old: needs at least 2.32");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 34))
GBytes* g_mapped_file_get_bytes(GMappedFile* v) {
	goPanic("g_mapped_file_get_bytes: library too old: needs at least 2.34");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 34))
gboolean g_regex_get_has_cr_or_lf(const GRegex* v) {
	goPanic("g_regex_get_has_cr_or_lf: library too old: needs at least 2.34");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 34))
gboolean g_spawn_check_exit_status(gint _0) {
	goPanic("g_spawn_check_exit_status: library too old: needs at least 2.34");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 34))
gboolean g_variant_check_format_string(GVariant* v, const gchar* _0, gboolean _1) {
	goPanic("g_variant_check_format_string: library too old: needs at least 2.34");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 34))
gchar* g_compute_checksum_for_bytes(GChecksumType _0, GBytes* _1) {
	goPanic("g_compute_checksum_for_bytes: library too old: needs at least 2.34");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 36))
GBytes* g_variant_get_data_as_bytes(GVariant* v) {
	goPanic("g_variant_get_data_as_bytes: library too old: needs at least 2.36");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 36))
GIOCondition g_source_query_unix_fd(GSource* v, gpointer _0) {
	goPanic("g_source_query_unix_fd: library too old: needs at least 2.36");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 36))
GVariant* g_variant_new_from_bytes(const GVariantType* _0, GBytes* _1, gboolean _2) {
	goPanic("g_variant_new_from_bytes: library too old: needs at least 2.36");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 36))
gpointer g_source_add_unix_fd(GSource* v, gint _0, GIOCondition _1) {
	goPanic("g_source_add_unix_fd: library too old: needs at least 2.36");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 36))
void g_source_modify_unix_fd(GSource* v, gpointer _0, GIOCondition _1) {
	goPanic("g_source_modify_unix_fd: library too old: needs at least 2.36");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 36))
void g_source_remove_unix_fd(GSource* v, gpointer _0) {
	goPanic("g_source_remove_unix_fd: library too old: needs at least 2.36");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 36))
void g_source_set_ready_time(GSource* v, gint64 _0) {
	goPanic("g_source_set_ready_time: library too old: needs at least 2.36");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 38))
gint g_regex_get_max_lookbehind(const GRegex* v) {
	goPanic("g_regex_get_max_lookbehind: library too old: needs at least 2.38");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
GSource* g_child_watch_source_new(GPid _0) {
	goPanic("g_child_watch_source_new: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
const gchar* g_strip_context(const gchar* _0, const gchar* _1) {
	goPanic("g_strip_context: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
gboolean g_queue_remove(GQueue* v, gconstpointer _0) {
	goPanic("g_queue_remove: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
gboolean g_setenv(const gchar* _0, const gchar* _1, gboolean _2) {
	goPanic("g_setenv: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
gboolean g_unichar_get_mirror_char(gunichar _0, gunichar* _1) {
	goPanic("g_unichar_get_mirror_char: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
gchar* g_file_read_link(const gchar* _0) {
	goPanic("g_file_read_link: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
gint g_queue_index(GQueue* v, gconstpointer _0) {
	goPanic("g_queue_index: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
gpointer g_queue_peek_nth(GQueue* v, guint _0) {
	goPanic("g_queue_peek_nth: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
gpointer g_queue_pop_nth(GQueue* v, guint _0) {
	goPanic("g_queue_pop_nth: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
guint g_queue_get_length(GQueue* v) {
	goPanic("g_queue_get_length: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
guint g_queue_remove_all(GQueue* v, gconstpointer _0) {
	goPanic("g_queue_remove_all: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
void g_queue_push_nth(GQueue* v, gpointer _0, gint _1) {
	goPanic("g_queue_push_nth: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
void g_queue_reverse(GQueue* v) {
	goPanic("g_queue_reverse: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 4))
void g_unsetenv(const gchar* _0) {
	goPanic("g_unsetenv: library too old: needs at least 2.4");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 40))
GVariant* g_variant_dict_end(GVariantDict* v) {
	goPanic("g_variant_dict_end: library too old: needs at least 2.40");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 40))
GVariant* g_variant_dict_lookup_value(GVariantDict* v, const gchar* _0, const GVariantType* _1) {
	goPanic("g_variant_dict_lookup_value: library too old: needs at least 2.40");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 40))
GVariantDict* g_variant_dict_new(GVariant* _0) {
	goPanic("g_variant_dict_new: library too old: needs at least 2.40");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 40))
gboolean g_key_file_save_to_file(GKeyFile* v, const gchar* _0) {
	goPanic("g_key_file_save_to_file: library too old: needs at least 2.40");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 40))
gboolean g_variant_dict_contains(GVariantDict* v, const gchar* _0) {
	goPanic("g_variant_dict_contains: library too old: needs at least 2.40");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 40))
gboolean g_variant_dict_remove(GVariantDict* v, const gchar* _0) {
	goPanic("g_variant_dict_remove: library too old: needs at least 2.40");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 40))
gchar* g_variant_parse_error_print_context(GError* _0, const gchar* _1) {
	goPanic("g_variant_parse_error_print_context: library too old: needs at least 2.40");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 40))
void g_variant_dict_clear(GVariantDict* v) {
	goPanic("g_variant_dict_clear: library too old: needs at least 2.40");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 40))
void g_variant_dict_insert_value(GVariantDict* v, const gchar* _0, GVariant* _1) {
	goPanic("g_variant_dict_insert_value: library too old: needs at least 2.40");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 50))
GLogWriterOutput g_log_writer_default(GLogLevelFlags _0, const GLogField* _1, gsize _2, gpointer _3) {
	goPanic("g_log_writer_default: library too old: needs at least 2.50");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 50))
GLogWriterOutput g_log_writer_journald(GLogLevelFlags _0, const GLogField* _1, gsize _2, gpointer _3) {
	goPanic("g_log_writer_journald: library too old: needs at least 2.50");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 50))
GLogWriterOutput g_log_writer_standard_streams(GLogLevelFlags _0, const GLogField* _1, gsize _2, gpointer _3) {
	goPanic("g_log_writer_standard_streams: library too old: needs at least 2.50");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 50))
gboolean g_key_file_load_from_bytes(GKeyFile* v, GBytes* _0, GKeyFileFlags _1) {
	goPanic("g_key_file_load_from_bytes: library too old: needs at least 2.50");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 50))
gboolean g_log_writer_is_journald(gint _0) {
	goPanic("g_log_writer_is_journald: library too old: needs at least 2.50");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 50))
gboolean g_log_writer_supports_color(gint _0) {
	goPanic("g_log_writer_supports_color: library too old: needs at least 2.50");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 50))
gchar* g_compute_hmac_for_bytes(GChecksumType _0, GBytes* _1, GBytes* _2) {
	goPanic("g_compute_hmac_for_bytes: library too old: needs at least 2.50");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 50))
gchar* g_log_writer_format_fields(GLogLevelFlags _0, const GLogField* _1, gsize _2, gboolean _3) {
	goPanic("g_log_writer_format_fields: library too old: needs at least 2.50");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 50))
void g_log_structured_array(GLogLevelFlags _0, const GLogField* _1, gsize _2) {
	goPanic("g_log_structured_array: library too old: needs at least 2.50");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 50))
void g_log_variant(const gchar* _0, GLogLevelFlags _1, GVariant* _2) {
	goPanic("g_log_variant: library too old: needs at least 2.50");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 52))
gboolean g_uuid_string_is_valid(const gchar* _0) {
	goPanic("g_uuid_string_is_valid: library too old: needs at least 2.52");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 52))
gchar* g_utf8_make_valid(const gchar* _0, gssize _1) {
	goPanic("g_utf8_make_valid: library too old: needs at least 2.52");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 52))
gchar* g_uuid_string_random(void) {
	goPanic("g_uuid_string_random: library too old: needs at least 2.52");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 56))
gchar* g_key_file_get_locale_for_key(GKeyFile* v, const gchar* _0, const gchar* _1, const gchar* _2) {
	goPanic("g_key_file_get_locale_for_key: library too old: needs at least 2.56");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 58))
GTimeZone* g_time_zone_new_offset(gint32 _0) {
	goPanic("g_time_zone_new_offset: library too old: needs at least 2.58");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 58))
const gchar* const* g_get_language_names_with_category(const gchar* _0) {
	goPanic("g_get_language_names_with_category: library too old: needs at least 2.58");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 58))
const gchar* g_time_zone_get_identifier(GTimeZone* v) {
	goPanic("g_time_zone_get_identifier: library too old: needs at least 2.58");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 58))
gboolean g_hash_table_steal_extended(GHashTable* _0, gconstpointer _1, gpointer* _2, gpointer* _3) {
	goPanic("g_hash_table_steal_extended: library too old: needs at least 2.58");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 58))
gboolean g_spawn_async_with_fds(const gchar* _0, gchar** _1, gchar** _2, GSpawnFlags _3, GSpawnChildSetupFunc _4, gpointer _5, GPid* _6, gint _7, gint _8, gint _9) {
	goPanic("g_spawn_async_with_fds: library too old: needs at least 2.58");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 58))
gchar* g_canonicalize_filename(const gchar* _0, const gchar* _1) {
	goPanic("g_canonicalize_filename: library too old: needs at least 2.58");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
GKeyFile* g_key_file_new(void) {
	goPanic("g_key_file_new: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
const gchar* const* g_get_language_names(void) {
	goPanic("g_get_language_names: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
const gchar* const* g_get_system_config_dirs(void) {
	goPanic("g_get_system_config_dirs: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
const gchar* const* g_get_system_data_dirs(void) {
	goPanic("g_get_system_data_dirs: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
const gchar* g_get_user_cache_dir(void) {
	goPanic("g_get_user_cache_dir: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
const gchar* g_get_user_config_dir(void) {
	goPanic("g_get_user_config_dir: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
const gchar* g_get_user_data_dir(void) {
	goPanic("g_get_user_data_dir: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
const gchar* glib_check_version(guint _0, guint _1, guint _2) {
	goPanic("glib_check_version: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean g_get_filename_charsets(const gchar*** _0) {
	goPanic("g_get_filename_charsets: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean g_key_file_get_boolean(GKeyFile* v, const gchar* _0, const gchar* _1) {
	goPanic("g_key_file_get_boolean: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean g_key_file_has_group(GKeyFile* v, const gchar* _0) {
	goPanic("g_key_file_has_group: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean g_key_file_load_from_data(GKeyFile* v, const gchar* _0, gsize _1, GKeyFileFlags _2) {
	goPanic("g_key_file_load_from_data: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean g_key_file_load_from_data_dirs(GKeyFile* v, const gchar* _0, gchar** _1, GKeyFileFlags _2) {
	goPanic("g_key_file_load_from_data_dirs: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean g_key_file_load_from_file(GKeyFile* v, const gchar* _0, GKeyFileFlags _1) {
	goPanic("g_key_file_load_from_file: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean g_key_file_remove_comment(GKeyFile* v, const gchar* _0, const gchar* _1) {
	goPanic("g_key_file_remove_comment: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean g_key_file_remove_group(GKeyFile* v, const gchar* _0) {
	goPanic("g_key_file_remove_group: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean g_key_file_remove_key(GKeyFile* v, const gchar* _0, const gchar* _1) {
	goPanic("g_key_file_remove_key: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean g_key_file_set_comment(GKeyFile* v, const gchar* _0, const gchar* _1, const gchar* _2) {
	goPanic("g_key_file_set_comment: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gboolean* g_key_file_get_boolean_list(GKeyFile* v, const gchar* _0, const gchar* _1, gsize* _2) {
	goPanic("g_key_file_get_boolean_list: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar* g_filename_display_basename(const gchar* _0) {
	goPanic("g_filename_display_basename: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar* g_filename_display_name(const gchar* _0) {
	goPanic("g_filename_display_name: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar* g_key_file_get_comment(GKeyFile* v, const gchar* _0, const gchar* _1) {
	goPanic("g_key_file_get_comment: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar* g_key_file_get_locale_string(GKeyFile* v, const gchar* _0, const gchar* _1, const gchar* _2) {
	goPanic("g_key_file_get_locale_string: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar* g_key_file_get_start_group(GKeyFile* v) {
	goPanic("g_key_file_get_start_group: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar* g_key_file_get_string(GKeyFile* v, const gchar* _0, const gchar* _1) {
	goPanic("g_key_file_get_string: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar* g_key_file_get_value(GKeyFile* v, const gchar* _0, const gchar* _1) {
	goPanic("g_key_file_get_value: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar* g_key_file_to_data(GKeyFile* v, gsize* _0) {
	goPanic("g_key_file_to_data: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar** g_key_file_get_groups(GKeyFile* v, gsize* _0) {
	goPanic("g_key_file_get_groups: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar** g_key_file_get_keys(GKeyFile* v, const gchar* _0, gsize* _1) {
	goPanic("g_key_file_get_keys: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar** g_key_file_get_locale_string_list(GKeyFile* v, const gchar* _0, const gchar* _1, const gchar* _2, gsize* _3) {
	goPanic("g_key_file_get_locale_string_list: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar** g_key_file_get_string_list(GKeyFile* v, const gchar* _0, const gchar* _1, gsize* _2) {
	goPanic("g_key_file_get_string_list: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gchar** g_uri_list_extract_uris(const gchar* _0) {
	goPanic("g_uri_list_extract_uris: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gint g_key_file_get_integer(GKeyFile* v, const gchar* _0, const gchar* _1) {
	goPanic("g_key_file_get_integer: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
gint* g_key_file_get_integer_list(GKeyFile* v, const gchar* _0, const gchar* _1, gsize* _2) {
	goPanic("g_key_file_get_integer_list: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_key_file_set_boolean(GKeyFile* v, const gchar* _0, const gchar* _1, gboolean _2) {
	goPanic("g_key_file_set_boolean: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_key_file_set_boolean_list(GKeyFile* v, const gchar* _0, const gchar* _1, gboolean* _2, gsize _3) {
	goPanic("g_key_file_set_boolean_list: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_key_file_set_integer(GKeyFile* v, const gchar* _0, const gchar* _1, gint _2) {
	goPanic("g_key_file_set_integer: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_key_file_set_integer_list(GKeyFile* v, const gchar* _0, const gchar* _1, gint* _2, gsize _3) {
	goPanic("g_key_file_set_integer_list: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_key_file_set_list_separator(GKeyFile* v, gchar _0) {
	goPanic("g_key_file_set_list_separator: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_key_file_set_locale_string(GKeyFile* v, const gchar* _0, const gchar* _1, const gchar* _2, const gchar* _3) {
	goPanic("g_key_file_set_locale_string: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_key_file_set_locale_string_list(GKeyFile* v, const gchar* _0, const gchar* _1, const gchar* _2, const gchar* const* _3, gsize _4) {
	goPanic("g_key_file_set_locale_string_list: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_key_file_set_string(GKeyFile* v, const gchar* _0, const gchar* _1, const gchar* _2) {
	goPanic("g_key_file_set_string: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_key_file_set_string_list(GKeyFile* v, const gchar* _0, const gchar* _1, const gchar* const* _2, gsize _3) {
	goPanic("g_key_file_set_string_list: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_key_file_set_value(GKeyFile* v, const gchar* _0, const gchar* _1, const gchar* _2) {
	goPanic("g_key_file_set_value: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_option_group_add_entries(GOptionGroup* v, const GOptionEntry* _0) {
	goPanic("g_option_group_add_entries: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 6))
void g_option_group_set_translation_domain(GOptionGroup* v, const gchar* _0) {
	goPanic("g_option_group_set_translation_domain: library too old: needs at least 2.6");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 60))
gboolean g_utf8_validate_len(const gchar* _0, gsize _1, const gchar** _2) {
	goPanic("g_utf8_validate_len: library too old: needs at least 2.60");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 62))
gboolean g_get_console_charset(const char** _0) {
	goPanic("g_get_console_charset: library too old: needs at least 2.62");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 64))
gchar* g_get_os_info(const gchar* _0) {
	goPanic("g_get_os_info: library too old: needs at least 2.64");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
GBytes* g_uri_unescape_bytes(const char* _0, gssize _1, const char* _2) {
	goPanic("g_uri_unescape_bytes: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
GHashTable* g_uri_parse_params(const gchar* _0, gssize _1, const gchar* _2, GUriParamsFlags _3) {
	goPanic("g_uri_parse_params: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
GUri* g_uri_build(GUriFlags _0, const gchar* _1, const gchar* _2, const gchar* _3, gint _4, const gchar* _5, const gchar* _6, const gchar* _7) {
	goPanic("g_uri_build: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
GUri* g_uri_build_with_user(GUriFlags _0, const gchar* _1, const gchar* _2, const gchar* _3, const gchar* _4, const gchar* _5, gint _6, const gchar* _7, const gchar* _8, const gchar* _9) {
	goPanic("g_uri_build_with_user: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
GUri* g_uri_parse(const gchar* _0, GUriFlags _1) {
	goPanic("g_uri_parse: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
GUri* g_uri_parse_relative(GUri* v, const gchar* _0, GUriFlags _1) {
	goPanic("g_uri_parse_relative: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
GUriFlags g_uri_get_flags(GUri* v) {
	goPanic("g_uri_get_flags: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
char* g_uri_escape_bytes(const guint8* _0, gsize _1, const char* _2) {
	goPanic("g_uri_escape_bytes: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
char* g_uri_to_string(GUri* v) {
	goPanic("g_uri_to_string: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
char* g_uri_to_string_partial(GUri* v, GUriHideFlags _0) {
	goPanic("g_uri_to_string_partial: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
const char* g_uri_peek_scheme(const char* _0) {
	goPanic("g_uri_peek_scheme: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
const gchar* g_uri_get_auth_params(GUri* v) {
	goPanic("g_uri_get_auth_params: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
const gchar* g_uri_get_fragment(GUri* v) {
	goPanic("g_uri_get_fragment: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
const gchar* g_uri_get_host(GUri* v) {
	goPanic("g_uri_get_host: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
const gchar* g_uri_get_password(GUri* v) {
	goPanic("g_uri_get_password: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
const gchar* g_uri_get_path(GUri* v) {
	goPanic("g_uri_get_path: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
const gchar* g_uri_get_query(GUri* v) {
	goPanic("g_uri_get_query: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
const gchar* g_uri_get_scheme(GUri* v) {
	goPanic("g_uri_get_scheme: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
const gchar* g_uri_get_user(GUri* v) {
	goPanic("g_uri_get_user: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
const gchar* g_uri_get_userinfo(GUri* v) {
	goPanic("g_uri_get_userinfo: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
gboolean g_file_set_contents_full(const gchar* _0, const gchar* _1, gssize _2, GFileSetContentsFlags _3, int _4) {
	goPanic("g_file_set_contents_full: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
gboolean g_uri_is_valid(const gchar* _0, GUriFlags _1) {
	goPanic("g_uri_is_valid: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
gboolean g_uri_params_iter_next(GUriParamsIter* v, gchar** _0, gchar** _1) {
	goPanic("g_uri_params_iter_next: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
gboolean g_uri_split(const gchar* _0, GUriFlags _1, gchar** _2, gchar** _3, gchar** _4, gint* _5, gchar** _6, gchar** _7, gchar** _8) {
	goPanic("g_uri_split: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
gboolean g_uri_split_network(const gchar* _0, GUriFlags _1, gchar** _2, gchar** _3, gint* _4) {
	goPanic("g_uri_split_network: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
gboolean g_uri_split_with_user(const gchar* _0, GUriFlags _1, gchar** _2, gchar** _3, gchar** _4, gchar** _5, gchar** _6, gint* _7, gchar** _8, gchar** _9, gchar** _10) {
	goPanic("g_uri_split_with_user: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
gchar* g_uri_join(GUriFlags _0, const gchar* _1, const gchar* _2, const gchar* _3, gint _4, const gchar* _5, const gchar* _6, const gchar* _7) {
	goPanic("g_uri_join: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
gchar* g_uri_join_with_user(GUriFlags _0, const gchar* _1, const gchar* _2, const gchar* _3, const gchar* _4, const gchar* _5, gint _6, const gchar* _7, const gchar* _8, const gchar* _9) {
	goPanic("g_uri_join_with_user: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
gchar* g_uri_resolve_relative(const gchar* _0, const gchar* _1, GUriFlags _2) {
	goPanic("g_uri_resolve_relative: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
gint g_uri_get_port(GUri* v) {
	goPanic("g_uri_get_port: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 66))
void g_uri_params_iter_init(GUriParamsIter* v, const gchar* _0, gssize _1, const gchar* _2, GUriParamsFlags _3) {
	goPanic("g_uri_params_iter_init: library too old: needs at least 2.66");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 68))
GTimeZone* g_time_zone_new_identifier(const gchar* _0) {
	goPanic("g_time_zone_new_identifier: library too old: needs at least 2.68");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 68))
gboolean g_log_writer_default_would_drop(GLogLevelFlags _0, const char* _1) {
	goPanic("g_log_writer_default_would_drop: library too old: needs at least 2.68");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 68))
gboolean g_spawn_async_with_pipes_and_fds(const gchar* _0, const gchar* const* _1, const gchar* const* _2, GSpawnFlags _3, GSpawnChildSetupFunc _4, gpointer _5, gint _6, gint _7, gint _8, const gint* _9, const gint* _10, gsize _11, GPid* _12, gint* _13, gint* _14, gint* _15) {
	goPanic("g_spawn_async_with_pipes_and_fds: library too old: needs at least 2.68");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 68))
void g_log_writer_default_set_use_stderr(gboolean _0) {
	goPanic("g_log_writer_default_set_use_stderr: library too old: needs at least 2.68");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 8))
GMappedFile* g_mapped_file_new(const gchar* _0, gboolean _1) {
	goPanic("g_mapped_file_new: library too old: needs at least 2.8");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 8))
const gchar* g_get_host_name(void) {
	goPanic("g_get_host_name: library too old: needs at least 2.8");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 8))
gboolean g_file_set_contents(const gchar* _0, const gchar* _1, gssize _2) {
	goPanic("g_file_set_contents: library too old: needs at least 2.8");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 8))
gchar* g_build_filenamev(gchar** _0) {
	goPanic("g_build_filenamev: library too old: needs at least 2.8");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 8))
gchar* g_build_pathv(const gchar* _0, gchar** _1) {
	goPanic("g_build_pathv: library too old: needs at least 2.8");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 8))
gchar* g_mapped_file_get_contents(GMappedFile* v) {
	goPanic("g_mapped_file_get_contents: library too old: needs at least 2.8");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 8))
gchar* g_utf8_collate_key_for_filename(const gchar* _0, gssize _1) {
	goPanic("g_utf8_collate_key_for_filename: library too old: needs at least 2.8");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 8))
gchar** g_listenv(void) {
	goPanic("g_listenv: library too old: needs at least 2.8");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 8))
gint g_mkdir_with_parents(const gchar* _0, gint _1) {
	goPanic("g_mkdir_with_parents: library too old: needs at least 2.8");
}
#endif

#if (GLIB_MAJOR_VERSION < 2 || (GLIB_MAJOR_VERSION == 2 && GLIB_MINOR_VERSION < 8))
gsize g_mapped_file_get_length(GMappedFile* v) {
	goPanic("g_mapped_file_get_length: library too old: needs at least 2.8");
}
#endif
