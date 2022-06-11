#include <stdlib.h>
#include <glib.h>

struct ContentFormats {

};
struct ContentFormatsBuilder {

};
struct ContentProviderClass {
          parent_class;
    void* content_changed;
    void* attach_clipboard;
    void* detach_clipboard;
    void* ref_formats;
    void* ref_storable_formats;
    void* write_mime_type_async;
    void* write_mime_type_finish;
    void* get_value;
    void  padding;
};
struct DevicePadInterface {

};
struct DragSurfaceInterface {

};
struct EventSequence {

};
struct FrameClockClass {

};
struct FrameClockPrivate {

};
struct FrameTimings {

};
struct GLTextureClass {

};
struct KeymapKey {
    guint keycode;
    int   group;
    int   level;
};
struct MemoryTextureClass {

};
struct MonitorClass {

};
struct PaintableInterface {
          g_iface;
    void* snapshot;
    void* get_current_image;
    void* get_flags;
    void* get_intrinsic_width;
    void* get_intrinsic_height;
    void* get_intrinsic_aspect_ratio;
};
struct PopupInterface {

};
struct PopupLayout {

};
struct RGBA {
    float red;
    float green;
    float blue;
    float alpha;
};
struct Rectangle {
    int x;
    int y;
    int width;
    int height;
};
struct SnapshotClass {

};
struct SurfaceClass {

};
struct TextureClass {

};
struct TimeCoord {
    guint32 time;
            flags;
    void    axes;
};
struct ToplevelInterface {

};
struct ToplevelLayout {

};
struct ToplevelSize {

};