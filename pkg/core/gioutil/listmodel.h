#pragma once

#include <gio/gio.h>
#include <glib-object.h>
#include <glib.h>

G_BEGIN_DECLS

#define GOTK4_TYPE_GBOX_OBJECT (gotk4_gbox_object_get_type())

G_DECLARE_FINAL_TYPE(Gotk4GboxObject, gotk4_gbox_object, GOTK4, GBOX_OBJECT,
                     GObject)

Gotk4GboxObject *gotk4_gbox_object_new(guintptr id);
guintptr gotk4_gbox_object_get_id(Gotk4GboxObject *self);

#define GOTK4_TYPE_GBOX_LIST (gotk4_gbox_list_get_type())

G_DECLARE_FINAL_TYPE(Gotk4GboxList, gotk4_gbox_list, GOTK4, GBOX_LIST, GObject)

Gotk4GboxList *gotk4_gbox_list_new(void);
void gotk4_gbox_list_splice(Gotk4GboxList *self, guint position,
                            guint n_removals, const guintptr *additions);
void gotk4_gbox_list_append(Gotk4GboxList *self, guintptr id);
void gotk4_gbox_list_remove(Gotk4GboxList *self, guint position);
guintptr gotk4_gbox_list_get_id(Gotk4GboxList *self, guint position);

G_END_DECLS
