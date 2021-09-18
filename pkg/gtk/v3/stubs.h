#include <glib-object.h>
#include <gtk/gtk-a11y.h>
#include <gtk/gtk.h>
#include <gtk/gtkx.h>

extern void goPanic(const char*);

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GList* gtk_recent_chooser_get_items(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_items: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GList* gtk_recent_manager_get_items(GtkRecentManager* v) __attribute__((weak)) {
	goPanic("gtk_recent_manager_get_items: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GQuark gtk_print_error_quark(void) __attribute__((weak)) {
	goPanic("gtk_print_error_quark: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GSList* gtk_recent_chooser_list_filters(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_list_filters: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GSList* gtk_size_group_get_widgets(GtkSizeGroup* v) __attribute__((weak)) {
	goPanic("gtk_size_group_get_widgets: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GdkPixbuf* gtk_assistant_get_page_header_image(GtkAssistant* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_assistant_get_page_header_image: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GdkPixbuf* gtk_assistant_get_page_side_image(GtkAssistant* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_assistant_get_page_side_image: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GdkPixbuf* gtk_recent_info_get_icon(GtkRecentInfo* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_icon: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GdkPixbuf* gtk_status_icon_get_pixbuf(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_pixbuf: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkAssistantPageType gtk_assistant_get_page_type(GtkAssistant* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_assistant_get_page_type: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkCellRenderer* gtk_cell_renderer_accel_new(void) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_accel_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkCellRenderer* gtk_cell_renderer_spin_new(void) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_spin_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkEntry* gtk_tree_view_get_search_entry(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_search_entry: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkImageType gtk_status_icon_get_storage_type(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_storage_type: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPageOrientation gtk_page_setup_get_orientation(GtkPageSetup* v) __attribute__((weak)) {
	goPanic("gtk_page_setup_get_orientation: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPageOrientation gtk_print_settings_get_orientation(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_orientation: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPageRange* gtk_print_settings_get_page_ranges(GtkPrintSettings* v, gint* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_page_ranges: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPageSet gtk_print_settings_get_page_set(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_page_set: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPageSetup* gtk_page_setup_copy(GtkPageSetup* v) __attribute__((weak)) {
	goPanic("gtk_page_setup_copy: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPageSetup* gtk_page_setup_new(void) __attribute__((weak)) {
	goPanic("gtk_page_setup_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPageSetup* gtk_print_context_get_page_setup(GtkPrintContext* v) __attribute__((weak)) {
	goPanic("gtk_print_context_get_page_setup: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPageSetup* gtk_print_operation_get_default_page_setup(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_get_default_page_setup: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPageSetup* gtk_print_run_page_setup_dialog(GtkWindow* _0, GtkPageSetup* _1, GtkPrintSettings* _2) __attribute__((weak)) {
	goPanic("gtk_print_run_page_setup_dialog: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPaperSize* gtk_page_setup_get_paper_size(GtkPageSetup* v) __attribute__((weak)) {
	goPanic("gtk_page_setup_get_paper_size: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPaperSize* gtk_paper_size_copy(GtkPaperSize* v) __attribute__((weak)) {
	goPanic("gtk_paper_size_copy: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPaperSize* gtk_paper_size_new(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_paper_size_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPaperSize* gtk_paper_size_new_custom(const gchar* _0, const gchar* _1, gdouble _2, gdouble _3, GtkUnit _4) __attribute__((weak)) {
	goPanic("gtk_paper_size_new_custom: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPaperSize* gtk_paper_size_new_from_ppd(const gchar* _0, const gchar* _1, gdouble _2, gdouble _3) __attribute__((weak)) {
	goPanic("gtk_paper_size_new_from_ppd: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPaperSize* gtk_print_settings_get_paper_size(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_paper_size: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPositionType gtk_button_get_image_position(GtkButton* v) __attribute__((weak)) {
	goPanic("gtk_button_get_image_position: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPrintDuplex gtk_print_settings_get_duplex(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_duplex: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPrintOperation* gtk_print_operation_new(void) __attribute__((weak)) {
	goPanic("gtk_print_operation_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPrintOperationResult gtk_print_operation_run(GtkPrintOperation* v, GtkPrintOperationAction _0, GtkWindow* _1) __attribute__((weak)) {
	goPanic("gtk_print_operation_run: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPrintPages gtk_print_settings_get_print_pages(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_print_pages: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPrintQuality gtk_print_settings_get_quality(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_quality: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPrintSettings* gtk_print_operation_get_print_settings(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_get_print_settings: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPrintSettings* gtk_print_settings_copy(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_copy: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPrintSettings* gtk_print_settings_new(void) __attribute__((weak)) {
	goPanic("gtk_print_settings_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkPrintStatus gtk_print_operation_get_status(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_get_status: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkRecentFilter* gtk_recent_chooser_get_filter(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_filter: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkRecentFilter* gtk_recent_filter_new(void) __attribute__((weak)) {
	goPanic("gtk_recent_filter_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkRecentFilterFlags gtk_recent_filter_get_needed(GtkRecentFilter* v) __attribute__((weak)) {
	goPanic("gtk_recent_filter_get_needed: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkRecentInfo* gtk_recent_chooser_get_current_item(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_current_item: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkRecentInfo* gtk_recent_manager_lookup_item(GtkRecentManager* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_manager_lookup_item: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkRecentManager* gtk_recent_manager_get_default(void) __attribute__((weak)) {
	goPanic("gtk_recent_manager_get_default: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkRecentManager* gtk_recent_manager_new(void) __attribute__((weak)) {
	goPanic("gtk_recent_manager_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkRecentSortType gtk_recent_chooser_get_sort_type(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_sort_type: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkSensitivityType gtk_range_get_lower_stepper_sensitivity(GtkRange* v) __attribute__((weak)) {
	goPanic("gtk_range_get_lower_stepper_sensitivity: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkSensitivityType gtk_range_get_upper_stepper_sensitivity(GtkRange* v) __attribute__((weak)) {
	goPanic("gtk_range_get_upper_stepper_sensitivity: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkStatusIcon* gtk_status_icon_new(void) __attribute__((weak)) {
	goPanic("gtk_status_icon_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkStatusIcon* gtk_status_icon_new_from_file(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_new_from_file: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkStatusIcon* gtk_status_icon_new_from_icon_name(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_new_from_icon_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkStatusIcon* gtk_status_icon_new_from_pixbuf(GdkPixbuf* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_new_from_pixbuf: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkStatusIcon* gtk_status_icon_new_from_stock(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_new_from_stock: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkTargetEntry* gtk_target_table_new_from_list(GtkTargetList* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_target_table_new_from_list: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkTargetList* gtk_text_buffer_get_copy_target_list(GtkTextBuffer* v) __attribute__((weak)) {
	goPanic("gtk_text_buffer_get_copy_target_list: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkTargetList* gtk_text_buffer_get_paste_target_list(GtkTextBuffer* v) __attribute__((weak)) {
	goPanic("gtk_text_buffer_get_paste_target_list: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkTreeViewGridLines gtk_tree_view_get_grid_lines(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_grid_lines: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_assistant_get_nth_page(GtkAssistant* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_assistant_get_nth_page: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_assistant_new(void) __attribute__((weak)) {
	goPanic("gtk_assistant_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_link_button_new(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_link_button_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_link_button_new_with_label(const gchar* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_link_button_new_with_label: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_recent_chooser_menu_new(void) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_menu_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_recent_chooser_menu_new_for_manager(GtkRecentManager* _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_menu_new_for_manager: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_recent_chooser_widget_new(void) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_widget_new: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_recent_chooser_widget_new_for_manager(GtkRecentManager* _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_widget_new_for_manager: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
GtkWindowGroup* gtk_window_get_group(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_group: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
PangoContext* gtk_print_context_create_pango_context(GtkPrintContext* v) __attribute__((weak)) {
	goPanic("gtk_print_context_create_pango_context: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
PangoFontMap* gtk_print_context_get_pango_fontmap(GtkPrintContext* v) __attribute__((weak)) {
	goPanic("gtk_print_context_get_pango_fontmap: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
PangoLayout* gtk_print_context_create_pango_layout(GtkPrintContext* v) __attribute__((weak)) {
	goPanic("gtk_print_context_create_pango_layout: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
PangoWrapMode gtk_label_get_line_wrap_mode(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_line_wrap_mode: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
cairo_t* gtk_print_context_get_cairo_context(GtkPrintContext* v) __attribute__((weak)) {
	goPanic("gtk_print_context_get_cairo_context: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const GtkBorder* gtk_entry_get_inner_border(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_inner_border: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_assistant_get_page_title(GtkAssistant* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_assistant_get_page_title: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_combo_box_get_title(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_title: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_link_button_get_uri(GtkLinkButton* v) __attribute__((weak)) {
	goPanic("gtk_link_button_get_uri: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_paper_size_get_default(void) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_default: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_paper_size_get_display_name(GtkPaperSize* v) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_display_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_paper_size_get_name(GtkPaperSize* v) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_paper_size_get_ppd_name(GtkPaperSize* v) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_ppd_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_print_operation_get_status_string(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_get_status_string: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_print_settings_get(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_get: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_print_settings_get_default_source(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_default_source: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_print_settings_get_dither(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_dither: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_print_settings_get_finishings(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_finishings: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_print_settings_get_media_type(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_media_type: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_print_settings_get_output_bin(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_output_bin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_print_settings_get_printer(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_printer: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_recent_filter_get_name(GtkRecentFilter* v) __attribute__((weak)) {
	goPanic("gtk_recent_filter_get_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_recent_info_get_description(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_description: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_recent_info_get_display_name(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_display_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_recent_info_get_mime_type(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_mime_type: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_recent_info_get_uri(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_uri: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_status_icon_get_icon_name(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_icon_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
const gchar* gtk_status_icon_get_stock(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_stock: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_assistant_get_page_complete(GtkAssistant* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_assistant_get_page_complete: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_clipboard_wait_is_rich_text_available(GtkClipboard* v, GtkTextBuffer* _0) __attribute__((weak)) {
	goPanic("gtk_clipboard_wait_is_rich_text_available: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_drag_dest_get_track_motion(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_drag_dest_get_track_motion: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_file_chooser_button_get_focus_on_click(GtkFileChooserButton* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_button_get_focus_on_click: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_notebook_get_tab_detachable(GtkNotebook* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_notebook_get_tab_detachable: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_notebook_get_tab_reorderable(GtkNotebook* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_notebook_get_tab_reorderable: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_paper_size_is_equal(GtkPaperSize* v, GtkPaperSize* _0) __attribute__((weak)) {
	goPanic("gtk_paper_size_is_equal: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_print_operation_is_finished(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_is_finished: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_print_operation_preview_is_selected(GtkPrintOperationPreview* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_preview_is_selected: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_print_settings_get_bool(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_bool: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_print_settings_get_collate(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_collate: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_print_settings_get_reverse(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_reverse: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_print_settings_get_use_color(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_use_color: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_print_settings_has_key(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_has_key: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_chooser_get_local_only(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_local_only: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_chooser_get_select_multiple(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_select_multiple: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_chooser_get_show_icons(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_show_icons: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_chooser_get_show_not_found(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_show_not_found: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_chooser_get_show_private(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_show_private: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_chooser_get_show_tips(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_show_tips: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_chooser_menu_get_show_numbers(GtkRecentChooserMenu* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_menu_get_show_numbers: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_chooser_select_uri(GtkRecentChooser* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_select_uri: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_chooser_set_current_uri(GtkRecentChooser* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_current_uri: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_filter_filter(GtkRecentFilter* v, const GtkRecentFilterInfo* _0) __attribute__((weak)) {
	goPanic("gtk_recent_filter_filter: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_info_exists(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_exists: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_info_get_application_info(GtkRecentInfo* v, const gchar* _0, const gchar** _1, guint* _2, time_t* _3) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_application_info: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_info_get_private_hint(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_private_hint: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_info_has_application(GtkRecentInfo* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_info_has_application: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_info_has_group(GtkRecentInfo* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_info_has_group: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_info_is_local(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_is_local: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_info_match(GtkRecentInfo* v, GtkRecentInfo* _0) __attribute__((weak)) {
	goPanic("gtk_recent_info_match: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_manager_add_full(GtkRecentManager* v, const gchar* _0, const GtkRecentData* _1) __attribute__((weak)) {
	goPanic("gtk_recent_manager_add_full: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_manager_add_item(GtkRecentManager* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_manager_add_item: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_manager_has_item(GtkRecentManager* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_manager_has_item: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_manager_move_item(GtkRecentManager* v, const gchar* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_recent_manager_move_item: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_recent_manager_remove_item(GtkRecentManager* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_manager_remove_item: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_selection_data_targets_include_rich_text(const GtkSelectionData* v, GtkTextBuffer* _0) __attribute__((weak)) {
	goPanic("gtk_selection_data_targets_include_rich_text: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_selection_data_targets_include_uri(const GtkSelectionData* v) __attribute__((weak)) {
	goPanic("gtk_selection_data_targets_include_uri: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_status_icon_get_geometry(GtkStatusIcon* v, GdkScreen** _0, GdkRectangle* _1, GtkOrientation* _2) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_geometry: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_status_icon_get_visible(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_visible: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_status_icon_is_embedded(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_is_embedded: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_style_lookup_color(GtkStyle* v, const gchar* _0, GdkColor* _1) __attribute__((weak)) {
	goPanic("gtk_style_lookup_color: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_text_buffer_get_has_selection(GtkTextBuffer* v) __attribute__((weak)) {
	goPanic("gtk_text_buffer_get_has_selection: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_tree_view_get_enable_tree_lines(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_enable_tree_lines: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_tree_view_get_headers_clickable(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_headers_clickable: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_tree_view_get_rubber_banding(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_rubber_banding: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_widget_is_composited(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_is_composited: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gboolean gtk_window_get_deletable(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_deletable: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gchar* gtk_recent_chooser_get_current_uri(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_current_uri: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gchar* gtk_recent_info_get_short_name(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_short_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gchar* gtk_recent_info_get_uri_display(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_uri_display: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gchar* gtk_recent_info_last_application(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_last_application: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gchar** gtk_recent_chooser_get_uris(GtkRecentChooser* v, gsize* _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_uris: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gchar** gtk_recent_info_get_applications(GtkRecentInfo* v, gsize* _0) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_applications: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gchar** gtk_recent_info_get_groups(GtkRecentInfo* v, gsize* _0) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_groups: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_page_setup_get_bottom_margin(GtkPageSetup* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_get_bottom_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_page_setup_get_left_margin(GtkPageSetup* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_get_left_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_page_setup_get_page_height(GtkPageSetup* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_get_page_height: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_page_setup_get_page_width(GtkPageSetup* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_get_page_width: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_page_setup_get_paper_height(GtkPageSetup* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_get_paper_height: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_page_setup_get_paper_width(GtkPageSetup* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_get_paper_width: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_page_setup_get_right_margin(GtkPageSetup* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_get_right_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_page_setup_get_top_margin(GtkPageSetup* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_get_top_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_paper_size_get_default_bottom_margin(GtkPaperSize* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_default_bottom_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_paper_size_get_default_left_margin(GtkPaperSize* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_default_left_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_paper_size_get_default_right_margin(GtkPaperSize* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_default_right_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_paper_size_get_default_top_margin(GtkPaperSize* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_default_top_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_paper_size_get_height(GtkPaperSize* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_height: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_paper_size_get_width(GtkPaperSize* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_width: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_print_context_get_dpi_x(GtkPrintContext* v) __attribute__((weak)) {
	goPanic("gtk_print_context_get_dpi_x: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_print_context_get_dpi_y(GtkPrintContext* v) __attribute__((weak)) {
	goPanic("gtk_print_context_get_dpi_y: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_print_context_get_height(GtkPrintContext* v) __attribute__((weak)) {
	goPanic("gtk_print_context_get_height: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_print_context_get_width(GtkPrintContext* v) __attribute__((weak)) {
	goPanic("gtk_print_context_get_width: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_print_settings_get_double(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_double: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_print_settings_get_double_with_default(GtkPrintSettings* v, const gchar* _0, gdouble _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_double_with_default: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_print_settings_get_length(GtkPrintSettings* v, const gchar* _0, GtkUnit _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_length: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_print_settings_get_paper_height(GtkPrintSettings* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_paper_height: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_print_settings_get_paper_width(GtkPrintSettings* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_paper_width: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gdouble gtk_print_settings_get_scale(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_scale: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_assistant_append_page(GtkAssistant* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_assistant_append_page: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_assistant_get_current_page(GtkAssistant* v) __attribute__((weak)) {
	goPanic("gtk_assistant_get_current_page: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_assistant_get_n_pages(GtkAssistant* v) __attribute__((weak)) {
	goPanic("gtk_assistant_get_n_pages: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_assistant_insert_page(GtkAssistant* v, GtkWidget* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_assistant_insert_page: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_assistant_prepend_page(GtkAssistant* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_assistant_prepend_page: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_print_settings_get_int(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_int: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_print_settings_get_int_with_default(GtkPrintSettings* v, const gchar* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_int_with_default: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_print_settings_get_n_copies(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_n_copies: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_print_settings_get_number_up(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_number_up: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_print_settings_get_resolution(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_resolution: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_recent_chooser_get_limit(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_get_limit: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_recent_info_get_age(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_age: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_recent_manager_purge_items(GtkRecentManager* v) __attribute__((weak)) {
	goPanic("gtk_recent_manager_purge_items: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
gint gtk_status_icon_get_size(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_size: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
time_t gtk_recent_info_get_added(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_added: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
time_t gtk_recent_info_get_modified(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_modified: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
time_t gtk_recent_info_get_visited(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_visited: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_assistant_add_action_widget(GtkAssistant* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_assistant_add_action_widget: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_assistant_remove_action_widget(GtkAssistant* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_assistant_remove_action_widget: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_assistant_set_current_page(GtkAssistant* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_assistant_set_current_page: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_assistant_set_forward_page_func(GtkAssistant* v, GtkAssistantPageFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_assistant_set_forward_page_func: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_assistant_set_page_complete(GtkAssistant* v, GtkWidget* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_assistant_set_page_complete: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_assistant_set_page_header_image(GtkAssistant* v, GtkWidget* _0, GdkPixbuf* _1) __attribute__((weak)) {
	goPanic("gtk_assistant_set_page_header_image: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_assistant_set_page_side_image(GtkAssistant* v, GtkWidget* _0, GdkPixbuf* _1) __attribute__((weak)) {
	goPanic("gtk_assistant_set_page_side_image: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_assistant_set_page_title(GtkAssistant* v, GtkWidget* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_assistant_set_page_title: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_assistant_set_page_type(GtkAssistant* v, GtkWidget* _0, GtkAssistantPageType _1) __attribute__((weak)) {
	goPanic("gtk_assistant_set_page_type: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_assistant_update_buttons_state(GtkAssistant* v) __attribute__((weak)) {
	goPanic("gtk_assistant_update_buttons_state: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_button_set_image_position(GtkButton* v, GtkPositionType _0) __attribute__((weak)) {
	goPanic("gtk_button_set_image_position: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_combo_box_set_title(GtkComboBox* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_title: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_drag_dest_set_track_motion(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_drag_dest_set_track_motion: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_entry_set_inner_border(GtkEntry* v, const GtkBorder* _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_inner_border: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_file_chooser_button_set_focus_on_click(GtkFileChooserButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_button_set_focus_on_click: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_label_set_line_wrap_mode(GtkLabel* v, PangoWrapMode _0) __attribute__((weak)) {
	goPanic("gtk_label_set_line_wrap_mode: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_link_button_set_uri(GtkLinkButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_link_button_set_uri: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_message_dialog_set_image(GtkMessageDialog* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_message_dialog_set_image: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_notebook_set_tab_detachable(GtkNotebook* v, GtkWidget* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_notebook_set_tab_detachable: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_notebook_set_tab_reorderable(GtkNotebook* v, GtkWidget* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_notebook_set_tab_reorderable: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_page_setup_set_bottom_margin(GtkPageSetup* v, gdouble _0, GtkUnit _1) __attribute__((weak)) {
	goPanic("gtk_page_setup_set_bottom_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_page_setup_set_left_margin(GtkPageSetup* v, gdouble _0, GtkUnit _1) __attribute__((weak)) {
	goPanic("gtk_page_setup_set_left_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_page_setup_set_orientation(GtkPageSetup* v, GtkPageOrientation _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_set_orientation: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_page_setup_set_paper_size(GtkPageSetup* v, GtkPaperSize* _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_set_paper_size: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_page_setup_set_paper_size_and_default_margins(GtkPageSetup* v, GtkPaperSize* _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_set_paper_size_and_default_margins: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_page_setup_set_right_margin(GtkPageSetup* v, gdouble _0, GtkUnit _1) __attribute__((weak)) {
	goPanic("gtk_page_setup_set_right_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_page_setup_set_top_margin(GtkPageSetup* v, gdouble _0, GtkUnit _1) __attribute__((weak)) {
	goPanic("gtk_page_setup_set_top_margin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_paper_size_set_size(GtkPaperSize* v, gdouble _0, gdouble _1, GtkUnit _2) __attribute__((weak)) {
	goPanic("gtk_paper_size_set_size: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_context_set_cairo_context(GtkPrintContext* v, cairo_t* _0, double _1, double _2) __attribute__((weak)) {
	goPanic("gtk_print_context_set_cairo_context: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_cancel(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_cancel: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_get_error(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_get_error: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_preview_end_preview(GtkPrintOperationPreview* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_preview_end_preview: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_preview_render_page(GtkPrintOperationPreview* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_preview_render_page: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_allow_async(GtkPrintOperation* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_allow_async: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_current_page(GtkPrintOperation* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_current_page: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_custom_tab_label(GtkPrintOperation* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_custom_tab_label: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_default_page_setup(GtkPrintOperation* v, GtkPageSetup* _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_default_page_setup: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_export_filename(GtkPrintOperation* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_export_filename: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_job_name(GtkPrintOperation* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_job_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_n_pages(GtkPrintOperation* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_n_pages: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_print_settings(GtkPrintOperation* v, GtkPrintSettings* _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_print_settings: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_show_progress(GtkPrintOperation* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_show_progress: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_track_print_status(GtkPrintOperation* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_track_print_status: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_unit(GtkPrintOperation* v, GtkUnit _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_unit: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_operation_set_use_full_page(GtkPrintOperation* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_use_full_page: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_run_page_setup_dialog_async(GtkWindow* _0, GtkPageSetup* _1, GtkPrintSettings* _2, GtkPageSetupDoneFunc _3, gpointer _4) __attribute__((weak)) {
	goPanic("gtk_print_run_page_setup_dialog_async: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_foreach(GtkPrintSettings* v, GtkPrintSettingsFunc _0, gpointer _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_foreach: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set(GtkPrintSettings* v, const gchar* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_set: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_bool(GtkPrintSettings* v, const gchar* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_bool: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_collate(GtkPrintSettings* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_collate: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_default_source(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_default_source: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_dither(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_dither: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_double(GtkPrintSettings* v, const gchar* _0, gdouble _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_double: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_duplex(GtkPrintSettings* v, GtkPrintDuplex _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_duplex: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_finishings(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_finishings: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_int(GtkPrintSettings* v, const gchar* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_int: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_length(GtkPrintSettings* v, const gchar* _0, gdouble _1, GtkUnit _2) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_length: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_media_type(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_media_type: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_n_copies(GtkPrintSettings* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_n_copies: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_number_up(GtkPrintSettings* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_number_up: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_orientation(GtkPrintSettings* v, GtkPageOrientation _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_orientation: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_output_bin(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_output_bin: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_page_ranges(GtkPrintSettings* v, GtkPageRange* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_page_ranges: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_page_set(GtkPrintSettings* v, GtkPageSet _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_page_set: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_paper_height(GtkPrintSettings* v, gdouble _0, GtkUnit _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_paper_height: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_paper_size(GtkPrintSettings* v, GtkPaperSize* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_paper_size: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_paper_width(GtkPrintSettings* v, gdouble _0, GtkUnit _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_paper_width: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_print_pages(GtkPrintSettings* v, GtkPrintPages _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_print_pages: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_printer(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_printer: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_quality(GtkPrintSettings* v, GtkPrintQuality _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_quality: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_resolution(GtkPrintSettings* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_resolution: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_reverse(GtkPrintSettings* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_reverse: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_scale(GtkPrintSettings* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_scale: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_set_use_color(GtkPrintSettings* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_use_color: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_print_settings_unset(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_unset: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_radio_action_set_current_value(GtkRadioAction* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_radio_action_set_current_value: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_range_set_lower_stepper_sensitivity(GtkRange* v, GtkSensitivityType _0) __attribute__((weak)) {
	goPanic("gtk_range_set_lower_stepper_sensitivity: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_range_set_upper_stepper_sensitivity(GtkRange* v, GtkSensitivityType _0) __attribute__((weak)) {
	goPanic("gtk_range_set_upper_stepper_sensitivity: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_add_filter(GtkRecentChooser* v, GtkRecentFilter* _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_add_filter: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_menu_set_show_numbers(GtkRecentChooserMenu* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_menu_set_show_numbers: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_remove_filter(GtkRecentChooser* v, GtkRecentFilter* _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_remove_filter: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_select_all(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_select_all: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_set_filter(GtkRecentChooser* v, GtkRecentFilter* _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_filter: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_set_limit(GtkRecentChooser* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_limit: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_set_local_only(GtkRecentChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_local_only: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_set_select_multiple(GtkRecentChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_select_multiple: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_set_show_icons(GtkRecentChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_show_icons: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_set_show_not_found(GtkRecentChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_show_not_found: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_set_show_private(GtkRecentChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_show_private: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_set_show_tips(GtkRecentChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_show_tips: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_set_sort_func(GtkRecentChooser* v, GtkRecentSortFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_sort_func: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_set_sort_type(GtkRecentChooser* v, GtkRecentSortType _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_set_sort_type: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_unselect_all(GtkRecentChooser* v) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_unselect_all: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_chooser_unselect_uri(GtkRecentChooser* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_chooser_unselect_uri: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_filter_add_age(GtkRecentFilter* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_recent_filter_add_age: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_filter_add_application(GtkRecentFilter* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_filter_add_application: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_filter_add_custom(GtkRecentFilter* v, GtkRecentFilterFlags _0, GtkRecentFilterFunc _1, gpointer _2, GDestroyNotify _3) __attribute__((weak)) {
	goPanic("gtk_recent_filter_add_custom: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_filter_add_group(GtkRecentFilter* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_filter_add_group: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_filter_add_mime_type(GtkRecentFilter* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_filter_add_mime_type: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_filter_add_pattern(GtkRecentFilter* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_filter_add_pattern: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_filter_add_pixbuf_formats(GtkRecentFilter* v) __attribute__((weak)) {
	goPanic("gtk_recent_filter_add_pixbuf_formats: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_recent_filter_set_name(GtkRecentFilter* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_recent_filter_set_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_scrolled_window_unset_placement(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_unset_placement: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_status_icon_set_from_file(GtkStatusIcon* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_from_file: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_status_icon_set_from_icon_name(GtkStatusIcon* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_from_icon_name: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_status_icon_set_from_pixbuf(GtkStatusIcon* v, GdkPixbuf* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_from_pixbuf: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_status_icon_set_from_stock(GtkStatusIcon* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_from_stock: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_status_icon_set_visible(GtkStatusIcon* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_visible: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_target_list_add_rich_text_targets(GtkTargetList* v, guint _0, gboolean _1, GtkTextBuffer* _2) __attribute__((weak)) {
	goPanic("gtk_target_list_add_rich_text_targets: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_tree_store_insert_with_valuesv(GtkTreeStore* v, GtkTreeIter* _0, GtkTreeIter* _1, gint _2, gint* _3, GValue* _4, gint _5) __attribute__((weak)) {
	goPanic("gtk_tree_store_insert_with_valuesv: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_tree_view_set_enable_tree_lines(GtkTreeView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_enable_tree_lines: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_tree_view_set_grid_lines(GtkTreeView* v, GtkTreeViewGridLines _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_grid_lines: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_tree_view_set_rubber_banding(GtkTreeView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_rubber_banding: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_tree_view_set_search_entry(GtkTreeView* v, GtkEntry* _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_search_entry: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_tree_view_set_search_position_func(GtkTreeView* v, GtkTreeViewSearchPositionFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_search_position_func: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 10))
void gtk_window_set_deletable(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_deletable: library too old: needs at least 2.10");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GList* gtk_cell_layout_get_cells(GtkCellLayout* v) __attribute__((weak)) {
	goPanic("gtk_cell_layout_get_cells: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GList* gtk_icon_theme_list_contexts(GtkIconTheme* v) __attribute__((weak)) {
	goPanic("gtk_icon_theme_list_contexts: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GList* gtk_paper_size_get_paper_sizes(gboolean _0) __attribute__((weak)) {
	goPanic("gtk_paper_size_get_paper_sizes: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GObject* gtk_buildable_construct_child(GtkBuildable* v, GtkBuilder* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_buildable_construct_child: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GObject* gtk_buildable_get_internal_child(GtkBuildable* v, GtkBuilder* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_buildable_get_internal_child: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GObject* gtk_builder_get_object(GtkBuilder* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_builder_get_object: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GSList* gtk_builder_get_objects(GtkBuilder* v) __attribute__((weak)) {
	goPanic("gtk_builder_get_objects: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GType gtk_builder_get_type_from_name(GtkBuilder* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_builder_get_type_from_name: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GdkScreen* gtk_status_icon_get_screen(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_screen: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkAction* gtk_recent_action_new(const gchar* _0, const gchar* _1, const gchar* _2, const gchar* _3) __attribute__((weak)) {
	goPanic("gtk_recent_action_new: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkAction* gtk_recent_action_new_for_manager(const gchar* _0, const gchar* _1, const gchar* _2, const gchar* _3, GtkRecentManager* _4) __attribute__((weak)) {
	goPanic("gtk_recent_action_new_for_manager: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkAdjustment* gtk_entry_get_cursor_hadjustment(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_cursor_hadjustment: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkAdjustment* gtk_scale_button_get_adjustment(GtkScaleButton* v) __attribute__((weak)) {
	goPanic("gtk_scale_button_get_adjustment: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkBuilder* gtk_builder_new(void) __attribute__((weak)) {
	goPanic("gtk_builder_new: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkIconInfo* gtk_icon_theme_choose_icon(GtkIconTheme* v, const gchar** _0, gint _1, GtkIconLookupFlags _2) __attribute__((weak)) {
	goPanic("gtk_icon_theme_choose_icon: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkPageSetup* gtk_page_setup_new_from_file(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_new_from_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkPageSetup* gtk_page_setup_new_from_key_file(GKeyFile* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_page_setup_new_from_key_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkPaperSize* gtk_paper_size_new_from_key_file(GKeyFile* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_paper_size_new_from_key_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkPrintSettings* gtk_print_settings_new_from_file(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_new_from_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkPrintSettings* gtk_print_settings_new_from_key_file(GKeyFile* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_new_from_key_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkTextMark* gtk_text_mark_new(const gchar* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_text_mark_new: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_action_create_menu(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_create_menu: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_scale_button_new(GtkIconSize _0, gdouble _1, gdouble _2, gdouble _3, const gchar** _4) __attribute__((weak)) {
	goPanic("gtk_scale_button_new: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_tree_view_column_get_tree_view(GtkTreeViewColumn* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_column_get_tree_view: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_volume_button_new(void) __attribute__((weak)) {
	goPanic("gtk_volume_button_new: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
GtkWindow* gtk_widget_get_tooltip_window(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_tooltip_window: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
const gchar* gtk_about_dialog_get_program_name(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_program_name: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
const gchar* gtk_buildable_get_name(GtkBuildable* v) __attribute__((weak)) {
	goPanic("gtk_buildable_get_name: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
const gchar* gtk_builder_get_translation_domain(GtkBuilder* v) __attribute__((weak)) {
	goPanic("gtk_builder_get_translation_domain: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
const gchar* gtk_entry_completion_get_completion_prefix(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_get_completion_prefix: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_buildable_custom_tag_start(GtkBuildable* v, GtkBuilder* _0, GObject* _1, const gchar* _2, GMarkupParser* _3, gpointer* _4) __attribute__((weak)) {
	goPanic("gtk_buildable_custom_tag_start: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_builder_value_from_string_type(GtkBuilder* v, GType _0, const gchar* _1, GValue* _2) __attribute__((weak)) {
	goPanic("gtk_builder_value_from_string_type: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_entry_completion_get_inline_selection(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_get_inline_selection: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_page_setup_to_file(GtkPageSetup* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_to_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_print_settings_to_file(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_to_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_range_get_restrict_to_fill_level(GtkRange* v) __attribute__((weak)) {
	goPanic("gtk_range_get_restrict_to_fill_level: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_range_get_show_fill_level(GtkRange* v) __attribute__((weak)) {
	goPanic("gtk_range_get_show_fill_level: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_recent_action_get_show_numbers(GtkRecentAction* v) __attribute__((weak)) {
	goPanic("gtk_recent_action_get_show_numbers: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_tree_view_get_show_expanders(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_show_expanders: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_tree_view_is_rubber_banding_active(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_is_rubber_banding_active: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_widget_get_has_tooltip(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_has_tooltip: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gboolean gtk_widget_keynav_failed(GtkWidget* v, GtkDirectionType _0) __attribute__((weak)) {
	goPanic("gtk_widget_keynav_failed: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gchar* gtk_widget_get_tooltip_markup(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_tooltip_markup: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gchar* gtk_widget_get_tooltip_text(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_tooltip_text: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gdouble gtk_range_get_fill_level(GtkRange* v) __attribute__((weak)) {
	goPanic("gtk_range_get_fill_level: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gdouble gtk_scale_button_get_value(GtkScaleButton* v) __attribute__((weak)) {
	goPanic("gtk_scale_button_get_value: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gdouble gtk_window_get_opacity(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_opacity: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gint gtk_icon_view_get_tooltip_column(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_tooltip_column: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gint gtk_tree_view_get_level_indentation(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_level_indentation: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
gint gtk_tree_view_get_tooltip_column(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_tooltip_column: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
guint gtk_builder_add_from_file(GtkBuilder* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_builder_add_from_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
guint gtk_builder_add_from_string(GtkBuilder* v, const gchar* _0, gsize _1) __attribute__((weak)) {
	goPanic("gtk_builder_add_from_string: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
guint gtk_rc_parse_color_full(GScanner* _0, GtkRcStyle* _1, GdkColor* _2) __attribute__((weak)) {
	goPanic("gtk_rc_parse_color_full: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_about_dialog_set_program_name(GtkAboutDialog* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_program_name: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_binding_entry_skip(GtkBindingSet* _0, guint _1, GdkModifierType _2) __attribute__((weak)) {
	goPanic("gtk_binding_entry_skip: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_buildable_add_child(GtkBuildable* v, GtkBuilder* _0, GObject* _1, const gchar* _2) __attribute__((weak)) {
	goPanic("gtk_buildable_add_child: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_buildable_custom_finished(GtkBuildable* v, GtkBuilder* _0, GObject* _1, const gchar* _2, gpointer _3) __attribute__((weak)) {
	goPanic("gtk_buildable_custom_finished: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_buildable_custom_tag_end(GtkBuildable* v, GtkBuilder* _0, GObject* _1, const gchar* _2, gpointer* _3) __attribute__((weak)) {
	goPanic("gtk_buildable_custom_tag_end: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_buildable_parser_finished(GtkBuildable* v, GtkBuilder* _0) __attribute__((weak)) {
	goPanic("gtk_buildable_parser_finished: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_buildable_set_buildable_property(GtkBuildable* v, GtkBuilder* _0, const gchar* _1, const GValue* _2) __attribute__((weak)) {
	goPanic("gtk_buildable_set_buildable_property: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_buildable_set_name(GtkBuildable* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_buildable_set_name: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_builder_connect_signals(GtkBuilder* v, gpointer _0) __attribute__((weak)) {
	goPanic("gtk_builder_connect_signals: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_builder_set_translation_domain(GtkBuilder* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_builder_set_translation_domain: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_entry_completion_set_inline_selection(GtkEntryCompletion* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_set_inline_selection: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_entry_set_cursor_hadjustment(GtkEntry* v, GtkAdjustment* _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_cursor_hadjustment: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_icon_view_convert_widget_to_bin_window_coords(GtkIconView* v, gint _0, gint _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_icon_view_convert_widget_to_bin_window_coords: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_icon_view_set_tooltip_cell(GtkIconView* v, GtkTooltip* _0, GtkTreePath* _1, GtkCellRenderer* _2) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_tooltip_cell: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_icon_view_set_tooltip_column(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_tooltip_column: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_icon_view_set_tooltip_item(GtkIconView* v, GtkTooltip* _0, GtkTreePath* _1) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_tooltip_item: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_list_store_set_valuesv(GtkListStore* v, GtkTreeIter* _0, gint* _1, GValue* _2, gint _3) __attribute__((weak)) {
	goPanic("gtk_list_store_set_valuesv: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_menu_tool_button_set_arrow_tooltip_markup(GtkMenuToolButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_menu_tool_button_set_arrow_tooltip_markup: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_menu_tool_button_set_arrow_tooltip_text(GtkMenuToolButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_menu_tool_button_set_arrow_tooltip_text: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_page_setup_to_key_file(GtkPageSetup* v, GKeyFile* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_page_setup_to_key_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_paper_size_to_key_file(GtkPaperSize* v, GKeyFile* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_paper_size_to_key_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_print_settings_to_key_file(GtkPrintSettings* v, GKeyFile* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_to_key_file: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_range_set_fill_level(GtkRange* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_range_set_fill_level: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_range_set_restrict_to_fill_level(GtkRange* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_range_set_restrict_to_fill_level: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_range_set_show_fill_level(GtkRange* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_range_set_show_fill_level: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_recent_action_set_show_numbers(GtkRecentAction* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_recent_action_set_show_numbers: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_scale_button_set_adjustment(GtkScaleButton* v, GtkAdjustment* _0) __attribute__((weak)) {
	goPanic("gtk_scale_button_set_adjustment: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_scale_button_set_icons(GtkScaleButton* v, const gchar** _0) __attribute__((weak)) {
	goPanic("gtk_scale_button_set_icons: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_scale_button_set_value(GtkScaleButton* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_scale_button_set_value: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_status_icon_set_screen(GtkStatusIcon* v, GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_screen: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_text_buffer_add_mark(GtkTextBuffer* v, GtkTextMark* _0, const GtkTextIter* _1) __attribute__((weak)) {
	goPanic("gtk_text_buffer_add_mark: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tool_item_set_tooltip_markup(GtkToolItem* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_set_tooltip_markup: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tool_item_set_tooltip_text(GtkToolItem* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_set_tooltip_text: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tooltip_set_custom(GtkTooltip* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_tooltip_set_custom: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tooltip_set_icon(GtkTooltip* v, GdkPixbuf* _0) __attribute__((weak)) {
	goPanic("gtk_tooltip_set_icon: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tooltip_set_icon_from_stock(GtkTooltip* v, const gchar* _0, GtkIconSize _1) __attribute__((weak)) {
	goPanic("gtk_tooltip_set_icon_from_stock: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tooltip_set_markup(GtkTooltip* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tooltip_set_markup: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tooltip_set_text(GtkTooltip* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tooltip_set_text: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tooltip_set_tip_area(GtkTooltip* v, const GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gtk_tooltip_set_tip_area: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tooltip_trigger_tooltip_query(GdkDisplay* _0) __attribute__((weak)) {
	goPanic("gtk_tooltip_trigger_tooltip_query: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_store_set_valuesv(GtkTreeStore* v, GtkTreeIter* _0, gint* _1, GValue* _2, gint _3) __attribute__((weak)) {
	goPanic("gtk_tree_store_set_valuesv: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_convert_bin_window_to_tree_coords(GtkTreeView* v, gint _0, gint _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_tree_view_convert_bin_window_to_tree_coords: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_convert_bin_window_to_widget_coords(GtkTreeView* v, gint _0, gint _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_tree_view_convert_bin_window_to_widget_coords: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_convert_tree_to_bin_window_coords(GtkTreeView* v, gint _0, gint _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_tree_view_convert_tree_to_bin_window_coords: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_convert_tree_to_widget_coords(GtkTreeView* v, gint _0, gint _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_tree_view_convert_tree_to_widget_coords: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_convert_widget_to_bin_window_coords(GtkTreeView* v, gint _0, gint _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_tree_view_convert_widget_to_bin_window_coords: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_convert_widget_to_tree_coords(GtkTreeView* v, gint _0, gint _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_tree_view_convert_widget_to_tree_coords: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_set_level_indentation(GtkTreeView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_level_indentation: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_set_show_expanders(GtkTreeView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_show_expanders: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_set_tooltip_cell(GtkTreeView* v, GtkTooltip* _0, GtkTreePath* _1, GtkTreeViewColumn* _2, GtkCellRenderer* _3) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_tooltip_cell: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_set_tooltip_column(GtkTreeView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_tooltip_column: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_tree_view_set_tooltip_row(GtkTreeView* v, GtkTooltip* _0, GtkTreePath* _1) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_tooltip_row: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_widget_error_bell(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_error_bell: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_widget_modify_cursor(GtkWidget* v, const GdkColor* _0, const GdkColor* _1) __attribute__((weak)) {
	goPanic("gtk_widget_modify_cursor: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_widget_set_has_tooltip(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_has_tooltip: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_widget_set_tooltip_markup(GtkWidget* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_tooltip_markup: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_widget_set_tooltip_text(GtkWidget* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_tooltip_text: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_widget_set_tooltip_window(GtkWidget* v, GtkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_tooltip_window: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_widget_trigger_tooltip_query(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_trigger_tooltip_query: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_window_set_opacity(GtkWindow* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_window_set_opacity: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 12))
void gtk_window_set_startup_id(GtkWindow* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_window_set_startup_id: library too old: needs at least 2.12");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GFile* gtk_file_chooser_get_current_folder_file(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_current_folder_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GFile* gtk_file_chooser_get_file(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GFile* gtk_file_chooser_get_preview_file(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_preview_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GIcon* gtk_status_icon_get_gicon(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_gicon: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GList* gtk_window_group_list_windows(GtkWindowGroup* v) __attribute__((weak)) {
	goPanic("gtk_window_group_list_windows: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GMountOperation* gtk_mount_operation_new(GtkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_mount_operation_new: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GSList* gtk_file_chooser_get_files(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_files: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GdkDisplay* gtk_selection_data_get_display(const GtkSelectionData* v) __attribute__((weak)) {
	goPanic("gtk_selection_data_get_display: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GdkModifierType gtk_accel_group_get_modifier_mask(GtkAccelGroup* v) __attribute__((weak)) {
	goPanic("gtk_accel_group_get_modifier_mask: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GdkScreen* gtk_mount_operation_get_screen(GtkMountOperation* v) __attribute__((weak)) {
	goPanic("gtk_mount_operation_get_screen: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GdkWindow* gtk_layout_get_bin_window(GtkLayout* v) __attribute__((weak)) {
	goPanic("gtk_layout_get_bin_window: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GdkWindow* gtk_plug_get_socket_window(GtkPlug* v) __attribute__((weak)) {
	goPanic("gtk_plug_get_socket_window: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GdkWindow* gtk_socket_get_plug_window(GtkSocket* v) __attribute__((weak)) {
	goPanic("gtk_socket_get_plug_window: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GdkWindow* gtk_widget_get_window(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_window: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkBorder* gtk_border_new(void) __attribute__((weak)) {
	goPanic("gtk_border_new: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkIconInfo* gtk_icon_info_new_for_pixbuf(GtkIconTheme* _0, GdkPixbuf* _1) __attribute__((weak)) {
	goPanic("gtk_icon_info_new_for_pixbuf: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkIconInfo* gtk_icon_theme_lookup_by_gicon(GtkIconTheme* v, GIcon* _0, gint _1, GtkIconLookupFlags _2) __attribute__((weak)) {
	goPanic("gtk_icon_theme_lookup_by_gicon: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkIconSize gtk_tool_shell_get_icon_size(GtkToolShell* v) __attribute__((weak)) {
	goPanic("gtk_tool_shell_get_icon_size: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkNumberUpLayout gtk_print_settings_get_number_up_layout(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_number_up_layout: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkOrientation gtk_tool_shell_get_orientation(GtkToolShell* v) __attribute__((weak)) {
	goPanic("gtk_tool_shell_get_orientation: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkReliefStyle gtk_tool_shell_get_relief_style(GtkToolShell* v) __attribute__((weak)) {
	goPanic("gtk_tool_shell_get_relief_style: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkSensitivityType gtk_combo_box_get_button_sensitivity(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_button_sensitivity: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkStatusIcon* gtk_status_icon_new_from_gicon(GIcon* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_new_from_gicon: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkToolbarStyle gtk_tool_shell_get_style(GtkToolShell* v) __attribute__((weak)) {
	goPanic("gtk_tool_shell_get_style: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_color_selection_dialog_get_color_selection(GtkColorSelectionDialog* v) __attribute__((weak)) {
	goPanic("gtk_color_selection_dialog_get_color_selection: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_container_get_focus_child(GtkContainer* v) __attribute__((weak)) {
	goPanic("gtk_container_get_focus_child: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_dialog_get_action_area(GtkDialog* v) __attribute__((weak)) {
	goPanic("gtk_dialog_get_action_area: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_dialog_get_content_area(GtkDialog* v) __attribute__((weak)) {
	goPanic("gtk_dialog_get_content_area: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_font_selection_dialog_get_cancel_button(GtkFontSelectionDialog* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_dialog_get_cancel_button: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_font_selection_dialog_get_ok_button(GtkFontSelectionDialog* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_dialog_get_ok_button: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_font_selection_get_face_list(GtkFontSelection* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_get_face_list: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_font_selection_get_family_list(GtkFontSelection* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_get_family_list: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_font_selection_get_preview_entry(GtkFontSelection* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_get_preview_entry: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_font_selection_get_size_entry(GtkFontSelection* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_get_size_entry: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_font_selection_get_size_list(GtkFontSelection* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_get_size_list: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_hsv_new(void) __attribute__((weak)) {
	goPanic("gtk_hsv_new: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_image_new_from_gicon(GIcon* _0, GtkIconSize _1) __attribute__((weak)) {
	goPanic("gtk_image_new_from_gicon: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_message_dialog_get_image(GtkMessageDialog* v) __attribute__((weak)) {
	goPanic("gtk_message_dialog_get_image: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_scale_button_get_minus_button(GtkScaleButton* v) __attribute__((weak)) {
	goPanic("gtk_scale_button_get_minus_button: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_scale_button_get_plus_button(GtkScaleButton* v) __attribute__((weak)) {
	goPanic("gtk_scale_button_get_plus_button: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_scale_button_get_popup(GtkScaleButton* v) __attribute__((weak)) {
	goPanic("gtk_scale_button_get_popup: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_test_create_simple_window(const gchar* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_test_create_simple_window: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_test_find_label(GtkWidget* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_test_find_label: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_test_find_sibling(GtkWidget* _0, GType _1) __attribute__((weak)) {
	goPanic("gtk_test_find_sibling: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_test_find_widget(GtkWidget* _0, const gchar* _1, GType _2) __attribute__((weak)) {
	goPanic("gtk_test_find_widget: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_window_get_default_widget(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_default_widget: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
GtkWindow* gtk_mount_operation_get_parent(GtkMountOperation* v) __attribute__((weak)) {
	goPanic("gtk_mount_operation_get_parent: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
PangoFontFace* gtk_font_selection_get_face(GtkFontSelection* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_get_face: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
PangoFontFamily* gtk_font_selection_get_family(GtkFontSelection* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_get_family: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
const GType* gtk_test_list_all_types(guint* _0) __attribute__((weak)) {
	goPanic("gtk_test_list_all_types: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
const gchar* gtk_menu_get_accel_path(GtkMenu* v) __attribute__((weak)) {
	goPanic("gtk_menu_get_accel_path: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
const gchar* gtk_menu_item_get_accel_path(GtkMenuItem* v) __attribute__((weak)) {
	goPanic("gtk_menu_item_get_accel_path: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
double gtk_test_slider_get_value(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_test_slider_get_value: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_accel_group_get_is_locked(GtkAccelGroup* v) __attribute__((weak)) {
	goPanic("gtk_accel_group_get_is_locked: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_clipboard_wait_is_uris_available(GtkClipboard* v) __attribute__((weak)) {
	goPanic("gtk_clipboard_wait_is_uris_available: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_entry_get_overwrite_mode(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_overwrite_mode: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_file_chooser_select_file(GtkFileChooser* v, GFile* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_select_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_file_chooser_set_current_folder_file(GtkFileChooser* v, GFile* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_current_folder_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_file_chooser_set_file(GtkFileChooser* v, GFile* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_handle_box_get_child_detached(GtkHandleBox* v) __attribute__((weak)) {
	goPanic("gtk_handle_box_get_child_detached: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_hsv_is_adjusting(GtkHSV* v) __attribute__((weak)) {
	goPanic("gtk_hsv_is_adjusting: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_link_button_get_visited(GtkLinkButton* v) __attribute__((weak)) {
	goPanic("gtk_link_button_get_visited: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_mount_operation_is_showing(GtkMountOperation* v) __attribute__((weak)) {
	goPanic("gtk_mount_operation_is_showing: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_page_setup_load_file(GtkPageSetup* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_load_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_page_setup_load_key_file(GtkPageSetup* v, GKeyFile* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_page_setup_load_key_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_plug_get_embedded(GtkPlug* v) __attribute__((weak)) {
	goPanic("gtk_plug_get_embedded: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_print_settings_load_file(GtkPrintSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_load_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_print_settings_load_key_file(GtkPrintSettings* v, GKeyFile* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_load_key_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_show_uri(GdkScreen* _0, const gchar* _1, guint32 _2) __attribute__((weak)) {
	goPanic("gtk_show_uri: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_test_spin_button_click(GtkSpinButton* _0, guint _1, gboolean _2) __attribute__((weak)) {
	goPanic("gtk_test_spin_button_click: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_test_widget_click(GtkWidget* _0, guint _1, GdkModifierType _2) __attribute__((weak)) {
	goPanic("gtk_test_widget_click: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gboolean gtk_test_widget_send_key(GtkWidget* _0, guint _1, GdkModifierType _2) __attribute__((weak)) {
	goPanic("gtk_test_widget_send_key: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gchar* gtk_test_text_get(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_test_text_get: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gchar** gtk_clipboard_wait_for_uris(GtkClipboard* v) __attribute__((weak)) {
	goPanic("gtk_clipboard_wait_for_uris: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gdouble gtk_adjustment_get_lower(GtkAdjustment* v) __attribute__((weak)) {
	goPanic("gtk_adjustment_get_lower: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gdouble gtk_adjustment_get_page_increment(GtkAdjustment* v) __attribute__((weak)) {
	goPanic("gtk_adjustment_get_page_increment: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gdouble gtk_adjustment_get_page_size(GtkAdjustment* v) __attribute__((weak)) {
	goPanic("gtk_adjustment_get_page_size: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gdouble gtk_adjustment_get_step_increment(GtkAdjustment* v) __attribute__((weak)) {
	goPanic("gtk_adjustment_get_step_increment: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gdouble gtk_adjustment_get_upper(GtkAdjustment* v) __attribute__((weak)) {
	goPanic("gtk_adjustment_get_upper: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gint gtk_calendar_get_detail_height_rows(GtkCalendar* v) __attribute__((weak)) {
	goPanic("gtk_calendar_get_detail_height_rows: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gint gtk_calendar_get_detail_width_chars(GtkCalendar* v) __attribute__((weak)) {
	goPanic("gtk_calendar_get_detail_width_chars: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gint gtk_font_selection_get_size(GtkFontSelection* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_get_size: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gint gtk_menu_get_monitor(GtkMenu* v) __attribute__((weak)) {
	goPanic("gtk_menu_get_monitor: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gint gtk_selection_data_get_format(const GtkSelectionData* v) __attribute__((weak)) {
	goPanic("gtk_selection_data_get_format: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
gint gtk_selection_data_get_length(const GtkSelectionData* v) __attribute__((weak)) {
	goPanic("gtk_selection_data_get_length: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
guint gtk_builder_add_objects_from_file(GtkBuilder* v, const gchar* _0, gchar** _1) __attribute__((weak)) {
	goPanic("gtk_builder_add_objects_from_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
guint gtk_builder_add_objects_from_string(GtkBuilder* v, const gchar* _0, gsize _1, gchar** _2) __attribute__((weak)) {
	goPanic("gtk_builder_add_objects_from_string: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
guint16 gtk_entry_get_text_length(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_text_length: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
guint32 gtk_status_icon_get_x11_window_id(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_x11_window_id: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_adjustment_configure(GtkAdjustment* v, gdouble _0, gdouble _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_adjustment_configure: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_adjustment_set_lower(GtkAdjustment* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_adjustment_set_lower: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_adjustment_set_page_increment(GtkAdjustment* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_adjustment_set_page_increment: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_adjustment_set_page_size(GtkAdjustment* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_adjustment_set_page_size: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_adjustment_set_step_increment(GtkAdjustment* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_adjustment_set_step_increment: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_adjustment_set_upper(GtkAdjustment* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_adjustment_set_upper: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_calendar_set_detail_func(GtkCalendar* v, GtkCalendarDetailFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_calendar_set_detail_func: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_calendar_set_detail_height_rows(GtkCalendar* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_calendar_set_detail_height_rows: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_calendar_set_detail_width_chars(GtkCalendar* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_calendar_set_detail_width_chars: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_clipboard_request_uris(GtkClipboard* v, GtkClipboardURIReceivedFunc _0, gpointer _1) __attribute__((weak)) {
	goPanic("gtk_clipboard_request_uris: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_combo_box_set_button_sensitivity(GtkComboBox* v, GtkSensitivityType _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_button_sensitivity: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_entry_set_overwrite_mode(GtkEntry* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_overwrite_mode: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_file_chooser_unselect_file(GtkFileChooser* v, GFile* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_unselect_file: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_hsv_get_color(GtkHSV* v, gdouble* _0, gdouble* _1, gdouble* _2) __attribute__((weak)) {
	goPanic("gtk_hsv_get_color: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_hsv_get_metrics(GtkHSV* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_hsv_get_metrics: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_hsv_set_color(GtkHSV* v, double _0, double _1, double _2) __attribute__((weak)) {
	goPanic("gtk_hsv_set_color: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_hsv_set_metrics(GtkHSV* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_hsv_set_metrics: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_hsv_to_rgb(gdouble _0, gdouble _1, gdouble _2, gdouble* _3, gdouble* _4, gdouble* _5) __attribute__((weak)) {
	goPanic("gtk_hsv_to_rgb: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_image_get_gicon(GtkImage* v, GIcon** _0, GtkIconSize* _1) __attribute__((weak)) {
	goPanic("gtk_image_get_gicon: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_image_set_from_gicon(GtkImage* v, GIcon* _0, GtkIconSize _1) __attribute__((weak)) {
	goPanic("gtk_image_set_from_gicon: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_link_button_set_visited(GtkLinkButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_link_button_set_visited: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_mount_operation_set_parent(GtkMountOperation* v, GtkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_mount_operation_set_parent: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_mount_operation_set_screen(GtkMountOperation* v, GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_mount_operation_set_screen: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_print_settings_set_number_up_layout(GtkPrintSettings* v, GtkNumberUpLayout _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_number_up_layout: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_rgb_to_hsv(gdouble _0, gdouble _1, gdouble _2, gdouble* _3, gdouble* _4, gdouble* _5) __attribute__((weak)) {
	goPanic("gtk_rgb_to_hsv: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_status_icon_set_from_gicon(GtkStatusIcon* v, GIcon* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_from_gicon: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_test_register_all_types(void) __attribute__((weak)) {
	goPanic("gtk_test_register_all_types: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_test_slider_set_perc(GtkWidget* _0, double _1) __attribute__((weak)) {
	goPanic("gtk_test_slider_set_perc: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_test_text_set(GtkWidget* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_test_text_set: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_tool_item_toolbar_reconfigured(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_toolbar_reconfigured: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_tool_shell_rebuild_menu(GtkToolShell* v) __attribute__((weak)) {
	goPanic("gtk_tool_shell_rebuild_menu: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 14))
void gtk_tooltip_set_icon_from_icon_name(GtkTooltip* v, const gchar* _0, GtkIconSize _1) __attribute__((weak)) {
	goPanic("gtk_tooltip_set_icon_from_icon_name: library too old: needs at least 2.14");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
GIcon* gtk_action_get_gicon(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_gicon: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
GIcon* gtk_entry_get_icon_gicon(GtkEntry* v, GtkEntryIconPosition _0) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_gicon: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
GdkPixbuf* gtk_entry_get_icon_pixbuf(GtkEntry* v, GtkEntryIconPosition _0) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_pixbuf: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
GtkAction* gtk_activatable_get_related_action(GtkActivatable* v) __attribute__((weak)) {
	goPanic("gtk_activatable_get_related_action: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
GtkImageType gtk_entry_get_icon_storage_type(GtkEntry* v, GtkEntryIconPosition _0) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_storage_type: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
GtkOrientation gtk_orientable_get_orientation(GtkOrientable* v) __attribute__((weak)) {
	goPanic("gtk_orientable_get_orientation: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
GtkTreeModel* gtk_cell_view_get_model(GtkCellView* v) __attribute__((weak)) {
	goPanic("gtk_cell_view_get_model: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
const char* gtk_im_multicontext_get_context_id(GtkIMMulticontext* v) __attribute__((weak)) {
	goPanic("gtk_im_multicontext_get_context_id: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
const gchar* gtk_action_get_icon_name(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_icon_name: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
const gchar* gtk_action_get_label(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_label: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
const gchar* gtk_action_get_short_label(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_short_label: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
const gchar* gtk_action_get_stock_id(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_stock_id: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
const gchar* gtk_action_get_tooltip(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_tooltip: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
const gchar* gtk_entry_get_icon_name(GtkEntry* v, GtkEntryIconPosition _0) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_name: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
const gchar* gtk_entry_get_icon_stock(GtkEntry* v, GtkEntryIconPosition _0) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_stock: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
const gchar* gtk_menu_item_get_label(GtkMenuItem* v) __attribute__((weak)) {
	goPanic("gtk_menu_item_get_label: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
const gchar* gtk_window_get_default_icon_name(void) __attribute__((weak)) {
	goPanic("gtk_window_get_default_icon_name: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gboolean gtk_action_get_is_important(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_is_important: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gboolean gtk_action_get_visible_horizontal(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_visible_horizontal: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gboolean gtk_action_get_visible_vertical(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_visible_vertical: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gboolean gtk_activatable_get_use_action_appearance(GtkActivatable* v) __attribute__((weak)) {
	goPanic("gtk_activatable_get_use_action_appearance: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gboolean gtk_entry_get_icon_activatable(GtkEntry* v, GtkEntryIconPosition _0) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_activatable: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gboolean gtk_entry_get_icon_sensitive(GtkEntry* v, GtkEntryIconPosition _0) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_sensitive: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gboolean gtk_image_menu_item_get_always_show_image(GtkImageMenuItem* v) __attribute__((weak)) {
	goPanic("gtk_image_menu_item_get_always_show_image: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gboolean gtk_image_menu_item_get_use_stock(GtkImageMenuItem* v) __attribute__((weak)) {
	goPanic("gtk_image_menu_item_get_use_stock: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gboolean gtk_menu_item_get_use_underline(GtkMenuItem* v) __attribute__((weak)) {
	goPanic("gtk_menu_item_get_use_underline: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gboolean gtk_status_icon_get_has_tooltip(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_has_tooltip: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gchar* gtk_entry_get_icon_tooltip_markup(GtkEntry* v, GtkEntryIconPosition _0) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_tooltip_markup: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gchar* gtk_entry_get_icon_tooltip_text(GtkEntry* v, GtkEntryIconPosition _0) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_tooltip_text: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gchar* gtk_status_icon_get_tooltip_markup(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_tooltip_markup: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gchar* gtk_status_icon_get_tooltip_text(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_tooltip_text: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gdouble gtk_entry_get_progress_fraction(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_progress_fraction: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gdouble gtk_entry_get_progress_pulse_step(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_progress_pulse_step: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gdouble gtk_print_settings_get_printer_lpi(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_printer_lpi: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gint gtk_entry_get_current_icon_drag_source(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_current_icon_drag_source: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gint gtk_entry_get_icon_at_pos(GtkEntry* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_at_pos: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gint gtk_print_settings_get_resolution_x(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_resolution_x: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
gint gtk_print_settings_get_resolution_y(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_get_resolution_y: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_block_activate(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_block_activate: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_set_gicon(GtkAction* v, GIcon* _0) __attribute__((weak)) {
	goPanic("gtk_action_set_gicon: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_set_icon_name(GtkAction* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_action_set_icon_name: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_set_is_important(GtkAction* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_action_set_is_important: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_set_label(GtkAction* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_action_set_label: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_set_short_label(GtkAction* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_action_set_short_label: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_set_stock_id(GtkAction* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_action_set_stock_id: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_set_tooltip(GtkAction* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_action_set_tooltip: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_set_visible_horizontal(GtkAction* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_action_set_visible_horizontal: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_set_visible_vertical(GtkAction* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_action_set_visible_vertical: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_action_unblock_activate(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_unblock_activate: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_activatable_do_set_related_action(GtkActivatable* v, GtkAction* _0) __attribute__((weak)) {
	goPanic("gtk_activatable_do_set_related_action: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_activatable_set_related_action(GtkActivatable* v, GtkAction* _0) __attribute__((weak)) {
	goPanic("gtk_activatable_set_related_action: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_activatable_set_use_action_appearance(GtkActivatable* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_activatable_set_use_action_appearance: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_activatable_sync_action_properties(GtkActivatable* v, GtkAction* _0) __attribute__((weak)) {
	goPanic("gtk_activatable_sync_action_properties: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_progress_pulse(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_progress_pulse: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_icon_activatable(GtkEntry* v, GtkEntryIconPosition _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_entry_set_icon_activatable: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_icon_drag_source(GtkEntry* v, GtkEntryIconPosition _0, GtkTargetList* _1, GdkDragAction _2) __attribute__((weak)) {
	goPanic("gtk_entry_set_icon_drag_source: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_icon_from_gicon(GtkEntry* v, GtkEntryIconPosition _0, GIcon* _1) __attribute__((weak)) {
	goPanic("gtk_entry_set_icon_from_gicon: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_icon_from_icon_name(GtkEntry* v, GtkEntryIconPosition _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_entry_set_icon_from_icon_name: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_icon_from_pixbuf(GtkEntry* v, GtkEntryIconPosition _0, GdkPixbuf* _1) __attribute__((weak)) {
	goPanic("gtk_entry_set_icon_from_pixbuf: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_icon_from_stock(GtkEntry* v, GtkEntryIconPosition _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_entry_set_icon_from_stock: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_icon_sensitive(GtkEntry* v, GtkEntryIconPosition _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_entry_set_icon_sensitive: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_icon_tooltip_markup(GtkEntry* v, GtkEntryIconPosition _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_entry_set_icon_tooltip_markup: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_icon_tooltip_text(GtkEntry* v, GtkEntryIconPosition _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_entry_set_icon_tooltip_text: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_progress_fraction(GtkEntry* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_progress_fraction: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_set_progress_pulse_step(GtkEntry* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_progress_pulse_step: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_entry_unset_invisible_char(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_unset_invisible_char: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_im_multicontext_set_context_id(GtkIMMulticontext* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_im_multicontext_set_context_id: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_image_menu_item_set_accel_group(GtkImageMenuItem* v, GtkAccelGroup* _0) __attribute__((weak)) {
	goPanic("gtk_image_menu_item_set_accel_group: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_image_menu_item_set_always_show_image(GtkImageMenuItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_image_menu_item_set_always_show_image: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_image_menu_item_set_use_stock(GtkImageMenuItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_image_menu_item_set_use_stock: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_menu_item_set_label(GtkMenuItem* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_menu_item_set_label: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_menu_item_set_use_underline(GtkMenuItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_menu_item_set_use_underline: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_orientable_set_orientation(GtkOrientable* v, GtkOrientation _0) __attribute__((weak)) {
	goPanic("gtk_orientable_set_orientation: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_print_operation_draw_page_finish(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_draw_page_finish: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_print_operation_set_defer_drawing(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_defer_drawing: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_print_settings_set_printer_lpi(GtkPrintSettings* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_printer_lpi: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_print_settings_set_resolution_xy(GtkPrintSettings* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_print_settings_set_resolution_xy: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_scale_add_mark(GtkScale* v, gdouble _0, GtkPositionType _1, const gchar* _2) __attribute__((weak)) {
	goPanic("gtk_scale_add_mark: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_scale_clear_marks(GtkScale* v) __attribute__((weak)) {
	goPanic("gtk_scale_clear_marks: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_status_icon_set_has_tooltip(GtkStatusIcon* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_has_tooltip: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_status_icon_set_tooltip_markup(GtkStatusIcon* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_tooltip_markup: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_status_icon_set_tooltip_text(GtkStatusIcon* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_tooltip_text: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 16))
void gtk_style_get_style_property(GtkStyle* v, GType _0, const gchar* _1, GValue* _2) __attribute__((weak)) {
	goPanic("gtk_style_get_style_property: library too old: needs at least 2.16");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
GtkEntryBuffer* gtk_entry_buffer_new(const gchar* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_new: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
GtkEntryBuffer* gtk_entry_get_buffer(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_buffer: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
GtkMessageType gtk_info_bar_get_message_type(GtkInfoBar* v) __attribute__((weak)) {
	goPanic("gtk_info_bar_get_message_type: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
GtkStateType gtk_widget_get_state(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_state: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
GtkWidget* gtk_entry_new_with_buffer(GtkEntryBuffer* _0) __attribute__((weak)) {
	goPanic("gtk_entry_new_with_buffer: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
GtkWidget* gtk_info_bar_add_button(GtkInfoBar* v, const gchar* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_info_bar_add_button: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
GtkWidget* gtk_info_bar_get_action_area(GtkInfoBar* v) __attribute__((weak)) {
	goPanic("gtk_info_bar_get_action_area: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
GtkWidget* gtk_info_bar_get_content_area(GtkInfoBar* v) __attribute__((weak)) {
	goPanic("gtk_info_bar_get_content_area: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
GtkWidget* gtk_info_bar_new(void) __attribute__((weak)) {
	goPanic("gtk_info_bar_new: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
const gchar* gtk_entry_buffer_get_text(GtkEntryBuffer* v) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_get_text: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
const gchar* gtk_label_get_current_uri(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_current_uri: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
const gchar* gtk_status_icon_get_title(GtkStatusIcon* v) __attribute__((weak)) {
	goPanic("gtk_status_icon_get_title: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_cell_renderer_get_sensitive(GtkCellRenderer* v) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_sensitive: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_cell_renderer_get_visible(GtkCellRenderer* v) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_visible: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_cell_renderer_toggle_get_activatable(GtkCellRendererToggle* v) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_toggle_get_activatable: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_file_chooser_get_create_folders(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_create_folders: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_label_get_track_visited_links(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_track_visited_links: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_menu_get_reserve_toggle_size(GtkMenu* v) __attribute__((weak)) {
	goPanic("gtk_menu_get_reserve_toggle_size: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_print_operation_get_embed_page_setup(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_get_embed_page_setup: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_print_operation_get_has_selection(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_get_has_selection: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_print_operation_get_support_selection(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_get_support_selection: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_range_get_flippable(GtkRange* v) __attribute__((weak)) {
	goPanic("gtk_range_get_flippable: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_get_app_paintable(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_app_paintable: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_get_can_default(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_can_default: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_get_can_focus(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_can_focus: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_get_double_buffered(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_double_buffered: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_get_has_window(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_has_window: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_get_receives_default(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_receives_default: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_get_sensitive(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_sensitive: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_get_visible(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_visible: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_has_default(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_has_default: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_has_focus(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_has_focus: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_has_grab(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_has_grab: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_is_drawable(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_is_drawable: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_is_sensitive(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_is_sensitive: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gboolean gtk_widget_is_toplevel(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_is_toplevel: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gint gtk_entry_buffer_get_max_length(GtkEntryBuffer* v) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_get_max_length: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gint gtk_icon_view_get_item_padding(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_item_padding: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gint gtk_print_operation_get_n_pages_to_print(GtkPrintOperation* v) __attribute__((weak)) {
	goPanic("gtk_print_operation_get_n_pages_to_print: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
gsize gtk_entry_buffer_get_bytes(GtkEntryBuffer* v) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_get_bytes: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
guint gtk_entry_buffer_delete_text(GtkEntryBuffer* v, guint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_delete_text: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
guint gtk_entry_buffer_get_length(GtkEntryBuffer* v) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_get_length: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
guint gtk_entry_buffer_insert_text(GtkEntryBuffer* v, guint _0, const gchar* _1, gint _2) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_insert_text: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_cell_renderer_get_alignment(GtkCellRenderer* v, gfloat* _0, gfloat* _1) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_alignment: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_cell_renderer_get_padding(GtkCellRenderer* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_padding: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_cell_renderer_set_alignment(GtkCellRenderer* v, gfloat _0, gfloat _1) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_set_alignment: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_cell_renderer_set_padding(GtkCellRenderer* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_set_padding: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_cell_renderer_set_sensitive(GtkCellRenderer* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_set_sensitive: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_cell_renderer_set_visible(GtkCellRenderer* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_set_visible: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_cell_renderer_toggle_set_activatable(GtkCellRendererToggle* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_toggle_set_activatable: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_entry_buffer_emit_deleted_text(GtkEntryBuffer* v, guint _0, guint _1) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_emit_deleted_text: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_entry_buffer_emit_inserted_text(GtkEntryBuffer* v, guint _0, const gchar* _1, guint _2) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_emit_inserted_text: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_entry_buffer_set_max_length(GtkEntryBuffer* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_set_max_length: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_entry_buffer_set_text(GtkEntryBuffer* v, const gchar* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_entry_buffer_set_text: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_entry_set_buffer(GtkEntry* v, GtkEntryBuffer* _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_buffer: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_file_chooser_set_create_folders(GtkFileChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_create_folders: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_icon_view_set_item_padding(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_item_padding: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_info_bar_add_action_widget(GtkInfoBar* v, GtkWidget* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_info_bar_add_action_widget: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_info_bar_response(GtkInfoBar* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_info_bar_response: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_info_bar_set_default_response(GtkInfoBar* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_info_bar_set_default_response: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_info_bar_set_message_type(GtkInfoBar* v, GtkMessageType _0) __attribute__((weak)) {
	goPanic("gtk_info_bar_set_message_type: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_info_bar_set_response_sensitive(GtkInfoBar* v, gint _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_info_bar_set_response_sensitive: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_label_set_track_visited_links(GtkLabel* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_label_set_track_visited_links: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_menu_set_reserve_toggle_size(GtkMenu* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_menu_set_reserve_toggle_size: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_print_operation_set_embed_page_setup(GtkPrintOperation* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_embed_page_setup: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_print_operation_set_has_selection(GtkPrintOperation* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_has_selection: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_print_operation_set_support_selection(GtkPrintOperation* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_print_operation_set_support_selection: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_range_set_flippable(GtkRange* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_range_set_flippable: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_status_icon_set_title(GtkStatusIcon* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_title: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_widget_get_allocation(GtkWidget* v, GtkAllocation* _0) __attribute__((weak)) {
	goPanic("gtk_widget_get_allocation: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_widget_set_allocation(GtkWidget* v, const GtkAllocation* _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_allocation: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_widget_set_can_default(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_can_default: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_widget_set_can_focus(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_can_focus: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_widget_set_has_window(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_has_window: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_widget_set_receives_default(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_receives_default: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_widget_set_visible(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_visible: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 18))
void gtk_widget_set_window(GtkWidget* v, GdkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_window: library too old: needs at least 2.18");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
GList* gtk_tree_selection_get_selected_rows(GtkTreeSelection* v, GtkTreeModel** _0) __attribute__((weak)) {
	goPanic("gtk_tree_selection_get_selected_rows: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
GdkDisplay* gtk_clipboard_get_display(GtkClipboard* v) __attribute__((weak)) {
	goPanic("gtk_clipboard_get_display: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
GdkDisplay* gtk_widget_get_display(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_display: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
GdkScreen* gtk_invisible_get_screen(GtkInvisible* v) __attribute__((weak)) {
	goPanic("gtk_invisible_get_screen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
GdkScreen* gtk_widget_get_screen(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_screen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
GdkScreen* gtk_window_get_screen(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_screen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
GdkWindow* gtk_widget_get_root_window(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_root_window: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
GtkSettings* gtk_settings_get_for_screen(GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_settings_get_for_screen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
GtkTreeRowReference* gtk_tree_row_reference_copy(GtkTreeRowReference* v) __attribute__((weak)) {
	goPanic("gtk_tree_row_reference_copy: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
GtkWidget* gtk_invisible_new_for_screen(GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_invisible_new_for_screen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gboolean gtk_icon_size_lookup_for_settings(GtkSettings* _0, GtkIconSize _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_icon_size_lookup_for_settings: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gboolean gtk_list_store_iter_is_valid(GtkListStore* v, GtkTreeIter* _0) __attribute__((weak)) {
	goPanic("gtk_list_store_iter_is_valid: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gboolean gtk_tree_model_sort_iter_is_valid(GtkTreeModelSort* v, GtkTreeIter* _0) __attribute__((weak)) {
	goPanic("gtk_tree_model_sort_iter_is_valid: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gboolean gtk_tree_store_iter_is_valid(GtkTreeStore* v, GtkTreeIter* _0) __attribute__((weak)) {
	goPanic("gtk_tree_store_iter_is_valid: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gboolean gtk_widget_has_screen(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_has_screen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gboolean gtk_window_get_skip_pager_hint(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_skip_pager_hint: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gboolean gtk_window_get_skip_taskbar_hint(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_skip_taskbar_hint: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gboolean gtk_window_set_default_icon_from_file(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_window_set_default_icon_from_file: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gboolean gtk_window_set_icon_from_file(GtkWindow* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_window_set_icon_from_file: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gchar* gtk_tree_model_get_string_from_iter(GtkTreeModel* v, GtkTreeIter* _0) __attribute__((weak)) {
	goPanic("gtk_tree_model_get_string_from_iter: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gint gtk_notebook_get_n_pages(GtkNotebook* v) __attribute__((weak)) {
	goPanic("gtk_notebook_get_n_pages: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
gint gtk_tree_selection_count_selected_rows(GtkTreeSelection* v) __attribute__((weak)) {
	goPanic("gtk_tree_selection_count_selected_rows: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_invisible_set_screen(GtkInvisible* v, GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_invisible_set_screen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_list_store_move_after(GtkListStore* v, GtkTreeIter* _0, GtkTreeIter* _1) __attribute__((weak)) {
	goPanic("gtk_list_store_move_after: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_list_store_move_before(GtkListStore* v, GtkTreeIter* _0, GtkTreeIter* _1) __attribute__((weak)) {
	goPanic("gtk_list_store_move_before: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_list_store_reorder(GtkListStore* v, gint* _0) __attribute__((weak)) {
	goPanic("gtk_list_store_reorder: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_list_store_swap(GtkListStore* v, GtkTreeIter* _0, GtkTreeIter* _1) __attribute__((weak)) {
	goPanic("gtk_list_store_swap: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_menu_set_screen(GtkMenu* v, GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_menu_set_screen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_menu_shell_select_first(GtkMenuShell* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_menu_shell_select_first: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_tree_selection_unselect_range(GtkTreeSelection* v, GtkTreePath* _0, GtkTreePath* _1) __attribute__((weak)) {
	goPanic("gtk_tree_selection_unselect_range: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_tree_store_move_after(GtkTreeStore* v, GtkTreeIter* _0, GtkTreeIter* _1) __attribute__((weak)) {
	goPanic("gtk_tree_store_move_after: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_tree_store_move_before(GtkTreeStore* v, GtkTreeIter* _0, GtkTreeIter* _1) __attribute__((weak)) {
	goPanic("gtk_tree_store_move_before: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_tree_store_swap(GtkTreeStore* v, GtkTreeIter* _0, GtkTreeIter* _1) __attribute__((weak)) {
	goPanic("gtk_tree_store_swap: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_tree_view_column_focus_cell(GtkTreeViewColumn* v, GtkCellRenderer* _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_column_focus_cell: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_tree_view_expand_to_path(GtkTreeView* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_expand_to_path: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_tree_view_set_cursor_on_cell(GtkTreeView* v, GtkTreePath* _0, GtkTreeViewColumn* _1, GtkCellRenderer* _2, gboolean _3) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_cursor_on_cell: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_window_fullscreen(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_fullscreen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_window_set_auto_startup_notification(gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_auto_startup_notification: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_window_set_screen(GtkWindow* v, GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_window_set_screen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_window_set_skip_pager_hint(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_skip_pager_hint: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_window_set_skip_taskbar_hint(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_skip_taskbar_hint: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 2))
void gtk_window_unfullscreen(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_unfullscreen: library too old: needs at least 2.2");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GdkPixbuf* gtk_offscreen_window_get_pixbuf(GtkOffscreenWindow* v) __attribute__((weak)) {
	goPanic("gtk_offscreen_window_get_pixbuf: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GdkWindow* gtk_paned_get_handle_window(GtkPaned* v) __attribute__((weak)) {
	goPanic("gtk_paned_get_handle_window: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GdkWindow* gtk_viewport_get_bin_window(GtkViewport* v) __attribute__((weak)) {
	goPanic("gtk_viewport_get_bin_window: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkAdjustment* gtk_tool_palette_get_hadjustment(GtkToolPalette* v) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_hadjustment: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkAdjustment* gtk_tool_palette_get_vadjustment(GtkToolPalette* v) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_vadjustment: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkCellRenderer* gtk_cell_renderer_spinner_new(void) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_spinner_new: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkIconSize gtk_tool_palette_get_icon_size(GtkToolPalette* v) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_icon_size: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkOrientation gtk_tool_item_get_text_orientation(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_text_orientation: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkOrientation gtk_tool_shell_get_text_orientation(GtkToolShell* v) __attribute__((weak)) {
	goPanic("gtk_tool_shell_get_text_orientation: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkReliefStyle gtk_tool_item_group_get_header_relief(GtkToolItemGroup* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_get_header_relief: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkSizeGroup* gtk_tool_item_get_text_size_group(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_text_size_group: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkSizeGroup* gtk_tool_shell_get_text_size_group(GtkToolShell* v) __attribute__((weak)) {
	goPanic("gtk_tool_shell_get_text_size_group: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkToolItem* gtk_tool_item_group_get_drop_item(GtkToolItemGroup* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_get_drop_item: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkToolItem* gtk_tool_item_group_get_nth_item(GtkToolItemGroup* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_get_nth_item: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkToolItem* gtk_tool_palette_get_drop_item(GtkToolPalette* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_drop_item: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkToolItemGroup* gtk_tool_palette_get_drop_group(GtkToolPalette* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_drop_group: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkToolbarStyle gtk_tool_palette_get_style(GtkToolPalette* v) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_style: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkWidget* gtk_dialog_get_widget_for_response(GtkDialog* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_dialog_get_widget_for_response: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkWidget* gtk_notebook_get_action_widget(GtkNotebook* v, GtkPackType _0) __attribute__((weak)) {
	goPanic("gtk_notebook_get_action_widget: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkWidget* gtk_offscreen_window_new(void) __attribute__((weak)) {
	goPanic("gtk_offscreen_window_new: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkWidget* gtk_spinner_new(void) __attribute__((weak)) {
	goPanic("gtk_spinner_new: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkWidget* gtk_statusbar_get_message_area(GtkStatusbar* v) __attribute__((weak)) {
	goPanic("gtk_statusbar_get_message_area: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkWidget* gtk_tool_item_group_get_label_widget(GtkToolItemGroup* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_get_label_widget: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkWidget* gtk_tool_item_group_new(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_new: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkWidget* gtk_tool_palette_get_drag_item(GtkToolPalette* v, const GtkSelectionData* _0) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_drag_item: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkWidget* gtk_tool_palette_new(void) __attribute__((weak)) {
	goPanic("gtk_tool_palette_new: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
GtkWindowType gtk_window_get_window_type(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_window_type: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
PangoEllipsizeMode gtk_tool_item_get_ellipsize_mode(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_ellipsize_mode: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
PangoEllipsizeMode gtk_tool_item_group_get_ellipsize(GtkToolItemGroup* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_get_ellipsize: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
PangoEllipsizeMode gtk_tool_shell_get_ellipsize_mode(GtkToolShell* v) __attribute__((weak)) {
	goPanic("gtk_tool_shell_get_ellipsize_mode: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
cairo_surface_t* gtk_offscreen_window_get_surface(GtkOffscreenWindow* v) __attribute__((weak)) {
	goPanic("gtk_offscreen_window_get_surface: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
const GtkTargetEntry* gtk_tool_palette_get_drag_target_group(void) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_drag_target_group: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
const GtkTargetEntry* gtk_tool_palette_get_drag_target_item(void) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_drag_target_item: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
const gchar* gtk_tool_item_group_get_label(GtkToolItemGroup* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_get_label: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gboolean gtk_action_get_always_show_image(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_always_show_image: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gboolean gtk_print_context_get_hard_margins(GtkPrintContext* v, gdouble* _0, gdouble* _1, gdouble* _2, gdouble* _3) __attribute__((weak)) {
	goPanic("gtk_print_context_get_hard_margins: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gboolean gtk_range_get_slider_size_fixed(GtkRange* v) __attribute__((weak)) {
	goPanic("gtk_range_get_slider_size_fixed: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gboolean gtk_tool_item_group_get_collapsed(GtkToolItemGroup* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_get_collapsed: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gboolean gtk_tool_palette_get_exclusive(GtkToolPalette* v, GtkToolItemGroup* _0) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_exclusive: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gboolean gtk_tool_palette_get_expand(GtkToolPalette* v, GtkToolItemGroup* _0) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_expand: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gboolean gtk_widget_get_mapped(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_mapped: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gboolean gtk_widget_get_realized(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_realized: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gboolean gtk_widget_has_rc_style(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_has_rc_style: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gboolean gtk_window_get_mnemonics_visible(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_mnemonics_visible: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gfloat gtk_tool_item_get_text_alignment(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_text_alignment: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gfloat gtk_tool_shell_get_text_alignment(GtkToolShell* v) __attribute__((weak)) {
	goPanic("gtk_tool_shell_get_text_alignment: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gint gtk_range_get_min_slider_size(GtkRange* v) __attribute__((weak)) {
	goPanic("gtk_range_get_min_slider_size: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gint gtk_tool_item_group_get_item_position(GtkToolItemGroup* v, GtkToolItem* _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_get_item_position: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
gint gtk_tool_palette_get_group_position(GtkToolPalette* v, GtkToolItemGroup* _0) __attribute__((weak)) {
	goPanic("gtk_tool_palette_get_group_position: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
guint gtk_tool_item_group_get_n_items(GtkToolItemGroup* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_get_n_items: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_action_set_always_show_image(GtkAction* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_action_set_always_show_image: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_notebook_set_action_widget(GtkNotebook* v, GtkWidget* _0, GtkPackType _1) __attribute__((weak)) {
	goPanic("gtk_notebook_set_action_widget: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_range_get_range_rect(GtkRange* v, GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gtk_range_get_range_rect: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_range_get_slider_range(GtkRange* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_range_get_slider_range: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_range_set_min_slider_size(GtkRange* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_range_set_min_slider_size: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_range_set_slider_size_fixed(GtkRange* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_range_set_slider_size_fixed: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_spinner_start(GtkSpinner* v) __attribute__((weak)) {
	goPanic("gtk_spinner_start: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_spinner_stop(GtkSpinner* v) __attribute__((weak)) {
	goPanic("gtk_spinner_stop: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_status_icon_set_name(GtkStatusIcon* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_status_icon_set_name: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_item_group_insert(GtkToolItemGroup* v, GtkToolItem* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_insert: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_item_group_set_collapsed(GtkToolItemGroup* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_set_collapsed: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_item_group_set_ellipsize(GtkToolItemGroup* v, PangoEllipsizeMode _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_set_ellipsize: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_item_group_set_header_relief(GtkToolItemGroup* v, GtkReliefStyle _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_set_header_relief: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_item_group_set_item_position(GtkToolItemGroup* v, GtkToolItem* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_set_item_position: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_item_group_set_label(GtkToolItemGroup* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_set_label: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_item_group_set_label_widget(GtkToolItemGroup* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_group_set_label_widget: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_palette_add_drag_dest(GtkToolPalette* v, GtkWidget* _0, GtkDestDefaults _1, GtkToolPaletteDragTargets _2, GdkDragAction _3) __attribute__((weak)) {
	goPanic("gtk_tool_palette_add_drag_dest: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_palette_set_drag_source(GtkToolPalette* v, GtkToolPaletteDragTargets _0) __attribute__((weak)) {
	goPanic("gtk_tool_palette_set_drag_source: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_palette_set_exclusive(GtkToolPalette* v, GtkToolItemGroup* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_tool_palette_set_exclusive: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_palette_set_expand(GtkToolPalette* v, GtkToolItemGroup* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_tool_palette_set_expand: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_palette_set_group_position(GtkToolPalette* v, GtkToolItemGroup* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_tool_palette_set_group_position: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_palette_set_icon_size(GtkToolPalette* v, GtkIconSize _0) __attribute__((weak)) {
	goPanic("gtk_tool_palette_set_icon_size: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_palette_set_style(GtkToolPalette* v, GtkToolbarStyle _0) __attribute__((weak)) {
	goPanic("gtk_tool_palette_set_style: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_palette_unset_icon_size(GtkToolPalette* v) __attribute__((weak)) {
	goPanic("gtk_tool_palette_unset_icon_size: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tool_palette_unset_style(GtkToolPalette* v) __attribute__((weak)) {
	goPanic("gtk_tool_palette_unset_style: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_tooltip_set_icon_from_gicon(GtkTooltip* v, GIcon* _0, GtkIconSize _1) __attribute__((weak)) {
	goPanic("gtk_tooltip_set_icon_from_gicon: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_widget_get_requisition(GtkWidget* v, GtkRequisition* _0) __attribute__((weak)) {
	goPanic("gtk_widget_get_requisition: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_widget_set_mapped(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_mapped: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_widget_set_realized(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_realized: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_widget_style_attach(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_style_attach: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 20))
void gtk_window_set_mnemonics_visible(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_mnemonics_visible: library too old: needs at least 2.20");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
GIcon* gtk_recent_info_get_gicon(GtkRecentInfo* v) __attribute__((weak)) {
	goPanic("gtk_recent_info_get_gicon: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
GdkWindow* gtk_button_get_event_window(GtkButton* v) __attribute__((weak)) {
	goPanic("gtk_button_get_event_window: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
GdkWindow* gtk_viewport_get_view_window(GtkViewport* v) __attribute__((weak)) {
	goPanic("gtk_viewport_get_view_window: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
GtkAdjustment* gtk_text_view_get_hadjustment(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_get_hadjustment: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
GtkAdjustment* gtk_text_view_get_vadjustment(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_get_vadjustment: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
GtkWidget* gtk_accessible_get_widget(GtkAccessible* v) __attribute__((weak)) {
	goPanic("gtk_accessible_get_widget: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
GtkWidget* gtk_font_selection_dialog_get_font_selection(GtkFontSelectionDialog* v) __attribute__((weak)) {
	goPanic("gtk_font_selection_dialog_get_font_selection: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
GtkWidget* gtk_message_dialog_get_message_area(GtkMessageDialog* v) __attribute__((weak)) {
	goPanic("gtk_message_dialog_get_message_area: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
GtkWidget* gtk_window_group_get_current_grab(GtkWindowGroup* v) __attribute__((weak)) {
	goPanic("gtk_window_group_get_current_grab: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
gboolean gtk_entry_im_context_filter_keypress(GtkEntry* v, GdkEventKey* _0) __attribute__((weak)) {
	goPanic("gtk_entry_im_context_filter_keypress: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
gboolean gtk_expander_get_label_fill(GtkExpander* v) __attribute__((weak)) {
	goPanic("gtk_expander_get_label_fill: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
gboolean gtk_text_view_im_context_filter_keypress(GtkTextView* v, GdkEventKey* _0) __attribute__((weak)) {
	goPanic("gtk_text_view_im_context_filter_keypress: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
gint gtk_icon_view_get_item_column(GtkIconView* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_item_column: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
gint gtk_icon_view_get_item_row(GtkIconView* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_item_row: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
guint16 gtk_notebook_get_tab_hborder(GtkNotebook* v) __attribute__((weak)) {
	goPanic("gtk_notebook_get_tab_hborder: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
guint16 gtk_notebook_get_tab_vborder(GtkNotebook* v) __attribute__((weak)) {
	goPanic("gtk_notebook_get_tab_vborder: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
void gtk_accessible_set_widget(GtkAccessible* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_accessible_set_widget: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
void gtk_assistant_commit(GtkAssistant* v) __attribute__((weak)) {
	goPanic("gtk_assistant_commit: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
void gtk_entry_reset_im_context(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_reset_im_context: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
void gtk_expander_set_label_fill(GtkExpander* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_expander_set_label_fill: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
void gtk_statusbar_remove_all(GtkStatusbar* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_statusbar_remove_all: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
void gtk_table_get_size(GtkTable* v, guint* _0, guint* _1) __attribute__((weak)) {
	goPanic("gtk_table_get_size: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 22))
void gtk_text_view_reset_im_context(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_reset_im_context: library too old: needs at least 2.22");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
GtkWidget* gtk_combo_box_new_with_entry(void) __attribute__((weak)) {
	goPanic("gtk_combo_box_new_with_entry: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
GtkWidget* gtk_combo_box_new_with_model_and_entry(GtkTreeModel* _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_new_with_model_and_entry: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
GtkWidget* gtk_combo_box_text_new(void) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_new: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
GtkWidget* gtk_combo_box_text_new_with_entry(void) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_new_with_entry: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
const gchar* gtk_notebook_get_group_name(GtkNotebook* v) __attribute__((weak)) {
	goPanic("gtk_notebook_get_group_name: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
gboolean gtk_combo_box_get_has_entry(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_has_entry: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
gchar* gtk_combo_box_text_get_active_text(GtkComboBoxText* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_get_active_text: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
gint gtk_combo_box_get_entry_text_column(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_entry_text_column: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
gint gtk_range_get_round_digits(GtkRange* v) __attribute__((weak)) {
	goPanic("gtk_range_get_round_digits: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
void gtk_combo_box_set_entry_text_column(GtkComboBox* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_entry_text_column: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
void gtk_combo_box_text_append(GtkComboBoxText* v, const gchar* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_append: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
void gtk_combo_box_text_append_text(GtkComboBoxText* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_append_text: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
void gtk_combo_box_text_insert_text(GtkComboBoxText* v, gint _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_insert_text: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
void gtk_combo_box_text_prepend(GtkComboBoxText* v, const gchar* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_prepend: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
void gtk_combo_box_text_prepend_text(GtkComboBoxText* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_prepend_text: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
void gtk_combo_box_text_remove(GtkComboBoxText* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_remove: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
void gtk_notebook_set_group_name(GtkNotebook* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_notebook_set_group_name: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 24))
void gtk_range_set_round_digits(GtkRange* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_range_set_round_digits: library too old: needs at least 2.24");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GList* gtk_action_group_list_actions(GtkActionGroup* v) __attribute__((weak)) {
	goPanic("gtk_action_group_list_actions: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GList* gtk_icon_theme_list_icons(GtkIconTheme* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_icon_theme_list_icons: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GList* gtk_ui_manager_get_action_groups(GtkUIManager* v) __attribute__((weak)) {
	goPanic("gtk_ui_manager_get_action_groups: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GList* gtk_widget_list_mnemonic_labels(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_list_mnemonic_labels: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GQuark gtk_file_chooser_error_quark(void) __attribute__((weak)) {
	goPanic("gtk_file_chooser_error_quark: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GSList* gtk_action_get_proxies(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_proxies: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GSList* gtk_file_chooser_get_filenames(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_filenames: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GSList* gtk_file_chooser_get_uris(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_uris: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GSList* gtk_file_chooser_list_filters(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_list_filters: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GSList* gtk_file_chooser_list_shortcut_folder_uris(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_list_shortcut_folder_uris: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GSList* gtk_file_chooser_list_shortcut_folders(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_list_shortcut_folders: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GSList* gtk_radio_action_get_group(GtkRadioAction* v) __attribute__((weak)) {
	goPanic("gtk_radio_action_get_group: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GSList* gtk_radio_tool_button_get_group(GtkRadioToolButton* v) __attribute__((weak)) {
	goPanic("gtk_radio_tool_button_get_group: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GSList* gtk_ui_manager_get_toplevels(GtkUIManager* v, GtkUIManagerItemType _0) __attribute__((weak)) {
	goPanic("gtk_ui_manager_get_toplevels: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GdkPixbuf* gtk_icon_info_get_builtin_pixbuf(GtkIconInfo* v) __attribute__((weak)) {
	goPanic("gtk_icon_info_get_builtin_pixbuf: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GdkPixbuf* gtk_icon_info_load_icon(GtkIconInfo* v) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_icon: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GdkPixbuf* gtk_icon_theme_load_icon(GtkIconTheme* v, const gchar* _0, gint _1, GtkIconLookupFlags _2) __attribute__((weak)) {
	goPanic("gtk_icon_theme_load_icon: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkAccelGroup* gtk_ui_manager_get_accel_group(GtkUIManager* v) __attribute__((weak)) {
	goPanic("gtk_ui_manager_get_accel_group: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkAccelMap* gtk_accel_map_get(void) __attribute__((weak)) {
	goPanic("gtk_accel_map_get: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkAction* gtk_action_group_get_action(GtkActionGroup* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_action_group_get_action: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkAction* gtk_action_new(const gchar* _0, const gchar* _1, const gchar* _2, const gchar* _3) __attribute__((weak)) {
	goPanic("gtk_action_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkAction* gtk_ui_manager_get_action(GtkUIManager* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_ui_manager_get_action: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkActionGroup* gtk_action_group_new(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_action_group_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkCalendarDisplayOptions gtk_calendar_get_display_options(GtkCalendar* v) __attribute__((weak)) {
	goPanic("gtk_calendar_get_display_options: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkEntryCompletion* gtk_entry_completion_new(void) __attribute__((weak)) {
	goPanic("gtk_entry_completion_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkEntryCompletion* gtk_entry_get_completion(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_completion: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkFileChooserAction gtk_file_chooser_get_action(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_action: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkFileFilter* gtk_file_chooser_get_filter(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_filter: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkFileFilter* gtk_file_filter_new(void) __attribute__((weak)) {
	goPanic("gtk_file_filter_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkFileFilterFlags gtk_file_filter_get_needed(GtkFileFilter* v) __attribute__((weak)) {
	goPanic("gtk_file_filter_get_needed: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkIconInfo* gtk_icon_theme_lookup_icon(GtkIconTheme* v, const gchar* _0, gint _1, GtkIconLookupFlags _2) __attribute__((weak)) {
	goPanic("gtk_icon_theme_lookup_icon: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkIconSize gtk_tool_item_get_icon_size(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_icon_size: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkIconTheme* gtk_icon_theme_get_default(void) __attribute__((weak)) {
	goPanic("gtk_icon_theme_get_default: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkIconTheme* gtk_icon_theme_get_for_screen(GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_icon_theme_get_for_screen: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkIconTheme* gtk_icon_theme_new(void) __attribute__((weak)) {
	goPanic("gtk_icon_theme_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkOrientation gtk_tool_item_get_orientation(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_orientation: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkRadioAction* gtk_radio_action_new(const gchar* _0, const gchar* _1, const gchar* _2, const gchar* _3, gint _4) __attribute__((weak)) {
	goPanic("gtk_radio_action_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkReliefStyle gtk_tool_item_get_relief_style(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_relief_style: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkReliefStyle gtk_toolbar_get_relief_style(GtkToolbar* v) __attribute__((weak)) {
	goPanic("gtk_toolbar_get_relief_style: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkTargetList* gtk_drag_source_get_target_list(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_drag_source_get_target_list: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToggleAction* gtk_toggle_action_new(const gchar* _0, const gchar* _1, const gchar* _2, const gchar* _3) __attribute__((weak)) {
	goPanic("gtk_toggle_action_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_radio_tool_button_new(GSList* _0) __attribute__((weak)) {
	goPanic("gtk_radio_tool_button_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_radio_tool_button_new_from_stock(GSList* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_radio_tool_button_new_from_stock: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_radio_tool_button_new_from_widget(GtkRadioToolButton* _0) __attribute__((weak)) {
	goPanic("gtk_radio_tool_button_new_from_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_radio_tool_button_new_with_stock_from_widget(GtkRadioToolButton* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_radio_tool_button_new_with_stock_from_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_separator_tool_item_new(void) __attribute__((weak)) {
	goPanic("gtk_separator_tool_item_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_toggle_tool_button_new(void) __attribute__((weak)) {
	goPanic("gtk_toggle_tool_button_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_toggle_tool_button_new_from_stock(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_toggle_tool_button_new_from_stock: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_tool_button_new(GtkWidget* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_tool_button_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_tool_button_new_from_stock(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tool_button_new_from_stock: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_tool_item_new(void) __attribute__((weak)) {
	goPanic("gtk_tool_item_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolItem* gtk_toolbar_get_nth_item(GtkToolbar* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_toolbar_get_nth_item: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkToolbarStyle gtk_tool_item_get_toolbar_style(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_toolbar_style: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkTreeModel* gtk_combo_box_get_model(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_model: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkTreeModel* gtk_entry_completion_get_model(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_get_model: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkTreeModel* gtk_tree_model_filter_get_model(GtkTreeModelFilter* v) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_get_model: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkTreeModel* gtk_tree_model_filter_new(GtkTreeModel* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkTreePath* gtk_tree_model_filter_convert_child_path_to_path(GtkTreeModelFilter* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_convert_child_path_to_path: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkTreePath* gtk_tree_model_filter_convert_path_to_child_path(GtkTreeModelFilter* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_convert_path_to_child_path: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkUIManager* gtk_ui_manager_new(void) __attribute__((weak)) {
	goPanic("gtk_ui_manager_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_action_create_icon(GtkAction* v, GtkIconSize _0) __attribute__((weak)) {
	goPanic("gtk_action_create_icon: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_action_create_menu_item(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_create_menu_item: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_action_create_tool_item(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_create_tool_item: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_color_button_new(void) __attribute__((weak)) {
	goPanic("gtk_color_button_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_color_button_new_with_color(const GdkColor* _0) __attribute__((weak)) {
	goPanic("gtk_color_button_new_with_color: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_combo_box_new(void) __attribute__((weak)) {
	goPanic("gtk_combo_box_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_combo_box_new_with_model(GtkTreeModel* _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_new_with_model: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_entry_completion_get_entry(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_get_entry: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_expander_get_label_widget(GtkExpander* v) __attribute__((weak)) {
	goPanic("gtk_expander_get_label_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_expander_new(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_expander_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_expander_new_with_mnemonic(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_expander_new_with_mnemonic: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_file_chooser_get_extra_widget(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_extra_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_file_chooser_get_preview_widget(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_preview_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_file_chooser_widget_new(GtkFileChooserAction _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_widget_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_font_button_new(void) __attribute__((weak)) {
	goPanic("gtk_font_button_new: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_font_button_new_with_font(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_font_button_new_with_font: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_paned_get_child1(GtkPaned* v) __attribute__((weak)) {
	goPanic("gtk_paned_get_child1: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_paned_get_child2(GtkPaned* v) __attribute__((weak)) {
	goPanic("gtk_paned_get_child2: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_radio_menu_item_new_from_widget(GtkRadioMenuItem* _0) __attribute__((weak)) {
	goPanic("gtk_radio_menu_item_new_from_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_radio_menu_item_new_with_label_from_widget(GtkRadioMenuItem* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_radio_menu_item_new_with_label_from_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_radio_menu_item_new_with_mnemonic_from_widget(GtkRadioMenuItem* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_radio_menu_item_new_with_mnemonic_from_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_tool_button_get_icon_widget(GtkToolButton* v) __attribute__((weak)) {
	goPanic("gtk_tool_button_get_icon_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_tool_button_get_label_widget(GtkToolButton* v) __attribute__((weak)) {
	goPanic("gtk_tool_button_get_label_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_tool_item_get_proxy_menu_item(GtkToolItem* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_proxy_menu_item: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_tool_item_retrieve_proxy_menu_item(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_retrieve_proxy_menu_item: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_ui_manager_get_widget(GtkUIManager* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_ui_manager_get_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
PangoLayout* gtk_scale_get_layout(GtkScale* v) __attribute__((weak)) {
	goPanic("gtk_scale_get_layout: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
char* gtk_file_chooser_get_preview_filename(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_preview_filename: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
char* gtk_file_chooser_get_preview_uri(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_preview_uri: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
char* gtk_icon_theme_get_example_icon_name(GtkIconTheme* v) __attribute__((weak)) {
	goPanic("gtk_icon_theme_get_example_icon_name: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_action_get_name(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_name: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_action_group_get_name(GtkActionGroup* v) __attribute__((weak)) {
	goPanic("gtk_action_group_get_name: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_color_button_get_title(GtkColorButton* v) __attribute__((weak)) {
	goPanic("gtk_color_button_get_title: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_expander_get_label(GtkExpander* v) __attribute__((weak)) {
	goPanic("gtk_expander_get_label: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_file_filter_get_name(GtkFileFilter* v) __attribute__((weak)) {
	goPanic("gtk_file_filter_get_name: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_font_button_get_font_name(GtkFontButton* v) __attribute__((weak)) {
	goPanic("gtk_font_button_get_font_name: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_font_button_get_title(GtkFontButton* v) __attribute__((weak)) {
	goPanic("gtk_font_button_get_title: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_icon_info_get_display_name(GtkIconInfo* v) __attribute__((weak)) {
	goPanic("gtk_icon_info_get_display_name: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_icon_info_get_filename(GtkIconInfo* v) __attribute__((weak)) {
	goPanic("gtk_icon_info_get_filename: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_tool_button_get_label(GtkToolButton* v) __attribute__((weak)) {
	goPanic("gtk_tool_button_get_label: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
const gchar* gtk_tool_button_get_stock_id(GtkToolButton* v) __attribute__((weak)) {
	goPanic("gtk_tool_button_get_stock_id: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_action_get_sensitive(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_sensitive: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_action_get_visible(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_visible: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_action_group_get_sensitive(GtkActionGroup* v) __attribute__((weak)) {
	goPanic("gtk_action_group_get_sensitive: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_action_group_get_visible(GtkActionGroup* v) __attribute__((weak)) {
	goPanic("gtk_action_group_get_visible: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_action_is_sensitive(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_is_sensitive: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_action_is_visible(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_is_visible: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_bindings_activate_event(GObject* _0, GdkEventKey* _1) __attribute__((weak)) {
	goPanic("gtk_bindings_activate_event: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_button_box_get_child_secondary(GtkButtonBox* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_button_box_get_child_secondary: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_button_get_focus_on_click(GtkButton* v) __attribute__((weak)) {
	goPanic("gtk_button_get_focus_on_click: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_check_menu_item_get_draw_as_radio(GtkCheckMenuItem* v) __attribute__((weak)) {
	goPanic("gtk_check_menu_item_get_draw_as_radio: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_color_button_get_use_alpha(GtkColorButton* v) __attribute__((weak)) {
	goPanic("gtk_color_button_get_use_alpha: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_combo_box_get_active_iter(GtkComboBox* v, GtkTreeIter* _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_active_iter: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_event_box_get_above_child(GtkEventBox* v) __attribute__((weak)) {
	goPanic("gtk_event_box_get_above_child: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_event_box_get_visible_window(GtkEventBox* v) __attribute__((weak)) {
	goPanic("gtk_event_box_get_visible_window: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_expander_get_expanded(GtkExpander* v) __attribute__((weak)) {
	goPanic("gtk_expander_get_expanded: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_expander_get_use_markup(GtkExpander* v) __attribute__((weak)) {
	goPanic("gtk_expander_get_use_markup: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_expander_get_use_underline(GtkExpander* v) __attribute__((weak)) {
	goPanic("gtk_expander_get_use_underline: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_add_shortcut_folder(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_add_shortcut_folder: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_add_shortcut_folder_uri(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_add_shortcut_folder_uri: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_get_local_only(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_local_only: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_get_preview_widget_active(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_preview_widget_active: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_get_select_multiple(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_select_multiple: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_remove_shortcut_folder(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_remove_shortcut_folder: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_remove_shortcut_folder_uri(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_remove_shortcut_folder_uri: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_select_filename(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_select_filename: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_select_uri(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_select_uri: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_set_current_folder(GtkFileChooser* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_current_folder: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_set_current_folder_uri(GtkFileChooser* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_current_folder_uri: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_set_filename(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_filename: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_chooser_set_uri(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_uri: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_file_filter_filter(GtkFileFilter* v, const GtkFileFilterInfo* _0) __attribute__((weak)) {
	goPanic("gtk_file_filter_filter: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_font_button_get_show_size(GtkFontButton* v) __attribute__((weak)) {
	goPanic("gtk_font_button_get_show_size: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_font_button_get_show_style(GtkFontButton* v) __attribute__((weak)) {
	goPanic("gtk_font_button_get_show_style: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_font_button_get_use_font(GtkFontButton* v) __attribute__((weak)) {
	goPanic("gtk_font_button_get_use_font: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_font_button_get_use_size(GtkFontButton* v) __attribute__((weak)) {
	goPanic("gtk_font_button_get_use_size: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_font_button_set_font_name(GtkFontButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_font_button_set_font_name: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_icon_info_get_attach_points(GtkIconInfo* v, GdkPoint** _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_icon_info_get_attach_points: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_icon_info_get_embedded_rect(GtkIconInfo* v, GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gtk_icon_info_get_embedded_rect: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_icon_theme_has_icon(GtkIconTheme* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_icon_theme_has_icon: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_icon_theme_rescan_if_needed(GtkIconTheme* v) __attribute__((weak)) {
	goPanic("gtk_icon_theme_rescan_if_needed: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_separator_tool_item_get_draw(GtkSeparatorToolItem* v) __attribute__((weak)) {
	goPanic("gtk_separator_tool_item_get_draw: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_text_iter_backward_visible_cursor_position(GtkTextIter* v) __attribute__((weak)) {
	goPanic("gtk_text_iter_backward_visible_cursor_position: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_text_iter_backward_visible_cursor_positions(GtkTextIter* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_text_iter_backward_visible_cursor_positions: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_text_iter_backward_visible_word_start(GtkTextIter* v) __attribute__((weak)) {
	goPanic("gtk_text_iter_backward_visible_word_start: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_text_iter_backward_visible_word_starts(GtkTextIter* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_text_iter_backward_visible_word_starts: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_text_iter_forward_visible_cursor_position(GtkTextIter* v) __attribute__((weak)) {
	goPanic("gtk_text_iter_forward_visible_cursor_position: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_text_iter_forward_visible_cursor_positions(GtkTextIter* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_text_iter_forward_visible_cursor_positions: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_text_iter_forward_visible_word_end(GtkTextIter* v) __attribute__((weak)) {
	goPanic("gtk_text_iter_forward_visible_word_end: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_text_iter_forward_visible_word_ends(GtkTextIter* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_text_iter_forward_visible_word_ends: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_text_view_get_accepts_tab(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_get_accepts_tab: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_text_view_get_overwrite(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_get_overwrite: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_toggle_action_get_active(GtkToggleAction* v) __attribute__((weak)) {
	goPanic("gtk_toggle_action_get_active: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_toggle_action_get_draw_as_radio(GtkToggleAction* v) __attribute__((weak)) {
	goPanic("gtk_toggle_action_get_draw_as_radio: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_toggle_tool_button_get_active(GtkToggleToolButton* v) __attribute__((weak)) {
	goPanic("gtk_toggle_tool_button_get_active: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_tool_button_get_use_underline(GtkToolButton* v) __attribute__((weak)) {
	goPanic("gtk_tool_button_get_use_underline: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_tool_item_get_expand(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_expand: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_tool_item_get_homogeneous(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_homogeneous: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_tool_item_get_is_important(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_is_important: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_tool_item_get_use_drag_window(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_use_drag_window: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_tool_item_get_visible_horizontal(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_visible_horizontal: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_tool_item_get_visible_vertical(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_get_visible_vertical: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_toolbar_get_show_arrow(GtkToolbar* v) __attribute__((weak)) {
	goPanic("gtk_toolbar_get_show_arrow: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_tree_model_filter_convert_child_iter_to_iter(GtkTreeModelFilter* v, GtkTreeIter* _0, GtkTreeIter* _1) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_convert_child_iter_to_iter: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_tree_view_column_get_expand(GtkTreeViewColumn* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_column_get_expand: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_ui_manager_get_add_tearoffs(GtkUIManager* v) __attribute__((weak)) {
	goPanic("gtk_ui_manager_get_add_tearoffs: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_widget_can_activate_accel(GtkWidget* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_widget_can_activate_accel: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_widget_get_no_show_all(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_no_show_all: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_window_activate_key(GtkWindow* v, GdkEventKey* _0) __attribute__((weak)) {
	goPanic("gtk_window_activate_key: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_window_get_accept_focus(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_accept_focus: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_window_has_toplevel_focus(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_has_toplevel_focus: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_window_is_active(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_is_active: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gboolean gtk_window_propagate_key_event(GtkWindow* v, GdkEventKey* _0) __attribute__((weak)) {
	goPanic("gtk_window_propagate_key_event: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gchar* gtk_file_chooser_get_current_folder(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_current_folder: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gchar* gtk_file_chooser_get_current_folder_uri(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_current_folder_uri: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gchar* gtk_file_chooser_get_filename(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_filename: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gchar* gtk_file_chooser_get_uri(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_uri: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gchar* gtk_ui_manager_get_ui(GtkUIManager* v) __attribute__((weak)) {
	goPanic("gtk_ui_manager_get_ui: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gfloat gtk_entry_get_alignment(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_alignment: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gint gtk_combo_box_get_active(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_active: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gint gtk_entry_completion_get_minimum_key_length(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_get_minimum_key_length: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gint gtk_expander_get_spacing(GtkExpander* v) __attribute__((weak)) {
	goPanic("gtk_expander_get_spacing: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gint gtk_icon_info_get_base_size(GtkIconInfo* v) __attribute__((weak)) {
	goPanic("gtk_icon_info_get_base_size: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gint gtk_radio_action_get_current_value(GtkRadioAction* v) __attribute__((weak)) {
	goPanic("gtk_radio_action_get_current_value: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gint gtk_toolbar_get_drop_index(GtkToolbar* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_toolbar_get_drop_index: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gint gtk_toolbar_get_item_index(GtkToolbar* v, GtkToolItem* _0) __attribute__((weak)) {
	goPanic("gtk_toolbar_get_item_index: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
gint gtk_toolbar_get_n_items(GtkToolbar* v) __attribute__((weak)) {
	goPanic("gtk_toolbar_get_n_items: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
guint gtk_ui_manager_add_ui_from_file(GtkUIManager* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_ui_manager_add_ui_from_file: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
guint gtk_ui_manager_add_ui_from_string(GtkUIManager* v, const gchar* _0, gssize _1) __attribute__((weak)) {
	goPanic("gtk_ui_manager_add_ui_from_string: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
guint gtk_ui_manager_new_merge_id(GtkUIManager* v) __attribute__((weak)) {
	goPanic("gtk_ui_manager_new_merge_id: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
guint16 gtk_color_button_get_alpha(GtkColorButton* v) __attribute__((weak)) {
	goPanic("gtk_color_button_get_alpha: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_accel_map_lock_path(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_accel_map_lock_path: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_accel_map_unlock_path(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_accel_map_unlock_path: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_activate(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_activate: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_connect_accelerator(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_connect_accelerator: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_disconnect_accelerator(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_disconnect_accelerator: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_group_add_action(GtkActionGroup* v, GtkAction* _0) __attribute__((weak)) {
	goPanic("gtk_action_group_add_action: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_group_add_action_with_accel(GtkActionGroup* v, GtkAction* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_action_group_add_action_with_accel: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_group_remove_action(GtkActionGroup* v, GtkAction* _0) __attribute__((weak)) {
	goPanic("gtk_action_group_remove_action: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_group_set_sensitive(GtkActionGroup* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_action_group_set_sensitive: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_group_set_translate_func(GtkActionGroup* v, GtkTranslateFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_action_group_set_translate_func: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_group_set_translation_domain(GtkActionGroup* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_action_group_set_translation_domain: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_group_set_visible(GtkActionGroup* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_action_group_set_visible: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_set_accel_group(GtkAction* v, GtkAccelGroup* _0) __attribute__((weak)) {
	goPanic("gtk_action_set_accel_group: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_action_set_accel_path(GtkAction* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_action_set_accel_path: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_alignment_get_padding(GtkAlignment* v, guint* _0, guint* _1, guint* _2, guint* _3) __attribute__((weak)) {
	goPanic("gtk_alignment_get_padding: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_alignment_set_padding(GtkAlignment* v, guint _0, guint _1, guint _2, guint _3) __attribute__((weak)) {
	goPanic("gtk_alignment_set_padding: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_button_get_alignment(GtkButton* v, gfloat* _0, gfloat* _1) __attribute__((weak)) {
	goPanic("gtk_button_get_alignment: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_button_set_alignment(GtkButton* v, gfloat _0, gfloat _1) __attribute__((weak)) {
	goPanic("gtk_button_set_alignment: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_button_set_focus_on_click(GtkButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_button_set_focus_on_click: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_calendar_set_display_options(GtkCalendar* v, GtkCalendarDisplayOptions _0) __attribute__((weak)) {
	goPanic("gtk_calendar_set_display_options: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_cell_layout_add_attribute(GtkCellLayout* v, GtkCellRenderer* _0, const gchar* _1, gint _2) __attribute__((weak)) {
	goPanic("gtk_cell_layout_add_attribute: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_cell_layout_clear(GtkCellLayout* v) __attribute__((weak)) {
	goPanic("gtk_cell_layout_clear: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_cell_layout_clear_attributes(GtkCellLayout* v, GtkCellRenderer* _0) __attribute__((weak)) {
	goPanic("gtk_cell_layout_clear_attributes: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_cell_layout_pack_end(GtkCellLayout* v, GtkCellRenderer* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_cell_layout_pack_end: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_cell_layout_pack_start(GtkCellLayout* v, GtkCellRenderer* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_cell_layout_pack_start: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_cell_layout_reorder(GtkCellLayout* v, GtkCellRenderer* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_cell_layout_reorder: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_cell_layout_set_cell_data_func(GtkCellLayout* v, GtkCellRenderer* _0, GtkCellLayoutDataFunc _1, gpointer _2, GDestroyNotify _3) __attribute__((weak)) {
	goPanic("gtk_cell_layout_set_cell_data_func: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_check_menu_item_set_draw_as_radio(GtkCheckMenuItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_check_menu_item_set_draw_as_radio: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_color_button_get_color(GtkColorButton* v, GdkColor* _0) __attribute__((weak)) {
	goPanic("gtk_color_button_get_color: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_color_button_set_alpha(GtkColorButton* v, guint16 _0) __attribute__((weak)) {
	goPanic("gtk_color_button_set_alpha: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_color_button_set_color(GtkColorButton* v, const GdkColor* _0) __attribute__((weak)) {
	goPanic("gtk_color_button_set_color: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_color_button_set_title(GtkColorButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_color_button_set_title: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_color_button_set_use_alpha(GtkColorButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_color_button_set_use_alpha: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_combo_box_popdown(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_popdown: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_combo_box_popup(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_popup: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_combo_box_set_active(GtkComboBox* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_active: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_combo_box_set_active_iter(GtkComboBox* v, GtkTreeIter* _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_active_iter: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_combo_box_set_column_span_column(GtkComboBox* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_column_span_column: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_combo_box_set_model(GtkComboBox* v, GtkTreeModel* _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_model: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_combo_box_set_row_span_column(GtkComboBox* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_row_span_column: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_combo_box_set_wrap_width(GtkComboBox* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_wrap_width: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_drag_source_set_target_list(GtkWidget* v, GtkTargetList* _0) __attribute__((weak)) {
	goPanic("gtk_drag_source_set_target_list: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_entry_completion_complete(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_complete: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_entry_completion_delete_action(GtkEntryCompletion* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_delete_action: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_entry_completion_insert_action_markup(GtkEntryCompletion* v, gint _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_entry_completion_insert_action_markup: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_entry_completion_insert_action_text(GtkEntryCompletion* v, gint _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_entry_completion_insert_action_text: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_entry_completion_set_match_func(GtkEntryCompletion* v, GtkEntryCompletionMatchFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_entry_completion_set_match_func: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_entry_completion_set_minimum_key_length(GtkEntryCompletion* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_set_minimum_key_length: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_entry_completion_set_model(GtkEntryCompletion* v, GtkTreeModel* _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_set_model: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_entry_completion_set_text_column(GtkEntryCompletion* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_set_text_column: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_entry_set_alignment(GtkEntry* v, gfloat _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_alignment: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_entry_set_completion(GtkEntry* v, GtkEntryCompletion* _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_completion: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_event_box_set_above_child(GtkEventBox* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_event_box_set_above_child: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_event_box_set_visible_window(GtkEventBox* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_event_box_set_visible_window: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_expander_set_expanded(GtkExpander* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_expander_set_expanded: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_expander_set_label(GtkExpander* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_expander_set_label: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_expander_set_label_widget(GtkExpander* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_expander_set_label_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_expander_set_spacing(GtkExpander* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_expander_set_spacing: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_expander_set_use_markup(GtkExpander* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_expander_set_use_markup: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_expander_set_use_underline(GtkExpander* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_expander_set_use_underline: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_add_filter(GtkFileChooser* v, GtkFileFilter* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_add_filter: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_remove_filter(GtkFileChooser* v, GtkFileFilter* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_remove_filter: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_select_all(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_select_all: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_set_action(GtkFileChooser* v, GtkFileChooserAction _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_action: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_set_current_name(GtkFileChooser* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_current_name: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_set_extra_widget(GtkFileChooser* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_extra_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_set_filter(GtkFileChooser* v, GtkFileFilter* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_filter: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_set_local_only(GtkFileChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_local_only: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_set_preview_widget(GtkFileChooser* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_preview_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_set_preview_widget_active(GtkFileChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_preview_widget_active: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_set_select_multiple(GtkFileChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_select_multiple: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_set_use_preview_label(GtkFileChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_use_preview_label: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_unselect_all(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_unselect_all: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_unselect_filename(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_unselect_filename: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_chooser_unselect_uri(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_unselect_uri: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_filter_add_custom(GtkFileFilter* v, GtkFileFilterFlags _0, GtkFileFilterFunc _1, gpointer _2, GDestroyNotify _3) __attribute__((weak)) {
	goPanic("gtk_file_filter_add_custom: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_filter_add_mime_type(GtkFileFilter* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_file_filter_add_mime_type: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_filter_add_pattern(GtkFileFilter* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_file_filter_add_pattern: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_file_filter_set_name(GtkFileFilter* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_file_filter_set_name: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_font_button_set_show_size(GtkFontButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_font_button_set_show_size: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_font_button_set_show_style(GtkFontButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_font_button_set_show_style: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_font_button_set_title(GtkFontButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_font_button_set_title: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_font_button_set_use_font(GtkFontButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_font_button_set_use_font: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_font_button_set_use_size(GtkFontButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_font_button_set_use_size: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_icon_info_set_raw_coordinates(GtkIconInfo* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_icon_info_set_raw_coordinates: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_icon_theme_add_builtin_icon(const gchar* _0, gint _1, GdkPixbuf* _2) __attribute__((weak)) {
	goPanic("gtk_icon_theme_add_builtin_icon: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_icon_theme_append_search_path(GtkIconTheme* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_icon_theme_append_search_path: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_icon_theme_get_search_path(GtkIconTheme* v, gchar*** _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_icon_theme_get_search_path: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_icon_theme_prepend_search_path(GtkIconTheme* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_icon_theme_prepend_search_path: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_icon_theme_set_custom_theme(GtkIconTheme* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_icon_theme_set_custom_theme: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_icon_theme_set_screen(GtkIconTheme* v, GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_icon_theme_set_screen: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_icon_theme_set_search_path(GtkIconTheme* v, const gchar** _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_icon_theme_set_search_path: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_menu_attach(GtkMenu* v, GtkWidget* _0, guint _1, guint _2, guint _3, guint _4) __attribute__((weak)) {
	goPanic("gtk_menu_attach: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_menu_set_monitor(GtkMenu* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_menu_set_monitor: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_menu_shell_cancel(GtkMenuShell* v) __attribute__((weak)) {
	goPanic("gtk_menu_shell_cancel: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_message_dialog_set_markup(GtkMessageDialog* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_message_dialog_set_markup: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_radio_action_set_group(GtkRadioAction* v, GSList* _0) __attribute__((weak)) {
	goPanic("gtk_radio_action_set_group: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_radio_tool_button_set_group(GtkRadioToolButton* v, GSList* _0) __attribute__((weak)) {
	goPanic("gtk_radio_tool_button_set_group: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_rc_reset_styles(GtkSettings* _0) __attribute__((weak)) {
	goPanic("gtk_rc_reset_styles: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_scale_get_layout_offsets(GtkScale* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_scale_get_layout_offsets: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_separator_tool_item_set_draw(GtkSeparatorToolItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_separator_tool_item_set_draw: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_text_buffer_select_range(GtkTextBuffer* v, const GtkTextIter* _0, const GtkTextIter* _1) __attribute__((weak)) {
	goPanic("gtk_text_buffer_select_range: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_text_view_set_accepts_tab(GtkTextView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_text_view_set_accepts_tab: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_text_view_set_overwrite(GtkTextView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_text_view_set_overwrite: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_toggle_action_set_active(GtkToggleAction* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_toggle_action_set_active: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_toggle_action_set_draw_as_radio(GtkToggleAction* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_toggle_action_set_draw_as_radio: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_toggle_action_toggled(GtkToggleAction* v) __attribute__((weak)) {
	goPanic("gtk_toggle_action_toggled: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_toggle_tool_button_set_active(GtkToggleToolButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_toggle_tool_button_set_active: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_button_set_icon_widget(GtkToolButton* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_tool_button_set_icon_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_button_set_label(GtkToolButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tool_button_set_label: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_button_set_label_widget(GtkToolButton* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_tool_button_set_label_widget: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_button_set_stock_id(GtkToolButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tool_button_set_stock_id: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_button_set_use_underline(GtkToolButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tool_button_set_use_underline: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_item_set_expand(GtkToolItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_set_expand: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_item_set_homogeneous(GtkToolItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_set_homogeneous: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_item_set_is_important(GtkToolItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_set_is_important: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_item_set_proxy_menu_item(GtkToolItem* v, const gchar* _0, GtkWidget* _1) __attribute__((weak)) {
	goPanic("gtk_tool_item_set_proxy_menu_item: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_item_set_use_drag_window(GtkToolItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_set_use_drag_window: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_item_set_visible_horizontal(GtkToolItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_set_visible_horizontal: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tool_item_set_visible_vertical(GtkToolItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tool_item_set_visible_vertical: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_toolbar_insert(GtkToolbar* v, GtkToolItem* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_toolbar_insert: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_toolbar_set_drop_highlight_item(GtkToolbar* v, GtkToolItem* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_toolbar_set_drop_highlight_item: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_toolbar_set_show_arrow(GtkToolbar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_toolbar_set_show_arrow: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tree_model_filter_clear_cache(GtkTreeModelFilter* v) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_clear_cache: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tree_model_filter_convert_iter_to_child_iter(GtkTreeModelFilter* v, GtkTreeIter* _0, GtkTreeIter* _1) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_convert_iter_to_child_iter: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tree_model_filter_refilter(GtkTreeModelFilter* v) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_refilter: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tree_model_filter_set_modify_func(GtkTreeModelFilter* v, gint _0, GType* _1, GtkTreeModelFilterModifyFunc _2, gpointer _3, GDestroyNotify _4) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_set_modify_func: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tree_model_filter_set_visible_column(GtkTreeModelFilter* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_set_visible_column: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tree_model_filter_set_visible_func(GtkTreeModelFilter* v, GtkTreeModelFilterVisibleFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_tree_model_filter_set_visible_func: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_tree_view_column_set_expand(GtkTreeViewColumn* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_column_set_expand: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_ui_manager_add_ui(GtkUIManager* v, guint _0, const gchar* _1, const gchar* _2, const gchar* _3, GtkUIManagerItemType _4, gboolean _5) __attribute__((weak)) {
	goPanic("gtk_ui_manager_add_ui: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_ui_manager_ensure_update(GtkUIManager* v) __attribute__((weak)) {
	goPanic("gtk_ui_manager_ensure_update: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_ui_manager_insert_action_group(GtkUIManager* v, GtkActionGroup* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_ui_manager_insert_action_group: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_ui_manager_remove_action_group(GtkUIManager* v, GtkActionGroup* _0) __attribute__((weak)) {
	goPanic("gtk_ui_manager_remove_action_group: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_ui_manager_remove_ui(GtkUIManager* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_ui_manager_remove_ui: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_ui_manager_set_add_tearoffs(GtkUIManager* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_ui_manager_set_add_tearoffs: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_widget_add_mnemonic_label(GtkWidget* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_widget_add_mnemonic_label: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_widget_queue_resize_no_redraw(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_queue_resize_no_redraw: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_widget_remove_mnemonic_label(GtkWidget* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_widget_remove_mnemonic_label: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_widget_set_no_show_all(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_no_show_all: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_window_set_accept_focus(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_accept_focus: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_window_set_default_icon(GdkPixbuf* _0) __attribute__((weak)) {
	goPanic("gtk_window_set_default_icon: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_window_set_keep_above(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_keep_above: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 4))
void gtk_window_set_keep_below(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_keep_below: library too old: needs at least 2.4");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
AtkObject* gtk_combo_box_get_popup_accessible(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_popup_accessible: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GList* gtk_icon_view_get_selected_items(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_selected_items: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GList* gtk_menu_get_for_attach_widget(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_menu_get_for_attach_widget: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GOptionGroup* gtk_get_option_group(gboolean _0) __attribute__((weak)) {
	goPanic("gtk_get_option_group: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GdkPixbuf* gtk_about_dialog_get_logo(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_logo: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GdkPixbuf* gtk_clipboard_wait_for_image(GtkClipboard* v) __attribute__((weak)) {
	goPanic("gtk_clipboard_wait_for_image: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GdkPixbuf* gtk_selection_data_get_pixbuf(const GtkSelectionData* v) __attribute__((weak)) {
	goPanic("gtk_selection_data_get_pixbuf: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkCellRenderer* gtk_cell_renderer_combo_new(void) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_combo_new: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkCellRenderer* gtk_cell_renderer_progress_new(void) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_progress_new: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkOrientation gtk_icon_view_get_item_orientation(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_item_orientation: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkSelectionMode gtk_icon_view_get_selection_mode(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_selection_mode: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkToolItem* gtk_menu_tool_button_new(GtkWidget* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_menu_tool_button_new: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkToolItem* gtk_menu_tool_button_new_from_stock(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_menu_tool_button_new_from_stock: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkTreeModel* gtk_icon_view_get_model(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_model: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkTreePath* gtk_cell_view_get_displayed_row(GtkCellView* v) __attribute__((weak)) {
	goPanic("gtk_cell_view_get_displayed_row: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkTreePath* gtk_icon_view_get_path_at_pos(GtkIconView* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_path_at_pos: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_about_dialog_new(void) __attribute__((weak)) {
	goPanic("gtk_about_dialog_new: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_button_get_image(GtkButton* v) __attribute__((weak)) {
	goPanic("gtk_button_get_image: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_cell_view_new(void) __attribute__((weak)) {
	goPanic("gtk_cell_view_new: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_cell_view_new_with_context(GtkCellArea* _0, GtkCellAreaContext* _1) __attribute__((weak)) {
	goPanic("gtk_cell_view_new_with_context: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_cell_view_new_with_markup(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_cell_view_new_with_markup: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_cell_view_new_with_pixbuf(GdkPixbuf* _0) __attribute__((weak)) {
	goPanic("gtk_cell_view_new_with_pixbuf: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_cell_view_new_with_text(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_cell_view_new_with_text: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_file_chooser_button_new(const gchar* _0, GtkFileChooserAction _1) __attribute__((weak)) {
	goPanic("gtk_file_chooser_button_new: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_file_chooser_button_new_with_dialog(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_button_new_with_dialog: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_icon_view_new(void) __attribute__((weak)) {
	goPanic("gtk_icon_view_new: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_icon_view_new_with_model(GtkTreeModel* _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_new_with_model: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_image_new_from_icon_name(const gchar* _0, GtkIconSize _1) __attribute__((weak)) {
	goPanic("gtk_image_new_from_icon_name: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_menu_tool_button_get_menu(GtkMenuToolButton* v) __attribute__((weak)) {
	goPanic("gtk_menu_tool_button_get_menu: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
PangoEllipsizeMode gtk_label_get_ellipsize(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_ellipsize: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
PangoEllipsizeMode gtk_progress_bar_get_ellipsize(GtkProgressBar* v) __attribute__((weak)) {
	goPanic("gtk_progress_bar_get_ellipsize: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* const* gtk_about_dialog_get_artists(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_artists: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* const* gtk_about_dialog_get_authors(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_authors: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* const* gtk_about_dialog_get_documenters(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_documenters: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_about_dialog_get_comments(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_comments: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_about_dialog_get_copyright(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_copyright: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_about_dialog_get_license(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_license: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_about_dialog_get_logo_icon_name(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_logo_icon_name: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_about_dialog_get_translator_credits(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_translator_credits: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_about_dialog_get_version(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_version: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_about_dialog_get_website(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_website: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_about_dialog_get_website_label(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_website_label: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_action_get_accel_path(GtkAction* v) __attribute__((weak)) {
	goPanic("gtk_action_get_accel_path: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_action_group_translate_string(GtkActionGroup* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_action_group_translate_string: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_file_chooser_button_get_title(GtkFileChooserButton* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_button_get_title: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
const gchar* gtk_window_get_icon_name(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_icon_name: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_alternative_dialog_button_order(GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_alternative_dialog_button_order: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_cell_view_get_size_of_row(GtkCellView* v, GtkTreePath* _0, GtkRequisition* _1) __attribute__((weak)) {
	goPanic("gtk_cell_view_get_size_of_row: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_clipboard_wait_is_image_available(GtkClipboard* v) __attribute__((weak)) {
	goPanic("gtk_clipboard_wait_is_image_available: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_combo_box_get_focus_on_click(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_focus_on_click: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_entry_completion_get_inline_completion(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_get_inline_completion: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_entry_completion_get_popup_completion(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_get_popup_completion: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_file_chooser_get_show_hidden(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_show_hidden: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_icon_view_path_is_selected(GtkIconView* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_path_is_selected: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_label_get_single_line_mode(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_single_line_mode: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_selection_data_set_pixbuf(GtkSelectionData* v, GdkPixbuf* _0) __attribute__((weak)) {
	goPanic("gtk_selection_data_set_pixbuf: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_selection_data_set_uris(GtkSelectionData* v, gchar** _0) __attribute__((weak)) {
	goPanic("gtk_selection_data_set_uris: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_selection_data_targets_include_image(const GtkSelectionData* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_selection_data_targets_include_image: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_text_buffer_backspace(GtkTextBuffer* v, GtkTextIter* _0, gboolean _1, gboolean _2) __attribute__((weak)) {
	goPanic("gtk_text_buffer_backspace: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_text_view_get_iter_at_position(GtkTextView* v, GtkTextIter* _0, gint* _1, gint _2, gint _3) __attribute__((weak)) {
	goPanic("gtk_text_view_get_iter_at_position: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_tree_view_get_fixed_height_mode(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_fixed_height_mode: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_tree_view_get_hover_expand(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_hover_expand: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_tree_view_get_hover_selection(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_hover_selection: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gboolean gtk_window_get_focus_on_map(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_focus_on_map: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gchar* gtk_accelerator_get_label(guint _0, GdkModifierType _1) __attribute__((weak)) {
	goPanic("gtk_accelerator_get_label: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gchar** gtk_selection_data_get_uris(const GtkSelectionData* v) __attribute__((weak)) {
	goPanic("gtk_selection_data_get_uris: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gdouble gtk_label_get_angle(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_angle: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_combo_box_get_column_span_column(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_column_span_column: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_combo_box_get_row_span_column(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_row_span_column: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_combo_box_get_wrap_width(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_wrap_width: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_entry_completion_get_text_column(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_get_text_column: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_file_chooser_button_get_width_chars(GtkFileChooserButton* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_button_get_width_chars: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_icon_view_get_column_spacing(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_column_spacing: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_icon_view_get_columns(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_columns: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_icon_view_get_item_width(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_item_width: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_icon_view_get_margin(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_margin: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_icon_view_get_markup_column(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_markup_column: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_icon_view_get_pixbuf_column(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_pixbuf_column: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_icon_view_get_row_spacing(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_row_spacing: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_icon_view_get_spacing(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_spacing: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_icon_view_get_text_column(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_text_column: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_image_get_pixel_size(GtkImage* v) __attribute__((weak)) {
	goPanic("gtk_image_get_pixel_size: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_label_get_max_width_chars(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_max_width_chars: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint gtk_label_get_width_chars(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_width_chars: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
gint* gtk_icon_theme_get_icon_sizes(GtkIconTheme* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_icon_theme_get_icon_sizes: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_artists(GtkAboutDialog* v, const gchar** _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_artists: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_authors(GtkAboutDialog* v, const gchar** _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_authors: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_comments(GtkAboutDialog* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_comments: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_copyright(GtkAboutDialog* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_copyright: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_documenters(GtkAboutDialog* v, const gchar** _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_documenters: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_license(GtkAboutDialog* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_license: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_logo(GtkAboutDialog* v, GdkPixbuf* _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_logo: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_logo_icon_name(GtkAboutDialog* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_logo_icon_name: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_translator_credits(GtkAboutDialog* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_translator_credits: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_version(GtkAboutDialog* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_version: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_website(GtkAboutDialog* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_website: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_about_dialog_set_website_label(GtkAboutDialog* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_website_label: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_action_set_sensitive(GtkAction* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_action_set_sensitive: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_action_set_visible(GtkAction* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_action_set_visible: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_button_set_image(GtkButton* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_button_set_image: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_cell_renderer_stop_editing(GtkCellRenderer* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_stop_editing: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_cell_view_set_background_color(GtkCellView* v, const GdkColor* _0) __attribute__((weak)) {
	goPanic("gtk_cell_view_set_background_color: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_cell_view_set_displayed_row(GtkCellView* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_cell_view_set_displayed_row: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_cell_view_set_model(GtkCellView* v, GtkTreeModel* _0) __attribute__((weak)) {
	goPanic("gtk_cell_view_set_model: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_clipboard_request_image(GtkClipboard* v, GtkClipboardImageReceivedFunc _0, gpointer _1) __attribute__((weak)) {
	goPanic("gtk_clipboard_request_image: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_clipboard_set_can_store(GtkClipboard* v, const GtkTargetEntry* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_clipboard_set_can_store: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_clipboard_set_image(GtkClipboard* v, GdkPixbuf* _0) __attribute__((weak)) {
	goPanic("gtk_clipboard_set_image: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_clipboard_store(GtkClipboard* v) __attribute__((weak)) {
	goPanic("gtk_clipboard_store: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_combo_box_set_add_tearoffs(GtkComboBox* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_add_tearoffs: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_combo_box_set_focus_on_click(GtkComboBox* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_focus_on_click: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_combo_box_set_row_separator_func(GtkComboBox* v, GtkTreeViewRowSeparatorFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_row_separator_func: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_dialog_set_alternative_button_order_from_array(GtkDialog* v, gint _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_dialog_set_alternative_button_order_from_array: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_drag_dest_add_image_targets(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_drag_dest_add_image_targets: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_drag_dest_add_text_targets(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_drag_dest_add_text_targets: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_drag_dest_add_uri_targets(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_drag_dest_add_uri_targets: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_drag_source_add_image_targets(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_drag_source_add_image_targets: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_drag_source_add_text_targets(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_drag_source_add_text_targets: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_drag_source_add_uri_targets(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_drag_source_add_uri_targets: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_entry_completion_insert_prefix(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_insert_prefix: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_entry_completion_set_inline_completion(GtkEntryCompletion* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_set_inline_completion: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_entry_completion_set_popup_completion(GtkEntryCompletion* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_set_popup_completion: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_file_chooser_button_set_title(GtkFileChooserButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_button_set_title: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_file_chooser_button_set_width_chars(GtkFileChooserButton* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_button_set_width_chars: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_file_chooser_set_show_hidden(GtkFileChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_show_hidden: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_file_filter_add_pixbuf_formats(GtkFileFilter* v) __attribute__((weak)) {
	goPanic("gtk_file_filter_add_pixbuf_formats: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_item_activated(GtkIconView* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_item_activated: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_select_all(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_select_all: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_select_path(GtkIconView* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_select_path: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_selected_foreach(GtkIconView* v, GtkIconViewForeachFunc _0, gpointer _1) __attribute__((weak)) {
	goPanic("gtk_icon_view_selected_foreach: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_column_spacing(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_column_spacing: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_columns(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_columns: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_item_orientation(GtkIconView* v, GtkOrientation _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_item_orientation: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_item_width(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_item_width: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_margin(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_margin: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_markup_column(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_markup_column: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_model(GtkIconView* v, GtkTreeModel* _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_model: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_pixbuf_column(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_pixbuf_column: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_row_spacing(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_row_spacing: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_selection_mode(GtkIconView* v, GtkSelectionMode _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_selection_mode: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_spacing(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_spacing: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_set_text_column(GtkIconView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_text_column: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_unselect_all(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_unselect_all: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_icon_view_unselect_path(GtkIconView* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_unselect_path: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_image_get_icon_name(GtkImage* v, const gchar** _0, GtkIconSize* _1) __attribute__((weak)) {
	goPanic("gtk_image_get_icon_name: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_image_set_from_icon_name(GtkImage* v, const gchar* _0, GtkIconSize _1) __attribute__((weak)) {
	goPanic("gtk_image_set_from_icon_name: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_image_set_pixel_size(GtkImage* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_image_set_pixel_size: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_label_set_angle(GtkLabel* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_label_set_angle: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_label_set_ellipsize(GtkLabel* v, PangoEllipsizeMode _0) __attribute__((weak)) {
	goPanic("gtk_label_set_ellipsize: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_label_set_max_width_chars(GtkLabel* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_label_set_max_width_chars: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_label_set_single_line_mode(GtkLabel* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_label_set_single_line_mode: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_label_set_width_chars(GtkLabel* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_label_set_width_chars: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_list_store_insert_with_valuesv(GtkListStore* v, GtkTreeIter* _0, gint _1, gint* _2, GValue* _3, gint _4) __attribute__((weak)) {
	goPanic("gtk_list_store_insert_with_valuesv: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_menu_tool_button_set_menu(GtkMenuToolButton* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_menu_tool_button_set_menu: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_progress_bar_set_ellipsize(GtkProgressBar* v, PangoEllipsizeMode _0) __attribute__((weak)) {
	goPanic("gtk_progress_bar_set_ellipsize: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_target_list_add_image_targets(GtkTargetList* v, guint _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_target_list_add_image_targets: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_target_list_add_text_targets(GtkTargetList* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_target_list_add_text_targets: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_target_list_add_uri_targets(GtkTargetList* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_target_list_add_uri_targets: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_tool_item_rebuild_menu(GtkToolItem* v) __attribute__((weak)) {
	goPanic("gtk_tool_item_rebuild_menu: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_tree_view_set_fixed_height_mode(GtkTreeView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_fixed_height_mode: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_tree_view_set_hover_expand(GtkTreeView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_hover_expand: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_tree_view_set_hover_selection(GtkTreeView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_hover_selection: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_tree_view_set_row_separator_func(GtkTreeView* v, GtkTreeViewRowSeparatorFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_row_separator_func: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_window_set_default_icon_name(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_window_set_default_icon_name: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_window_set_focus_on_map(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_focus_on_map: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 6))
void gtk_window_set_icon_name(GtkWindow* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_window_set_icon_name: library too old: needs at least 2.6");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
GtkPackDirection gtk_menu_bar_get_child_pack_direction(GtkMenuBar* v) __attribute__((weak)) {
	goPanic("gtk_menu_bar_get_child_pack_direction: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
GtkPackDirection gtk_menu_bar_get_pack_direction(GtkMenuBar* v) __attribute__((weak)) {
	goPanic("gtk_menu_bar_get_pack_direction: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
GtkTreeModel* gtk_tree_row_reference_get_model(GtkTreeRowReference* v) __attribute__((weak)) {
	goPanic("gtk_tree_row_reference_get_model: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
GtkWidget* gtk_scrolled_window_get_hscrollbar(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_hscrollbar: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
GtkWidget* gtk_scrolled_window_get_vscrollbar(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_vscrollbar: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
cairo_surface_t* gtk_icon_view_create_drag_icon(GtkIconView* v, GtkTreePath* _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_create_drag_icon: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
const gchar* gtk_tool_button_get_icon_name(GtkToolButton* v) __attribute__((weak)) {
	goPanic("gtk_tool_button_get_icon_name: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_about_dialog_get_wrap_license(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_wrap_license: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_entry_completion_get_popup_set_width(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_get_popup_set_width: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_entry_completion_get_popup_single_match(GtkEntryCompletion* v) __attribute__((weak)) {
	goPanic("gtk_entry_completion_get_popup_single_match: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_file_chooser_get_do_overwrite_confirmation(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_do_overwrite_confirmation: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_icon_view_get_cursor(GtkIconView* v, GtkTreePath** _0, GtkCellRenderer** _1) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_cursor: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_icon_view_get_dest_item_at_pos(GtkIconView* v, gint _0, gint _1, GtkTreePath** _2, GtkIconViewDropPosition* _3) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_dest_item_at_pos: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_icon_view_get_item_at_pos(GtkIconView* v, gint _0, gint _1, GtkTreePath** _2, GtkCellRenderer** _3) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_item_at_pos: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_icon_view_get_reorderable(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_reorderable: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_icon_view_get_visible_range(GtkIconView* v, GtkTreePath** _0, GtkTreePath** _1) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_visible_range: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_menu_shell_get_take_focus(GtkMenuShell* v) __attribute__((weak)) {
	goPanic("gtk_menu_shell_get_take_focus: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_size_group_get_ignore_hidden(GtkSizeGroup* v) __attribute__((weak)) {
	goPanic("gtk_size_group_get_ignore_hidden: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_text_iter_backward_visible_line(GtkTextIter* v) __attribute__((weak)) {
	goPanic("gtk_text_iter_backward_visible_line: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_text_iter_backward_visible_lines(GtkTextIter* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_text_iter_backward_visible_lines: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_text_iter_forward_visible_line(GtkTextIter* v) __attribute__((weak)) {
	goPanic("gtk_text_iter_forward_visible_line: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_text_iter_forward_visible_lines(GtkTextIter* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_text_iter_forward_visible_lines: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_tree_view_get_visible_range(GtkTreeView* v, GtkTreePath** _0, GtkTreePath** _1) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_visible_range: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gboolean gtk_window_get_urgency_hint(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_urgency_hint: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
gint gtk_dialog_get_response_for_widget(GtkDialog* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_dialog_get_response_for_widget: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_about_dialog_set_wrap_license(GtkAboutDialog* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_wrap_license: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_drag_set_icon_name(GdkDragContext* _0, const gchar* _1, gint _2, gint _3) __attribute__((weak)) {
	goPanic("gtk_drag_set_icon_name: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_drag_source_set_icon_name(GtkWidget* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_drag_source_set_icon_name: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_entry_completion_set_popup_set_width(GtkEntryCompletion* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_set_popup_set_width: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_entry_completion_set_popup_single_match(GtkEntryCompletion* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_set_popup_single_match: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_file_chooser_set_do_overwrite_confirmation(GtkFileChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_do_overwrite_confirmation: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_icon_view_enable_model_drag_dest(GtkIconView* v, const GtkTargetEntry* _0, gint _1, GdkDragAction _2) __attribute__((weak)) {
	goPanic("gtk_icon_view_enable_model_drag_dest: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_icon_view_enable_model_drag_source(GtkIconView* v, GdkModifierType _0, const GtkTargetEntry* _1, gint _2, GdkDragAction _3) __attribute__((weak)) {
	goPanic("gtk_icon_view_enable_model_drag_source: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_icon_view_get_drag_dest_item(GtkIconView* v, GtkTreePath** _0, GtkIconViewDropPosition* _1) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_drag_dest_item: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_icon_view_scroll_to_path(GtkIconView* v, GtkTreePath* _0, gboolean _1, gfloat _2, gfloat _3) __attribute__((weak)) {
	goPanic("gtk_icon_view_scroll_to_path: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_icon_view_set_cursor(GtkIconView* v, GtkTreePath* _0, GtkCellRenderer* _1, gboolean _2) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_cursor: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_icon_view_set_drag_dest_item(GtkIconView* v, GtkTreePath* _0, GtkIconViewDropPosition _1) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_drag_dest_item: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_icon_view_set_reorderable(GtkIconView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_reorderable: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_icon_view_unset_model_drag_dest(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_unset_model_drag_dest: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_icon_view_unset_model_drag_source(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_unset_model_drag_source: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_image_clear(GtkImage* v) __attribute__((weak)) {
	goPanic("gtk_image_clear: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_menu_bar_set_child_pack_direction(GtkMenuBar* v, GtkPackDirection _0) __attribute__((weak)) {
	goPanic("gtk_menu_bar_set_child_pack_direction: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_menu_bar_set_pack_direction(GtkMenuBar* v, GtkPackDirection _0) __attribute__((weak)) {
	goPanic("gtk_menu_bar_set_pack_direction: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_menu_shell_set_take_focus(GtkMenuShell* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_menu_shell_set_take_focus: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_size_group_set_ignore_hidden(GtkSizeGroup* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_size_group_set_ignore_hidden: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_stock_set_translate_func(const gchar* _0, GtkTranslateFunc _1, gpointer _2, GDestroyNotify _3) __attribute__((weak)) {
	goPanic("gtk_stock_set_translate_func: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_tool_button_set_icon_name(GtkToolButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_tool_button_set_icon_name: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_tree_view_column_queue_resize(GtkTreeViewColumn* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_column_queue_resize: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_window_present_with_time(GtkWindow* v, guint32 _0) __attribute__((weak)) {
	goPanic("gtk_window_present_with_time: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 2 || (GTK_MAJOR_VERSION == 2 && GTK_MINOR_VERSION < 8))
void gtk_window_set_urgency_hint(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_urgency_hint: library too old: needs at least 2.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GAppInfo* gtk_app_chooser_get_app_info(GtkAppChooser* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_get_app_info: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GIcon* gtk_numerable_icon_get_background_gicon(GtkNumerableIcon* v) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_get_background_gicon: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GIcon* gtk_numerable_icon_new(GIcon* _0) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GIcon* gtk_numerable_icon_new_with_style_context(GIcon* _0, GtkStyleContext* _1) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_new_with_style_context: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GList* gtk_application_get_windows(GtkApplication* v) __attribute__((weak)) {
	goPanic("gtk_application_get_windows: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GList* gtk_style_context_list_classes(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_list_classes: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GList* gtk_style_context_list_regions(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_list_regions: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GSList* gtk_widget_path_iter_list_classes(const GtkWidgetPath* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_list_classes: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GSList* gtk_widget_path_iter_list_regions(const GtkWidgetPath* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_list_regions: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GTokenType gtk_binding_entry_add_signal_from_string(GtkBindingSet* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_binding_entry_add_signal_from_string: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GType gtk_widget_path_get_object_type(const GtkWidgetPath* v) __attribute__((weak)) {
	goPanic("gtk_widget_path_get_object_type: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GType gtk_widget_path_iter_get_object_type(const GtkWidgetPath* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_get_object_type: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GdkEventMask gtk_widget_get_device_events(GtkWidget* v, GdkDevice* _0) __attribute__((weak)) {
	goPanic("gtk_widget_get_device_events: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GdkPixbuf* gtk_icon_info_load_symbolic(GtkIconInfo* v, const GdkRGBA* _0, const GdkRGBA* _1, const GdkRGBA* _2, const GdkRGBA* _3, gboolean* _4) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_symbolic: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GdkPixbuf* gtk_icon_info_load_symbolic_for_context(GtkIconInfo* v, GtkStyleContext* _0, gboolean* _1) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_symbolic_for_context: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GdkPixbuf* gtk_icon_info_load_symbolic_for_style(GtkIconInfo* v, GtkStyle* _0, GtkStateType _1, gboolean* _2) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_symbolic_for_style: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GdkPixbuf* gtk_icon_set_render_icon_pixbuf(GtkIconSet* v, GtkStyleContext* _0, GtkIconSize _1) __attribute__((weak)) {
	goPanic("gtk_icon_set_render_icon_pixbuf: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GdkPixbuf* gtk_render_icon_pixbuf(GtkStyleContext* _0, const GtkIconSource* _1, GtkIconSize _2) __attribute__((weak)) {
	goPanic("gtk_render_icon_pixbuf: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GdkPixbuf* gtk_widget_render_icon_pixbuf(GtkWidget* v, const gchar* _0, GtkIconSize _1) __attribute__((weak)) {
	goPanic("gtk_widget_render_icon_pixbuf: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkAdjustment* gtk_scrollable_get_hadjustment(GtkScrollable* v) __attribute__((weak)) {
	goPanic("gtk_scrollable_get_hadjustment: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkAdjustment* gtk_scrollable_get_vadjustment(GtkScrollable* v) __attribute__((weak)) {
	goPanic("gtk_scrollable_get_vadjustment: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkApplication* gtk_application_new(const gchar* _0, GApplicationFlags _1) __attribute__((weak)) {
	goPanic("gtk_application_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkApplication* gtk_window_get_application(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_application: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkCellArea* gtk_cell_area_box_new(void) __attribute__((weak)) {
	goPanic("gtk_cell_area_box_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkCellArea* gtk_cell_area_context_get_area(GtkCellAreaContext* v) __attribute__((weak)) {
	goPanic("gtk_cell_area_context_get_area: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkCellArea* gtk_cell_layout_get_area(GtkCellLayout* v) __attribute__((weak)) {
	goPanic("gtk_cell_layout_get_area: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkCellAreaContext* gtk_cell_area_copy_context(GtkCellArea* v, GtkCellAreaContext* _0) __attribute__((weak)) {
	goPanic("gtk_cell_area_copy_context: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkCellAreaContext* gtk_cell_area_create_context(GtkCellArea* v) __attribute__((weak)) {
	goPanic("gtk_cell_area_create_context: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkCellEditable* gtk_cell_area_get_edit_widget(GtkCellArea* v) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_edit_widget: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkCellRenderer* gtk_cell_area_get_cell_at_position(GtkCellArea* v, GtkCellAreaContext* _0, GtkWidget* _1, const GdkRectangle* _2, gint _3, gint _4, GdkRectangle* _5) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_cell_at_position: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkCellRenderer* gtk_cell_area_get_edited_cell(GtkCellArea* v) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_edited_cell: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkCellRenderer* gtk_cell_area_get_focus_cell(GtkCellArea* v) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_focus_cell: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkCellRenderer* gtk_cell_area_get_focus_from_sibling(GtkCellArea* v, GtkCellRenderer* _0) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_focus_from_sibling: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkEntryCompletion* gtk_entry_completion_new_with_area(GtkCellArea* _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_new_with_area: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkGradient* gtk_gradient_new_linear(gdouble _0, gdouble _1, gdouble _2, gdouble _3) __attribute__((weak)) {
	goPanic("gtk_gradient_new_linear: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkGradient* gtk_gradient_new_radial(gdouble _0, gdouble _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_gradient_new_radial: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkIconFactory* gtk_style_provider_get_icon_factory(GtkStyleProvider* v, GtkWidgetPath* _0) __attribute__((weak)) {
	goPanic("gtk_style_provider_get_icon_factory: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkJunctionSides gtk_style_context_get_junction_sides(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_get_junction_sides: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkJunctionSides gtk_theming_engine_get_junction_sides(GtkThemingEngine* v) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_junction_sides: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkLicense gtk_about_dialog_get_license_type(GtkAboutDialog* v) __attribute__((weak)) {
	goPanic("gtk_about_dialog_get_license_type: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkRequisition* gtk_requisition_new(void) __attribute__((weak)) {
	goPanic("gtk_requisition_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkScrollablePolicy gtk_scrollable_get_hscroll_policy(GtkScrollable* v) __attribute__((weak)) {
	goPanic("gtk_scrollable_get_hscroll_policy: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkScrollablePolicy gtk_scrollable_get_vscroll_policy(GtkScrollable* v) __attribute__((weak)) {
	goPanic("gtk_scrollable_get_vscroll_policy: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkSizeRequestMode gtk_cell_area_get_request_mode(GtkCellArea* v) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_request_mode: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkSizeRequestMode gtk_cell_renderer_get_request_mode(GtkCellRenderer* v) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_request_mode: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkSizeRequestMode gtk_widget_get_request_mode(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_request_mode: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkStateFlags gtk_cell_renderer_get_state(GtkCellRenderer* v, GtkWidget* _0, GtkCellRendererState _1) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_state: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkStateFlags gtk_style_context_get_state(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_get_state: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkStateFlags gtk_theming_engine_get_state(GtkThemingEngine* v) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_state: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkStateFlags gtk_widget_get_state_flags(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_state_flags: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkStyleContext* gtk_numerable_icon_get_style_context(GtkNumerableIcon* v) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_get_style_context: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkStyleProperties* gtk_style_provider_get_style(GtkStyleProvider* v, GtkWidgetPath* _0) __attribute__((weak)) {
	goPanic("gtk_style_provider_get_style: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkSymbolicColor* gtk_style_properties_lookup_color(GtkStyleProperties* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_style_properties_lookup_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkSymbolicColor* gtk_symbolic_color_new_alpha(GtkSymbolicColor* _0, gdouble _1) __attribute__((weak)) {
	goPanic("gtk_symbolic_color_new_alpha: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkSymbolicColor* gtk_symbolic_color_new_literal(const GdkRGBA* _0) __attribute__((weak)) {
	goPanic("gtk_symbolic_color_new_literal: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkSymbolicColor* gtk_symbolic_color_new_mix(GtkSymbolicColor* _0, GtkSymbolicColor* _1, gdouble _2) __attribute__((weak)) {
	goPanic("gtk_symbolic_color_new_mix: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkSymbolicColor* gtk_symbolic_color_new_name(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_symbolic_color_new_name: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkSymbolicColor* gtk_symbolic_color_new_shade(GtkSymbolicColor* _0, gdouble _1) __attribute__((weak)) {
	goPanic("gtk_symbolic_color_new_shade: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkTextDirection gtk_style_context_get_direction(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_get_direction: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkTextDirection gtk_theming_engine_get_direction(GtkThemingEngine* v) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_direction: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkTreeViewColumn* gtk_tree_view_column_new_with_area(GtkCellArea* _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_column_new_with_area: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_app_chooser_button_new(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_app_chooser_button_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_app_chooser_dialog_get_widget(GtkAppChooserDialog* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_dialog_get_widget: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_app_chooser_dialog_new(GtkWindow* _0, GtkDialogFlags _1, GFile* _2) __attribute__((weak)) {
	goPanic("gtk_app_chooser_dialog_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_app_chooser_dialog_new_for_content_type(GtkWindow* _0, GtkDialogFlags _1, const gchar* _2) __attribute__((weak)) {
	goPanic("gtk_app_chooser_dialog_new_for_content_type: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_app_chooser_widget_new(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_box_new(GtkOrientation _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_box_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_button_box_new(GtkOrientation _0) __attribute__((weak)) {
	goPanic("gtk_button_box_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_color_button_new_with_rgba(const GdkRGBA* _0) __attribute__((weak)) {
	goPanic("gtk_color_button_new_with_rgba: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_icon_view_new_with_area(GtkCellArea* _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_new_with_area: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_menu_shell_get_parent_shell(GtkMenuShell* v) __attribute__((weak)) {
	goPanic("gtk_menu_shell_get_parent_shell: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_menu_shell_get_selected_item(GtkMenuShell* v) __attribute__((weak)) {
	goPanic("gtk_menu_shell_get_selected_item: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_paned_new(GtkOrientation _0) __attribute__((weak)) {
	goPanic("gtk_paned_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_scale_new(GtkOrientation _0, GtkAdjustment* _1) __attribute__((weak)) {
	goPanic("gtk_scale_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_scale_new_with_range(GtkOrientation _0, gdouble _1, gdouble _2, gdouble _3) __attribute__((weak)) {
	goPanic("gtk_scale_new_with_range: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_scrollbar_new(GtkOrientation _0, GtkAdjustment* _1) __attribute__((weak)) {
	goPanic("gtk_scrollbar_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_separator_new(GtkOrientation _0) __attribute__((weak)) {
	goPanic("gtk_separator_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_switch_new(void) __attribute__((weak)) {
	goPanic("gtk_switch_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_tree_view_column_get_button(GtkTreeViewColumn* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_column_get_button: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidget* gtk_window_group_get_current_device_grab(GtkWindowGroup* v, GdkDevice* _0) __attribute__((weak)) {
	goPanic("gtk_window_group_get_current_device_grab: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidgetPath* gtk_widget_path_copy(const GtkWidgetPath* v) __attribute__((weak)) {
	goPanic("gtk_widget_path_copy: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
GtkWidgetPath* gtk_widget_path_new(void) __attribute__((weak)) {
	goPanic("gtk_widget_path_new: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const GList* gtk_cell_area_get_focus_siblings(GtkCellArea* v, GtkCellRenderer* _0) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_focus_siblings: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const GtkWidgetPath* gtk_style_context_get_path(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_get_path: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const GtkWidgetPath* gtk_theming_engine_get_path(GtkThemingEngine* v) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_path: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const PangoFontDescription* gtk_style_context_get_font(GtkStyleContext* v, GtkStateFlags _0) __attribute__((weak)) {
	goPanic("gtk_style_context_get_font: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const PangoFontDescription* gtk_theming_engine_get_font(GtkThemingEngine* v, GtkStateFlags _0) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_font: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const gchar* gtk_app_chooser_widget_get_default_text(GtkAppChooserWidget* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_get_default_text: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const gchar* gtk_cell_area_get_current_path_string(GtkCellArea* v) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_current_path_string: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const gchar* gtk_combo_box_get_active_id(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_active_id: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const gchar* gtk_numerable_icon_get_background_icon_name(GtkNumerableIcon* v) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_get_background_icon_name: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const gchar* gtk_numerable_icon_get_label(GtkNumerableIcon* v) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_get_label: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
const guchar* gtk_selection_data_get_data_with_length(const GtkSelectionData* v, gint* _0) __attribute__((weak)) {
	goPanic("gtk_selection_data_get_data_with_length: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_app_chooser_button_get_show_dialog_item(GtkAppChooserButton* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_button_get_show_dialog_item: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_app_chooser_widget_get_show_all(GtkAppChooserWidget* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_get_show_all: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_app_chooser_widget_get_show_default(GtkAppChooserWidget* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_get_show_default: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_app_chooser_widget_get_show_fallback(GtkAppChooserWidget* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_get_show_fallback: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_app_chooser_widget_get_show_other(GtkAppChooserWidget* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_get_show_other: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_app_chooser_widget_get_show_recommended(GtkAppChooserWidget* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_get_show_recommended: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_cairo_should_draw_window(cairo_t* _0, GdkWindow* _1) __attribute__((weak)) {
	goPanic("gtk_cairo_should_draw_window: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_calendar_get_day_is_marked(GtkCalendar* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_calendar_get_day_is_marked: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_cell_area_activate(GtkCellArea* v, GtkCellAreaContext* _0, GtkWidget* _1, const GdkRectangle* _2, GtkCellRendererState _3, gboolean _4) __attribute__((weak)) {
	goPanic("gtk_cell_area_activate: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_cell_area_focus(GtkCellArea* v, GtkDirectionType _0) __attribute__((weak)) {
	goPanic("gtk_cell_area_focus: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_cell_area_has_renderer(GtkCellArea* v, GtkCellRenderer* _0) __attribute__((weak)) {
	goPanic("gtk_cell_area_has_renderer: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_cell_area_is_activatable(GtkCellArea* v) __attribute__((weak)) {
	goPanic("gtk_cell_area_is_activatable: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_cell_area_is_focus_sibling(GtkCellArea* v, GtkCellRenderer* _0, GtkCellRenderer* _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_is_focus_sibling: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_cell_renderer_is_activatable(GtkCellRenderer* v) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_is_activatable: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_cell_view_get_draw_sensitive(GtkCellView* v) __attribute__((weak)) {
	goPanic("gtk_cell_view_get_draw_sensitive: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_cell_view_get_fit_model(GtkCellView* v) __attribute__((weak)) {
	goPanic("gtk_cell_view_get_fit_model: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_combo_box_get_popup_fixed_width(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_popup_fixed_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_combo_box_set_active_id(GtkComboBox* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_active_id: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_gradient_resolve(GtkGradient* v, GtkStyleProperties* _0, cairo_pattern_t** _1) __attribute__((weak)) {
	goPanic("gtk_gradient_resolve: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_menu_item_get_reserve_indicator(GtkMenuItem* v) __attribute__((weak)) {
	goPanic("gtk_menu_item_get_reserve_indicator: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_progress_bar_get_show_text(GtkProgressBar* v) __attribute__((weak)) {
	goPanic("gtk_progress_bar_get_show_text: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_style_context_has_class(GtkStyleContext* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_has_class: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_style_context_has_region(GtkStyleContext* v, const gchar* _0, GtkRegionFlags* _1) __attribute__((weak)) {
	goPanic("gtk_style_context_has_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_style_context_state_is_running(GtkStyleContext* v, GtkStateType _0, gdouble* _1) __attribute__((weak)) {
	goPanic("gtk_style_context_state_is_running: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_style_has_context(GtkStyle* v) __attribute__((weak)) {
	goPanic("gtk_style_has_context: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_style_properties_get_property(GtkStyleProperties* v, const gchar* _0, GtkStateFlags _1, GValue* _2) __attribute__((weak)) {
	goPanic("gtk_style_properties_get_property: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_switch_get_active(GtkSwitch* v) __attribute__((weak)) {
	goPanic("gtk_switch_get_active: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_symbolic_color_resolve(GtkSymbolicColor* v, GtkStyleProperties* _0, GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_symbolic_color_resolve: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_theming_engine_has_class(GtkThemingEngine* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_theming_engine_has_class: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_theming_engine_has_region(GtkThemingEngine* v, const gchar* _0, GtkRegionFlags* _1) __attribute__((weak)) {
	goPanic("gtk_theming_engine_has_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_theming_engine_lookup_color(GtkThemingEngine* v, const gchar* _0, GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_theming_engine_lookup_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_theming_engine_state_is_running(GtkThemingEngine* v, GtkStateType _0, gdouble* _1) __attribute__((weak)) {
	goPanic("gtk_theming_engine_state_is_running: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_tree_model_iter_previous(GtkTreeModel* v, GtkTreeIter* _0) __attribute__((weak)) {
	goPanic("gtk_tree_model_iter_previous: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_tree_view_is_blank_at_pos(GtkTreeView* v, gint _0, gint _1, GtkTreePath** _2, GtkTreeViewColumn** _3, gint* _4, gint* _5) __attribute__((weak)) {
	goPanic("gtk_tree_view_is_blank_at_pos: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_widget_device_is_shadowed(GtkWidget* v, GdkDevice* _0) __attribute__((weak)) {
	goPanic("gtk_widget_device_is_shadowed: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_widget_get_device_enabled(GtkWidget* v, GdkDevice* _0) __attribute__((weak)) {
	goPanic("gtk_widget_get_device_enabled: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_widget_path_has_parent(const GtkWidgetPath* v, GType _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_has_parent: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_widget_path_is_type(const GtkWidgetPath* v, GType _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_is_type: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_widget_path_iter_has_class(const GtkWidgetPath* v, gint _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_has_class: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_widget_path_iter_has_name(const GtkWidgetPath* v, gint _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_has_name: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_widget_path_iter_has_qclass(const GtkWidgetPath* v, gint _0, GQuark _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_has_qclass: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_widget_path_iter_has_qname(const GtkWidgetPath* v, gint _0, GQuark _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_has_qname: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_widget_path_iter_has_qregion(const GtkWidgetPath* v, gint _0, GQuark _1, GtkRegionFlags* _2) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_has_qregion: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_widget_path_iter_has_region(const GtkWidgetPath* v, gint _0, const gchar* _1, GtkRegionFlags* _2) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_has_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_window_get_has_resize_grip(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_has_resize_grip: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_window_get_resize_grip_area(GtkWindow* v, GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gtk_window_get_resize_grip_area: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gboolean gtk_window_resize_grip_is_visible(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_resize_grip_is_visible: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gchar* gtk_app_chooser_get_content_type(GtkAppChooser* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_get_content_type: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_cell_area_box_get_spacing(GtkCellAreaBox* v) __attribute__((weak)) {
	goPanic("gtk_cell_area_box_get_spacing: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_combo_box_get_id_column(GtkComboBox* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_get_id_column: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_numerable_icon_get_count(GtkNumerableIcon* v) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_get_count: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_scrolled_window_get_min_content_height(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_min_content_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_scrolled_window_get_min_content_width(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_min_content_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_widget_get_margin_bottom(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_margin_bottom: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_widget_get_margin_left(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_margin_left: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_widget_get_margin_right(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_margin_right: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_widget_get_margin_top(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_margin_top: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_widget_path_append_type(GtkWidgetPath* v, GType _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_append_type: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint gtk_widget_path_length(const GtkWidgetPath* v) __attribute__((weak)) {
	goPanic("gtk_widget_path_length: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
gint* gtk_tree_path_get_indices_with_depth(GtkTreePath* v, gint* _0) __attribute__((weak)) {
	goPanic("gtk_tree_path_get_indices_with_depth: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
guint gtk_get_binary_age(void) __attribute__((weak)) {
	goPanic("gtk_get_binary_age: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
guint gtk_get_interface_age(void) __attribute__((weak)) {
	goPanic("gtk_get_interface_age: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
guint gtk_get_major_version(void) __attribute__((weak)) {
	goPanic("gtk_get_major_version: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
guint gtk_get_micro_version(void) __attribute__((weak)) {
	goPanic("gtk_get_micro_version: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
guint gtk_get_minor_version(void) __attribute__((weak)) {
	goPanic("gtk_get_minor_version: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_about_dialog_set_license_type(GtkAboutDialog* v, GtkLicense _0) __attribute__((weak)) {
	goPanic("gtk_about_dialog_set_license_type: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_app_chooser_button_append_custom_item(GtkAppChooserButton* v, const gchar* _0, const gchar* _1, GIcon* _2) __attribute__((weak)) {
	goPanic("gtk_app_chooser_button_append_custom_item: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_app_chooser_button_append_separator(GtkAppChooserButton* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_button_append_separator: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_app_chooser_button_set_active_custom_item(GtkAppChooserButton* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_app_chooser_button_set_active_custom_item: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_app_chooser_button_set_show_dialog_item(GtkAppChooserButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_app_chooser_button_set_show_dialog_item: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_app_chooser_refresh(GtkAppChooser* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_refresh: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_app_chooser_widget_set_show_all(GtkAppChooserWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_set_show_all: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_app_chooser_widget_set_show_default(GtkAppChooserWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_set_show_default: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_app_chooser_widget_set_show_fallback(GtkAppChooserWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_set_show_fallback: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_app_chooser_widget_set_show_other(GtkAppChooserWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_set_show_other: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_app_chooser_widget_set_show_recommended(GtkAppChooserWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_app_chooser_widget_set_show_recommended: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_application_add_window(GtkApplication* v, GtkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_application_add_window: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_application_remove_window(GtkApplication* v, GtkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_application_remove_window: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_assistant_next_page(GtkAssistant* v) __attribute__((weak)) {
	goPanic("gtk_assistant_next_page: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_assistant_previous_page(GtkAssistant* v) __attribute__((weak)) {
	goPanic("gtk_assistant_previous_page: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cairo_transform_to_window(cairo_t* _0, GtkWidget* _1, GdkWindow* _2) __attribute__((weak)) {
	goPanic("gtk_cairo_transform_to_window: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_add(GtkCellArea* v, GtkCellRenderer* _0) __attribute__((weak)) {
	goPanic("gtk_cell_area_add: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_add_focus_sibling(GtkCellArea* v, GtkCellRenderer* _0, GtkCellRenderer* _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_add_focus_sibling: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_apply_attributes(GtkCellArea* v, GtkTreeModel* _0, GtkTreeIter* _1, gboolean _2, gboolean _3) __attribute__((weak)) {
	goPanic("gtk_cell_area_apply_attributes: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_attribute_connect(GtkCellArea* v, GtkCellRenderer* _0, const gchar* _1, gint _2) __attribute__((weak)) {
	goPanic("gtk_cell_area_attribute_connect: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_attribute_disconnect(GtkCellArea* v, GtkCellRenderer* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_attribute_disconnect: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_box_pack_end(GtkCellAreaBox* v, GtkCellRenderer* _0, gboolean _1, gboolean _2, gboolean _3) __attribute__((weak)) {
	goPanic("gtk_cell_area_box_pack_end: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_box_pack_start(GtkCellAreaBox* v, GtkCellRenderer* _0, gboolean _1, gboolean _2, gboolean _3) __attribute__((weak)) {
	goPanic("gtk_cell_area_box_pack_start: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_box_set_spacing(GtkCellAreaBox* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_cell_area_box_set_spacing: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_cell_get_property(GtkCellArea* v, GtkCellRenderer* _0, const gchar* _1, GValue* _2) __attribute__((weak)) {
	goPanic("gtk_cell_area_cell_get_property: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_cell_set_property(GtkCellArea* v, GtkCellRenderer* _0, const gchar* _1, const GValue* _2) __attribute__((weak)) {
	goPanic("gtk_cell_area_cell_set_property: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_context_get_allocation(GtkCellAreaContext* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_context_get_allocation: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_context_get_preferred_height(GtkCellAreaContext* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_context_get_preferred_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_context_get_preferred_height_for_width(GtkCellAreaContext* v, gint _0, gint* _1, gint* _2) __attribute__((weak)) {
	goPanic("gtk_cell_area_context_get_preferred_height_for_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_context_get_preferred_width(GtkCellAreaContext* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_context_get_preferred_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_context_get_preferred_width_for_height(GtkCellAreaContext* v, gint _0, gint* _1, gint* _2) __attribute__((weak)) {
	goPanic("gtk_cell_area_context_get_preferred_width_for_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_context_push_preferred_height(GtkCellAreaContext* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_context_push_preferred_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_context_push_preferred_width(GtkCellAreaContext* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_context_push_preferred_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_foreach(GtkCellArea* v, GtkCellCallback _0, gpointer _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_foreach: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_foreach_alloc(GtkCellArea* v, GtkCellAreaContext* _0, GtkWidget* _1, const GdkRectangle* _2, const GdkRectangle* _3, GtkCellAllocCallback _4, gpointer _5) __attribute__((weak)) {
	goPanic("gtk_cell_area_foreach_alloc: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_get_cell_allocation(GtkCellArea* v, GtkCellAreaContext* _0, GtkWidget* _1, GtkCellRenderer* _2, const GdkRectangle* _3, GdkRectangle* _4) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_cell_allocation: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_get_preferred_height(GtkCellArea* v, GtkCellAreaContext* _0, GtkWidget* _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_preferred_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_get_preferred_height_for_width(GtkCellArea* v, GtkCellAreaContext* _0, GtkWidget* _1, gint _2, gint* _3, gint* _4) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_preferred_height_for_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_get_preferred_width(GtkCellArea* v, GtkCellAreaContext* _0, GtkWidget* _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_preferred_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_get_preferred_width_for_height(GtkCellArea* v, GtkCellAreaContext* _0, GtkWidget* _1, gint _2, gint* _3, gint* _4) __attribute__((weak)) {
	goPanic("gtk_cell_area_get_preferred_width_for_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_inner_cell_area(GtkCellArea* v, GtkWidget* _0, const GdkRectangle* _1, GdkRectangle* _2) __attribute__((weak)) {
	goPanic("gtk_cell_area_inner_cell_area: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_remove(GtkCellArea* v, GtkCellRenderer* _0) __attribute__((weak)) {
	goPanic("gtk_cell_area_remove: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_remove_focus_sibling(GtkCellArea* v, GtkCellRenderer* _0, GtkCellRenderer* _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_remove_focus_sibling: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_render(GtkCellArea* v, GtkCellAreaContext* _0, GtkWidget* _1, cairo_t* _2, const GdkRectangle* _3, const GdkRectangle* _4, GtkCellRendererState _5, gboolean _6) __attribute__((weak)) {
	goPanic("gtk_cell_area_render: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_request_renderer(GtkCellArea* v, GtkCellRenderer* _0, GtkOrientation _1, GtkWidget* _2, gint _3, gint* _4, gint* _5) __attribute__((weak)) {
	goPanic("gtk_cell_area_request_renderer: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_set_focus_cell(GtkCellArea* v, GtkCellRenderer* _0) __attribute__((weak)) {
	goPanic("gtk_cell_area_set_focus_cell: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_area_stop_editing(GtkCellArea* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_cell_area_stop_editing: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_renderer_get_aligned_area(GtkCellRenderer* v, GtkWidget* _0, GtkCellRendererState _1, const GdkRectangle* _2, GdkRectangle* _3) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_aligned_area: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_renderer_get_preferred_height(GtkCellRenderer* v, GtkWidget* _0, gint* _1, gint* _2) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_preferred_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_renderer_get_preferred_height_for_width(GtkCellRenderer* v, GtkWidget* _0, gint _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_preferred_height_for_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_renderer_get_preferred_size(GtkCellRenderer* v, GtkWidget* _0, GtkRequisition* _1, GtkRequisition* _2) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_preferred_size: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_renderer_get_preferred_width(GtkCellRenderer* v, GtkWidget* _0, gint* _1, gint* _2) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_preferred_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_renderer_get_preferred_width_for_height(GtkCellRenderer* v, GtkWidget* _0, gint _1, gint* _2, gint* _3) __attribute__((weak)) {
	goPanic("gtk_cell_renderer_get_preferred_width_for_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_view_set_background_rgba(GtkCellView* v, const GdkRGBA* _0) __attribute__((weak)) {
	goPanic("gtk_cell_view_set_background_rgba: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_view_set_draw_sensitive(GtkCellView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_cell_view_set_draw_sensitive: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_cell_view_set_fit_model(GtkCellView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_cell_view_set_fit_model: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_color_selection_get_current_rgba(GtkColorSelection* v, GdkRGBA* _0) __attribute__((weak)) {
	goPanic("gtk_color_selection_get_current_rgba: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_color_selection_get_previous_rgba(GtkColorSelection* v, GdkRGBA* _0) __attribute__((weak)) {
	goPanic("gtk_color_selection_get_previous_rgba: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_color_selection_set_current_rgba(GtkColorSelection* v, const GdkRGBA* _0) __attribute__((weak)) {
	goPanic("gtk_color_selection_set_current_rgba: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_color_selection_set_previous_rgba(GtkColorSelection* v, const GdkRGBA* _0) __attribute__((weak)) {
	goPanic("gtk_color_selection_set_previous_rgba: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_combo_box_popup_for_device(GtkComboBox* v, GdkDevice* _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_popup_for_device: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_combo_box_set_id_column(GtkComboBox* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_id_column: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_combo_box_set_popup_fixed_width(GtkComboBox* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_combo_box_set_popup_fixed_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_combo_box_text_insert(GtkComboBoxText* v, gint _0, const gchar* _1, const gchar* _2) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_insert: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_combo_box_text_remove_all(GtkComboBoxText* v) __attribute__((weak)) {
	goPanic("gtk_combo_box_text_remove_all: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_device_grab_add(GtkWidget* _0, GdkDevice* _1, gboolean _2) __attribute__((weak)) {
	goPanic("gtk_device_grab_add: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_device_grab_remove(GtkWidget* _0, GdkDevice* _1) __attribute__((weak)) {
	goPanic("gtk_device_grab_remove: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_draw_insertion_cursor(GtkWidget* _0, cairo_t* _1, const GdkRectangle* _2, gboolean _3, GtkTextDirection _4, gboolean _5) __attribute__((weak)) {
	goPanic("gtk_draw_insertion_cursor: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_entry_get_icon_area(GtkEntry* v, GtkEntryIconPosition _0, GdkRectangle* _1) __attribute__((weak)) {
	goPanic("gtk_entry_get_icon_area: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_entry_get_text_area(GtkEntry* v, GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gtk_entry_get_text_area: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_gradient_add_color_stop(GtkGradient* v, gdouble _0, GtkSymbolicColor* _1) __attribute__((weak)) {
	goPanic("gtk_gradient_add_color_stop: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_menu_item_set_reserve_indicator(GtkMenuItem* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_menu_item_set_reserve_indicator: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_numerable_icon_set_background_gicon(GtkNumerableIcon* v, GIcon* _0) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_set_background_gicon: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_numerable_icon_set_background_icon_name(GtkNumerableIcon* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_set_background_icon_name: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_numerable_icon_set_count(GtkNumerableIcon* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_set_count: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_numerable_icon_set_label(GtkNumerableIcon* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_set_label: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_numerable_icon_set_style_context(GtkNumerableIcon* v, GtkStyleContext* _0) __attribute__((weak)) {
	goPanic("gtk_numerable_icon_set_style_context: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_progress_bar_set_show_text(GtkProgressBar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_progress_bar_set_show_text: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_radio_action_join_group(GtkRadioAction* v, GtkRadioAction* _0) __attribute__((weak)) {
	goPanic("gtk_radio_action_join_group: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_radio_button_join_group(GtkRadioButton* v, GtkRadioButton* _0) __attribute__((weak)) {
	goPanic("gtk_radio_button_join_group: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_activity(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_render_activity: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_arrow(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_render_arrow: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_background(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_render_background: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_check(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_render_check: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_expander(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_render_expander: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_extension(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5, GtkPositionType _6) __attribute__((weak)) {
	goPanic("gtk_render_extension: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_focus(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_render_focus: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_frame(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_render_frame: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_frame_gap(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5, GtkPositionType _6, gdouble _7, gdouble _8) __attribute__((weak)) {
	goPanic("gtk_render_frame_gap: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_handle(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_render_handle: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_layout(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, PangoLayout* _4) __attribute__((weak)) {
	goPanic("gtk_render_layout: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_line(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_render_line: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_option(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5) __attribute__((weak)) {
	goPanic("gtk_render_option: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_render_slider(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, gdouble _4, gdouble _5, GtkOrientation _6) __attribute__((weak)) {
	goPanic("gtk_render_slider: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_scrollable_set_hadjustment(GtkScrollable* v, GtkAdjustment* _0) __attribute__((weak)) {
	goPanic("gtk_scrollable_set_hadjustment: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_scrollable_set_hscroll_policy(GtkScrollable* v, GtkScrollablePolicy _0) __attribute__((weak)) {
	goPanic("gtk_scrollable_set_hscroll_policy: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_scrollable_set_vadjustment(GtkScrollable* v, GtkAdjustment* _0) __attribute__((weak)) {
	goPanic("gtk_scrollable_set_vadjustment: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_scrollable_set_vscroll_policy(GtkScrollable* v, GtkScrollablePolicy _0) __attribute__((weak)) {
	goPanic("gtk_scrollable_set_vscroll_policy: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_scrolled_window_set_min_content_height(GtkScrolledWindow* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_set_min_content_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_scrolled_window_set_min_content_width(GtkScrolledWindow* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_set_min_content_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_add_class(GtkStyleContext* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_add_class: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_add_provider(GtkStyleContext* v, GtkStyleProvider* _0, guint _1) __attribute__((weak)) {
	goPanic("gtk_style_context_add_provider: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_add_provider_for_screen(GdkScreen* _0, GtkStyleProvider* _1, guint _2) __attribute__((weak)) {
	goPanic("gtk_style_context_add_provider_for_screen: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_add_region(GtkStyleContext* v, const gchar* _0, GtkRegionFlags _1) __attribute__((weak)) {
	goPanic("gtk_style_context_add_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_cancel_animations(GtkStyleContext* v, gpointer _0) __attribute__((weak)) {
	goPanic("gtk_style_context_cancel_animations: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_get_background_color(GtkStyleContext* v, GtkStateFlags _0, GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_style_context_get_background_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_get_border(GtkStyleContext* v, GtkStateFlags _0, GtkBorder* _1) __attribute__((weak)) {
	goPanic("gtk_style_context_get_border: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_get_border_color(GtkStyleContext* v, GtkStateFlags _0, GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_style_context_get_border_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_get_color(GtkStyleContext* v, GtkStateFlags _0, GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_style_context_get_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_get_margin(GtkStyleContext* v, GtkStateFlags _0, GtkBorder* _1) __attribute__((weak)) {
	goPanic("gtk_style_context_get_margin: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_get_padding(GtkStyleContext* v, GtkStateFlags _0, GtkBorder* _1) __attribute__((weak)) {
	goPanic("gtk_style_context_get_padding: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_get_property(GtkStyleContext* v, const gchar* _0, GtkStateFlags _1, GValue* _2) __attribute__((weak)) {
	goPanic("gtk_style_context_get_property: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_invalidate(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_invalidate: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_notify_state_change(GtkStyleContext* v, GdkWindow* _0, gpointer _1, GtkStateType _2, gboolean _3) __attribute__((weak)) {
	goPanic("gtk_style_context_notify_state_change: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_pop_animatable_region(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_pop_animatable_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_push_animatable_region(GtkStyleContext* v, gpointer _0) __attribute__((weak)) {
	goPanic("gtk_style_context_push_animatable_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_remove_class(GtkStyleContext* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_remove_class: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_remove_provider(GtkStyleContext* v, GtkStyleProvider* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_remove_provider: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_remove_provider_for_screen(GdkScreen* _0, GtkStyleProvider* _1) __attribute__((weak)) {
	goPanic("gtk_style_context_remove_provider_for_screen: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_remove_region(GtkStyleContext* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_remove_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_reset_widgets(GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_reset_widgets: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_restore(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_restore: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_save(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_save: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_scroll_animations(GtkStyleContext* v, GdkWindow* _0, gint _1, gint _2) __attribute__((weak)) {
	goPanic("gtk_style_context_scroll_animations: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_set_background(GtkStyleContext* v, GdkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_set_background: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_set_direction(GtkStyleContext* v, GtkTextDirection _0) __attribute__((weak)) {
	goPanic("gtk_style_context_set_direction: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_set_junction_sides(GtkStyleContext* v, GtkJunctionSides _0) __attribute__((weak)) {
	goPanic("gtk_style_context_set_junction_sides: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_set_path(GtkStyleContext* v, GtkWidgetPath* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_set_path: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_set_screen(GtkStyleContext* v, GdkScreen* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_set_screen: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_context_set_state(GtkStyleContext* v, GtkStateFlags _0) __attribute__((weak)) {
	goPanic("gtk_style_context_set_state: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_properties_map_color(GtkStyleProperties* v, const gchar* _0, GtkSymbolicColor* _1) __attribute__((weak)) {
	goPanic("gtk_style_properties_map_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_properties_merge(GtkStyleProperties* v, const GtkStyleProperties* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_style_properties_merge: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_properties_set_property(GtkStyleProperties* v, const gchar* _0, GtkStateFlags _1, const GValue* _2) __attribute__((weak)) {
	goPanic("gtk_style_properties_set_property: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_style_properties_unset_property(GtkStyleProperties* v, const gchar* _0, GtkStateFlags _1) __attribute__((weak)) {
	goPanic("gtk_style_properties_unset_property: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_switch_set_active(GtkSwitch* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_switch_set_active: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_text_view_get_cursor_locations(GtkTextView* v, const GtkTextIter* _0, GdkRectangle* _1, GdkRectangle* _2) __attribute__((weak)) {
	goPanic("gtk_text_view_get_cursor_locations: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_theming_engine_get_background_color(GtkThemingEngine* v, GtkStateFlags _0, GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_background_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_theming_engine_get_border(GtkThemingEngine* v, GtkStateFlags _0, GtkBorder* _1) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_border: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_theming_engine_get_border_color(GtkThemingEngine* v, GtkStateFlags _0, GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_border_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_theming_engine_get_color(GtkThemingEngine* v, GtkStateFlags _0, GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_theming_engine_get_margin(GtkThemingEngine* v, GtkStateFlags _0, GtkBorder* _1) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_margin: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_theming_engine_get_padding(GtkThemingEngine* v, GtkStateFlags _0, GtkBorder* _1) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_padding: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_theming_engine_get_property(GtkThemingEngine* v, const gchar* _0, GtkStateFlags _1, GValue* _2) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_property: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_theming_engine_get_style_property(GtkThemingEngine* v, const gchar* _0, GValue* _1) __attribute__((weak)) {
	goPanic("gtk_theming_engine_get_style_property: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_add_device_events(GtkWidget* v, GdkDevice* _0, GdkEventMask _1) __attribute__((weak)) {
	goPanic("gtk_widget_add_device_events: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_draw(GtkWidget* v, cairo_t* _0) __attribute__((weak)) {
	goPanic("gtk_widget_draw: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_get_preferred_height(GtkWidget* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_widget_get_preferred_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_get_preferred_height_for_width(GtkWidget* v, gint _0, gint* _1, gint* _2) __attribute__((weak)) {
	goPanic("gtk_widget_get_preferred_height_for_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_get_preferred_size(GtkWidget* v, GtkRequisition* _0, GtkRequisition* _1) __attribute__((weak)) {
	goPanic("gtk_widget_get_preferred_size: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_get_preferred_width(GtkWidget* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_widget_get_preferred_width: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_get_preferred_width_for_height(GtkWidget* v, gint _0, gint* _1, gint* _2) __attribute__((weak)) {
	goPanic("gtk_widget_get_preferred_width_for_height: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_input_shape_combine_region(GtkWidget* v, cairo_region_t* _0) __attribute__((weak)) {
	goPanic("gtk_widget_input_shape_combine_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_override_background_color(GtkWidget* v, GtkStateFlags _0, const GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_widget_override_background_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_override_color(GtkWidget* v, GtkStateFlags _0, const GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_widget_override_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_override_cursor(GtkWidget* v, const GdkRGBA* _0, const GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_widget_override_cursor: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_override_font(GtkWidget* v, const PangoFontDescription* _0) __attribute__((weak)) {
	goPanic("gtk_widget_override_font: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_override_symbolic_color(GtkWidget* v, const gchar* _0, const GdkRGBA* _1) __attribute__((weak)) {
	goPanic("gtk_widget_override_symbolic_color: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_path_iter_add_class(GtkWidgetPath* v, gint _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_add_class: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_path_iter_add_region(GtkWidgetPath* v, gint _0, const gchar* _1, GtkRegionFlags _2) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_add_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_path_iter_clear_classes(GtkWidgetPath* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_clear_classes: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_path_iter_clear_regions(GtkWidgetPath* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_clear_regions: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_path_iter_remove_class(GtkWidgetPath* v, gint _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_remove_class: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_path_iter_remove_region(GtkWidgetPath* v, gint _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_remove_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_path_iter_set_name(GtkWidgetPath* v, gint _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_set_name: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_path_iter_set_object_type(GtkWidgetPath* v, gint _0, GType _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_set_object_type: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_path_prepend_type(GtkWidgetPath* v, GType _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_prepend_type: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_queue_draw_region(GtkWidget* v, const cairo_region_t* _0) __attribute__((weak)) {
	goPanic("gtk_widget_queue_draw_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_reset_style(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_reset_style: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_set_device_enabled(GtkWidget* v, GdkDevice* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_widget_set_device_enabled: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_set_device_events(GtkWidget* v, GdkDevice* _0, GdkEventMask _1) __attribute__((weak)) {
	goPanic("gtk_widget_set_device_events: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_set_margin_bottom(GtkWidget* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_margin_bottom: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_set_margin_left(GtkWidget* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_margin_left: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_set_margin_right(GtkWidget* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_margin_right: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_set_margin_top(GtkWidget* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_margin_top: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_set_state_flags(GtkWidget* v, GtkStateFlags _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_widget_set_state_flags: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_set_support_multidevice(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_support_multidevice: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_shape_combine_region(GtkWidget* v, cairo_region_t* _0) __attribute__((weak)) {
	goPanic("gtk_widget_shape_combine_region: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_widget_unset_state_flags(GtkWidget* v, GtkStateFlags _0) __attribute__((weak)) {
	goPanic("gtk_widget_unset_state_flags: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_window_resize_to_geometry(GtkWindow* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_window_resize_to_geometry: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_window_set_application(GtkWindow* v, GtkApplication* _0) __attribute__((weak)) {
	goPanic("gtk_window_set_application: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_window_set_default_geometry(GtkWindow* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_window_set_default_geometry: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_window_set_has_resize_grip(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_has_resize_grip: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 0))
void gtk_window_set_has_user_ref_count(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_has_user_ref_count: library too old: needs at least 3.0");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GFile* gtk_places_sidebar_get_location(GtkPlacesSidebar* v) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_get_location: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GFile* gtk_places_sidebar_get_nth_bookmark(GtkPlacesSidebar* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_get_nth_bookmark: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GSList* gtk_places_sidebar_list_shortcuts(GtkPlacesSidebar* v) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_list_shortcuts: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GdkPixbuf* gtk_icon_theme_load_icon_for_scale(GtkIconTheme* v, const gchar* _0, gint _1, gint _2, GtkIconLookupFlags _3) __attribute__((weak)) {
	goPanic("gtk_icon_theme_load_icon_for_scale: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkAdjustment* gtk_list_box_get_adjustment(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_get_adjustment: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkAlign gtk_widget_get_valign_with_baseline(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_valign_with_baseline: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkApplication* gtk_builder_get_application(GtkBuilder* v) __attribute__((weak)) {
	goPanic("gtk_builder_get_application: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkBaselinePosition gtk_box_get_baseline_position(GtkBox* v) __attribute__((weak)) {
	goPanic("gtk_box_get_baseline_position: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkBaselinePosition gtk_grid_get_row_baseline_position(GtkGrid* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_grid_get_row_baseline_position: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkBuilder* gtk_builder_new_from_file(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_builder_new_from_file: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkBuilder* gtk_builder_new_from_resource(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_builder_new_from_resource: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkBuilder* gtk_builder_new_from_string(const gchar* _0, gssize _1) __attribute__((weak)) {
	goPanic("gtk_builder_new_from_string: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkIconInfo* gtk_icon_theme_choose_icon_for_scale(GtkIconTheme* v, const gchar** _0, gint _1, gint _2, GtkIconLookupFlags _3) __attribute__((weak)) {
	goPanic("gtk_icon_theme_choose_icon_for_scale: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkIconInfo* gtk_icon_theme_lookup_by_gicon_for_scale(GtkIconTheme* v, GIcon* _0, gint _1, gint _2, GtkIconLookupFlags _3) __attribute__((weak)) {
	goPanic("gtk_icon_theme_lookup_by_gicon_for_scale: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkIconInfo* gtk_icon_theme_lookup_icon_for_scale(GtkIconTheme* v, const gchar* _0, gint _1, gint _2, GtkIconLookupFlags _3) __attribute__((weak)) {
	goPanic("gtk_icon_theme_lookup_icon_for_scale: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkListBoxRow* gtk_list_box_get_row_at_index(GtkListBox* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_list_box_get_row_at_index: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkListBoxRow* gtk_list_box_get_row_at_y(GtkListBox* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_list_box_get_row_at_y: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkListBoxRow* gtk_list_box_get_selected_row(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_get_selected_row: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkPlacesOpenFlags gtk_places_sidebar_get_open_flags(GtkPlacesSidebar* v) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_get_open_flags: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkRevealerTransitionType gtk_revealer_get_transition_type(GtkRevealer* v) __attribute__((weak)) {
	goPanic("gtk_revealer_get_transition_type: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkSelectionMode gtk_list_box_get_selection_mode(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_get_selection_mode: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkStack* gtk_stack_switcher_get_stack(GtkStackSwitcher* v) __attribute__((weak)) {
	goPanic("gtk_stack_switcher_get_stack: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkStackTransitionType gtk_stack_get_transition_type(GtkStack* v) __attribute__((weak)) {
	goPanic("gtk_stack_get_transition_type: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_button_new_from_icon_name(const gchar* _0, GtkIconSize _1) __attribute__((weak)) {
	goPanic("gtk_button_new_from_icon_name: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_header_bar_get_custom_title(GtkHeaderBar* v) __attribute__((weak)) {
	goPanic("gtk_header_bar_get_custom_title: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_header_bar_new(void) __attribute__((weak)) {
	goPanic("gtk_header_bar_new: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_image_new_from_surface(cairo_surface_t* _0) __attribute__((weak)) {
	goPanic("gtk_image_new_from_surface: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_list_box_new(void) __attribute__((weak)) {
	goPanic("gtk_list_box_new: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_list_box_row_get_header(GtkListBoxRow* v) __attribute__((weak)) {
	goPanic("gtk_list_box_row_get_header: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_list_box_row_new(void) __attribute__((weak)) {
	goPanic("gtk_list_box_row_new: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_places_sidebar_new(void) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_new: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_revealer_new(void) __attribute__((weak)) {
	goPanic("gtk_revealer_new: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_search_bar_new(void) __attribute__((weak)) {
	goPanic("gtk_search_bar_new: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_stack_get_visible_child(GtkStack* v) __attribute__((weak)) {
	goPanic("gtk_stack_get_visible_child: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_stack_new(void) __attribute__((weak)) {
	goPanic("gtk_stack_new: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
GtkWidget* gtk_stack_switcher_new(void) __attribute__((weak)) {
	goPanic("gtk_stack_switcher_new: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
PangoTabArray* gtk_entry_get_tabs(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_tabs: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
cairo_surface_t* gtk_icon_info_load_surface(GtkIconInfo* v, GdkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_surface: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
cairo_surface_t* gtk_icon_set_render_icon_surface(GtkIconSet* v, GtkStyleContext* _0, GtkIconSize _1, int _2, GdkWindow* _3) __attribute__((weak)) {
	goPanic("gtk_icon_set_render_icon_surface: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
cairo_surface_t* gtk_icon_theme_load_surface(GtkIconTheme* v, const gchar* _0, gint _1, gint _2, GdkWindow* _3, GtkIconLookupFlags _4) __attribute__((weak)) {
	goPanic("gtk_icon_theme_load_surface: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
const gchar* gtk_header_bar_get_subtitle(GtkHeaderBar* v) __attribute__((weak)) {
	goPanic("gtk_header_bar_get_subtitle: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
const gchar* gtk_header_bar_get_title(GtkHeaderBar* v) __attribute__((weak)) {
	goPanic("gtk_header_bar_get_title: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
const gchar* gtk_stack_get_visible_child_name(GtkStack* v) __attribute__((weak)) {
	goPanic("gtk_stack_get_visible_child_name: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gboolean gtk_header_bar_get_show_close_button(GtkHeaderBar* v) __attribute__((weak)) {
	goPanic("gtk_header_bar_get_show_close_button: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gboolean gtk_info_bar_get_show_close_button(GtkInfoBar* v) __attribute__((weak)) {
	goPanic("gtk_info_bar_get_show_close_button: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gboolean gtk_list_box_get_activate_on_single_click(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_get_activate_on_single_click: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gboolean gtk_places_sidebar_get_show_desktop(GtkPlacesSidebar* v) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_get_show_desktop: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gboolean gtk_revealer_get_child_revealed(GtkRevealer* v) __attribute__((weak)) {
	goPanic("gtk_revealer_get_child_revealed: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gboolean gtk_revealer_get_reveal_child(GtkRevealer* v) __attribute__((weak)) {
	goPanic("gtk_revealer_get_reveal_child: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gboolean gtk_search_bar_get_search_mode(GtkSearchBar* v) __attribute__((weak)) {
	goPanic("gtk_search_bar_get_search_mode: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gboolean gtk_search_bar_get_show_close_button(GtkSearchBar* v) __attribute__((weak)) {
	goPanic("gtk_search_bar_get_show_close_button: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gboolean gtk_stack_get_homogeneous(GtkStack* v) __attribute__((weak)) {
	goPanic("gtk_stack_get_homogeneous: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gchar* gtk_file_chooser_get_current_name(GtkFileChooser* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_current_name: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gint gtk_grid_get_baseline_row(GtkGrid* v) __attribute__((weak)) {
	goPanic("gtk_grid_get_baseline_row: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gint gtk_icon_info_get_base_scale(GtkIconInfo* v) __attribute__((weak)) {
	goPanic("gtk_icon_info_get_base_scale: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gint gtk_label_get_lines(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_lines: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gint gtk_list_box_row_get_index(GtkListBoxRow* v) __attribute__((weak)) {
	goPanic("gtk_list_box_row_get_index: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gint gtk_style_context_get_scale(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_get_scale: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
gint gtk_widget_get_scale_factor(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_scale_factor: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
guint gtk_revealer_get_transition_duration(GtkRevealer* v) __attribute__((weak)) {
	goPanic("gtk_revealer_get_transition_duration: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
guint gtk_stack_get_transition_duration(GtkStack* v) __attribute__((weak)) {
	goPanic("gtk_stack_get_transition_duration: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
int gtk_widget_get_allocated_baseline(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_allocated_baseline: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_box_set_baseline_position(GtkBox* v, GtkBaselinePosition _0) __attribute__((weak)) {
	goPanic("gtk_box_set_baseline_position: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_builder_set_application(GtkBuilder* v, GtkApplication* _0) __attribute__((weak)) {
	goPanic("gtk_builder_set_application: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_entry_set_tabs(GtkEntry* v, PangoTabArray* _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_tabs: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_grid_remove_column(GtkGrid* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_grid_remove_column: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_grid_remove_row(GtkGrid* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_grid_remove_row: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_grid_set_baseline_row(GtkGrid* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_grid_set_baseline_row: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_grid_set_row_baseline_position(GtkGrid* v, gint _0, GtkBaselinePosition _1) __attribute__((weak)) {
	goPanic("gtk_grid_set_row_baseline_position: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_header_bar_pack_end(GtkHeaderBar* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_header_bar_pack_end: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_header_bar_pack_start(GtkHeaderBar* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_header_bar_pack_start: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_header_bar_set_custom_title(GtkHeaderBar* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_header_bar_set_custom_title: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_header_bar_set_show_close_button(GtkHeaderBar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_header_bar_set_show_close_button: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_header_bar_set_subtitle(GtkHeaderBar* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_header_bar_set_subtitle: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_header_bar_set_title(GtkHeaderBar* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_header_bar_set_title: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_image_set_from_surface(GtkImage* v, cairo_surface_t* _0) __attribute__((weak)) {
	goPanic("gtk_image_set_from_surface: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_info_bar_set_show_close_button(GtkInfoBar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_info_bar_set_show_close_button: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_label_set_lines(GtkLabel* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_label_set_lines: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_drag_highlight_row(GtkListBox* v, GtkListBoxRow* _0) __attribute__((weak)) {
	goPanic("gtk_list_box_drag_highlight_row: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_drag_unhighlight_row(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_drag_unhighlight_row: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_insert(GtkListBox* v, GtkWidget* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_list_box_insert: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_invalidate_filter(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_invalidate_filter: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_invalidate_headers(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_invalidate_headers: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_invalidate_sort(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_invalidate_sort: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_prepend(GtkListBox* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_list_box_prepend: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_row_changed(GtkListBoxRow* v) __attribute__((weak)) {
	goPanic("gtk_list_box_row_changed: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_row_set_header(GtkListBoxRow* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_list_box_row_set_header: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_select_row(GtkListBox* v, GtkListBoxRow* _0) __attribute__((weak)) {
	goPanic("gtk_list_box_select_row: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_set_activate_on_single_click(GtkListBox* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_list_box_set_activate_on_single_click: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_set_adjustment(GtkListBox* v, GtkAdjustment* _0) __attribute__((weak)) {
	goPanic("gtk_list_box_set_adjustment: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_set_filter_func(GtkListBox* v, GtkListBoxFilterFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_list_box_set_filter_func: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_set_header_func(GtkListBox* v, GtkListBoxUpdateHeaderFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_list_box_set_header_func: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_set_placeholder(GtkListBox* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_list_box_set_placeholder: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_set_selection_mode(GtkListBox* v, GtkSelectionMode _0) __attribute__((weak)) {
	goPanic("gtk_list_box_set_selection_mode: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_list_box_set_sort_func(GtkListBox* v, GtkListBoxSortFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_list_box_set_sort_func: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_places_sidebar_add_shortcut(GtkPlacesSidebar* v, GFile* _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_add_shortcut: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_places_sidebar_remove_shortcut(GtkPlacesSidebar* v, GFile* _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_remove_shortcut: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_places_sidebar_set_location(GtkPlacesSidebar* v, GFile* _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_location: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_places_sidebar_set_open_flags(GtkPlacesSidebar* v, GtkPlacesOpenFlags _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_open_flags: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_places_sidebar_set_show_connect_to_server(GtkPlacesSidebar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_show_connect_to_server: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_places_sidebar_set_show_desktop(GtkPlacesSidebar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_show_desktop: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_render_icon_surface(GtkStyleContext* _0, cairo_t* _1, cairo_surface_t* _2, gdouble _3, gdouble _4) __attribute__((weak)) {
	goPanic("gtk_render_icon_surface: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_revealer_set_reveal_child(GtkRevealer* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_revealer_set_reveal_child: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_revealer_set_transition_duration(GtkRevealer* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_revealer_set_transition_duration: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_revealer_set_transition_type(GtkRevealer* v, GtkRevealerTransitionType _0) __attribute__((weak)) {
	goPanic("gtk_revealer_set_transition_type: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_search_bar_connect_entry(GtkSearchBar* v, GtkEntry* _0) __attribute__((weak)) {
	goPanic("gtk_search_bar_connect_entry: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_search_bar_set_search_mode(GtkSearchBar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_search_bar_set_search_mode: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_search_bar_set_show_close_button(GtkSearchBar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_search_bar_set_show_close_button: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_stack_add_named(GtkStack* v, GtkWidget* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_stack_add_named: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_stack_add_titled(GtkStack* v, GtkWidget* _0, const gchar* _1, const gchar* _2) __attribute__((weak)) {
	goPanic("gtk_stack_add_titled: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_stack_set_homogeneous(GtkStack* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_stack_set_homogeneous: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_stack_set_transition_duration(GtkStack* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_stack_set_transition_duration: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_stack_set_transition_type(GtkStack* v, GtkStackTransitionType _0) __attribute__((weak)) {
	goPanic("gtk_stack_set_transition_type: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_stack_set_visible_child(GtkStack* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_stack_set_visible_child: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_stack_set_visible_child_full(GtkStack* v, const gchar* _0, GtkStackTransitionType _1) __attribute__((weak)) {
	goPanic("gtk_stack_set_visible_child_full: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_stack_set_visible_child_name(GtkStack* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_stack_set_visible_child_name: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_stack_switcher_set_stack(GtkStackSwitcher* v, GtkStack* _0) __attribute__((weak)) {
	goPanic("gtk_stack_switcher_set_stack: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_style_context_set_scale(GtkStyleContext* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_style_context_set_scale: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_test_widget_wait_for_draw(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_test_widget_wait_for_draw: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_tree_model_rows_reordered_with_length(GtkTreeModel* v, GtkTreePath* _0, GtkTreeIter* _1, gint* _2, gint _3) __attribute__((weak)) {
	goPanic("gtk_tree_model_rows_reordered_with_length: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_widget_get_preferred_height_and_baseline_for_width(GtkWidget* v, gint _0, gint* _1, gint* _2, gint* _3, gint* _4) __attribute__((weak)) {
	goPanic("gtk_widget_get_preferred_height_and_baseline_for_width: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_widget_init_template(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_init_template: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_widget_size_allocate_with_baseline(GtkWidget* v, GtkAllocation* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_widget_size_allocate_with_baseline: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_window_close(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_close: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 10))
void gtk_window_set_titlebar(GtkWindow* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_window_set_titlebar: library too old: needs at least 3.10");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GList* gtk_flow_box_get_selected_children(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_get_selected_children: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkFlowBoxChild* gtk_flow_box_get_child_at_index(GtkFlowBox* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_get_child_at_index: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkPopover* gtk_menu_button_get_popover(GtkMenuButton* v) __attribute__((weak)) {
	goPanic("gtk_menu_button_get_popover: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkSelectionMode gtk_flow_box_get_selection_mode(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_get_selection_mode: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkTextDirection gtk_get_locale_direction(void) __attribute__((weak)) {
	goPanic("gtk_get_locale_direction: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkTreePath* gtk_tree_path_new_from_indicesv(gint* _0, gsize _1) __attribute__((weak)) {
	goPanic("gtk_tree_path_new_from_indicesv: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_action_bar_get_center_widget(GtkActionBar* v) __attribute__((weak)) {
	goPanic("gtk_action_bar_get_center_widget: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_action_bar_new(void) __attribute__((weak)) {
	goPanic("gtk_action_bar_new: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_box_get_center_widget(GtkBox* v) __attribute__((weak)) {
	goPanic("gtk_box_get_center_widget: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_dialog_get_header_bar(GtkDialog* v) __attribute__((weak)) {
	goPanic("gtk_dialog_get_header_bar: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_flow_box_child_new(void) __attribute__((weak)) {
	goPanic("gtk_flow_box_child_new: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_flow_box_new(void) __attribute__((weak)) {
	goPanic("gtk_flow_box_new: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_popover_get_relative_to(GtkPopover* v) __attribute__((weak)) {
	goPanic("gtk_popover_get_relative_to: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_popover_new(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_popover_new: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_popover_new_from_model(GtkWidget* _0, GMenuModel* _1) __attribute__((weak)) {
	goPanic("gtk_popover_new_from_model: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
GtkWidget* gtk_stack_get_child_by_name(GtkStack* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_stack_get_child_by_name: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
const gchar* gtk_header_bar_get_decoration_layout(GtkHeaderBar* v) __attribute__((weak)) {
	goPanic("gtk_header_bar_get_decoration_layout: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gboolean gtk_flow_box_child_is_selected(GtkFlowBoxChild* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_child_is_selected: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gboolean gtk_flow_box_get_activate_on_single_click(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_get_activate_on_single_click: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gboolean gtk_flow_box_get_homogeneous(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_get_homogeneous: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gboolean gtk_header_bar_get_has_subtitle(GtkHeaderBar* v) __attribute__((weak)) {
	goPanic("gtk_header_bar_get_has_subtitle: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gboolean gtk_icon_info_is_symbolic(GtkIconInfo* v) __attribute__((weak)) {
	goPanic("gtk_icon_info_is_symbolic: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gboolean gtk_menu_button_get_use_popover(GtkMenuButton* v) __attribute__((weak)) {
	goPanic("gtk_menu_button_get_use_popover: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gboolean gtk_places_sidebar_get_local_only(GtkPlacesSidebar* v) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_get_local_only: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gboolean gtk_popover_get_modal(GtkPopover* v) __attribute__((weak)) {
	goPanic("gtk_popover_get_modal: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gboolean gtk_stack_get_transition_running(GtkStack* v) __attribute__((weak)) {
	goPanic("gtk_stack_get_transition_running: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gboolean gtk_window_is_maximized(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_is_maximized: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gchar** gtk_application_get_accels_for_action(GtkApplication* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_application_get_accels_for_action: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gchar** gtk_application_list_action_descriptions(GtkApplication* v) __attribute__((weak)) {
	goPanic("gtk_application_list_action_descriptions: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gint gtk_entry_get_max_width_chars(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_max_width_chars: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gint gtk_flow_box_child_get_index(GtkFlowBoxChild* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_child_get_index: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gint gtk_widget_get_margin_end(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_margin_end: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
gint gtk_widget_get_margin_start(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_margin_start: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
guint gtk_flow_box_get_column_spacing(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_get_column_spacing: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
guint gtk_flow_box_get_max_children_per_line(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_get_max_children_per_line: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
guint gtk_flow_box_get_min_children_per_line(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_get_min_children_per_line: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
guint gtk_flow_box_get_row_spacing(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_get_row_spacing: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_accel_label_get_accel(GtkAccelLabel* v, guint* _0, GdkModifierType* _1) __attribute__((weak)) {
	goPanic("gtk_accel_label_get_accel: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_action_bar_pack_end(GtkActionBar* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_action_bar_pack_end: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_action_bar_pack_start(GtkActionBar* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_action_bar_pack_start: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_action_bar_set_center_widget(GtkActionBar* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_action_bar_set_center_widget: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_application_set_accels_for_action(GtkApplication* v, const gchar* _0, const gchar* const* _1) __attribute__((weak)) {
	goPanic("gtk_application_set_accels_for_action: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_box_set_center_widget(GtkBox* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_box_set_center_widget: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_entry_set_max_width_chars(GtkEntry* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_max_width_chars: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_child_changed(GtkFlowBoxChild* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_child_changed: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_insert(GtkFlowBox* v, GtkWidget* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_flow_box_insert: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_invalidate_filter(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_invalidate_filter: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_invalidate_sort(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_invalidate_sort: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_select_all(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_select_all: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_select_child(GtkFlowBox* v, GtkFlowBoxChild* _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_select_child: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_selected_foreach(GtkFlowBox* v, GtkFlowBoxForeachFunc _0, gpointer _1) __attribute__((weak)) {
	goPanic("gtk_flow_box_selected_foreach: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_activate_on_single_click(GtkFlowBox* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_activate_on_single_click: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_column_spacing(GtkFlowBox* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_column_spacing: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_filter_func(GtkFlowBox* v, GtkFlowBoxFilterFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_filter_func: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_hadjustment(GtkFlowBox* v, GtkAdjustment* _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_hadjustment: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_homogeneous(GtkFlowBox* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_homogeneous: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_max_children_per_line(GtkFlowBox* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_max_children_per_line: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_min_children_per_line(GtkFlowBox* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_min_children_per_line: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_row_spacing(GtkFlowBox* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_row_spacing: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_selection_mode(GtkFlowBox* v, GtkSelectionMode _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_selection_mode: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_sort_func(GtkFlowBox* v, GtkFlowBoxSortFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_sort_func: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_set_vadjustment(GtkFlowBox* v, GtkAdjustment* _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_set_vadjustment: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_unselect_all(GtkFlowBox* v) __attribute__((weak)) {
	goPanic("gtk_flow_box_unselect_all: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_flow_box_unselect_child(GtkFlowBox* v, GtkFlowBoxChild* _0) __attribute__((weak)) {
	goPanic("gtk_flow_box_unselect_child: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_header_bar_set_decoration_layout(GtkHeaderBar* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_header_bar_set_decoration_layout: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_header_bar_set_has_subtitle(GtkHeaderBar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_header_bar_set_has_subtitle: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_menu_button_set_popover(GtkMenuButton* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_menu_button_set_popover: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_menu_button_set_use_popover(GtkMenuButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_menu_button_set_use_popover: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_places_sidebar_set_local_only(GtkPlacesSidebar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_local_only: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_popover_bind_model(GtkPopover* v, GMenuModel* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_popover_bind_model: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_popover_set_modal(GtkPopover* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_popover_set_modal: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_popover_set_pointing_to(GtkPopover* v, const GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gtk_popover_set_pointing_to: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_popover_set_position(GtkPopover* v, GtkPositionType _0) __attribute__((weak)) {
	goPanic("gtk_popover_set_position: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_popover_set_relative_to(GtkPopover* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_popover_set_relative_to: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_widget_set_margin_end(GtkWidget* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_margin_end: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 12))
void gtk_widget_set_margin_start(GtkWidget* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_margin_start: library too old: needs at least 3.12");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GList* gtk_gesture_get_group(GtkGesture* v) __attribute__((weak)) {
	goPanic("gtk_gesture_get_group: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GList* gtk_gesture_get_sequences(GtkGesture* v) __attribute__((weak)) {
	goPanic("gtk_gesture_get_sequences: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GList* gtk_list_box_get_selected_rows(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_get_selected_rows: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GMenu* gtk_application_get_menu_by_id(GtkApplication* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_application_get_menu_by_id: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GdkDevice* gtk_gesture_get_device(GtkGesture* v) __attribute__((weak)) {
	goPanic("gtk_gesture_get_device: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GdkEventSequence* gtk_gesture_get_last_updated_sequence(GtkGesture* v) __attribute__((weak)) {
	goPanic("gtk_gesture_get_last_updated_sequence: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GdkEventSequence* gtk_gesture_single_get_current_sequence(GtkGestureSingle* v) __attribute__((weak)) {
	goPanic("gtk_gesture_single_get_current_sequence: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GdkWindow* gtk_gesture_get_window(GtkGesture* v) __attribute__((weak)) {
	goPanic("gtk_gesture_get_window: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkEventSequenceState gtk_gesture_get_sequence_state(GtkGesture* v, GdkEventSequence* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_get_sequence_state: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkGesture* gtk_gesture_drag_new(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_drag_new: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkGesture* gtk_gesture_long_press_new(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_long_press_new: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkGesture* gtk_gesture_multi_press_new(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_multi_press_new: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkGesture* gtk_gesture_pan_new(GtkWidget* _0, GtkOrientation _1) __attribute__((weak)) {
	goPanic("gtk_gesture_pan_new: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkGesture* gtk_gesture_rotate_new(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_rotate_new: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkGesture* gtk_gesture_swipe_new(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_swipe_new: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkGesture* gtk_gesture_zoom_new(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_zoom_new: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkOrientation gtk_gesture_pan_get_orientation(GtkGesturePan* v) __attribute__((weak)) {
	goPanic("gtk_gesture_pan_get_orientation: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkPropagationPhase gtk_event_controller_get_propagation_phase(GtkEventController* v) __attribute__((weak)) {
	goPanic("gtk_event_controller_get_propagation_phase: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkStateFlags gtk_widget_path_iter_get_state(const GtkWidgetPath* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_get_state: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
GtkWidget* gtk_event_controller_get_widget(GtkEventController* v) __attribute__((weak)) {
	goPanic("gtk_event_controller_get_widget: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_application_prefers_app_menu(GtkApplication* v) __attribute__((weak)) {
	goPanic("gtk_application_prefers_app_menu: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_drag_get_offset(GtkGestureDrag* v, gdouble* _0, gdouble* _1) __attribute__((weak)) {
	goPanic("gtk_gesture_drag_get_offset: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_drag_get_start_point(GtkGestureDrag* v, gdouble* _0, gdouble* _1) __attribute__((weak)) {
	goPanic("gtk_gesture_drag_get_start_point: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_get_bounding_box(GtkGesture* v, GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_get_bounding_box: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_get_bounding_box_center(GtkGesture* v, gdouble* _0, gdouble* _1) __attribute__((weak)) {
	goPanic("gtk_gesture_get_bounding_box_center: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_get_point(GtkGesture* v, GdkEventSequence* _0, gdouble* _1, gdouble* _2) __attribute__((weak)) {
	goPanic("gtk_gesture_get_point: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_handles_sequence(GtkGesture* v, GdkEventSequence* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_handles_sequence: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_is_active(GtkGesture* v) __attribute__((weak)) {
	goPanic("gtk_gesture_is_active: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_is_grouped_with(GtkGesture* v, GtkGesture* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_is_grouped_with: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_is_recognized(GtkGesture* v) __attribute__((weak)) {
	goPanic("gtk_gesture_is_recognized: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_multi_press_get_area(GtkGestureMultiPress* v, GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_multi_press_get_area: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_set_sequence_state(GtkGesture* v, GdkEventSequence* _0, GtkEventSequenceState _1) __attribute__((weak)) {
	goPanic("gtk_gesture_set_sequence_state: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_set_state(GtkGesture* v, GtkEventSequenceState _0) __attribute__((weak)) {
	goPanic("gtk_gesture_set_state: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_single_get_exclusive(GtkGestureSingle* v) __attribute__((weak)) {
	goPanic("gtk_gesture_single_get_exclusive: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_single_get_touch_only(GtkGestureSingle* v) __attribute__((weak)) {
	goPanic("gtk_gesture_single_get_touch_only: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_gesture_swipe_get_velocity(GtkGestureSwipe* v, gdouble* _0, gdouble* _1) __attribute__((weak)) {
	goPanic("gtk_gesture_swipe_get_velocity: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_list_box_row_get_activatable(GtkListBoxRow* v) __attribute__((weak)) {
	goPanic("gtk_list_box_row_get_activatable: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_list_box_row_get_selectable(GtkListBoxRow* v) __attribute__((weak)) {
	goPanic("gtk_list_box_row_get_selectable: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_list_box_row_is_selected(GtkListBoxRow* v) __attribute__((weak)) {
	goPanic("gtk_list_box_row_is_selected: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_places_sidebar_get_show_enter_location(GtkPlacesSidebar* v) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_get_show_enter_location: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gboolean gtk_switch_get_state(GtkSwitch* v) __attribute__((weak)) {
	goPanic("gtk_switch_get_state: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gchar** gtk_application_get_actions_for_accel(GtkApplication* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_application_get_actions_for_accel: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gdouble gtk_gesture_rotate_get_angle_delta(GtkGestureRotate* v) __attribute__((weak)) {
	goPanic("gtk_gesture_rotate_get_angle_delta: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gdouble gtk_gesture_zoom_get_scale_delta(GtkGestureZoom* v) __attribute__((weak)) {
	goPanic("gtk_gesture_zoom_get_scale_delta: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
gint gtk_cell_area_attribute_get_column(GtkCellArea* v, GtkCellRenderer* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_cell_area_attribute_get_column: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
guint gtk_gesture_single_get_button(GtkGestureSingle* v) __attribute__((weak)) {
	goPanic("gtk_gesture_single_get_button: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
guint gtk_gesture_single_get_current_button(GtkGestureSingle* v) __attribute__((weak)) {
	goPanic("gtk_gesture_single_get_current_button: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_event_controller_reset(GtkEventController* v) __attribute__((weak)) {
	goPanic("gtk_event_controller_reset: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_event_controller_set_propagation_phase(GtkEventController* v, GtkPropagationPhase _0) __attribute__((weak)) {
	goPanic("gtk_event_controller_set_propagation_phase: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_gesture_group(GtkGesture* v, GtkGesture* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_group: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_gesture_multi_press_set_area(GtkGestureMultiPress* v, const GdkRectangle* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_multi_press_set_area: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_gesture_pan_set_orientation(GtkGesturePan* v, GtkOrientation _0) __attribute__((weak)) {
	goPanic("gtk_gesture_pan_set_orientation: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_gesture_set_window(GtkGesture* v, GdkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_set_window: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_gesture_single_set_button(GtkGestureSingle* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_gesture_single_set_button: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_gesture_single_set_exclusive(GtkGestureSingle* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_gesture_single_set_exclusive: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_gesture_single_set_touch_only(GtkGestureSingle* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_gesture_single_set_touch_only: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_gesture_ungroup(GtkGesture* v) __attribute__((weak)) {
	goPanic("gtk_gesture_ungroup: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_icon_theme_add_resource_path(GtkIconTheme* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_icon_theme_add_resource_path: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_list_box_row_set_activatable(GtkListBoxRow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_list_box_row_set_activatable: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_list_box_row_set_selectable(GtkListBoxRow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_list_box_row_set_selectable: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_list_box_select_all(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_select_all: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_list_box_selected_foreach(GtkListBox* v, GtkListBoxForeachFunc _0, gpointer _1) __attribute__((weak)) {
	goPanic("gtk_list_box_selected_foreach: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_list_box_unselect_all(GtkListBox* v) __attribute__((weak)) {
	goPanic("gtk_list_box_unselect_all: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_list_box_unselect_row(GtkListBox* v, GtkListBoxRow* _0) __attribute__((weak)) {
	goPanic("gtk_list_box_unselect_row: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_places_sidebar_set_show_enter_location(GtkPlacesSidebar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_show_enter_location: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_switch_set_state(GtkSwitch* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_switch_set_state: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_widget_get_clip(GtkWidget* v, GtkAllocation* _0) __attribute__((weak)) {
	goPanic("gtk_widget_get_clip: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_widget_path_iter_set_state(GtkWidgetPath* v, gint _0, GtkStateFlags _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_set_state: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_widget_set_clip(GtkWidget* v, const GtkAllocation* _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_clip: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 14))
void gtk_window_set_interactive_debugging(gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_interactive_debugging: library too old: needs at least 3.14");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GActionGroup* gtk_widget_get_action_group(GtkWidget* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_widget_get_action_group: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GError* gtk_gl_area_get_error(GtkGLArea* v) __attribute__((weak)) {
	goPanic("gtk_gl_area_get_error: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GdkGLContext* gtk_gl_area_get_context(GtkGLArea* v) __attribute__((weak)) {
	goPanic("gtk_gl_area_get_context: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GtkClipboard* gtk_clipboard_get_default(GdkDisplay* _0) __attribute__((weak)) {
	goPanic("gtk_clipboard_get_default: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GtkPaperSize* gtk_paper_size_new_from_ipp(const gchar* _0, gdouble _1, gdouble _2) __attribute__((weak)) {
	goPanic("gtk_paper_size_new_from_ipp: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GtkStack* gtk_stack_sidebar_get_stack(GtkStackSidebar* v) __attribute__((weak)) {
	goPanic("gtk_stack_sidebar_get_stack: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GtkWidget* gtk_gl_area_new(void) __attribute__((weak)) {
	goPanic("gtk_gl_area_new: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GtkWidget* gtk_model_button_new(void) __attribute__((weak)) {
	goPanic("gtk_model_button_new: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GtkWidget* gtk_popover_menu_new(void) __attribute__((weak)) {
	goPanic("gtk_popover_menu_new: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GtkWidget* gtk_stack_sidebar_new(void) __attribute__((weak)) {
	goPanic("gtk_stack_sidebar_new: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
GtkWidget* gtk_window_get_titlebar(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_titlebar: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
const gchar** gtk_widget_list_action_prefixes(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_list_action_prefixes: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_gl_area_get_auto_render(GtkGLArea* v) __attribute__((weak)) {
	goPanic("gtk_gl_area_get_auto_render: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_gl_area_get_has_alpha(GtkGLArea* v) __attribute__((weak)) {
	goPanic("gtk_gl_area_get_has_alpha: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_gl_area_get_has_depth_buffer(GtkGLArea* v) __attribute__((weak)) {
	goPanic("gtk_gl_area_get_has_depth_buffer: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_gl_area_get_has_stencil_buffer(GtkGLArea* v) __attribute__((weak)) {
	goPanic("gtk_gl_area_get_has_stencil_buffer: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_paned_get_wide_handle(GtkPaned* v) __attribute__((weak)) {
	goPanic("gtk_paned_get_wide_handle: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_popover_get_transitions_enabled(GtkPopover* v) __attribute__((weak)) {
	goPanic("gtk_popover_get_transitions_enabled: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_scrollable_get_border(GtkScrollable* v, GtkBorder* _0) __attribute__((weak)) {
	goPanic("gtk_scrollable_get_border: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_scrolled_window_get_overlay_scrolling(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_overlay_scrolling: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_stack_get_hhomogeneous(GtkStack* v) __attribute__((weak)) {
	goPanic("gtk_stack_get_hhomogeneous: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_stack_get_vhomogeneous(GtkStack* v) __attribute__((weak)) {
	goPanic("gtk_stack_get_vhomogeneous: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gboolean gtk_text_view_get_monospace(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_get_monospace: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gfloat gtk_label_get_xalign(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_xalign: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
gfloat gtk_label_get_yalign(GtkLabel* v) __attribute__((weak)) {
	goPanic("gtk_label_get_yalign: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_css_provider_load_from_resource(GtkCssProvider* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_css_provider_load_from_resource: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_drag_cancel(GdkDragContext* _0) __attribute__((weak)) {
	goPanic("gtk_drag_cancel: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_entry_grab_focus_without_selecting(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_grab_focus_without_selecting: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_gl_area_attach_buffers(GtkGLArea* v) __attribute__((weak)) {
	goPanic("gtk_gl_area_attach_buffers: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_gl_area_get_required_version(GtkGLArea* v, gint* _0, gint* _1) __attribute__((weak)) {
	goPanic("gtk_gl_area_get_required_version: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_gl_area_make_current(GtkGLArea* v) __attribute__((weak)) {
	goPanic("gtk_gl_area_make_current: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_gl_area_queue_render(GtkGLArea* v) __attribute__((weak)) {
	goPanic("gtk_gl_area_queue_render: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_gl_area_set_auto_render(GtkGLArea* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_gl_area_set_auto_render: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_gl_area_set_error(GtkGLArea* v, const GError* _0) __attribute__((weak)) {
	goPanic("gtk_gl_area_set_error: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_gl_area_set_has_alpha(GtkGLArea* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_gl_area_set_has_alpha: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_gl_area_set_has_depth_buffer(GtkGLArea* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_gl_area_set_has_depth_buffer: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_gl_area_set_has_stencil_buffer(GtkGLArea* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_gl_area_set_has_stencil_buffer: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_gl_area_set_required_version(GtkGLArea* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_gl_area_set_required_version: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_label_set_xalign(GtkLabel* v, gfloat _0) __attribute__((weak)) {
	goPanic("gtk_label_set_xalign: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_label_set_yalign(GtkLabel* v, gfloat _0) __attribute__((weak)) {
	goPanic("gtk_label_set_yalign: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_list_box_bind_model(GtkListBox* v, GListModel* _0, GtkListBoxCreateWidgetFunc _1, gpointer _2, GDestroyNotify _3) __attribute__((weak)) {
	goPanic("gtk_list_box_bind_model: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_notebook_detach_tab(GtkNotebook* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_notebook_detach_tab: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_paned_set_wide_handle(GtkPaned* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_paned_set_wide_handle: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_popover_menu_open_submenu(GtkPopoverMenu* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_popover_menu_open_submenu: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_popover_set_transitions_enabled(GtkPopover* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_popover_set_transitions_enabled: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_scrolled_window_set_overlay_scrolling(GtkScrolledWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_set_overlay_scrolling: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_stack_set_hhomogeneous(GtkStack* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_stack_set_hhomogeneous: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_stack_set_vhomogeneous(GtkStack* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_stack_set_vhomogeneous: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_stack_sidebar_set_stack(GtkStackSidebar* v, GtkStack* _0) __attribute__((weak)) {
	goPanic("gtk_stack_sidebar_set_stack: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_text_buffer_insert_markup(GtkTextBuffer* v, GtkTextIter* _0, const gchar* _1, gint _2) __attribute__((weak)) {
	goPanic("gtk_text_buffer_insert_markup: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 16))
void gtk_text_view_set_monospace(GtkTextView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_text_view_set_monospace: library too old: needs at least 3.16");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
GtkWidget* gtk_popover_get_default_widget(GtkPopover* v) __attribute__((weak)) {
	goPanic("gtk_popover_get_default_widget: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
PangoFontMap* gtk_font_chooser_get_font_map(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_font_map: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
PangoFontMap* gtk_widget_get_font_map(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_font_map: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
const cairo_font_options_t* gtk_widget_get_font_options(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_font_options: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
gboolean gtk_assistant_get_page_has_padding(GtkAssistant* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_assistant_get_page_has_padding: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
gboolean gtk_overlay_get_overlay_pass_through(GtkOverlay* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_overlay_get_overlay_pass_through: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
gboolean gtk_places_sidebar_get_show_other_locations(GtkPlacesSidebar* v) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_get_show_other_locations: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
gboolean gtk_places_sidebar_get_show_recent(GtkPlacesSidebar* v) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_get_show_recent: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
gboolean gtk_places_sidebar_get_show_trash(GtkPlacesSidebar* v) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_get_show_trash: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
gboolean gtk_stack_get_interpolate_size(GtkStack* v) __attribute__((weak)) {
	goPanic("gtk_stack_get_interpolate_size: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
gint gtk_text_view_get_bottom_margin(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_get_bottom_margin: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
gint gtk_text_view_get_top_margin(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_get_top_margin: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_assistant_set_page_has_padding(GtkAssistant* v, GtkWidget* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_assistant_set_page_has_padding: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_flow_box_bind_model(GtkFlowBox* v, GListModel* _0, GtkFlowBoxCreateWidgetFunc _1, gpointer _2, GDestroyNotify _3) __attribute__((weak)) {
	goPanic("gtk_flow_box_bind_model: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_font_chooser_set_font_map(GtkFontChooser* v, PangoFontMap* _0) __attribute__((weak)) {
	goPanic("gtk_font_chooser_set_font_map: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_overlay_reorder_overlay(GtkOverlay* v, GtkWidget* _0, int _1) __attribute__((weak)) {
	goPanic("gtk_overlay_reorder_overlay: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_overlay_set_overlay_pass_through(GtkOverlay* v, GtkWidget* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_overlay_set_overlay_pass_through: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_places_sidebar_set_drop_targets_visible(GtkPlacesSidebar* v, gboolean _0, GdkDragContext* _1) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_drop_targets_visible: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_places_sidebar_set_show_other_locations(GtkPlacesSidebar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_show_other_locations: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_places_sidebar_set_show_recent(GtkPlacesSidebar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_show_recent: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_places_sidebar_set_show_trash(GtkPlacesSidebar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_show_trash: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_popover_set_default_widget(GtkPopover* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_popover_set_default_widget: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_radio_menu_item_join_group(GtkRadioMenuItem* v, GtkRadioMenuItem* _0) __attribute__((weak)) {
	goPanic("gtk_radio_menu_item_join_group: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_stack_set_interpolate_size(GtkStack* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_stack_set_interpolate_size: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_text_view_set_bottom_margin(GtkTextView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_text_view_set_bottom_margin: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_text_view_set_top_margin(GtkTextView* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_text_view_set_top_margin: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_widget_set_font_map(GtkWidget* v, PangoFontMap* _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_font_map: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_widget_set_font_options(GtkWidget* v, const cairo_font_options_t* _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_font_options: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 18))
void gtk_window_fullscreen_on_monitor(GtkWindow* v, GdkScreen* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_window_fullscreen_on_monitor: library too old: needs at least 3.18");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
GFile* gtk_css_section_get_file(const GtkCssSection* v) __attribute__((weak)) {
	goPanic("gtk_css_section_get_file: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
GPermission* gtk_lock_button_get_permission(GtkLockButton* v) __attribute__((weak)) {
	goPanic("gtk_lock_button_get_permission: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
GtkCssSection* gtk_css_section_get_parent(const GtkCssSection* v) __attribute__((weak)) {
	goPanic("gtk_css_section_get_parent: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
GtkCssSectionType gtk_css_section_get_section_type(const GtkCssSection* v) __attribute__((weak)) {
	goPanic("gtk_css_section_get_section_type: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
GtkWidget* gtk_font_chooser_dialog_new(const gchar* _0, GtkWindow* _1) __attribute__((weak)) {
	goPanic("gtk_font_chooser_dialog_new: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
GtkWidget* gtk_font_chooser_widget_new(void) __attribute__((weak)) {
	goPanic("gtk_font_chooser_widget_new: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
GtkWidget* gtk_grid_get_child_at(GtkGrid* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_grid_get_child_at: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
GtkWidget* gtk_lock_button_new(GPermission* _0) __attribute__((weak)) {
	goPanic("gtk_lock_button_new: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
GtkWidget* gtk_overlay_new(void) __attribute__((weak)) {
	goPanic("gtk_overlay_new: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
PangoFontDescription* gtk_font_chooser_get_font_desc(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_font_desc: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
PangoFontFace* gtk_font_chooser_get_font_face(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_font_face: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
PangoFontFamily* gtk_font_chooser_get_font_family(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_font_family: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
char* gtk_css_provider_to_string(GtkCssProvider* v) __attribute__((weak)) {
	goPanic("gtk_css_provider_to_string: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
char* gtk_widget_path_to_string(const GtkWidgetPath* v) __attribute__((weak)) {
	goPanic("gtk_widget_path_to_string: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
const gchar* gtk_entry_get_placeholder_text(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_placeholder_text: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gboolean gtk_app_chooser_button_get_show_default_item(GtkAppChooserButton* v) __attribute__((weak)) {
	goPanic("gtk_app_chooser_button_get_show_default_item: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gboolean gtk_button_box_get_child_non_homogeneous(GtkButtonBox* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_button_box_get_child_non_homogeneous: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gboolean gtk_expander_get_resize_toplevel(GtkExpander* v) __attribute__((weak)) {
	goPanic("gtk_expander_get_resize_toplevel: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gboolean gtk_font_chooser_get_show_preview_entry(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_show_preview_entry: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gboolean gtk_widget_has_visible_focus(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_has_visible_focus: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gboolean gtk_window_get_focus_visible(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_focus_visible: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gchar* gtk_font_chooser_get_font(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_font: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gchar* gtk_font_chooser_get_preview_text(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_preview_text: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gdouble gtk_adjustment_get_minimum_increment(GtkAdjustment* v) __attribute__((weak)) {
	goPanic("gtk_adjustment_get_minimum_increment: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gint gtk_font_chooser_get_font_size(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_font_size: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gint gtk_tree_view_column_get_x_offset(GtkTreeViewColumn* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_column_get_x_offset: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gint gtk_widget_path_append_for_widget(GtkWidgetPath* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_append_for_widget: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
gint gtk_widget_path_append_with_siblings(GtkWidgetPath* v, GtkWidgetPath* _0, guint _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_append_with_siblings: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
guint gtk_css_section_get_end_line(const GtkCssSection* v) __attribute__((weak)) {
	goPanic("gtk_css_section_get_end_line: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
guint gtk_css_section_get_end_position(const GtkCssSection* v) __attribute__((weak)) {
	goPanic("gtk_css_section_get_end_position: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
guint gtk_css_section_get_start_line(const GtkCssSection* v) __attribute__((weak)) {
	goPanic("gtk_css_section_get_start_line: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
guint gtk_css_section_get_start_position(const GtkCssSection* v) __attribute__((weak)) {
	goPanic("gtk_css_section_get_start_position: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_app_chooser_button_set_show_default_item(GtkAppChooserButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_app_chooser_button_set_show_default_item: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_assistant_remove_page(GtkAssistant* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_assistant_remove_page: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_button_box_set_child_non_homogeneous(GtkButtonBox* v, GtkWidget* _0, gboolean _1) __attribute__((weak)) {
	goPanic("gtk_button_box_set_child_non_homogeneous: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_container_child_notify(GtkContainer* v, GtkWidget* _0, const gchar* _1) __attribute__((weak)) {
	goPanic("gtk_container_child_notify: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_drag_set_icon_gicon(GdkDragContext* _0, GIcon* _1, gint _2, gint _3) __attribute__((weak)) {
	goPanic("gtk_drag_set_icon_gicon: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_drag_source_set_icon_gicon(GtkWidget* v, GIcon* _0) __attribute__((weak)) {
	goPanic("gtk_drag_source_set_icon_gicon: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_entry_set_placeholder_text(GtkEntry* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_placeholder_text: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_expander_set_resize_toplevel(GtkExpander* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_expander_set_resize_toplevel: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_font_chooser_set_filter_func(GtkFontChooser* v, GtkFontFilterFunc _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_font_chooser_set_filter_func: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_font_chooser_set_font(GtkFontChooser* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_font_chooser_set_font: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_font_chooser_set_font_desc(GtkFontChooser* v, const PangoFontDescription* _0) __attribute__((weak)) {
	goPanic("gtk_font_chooser_set_font_desc: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_font_chooser_set_preview_text(GtkFontChooser* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_font_chooser_set_preview_text: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_font_chooser_set_show_preview_entry(GtkFontChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_font_chooser_set_show_preview_entry: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_grid_insert_column(GtkGrid* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_grid_insert_column: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_grid_insert_next_to(GtkGrid* v, GtkWidget* _0, GtkPositionType _1) __attribute__((weak)) {
	goPanic("gtk_grid_insert_next_to: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_grid_insert_row(GtkGrid* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_grid_insert_row: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_lock_button_set_permission(GtkLockButton* v, GPermission* _0) __attribute__((weak)) {
	goPanic("gtk_lock_button_set_permission: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_overlay_add_overlay(GtkOverlay* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_overlay_add_overlay: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_render_icon(GtkStyleContext* _0, cairo_t* _1, GdkPixbuf* _2, gdouble _3, gdouble _4) __attribute__((weak)) {
	goPanic("gtk_render_icon: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_text_iter_assign(GtkTextIter* v, const GtkTextIter* _0) __attribute__((weak)) {
	goPanic("gtk_text_iter_assign: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 2))
void gtk_window_set_focus_visible(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_focus_visible: library too old: needs at least 3.2");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
GtkFileChooserNative* gtk_file_chooser_native_new(const gchar* _0, GtkWindow* _1, GtkFileChooserAction _2, const gchar* _3, const gchar* _4) __attribute__((weak)) {
	goPanic("gtk_file_chooser_native_new: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
GtkPopoverConstraint gtk_popover_get_constrain_to(GtkPopover* v) __attribute__((weak)) {
	goPanic("gtk_popover_get_constrain_to: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
GtkShortcutsWindow* gtk_application_window_get_help_overlay(GtkApplicationWindow* v) __attribute__((weak)) {
	goPanic("gtk_application_window_get_help_overlay: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
GtkWindow* gtk_native_dialog_get_transient_for(GtkNativeDialog* v) __attribute__((weak)) {
	goPanic("gtk_native_dialog_get_transient_for: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
char* gtk_style_context_to_string(GtkStyleContext* v, GtkStyleContextPrintFlags _0) __attribute__((weak)) {
	goPanic("gtk_style_context_to_string: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
const char* gtk_file_chooser_native_get_accept_label(GtkFileChooserNative* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_native_get_accept_label: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
const char* gtk_file_chooser_native_get_cancel_label(GtkFileChooserNative* v) __attribute__((weak)) {
	goPanic("gtk_file_chooser_native_get_cancel_label: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
const char* gtk_native_dialog_get_title(GtkNativeDialog* v) __attribute__((weak)) {
	goPanic("gtk_native_dialog_get_title: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
const char* gtk_widget_path_iter_get_object_name(const GtkWidgetPath* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_get_object_name: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
gboolean gtk_native_dialog_get_modal(GtkNativeDialog* v) __attribute__((weak)) {
	goPanic("gtk_native_dialog_get_modal: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
gboolean gtk_native_dialog_get_visible(GtkNativeDialog* v) __attribute__((weak)) {
	goPanic("gtk_native_dialog_get_visible: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
gboolean gtk_text_iter_starts_tag(const GtkTextIter* v, GtkTextTag* _0) __attribute__((weak)) {
	goPanic("gtk_text_iter_starts_tag: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
gboolean gtk_widget_get_focus_on_click(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_focus_on_click: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
gint gtk_native_dialog_run(GtkNativeDialog* v) __attribute__((weak)) {
	goPanic("gtk_native_dialog_run: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_application_window_set_help_overlay(GtkApplicationWindow* v, GtkShortcutsWindow* _0) __attribute__((weak)) {
	goPanic("gtk_application_window_set_help_overlay: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_file_chooser_native_set_accept_label(GtkFileChooserNative* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_native_set_accept_label: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_file_chooser_native_set_cancel_label(GtkFileChooserNative* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_native_set_cancel_label: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_native_dialog_destroy(GtkNativeDialog* v) __attribute__((weak)) {
	goPanic("gtk_native_dialog_destroy: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_native_dialog_hide(GtkNativeDialog* v) __attribute__((weak)) {
	goPanic("gtk_native_dialog_hide: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_native_dialog_set_modal(GtkNativeDialog* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_native_dialog_set_modal: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_native_dialog_set_title(GtkNativeDialog* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_native_dialog_set_title: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_native_dialog_set_transient_for(GtkNativeDialog* v, GtkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_native_dialog_set_transient_for: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_native_dialog_show(GtkNativeDialog* v) __attribute__((weak)) {
	goPanic("gtk_native_dialog_show: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_popover_set_constrain_to(GtkPopover* v, GtkPopoverConstraint _0) __attribute__((weak)) {
	goPanic("gtk_popover_set_constrain_to: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_render_background_get_clip(GtkStyleContext* _0, gdouble _1, gdouble _2, gdouble _3, gdouble _4, GdkRectangle* _5) __attribute__((weak)) {
	goPanic("gtk_render_background_get_clip: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_settings_reset_property(GtkSettings* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_settings_reset_property: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_text_tag_changed(GtkTextTag* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_text_tag_changed: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_text_view_reset_cursor_blink(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_reset_cursor_blink: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_widget_get_allocated_size(GtkWidget* v, GtkAllocation* _0, int* _1) __attribute__((weak)) {
	goPanic("gtk_widget_get_allocated_size: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_widget_path_iter_set_object_name(GtkWidgetPath* v, gint _0, const char* _1) __attribute__((weak)) {
	goPanic("gtk_widget_path_iter_set_object_name: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_widget_queue_allocate(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_queue_allocate: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 20))
void gtk_widget_set_focus_on_click(GtkWidget* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_focus_on_click: library too old: needs at least 3.20");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GVariant* gtk_file_filter_to_gvariant(GtkFileFilter* v) __attribute__((weak)) {
	goPanic("gtk_file_filter_to_gvariant: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GVariant* gtk_page_setup_to_gvariant(GtkPageSetup* v) __attribute__((weak)) {
	goPanic("gtk_page_setup_to_gvariant: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GVariant* gtk_paper_size_to_gvariant(GtkPaperSize* v) __attribute__((weak)) {
	goPanic("gtk_paper_size_to_gvariant: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GVariant* gtk_print_settings_to_gvariant(GtkPrintSettings* v) __attribute__((weak)) {
	goPanic("gtk_print_settings_to_gvariant: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GtkFileFilter* gtk_file_filter_new_from_gvariant(GVariant* _0) __attribute__((weak)) {
	goPanic("gtk_file_filter_new_from_gvariant: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GtkFlowBoxChild* gtk_flow_box_get_child_at_pos(GtkFlowBox* v, gint _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_flow_box_get_child_at_pos: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GtkPadController* gtk_pad_controller_new(GtkWindow* _0, GActionGroup* _1, GdkDevice* _2) __attribute__((weak)) {
	goPanic("gtk_pad_controller_new: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GtkPageSetup* gtk_page_setup_new_from_gvariant(GVariant* _0) __attribute__((weak)) {
	goPanic("gtk_page_setup_new_from_gvariant: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GtkPaperSize* gtk_paper_size_new_from_gvariant(GVariant* _0) __attribute__((weak)) {
	goPanic("gtk_paper_size_new_from_gvariant: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GtkPrintSettings* gtk_print_settings_new_from_gvariant(GVariant* _0) __attribute__((weak)) {
	goPanic("gtk_print_settings_new_from_gvariant: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
GtkWidget* gtk_shortcut_label_new(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_shortcut_label_new: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
const char* gtk_file_chooser_get_choice(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_get_choice: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
const gchar* gtk_shortcut_label_get_accelerator(GtkShortcutLabel* v) __attribute__((weak)) {
	goPanic("gtk_shortcut_label_get_accelerator: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
const gchar* gtk_shortcut_label_get_disabled_text(GtkShortcutLabel* v) __attribute__((weak)) {
	goPanic("gtk_shortcut_label_get_disabled_text: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
gboolean gtk_gl_area_get_use_es(GtkGLArea* v) __attribute__((weak)) {
	goPanic("gtk_gl_area_get_use_es: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
gboolean gtk_info_bar_get_revealed(GtkInfoBar* v) __attribute__((weak)) {
	goPanic("gtk_info_bar_get_revealed: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
gboolean gtk_places_sidebar_get_show_starred_location(GtkPlacesSidebar* v) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_get_show_starred_location: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
gboolean gtk_scrolled_window_get_propagate_natural_height(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_propagate_natural_height: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
gboolean gtk_scrolled_window_get_propagate_natural_width(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_propagate_natural_width: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
gboolean gtk_show_uri_on_window(GtkWindow* _0, const char* _1, guint32 _2) __attribute__((weak)) {
	goPanic("gtk_show_uri_on_window: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
gint gtk_scrolled_window_get_max_content_height(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_max_content_height: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
gint gtk_scrolled_window_get_max_content_width(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_max_content_width: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_file_chooser_add_choice(GtkFileChooser* v, const char* _0, const char* _1, const char** _2, const char** _3) __attribute__((weak)) {
	goPanic("gtk_file_chooser_add_choice: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_file_chooser_remove_choice(GtkFileChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_file_chooser_remove_choice: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_file_chooser_set_choice(GtkFileChooser* v, const char* _0, const char* _1) __attribute__((weak)) {
	goPanic("gtk_file_chooser_set_choice: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_gl_area_set_use_es(GtkGLArea* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_gl_area_set_use_es: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_info_bar_set_revealed(GtkInfoBar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_info_bar_set_revealed: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_menu_place_on_monitor(GtkMenu* v, GdkMonitor* _0) __attribute__((weak)) {
	goPanic("gtk_menu_place_on_monitor: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_pad_controller_set_action(GtkPadController* v, GtkPadActionType _0, gint _1, gint _2, const gchar* _3, const gchar* _4) __attribute__((weak)) {
	goPanic("gtk_pad_controller_set_action: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_pad_controller_set_action_entries(GtkPadController* v, const GtkPadActionEntry* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_pad_controller_set_action_entries: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_places_sidebar_set_show_starred_location(GtkPlacesSidebar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_places_sidebar_set_show_starred_location: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_popover_popdown(GtkPopover* v) __attribute__((weak)) {
	goPanic("gtk_popover_popdown: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_popover_popup(GtkPopover* v) __attribute__((weak)) {
	goPanic("gtk_popover_popup: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_scrolled_window_set_max_content_height(GtkScrolledWindow* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_set_max_content_height: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_scrolled_window_set_max_content_width(GtkScrolledWindow* v, gint _0) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_set_max_content_width: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_scrolled_window_set_propagate_natural_height(GtkScrolledWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_set_propagate_natural_height: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_scrolled_window_set_propagate_natural_width(GtkScrolledWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_set_propagate_natural_width: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_shortcut_label_set_accelerator(GtkShortcutLabel* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_shortcut_label_set_accelerator: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 22))
void gtk_shortcut_label_set_disabled_text(GtkShortcutLabel* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_shortcut_label_set_disabled_text: library too old: needs at least 3.22");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
GdkDeviceTool* gtk_gesture_stylus_get_device_tool(GtkGestureStylus* v) __attribute__((weak)) {
	goPanic("gtk_gesture_stylus_get_device_tool: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
GtkEventController* gtk_event_controller_motion_new(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_event_controller_motion_new: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
GtkEventController* gtk_event_controller_scroll_new(GtkWidget* _0, GtkEventControllerScrollFlags _1) __attribute__((weak)) {
	goPanic("gtk_event_controller_scroll_new: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
GtkEventControllerScrollFlags gtk_event_controller_scroll_get_flags(GtkEventControllerScroll* v) __attribute__((weak)) {
	goPanic("gtk_event_controller_scroll_get_flags: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
GtkFontChooserLevel gtk_font_chooser_get_level(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_level: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
GtkGesture* gtk_gesture_stylus_new(GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_gesture_stylus_new: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
GtkIMContext* gtk_event_controller_key_get_im_context(GtkEventControllerKey* v) __attribute__((weak)) {
	goPanic("gtk_event_controller_key_get_im_context: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
char* gtk_font_chooser_get_font_features(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_font_features: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
char* gtk_font_chooser_get_language(GtkFontChooser* v) __attribute__((weak)) {
	goPanic("gtk_font_chooser_get_language: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
gboolean gtk_gesture_stylus_get_axis(GtkGestureStylus* v, GdkAxisUse _0, gdouble* _1) __attribute__((weak)) {
	goPanic("gtk_gesture_stylus_get_axis: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
void gtk_event_controller_scroll_set_flags(GtkEventControllerScroll* v, GtkEventControllerScrollFlags _0) __attribute__((weak)) {
	goPanic("gtk_event_controller_scroll_set_flags: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
void gtk_font_chooser_set_language(GtkFontChooser* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_font_chooser_set_language: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 24))
void gtk_font_chooser_set_level(GtkFontChooser* v, GtkFontChooserLevel _0) __attribute__((weak)) {
	goPanic("gtk_font_chooser_set_level: library too old: needs at least 3.24");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GMenuModel* gtk_application_get_app_menu(GtkApplication* v) __attribute__((weak)) {
	goPanic("gtk_application_get_app_menu: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GMenuModel* gtk_application_get_menubar(GtkApplication* v) __attribute__((weak)) {
	goPanic("gtk_application_get_menubar: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GVariant* gtk_actionable_get_action_target_value(GtkActionable* v) __attribute__((weak)) {
	goPanic("gtk_actionable_get_action_target_value: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GdkModifierType gtk_widget_get_modifier_mask(GtkWidget* v, GdkModifierIntent _0) __attribute__((weak)) {
	goPanic("gtk_widget_get_modifier_mask: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GtkStyleContext* gtk_style_context_get_parent(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_get_parent: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GtkSymbolicColor* gtk_symbolic_color_new_win32(const gchar* _0, gint _1) __attribute__((weak)) {
	goPanic("gtk_symbolic_color_new_win32: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_application_window_new(GtkApplication* _0) __attribute__((weak)) {
	goPanic("gtk_application_window_new: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_color_chooser_dialog_new(const gchar* _0, GtkWindow* _1) __attribute__((weak)) {
	goPanic("gtk_color_chooser_dialog_new: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_color_chooser_widget_new(void) __attribute__((weak)) {
	goPanic("gtk_color_chooser_widget_new: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_image_new_from_resource(const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_image_new_from_resource: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_menu_bar_new_from_model(GMenuModel* _0) __attribute__((weak)) {
	goPanic("gtk_menu_bar_new_from_model: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_menu_new_from_model(GMenuModel* _0) __attribute__((weak)) {
	goPanic("gtk_menu_new_from_model: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
GtkWidget* gtk_window_get_attached_to(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_attached_to: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
const gchar* gtk_actionable_get_action_name(GtkActionable* v) __attribute__((weak)) {
	goPanic("gtk_actionable_get_action_name: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
gboolean gtk_application_is_inhibited(GtkApplication* v, GtkApplicationInhibitFlags _0) __attribute__((weak)) {
	goPanic("gtk_application_is_inhibited: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
gboolean gtk_application_window_get_show_menubar(GtkApplicationWindow* v) __attribute__((weak)) {
	goPanic("gtk_application_window_get_show_menubar: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
gboolean gtk_color_chooser_get_use_alpha(GtkColorChooser* v) __attribute__((weak)) {
	goPanic("gtk_color_chooser_get_use_alpha: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
gboolean gtk_scale_get_has_origin(GtkScale* v) __attribute__((weak)) {
	goPanic("gtk_scale_get_has_origin: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
gboolean gtk_scrolled_window_get_capture_button_press(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_capture_button_press: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
gboolean gtk_scrolled_window_get_kinetic_scrolling(GtkScrolledWindow* v) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_get_kinetic_scrolling: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
gboolean gtk_window_get_hide_titlebar_when_maximized(GtkWindow* v) __attribute__((weak)) {
	goPanic("gtk_window_get_hide_titlebar_when_maximized: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
gchar* gtk_accelerator_get_label_with_keycode(GdkDisplay* _0, guint _1, guint _2, GdkModifierType _3) __attribute__((weak)) {
	goPanic("gtk_accelerator_get_label_with_keycode: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
gchar* gtk_accelerator_name_with_keycode(GdkDisplay* _0, guint _1, guint _2, GdkModifierType _3) __attribute__((weak)) {
	goPanic("gtk_accelerator_name_with_keycode: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
gchar* gtk_entry_completion_compute_prefix(GtkEntryCompletion* v, const char* _0) __attribute__((weak)) {
	goPanic("gtk_entry_completion_compute_prefix: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
guint gtk_application_inhibit(GtkApplication* v, GtkWindow* _0, GtkApplicationInhibitFlags _1, const gchar* _2) __attribute__((weak)) {
	goPanic("gtk_application_inhibit: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
guint gtk_builder_add_from_resource(GtkBuilder* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_builder_add_from_resource: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
guint gtk_builder_add_objects_from_resource(GtkBuilder* v, const gchar* _0, gchar** _1) __attribute__((weak)) {
	goPanic("gtk_builder_add_objects_from_resource: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
guint gtk_tree_view_get_n_columns(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_n_columns: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
guint gtk_ui_manager_add_ui_from_resource(GtkUIManager* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_ui_manager_add_ui_from_resource: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_about_dialog_add_credit_section(GtkAboutDialog* v, const gchar* _0, const gchar** _1) __attribute__((weak)) {
	goPanic("gtk_about_dialog_add_credit_section: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_accelerator_parse_with_keycode(const gchar* _0, guint* _1, guint** _2, GdkModifierType* _3) __attribute__((weak)) {
	goPanic("gtk_accelerator_parse_with_keycode: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_actionable_set_action_name(GtkActionable* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_actionable_set_action_name: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_actionable_set_action_target_value(GtkActionable* v, GVariant* _0) __attribute__((weak)) {
	goPanic("gtk_actionable_set_action_target_value: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_actionable_set_detailed_action_name(GtkActionable* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_actionable_set_detailed_action_name: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_application_add_accelerator(GtkApplication* v, const gchar* _0, const gchar* _1, GVariant* _2) __attribute__((weak)) {
	goPanic("gtk_application_add_accelerator: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_application_remove_accelerator(GtkApplication* v, const gchar* _0, GVariant* _1) __attribute__((weak)) {
	goPanic("gtk_application_remove_accelerator: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_application_set_app_menu(GtkApplication* v, GMenuModel* _0) __attribute__((weak)) {
	goPanic("gtk_application_set_app_menu: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_application_set_menubar(GtkApplication* v, GMenuModel* _0) __attribute__((weak)) {
	goPanic("gtk_application_set_menubar: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_application_uninhibit(GtkApplication* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_application_uninhibit: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_application_window_set_show_menubar(GtkApplicationWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_application_window_set_show_menubar: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_color_chooser_add_palette(GtkColorChooser* v, GtkOrientation _0, gint _1, gint _2, GdkRGBA* _3) __attribute__((weak)) {
	goPanic("gtk_color_chooser_add_palette: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_color_chooser_get_rgba(GtkColorChooser* v, GdkRGBA* _0) __attribute__((weak)) {
	goPanic("gtk_color_chooser_get_rgba: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_color_chooser_set_rgba(GtkColorChooser* v, const GdkRGBA* _0) __attribute__((weak)) {
	goPanic("gtk_color_chooser_set_rgba: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_color_chooser_set_use_alpha(GtkColorChooser* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_color_chooser_set_use_alpha: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_render_insertion_cursor(GtkStyleContext* _0, cairo_t* _1, gdouble _2, gdouble _3, PangoLayout* _4, int _5, PangoDirection _6) __attribute__((weak)) {
	goPanic("gtk_render_insertion_cursor: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_scale_set_has_origin(GtkScale* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_scale_set_has_origin: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_scrolled_window_set_capture_button_press(GtkScrolledWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_set_capture_button_press: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_scrolled_window_set_kinetic_scrolling(GtkScrolledWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_scrolled_window_set_kinetic_scrolling: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_style_context_set_parent(GtkStyleContext* v, GtkStyleContext* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_set_parent: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_window_set_attached_to(GtkWindow* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_window_set_attached_to: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 4))
void gtk_window_set_hide_titlebar_when_maximized(GtkWindow* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_window_set_hide_titlebar_when_maximized: library too old: needs at least 3.4");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GMenuModel* gtk_menu_button_get_menu_model(GtkMenuButton* v) __attribute__((weak)) {
	goPanic("gtk_menu_button_get_menu_model: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkAccelGroup* gtk_action_group_get_accel_group(GtkActionGroup* v) __attribute__((weak)) {
	goPanic("gtk_action_group_get_accel_group: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkArrowType gtk_menu_button_get_direction(GtkMenuButton* v) __attribute__((weak)) {
	goPanic("gtk_menu_button_get_direction: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkInputHints gtk_entry_get_input_hints(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_input_hints: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkInputHints gtk_text_view_get_input_hints(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_get_input_hints: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkInputPurpose gtk_entry_get_input_purpose(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_input_purpose: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkInputPurpose gtk_text_view_get_input_purpose(GtkTextView* v) __attribute__((weak)) {
	goPanic("gtk_text_view_get_input_purpose: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkLevelBarMode gtk_level_bar_get_mode(GtkLevelBar* v) __attribute__((weak)) {
	goPanic("gtk_level_bar_get_mode: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkMenu* gtk_menu_button_get_popup(GtkMenuButton* v) __attribute__((weak)) {
	goPanic("gtk_menu_button_get_popup: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_level_bar_new(void) __attribute__((weak)) {
	goPanic("gtk_level_bar_new: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_level_bar_new_for_interval(gdouble _0, gdouble _1) __attribute__((weak)) {
	goPanic("gtk_level_bar_new_for_interval: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_menu_button_get_align_widget(GtkMenuButton* v) __attribute__((weak)) {
	goPanic("gtk_menu_button_get_align_widget: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_menu_button_new(void) __attribute__((weak)) {
	goPanic("gtk_menu_button_new: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkWidget* gtk_search_entry_new(void) __attribute__((weak)) {
	goPanic("gtk_search_entry_new: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkWindow* gtk_application_get_active_window(GtkApplication* v) __attribute__((weak)) {
	goPanic("gtk_application_get_active_window: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
GtkWindow* gtk_application_get_window_by_id(GtkApplication* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_application_get_window_by_id: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
PangoAttrList* gtk_entry_get_attributes(GtkEntry* v) __attribute__((weak)) {
	goPanic("gtk_entry_get_attributes: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
gboolean gtk_button_get_always_show_image(GtkButton* v) __attribute__((weak)) {
	goPanic("gtk_button_get_always_show_image: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
gboolean gtk_icon_view_get_cell_rect(GtkIconView* v, GtkTreePath* _0, GtkCellRenderer* _1, GdkRectangle* _2) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_cell_rect: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
gboolean gtk_level_bar_get_offset_value(GtkLevelBar* v, const gchar* _0, gdouble* _1) __attribute__((weak)) {
	goPanic("gtk_level_bar_get_offset_value: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
gdouble gtk_level_bar_get_max_value(GtkLevelBar* v) __attribute__((weak)) {
	goPanic("gtk_level_bar_get_max_value: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
gdouble gtk_level_bar_get_min_value(GtkLevelBar* v) __attribute__((weak)) {
	goPanic("gtk_level_bar_get_min_value: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
gdouble gtk_level_bar_get_value(GtkLevelBar* v) __attribute__((weak)) {
	goPanic("gtk_level_bar_get_value: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
guint gtk_application_window_get_id(GtkApplicationWindow* v) __attribute__((weak)) {
	goPanic("gtk_application_window_get_id: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_accel_label_set_accel(GtkAccelLabel* v, guint _0, GdkModifierType _1) __attribute__((weak)) {
	goPanic("gtk_accel_label_set_accel: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_action_group_set_accel_group(GtkActionGroup* v, GtkAccelGroup* _0) __attribute__((weak)) {
	goPanic("gtk_action_group_set_accel_group: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_button_set_always_show_image(GtkButton* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_button_set_always_show_image: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_entry_set_attributes(GtkEntry* v, PangoAttrList* _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_attributes: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_entry_set_input_hints(GtkEntry* v, GtkInputHints _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_input_hints: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_entry_set_input_purpose(GtkEntry* v, GtkInputPurpose _0) __attribute__((weak)) {
	goPanic("gtk_entry_set_input_purpose: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_level_bar_add_offset_value(GtkLevelBar* v, const gchar* _0, gdouble _1) __attribute__((weak)) {
	goPanic("gtk_level_bar_add_offset_value: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_level_bar_remove_offset_value(GtkLevelBar* v, const gchar* _0) __attribute__((weak)) {
	goPanic("gtk_level_bar_remove_offset_value: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_level_bar_set_max_value(GtkLevelBar* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_level_bar_set_max_value: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_level_bar_set_min_value(GtkLevelBar* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_level_bar_set_min_value: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_level_bar_set_mode(GtkLevelBar* v, GtkLevelBarMode _0) __attribute__((weak)) {
	goPanic("gtk_level_bar_set_mode: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_level_bar_set_value(GtkLevelBar* v, gdouble _0) __attribute__((weak)) {
	goPanic("gtk_level_bar_set_value: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_menu_button_set_align_widget(GtkMenuButton* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_menu_button_set_align_widget: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_menu_button_set_direction(GtkMenuButton* v, GtkArrowType _0) __attribute__((weak)) {
	goPanic("gtk_menu_button_set_direction: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_menu_button_set_menu_model(GtkMenuButton* v, GMenuModel* _0) __attribute__((weak)) {
	goPanic("gtk_menu_button_set_menu_model: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_menu_button_set_popup(GtkMenuButton* v, GtkWidget* _0) __attribute__((weak)) {
	goPanic("gtk_menu_button_set_popup: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_menu_shell_bind_model(GtkMenuShell* v, GMenuModel* _0, const gchar* _1, gboolean _2) __attribute__((weak)) {
	goPanic("gtk_menu_shell_bind_model: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_text_view_set_input_hints(GtkTextView* v, GtkInputHints _0) __attribute__((weak)) {
	goPanic("gtk_text_view_set_input_hints: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_text_view_set_input_purpose(GtkTextView* v, GtkInputPurpose _0) __attribute__((weak)) {
	goPanic("gtk_text_view_set_input_purpose: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 6))
void gtk_widget_insert_action_group(GtkWidget* v, const gchar* _0, GActionGroup* _1) __attribute__((weak)) {
	goPanic("gtk_widget_insert_action_group: library too old: needs at least 3.6");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
GdkFrameClock* gtk_style_context_get_frame_clock(GtkStyleContext* v) __attribute__((weak)) {
	goPanic("gtk_style_context_get_frame_clock: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
GdkFrameClock* gtk_widget_get_frame_clock(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_frame_clock: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
GdkPixbuf* gtk_icon_info_load_icon_finish(GtkIconInfo* v, GAsyncResult* _0) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_icon_finish: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
GdkPixbuf* gtk_icon_info_load_symbolic_finish(GtkIconInfo* v, GAsyncResult* _0, gboolean* _1) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_symbolic_finish: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
GdkPixbuf* gtk_icon_info_load_symbolic_for_context_finish(GtkIconInfo* v, GAsyncResult* _0, gboolean* _1) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_symbolic_for_context_finish: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
double gtk_widget_get_opacity(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_get_opacity: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
gboolean gtk_icon_view_get_activate_on_single_click(GtkIconView* v) __attribute__((weak)) {
	goPanic("gtk_icon_view_get_activate_on_single_click: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
gboolean gtk_level_bar_get_inverted(GtkLevelBar* v) __attribute__((weak)) {
	goPanic("gtk_level_bar_get_inverted: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
gboolean gtk_tree_view_get_activate_on_single_click(GtkTreeView* v) __attribute__((weak)) {
	goPanic("gtk_tree_view_get_activate_on_single_click: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
gboolean gtk_widget_is_visible(GtkWidget* v) __attribute__((weak)) {
	goPanic("gtk_widget_is_visible: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
guint gtk_widget_add_tick_callback(GtkWidget* v, GtkTickCallback _0, gpointer _1, GDestroyNotify _2) __attribute__((weak)) {
	goPanic("gtk_widget_add_tick_callback: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_builder_expose_object(GtkBuilder* v, const gchar* _0, GObject* _1) __attribute__((weak)) {
	goPanic("gtk_builder_expose_object: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_icon_info_load_icon_async(GtkIconInfo* v, GCancellable* _0, GAsyncReadyCallback _1, gpointer _2) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_icon_async: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_icon_info_load_symbolic_async(GtkIconInfo* v, const GdkRGBA* _0, const GdkRGBA* _1, const GdkRGBA* _2, const GdkRGBA* _3, GCancellable* _4, GAsyncReadyCallback _5, gpointer _6) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_symbolic_async: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_icon_info_load_symbolic_for_context_async(GtkIconInfo* v, GtkStyleContext* _0, GCancellable* _1, GAsyncReadyCallback _2, gpointer _3) __attribute__((weak)) {
	goPanic("gtk_icon_info_load_symbolic_for_context_async: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_icon_view_set_activate_on_single_click(GtkIconView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_icon_view_set_activate_on_single_click: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_level_bar_set_inverted(GtkLevelBar* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_level_bar_set_inverted: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_style_context_set_frame_clock(GtkStyleContext* v, GdkFrameClock* _0) __attribute__((weak)) {
	goPanic("gtk_style_context_set_frame_clock: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_tree_view_set_activate_on_single_click(GtkTreeView* v, gboolean _0) __attribute__((weak)) {
	goPanic("gtk_tree_view_set_activate_on_single_click: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_widget_register_window(GtkWidget* v, GdkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_widget_register_window: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_widget_remove_tick_callback(GtkWidget* v, guint _0) __attribute__((weak)) {
	goPanic("gtk_widget_remove_tick_callback: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_widget_set_opacity(GtkWidget* v, double _0) __attribute__((weak)) {
	goPanic("gtk_widget_set_opacity: library too old: needs at least 3.8");
}
#endif

#if (GTK_MAJOR_VERSION < 3 || (GTK_MAJOR_VERSION == 3 && GTK_MINOR_VERSION < 8))
void gtk_widget_unregister_window(GtkWidget* v, GdkWindow* _0) __attribute__((weak)) {
	goPanic("gtk_widget_unregister_window: library too old: needs at least 3.8");
}
#endif
