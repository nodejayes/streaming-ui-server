import { BaseEvent } from "./base.event";
import { KeyboardEventData } from "./types";

export class KeyUpEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as KeyboardEvent;
        if (!event) {
            console.warn(`${KeyUpEvent.typ} event is not of type PointerEvent`, e);
            return;
        }
        if (!event.target) {
            console.warn(`missing ${KeyUpEvent.typ} event target`, e);
            return;
        }
        KeyUpEvent.base?.handleEvent(KeyUpEvent.typ, event.target as HTMLElement, () => ({
            typ: KeyUpEvent.typ,
            ctrlKey: event.ctrlKey,
            altKey: event.altKey,
            shiftKey: event.shiftKey,
            key: event.key,
            repeat: event.repeat,
            code: event.code,
            isComposing: event.isComposing,
            location: event.location,
            metaKey: event.metaKey,
            detail: event.detail,
            type: event.type,
            timeStamp: event.timeStamp,
            bubbles: event.bubbles,
            cancelable: event.cancelable,
            composed: event.composed,
            eventPhase: event.eventPhase,
            isTrusted: event.isTrusted,
            defaultPrevented: event.defaultPrevented,
        } as KeyboardEventData), e);
    };
    private static get typ() {
        return "keyup"
    }
    public static register(api: WsApi) {
        KeyUpEvent.base.reattach(api, KeyUpEvent.typ, KeyUpEvent.eventHandler);
    }
}