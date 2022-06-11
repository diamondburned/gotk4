#include <stdlib.h>
#include <glib.h>

struct Pixdata {
    guint32 magic;
    gint32  length;
    guint32 pixdata_type;
    guint32 rowstride;
    guint32 width;
    guint32 height;
    void*   pixel_data;
};