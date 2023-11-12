import { BaseEvent } from "./base.event";
import { FocusEventData } from "./types";

export class FocusOutEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as Event;
        if (!event) {
            console.warn(`${FocusOutEvent.typ} event is not of type PointerEvent`, e);
            return;
        }
        if (!event.target) {
            console.warn(`missing ${FocusOutEvent.typ} event target`, e);
            return;
        }
        FocusOutEvent.base?.handleEvent(FocusOutEvent.typ, event.target as HTMLElement, () => ({
            typ: FocusOutEvent.typ,
        } as FocusEventData), e);
    };
    private static get typ() {
        return "focusout"
    }
    public static register(api: WsApi) {
        FocusOutEvent.base.reattach(api, FocusOutEvent.typ, FocusOutEvent.eventHandler);
    }
}