#include <gdk/gdk.h>
#include <glib-object.h>

extern void goPanic(const char*);

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
GList* gdk_screen_get_window_stack(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_window_stack: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
GdkWindow* gdk_screen_get_active_window(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_active_window: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
GdkWindowTypeHint gdk_window_get_type_hint(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_type_hint: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
const cairo_font_options_t* gdk_screen_get_font_options(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_font_options: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
gboolean gdk_display_supports_input_shapes(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_supports_input_shapes: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
gboolean gdk_display_supports_shapes(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_supports_shapes: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
gboolean gdk_screen_is_composited(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_is_composited: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
gdouble gdk_screen_get_resolution(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_resolution: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
void gdk_screen_set_font_options(GdkScreen* v, const cairo_font_options_t* _0) __attribute__((weak)) {
	goPanic("gdk_screen_set_font_options: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
void gdk_screen_set_resolution(GdkScreen* v, gdouble _0) __attribute__((weak)) {
	goPanic("gdk_screen_set_resolution: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
void gdk_window_input_shape_combine_region(GdkWindow* v, const cairo_region_t* _0, gint _1, gint _2) __attribute__((weak)) {
	goPanic("gdk_window_input_shape_combine_region: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
void gdk_window_merge_child_input_shapes(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_merge_child_input_shapes: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 10))
void gdk_window_set_child_input_shapes(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_set_child_input_shapes: library too old: needs at least 2.10");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 12))
gboolean gdk_display_supports_composite(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_supports_composite: library too old: needs at least 2.12");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 12))
gboolean gdk_keymap_have_bidi_layouts(GdkKeymap* v) __attribute__((weak)) {
	goPanic("gdk_keymap_have_bidi_layouts: library too old: needs at least 2.12");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 12))
gchar* gdk_color_to_string(const GdkColor* v) __attribute__((weak)) {
	goPanic("gdk_color_to_string: library too old: needs at least 2.12");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 12))
guint gdk_threads_add_idle_full(gint _0, GSourceFunc _1, gpointer _2, GDestroyNotify _3) __attribute__((weak)) {
	goPanic("gdk_threads_add_idle_full: library too old: needs at least 2.12");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 12))
guint gdk_threads_add_timeout_full(gint _0, guint _1, GSourceFunc _2, gpointer _3, GDestroyNotify _4) __attribute__((weak)) {
	goPanic("gdk_threads_add_timeout_full: library too old: needs at least 2.12");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 12))
void gdk_notify_startup_complete_with_id(const gchar* _0) __attribute__((weak)) {
	goPanic("gdk_notify_startup_complete_with_id: library too old: needs at least 2.12");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 12))
void gdk_window_beep(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_beep: library too old: needs at least 2.12");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 12))
void gdk_window_set_composited(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_composited: library too old: needs at least 2.12");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 12))
void gdk_window_set_opacity(GdkWindow* v, gdouble _0) __attribute__((weak)) {
	goPanic("gdk_window_set_opacity: library too old: needs at least 2.12");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 12))
void gdk_window_set_startup_id(GdkWindow* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gdk_window_set_startup_id: library too old: needs at least 2.12");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
GdkAppLaunchContext* gdk_app_launch_context_new(void) __attribute__((weak)) {
	goPanic("gdk_app_launch_context_new: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
gboolean gdk_test_simulate_button(GdkWindow* _0, gint _1, gint _2, guint _3, GdkModifierType _4, GdkEventType _5) __attribute__((weak)) {
	goPanic("gdk_test_simulate_button: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
gboolean gdk_test_simulate_key(GdkWindow* _0, gint _1, gint _2, guint _3, GdkModifierType _4, GdkEventType _5) __attribute__((weak)) {
	goPanic("gdk_test_simulate_key: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
gchar* gdk_screen_get_monitor_plug_name(GdkScreen* v, gint _0) __attribute__((weak)) {
	goPanic("gdk_screen_get_monitor_plug_name: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
gint gdk_screen_get_monitor_height_mm(GdkScreen* v, gint _0) __attribute__((weak)) {
	goPanic("gdk_screen_get_monitor_height_mm: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
gint gdk_screen_get_monitor_width_mm(GdkScreen* v, gint _0) __attribute__((weak)) {
	goPanic("gdk_screen_get_monitor_width_mm: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
guint gdk_threads_add_timeout_seconds_full(gint _0, guint _1, GSourceFunc _2, gpointer _3, GDestroyNotify _4) __attribute__((weak)) {
	goPanic("gdk_threads_add_timeout_seconds_full: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
void gdk_app_launch_context_set_desktop(GdkAppLaunchContext* v, gint _0) __attribute__((weak)) {
	goPanic("gdk_app_launch_context_set_desktop: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
void gdk_app_launch_context_set_display(GdkAppLaunchContext* v, GdkDisplay* _0) __attribute__((weak)) {
	goPanic("gdk_app_launch_context_set_display: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
void gdk_app_launch_context_set_icon(GdkAppLaunchContext* v, GIcon* _0) __attribute__((weak)) {
	goPanic("gdk_app_launch_context_set_icon: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
void gdk_app_launch_context_set_icon_name(GdkAppLaunchContext* v, const char* _0) __attribute__((weak)) {
	goPanic("gdk_app_launch_context_set_icon_name: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
void gdk_app_launch_context_set_screen(GdkAppLaunchContext* v, GdkScreen* _0) __attribute__((weak)) {
	goPanic("gdk_app_launch_context_set_screen: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
void gdk_app_launch_context_set_timestamp(GdkAppLaunchContext* v, guint32 _0) __attribute__((weak)) {
	goPanic("gdk_app_launch_context_set_timestamp: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 14))
void gdk_test_render_sync(GdkWindow* _0) __attribute__((weak)) {
	goPanic("gdk_test_render_sync: library too old: needs at least 2.14");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 16))
gboolean gdk_keymap_get_caps_lock_state(GdkKeymap* v) __attribute__((weak)) {
	goPanic("gdk_keymap_get_caps_lock_state: library too old: needs at least 2.16");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 18))
GdkCursor* gdk_window_get_cursor(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_cursor: library too old: needs at least 2.18");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 18))
GdkWindow* gdk_offscreen_window_get_embedder(GdkWindow* _0) __attribute__((weak)) {
	goPanic("gdk_offscreen_window_get_embedder: library too old: needs at least 2.18");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 18))
gboolean gdk_window_ensure_native(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_ensure_native: library too old: needs at least 2.18");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 18))
gboolean gdk_window_is_destroyed(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_is_destroyed: library too old: needs at least 2.18");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 18))
void gdk_offscreen_window_set_embedder(GdkWindow* _0, GdkWindow* _1) __attribute__((weak)) {
	goPanic("gdk_offscreen_window_set_embedder: library too old: needs at least 2.18");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 18))
void gdk_window_flush(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_flush: library too old: needs at least 2.18");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 18))
void gdk_window_geometry_changed(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_geometry_changed: library too old: needs at least 2.18");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 18))
void gdk_window_get_root_coords(GdkWindow* v, gint _0, gint _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gdk_window_get_root_coords: library too old: needs at least 2.18");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 18))
void gdk_window_restack(GdkWindow* v, GdkWindow* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gdk_window_restack: library too old: needs at least 2.18");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GList* gdk_display_list_devices(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_list_devices: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GList* gdk_screen_get_toplevel_windows(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_toplevel_windows: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GList* gdk_screen_list_visuals(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_list_visuals: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GSList* gdk_display_manager_list_displays(GdkDisplayManager* v) __attribute__((weak)) {
	goPanic("gdk_display_manager_list_displays: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkCursor* gdk_cursor_new_for_display(GdkDisplay* _0, GdkCursorType _1) __attribute__((weak)) {
	goPanic("gdk_cursor_new_for_display: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkDisplay* gdk_cursor_get_display(GdkCursor* v) __attribute__((weak)) {
	goPanic("gdk_cursor_get_display: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkDisplay* gdk_display_get_default(void) __attribute__((weak)) {
	goPanic("gdk_display_get_default: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkDisplay* gdk_display_manager_get_default_display(GdkDisplayManager* v) __attribute__((weak)) {
	goPanic("gdk_display_manager_get_default_display: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkDisplay* gdk_display_open(const gchar* _0) __attribute__((weak)) {
	goPanic("gdk_display_open: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkDisplay* gdk_screen_get_display(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_display: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkDisplayManager* gdk_display_manager_get(void) __attribute__((weak)) {
	goPanic("gdk_display_manager_get: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkKeymap* gdk_keymap_get_for_display(GdkDisplay* _0) __attribute__((weak)) {
	goPanic("gdk_keymap_get_for_display: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkScreen* gdk_display_get_default_screen(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_get_default_screen: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkScreen* gdk_display_get_screen(GdkDisplay* v, gint _0) __attribute__((weak)) {
	goPanic("gdk_display_get_screen: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkScreen* gdk_screen_get_default(void) __attribute__((weak)) {
	goPanic("gdk_screen_get_default: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkScreen* gdk_visual_get_screen(GdkVisual* v) __attribute__((weak)) {
	goPanic("gdk_visual_get_screen: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkVisual* gdk_screen_get_system_visual(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_system_visual: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkWindow* gdk_display_get_window_at_pointer(GdkDisplay* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gdk_display_get_window_at_pointer: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
GdkWindow* gdk_screen_get_root_window(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_root_window: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
PangoContext* gdk_pango_context_get_for_screen(GdkScreen* _0) __attribute__((weak)) {
	goPanic("gdk_pango_context_get_for_screen: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
const gchar* gdk_display_get_name(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_get_name: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
const gchar* gdk_get_display_arg_name(void) __attribute__((weak)) {
	goPanic("gdk_get_display_arg_name: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gboolean gdk_display_pointer_is_grabbed(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_pointer_is_grabbed: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gboolean gdk_screen_get_setting(GdkScreen* v, const gchar* _0, GValue* _1) __attribute__((weak)) {
	goPanic("gdk_screen_get_setting: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gchar* gdk_screen_make_display_name(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_make_display_name: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gint gdk_display_get_n_screens(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_get_n_screens: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gint gdk_screen_get_height(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_height: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gint gdk_screen_get_height_mm(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_height_mm: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gint gdk_screen_get_monitor_at_point(GdkScreen* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gdk_screen_get_monitor_at_point: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gint gdk_screen_get_monitor_at_window(GdkScreen* v, GdkWindow* _0) __attribute__((weak)) {
	goPanic("gdk_screen_get_monitor_at_window: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gint gdk_screen_get_n_monitors(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_n_monitors: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gint gdk_screen_get_number(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_number: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gint gdk_screen_get_width(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_width: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
gint gdk_screen_get_width_mm(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_width_mm: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_display_beep(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_beep: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_display_close(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_close: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_display_get_pointer(GdkDisplay* v, GdkScreen** _0, gint* _1, gint* _2, GdkModifierType* _3) __attribute__((weak)) {
	goPanic("gdk_display_get_pointer: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_display_keyboard_ungrab(GdkDisplay* v, guint32 _0) __attribute__((weak)) {
	goPanic("gdk_display_keyboard_ungrab: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_display_manager_set_default_display(GdkDisplayManager* v, GdkDisplay* _0) __attribute__((weak)) {
	goPanic("gdk_display_manager_set_default_display: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_display_pointer_ungrab(GdkDisplay* v, guint32 _0) __attribute__((weak)) {
	goPanic("gdk_display_pointer_ungrab: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_display_set_double_click_time(GdkDisplay* v, guint _0) __attribute__((weak)) {
	goPanic("gdk_display_set_double_click_time: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_display_sync(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_sync: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_drag_find_window_for_screen(GdkDragContext* _0, GdkWindow* _1, GdkScreen* _2, gint _3, gint _4, GdkWindow** _5, GdkDragProtocol* _6) __attribute__((weak)) {
	goPanic("gdk_drag_find_window_for_screen: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_notify_startup_complete(void) __attribute__((weak)) {
	goPanic("gdk_notify_startup_complete: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_screen_get_monitor_geometry(GdkScreen* v, gint _0, GdkRectangle* _1) __attribute__((weak)) {
	goPanic("gdk_screen_get_monitor_geometry: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_window_fullscreen(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_fullscreen: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_window_set_skip_pager_hint(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_skip_pager_hint: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_window_set_skip_taskbar_hint(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_skip_taskbar_hint: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 2))
void gdk_window_unfullscreen(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_unfullscreen: library too old: needs at least 2.2");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 20))
GdkAxisUse gdk_device_get_axis_use(GdkDevice* v, guint _0) __attribute__((weak)) {
	goPanic("gdk_device_get_axis_use: library too old: needs at least 2.20");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 20))
GdkInputMode gdk_device_get_mode(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_mode: library too old: needs at least 2.20");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 20))
GdkInputSource gdk_device_get_source(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_source: library too old: needs at least 2.20");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 20))
const gchar* gdk_device_get_name(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_name: library too old: needs at least 2.20");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 20))
gboolean gdk_device_get_has_cursor(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_has_cursor: library too old: needs at least 2.20");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 20))
gboolean gdk_device_get_key(GdkDevice* v, guint _0, guint* _1, GdkModifierType* _2) __attribute__((weak)) {
	goPanic("gdk_device_get_key: library too old: needs at least 2.20");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 20))
gint gdk_screen_get_primary_monitor(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_primary_monitor: library too old: needs at least 2.20");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
GdkByteOrder gdk_visual_get_byte_order(GdkVisual* v) __attribute__((weak)) {
	goPanic("gdk_visual_get_byte_order: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
GdkCursorType gdk_cursor_get_cursor_type(GdkCursor* v) __attribute__((weak)) {
	goPanic("gdk_cursor_get_cursor_type: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
GdkDragAction gdk_drag_context_get_actions(GdkDragContext* v) __attribute__((weak)) {
	goPanic("gdk_drag_context_get_actions: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
GdkDragAction gdk_drag_context_get_selected_action(GdkDragContext* v) __attribute__((weak)) {
	goPanic("gdk_drag_context_get_selected_action: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
GdkDragAction gdk_drag_context_get_suggested_action(GdkDragContext* v) __attribute__((weak)) {
	goPanic("gdk_drag_context_get_suggested_action: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
GdkVisualType gdk_visual_get_visual_type(GdkVisual* v) __attribute__((weak)) {
	goPanic("gdk_visual_get_visual_type: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
GdkWindow* gdk_drag_context_get_source_window(GdkDragContext* v) __attribute__((weak)) {
	goPanic("gdk_drag_context_get_source_window: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
GdkWindow* gdk_window_get_effective_parent(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_effective_parent: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
GdkWindow* gdk_window_get_effective_toplevel(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_effective_toplevel: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
cairo_pattern_t* gdk_window_get_background_pattern(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_background_pattern: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
cairo_surface_t* gdk_window_create_similar_surface(GdkWindow* v, cairo_content_t _0, int _1, int _2) __attribute__((weak)) {
	goPanic("gdk_window_create_similar_surface: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gboolean gdk_display_is_closed(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_is_closed: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gboolean gdk_window_get_accept_focus(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_accept_focus: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gboolean gdk_window_get_composited(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_composited: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gboolean gdk_window_get_focus_on_map(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_focus_on_map: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gboolean gdk_window_get_modal_hint(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_modal_hint: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gboolean gdk_window_has_native(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_has_native: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gboolean gdk_window_is_input_only(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_is_input_only: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gboolean gdk_window_is_shaped(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_is_shaped: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gint gdk_visual_get_bits_per_rgb(GdkVisual* v) __attribute__((weak)) {
	goPanic("gdk_visual_get_bits_per_rgb: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gint gdk_visual_get_colormap_size(GdkVisual* v) __attribute__((weak)) {
	goPanic("gdk_visual_get_colormap_size: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
gint gdk_visual_get_depth(GdkVisual* v) __attribute__((weak)) {
	goPanic("gdk_visual_get_depth: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
void gdk_visual_get_blue_pixel_details(GdkVisual* v, guint32* _0, gint* _1, gint* _2) __attribute__((weak)) {
	goPanic("gdk_visual_get_blue_pixel_details: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
void gdk_visual_get_green_pixel_details(GdkVisual* v, guint32* _0, gint* _1, gint* _2) __attribute__((weak)) {
	goPanic("gdk_visual_get_green_pixel_details: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
void gdk_visual_get_red_pixel_details(GdkVisual* v, guint32* _0, gint* _1, gint* _2) __attribute__((weak)) {
	goPanic("gdk_visual_get_red_pixel_details: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
void gdk_window_coords_from_parent(GdkWindow* v, gdouble _0, gdouble _1, gdouble* _2, gdouble* _3) __attribute__((weak)) {
	goPanic("gdk_window_coords_from_parent: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 22))
void gdk_window_coords_to_parent(GdkWindow* v, gdouble _0, gdouble _1, gdouble* _2, gdouble* _3) __attribute__((weak)) {
	goPanic("gdk_window_coords_to_parent: library too old: needs at least 2.22");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 24))
GdkDisplay* gdk_window_get_display(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_display: library too old: needs at least 2.24");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 24))
GdkScreen* gdk_window_get_screen(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_screen: library too old: needs at least 2.24");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 24))
GdkVisual* gdk_window_get_visual(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_visual: library too old: needs at least 2.24");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 24))
gint gdk_device_get_n_keys(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_n_keys: library too old: needs at least 2.24");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 24))
int gdk_window_get_height(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_height: library too old: needs at least 2.24");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 24))
int gdk_window_get_width(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_width: library too old: needs at least 2.24");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 24))
void gdk_cairo_set_source_window(cairo_t* _0, GdkWindow* _1, gdouble _2, gdouble _3) __attribute__((weak)) {
	goPanic("gdk_cairo_set_source_window: library too old: needs at least 2.24");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
GdkCursor* gdk_cursor_new_from_pixbuf(GdkDisplay* _0, GdkPixbuf* _1, gint _2, gint _3) __attribute__((weak)) {
	goPanic("gdk_cursor_new_from_pixbuf: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
GdkWindow* gdk_display_get_default_group(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_get_default_group: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
GdkWindow* gdk_window_get_group(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_group: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
gboolean gdk_display_supports_cursor_alpha(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_supports_cursor_alpha: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
gboolean gdk_display_supports_cursor_color(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_supports_cursor_color: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
guint gdk_display_get_default_cursor_size(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_get_default_cursor_size: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
void gdk_display_flush(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_flush: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
void gdk_display_get_maximal_cursor_size(GdkDisplay* v, guint* _0, guint* _1) __attribute__((weak)) {
	goPanic("gdk_display_get_maximal_cursor_size: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
void gdk_display_set_double_click_distance(GdkDisplay* v, guint _0) __attribute__((weak)) {
	goPanic("gdk_display_set_double_click_distance: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
void gdk_window_set_accept_focus(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_accept_focus: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
void gdk_window_set_keep_above(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_keep_above: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 4))
void gdk_window_set_keep_below(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_keep_below: library too old: needs at least 2.4");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 6))
gboolean gdk_display_supports_clipboard_persistence(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_supports_clipboard_persistence: library too old: needs at least 2.6");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 6))
gboolean gdk_display_supports_selection_notification(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_supports_selection_notification: library too old: needs at least 2.6");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 6))
gboolean gdk_drag_drop_succeeded(GdkDragContext* _0) __attribute__((weak)) {
	goPanic("gdk_drag_drop_succeeded: library too old: needs at least 2.6");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 6))
void gdk_window_configure_finished(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_configure_finished: library too old: needs at least 2.6");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 6))
void gdk_window_enable_synchronized_configure(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_enable_synchronized_configure: library too old: needs at least 2.6");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 6))
void gdk_window_set_focus_on_map(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_focus_on_map: library too old: needs at least 2.6");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
GdkCursor* gdk_cursor_new_from_name(GdkDisplay* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gdk_cursor_new_from_name: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
GdkPixbuf* gdk_cursor_get_image(GdkCursor* v) __attribute__((weak)) {
	goPanic("gdk_cursor_get_image: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
GdkVisual* gdk_screen_get_rgba_visual(GdkScreen* v) __attribute__((weak)) {
	goPanic("gdk_screen_get_rgba_visual: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
cairo_t* gdk_cairo_create(GdkWindow* _0) __attribute__((weak)) {
	goPanic("gdk_cairo_create: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
void gdk_cairo_rectangle(cairo_t* _0, const GdkRectangle* _1) __attribute__((weak)) {
	goPanic("gdk_cairo_rectangle: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
void gdk_cairo_region(cairo_t* _0, const cairo_region_t* _1) __attribute__((weak)) {
	goPanic("gdk_cairo_region: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
void gdk_cairo_set_source_color(cairo_t* _0, const GdkColor* _1) __attribute__((weak)) {
	goPanic("gdk_cairo_set_source_color: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
void gdk_cairo_set_source_pixbuf(cairo_t* _0, const GdkPixbuf* _1, gdouble _2, gdouble _3) __attribute__((weak)) {
	goPanic("gdk_cairo_set_source_pixbuf: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
void gdk_display_warp_pointer(GdkDisplay* v, GdkScreen* _0, gint _1, gint _2) __attribute__((weak)) {
	goPanic("gdk_display_warp_pointer: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
void gdk_window_move_region(GdkWindow* v, const cairo_region_t* _0, gint _1, gint _2) __attribute__((weak)) {
	goPanic("gdk_window_move_region: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 2 || (GDK_MAJOR_VERSION == 2 && GDK_MINOR_VERSION < 8))
void gdk_window_set_urgency_hint(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_urgency_hint: library too old: needs at least 2.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GList* gdk_device_manager_list_devices(GdkDeviceManager* v, GdkDeviceType _0) __attribute__((weak)) {
	goPanic("gdk_device_manager_list_devices: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkAppLaunchContext* gdk_display_get_app_launch_context(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_get_app_launch_context: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkCursor* gdk_window_get_device_cursor(GdkWindow* v, GdkDevice* _0) __attribute__((weak)) {
	goPanic("gdk_window_get_device_cursor: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkDevice* gdk_device_get_associated_device(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_associated_device: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkDevice* gdk_device_manager_get_client_pointer(GdkDeviceManager* v) __attribute__((weak)) {
	goPanic("gdk_device_manager_get_client_pointer: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkDeviceManager* gdk_display_get_device_manager(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_get_device_manager: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkDeviceType gdk_device_get_device_type(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_device_type: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkDisplay* gdk_device_get_display(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_display: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkDisplay* gdk_device_manager_get_display(GdkDeviceManager* v) __attribute__((weak)) {
	goPanic("gdk_device_manager_get_display: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkDisplay* gdk_display_manager_open_display(GdkDisplayManager* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gdk_display_manager_open_display: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkDragProtocol gdk_drag_context_get_protocol(GdkDragContext* v) __attribute__((weak)) {
	goPanic("gdk_drag_context_get_protocol: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkDragProtocol gdk_window_get_drag_protocol(GdkWindow* v, GdkWindow** _0) __attribute__((weak)) {
	goPanic("gdk_window_get_drag_protocol: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkEventMask gdk_window_get_device_events(GdkWindow* v, GdkDevice* _0) __attribute__((weak)) {
	goPanic("gdk_window_get_device_events: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkGrabStatus gdk_device_grab(GdkDevice* v, GdkWindow* _0, GdkGrabOwnership _1, gboolean _2, GdkEventMask _3, GdkCursor* _4, guint32 _5) __attribute__((weak)) {
	goPanic("gdk_device_grab: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkRGBA* gdk_rgba_copy(const GdkRGBA* v) __attribute__((weak)) {
	goPanic("gdk_rgba_copy: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkWindow* gdk_device_get_window_at_position(GdkDevice* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gdk_device_get_window_at_position: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkWindow* gdk_device_get_window_at_position_double(GdkDevice* v, gdouble* _0, gdouble* _1) __attribute__((weak)) {
	goPanic("gdk_device_get_window_at_position_double: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkWindow* gdk_drag_context_get_dest_window(GdkDragContext* v) __attribute__((weak)) {
	goPanic("gdk_drag_context_get_dest_window: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
GdkWindow* gdk_window_get_device_position(GdkWindow* v, GdkDevice* _0, gint* _1, gint* _2, GdkModifierType* _3) __attribute__((weak)) {
	goPanic("gdk_window_get_device_position: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
gboolean gdk_display_has_pending(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_has_pending: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
gboolean gdk_keymap_get_num_lock_state(GdkKeymap* v) __attribute__((weak)) {
	goPanic("gdk_keymap_get_num_lock_state: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
gboolean gdk_rgba_equal(gconstpointer v, gconstpointer _0) __attribute__((weak)) {
	goPanic("gdk_rgba_equal: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
gboolean gdk_rgba_parse(GdkRGBA* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gdk_rgba_parse: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
gboolean gdk_window_get_support_multidevice(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_support_multidevice: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
gchar* gdk_rgba_to_string(const GdkRGBA* v) __attribute__((weak)) {
	goPanic("gdk_rgba_to_string: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
gint gdk_device_get_n_axes(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_n_axes: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
guint gdk_rgba_hash(gconstpointer v) __attribute__((weak)) {
	goPanic("gdk_rgba_hash: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_cairo_set_source_rgba(cairo_t* _0, const GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gdk_cairo_set_source_rgba: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_device_get_position(GdkDevice* v, GdkScreen** _0, gint* _1, gint* _2) __attribute__((weak)) {
	goPanic("gdk_device_get_position: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_device_ungrab(GdkDevice* v, guint32 _0) __attribute__((weak)) {
	goPanic("gdk_device_ungrab: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_device_warp(GdkDevice* v, GdkScreen* _0, gint _1, gint _2) __attribute__((weak)) {
	goPanic("gdk_device_warp: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_disable_multidevice(void) __attribute__((weak)) {
	goPanic("gdk_disable_multidevice: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_display_notify_startup_complete(GdkDisplay* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gdk_display_notify_startup_complete: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_error_trap_pop_ignored(void) __attribute__((weak)) {
	goPanic("gdk_error_trap_pop_ignored: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_window_set_device_cursor(GdkWindow* v, GdkDevice* _0, GdkCursor* _1) __attribute__((weak)) {
	goPanic("gdk_window_set_device_cursor: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_window_set_device_events(GdkWindow* v, GdkDevice* _0, GdkEventMask _1) __attribute__((weak)) {
	goPanic("gdk_window_set_device_events: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_window_set_source_events(GdkWindow* v, GdkInputSource _0, GdkEventMask _1) __attribute__((weak)) {
	goPanic("gdk_window_set_source_events: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 0))
void gdk_window_set_support_multidevice(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_support_multidevice: library too old: needs at least 3.0");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
GList* gdk_window_get_children_with_user_data(GdkWindow* v, gpointer _0) __attribute__((weak)) {
	goPanic("gdk_window_get_children_with_user_data: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
GdkCursor* gdk_cursor_new_from_surface(GdkDisplay* _0, cairo_surface_t* _1, gdouble _2, gdouble _3) __attribute__((weak)) {
	goPanic("gdk_cursor_new_from_surface: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
GdkWindow* gdk_window_get_device_position_double(GdkWindow* v, GdkDevice* _0, gdouble* _1, gdouble* _2, GdkModifierType* _3) __attribute__((weak)) {
	goPanic("gdk_window_get_device_position_double: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
cairo_surface_t* gdk_cairo_surface_create_from_pixbuf(const GdkPixbuf* _0, int _1, GdkWindow* _2) __attribute__((weak)) {
	goPanic("gdk_cairo_surface_create_from_pixbuf: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
cairo_surface_t* gdk_cursor_get_surface(GdkCursor* v, gdouble* _0, gdouble* _1) __attribute__((weak)) {
	goPanic("gdk_cursor_get_surface: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
cairo_surface_t* gdk_window_create_similar_image_surface(GdkWindow* v, cairo_format_t _0, int _1, int _2, int _3) __attribute__((weak)) {
	goPanic("gdk_window_create_similar_image_surface: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
gint gdk_screen_get_monitor_scale_factor(GdkScreen* v, gint _0) __attribute__((weak)) {
	goPanic("gdk_screen_get_monitor_scale_factor: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
gint gdk_window_get_scale_factor(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_scale_factor: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
void gdk_device_get_position_double(GdkDevice* v, GdkScreen** _0, gdouble* _1, gdouble* _2) __attribute__((weak)) {
	goPanic("gdk_device_get_position_double: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
void gdk_set_allowed_backends(const gchar* _0) __attribute__((weak)) {
	goPanic("gdk_set_allowed_backends: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 10))
void gdk_window_set_opaque_region(GdkWindow* v, cairo_region_t* _0) __attribute__((weak)) {
	goPanic("gdk_window_set_opaque_region: library too old: needs at least 3.10");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 12))
GdkWindow* gdk_device_get_last_event_window(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_last_event_window: library too old: needs at least 3.12");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 12))
gboolean gdk_window_get_event_compression(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_event_compression: library too old: needs at least 3.12");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 12))
void gdk_window_set_event_compression(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_event_compression: library too old: needs at least 3.12");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 12))
void gdk_window_set_shadow_width(GdkWindow* v, gint _0, gint _1, gint _2, gint _3) __attribute__((weak)) {
	goPanic("gdk_window_set_shadow_width: library too old: needs at least 3.12");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
GdkDisplay* gdk_gl_context_get_display(GdkGLContext* v) __attribute__((weak)) {
	goPanic("gdk_gl_context_get_display: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
GdkGLContext* gdk_gl_context_get_current(void) __attribute__((weak)) {
	goPanic("gdk_gl_context_get_current: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
GdkGLContext* gdk_gl_context_get_shared_context(GdkGLContext* v) __attribute__((weak)) {
	goPanic("gdk_gl_context_get_shared_context: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
GdkGLContext* gdk_window_create_gl_context(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_create_gl_context: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
GdkWindow* gdk_gl_context_get_window(GdkGLContext* v) __attribute__((weak)) {
	goPanic("gdk_gl_context_get_window: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
const gchar* gdk_device_get_product_id(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_product_id: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
const gchar* gdk_device_get_vendor_id(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_vendor_id: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
gboolean gdk_gl_context_get_debug_enabled(GdkGLContext* v) __attribute__((weak)) {
	goPanic("gdk_gl_context_get_debug_enabled: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
gboolean gdk_gl_context_get_forward_compatible(GdkGLContext* v) __attribute__((weak)) {
	goPanic("gdk_gl_context_get_forward_compatible: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
gboolean gdk_gl_context_realize(GdkGLContext* v) __attribute__((weak)) {
	goPanic("gdk_gl_context_realize: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
void gdk_cairo_draw_from_gl(cairo_t* _0, GdkWindow* _1, int _2, int _3, int _4, int _5, int _6, int _7, int _8) __attribute__((weak)) {
	goPanic("gdk_cairo_draw_from_gl: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
void gdk_gl_context_clear_current(void) __attribute__((weak)) {
	goPanic("gdk_gl_context_clear_current: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
void gdk_gl_context_get_required_version(GdkGLContext* v, int* _0, int* _1) __attribute__((weak)) {
	goPanic("gdk_gl_context_get_required_version: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
void gdk_gl_context_get_version(GdkGLContext* v, int* _0, int* _1) __attribute__((weak)) {
	goPanic("gdk_gl_context_get_version: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
void gdk_gl_context_make_current(GdkGLContext* v) __attribute__((weak)) {
	goPanic("gdk_gl_context_make_current: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
void gdk_gl_context_set_debug_enabled(GdkGLContext* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_gl_context_set_debug_enabled: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
void gdk_gl_context_set_forward_compatible(GdkGLContext* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_gl_context_set_forward_compatible: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
void gdk_gl_context_set_required_version(GdkGLContext* v, int _0, int _1) __attribute__((weak)) {
	goPanic("gdk_gl_context_set_required_version: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 16))
void gdk_window_mark_paint_from_clip(GdkWindow* v, cairo_t* _0) __attribute__((weak)) {
	goPanic("gdk_window_mark_paint_from_clip: library too old: needs at least 3.16");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 18))
gboolean gdk_keymap_get_scroll_lock_state(GdkKeymap* v) __attribute__((weak)) {
	goPanic("gdk_keymap_get_scroll_lock_state: library too old: needs at least 3.18");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 18))
gboolean gdk_window_get_pass_through(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_pass_through: library too old: needs at least 3.18");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 18))
void gdk_window_set_pass_through(GdkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gdk_window_set_pass_through: library too old: needs at least 3.18");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
GList* gdk_display_list_seats(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_list_seats: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
GList* gdk_seat_get_slaves(GdkSeat* v, GdkSeatCapabilities _0) __attribute__((weak)) {
	goPanic("gdk_seat_get_slaves: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
GdkDevice* gdk_seat_get_keyboard(GdkSeat* v) __attribute__((weak)) {
	goPanic("gdk_seat_get_keyboard: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
GdkDevice* gdk_seat_get_pointer(GdkSeat* v) __attribute__((weak)) {
	goPanic("gdk_seat_get_pointer: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
GdkSeat* gdk_device_get_seat(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_seat: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
GdkSeat* gdk_display_get_default_seat(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_get_default_seat: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
GdkSeatCapabilities gdk_seat_get_capabilities(GdkSeat* v) __attribute__((weak)) {
	goPanic("gdk_seat_get_capabilities: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
GdkWindow* gdk_drag_context_get_drag_window(GdkDragContext* v) __attribute__((weak)) {
	goPanic("gdk_drag_context_get_drag_window: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
gboolean gdk_drag_context_manage_dnd(GdkDragContext* v, GdkWindow* _0, GdkDragAction _1) __attribute__((weak)) {
	goPanic("gdk_drag_context_manage_dnd: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
gboolean gdk_gl_context_is_legacy(GdkGLContext* v) __attribute__((weak)) {
	goPanic("gdk_gl_context_is_legacy: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
gboolean gdk_rectangle_equal(const GdkRectangle* v, const GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gdk_rectangle_equal: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
void gdk_drag_context_set_hotspot(GdkDragContext* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gdk_drag_context_set_hotspot: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
void gdk_drag_drop_done(GdkDragContext* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gdk_drag_drop_done: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 20))
void gdk_seat_ungrab(GdkSeat* v) __attribute__((weak)) {
	goPanic("gdk_seat_ungrab: library too old: needs at least 3.20");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkAxisFlags gdk_device_get_axes(GdkDevice* v) __attribute__((weak)) {
	goPanic("gdk_device_get_axes: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkDeviceToolType gdk_device_tool_get_tool_type(GdkDeviceTool* v) __attribute__((weak)) {
	goPanic("gdk_device_tool_get_tool_type: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkDisplay* gdk_monitor_get_display(GdkMonitor* v) __attribute__((weak)) {
	goPanic("gdk_monitor_get_display: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkDrawingContext* gdk_cairo_get_drawing_context(cairo_t* _0) __attribute__((weak)) {
	goPanic("gdk_cairo_get_drawing_context: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkDrawingContext* gdk_window_begin_draw_frame(GdkWindow* v, const cairo_region_t* _0) __attribute__((weak)) {
	goPanic("gdk_window_begin_draw_frame: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkMonitor* gdk_display_get_monitor(GdkDisplay* v, int _0) __attribute__((weak)) {
	goPanic("gdk_display_get_monitor: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkMonitor* gdk_display_get_monitor_at_point(GdkDisplay* v, int _0, int _1) __attribute__((weak)) {
	goPanic("gdk_display_get_monitor_at_point: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkMonitor* gdk_display_get_monitor_at_window(GdkDisplay* v, GdkWindow* _0) __attribute__((weak)) {
	goPanic("gdk_display_get_monitor_at_window: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkMonitor* gdk_display_get_primary_monitor(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_get_primary_monitor: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkSubpixelLayout gdk_monitor_get_subpixel_layout(GdkMonitor* v) __attribute__((weak)) {
	goPanic("gdk_monitor_get_subpixel_layout: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
GdkWindow* gdk_drawing_context_get_window(GdkDrawingContext* v) __attribute__((weak)) {
	goPanic("gdk_drawing_context_get_window: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
PangoContext* gdk_pango_context_get_for_display(GdkDisplay* _0) __attribute__((weak)) {
	goPanic("gdk_pango_context_get_for_display: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
cairo_region_t* gdk_drawing_context_get_clip(GdkDrawingContext* v) __attribute__((weak)) {
	goPanic("gdk_drawing_context_get_clip: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
cairo_t* gdk_drawing_context_get_cairo_context(GdkDrawingContext* v) __attribute__((weak)) {
	goPanic("gdk_drawing_context_get_cairo_context: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
gboolean gdk_drawing_context_is_valid(GdkDrawingContext* v) __attribute__((weak)) {
	goPanic("gdk_drawing_context_is_valid: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
gboolean gdk_gl_context_get_use_es(GdkGLContext* v) __attribute__((weak)) {
	goPanic("gdk_gl_context_get_use_es: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
gboolean gdk_monitor_is_primary(GdkMonitor* v) __attribute__((weak)) {
	goPanic("gdk_monitor_is_primary: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
gint gdk_device_pad_get_feature_group(GdkDevicePad* v, GdkDevicePadFeature _0, gint _1) __attribute__((weak)) {
	goPanic("gdk_device_pad_get_feature_group: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
gint gdk_device_pad_get_group_n_modes(GdkDevicePad* v, gint _0) __attribute__((weak)) {
	goPanic("gdk_device_pad_get_group_n_modes: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
gint gdk_device_pad_get_n_features(GdkDevicePad* v, GdkDevicePadFeature _0) __attribute__((weak)) {
	goPanic("gdk_device_pad_get_n_features: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
gint gdk_device_pad_get_n_groups(GdkDevicePad* v) __attribute__((weak)) {
	goPanic("gdk_device_pad_get_n_groups: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
guint64 gdk_device_tool_get_hardware_id(GdkDeviceTool* v) __attribute__((weak)) {
	goPanic("gdk_device_tool_get_hardware_id: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
guint64 gdk_device_tool_get_serial(GdkDeviceTool* v) __attribute__((weak)) {
	goPanic("gdk_device_tool_get_serial: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
int gdk_display_get_n_monitors(GdkDisplay* v) __attribute__((weak)) {
	goPanic("gdk_display_get_n_monitors: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
int gdk_monitor_get_height_mm(GdkMonitor* v) __attribute__((weak)) {
	goPanic("gdk_monitor_get_height_mm: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
int gdk_monitor_get_refresh_rate(GdkMonitor* v) __attribute__((weak)) {
	goPanic("gdk_monitor_get_refresh_rate: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
int gdk_monitor_get_scale_factor(GdkMonitor* v) __attribute__((weak)) {
	goPanic("gdk_monitor_get_scale_factor: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
int gdk_monitor_get_width_mm(GdkMonitor* v) __attribute__((weak)) {
	goPanic("gdk_monitor_get_width_mm: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
void gdk_gl_context_set_use_es(GdkGLContext* v, int _0) __attribute__((weak)) {
	goPanic("gdk_gl_context_set_use_es: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
void gdk_monitor_get_geometry(GdkMonitor* v, GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gdk_monitor_get_geometry: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
void gdk_monitor_get_workarea(GdkMonitor* v, GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gdk_monitor_get_workarea: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 22))
void gdk_window_end_draw_frame(GdkWindow* v, GdkDrawingContext* _0) __attribute__((weak)) {
	goPanic("gdk_window_end_draw_frame: library too old: needs at least 3.22");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 24))
void gdk_window_move_to_rect(GdkWindow* v, const GdkRectangle* _0, GdkGravity _1, GdkGravity _2, GdkAnchorHints _3, gint _4, gint _5) __attribute__((weak)) {
	goPanic("gdk_window_move_to_rect: library too old: needs at least 3.24");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 4))
GdkModifierType gdk_keymap_get_modifier_mask(GdkKeymap* v, GdkModifierIntent _0) __attribute__((weak)) {
	goPanic("gdk_keymap_get_modifier_mask: library too old: needs at least 3.4");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 4))
guint gdk_keymap_get_modifier_state(GdkKeymap* v) __attribute__((weak)) {
	goPanic("gdk_keymap_get_modifier_state: library too old: needs at least 3.4");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 4))
void gdk_screen_get_monitor_workarea(GdkScreen* v, gint _0, GdkRectangle* _1) __attribute__((weak)) {
	goPanic("gdk_screen_get_monitor_workarea: library too old: needs at least 3.4");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 4))
void gdk_window_begin_move_drag_for_device(GdkWindow* v, GdkDevice* _0, gint _1, gint _2, gint _3, guint32 _4) __attribute__((weak)) {
	goPanic("gdk_window_begin_move_drag_for_device: library too old: needs at least 3.4");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 4))
void gdk_window_begin_resize_drag_for_device(GdkWindow* v, GdkWindowEdge _0, GdkDevice* _1, gint _2, gint _3, gint _4, guint32 _5) __attribute__((weak)) {
	goPanic("gdk_window_begin_resize_drag_for_device: library too old: needs at least 3.4");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
GdkFrameClock* gdk_window_get_frame_clock(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_frame_clock: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
GdkFrameTimings* gdk_frame_clock_get_current_timings(GdkFrameClock* v) __attribute__((weak)) {
	goPanic("gdk_frame_clock_get_current_timings: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
GdkFrameTimings* gdk_frame_clock_get_timings(GdkFrameClock* v, gint64 _0) __attribute__((weak)) {
	goPanic("gdk_frame_clock_get_timings: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
GdkFullscreenMode gdk_window_get_fullscreen_mode(GdkWindow* v) __attribute__((weak)) {
	goPanic("gdk_window_get_fullscreen_mode: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
gboolean gdk_frame_timings_get_complete(GdkFrameTimings* v) __attribute__((weak)) {
	goPanic("gdk_frame_timings_get_complete: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
gint64 gdk_frame_clock_get_frame_counter(GdkFrameClock* v) __attribute__((weak)) {
	goPanic("gdk_frame_clock_get_frame_counter: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
gint64 gdk_frame_clock_get_frame_time(GdkFrameClock* v) __attribute__((weak)) {
	goPanic("gdk_frame_clock_get_frame_time: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
gint64 gdk_frame_clock_get_history_start(GdkFrameClock* v) __attribute__((weak)) {
	goPanic("gdk_frame_clock_get_history_start: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
gint64 gdk_frame_timings_get_frame_counter(GdkFrameTimings* v) __attribute__((weak)) {
	goPanic("gdk_frame_timings_get_frame_counter: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
gint64 gdk_frame_timings_get_predicted_presentation_time(GdkFrameTimings* v) __attribute__((weak)) {
	goPanic("gdk_frame_timings_get_predicted_presentation_time: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
gint64 gdk_frame_timings_get_presentation_time(GdkFrameTimings* v) __attribute__((weak)) {
	goPanic("gdk_frame_timings_get_presentation_time: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
gint64 gdk_frame_timings_get_refresh_interval(GdkFrameTimings* v) __attribute__((weak)) {
	goPanic("gdk_frame_timings_get_refresh_interval: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
void gdk_frame_clock_begin_updating(GdkFrameClock* v) __attribute__((weak)) {
	goPanic("gdk_frame_clock_begin_updating: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
void gdk_frame_clock_end_updating(GdkFrameClock* v) __attribute__((weak)) {
	goPanic("gdk_frame_clock_end_updating: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
void gdk_frame_clock_get_refresh_info(GdkFrameClock* v, gint64 _0, gint64* _1, gint64* _2) __attribute__((weak)) {
	goPanic("gdk_frame_clock_get_refresh_info: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
void gdk_frame_clock_request_phase(GdkFrameClock* v, GdkFrameClockPhase _0) __attribute__((weak)) {
	goPanic("gdk_frame_clock_request_phase: library too old: needs at least 3.8");
}
#endif

#if (GDK_MAJOR_VERSION < 3 || (GDK_MAJOR_VERSION == 3 && GDK_MINOR_VERSION < 8))
void gdk_window_set_fullscreen_mode(GdkWindow* v, GdkFullscreenMode _0) __attribute__((weak)) {
	goPanic("gdk_window_set_fullscreen_mode: library too old: needs at least 3.8");
}
#endif
