import { BaseEvent } from "./base.event";
import { MouseEventData } from "./types";

export class MouseUpEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as MouseEvent;
        if (!event) {
            console.warn(`${MouseUpEvent.typ} event is not of type PointerEvent`, e);
            return;
        }
        if (!event.target) {
            console.warn(`missing ${MouseUpEvent.typ} event target`, e);
            return;
        }
        MouseUpEvent.base?.handleEvent(MouseUpEvent.typ, event.target as HTMLElement, () => ({
            typ: MouseUpEvent.typ,
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
        return "mouseup"
    }
    public static register(api: WsApi) {
        MouseUpEvent.base.reattach(api, MouseUpEvent.typ, MouseUpEvent.eventHandler);
    }
}