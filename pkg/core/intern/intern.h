#include <glib-object.h>

extern void goToggleNotify(guintptr, GObject *, gboolean);
extern void goFinishRemovingToggleRef(guintptr);
extern void goRemoveQueuedToggleRef(guintptr);
const gchar *gotk4_object_type_name(gpointer obj);
void gotk4_intern_add_toggle_ref(gpointer obj, guintptr token);
void gotk4_intern_remove_toggle_ref(gpointer obj, guintptr token);
void gotk4_intern_queue_remove_toggle_ref(guintptr token);
void gotk4_intern_disconnect_closure(gpointer obj, GClosure *closure);
