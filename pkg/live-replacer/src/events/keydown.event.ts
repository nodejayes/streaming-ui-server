import { BaseEvent } from "./base.event";
import { KeyboardEventData } from "./types";

export class KeyDownEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as KeyboardEvent;
        if (!event) {
            console.warn(`${KeyDownEvent.typ} event is not of type PointerEvent`, e);
            return;
        }
        if (!event.target) {
            console.warn(`missing ${KeyDownEvent.typ} event target`, e);
            return;
        }
        KeyDownEvent.base?.handleEvent(KeyDownEvent.typ, event.target as HTMLElement, () => ({
            typ: KeyDownEvent.typ,
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
        return "keydown"
    }
    public static register(api: WsApi) {
        KeyDownEvent.base.reattach(api, KeyDownEvent.typ, KeyDownEvent.eventHandler);
    }
}