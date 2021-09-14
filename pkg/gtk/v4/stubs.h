#include <glib-object.h>
#include <gtk/gtk.h>

extern void goPanic(const char*);

#if (GTK_MAJOR_VERSION < 4 || (GTK_MAJOR_VERSION == 4 && GTK_MINOR_VERSION < 2))
gboolean gtk_im_context_get_surrounding_with_selection(GtkIMContext* v, char** _0, int* _1, int* _2) {
	goPanic("gtk_im_context_get_surrounding_with_selection: library too old: needs at least 4.2");
}
#endif

#if (GTK_MAJOR_VERSION < 4 || (GTK_MAJOR_VERSION == 4 && GTK_MINOR_VERSION < 2))
gboolean gtk_window_get_handle_menubar_accel(GtkWindow* v) {
	goPanic("gtk_window_get_handle_menubar_accel: library too old: needs at least 4.2");
}
#endif

#if (GTK_MAJOR_VERSION < 4 || (GTK_MAJOR_VERSION == 4 && GTK_MINOR_VERSION < 2))
void gtk_im_context_set_surrounding_with_selection(GtkIMContext* v, const char* _0, int _1, int _2, int _3) {
	goPanic("gtk_im_context_set_surrounding_with_selection: library too old: needs at least 4.2");
}
#endif

#if (GTK_MAJOR_VERSION < 4 || (GTK_MAJOR_VERSION == 4 && GTK_MINOR_VERSION < 2))
void gtk_window_set_handle_menubar_accel(GtkWindow* v, gboolean _0) {
	goPanic("gtk_window_set_handle_menubar_accel: library too old: needs at least 4.2");
}
#endif
