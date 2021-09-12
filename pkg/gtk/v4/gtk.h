#include <glib-object.h>
#include <gtk/gtk.h>
GListModel* _gotk4_gtk4_TreeListModelCreateModelFunc(gpointer, gpointer);
GtkWidget* _gotk4_gtk4_FlowBoxCreateWidgetFunc(gpointer, gpointer);
GtkWidget* _gotk4_gtk4_ListBoxCreateWidgetFunc(gpointer, gpointer);
char* _gotk4_gtk4_ScaleFormatValueFunc(GtkScale*, double, gpointer);
extern void callbackDelete(gpointer);
extern void goPanic(char*);
gboolean _gotk4_gtk4_CellAllocCallback(GtkCellRenderer*, GdkRectangle*, GdkRectangle*, gpointer);
gboolean _gotk4_gtk4_CellCallback(GtkCellRenderer*, gpointer);
gboolean _gotk4_gtk4_CustomFilterFunc(gpointer, gpointer);
gboolean _gotk4_gtk4_EntryCompletionMatchFunc(GtkEntryCompletion*, char*, GtkTreeIter*, gpointer);
gboolean _gotk4_gtk4_FlowBoxFilterFunc(GtkFlowBoxChild*, gpointer);
gboolean _gotk4_gtk4_FontFilterFunc(PangoFontFamily*, PangoFontFace*, gpointer);
gboolean _gotk4_gtk4_ListBoxFilterFunc(GtkListBoxRow*, gpointer);
gboolean _gotk4_gtk4_ShortcutFunc(GtkWidget*, GVariant*, gpointer);
gboolean _gotk4_gtk4_TextCharPredicate(gunichar, gpointer);
gboolean _gotk4_gtk4_TickCallback(GtkWidget*, GdkFrameClock*, gpointer);
gboolean _gotk4_gtk4_TreeModelFilterVisibleFunc(GtkTreeModel*, GtkTreeIter*, gpointer);
gboolean _gotk4_gtk4_TreeModelForeachFunc(GtkTreeModel*, GtkTreePath*, GtkTreeIter*, gpointer);
gboolean _gotk4_gtk4_TreeSelectionFunc(GtkTreeSelection*, GtkTreeModel*, GtkTreePath*, gboolean, gpointer);
gboolean _gotk4_gtk4_TreeViewColumnDropFunc(GtkTreeView*, GtkTreeViewColumn*, GtkTreeViewColumn*, GtkTreeViewColumn*, gpointer);
gboolean _gotk4_gtk4_TreeViewRowSeparatorFunc(GtkTreeModel*, GtkTreeIter*, gpointer);
gboolean _gotk4_gtk4_TreeViewSearchEqualFunc(GtkTreeModel*, int, char*, GtkTreeIter*, gpointer);
gint _gotk4_glib2_CompareDataFunc(gconstpointer, gconstpointer, gpointer);
gpointer _gotk4_gtk4_MapListModelMapFunc(gpointer, gpointer);
int _gotk4_gtk4_AssistantPageFunc(int, gpointer);
int _gotk4_gtk4_FlowBoxSortFunc(GtkFlowBoxChild*, GtkFlowBoxChild*, gpointer);
int _gotk4_gtk4_ListBoxSortFunc(GtkListBoxRow*, GtkListBoxRow*, gpointer);
int _gotk4_gtk4_TreeIterCompareFunc(GtkTreeModel*, GtkTreeIter*, GtkTreeIter*, gpointer);
void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
void _gotk4_gtk4_CellLayoutDataFunc(GtkCellLayout*, GtkCellRenderer*, GtkTreeModel*, GtkTreeIter*, gpointer);
void _gotk4_gtk4_DrawingAreaDrawFunc(GtkDrawingArea*, cairo_t*, int, int, gpointer);
void _gotk4_gtk4_ExpressionNotify(gpointer);
void _gotk4_gtk4_FlowBoxForeachFunc(GtkFlowBox*, GtkFlowBoxChild*, gpointer);
void _gotk4_gtk4_IconViewForeachFunc(GtkIconView*, GtkTreePath*, gpointer);
void _gotk4_gtk4_ListBoxForeachFunc(GtkListBox*, GtkListBoxRow*, gpointer);
void _gotk4_gtk4_ListBoxUpdateHeaderFunc(GtkListBoxRow*, GtkListBoxRow*, gpointer);
void _gotk4_gtk4_MenuButtonCreatePopupFunc(GtkMenuButton*, gpointer);
void _gotk4_gtk4_PageSetupDoneFunc(GtkPageSetup*, gpointer);
void _gotk4_gtk4_PrintSettingsFunc(char*, char*, gpointer);
void _gotk4_gtk4_TextTagTableForeach(GtkTextTag*, gpointer);
void _gotk4_gtk4_TreeCellDataFunc(GtkTreeViewColumn*, GtkCellRenderer*, GtkTreeModel*, GtkTreeIter*, gpointer);
void _gotk4_gtk4_TreeModelFilterModifyFunc(GtkTreeModel*, GtkTreeIter*, GValue*, int, gpointer);
void _gotk4_gtk4_TreeSelectionForeachFunc(GtkTreeModel*, GtkTreePath*, GtkTreeIter*, gpointer);
void _gotk4_gtk4_TreeViewMappingFunc(GtkTreeView*, GtkTreePath*, gpointer);
#if !(GTK_CHECK_VERSION(4, 2, 0))
gboolean gtk_im_context_get_surrounding_with_selection(GtkIMContext* v, char** _0, int* _1, int* _2) {
	goPanic("gtk_im_context_get_surrounding_with_selection: library too old: needs at least 4.2");
}
#endif
#if !(GTK_CHECK_VERSION(4, 2, 0))
gboolean gtk_window_get_handle_menubar_accel(GtkWindow* v) {
	goPanic("gtk_window_get_handle_menubar_accel: library too old: needs at least 4.2");
}
#endif
#if !(GTK_CHECK_VERSION(4, 2, 0))
void gtk_im_context_set_surrounding_with_selection(GtkIMContext* v, const char* _0, int _1, int _2, int _3) {
	goPanic("gtk_im_context_set_surrounding_with_selection: library too old: needs at least 4.2");
}
#endif
#if !(GTK_CHECK_VERSION(4, 2, 0))
void gtk_window_set_handle_menubar_accel(GtkWindow* v, gboolean _0) {
	goPanic("gtk_window_set_handle_menubar_accel: library too old: needs at least 4.2");
}
#endif