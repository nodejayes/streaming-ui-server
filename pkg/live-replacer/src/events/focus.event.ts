import { BaseEvent } from "./base.event";
import { FocusEventData } from "./types";

export class FocusEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as Event;
        if (!event) {
            console.warn(`${FocusEvent.typ} event is not of type PointerEvent`, e);
            return;
        }
        if (!event.target) {
            console.warn(`missing ${FocusEvent.typ} event target`, e);
            return;
        }
        FocusEvent.base?.handleEvent(FocusEvent.typ, event.target as HTMLElement, () => ({
            typ: FocusEvent.typ,
        } as FocusEventData), e);
    };
    private static get typ() {
        return "focus"
    }
    public static register(api: WsApi) {
        FocusEvent.base.reattach(api, FocusEvent.typ, FocusEvent.eventHandler);
    }
}