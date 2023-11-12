import { BaseEvent } from "./base.event";
import { ScrollEventData } from "./types";

export class ScrollEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as Event;
        if (!event) {
            console.warn(`${ScrollEvent.typ} event is not of type PointerEvent`, e);
            return;
        }
        if (!event.target) {
            console.warn(`missing ${ScrollEvent.typ} event target`, e);
            return;
        }
        ScrollEvent.base?.handleEvent(ScrollEvent.typ, event.target as HTMLElement, () => ({
            typ: ScrollEvent.typ,
        } as ScrollEventData), e);
    };
    private static get typ() {
        return "scroll"
    }
    public static register(api: WsApi) {
        ScrollEvent.base.reattach(api, ScrollEvent.typ, ScrollEvent.eventHandler);
    }
}