#include "listmodel.h"

// defined in gbox.
extern void callbackDelete(guintptr id);

#define GDK_ARRAY_ELEMENT_TYPE Gotk4GboxObject *
#define GDK_ARRAY_NAME objects
#define GDK_ARRAY_TYPE_NAME Objects
#define GDK_ARRAY_FREE_FUNC g_object_unref
#include "gdkarrayimpl.c"

struct _Gotk4GboxObject {
  GObject parent_instance;
  guintptr id;
};

G_DEFINE_TYPE(Gotk4GboxObject, gotk4_gbox_object, G_TYPE_OBJECT);

static void gotk4_gbox_object_init(Gotk4GboxObject *self) { self->id = 0; }

static void gotk4_gbox_object_finalize(GObject *object) {
  Gotk4GboxObject *self = GOTK4_GBOX_OBJECT(object);
  if (self->id != 0) {
    callbackDelete(self->id);
  }
  G_OBJECT_CLASS(gotk4_gbox_object_parent_class)->finalize(object);
}

static void gotk4_gbox_object_class_init(Gotk4GboxObjectClass *klass) {
  GObjectClass *object_class = G_OBJECT_CLASS(klass);
  object_class->finalize = gotk4_gbox_object_finalize;
}

Gotk4GboxObject *gotk4_gbox_object_new(guintptr id) {
  Gotk4GboxObject *self = g_object_new(GOTK4_TYPE_GBOX_OBJECT, NULL);
  self->id = id;
  return self;
}

guintptr gotk4_gbox_object_get_id(Gotk4GboxObject *self) { return self->id; }

struct _Gotk4GboxList {
  GObject parent_instance;
  Objects items;
};

struct _Gotk4GboxClass {
  GObjectClass parent_class;
};

static GType gotk4_gbox_list_get_item_type(GListModel *list) {
  return G_TYPE_OBJECT;
}

static guint gotk4_gbox_list_get_n_items(GListModel *list) {
  Gotk4GboxList *self = GOTK4_GBOX_LIST(list);
  return objects_get_size(&self->items);
}

static gpointer gotk4_gbox_list_get_item(GListModel *list, guint index) {
  Gotk4GboxList *self = GOTK4_GBOX_LIST(list);
  if (index >= objects_get_size(&self->items)) {
    return NULL;
  }
  return g_object_ref(objects_get(&self->items, index));
}

static void gotk4_gbox_list_list_model_init(GListModelInterface *iface) {
  iface->get_item_type = gotk4_gbox_list_get_item_type;
  iface->get_n_items = gotk4_gbox_list_get_n_items;
  iface->get_item = gotk4_gbox_list_get_item;
}

G_DEFINE_TYPE_WITH_CODE(Gotk4GboxList, gotk4_gbox_list, G_TYPE_OBJECT,
                        G_IMPLEMENT_INTERFACE(G_TYPE_LIST_MODEL,
                                              gotk4_gbox_list_list_model_init))

static void gotk4_gbox_list_dispose(GObject *object) {
  Gotk4GboxList *self = GOTK4_GBOX_LIST(object);
  objects_clear(&self->items);
  G_OBJECT_CLASS(gotk4_gbox_list_parent_class)->dispose(object);
}

static void gotk4_gbox_list_class_init(Gotk4GboxListClass *klass) {
  GObjectClass *object_class = G_OBJECT_CLASS(klass);
  object_class->dispose = gotk4_gbox_list_dispose;
}

static void gotk4_gbox_list_init(Gotk4GboxList *self) {
  objects_init(&self->items);
}

Gotk4GboxList *gotk4_gbox_list_new() {
  return GOTK4_GBOX_LIST(g_object_new(GOTK4_TYPE_GBOX_LIST, NULL));
}

void gotk4_gbox_list_splice(Gotk4GboxList *self, guint position,
                            guint n_removals, const guintptr *additions) {
  g_return_if_fail(GOTK4_IS_GBOX_LIST(self));
  g_return_if_fail(position + n_removals >= position); // overflow
  g_return_if_fail(position + n_removals <= objects_get_size(&self->items));

  guint n_additions = 0;
  if (additions) {
    for (n_additions = 0; additions[n_additions] != 0; n_additions++) {
    }
  }

  objects_splice(&self->items, position, n_removals, FALSE, NULL, n_additions);
  for (guint i = 0; i < n_additions; i++) {
    *objects_index(&self->items, position + i) =
        gotk4_gbox_object_new(additions[i]);
  }

  if (n_removals || n_additions) {
    g_list_model_items_changed(G_LIST_MODEL(self), position, n_removals,
                               n_additions);
  }
}

void gotk4_gbox_list_append(Gotk4GboxList *self, guintptr id) {
  g_return_if_fail(GOTK4_IS_GBOX_LIST(self));
  objects_append(&self->items, gotk4_gbox_object_new(id));
  g_list_model_items_changed(G_LIST_MODEL(self),
                             objects_get_size(&self->items) - 1, 0, 1);
}

guintptr gotk4_gbox_list_get_id(Gotk4GboxList *self, guint position) {
  g_return_val_if_fail(GOTK4_IS_GBOX_LIST(self), 0);
  g_return_val_if_fail(position < objects_get_size(&self->items), 0);
  return gotk4_gbox_object_get_id(*objects_index(&self->items, position));
}
