import { BaseEvent } from "./base.event";
import { ScrollEventData } from "./types";

export class ScrollEndEvent {
    private static base = new BaseEvent();
    private static eventHandler = (e: Event) => {
        const event = e as Event;
        if (!event) {
            console.warn(`${ScrollEndEvent.typ} event is not of type PointerEvent`, e);
            return;
        }
        if (!event.target) {
            console.warn(`missing ${ScrollEndEvent.typ} event target`, e);
            return;
        }
        ScrollEndEvent.base?.handleEvent(ScrollEndEvent.typ, event.target as HTMLElement, () => ({
            typ: ScrollEndEvent.typ,
        } as ScrollEventData), e);
    };
    private static get typ() {
        return "scrollend"
    }
    public static register(api: WsApi) {
        ScrollEndEvent.base.reattach(api, ScrollEndEvent.typ, ScrollEndEvent.eventHandler);
    }
}