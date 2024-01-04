import { ClickEventData } from "./types";
import { BaseEvent } from "./base.event";

export class ContextMenuEvent {
  private static base = new BaseEvent();
  private static eventHandler = (e: Event) => {
    const event = e as PointerEvent;
    event.preventDefault();
    if (!event) {
      console.warn(
        `${ContextMenuEvent.typ} event is not of type PointerEvent`,
        e
      );
      return;
    }
    if (!event.target) {
      console.warn(`missing ${ContextMenuEvent.typ} event target`, e);
      return;
    }
    ContextMenuEvent.base?.handleEvent(
      ContextMenuEvent.typ,
      event.target as HTMLElement,
      () =>
        ({
          typ: ContextMenuEvent.typ,
          ctrlKey: event.ctrlKey,
          altKey: event.altKey,
          shiftKey: event.shiftKey,
          isPrimary: event.isPrimary,
          clientX: event.clientX,
          clientY: event.clientY,
          height: event.height,
          width: event.width,
          pointerType: event.pointerType,
          pressure: event.pressure,
          tangentialPressure: event.tangentialPressure,
          tiltX: event.tiltX,
          tiltY: event.tiltY,
          twist: event.twist,
          button: event.button,
          buttons: event.buttons,
          metaKey: event.metaKey,
          movementX: event.movementX,
          movementY: event.movementY,
          offsetX: event.offsetX,
          offsetY: event.offsetY,
          pageX: event.pageX,
          pageY: event.pageY,
          screenX: event.screenX,
          screenY: event.screenY,
          x: event.x,
          y: event.y,
        } as ClickEventData),
      e
    );
  };

  private static get typ() {
    return "contextmenu";
  }

  public static register(api: WsApi) {
    ContextMenuEvent.base.reattach(
      api,
      ContextMenuEvent.typ,
      ContextMenuEvent.eventHandler
    );
  }
}
