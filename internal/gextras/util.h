#include <glib-2.0/glib.h>

// Helper function because cgo thinks gpointer is a different type, even though
// it's just a typedef to void*.
static gpointer conptr(void *ptr) { return ptr; };

static const gchar *object_get_class_name(GObject *object) {
  return G_OBJECT_CLASS_NAME(G_OBJECT_GET_CLASS(object));
};

static GObject *toGObject(void *p) { return G_OBJECT(p); };
