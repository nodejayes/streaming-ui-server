import { BaseEvent } from "./base.event";
import { MouseEventData } from "./types";

export class MouseMoveEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as MouseEvent;
        if (!event) {
            console.warn(`${MouseMoveEvent.typ} event is not of type PointerEvent`, e);
            return;
        }
        if (!event.target) {
            console.warn(`missing ${MouseMoveEvent.typ} event target`, e);
            return;
        }
        MouseMoveEvent.base?.handleEvent(MouseMoveEvent.typ, event.target as HTMLElement, () => ({
            typ: MouseMoveEvent.typ,
            ctrlKey: event.ctrlKey,
            altKey: event.altKey,
            shiftKey: event.shiftKey,
            clientX: event.clientX,
            clientY: event.clientY,
            pageX: event.pageX,
            pageY: event.pageY,
            offsetX: event.offsetX,
            offsetY: event.offsetY,
            screenX: event.screenX,
            screenY: event.screenY,
            buttons: event.buttons,
            button: event.button,
            movementX: event.movementX,
            movementY: event.movementY,
            x: event.x,
            y: event.y,
            detail: event.detail,
            type: event.type,
            timeStamp: event.timeStamp,
            bubbles: event.bubbles,
            cancelable: event.cancelable,
            composed: event.composed,
            eventPhase: event.eventPhase,
            isTrusted: event.isTrusted,
            defaultPrevented: event.defaultPrevented,
        } as MouseEventData), e);
    };
    private static get typ() {
        return "mousemove"
    }
    public static register(api: WsApi) {
        MouseMoveEvent.base.reattach(api, MouseMoveEvent.typ, MouseMoveEvent.eventHandler);
    }
}