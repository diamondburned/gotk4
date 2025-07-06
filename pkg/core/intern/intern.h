#include <glib-object.h>

extern void goToggleNotify(gpointer, GObject *, gboolean);
extern void goFinishRemovingToggleRef(gpointer);
gboolean gotk4_intern_remove_toggle_ref(gpointer obj);
