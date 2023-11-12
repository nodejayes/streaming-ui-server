import { BaseEvent } from "./base.event";
import { FocusEventData } from "./types";

export class FocusInEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as Event;
        if (!event) {
            console.warn(`${FocusInEvent.typ} event is not of type PointerEvent`, e);
            return;
        }
        if (!event.target) {
            console.warn(`missing ${FocusInEvent.typ} event target`, e);
            return;
        }
        FocusInEvent.base?.handleEvent(FocusInEvent.typ, event.target as HTMLElement, () => ({
            typ: FocusInEvent.typ,
        } as FocusEventData), e);
    };
    private static get typ() {
        return "focusin"
    }
    public static register(api: WsApi) {
        FocusInEvent.base.reattach(api, FocusInEvent.typ, FocusInEvent.eventHandler);
    }
}