#include "intern.h"

const gchar *gotk4_object_type_name(gpointer obj) {
  return G_OBJECT_TYPE_NAME(obj);
};

static void gotk4_intern_toggle_notify(gpointer data, GObject *obj,
                                       gboolean is_last_ref) {
  goToggleNotify((guintptr)data, obj, is_last_ref);
}

void gotk4_intern_add_toggle_ref(gpointer obj, guintptr token) {
  g_object_add_toggle_ref(G_OBJECT(obj), gotk4_intern_toggle_notify,
                          (gpointer)token);
}

void gotk4_intern_remove_toggle_ref(gpointer obj, guintptr token) {
  // Remove the toggle reference first. This may run GObject destruction
  // callbacks, so the Go intern entry is kept until this call returns.
  g_object_remove_toggle_ref(G_OBJECT(obj), gotk4_intern_toggle_notify,
                             (gpointer)token);

  goFinishRemovingToggleRef(token);
}

static gboolean remove_toggle_ref_idle(gpointer data) {
  goRemoveQueuedToggleRef((guintptr)data);
  return G_SOURCE_REMOVE;
}

void gotk4_intern_queue_remove_toggle_ref(guintptr token) {
  GSource *source = g_idle_source_new();
  g_source_set_priority(source, G_PRIORITY_DEFAULT);
  g_source_set_callback(source, remove_toggle_ref_idle, (gpointer)token, NULL);
  g_source_attach(source, g_main_context_default());
  g_source_unref(source);
}

void gotk4_intern_disconnect_closure(gpointer obj, GClosure *closure) {
  g_signal_handlers_disconnect_matched(obj, G_SIGNAL_MATCH_CLOSURE, 0, 0,
                                        closure, NULL, NULL);
}
