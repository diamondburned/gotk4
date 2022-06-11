#include <stdlib.h>
#include <glib.h>

struct Atom {

};
struct Color {
    guint32 pixel;
    guint16 red;
    guint16 green;
    guint16 blue;
};
struct DevicePadInterface {

};
struct DrawingContextClass {

};
struct EventAny {
          type;
    void* window;
    gint8 send_event;
};
struct EventButton {
            type;
    void*   window;
    gint8   send_event;
    guint32 time;
    gdouble x;
    gdouble y;
    void*   axes;
            state;
    guint   button;
    void*   device;
    gdouble x_root;
    gdouble y_root;
};
struct EventConfigure {
          type;
    void* window;
    gint8 send_event;
    gint  x;
    gint  y;
    gint  width;
    gint  height;
};
struct EventCrossing {
             type;
    void*    window;
    gint8    send_event;
    void*    subwindow;
    guint32  time;
    gdouble  x;
    gdouble  y;
    gdouble  x_root;
    gdouble  y_root;
             mode;
             detail;
    gboolean focus;
             state;
};
struct EventDND {
            type;
    void*   window;
    gint8   send_event;
    void*   context;
    guint32 time;
    gshort  x_root;
    gshort  y_root;
};
struct EventExpose {
          type;
    void* window;
    gint8 send_event;
          area;
    void* region;
    gint  count;
};
struct EventFocus {
           type;
    void*  window;
    gint8  send_event;
    gint16 in;
};
struct EventGrabBroken {
             type;
    void*    window;
    gint8    send_event;
    gboolean keyboard;
    gboolean implicit;
    void*    grab_window;
};
struct EventKey {
            type;
    void*   window;
    gint8   send_event;
    guint32 time;
            state;
    guint   keyval;
    gint    length;
    void*   string;
    guint16 hardware_keycode;
    guint8  group;
    guint   is_modifier  : 1;
};
struct EventMotion {
            type;
    void*   window;
    gint8   send_event;
    guint32 time;
    gdouble x;
    gdouble y;
    void*   axes;
            state;
    gint16  is_hint;
    void*   device;
    gdouble x_root;
    gdouble y_root;
};
struct EventOwnerChange {
            type;
    void*   window;
    gint8   send_event;
    void*   owner;
            reason;
            selection;
    guint32 time;
    guint32 selection_time;
};
struct EventPadAxis {
            type;
    void*   window;
    gint8   send_event;
    guint32 time;
    guint   group;
    guint   index;
    guint   mode;
    gdouble value;
};
struct EventPadButton {
            type;
    void*   window;
    gint8   send_event;
    guint32 time;
    guint   group;
    guint   button;
    guint   mode;
};
struct EventPadGroupMode {
            type;
    void*   window;
    gint8   send_event;
    guint32 time;
    guint   group;
    guint   mode;
};
struct EventProperty {
            type;
    void*   window;
    gint8   send_event;
            atom;
    guint32 time;
            state;
};
struct EventProximity {
            type;
    void*   window;
    gint8   send_event;
    guint32 time;
    void*   device;
};
struct EventScroll {
            type;
    void*   window;
    gint8   send_event;
    guint32 time;
    gdouble x;
    gdouble y;
            state;
            direction;
    void*   device;
    gdouble x_root;
    gdouble y_root;
    gdouble delta_x;
    gdouble delta_y;
    guint   is_stop  : 1;
};
struct EventSelection {
            type;
    void*   window;
    gint8   send_event;
            selection;
            target;
            property;
    guint32 time;
    void*   requestor;
};
struct EventSequence {

};
struct EventSetting {
          type;
    void* window;
    gint8 send_event;
          action;
    void* name;
};
struct EventTouch {
             type;
    void*    window;
    gint8    send_event;
    guint32  time;
    gdouble  x;
    gdouble  y;
    void*    axes;
             state;
    void*    sequence;
    gboolean emulating_pointer;
    void*    device;
    gdouble  x_root;
    gdouble  y_root;
};
struct EventTouchpadPinch {
            type;
    void*   window;
    gint8   send_event;
    gint8   phase;
    gint8   n_fingers;
    guint32 time;
    gdouble x;
    gdouble y;
    gdouble dx;
    gdouble dy;
    gdouble angle_delta;
    gdouble scale;
    gdouble x_root;
    gdouble y_root;
            state;
};
struct EventTouchpadSwipe {
            type;
    void*   window;
    gint8   send_event;
    gint8   phase;
    gint8   n_fingers;
    guint32 time;
    gdouble x;
    gdouble y;
    gdouble dx;
    gdouble dy;
    gdouble x_root;
    gdouble y_root;
            state;
};
struct EventVisibility {
          type;
    void* window;
    gint8 send_event;
          state;
};
struct EventWindowState {
          type;
    void* window;
    gint8 send_event;
          changed_mask;
          new_window_state;
};
struct FrameClockClass {

};
struct FrameClockPrivate {

};
struct FrameTimings {

};
struct Geometry {
    gint    min_width;
    gint    min_height;
    gint    max_width;
    gint    max_height;
    gint    base_width;
    gint    base_height;
    gint    width_inc;
    gint    height_inc;
    gdouble min_aspect;
    gdouble max_aspect;
            win_gravity;
};
struct KeymapKey {
    guint keycode;
    gint  group;
    gint  level;
};
struct MonitorClass {

};
struct Point {
    gint x;
    gint y;
};
struct RGBA {
    gdouble red;
    gdouble green;
    gdouble blue;
    gdouble alpha;
};
struct Rectangle {
    int x;
    int y;
    int width;
    int height;
};
struct TimeCoord {
    guint32 time;
    void    axes;
};
struct WindowAttr {
    void*    title;
    gint     event_mask;
    gint     x;
    gint     y;
    gint     width;
    gint     height;
             wclass;
    void*    visual;
             window_type;
    void*    cursor;
    void*    wmclass_name;
    void*    wmclass_class;
    gboolean override_redirect;
             type_hint;
};
struct WindowClass {
          parent_class;
    void* pick_embedded_child;
    void* to_embedder;
    void* from_embedder;
    void* create_surface;
    void* _gdk_reserved1;
    void* _gdk_reserved2;
    void* _gdk_reserved3;
    void* _gdk_reserved4;
    void* _gdk_reserved5;
    void* _gdk_reserved6;
    void* _gdk_reserved7;
    void* _gdk_reserved8;
};
struct WindowRedirect {

};